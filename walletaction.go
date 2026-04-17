// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"
	"time"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// WalletActionService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletActionService] method instead.
type WalletActionService struct {
	Options []option.RequestOption
}

// NewWalletActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWalletActionService(opts ...option.RequestOption) (r WalletActionService) {
	r = WalletActionService{}
	r.Options = opts
	return
}

// Status of a wallet action.
type WalletActionStatus string

const (
	WalletActionStatusPending   WalletActionStatus = "pending"
	WalletActionStatusSucceeded WalletActionStatus = "succeeded"
	WalletActionStatusRejected  WalletActionStatus = "rejected"
	WalletActionStatusFailed    WalletActionStatus = "failed"
)

// Status of an EVM step in a wallet action.
type EvmWalletActionStepStatus string

const (
	EvmWalletActionStepStatusPreparing EvmWalletActionStepStatus = "preparing"
	EvmWalletActionStepStatusQueued    EvmWalletActionStepStatus = "queued"
	EvmWalletActionStepStatusPending   EvmWalletActionStepStatus = "pending"
	EvmWalletActionStepStatusRetrying  EvmWalletActionStepStatus = "retrying"
	EvmWalletActionStepStatusConfirmed EvmWalletActionStepStatus = "confirmed"
	EvmWalletActionStepStatusRejected  EvmWalletActionStepStatus = "rejected"
	EvmWalletActionStepStatusReverted  EvmWalletActionStepStatus = "reverted"
	EvmWalletActionStepStatusReplaced  EvmWalletActionStepStatus = "replaced"
	EvmWalletActionStepStatusAbandoned EvmWalletActionStepStatus = "abandoned"
)

// A description of why a wallet action (or a step within a wallet action) failed.
type FailureReason struct {
	// Human-readable failure message.
	Message string `json:"message" api:"required"`
	// Additional error details, if available.
	Details any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FailureReason) RawJSON() string { return r.JSON.raw }
func (r *FailureReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A wallet action step consisting of an EVM transaction.
type EvmTransactionWalletActionStep struct {
	// CAIP-2 chain identifier of the transaction, containing the chain ID.
	Caip2 string `json:"caip2" api:"required"`
	// Status of an EVM step in a wallet action.
	//
	// Any of "preparing", "queued", "pending", "retrying", "confirmed", "rejected",
	// "reverted", "replaced", "abandoned".
	Status EvmWalletActionStepStatus `json:"status" api:"required"`
	// The transaction hash for this step. May change while the step status is
	// non-terminal.
	TransactionHash string `json:"transaction_hash" api:"required"`
	// Any of "evm_transaction".
	Type EvmTransactionWalletActionStepType `json:"type" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		Status          respjson.Field
		TransactionHash respjson.Field
		Type            respjson.Field
		FailureReason   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvmTransactionWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *EvmTransactionWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvmTransactionWalletActionStepType string

const (
	EvmTransactionWalletActionStepTypeEvmTransaction EvmTransactionWalletActionStepType = "evm_transaction"
)

// A wallet action step consisting of an EVM user operation.
type EvmUserOperationWalletActionStep struct {
	// Transaction hash of the bundle in which this user operation was included. Null
	// until included by a bundler.
	BundleTransactionHash string `json:"bundle_transaction_hash" api:"required"`
	// CAIP-2 network identifier, containing the chain ID of the user operation.
	Caip2 string `json:"caip2" api:"required"`
	// The entrypoint version of the user operation.
	//
	// Any of "0.6", "0.7", "0.8", "0.9".
	EntrypointVersion EvmUserOperationWalletActionStepEntrypointVersion `json:"entrypoint_version" api:"required"`
	// Status of an EVM step in a wallet action.
	//
	// Any of "preparing", "queued", "pending", "retrying", "confirmed", "rejected",
	// "reverted", "replaced", "abandoned".
	Status EvmWalletActionStepStatus `json:"status" api:"required"`
	// Any of "evm_user_operation".
	Type EvmUserOperationWalletActionStepType `json:"type" api:"required"`
	// The user operation hash for this step. May change while the step status is
	// non-terminal.
	UserOperationHash string `json:"user_operation_hash" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BundleTransactionHash respjson.Field
		Caip2                 respjson.Field
		EntrypointVersion     respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		UserOperationHash     respjson.Field
		FailureReason         respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvmUserOperationWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *EvmUserOperationWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The entrypoint version of the user operation.
type EvmUserOperationWalletActionStepEntrypointVersion string

const (
	EvmUserOperationWalletActionStepEntrypointVersion0_6 EvmUserOperationWalletActionStepEntrypointVersion = "0.6"
	EvmUserOperationWalletActionStepEntrypointVersion0_7 EvmUserOperationWalletActionStepEntrypointVersion = "0.7"
	EvmUserOperationWalletActionStepEntrypointVersion0_8 EvmUserOperationWalletActionStepEntrypointVersion = "0.8"
	EvmUserOperationWalletActionStepEntrypointVersion0_9 EvmUserOperationWalletActionStepEntrypointVersion = "0.9"
)

type EvmUserOperationWalletActionStepType string

const (
	EvmUserOperationWalletActionStepTypeEvmUserOperation EvmUserOperationWalletActionStepType = "evm_user_operation"
)

// Status of an SVM step in a wallet action.
type SvmWalletActionStepStatus string

const (
	SvmWalletActionStepStatusPreparing SvmWalletActionStepStatus = "preparing"
	SvmWalletActionStepStatusQueued    SvmWalletActionStepStatus = "queued"
	SvmWalletActionStepStatusPending   SvmWalletActionStepStatus = "pending"
	SvmWalletActionStepStatusConfirmed SvmWalletActionStepStatus = "confirmed"
	SvmWalletActionStepStatusFinalized SvmWalletActionStepStatus = "finalized"
	SvmWalletActionStepStatusRejected  SvmWalletActionStepStatus = "rejected"
	SvmWalletActionStepStatusReverted  SvmWalletActionStepStatus = "reverted"
	SvmWalletActionStepStatusFailed    SvmWalletActionStepStatus = "failed"
)

// A wallet action step consisting of an SVM (Solana) transaction.
type SvmTransactionWalletActionStep struct {
	// CAIP-2 chain identifier for the Solana network.
	Caip2 string `json:"caip2" api:"required"`
	// Status of an SVM step in a wallet action.
	//
	// Any of "preparing", "queued", "pending", "confirmed", "finalized", "rejected",
	// "reverted", "failed".
	Status SvmWalletActionStepStatus `json:"status" api:"required"`
	// The Solana transaction signature (base58-encoded). Null until broadcast.
	TransactionSignature string `json:"transaction_signature" api:"required"`
	// Any of "svm_transaction".
	Type SvmTransactionWalletActionStepType `json:"type" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2                respjson.Field
		Status               respjson.Field
		TransactionSignature respjson.Field
		Type                 respjson.Field
		FailureReason        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SvmTransactionWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *SvmTransactionWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SvmTransactionWalletActionStepType string

const (
	SvmTransactionWalletActionStepTypeSvmTransaction SvmTransactionWalletActionStepType = "svm_transaction"
)

// WalletActionStepUnion contains all possible properties and values from
// [EvmTransactionWalletActionStep], [EvmUserOperationWalletActionStep],
// [SvmTransactionWalletActionStep].
//
// Use the [WalletActionStepUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletActionStepUnion struct {
	Caip2  string `json:"caip2"`
	Status string `json:"status"`
	// This field is from variant [EvmTransactionWalletActionStep].
	TransactionHash string `json:"transaction_hash"`
	// Any of "evm_transaction", "evm_user_operation", "svm_transaction".
	Type string `json:"type"`
	// This field is from variant [EvmTransactionWalletActionStep].
	FailureReason FailureReason `json:"failure_reason"`
	// This field is from variant [EvmUserOperationWalletActionStep].
	BundleTransactionHash string `json:"bundle_transaction_hash"`
	// This field is from variant [EvmUserOperationWalletActionStep].
	EntrypointVersion EvmUserOperationWalletActionStepEntrypointVersion `json:"entrypoint_version"`
	// This field is from variant [EvmUserOperationWalletActionStep].
	UserOperationHash string `json:"user_operation_hash"`
	// This field is from variant [SvmTransactionWalletActionStep].
	TransactionSignature string `json:"transaction_signature"`
	JSON                 struct {
		Caip2                 respjson.Field
		Status                respjson.Field
		TransactionHash       respjson.Field
		Type                  respjson.Field
		FailureReason         respjson.Field
		BundleTransactionHash respjson.Field
		EntrypointVersion     respjson.Field
		UserOperationHash     respjson.Field
		TransactionSignature  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyWalletActionStep is implemented by each variant of [WalletActionStepUnion] to
// add type safety for the return type of [WalletActionStepUnion.AsAny]
type anyWalletActionStep interface {
	implWalletActionStepUnion()
}

func (EvmTransactionWalletActionStep) implWalletActionStepUnion()   {}
func (EvmUserOperationWalletActionStep) implWalletActionStepUnion() {}
func (SvmTransactionWalletActionStep) implWalletActionStepUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletActionStepUnion.AsAny().(type) {
//	case privyclient.EvmTransactionWalletActionStep:
//	case privyclient.EvmUserOperationWalletActionStep:
//	case privyclient.SvmTransactionWalletActionStep:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletActionStepUnion) AsAny() anyWalletActionStep {
	switch u.Type {
	case "evm_transaction":
		return u.AsEvmTransaction()
	case "evm_user_operation":
		return u.AsEvmUserOperation()
	case "svm_transaction":
		return u.AsSvmTransaction()
	}
	return nil
}

func (u WalletActionStepUnion) AsEvmTransaction() (v EvmTransactionWalletActionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionStepUnion) AsEvmUserOperation() (v EvmUserOperationWalletActionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionStepUnion) AsSvmTransaction() (v SvmTransactionWalletActionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletActionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletActionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a transfer action.
type TransferActionResponse struct {
	// The ID of the wallet action.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// Decimal amount as the user provided (e.g. "1.5").
	SourceAmount string `json:"source_amount" api:"required"`
	// Asset identifier (e.g. "usdc", "eth").
	SourceAsset string `json:"source_asset" api:"required"`
	// Chain name (e.g. "base", "ethereum").
	SourceChain string `json:"source_chain" api:"required"`
	// Status of a wallet action.
	//
	// Any of "pending", "succeeded", "rejected", "failed".
	Status WalletActionStatus `json:"status" api:"required"`
	// Any of "transfer".
	Type TransferActionResponseType `json:"type" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// The steps of the wallet action. Only returned if `?include=steps` is provided.
	Steps []WalletActionStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		DestinationAddress respjson.Field
		SourceAmount       respjson.Field
		SourceAsset        respjson.Field
		SourceChain        respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		WalletID           respjson.Field
		FailureReason      respjson.Field
		Steps              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferActionResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferActionResponseType string

const (
	TransferActionResponseTypeTransfer TransferActionResponseType = "transfer"
)

// A specific reward token and amount associated with an earn incentive claim.
type EarnIncetiveClaimRewardEntry struct {
	// Claimable amount in base units.
	Amount string `json:"amount" api:"required"`
	// Address of the reward token.
	TokenAddress string `json:"token_address" api:"required"`
	// Symbol of the reward token (e.g. "MORPHO").
	TokenSymbol string `json:"token_symbol" api:"required"`
	// Number of decimal places for the reward token.
	TokenDecimals int64 `json:"token_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		TokenAddress  respjson.Field
		TokenSymbol   respjson.Field
		TokenDecimals respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EarnIncetiveClaimRewardEntry) RawJSON() string { return r.JSON.raw }
func (r *EarnIncetiveClaimRewardEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for an earn deposit action.
type EarnDepositActionResponse struct {
	// The ID of the wallet action.
	ID string `json:"id" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Base-unit amount of asset deposited (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// Vault shares received in base units. Populated after on-chain confirmation.
	ShareAmount string `json:"share_amount" api:"required"`
	// Status of a wallet action.
	//
	// Any of "pending", "succeeded", "rejected", "failed".
	Status WalletActionStatus `json:"status" api:"required"`
	// Any of "earn_deposit".
	Type EarnDepositActionResponseType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
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
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// The steps of the wallet action. Only returned if `?include=steps` is provided.
	Steps []WalletActionStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AssetAddress  respjson.Field
		Caip2         respjson.Field
		CreatedAt     respjson.Field
		RawAmount     respjson.Field
		ShareAmount   respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		VaultAddress  respjson.Field
		VaultID       respjson.Field
		WalletID      respjson.Field
		Amount        respjson.Field
		Asset         respjson.Field
		Decimals      respjson.Field
		FailureReason respjson.Field
		Steps         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EarnDepositActionResponse) RawJSON() string { return r.JSON.raw }
func (r *EarnDepositActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EarnDepositActionResponseType string

const (
	EarnDepositActionResponseTypeEarnDeposit EarnDepositActionResponseType = "earn_deposit"
)

// Response for an earn withdraw action.
type EarnWithdrawActionResponse struct {
	// The ID of the wallet action.
	ID string `json:"id" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Base-unit amount of asset withdrawn (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
	// Vault shares burned in base units. Populated after on-chain confirmation.
	ShareAmount string `json:"share_amount" api:"required"`
	// Status of a wallet action.
	//
	// Any of "pending", "succeeded", "rejected", "failed".
	Status WalletActionStatus `json:"status" api:"required"`
	// Any of "earn_withdraw".
	Type EarnWithdrawActionResponseType `json:"type" api:"required"`
	// ERC-4626 vault contract address.
	VaultAddress string `json:"vault_address" api:"required"`
	// The vault ID.
	VaultID string `json:"vault_id" api:"required"`
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
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// The steps of the wallet action. Only returned if `?include=steps` is provided.
	Steps []WalletActionStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AssetAddress  respjson.Field
		Caip2         respjson.Field
		CreatedAt     respjson.Field
		RawAmount     respjson.Field
		ShareAmount   respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		VaultAddress  respjson.Field
		VaultID       respjson.Field
		WalletID      respjson.Field
		Amount        respjson.Field
		Asset         respjson.Field
		Decimals      respjson.Field
		FailureReason respjson.Field
		Steps         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EarnWithdrawActionResponse) RawJSON() string { return r.JSON.raw }
func (r *EarnWithdrawActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EarnWithdrawActionResponseType string

const (
	EarnWithdrawActionResponseTypeEarnWithdraw EarnWithdrawActionResponseType = "earn_withdraw"
)

// Response for an earn incentive claim action.
type EarnIncentiveClaimActionResponse struct {
	// The ID of the wallet action.
	ID string `json:"id" api:"required"`
	// EVM chain name (e.g. "base", "ethereum").
	Chain string `json:"chain" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Claimed reward tokens. Populated after the preparation step fetches from Merkl.
	Rewards []EarnIncetiveClaimRewardEntry `json:"rewards" api:"required"`
	// Status of a wallet action.
	//
	// Any of "pending", "succeeded", "rejected", "failed".
	Status WalletActionStatus `json:"status" api:"required"`
	// Any of "earn_incentive_claim".
	Type EarnIncentiveClaimActionResponseType `json:"type" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// The steps of the wallet action. Only returned if `?include=steps` is provided.
	Steps []WalletActionStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Chain         respjson.Field
		CreatedAt     respjson.Field
		Rewards       respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		WalletID      respjson.Field
		FailureReason respjson.Field
		Steps         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EarnIncentiveClaimActionResponse) RawJSON() string { return r.JSON.raw }
func (r *EarnIncentiveClaimActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EarnIncentiveClaimActionResponseType string

const (
	EarnIncentiveClaimActionResponseTypeEarnIncentiveClaim EarnIncentiveClaimActionResponseType = "earn_incentive_claim"
)

// Input for depositing assets into an ERC-4626 vault. Exactly one of `amount` or
// `raw_amount` must be provided.
//
// The property VaultID is required.
type EarnDepositRequestBody struct {
	// The ID of the vault to deposit into.
	VaultID string `json:"vault_id" api:"required"`
	// Human-readable decimal amount to deposit (e.g. "1.5" for 1.5 USDC). Exactly one
	// of `amount` or `raw_amount` must be provided.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Amount in smallest unit to deposit (e.g. "1500000" for 1.5 USDC with 6
	// decimals). Exactly one of `amount` or `raw_amount` must be provided.
	RawAmount param.Opt[string] `json:"raw_amount,omitzero"`
	paramObj
}

func (r EarnDepositRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow EarnDepositRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EarnDepositRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for withdrawing assets from an ERC-4626 vault. Exactly one of `amount` or
// `raw_amount` must be provided.
//
// The property VaultID is required.
type EarnWithdrawRequestBody struct {
	// The ID of the vault to withdraw from.
	VaultID string `json:"vault_id" api:"required"`
	// Human-readable decimal amount to withdraw (e.g. "1.5" for 1.5 USDC). Exactly one
	// of `amount` or `raw_amount` must be provided.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Amount in smallest unit to withdraw (e.g. "1500000" for 1.5 USDC with 6
	// decimals). Exactly one of `amount` or `raw_amount` must be provided.
	RawAmount param.Opt[string] `json:"raw_amount,omitzero"`
	paramObj
}

func (r EarnWithdrawRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow EarnWithdrawRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EarnWithdrawRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for claiming incentive rewards.
//
// The property Chain is required.
type EarnIncentiveClaimRequestBody struct {
	// The blockchain network on which to perform the incentive claim. Supported chains
	// include: 'ethereum', 'base', 'arbitrum', 'polygon', 'solana', and more, along
	// with their respective testnets.
	Chain string `json:"chain" api:"required"`
	paramObj
}

func (r EarnIncentiveClaimRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow EarnIncentiveClaimRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EarnIncentiveClaimRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
