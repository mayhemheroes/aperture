{
  /**
  * @section Common
  *
  * @param (common.policy_name: string required) Name of the policy.
  */
  common: {
    policy_name: error 'policyName is not set',
  },
  /**
  * @section Policy
  *
  * @param (policy.flux_meter: aperture.spec.v1.FluxMeter required) Flux Meter.
  * @param (policy.classifiers: []aperture.spec.v1.Classifier) List of classification rules.
  * @param (policy.components: []aperture.spec.v1.Component) List of additional circuit components.
  */
  policy: {
    flux_meter: error 'flux_meter is not set',
    classifiers: [],
    components: [],
    /**
    * @section Policy
    * @subsection Latency Baseliner
    *
    * @param (policy.latency_baseliner.ema: aperture.spec.v1.EMAParameters) EMA parameters.
    * @param (policy.latency_baseliner.latency_tolerance_multiplier: float64) Tolerance factor beyond which the service is considered to be in overloaded state. E.g. if EMA of latency is 50ms and if Tolerance is 1.1, then service is considered to be in overloaded state if current latency is more than 55ms.
    * @param (policy.latency_baseliner.latency_ema_limit_multiplier: float64) Current latency value is multiplied with this factor to calculate maximum envelope of Latency EMA.
    */
    latency_baseliner: {
      ema: {
        ema_window: '1500s',
        warmup_window: '60s',
        correction_factor_on_max_envelope_violation: '0.95',
      },
      latency_tolerance_multiplier: '1.1',
      latency_ema_limit_multiplier: '2.0',
    },
    /**
    * @section Policy
    * @subsection Concurrency Controller
    *
    * @param (policy.concurrency_controller.flow_selector: aperture.spec.v1.FlowSelector required) Concurrency Limiter flow selector.
    * @param (policy.concurrency_controller.scheduler: aperture.spec.v1.SchedulerParameters) Scheduler parameters.
    * @param (policy.concurrency_controller.gradient: aperture.spec.v1.GradientParameters) Gradient parameters.
    * @param (policy.concurrency_controller.alerter: aperture.spec.v1.AlerterParameters) Whether tokens for workloads are computed dynamically or set statically by the user.
    * @param (policy.concurrency_controller.concurrency_limit_multiplier: float64) Current accepted concurrency is multiplied with this number to dynamically calculate the upper concurrency limit of a Service during normal (non-overload) state. This protects the Service from sudden spikes.
    * @param (policy.concurrency_controller.concurrency_linear_increment: float64) Linear increment to concurrency in each execution tick when the system is not in overloaded state.
    * @param (policy.concurrency_controller.concurrency_sqrt_increment_multiplier: float64) Scale factor to multiply square root of current accepted concurrrency. This, along with concurrency_linear_increment helps calculate overall concurrency increment in each tick. Concurrency is rapidly ramped up in each execution cycle during normal (non-overload) state (integral effect).
    * @param (policy.concurrency_controller.dynamic_config: aperture.v1.LoadActuatorDynamicConfig) Dynamic configuration for concurrency controller.
    */
    concurrency_controller: {
      flow_selector: error 'flow_selector for concurrencyController is not set',
      scheduler: {
        auto_tokens: true,
        timeout_factor: '0.5',
        default_workload_parameters: {
          priority: 20,
        },
        workloads: [],
      },
      gradient: {
        slope: '-1',
        min_gradient: '0.1',
        max_gradient: '1.0',
      },
      alerter: {
        alert_name: 'Load Shed Event',
        alert_channels: [],
        resolve_timeout: '5s',
      },
      concurrency_limit_multiplier: '2.0',
      concurrency_linear_increment: '5.0',
      concurrency_sqrt_increment_multiplier: '1',
      dynamic_config: {
        dry_run: false,
      },
    },
  },
  /**
  * @section Dashboard
  *
  * @param (dashboard.refresh_interval: string) Refresh interval for dashboard panels.
  */
  dashboard: {
    refresh_interval: '10s',
    /**
    * @section Dashboard
    * @subsection Datasource
    *
    * @param (dashboard.datasource.name: string) Datasource name.
    * @param (dashboard.datasource.filter_regex: string) Datasource filter regex.
    */
    datasource: {
      name: '$datasource',
      filter_regex: '',
    },
  },
}
