package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Port definition
type Port struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

// EnvVar defines Environment variable
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// SecCtx defines Security Context fot Control services deployment
type SecCtx struct {
	Capabilities []corev1.Capability `json:"capabilities"`
}

// TFControlControlSpec defines the desired state of control service from TF Control group
type TFControlControlSpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFControlNamedSpec defines the desired state of named service from TF Control group
type TFControlNamedSpec struct {
	Enabled         bool    `json:"enabled,omitempty"`
	Replicas        *int32  `json:"replicas"`
	Image           string  `json:"image"`
	Ports           []Port  `json:"ports,omitempty"`
	SecurityContext SecCtx  `json:"securityContext,omitempty"`
}

// TFControlDnsSpec defines the desired state of dns service from TF Control group
type TFControlDnsSpec struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`
	Ports    []Port `json:"ports,omitempty"`
}

// TFConfigSpec defines the desired state of TFConfig
type TFControlSpec struct {
	ControlSpec    TFControlControlSpec   `json:"control"`
	NamedSpec      TFControlNamedSpec     `json:"named"`
	DnsSpec        TFControlDnsSpec       `json:"dns"`
	ConfigMapList  []string               `json:"configmaps,omitempty"`
}

// TFControlStatus defines the observed state of TFControl
type TFControlStatus struct {
	ConfigMapList []string `json:"config-map-list,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFControl is the Schema for the tfcontrols API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type TFControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TFControlSpec   `json:"spec,omitempty"`
	Status TFControlStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFControlList contains a list of TFControl
type TFControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TFControl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TFControl{}, &TFControlList{})
}
