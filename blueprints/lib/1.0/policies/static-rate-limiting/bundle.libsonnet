local blueprint = import './static-rate-limiting.libsonnet';

local policy = blueprint.policy;
local dashboard = blueprint.dashboard;
local config = blueprint.config;

{
  dashboards: {
    [std.format('%s.json', $._config.common.policy_name)]: dashboard($._config.common + $._config.dashboard).dashboard,
  },
  policies: {
    [std.format('%s.yaml', $._config.common.policy_name)]: policy($._config.common + $._config.policy).policyResource,
  },
} + { _config:: config }
