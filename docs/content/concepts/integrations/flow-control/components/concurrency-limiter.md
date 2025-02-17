---
title: Concurrency Limiter
keywords:
  - scheduler
  - tokens
  - priority
  - fairness
  - queuing
  - actuators
sidebar_position: 6
---

:::info

See also
[Concurrency Limiter reference](/references/configuration/policy.md#v1-concurrency-limiter).

:::

Concurrency Limiter is about protecting your services from overload. Its goal is
to limit number of concurrent requests to service to a level the service can
handle. It's implemented by configurable load-shedding. Thanks to the ability to
define workloads of different priorities and weights, it allows to shed some
“less useful” flows, while not affecting the more important ones.

Concurrency Limiter is configured as a [policy][policies] component.

## Scheduler {#scheduler}

Each Aperture Agent instantiates a
[Weighted Fair Queueing](https://en.wikipedia.org/wiki/Weighted_fair_queueing)
based Scheduler as a way to prioritize flows based on their weights (priority)
and size(tokens). Concurrency Limiter applies a load-shed-factor that the
Scheduler uses to compute a level of [tokens](#tokens) per second, which it
tries to maintain.

If rate of tokens in flows entering the scheduler exceeds the desired rate,
flows are queued in the scheduler. If a flow can't be scheduled within its
specified timeout, it will be rejected.

### Workload {#workload}

Workloads are groups of flows based on common attributes. Workloads are
expressed by [label matcher][label-matcher] rules in Aperture. Aperture Agents
schedule workloads based on their priorities and by estimating their
[tokens](#tokens).

### Priority {#priority}

Priority represents the importance of a flow with respect to other flows in the
queue.

:::note

Priority levels are in the range `0 to 255`. `0` is the lowest priority and
`255` is the highest priority.

:::

### Tokens {#tokens}

Tokens represent the cost of admitting a flow in the system. Most commonly,
tokens are estimated based on milliseconds of response time observed when a flow
is processed. Token estimation of flows within a workload is crucial when making
flow control decisions. The concept of tokens is aligned with
[Little's Law](https://en.wikipedia.org/wiki/Little%27s_law), which defines the
relationship between response times, arrival rate and the number of requests
currently in the system (concurrency).

In some cases, tokens can be represented as the number of requests instead of
response times, e.g. when performing flow control on external APIs that have
hard rate-limits.

Aperture can be configured to automatically estimate the tokens for each
workload. See `auto-tokens`
[configuration](/references/configuration/policy.md#v1-scheduler).

### Token bucket {#token-bucket}

Aperture Agents use a variant of a
[token bucket algorithm](https://en.wikipedia.org/wiki/Token_bucket) to control
the flows entering the system. Each flow has to acquire tokens from the bucket
within a deadline period in order to be admitted.

### Timeout Factor {#timeout-factor}

The timeout factor parameter decides how long a request in the workload can wait
for tokens. This value impacts fairness because the larger the timeout the
higher the chance a request has to get scheduled.

The timeout is calculated as `timeout = timeout_factor * tokens`.

:::info

It's advisable to configure the timeouts in the same order of magnitude as the
normal latency of the workload flows in order to protect from retry storms
during overload scenarios.

:::

[label-matcher]: ../flow-selector.md#label-matcher
[policies]: /concepts/policy/policy.md
