/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WholepicSpec defines the desired state of Wholepic

type WholepicSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Wholepic. Edit Wholepic_types.go to remove/update
	//Foo string `json:"foo,omitempty"`

}

// WholepicStatus defines the observed state of Wholepic
type WholepicStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	DeploymentCount int64 `json:"deployment_count,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// Wholepic is the Schema for the wholepics API
type Wholepic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WholepicSpec   `json:"spec,omitempty"`
	Status WholepicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WholepicList contains a list of Wholepic
type WholepicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wholepic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Wholepic{}, &WholepicList{})
}
