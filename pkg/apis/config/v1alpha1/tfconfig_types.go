package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Port definition
type Port struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

// EnvVar defines Environment sariable
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TFConfigAPISpec defines the desired state of config api service
type TFConfigAPISpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFConfigSVCMonitorSpec defines the desired state of config service monitor
type TFConfigSVCMonitorSpec struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`
	Ports    []Port `json:"ports,omitempty"`
}

// TFConfigSchemaSpec defines the desired state of config schema
type TFConfigSchemaSpec struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`
}

// TFConfigDeviceMgrSpec defines the desired state of config device manager
type TFConfigDeviceMgrSpec struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`
	Ports    []Port `json:"ports,omitempty"`
}

// TFConfigSpec defines the desired state of TFConfig
type TFConfigSpec struct {
	APISpec        TFConfigAPISpec        `json:"api"`
	SVCMonitorSpec TFConfigSVCMonitorSpec `json:"svc-monitor"`
	SchemaSpec     TFConfigSchemaSpec     `json:"schema"`
	DeviceMgrSpec  TFConfigDeviceMgrSpec  `json:"devicemgr"`
	CofigMapList   []string               `json:"configmaps,omitempty"`
}

// TFConfigStatus defines the observed state of TFConfig
type TFConfigStatus struct {
	ConfigMapList []string `json:"config-map-list,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFConfig is the Schema for the tfconfigs API
// +k8s:openapi-gen=true
type TFConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TFConfigSpec   `json:"spec,omitempty"`
	Status TFConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFConfigList contains a list of TFConfig
type TFConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TFConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TFConfig{}, &TFConfigList{})
}
