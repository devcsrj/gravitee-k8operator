package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GatewayServiceSpec defines the desired state of GatewayService
type GatewayServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Collects 'Service's with label keys and values matching this selector.
	Selector map[string]string `json:"selector"`

	// The path where the HTTP GET request should be made on the Service to retrieve the OAS content
	OasPath string `json:"oasPath"`
}

// GatewayServiceStatus defines the observed state of GatewayService
type GatewayServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayService is the Schema for the gatewayservices API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=gatewayservices,scope=Namespaced
type GatewayService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewayServiceSpec   `json:"spec,omitempty"`
	Status GatewayServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayServiceList contains a list of GatewayService
type GatewayServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GatewayService{}, &GatewayServiceList{})
}
