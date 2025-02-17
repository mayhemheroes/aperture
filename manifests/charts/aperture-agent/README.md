# Aperture Agent

## Parameters

### Global Parameters

| Name                      | Description                                     | Value |
| ------------------------- | ----------------------------------------------- | ----- |
| `global.imageRegistry`    | Global Docker image registry                    | `nil` |
| `global.imagePullSecrets` | Global Docker registry secret names as an array | `[]`  |


### Common Parameters

| Name                | Description                                                          | Value           |
| ------------------- | -------------------------------------------------------------------- | --------------- |
| `kubeVersion`       | Force target Kubernetes version (using Helm capabilities if not set) | `""`            |
| `nameOverride`      | String to partially override common.names.name                       | `""`            |
| `fullnameOverride`  | String to fully override common.names.fullname                       | `""`            |
| `namespaceOverride` | String to fully override common.names.namespace                      | `""`            |
| `extraDeploy`       | Array of extra objects to deploy with the release                    | `[]`            |
| `commonLabels`      | Labels to add to all deployed objects                                | `{}`            |
| `commonAnnotations` | Annotations to add to all deployed objects                           | `{}`            |
| `clusterDomain`     | Kubernetes cluster domain name                                       | `cluster.local` |


### Operator Parameters

| Name                                                         | Description                                                                                                            | Value                 |
| ------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------- | --------------------- |
| `operator.image.registry`                                    | Operator image registry                                                                                                | `docker.io/fluxninja` |
| `operator.image.repository`                                  | Operator image repository                                                                                              | `aperture-operator`   |
| `operator.image.tag`                                         | Operator image tag (immutable tags are recommended)                                                                    | `nil`                 |
| `operator.image.pullPolicy`                                  | Operator image pull policy                                                                                             | `Always`              |
| `operator.image.pullSecrets`                                 | Operator image pull secrets                                                                                            | `[]`                  |
| `operator.replicaCount`                                      | Number of replicas for Operator deployment                                                                             | `1`                   |
| `operator.podAnnotations`                                    | Pod annotations                                                                                                        | `{}`                  |
| `operator.podLabels`                                         | Additional pod labels                                                                                                  | `{}`                  |
| `operator.updateStrategy.type`                               | Set up update strategy for Aperture Operator installation.                                                             | `RollingUpdate`       |
| `operator.updateStrategy.rollingUpdate.maxSurge`             | Set up maximum number of Pods that can be created over the desired number of Pods.                                     | `25%`                 |
| `operator.updateStrategy.rollingUpdate.maxUnavailable`       | Set up maximum number of Pods that can be unavailable during the update process.                                       | `25%`                 |
| `operator.priorityClassName`                                 | pods' priorityClassName                                                                                                | `""`                  |
| `operator.topologySpreadConstraints`                         | Topology Spread Constraints for pod assignment                                                                         | `[]`                  |
| `operator.schedulerName`                                     | Alternative scheduler                                                                                                  | `""`                  |
| `operator.hostAliases`                                       | Add deployment host aliases                                                                                            | `[]`                  |
| `operator.nodeSelector`                                      | Node labels for pod assignment                                                                                         | `{}`                  |
| `operator.podAffinityPreset`                                 | Pod affinity preset                                                                                                    | `""`                  |
| `operator.podAntiAffinityPreset`                             | Pod anti-affinity preset                                                                                               | `soft`                |
| `operator.nodeAffinityPreset.type`                           | Set nodeAffinity preset type                                                                                           | `""`                  |
| `operator.nodeAffinityPreset.key`                            | Set nodeAffinity preset key                                                                                            | `""`                  |
| `operator.nodeAffinityPreset.values`                         | Set nodeAffinity preset values                                                                                         | `[]`                  |
| `operator.affinity`                                          | Affinity for controller pod assignment                                                                                 | `{}`                  |
| `operator.tolerations`                                       | Tolerations for controller pod assignment                                                                              | `[]`                  |
| `operator.podSecurityContext.enabled`                        | Enable pods security context                                                                                           | `true`                |
| `operator.podSecurityContext.runAsUser`                      | User ID for the pods                                                                                                   | `1001`                |
| `operator.podSecurityContext.runAsGroup`                     | User ID for the pods                                                                                                   | `1001`                |
| `operator.podSecurityContext.runAsNonRoot`                   | Aperture Operator must run as nonRoot                                                                                  | `true`                |
| `operator.podSecurityContext.fsGroup`                        | Group ID for the pods                                                                                                  | `1001`                |
| `operator.podSecurityContext.supplementalGroups`             | Which group IDs containers add                                                                                         | `[]`                  |
| `operator.containerSecurityContext.enabled`                  | Enable container security context                                                                                      | `true`                |
| `operator.containerSecurityContext.runAsUser`                | User ID for the operator container                                                                                     | `1001`                |
| `operator.containerSecurityContext.runAsGroup`               | User ID for the operator container                                                                                     | `1001`                |
| `operator.containerSecurityContext.runAsNonRoot`             | Force the container to be run as non-root                                                                              | `true`                |
| `operator.containerSecurityContext.privileged`               | Decide if the container runs privileged.                                                                               | `false`               |
| `operator.containerSecurityContext.readOnlyRootFilesystem`   | ReadOnlyRootFilesystem for the operator container                                                                      | `false`               |
| `operator.containerSecurityContext.allowPrivilegeEscalation` | Allow Privilege Escalation for the operator container                                                                  | `false`               |
| `operator.terminationGracePeriodSeconds`                     | In seconds, time the given to the pod needs to terminate gracefully                                                    | `10`                  |
| `operator.initContainers`                                    | Add additional init containers to the etcd pods                                                                        | `[]`                  |
| `operator.command`                                           | Default container command (useful when using custom images)                                                            | `[]`                  |
| `operator.args`                                              | Default container args (useful when using custom images)                                                               | `[]`                  |
| `operator.lifecycleHooks`                                    | for the aperture-operator container to automate configuration before or after startup                                  | `{}`                  |
| `operator.extraEnvVars`                                      | Array with extra environment variables to add to RabbitMQ Cluster Operator nodes                                       | `[]`                  |
| `operator.extraEnvVarsCM`                                    | Name of existing ConfigMap containing extra env vars for RabbitMQ Cluster Operator nodes                               | `""`                  |
| `operator.extraEnvVarsSecret`                                | Name of existing Secret containing extra env vars for RabbitMQ Cluster Operator nodes                                  | `""`                  |
| `operator.resources`                                         | Container resource requests and limits                                                                                 | `{}`                  |
| `operator.livenessProbe.enabled`                             | Enable livenessProbe                                                                                                   | `true`                |
| `operator.livenessProbe.initialDelaySeconds`                 | Initial delay seconds for livenessProbe                                                                                | `10`                  |
| `operator.livenessProbe.periodSeconds`                       | Period seconds for livenessProbe                                                                                       | `10`                  |
| `operator.livenessProbe.timeoutSeconds`                      | Timeout seconds for livenessProbe                                                                                      | `1`                   |
| `operator.livenessProbe.failureThreshold`                    | Failure threshold for livenessProbe                                                                                    | `3`                   |
| `operator.livenessProbe.successThreshold`                    | Success threshold for livenessProbe                                                                                    | `1`                   |
| `operator.readinessProbe.enabled`                            | Enable readinessProbe                                                                                                  | `true`                |
| `operator.readinessProbe.initialDelaySeconds`                | Initial delay seconds for readinessProbe                                                                               | `10`                  |
| `operator.readinessProbe.periodSeconds`                      | Period seconds for readinessProbe                                                                                      | `10`                  |
| `operator.readinessProbe.timeoutSeconds`                     | Timeout seconds for readinessProbe                                                                                     | `1`                   |
| `operator.readinessProbe.failureThreshold`                   | Failure threshold for readinessProbe                                                                                   | `3`                   |
| `operator.readinessProbe.successThreshold`                   | Success threshold for readinessProbe                                                                                   | `1`                   |
| `operator.startupProbe.enabled`                              | Enable startupProbe                                                                                                    | `true`                |
| `operator.startupProbe.initialDelaySeconds`                  | Initial delay seconds for startupProbe                                                                                 | `10`                  |
| `operator.startupProbe.periodSeconds`                        | Period seconds for startupProbe                                                                                        | `10`                  |
| `operator.startupProbe.timeoutSeconds`                       | Timeout seconds for startupProbe                                                                                       | `1`                   |
| `operator.startupProbe.failureThreshold`                     | Failure threshold for startupProbe                                                                                     | `3`                   |
| `operator.startupProbe.successThreshold`                     | Success threshold for startupProbe                                                                                     | `1`                   |
| `operator.customLivenessProbe`                               | Override default liveness probe                                                                                        | `{}`                  |
| `operator.customReadinessProbe`                              | Override default readiness probe                                                                                       | `{}`                  |
| `operator.customStartupProbe`                                | Override default startup probe                                                                                         | `{}`                  |
| `operator.extraVolumes`                                      | Optionally specify extra list of additional volumes                                                                    | `[]`                  |
| `operator.extraVolumeMounts`                                 | Optionally specify extra list of additional volumeMounts                                                               | `[]`                  |
| `operator.rbac.create`                                       | Create specifies whether to install and use RBAC rules                                                                 | `true`                |
| `operator.serviceAccount.create`                             | Specifies whether a service account should be created                                                                  | `true`                |
| `operator.serviceAccount.name`                               | The name of the service account to use. If not set and create is true, a name is generated using the fullname template | `""`                  |
| `operator.serviceAccount.annotations`                        | Add annotations                                                                                                        | `{}`                  |
| `operator.serviceAccount.automountServiceAccountToken`       | Automount API credentials for a service account.                                                                       | `true`                |


### Agent Custom Resource Parameters

| Name                                                    | Description                                                                                                     | Value    |
| ------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | -------- |
| `agent.create`                                          | Specifies whether a CR for Agent should be created                                                              | `true`   |
| `agent.image.registry`                                  | Agent image registry. Defaults to 'docker.io/fluxninja'.                                                        | `nil`    |
| `agent.image.repository`                                | Agent image repository. Defaults to 'aperture-agent'.                                                           | `nil`    |
| `agent.image.tag`                                       | Agent image tag (immutable tags are recommended). Defaults to 'latest'.                                         | `nil`    |
| `agent.image.pullPolicy`                                | Agent image pull policy. Defaults to 'IfNotPresent'.                                                            | `nil`    |
| `agent.image.pullSecrets`                               | Agent image pull secrets                                                                                        | `[]`     |
| `agent.service.annotations`                             | Additional custom annotations for Agent service                                                                 | `{}`     |
| `agent.serviceAccount.create`                           | Specifies whether a ServiceAccount should be created                                                            | `true`   |
| `agent.serviceAccount.annotations`                      | Additional Service Account annotations (evaluated as a template)                                                | `{}`     |
| `agent.serviceAccount.automountServiceAccountToken`     | Automount service account token for the server service account. Defaults to true                                | `nil`    |
| `agent.livenessProbe.enabled`                           | Enable livenessProbe on Agent containers                                                                        | `true`   |
| `agent.livenessProbe.initialDelaySeconds`               | Initial delay seconds for livenessProbe. Defaults to 15.                                                        | `nil`    |
| `agent.livenessProbe.periodSeconds`                     | Period seconds for livenessProbe. Defaults to 15.                                                               | `nil`    |
| `agent.livenessProbe.timeoutSeconds`                    | Timeout seconds for livenessProbe. Defaults to 5.                                                               | `nil`    |
| `agent.livenessProbe.failureThreshold`                  | Failure threshold for livenessProbe. Defaults to 6.                                                             | `nil`    |
| `agent.livenessProbe.successThreshold`                  | Success threshold for livenessProbe. Defaults to 1.                                                             | `nil`    |
| `agent.readinessProbe.enabled`                          | Enable readinessProbe on Agent containers                                                                       | `true`   |
| `agent.readinessProbe.initialDelaySeconds`              | Initial delay seconds for readinessProbe. Defaults to 15.                                                       | `nil`    |
| `agent.readinessProbe.periodSeconds`                    | Period seconds for readinessProbe. Defaults to 15.                                                              | `nil`    |
| `agent.readinessProbe.timeoutSeconds`                   | Timeout seconds for readinessProbe. Defaults to 5.                                                              | `nil`    |
| `agent.readinessProbe.failureThreshold`                 | Failure threshold for readinessProbe. Defaults to 6.                                                            | `nil`    |
| `agent.readinessProbe.successThreshold`                 | Success threshold for readinessProbe. Defaults to 1.                                                            | `nil`    |
| `agent.customLivenessProbe`                             | Custom livenessProbe that overrides the default one                                                             | `{}`     |
| `agent.customReadinessProbe`                            | Custom readinessProbe that overrides the default one                                                            | `{}`     |
| `agent.resources.limits`                                | The resources limits for the Agent containers                                                                   | `{}`     |
| `agent.resources.requests`                              | The requested resources for the Agent containers                                                                | `{}`     |
| `agent.podSecurityContext.enabled`                      | Enabled Agent pods' Security Context                                                                            | `false`  |
| `agent.podSecurityContext.fsGroup`                      | Set Agent pod's Security Context fsGroup. Defaults to 1001.                                                     | `nil`    |
| `agent.containerSecurityContext.enabled`                | Enabled Agent containers' Security Context. Defaults to false.                                                  | `false`  |
| `agent.containerSecurityContext.runAsUser`              | Set Agent containers' Security Context runAsUser. Defaults to 1001.                                             | `nil`    |
| `agent.containerSecurityContext.runAsNonRoot`           | Set Agent containers' Security Context runAsNonRoot. Defaults to false.                                         | `nil`    |
| `agent.containerSecurityContext.readOnlyRootFilesystem` | Set Agent containers' Security Context runAsNonRoot. Defaults to false.                                         | `nil`    |
| `agent.command`                                         | Override default container command (useful when using custom images)                                            | `[]`     |
| `agent.args`                                            | Override default container args (useful when using custom images)                                               | `[]`     |
| `agent.podLabels`                                       | Extra labels for Agent pods                                                                                     | `{}`     |
| `agent.podAnnotations`                                  | Annotations for Agent pods                                                                                      | `{}`     |
| `agent.affinity`                                        | Affinity for Agent pods assignment                                                                              | `{}`     |
| `agent.nodeSelector`                                    | Node labels for Agent pods assignment                                                                           | `{}`     |
| `agent.tolerations`                                     | Tolerations for Agent pods assignment                                                                           | `[]`     |
| `agent.terminationGracePeriodSeconds`                   | configures how long kubelet gives Agent chart to terminate cleanly                                              | `nil`    |
| `agent.lifecycleHooks`                                  | for the Agent container(s) to automate configuration before or after startup                                    | `{}`     |
| `agent.extraEnvVars`                                    | Array with extra environment variables to add to Agent nodes                                                    | `[]`     |
| `agent.extraEnvVarsCM`                                  | Name of existing ConfigMap containing extra env vars for Agent nodes                                            | `""`     |
| `agent.extraEnvVarsSecret`                              | Name of existing Secret containing extra env vars for Agent nodes                                               | `""`     |
| `agent.extraVolumes`                                    | Optionally specify extra list of additional volumes for the Agent pod(s)                                        | `[]`     |
| `agent.extraVolumeMounts`                               | Optionally specify extra list of additional volumeMounts for the Agent container(s)                             | `[]`     |
| `agent.sidecars`                                        | Add additional sidecar containers to the Agent pod(s)                                                           | `[]`     |
| `agent.initContainers`                                  | Add additional init containers to the Agent pod(s)                                                              | `[]`     |
| `agent.secrets.fluxNinjaPlugin.create`                  | Whether to create Kubernetes Secret with provided Agent API Key.                                                | `false`  |
| `agent.secrets.fluxNinjaPlugin.secretKeyRef.name`       | specifies a name of the Secret for Agent API Key to be used. This defaults to {{ .Release.Name }}-agent-apikey  | `nil`    |
| `agent.secrets.fluxNinjaPlugin.secretKeyRef.key`        | specifies which key from the Secret for Agent API Key to use                                                    | `apiKey` |
| `agent.secrets.fluxNinjaPlugin.value`                   | API Key to use when creating a new Agent API Key Secret                                                         | `nil`    |
| `agent.sidecar.enabled`                                 | Enables sidecar mode for the Agent                                                                              | `false`  |
| `agent.sidecar.enableNamespacesByDefault`               | List of namespaces in which sidecar injection will be enabled when Sidecar mode is enabled.                     | `[]`     |
| `agent.config.etcd.endpoints`                           | List of Etcd server endpoints. Example, ["http://etcd:2379"]. This must not be empty.                           | `[]`     |
| `agent.config.etcd.lease_ttl`                           | Lease time-to-live.                                                                                             | `60s`    |
| `agent.config.prometheus.address`                       | specifies the address of the Prometheus server. Example, "http://prometheus-server:80". This must not be empty. | `nil`    |


