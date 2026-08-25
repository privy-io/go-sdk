// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// WalletDepositAccountService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletDepositAccountService] method instead.
type WalletDepositAccountService struct {
	Options []option.RequestOption
	// Operations related to wallets
	Crypto WalletDepositAccountCryptoService
	// Operations related to fiat onramping and offramping
	Fiat WalletDepositAccountFiatService
}

// NewWalletDepositAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWalletDepositAccountService(opts ...option.RequestOption) (r WalletDepositAccountService) {
	r = WalletDepositAccountService{}
	r.Options = opts
	r.Crypto = NewWalletDepositAccountCryptoService(opts...)
	r.Fiat = NewWalletDepositAccountFiatService(opts...)
	return
}
