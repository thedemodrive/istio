// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operator/v1alpha1/component.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	v1 "k8s.io/api/core/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// IstioComponentSpec defines the desired installed state of Istio components.
type IstioComponentSetSpec struct {
	Base                 *BaseComponentSpec `protobuf:"bytes,29,opt,name=base,proto3" json:"base,omitempty"`
	Pilot                *ComponentSpec     `protobuf:"bytes,30,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Policy               *ComponentSpec     `protobuf:"bytes,33,opt,name=policy,proto3" json:"policy,omitempty"`
	Telemetry            *ComponentSpec     `protobuf:"bytes,34,opt,name=telemetry,proto3" json:"telemetry,omitempty"`
	Cni                  *ComponentSpec     `protobuf:"bytes,38,opt,name=cni,proto3" json:"cni,omitempty"`
	IstiodRemote         *ComponentSpec     `protobuf:"bytes,39,opt,name=istiod_remote,json=istiodRemote,proto3" json:"istiodRemote,omitempty"`
	IngressGateways      []*GatewaySpec     `protobuf:"bytes,40,rep,name=ingress_gateways,json=ingressGateways,proto3" json:"ingressGateways,omitempty"`
	EgressGateways       []*GatewaySpec     `protobuf:"bytes,41,rep,name=egress_gateways,json=egressGateways,proto3" json:"egressGateways,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IstioComponentSetSpec) Reset()         { *m = IstioComponentSetSpec{} }
func (m *IstioComponentSetSpec) String() string { return proto.CompactTextString(m) }
func (*IstioComponentSetSpec) ProtoMessage()    {}
func (*IstioComponentSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{0}
}

func (m *IstioComponentSetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioComponentSetSpec.Unmarshal(m, b)
}
func (m *IstioComponentSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioComponentSetSpec.Marshal(b, m, deterministic)
}
func (m *IstioComponentSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioComponentSetSpec.Merge(m, src)
}
func (m *IstioComponentSetSpec) XXX_Size() int {
	return xxx_messageInfo_IstioComponentSetSpec.Size(m)
}
func (m *IstioComponentSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioComponentSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioComponentSetSpec proto.InternalMessageInfo

func (m *IstioComponentSetSpec) GetBase() *BaseComponentSpec {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *IstioComponentSetSpec) GetPilot() *ComponentSpec {
	if m != nil {
		return m.Pilot
	}
	return nil
}

func (m *IstioComponentSetSpec) GetPolicy() *ComponentSpec {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *IstioComponentSetSpec) GetTelemetry() *ComponentSpec {
	if m != nil {
		return m.Telemetry
	}
	return nil
}

func (m *IstioComponentSetSpec) GetCni() *ComponentSpec {
	if m != nil {
		return m.Cni
	}
	return nil
}

func (m *IstioComponentSetSpec) GetIstiodRemote() *ComponentSpec {
	if m != nil {
		return m.IstiodRemote
	}
	return nil
}

func (m *IstioComponentSetSpec) GetIngressGateways() []*GatewaySpec {
	if m != nil {
		return m.IngressGateways
	}
	return nil
}

func (m *IstioComponentSetSpec) GetEgressGateways() []*GatewaySpec {
	if m != nil {
		return m.EgressGateways
	}
	return nil
}

// Configuration for base component.
type BaseComponentSpec struct {
	// Selects whether this component is installed.
	Enabled              *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BaseComponentSpec) Reset()         { *m = BaseComponentSpec{} }
func (m *BaseComponentSpec) String() string { return proto.CompactTextString(m) }
func (*BaseComponentSpec) ProtoMessage()    {}
func (*BaseComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{1}
}

func (m *BaseComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseComponentSpec.Unmarshal(m, b)
}
func (m *BaseComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseComponentSpec.Marshal(b, m, deterministic)
}
func (m *BaseComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseComponentSpec.Merge(m, src)
}
func (m *BaseComponentSpec) XXX_Size() int {
	return xxx_messageInfo_BaseComponentSpec.Size(m)
}
func (m *BaseComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BaseComponentSpec proto.InternalMessageInfo


// Configuration for internal components.
type ComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the component.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Hub for the component (overrides top level hub setting).
	Hub string `protobuf:"bytes,10,opt,name=hub,proto3" json:"hub,omitempty"`
	// Tag for the component (overrides top level tag setting).
	Tag interface{} `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag,omitempty"`
	// Arbitrary install time configuration for the component.
	Spec interface{} `protobuf:"bytes,30,opt,name=spec,proto3" json:"spec,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ComponentSpec) Reset()         { *m = ComponentSpec{} }
func (m *ComponentSpec) String() string { return proto.CompactTextString(m) }
func (*ComponentSpec) ProtoMessage()    {}
func (*ComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{2}
}

func (m *ComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentSpec.Unmarshal(m, b)
}
func (m *ComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentSpec.Marshal(b, m, deterministic)
}
func (m *ComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentSpec.Merge(m, src)
}
func (m *ComponentSpec) XXX_Size() int {
	return xxx_messageInfo_ComponentSpec.Size(m)
}
func (m *ComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentSpec proto.InternalMessageInfo


func (m *ComponentSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ComponentSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}



func (m *ComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for external components.
type ExternalComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the component.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Arbitrary install time configuration for the component.
	Spec interface{} `protobuf:"bytes,10,opt,name=spec,proto3" json:"spec,omitempty"`
	// Chart path for addon components.
	ChartPath string `protobuf:"bytes,30,opt,name=chart_path,json=chartPath,proto3" json:"chartPath,omitempty"`
	// Optional schema to validate spec against.
	Schema *any.Any `protobuf:"bytes,35,opt,name=schema,proto3" json:"schema,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExternalComponentSpec) Reset()         { *m = ExternalComponentSpec{} }
func (m *ExternalComponentSpec) String() string { return proto.CompactTextString(m) }
func (*ExternalComponentSpec) ProtoMessage()    {}
func (*ExternalComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{3}
}

func (m *ExternalComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalComponentSpec.Unmarshal(m, b)
}
func (m *ExternalComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalComponentSpec.Marshal(b, m, deterministic)
}
func (m *ExternalComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalComponentSpec.Merge(m, src)
}
func (m *ExternalComponentSpec) XXX_Size() int {
	return xxx_messageInfo_ExternalComponentSpec.Size(m)
}
func (m *ExternalComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalComponentSpec proto.InternalMessageInfo


func (m *ExternalComponentSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}


func (m *ExternalComponentSpec) GetChartPath() string {
	if m != nil {
		return m.ChartPath
	}
	return ""
}

func (m *ExternalComponentSpec) GetSchema() *any.Any {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *ExternalComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for gateways.
type GatewaySpec struct {
	// Selects whether this gateway is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the gateway.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name for the gateway.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Labels for the gateway.
	Label map[string]string `protobuf:"bytes,4,rep,name=label,proto3" json:"label,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Hub for the component (overrides top level hub setting).
	Hub string `protobuf:"bytes,10,opt,name=hub,proto3" json:"hub,omitempty"`
	// Tag for the component (overrides top level tag setting).
	Tag interface{} `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GatewaySpec) Reset()         { *m = GatewaySpec{} }
func (m *GatewaySpec) String() string { return proto.CompactTextString(m) }
func (*GatewaySpec) ProtoMessage()    {}
func (*GatewaySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{4}
}

func (m *GatewaySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewaySpec.Unmarshal(m, b)
}
func (m *GatewaySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewaySpec.Marshal(b, m, deterministic)
}
func (m *GatewaySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewaySpec.Merge(m, src)
}
func (m *GatewaySpec) XXX_Size() int {
	return xxx_messageInfo_GatewaySpec.Size(m)
}
func (m *GatewaySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewaySpec.DiscardUnknown(m)
}

var xxx_messageInfo_GatewaySpec proto.InternalMessageInfo


func (m *GatewaySpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GatewaySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewaySpec) GetLabel() map[string]string {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *GatewaySpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}


func (m *GatewaySpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// KubernetesResourcesConfig is a common set of k8s resource configs for components.
type KubernetesResourcesSpec struct {
	// k8s affinity.
	// [https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity)
	Affinity *Affinity `protobuf:"bytes,1,opt,name=affinity,proto3" json:"affinity,omitempty"`
	// Deployment environment variables.
	// [https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/)
	Env []*EnvVar `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// k8s HorizontalPodAutoscaler settings.
	// [https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
	HpaSpec *HorizontalPodAutoscalerSpec `protobuf:"bytes,3,opt,name=hpa_spec,json=hpaSpec,proto3" json:"hpaSpec,omitempty"`
	// k8s imagePullPolicy.
	// [https://kubernetes.io/docs/concepts/containers/images/](https://kubernetes.io/docs/concepts/containers/images/)
	ImagePullPolicy string `protobuf:"bytes,4,opt,name=image_pull_policy,json=imagePullPolicy,proto3" json:"imagePullPolicy,omitempty"`
	// k8s nodeSelector.
	// [https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector)
	NodeSelector map[string]string `protobuf:"bytes,5,rep,name=node_selector,json=nodeSelector,proto3" json:"nodeSelector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s PodDisruptionBudget settings.
	// [https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work)
	PodDisruptionBudget *PodDisruptionBudgetSpec `protobuf:"bytes,6,opt,name=pod_disruption_budget,json=podDisruptionBudget,proto3" json:"podDisruptionBudget,omitempty"`
	// k8s pod annotations.
	// [https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	PodAnnotations map[string]string `protobuf:"bytes,7,rep,name=pod_annotations,json=podAnnotations,proto3" json:"podAnnotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s priority_class_name. Default for all resources unless overridden.
	// [https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass)
	PriorityClassName string `protobuf:"bytes,8,opt,name=priority_class_name,json=priorityClassName,proto3" json:"priorityClassName,omitempty"`
	// k8s readinessProbe settings.
	// [https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)
	// k8s.io.api.core.v1.Probe readiness_probe = 9;
	ReadinessProbe *ReadinessProbe `protobuf:"bytes,9,opt,name=readiness_probe,json=readinessProbe,proto3" json:"readinessProbe,omitempty"`
	// k8s Deployment replicas setting.
	// [https://kubernetes.io/docs/concepts/workloads/controllers/deployment/](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
	ReplicaCount uint32 `protobuf:"varint,10,opt,name=replica_count,json=replicaCount,proto3" json:"replicaCount,omitempty"`
	// k8s resources settings.
	// [https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container)
	Resources *Resources `protobuf:"bytes,11,opt,name=resources,proto3" json:"resources,omitempty"`
	// k8s Service settings.
	// [https://kubernetes.io/docs/concepts/services-networking/service/](https://kubernetes.io/docs/concepts/services-networking/service/)
	Service *ServiceSpec `protobuf:"bytes,12,opt,name=service,proto3" json:"service,omitempty"`
	// k8s deployment strategy.
	// [https://kubernetes.io/docs/concepts/workloads/controllers/deployment/](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
	Strategy *DeploymentStrategy `protobuf:"bytes,13,opt,name=strategy,proto3" json:"strategy,omitempty"`
	// k8s toleration
	// [https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
	Tolerations []*v1.Toleration `protobuf:"bytes,14,rep,name=tolerations,proto3" json:"tolerations,omitempty"`
	// k8s service annotations.
	// [https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	ServiceAnnotations map[string]string `protobuf:"bytes,15,rep,name=service_annotations,json=serviceAnnotations,proto3" json:"serviceAnnotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Overlays for k8s resources in rendered manifests.
	Overlays             []*K8SObjectOverlay `protobuf:"bytes,100,rep,name=overlays,proto3" json:"overlays,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KubernetesResourcesSpec) Reset()         { *m = KubernetesResourcesSpec{} }
func (m *KubernetesResourcesSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesResourcesSpec) ProtoMessage()    {}
func (*KubernetesResourcesSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{5}
}

func (m *KubernetesResourcesSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesResourcesSpec.Unmarshal(m, b)
}
func (m *KubernetesResourcesSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesResourcesSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesResourcesSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesResourcesSpec.Merge(m, src)
}
func (m *KubernetesResourcesSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesResourcesSpec.Size(m)
}
func (m *KubernetesResourcesSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesResourcesSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesResourcesSpec proto.InternalMessageInfo

func (m *KubernetesResourcesSpec) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetEnv() []*EnvVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetHpaSpec() *HorizontalPodAutoscalerSpec {
	if m != nil {
		return m.HpaSpec
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetImagePullPolicy() string {
	if m != nil {
		return m.ImagePullPolicy
	}
	return ""
}

func (m *KubernetesResourcesSpec) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPodDisruptionBudget() *PodDisruptionBudgetSpec {
	if m != nil {
		return m.PodDisruptionBudget
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPodAnnotations() map[string]string {
	if m != nil {
		return m.PodAnnotations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPriorityClassName() string {
	if m != nil {
		return m.PriorityClassName
	}
	return ""
}

func (m *KubernetesResourcesSpec) GetReadinessProbe() *ReadinessProbe {
	if m != nil {
		return m.ReadinessProbe
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetReplicaCount() uint32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *KubernetesResourcesSpec) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetService() *ServiceSpec {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetStrategy() *DeploymentStrategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetTolerations() []*v1.Toleration {
	if m != nil {
		return m.Tolerations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetServiceAnnotations() map[string]string {
	if m != nil {
		return m.ServiceAnnotations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetOverlays() []*K8SObjectOverlay {
	if m != nil {
		return m.Overlays
	}
	return nil
}

// Patch for an existing k8s resource.
type K8SObjectOverlay struct {
	// Resource API version.
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"apiVersion,omitempty"`
	// Resource kind.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of resource.
	// Namespace is always the component namespace.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// List of patches to apply to resource.
	Patches              []*K8SObjectOverlay_PathValue `protobuf:"bytes,4,rep,name=patches,proto3" json:"patches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *K8SObjectOverlay) Reset()         { *m = K8SObjectOverlay{} }
func (m *K8SObjectOverlay) String() string { return proto.CompactTextString(m) }
func (*K8SObjectOverlay) ProtoMessage()    {}
func (*K8SObjectOverlay) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{6}
}

func (m *K8SObjectOverlay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SObjectOverlay.Unmarshal(m, b)
}
func (m *K8SObjectOverlay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SObjectOverlay.Marshal(b, m, deterministic)
}
func (m *K8SObjectOverlay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SObjectOverlay.Merge(m, src)
}
func (m *K8SObjectOverlay) XXX_Size() int {
	return xxx_messageInfo_K8SObjectOverlay.Size(m)
}
func (m *K8SObjectOverlay) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SObjectOverlay.DiscardUnknown(m)
}

var xxx_messageInfo_K8SObjectOverlay proto.InternalMessageInfo

func (m *K8SObjectOverlay) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *K8SObjectOverlay) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *K8SObjectOverlay) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *K8SObjectOverlay) GetPatches() []*K8SObjectOverlay_PathValue {
	if m != nil {
		return m.Patches
	}
	return nil
}

type K8SObjectOverlay_PathValue struct {
	// Path of the form a.[key1:value1].b.[:value2]
	// Where [key1:value1] is a selector for a key-value pair to identify a list element and [:value] is a value
	// selector to identify a list element in a leaf list.
	// All path intermediate nodes must exist.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Value to add, delete or replace.
	// For add, the path should be a new leaf.
	// For delete, value should be unset.
	// For replace, path should reference an existing node.
	// All values are strings but are converted into appropriate type based on schema.
	Value                interface{} `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *K8SObjectOverlay_PathValue) Reset()         { *m = K8SObjectOverlay_PathValue{} }
func (m *K8SObjectOverlay_PathValue) String() string { return proto.CompactTextString(m) }
func (*K8SObjectOverlay_PathValue) ProtoMessage()    {}
func (*K8SObjectOverlay_PathValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{6, 0}
}

func (m *K8SObjectOverlay_PathValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Unmarshal(m, b)
}
func (m *K8SObjectOverlay_PathValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Marshal(b, m, deterministic)
}
func (m *K8SObjectOverlay_PathValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SObjectOverlay_PathValue.Merge(m, src)
}
func (m *K8SObjectOverlay_PathValue) XXX_Size() int {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Size(m)
}
func (m *K8SObjectOverlay_PathValue) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SObjectOverlay_PathValue.DiscardUnknown(m)
}

var xxx_messageInfo_K8SObjectOverlay_PathValue proto.InternalMessageInfo

func (m *K8SObjectOverlay_PathValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}





func init() {
	proto.RegisterType((*IstioComponentSetSpec)(nil), "istio.operator.v1alpha1.IstioComponentSetSpec")
	proto.RegisterType((*BaseComponentSpec)(nil), "istio.operator.v1alpha1.BaseComponentSpec")
	proto.RegisterType((*ComponentSpec)(nil), "istio.operator.v1alpha1.ComponentSpec")
	proto.RegisterType((*ExternalComponentSpec)(nil), "istio.operator.v1alpha1.ExternalComponentSpec")
	proto.RegisterType((*GatewaySpec)(nil), "istio.operator.v1alpha1.GatewaySpec")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.GatewaySpec.LabelEntry")
	proto.RegisterType((*KubernetesResourcesSpec)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.NodeSelectorEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.PodAnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.ServiceAnnotationsEntry")
	proto.RegisterType((*K8SObjectOverlay)(nil), "istio.operator.v1alpha1.K8sObjectOverlay")
	proto.RegisterType((*K8SObjectOverlay_PathValue)(nil), "istio.operator.v1alpha1.K8sObjectOverlay.PathValue")
}

func init() { proto.RegisterFile("operator/v1alpha1/component.proto", fileDescriptor_6ed34a579e9b43a2) }

var fileDescriptor_6ed34a579e9b43a2 = []byte{
	// 1252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x6d, 0x6f, 0x13, 0x47,
	0x10, 0x56, 0x62, 0x27, 0xb1, 0x27, 0x6f, 0xce, 0x06, 0xca, 0x35, 0x2a, 0x90, 0x98, 0xb7, 0x40,
	0xab, 0x73, 0x03, 0xfd, 0x10, 0xa1, 0x96, 0x12, 0x13, 0x17, 0x08, 0x85, 0x58, 0x17, 0xc4, 0x07,
	0xa4, 0xea, 0xb4, 0xbe, 0x9b, 0xd8, 0xdb, 0x9c, 0x77, 0x57, 0xbb, 0x6b, 0x97, 0xeb, 0x1f, 0xa8,
	0xfa, 0x07, 0xda, 0xbf, 0xd3, 0xaf, 0xfd, 0x57, 0xd5, 0xee, 0x9d, 0x6d, 0x8c, 0x71, 0x83, 0x2b,
	0xda, 0x7e, 0xbb, 0x9b, 0x9d, 0xe7, 0x99, 0x97, 0x9b, 0x9d, 0x99, 0x83, 0x1d, 0x21, 0x51, 0x51,
	0x23, 0x54, 0xad, 0xbf, 0x47, 0x13, 0xd9, 0xa1, 0x7b, 0xb5, 0x48, 0x74, 0xa5, 0xe0, 0xc8, 0x8d,
	0x2f, 0x95, 0x30, 0x82, 0x5c, 0x62, 0xda, 0x30, 0xe1, 0x0f, 0x14, 0xfd, 0x81, 0xe2, 0xd6, 0xa7,
	0x6d, 0x21, 0xda, 0x09, 0xd6, 0x9c, 0x5a, 0xab, 0x77, 0x5a, 0xa3, 0x3c, 0xcd, 0x30, 0x5b, 0xd5,
	0xb3, 0x7d, 0xed, 0x33, 0x51, 0xa3, 0x92, 0xd5, 0x22, 0xa1, 0xb0, 0xd6, 0xdf, 0xab, 0xb5, 0x91,
	0x5b, 0x06, 0x8c, 0x07, 0x3a, 0x93, 0xa6, 0xcf, 0x7a, 0x2d, 0x54, 0x1c, 0x0d, 0xea, 0x4c, 0xa7,
	0xfa, 0xdb, 0x02, 0x5c, 0x7c, 0x6a, 0xcd, 0x3f, 0x1a, 0x38, 0x75, 0x82, 0xe6, 0x44, 0x62, 0x44,
	0x1e, 0x40, 0xb1, 0x45, 0x35, 0x7a, 0x97, 0xb7, 0xe7, 0x76, 0x97, 0xef, 0xde, 0xf1, 0xa7, 0x38,
	0xe9, 0xd7, 0xa9, 0xc6, 0x11, 0x58, 0x62, 0x14, 0x38, 0x1c, 0xf9, 0x1a, 0x16, 0x24, 0x4b, 0x84,
	0xf1, 0xae, 0x38, 0x82, 0x9b, 0x53, 0x09, 0xc6, 0xc1, 0x19, 0x88, 0x3c, 0x80, 0x45, 0x29, 0x12,
	0x16, 0xa5, 0xde, 0xce, 0x4c, 0xf0, 0x1c, 0x45, 0x0e, 0xa1, 0x6c, 0x30, 0xc1, 0x2e, 0x1a, 0x95,
	0x7a, 0xd5, 0x99, 0x28, 0x46, 0x40, 0xb2, 0x0f, 0x85, 0x88, 0x33, 0xef, 0xe6, 0x4c, 0x78, 0x0b,
	0x21, 0xcf, 0x60, 0xd5, 0x69, 0xc7, 0xa1, 0xc2, 0xae, 0x30, 0xe8, 0xdd, 0x9a, 0x89, 0x63, 0x25,
	0x03, 0x07, 0x0e, 0x4b, 0x8e, 0xa1, 0xc2, 0x78, 0x5b, 0xa1, 0xd6, 0x61, 0x9b, 0x1a, 0xfc, 0x89,
	0xa6, 0xda, 0xdb, 0xdd, 0x2e, 0xec, 0x2e, 0xdf, 0xbd, 0x3e, 0x95, 0xef, 0x71, 0xa6, 0xe8, 0xd8,
	0xd6, 0x73, 0x74, 0x2e, 0xd3, 0xe4, 0x39, 0xac, 0xe3, 0x3b, 0x7c, 0xb7, 0x67, 0xe0, 0x5b, 0xc3,
	0x31, 0xba, 0xa3, 0x62, 0xe9, 0x6a, 0x65, 0xfb, 0xa8, 0x58, 0xda, 0xae, 0xec, 0x1c, 0x15, 0x4b,
	0xd7, 0x2a, 0xd7, 0x8f, 0x8a, 0xa5, 0xeb, 0x95, 0x1b, 0x47, 0xc5, 0xd2, 0x8d, 0xca, 0xcd, 0x60,
	0x29, 0x62, 0x86, 0xc6, 0x98, 0x04, 0x8b, 0x6d, 0x9a, 0x24, 0x98, 0x06, 0xc0, 0x45, 0x8c, 0x21,
	0x6d, 0x23, 0x37, 0xc1, 0x82, 0x54, 0xe2, 0x4d, 0x1a, 0x54, 0x34, 0x8b, 0x31, 0xa2, 0x2a, 0x64,
	0xfc, 0x47, 0x8c, 0x8c, 0x50, 0xd5, 0xd7, 0xb0, 0x31, 0x51, 0x59, 0xa4, 0x01, 0x4b, 0xc8, 0x69,
	0x2b, 0xc1, 0xd8, 0x9b, 0x73, 0xf9, 0xfc, 0x7c, 0xaa, 0xbf, 0x2f, 0x53, 0x89, 0x75, 0x21, 0x92,
	0x57, 0x34, 0xe9, 0xe1, 0x77, 0x42, 0x35, 0xeb, 0xc1, 0x00, 0x5b, 0xfd, 0x63, 0x1e, 0x56, 0xff,
	0x0d, 0x62, 0xf2, 0x19, 0x94, 0x39, 0xed, 0xa2, 0x96, 0x34, 0x42, 0x6f, 0x7e, 0x7b, 0x6e, 0xb7,
	0x1c, 0x8c, 0x04, 0xa4, 0x02, 0x85, 0x4e, 0xaf, 0xe5, 0x81, 0x93, 0xdb, 0x47, 0x5b, 0x5f, 0x86,
	0xb6, 0xbd, 0xe5, 0x73, 0x6a, 0xc3, 0x9a, 0x7c, 0xca, 0x0d, 0xaa, 0x53, 0x1a, 0x61, 0x60, 0x21,
	0xe4, 0x3e, 0x14, 0xb5, 0xc4, 0xe8, 0xdc, 0xcb, 0x35, 0x0e, 0x75, 0x18, 0x52, 0x87, 0xc2, 0xd9,
	0xbe, 0xf6, 0xee, 0x3a, 0xe8, 0x97, 0x53, 0xa1, 0xcf, 0x86, 0xbd, 0x22, 0x40, 0x2d, 0x7a, 0x2a,
	0x42, 0x9d, 0xd5, 0xf7, 0xd9, 0xbe, 0xae, 0xfe, 0x39, 0x0f, 0x17, 0x1b, 0x6f, 0x0c, 0x2a, 0x4e,
	0x93, 0xff, 0x21, 0x95, 0x83, 0xf0, 0xe1, 0x1f, 0x84, 0x7f, 0x19, 0x20, 0xea, 0x50, 0x65, 0x42,
	0x49, 0x4d, 0xc7, 0x25, 0xb0, 0x1c, 0x94, 0x9d, 0xa4, 0x49, 0x4d, 0x87, 0x7c, 0x01, 0x8b, 0x3a,
	0xea, 0x60, 0x97, 0x7a, 0xd7, 0x1c, 0xf9, 0x05, 0x3f, 0xeb, 0xc2, 0xfe, 0xa0, 0x0b, 0xfb, 0x07,
	0x3c, 0x0d, 0x72, 0x9d, 0x8f, 0x92, 0xcb, 0xdf, 0x0b, 0xb0, 0xfc, 0xd6, 0xf5, 0xfa, 0x6f, 0x32,
	0x48, 0xa0, 0x68, 0x5f, 0xbc, 0x82, 0x3b, 0x70, 0xcf, 0xa4, 0x01, 0x0b, 0x09, 0x6d, 0x61, 0xe2,
	0x15, 0x5d, 0x33, 0xa8, 0x7d, 0x48, 0x33, 0xf0, 0xbf, 0xb7, 0x88, 0x06, 0x37, 0x2a, 0x0d, 0x32,
	0xf4, 0x47, 0xad, 0xf3, 0x8f, 0x90, 0xdf, 0xad, 0x7d, 0x80, 0x91, 0x93, 0xd6, 0xbb, 0x33, 0x4c,
	0x5d, 0x66, 0xcb, 0x81, 0x7d, 0x24, 0x17, 0x60, 0xa1, 0x6f, 0xf3, 0x97, 0x27, 0x29, 0x7b, 0xb9,
	0x3f, 0xbf, 0x3f, 0x57, 0xfd, 0x65, 0x19, 0x2e, 0x4d, 0xa1, 0x26, 0xdf, 0x40, 0x89, 0x9e, 0x9e,
	0x32, 0xce, 0x4c, 0x9a, 0x7f, 0xa6, 0x9d, 0xa9, 0xee, 0x1d, 0xe4, 0x8a, 0xc1, 0x10, 0x42, 0xf6,
	0xa0, 0x80, 0xbc, 0xef, 0xcd, 0xbb, 0x4c, 0x5f, 0x9d, 0x8a, 0x6c, 0xf0, 0xfe, 0x2b, 0xaa, 0x02,
	0xab, 0x4b, 0x8e, 0xa1, 0xd4, 0x91, 0x34, 0x74, 0x85, 0x5f, 0x70, 0x16, 0xbf, 0x9a, 0x8a, 0x7b,
	0x22, 0x14, 0xfb, 0x59, 0x70, 0x43, 0x93, 0xa6, 0x88, 0x0f, 0x7a, 0x46, 0xe8, 0x88, 0x26, 0xa8,
	0x5c, 0x52, 0x96, 0x3a, 0x92, 0xba, 0x10, 0xee, 0xc0, 0x06, 0xeb, 0xd2, 0x36, 0x86, 0xb2, 0x97,
	0x24, 0x61, 0x3e, 0x6f, 0x8b, 0x2e, 0x09, 0xeb, 0xee, 0xa0, 0xd9, 0x4b, 0x92, 0x66, 0x36, 0x50,
	0xdb, 0xb0, 0xea, 0xda, 0xb6, 0xc6, 0xc4, 0x35, 0x68, 0x6f, 0xc1, 0x79, 0x5e, 0x9f, 0xf5, 0x93,
	0xf8, 0x2f, 0x44, 0x8c, 0x27, 0x39, 0x49, 0x56, 0x36, 0x2b, 0xfc, 0x2d, 0x11, 0x89, 0xe1, 0xa2,
	0x14, 0x71, 0x18, 0x33, 0xad, 0x7a, 0xd2, 0x30, 0xc1, 0xc3, 0x56, 0x2f, 0x6e, 0xa3, 0xf1, 0x16,
	0xcf, 0xa9, 0x81, 0xa6, 0x88, 0x0f, 0x87, 0xa0, 0xba, 0xc3, 0xb8, 0x70, 0x37, 0xe5, 0xe4, 0x01,
	0xe9, 0xc2, 0xba, 0xb5, 0x42, 0x39, 0x17, 0x86, 0x5a, 0xb9, 0xf6, 0x96, 0x5c, 0x40, 0x87, 0x33,
	0x07, 0x64, 0x13, 0x3c, 0xa2, 0xc9, 0x42, 0x5a, 0x93, 0x63, 0x42, 0xe2, 0xc3, 0xa6, 0x54, 0x4c,
	0x28, 0x66, 0xd2, 0x30, 0x4a, 0xa8, 0xd6, 0xa1, 0xbb, 0x7c, 0x25, 0x97, 0xeb, 0x8d, 0xc1, 0xd1,
	0x23, 0x7b, 0xf2, 0xc2, 0xde, 0xc4, 0x26, 0xac, 0x2b, 0xa4, 0x31, 0xe3, 0x76, 0x46, 0x4b, 0x25,
	0x5a, 0xe8, 0x95, 0x5d, 0xf8, 0xb7, 0xa6, 0xba, 0x17, 0x0c, 0xf4, 0x9b, 0x56, 0x3d, 0x58, 0x53,
	0x63, 0xef, 0xe4, 0x1a, 0xac, 0x2a, 0x94, 0x09, 0x8b, 0x68, 0x18, 0x89, 0x1e, 0x37, 0xee, 0x7a,
	0xae, 0x06, 0x2b, 0xb9, 0xf0, 0x91, 0x95, 0x91, 0x87, 0x50, 0x56, 0x83, 0xd8, 0xf2, 0xdb, 0x5a,
	0xfd, 0x1b, 0x83, 0xb9, 0x66, 0x30, 0x02, 0x91, 0x07, 0xb0, 0xa4, 0x51, 0xf5, 0x59, 0x84, 0xde,
	0x8a, 0xc3, 0x4f, 0xdf, 0x28, 0x4e, 0x32, 0xbd, 0xac, 0x24, 0x73, 0x10, 0x79, 0x0c, 0x25, 0x6d,
	0xec, 0x12, 0xdb, 0x4e, 0xbd, 0xd5, 0x73, 0x9a, 0xdf, 0x21, 0xca, 0x44, 0xa4, 0x5d, 0x3b, 0x78,
	0x72, 0x48, 0x30, 0x04, 0x93, 0x87, 0xb0, 0x6c, 0x44, 0x62, 0x21, 0xee, 0xe3, 0xae, 0xb9, 0x8f,
	0x7b, 0xc5, 0xcf, 0xd6, 0x66, 0x9f, 0x4a, 0xe6, 0xdb, 0xb5, 0xd9, 0xef, 0xef, 0xf9, 0x2f, 0x87,
	0x6a, 0xc1, 0xdb, 0x10, 0x92, 0xc2, 0x66, 0xee, 0xd5, 0x58, 0x99, 0xac, 0x3b, 0xa6, 0x27, 0x33,
	0x97, 0x49, 0x1e, 0xee, 0x44, 0xa9, 0x10, 0x3d, 0x71, 0x40, 0x1a, 0x50, 0x12, 0x7d, 0x54, 0x89,
	0x5d, 0xcc, 0x62, 0x67, 0xef, 0xf6, 0x74, 0x7b, 0xfb, 0xfa, 0xb8, 0x65, 0x97, 0xa6, 0xe3, 0x0c,
	0x11, 0x0c, 0xa1, 0x5b, 0xdf, 0xc2, 0xc6, 0xc4, 0x6d, 0x9b, 0xa5, 0xff, 0x6d, 0x1d, 0xc0, 0xe6,
	0x7b, 0xaa, 0x7b, 0x26, 0x8a, 0x06, 0x5c, 0x9a, 0x12, 0xf9, 0x4c, 0x9d, 0xf8, 0xd7, 0x79, 0xa8,
	0xbc, 0x1b, 0x29, 0xb9, 0x0a, 0xcb, 0x54, 0xb2, 0xb0, 0x8f, 0x4a, 0x33, 0xc1, 0x73, 0x22, 0xa0,
	0x92, 0xbd, 0xca, 0x24, 0x76, 0xc8, 0x9d, 0x31, 0x1e, 0xe7, 0x74, 0xee, 0xf9, 0xbd, 0x83, 0xef,
	0x39, 0x2c, 0x49, 0x6a, 0xa2, 0x0e, 0xea, 0x7c, 0xf4, 0xdd, 0xfb, 0xe0, 0x74, 0xfb, 0x76, 0x69,
	0x70, 0xe3, 0x37, 0x18, 0x70, 0x6c, 0xfd, 0x00, 0xe5, 0xa1, 0xd4, 0xda, 0x73, 0x8b, 0x46, 0xe6,
	0x9d, 0x7b, 0xb6, 0xff, 0x46, 0xa3, 0x38, 0x3f, 0x7c, 0x22, 0x66, 0xa0, 0xaa, 0x07, 0x9f, 0x58,
	0xf9, 0x73, 0x2a, 0x4f, 0x8c, 0x62, 0xbc, 0x3d, 0x54, 0xa8, 0xae, 0xc3, 0xea, 0x18, 0xa2, 0x7a,
	0x01, 0xc8, 0xe4, 0x8a, 0x50, 0xdf, 0x7e, 0x7d, 0x25, 0x33, 0x98, 0xff, 0x40, 0x4e, 0xfc, 0x27,
	0xb6, 0x16, 0xdd, 0xb2, 0x73, 0xef, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x21, 0x4f, 0x0d,
	0xbe, 0x0e, 0x00, 0x00,
}