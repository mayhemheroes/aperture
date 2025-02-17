/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// +kubebuilder:object:generate=true
package common

import (
	amclient "github.com/fluxninja/aperture/pkg/alertmanager/client"
	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/discovery/kubernetes"
	"github.com/fluxninja/aperture/pkg/discovery/static"
	etcd "github.com/fluxninja/aperture/pkg/etcd/client"
	"github.com/fluxninja/aperture/pkg/jobs"
	"github.com/fluxninja/aperture/pkg/metrics"
	"github.com/fluxninja/aperture/pkg/net/grpc"
	"github.com/fluxninja/aperture/pkg/net/grpcgateway"
	"github.com/fluxninja/aperture/pkg/net/http"
	"github.com/fluxninja/aperture/pkg/net/listener"
	"github.com/fluxninja/aperture/pkg/net/tlsconfig"
	"github.com/fluxninja/aperture/pkg/plugins"
	"github.com/fluxninja/aperture/pkg/profilers"
	"github.com/fluxninja/aperture/pkg/prometheus"
	"github.com/fluxninja/aperture/pkg/watchdog"
	"github.com/fluxninja/aperture/plugins/service/aperture-plugin-fluxninja/pluginconfig"
	"github.com/fluxninja/aperture/plugins/service/aperture-plugin-sentry/sentry"

	corev1 "k8s.io/api/core/v1"
)

// Image defines the Registry, Repository, Tag, PullPolicy, PullSecrets and Debug.
type Image struct {
	// The registry of the image
	//+kubebuilder:validation:Optional
	Registry string `json:"registry" default:"docker.io/fluxninja"`

	// The tag (version) of the image
	//+kubebuilder:validation:Optional
	Tag string `json:"tag" default:"latest"`

	// The ImagePullPolicy of the image
	//+kubebuilder:validation:Optional
	PullPolicy string `json:"pullPolicy" default:"IfNotPresent" validate:"oneof=Never Always IfNotPresent"`

	// The PullSecrets for the image
	//+kubebuilder:validation:Optional
	PullSecrets []string `json:"pullSecrets,omitempty"`
}

// AgentImage defines Image spec for Aperture Agent.
type AgentImage struct {
	// Image specs for Agent
	Image `json:",inline"`

	// The repository of the image
	//+kubebuilder:validation:Optional
	Repository string `json:"repository" default:"aperture-agent"`
}

// ControllerImage defines Image spec for Aperture Controller.
type ControllerImage struct {
	// Image specs for Controller
	Image `json:",inline"`

	// The repository of the image
	//+kubebuilder:validation:Optional
	Repository string `json:"repository" default:"aperture-controller"`
}

// Probe defines Enabled, InitialDelaySeconds, PeriodSeconds, TimeoutSeconds, FailureThreshold and SuccessThreshold for probes like livenessProbe.
type Probe struct {
	// Enable probe on agent containers
	Enabled bool `json:"enabled" default:"true"`

	// Initial delay seconds for probe
	//+kubebuilder:validation:Optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds" default:"15" validate:"gte=0"`

	// Period delay seconds for probe
	//+kubebuilder:validation:Optional
	PeriodSeconds int32 `json:"periodSeconds" default:"15" validate:"gte=1"`

	// Timeout delay seconds for probe
	//+kubebuilder:validation:Optional
	TimeoutSeconds int32 `json:"timeoutSeconds" default:"5" validate:"gte=1"`

	// Failure threshold for probe
	//+kubebuilder:validation:Optional
	FailureThreshold int32 `json:"failureThreshold" default:"6" validate:"gte=1"`

	// Success threshold for probe
	//+kubebuilder:validation:Optional
	SuccessThreshold int32 `json:"successThreshold" default:"1" validate:"gte=1"`
}

// PodSecurityContext defines Enabled and FsGroup for the Pods' security context.
type PodSecurityContext struct {
	// Enable PodSecurityContext on Pod
	Enabled bool `json:"enabled" default:"false"`

	// fsGroup to define the Group ID for the Pod
	//+kubebuilder:validation:Optional
	FsGroup int64 `json:"fsGroup" default:"1001" validate:"gte=0"`
}

// ContainerSecurityContext defines Enabled, RunAsUser, RunAsNonRootUser and ReadOnlyRootFilesystem for the containers' security context.
type ContainerSecurityContext struct {
	// Enable ContainerSecurityContext on containers
	Enabled bool `json:"enabled" default:"false"`

	// Set containers' Security Context runAsUser
	//+kubebuilder:validation:Optional
	RunAsUser int64 `json:"runAsUser" default:"1001" validate:"gte=0"`

	// Set containers' Security Context runAsNonRoot
	//+kubebuilder:validation:Optional
	RunAsNonRootUser bool `json:"runAsNonRoot" default:"false"`

	// Set agent containers' Security Context runAsNonRoot
	//+kubebuilder:validation:Optional
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem" default:"false"`
}

// CommonSpec defines the desired the common state of Agent and Controller.
type CommonSpec struct {
	// Labels to add to all deployed objects
	//+mapType=atomic
	//+kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to add to all deployed objects
	//+kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Configuration for Agent or Controller service
	//+kubebuilder:validation:Optional
	Service Service `json:"service"`

	// ServiceAccountSpec defines the the configuration pf Service account for Agent or Controller
	//+kubebuilder:validation:Optional
	ServiceAccountSpec ServiceAccountSpec `json:"serviceAccount"`

	// livenessProbe related configuration
	//+kubebuilder:validation:Optional
	LivenessProbe Probe `json:"livenessProbe"`

	// readinessProbe related configuration
	//+kubebuilder:validation:Optional
	ReadinessProbe Probe `json:"readinessProbe"`

	// Custom livenessProbe that overrides the default one
	//+kubebuilder:validation:Optional
	CustomLivenessProbe *corev1.Probe `json:"customLivenessProbe,omitempty"`

	// Custom readinessProbe that overrides the default one
	//+kubebuilder:validation:Optional
	CustomReadinessProbe *corev1.Probe `json:"customReadinessProbe,omitempty"`

	// Resource requests and limits
	//+kubebuilder:validation:Optional
	Resources corev1.ResourceRequirements `json:"resources"`

	// Configure Pods' Security Context
	//+kubebuilder:validation:Optional
	PodSecurityContext PodSecurityContext `json:"podSecurityContext"`

	// Configure Containers' Security Context
	//+kubebuilder:validation:Optional
	ContainerSecurityContext ContainerSecurityContext `json:"containerSecurityContext"`

	// Override default container command
	//+kubebuilder:validation:Optional
	Command []string `json:"command,omitempty"`

	// Override default container args
	//+kubebuilder:validation:Optional
	Args []string `json:"args,omitempty"`

	// Extra labels for pods
	//+mapType=atomic
	//+kubebuilder:validation:Optional
	PodLabels map[string]string `json:"podLabels,omitempty"`

	// Extra Annotations for pods
	//+kubebuilder:validation:Optional
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`

	// Affinity for pods assignment.
	//+kubebuilder:validation:Optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// Node labels for pods assignment
	//+kubebuilder:validation:Optional
	//+mapType=atomic
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Tolerations for pods assignment
	//+kubebuilder:validation:Optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Seconds Redmine pod needs to terminate gracefully
	//+kubebuilder:validation:Optional
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds" validate:"gte=0"`

	// For the container(s) to automate configuration before or after startup
	//+kubebuilder:validation:Optional
	LifecycleHooks *corev1.Lifecycle `json:"lifecycleHooks,omitempty" validate:"omitempty"`

	// Array with extra environment variables to add
	//+kubebuilder:validation:Optional
	ExtraEnvVars []corev1.EnvVar `json:"extraEnvVars,omitempty"`

	// Name of existing ConfigMap containing extra env vars
	//+kubebuilder:validation:Optional
	ExtraEnvVarsCM string `json:"extraEnvVarsCM"`

	// Name of existing Secret containing extra env vars
	//+kubebuilder:validation:Optional
	ExtraEnvVarsSecret string `json:"extraEnvVarsSecret"`

	// Optionally specify extra list of additional volumes for the pod(s)
	//+kubebuilder:validation:Optional
	ExtraVolumes []corev1.Volume `json:"extraVolumes,omitempty"`

	// Optionally specify extra list of additional volumeMounts
	//+kubebuilder:validation:Optional
	ExtraVolumeMounts []corev1.VolumeMount `json:"extraVolumeMounts,omitempty"`

	// Add additional sidecar containers
	//+kubebuilder:validation:Optional
	Sidecars []corev1.Container `json:"sidecars,omitempty"`

	// Add additional init containers
	//+kubebuilder:validation:Optional
	InitContainers []corev1.Container `json:"initContainers,omitempty"`

	// Secrets
	//+kubebuilder:validation:Optional
	Secrets Secrets `json:"secrets"`
}

// Secrets for Agent or Controller.
type Secrets struct {
	// FluxNinja plugin.
	//+kubebuilder:validation:Optional
	FluxNinjaPlugin APIKeySecret `json:"fluxNinjaPlugin"`
}

// APIKeySecret defines fields required for creation/usage of secret for the ApiKey of Agent and Controller.
type APIKeySecret struct {
	// Create new secret or not
	Create bool `json:"create" default:"false"`

	// Secret details
	//+kubebuilder:validation:Optional
	SecretKeyRef SecretKeyRef `json:"secretKeyRef"`

	// Value for the ApiKey
	Value string `json:"value"`
}

// SecretKeyRef defines fields for details of the ApiKey secret.
type SecretKeyRef struct {
	// Name of the secret
	//+kubebuilder:validation:Optional
	Name string `json:"name"`

	// Key of the secret in Data
	//+kubebuilder:validation:Optional
	Key string `json:"key" default:"apiKey"`
}

// APIKeySecretSpec defines API Key secret details for Agent and Controller.
type APIKeySecretSpec struct {
	// API Key secret reference for Agent
	//+kubebuilder:validation:Optional
	Agent APIKeySecret `json:"agent"`

	// API Key secret reference for Controller
	//+kubebuilder:validation:Optional
	Controller APIKeySecret `json:"controller"`
}

// CommonConfigSpec defines common configuration for agent and controller.
type CommonConfigSpec struct {
	// Client configuration such as proxy settings.
	//+kubebuilder:validation:Optional
	Client ClientConfigSpec `json:"client"`

	// Etcd configuration.
	//+kubebuilder:validation:Required
	Etcd etcd.EtcdConfig `json:"etcd"`

	// Liveness probe configuration.
	//+kubebuilder:validation:Optional
	Liveness ProbeConfigSpec `json:"liveness"`

	// Readiness probe configuration.
	//+kubebuilder:validation:Optional
	Readiness ProbeConfigSpec `json:"readiness"`

	// Log configuration.
	//+kubebuilder:validation:Optional
	Log config.LogConfig `json:"log"`

	// Metrics configuration.
	//+kubebuilder:validation:Optional
	Metrics metrics.MetricsConfig `json:"metrics"`

	// Plugins configuration.
	//+kubebuilder:validation:Optional
	Plugins plugins.PluginsConfig `json:"plugins"`

	// Profilers configuration.
	//+kubebuilder:validation:Optional
	Profilers profilers.ProfilersConfig `json:"profilers"`

	// Prometheus configuration.
	//+kubebuilder:validation:Required
	Prometheus prometheus.PrometheusConfig `json:"prometheus"`

	// Server configuration.
	//+kubebuilder:validation:Optional
	Server ServerConfigSpec `json:"server"`

	// Watchdog configuration.
	//+kubebuilder:validation:Optional
	Watchdog watchdog.WatchdogConfig `json:"watchdog"`

	// Alert Managers configuration.
	//+kubebuilder:validation:Optional
	Alertmanagers amclient.AlertManagerConfig `json:"alertmanagers,omitempty"`

	// BundledPluginsSpec defines configuration for bundled plugins.
	//+kubebuilder:validation:Optional
	BundledPluginsSpec `json:",inline"`
}

// ServerConfigSpec configures main server.
type ServerConfigSpec struct {
	// Listener configuration.
	//+kubebuilder:validation:Optional
	listener.ListenerConfig `json:",inline"`

	// GRPC server configuration.
	//+kubebuilder:validation:Optional
	Grpc grpc.GRPCServerConfig `json:"grpc"`

	// GRPC Gateway configuration.
	//+kubebuilder:validation:Optional
	GrpcGateway grpcgateway.GRPCGatewayConfig `json:"grpc_gateway"`

	// HTTP server configuration.
	//+kubebuilder:validation:Optional
	HTTP http.HTTPServerConfig `json:"http"`

	// TLS configuration.
	//+kubebuilder:validation:Optional
	TLS tlsconfig.ServerTLSConfig `json:"tls"`
}

// ProbeConfigSpec defines liveness and readiness probe configuration.
type ProbeConfigSpec struct {
	// Scheduler settings.
	//+kubebuilder:validation:Optional
	Scheduler jobs.JobGroupConfig `json:"scheduler"`

	// Service settings.
	//+kubebuilder:validation:Optional
	Service jobs.JobConfig `json:"service"`
}

// ClientConfigSpec defines client configuration.
type ClientConfigSpec struct {
	// Proxy settings for the client.
	//+kubebuilder:validation:Optional
	Proxy http.ProxyConfig `json:"proxy"`
}

// BundledPluginsSpec defines configuration for bundled plugins.
type BundledPluginsSpec struct {
	// FluxNinja ARC plugin configuration.
	//+kubebuilder:validation:Optional
	FluxNinjaPlugin pluginconfig.FluxNinjaPluginConfig `json:"fluxninja_plugin"`

	// Sentry plugin configuration.
	//+kubebuilder:validation:Optional
	SentryPlugin sentry.SentryConfig `json:"sentry_plugin"`
}

// ServiceDiscoverySpec defines configuration for Service discoveru.
type ServiceDiscoverySpec struct {
	// KubernetesDiscoveryConfig for Kubernetes service discovery.
	KubernetesDiscoveryConfig kubernetes.KubernetesDiscoveryConfig `json:"kubernetes"`

	// StaticDiscoveryConfig for pre-determined list of services.
	StaticDiscoveryConfig static.StaticDiscoveryConfig `json:"static"`
}
