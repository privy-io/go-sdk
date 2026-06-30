// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/apiquery"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to wallet actions
//
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

// Get the current status of a wallet action by its ID. Use `?include=steps` to
// include step-level details.
func (r *WalletActionService) Get(ctx context.Context, actionID string, params WalletActionGetParams, opts ...option.RequestOption) (res *WalletActionResponseUnion, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.WalletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	if actionID == "" {
		err = errors.New("missing required action_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/actions/%s", url.PathEscape(params.WalletID), actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// A wallet action step representing a transaction executed by a custodian (e.g.
// Bridge).
type CustodianTransactionWalletActionStep struct {
	// Identifier of the custodian executing this transaction (e.g. "bridge").
	Custodian string `json:"custodian" api:"required"`
	// Status of a custodian transaction step in a wallet action.
	//
	// Any of "preparing", "queued", "custodian_reviewing", "pending", "confirmed",
	// "rejected", "failed".
	Status CustodianTransactionWalletActionStepStatus `json:"status" api:"required"`
	// Any of "custodian_transaction".
	Type CustodianTransactionWalletActionStepType `json:"type" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Custodian     respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		FailureReason respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustodianTransactionWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *CustodianTransactionWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustodianTransactionWalletActionStepType string

const (
	CustodianTransactionWalletActionStepTypeCustodianTransaction CustodianTransactionWalletActionStepType = "custodian_transaction"
)

// Status of a custodian transaction step in a wallet action.
type CustodianTransactionWalletActionStepStatus string

const (
	CustodianTransactionWalletActionStepStatusPreparing          CustodianTransactionWalletActionStepStatus = "preparing"
	CustodianTransactionWalletActionStepStatusQueued             CustodianTransactionWalletActionStepStatus = "queued"
	CustodianTransactionWalletActionStepStatusCustodianReviewing CustodianTransactionWalletActionStepStatus = "custodian_reviewing"
	CustodianTransactionWalletActionStepStatusPending            CustodianTransactionWalletActionStepStatus = "pending"
	CustodianTransactionWalletActionStepStatusConfirmed          CustodianTransactionWalletActionStepStatus = "confirmed"
	CustodianTransactionWalletActionStepStatusRejected           CustodianTransactionWalletActionStepStatus = "rejected"
	CustodianTransactionWalletActionStepStatusFailed             CustodianTransactionWalletActionStepStatus = "failed"
)

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
	// Whether this step has reached on-chain finality. Absent until finality is
	// confirmed.
	Finalized bool `json:"finalized"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2           respjson.Field
		Status          respjson.Field
		TransactionHash respjson.Field
		Type            respjson.Field
		FailureReason   respjson.Field
		Finalized       respjson.Field
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

// The ERC-4337 entrypoint contract version used by the user operation.
type EvmUserOperationEntrypointVersion string

const (
	EvmUserOperationEntrypointVersion0_6 EvmUserOperationEntrypointVersion = "0.6"
	EvmUserOperationEntrypointVersion0_7 EvmUserOperationEntrypointVersion = "0.7"
	EvmUserOperationEntrypointVersion0_8 EvmUserOperationEntrypointVersion = "0.8"
	EvmUserOperationEntrypointVersion0_9 EvmUserOperationEntrypointVersion = "0.9"
)

// A wallet action step consisting of an EVM user operation.
type EvmUserOperationWalletActionStep struct {
	// Transaction hash of the bundle in which this user operation was included. Null
	// until included by a bundler.
	BundleTransactionHash string `json:"bundle_transaction_hash" api:"required"`
	// CAIP-2 network identifier, containing the chain ID of the user operation.
	Caip2 string `json:"caip2" api:"required"`
	// The ERC-4337 entrypoint contract version used by the user operation.
	//
	// Any of "0.6", "0.7", "0.8", "0.9".
	EntrypointVersion EvmUserOperationEntrypointVersion `json:"entrypoint_version" api:"required"`
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
	// Whether this step has reached on-chain finality. Absent until finality is
	// confirmed.
	Finalized bool `json:"finalized"`
	// Amount charged in USD for gas sponsorship on this step.
	GasCreditsChargedUsd string `json:"gas_credits_charged_usd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BundleTransactionHash respjson.Field
		Caip2                 respjson.Field
		EntrypointVersion     respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		UserOperationHash     respjson.Field
		FailureReason         respjson.Field
		Finalized             respjson.Field
		GasCreditsChargedUsd  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvmUserOperationWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *EvmUserOperationWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvmUserOperationWalletActionStepType string

const (
	EvmUserOperationWalletActionStepTypeEvmUserOperation EvmUserOperationWalletActionStepType = "evm_user_operation"
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

// A wallet action step representing a cross-chain/cross-asset fill by an external
// provider.
type ExternalTransactionWalletActionStep struct {
	// Status of an external transaction step in a wallet action.
	//
	// Any of "preparing", "queued", "pending", "confirmed", "rejected", "failed".
	Status ExternalTransactionWalletActionStepStatus `json:"status" api:"required"`
	// Any of "external_transaction".
	Type ExternalTransactionWalletActionStepType `json:"type" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status        respjson.Field
		Type          respjson.Field
		FailureReason respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalTransactionWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *ExternalTransactionWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalTransactionWalletActionStepType string

const (
	ExternalTransactionWalletActionStepTypeExternalTransaction ExternalTransactionWalletActionStepType = "external_transaction"
)

// Status of an external transaction step in a wallet action.
type ExternalTransactionWalletActionStepStatus string

const (
	ExternalTransactionWalletActionStepStatusPreparing ExternalTransactionWalletActionStepStatus = "preparing"
	ExternalTransactionWalletActionStepStatusQueued    ExternalTransactionWalletActionStepStatus = "queued"
	ExternalTransactionWalletActionStepStatusPending   ExternalTransactionWalletActionStepStatus = "pending"
	ExternalTransactionWalletActionStepStatusConfirmed ExternalTransactionWalletActionStepStatus = "confirmed"
	ExternalTransactionWalletActionStepStatusRejected  ExternalTransactionWalletActionStepStatus = "rejected"
	ExternalTransactionWalletActionStepStatusFailed    ExternalTransactionWalletActionStepStatus = "failed"
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

// A wallet action step consisting of an SVM (Solana) transaction.
type SvmTransactionWalletActionStep struct {
	// CAIP-2 chain identifier for the Solana network.
	Caip2 string `json:"caip2" api:"required"`
	// Status of an SVM step in a wallet action.
	//
	// Any of "preparing", "queued", "pending", "confirmed", "rejected", "reverted",
	// "failed".
	Status SvmWalletActionStepStatus `json:"status" api:"required"`
	// The Solana transaction signature (base58-encoded). Null until broadcast.
	TransactionSignature string `json:"transaction_signature" api:"required"`
	// Any of "svm_transaction".
	Type SvmTransactionWalletActionStepType `json:"type" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// Whether this step has reached on-chain finality. Absent until finality is
	// confirmed.
	Finalized bool `json:"finalized"`
	// Amount charged in USD for gas sponsorship on this step.
	GasCreditsChargedUsd string `json:"gas_credits_charged_usd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2                respjson.Field
		Status               respjson.Field
		TransactionSignature respjson.Field
		Type                 respjson.Field
		FailureReason        respjson.Field
		Finalized            respjson.Field
		GasCreditsChargedUsd respjson.Field
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

// Status of an SVM step in a wallet action.
type SvmWalletActionStepStatus string

const (
	SvmWalletActionStepStatusPreparing SvmWalletActionStepStatus = "preparing"
	SvmWalletActionStepStatusQueued    SvmWalletActionStepStatus = "queued"
	SvmWalletActionStepStatusPending   SvmWalletActionStepStatus = "pending"
	SvmWalletActionStepStatusConfirmed SvmWalletActionStepStatus = "confirmed"
	SvmWalletActionStepStatusRejected  SvmWalletActionStepStatus = "rejected"
	SvmWalletActionStepStatusReverted  SvmWalletActionStepStatus = "reverted"
	SvmWalletActionStepStatusFailed    SvmWalletActionStepStatus = "failed"
)

// Response for a swap action.
type SwapActionResponse struct {
	// The ID of the wallet action.
	ID string `json:"id" api:"required"`
	// CAIP-2 chain identifier for the swap.
	Caip2 string `json:"caip2" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Exact base-unit amount of input token. Populated after on-chain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address or "native" for the token being sold.
	InputToken string `json:"input_token" api:"required"`
	// Exact base-unit amount of output token. Populated after on-chain confirmation.
	OutputAmount string `json:"output_amount" api:"required"`
	// Token address or "native" for the token being bought.
	OutputToken string `json:"output_token" api:"required"`
	// Status of a wallet action.
	//
	// Any of "pending", "succeeded", "rejected", "failed".
	Status WalletActionStatus `json:"status" api:"required"`
	// Any of "swap".
	Type SwapActionResponseType `json:"type" api:"required"`
	// The ID of the wallet involved in the action.
	WalletID string `json:"wallet_id" api:"required"`
	// Recipient address on the destination chain. Present for cross-chain swaps. May
	// differ from the source wallet address when swapping between chain types (e.g.
	// EVM to Solana).
	DestinationAddress string `json:"destination_address"`
	// Destination chain CAIP-2 identifier. Present for cross-chain swaps.
	DestinationCaip2 string `json:"destination_caip2"`
	// Estimated fee breakdown from the provider quote. Only present for cross-chain
	// swaps. Populated after on-chain confirmation.
	EstimatedFees []FeeLineItemUnion `json:"estimated_fees" api:"nullable"`
	// Gas cost for a blockchain action. Includes both raw base-unit amount and a
	// human-readable decimal string, plus the gas token symbol.
	EstimatedGas Gas `json:"estimated_gas" api:"nullable"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// Actual fees paid for the swap. Populated after on-chain confirmation. Only
	// present for cross-chain swaps.
	Fees []FeeLineItemUnion `json:"fees" api:"nullable"`
	// Gas cost for a blockchain action. Includes both raw base-unit amount and a
	// human-readable decimal string, plus the gas token symbol.
	Gas Gas `json:"gas" api:"nullable"`
	// The steps of the wallet action. Only returned if `?include=steps` is provided.
	Steps []WalletActionStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Caip2              respjson.Field
		CreatedAt          respjson.Field
		InputAmount        respjson.Field
		InputToken         respjson.Field
		OutputAmount       respjson.Field
		OutputToken        respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		WalletID           respjson.Field
		DestinationAddress respjson.Field
		DestinationCaip2   respjson.Field
		EstimatedFees      respjson.Field
		EstimatedGas       respjson.Field
		FailureReason      respjson.Field
		Fees               respjson.Field
		Gas                respjson.Field
		Steps              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwapActionResponse) RawJSON() string { return r.JSON.raw }
func (r *SwapActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapActionResponseType string

const (
	SwapActionResponseTypeSwap SwapActionResponseType = "swap"
)

// A wallet action step consisting of a TVM (Tron) transaction.
type TvmTransactionWalletActionStep struct {
	// CAIP-2 chain identifier for the Tron network.
	Caip2 string `json:"caip2" api:"required"`
	// Status of a TVM (Tron) step in a wallet action.
	//
	// Any of "preparing", "queued", "pending", "confirmed", "rejected", "reverted",
	// "failed".
	Status TvmWalletActionStepStatus `json:"status" api:"required"`
	// The Tron transaction ID. Null until broadcast.
	TransactionID string `json:"transaction_id" api:"required"`
	// Any of "tvm_transaction".
	Type TvmTransactionWalletActionStepType `json:"type" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2         respjson.Field
		Status        respjson.Field
		TransactionID respjson.Field
		Type          respjson.Field
		FailureReason respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TvmTransactionWalletActionStep) RawJSON() string { return r.JSON.raw }
func (r *TvmTransactionWalletActionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TvmTransactionWalletActionStepType string

const (
	TvmTransactionWalletActionStepTypeTvmTransaction TvmTransactionWalletActionStepType = "tvm_transaction"
)

// Status of a TVM (Tron) step in a wallet action.
type TvmWalletActionStepStatus string

const (
	TvmWalletActionStepStatusPreparing TvmWalletActionStepStatus = "preparing"
	TvmWalletActionStepStatusQueued    TvmWalletActionStepStatus = "queued"
	TvmWalletActionStepStatusPending   TvmWalletActionStepStatus = "pending"
	TvmWalletActionStepStatusConfirmed TvmWalletActionStepStatus = "confirmed"
	TvmWalletActionStepStatusRejected  TvmWalletActionStepStatus = "rejected"
	TvmWalletActionStepStatusReverted  TvmWalletActionStepStatus = "reverted"
	TvmWalletActionStepStatusFailed    TvmWalletActionStepStatus = "failed"
)

// Response for a transfer action.
type TransferActionResponse struct {
	// The ID of the wallet action.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the wallet action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// Amount received on the destination chain. For exact_output cross-chain
	// transfers, set at creation (the guaranteed exact amount). For exact_input
	// cross-chain transfers, null until fill confirmation.
	DestinationAmount string `json:"destination_amount" api:"required"`
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
	// Whether the amount refers to the input token or output token.
	//
	// Any of "exact_input", "exact_output".
	AmountType AmountType `json:"amount_type"`
	// Destination asset for cross-asset transfers. Omitted for same-asset transfers.
	DestinationAsset string `json:"destination_asset"`
	// Destination chain for cross-chain transfers. Omitted for same-chain transfers.
	DestinationChain string `json:"destination_chain"`
	// Estimated fee breakdown from the provider quote. Only present for cross-chain or
	// cross-asset transfers. Populated after on-chain confirmation.
	EstimatedFees []FeeLineItemUnion `json:"estimated_fees" api:"nullable"`
	// Gas cost for a blockchain action. Includes both raw base-unit amount and a
	// human-readable decimal string, plus the gas token symbol.
	EstimatedGas Gas `json:"estimated_gas" api:"nullable"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason"`
	// Actual fees paid for the transfer. Populated after on-chain confirmation. Only
	// present for cross-chain transfers.
	Fees []FeeLineItemUnion `json:"fees" api:"nullable"`
	// Gas cost for a blockchain action. Includes both raw base-unit amount and a
	// human-readable decimal string, plus the gas token symbol.
	Gas Gas `json:"gas" api:"nullable"`
	// Decimal amount sent on the source chain (e.g. "1.5"). For exact_output
	// cross-chain transfers, null until fill confirmation.
	SourceAmount string `json:"source_amount"`
	// Asset identifier (e.g. "usdc", "eth"). Present when the transfer was initiated
	// with a named asset; omitted for custom-token transfers.
	SourceAsset string `json:"source_asset"`
	// Token contract address (EVM) or mint address (Solana). Present when the transfer
	// was initiated with `asset_address`.
	SourceAssetAddress string `json:"source_asset_address"`
	// Number of decimals for the transferred token. Present when the transfer was
	// initiated with `asset_address` and the decimals were resolved on-chain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// The steps of the wallet action. Only returned if `?include=steps` is provided.
	Steps []WalletActionStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		DestinationAddress  respjson.Field
		DestinationAmount   respjson.Field
		SourceChain         respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		WalletID            respjson.Field
		AmountType          respjson.Field
		DestinationAsset    respjson.Field
		DestinationChain    respjson.Field
		EstimatedFees       respjson.Field
		EstimatedGas        respjson.Field
		FailureReason       respjson.Field
		Fees                respjson.Field
		Gas                 respjson.Field
		SourceAmount        respjson.Field
		SourceAsset         respjson.Field
		SourceAssetAddress  respjson.Field
		SourceAssetDecimals respjson.Field
		Steps               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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

// Expandable relations to include on a wallet action response.
type WalletActionInclude string

const (
	WalletActionIncludeSteps WalletActionInclude = "steps"
)

// WalletActionResponseUnion contains all possible properties and values from
// [SwapActionResponse], [TransferActionResponse], [EarnDepositActionResponse],
// [EarnWithdrawActionResponse], [EarnIncentiveClaimActionResponse].
//
// Use the [WalletActionResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletActionResponseUnion struct {
	ID        string    `json:"id"`
	Caip2     string    `json:"caip2"`
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [SwapActionResponse].
	InputAmount string `json:"input_amount"`
	// This field is from variant [SwapActionResponse].
	InputToken string `json:"input_token"`
	// This field is from variant [SwapActionResponse].
	OutputAmount string `json:"output_amount"`
	// This field is from variant [SwapActionResponse].
	OutputToken string `json:"output_token"`
	// This field is from variant [SwapActionResponse].
	Status WalletActionStatus `json:"status"`
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	Type               string `json:"type"`
	WalletID           string `json:"wallet_id"`
	DestinationAddress string `json:"destination_address"`
	// This field is from variant [SwapActionResponse].
	DestinationCaip2 string             `json:"destination_caip2"`
	EstimatedFees    []FeeLineItemUnion `json:"estimated_fees"`
	// This field is from variant [SwapActionResponse].
	EstimatedGas Gas `json:"estimated_gas"`
	// This field is from variant [SwapActionResponse].
	FailureReason FailureReason      `json:"failure_reason"`
	Fees          []FeeLineItemUnion `json:"fees"`
	// This field is from variant [SwapActionResponse].
	Gas   Gas                     `json:"gas"`
	Steps []WalletActionStepUnion `json:"steps"`
	// This field is from variant [TransferActionResponse].
	DestinationAmount string `json:"destination_amount"`
	// This field is from variant [TransferActionResponse].
	SourceChain string `json:"source_chain"`
	// This field is from variant [TransferActionResponse].
	AmountType AmountType `json:"amount_type"`
	// This field is from variant [TransferActionResponse].
	DestinationAsset string `json:"destination_asset"`
	// This field is from variant [TransferActionResponse].
	DestinationChain string `json:"destination_chain"`
	// This field is from variant [TransferActionResponse].
	SourceAmount string `json:"source_amount"`
	// This field is from variant [TransferActionResponse].
	SourceAsset string `json:"source_asset"`
	// This field is from variant [TransferActionResponse].
	SourceAssetAddress string `json:"source_asset_address"`
	// This field is from variant [TransferActionResponse].
	SourceAssetDecimals int64  `json:"source_asset_decimals"`
	AssetAddress        string `json:"asset_address"`
	RawAmount           string `json:"raw_amount"`
	ShareAmount         string `json:"share_amount"`
	VaultAddress        string `json:"vault_address"`
	VaultID             string `json:"vault_id"`
	Amount              string `json:"amount"`
	Asset               string `json:"asset"`
	Decimals            int64  `json:"decimals"`
	// This field is from variant [EarnIncentiveClaimActionResponse].
	Chain string `json:"chain"`
	// This field is from variant [EarnIncentiveClaimActionResponse].
	Rewards []EarnIncetiveClaimRewardEntry `json:"rewards"`
	JSON    struct {
		ID                  respjson.Field
		Caip2               respjson.Field
		CreatedAt           respjson.Field
		InputAmount         respjson.Field
		InputToken          respjson.Field
		OutputAmount        respjson.Field
		OutputToken         respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		WalletID            respjson.Field
		DestinationAddress  respjson.Field
		DestinationCaip2    respjson.Field
		EstimatedFees       respjson.Field
		EstimatedGas        respjson.Field
		FailureReason       respjson.Field
		Fees                respjson.Field
		Gas                 respjson.Field
		Steps               respjson.Field
		DestinationAmount   respjson.Field
		SourceChain         respjson.Field
		AmountType          respjson.Field
		DestinationAsset    respjson.Field
		DestinationChain    respjson.Field
		SourceAmount        respjson.Field
		SourceAsset         respjson.Field
		SourceAssetAddress  respjson.Field
		SourceAssetDecimals respjson.Field
		AssetAddress        respjson.Field
		RawAmount           respjson.Field
		ShareAmount         respjson.Field
		VaultAddress        respjson.Field
		VaultID             respjson.Field
		Amount              respjson.Field
		Asset               respjson.Field
		Decimals            respjson.Field
		Chain               respjson.Field
		Rewards             respjson.Field
		raw                 string
	} `json:"-"`
}

// anyWalletActionResponse is implemented by each variant of
// [WalletActionResponseUnion] to add type safety for the return type of
// [WalletActionResponseUnion.AsAny]
type anyWalletActionResponse interface {
	implWalletActionResponseUnion()
}

func (SwapActionResponse) implWalletActionResponseUnion()               {}
func (TransferActionResponse) implWalletActionResponseUnion()           {}
func (EarnDepositActionResponse) implWalletActionResponseUnion()        {}
func (EarnWithdrawActionResponse) implWalletActionResponseUnion()       {}
func (EarnIncentiveClaimActionResponse) implWalletActionResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletActionResponseUnion.AsAny().(type) {
//	case privyclient.SwapActionResponse:
//	case privyclient.TransferActionResponse:
//	case privyclient.EarnDepositActionResponse:
//	case privyclient.EarnWithdrawActionResponse:
//	case privyclient.EarnIncentiveClaimActionResponse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletActionResponseUnion) AsAny() anyWalletActionResponse {
	switch u.Type {
	case "swap":
		return u.AsSwap()
	case "transfer":
		return u.AsTransfer()
	case "earn_deposit":
		return u.AsEarnDeposit()
	case "earn_withdraw":
		return u.AsEarnWithdraw()
	case "earn_incentive_claim":
		return u.AsEarnIncentiveClaim()
	}
	return nil
}

func (u WalletActionResponseUnion) AsSwap() (v SwapActionResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionResponseUnion) AsTransfer() (v TransferActionResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionResponseUnion) AsEarnDeposit() (v EarnDepositActionResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionResponseUnion) AsEarnWithdraw() (v EarnWithdrawActionResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionResponseUnion) AsEarnIncentiveClaim() (v EarnIncentiveClaimActionResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletActionResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletActionResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a wallet action.
type WalletActionStatus string

const (
	WalletActionStatusPending   WalletActionStatus = "pending"
	WalletActionStatusSucceeded WalletActionStatus = "succeeded"
	WalletActionStatusRejected  WalletActionStatus = "rejected"
	WalletActionStatusFailed    WalletActionStatus = "failed"
)

// WalletActionStepUnion contains all possible properties and values from
// [EvmTransactionWalletActionStep], [EvmUserOperationWalletActionStep],
// [SvmTransactionWalletActionStep], [TvmTransactionWalletActionStep],
// [ExternalTransactionWalletActionStep], [CustodianTransactionWalletActionStep].
//
// Use the [WalletActionStepUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletActionStepUnion struct {
	Caip2  string `json:"caip2"`
	Status string `json:"status"`
	// This field is from variant [EvmTransactionWalletActionStep].
	TransactionHash string `json:"transaction_hash"`
	// Any of "evm_transaction", "evm_user_operation", "svm_transaction",
	// "tvm_transaction", "external_transaction", "custodian_transaction".
	Type string `json:"type"`
	// This field is from variant [EvmTransactionWalletActionStep].
	FailureReason FailureReason `json:"failure_reason"`
	Finalized     bool          `json:"finalized"`
	// This field is from variant [EvmUserOperationWalletActionStep].
	BundleTransactionHash string `json:"bundle_transaction_hash"`
	// This field is from variant [EvmUserOperationWalletActionStep].
	EntrypointVersion EvmUserOperationEntrypointVersion `json:"entrypoint_version"`
	// This field is from variant [EvmUserOperationWalletActionStep].
	UserOperationHash    string `json:"user_operation_hash"`
	GasCreditsChargedUsd string `json:"gas_credits_charged_usd"`
	// This field is from variant [SvmTransactionWalletActionStep].
	TransactionSignature string `json:"transaction_signature"`
	// This field is from variant [TvmTransactionWalletActionStep].
	TransactionID string `json:"transaction_id"`
	// This field is from variant [CustodianTransactionWalletActionStep].
	Custodian string `json:"custodian"`
	JSON      struct {
		Caip2                 respjson.Field
		Status                respjson.Field
		TransactionHash       respjson.Field
		Type                  respjson.Field
		FailureReason         respjson.Field
		Finalized             respjson.Field
		BundleTransactionHash respjson.Field
		EntrypointVersion     respjson.Field
		UserOperationHash     respjson.Field
		GasCreditsChargedUsd  respjson.Field
		TransactionSignature  respjson.Field
		TransactionID         respjson.Field
		Custodian             respjson.Field
		raw                   string
	} `json:"-"`
}

// anyWalletActionStep is implemented by each variant of [WalletActionStepUnion] to
// add type safety for the return type of [WalletActionStepUnion.AsAny]
type anyWalletActionStep interface {
	implWalletActionStepUnion()
}

func (EvmTransactionWalletActionStep) implWalletActionStepUnion()       {}
func (EvmUserOperationWalletActionStep) implWalletActionStepUnion()     {}
func (SvmTransactionWalletActionStep) implWalletActionStepUnion()       {}
func (TvmTransactionWalletActionStep) implWalletActionStepUnion()       {}
func (ExternalTransactionWalletActionStep) implWalletActionStepUnion()  {}
func (CustodianTransactionWalletActionStep) implWalletActionStepUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletActionStepUnion.AsAny().(type) {
//	case privyclient.EvmTransactionWalletActionStep:
//	case privyclient.EvmUserOperationWalletActionStep:
//	case privyclient.SvmTransactionWalletActionStep:
//	case privyclient.TvmTransactionWalletActionStep:
//	case privyclient.ExternalTransactionWalletActionStep:
//	case privyclient.CustodianTransactionWalletActionStep:
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
	case "tvm_transaction":
		return u.AsTvmTransaction()
	case "external_transaction":
		return u.AsExternalTransaction()
	case "custodian_transaction":
		return u.AsCustodianTransaction()
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

func (u WalletActionStepUnion) AsTvmTransaction() (v TvmTransactionWalletActionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionStepUnion) AsExternalTransaction() (v ExternalTransactionWalletActionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletActionStepUnion) AsCustodianTransaction() (v CustodianTransactionWalletActionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletActionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletActionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of wallet action
type WalletActionType string

const (
	WalletActionTypeSwap               WalletActionType = "swap"
	WalletActionTypeTransfer           WalletActionType = "transfer"
	WalletActionTypeEarnDeposit        WalletActionType = "earn_deposit"
	WalletActionTypeEarnWithdraw       WalletActionType = "earn_withdraw"
	WalletActionTypeEarnIncentiveClaim WalletActionType = "earn_incentive_claim"
)

type WalletActionGetParams struct {
	// ID of the wallet.
	WalletID string `path:"wallet_id" api:"required" json:"-"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Expandable relations to include on a wallet action response.
	//
	// Any of "steps".
	Include WalletActionInclude `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletActionGetParams]'s query parameters as `url.Values`.
func (r WalletActionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
