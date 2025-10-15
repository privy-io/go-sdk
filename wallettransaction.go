// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

import (
	"context"
	"encoding/json"
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
		return
	}
	path := fmt.Sprintf("v1/wallets/%s/transactions", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WalletTransactionGetResponse struct {
	NextCursor   string                                    `json:"next_cursor,required"`
	Transactions []WalletTransactionGetResponseTransaction `json:"transactions,required"`
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
	Caip2              string                                              `json:"caip2,required"`
	CreatedAt          float64                                             `json:"created_at,required"`
	Details            WalletTransactionGetResponseTransactionDetailsUnion `json:"details,required"`
	PrivyTransactionID string                                              `json:"privy_transaction_id,required"`
	// Any of "broadcasted", "confirmed", "execution_reverted", "failed", "replaced",
	// "finalized", "provider_error", "pending".
	Status          string `json:"status,required"`
	TransactionHash string `json:"transaction_hash,required"`
	WalletID        string `json:"wallet_id,required"`
	Sponsored       bool   `json:"sponsored"`
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
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletTransactionGetResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *WalletTransactionGetResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletTransactionGetResponseTransactionDetailsUnion contains all possible
// properties and values from
// [WalletTransactionGetResponseTransactionDetailsObject],
// [WalletTransactionGetResponseTransactionDetailsObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletTransactionGetResponseTransactionDetailsUnion struct {
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	Asset string `json:"asset"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	Chain string `json:"chain"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	DisplayValues map[string]string `json:"display_values"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	RawValue string `json:"raw_value"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	RawValueDecimals float64 `json:"raw_value_decimals"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	Recipient string `json:"recipient"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	RecipientPrivyUserID string `json:"recipient_privy_user_id"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	Sender string `json:"sender"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	SenderPrivyUserID string `json:"sender_privy_user_id"`
	// This field is from variant
	// [WalletTransactionGetResponseTransactionDetailsObject].
	Type string `json:"type"`
	JSON struct {
		Asset                respjson.Field
		Chain                respjson.Field
		DisplayValues        respjson.Field
		RawValue             respjson.Field
		RawValueDecimals     respjson.Field
		Recipient            respjson.Field
		RecipientPrivyUserID respjson.Field
		Sender               respjson.Field
		SenderPrivyUserID    respjson.Field
		Type                 respjson.Field
		raw                  string
	} `json:"-"`
}

func (u WalletTransactionGetResponseTransactionDetailsUnion) AsWalletTransactionGetResponseTransactionDetailsObject() (v WalletTransactionGetResponseTransactionDetailsObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletTransactionGetResponseTransactionDetailsUnion) AsVariant2() (v WalletTransactionGetResponseTransactionDetailsObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletTransactionGetResponseTransactionDetailsUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletTransactionGetResponseTransactionDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletTransactionGetResponseTransactionDetailsObject struct {
	// Any of "usdc", "eth", "pol", "usdt", "sol".
	Asset string `json:"asset,required"`
	// Any of "ethereum", "arbitrum", "base", "linea", "optimism", "polygon", "solana",
	// "zksync_era", "sepolia", "arbitrum_sepolia", "base_sepolia", "linea_testnet",
	// "optimism_sepolia", "polygon_amoy".
	Chain                string            `json:"chain,required"`
	DisplayValues        map[string]string `json:"display_values,required"`
	RawValue             string            `json:"raw_value,required"`
	RawValueDecimals     float64           `json:"raw_value_decimals,required"`
	Recipient            string            `json:"recipient,required"`
	RecipientPrivyUserID string            `json:"recipient_privy_user_id,required"`
	Sender               string            `json:"sender,required"`
	SenderPrivyUserID    string            `json:"sender_privy_user_id,required"`
	// Any of "transfer_sent".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset                respjson.Field
		Chain                respjson.Field
		DisplayValues        respjson.Field
		RawValue             respjson.Field
		RawValueDecimals     respjson.Field
		Recipient            respjson.Field
		RecipientPrivyUserID respjson.Field
		Sender               respjson.Field
		SenderPrivyUserID    respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletTransactionGetResponseTransactionDetailsObject) RawJSON() string { return r.JSON.raw }
func (r *WalletTransactionGetResponseTransactionDetailsObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletTransactionGetParams struct {
	Asset WalletTransactionGetParamsAssetUnion `query:"asset,omitzero,required" json:"-"`
	// Any of "base".
	Chain  WalletTransactionGetParamsChain `query:"chain,omitzero,required" json:"-"`
	Limit  param.Opt[float64]              `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]               `query:"cursor,omitzero" json:"-"`
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
type WalletTransactionGetParamsAssetUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWalletTransactionGetsAssetString)
	OfWalletTransactionGetsAssetString         param.Opt[string] `query:",omitzero,inline"`
	OfWalletTransactionGetsAssetArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *WalletTransactionGetParamsAssetUnion) asAny() any {
	if !param.IsOmitted(u.OfWalletTransactionGetsAssetString) {
		return &u.OfWalletTransactionGetsAssetString
	} else if !param.IsOmitted(u.OfWalletTransactionGetsAssetArrayItemArray) {
		return &u.OfWalletTransactionGetsAssetArrayItemArray
	}
	return nil
}

type WalletTransactionGetParamsAssetString string

const (
	WalletTransactionGetParamsAssetStringUsdc WalletTransactionGetParamsAssetString = "usdc"
	WalletTransactionGetParamsAssetStringEth  WalletTransactionGetParamsAssetString = "eth"
	WalletTransactionGetParamsAssetStringPol  WalletTransactionGetParamsAssetString = "pol"
	WalletTransactionGetParamsAssetStringUsdt WalletTransactionGetParamsAssetString = "usdt"
	WalletTransactionGetParamsAssetStringSol  WalletTransactionGetParamsAssetString = "sol"
)

type WalletTransactionGetParamsChain string

const (
	WalletTransactionGetParamsChainBase WalletTransactionGetParamsChain = "base"
)
