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

// OAuth2Grant maps a grant to an authorized credential service action.
type OAuth2Grant struct {
	// GrantType is the exact OAuth grant_type, including extension grant URIs.
	// +kubebuilder:validation:MinLength=1
	GrantType string `json:"grantType"`
	// Action is the credential service's X-Api-Action-Name.
	// +kubebuilder:validation:MinLength=1
	Action string `json:"action"`
	// RequestParameters maps public OAuth request parameter names to credential
	// service JSON fields. Client IDs/secrets and refresh tokens cannot be mapped.
	// Trusted credentialProviderName and resourceId cannot be overridden.
	// +optional
	RequestParameters map[string]string `json:"requestParameters,omitempty"`
	// RequiredParameters declares required inputs for an extension grant.
	// Core grants also validate their standard required parameters.
	// +optional
	RequiredParameters []string `json:"requiredParameters,omitempty"`
	// Response overrides the endpoint's public response extensions for this grant.
	// +optional
	Response *OAuth2Response `json:"response,omitempty"`
}

// OAuth2Config describes an OAuth endpoint independently of its URL or vendor.
type OAuth2Config struct {
	// Operation selects the endpoint role. Resource injects a Bearer credential.
	// Authorization delegates authorization-code/PKCE session setup and callback
	// binding to the trusted credential service and returns its authorization URL.
	// +kubebuilder:validation:Enum=Authorization;DeviceAuthorization;Token;Resource
	Operation string `json:"operation"`
	// Action is required for non-Token operations.
	// +optional
	Action string `json:"action,omitempty"`
	// RequestFormat defaults to Form. JSON is an explicit client compatibility
	// option. Authorization also accepts a standard GET query.
	// +optional
	// +kubebuilder:validation:Enum=Form;JSON
	RequestFormat string `json:"requestFormat,omitempty"`
	// RequestParameters maps public endpoint request parameters to credential
	// service fields. The same restrictions as OAuth2Grant apply.
	// +optional
	RequestParameters map[string]string `json:"requestParameters,omitempty"`
	// Grants selects the supported token grants and their service actions.
	// +optional
	// +listType=map
	// +listMapKey=grantType
	Grants []OAuth2Grant `json:"grants,omitempty"`
	// AllowAuthorizedSession permits a broker-confirmed already-authorized
	// device session with no verification URI. It is not an RFC 8628 response
	// and must be explicitly selected for clients that support this extension.
	// +optional
	AllowAuthorizedSession bool `json:"allowAuthorizedSession,omitempty"`
	// Response adds public response fields. Token grants inherit it unless they
	// declare their own Response.
	// +optional
	Response *OAuth2Response `json:"response,omitempty"`
}
