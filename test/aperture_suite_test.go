package test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"

	"github.com/fluxninja/aperture/cmd/aperture-agent/agent"
	"github.com/fluxninja/aperture/pkg/agentinfo"
	"github.com/fluxninja/aperture/pkg/alerts"
	"github.com/fluxninja/aperture/pkg/cache"
	"github.com/fluxninja/aperture/pkg/entitycache"
	etcdclient "github.com/fluxninja/aperture/pkg/etcd/client"
	etcdwatcher "github.com/fluxninja/aperture/pkg/etcd/watcher"
	"github.com/fluxninja/aperture/pkg/jobs"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/net/grpc"
	"github.com/fluxninja/aperture/pkg/notifiers"
	"github.com/fluxninja/aperture/pkg/otelcollector"
	otelconfig "github.com/fluxninja/aperture/pkg/otelcollector/config"
	"github.com/fluxninja/aperture/pkg/platform"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/resources/classifier"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/selectors"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/service"
	"github.com/fluxninja/aperture/pkg/policies/flowcontrol/servicegetter"
	"github.com/fluxninja/aperture/pkg/status"
	"github.com/fluxninja/aperture/pkg/utils"
	"github.com/fluxninja/aperture/test/harness"
)

const (
	jobGroupName = "job-group"
)

var (
	project     string
	app         *fx.App
	addr        string
	configDir   string
	l           *utils.GoLeakDetector
	eh          *harness.EtcdHarness
	ehStarted   bool
	etcdClient  *etcdclient.Client
	etcdWatcher notifiers.Watcher
	ph          *harness.PrometheusHarness
	phStarted   bool
	jgIn        *JobGroupIn
)

type JobGroupIn struct {
	fx.In
	JobGroup *jobs.JobGroup `name:"job-group"`
	Registry status.Registry
}

func TestCore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aperture Suite")
}

var _ = BeforeSuite(func() {
	var err error

	addr = ":18081"

	configDir = "aperturetest-config"

	l = utils.NewGoLeakDetector()

	_, err = harness.LocalBinAvailable(harness.EtcdBinPath)
	if err == nil {
		eh, err = harness.NewEtcdHarness(os.Stderr)
		Expect(err).NotTo(HaveOccurred())
		ehStarted = true
	}

	_, err = harness.LocalBinAvailable(harness.PrometheusBinPath)
	if err == nil {
		ph, err = harness.NewPrometheusHarness(os.Stderr)
		Expect(err).NotTo(HaveOccurred())
		phStarted = true
	}

	// reset config dir
	err = os.RemoveAll(configDir)
	Expect(err).NotTo(HaveOccurred())
	err = os.MkdirAll(configDir, 0o777)
	Expect(err).NotTo(HaveOccurred())

	jgIn = &JobGroupIn{}

	apertureConfig := map[string]interface{}{
		"plugins": map[string]interface{}{
			"disabled_plugins": []string{"aperture-plugin-fluxninja"},
		},
		"log": map[string]interface{}{
			// for cleaner logs and for testing config
			"level":          "debug",
			"pretty_console": true,
		},
		"server": map[string]interface{}{
			"addr": addr,
			"grpc": map[string]interface{}{
				"enable_reflection": true,
			},
		},
		"config_path": configDir,
		"sentrywriter": map[string]interface{}{
			"disabled": true,
		},
	}

	if ehStarted {
		apertureConfig["etcd"] = map[string]interface{}{
			"endpoints": []string{eh.Endpoint},
			"tls": map[string]interface{}{
				"insecure_skip_verify": true,
			},
			"username": harness.EtcdUsername,
			"password": harness.EtcdPassword,
		}
	}

	if phStarted {
		apertureConfig["prometheus"] = map[string]interface{}{
			"address": ph.Endpoint,
		}
	}

	apertureOpts := fx.Options(
		platform.Config{MergeConfig: apertureConfig}.Module(),
		fx.Option(
			fx.Provide(
				fx.Annotate(
					provideOTELConfig,
					fx.ResultTags(`name:"base"`),
				),
			),
		),
		classifier.Module(),
		service.Module(),
		fx.Provide(
			clockwork.NewRealClock,
			fx.Annotate(
				agent.AgentOTELComponents,
				fx.ParamTags(alerts.AlertsFxTag),
			),
			entitycache.NewEntityCache,
			servicegetter.NewEmpty,
			agentinfo.ProvideAgentInfo,
			flowcontrol.NewEngine,
			cache.NewCache[selectors.ControlPointID],
		),
		otelcollector.Module(),
		grpc.ClientConstructor{Name: "flowcontrol-grpc-client", ConfigKey: "flowcontrol.client.grpc"}.Annotate(),
		jobs.JobGroupConstructor{Name: jobGroupName}.Annotate(),
		fx.Populate(jgIn),
	)

	if ehStarted {
		apertureOpts = fx.Options(
			apertureOpts,
			etcdwatcher.Constructor{Name: "test-etcd-watcher", EtcdPath: "foo/"}.Annotate(),
			fx.Populate(&etcdClient),
		)
	}

	app = platform.New(apertureOpts)

	err = app.Err()

	if err != nil {
		visualize, _ := fx.VisualizeError(err)
		log.Error().Err(err).Msg("fx.New failed: " + visualize)
	}

	Expect(err).NotTo(HaveOccurred())

	startCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err = app.Start(startCtx)
	Expect(err).NotTo(HaveOccurred())

	etcdWatcher, err = etcdwatcher.NewWatcher(etcdClient, "foo/")
	Expect(err).NotTo(HaveOccurred())
	etcdWatcher.Start()

	project = "staging"
	Eventually(func() bool {
		_, err := http.Get(fmt.Sprintf("http://%v/version", addr))
		return err == nil
	}).Should(BeTrue())
})

var _ = AfterSuite(func() {
	stopCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var err error

	if etcdWatcher != nil {
		etcdWatcher.Stop()
	}

	err = app.Stop(stopCtx)
	Expect(err).NotTo(HaveOccurred())

	_ = os.RemoveAll(configDir)

	if ehStarted {
		eh.Stop()
	}

	if phStarted {
		ph.Stop()
	}

	err = l.FindLeaks()
	Expect(err).NotTo(HaveOccurred())
})

func provideOTELConfig() *otelconfig.OTELConfig {
	cfg := otelconfig.NewOTELConfig()
	cfg.AddReceiver("prometheus", map[string]interface{}{
		"config": map[string]interface{}{
			"scrape_configs": []map[string]interface{}{
				{
					"job_name":        "aperture-agent",
					"scrape_interval": "5s",
					"static_configs": []map[string]interface{}{
						{
							"targets": []string{addr},
						},
					},
				},
			},
		},
	})
	cfg.AddExporter("prometheusremotewrite", map[string]interface{}{
		"endpoint": ph.Endpoint + "/api/v1/write",
	})
	cfg.Service.AddPipeline("metrics", otelconfig.Pipeline{
		Receivers: []string{"prometheus"},
		Exporters: []string{"prometheusremotewrite"},
	})
	return cfg
}
