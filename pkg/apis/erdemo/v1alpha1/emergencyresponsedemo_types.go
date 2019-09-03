package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type EmergencyResponseDemoConditionType string

var (
	EmergencyResponseDemoInit    EmergencyResponseDemoConditionType = "Init"
	EmergencyResponseDemoInstall EmergencyResponseDemoConditionType = "Install"
	EmergencyResponseDemoReady   EmergencyResponseDemoConditionType = "Ready"
	EmergencyResponseDemoError   EmergencyResponseDemoConditionType = "Error"
)

// EmergencyResponseDemoSpec defines the desired state of EmergencyResponseDemo
// +k8s:openapi-gen=true
type EmergencyResponseDemoSpec struct {
	SecretName string `json:"secretName,omitempty" description:"Secret name to store erd components"`
	SelfSignedCerts bool `json:"selfSignedCerts" description:"boolean value to indicate the use of self signed cert"`
	SubDomain string `json:"subDomain" description:"cluster app sub domain host"`
	MasterUrl string `json:"masterUrl" description:"cluster master url"`
}

// EmergencyResponseDemoStatus defines the observed state of EmergencyResponseDemo
// +k8s:openapi-gen=true
type EmergencyResponseDemoStatus struct {
	Type EmergencyResponseDemoConditionType `json:"type"`
	Status v1.ConditionStatus `json:"status"`
	// +optional
	Reason  *string `json:"reason,omitempty" description:"one-word CamelCase reason for the condition's last transition"`
	// +optional
	Message *string `json:"message,omitempty" description:"human-readable message indicating details about last transition"`
	// +optional
	LastHeartbeatTime  *metav1.Time `json:"lastHeartbeatTime,omitempty" description:"last time we got an update on a given condition"`
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" description:"last time the condition transit from one status to another"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EmergencyResponseDemo is the Schema for the emergencyresponsedemos API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type EmergencyResponseDemo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmergencyResponseDemoSpec   `json:"spec,omitempty"`
	Status EmergencyResponseDemoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EmergencyResponseDemoList contains a list of EmergencyResponseDemo
type EmergencyResponseDemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmergencyResponseDemo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EmergencyResponseDemo{}, &EmergencyResponseDemoList{})
}
