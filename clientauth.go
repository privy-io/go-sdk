// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// ClientAuthService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientAuthService] method instead.
type ClientAuthService struct {
	Options []option.RequestOption
}

// NewClientAuthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientAuthService(opts ...option.RequestOption) (r ClientAuthService) {
	r = ClientAuthService{}
	r.Options = opts
	return
}

type CustomOAuthProviderID = string
