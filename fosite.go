/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package fosite

import (
	"html/template"
	"net/http"
	"reflect"

	"github.com/ory/fosite/i18n"
)

// AuthorizeEndpointHandlers is a list of AuthorizeEndpointHandler
type AuthorizeEndpointHandlers []AuthorizeEndpointHandler

// Append adds an AuthorizeEndpointHandler to this list. Ignores duplicates based on reflect.TypeOf.
func (a *AuthorizeEndpointHandlers) Append(h AuthorizeEndpointHandler) {
	for _, this := range *a {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*a = append(*a, h)
}

// TokenEndpointHandlers is a list of TokenEndpointHandler
type TokenEndpointHandlers []TokenEndpointHandler

// Append adds an TokenEndpointHandler to this list. Ignores duplicates based on reflect.TypeOf.
func (t *TokenEndpointHandlers) Append(h TokenEndpointHandler) {
	for _, this := range *t {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*t = append(*t, h)
}

// TokenIntrospectionHandlers is a list of TokenValidator
type TokenIntrospectionHandlers []TokenIntrospector

// Append adds an AccessTokenValidator to this list. Ignores duplicates based on reflect.TypeOf.
func (t *TokenIntrospectionHandlers) Append(h TokenIntrospector) {
	for _, this := range *t {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*t = append(*t, h)
}

// RevocationHandlers is a list of RevocationHandler
type RevocationHandlers []RevocationHandler

// Append adds an RevocationHandler to this list. Ignores duplicates based on reflect.TypeOf.
func (t *RevocationHandlers) Append(h RevocationHandler) {
	for _, this := range *t {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*t = append(*t, h)
}

// Fosite implements OAuth2Provider.
type Fosite struct {
	Store                      Storage
	AuthorizeEndpointHandlers  AuthorizeEndpointHandlers
	TokenEndpointHandlers      TokenEndpointHandlers
	TokenIntrospectionHandlers TokenIntrospectionHandlers
	RevocationHandlers         RevocationHandlers
	Hasher                     Hasher
	ScopeStrategy              ScopeStrategy
	AudienceMatchingStrategy   AudienceMatchingStrategy
	JWKSFetcherStrategy        JWKSFetcherStrategy
	HTTPClient                 *http.Client
	UseLegacyErrorFormat       bool

	// TokenURL is the the URL of the Authorization Server's Token Endpoint.
	TokenURL string

	// SendDebugMessagesToClients if set to true, includes error debug messages in response payloads. Be aware that sensitive
	// data may be exposed, depending on your implementation of Fosite. Such sensitive data might include database error
	// codes or other information. Proceed with caution!
	SendDebugMessagesToClients bool

	// MinParameterEntropy controls the minimum size of state and nonce parameters. Defaults to fosite.MinParameterEntropy.
	MinParameterEntropy int

	// FormPostHTMLTemplate sets html template for rendering the authorization response when the request has response_mode=form_post. Defaults to fosite.FormPostDefaultTemplate
	FormPostHTMLTemplate *template.Template

	// WebMessageHTMLTemplate sets the html template used for response_mode=web_message. Defaults to fosite.WebMessageHTMLTemplate
	WebMessageHTMLTemplate *template.Template

	// ClientAuthenticationStrategy provides an extension point to plug a strategy to authenticate clients
	ClientAuthenticationStrategy ClientAuthenticationStrategy

	ResponseModeHandlerExtension ResponseModeHandler

	// MessageCatalog is the catalog of messages used for i18n
	MessageCatalog i18n.MessageCatalog
}

const MinParameterEntropy = 8

// GetMinParameterEntropy returns MinParameterEntropy if set. Defaults to fosite.MinParameterEntropy.
func (f *Fosite) GetMinParameterEntropy() int {
	if f.MinParameterEntropy == 0 {
		return MinParameterEntropy
	} else {
		return f.MinParameterEntropy
	}
}

var defaultResponseModeHandler = &DefaultResponseModeHandler{}

func (f *Fosite) ResponseModeHandler() ResponseModeHandler {
	if f.ResponseModeHandlerExtension == nil {
		return defaultResponseModeHandler
	}
	return f.ResponseModeHandlerExtension
}
