package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Redundancy is a specification for a Redundancy resource
type Redundancy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedundancySpec   `json:"spec"`
	Status RedundancyStatus `json:"status"`
}

// RedundancySpec is the spec for a Redundancy resource
type RedundancySpec struct {
	Replicas   *int32                 `json:"replicas"`
	Redundancy *int32                 `json:"redundancy"`
	Template   corev1.PodTemplateSpec `json:"template"`
}

// RedundancyStatus is the status for a Redundancy resource
type RedundancyStatus struct {
	Replicas       int32 `json:"replicas"`
	ReadyReplicas  int32 `json:"readyReplicas"`
	BackupReplicas int32 `json:"backupReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedundancyList is a list of Redundancy resources
type RedundancyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Redundancy `json:"items"`
}
