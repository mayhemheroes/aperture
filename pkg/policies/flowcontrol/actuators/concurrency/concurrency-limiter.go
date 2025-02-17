package concurrency

import (
	"context"
	"fmt"
	"math"
	"path"
	"strconv"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
	"go.uber.org/multierr"

	flowcontrolv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/flowcontrol/check/v1"
	policylangv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	policysyncv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/sync/v1"
	"github.com/fluxninja/aperture/pkg/agentinfo"
	"github.com/fluxninja/aperture/pkg/config"
	etcdclient "github.com/fluxninja/aperture/pkg/etcd/client"
	etcdwatcher "github.com/fluxninja/aperture/pkg/etcd/watcher"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/metrics"
	"github.com/fluxninja/aperture/pkg/multimatcher"
	"github.com/fluxninja/aperture/pkg/notifiers"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/actuators/concurrency/scheduler"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/iface"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/selectors"
	"github.com/fluxninja/aperture/pkg/policies/paths"
	"github.com/fluxninja/aperture/pkg/status"
)

var (
	// FxNameTag is Concurrency Limiter Watcher's Fx Tag.
	fxNameTag = config.NameTag("concurrency_limiter_watcher")

	// Array of Label Keys for WFQ and Token Bucket Metrics.
	metricLabelKeys = []string{metrics.PolicyNameLabel, metrics.PolicyHashLabel, metrics.ComponentIDLabel}
)

// concurrencyLimiterModule returns the fx options for flowcontrol side pieces of concurrency limiter in the main fx app.
func concurrencyLimiterModule() fx.Option {
	return fx.Options(
		// Tag the watcher so that other modules can find it.
		fx.Provide(
			fx.Annotate(
				provideWatcher,
				fx.ResultTags(fxNameTag),
			),
		),
		fx.Invoke(
			fx.Annotate(
				setupConcurrencyLimiterFactory,
				fx.ParamTags(fxNameTag),
			),
		),
	)
}

// provideWatcher provides pointer to concurrency limiter watcher.
func provideWatcher(
	etcdClient *etcdclient.Client,
	ai *agentinfo.AgentInfo,
) (notifiers.Watcher, error) {
	// Get Agent Group from host info gatherer
	agentGroupName := ai.GetAgentGroup()
	// Scope the sync to the agent group.
	etcdPath := path.Join(paths.ConcurrencyLimiterConfigPath,
		paths.AgentGroupPrefix(agentGroupName))
	watcher, err := etcdwatcher.NewWatcher(etcdClient, etcdPath)
	if err != nil {
		return nil, err
	}
	return watcher, nil
}

type concurrencyLimiterFactory struct {
	engineAPI iface.Engine
	registry  status.Registry

	autoTokensFactory   *autoTokensFactory
	loadActuatorFactory *loadActuatorFactory

	// WFQ Metrics.
	wfqFlowsGaugeVec    *prometheus.GaugeVec
	wfqRequestsGaugeVec *prometheus.GaugeVec

	// TODO: following will be moved to scheduler.
	incomingWorkSecondsCounterVec *prometheus.CounterVec
	acceptedWorkSecondsCounterVec *prometheus.CounterVec

	workloadLatencySummaryVec *prometheus.SummaryVec
	workloadCounterVec        *prometheus.CounterVec
}

// setupConcurrencyLimiterFactory sets up the concurrency limiter module in the main fx app.
func setupConcurrencyLimiterFactory(
	watcher notifiers.Watcher,
	lifecycle fx.Lifecycle,
	e iface.Engine,
	statusRegistry status.Registry,
	prometheusRegistry *prometheus.Registry,
	etcdClient *etcdclient.Client,
	ai *agentinfo.AgentInfo,
) error {
	agentGroup := ai.GetAgentGroup()

	// Create factories
	loadActuatorFactory, err := newLoadActuatorFactory(lifecycle, etcdClient, agentGroup, prometheusRegistry)
	if err != nil {
		return err
	}

	autoTokensFactory, err := newAutoTokensFactory(lifecycle, etcdClient, agentGroup)
	if err != nil {
		return err
	}

	reg := statusRegistry.Child("component", "concurrency_limiter")

	conLimiterFactory := &concurrencyLimiterFactory{
		engineAPI:           e,
		autoTokensFactory:   autoTokensFactory,
		loadActuatorFactory: loadActuatorFactory,
		registry:            reg,
	}

	conLimiterFactory.wfqFlowsGaugeVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metrics.WFQFlowsMetricName,
			Help: "A gauge that tracks the number of flows in the WFQScheduler",
		},
		metricLabelKeys,
	)
	conLimiterFactory.wfqRequestsGaugeVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metrics.WFQRequestsMetricName,
			Help: "A gauge that tracks the number of queued requests in the WFQScheduler",
		},
		metricLabelKeys,
	)
	conLimiterFactory.incomingWorkSecondsCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metrics.IncomingWorkSecondsMetricName,
			Help: "A counter measuring work incoming into Scheduler",
		},
		metricLabelKeys,
	)
	conLimiterFactory.acceptedWorkSecondsCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metrics.AcceptedWorkSecondsMetricName,
			Help: "A counter measuring work admitted by Scheduler",
		},
		metricLabelKeys,
	)

	conLimiterFactory.workloadLatencySummaryVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: metrics.WorkloadLatencyMetricName,
		Help: "Latency summary of workload",
	}, []string{
		metrics.PolicyNameLabel,
		metrics.PolicyHashLabel,
		metrics.ComponentIDLabel,
		metrics.WorkloadIndexLabel,
	})

	conLimiterFactory.workloadCounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: metrics.WorkloadCounterMetricName,
		Help: "Counter of workload requests",
	}, []string{
		metrics.PolicyNameLabel,
		metrics.PolicyHashLabel,
		metrics.ComponentIDLabel,
		metrics.DecisionTypeLabel,
		metrics.WorkloadIndexLabel,
		metrics.LimiterDroppedLabel,
	})

	fxDriver := &notifiers.FxDriver{
		FxOptionsFuncs: []notifiers.FxOptionsFunc{conLimiterFactory.newConcurrencyLimiterOptions},
		UnmarshalPrefixNotifier: notifiers.UnmarshalPrefixNotifier{
			GetUnmarshallerFunc: config.NewProtobufUnmarshaller,
		},
		StatusRegistry:     reg,
		PrometheusRegistry: prometheusRegistry,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			var merr error

			err := prometheusRegistry.Register(conLimiterFactory.wfqFlowsGaugeVec)
			if err != nil {
				merr = multierr.Append(merr, err)
			}
			err = prometheusRegistry.Register(conLimiterFactory.wfqRequestsGaugeVec)
			if err != nil {
				merr = multierr.Append(merr, err)
			}
			err = prometheusRegistry.Register(conLimiterFactory.incomingWorkSecondsCounterVec)
			if err != nil {
				merr = multierr.Append(merr, err)
			}
			err = prometheusRegistry.Register(conLimiterFactory.acceptedWorkSecondsCounterVec)
			if err != nil {
				merr = multierr.Append(merr, err)
			}
			err = prometheusRegistry.Register(conLimiterFactory.workloadLatencySummaryVec)
			if err != nil {
				merr = multierr.Append(merr, err)
			}
			err = prometheusRegistry.Register(conLimiterFactory.workloadCounterVec)
			if err != nil {
				merr = multierr.Append(merr, err)
			}

			return merr
		},
		OnStop: func(_ context.Context) error {
			var merr error

			if !prometheusRegistry.Unregister(conLimiterFactory.wfqFlowsGaugeVec) {
				err := fmt.Errorf("failed to unregister wfq_flows metric")
				merr = multierr.Append(merr, err)
			}
			if !prometheusRegistry.Unregister(conLimiterFactory.wfqRequestsGaugeVec) {
				err := fmt.Errorf("failed to unregister wfq_requests metric")
				merr = multierr.Append(merr, err)
			}
			if !prometheusRegistry.Unregister(conLimiterFactory.incomingWorkSecondsCounterVec) {
				err := fmt.Errorf("failed to unregister incoming_concurrency metric")
				merr = multierr.Append(merr, err)
			}
			if !prometheusRegistry.Unregister(conLimiterFactory.acceptedWorkSecondsCounterVec) {
				err := fmt.Errorf("failed to unregister accepted_concurrency metric")
				merr = multierr.Append(merr, err)
			}
			if !prometheusRegistry.Unregister(conLimiterFactory.workloadLatencySummaryVec) {
				err := fmt.Errorf("failed to unregister workload_latency_ms metric")
				merr = multierr.Append(merr, err)
			}
			if !prometheusRegistry.Unregister(conLimiterFactory.workloadCounterVec) {
				err := fmt.Errorf("failed to unregister workload_counter metric")
				merr = multierr.Append(merr, err)
			}

			return merr
		},
	})

	notifiers.WatcherLifecycle(lifecycle, watcher, []notifiers.PrefixNotifier{fxDriver})

	return nil
}

// multiMatchResult is used as return value of PolicyConfigAPI.GetMatches.
type multiMatchResult struct {
	matchedWorkloads map[int]*policylangv1.Scheduler_Workload_Parameters
}

// multiMatcher is MultiMatcher instantiation used in this package.
type multiMatcher = multimatcher.MultiMatcher[int, multiMatchResult]

// newConcurrencyLimiterOptions returns fx options for the concurrency limiter fx app.
func (conLimiterFactory *concurrencyLimiterFactory) newConcurrencyLimiterOptions(
	key notifiers.Key,
	unmarshaller config.Unmarshaller,
	reg status.Registry,
) (fx.Option, error) {
	logger := conLimiterFactory.registry.GetLogger()
	wrapperMessage := &policysyncv1.ConcurrencyLimiterWrapper{}
	err := unmarshaller.Unmarshal(wrapperMessage)
	concurrencyLimiterMessage := wrapperMessage.ConcurrencyLimiter
	if err != nil || concurrencyLimiterMessage == nil {
		reg.SetStatus(status.NewStatus(nil, err))
		logger.Warn().Err(err).Msg("Failed to unmarshal concurrency limiter config wrapper")
		return fx.Options(), err
	}

	// Scheduler config
	schedulerMsg := concurrencyLimiterMessage.Scheduler
	if schedulerMsg == nil {
		err = fmt.Errorf("no scheduler specified")
		reg.SetStatus(status.NewStatus(nil, err))
		return fx.Options(), err
	}
	schedulerParams := schedulerMsg.Parameters
	if schedulerParams == nil {
		err = fmt.Errorf("no scheduler parameters specified")
		reg.SetStatus(status.NewStatus(nil, err))
		return fx.Options(), err
	}
	mm := multimatcher.New[int, multiMatchResult]()
	// Loop through the workloads
	for workloadIndex, workloadProto := range schedulerParams.Workloads {
		labelMatcher, err := selectors.MMExprFromLabelMatcher(workloadProto.GetLabelMatcher())
		if err != nil {
			return fx.Options(), err
		}
		wm := &workloadMatcher{
			workloadIndex: workloadIndex,
			workloadProto: workloadProto,
		}
		err = mm.AddEntry(workloadIndex, labelMatcher, wm.matchCallback)
		if err != nil {
			return fx.Options(), err
		}
	}

	conLimiter := &concurrencyLimiter{
		Component:                    wrapperMessage.GetCommonAttributes(),
		concurrencyLimiterMsg:        concurrencyLimiterMessage,
		registry:                     reg,
		concurrencyLimiterFactory:    conLimiterFactory,
		workloadMultiMatcher:         mm,
		defaultWorkloadParametersMsg: schedulerParams.DefaultWorkloadParameters,
		schedulerParameters:          schedulerParams,
	}

	return fx.Options(
		fx.Invoke(
			conLimiter.setup,
		),
	), nil
}

type workloadMatcher struct {
	workloadProto *policylangv1.Scheduler_Workload
	workloadIndex int
}

func (wm *workloadMatcher) matchCallback(mmr multiMatchResult) multiMatchResult {
	// mmr.matchedWorkloads is nil on first match.
	if mmr.matchedWorkloads == nil {
		mmr.matchedWorkloads = make(map[int]*policylangv1.Scheduler_Workload_Parameters)
	}
	mmr.matchedWorkloads[wm.workloadIndex] = wm.workloadProto.GetParameters()
	return mmr
}

// concurrencyLimiter implements concurrency limiter on the flowcontrol side.
type concurrencyLimiter struct {
	iface.Component
	scheduler                    scheduler.Scheduler
	registry                     status.Registry
	incomingWorkSecondsCounter   prometheus.Counter
	acceptedWorkSecondsCounter   prometheus.Counter
	concurrencyLimiterMsg        *policylangv1.ConcurrencyLimiter
	concurrencyLimiterFactory    *concurrencyLimiterFactory
	autoTokens                   *autoTokens
	workloadMultiMatcher         *multiMatcher
	defaultWorkloadParametersMsg *policylangv1.Scheduler_Workload_Parameters
	schedulerParameters          *policylangv1.Scheduler_Parameters
}

// Make sure ConcurrencyLimiter implements the iface.ConcurrencyLimiter.
var _ iface.Limiter = &concurrencyLimiter{}

func (conLimiter *concurrencyLimiter) setup(lifecycle fx.Lifecycle) error {
	// Factories
	conLimiterFactory := conLimiter.concurrencyLimiterFactory
	loadActuatorFactory := conLimiterFactory.loadActuatorFactory
	autoTokensFactory := conLimiterFactory.autoTokensFactory
	// Form metric labels
	metricLabels := make(prometheus.Labels)
	metricLabels[metrics.PolicyNameLabel] = conLimiter.GetPolicyName()
	metricLabels[metrics.PolicyHashLabel] = conLimiter.GetPolicyHash()
	metricLabels[metrics.ComponentIDLabel] = conLimiter.GetComponentId()
	// Create sub components.
	clock := clockwork.NewRealClock()
	loadActuator, err := loadActuatorFactory.newLoadActuator(conLimiter.concurrencyLimiterMsg.GetLoadActuator(), conLimiter, conLimiter.registry, clock, lifecycle, metricLabels)
	if err != nil {
		return err
	}
	if conLimiter.schedulerParameters.AutoTokens {
		autoTokens, err := autoTokensFactory.newAutoTokens(
			conLimiter.GetPolicyName(), conLimiter.GetPolicyHash(),
			lifecycle, conLimiter.GetComponentId(), conLimiter.registry)
		if err != nil {
			return err
		}
		conLimiter.autoTokens = autoTokens
	}

	engineAPI := conLimiterFactory.engineAPI
	wfqFlowsGaugeVec := conLimiterFactory.wfqFlowsGaugeVec
	wfqRequestsGaugeVec := conLimiterFactory.wfqRequestsGaugeVec
	incomingWorkSecondsCounterVec := conLimiterFactory.incomingWorkSecondsCounterVec
	acceptedWorkSecondsCounterVec := conLimiterFactory.acceptedWorkSecondsCounterVec

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			retErr := func(err error) error {
				conLimiter.registry.SetStatus(status.NewStatus(nil, err))
				return err
			}

			wfqFlowsGauge, err := wfqFlowsGaugeVec.GetMetricWith(metricLabels)
			if err != nil {
				return errors.Wrap(err, "failed to get wfq flows gauge")
			}

			wfqRequestsGauge, err := wfqRequestsGaugeVec.GetMetricWith(metricLabels)
			if err != nil {
				return errors.Wrap(err, "failed to get wfq requests gauge")
			}

			wfqMetrics := &scheduler.WFQMetrics{
				FlowsGauge:        wfqFlowsGauge,
				HeapRequestsGauge: wfqRequestsGauge,
			}

			// setup scheduler
			conLimiter.scheduler = scheduler.NewWFQScheduler(loadActuator.tokenBucketLoadMultiplier, clock, wfqMetrics)

			conLimiter.incomingWorkSecondsCounter, err = incomingWorkSecondsCounterVec.GetMetricWith(metricLabels)
			if err != nil {
				return err
			}
			conLimiter.acceptedWorkSecondsCounter, err = acceptedWorkSecondsCounterVec.GetMetricWith(metricLabels)
			if err != nil {
				return err
			}

			err = engineAPI.RegisterConcurrencyLimiter(conLimiter)
			if err != nil {
				return retErr(err)
			}

			return retErr(nil)
		},
		OnStop: func(context.Context) error {
			var errMulti error

			err := engineAPI.UnregisterConcurrencyLimiter(conLimiter)
			if err != nil {
				errMulti = multierr.Append(errMulti, err)
			}

			// Remove metrics from metric vectors
			deleted := wfqFlowsGaugeVec.Delete(metricLabels)
			if !deleted {
				errMulti = multierr.Append(errMulti, errors.New("failed to delete wfq_flows gauge from its metric vector"))
			}
			deleted = wfqRequestsGaugeVec.Delete(metricLabels)
			if !deleted {
				errMulti = multierr.Append(errMulti, errors.New("failed to delete wfq_requests gauge from its metric vector"))
			}
			deleted = incomingWorkSecondsCounterVec.Delete(metricLabels)
			if !deleted {
				errMulti = multierr.Append(errMulti, errors.New("failed to delete incoming_concurrency counter from its metric vector"))
			}
			deleted = acceptedWorkSecondsCounterVec.Delete(metricLabels)
			if !deleted {
				errMulti = multierr.Append(errMulti, errors.New("failed to delete accepted_concurrency counter from its metric vector"))
			}
			deletedCount := conLimiter.concurrencyLimiterFactory.workloadLatencySummaryVec.DeletePartialMatch(metricLabels)
			if deletedCount == 0 {
				log.Warn().Msg("Could not delete workload_latency_ms summary from its metric vector. No traffic to generate metrics?")
			}
			deletedCount = conLimiter.concurrencyLimiterFactory.workloadCounterVec.DeletePartialMatch(metricLabels)
			if deletedCount == 0 {
				log.Warn().Msg("Could not delete workload_requests_total counter from its metric vector. No traffic to generate metrics?")
			}

			conLimiter.registry.SetStatus(status.NewStatus(nil, errMulti))
			return errMulti
		},
	})

	return nil
}

// GetFlowSelector returns selector.
func (conLimiter *concurrencyLimiter) GetFlowSelector() *policylangv1.FlowSelector {
	return conLimiter.concurrencyLimiterMsg.GetFlowSelector()
}

// RunLimiter processes a single flow by concurrency limiter in a blocking manner.
//
// Context is used to ensure that requests are not scheduled for longer than its deadline allows.
func (conLimiter *concurrencyLimiter) RunLimiter(ctx context.Context, labels map[string]string) *flowcontrolv1.LimiterDecision {
	var matchedWorkloadProto *policylangv1.Scheduler_Workload_Parameters
	var matchedWorkloadIndex string
	// match labels against conLimiter.workloadMultiMatcher
	mmr := conLimiter.workloadMultiMatcher.Match(multimatcher.Labels(labels))
	// if at least one match, return workload with lowest index
	if len(mmr.matchedWorkloads) > 0 {
		// select the smallest workloadIndex
		smallestWorkloadIndex := math.MaxInt32
		for workloadIndex := range mmr.matchedWorkloads {
			if workloadIndex < smallestWorkloadIndex {
				smallestWorkloadIndex = workloadIndex
			}
		}
		matchedWorkloadProto = mmr.matchedWorkloads[smallestWorkloadIndex]
		matchedWorkloadIndex = strconv.Itoa(smallestWorkloadIndex)
	} else {
		// no match, return default workload
		matchedWorkloadProto = conLimiter.defaultWorkloadParametersMsg
		matchedWorkloadIndex = metrics.DefaultWorkloadIndex
	}

	fairnessLabel := "workload:" + matchedWorkloadIndex

	if val, ok := labels[matchedWorkloadProto.FairnessKey]; ok {
		fairnessLabel = fairnessLabel + "," + val
	}
	// Lookup tokens for the workload
	var tokens uint64
	if conLimiter.schedulerParameters.AutoTokens {
		tokensAuto, ok := conLimiter.autoTokens.GetTokensForWorkload(matchedWorkloadIndex)
		if !ok {
			// default to 1 if auto tokens not found
			tokens = 1
		} else {
			tokens = tokensAuto
		}
	} else {
		tokens = matchedWorkloadProto.Tokens
	}

	// timeout is tokens(which is in milliseconds) * conLimiter.schedulerProto.TimeoutFactor(float64)
	timeout := time.Duration(float64(tokens)*conLimiter.schedulerParameters.TimeoutFactor) * time.Millisecond

	if timeout > conLimiter.schedulerParameters.MaxTimeout.AsDuration() {
		timeout = conLimiter.schedulerParameters.MaxTimeout.AsDuration()
	}

	if clientDeadline, hasDeadline := ctx.Deadline(); hasDeadline {
		// The clientDeadline is calculated based on client's timeout, passed
		// as grpc-timeout. Our goal is for the response to be received by the
		// client before its deadline passes (otherwise we risk fail-open on
		// timeout). To allow some headroom for transmitting the response to
		// the client, we set an "internal" deadline to a bit before client's
		// deadline, subtracting:
		// * 2 * 1ms to account for deadline inaccuracies (observed
		//   that Deadline() - Now() delta might end up longer than
		//   grpc-timeout (!), usually within 1ms),
		// * 1ms for response overhead,
		// * 7ms so that we don't always operate on the edge of the time budget.
		clientTimeout := time.Until(clientDeadline)
		internalTimeout := clientTimeout - 10*time.Millisecond
		if internalTimeout < timeout {
			timeout = internalTimeout
		}
		if timeout < 0 {
			timeout = 0
		}
	}

	reqContext := scheduler.RequestContext{
		FairnessLabel: fairnessLabel,
		Priority:      uint8(matchedWorkloadProto.Priority),
		Timeout:       timeout,
		WorkMillis:    tokens,
	}

	accepted := conLimiter.scheduler.Schedule(reqContext)

	// update concurrency metrics and decisionType
	conLimiter.incomingWorkSecondsCounter.Add(float64(reqContext.WorkMillis) / 1000)

	if accepted {
		conLimiter.acceptedWorkSecondsCounter.Add(float64(reqContext.WorkMillis) / 1000)
	}

	return &flowcontrolv1.LimiterDecision{
		PolicyName:  conLimiter.GetPolicyName(),
		PolicyHash:  conLimiter.GetPolicyHash(),
		ComponentId: conLimiter.GetComponentId(),
		Dropped:     !accepted,
		Details: &flowcontrolv1.LimiterDecision_ConcurrencyLimiterInfo_{
			ConcurrencyLimiterInfo: &flowcontrolv1.LimiterDecision_ConcurrencyLimiterInfo{
				WorkloadIndex: matchedWorkloadIndex,
			},
		},
	}
}

// GetLimiterID returns the limiter ID.
func (conLimiter *concurrencyLimiter) GetLimiterID() iface.LimiterID {
	// TODO: move this to limiter base.
	return iface.LimiterID{
		PolicyName:  conLimiter.GetPolicyName(),
		PolicyHash:  conLimiter.GetPolicyHash(),
		ComponentID: conLimiter.GetComponentId(),
	}
}

// GetLatencyObserver returns histogram for specific workload.
func (conLimiter *concurrencyLimiter) GetLatencyObserver(labels map[string]string) prometheus.Observer {
	latencySummary, err := conLimiter.concurrencyLimiterFactory.workloadLatencySummaryVec.GetMetricWith(labels)
	if err != nil {
		log.Warn().Err(err).Msg("Getting latency histogram")
		return nil
	}

	return latencySummary
}

// GetRequestCounter returns request counter for specific workload.
func (conLimiter *concurrencyLimiter) GetRequestCounter(labels map[string]string) prometheus.Counter {
	counter, err := conLimiter.concurrencyLimiterFactory.workloadCounterVec.GetMetricWith(labels)
	if err != nil {
		log.Warn().Err(err).Msg("Getting counter")
		return nil
	}

	return counter
}
