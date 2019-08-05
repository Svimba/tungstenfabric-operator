package v1alpha1

import (
	analyticsv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/analytics/v1alpha1"
	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	controlv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/control/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

// TFOperatorSpec defines the desired state of TFOperator
type TFOperatorSpec struct {
	TFConfig    *configv1alpha1.TFConfigSpec       `json:"tf-config"`
	TFControl   *controlv1alpha1.TFControlSpec     `json:"tf-control"`
	TFAnalytics *analyticsv1alpha1.TFAnalyticsSpec `json:"tf-analytics"`
}

// TFOperatorStatus defines the observed state of TFOperator
type TFOperatorStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFOperator is the Schema for the tfoperators API
// +k8s:openapi-gen=true
type TFOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TFOperatorSpec   `json:"spec,omitempty"`
	Status TFOperatorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFOperatorList contains a list of TFOperator
type TFOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TFOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TFOperator{}, &TFOperatorList{})
}
