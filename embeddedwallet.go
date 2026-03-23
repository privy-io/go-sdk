// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// EmbeddedWalletService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddedWalletService] method instead.
type EmbeddedWalletService struct {
	Options []option.RequestOption
}

// NewEmbeddedWalletService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddedWalletService(opts ...option.RequestOption) (r EmbeddedWalletService) {
	r = EmbeddedWalletService{}
	r.Options = opts
	return
}
