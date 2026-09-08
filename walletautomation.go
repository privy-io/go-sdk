// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"net/http"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to wallet automations
//
// WalletAutomationService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletAutomationService] method instead.
type WalletAutomationService struct {
	Options []option.RequestOption
}

// NewWalletAutomationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWalletAutomationService(opts ...option.RequestOption) (r WalletAutomationService) {
	r = WalletAutomationService{}
	r.Options = opts
	return
}

// Re-checks a wallet (identified by wallet_id or deposit_address) for funds
// matching its wallet automation configs and triggers an automation run if a match
// is found. Use this to recover a deposit whose automation trigger was missed or
// failed.
func (r *WalletAutomationService) Reindex(ctx context.Context, body WalletAutomationReindexParams, opts ...option.RequestOption) (res *WalletAutomationReindexResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallet_automations/reindex"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The outcome of checking one asset on the requested chain during a reindex.
type WalletAutomationReindexAssetResult struct {
	// Asset contract address; the native asset uses `native`.
	AssetAddress string `json:"asset_address" api:"required"`
	// EVM CAIP-2 chain identifier (e.g. "eip155:4217" for Tempo, "eip155:1" for
	// Ethereum).
	Caip2 WalletAutomationReindexCaip2 `json:"caip2" api:"required"`
	// ID of the in-flight execution blocking a re-trigger. Populated only when
	// `status` is `skipped_existing_execution`; `null` otherwise.
	ExistingExecutionID string `json:"existing_execution_id" api:"required"`
	// On-chain balance in base units. Populated when `status` is `triggered` or
	// `skipped_zero_balance`; `null` otherwise. For example, 1 USDC is `1000000`.
	RawBalance string `json:"raw_balance" api:"required"`
	// Outcome of checking a single asset during a wallet automation reindex. One of
	// `triggered`, `skipped_zero_balance`, `skipped_no_match`,
	// `skipped_existing_execution`, or `failed`.
	Status WalletAutomationReindexAssetStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetAddress        respjson.Field
		Caip2               respjson.Field
		ExistingExecutionID respjson.Field
		RawBalance          respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationReindexAssetResult) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationReindexAssetResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAutomationReindexAssetStatus = string

type WalletAutomationReindexCaip2 = string

// Request body for re-checking a wallet against its wallet automations. Identify
// the wallet by wallet_id or deposit_address (at least one is required). If both
// are provided, wallet_id takes precedence and deposit_address must match that
// wallet's address. Specify exactly one of caip2 or chain, and the asset_address
// to check. Useful for recovering a deposit that was missed or failed to trigger
// its automation.
//
// The property AssetAddress is required.
type WalletAutomationReindexRequestBody struct {
	// Asset contract address to check; the native asset uses `native`.
	AssetAddress string `json:"asset_address" api:"required"`
	// EVM CAIP-2 chain identifier (e.g. "eip155:4217" for Tempo, "eip155:1" for
	// Ethereum).
	Caip2 param.Opt[WalletAutomationReindexCaip2] `json:"caip2,omitzero"`
	// Human-readable chain name to check. Specify exactly one of `caip2` or `chain`.
	Chain param.Opt[string] `json:"chain,omitzero"`
	// On-chain deposit address of the wallet to reindex. Must match the resolved
	// wallet's address if `wallet_id` is also provided.
	DepositAddress param.Opt[string] `json:"deposit_address,omitzero"`
	// Privy wallet ID to reindex. Takes precedence over `deposit_address` when both
	// are supplied.
	WalletID param.Opt[string] `json:"wallet_id,omitzero"`
	paramObj
}

func (r WalletAutomationReindexRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow WalletAutomationReindexRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletAutomationReindexRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of re-checking a wallet against its wallet automations.
type WalletAutomationReindexResponse struct {
	Results  []WalletAutomationReindexAssetResult `json:"results" api:"required"`
	WalletID string                               `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationReindexResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationReindexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAutomationReindexParams struct {
	// Request body for re-checking a wallet against its wallet automations. Identify
	// the wallet by wallet_id or deposit_address (at least one is required). If both
	// are provided, wallet_id takes precedence and deposit_address must match that
	// wallet's address. Specify exactly one of caip2 or chain, and the asset_address
	// to check. Useful for recovering a deposit that was missed or failed to trigger
	// its automation.
	WalletAutomationReindexRequestBody WalletAutomationReindexRequestBody
	paramObj
}

func (r WalletAutomationReindexParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletAutomationReindexRequestBody)
}
func (r *WalletAutomationReindexParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
