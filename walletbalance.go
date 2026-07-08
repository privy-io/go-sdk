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
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to wallets
//
// WalletBalanceService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletBalanceService] method instead.
type WalletBalanceService struct {
	Options []option.RequestOption
}

// NewWalletBalanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWalletBalanceService(opts ...option.RequestOption) (r WalletBalanceService) {
	r = WalletBalanceService{}
	r.Options = opts
	return
}

// Get the balance of a wallet by wallet ID.
func (r *WalletBalanceService) Get(ctx context.Context, walletID string, query WalletBalanceGetParams, opts ...option.RequestOption) (res *WalletBalanceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/balance", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WalletBalanceGetResponse struct {
	Balances []WalletBalanceGetResponseBalance `json:"balances" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balances    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletBalanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletBalanceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletBalanceGetResponseBalance struct {
	Asset WalletBalanceGetResponseBalanceAsset `json:"asset" api:"required"`
	// Supported blockchain network names for wallet balance and transaction queries.
	//
	// Any of "ethereum", "arbitrum", "avalanche", "base", "tempo", "linea",
	// "optimism", "polygon", "bsc", "solana", "tron", "zksync_era", "sepolia",
	// "arbitrum_sepolia", "avalanche_fuji", "base_sepolia", "linea_testnet",
	// "optimism_sepolia", "polygon_amoy", "solana_devnet", "solana_testnet",
	// "tron_nile".
	Chain            WalletAssetChainNameInput `json:"chain" api:"required"`
	DisplayValues    map[string]string         `json:"display_values" api:"required"`
	RawValue         string                    `json:"raw_value" api:"required"`
	RawValueDecimals float64                   `json:"raw_value_decimals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset            respjson.Field
		Chain            respjson.Field
		DisplayValues    respjson.Field
		RawValue         respjson.Field
		RawValueDecimals respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletBalanceGetResponseBalance) RawJSON() string { return r.JSON.raw }
func (r *WalletBalanceGetResponseBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletBalanceGetResponseBalanceAsset string

const (
	WalletBalanceGetResponseBalanceAssetUsdc  WalletBalanceGetResponseBalanceAsset = "usdc"
	WalletBalanceGetResponseBalanceAssetUsdcE WalletBalanceGetResponseBalanceAsset = "usdc.e"
	WalletBalanceGetResponseBalanceAssetEth   WalletBalanceGetResponseBalanceAsset = "eth"
	WalletBalanceGetResponseBalanceAssetAvax  WalletBalanceGetResponseBalanceAsset = "avax"
	WalletBalanceGetResponseBalanceAssetPol   WalletBalanceGetResponseBalanceAsset = "pol"
	WalletBalanceGetResponseBalanceAssetBnb   WalletBalanceGetResponseBalanceAsset = "bnb"
	WalletBalanceGetResponseBalanceAssetUsdt  WalletBalanceGetResponseBalanceAsset = "usdt"
	WalletBalanceGetResponseBalanceAssetEurc  WalletBalanceGetResponseBalanceAsset = "eurc"
	WalletBalanceGetResponseBalanceAssetUsdb  WalletBalanceGetResponseBalanceAsset = "usdb"
	WalletBalanceGetResponseBalanceAssetSol   WalletBalanceGetResponseBalanceAsset = "sol"
	WalletBalanceGetResponseBalanceAssetTrx   WalletBalanceGetResponseBalanceAsset = "trx"
)

type WalletBalanceGetParams struct {
	// Include archived wallets in lookup. Defaults to false.
	IncludeArchived param.Opt[bool] `query:"include_archived,omitzero" json:"-"`
	// The token contract address(es) to query in format "chain:address" (e.g.,
	// "base:0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913" or
	// "solana:EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"). Cannot be used together
	// with `asset`/`chain` or with `include_currency`.
	Token WalletBalanceGetParamsTokenUnion `query:"token,omitzero" json:"-"`
	// Named asset(s) to query (e.g. `eth`, `usdc`). Use together with `chain` to scope
	// the query. Cannot be used with `token`.
	Asset WalletBalanceGetParamsAssetUnion `query:"asset,omitzero" json:"-"`
	// Chain(s) to query named assets on (e.g. `base`, `ethereum`). Use together with
	// `asset`. Cannot be used with `token`.
	Chain WalletBalanceGetParamsChainUnion `query:"chain,omitzero" json:"-"`
	// If set, balances are converted to the specified fiat currency. Not supported
	// when `token` is provided.
	//
	// Any of "usd", "eur".
	IncludeCurrency WalletBalanceGetParamsIncludeCurrency `query:"include_currency,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletBalanceGetParams]'s query parameters as `url.Values`.
func (r WalletBalanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletBalanceGetParamsTokenUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletBalanceGetParamsAssetUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWalletBalanceGetsAssetString)
	OfWalletBalanceGetsAssetString param.Opt[WalletBalanceGetParamsAssetString] `query:",omitzero,inline"`
	OfWalletAssetArray             []WalletAsset                                `query:",omitzero,inline"`
	paramUnion
}

type WalletBalanceGetParamsAssetString string

const (
	WalletBalanceGetParamsAssetStringUsdc  WalletBalanceGetParamsAssetString = "usdc"
	WalletBalanceGetParamsAssetStringUsdcE WalletBalanceGetParamsAssetString = "usdc.e"
	WalletBalanceGetParamsAssetStringEth   WalletBalanceGetParamsAssetString = "eth"
	WalletBalanceGetParamsAssetStringAvax  WalletBalanceGetParamsAssetString = "avax"
	WalletBalanceGetParamsAssetStringPol   WalletBalanceGetParamsAssetString = "pol"
	WalletBalanceGetParamsAssetStringBnb   WalletBalanceGetParamsAssetString = "bnb"
	WalletBalanceGetParamsAssetStringUsdt  WalletBalanceGetParamsAssetString = "usdt"
	WalletBalanceGetParamsAssetStringEurc  WalletBalanceGetParamsAssetString = "eurc"
	WalletBalanceGetParamsAssetStringUsdb  WalletBalanceGetParamsAssetString = "usdb"
	WalletBalanceGetParamsAssetStringSol   WalletBalanceGetParamsAssetString = "sol"
	WalletBalanceGetParamsAssetStringTrx   WalletBalanceGetParamsAssetString = "trx"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletBalanceGetParamsChainUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWalletAssetChainNameInput)
	OfWalletAssetChainNameInput      param.Opt[WalletAssetChainNameInput] `query:",omitzero,inline"`
	OfWalletAssetChainNameInputArray []WalletAssetChainNameInput          `query:",omitzero,inline"`
	paramUnion
}

// If set, balances are converted to the specified fiat currency. Not supported
// when `token` is provided.
type WalletBalanceGetParamsIncludeCurrency string

const (
	WalletBalanceGetParamsIncludeCurrencyUsd WalletBalanceGetParamsIncludeCurrency = "usd"
	WalletBalanceGetParamsIncludeCurrencyEur WalletBalanceGetParamsIncludeCurrency = "eur"
)
