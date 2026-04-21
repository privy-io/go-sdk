// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// WalletEarnService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletEarnService] method instead.
type WalletEarnService struct {
	Options []option.RequestOption
	// Operations related to wallet actions
	Ethereum WalletEarnEthereumService
}

// NewWalletEarnService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWalletEarnService(opts ...option.RequestOption) (r WalletEarnService) {
	r = WalletEarnService{}
	r.Options = opts
	r.Ethereum = NewWalletEarnEthereumService(opts...)
	return
}
