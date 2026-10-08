/*
Copyright 2026.

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

import "k8s.io/apimachinery/pkg/runtime"

// OAuth2Response configures additional public JSON fields. Standard OAuth
// fields cannot be overridden, and real provider credentials are never in scope.
type OAuth2Response struct {
	// Success adds fields to a successful token response.
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Success *runtime.RawExtension `json:"success,omitempty"`
	// Error adds fields to an OAuth error response.
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Error *runtime.RawExtension `json:"error,omitempty"`
	// AuthorizationRequired adds fields when the credential service confirms
	// a non-retryable need for new authorization.
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	AuthorizationRequired *runtime.RawExtension `json:"authorizationRequired,omitempty"`
}

// OAuth2Config describes an OAuth endpoint independently of its URL or vendor.
// +kubebuilder:validation:XValidation:rule="self.operation == 'Token' || !has(self.allowedGrantTypes) || size(self.allowedGrantTypes) == 0",message="allowedGrantTypes requires the Token operation"
type OAuth2Config struct {
	// Operation selects the endpoint role. Resource injects a Bearer credential.
	// +kubebuilder:validation:Enum=DeviceAuthorization;Token;Resource
	Operation string `json:"operation"`
	// RequestFormat defaults to Form. JSON is an explicit client compatibility
	// option.
	// +optional
	// +kubebuilder:validation:Enum=Form;JSON
	// +kubebuilder:default=Form
	RequestFormat string `json:"requestFormat,omitempty"`
	// AllowedGrantTypes optionally restricts Token requests to these exact
	// grant_type values, including extension grant URIs. Omitted or empty means
	// no additional policy restriction; protocol support and service authorization
	// still apply. Non-empty lists are only valid for the Token operation.
	// +optional
	// +listType=set
	// +kubebuilder:validation:items:MinLength=1
	AllowedGrantTypes []string `json:"allowedGrantTypes,omitempty"`
	// AllowAuthorizedSession permits a broker-confirmed already-authorized
	// device session with no verification URI. Defaults to true and only affects
	// DeviceAuthorization. Set false for clients requiring an RFC 8628 response.
	// +optional
	// +kubebuilder:default=true
	AllowAuthorizedSession *bool `json:"allowAuthorizedSession,omitempty"`
	// CustomResponse adds public JSON fields to generated token success and
	// OAuth error responses. It cannot replace standard fields or alter an
	// upstream Resource response. Omitted means no additional fields.
	// +optional
	CustomResponse *OAuth2Response `json:"customResponse,omitempty"`
}
