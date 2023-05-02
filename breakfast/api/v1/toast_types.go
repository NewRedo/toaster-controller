/*
Copyright 2023.

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

// ToastSpec defines the desired state of Toast
type ToastSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Toastiness is how long to cook the toast, from 0 (raw) to 10 (blackend)
	Toastiness int `json:"toastiness,omitempty"`
}

// ToastStatus defines the observed state of Toast
type ToastStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Toastiness is the current level of toastiness
	Toastiness int `json:"toastiness,omitempty"`
	// Ready is whether the Toast has been fully cooked
	Ready bool `json:"ready,omitempty"`
	// LastUpdated is the last time the Toast was updated
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Toast is the Schema for the toasts API
type Toast struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToastSpec   `json:"spec,omitempty"`
	Status ToastStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ToastList contains a list of Toast
type ToastList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Toast `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Toast{}, &ToastList{})
}
