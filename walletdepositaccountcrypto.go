// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
)

// Operations related to wallets
//
// WalletDepositAccountCryptoService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletDepositAccountCryptoService] method instead.
type WalletDepositAccountCryptoService struct {
	Options []option.RequestOption
}

// NewWalletDepositAccountCryptoService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWalletDepositAccountCryptoService(opts ...option.RequestOption) (r WalletDepositAccountCryptoService) {
	r = WalletDepositAccountCryptoService{}
	r.Options = opts
	return
}

// Creates deposit source wallets and attaches them to a sweep into the path
// wallet. Requires a dest-owner privy-authorization-signature. Accepts a
// dest-owner user JWT or an app secret (app-secret callers use the dest owner).
// JWT-only requests 401 when the app requires an app secret for wallet actions.
func (r *WalletDepositAccountCryptoService) New(ctx context.Context, walletID string, params WalletDepositAccountCryptoNewParams, opts ...option.RequestOption) (res *CreateCryptoDepositAccountResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/crypto", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type WalletDepositAccountCryptoNewParams struct {
	// Request body for creating a crypto deposit account.
	CreateCryptoDepositAccountRequestBody CreateCryptoDepositAccountRequestBodyUnion
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r WalletDepositAccountCryptoNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCryptoDepositAccountRequestBody)
}
func (r *WalletDepositAccountCryptoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
