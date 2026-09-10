// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
)

// Operations related to wallets
//
// WalletDepositAccountCryptoOrderService contains methods and other services that
// help with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletDepositAccountCryptoOrderService] method instead.
type WalletDepositAccountCryptoOrderService struct {
	Options []option.RequestOption
}

// NewWalletDepositAccountCryptoOrderService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWalletDepositAccountCryptoOrderService(opts ...option.RequestOption) (r WalletDepositAccountCryptoOrderService) {
	r = WalletDepositAccountCryptoOrderService{}
	r.Options = opts
	return
}

// Fetch a crypto deposit-account sweep by wallet action ID. Returns
// `{id, status}`. The path wallet is the destination (same as create). Accepts an
// app secret or a user / wallet-signer JWT (`privy-app-id`).
func (r *WalletDepositAccountCryptoOrderService) Get(ctx context.Context, orderID string, query WalletDepositAccountCryptoOrderGetParams, opts ...option.RequestOption) (res *GetCryptoDepositAccountOrderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WalletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/crypto/orders/%s", url.PathEscape(query.WalletID), orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WalletDepositAccountCryptoOrderGetParams struct {
	// ID of the wallet.
	WalletID string `path:"wallet_id" api:"required" json:"-"`
	paramObj
}
