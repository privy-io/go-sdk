// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/apiquery"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/pagination"
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
	// Operations related to wallets
	Orders WalletDepositAccountCryptoOrderService
}

// NewWalletDepositAccountCryptoService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWalletDepositAccountCryptoService(opts ...option.RequestOption) (r WalletDepositAccountCryptoService) {
	r = WalletDepositAccountCryptoService{}
	r.Options = opts
	r.Orders = NewWalletDepositAccountCryptoOrderService(opts...)
	return
}

// Creates deposit source wallets that sweep into the path wallet.
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

// Returns active crypto deposit accounts that sweep into the path wallet. Requires
// an app secret or a JWT for a wallet signer, plus `privy-app-id`.
func (r *WalletDepositAccountCryptoService) List(ctx context.Context, walletID string, query WalletDepositAccountCryptoListParams, opts ...option.RequestOption) (res *pagination.Cursor[CryptoDepositAddressRoute], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/crypto", url.PathEscape(walletID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns active crypto deposit accounts that sweep into the path wallet. Requires
// an app secret or a JWT for a wallet signer, plus `privy-app-id`.
func (r *WalletDepositAccountCryptoService) ListAutoPaging(ctx context.Context, walletID string, query WalletDepositAccountCryptoListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CryptoDepositAddressRoute] {
	return pagination.NewCursorAutoPager(r.List(ctx, walletID, query, opts...))
}

// Returns the tokens and chains a user can send from when creating a crypto
// deposit account.
func (r *WalletDepositAccountCryptoService) GetConfig(ctx context.Context, opts ...option.RequestOption) (res *CryptoDepositAccountConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/deposit_accounts/crypto/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the earliest crypto deposit-account sweep into the path wallet after
// `after`. Returns `{order: {id, status} | null}` — the same order object as GET
// order. The path wallet is the destination (same as create). Accepts an app
// secret or a user / wallet-signer JWT (`privy-app-id`).
func (r *WalletDepositAccountCryptoService) GetNextOrder(ctx context.Context, walletID string, query WalletDepositAccountCryptoGetNextOrderParams, opts ...option.RequestOption) (res *GetCryptoDepositAccountNextOrderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/crypto/next_order", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns an indicative route quote without creating a wallet. Amounts use token
// standard units. Accepts an app secret or user token.
func (r *WalletDepositAccountCryptoService) Quote(ctx context.Context, body WalletDepositAccountCryptoQuoteParams, opts ...option.RequestOption) (res *DepositAccountCryptoQuoteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/deposit_accounts/crypto/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type WalletDepositAccountCryptoListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletDepositAccountCryptoListParams]'s query parameters as
// `url.Values`.
func (r WalletDepositAccountCryptoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletDepositAccountCryptoGetNextOrderParams struct {
	// Return the earliest sweep strictly after this timestamp.
	After time.Time `query:"after" api:"required" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [WalletDepositAccountCryptoGetNextOrderParams]'s query
// parameters as `url.Values`.
func (r WalletDepositAccountCryptoGetNextOrderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletDepositAccountCryptoQuoteParams struct {
	// Request body for an indicative crypto deposit-account route quote.
	DepositAccountCryptoQuoteRequestBody DepositAccountCryptoQuoteRequestBody
	paramObj
}

func (r WalletDepositAccountCryptoQuoteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DepositAccountCryptoQuoteRequestBody)
}
func (r *WalletDepositAccountCryptoQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
