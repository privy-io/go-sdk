// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// SwapService contains methods and other services that help with interacting with
// the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwapService] method instead.
type SwapService struct {
	Options []option.RequestOption
}

// NewSwapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSwapService(opts ...option.RequestOption) (r SwapService) {
	r = SwapService{}
	r.Options = opts
	return
}

// The output side of a swap execution request.
//
// The property AssetAddress is required.
type SwapDestination struct {
	// Token contract address to buy, or "native" for the chain's native token.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier for the destination. Defaults to source chain if
	// omitted. Specify a different chain for cross-chain swaps.
	Caip2 param.Opt[string] `json:"caip2,omitzero"`
	// Address to receive the output tokens. Defaults to the swapping wallet address.
	// Required when swapping between different chain types (e.g. EVM to Solana).
	DestinationAddress param.Opt[string] `json:"destination_address,omitzero"`
	paramObj
}

func (r SwapDestination) MarshalJSON() (data []byte, err error) {
	type shadow SwapDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwapDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The output side of a swap quote request.
//
// The property AssetAddress is required.
type SwapQuoteDestination struct {
	// Token contract address to buy, or "native" for the chain's native token.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier for the destination. Defaults to source chain if
	// omitted. Will result in a cross-chain swap if source and destination chains
	// differ.
	Caip2 param.Opt[string] `json:"caip2,omitzero"`
	// Address to receive the output tokens. Defaults to the swapping wallet address.
	// Required when swapping between chains with different address types (e.g. EVM to
	// Solana).
	DestinationAddress param.Opt[string] `json:"destination_address,omitzero"`
	paramObj
}

func (r SwapQuoteDestination) MarshalJSON() (data []byte, err error) {
	type shadow SwapQuoteDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwapQuoteDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for requesting a token swap quote.
//
// The properties BaseAmount, Destination, Source are required.
type SwapQuoteRequestBody struct {
	// Amount in base units (e.g., wei for ETH). Must be a non-negative integer string.
	BaseAmount string `json:"base_amount" api:"required"`
	// The output side of a swap quote request.
	Destination SwapQuoteDestination `json:"destination,omitzero" api:"required"`
	// The input side of a swap request, including token and chain.
	Source SwapSource `json:"source,omitzero" api:"required"`
	// Maximum slippage tolerance in basis points (e.g., 50 for 0.5%). If omitted,
	// auto-slippage is used.
	SlippageBps param.Opt[float64] `json:"slippage_bps,omitzero"`
	// Whether the amount refers to the input token or output token.
	//
	// Any of "exact_input", "exact_output".
	AmountType AmountType `json:"amount_type,omitzero"`
	// Total fees assessed on a transfer, in BPS
	FeeConfiguration FeeConfiguration `json:"fee_configuration,omitzero"`
	paramObj
}

func (r SwapQuoteRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow SwapQuoteRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwapQuoteRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pricing data for a token swap.
type SwapQuoteResponse struct {
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// Estimated amount of output token in base units.
	EstOutputAmount string `json:"est_output_amount" api:"required"`
	// Estimated gas cost in base units of the native token.
	GasEstimate string `json:"gas_estimate" api:"required"`
	// Amount of input token in base units.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Minimum output amount accounting for slippage, in base units.
	MinimumOutputAmount string `json:"minimum_output_amount" api:"required"`
	// Token address being bought.
	OutputToken string `json:"output_token" api:"required"`
	// Destination chain CAIP-2 identifier for cross-chain swaps.
	DestinationCaip2 string `json:"destination_caip2"`
	// Estimated fees in USD.
	EstimatedFees []FeeLineItemUnion `json:"estimated_fees"`
	// Quote expiry as Unix timestamp (seconds). Present for cross-chain quotes.
	ExpiresAt float64 `json:"expires_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2               respjson.Field
		EstOutputAmount     respjson.Field
		GasEstimate         respjson.Field
		InputAmount         respjson.Field
		InputToken          respjson.Field
		MinimumOutputAmount respjson.Field
		OutputToken         respjson.Field
		DestinationCaip2    respjson.Field
		EstimatedFees       respjson.Field
		ExpiresAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwapQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *SwapQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for executing a token swap.
//
// The properties BaseAmount, Destination, Source are required.
type SwapRequestBody struct {
	// Amount in base units (e.g., wei for ETH). Must be a non-negative integer string.
	BaseAmount string `json:"base_amount" api:"required"`
	// The output side of a swap execution request.
	Destination SwapDestination `json:"destination,omitzero" api:"required"`
	// The input side of a swap request, including token and chain.
	Source SwapSource `json:"source,omitzero" api:"required"`
	// Maximum slippage tolerance in basis points (e.g., 50 for 0.5%).
	SlippageBps param.Opt[float64] `json:"slippage_bps,omitzero"`
	// Whether the amount refers to the input token or output token.
	//
	// Any of "exact_input", "exact_output".
	AmountType AmountType `json:"amount_type,omitzero"`
	// Total fees assessed on a transfer, in BPS
	FeeConfiguration FeeConfiguration `json:"fee_configuration,omitzero"`
	paramObj
}

func (r SwapRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow SwapRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwapRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The input side of a swap request, including token and chain.
//
// The properties AssetAddress, Caip2 are required.
type SwapSource struct {
	// Token contract address to sell, or "native" for the chain's native token.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier (e.g., "eip155:1").
	Caip2 string `json:"caip2" api:"required"`
	paramObj
}

func (r SwapSource) MarshalJSON() (data []byte, err error) {
	type shadow SwapSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwapSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
