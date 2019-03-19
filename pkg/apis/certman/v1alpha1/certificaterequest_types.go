package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CertificateRequestSpec defines the desired state of CertificateRequest
// +k8s:openapi-gen=true
type CertificateRequestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// BaseDomain is the base domain to which the cluster should belong.
	BaseDomain string `json:"baseDomain,omitempty"`

	// CertificateSecret is the reference to the secret where certificates are stored.
	CertificateSecret corev1.LocalObjectReference `json:"certificateSecret,omitempty"`

	// AwsSecrets refers to a secret that contains the AWS account access credentials.
	AwsSecrets corev1.LocalObjectReference `json:"awsSecrets,omitempty"`

	// Request wildcard certificates.
	// +optional
	Wildcard bool `json:"wildcard,omitempty"`

	// Certificate renew before expiration duration in days.
	RenewBeforeDays int `json:"renewBeforeDays,omitempty"`

	// DNSNames is a list of subject alt names to be used on the Certificate.
	// +optional
	DnsNames []string `json:"dnsNames,omitempty"`
}

// CertificateRequestStatus defines the observed state of CertificateRequest
// +k8s:openapi-gen=true
type CertificateRequestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateRequest is the Schema for the certificaterequests API
// +k8s:openapi-gen=true
type CertificateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateRequestSpec   `json:"spec,omitempty"`
	Status CertificateRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateRequestList contains a list of CertificateRequest
type CertificateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateRequest{}, &CertificateRequestList{})
}
