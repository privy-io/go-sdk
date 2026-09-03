// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// WalletPayoutService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletPayoutService] method instead.
type WalletPayoutService struct {
	Options []option.RequestOption
	// Operations related to fiat onramping and offramping
	Fiat WalletPayoutFiatService
}

// NewWalletPayoutService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWalletPayoutService(opts ...option.RequestOption) (r WalletPayoutService) {
	r = WalletPayoutService{}
	r.Options = opts
	r.Fiat = NewWalletPayoutFiatService(opts...)
	return
}
