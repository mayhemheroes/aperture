local nestedsignalegressins = import './nestedsignalegressins.libsonnet';
{
  new():: {
    in_ports: {
      signal: error 'Port signal is missing',
    },
  },
  inPorts:: nestedsignalegressins,
  withInPorts(in_ports):: {
    in_ports: in_ports,
  },
  withInPortsMixin(in_ports):: {
    in_ports+: in_ports,
  },
  withPortName(port_name):: {
    port_name: port_name,
  },
  withPortNameMixin(port_name):: {
    port_name+: port_name,
  },
}
