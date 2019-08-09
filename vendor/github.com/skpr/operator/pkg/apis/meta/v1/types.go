package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Phase which indicates the status of an operation.
type Phase string

const (
	// PhaseFailed to be assigned when an operation fails.
	PhaseFailed Phase = "Failed"
	// PhaseReady to be assigned when an operation is ready to be progressed.
	PhaseReady Phase = "Ready"
	// PhaseInProgress to be assigned when an operation is in progress.
	PhaseInProgress Phase = "InProgress"
	// PhaseCompleted to be assigned when an operation has been completed.
	PhaseCompleted Phase = "Completed"
	// PhaseUnknown to be assigned when the above phases cannot be determined.
	PhaseUnknown Phase = "Unknown"
)

// +genclient
// +k8s:deepcopy-gen=true

// ScheduledStatus defines the observed state of a scheduled object.
type ScheduledStatus struct {
	// Last time an object object was created.
	LastCreated *metav1.Time `json:"lastCreated,omitempty"`
}
