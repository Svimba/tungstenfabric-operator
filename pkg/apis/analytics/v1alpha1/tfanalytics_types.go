package v1alpha1

import (
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

// TFAnalyticsAlarmGenSpec defines the desired state of AlarmGen
type TFAnalyticsAlarmGenSpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFAnalyticsAPISpec defines the desired state of API
type TFAnalyticsAPISpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFAnalyticsCollectorSpec defines the desired state of Collector
type TFAnalyticsCollectorSpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFAnalyticsQueryEngineSpec defines the desired state of Query Engine
type TFAnalyticsQueryEngineSpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFAnalyticsSNMPSpec defines the desired state of SNMP
type TFAnalyticsSNMPSpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFAnalyticsTopologySpec defines the desired state of Topology
type TFAnalyticsTopologySpec struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Replicas *int32   `json:"replicas"`
	Image    string   `json:"image"`
	Ports    []Port   `json:"ports,omitempty"`
	EnvList  []EnvVar `json:"env"`
}

// TFAnalyticsSpec defines the desired state of TFAnalytics
// +k8s:openapi-gen=true
type TFAnalyticsSpec struct {
	AlarmGenSpec  TFAnalyticsAlarmGenSpec    `json:"alarm-gen"`
	APISpec       TFAnalyticsAPISpec         `json:"api"`
	CollectorSpec TFAnalyticsCollectorSpec   `json:"collector"`
	QueryEngine   TFAnalyticsQueryEngineSpec `json:"query-engine"`
	SNMPSpec      TFAnalyticsSNMPSpec        `json:"snmp"`
	TopologySpec  TFAnalyticsTopologySpec    `json:"topology"`
	ConfigMapList []string                   `json:"configmaps,omitempty"`
}

// TFAnalyticsStatus defines the observed state of TFAnalytics
// +k8s:openapi-gen=true
type TFAnalyticsStatus struct {
	ConfigMapList []string `json:"config-map-list,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFAnalytics is the Schema for the tfanalytics API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type TFAnalytics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TFAnalyticsSpec   `json:"spec,omitempty"`
	Status TFAnalyticsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFAnalyticsList contains a list of TFAnalytics
type TFAnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TFAnalytics `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TFAnalytics{}, &TFAnalyticsList{})
}
