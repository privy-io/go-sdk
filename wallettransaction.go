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
// WalletTransactionService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletTransactionService] method instead.
type WalletTransactionService struct {
	Options []option.RequestOption
}

// NewWalletTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWalletTransactionService(opts ...option.RequestOption) (r WalletTransactionService) {
	r = WalletTransactionService{}
	r.Options = opts
	return
}

// Get incoming and outgoing transactions of a wallet by wallet ID.
func (r *WalletTransactionService) Get(ctx context.Context, walletID string, query WalletTransactionGetParams, opts ...option.RequestOption) (res *WalletTransactionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/transactions", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WalletTransactionGetResponse struct {
	NextCursor   string                                    `json:"next_cursor" api:"required"`
	Transactions []WalletTransactionGetResponseTransaction `json:"transactions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor   respjson.Field
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletTransactionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletTransactionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletTransactionGetResponseTransaction struct {
	Caip2     string  `json:"caip2" api:"required"`
	CreatedAt float64 `json:"created_at" api:"required"`
	// Details of a wallet transaction, varying by transaction type.
	Details            TransactionDetailUnion `json:"details" api:"required"`
	PrivyTransactionID string                 `json:"privy_transaction_id" api:"required"`
	// Status of a blockchain transaction submitted by Privy.
	//
	// Any of "broadcasted", "confirmed", "execution_reverted", "failed", "replaced",
	// "finalized", "provider_error", "pending".
	Status            BlockchainTransactionStatus `json:"status" api:"required"`
	TransactionHash   string                      `json:"transaction_hash" api:"required"`
	WalletID          string                      `json:"wallet_id" api:"required"`
	Sponsored         bool                        `json:"sponsored"`
	UserOperationHash string                      `json:"user_operation_hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2              respjson.Field
		CreatedAt          respjson.Field
		Details            respjson.Field
		PrivyTransactionID respjson.Field
		Status             respjson.Field
		TransactionHash    respjson.Field
		WalletID           respjson.Field
		Sponsored          respjson.Field
		UserOperationHash  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletTransactionGetResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *WalletTransactionGetResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletTransactionGetParams struct {
	// Chains supported for transaction history queries.
	//
	// Any of "ethereum", "arbitrum", "avalanche", "base", "base_sepolia", "bsc",
	// "tempo", "linea", "optimism", "polygon", "solana", "sepolia".
	Chain  TransactionChainNameInput `query:"chain,omitzero" api:"required" json:"-"`
	Limit  param.Opt[float64]        `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]         `query:"cursor,omitzero" json:"-"`
	// Include archived wallets in lookup. Defaults to false.
	IncludeArchived param.Opt[bool]   `query:"include_archived,omitzero" json:"-"`
	TxHash          param.Opt[string] `query:"tx_hash,omitzero" json:"-"`
	// Exactly one of `token` or `asset` is required. Cannot be used together with
	// `asset`.
	Token WalletTransactionGetParamsTokenUnion `query:"token,omitzero" json:"-"`
	// Exactly one of `asset` or `token` is required. Cannot be used together with
	// `token`.
	Asset WalletTransactionGetParamsAssetUnion `query:"asset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletTransactionGetParams]'s query parameters as
// `url.Values`.
func (r WalletTransactionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletTransactionGetParamsTokenUnion struct {
	OfString                            param.Opt[TransactionTokenAddressInput] `query:",omitzero,inline"`
	OfTransactionTokenAddressInputArray []TransactionTokenAddressInput          `query:",omitzero,inline"`
	paramUnion
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletTransactionGetParamsAssetUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWalletTransactionGetsAssetString)
	OfWalletTransactionGetsAssetString param.Opt[WalletTransactionGetParamsAssetString] `query:",omitzero,inline"`
	OfWalletAssetArray                 []WalletAsset                                    `query:",omitzero,inline"`
	paramUnion
}

type WalletTransactionGetParamsAssetString string

const (
	WalletTransactionGetParamsAssetStringUsdc    WalletTransactionGetParamsAssetString = "usdc"
	WalletTransactionGetParamsAssetStringUsdcE   WalletTransactionGetParamsAssetString = "usdc.e"
	WalletTransactionGetParamsAssetStringEth     WalletTransactionGetParamsAssetString = "eth"
	WalletTransactionGetParamsAssetStringAvax    WalletTransactionGetParamsAssetString = "avax"
	WalletTransactionGetParamsAssetStringPol     WalletTransactionGetParamsAssetString = "pol"
	WalletTransactionGetParamsAssetStringBnb     WalletTransactionGetParamsAssetString = "bnb"
	WalletTransactionGetParamsAssetStringUsdt    WalletTransactionGetParamsAssetString = "usdt"
	WalletTransactionGetParamsAssetStringEurc    WalletTransactionGetParamsAssetString = "eurc"
	WalletTransactionGetParamsAssetStringUsdb    WalletTransactionGetParamsAssetString = "usdb"
	WalletTransactionGetParamsAssetStringPathusd WalletTransactionGetParamsAssetString = "pathusd"
	WalletTransactionGetParamsAssetStringSol     WalletTransactionGetParamsAssetString = "sol"
	WalletTransactionGetParamsAssetStringTrx     WalletTransactionGetParamsAssetString = "trx"
)
