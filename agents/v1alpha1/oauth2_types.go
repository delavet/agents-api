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

// OAuth2ErrorResponse adds public JSON fields to generated OAuth errors.
// Standard fields cannot be overridden; provider credentials are never in scope.
type OAuth2ErrorResponse struct {
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

// OAuth2Response adds public JSON fields to generated token responses.
type OAuth2Response struct {
	OAuth2ErrorResponse `json:",inline"`
	// Success adds fields to a successful token response.
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Success *runtime.RawExtension `json:"success,omitempty"`
}

// OAuth2Config selects an OAuth endpoint independently of its URL or vendor.
type OAuth2Config struct {
	// Operation contains exactly one endpoint role and its options.
	Operation OAuth2Operation `json:"operation"`
}

// OAuth2Operation selects exactly one endpoint role. Empty role objects are valid.
// +kubebuilder:validation:XValidation:rule="(has(self.deviceAuthorization) ? 1 : 0) + (has(self.token) ? 1 : 0) + (has(self.resource) ? 1 : 0) == 1",message="exactly one OAuth2 operation is required"
type OAuth2Operation struct {
	// DeviceAuthorization creates a device authorization session.
	// +optional
	DeviceAuthorization *OAuth2DeviceAuthorizationConfig `json:"deviceAuthorization,omitempty"`
	// Token processes token grants and returns placeholder credentials.
	// +optional
	Token *OAuth2TokenConfig `json:"token,omitempty"`
	// Resource injects a service-issued Bearer credential into the request.
	// +optional
	Resource *OAuth2ResourceConfig `json:"resource,omitempty"`
}

// OAuth2DeviceAuthorizationConfig configures the device authorization endpoint.
type OAuth2DeviceAuthorizationConfig struct {
	// AllowAuthorizedSession permits a broker-confirmed already-authorized
	// device session with no verification URI. Defaults to true. Set false for
	// clients requiring an RFC 8628 response; it does not force a new session.
	// +optional
	// +kubebuilder:default=true
	AllowAuthorizedSession *bool `json:"allowAuthorizedSession,omitempty"`
	// CustomResponse adds public fields to generated OAuth errors.
	// +optional
	CustomResponse *OAuth2ErrorResponse `json:"customResponse,omitempty"`
}

// OAuth2TokenConfig configures the token endpoint. Form and JSON requests are
// decoded according to Content-Type; their response extensions are policy-selected.
type OAuth2TokenConfig struct {
	// AllowedGrantTypes optionally restricts requests to exact grant_type values,
	// including extension grant URIs. Omitted or empty adds no policy restriction;
	// protocol support and service authorization still apply.
	// +optional
	// +listType=set
	// +kubebuilder:validation:items:MinLength=1
	AllowedGrantTypes []string `json:"allowedGrantTypes,omitempty"`
	// CustomResponse adds public fields to generated token success and errors.
	// It cannot replace standard OAuth fields.
	// +optional
	CustomResponse *OAuth2Response `json:"customResponse,omitempty"`
}

// OAuth2ResourceConfig configures Bearer injection for resource requests.
type OAuth2ResourceConfig struct {
	// CustomResponse adds public fields to generated OAuth errors. Upstream
	// resource responses are not modified.
	// +optional
	CustomResponse *OAuth2ErrorResponse `json:"customResponse,omitempty"`
}
