package controlplane_test

import (
	"context"
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	admissionv1 "k8s.io/api/admission/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"

	policyv1alpha1 "github.com/fluxninja/aperture/operator/api/policy/v1alpha1"
	"github.com/fluxninja/aperture/pkg/policies/controlplane"
	"github.com/fluxninja/aperture/pkg/webhooks/policyvalidator"
)

var _ = Describe("Validator", Ordered, func() {
	policySpecValidator := &controlplane.PolicySpecValidator{}
	policyValidator := policyvalidator.NewPolicyValidator([]policyvalidator.PolicySpecValidator{policySpecValidator})

	validateExample := func(contents string) {
		os.Setenv("APERTURE_CONTROLLER_NAMESPACE", "aperture-controller")
		jsonPolicy, err := yaml.YAMLToJSON([]byte(contents))
		Expect(err).ToNot(HaveOccurred())
		var policy policyv1alpha1.Policy
		err = json.Unmarshal([]byte(jsonPolicy), &policy)
		Expect(err).NotTo(HaveOccurred())
		request := &admissionv1.AdmissionRequest{
			Name:      policy.Name,
			Namespace: policy.Namespace,
			Kind:      v1.GroupVersionKind(policy.GroupVersionKind()),
			Object:    runtime.RawExtension{Raw: []byte(jsonPolicy)},
		}

		ok, msg, err := policyValidator.ValidateObject(context.TODO(), request)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(BeEmpty())
		Expect(ok).To(BeTrue())
	}

	It("accepts example policy for demoapp", func() {
		validateExample(latencyGradientPolicy)
	})

	It("accepts example policy for rate limit", func() {
		validateExample(rateLimitPolicy)
	})

	It("accepts example policy for classification", func() {
		validateExample(classificationPolicy)
	})

	It("does not accept policy in other namespace than controller", func() {
		os.Setenv("APERTURE_CONTROLLER_NAMESPACE", "")
		jsonPolicy, err := yaml.YAMLToJSON([]byte(rateLimitPolicy))
		Expect(err).ToNot(HaveOccurred())
		var policy policyv1alpha1.Policy
		err = json.Unmarshal([]byte(jsonPolicy), &policy)
		Expect(err).NotTo(HaveOccurred())
		request := &admissionv1.AdmissionRequest{
			Name:      policy.Name,
			Namespace: policy.Namespace,
			Kind:      v1.GroupVersionKind(policy.GroupVersionKind()),
			Object:    runtime.RawExtension{Raw: []byte(jsonPolicy)},
		}

		ok, msg, err := policyValidator.ValidateObject(context.TODO(), request)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg).To(Equal("Policy should be created in the same namespace as Aperture Controller"))
		Expect(ok).To(BeFalse())
	})
})

const latencyGradientPolicy = `
apiVersion: fluxninja.com/v1alpha1
kind: Policy
metadata:
  name: policies
  namespace: aperture-controller
  labels:
    fluxninja.com/validate: "true"
spec:
  resources:
    flux_meters:
      "service_latency":
        flow_selector:
          service_selector:
            service: "service1-demo-app.demoapp.svc.cluster.local"
          flow_matcher:
            control_point: "ingress"
    classifiers:
      - flow_selector:
          service_selector:
            service: service1-demo-app.demoapp.svc.cluster.local
          flow_matcher:
            control_point: ingress
        rules:
          # An example rule using extractor.
          # See following RFC for list of available extractors and their syntax.
          ua:
            extractor:
              from: request.http.headers.user-agent
          # The same rule using raw rego. Requires specifying rego source code and a query
          also-ua:
            rego:
              source: |
                package my.rego.pkg
                import input.attributes.request.http
                ua = http.headers["user-agent"]
              query: data.my.rego.pkg.ua
          user_type:
            extractor:
              from: request.http.headers.user_type
  circuit:
    evaluation_interval: "0.5s"
    components:
      - query:
          promql:
            query_string: "sum(increase(flux_meter_sum{decision_type!=\"DECISION_TYPE_REJECTED\", policy_name=\"latency-gradient\", flux_meter_name=\"service_latency\"}[5s]))/sum(increase(flux_meter_count{decision_type!=\"DECISION_TYPE_REJECTED\", policy_name=\"latency-gradient\", flux_meter_name=\"service_latency\"}[5s]))"
            evaluation_interval: "1s"
            out_ports:
              output:
                signal_name: "LATENCY"
      - variable:
          default_config:
            constant_signal:
              value: 2.0
          out_ports:
            output:
              signal_name: "EMA_LIMIT_MULTIPLIER"
      - arithmetic_combinator:
          operator: "mul"
          in_ports:
            lhs:
              signal_name: "LATENCY"
            rhs:
              signal_name: "EMA_LIMIT_MULTIPLIER"
          out_ports:
            output:
              signal_name: "MAX_EMA"
      - ema:
          parameters:
            ema_window: "300s"
            warmup_window: "10s"
            correction_factor_on_max_envelope_violation: "0.95"
          in_ports:
            input:
              signal_name: "LATENCY"
            max_envelope:
              signal_name: "MAX_EMA"
          out_ports:
            output:
              signal_name: "LATENCY_EMA"
      - variable:
          default_config:
            constant_signal:
              value: 1.1
          out_ports:
            output:
              signal_name: "EMA_SETPOINT_MULTIPLIER"
      - arithmetic_combinator:
          operator: "mul"
          in_ports:
            lhs:
              signal_name: "LATENCY_EMA"
            rhs:
              signal_name: "EMA_SETPOINT_MULTIPLIER"
          out_ports:
            output:
              signal_name: "LATENCY_SETPOINT"
      - gradient_controller:
          parameters:
            slope: -1
            min_gradient: "0.1"
            max_gradient: "1.0"
          in_ports:
            signal:
              signal_name: "LATENCY"
            setpoint:
              signal_name: "LATENCY_SETPOINT"
            max:
              signal_name: "MAX_CONCURRENCY"
            control_variable:
              signal_name: "ACCEPTED_CONCURRENCY"
            optimize:
              signal_name: "CONCURRENCY_INCREMENT"
          out_ports:
            output:
              signal_name: "DESIRED_CONCURRENCY"
      - arithmetic_combinator:
          operator: "div"
          in_ports:
            lhs:
              signal_name: "DESIRED_CONCURRENCY"
            rhs:
              signal_name: "INCOMING_CONCURRENCY"
          out_ports:
            output:
              signal_name: "LOAD_MULTIPLIER"
      - flow_control:
          concurrency_limiter:
            flow_selector:
              service_selector:
                service: "service1-demo-app.demoapp.svc.cluster.local"
              flow_matcher:
                control_point: "ingress"
            scheduler:
              parameters:
                auto_tokens: true
                default_workload_parameters:
                  priority: 20
                workloads:
                  - parameters:
                      priority: 50
                    label_matcher:
                      match_labels:
                        user_type: "guest"
                  - parameters:
                      priority: 200
                    label_matcher:
                      match_labels:
                        http.request.header.user_type: "subscriber"
              out_ports:
                accepted_concurrency:
                  signal_name: "ACCEPTED_CONCURRENCY"
                incoming_concurrency:
                  signal_name: "INCOMING_CONCURRENCY"
            load_actuator:
              in_ports:
                load_multiplier:
                  signal_name: "LOAD_MULTIPLIER"
      - variable:
          default_config:
            constant_signal:
              value: 2.0
          out_ports:
            output:
              signal_name: "CONCURRENCY_LIMIT_MULTIPLIER"
      - variable:
          default_config:
            constant_signal:
              value: 10.0
          out_ports:
            output:
              signal_name: "MIN_CONCURRENCY"
      - variable:
          default_config:
            constant_signal:
              value: 5.0
          out_ports:
            output:
              signal_name: "LINEAR_CONCURRENCY_INCREMENT"
      - arithmetic_combinator:
          operator: "mul"
          in_ports:
            lhs:
              signal_name: "CONCURRENCY_LIMIT_MULTIPLIER"
            rhs:
              signal_name: "ACCEPTED_CONCURRENCY"
          out_ports:
            output:
              signal_name: "UPPER_CONCURRENCY_LIMIT"
      - max:
          in_ports:
            inputs:
              - signal_name: "UPPER_CONCURRENCY_LIMIT"
              - signal_name: "MIN_CONCURRENCY"
          out_ports:
            output:
              signal_name: "MAX_CONCURRENCY"
      - sqrt:
          scale: "0.5"
          in_ports:
            input:
              signal_name: "ACCEPTED_CONCURRENCY"
          out_ports:
            output:
              signal_name: "SQRT_CONCURRENCY_INCREMENT"
      - arithmetic_combinator:
          operator: "add"
          in_ports:
            lhs:
              signal_name: "LINEAR_CONCURRENCY_INCREMENT"
            rhs:
              signal_name: "SQRT_CONCURRENCY_INCREMENT"
          out_ports:
            output:
              signal_name: "CONCURRENCY_INCREMENT_NORMAL"
      - variable:
          default_config:
            constant_signal:
              value: 1.2
          out_ports:
            output:
              signal_name: "OVERLOAD_MULTIPLIER"
      - arithmetic_combinator:
          operator: "mul"
          in_ports:
            lhs:
              signal_name: "LATENCY_EMA"
            rhs:
              signal_name: "OVERLOAD_MULTIPLIER"
          out_ports:
            output:
              signal_name: "LATENCY_OVERLOAD"
      - variable:
          default_config:
            constant_signal:
              value: 10.0
          out_ports:
            output:
              signal_name: "CONCURRENCY_INCREMENT_OVERLOAD"
      - decider:
          operator: "gt"
          in_ports:
            lhs:
              signal_name: "LATENCY"
            rhs:
              signal_name: "LATENCY_OVERLOAD"
          out_ports:
            output:
              signal_name: "CONCURRENCY_INCREMENT_DECISION"
      - switcher:
          in_ports:
            on_true:
              signal_name: "CONCURRENCY_INCREMENT_OVERLOAD"
            on_false:
              signal_name: "CONCURRENCY_INCREMENT_NORMAL"
            switch:
              signal_name: "CONCURRENCY_INCREMENT_DECISION"
          out_ports:
            output:
              signal_name: "CONCURRENCY_INCREMENT"
  `

const classificationPolicy = `
apiVersion: fluxninja.com/v1alpha1
kind: Policy
metadata:
  name: policies
  namespace: aperture-controller
  labels:
    fluxninja.com/validate: "true"
spec:
  resources:
    classifiers:
      - flow_selector:
          service_selector:
            service: productpage.bookinfo.svc.cluster.local
          flow_matcher:
            control_point: ingress
        rules:
          ua:
            extractor:
              from: request.http.headers.user-agent
          user:
            rego:
              query: data.user_from_cookie.user
              source: |
                package user_from_cookie
                cookies := split(input.attributes.request.http.headers.cookie, "; ")
                user := user {
                    cookie := cookies[_]
                    startswith(cookie, "session=")
                    session := substring(cookie, count("session="), -1)
                    parts := split(session, ".")
                    object := json.unmarshal(base64url.decode(parts[0]))
                    user := object.user
                }
`

const rateLimitPolicy = `
apiVersion: fluxninja.com/v1alpha1
kind: Policy
metadata:
  name: policies
  namespace: aperture-controller
  labels:
    fluxninja.com/validate: "true"
spec:
  circuit:
    evaluation_interval: "0.5s"
    components:
      - variable:
          default_config:
            constant_signal:
              value: 250.0
          out_ports:
            output:
              signal_name: "RATE_LIMIT"
      - flow_control:
          rate_limiter:
            in_ports:
              limit:
                signal_name: "RATE_LIMIT"
            flow_selector:
              service_selector:
                service: "service1-demo-app.demoapp.svc.cluster.local"
              flow_matcher:
                control_point: "ingress"
            parameters:
              label_key: "http.request.header.user_type"
              limit_reset_interval: "1s"
`
