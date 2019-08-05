package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureQueue struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status AzureQueueStatus `json:"status,omitempty"`

	Spec AzureQueueSpec `json:"spec,omitempty"`
}

type AzureQueueSpec struct {
	// +optional
	StorageAccountName string `json:"storageAccountName,omitempty"`

	// +optional
	UsePodIdentity string `json:"usePodIdentity,omitempty"`

	// +optional
	ConnectionStringSecretRef AzureConnectionStringSecretRef `json:"connectionStringSecretRef,omitempty"`
}

type AzureConnectionStringSecretRef struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type AzureQueueStatus struct {
	Provisioned bool `json:"provisioned,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureQueueList struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AzureQueue `json:"items"`
}
