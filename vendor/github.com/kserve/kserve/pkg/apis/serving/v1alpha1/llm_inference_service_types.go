/*
Copyright 2024 The KServe Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	knativev1 "knative.dev/pkg/apis/duck/v1"
)

// LLMInferenceService is the Schema for the LLMInferenceService API
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=llminferenceservices,shortName=llmisvc,singular=llminferenceservice
type LLMInferenceService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LLMInferenceServiceSpec   `json:"spec,omitempty"`
	Status            LLMInferenceServiceStatus `json:"status,omitempty"`
}

// LLMInferenceServiceList contains a list of LLMInferenceService
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type LLMInferenceServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LLMInferenceService `json:"items"`
}

// LLMInferenceServiceModelRef identifies the model served by this LLMInferenceService.
type LLMInferenceServiceModelRef struct {
	// Name of the model
	// +optional
	Name string `json:"name,omitempty"`
	// URI is the storage location of the model artifacts
	// +optional
	URI string `json:"uri,omitempty"`
}

// LLMInferenceServiceSpec defines the desired state of LLMInferenceService
type LLMInferenceServiceSpec struct {
	// Model reference
	// +optional
	Model LLMInferenceServiceModelRef `json:"model,omitempty"`
	// Replicas is the desired number of replicas
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`
	// Router, Template, and other fields are captured as raw JSON to avoid
	// unmarshal errors from schema fields not used by this controller.
	// +optional
	Router *runtime.RawExtension `json:"router,omitempty"`
	// +optional
	Template *runtime.RawExtension `json:"template,omitempty"`
}

// LLMInferenceServiceAddress holds a named URL for the service.
type LLMInferenceServiceAddress struct {
	// Name identifies the address (e.g. "gateway-internal")
	// +optional
	Name string `json:"name,omitempty"`
	// URL is the endpoint address
	// +optional
	URL string `json:"url,omitempty"`
}

// LLMInferenceServiceStatus defines the observed state of LLMInferenceService
type LLMInferenceServiceStatus struct {
	knativev1.Status `json:",inline"`
	// URL of the primary inference service endpoint
	// +optional
	URL string `json:"url,omitempty"`
	// Addresses lists all named endpoints for this service
	// +optional
	Addresses []LLMInferenceServiceAddress `json:"addresses,omitempty"`
}
