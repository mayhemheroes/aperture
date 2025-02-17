package internal_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"google.golang.org/protobuf/types/known/timestamppb"

	flowcontrolv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/flowcontrol/check/v1"
	otelconsts "github.com/fluxninja/aperture/pkg/otelcollector/consts"
	"github.com/fluxninja/aperture/pkg/otelcollector/metricsprocessor/internal"
)

var _ = DescribeTable("Check Response labels", func(checkResponse *flowcontrolv1.CheckResponse, after map[string]interface{}) {
	attributes := pcommon.NewMap()
	internal.AddCheckResponseBasedLabels(attributes, checkResponse, "source")
	for k, v := range after {
		Expect(attributes.AsRaw()).To(HaveKeyWithValue(k, v))
	}
},
	Entry("Sets processing duration",
		&flowcontrolv1.CheckResponse{
			Start: timestamppb.New(time.Date(1969, time.Month(7), 20, 17, 0, 0, 0, time.UTC)),
			End:   timestamppb.New(time.Date(1969, time.Month(7), 20, 17, 0, 1, 0, time.UTC)),
		},
		map[string]interface{}{otelconsts.ApertureProcessingDurationLabel: float64(1000)},
	),

	Entry("Sets services",
		&flowcontrolv1.CheckResponse{
			Services: []string{"svc1", "svc2"},
		},
		map[string]interface{}{otelconsts.ApertureServicesLabel: []interface{}{"svc1", "svc2"}},
	),

	Entry("Sets control point",
		&flowcontrolv1.CheckResponse{
			ControlPoint: "ingress",
		},
		map[string]interface{}{otelconsts.ApertureControlPointLabel: "ingress"},
	),

	Entry("Sets rate limiters",
		&flowcontrolv1.CheckResponse{
			LimiterDecisions: []*flowcontrolv1.LimiterDecision{
				{
					PolicyName:  "foo",
					PolicyHash:  "foo-hash",
					ComponentId: "2",
					Dropped:     true,
					Details: &flowcontrolv1.LimiterDecision_RateLimiterInfo_{
						RateLimiterInfo: &flowcontrolv1.LimiterDecision_RateLimiterInfo{
							Remaining: 1,
							Current:   1,
							Label:     "test",
						},
					},
				},
			},
		},
		map[string]interface{}{
			otelconsts.ApertureRateLimitersLabel:         []interface{}{"policy_name:foo,component_id:2,policy_hash:foo-hash"},
			otelconsts.ApertureDroppingRateLimitersLabel: []interface{}{"policy_name:foo,component_id:2,policy_hash:foo-hash"},
		},
	),

	Entry("Sets concurrency limiters",
		&flowcontrolv1.CheckResponse{
			LimiterDecisions: []*flowcontrolv1.LimiterDecision{
				{
					PolicyName:  "foo",
					PolicyHash:  "foo-hash",
					ComponentId: "1",
					Dropped:     true,
					Details: &flowcontrolv1.LimiterDecision_ConcurrencyLimiterInfo_{
						ConcurrencyLimiterInfo: &flowcontrolv1.LimiterDecision_ConcurrencyLimiterInfo{
							WorkloadIndex: "0",
						},
					},
				},
			},
		},
		map[string]interface{}{
			otelconsts.ApertureConcurrencyLimitersLabel:         []interface{}{"policy_name:foo,component_id:1,policy_hash:foo-hash"},
			otelconsts.ApertureDroppingConcurrencyLimitersLabel: []interface{}{"policy_name:foo,component_id:1,policy_hash:foo-hash"},
		},
	),

	Entry("Sets flux meters",
		&flowcontrolv1.CheckResponse{
			FluxMeterInfos: []*flowcontrolv1.FluxMeterInfo{
				{FluxMeterName: "foo"},
				{FluxMeterName: "bar"},
			},
		},
		map[string]interface{}{otelconsts.ApertureFluxMetersLabel: []interface{}{"foo", "bar"}},
	),

	Entry("Sets flow labels",
		&flowcontrolv1.CheckResponse{
			FlowLabelKeys: []string{"someLabel", "otherLabel"},
		},
		map[string]interface{}{otelconsts.ApertureFlowLabelKeysLabel: []interface{}{"someLabel", "otherLabel"}},
	),

	Entry("Sets classifiers",
		&flowcontrolv1.CheckResponse{
			ClassifierInfos: []*flowcontrolv1.ClassifierInfo{
				{
					PolicyName:      "foo",
					PolicyHash:      "bar",
					ClassifierIndex: 42,
					LabelKey:        "timing",
					Error:           flowcontrolv1.ClassifierInfo_ERROR_MULTI_EXPRESSION,
				},
			},
		},
		map[string]interface{}{
			otelconsts.ApertureClassifiersLabel:      []interface{}{"policy_name:foo,classifier_index:42"},
			otelconsts.ApertureClassifierErrorsLabel: []interface{}{"ERROR_MULTI_EXPRESSION,policy_name:foo,classifier_index:42,policy_hash:bar"},
		},
	),
)

var _ = Describe("AddFlowLabels", func() {
	attributes := pcommon.NewMap()
	checkResponse := &flowcontrolv1.CheckResponse{
		TelemetryFlowLabels: map[string]string{
			"someLabel":  "someValue",
			"otherLabel": "otherValue",
		},
	}
	internal.AddFlowLabels(attributes, checkResponse)
	Expect(attributes.AsRaw()).To(HaveKeyWithValue("someLabel", "someValue"))
	Expect(attributes.AsRaw()).To(HaveKeyWithValue("otherLabel", "otherValue"))
})
