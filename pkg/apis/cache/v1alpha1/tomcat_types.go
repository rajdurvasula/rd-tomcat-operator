package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TomcatSpec defines the desired state of Tomcat
// +k8s:openapi-gen=true
type TomcatSpec struct {
	// count is the number of servers in deployment
	Count int32 `json:"size"`
}

// TomcatStatus defines the observed state of Tomcat server
// +k8s:openapi-gen=true
type TomcatStatus struct {
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tomcat is the Scheme for the tomcats API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tomcats,scope=Namespaced
type Tomcat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TomcatSpec   `json:"spec,omitempty"`
	Status TomcatStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TomcatList contains a list of Tomcat
type TomcatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tomcat `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tomcat{}, &TomcatList{})
}
