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

package agent

import (
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"

	"github.com/fluxninja/aperture/cmd/aperture-agent/agent"
	agentv1alpha1 "github.com/fluxninja/aperture/operator/api/agent/v1alpha1"
	"github.com/fluxninja/aperture/operator/api/common"
	. "github.com/fluxninja/aperture/operator/controllers"
	"github.com/fluxninja/aperture/pkg/distcache"
	"github.com/fluxninja/aperture/pkg/net/listener"
)

var _ = Describe("Service for Agent", func() {
	Context("Instance with default parameters", func() {
		It("returns correct Service", func() {
			instance := &agentv1alpha1.Agent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      AppName,
					Namespace: AppName,
				},
				Spec: agentv1alpha1.AgentSpec{
					ConfigSpec: agentv1alpha1.AgentConfigSpec{
						CommonConfigSpec: common.CommonConfigSpec{
							Server: common.ServerConfigSpec{
								ListenerConfig: listener.ListenerConfig{
									Addr: ":8080",
								},
							},
						},
						DistCache: distcache.DistCacheConfig{
							BindAddr:           ":3320",
							MemberlistBindAddr: ":3322",
						},
						OTEL: agent.AgentOTELConfig{},
					},
				},
			}

			expected := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      AgentServiceName,
					Namespace: AppName,
					Labels: map[string]string{
						"app.kubernetes.io/name":       AppName,
						"app.kubernetes.io/instance":   AppName,
						"app.kubernetes.io/managed-by": OperatorName,
						"app.kubernetes.io/component":  AgentServiceName,
					},
					Annotations: nil,
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "fluxninja.com/v1alpha1",
							Name:               instance.GetName(),
							Kind:               "Agent",
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
					},
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{
							Name:       Server,
							Protocol:   corev1.Protocol(TCP),
							Port:       int32(8080),
							TargetPort: intstr.FromString(Server),
						},
						{
							Name:       DistCache,
							Protocol:   corev1.Protocol(TCP),
							Port:       int32(3320),
							TargetPort: intstr.FromString(DistCache),
						},
						{
							Name:       MemberList,
							Protocol:   corev1.Protocol(TCP),
							Port:       int32(3322),
							TargetPort: intstr.FromString(MemberList),
						},
					},
					InternalTrafficPolicy: &[]corev1.ServiceInternalTrafficPolicyType{corev1.ServiceInternalTrafficPolicyLocal}[0],
					Selector: map[string]string{
						"app.kubernetes.io/name":       AppName,
						"app.kubernetes.io/instance":   AppName,
						"app.kubernetes.io/managed-by": OperatorName,
						"app.kubernetes.io/component":  AgentServiceName,
					},
				},
			}

			result, err := serviceForAgent(instance.DeepCopy(), logr.Logger{}, scheme.Scheme)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(expected))
		})
	})

	Context("Instance with all parameters", func() {
		It("returns correct Service", func() {
			instance := &agentv1alpha1.Agent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      AppName,
					Namespace: AppName,
				},
				Spec: agentv1alpha1.AgentSpec{
					CommonSpec: common.CommonSpec{
						Labels:      TestMap,
						Annotations: TestMap,
						Service: common.Service{
							Annotations: TestMapTwo,
						},
					},
					ConfigSpec: agentv1alpha1.AgentConfigSpec{
						CommonConfigSpec: common.CommonConfigSpec{
							Server: common.ServerConfigSpec{
								ListenerConfig: listener.ListenerConfig{
									Addr: ":8080",
								},
							},
						},
						DistCache: distcache.DistCacheConfig{
							BindAddr:           ":3320",
							MemberlistBindAddr: ":3322",
						},
						OTEL: agent.AgentOTELConfig{},
					},
				},
			}

			expected := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      AgentServiceName,
					Namespace: AppName,
					Labels: map[string]string{
						"app.kubernetes.io/name":       AppName,
						"app.kubernetes.io/instance":   AppName,
						"app.kubernetes.io/managed-by": OperatorName,
						"app.kubernetes.io/component":  AgentServiceName,
						Test:                           Test,
					},
					Annotations: map[string]string{
						Test:    Test,
						TestTwo: TestTwo,
					},
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "fluxninja.com/v1alpha1",
							Name:               instance.GetName(),
							Kind:               "Agent",
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
					},
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{
							Name:       Server,
							Protocol:   corev1.Protocol(TCP),
							Port:       int32(8080),
							TargetPort: intstr.FromString(Server),
						},
						{
							Name:       DistCache,
							Protocol:   corev1.Protocol(TCP),
							Port:       int32(3320),
							TargetPort: intstr.FromString(DistCache),
						},
						{
							Name:       MemberList,
							Protocol:   corev1.Protocol(TCP),
							Port:       int32(3322),
							TargetPort: intstr.FromString(MemberList),
						},
					},
					InternalTrafficPolicy: &[]corev1.ServiceInternalTrafficPolicyType{corev1.ServiceInternalTrafficPolicyLocal}[0],
					Selector: map[string]string{
						"app.kubernetes.io/name":       AppName,
						"app.kubernetes.io/instance":   AppName,
						"app.kubernetes.io/managed-by": OperatorName,
						"app.kubernetes.io/component":  AgentServiceName,
					},
				},
			}

			result, err := serviceForAgent(instance.DeepCopy(), logr.Logger{}, scheme.Scheme)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(expected))
		})
	})
})
