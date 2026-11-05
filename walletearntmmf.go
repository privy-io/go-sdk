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
	"github.com/privy-io/go-sdk/internal/apiquery"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
)

// Operations related to wallet actions
//
// WalletEarnTmmfService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletEarnTmmfService] method instead.
type WalletEarnTmmfService struct {
	Options []option.RequestOption
}

// NewWalletEarnTmmfService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWalletEarnTmmfService(opts ...option.RequestOption) (r WalletEarnTmmfService) {
	r = WalletEarnTmmfService{}
	r.Options = opts
	return
}

// Sends `deposit_asset` from the registered wallet to the fund provider and tracks
// delivery of the shares as a wallet action. The wallet owner or signers authorize
// the request.
func (r *WalletEarnTmmfService) Deposit(ctx context.Context, walletID string, params WalletEarnTmmfDepositParams, opts ...option.RequestOption) (res *TmmfDepositActionResponse, err error) {
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
	path := fmt.Sprintf("v1/wallets/%s/earn/tmmf/deposit", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a page of the wallet's tMMF positions for all providers: shares, their
// USD value at the latest NAV, pending dividends, and pending orders.
func (r *WalletEarnTmmfService) ListPositions(ctx context.Context, walletID string, query WalletEarnTmmfListPositionsParams, opts ...option.RequestOption) (res *TmmfPositionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/earn/tmmf/positions", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Registers the wallet with a tMMF provider on one chain. The provider must have
// approved the wallet's user, e.g. an organization; the response reports whether
// it also approved the wallet.
func (r *WalletEarnTmmfService) Register(ctx context.Context, walletID string, body WalletEarnTmmfRegisterParams, opts ...option.RequestOption) (res *TmmfWalletResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/earn/tmmf/register", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sends fund shares from the registered wallet to the fund provider and tracks
// delivery of `withdraw_asset` as a wallet action. The wallet owner or signers
// authorize the request.
func (r *WalletEarnTmmfService) Withdraw(ctx context.Context, walletID string, params WalletEarnTmmfWithdrawParams, opts ...option.RequestOption) (res *TmmfWithdrawActionResponse, err error) {
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
	path := fmt.Sprintf("v1/wallets/%s/earn/tmmf/withdraw", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type WalletEarnTmmfDepositParams struct {
	// Request body for buying tMMF shares with a currency the wallet holds.
	TmmfDepositRequestBody TmmfDepositRequestBody
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

func (r WalletEarnTmmfDepositParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TmmfDepositRequestBody)
}
func (r *WalletEarnTmmfDepositParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletEarnTmmfListPositionsParams struct {
	// Chain the fund is held on, e.g. "ethereum_sepolia".
	Chain param.Opt[TmmfChain] `query:"chain,omitzero" json:"-"`
	// Cursor returned by the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Fund ticker, e.g. "WTGXX". Case-insensitive.
	Fund  param.Opt[TmmfFund] `query:"fund,omitzero" json:"-"`
	Limit param.Opt[int64]    `query:"limit,omitzero" json:"-"`
	// The fund provider.
	Provider TmmfProvider `query:"provider,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletEarnTmmfListPositionsParams]'s query parameters as
// `url.Values`.
func (r WalletEarnTmmfListPositionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletEarnTmmfRegisterParams struct {
	// Request body for registering a wallet with a tMMF provider for the wallet's
	// user, e.g. an organization.
	TmmfRegisterRequestBody TmmfRegisterRequestBody
	paramObj
}

func (r WalletEarnTmmfRegisterParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TmmfRegisterRequestBody)
}
func (r *WalletEarnTmmfRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletEarnTmmfWithdrawParams struct {
	// Request body for redeeming tMMF shares the wallet holds.
	TmmfWithdrawRequestBody TmmfWithdrawRequestBody
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

func (r WalletEarnTmmfWithdrawParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TmmfWithdrawRequestBody)
}
func (r *WalletEarnTmmfWithdrawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
