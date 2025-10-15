// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/privy-api-client-go/internal/apijson"
	"github.com/stainless-sdks/privy-api-client-go/internal/apiquery"
	"github.com/stainless-sdks/privy-api-client-go/internal/requestconfig"
	"github.com/stainless-sdks/privy-api-client-go/option"
	"github.com/stainless-sdks/privy-api-client-go/packages/param"
	"github.com/stainless-sdks/privy-api-client-go/packages/respjson"
)

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
		return
	}
	path := fmt.Sprintf("v1/wallets/%s/balance", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WalletBalanceGetResponse struct {
	Balances []WalletBalanceGetResponseBalance `json:"balances,required"`
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
	// Any of "usdc", "eth", "pol", "usdt", "sol".
	Asset string `json:"asset,required"`
	// Any of "ethereum", "arbitrum", "base", "linea", "optimism", "polygon", "solana",
	// "zksync_era", "sepolia", "arbitrum_sepolia", "base_sepolia", "linea_testnet",
	// "optimism_sepolia", "polygon_amoy".
	Chain            string            `json:"chain,required"`
	DisplayValues    map[string]string `json:"display_values,required"`
	RawValue         string            `json:"raw_value,required"`
	RawValueDecimals float64           `json:"raw_value_decimals,required"`
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

type WalletBalanceGetParams struct {
	Asset WalletBalanceGetParamsAssetUnion `query:"asset,omitzero,required" json:"-"`
	Chain WalletBalanceGetParamsChainUnion `query:"chain,omitzero,required" json:"-"`
	// Any of "usd".
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
type WalletBalanceGetParamsAssetUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWalletBalanceGetsAssetString)
	OfWalletBalanceGetsAssetString         param.Opt[string] `query:",omitzero,inline"`
	OfWalletBalanceGetsAssetArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *WalletBalanceGetParamsAssetUnion) asAny() any {
	if !param.IsOmitted(u.OfWalletBalanceGetsAssetString) {
		return &u.OfWalletBalanceGetsAssetString
	} else if !param.IsOmitted(u.OfWalletBalanceGetsAssetArrayItemArray) {
		return &u.OfWalletBalanceGetsAssetArrayItemArray
	}
	return nil
}

type WalletBalanceGetParamsAssetString string

const (
	WalletBalanceGetParamsAssetStringUsdc WalletBalanceGetParamsAssetString = "usdc"
	WalletBalanceGetParamsAssetStringEth  WalletBalanceGetParamsAssetString = "eth"
	WalletBalanceGetParamsAssetStringPol  WalletBalanceGetParamsAssetString = "pol"
	WalletBalanceGetParamsAssetStringUsdt WalletBalanceGetParamsAssetString = "usdt"
	WalletBalanceGetParamsAssetStringSol  WalletBalanceGetParamsAssetString = "sol"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletBalanceGetParamsChainUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWalletBalanceGetsChainString)
	OfWalletBalanceGetsChainString         param.Opt[string] `query:",omitzero,inline"`
	OfWalletBalanceGetsChainArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *WalletBalanceGetParamsChainUnion) asAny() any {
	if !param.IsOmitted(u.OfWalletBalanceGetsChainString) {
		return &u.OfWalletBalanceGetsChainString
	} else if !param.IsOmitted(u.OfWalletBalanceGetsChainArrayItemArray) {
		return &u.OfWalletBalanceGetsChainArrayItemArray
	}
	return nil
}

type WalletBalanceGetParamsChainString string

const (
	WalletBalanceGetParamsChainStringEthereum        WalletBalanceGetParamsChainString = "ethereum"
	WalletBalanceGetParamsChainStringArbitrum        WalletBalanceGetParamsChainString = "arbitrum"
	WalletBalanceGetParamsChainStringBase            WalletBalanceGetParamsChainString = "base"
	WalletBalanceGetParamsChainStringLinea           WalletBalanceGetParamsChainString = "linea"
	WalletBalanceGetParamsChainStringOptimism        WalletBalanceGetParamsChainString = "optimism"
	WalletBalanceGetParamsChainStringPolygon         WalletBalanceGetParamsChainString = "polygon"
	WalletBalanceGetParamsChainStringSolana          WalletBalanceGetParamsChainString = "solana"
	WalletBalanceGetParamsChainStringZksyncEra       WalletBalanceGetParamsChainString = "zksync_era"
	WalletBalanceGetParamsChainStringSepolia         WalletBalanceGetParamsChainString = "sepolia"
	WalletBalanceGetParamsChainStringArbitrumSepolia WalletBalanceGetParamsChainString = "arbitrum_sepolia"
	WalletBalanceGetParamsChainStringBaseSepolia     WalletBalanceGetParamsChainString = "base_sepolia"
	WalletBalanceGetParamsChainStringLineaTestnet    WalletBalanceGetParamsChainString = "linea_testnet"
	WalletBalanceGetParamsChainStringOptimismSepolia WalletBalanceGetParamsChainString = "optimism_sepolia"
	WalletBalanceGetParamsChainStringPolygonAmoy     WalletBalanceGetParamsChainString = "polygon_amoy"
)

type WalletBalanceGetParamsIncludeCurrency string

const (
	WalletBalanceGetParamsIncludeCurrencyUsd WalletBalanceGetParamsIncludeCurrency = "usd"
)
