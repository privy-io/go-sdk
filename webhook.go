// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEventUnion, error) {
	res := &UnsafeUnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Block metadata for a wallet transfer event.
type BlockInfo struct {
	// The block number.
	Number float64 `json:"number" api:"required"`
	// The block timestamp.
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlockInfo) RawJSON() string { return r.JSON.raw }
func (r *BlockInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bridge metadata for a crypto deposit via liquidation address.
type BridgeCryptoDepositMetadata struct {
	DrainID string `json:"drain_id" api:"required"`
	// The crypto address of the liquidation address that received the deposit.
	LiquidationAddress   string `json:"liquidation_address" api:"required"`
	LiquidationAddressID string `json:"liquidation_address_id" api:"required"`
	// Any of "liquidation_address".
	Method BridgeCryptoDepositMetadataMethod `json:"method" api:"required"`
	// The address that sent the deposit.
	SourceWalletAddress string `json:"source_wallet_address" api:"required"`
	// Any of "crypto_deposit".
	Type BridgeCryptoDepositMetadataType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DrainID              respjson.Field
		LiquidationAddress   respjson.Field
		LiquidationAddressID respjson.Field
		Method               respjson.Field
		SourceWalletAddress  respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeCryptoDepositMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeCryptoDepositMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeCryptoDepositMetadataMethod string

const (
	BridgeCryptoDepositMetadataMethodLiquidationAddress BridgeCryptoDepositMetadataMethod = "liquidation_address"
)

type BridgeCryptoDepositMetadataType string

const (
	BridgeCryptoDepositMetadataTypeCryptoDeposit BridgeCryptoDepositMetadataType = "crypto_deposit"
)

// Bridge metadata for a crypto deposit via transfer.
type BridgeCryptoTransferMetadata struct {
	// Any of "transfer".
	Method BridgeCryptoTransferMetadataMethod `json:"method" api:"required"`
	// The wallet address that sent the transfer.
	SourceWalletAddress string `json:"source_wallet_address" api:"required"`
	TransferID          string `json:"transfer_id" api:"required"`
	// Any of "crypto_deposit".
	Type BridgeCryptoTransferMetadataType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method              respjson.Field
		SourceWalletAddress respjson.Field
		TransferID          respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeCryptoTransferMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeCryptoTransferMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeCryptoTransferMetadataMethod string

const (
	BridgeCryptoTransferMetadataMethodTransfer BridgeCryptoTransferMetadataMethod = "transfer"
)

type BridgeCryptoTransferMetadataType string

const (
	BridgeCryptoTransferMetadataTypeCryptoDeposit BridgeCryptoTransferMetadataType = "crypto_deposit"
)

// Bridge metadata for a fiat deposit via virtual account.
type BridgeFiatDepositMetadata struct {
	ActivityID string `json:"activity_id" api:"required"`
	// Any of "virtual_account".
	Method BridgeFiatDepositMetadataMethod `json:"method" api:"required"`
	// Any of "fiat_deposit".
	Type             BridgeFiatDepositMetadataType `json:"type" api:"required"`
	VirtualAccountID string                        `json:"virtual_account_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivityID       respjson.Field
		Method           respjson.Field
		Type             respjson.Field
		VirtualAccountID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeFiatDepositMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeFiatDepositMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeFiatDepositMetadataMethod string

const (
	BridgeFiatDepositMetadataMethodVirtualAccount BridgeFiatDepositMetadataMethod = "virtual_account"
)

type BridgeFiatDepositMetadataType string

const (
	BridgeFiatDepositMetadataTypeFiatDeposit BridgeFiatDepositMetadataType = "fiat_deposit"
)

// Bridge metadata for a fiat deposit via transfer.
type BridgeFiatTransferMetadata struct {
	// Any of "transfer".
	Method     BridgeFiatTransferMetadataMethod `json:"method" api:"required"`
	TransferID string                           `json:"transfer_id" api:"required"`
	// Any of "fiat_deposit".
	Type BridgeFiatTransferMetadataType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		TransferID  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeFiatTransferMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeFiatTransferMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeFiatTransferMetadataMethod string

const (
	BridgeFiatTransferMetadataMethodTransfer BridgeFiatTransferMetadataMethod = "transfer"
)

type BridgeFiatTransferMetadataType string

const (
	BridgeFiatTransferMetadataTypeFiatDeposit BridgeFiatTransferMetadataType = "fiat_deposit"
)

// BridgeMetadataUnion contains all possible properties and values from
// [BridgeCryptoDepositMetadata], [BridgeRefundMetadata],
// [BridgeFiatDepositMetadata], [BridgeCryptoTransferMetadata],
// [BridgeFiatTransferMetadata], [BridgeTransferRefundMetadata],
// [BridgeStaticMemoDepositMetadata].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BridgeMetadataUnion struct {
	DrainID string `json:"drain_id"`
	// This field is from variant [BridgeCryptoDepositMetadata].
	LiquidationAddress      string `json:"liquidation_address"`
	LiquidationAddressID    string `json:"liquidation_address_id"`
	Method                  string `json:"method"`
	SourceWalletAddress     string `json:"source_wallet_address"`
	Type                    string `json:"type"`
	OriginalTransactionHash string `json:"original_transaction_hash"`
	// This field is from variant [BridgeFiatDepositMetadata].
	ActivityID string `json:"activity_id"`
	// This field is from variant [BridgeFiatDepositMetadata].
	VirtualAccountID string `json:"virtual_account_id"`
	TransferID       string `json:"transfer_id"`
	// This field is from variant [BridgeStaticMemoDepositMetadata].
	StaticMemoEventID string `json:"static_memo_event_id"`
	// This field is from variant [BridgeStaticMemoDepositMetadata].
	StaticMemoID string `json:"static_memo_id"`
	JSON         struct {
		DrainID                 respjson.Field
		LiquidationAddress      respjson.Field
		LiquidationAddressID    respjson.Field
		Method                  respjson.Field
		SourceWalletAddress     respjson.Field
		Type                    respjson.Field
		OriginalTransactionHash respjson.Field
		ActivityID              respjson.Field
		VirtualAccountID        respjson.Field
		TransferID              respjson.Field
		StaticMemoEventID       respjson.Field
		StaticMemoID            respjson.Field
		raw                     string
	} `json:"-"`
}

func (u BridgeMetadataUnion) AsBridgeCryptoDepositMetadata() (v BridgeCryptoDepositMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BridgeMetadataUnion) AsBridgeRefundMetadata() (v BridgeRefundMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BridgeMetadataUnion) AsBridgeFiatDepositMetadata() (v BridgeFiatDepositMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BridgeMetadataUnion) AsBridgeCryptoTransferMetadata() (v BridgeCryptoTransferMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BridgeMetadataUnion) AsBridgeFiatTransferMetadata() (v BridgeFiatTransferMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BridgeMetadataUnion) AsBridgeTransferRefundMetadata() (v BridgeTransferRefundMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BridgeMetadataUnion) AsBridgeStaticMemoDepositMetadata() (v BridgeStaticMemoDepositMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BridgeMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BridgeMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bridge metadata for a refund via liquidation address.
type BridgeRefundMetadata struct {
	DrainID              string `json:"drain_id" api:"required"`
	LiquidationAddressID string `json:"liquidation_address_id" api:"required"`
	// Any of "liquidation_address".
	Method BridgeRefundMetadataMethod `json:"method" api:"required"`
	// The original deposit transaction hash that triggered the failed drain.
	OriginalTransactionHash string `json:"original_transaction_hash" api:"required"`
	// Any of "refund".
	Type BridgeRefundMetadataType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DrainID                 respjson.Field
		LiquidationAddressID    respjson.Field
		Method                  respjson.Field
		OriginalTransactionHash respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeRefundMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeRefundMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeRefundMetadataMethod string

const (
	BridgeRefundMetadataMethodLiquidationAddress BridgeRefundMetadataMethod = "liquidation_address"
)

type BridgeRefundMetadataType string

const (
	BridgeRefundMetadataTypeRefund BridgeRefundMetadataType = "refund"
)

// Bridge metadata for a fiat deposit via static memo.
type BridgeStaticMemoDepositMetadata struct {
	// Any of "static_memo".
	Method            BridgeStaticMemoDepositMetadataMethod `json:"method" api:"required"`
	StaticMemoEventID string                                `json:"static_memo_event_id" api:"required"`
	StaticMemoID      string                                `json:"static_memo_id" api:"required"`
	// Any of "fiat_deposit".
	Type BridgeStaticMemoDepositMetadataType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method            respjson.Field
		StaticMemoEventID respjson.Field
		StaticMemoID      respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeStaticMemoDepositMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeStaticMemoDepositMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeStaticMemoDepositMetadataMethod string

const (
	BridgeStaticMemoDepositMetadataMethodStaticMemo BridgeStaticMemoDepositMetadataMethod = "static_memo"
)

type BridgeStaticMemoDepositMetadataType string

const (
	BridgeStaticMemoDepositMetadataTypeFiatDeposit BridgeStaticMemoDepositMetadataType = "fiat_deposit"
)

// Bridge metadata for a transfer refund.
type BridgeTransferRefundMetadata struct {
	// Any of "transfer".
	Method     BridgeTransferRefundMetadataMethod `json:"method" api:"required"`
	TransferID string                             `json:"transfer_id" api:"required"`
	// Any of "refund".
	Type BridgeTransferRefundMetadataType `json:"type" api:"required"`
	// The original transfer transaction hash (if available).
	OriginalTransactionHash string `json:"original_transaction_hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method                  respjson.Field
		TransferID              respjson.Field
		Type                    respjson.Field
		OriginalTransactionHash respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeTransferRefundMetadata) RawJSON() string { return r.JSON.raw }
func (r *BridgeTransferRefundMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeTransferRefundMetadataMethod string

const (
	BridgeTransferRefundMetadataMethodTransfer BridgeTransferRefundMetadataMethod = "transfer"
)

type BridgeTransferRefundMetadataType string

const (
	BridgeTransferRefundMetadataTypeRefund BridgeTransferRefundMetadataType = "refund"
)

// Details of a fiat deposit that has finished converting and been delivered to the
// wallet.
type DepositCompletedData struct {
	CreatedAt string `json:"created_at" api:"required"`
	// The crypto asset, chain, delivered amount, and settlement transaction for a
	// completed deposit.
	Destination DepositCompletedDestination `json:"destination" api:"required"`
	// The fiat deposit that was received, including amount, currency, and originator.
	Source DepositStartedSource `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Destination respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DepositCompletedData) RawJSON() string { return r.JSON.raw }
func (r *DepositCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The crypto asset, chain, delivered amount, and settlement transaction for a
// completed deposit.
type DepositCompletedDestination struct {
	// The crypto amount delivered to the wallet, after conversion and fees.
	Amount string `json:"amount" api:"required"`
	// The crypto asset the deposit was converted into (e.g. "usdc").
	Asset string `json:"asset" api:"required"`
	// The chain the converted crypto was delivered on (e.g. "base").
	Chain string `json:"chain" api:"required"`
	// The on-chain settlement transaction for the delivered crypto.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		Asset           respjson.Field
		Chain           respjson.Field
		TransactionHash respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DepositCompletedDestination) RawJSON() string { return r.JSON.raw }
func (r *DepositCompletedDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of a fiat deposit that failed to convert and was refunded to the sender.
type DepositFailedData struct {
	CreatedAt string `json:"created_at" api:"required"`
	// The crypto asset and chain the fiat deposit is being converted into.
	Destination DepositStartedDestination `json:"destination" api:"required"`
	Reason      string                    `json:"reason" api:"required"`
	ReasonCode  string                    `json:"reason_code" api:"required"`
	RefundedAt  string                    `json:"refunded_at" api:"required"`
	// The fiat deposit that was received, including amount, currency, and originator.
	Source DepositStartedSource `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Destination respjson.Field
		Reason      respjson.Field
		ReasonCode  respjson.Field
		RefundedAt  respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DepositFailedData) RawJSON() string { return r.JSON.raw }
func (r *DepositFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of a fiat deposit that has begun processing into a deposit account.
type DepositStartedData struct {
	CreatedAt string `json:"created_at" api:"required"`
	// The crypto asset and chain the fiat deposit is being converted into.
	Destination DepositStartedDestination `json:"destination" api:"required"`
	// The fiat deposit that was received, including amount, currency, and originator.
	Source DepositStartedSource `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Destination respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DepositStartedData) RawJSON() string { return r.JSON.raw }
func (r *DepositStartedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The crypto asset and chain the fiat deposit is being converted into.
type DepositStartedDestination struct {
	// The crypto asset the deposit is converted into (e.g. "usdc").
	Asset string `json:"asset" api:"required"`
	// The chain the converted crypto is delivered on (e.g. "base").
	Chain string `json:"chain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		Chain       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DepositStartedDestination) RawJSON() string { return r.JSON.raw }
func (r *DepositStartedDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The fiat deposit that was received, including amount, currency, and originator.
type DepositStartedSource struct {
	// The fiat amount deposited.
	Amount string `json:"amount" api:"required"`
	// Supported fiat currencies.
	//
	// Any of "usd", "eur".
	Currency FiatCurrency `json:"currency" api:"required"`
	// Supported fiat payment rails.
	//
	// Any of "sepa", "ach_push", "wire", "fednow", "faster_payments".
	PaymentRail FiatPaymentRail `json:"payment_rail"`
	SenderName  string          `json:"sender_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		PaymentRail respjson.Field
		SenderName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DepositStartedSource) RawJSON() string { return r.JSON.raw }
func (r *DepositStartedSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the wallet.funds_deposited webhook event.
type FundsDepositedWebhookPayload struct {
	// The amount transferred, as a stringified bigint.
	Amount string `json:"amount" api:"required"`
	// An asset involved in a wallet transfer.
	Asset WalletFundsAssetUnion `json:"asset" api:"required"`
	// Block metadata for a wallet transfer event.
	Block BlockInfo `json:"block" api:"required"`
	// The CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// A unique key for this event.
	IdempotencyKey string `json:"idempotency_key" api:"required"`
	// The recipient address.
	Recipient string `json:"recipient" api:"required"`
	// The sender address.
	Sender string `json:"sender" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.funds_deposited".
	Type FundsDepositedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// Metadata about a Bridge transaction associated with a wallet event.
	BridgeMetadata BridgeMetadataUnion `json:"bridge_metadata"`
	// The transaction fee paid, as a stringified bigint in the chain's native token.
	TransactionFee string `json:"transaction_fee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		Asset           respjson.Field
		Block           respjson.Field
		Caip2           respjson.Field
		IdempotencyKey  respjson.Field
		Recipient       respjson.Field
		Sender          respjson.Field
		TransactionHash respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		BridgeMetadata  respjson.Field
		TransactionFee  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundsDepositedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *FundsDepositedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type FundsDepositedWebhookPayloadType string

const (
	FundsDepositedWebhookPayloadTypeWalletFundsDeposited FundsDepositedWebhookPayloadType = "wallet.funds_deposited"
)

// Payload for the wallet.funds_withdrawn webhook event.
type FundsWithdrawnWebhookPayload struct {
	// The amount transferred, as a stringified bigint.
	Amount string `json:"amount" api:"required"`
	// An asset involved in a wallet transfer.
	Asset WalletFundsAssetUnion `json:"asset" api:"required"`
	// Block metadata for a wallet transfer event.
	Block BlockInfo `json:"block" api:"required"`
	// The CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// A unique key for this event.
	IdempotencyKey string `json:"idempotency_key" api:"required"`
	// The recipient address.
	Recipient string `json:"recipient" api:"required"`
	// The sender address.
	Sender string `json:"sender" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.funds_withdrawn".
	Type FundsWithdrawnWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// The transaction fee paid, as a stringified bigint in the chain's native token.
	TransactionFee string `json:"transaction_fee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		Asset           respjson.Field
		Block           respjson.Field
		Caip2           respjson.Field
		IdempotencyKey  respjson.Field
		Recipient       respjson.Field
		Sender          respjson.Field
		TransactionHash respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		TransactionFee  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FundsWithdrawnWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *FundsWithdrawnWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type FundsWithdrawnWebhookPayloadType string

const (
	FundsWithdrawnWebhookPayloadTypeWalletFundsWithdrawn FundsWithdrawnWebhookPayloadType = "wallet.funds_withdrawn"
)

// Payload for the intent.authorized webhook event.
type IntentAuthorizedWebhookPayload struct {
	// Unix timestamp when the authorization was recorded.
	AuthorizedAt float64 `json:"authorized_at" api:"required"`
	// Unix timestamp when the intent was created.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Unix timestamp when the intent expires.
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// The unique ID of the intent.
	IntentID string `json:"intent_id" api:"required"`
	// Type of intent.
	//
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "TRANSFER", "WALLET".
	IntentType IntentType `json:"intent_type" api:"required"`
	// A leaf member (user or key) of a nested key quorum in an intent authorization.
	Member IntentAuthorizationKeyQuorumMemberUnion `json:"member" api:"required"`
	// The current status of the intent.
	Status string `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "intent.authorized".
	Type IntentAuthorizedWebhookPayloadType `json:"type" api:"required"`
	// Display name of the user who created the intent.
	CreatedByDisplayName string `json:"created_by_display_name"`
	// The ID of the user who created the intent.
	CreatedByID string `json:"created_by_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizedAt         respjson.Field
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		Member               respjson.Field
		Status               respjson.Field
		Type                 respjson.Field
		CreatedByDisplayName respjson.Field
		CreatedByID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorizedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorizedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type IntentAuthorizedWebhookPayloadType string

const (
	IntentAuthorizedWebhookPayloadTypeIntentAuthorized IntentAuthorizedWebhookPayloadType = "intent.authorized"
)

// Payload for the intent.created webhook event.
type IntentCreatedWebhookPayload struct {
	// Unix timestamp when the intent was created.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Unix timestamp when the intent expires.
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// The unique ID of the intent.
	IntentID string `json:"intent_id" api:"required"`
	// Type of intent.
	//
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "TRANSFER", "WALLET".
	IntentType IntentType `json:"intent_type" api:"required"`
	// The current status of the intent.
	Status string `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "intent.created".
	Type IntentCreatedWebhookPayloadType `json:"type" api:"required"`
	// Key quorums that can authorize this intent.
	AuthorizationDetails []IntentAuthorization `json:"authorization_details"`
	// Display name of the user who created the intent.
	CreatedByDisplayName string `json:"created_by_display_name"`
	// The ID of the user who created the intent.
	CreatedByID string `json:"created_by_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		Status               respjson.Field
		Type                 respjson.Field
		AuthorizationDetails respjson.Field
		CreatedByDisplayName respjson.Field
		CreatedByID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *IntentCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type IntentCreatedWebhookPayloadType string

const (
	IntentCreatedWebhookPayloadTypeIntentCreated IntentCreatedWebhookPayloadType = "intent.created"
)

// Payload for the intent.executed webhook event.
type IntentExecutedWebhookPayload struct {
	// Result of the successful intent execution.
	ActionResult BaseActionResult `json:"action_result" api:"required"`
	// Unix timestamp when the intent was created.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Unix timestamp when the intent expires.
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// The unique ID of the intent.
	IntentID string `json:"intent_id" api:"required"`
	// Type of intent.
	//
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "TRANSFER", "WALLET".
	IntentType IntentType `json:"intent_type" api:"required"`
	// The current status of the intent.
	Status string `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "intent.executed".
	Type IntentExecutedWebhookPayloadType `json:"type" api:"required"`
	// Display name of the user who created the intent.
	CreatedByDisplayName string `json:"created_by_display_name"`
	// The ID of the user who created the intent.
	CreatedByID string `json:"created_by_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionResult         respjson.Field
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		Status               respjson.Field
		Type                 respjson.Field
		CreatedByDisplayName respjson.Field
		CreatedByID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentExecutedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *IntentExecutedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type IntentExecutedWebhookPayloadType string

const (
	IntentExecutedWebhookPayloadTypeIntentExecuted IntentExecutedWebhookPayloadType = "intent.executed"
)

// Payload for the intent.failed webhook event.
type IntentFailedWebhookPayload struct {
	// Result of the failed intent execution.
	ActionResult BaseActionResult `json:"action_result" api:"required"`
	// Unix timestamp when the intent was created.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Unix timestamp when the intent expires.
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// The unique ID of the intent.
	IntentID string `json:"intent_id" api:"required"`
	// Type of intent.
	//
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "TRANSFER", "WALLET".
	IntentType IntentType `json:"intent_type" api:"required"`
	// The current status of the intent.
	Status string `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "intent.failed".
	Type IntentFailedWebhookPayloadType `json:"type" api:"required"`
	// Display name of the user who created the intent.
	CreatedByDisplayName string `json:"created_by_display_name"`
	// The ID of the user who created the intent.
	CreatedByID string `json:"created_by_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionResult         respjson.Field
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		Status               respjson.Field
		Type                 respjson.Field
		CreatedByDisplayName respjson.Field
		CreatedByID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *IntentFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type IntentFailedWebhookPayloadType string

const (
	IntentFailedWebhookPayloadTypeIntentFailed IntentFailedWebhookPayloadType = "intent.failed"
)

// Payload for the intent.rejected webhook event.
type IntentRejectedWebhookPayload struct {
	// Unix timestamp when the intent was created.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Unix timestamp when the intent expires.
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// The unique ID of the intent.
	IntentID string `json:"intent_id" api:"required"`
	// Type of intent.
	//
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "TRANSFER", "WALLET".
	IntentType IntentType `json:"intent_type" api:"required"`
	// Unix timestamp when the intent was rejected.
	RejectedAt float64 `json:"rejected_at" api:"required"`
	// The current status of the intent.
	Status string `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "intent.rejected".
	Type IntentRejectedWebhookPayloadType `json:"type" api:"required"`
	// Display name of the user who created the intent.
	CreatedByDisplayName string `json:"created_by_display_name"`
	// The ID of the user who created the intent.
	CreatedByID string `json:"created_by_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RejectedAt           respjson.Field
		Status               respjson.Field
		Type                 respjson.Field
		CreatedByDisplayName respjson.Field
		CreatedByID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *IntentRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type IntentRejectedWebhookPayloadType string

const (
	IntentRejectedWebhookPayloadTypeIntentRejected IntentRejectedWebhookPayloadType = "intent.rejected"
)

// Payload for the mfa.disabled webhook event.
type MfaDisabledWebhookPayload struct {
	// A multi-factor authentication method supported by the app.
	//
	// Any of "sms", "totp", "passkey", "email".
	Method MfaMethod `json:"method" api:"required"`
	// The type of webhook event.
	//
	// Any of "mfa.disabled".
	Type MfaDisabledWebhookPayloadType `json:"type" api:"required"`
	// The ID of the user who disabled MFA.
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaDisabledWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *MfaDisabledWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type MfaDisabledWebhookPayloadType string

const (
	MfaDisabledWebhookPayloadTypeMfaDisabled MfaDisabledWebhookPayloadType = "mfa.disabled"
)

// Payload for the mfa.enabled webhook event.
type MfaEnabledWebhookPayload struct {
	// A multi-factor authentication method supported by the app.
	//
	// Any of "sms", "totp", "passkey", "email".
	Method MfaMethod `json:"method" api:"required"`
	// The type of webhook event.
	//
	// Any of "mfa.enabled".
	Type MfaEnabledWebhookPayloadType `json:"type" api:"required"`
	// The ID of the user who enabled MFA.
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaEnabledWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *MfaEnabledWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type MfaEnabledWebhookPayloadType string

const (
	MfaEnabledWebhookPayloadTypeMfaEnabled MfaEnabledWebhookPayloadType = "mfa.enabled"
)

// Full KYB state snapshot in a KYB update event.
type OrganizationKYBUpdatedData struct {
	// Capability statuses for the customer.
	Capabilities KyxCapabilities  `json:"capabilities" api:"required"`
	Endorsements []KyxEndorsement `json:"endorsements" api:"required"`
	// KYB verification status in a KYB update event.
	KYB OrganizationKYBUpdatedKYBData `json:"kyb" api:"required"`
	// KYC/KYB status for the user.
	Status KyxProviderStatus `json:"status" api:"required"`
	// Terms of service status in a KYB update event.
	Tos OrganizationKYBUpdatedTosData `json:"tos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities respjson.Field
		Endorsements respjson.Field
		KYB          respjson.Field
		Status       respjson.Field
		Tos          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationKYBUpdatedData) RawJSON() string { return r.JSON.raw }
func (r *OrganizationKYBUpdatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYB verification status in a KYB update event.
type OrganizationKYBUpdatedKYBData struct {
	// Status of KYC/KYB verification. Passthrough from the provider.
	Status KyxVerificationStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationKYBUpdatedKYBData) RawJSON() string { return r.JSON.raw }
func (r *OrganizationKYBUpdatedKYBData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terms of service status in a KYB update event.
type OrganizationKYBUpdatedTosData struct {
	// Status of Terms of Service acceptance. Passthrough from the provider.
	Status KyxTosStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationKYBUpdatedTosData) RawJSON() string { return r.JSON.raw }
func (r *OrganizationKYBUpdatedTosData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the wallet.private_key_export webhook event.
type PrivateKeyExportWebhookPayload struct {
	// The type of webhook event.
	//
	// Any of "wallet.private_key_export".
	Type PrivateKeyExportWebhookPayloadType `json:"type" api:"required"`
	// The ID of the user who exported the key.
	UserID string `json:"user_id" api:"required"`
	// The address of the wallet.
	WalletAddress string `json:"wallet_address" api:"required"`
	// The ID of the wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// The export type. 'display' is for showing the key to the user in the UI,
	// 'client' is for exporting to the client application.
	//
	// Any of "display", "client".
	ExportSource ExportType `json:"export_source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		UserID        respjson.Field
		WalletAddress respjson.Field
		WalletID      respjson.Field
		ExportSource  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateKeyExportWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *PrivateKeyExportWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type PrivateKeyExportWebhookPayloadType string

const (
	PrivateKeyExportWebhookPayloadTypeWalletPrivateKeyExport PrivateKeyExportWebhookPayloadType = "wallet.private_key_export"
)

// Payload for the transaction.broadcasted webhook event.
type TransactionBroadcastedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.broadcasted".
	Type TransactionBroadcastedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		TransactionHash respjson.Field
		TransactionID   respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		ReferenceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionBroadcastedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionBroadcastedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionBroadcastedWebhookPayloadType string

const (
	TransactionBroadcastedWebhookPayloadTypeTransactionBroadcasted TransactionBroadcastedWebhookPayloadType = "transaction.broadcasted"
)

// Payload for the transaction.confirmed webhook event.
type TransactionConfirmedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.confirmed".
	Type TransactionConfirmedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		TransactionHash respjson.Field
		TransactionID   respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		ReferenceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionConfirmedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionConfirmedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionConfirmedWebhookPayloadType string

const (
	TransactionConfirmedWebhookPayloadTypeTransactionConfirmed TransactionConfirmedWebhookPayloadType = "transaction.confirmed"
)

// Payload for the transaction.execution_reverted webhook event.
type TransactionExecutionRevertedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.execution_reverted".
	Type TransactionExecutionRevertedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		TransactionHash respjson.Field
		TransactionID   respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		ReferenceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionExecutionRevertedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionExecutionRevertedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionExecutionRevertedWebhookPayloadType string

const (
	TransactionExecutionRevertedWebhookPayloadTypeTransactionExecutionReverted TransactionExecutionRevertedWebhookPayloadType = "transaction.execution_reverted"
)

// Payload for the transaction.failed webhook event.
type TransactionFailedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.failed".
	Type TransactionFailedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		TransactionHash respjson.Field
		TransactionID   respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		ReferenceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionFailedWebhookPayloadType string

const (
	TransactionFailedWebhookPayloadTypeTransactionFailed TransactionFailedWebhookPayloadType = "transaction.failed"
)

// Payload for the transaction.provider_error webhook event.
type TransactionProviderErrorWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.provider_error".
	Type TransactionProviderErrorWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		TransactionHash respjson.Field
		TransactionID   respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		ReferenceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionProviderErrorWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionProviderErrorWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionProviderErrorWebhookPayloadType string

const (
	TransactionProviderErrorWebhookPayloadTypeTransactionProviderError TransactionProviderErrorWebhookPayloadType = "transaction.provider_error"
)

// Payload for the transaction.replaced webhook event.
type TransactionReplacedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.replaced".
	Type TransactionReplacedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		TransactionHash respjson.Field
		TransactionID   respjson.Field
		Type            respjson.Field
		WalletID        respjson.Field
		ReferenceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionReplacedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionReplacedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionReplacedWebhookPayloadType string

const (
	TransactionReplacedWebhookPayloadTypeTransactionReplaced TransactionReplacedWebhookPayloadType = "transaction.replaced"
)

// Payload for the transaction.still_pending webhook event.
type TransactionStillPendingWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:4217 for Tempo, eip155:1 for Ethereum
	// mainnet).
	Caip2 string `json:"caip2" api:"required"`
	// The blockchain transaction hash.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// The Privy-assigned ID for this transaction.
	TransactionID string `json:"transaction_id" api:"required"`
	// An unsigned standard Ethereum transaction object. Supports EVM transaction types
	// 0, 1, 2, and 4.
	TransactionRequest UnsignedStandardEthereumTransactionResp `json:"transaction_request" api:"required"`
	// The type of webhook event.
	//
	// Any of "transaction.still_pending".
	Type TransactionStillPendingWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet that initiated the transaction.
	WalletID string `json:"wallet_id" api:"required"`
	// Developer-provided reference ID for transaction reconciliation, if one was
	// provided.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2              respjson.Field
		TransactionHash    respjson.Field
		TransactionID      respjson.Field
		TransactionRequest respjson.Field
		Type               respjson.Field
		WalletID           respjson.Field
		ReferenceID        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionStillPendingWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *TransactionStillPendingWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type TransactionStillPendingWebhookPayloadType string

const (
	TransactionStillPendingWebhookPayloadTypeTransactionStillPending TransactionStillPendingWebhookPayloadType = "transaction.still_pending"
)

// Payload for the usage.cross_chain_fee.recorded webhook event (Privy fee on a
// cross-chain transfer or swap).
type UsageCrossChainFeeRecordedWebhookPayload struct {
	AmountUsd string `json:"amount_usd" api:"required"`
	// An opaque, stable identifier for this charge. Use it to deduplicate webhook
	// deliveries.
	EventID    string `json:"event_id" api:"required"`
	RecordedAt int64  `json:"recorded_at" api:"required"`
	SourceID   string `json:"source_id" api:"required"`
	// The type of operation that incurred a usage charge.
	//
	// Any of "wallet-action-transfer", "wallet-action-swap", "rpc".
	SourceType UsageSourceType `json:"source_type" api:"required"`
	// The type of webhook event.
	//
	// Any of "usage.cross_chain_fee.recorded".
	Type UsageCrossChainFeeRecordedWebhookPayloadType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountUsd   respjson.Field
		EventID     respjson.Field
		RecordedAt  respjson.Field
		SourceID    respjson.Field
		SourceType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageCrossChainFeeRecordedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UsageCrossChainFeeRecordedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UsageCrossChainFeeRecordedWebhookPayloadType string

const (
	UsageCrossChainFeeRecordedWebhookPayloadTypeUsageCrossChainFeeRecorded UsageCrossChainFeeRecordedWebhookPayloadType = "usage.cross_chain_fee.recorded"
)

// Payload for the usage.gas_sponsorship.recorded webhook event (sponsored network
// gas).
type UsageGasSponsorshipRecordedWebhookPayload struct {
	AmountUsd string `json:"amount_usd" api:"required"`
	// An opaque, stable identifier for this charge. Use it to deduplicate webhook
	// deliveries.
	EventID    string `json:"event_id" api:"required"`
	RecordedAt int64  `json:"recorded_at" api:"required"`
	SourceID   string `json:"source_id" api:"required"`
	// The type of operation that incurred a usage charge.
	//
	// Any of "wallet-action-transfer", "wallet-action-swap", "rpc".
	SourceType UsageSourceType `json:"source_type" api:"required"`
	// The type of webhook event.
	//
	// Any of "usage.gas_sponsorship.recorded".
	Type UsageGasSponsorshipRecordedWebhookPayloadType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountUsd   respjson.Field
		EventID     respjson.Field
		RecordedAt  respjson.Field
		SourceID    respjson.Field
		SourceType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGasSponsorshipRecordedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UsageGasSponsorshipRecordedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UsageGasSponsorshipRecordedWebhookPayloadType string

const (
	UsageGasSponsorshipRecordedWebhookPayloadTypeUsageGasSponsorshipRecorded UsageGasSponsorshipRecordedWebhookPayloadType = "usage.gas_sponsorship.recorded"
)

// The type of operation that incurred a usage charge.
type UsageSourceType string

const (
	UsageSourceTypeWalletActionTransfer UsageSourceType = "wallet-action-transfer"
	UsageSourceTypeWalletActionSwap     UsageSourceType = "wallet-action-swap"
	UsageSourceTypeRpc                  UsageSourceType = "rpc"
)

// Payload for the user.authenticated webhook event.
type UserAuthenticatedWebhookPayload struct {
	// A linked account for the user.
	Account LinkedAccountUnion `json:"account" api:"required"`
	// The type of webhook event.
	//
	// Any of "user.authenticated".
	Type UserAuthenticatedWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAuthenticatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserAuthenticatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserAuthenticatedWebhookPayloadType string

const (
	UserAuthenticatedWebhookPayloadTypeUserAuthenticated UserAuthenticatedWebhookPayloadType = "user.authenticated"
)

// Payload for the user.created webhook event.
type UserCreatedWebhookPayload struct {
	// The type of webhook event.
	//
	// Any of "user.created".
	Type UserCreatedWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserCreatedWebhookPayloadType string

const (
	UserCreatedWebhookPayloadTypeUserCreated UserCreatedWebhookPayloadType = "user.created"
)

// Payload for the user.deleted webhook event.
type UserDeletedWebhookPayload struct {
	// The type of webhook event.
	//
	// Any of "user.deleted".
	Type UserDeletedWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserDeletedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserDeletedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserDeletedWebhookPayloadType string

const (
	UserDeletedWebhookPayloadTypeUserDeleted UserDeletedWebhookPayloadType = "user.deleted"
)

// Full KYC state snapshot in a KYC update event.
type UserKYCUpdatedData struct {
	// Capability statuses for the customer.
	Capabilities KyxCapabilities  `json:"capabilities" api:"required"`
	Endorsements []KyxEndorsement `json:"endorsements" api:"required"`
	// KYC verification status in a KYC update event.
	KYC UserKYCUpdatedKYCData `json:"kyc" api:"required"`
	// KYC/KYB status for the user.
	Status KyxProviderStatus `json:"status" api:"required"`
	// Terms of service status in a KYC update event.
	Tos UserKYCUpdatedTosData `json:"tos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities respjson.Field
		Endorsements respjson.Field
		KYC          respjson.Field
		Status       respjson.Field
		Tos          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserKYCUpdatedData) RawJSON() string { return r.JSON.raw }
func (r *UserKYCUpdatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC verification status in a KYC update event.
type UserKYCUpdatedKYCData struct {
	// Status of KYC/KYB verification. Passthrough from the provider.
	Status KyxVerificationStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserKYCUpdatedKYCData) RawJSON() string { return r.JSON.raw }
func (r *UserKYCUpdatedKYCData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terms of service status in a KYC update event.
type UserKYCUpdatedTosData struct {
	// Status of Terms of Service acceptance. Passthrough from the provider.
	Status KyxTosStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserKYCUpdatedTosData) RawJSON() string { return r.JSON.raw }
func (r *UserKYCUpdatedTosData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the user.linked_account webhook event.
type UserLinkedAccountWebhookPayload struct {
	// A linked account for the user.
	Account LinkedAccountUnion `json:"account" api:"required"`
	// The type of webhook event.
	//
	// Any of "user.linked_account".
	Type UserLinkedAccountWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserLinkedAccountWebhookPayloadType string

const (
	UserLinkedAccountWebhookPayloadTypeUserLinkedAccount UserLinkedAccountWebhookPayloadType = "user.linked_account"
)

// Payload for the user_operation.completed webhook event.
type UserOperationCompletedWebhookPayload struct {
	ActualGasCost   string  `json:"actual_gas_cost" api:"required"`
	ActualGasUsed   string  `json:"actual_gas_used" api:"required"`
	BlockNumber     float64 `json:"block_number" api:"required"`
	Caip2           string  `json:"caip2" api:"required"`
	LogIndex        float64 `json:"log_index" api:"required"`
	Nonce           string  `json:"nonce" api:"required"`
	Paymaster       string  `json:"paymaster" api:"required"`
	Sender          string  `json:"sender" api:"required"`
	Success         bool    `json:"success" api:"required"`
	TransactionHash string  `json:"transaction_hash" api:"required"`
	// The type of webhook event.
	//
	// Any of "user_operation.completed".
	Type       UserOperationCompletedWebhookPayloadType `json:"type" api:"required"`
	UserOpHash string                                   `json:"user_op_hash" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualGasCost   respjson.Field
		ActualGasUsed   respjson.Field
		BlockNumber     respjson.Field
		Caip2           respjson.Field
		LogIndex        respjson.Field
		Nonce           respjson.Field
		Paymaster       respjson.Field
		Sender          respjson.Field
		Success         respjson.Field
		TransactionHash respjson.Field
		Type            respjson.Field
		UserOpHash      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserOperationCompletedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserOperationCompletedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserOperationCompletedWebhookPayloadType string

const (
	UserOperationCompletedWebhookPayloadTypeUserOperationCompleted UserOperationCompletedWebhookPayloadType = "user_operation.completed"
)

// A reference to a user by their unique identifier.
type UserReference struct {
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserReference) RawJSON() string { return r.JSON.raw }
func (r *UserReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the user.transferred_account webhook event.
type UserTransferredAccountWebhookPayload struct {
	// A linked account for the user.
	Account LinkedAccountUnion `json:"account" api:"required"`
	// Any of true.
	DeletedUser bool `json:"deletedUser" api:"required"`
	// A reference to a user by their unique identifier.
	FromUser UserReference `json:"fromUser" api:"required"`
	// A Privy user object.
	ToUser User `json:"toUser" api:"required"`
	// The type of webhook event.
	//
	// Any of "user.transferred_account".
	Type UserTransferredAccountWebhookPayloadType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		DeletedUser respjson.Field
		FromUser    respjson.Field
		ToUser      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTransferredAccountWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserTransferredAccountWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserTransferredAccountWebhookPayloadType string

const (
	UserTransferredAccountWebhookPayloadTypeUserTransferredAccount UserTransferredAccountWebhookPayloadType = "user.transferred_account"
)

// Payload for the user.unlinked_account webhook event.
type UserUnlinkedAccountWebhookPayload struct {
	// A linked account for the user.
	Account LinkedAccountUnion `json:"account" api:"required"`
	// The type of webhook event.
	//
	// Any of "user.unlinked_account".
	Type UserUnlinkedAccountWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUnlinkedAccountWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserUnlinkedAccountWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserUnlinkedAccountWebhookPayloadType string

const (
	UserUnlinkedAccountWebhookPayloadTypeUserUnlinkedAccount UserUnlinkedAccountWebhookPayloadType = "user.unlinked_account"
)

// Payload for the user.updated_account webhook event.
type UserUpdatedAccountWebhookPayload struct {
	// A linked account for the user.
	Account LinkedAccountUnion `json:"account" api:"required"`
	// The type of webhook event.
	//
	// Any of "user.updated_account".
	Type UserUpdatedAccountWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUpdatedAccountWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserUpdatedAccountWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserUpdatedAccountWebhookPayloadType string

const (
	UserUpdatedAccountWebhookPayloadTypeUserUpdatedAccount UserUpdatedAccountWebhookPayloadType = "user.updated_account"
)

// Payload for the user.wallet_created webhook event.
type UserWalletCreatedWebhookPayload struct {
	// The type of webhook event.
	//
	// Any of "user.wallet_created".
	Type UserWalletCreatedWebhookPayloadType `json:"type" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// Base schema for wallet accounts linked to the user.
	Wallet LinkedAccountBaseWallet `json:"wallet" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		User        respjson.Field
		Wallet      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserWalletCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *UserWalletCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserWalletCreatedWebhookPayloadType string

const (
	UserWalletCreatedWebhookPayloadTypeUserWalletCreated UserWalletCreatedWebhookPayloadType = "user.wallet_created"
)

// Payload for the wallet_action.earn_deposit.created webhook event.
type WalletActionEarnDepositCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Base-unit amount of asset deposited (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionEarnDepositCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_deposit.created".
	Type WalletActionEarnDepositCreatedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset deposited (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnDepositCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnDepositCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnDepositCreatedWebhookPayloadStatus string

const (
	WalletActionEarnDepositCreatedWebhookPayloadStatusPending WalletActionEarnDepositCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionEarnDepositCreatedWebhookPayloadType string

const (
	WalletActionEarnDepositCreatedWebhookPayloadTypeWalletActionEarnDepositCreated WalletActionEarnDepositCreatedWebhookPayloadType = "wallet_action.earn_deposit.created"
)

// Payload for the wallet_action.earn_deposit.failed webhook event.
type WalletActionEarnDepositFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of asset deposited (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionEarnDepositFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_deposit.failed".
	Type WalletActionEarnDepositFailedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset deposited (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailedAt       respjson.Field
		FailureReason  respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnDepositFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnDepositFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnDepositFailedWebhookPayloadStatus string

const (
	WalletActionEarnDepositFailedWebhookPayloadStatusFailed WalletActionEarnDepositFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionEarnDepositFailedWebhookPayloadType string

const (
	WalletActionEarnDepositFailedWebhookPayloadTypeWalletActionEarnDepositFailed WalletActionEarnDepositFailedWebhookPayloadType = "wallet_action.earn_deposit.failed"
)

// Payload for the wallet_action.earn_deposit.rejected webhook event.
type WalletActionEarnDepositRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of asset deposited (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionEarnDepositRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_deposit.rejected".
	Type WalletActionEarnDepositRejectedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset deposited (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailureReason  respjson.Field
		RawAmount      respjson.Field
		RejectedAt     respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnDepositRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnDepositRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnDepositRejectedWebhookPayloadStatus string

const (
	WalletActionEarnDepositRejectedWebhookPayloadStatusRejected WalletActionEarnDepositRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionEarnDepositRejectedWebhookPayloadType string

const (
	WalletActionEarnDepositRejectedWebhookPayloadTypeWalletActionEarnDepositRejected WalletActionEarnDepositRejectedWebhookPayloadType = "wallet_action.earn_deposit.rejected"
)

// Payload for the wallet_action.earn_deposit.succeeded webhook event.
type WalletActionEarnDepositSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Base-unit amount of asset deposited (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// Vault shares received in base units.
	ShareAmount string `json:"share_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionEarnDepositSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_deposit.succeeded".
	Type WalletActionEarnDepositSucceededWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset deposited (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		RawAmount      respjson.Field
		ShareAmount    respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnDepositSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnDepositSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnDepositSucceededWebhookPayloadStatus string

const (
	WalletActionEarnDepositSucceededWebhookPayloadStatusSucceeded WalletActionEarnDepositSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionEarnDepositSucceededWebhookPayloadType string

const (
	WalletActionEarnDepositSucceededWebhookPayloadTypeWalletActionEarnDepositSucceeded WalletActionEarnDepositSucceededWebhookPayloadType = "wallet_action.earn_deposit.succeeded"
)

// Payload for the wallet_action.earn_fee_collect.created webhook event.
type WalletActionEarnFeeCollectCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Base-unit amount of fees collected (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionEarnFeeCollectCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_fee_collect.created".
	Type WalletActionEarnFeeCollectCreatedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of fees collected (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnFeeCollectCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnFeeCollectCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnFeeCollectCreatedWebhookPayloadStatus string

const (
	WalletActionEarnFeeCollectCreatedWebhookPayloadStatusPending WalletActionEarnFeeCollectCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionEarnFeeCollectCreatedWebhookPayloadType string

const (
	WalletActionEarnFeeCollectCreatedWebhookPayloadTypeWalletActionEarnFeeCollectCreated WalletActionEarnFeeCollectCreatedWebhookPayloadType = "wallet_action.earn_fee_collect.created"
)

// Payload for the wallet_action.earn_fee_collect.failed webhook event.
type WalletActionEarnFeeCollectFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of fees collected (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionEarnFeeCollectFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_fee_collect.failed".
	Type WalletActionEarnFeeCollectFailedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of fees collected (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailedAt       respjson.Field
		FailureReason  respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnFeeCollectFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnFeeCollectFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnFeeCollectFailedWebhookPayloadStatus string

const (
	WalletActionEarnFeeCollectFailedWebhookPayloadStatusFailed WalletActionEarnFeeCollectFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionEarnFeeCollectFailedWebhookPayloadType string

const (
	WalletActionEarnFeeCollectFailedWebhookPayloadTypeWalletActionEarnFeeCollectFailed WalletActionEarnFeeCollectFailedWebhookPayloadType = "wallet_action.earn_fee_collect.failed"
)

// Payload for the wallet_action.earn_fee_collect.rejected webhook event.
type WalletActionEarnFeeCollectRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of fees collected (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionEarnFeeCollectRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_fee_collect.rejected".
	Type WalletActionEarnFeeCollectRejectedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of fees collected (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailureReason  respjson.Field
		RawAmount      respjson.Field
		RejectedAt     respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnFeeCollectRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnFeeCollectRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnFeeCollectRejectedWebhookPayloadStatus string

const (
	WalletActionEarnFeeCollectRejectedWebhookPayloadStatusRejected WalletActionEarnFeeCollectRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionEarnFeeCollectRejectedWebhookPayloadType string

const (
	WalletActionEarnFeeCollectRejectedWebhookPayloadTypeWalletActionEarnFeeCollectRejected WalletActionEarnFeeCollectRejectedWebhookPayloadType = "wallet_action.earn_fee_collect.rejected"
)

// Payload for the wallet_action.earn_fee_collect.succeeded webhook event.
type WalletActionEarnFeeCollectSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Base-unit amount of fees collected (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionEarnFeeCollectSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_fee_collect.succeeded".
	Type WalletActionEarnFeeCollectSucceededWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of fees collected (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnFeeCollectSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnFeeCollectSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnFeeCollectSucceededWebhookPayloadStatus string

const (
	WalletActionEarnFeeCollectSucceededWebhookPayloadStatusSucceeded WalletActionEarnFeeCollectSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionEarnFeeCollectSucceededWebhookPayloadType string

const (
	WalletActionEarnFeeCollectSucceededWebhookPayloadTypeWalletActionEarnFeeCollectSucceeded WalletActionEarnFeeCollectSucceededWebhookPayloadType = "wallet_action.earn_fee_collect.succeeded"
)

// Payload for the wallet_action.earn_incentive_claim.created webhook event.
type WalletActionEarnIncentiveClaimCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "tempo", "base").
	Chain string `json:"chain" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Claimed reward tokens. Populated after the preparation step fetches from Merkl.
	Rewards []EarnIncetiveClaimRewardEntry `json:"rewards" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionEarnIncentiveClaimCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_incentive_claim.created".
	Type WalletActionEarnIncentiveClaimCreatedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Chain          respjson.Field
		CreatedAt      respjson.Field
		Rewards        respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnIncentiveClaimCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnIncentiveClaimCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnIncentiveClaimCreatedWebhookPayloadStatus string

const (
	WalletActionEarnIncentiveClaimCreatedWebhookPayloadStatusPending WalletActionEarnIncentiveClaimCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionEarnIncentiveClaimCreatedWebhookPayloadType string

const (
	WalletActionEarnIncentiveClaimCreatedWebhookPayloadTypeWalletActionEarnIncentiveClaimCreated WalletActionEarnIncentiveClaimCreatedWebhookPayloadType = "wallet_action.earn_incentive_claim.created"
)

// Payload for the wallet_action.earn_incentive_claim.failed webhook event.
type WalletActionEarnIncentiveClaimFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "tempo", "base").
	Chain string `json:"chain" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Claimed reward tokens. Populated after the preparation step fetches from Merkl.
	Rewards []EarnIncetiveClaimRewardEntry `json:"rewards" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionEarnIncentiveClaimFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_incentive_claim.failed".
	Type WalletActionEarnIncentiveClaimFailedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Chain          respjson.Field
		CreatedAt      respjson.Field
		FailedAt       respjson.Field
		FailureReason  respjson.Field
		Rewards        respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnIncentiveClaimFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnIncentiveClaimFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnIncentiveClaimFailedWebhookPayloadStatus string

const (
	WalletActionEarnIncentiveClaimFailedWebhookPayloadStatusFailed WalletActionEarnIncentiveClaimFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionEarnIncentiveClaimFailedWebhookPayloadType string

const (
	WalletActionEarnIncentiveClaimFailedWebhookPayloadTypeWalletActionEarnIncentiveClaimFailed WalletActionEarnIncentiveClaimFailedWebhookPayloadType = "wallet_action.earn_incentive_claim.failed"
)

// Payload for the wallet_action.earn_incentive_claim.rejected webhook event.
type WalletActionEarnIncentiveClaimRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "tempo", "base").
	Chain string `json:"chain" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// Claimed reward tokens. Populated after the preparation step fetches from Merkl.
	Rewards []EarnIncetiveClaimRewardEntry `json:"rewards" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionEarnIncentiveClaimRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_incentive_claim.rejected".
	Type WalletActionEarnIncentiveClaimRejectedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Chain          respjson.Field
		CreatedAt      respjson.Field
		FailureReason  respjson.Field
		RejectedAt     respjson.Field
		Rewards        respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnIncentiveClaimRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnIncentiveClaimRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnIncentiveClaimRejectedWebhookPayloadStatus string

const (
	WalletActionEarnIncentiveClaimRejectedWebhookPayloadStatusRejected WalletActionEarnIncentiveClaimRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionEarnIncentiveClaimRejectedWebhookPayloadType string

const (
	WalletActionEarnIncentiveClaimRejectedWebhookPayloadTypeWalletActionEarnIncentiveClaimRejected WalletActionEarnIncentiveClaimRejectedWebhookPayloadType = "wallet_action.earn_incentive_claim.rejected"
)

// Payload for the wallet_action.earn_incentive_claim.succeeded webhook event.
type WalletActionEarnIncentiveClaimSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "tempo", "base").
	Chain string `json:"chain" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Claimed reward tokens. Populated after the preparation step fetches from Merkl.
	Rewards []EarnIncetiveClaimRewardEntry `json:"rewards" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionEarnIncentiveClaimSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_incentive_claim.succeeded".
	Type WalletActionEarnIncentiveClaimSucceededWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Chain          respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		Rewards        respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnIncentiveClaimSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnIncentiveClaimSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnIncentiveClaimSucceededWebhookPayloadStatus string

const (
	WalletActionEarnIncentiveClaimSucceededWebhookPayloadStatusSucceeded WalletActionEarnIncentiveClaimSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionEarnIncentiveClaimSucceededWebhookPayloadType string

const (
	WalletActionEarnIncentiveClaimSucceededWebhookPayloadTypeWalletActionEarnIncentiveClaimSucceeded WalletActionEarnIncentiveClaimSucceededWebhookPayloadType = "wallet_action.earn_incentive_claim.succeeded"
)

// Payload for the wallet_action.earn_withdraw.created webhook event.
type WalletActionEarnWithdrawCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Base-unit amount of asset withdrawn (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionEarnWithdrawCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_withdraw.created".
	Type WalletActionEarnWithdrawCreatedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset withdrawn (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnWithdrawCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnWithdrawCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnWithdrawCreatedWebhookPayloadStatus string

const (
	WalletActionEarnWithdrawCreatedWebhookPayloadStatusPending WalletActionEarnWithdrawCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionEarnWithdrawCreatedWebhookPayloadType string

const (
	WalletActionEarnWithdrawCreatedWebhookPayloadTypeWalletActionEarnWithdrawCreated WalletActionEarnWithdrawCreatedWebhookPayloadType = "wallet_action.earn_withdraw.created"
)

// Payload for the wallet_action.earn_withdraw.failed webhook event.
type WalletActionEarnWithdrawFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of asset withdrawn (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionEarnWithdrawFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_withdraw.failed".
	Type WalletActionEarnWithdrawFailedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset withdrawn (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailedAt       respjson.Field
		FailureReason  respjson.Field
		RawAmount      respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnWithdrawFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnWithdrawFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnWithdrawFailedWebhookPayloadStatus string

const (
	WalletActionEarnWithdrawFailedWebhookPayloadStatusFailed WalletActionEarnWithdrawFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionEarnWithdrawFailedWebhookPayloadType string

const (
	WalletActionEarnWithdrawFailedWebhookPayloadTypeWalletActionEarnWithdrawFailed WalletActionEarnWithdrawFailedWebhookPayloadType = "wallet_action.earn_withdraw.failed"
)

// Payload for the wallet_action.earn_withdraw.rejected webhook event.
type WalletActionEarnWithdrawRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of asset withdrawn (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionEarnWithdrawRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_withdraw.rejected".
	Type WalletActionEarnWithdrawRejectedWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset withdrawn (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailureReason  respjson.Field
		RawAmount      respjson.Field
		RejectedAt     respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnWithdrawRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnWithdrawRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnWithdrawRejectedWebhookPayloadStatus string

const (
	WalletActionEarnWithdrawRejectedWebhookPayloadStatusRejected WalletActionEarnWithdrawRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionEarnWithdrawRejectedWebhookPayloadType string

const (
	WalletActionEarnWithdrawRejectedWebhookPayloadTypeWalletActionEarnWithdrawRejected WalletActionEarnWithdrawRejectedWebhookPayloadType = "wallet_action.earn_withdraw.rejected"
)

// Payload for the wallet_action.earn_withdraw.succeeded webhook event.
type WalletActionEarnWithdrawSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Base-unit amount of asset withdrawn (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// Vault shares burned in base units.
	ShareAmount string `json:"share_amount" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionEarnWithdrawSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.earn_withdraw.succeeded".
	Type WalletActionEarnWithdrawSucceededWebhookPayloadType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Human-readable decimal amount of asset withdrawn (e.g. "1.5"). Only present when
	// the token is known in the asset registry.
	Amount string `json:"amount"`
	// Asset identifier (e.g. "usdc", "eth"). Only present when the token is known in
	// the asset registry.
	Asset string `json:"asset"`
	// Number of decimals for the underlying asset (e.g. 6 for USDC, 18 for ETH). Only
	// present when the token is known in the asset registry.
	Decimals int64 `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		AssetAddress   respjson.Field
		Caip2          respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		RawAmount      respjson.Field
		ShareAmount    respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		VaultAddress   respjson.Field
		VaultID        respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		Amount         respjson.Field
		Asset          respjson.Field
		Decimals       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionEarnWithdrawSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionEarnWithdrawSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionEarnWithdrawSucceededWebhookPayloadStatus string

const (
	WalletActionEarnWithdrawSucceededWebhookPayloadStatusSucceeded WalletActionEarnWithdrawSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionEarnWithdrawSucceededWebhookPayloadType string

const (
	WalletActionEarnWithdrawSucceededWebhookPayloadTypeWalletActionEarnWithdrawSucceeded WalletActionEarnWithdrawSucceededWebhookPayloadType = "wallet_action.earn_withdraw.succeeded"
)

// Payload for the wallet_action.payout.created webhook event.
type WalletActionPayoutCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The fiat currency the payout settles in (e.g. "usd").
	DestinationCurrency string `json:"destination_currency" api:"required"`
	// The registered external fiat account the payout settles to.
	DestinationFiatAccountID string `json:"destination_fiat_account_id" api:"required"`
	// The fiat payment rail the payout settles over (e.g. "ach", "sepa", "wire").
	DestinationPaymentRail string `json:"destination_payment_rail" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// Decimal amount offramped, in the asset's standard units (e.g. "100.00").
	SourceAmount string `json:"source_amount" api:"required"`
	// Source crypto asset sent on-chain (e.g. "usdc").
	SourceAsset string `json:"source_asset" api:"required"`
	// Source chain the crypto was sent from (e.g. "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionPayoutCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.payout.created".
	Type WalletActionPayoutCreatedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType               respjson.Field
		CreatedAt                respjson.Field
		DestinationCurrency      respjson.Field
		DestinationFiatAccountID respjson.Field
		DestinationPaymentRail   respjson.Field
		Environment              respjson.Field
		Provider                 respjson.Field
		SourceAmount             respjson.Field
		SourceAsset              respjson.Field
		SourceChain              respjson.Field
		Status                   respjson.Field
		Type                     respjson.Field
		WalletActionID           respjson.Field
		WalletID                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionPayoutCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionPayoutCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionPayoutCreatedWebhookPayloadStatus string

const (
	WalletActionPayoutCreatedWebhookPayloadStatusPending WalletActionPayoutCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionPayoutCreatedWebhookPayloadType string

const (
	WalletActionPayoutCreatedWebhookPayloadTypeWalletActionPayoutCreated WalletActionPayoutCreatedWebhookPayloadType = "wallet_action.payout.created"
)

// Payload for the wallet_action.payout.failed webhook event.
type WalletActionPayoutFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The fiat currency the payout settles in (e.g. "usd").
	DestinationCurrency string `json:"destination_currency" api:"required"`
	// The registered external fiat account the payout settles to.
	DestinationFiatAccountID string `json:"destination_fiat_account_id" api:"required"`
	// The fiat payment rail the payout settles over (e.g. "ach", "sepa", "wire").
	DestinationPaymentRail string `json:"destination_payment_rail" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// Decimal amount offramped, in the asset's standard units (e.g. "100.00").
	SourceAmount string `json:"source_amount" api:"required"`
	// Source crypto asset sent on-chain (e.g. "usdc").
	SourceAsset string `json:"source_asset" api:"required"`
	// Source chain the crypto was sent from (e.g. "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionPayoutFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.payout.failed".
	Type WalletActionPayoutFailedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType               respjson.Field
		CreatedAt                respjson.Field
		DestinationCurrency      respjson.Field
		DestinationFiatAccountID respjson.Field
		DestinationPaymentRail   respjson.Field
		Environment              respjson.Field
		FailedAt                 respjson.Field
		FailureReason            respjson.Field
		Provider                 respjson.Field
		SourceAmount             respjson.Field
		SourceAsset              respjson.Field
		SourceChain              respjson.Field
		Status                   respjson.Field
		Steps                    respjson.Field
		Type                     respjson.Field
		WalletActionID           respjson.Field
		WalletID                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionPayoutFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionPayoutFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionPayoutFailedWebhookPayloadStatus string

const (
	WalletActionPayoutFailedWebhookPayloadStatusFailed WalletActionPayoutFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionPayoutFailedWebhookPayloadType string

const (
	WalletActionPayoutFailedWebhookPayloadTypeWalletActionPayoutFailed WalletActionPayoutFailedWebhookPayloadType = "wallet_action.payout.failed"
)

// Payload for the wallet_action.payout.rejected webhook event.
type WalletActionPayoutRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The fiat currency the payout settles in (e.g. "usd").
	DestinationCurrency string `json:"destination_currency" api:"required"`
	// The registered external fiat account the payout settles to.
	DestinationFiatAccountID string `json:"destination_fiat_account_id" api:"required"`
	// The fiat payment rail the payout settles over (e.g. "ach", "sepa", "wire").
	DestinationPaymentRail string `json:"destination_payment_rail" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// Decimal amount offramped, in the asset's standard units (e.g. "100.00").
	SourceAmount string `json:"source_amount" api:"required"`
	// Source crypto asset sent on-chain (e.g. "usdc").
	SourceAsset string `json:"source_asset" api:"required"`
	// Source chain the crypto was sent from (e.g. "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionPayoutRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.payout.rejected".
	Type WalletActionPayoutRejectedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType               respjson.Field
		CreatedAt                respjson.Field
		DestinationCurrency      respjson.Field
		DestinationFiatAccountID respjson.Field
		DestinationPaymentRail   respjson.Field
		Environment              respjson.Field
		FailureReason            respjson.Field
		Provider                 respjson.Field
		RejectedAt               respjson.Field
		SourceAmount             respjson.Field
		SourceAsset              respjson.Field
		SourceChain              respjson.Field
		Status                   respjson.Field
		Steps                    respjson.Field
		Type                     respjson.Field
		WalletActionID           respjson.Field
		WalletID                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionPayoutRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionPayoutRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionPayoutRejectedWebhookPayloadStatus string

const (
	WalletActionPayoutRejectedWebhookPayloadStatusRejected WalletActionPayoutRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionPayoutRejectedWebhookPayloadType string

const (
	WalletActionPayoutRejectedWebhookPayloadTypeWalletActionPayoutRejected WalletActionPayoutRejectedWebhookPayloadType = "wallet_action.payout.rejected"
)

// Payload for the wallet_action.payout.succeeded webhook event.
type WalletActionPayoutSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The fiat currency the payout settles in (e.g. "usd").
	DestinationCurrency string `json:"destination_currency" api:"required"`
	// The registered external fiat account the payout settles to.
	DestinationFiatAccountID string `json:"destination_fiat_account_id" api:"required"`
	// The fiat payment rail the payout settles over (e.g. "ach", "sepa", "wire").
	DestinationPaymentRail string `json:"destination_payment_rail" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// Decimal amount offramped, in the asset's standard units (e.g. "100.00").
	SourceAmount string `json:"source_amount" api:"required"`
	// Source crypto asset sent on-chain (e.g. "usdc").
	SourceAsset string `json:"source_asset" api:"required"`
	// Source chain the crypto was sent from (e.g. "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionPayoutSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.payout.succeeded".
	Type WalletActionPayoutSucceededWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType               respjson.Field
		CompletedAt              respjson.Field
		CreatedAt                respjson.Field
		DestinationCurrency      respjson.Field
		DestinationFiatAccountID respjson.Field
		DestinationPaymentRail   respjson.Field
		Environment              respjson.Field
		Provider                 respjson.Field
		SourceAmount             respjson.Field
		SourceAsset              respjson.Field
		SourceChain              respjson.Field
		Status                   respjson.Field
		Steps                    respjson.Field
		Type                     respjson.Field
		WalletActionID           respjson.Field
		WalletID                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionPayoutSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionPayoutSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionPayoutSucceededWebhookPayloadStatus string

const (
	WalletActionPayoutSucceededWebhookPayloadStatusSucceeded WalletActionPayoutSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionPayoutSucceededWebhookPayloadType string

const (
	WalletActionPayoutSucceededWebhookPayloadTypeWalletActionPayoutSucceeded WalletActionPayoutSucceededWebhookPayloadType = "wallet_action.payout.succeeded"
)

// Payload for the wallet_action.swap.created webhook event.
type WalletActionSwapCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Amount of input token in base units. Populated after onchain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Token address being bought.
	OutputToken string `json:"output_token" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionSwapCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.swap.created".
	Type WalletActionSwapCreatedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		InputAmount    respjson.Field
		InputToken     respjson.Field
		OutputToken    respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionSwapCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionSwapCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionSwapCreatedWebhookPayloadStatus string

const (
	WalletActionSwapCreatedWebhookPayloadStatusPending WalletActionSwapCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionSwapCreatedWebhookPayloadType string

const (
	WalletActionSwapCreatedWebhookPayloadTypeWalletActionSwapCreated WalletActionSwapCreatedWebhookPayloadType = "wallet_action.swap.created"
)

// Payload for the wallet_action.swap.failed webhook event.
type WalletActionSwapFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Amount of input token in base units. Populated after onchain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Token address being bought.
	OutputToken string `json:"output_token" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionSwapFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.swap.failed".
	Type WalletActionSwapFailedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailedAt       respjson.Field
		FailureReason  respjson.Field
		InputAmount    respjson.Field
		InputToken     respjson.Field
		OutputToken    respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionSwapFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionSwapFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionSwapFailedWebhookPayloadStatus string

const (
	WalletActionSwapFailedWebhookPayloadStatusFailed WalletActionSwapFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionSwapFailedWebhookPayloadType string

const (
	WalletActionSwapFailedWebhookPayloadTypeWalletActionSwapFailed WalletActionSwapFailedWebhookPayloadType = "wallet_action.swap.failed"
)

// Payload for the wallet_action.swap.rejected webhook event.
type WalletActionSwapRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Amount of input token in base units. Populated after onchain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Token address being bought.
	OutputToken string `json:"output_token" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionSwapRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.swap.rejected".
	Type WalletActionSwapRejectedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Caip2          respjson.Field
		CreatedAt      respjson.Field
		FailureReason  respjson.Field
		InputAmount    respjson.Field
		InputToken     respjson.Field
		OutputToken    respjson.Field
		RejectedAt     respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionSwapRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionSwapRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionSwapRejectedWebhookPayloadStatus string

const (
	WalletActionSwapRejectedWebhookPayloadStatusRejected WalletActionSwapRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionSwapRejectedWebhookPayloadType string

const (
	WalletActionSwapRejectedWebhookPayloadTypeWalletActionSwapRejected WalletActionSwapRejectedWebhookPayloadType = "wallet_action.swap.rejected"
)

// Payload for the wallet_action.swap.succeeded webhook event.
type WalletActionSwapSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Amount of input token in base units. Populated after onchain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Amount of output token received, in base units. Populated after onchain
	// confirmation.
	OutputAmount string `json:"output_amount" api:"required"`
	// Token address being bought.
	OutputToken string `json:"output_token" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionSwapSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.swap.succeeded".
	Type WalletActionSwapSucceededWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType     respjson.Field
		Caip2          respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		InputAmount    respjson.Field
		InputToken     respjson.Field
		OutputAmount   respjson.Field
		OutputToken    respjson.Field
		Status         respjson.Field
		Steps          respjson.Field
		Type           respjson.Field
		WalletActionID respjson.Field
		WalletID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionSwapSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionSwapSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionSwapSucceededWebhookPayloadStatus string

const (
	WalletActionSwapSucceededWebhookPayloadStatusSucceeded WalletActionSwapSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionSwapSucceededWebhookPayloadType string

const (
	WalletActionSwapSucceededWebhookPayloadTypeWalletActionSwapSucceeded WalletActionSwapSucceededWebhookPayloadType = "wallet_action.swap.succeeded"
)

// Payload for the wallet_action.transfer.created webhook event.
type WalletActionTransferCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// Chain name (e.g. "tempo", "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "pending".
	Status WalletActionTransferCreatedWebhookPayloadStatus `json:"status" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.transfer.created".
	Type WalletActionTransferCreatedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Decimal amount sent on the source chain (e.g. "1.5"). Omitted for exact_output
	// cross-chain transfers until the source amount is determined.
	SourceAmount string `json:"source_amount"`
	// Asset identifier (e.g. "usdc", "eth"). Present when the transfer was initiated
	// with a named asset; omitted for custom-token transfers.
	SourceAsset string `json:"source_asset"`
	// Token contract address (EVM) or mint address (Solana). Present when the transfer
	// was initiated with `asset_address`.
	SourceAssetAddress string `json:"source_asset_address"`
	// Number of decimals for the transferred token. Present when the transfer was
	// initiated with `asset_address` and the decimals were resolved onchain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
		CreatedAt           respjson.Field
		DestinationAddress  respjson.Field
		SourceChain         respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		WalletActionID      respjson.Field
		WalletID            respjson.Field
		SourceAmount        respjson.Field
		SourceAsset         respjson.Field
		SourceAssetAddress  respjson.Field
		SourceAssetDecimals respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionTransferCreatedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionTransferCreatedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionTransferCreatedWebhookPayloadStatus string

const (
	WalletActionTransferCreatedWebhookPayloadStatusPending WalletActionTransferCreatedWebhookPayloadStatus = "pending"
)

// The type of webhook event.
type WalletActionTransferCreatedWebhookPayloadType string

const (
	WalletActionTransferCreatedWebhookPayloadTypeWalletActionTransferCreated WalletActionTransferCreatedWebhookPayloadType = "wallet_action.transfer.created"
)

// Payload for the wallet_action.transfer.failed webhook event.
type WalletActionTransferFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// ISO 8601 timestamp of when the wallet action failed.
	FailedAt string `json:"failed_at" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Chain name (e.g. "tempo", "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "failed".
	Status WalletActionTransferFailedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action. Completed steps will have transaction hashes;
	// the failing step will have a failure_reason.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.transfer.failed".
	Type WalletActionTransferFailedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Decimal amount sent on the source chain (e.g. "1.5"). Omitted for exact_output
	// cross-chain transfers until the source amount is determined.
	SourceAmount string `json:"source_amount"`
	// Asset identifier (e.g. "usdc", "eth"). Present when the transfer was initiated
	// with a named asset; omitted for custom-token transfers.
	SourceAsset string `json:"source_asset"`
	// Token contract address (EVM) or mint address (Solana). Present when the transfer
	// was initiated with `asset_address`.
	SourceAssetAddress string `json:"source_asset_address"`
	// Number of decimals for the transferred token. Present when the transfer was
	// initiated with `asset_address` and the decimals were resolved onchain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
		CreatedAt           respjson.Field
		DestinationAddress  respjson.Field
		FailedAt            respjson.Field
		FailureReason       respjson.Field
		SourceChain         respjson.Field
		Status              respjson.Field
		Steps               respjson.Field
		Type                respjson.Field
		WalletActionID      respjson.Field
		WalletID            respjson.Field
		SourceAmount        respjson.Field
		SourceAsset         respjson.Field
		SourceAssetAddress  respjson.Field
		SourceAssetDecimals respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionTransferFailedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionTransferFailedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionTransferFailedWebhookPayloadStatus string

const (
	WalletActionTransferFailedWebhookPayloadStatusFailed WalletActionTransferFailedWebhookPayloadStatus = "failed"
)

// The type of webhook event.
type WalletActionTransferFailedWebhookPayloadType string

const (
	WalletActionTransferFailedWebhookPayloadTypeWalletActionTransferFailed WalletActionTransferFailedWebhookPayloadType = "wallet_action.transfer.failed"
)

// Payload for the wallet_action.transfer.rejected webhook event.
type WalletActionTransferRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// ISO 8601 timestamp of when the wallet action was rejected.
	RejectedAt string `json:"rejected_at" api:"required"`
	// Chain name (e.g. "tempo", "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "rejected".
	Status WalletActionTransferRejectedWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action at the time of rejection.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.transfer.rejected".
	Type WalletActionTransferRejectedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Decimal amount sent on the source chain (e.g. "1.5"). Omitted for exact_output
	// cross-chain transfers until the source amount is determined.
	SourceAmount string `json:"source_amount"`
	// Asset identifier (e.g. "usdc", "eth"). Present when the transfer was initiated
	// with a named asset; omitted for custom-token transfers.
	SourceAsset string `json:"source_asset"`
	// Token contract address (EVM) or mint address (Solana). Present when the transfer
	// was initiated with `asset_address`.
	SourceAssetAddress string `json:"source_asset_address"`
	// Number of decimals for the transferred token. Present when the transfer was
	// initiated with `asset_address` and the decimals were resolved onchain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
		CreatedAt           respjson.Field
		DestinationAddress  respjson.Field
		FailureReason       respjson.Field
		RejectedAt          respjson.Field
		SourceChain         respjson.Field
		Status              respjson.Field
		Steps               respjson.Field
		Type                respjson.Field
		WalletActionID      respjson.Field
		WalletID            respjson.Field
		SourceAmount        respjson.Field
		SourceAsset         respjson.Field
		SourceAssetAddress  respjson.Field
		SourceAssetDecimals respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionTransferRejectedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionTransferRejectedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionTransferRejectedWebhookPayloadStatus string

const (
	WalletActionTransferRejectedWebhookPayloadStatusRejected WalletActionTransferRejectedWebhookPayloadStatus = "rejected"
)

// The type of webhook event.
type WalletActionTransferRejectedWebhookPayloadType string

const (
	WalletActionTransferRejectedWebhookPayloadTypeWalletActionTransferRejected WalletActionTransferRejectedWebhookPayloadType = "wallet_action.transfer.rejected"
)

// Payload for the wallet_action.transfer.succeeded webhook event.
type WalletActionTransferSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim", "earn_fee_collect", "payout".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// ISO 8601 timestamp of when the wallet action completed successfully.
	CompletedAt string `json:"completed_at" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// Chain name (e.g. "tempo", "base").
	SourceChain string `json:"source_chain" api:"required"`
	// The status of the wallet action.
	//
	// Any of "succeeded".
	Status WalletActionTransferSucceededWebhookPayloadStatus `json:"status" api:"required"`
	// The steps of the wallet action, including transaction hashes.
	Steps []WalletActionStepUnion `json:"steps" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_action.transfer.succeeded".
	Type WalletActionTransferSucceededWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet action.
	WalletActionID string `json:"wallet_action_id" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Decimal amount sent on the source chain (e.g. "1.5"). Omitted for exact_output
	// cross-chain transfers until the source amount is determined.
	SourceAmount string `json:"source_amount"`
	// Asset identifier (e.g. "usdc", "eth"). Present when the transfer was initiated
	// with a named asset; omitted for custom-token transfers.
	SourceAsset string `json:"source_asset"`
	// Token contract address (EVM) or mint address (Solana). Present when the transfer
	// was initiated with `asset_address`.
	SourceAssetAddress string `json:"source_asset_address"`
	// Number of decimals for the transferred token. Present when the transfer was
	// initiated with `asset_address` and the decimals were resolved onchain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
		CompletedAt         respjson.Field
		CreatedAt           respjson.Field
		DestinationAddress  respjson.Field
		SourceChain         respjson.Field
		Status              respjson.Field
		Steps               respjson.Field
		Type                respjson.Field
		WalletActionID      respjson.Field
		WalletID            respjson.Field
		SourceAmount        respjson.Field
		SourceAsset         respjson.Field
		SourceAssetAddress  respjson.Field
		SourceAssetDecimals respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletActionTransferSucceededWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletActionTransferSucceededWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the wallet action.
type WalletActionTransferSucceededWebhookPayloadStatus string

const (
	WalletActionTransferSucceededWebhookPayloadStatusSucceeded WalletActionTransferSucceededWebhookPayloadStatus = "succeeded"
)

// The type of webhook event.
type WalletActionTransferSucceededWebhookPayloadType string

const (
	WalletActionTransferSucceededWebhookPayloadTypeWalletActionTransferSucceeded WalletActionTransferSucceededWebhookPayloadType = "wallet_action.transfer.succeeded"
)

// Payload for the wallet.archived webhook event.
type WalletArchivedWebhookPayload struct {
	// Unix timestamp of when the wallet was archived.
	ArchivedAt float64 `json:"archived_at" api:"required"`
	// The chain type of the archived wallet.
	ChainType string `json:"chain_type" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.archived".
	Type WalletArchivedWebhookPayloadType `json:"type" api:"required"`
	// The address of the archived wallet.
	WalletAddress string `json:"wallet_address" api:"required"`
	// The ID of the archived wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArchivedAt    respjson.Field
		ChainType     respjson.Field
		Type          respjson.Field
		WalletAddress respjson.Field
		WalletID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletArchivedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletArchivedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type WalletArchivedWebhookPayloadType string

const (
	WalletArchivedWebhookPayloadTypeWalletArchived WalletArchivedWebhookPayloadType = "wallet.archived"
)

// Payload for the wallet_automation.submitted webhook event.
type WalletAutomationSubmittedWebhookPayload struct {
	// The ID of the wallet action created to fulfill the automation.
	ActionID string `json:"action_id" api:"required"`
	// The ID of the automation that fired.
	AutomationID string `json:"automation_id" api:"required"`
	// ISO 8601 timestamp of when the automation was submitted.
	CreatedAt string `json:"created_at" api:"required"`
	// Contract address of the triggering deposit's asset, or 'native-token' for the
	// native asset.
	TriggerAssetAddress string `json:"trigger_asset_address" api:"required"`
	// CAIP-2 chain identifier of the triggering deposit (e.g., 'eip155:8453').
	TriggerCaip2 string `json:"trigger_caip2" api:"required"`
	// The ID of the automation execution that fired.
	TriggerID string `json:"trigger_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet_automation.submitted".
	Type WalletAutomationSubmittedWebhookPayloadType `json:"type" api:"required"`
	// The ID of the wallet the automation fired for.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID            respjson.Field
		AutomationID        respjson.Field
		CreatedAt           respjson.Field
		TriggerAssetAddress respjson.Field
		TriggerCaip2        respjson.Field
		TriggerID           respjson.Field
		Type                respjson.Field
		WalletID            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationSubmittedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationSubmittedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type WalletAutomationSubmittedWebhookPayloadType string

const (
	WalletAutomationSubmittedWebhookPayloadTypeWalletAutomationSubmitted WalletAutomationSubmittedWebhookPayloadType = "wallet_automation.submitted"
)

// WalletFundsAssetUnion contains all possible properties and values from
// [WalletFundsNativeTokenAsset], [WalletFundsErc20Asset], [WalletFundsSplAsset],
// [WalletFundsSacAsset], [WalletFundsTrc20Asset].
//
// Use the [WalletFundsAssetUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletFundsAssetUnion struct {
	// This field is a union of [any], [string], [string], [string]
	Address WalletFundsAssetUnionAddress `json:"address"`
	// Any of "native-token", "erc20", "spl", "sac", "trc20".
	Type string `json:"type"`
	// This field is from variant [WalletFundsSplAsset].
	Mint string `json:"mint"`
	JSON struct {
		Address respjson.Field
		Type    respjson.Field
		Mint    respjson.Field
		raw     string
	} `json:"-"`
}

// anyWalletFundsAsset is implemented by each variant of [WalletFundsAssetUnion] to
// add type safety for the return type of [WalletFundsAssetUnion.AsAny]
type anyWalletFundsAsset interface {
	implWalletFundsAssetUnion()
}

func (WalletFundsNativeTokenAsset) implWalletFundsAssetUnion() {}
func (WalletFundsErc20Asset) implWalletFundsAssetUnion()       {}
func (WalletFundsSplAsset) implWalletFundsAssetUnion()         {}
func (WalletFundsSacAsset) implWalletFundsAssetUnion()         {}
func (WalletFundsTrc20Asset) implWalletFundsAssetUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletFundsAssetUnion.AsAny().(type) {
//	case privyclient.WalletFundsNativeTokenAsset:
//	case privyclient.WalletFundsErc20Asset:
//	case privyclient.WalletFundsSplAsset:
//	case privyclient.WalletFundsSacAsset:
//	case privyclient.WalletFundsTrc20Asset:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletFundsAssetUnion) AsAny() anyWalletFundsAsset {
	switch u.Type {
	case "native-token":
		return u.AsNativeToken()
	case "erc20":
		return u.AsErc20()
	case "spl":
		return u.AsSpl()
	case "sac":
		return u.AsSac()
	case "trc20":
		return u.AsTrc20()
	}
	return nil
}

func (u WalletFundsAssetUnion) AsNativeToken() (v WalletFundsNativeTokenAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletFundsAssetUnion) AsErc20() (v WalletFundsErc20Asset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletFundsAssetUnion) AsSpl() (v WalletFundsSplAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletFundsAssetUnion) AsSac() (v WalletFundsSacAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletFundsAssetUnion) AsTrc20() (v WalletFundsTrc20Asset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletFundsAssetUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletFundsAssetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletFundsAssetUnionAddress is an implicit subunion of [WalletFundsAssetUnion].
// WalletFundsAssetUnionAddress provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletFundsAssetUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfWalletFundsNativeTokenAssetAddress OfString]
type WalletFundsAssetUnionAddress struct {
	// This field will be present if the value is a [any] instead of an object.
	OfWalletFundsNativeTokenAssetAddress any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfWalletFundsNativeTokenAssetAddress respjson.Field
		OfString                             respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *WalletFundsAssetUnionAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An ERC-20 token asset.
type WalletFundsErc20Asset struct {
	Address string `json:"address" api:"required"`
	// Any of "erc20".
	Type WalletFundsErc20AssetType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletFundsErc20Asset) RawJSON() string { return r.JSON.raw }
func (r *WalletFundsErc20Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletFundsErc20AssetType string

const (
	WalletFundsErc20AssetTypeErc20 WalletFundsErc20AssetType = "erc20"
)

// A native token asset (e.g. ETH, SOL).
type WalletFundsNativeTokenAsset struct {
	Address any `json:"address" api:"required"`
	// Any of "native-token".
	Type WalletFundsNativeTokenAssetType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletFundsNativeTokenAsset) RawJSON() string { return r.JSON.raw }
func (r *WalletFundsNativeTokenAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletFundsNativeTokenAssetType string

const (
	WalletFundsNativeTokenAssetTypeNativeToken WalletFundsNativeTokenAssetType = "native-token"
)

// A Stellar Asset Contract (SAC) asset.
type WalletFundsSacAsset struct {
	Address string `json:"address" api:"required"`
	// Any of "sac".
	Type WalletFundsSacAssetType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletFundsSacAsset) RawJSON() string { return r.JSON.raw }
func (r *WalletFundsSacAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletFundsSacAssetType string

const (
	WalletFundsSacAssetTypeSac WalletFundsSacAssetType = "sac"
)

// A Solana SPL token asset.
type WalletFundsSplAsset struct {
	Mint string `json:"mint" api:"required"`
	// Any of "spl".
	Type WalletFundsSplAssetType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mint        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletFundsSplAsset) RawJSON() string { return r.JSON.raw }
func (r *WalletFundsSplAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletFundsSplAssetType string

const (
	WalletFundsSplAssetTypeSpl WalletFundsSplAssetType = "spl"
)

// A Tron TRC-20 token asset.
type WalletFundsTrc20Asset struct {
	Address string `json:"address" api:"required"`
	// Any of "trc20".
	Type WalletFundsTrc20AssetType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletFundsTrc20Asset) RawJSON() string { return r.JSON.raw }
func (r *WalletFundsTrc20Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletFundsTrc20AssetType string

const (
	WalletFundsTrc20AssetTypeTrc20 WalletFundsTrc20AssetType = "trc20"
)

// Payload for the wallet.recovered webhook event.
type WalletRecoveredWebhookPayload struct {
	// The type of webhook event.
	//
	// Any of "wallet.recovered".
	Type WalletRecoveredWebhookPayloadType `json:"type" api:"required"`
	// The ID of the user.
	UserID string `json:"user_id" api:"required"`
	// The address of the wallet.
	WalletAddress string `json:"wallet_address" api:"required"`
	// The ID of the wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		UserID        respjson.Field
		WalletAddress respjson.Field
		WalletID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRecoveredWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletRecoveredWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type WalletRecoveredWebhookPayloadType string

const (
	WalletRecoveredWebhookPayloadTypeWalletRecovered WalletRecoveredWebhookPayloadType = "wallet.recovered"
)

// Recovery method types for embedded wallet recovery setup webhooks.
type WalletRecoverySetupMethod string

const (
	WalletRecoverySetupMethodUserPasscodeDerivedRecoveryKey  WalletRecoverySetupMethod = "user_passcode_derived_recovery_key"
	WalletRecoverySetupMethodPrivyPasscodeDerivedRecoveryKey WalletRecoverySetupMethod = "privy_passcode_derived_recovery_key"
	WalletRecoverySetupMethodPrivyGeneratedRecoveryKey       WalletRecoverySetupMethod = "privy_generated_recovery_key"
	WalletRecoverySetupMethodGoogleDriveRecoverySecret       WalletRecoverySetupMethod = "google_drive_recovery_secret"
	WalletRecoverySetupMethodICloudRecoverySecret            WalletRecoverySetupMethod = "icloud_recovery_secret"
	WalletRecoverySetupMethodRecoveryEncryptionKey           WalletRecoverySetupMethod = "recovery_encryption_key"
)

// Payload for the wallet.recovery_setup webhook event.
type WalletRecoverySetupWebhookPayload struct {
	// Recovery method types for embedded wallet recovery setup webhooks.
	//
	// Any of "user_passcode_derived_recovery_key",
	// "privy_passcode_derived_recovery_key", "privy_generated_recovery_key",
	// "google_drive_recovery_secret", "icloud_recovery_secret",
	// "recovery_encryption_key".
	Method WalletRecoverySetupMethod `json:"method" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.recovery_setup".
	Type WalletRecoverySetupWebhookPayloadType `json:"type" api:"required"`
	// The ID of the user.
	UserID string `json:"user_id" api:"required"`
	// The address of the wallet.
	WalletAddress string `json:"wallet_address" api:"required"`
	// The ID of the wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		Type          respjson.Field
		UserID        respjson.Field
		WalletAddress respjson.Field
		WalletID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRecoverySetupWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletRecoverySetupWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type WalletRecoverySetupWebhookPayloadType string

const (
	WalletRecoverySetupWebhookPayloadTypeWalletRecoverySetup WalletRecoverySetupWebhookPayloadType = "wallet.recovery_setup"
)

// Payload for the wallet.restored webhook event.
type WalletRestoredWebhookPayload struct {
	// The chain type of the restored wallet.
	ChainType string `json:"chain_type" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.restored".
	Type WalletRestoredWebhookPayloadType `json:"type" api:"required"`
	// The address of the restored wallet.
	WalletAddress string `json:"wallet_address" api:"required"`
	// The ID of the restored wallet.
	WalletID string `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType     respjson.Field
		Type          respjson.Field
		WalletAddress respjson.Field
		WalletID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRestoredWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *WalletRestoredWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type WalletRestoredWebhookPayloadType string

const (
	WalletRestoredWebhookPayloadTypeWalletRestored WalletRestoredWebhookPayloadType = "wallet.restored"
)

// Payload for the yield.claim.confirmed webhook event.
type YieldClaimConfirmedWebhookPayload struct {
	Caip2         string             `json:"caip2" api:"required"`
	Rewards       []YieldClaimReward `json:"rewards" api:"required"`
	TransactionID string             `json:"transaction_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "yield.claim.confirmed".
	Type     YieldClaimConfirmedWebhookPayloadType `json:"type" api:"required"`
	WalletID string                                `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2         respjson.Field
		Rewards       respjson.Field
		TransactionID respjson.Field
		Type          respjson.Field
		WalletID      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YieldClaimConfirmedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *YieldClaimConfirmedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type YieldClaimConfirmedWebhookPayloadType string

const (
	YieldClaimConfirmedWebhookPayloadTypeYieldClaimConfirmed YieldClaimConfirmedWebhookPayloadType = "yield.claim.confirmed"
)

// A single reward token claimed from a yield vault.
type YieldClaimReward struct {
	Amount       string `json:"amount" api:"required"`
	TokenAddress string `json:"token_address" api:"required"`
	TokenSymbol  string `json:"token_symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		TokenAddress respjson.Field
		TokenSymbol  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YieldClaimReward) RawJSON() string { return r.JSON.raw }
func (r *YieldClaimReward) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the yield.deposit.confirmed webhook event.
type YieldDepositConfirmedWebhookPayload struct {
	Assets string `json:"assets" api:"required"`
	Caip2  string `json:"caip2" api:"required"`
	Owner  string `json:"owner" api:"required"`
	Sender string `json:"sender" api:"required"`
	Shares string `json:"shares" api:"required"`
	// The type of webhook event.
	//
	// Any of "yield.deposit.confirmed".
	Type         YieldDepositConfirmedWebhookPayloadType `json:"type" api:"required"`
	VaultAddress string                                  `json:"vault_address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets       respjson.Field
		Caip2        respjson.Field
		Owner        respjson.Field
		Sender       respjson.Field
		Shares       respjson.Field
		Type         respjson.Field
		VaultAddress respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YieldDepositConfirmedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *YieldDepositConfirmedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type YieldDepositConfirmedWebhookPayloadType string

const (
	YieldDepositConfirmedWebhookPayloadTypeYieldDepositConfirmed YieldDepositConfirmedWebhookPayloadType = "yield.deposit.confirmed"
)

// Payload for the yield.withdraw.confirmed webhook event.
type YieldWithdrawConfirmedWebhookPayload struct {
	Assets   string `json:"assets" api:"required"`
	Caip2    string `json:"caip2" api:"required"`
	Owner    string `json:"owner" api:"required"`
	Receiver string `json:"receiver" api:"required"`
	Sender   string `json:"sender" api:"required"`
	Shares   string `json:"shares" api:"required"`
	// The type of webhook event.
	//
	// Any of "yield.withdraw.confirmed".
	Type         YieldWithdrawConfirmedWebhookPayloadType `json:"type" api:"required"`
	VaultAddress string                                   `json:"vault_address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets       respjson.Field
		Caip2        respjson.Field
		Owner        respjson.Field
		Receiver     respjson.Field
		Sender       respjson.Field
		Shares       respjson.Field
		Type         respjson.Field
		VaultAddress respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YieldWithdrawConfirmedWebhookPayload) RawJSON() string { return r.JSON.raw }
func (r *YieldWithdrawConfirmedWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type YieldWithdrawConfirmedWebhookPayloadType string

const (
	YieldWithdrawConfirmedWebhookPayloadTypeYieldWithdrawConfirmed YieldWithdrawConfirmedWebhookPayloadType = "yield.withdraw.confirmed"
)

type OrganizationKYBUpdatedWebhookEvent struct {
	Changes map[string][]any `json:"changes" api:"required"`
	// Full KYB state snapshot in a KYB update event.
	Data OrganizationKYBUpdatedData `json:"data" api:"required"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment    KyxEnvironment `json:"environment" api:"required"`
	OrganizationID string         `json:"organization_id" api:"required"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider" api:"required"`
	// The type of webhook event.
	//
	// Any of "organization.kyb.updated".
	Type OrganizationKYBUpdatedWebhookEventType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes        respjson.Field
		Data           respjson.Field
		Environment    respjson.Field
		OrganizationID respjson.Field
		Provider       respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationKYBUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *OrganizationKYBUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type OrganizationKYBUpdatedWebhookEventType string

const (
	OrganizationKYBUpdatedWebhookEventTypeOrganizationKYBUpdated OrganizationKYBUpdatedWebhookEventType = "organization.kyb.updated"
)

type UserKYCUpdatedWebhookEvent struct {
	Changes map[string][]any `json:"changes" api:"required"`
	// Full KYC state snapshot in a KYC update event.
	Data UserKYCUpdatedData `json:"data" api:"required"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment" api:"required"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider" api:"required"`
	// The type of webhook event.
	//
	// Any of "user.kyc.updated".
	Type   UserKYCUpdatedWebhookEventType `json:"type" api:"required"`
	UserID string                         `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		Data        respjson.Field
		Environment respjson.Field
		Provider    respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserKYCUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *UserKYCUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserKYCUpdatedWebhookEventType string

const (
	UserKYCUpdatedWebhookEventTypeUserKYCUpdated UserKYCUpdatedWebhookEventType = "user.kyc.updated"
)

type WalletDepositAccountDepositCompletedWebhookEvent struct {
	// Details of a fiat deposit that has finished converting and been delivered to the
	// wallet.
	Data             DepositCompletedData `json:"data" api:"required"`
	DepositAccountID string               `json:"deposit_account_id" api:"required"`
	// Any of "fiat".
	DepositType WalletDepositAccountDepositCompletedWebhookEventDepositType `json:"deposit_type" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// The deposit's ID in the provider's system (e.g. Bridge), not a Privy ID.
	ProviderDepositID string `json:"provider_deposit_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.deposit_account.deposit_completed".
	Type     WalletDepositAccountDepositCompletedWebhookEventType `json:"type" api:"required"`
	WalletID string                                               `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data              respjson.Field
		DepositAccountID  respjson.Field
		DepositType       respjson.Field
		Environment       respjson.Field
		Provider          respjson.Field
		ProviderDepositID respjson.Field
		Type              respjson.Field
		WalletID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletDepositAccountDepositCompletedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *WalletDepositAccountDepositCompletedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletDepositAccountDepositCompletedWebhookEventDepositType string

const (
	WalletDepositAccountDepositCompletedWebhookEventDepositTypeFiat WalletDepositAccountDepositCompletedWebhookEventDepositType = "fiat"
)

// The type of webhook event.
type WalletDepositAccountDepositCompletedWebhookEventType string

const (
	WalletDepositAccountDepositCompletedWebhookEventTypeWalletDepositAccountDepositCompleted WalletDepositAccountDepositCompletedWebhookEventType = "wallet.deposit_account.deposit_completed"
)

type WalletDepositAccountDepositFailedWebhookEvent struct {
	// Details of a fiat deposit that failed to convert and was refunded to the sender.
	Data             DepositFailedData `json:"data" api:"required"`
	DepositAccountID string            `json:"deposit_account_id" api:"required"`
	// Any of "fiat".
	DepositType WalletDepositAccountDepositFailedWebhookEventDepositType `json:"deposit_type" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.deposit_account.deposit_failed".
	Type     WalletDepositAccountDepositFailedWebhookEventType `json:"type" api:"required"`
	WalletID string                                            `json:"wallet_id" api:"required"`
	// The deposit's ID in the provider's system (e.g. Bridge), when the provider
	// assigned one.
	ProviderDepositID string `json:"provider_deposit_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data              respjson.Field
		DepositAccountID  respjson.Field
		DepositType       respjson.Field
		Environment       respjson.Field
		Provider          respjson.Field
		Type              respjson.Field
		WalletID          respjson.Field
		ProviderDepositID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletDepositAccountDepositFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *WalletDepositAccountDepositFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletDepositAccountDepositFailedWebhookEventDepositType string

const (
	WalletDepositAccountDepositFailedWebhookEventDepositTypeFiat WalletDepositAccountDepositFailedWebhookEventDepositType = "fiat"
)

// The type of webhook event.
type WalletDepositAccountDepositFailedWebhookEventType string

const (
	WalletDepositAccountDepositFailedWebhookEventTypeWalletDepositAccountDepositFailed WalletDepositAccountDepositFailedWebhookEventType = "wallet.deposit_account.deposit_failed"
)

type WalletDepositAccountDepositStartedWebhookEvent struct {
	// Details of a fiat deposit that has begun processing into a deposit account.
	Data             DepositStartedData `json:"data" api:"required"`
	DepositAccountID string             `json:"deposit_account_id" api:"required"`
	// Any of "fiat".
	DepositType WalletDepositAccountDepositStartedWebhookEventDepositType `json:"deposit_type" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `json:"provider" api:"required"`
	// The deposit's ID in the provider's system (e.g. Bridge), not a Privy ID.
	ProviderDepositID string `json:"provider_deposit_id" api:"required"`
	// The type of webhook event.
	//
	// Any of "wallet.deposit_account.deposit_started".
	Type     WalletDepositAccountDepositStartedWebhookEventType `json:"type" api:"required"`
	WalletID string                                             `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data              respjson.Field
		DepositAccountID  respjson.Field
		DepositType       respjson.Field
		Environment       respjson.Field
		Provider          respjson.Field
		ProviderDepositID respjson.Field
		Type              respjson.Field
		WalletID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletDepositAccountDepositStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *WalletDepositAccountDepositStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletDepositAccountDepositStartedWebhookEventDepositType string

const (
	WalletDepositAccountDepositStartedWebhookEventDepositTypeFiat WalletDepositAccountDepositStartedWebhookEventDepositType = "fiat"
)

// The type of webhook event.
type WalletDepositAccountDepositStartedWebhookEventType string

const (
	WalletDepositAccountDepositStartedWebhookEventTypeWalletDepositAccountDepositStarted WalletDepositAccountDepositStartedWebhookEventType = "wallet.deposit_account.deposit_started"
)

// UnsafeUnwrapWebhookEventUnion contains all possible properties and values from
// [IntentAuthorizedWebhookPayload], [IntentCreatedWebhookPayload],
// [IntentExecutedWebhookPayload], [IntentFailedWebhookPayload],
// [IntentRejectedWebhookPayload], [MfaDisabledWebhookPayload],
// [MfaEnabledWebhookPayload], [OrganizationKYBUpdatedWebhookEvent],
// [TransactionBroadcastedWebhookPayload], [TransactionConfirmedWebhookPayload],
// [TransactionExecutionRevertedWebhookPayload], [TransactionFailedWebhookPayload],
// [TransactionProviderErrorWebhookPayload], [TransactionReplacedWebhookPayload],
// [TransactionStillPendingWebhookPayload],
// [UsageCrossChainFeeRecordedWebhookPayload],
// [UsageGasSponsorshipRecordedWebhookPayload], [UserAuthenticatedWebhookPayload],
// [UserCreatedWebhookPayload], [UserDeletedWebhookPayload],
// [UserKYCUpdatedWebhookEvent], [UserLinkedAccountWebhookPayload],
// [UserTransferredAccountWebhookPayload], [UserUnlinkedAccountWebhookPayload],
// [UserUpdatedAccountWebhookPayload], [UserWalletCreatedWebhookPayload],
// [UserOperationCompletedWebhookPayload], [WalletArchivedWebhookPayload],
// [WalletDepositAccountDepositCompletedWebhookEvent],
// [WalletDepositAccountDepositFailedWebhookEvent],
// [WalletDepositAccountDepositStartedWebhookEvent],
// [FundsDepositedWebhookPayload], [FundsWithdrawnWebhookPayload],
// [PrivateKeyExportWebhookPayload], [WalletRecoveredWebhookPayload],
// [WalletRecoverySetupWebhookPayload], [WalletRestoredWebhookPayload],
// [WalletActionEarnDepositCreatedWebhookPayload],
// [WalletActionEarnDepositFailedWebhookPayload],
// [WalletActionEarnDepositRejectedWebhookPayload],
// [WalletActionEarnDepositSucceededWebhookPayload],
// [WalletActionEarnFeeCollectCreatedWebhookPayload],
// [WalletActionEarnFeeCollectFailedWebhookPayload],
// [WalletActionEarnFeeCollectRejectedWebhookPayload],
// [WalletActionEarnFeeCollectSucceededWebhookPayload],
// [WalletActionEarnIncentiveClaimCreatedWebhookPayload],
// [WalletActionEarnIncentiveClaimFailedWebhookPayload],
// [WalletActionEarnIncentiveClaimRejectedWebhookPayload],
// [WalletActionEarnIncentiveClaimSucceededWebhookPayload],
// [WalletActionEarnWithdrawCreatedWebhookPayload],
// [WalletActionEarnWithdrawFailedWebhookPayload],
// [WalletActionEarnWithdrawRejectedWebhookPayload],
// [WalletActionEarnWithdrawSucceededWebhookPayload],
// [WalletActionPayoutCreatedWebhookPayload],
// [WalletActionPayoutFailedWebhookPayload],
// [WalletActionPayoutRejectedWebhookPayload],
// [WalletActionPayoutSucceededWebhookPayload],
// [WalletActionSwapCreatedWebhookPayload], [WalletActionSwapFailedWebhookPayload],
// [WalletActionSwapRejectedWebhookPayload],
// [WalletActionSwapSucceededWebhookPayload],
// [WalletActionTransferCreatedWebhookPayload],
// [WalletActionTransferFailedWebhookPayload],
// [WalletActionTransferRejectedWebhookPayload],
// [WalletActionTransferSucceededWebhookPayload],
// [WalletAutomationSubmittedWebhookPayload], [YieldClaimConfirmedWebhookPayload],
// [YieldDepositConfirmedWebhookPayload], [YieldWithdrawConfirmedWebhookPayload].
//
// Use the [UnsafeUnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	// This field is from variant [IntentAuthorizedWebhookPayload].
	AuthorizedAt float64 `json:"authorized_at"`
	// This field is a union of [float64], [float64], [float64], [float64], [float64],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string]
	CreatedAt UnsafeUnwrapWebhookEventUnionCreatedAt `json:"created_at"`
	ExpiresAt float64                                `json:"expires_at"`
	IntentID  string                                 `json:"intent_id"`
	// This field is from variant [IntentAuthorizedWebhookPayload].
	IntentType IntentType `json:"intent_type"`
	// This field is from variant [IntentAuthorizedWebhookPayload].
	Member IntentAuthorizationKeyQuorumMemberUnion `json:"member"`
	Status string                                  `json:"status"`
	// Any of "intent.authorized", "intent.created", "intent.executed",
	// "intent.failed", "intent.rejected", "mfa.disabled", "mfa.enabled",
	// "organization.kyb.updated", "transaction.broadcasted", "transaction.confirmed",
	// "transaction.execution_reverted", "transaction.failed",
	// "transaction.provider_error", "transaction.replaced",
	// "transaction.still_pending", "usage.cross_chain_fee.recorded",
	// "usage.gas_sponsorship.recorded", "user.authenticated", "user.created",
	// "user.deleted", "user.kyc.updated", "user.linked_account",
	// "user.transferred_account", "user.unlinked_account", "user.updated_account",
	// "user.wallet_created", "user_operation.completed", "wallet.archived",
	// "wallet.deposit_account.deposit_completed",
	// "wallet.deposit_account.deposit_failed",
	// "wallet.deposit_account.deposit_started", "wallet.funds_deposited",
	// "wallet.funds_withdrawn", "wallet.private_key_export", "wallet.recovered",
	// "wallet.recovery_setup", "wallet.restored",
	// "wallet_action.earn_deposit.created", "wallet_action.earn_deposit.failed",
	// "wallet_action.earn_deposit.rejected", "wallet_action.earn_deposit.succeeded",
	// "wallet_action.earn_fee_collect.created",
	// "wallet_action.earn_fee_collect.failed",
	// "wallet_action.earn_fee_collect.rejected",
	// "wallet_action.earn_fee_collect.succeeded",
	// "wallet_action.earn_incentive_claim.created",
	// "wallet_action.earn_incentive_claim.failed",
	// "wallet_action.earn_incentive_claim.rejected",
	// "wallet_action.earn_incentive_claim.succeeded",
	// "wallet_action.earn_withdraw.created", "wallet_action.earn_withdraw.failed",
	// "wallet_action.earn_withdraw.rejected", "wallet_action.earn_withdraw.succeeded",
	// "wallet_action.payout.created", "wallet_action.payout.failed",
	// "wallet_action.payout.rejected", "wallet_action.payout.succeeded",
	// "wallet_action.swap.created", "wallet_action.swap.failed",
	// "wallet_action.swap.rejected", "wallet_action.swap.succeeded",
	// "wallet_action.transfer.created", "wallet_action.transfer.failed",
	// "wallet_action.transfer.rejected", "wallet_action.transfer.succeeded",
	// "wallet_automation.submitted", "yield.claim.confirmed",
	// "yield.deposit.confirmed", "yield.withdraw.confirmed".
	Type                 string `json:"type"`
	CreatedByDisplayName string `json:"created_by_display_name"`
	CreatedByID          string `json:"created_by_id"`
	// This field is from variant [IntentCreatedWebhookPayload].
	AuthorizationDetails []IntentAuthorization `json:"authorization_details"`
	// This field is from variant [IntentExecutedWebhookPayload].
	ActionResult BaseActionResult `json:"action_result"`
	// This field is a union of [float64], [string], [string], [string], [string],
	// [string], [string], [string]
	RejectedAt UnsafeUnwrapWebhookEventUnionRejectedAt `json:"rejected_at"`
	Method     string                                  `json:"method"`
	UserID     string                                  `json:"user_id"`
	Changes    []any                                   `json:"changes"`
	// This field is a union of [OrganizationKYBUpdatedData], [UserKYCUpdatedData],
	// [DepositCompletedData], [DepositFailedData], [DepositStartedData]
	Data        UnsafeUnwrapWebhookEventUnionData `json:"data"`
	Environment string                            `json:"environment"`
	// This field is from variant [OrganizationKYBUpdatedWebhookEvent].
	OrganizationID  string `json:"organization_id"`
	Provider        string `json:"provider"`
	Caip2           string `json:"caip2"`
	TransactionHash string `json:"transaction_hash"`
	TransactionID   string `json:"transaction_id"`
	WalletID        string `json:"wallet_id"`
	ReferenceID     string `json:"reference_id"`
	// This field is from variant [TransactionStillPendingWebhookPayload].
	TransactionRequest UnsignedStandardEthereumTransactionResp `json:"transaction_request"`
	AmountUsd          string                                  `json:"amount_usd"`
	EventID            string                                  `json:"event_id"`
	RecordedAt         int64                                   `json:"recorded_at"`
	SourceID           string                                  `json:"source_id"`
	// This field is from variant [UsageCrossChainFeeRecordedWebhookPayload].
	SourceType UsageSourceType `json:"source_type"`
	// This field is from variant [UserAuthenticatedWebhookPayload].
	Account LinkedAccountUnion `json:"account"`
	// This field is from variant [UserAuthenticatedWebhookPayload].
	User User `json:"user"`
	// This field is from variant [UserTransferredAccountWebhookPayload].
	DeletedUser bool `json:"deletedUser"`
	// This field is from variant [UserTransferredAccountWebhookPayload].
	FromUser UserReference `json:"fromUser"`
	// This field is from variant [UserTransferredAccountWebhookPayload].
	ToUser User `json:"toUser"`
	// This field is from variant [UserWalletCreatedWebhookPayload].
	Wallet LinkedAccountBaseWallet `json:"wallet"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	ActualGasCost string `json:"actual_gas_cost"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	ActualGasUsed string `json:"actual_gas_used"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	BlockNumber float64 `json:"block_number"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	LogIndex float64 `json:"log_index"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	Nonce string `json:"nonce"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	Paymaster string `json:"paymaster"`
	Sender    string `json:"sender"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	Success bool `json:"success"`
	// This field is from variant [UserOperationCompletedWebhookPayload].
	UserOpHash string `json:"user_op_hash"`
	// This field is from variant [WalletArchivedWebhookPayload].
	ArchivedAt        float64 `json:"archived_at"`
	ChainType         string  `json:"chain_type"`
	WalletAddress     string  `json:"wallet_address"`
	DepositAccountID  string  `json:"deposit_account_id"`
	DepositType       string  `json:"deposit_type"`
	ProviderDepositID string  `json:"provider_deposit_id"`
	Amount            string  `json:"amount"`
	// This field is a union of [WalletFundsAssetUnion], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string]
	Asset UnsafeUnwrapWebhookEventUnionAsset `json:"asset"`
	// This field is from variant [FundsDepositedWebhookPayload].
	Block          BlockInfo `json:"block"`
	IdempotencyKey string    `json:"idempotency_key"`
	Recipient      string    `json:"recipient"`
	// This field is from variant [FundsDepositedWebhookPayload].
	BridgeMetadata BridgeMetadataUnion `json:"bridge_metadata"`
	TransactionFee string              `json:"transaction_fee"`
	// This field is from variant [PrivateKeyExportWebhookPayload].
	ExportSource ExportType `json:"export_source"`
	// This field is from variant [WalletActionEarnDepositCreatedWebhookPayload].
	ActionType     WalletActionType `json:"action_type"`
	AssetAddress   string           `json:"asset_address"`
	RawAmount      string           `json:"raw_amount"`
	VaultAddress   string           `json:"vault_address"`
	VaultID        string           `json:"vault_id"`
	WalletActionID string           `json:"wallet_action_id"`
	Decimals       int64            `json:"decimals"`
	FailedAt       string           `json:"failed_at"`
	// This field is from variant [WalletActionEarnDepositFailedWebhookPayload].
	FailureReason FailureReason           `json:"failure_reason"`
	Steps         []WalletActionStepUnion `json:"steps"`
	CompletedAt   string                  `json:"completed_at"`
	ShareAmount   string                  `json:"share_amount"`
	Chain         string                  `json:"chain"`
	// This field is a union of [[]EarnIncetiveClaimRewardEntry],
	// [[]EarnIncetiveClaimRewardEntry], [[]EarnIncetiveClaimRewardEntry],
	// [[]EarnIncetiveClaimRewardEntry], [[]YieldClaimReward]
	Rewards                  UnsafeUnwrapWebhookEventUnionRewards `json:"rewards"`
	DestinationCurrency      string                               `json:"destination_currency"`
	DestinationFiatAccountID string                               `json:"destination_fiat_account_id"`
	DestinationPaymentRail   string                               `json:"destination_payment_rail"`
	SourceAmount             string                               `json:"source_amount"`
	SourceAsset              string                               `json:"source_asset"`
	SourceChain              string                               `json:"source_chain"`
	InputAmount              string                               `json:"input_amount"`
	InputToken               string                               `json:"input_token"`
	OutputToken              string                               `json:"output_token"`
	// This field is from variant [WalletActionSwapSucceededWebhookPayload].
	OutputAmount        string `json:"output_amount"`
	DestinationAddress  string `json:"destination_address"`
	SourceAssetAddress  string `json:"source_asset_address"`
	SourceAssetDecimals int64  `json:"source_asset_decimals"`
	// This field is from variant [WalletAutomationSubmittedWebhookPayload].
	ActionID string `json:"action_id"`
	// This field is from variant [WalletAutomationSubmittedWebhookPayload].
	AutomationID string `json:"automation_id"`
	// This field is from variant [WalletAutomationSubmittedWebhookPayload].
	TriggerAssetAddress string `json:"trigger_asset_address"`
	// This field is from variant [WalletAutomationSubmittedWebhookPayload].
	TriggerCaip2 string `json:"trigger_caip2"`
	// This field is from variant [WalletAutomationSubmittedWebhookPayload].
	TriggerID string `json:"trigger_id"`
	Assets    string `json:"assets"`
	Owner     string `json:"owner"`
	Shares    string `json:"shares"`
	// This field is from variant [YieldWithdrawConfirmedWebhookPayload].
	Receiver string `json:"receiver"`
	JSON     struct {
		AuthorizedAt             respjson.Field
		CreatedAt                respjson.Field
		ExpiresAt                respjson.Field
		IntentID                 respjson.Field
		IntentType               respjson.Field
		Member                   respjson.Field
		Status                   respjson.Field
		Type                     respjson.Field
		CreatedByDisplayName     respjson.Field
		CreatedByID              respjson.Field
		AuthorizationDetails     respjson.Field
		ActionResult             respjson.Field
		RejectedAt               respjson.Field
		Method                   respjson.Field
		UserID                   respjson.Field
		Changes                  respjson.Field
		Data                     respjson.Field
		Environment              respjson.Field
		OrganizationID           respjson.Field
		Provider                 respjson.Field
		Caip2                    respjson.Field
		TransactionHash          respjson.Field
		TransactionID            respjson.Field
		WalletID                 respjson.Field
		ReferenceID              respjson.Field
		TransactionRequest       respjson.Field
		AmountUsd                respjson.Field
		EventID                  respjson.Field
		RecordedAt               respjson.Field
		SourceID                 respjson.Field
		SourceType               respjson.Field
		Account                  respjson.Field
		User                     respjson.Field
		DeletedUser              respjson.Field
		FromUser                 respjson.Field
		ToUser                   respjson.Field
		Wallet                   respjson.Field
		ActualGasCost            respjson.Field
		ActualGasUsed            respjson.Field
		BlockNumber              respjson.Field
		LogIndex                 respjson.Field
		Nonce                    respjson.Field
		Paymaster                respjson.Field
		Sender                   respjson.Field
		Success                  respjson.Field
		UserOpHash               respjson.Field
		ArchivedAt               respjson.Field
		ChainType                respjson.Field
		WalletAddress            respjson.Field
		DepositAccountID         respjson.Field
		DepositType              respjson.Field
		ProviderDepositID        respjson.Field
		Amount                   respjson.Field
		Asset                    respjson.Field
		Block                    respjson.Field
		IdempotencyKey           respjson.Field
		Recipient                respjson.Field
		BridgeMetadata           respjson.Field
		TransactionFee           respjson.Field
		ExportSource             respjson.Field
		ActionType               respjson.Field
		AssetAddress             respjson.Field
		RawAmount                respjson.Field
		VaultAddress             respjson.Field
		VaultID                  respjson.Field
		WalletActionID           respjson.Field
		Decimals                 respjson.Field
		FailedAt                 respjson.Field
		FailureReason            respjson.Field
		Steps                    respjson.Field
		CompletedAt              respjson.Field
		ShareAmount              respjson.Field
		Chain                    respjson.Field
		Rewards                  respjson.Field
		DestinationCurrency      respjson.Field
		DestinationFiatAccountID respjson.Field
		DestinationPaymentRail   respjson.Field
		SourceAmount             respjson.Field
		SourceAsset              respjson.Field
		SourceChain              respjson.Field
		InputAmount              respjson.Field
		InputToken               respjson.Field
		OutputToken              respjson.Field
		OutputAmount             respjson.Field
		DestinationAddress       respjson.Field
		SourceAssetAddress       respjson.Field
		SourceAssetDecimals      respjson.Field
		ActionID                 respjson.Field
		AutomationID             respjson.Field
		TriggerAssetAddress      respjson.Field
		TriggerCaip2             respjson.Field
		TriggerID                respjson.Field
		Assets                   respjson.Field
		Owner                    respjson.Field
		Shares                   respjson.Field
		Receiver                 respjson.Field
		raw                      string
	} `json:"-"`
}

// anyUnsafeUnwrapWebhookEvent is implemented by each variant of
// [UnsafeUnwrapWebhookEventUnion] to add type safety for the return type of
// [UnsafeUnwrapWebhookEventUnion.AsAny]
type anyUnsafeUnwrapWebhookEvent interface {
	implUnsafeUnwrapWebhookEventUnion()
}

func (IntentAuthorizedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                        {}
func (IntentCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                           {}
func (IntentExecutedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (IntentFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                            {}
func (IntentRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (MfaDisabledWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                             {}
func (MfaEnabledWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                              {}
func (OrganizationKYBUpdatedWebhookEvent) implUnsafeUnwrapWebhookEventUnion()                    {}
func (TransactionBroadcastedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (TransactionConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                    {}
func (TransactionExecutionRevertedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()            {}
func (TransactionFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (TransactionProviderErrorWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                {}
func (TransactionReplacedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (TransactionStillPendingWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                 {}
func (UsageCrossChainFeeRecordedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()              {}
func (UsageGasSponsorshipRecordedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()             {}
func (UserAuthenticatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (UserCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                             {}
func (UserDeletedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                             {}
func (UserKYCUpdatedWebhookEvent) implUnsafeUnwrapWebhookEventUnion()                            {}
func (UserLinkedAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (UserTransferredAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (UserUnlinkedAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (UserUpdatedAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                      {}
func (UserWalletCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (UserOperationCompletedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (WalletArchivedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (WalletDepositAccountDepositCompletedWebhookEvent) implUnsafeUnwrapWebhookEventUnion()      {}
func (WalletDepositAccountDepositFailedWebhookEvent) implUnsafeUnwrapWebhookEventUnion()         {}
func (WalletDepositAccountDepositStartedWebhookEvent) implUnsafeUnwrapWebhookEventUnion()        {}
func (FundsDepositedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (FundsWithdrawnWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (PrivateKeyExportWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                        {}
func (WalletRecoveredWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                         {}
func (WalletRecoverySetupWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (WalletRestoredWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (WalletActionEarnDepositCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()          {}
func (WalletActionEarnDepositFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()           {}
func (WalletActionEarnDepositRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()         {}
func (WalletActionEarnDepositSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()        {}
func (WalletActionEarnFeeCollectCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()       {}
func (WalletActionEarnFeeCollectFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()        {}
func (WalletActionEarnFeeCollectRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()      {}
func (WalletActionEarnFeeCollectSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()     {}
func (WalletActionEarnIncentiveClaimCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()   {}
func (WalletActionEarnIncentiveClaimFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()    {}
func (WalletActionEarnIncentiveClaimRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()  {}
func (WalletActionEarnIncentiveClaimSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion() {}
func (WalletActionEarnWithdrawCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()         {}
func (WalletActionEarnWithdrawFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()          {}
func (WalletActionEarnWithdrawRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()        {}
func (WalletActionEarnWithdrawSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()       {}
func (WalletActionPayoutCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()               {}
func (WalletActionPayoutFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                {}
func (WalletActionPayoutRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()              {}
func (WalletActionPayoutSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()             {}
func (WalletActionSwapCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                 {}
func (WalletActionSwapFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (WalletActionSwapRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                {}
func (WalletActionSwapSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()               {}
func (WalletActionTransferCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()             {}
func (WalletActionTransferFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()              {}
func (WalletActionTransferRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()            {}
func (WalletActionTransferSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()           {}
func (WalletAutomationSubmittedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()               {}
func (YieldClaimConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (YieldDepositConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                   {}
func (YieldWithdrawConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UnsafeUnwrapWebhookEventUnion.AsAny().(type) {
//	case privyclient.IntentAuthorizedWebhookPayload:
//	case privyclient.IntentCreatedWebhookPayload:
//	case privyclient.IntentExecutedWebhookPayload:
//	case privyclient.IntentFailedWebhookPayload:
//	case privyclient.IntentRejectedWebhookPayload:
//	case privyclient.MfaDisabledWebhookPayload:
//	case privyclient.MfaEnabledWebhookPayload:
//	case privyclient.OrganizationKYBUpdatedWebhookEvent:
//	case privyclient.TransactionBroadcastedWebhookPayload:
//	case privyclient.TransactionConfirmedWebhookPayload:
//	case privyclient.TransactionExecutionRevertedWebhookPayload:
//	case privyclient.TransactionFailedWebhookPayload:
//	case privyclient.TransactionProviderErrorWebhookPayload:
//	case privyclient.TransactionReplacedWebhookPayload:
//	case privyclient.TransactionStillPendingWebhookPayload:
//	case privyclient.UsageCrossChainFeeRecordedWebhookPayload:
//	case privyclient.UsageGasSponsorshipRecordedWebhookPayload:
//	case privyclient.UserAuthenticatedWebhookPayload:
//	case privyclient.UserCreatedWebhookPayload:
//	case privyclient.UserDeletedWebhookPayload:
//	case privyclient.UserKYCUpdatedWebhookEvent:
//	case privyclient.UserLinkedAccountWebhookPayload:
//	case privyclient.UserTransferredAccountWebhookPayload:
//	case privyclient.UserUnlinkedAccountWebhookPayload:
//	case privyclient.UserUpdatedAccountWebhookPayload:
//	case privyclient.UserWalletCreatedWebhookPayload:
//	case privyclient.UserOperationCompletedWebhookPayload:
//	case privyclient.WalletArchivedWebhookPayload:
//	case privyclient.WalletDepositAccountDepositCompletedWebhookEvent:
//	case privyclient.WalletDepositAccountDepositFailedWebhookEvent:
//	case privyclient.WalletDepositAccountDepositStartedWebhookEvent:
//	case privyclient.FundsDepositedWebhookPayload:
//	case privyclient.FundsWithdrawnWebhookPayload:
//	case privyclient.PrivateKeyExportWebhookPayload:
//	case privyclient.WalletRecoveredWebhookPayload:
//	case privyclient.WalletRecoverySetupWebhookPayload:
//	case privyclient.WalletRestoredWebhookPayload:
//	case privyclient.WalletActionEarnDepositCreatedWebhookPayload:
//	case privyclient.WalletActionEarnDepositFailedWebhookPayload:
//	case privyclient.WalletActionEarnDepositRejectedWebhookPayload:
//	case privyclient.WalletActionEarnDepositSucceededWebhookPayload:
//	case privyclient.WalletActionEarnFeeCollectCreatedWebhookPayload:
//	case privyclient.WalletActionEarnFeeCollectFailedWebhookPayload:
//	case privyclient.WalletActionEarnFeeCollectRejectedWebhookPayload:
//	case privyclient.WalletActionEarnFeeCollectSucceededWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimCreatedWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimFailedWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimRejectedWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimSucceededWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawCreatedWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawFailedWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawRejectedWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawSucceededWebhookPayload:
//	case privyclient.WalletActionPayoutCreatedWebhookPayload:
//	case privyclient.WalletActionPayoutFailedWebhookPayload:
//	case privyclient.WalletActionPayoutRejectedWebhookPayload:
//	case privyclient.WalletActionPayoutSucceededWebhookPayload:
//	case privyclient.WalletActionSwapCreatedWebhookPayload:
//	case privyclient.WalletActionSwapFailedWebhookPayload:
//	case privyclient.WalletActionSwapRejectedWebhookPayload:
//	case privyclient.WalletActionSwapSucceededWebhookPayload:
//	case privyclient.WalletActionTransferCreatedWebhookPayload:
//	case privyclient.WalletActionTransferFailedWebhookPayload:
//	case privyclient.WalletActionTransferRejectedWebhookPayload:
//	case privyclient.WalletActionTransferSucceededWebhookPayload:
//	case privyclient.WalletAutomationSubmittedWebhookPayload:
//	case privyclient.YieldClaimConfirmedWebhookPayload:
//	case privyclient.YieldDepositConfirmedWebhookPayload:
//	case privyclient.YieldWithdrawConfirmedWebhookPayload:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UnsafeUnwrapWebhookEventUnion) AsAny() anyUnsafeUnwrapWebhookEvent {
	switch u.Type {
	case "intent.authorized":
		return u.AsIntentAuthorized()
	case "intent.created":
		return u.AsIntentCreated()
	case "intent.executed":
		return u.AsIntentExecuted()
	case "intent.failed":
		return u.AsIntentFailed()
	case "intent.rejected":
		return u.AsIntentRejected()
	case "mfa.disabled":
		return u.AsMfaDisabled()
	case "mfa.enabled":
		return u.AsMfaEnabled()
	case "organization.kyb.updated":
		return u.AsOrganizationKYBUpdated()
	case "transaction.broadcasted":
		return u.AsTransactionBroadcasted()
	case "transaction.confirmed":
		return u.AsTransactionConfirmed()
	case "transaction.execution_reverted":
		return u.AsTransactionExecutionReverted()
	case "transaction.failed":
		return u.AsTransactionFailed()
	case "transaction.provider_error":
		return u.AsTransactionProviderError()
	case "transaction.replaced":
		return u.AsTransactionReplaced()
	case "transaction.still_pending":
		return u.AsTransactionStillPending()
	case "usage.cross_chain_fee.recorded":
		return u.AsUsageCrossChainFeeRecorded()
	case "usage.gas_sponsorship.recorded":
		return u.AsUsageGasSponsorshipRecorded()
	case "user.authenticated":
		return u.AsUserAuthenticated()
	case "user.created":
		return u.AsUserCreated()
	case "user.deleted":
		return u.AsUserDeleted()
	case "user.kyc.updated":
		return u.AsUserKYCUpdated()
	case "user.linked_account":
		return u.AsUserLinkedAccount()
	case "user.transferred_account":
		return u.AsUserTransferredAccount()
	case "user.unlinked_account":
		return u.AsUserUnlinkedAccount()
	case "user.updated_account":
		return u.AsUserUpdatedAccount()
	case "user.wallet_created":
		return u.AsUserWalletCreated()
	case "user_operation.completed":
		return u.AsUserOperationCompleted()
	case "wallet.archived":
		return u.AsWalletArchived()
	case "wallet.deposit_account.deposit_completed":
		return u.AsWalletDepositAccountDepositCompleted()
	case "wallet.deposit_account.deposit_failed":
		return u.AsWalletDepositAccountDepositFailed()
	case "wallet.deposit_account.deposit_started":
		return u.AsWalletDepositAccountDepositStarted()
	case "wallet.funds_deposited":
		return u.AsWalletFundsDeposited()
	case "wallet.funds_withdrawn":
		return u.AsWalletFundsWithdrawn()
	case "wallet.private_key_export":
		return u.AsWalletPrivateKeyExport()
	case "wallet.recovered":
		return u.AsWalletRecovered()
	case "wallet.recovery_setup":
		return u.AsWalletRecoverySetup()
	case "wallet.restored":
		return u.AsWalletRestored()
	case "wallet_action.earn_deposit.created":
		return u.AsWalletActionEarnDepositCreated()
	case "wallet_action.earn_deposit.failed":
		return u.AsWalletActionEarnDepositFailed()
	case "wallet_action.earn_deposit.rejected":
		return u.AsWalletActionEarnDepositRejected()
	case "wallet_action.earn_deposit.succeeded":
		return u.AsWalletActionEarnDepositSucceeded()
	case "wallet_action.earn_fee_collect.created":
		return u.AsWalletActionEarnFeeCollectCreated()
	case "wallet_action.earn_fee_collect.failed":
		return u.AsWalletActionEarnFeeCollectFailed()
	case "wallet_action.earn_fee_collect.rejected":
		return u.AsWalletActionEarnFeeCollectRejected()
	case "wallet_action.earn_fee_collect.succeeded":
		return u.AsWalletActionEarnFeeCollectSucceeded()
	case "wallet_action.earn_incentive_claim.created":
		return u.AsWalletActionEarnIncentiveClaimCreated()
	case "wallet_action.earn_incentive_claim.failed":
		return u.AsWalletActionEarnIncentiveClaimFailed()
	case "wallet_action.earn_incentive_claim.rejected":
		return u.AsWalletActionEarnIncentiveClaimRejected()
	case "wallet_action.earn_incentive_claim.succeeded":
		return u.AsWalletActionEarnIncentiveClaimSucceeded()
	case "wallet_action.earn_withdraw.created":
		return u.AsWalletActionEarnWithdrawCreated()
	case "wallet_action.earn_withdraw.failed":
		return u.AsWalletActionEarnWithdrawFailed()
	case "wallet_action.earn_withdraw.rejected":
		return u.AsWalletActionEarnWithdrawRejected()
	case "wallet_action.earn_withdraw.succeeded":
		return u.AsWalletActionEarnWithdrawSucceeded()
	case "wallet_action.payout.created":
		return u.AsWalletActionPayoutCreated()
	case "wallet_action.payout.failed":
		return u.AsWalletActionPayoutFailed()
	case "wallet_action.payout.rejected":
		return u.AsWalletActionPayoutRejected()
	case "wallet_action.payout.succeeded":
		return u.AsWalletActionPayoutSucceeded()
	case "wallet_action.swap.created":
		return u.AsWalletActionSwapCreated()
	case "wallet_action.swap.failed":
		return u.AsWalletActionSwapFailed()
	case "wallet_action.swap.rejected":
		return u.AsWalletActionSwapRejected()
	case "wallet_action.swap.succeeded":
		return u.AsWalletActionSwapSucceeded()
	case "wallet_action.transfer.created":
		return u.AsWalletActionTransferCreated()
	case "wallet_action.transfer.failed":
		return u.AsWalletActionTransferFailed()
	case "wallet_action.transfer.rejected":
		return u.AsWalletActionTransferRejected()
	case "wallet_action.transfer.succeeded":
		return u.AsWalletActionTransferSucceeded()
	case "wallet_automation.submitted":
		return u.AsWalletAutomationSubmitted()
	case "yield.claim.confirmed":
		return u.AsYieldClaimConfirmed()
	case "yield.deposit.confirmed":
		return u.AsYieldDepositConfirmed()
	case "yield.withdraw.confirmed":
		return u.AsYieldWithdrawConfirmed()
	}
	return nil
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentAuthorized() (v IntentAuthorizedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentCreated() (v IntentCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentExecuted() (v IntentExecutedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentFailed() (v IntentFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentRejected() (v IntentRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsMfaDisabled() (v MfaDisabledWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsMfaEnabled() (v MfaEnabledWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsOrganizationKYBUpdated() (v OrganizationKYBUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionBroadcasted() (v TransactionBroadcastedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionConfirmed() (v TransactionConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionExecutionReverted() (v TransactionExecutionRevertedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionFailed() (v TransactionFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionProviderError() (v TransactionProviderErrorWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionReplaced() (v TransactionReplacedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionStillPending() (v TransactionStillPendingWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUsageCrossChainFeeRecorded() (v UsageCrossChainFeeRecordedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUsageGasSponsorshipRecorded() (v UsageGasSponsorshipRecordedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserAuthenticated() (v UserAuthenticatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserCreated() (v UserCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserDeleted() (v UserDeletedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserKYCUpdated() (v UserKYCUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserLinkedAccount() (v UserLinkedAccountWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserTransferredAccount() (v UserTransferredAccountWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserUnlinkedAccount() (v UserUnlinkedAccountWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserUpdatedAccount() (v UserUpdatedAccountWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserWalletCreated() (v UserWalletCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserOperationCompleted() (v UserOperationCompletedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletArchived() (v WalletArchivedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletDepositAccountDepositCompleted() (v WalletDepositAccountDepositCompletedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletDepositAccountDepositFailed() (v WalletDepositAccountDepositFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletDepositAccountDepositStarted() (v WalletDepositAccountDepositStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletFundsDeposited() (v FundsDepositedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletFundsWithdrawn() (v FundsWithdrawnWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletPrivateKeyExport() (v PrivateKeyExportWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletRecovered() (v WalletRecoveredWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletRecoverySetup() (v WalletRecoverySetupWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletRestored() (v WalletRestoredWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositCreated() (v WalletActionEarnDepositCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositFailed() (v WalletActionEarnDepositFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositRejected() (v WalletActionEarnDepositRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositSucceeded() (v WalletActionEarnDepositSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnFeeCollectCreated() (v WalletActionEarnFeeCollectCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnFeeCollectFailed() (v WalletActionEarnFeeCollectFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnFeeCollectRejected() (v WalletActionEarnFeeCollectRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnFeeCollectSucceeded() (v WalletActionEarnFeeCollectSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimCreated() (v WalletActionEarnIncentiveClaimCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimFailed() (v WalletActionEarnIncentiveClaimFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimRejected() (v WalletActionEarnIncentiveClaimRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimSucceeded() (v WalletActionEarnIncentiveClaimSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawCreated() (v WalletActionEarnWithdrawCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawFailed() (v WalletActionEarnWithdrawFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawRejected() (v WalletActionEarnWithdrawRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawSucceeded() (v WalletActionEarnWithdrawSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionPayoutCreated() (v WalletActionPayoutCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionPayoutFailed() (v WalletActionPayoutFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionPayoutRejected() (v WalletActionPayoutRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionPayoutSucceeded() (v WalletActionPayoutSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapCreated() (v WalletActionSwapCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapFailed() (v WalletActionSwapFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapRejected() (v WalletActionSwapRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapSucceeded() (v WalletActionSwapSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferCreated() (v WalletActionTransferCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferFailed() (v WalletActionTransferFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferRejected() (v WalletActionTransferRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferSucceeded() (v WalletActionTransferSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletAutomationSubmitted() (v WalletAutomationSubmittedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsYieldClaimConfirmed() (v YieldClaimConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsYieldDepositConfirmed() (v YieldDepositConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsYieldWithdrawConfirmed() (v YieldWithdrawConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnsafeUnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnsafeUnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionCreatedAt is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionCreatedAt provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type UnsafeUnwrapWebhookEventUnionCreatedAt struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionCreatedAt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionRejectedAt is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionRejectedAt
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type UnsafeUnwrapWebhookEventUnionRejectedAt struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionRejectedAt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionData is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionData struct {
	// This field is from variant [OrganizationKYBUpdatedData].
	Capabilities KyxCapabilities  `json:"capabilities"`
	Endorsements []KyxEndorsement `json:"endorsements"`
	// This field is from variant [OrganizationKYBUpdatedData].
	KYB OrganizationKYBUpdatedKYBData `json:"kyb"`
	// This field is from variant [OrganizationKYBUpdatedData].
	Status KyxProviderStatus `json:"status"`
	// This field is a union of [OrganizationKYBUpdatedTosData],
	// [UserKYCUpdatedTosData]
	Tos UnsafeUnwrapWebhookEventUnionDataTos `json:"tos"`
	// This field is from variant [UserKYCUpdatedData].
	KYC       UserKYCUpdatedKYCData `json:"kyc"`
	CreatedAt string                `json:"created_at"`
	// This field is a union of [DepositCompletedDestination],
	// [DepositStartedDestination]
	Destination UnsafeUnwrapWebhookEventUnionDataDestination `json:"destination"`
	// This field is from variant [DepositCompletedData].
	Source DepositStartedSource `json:"source"`
	// This field is from variant [DepositFailedData].
	Reason string `json:"reason"`
	// This field is from variant [DepositFailedData].
	ReasonCode string `json:"reason_code"`
	// This field is from variant [DepositFailedData].
	RefundedAt string `json:"refunded_at"`
	JSON       struct {
		Capabilities respjson.Field
		Endorsements respjson.Field
		KYB          respjson.Field
		Status       respjson.Field
		Tos          respjson.Field
		KYC          respjson.Field
		CreatedAt    respjson.Field
		Destination  respjson.Field
		Source       respjson.Field
		Reason       respjson.Field
		ReasonCode   respjson.Field
		RefundedAt   respjson.Field
		raw          string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataTos is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataTos provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataTos struct {
	// This field is from variant [OrganizationKYBUpdatedTosData].
	Status KyxTosStatus `json:"status"`
	JSON   struct {
		Status respjson.Field
		raw    string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataTos) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataDestination is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataDestination
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataDestination struct {
	// This field is from variant [DepositCompletedDestination].
	Amount string `json:"amount"`
	Asset  string `json:"asset"`
	Chain  string `json:"chain"`
	// This field is from variant [DepositCompletedDestination].
	TransactionHash string `json:"transaction_hash"`
	JSON            struct {
		Amount          respjson.Field
		Asset           respjson.Field
		Chain           respjson.Field
		TransactionHash respjson.Field
		raw             string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionAsset is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionAsset provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type UnsafeUnwrapWebhookEventUnionAsset struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is a union of [any], [string], [string], [string]
	Address UnsafeUnwrapWebhookEventUnionAssetAddress `json:"address"`
	Type    string                                    `json:"type"`
	// This field is from variant [WalletFundsAssetUnion].
	Mint string `json:"mint"`
	JSON struct {
		OfString respjson.Field
		Address  respjson.Field
		Type     respjson.Field
		Mint     respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionAssetAddress is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionAssetAddress
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfWalletFundsNativeTokenAssetAddress OfString]
type UnsafeUnwrapWebhookEventUnionAssetAddress struct {
	// This field will be present if the value is a [any] instead of an object.
	OfWalletFundsNativeTokenAssetAddress any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfWalletFundsNativeTokenAssetAddress respjson.Field
		OfString                             respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionAssetAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionRewards is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionRewards provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEarnIncetiveClaimRewardEntryArray OfYieldClaimRewardArray]
type UnsafeUnwrapWebhookEventUnionRewards struct {
	// This field will be present if the value is a [[]EarnIncetiveClaimRewardEntry]
	// instead of an object.
	OfEarnIncetiveClaimRewardEntryArray []EarnIncetiveClaimRewardEntry `json:",inline"`
	// This field will be present if the value is a [[]YieldClaimReward] instead of an
	// object.
	OfYieldClaimRewardArray []YieldClaimReward `json:",inline"`
	JSON                    struct {
		OfEarnIncetiveClaimRewardEntryArray respjson.Field
		OfYieldClaimRewardArray             respjson.Field
		raw                                 string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionRewards) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
