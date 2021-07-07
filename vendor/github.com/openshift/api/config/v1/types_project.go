package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Project holds cluster-wide information about Project.  The canonical name is `cluster`
type Project struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 03397665 (update api)
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +kubebuilder:validation:Required
<<<<<<< HEAD
=======
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
>>>>>>> 79bfea2d (update vendor)
=======
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +kubebuilder:validation:Required
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
	// +required
	Spec ProjectSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status ProjectStatus `json:"status"`
}

// TemplateReference references a template in a specific namespace.
// The namespace must be specified at the point of use.
type TemplateReference struct {
	// name is the metadata.name of the referenced project request template
	Name string `json:"name"`
}

// ProjectSpec holds the project creation configuration.
type ProjectSpec struct {
	// projectRequestMessage is the string presented to a user if they are unable to request a project via the projectrequest api endpoint
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// +optional
=======
>>>>>>> 79bfea2d (update vendor)
=======
	// +optional
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
	// +optional
>>>>>>> 03397665 (update api)
	ProjectRequestMessage string `json:"projectRequestMessage"`

	// projectRequestTemplate is the template to use for creating projects in response to projectrequest.
	// This must point to a template in 'openshift-config' namespace. It is optional.
	// If it is not specified, a default template is used.
	//
	// +optional
	ProjectRequestTemplate TemplateReference `json:"projectRequestTemplate"`
}

type ProjectStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	metav1.ListMeta `json:"metadata"`

	Items []Project `json:"items"`
=======
	// Standard object's metadata.
	metav1.ListMeta `json:"metadata"`
	Items           []Project `json:"items"`
>>>>>>> 79bfea2d (update vendor)
=======
	metav1.ListMeta `json:"metadata"`

	Items []Project `json:"items"`
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
	metav1.ListMeta `json:"metadata"`

	Items []Project `json:"items"`
>>>>>>> 03397665 (update api)
}
