local schedulerouts = import './schedulerouts.libsonnet';
{
  new():: {
    out_ports: {
      accepted_concurrency: error 'Port accepted_concurrency is missing',
      incoming_concurrency: error 'Port incoming_concurrency is missing',
    },
  },
  outPorts:: schedulerouts,
  withOutPorts(out_ports):: {
    out_ports: out_ports,
  },
  withOutPortsMixin(out_ports):: {
    out_ports+: out_ports,
  },
  withSchedulerParameters(scheduler_parameters):: {
    scheduler_parameters: scheduler_parameters,
  },
  withSchedulerParametersMixin(scheduler_parameters):: {
    scheduler_parameters+: scheduler_parameters,
  },
}
