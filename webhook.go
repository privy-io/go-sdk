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

// Payload for the mfa.enabled webhook event.
type MfaEnabledWebhookPayload struct {
	// The MFA method that was enabled.
	//
	// Any of "sms", "totp", "passkey".
	Method MfaEnabledWebhookPayloadMethod `json:"method" api:"required"`
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

// The MFA method that was enabled.
type MfaEnabledWebhookPayloadMethod string

const (
	MfaEnabledWebhookPayloadMethodSMS     MfaEnabledWebhookPayloadMethod = "sms"
	MfaEnabledWebhookPayloadMethodTotp    MfaEnabledWebhookPayloadMethod = "totp"
	MfaEnabledWebhookPayloadMethodPasskey MfaEnabledWebhookPayloadMethod = "passkey"
)

// The type of webhook event.
type MfaEnabledWebhookPayloadType string

const (
	MfaEnabledWebhookPayloadTypeMfaEnabled MfaEnabledWebhookPayloadType = "mfa.enabled"
)

// Payload for the mfa.disabled webhook event.
type MfaDisabledWebhookPayload struct {
	// The MFA method that was disabled.
	//
	// Any of "sms", "totp", "passkey".
	Method MfaDisabledWebhookPayloadMethod `json:"method" api:"required"`
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

// The MFA method that was disabled.
type MfaDisabledWebhookPayloadMethod string

const (
	MfaDisabledWebhookPayloadMethodSMS     MfaDisabledWebhookPayloadMethod = "sms"
	MfaDisabledWebhookPayloadMethodTotp    MfaDisabledWebhookPayloadMethod = "totp"
	MfaDisabledWebhookPayloadMethodPasskey MfaDisabledWebhookPayloadMethod = "passkey"
)

// The type of webhook event.
type MfaDisabledWebhookPayloadType string

const (
	MfaDisabledWebhookPayloadTypeMfaDisabled MfaDisabledWebhookPayloadType = "mfa.disabled"
)

// Payload for the transaction.broadcasted webhook event.
type TransactionBroadcastedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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

// Payload for the transaction.still_pending webhook event.
type TransactionStillPendingWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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

// Payload for the transaction.failed webhook event.
type TransactionFailedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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

// Payload for the transaction.replaced webhook event.
type TransactionReplacedWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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

// Payload for the transaction.provider_error webhook event.
type TransactionProviderErrorWebhookPayload struct {
	// The CAIP-2 chain identifier (e.g., eip155:1 for Ethereum mainnet).
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

// Payload for the user.transferred_account webhook event.
type UserTransferredAccountWebhookPayload struct {
	// A linked account for the user.
	Account LinkedAccountUnion `json:"account" api:"required"`
	// Any of true.
	DeletedUser bool                                         `json:"deletedUser" api:"required"`
	FromUser    UserTransferredAccountWebhookPayloadFromUser `json:"fromUser" api:"required"`
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

type UserTransferredAccountWebhookPayloadFromUser struct {
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTransferredAccountWebhookPayloadFromUser) RawJSON() string { return r.JSON.raw }
func (r *UserTransferredAccountWebhookPayloadFromUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type UserTransferredAccountWebhookPayloadType string

const (
	UserTransferredAccountWebhookPayloadTypeUserTransferredAccount UserTransferredAccountWebhookPayloadType = "user.transferred_account"
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

// Payload for the wallet_action.swap.created webhook event.
type WalletActionSwapCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// Amount of input token in base units. Populated after on-chain confirmation.
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

// Payload for the wallet_action.swap.succeeded webhook event.
type WalletActionSwapSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// Amount of input token in base units. Populated after on-chain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Amount of output token received, in base units. Populated after on-chain
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

// Payload for the wallet_action.swap.rejected webhook event.
type WalletActionSwapRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Amount of input token in base units. Populated after on-chain confirmation.
	InputAmount string `json:"input_amount" api:"required"`
	// Token address being sold.
	InputToken string `json:"input_token" api:"required"`
	// Token address being bought.
	OutputToken string `json:"output_token" api:"required"`
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

// Payload for the wallet_action.swap.failed webhook event.
type WalletActionSwapFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Amount of input token in base units. Populated after on-chain confirmation.
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

// Payload for the wallet_action.transfer.created webhook event.
type WalletActionTransferCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// Chain name (e.g. "base", "ethereum").
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
	// initiated with `asset_address` and the decimals were resolved on-chain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
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

// Payload for the wallet_action.transfer.succeeded webhook event.
type WalletActionTransferSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// Chain name (e.g. "base", "ethereum").
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
	// initiated with `asset_address` and the decimals were resolved on-chain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
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

// Payload for the wallet_action.transfer.rejected webhook event.
type WalletActionTransferRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Chain name (e.g. "base", "ethereum").
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
	// initiated with `asset_address` and the decimals were resolved on-chain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
		DestinationAddress  respjson.Field
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

// Payload for the wallet_action.transfer.failed webhook event.
type WalletActionTransferFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Recipient address.
	DestinationAddress string `json:"destination_address" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Chain name (e.g. "base", "ethereum").
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
	// initiated with `asset_address` and the decimals were resolved on-chain.
	SourceAssetDecimals int64 `json:"source_asset_decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType          respjson.Field
		DestinationAddress  respjson.Field
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

// Payload for the wallet_action.earn_deposit.created webhook event.
type WalletActionEarnDepositCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
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

// Payload for the wallet_action.earn_deposit.succeeded webhook event.
type WalletActionEarnDepositSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
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

// Payload for the wallet_action.earn_deposit.rejected webhook event.
type WalletActionEarnDepositRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of asset deposited (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
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

// Payload for the wallet_action.earn_deposit.failed webhook event.
type WalletActionEarnDepositFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
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

// Payload for the wallet_action.earn_withdraw.created webhook event.
type WalletActionEarnWithdrawCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
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

// Payload for the wallet_action.earn_withdraw.succeeded webhook event.
type WalletActionEarnWithdrawSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
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

// Payload for the wallet_action.earn_withdraw.rejected webhook event.
type WalletActionEarnWithdrawRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
	// Base-unit amount of asset withdrawn (e.g. "1500000").
	RawAmount string `json:"raw_amount" api:"required"`
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

// Payload for the wallet_action.earn_withdraw.failed webhook event.
type WalletActionEarnWithdrawFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// Underlying asset token address.
	AssetAddress string `json:"asset_address" api:"required"`
	// CAIP-2 chain identifier.
	Caip2 string `json:"caip2" api:"required"`
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

// Payload for the wallet_action.earn_incentive_claim.created webhook event.
type WalletActionEarnIncentiveClaimCreatedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "base", "ethereum").
	Chain string `json:"chain" api:"required"`
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

// Payload for the wallet_action.earn_incentive_claim.succeeded webhook event.
type WalletActionEarnIncentiveClaimSucceededWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "base", "ethereum").
	Chain string `json:"chain" api:"required"`
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

// Payload for the wallet_action.earn_incentive_claim.rejected webhook event.
type WalletActionEarnIncentiveClaimRejectedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "base", "ethereum").
	Chain string `json:"chain" api:"required"`
	// A description of why a wallet action (or a step within a wallet action) failed.
	FailureReason FailureReason `json:"failure_reason" api:"required"`
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

// Payload for the wallet_action.earn_incentive_claim.failed webhook event.
type WalletActionEarnIncentiveClaimFailedWebhookPayload struct {
	// Type of wallet action
	//
	// Any of "swap", "transfer", "earn_deposit", "earn_withdraw",
	// "earn_incentive_claim".
	ActionType WalletActionType `json:"action_type" api:"required"`
	// EVM chain name (e.g. "base", "ethereum").
	Chain string `json:"chain" api:"required"`
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

// WalletFundsAssetUnion contains all possible properties and values from
// [WalletFundsNativeTokenAsset], [WalletFundsErc20Asset], [WalletFundsSplAsset],
// [WalletFundsSacAsset].
//
// Use the [WalletFundsAssetUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletFundsAssetUnion struct {
	// This field is a union of [any], [string], [string]
	Address WalletFundsAssetUnionAddress `json:"address"`
	// Any of "native-token", "erc20", "spl", "sac".
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

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletFundsAssetUnion.AsAny().(type) {
//	case privyclient.WalletFundsNativeTokenAsset:
//	case privyclient.WalletFundsErc20Asset:
//	case privyclient.WalletFundsSplAsset:
//	case privyclient.WalletFundsSacAsset:
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

// Payload for the wallet.funds_deposited webhook event.
type FundsDepositedWebhookPayload struct {
	// The amount transferred, as a stringified bigint.
	Amount string `json:"amount" api:"required"`
	// An asset involved in a wallet transfer.
	Asset WalletFundsAssetUnion             `json:"asset" api:"required"`
	Block FundsDepositedWebhookPayloadBlock `json:"block" api:"required"`
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

type FundsDepositedWebhookPayloadBlock struct {
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
func (r FundsDepositedWebhookPayloadBlock) RawJSON() string { return r.JSON.raw }
func (r *FundsDepositedWebhookPayloadBlock) UnmarshalJSON(data []byte) error {
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
	Asset WalletFundsAssetUnion             `json:"asset" api:"required"`
	Block FundsWithdrawnWebhookPayloadBlock `json:"block" api:"required"`
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

type FundsWithdrawnWebhookPayloadBlock struct {
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
func (r FundsWithdrawnWebhookPayloadBlock) RawJSON() string { return r.JSON.raw }
func (r *FundsWithdrawnWebhookPayloadBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type FundsWithdrawnWebhookPayloadType string

const (
	FundsWithdrawnWebhookPayloadTypeWalletFundsWithdrawn FundsWithdrawnWebhookPayloadType = "wallet.funds_withdrawn"
)

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
	// Any of "display", "client".
	ExportSource PrivateKeyExportWebhookPayloadExportSource `json:"export_source"`
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

type PrivateKeyExportWebhookPayloadExportSource string

const (
	PrivateKeyExportWebhookPayloadExportSourceDisplay PrivateKeyExportWebhookPayloadExportSource = "display"
	PrivateKeyExportWebhookPayloadExportSourceClient  PrivateKeyExportWebhookPayloadExportSource = "client"
)

// Payload for the wallet.recovery_setup webhook event.
type WalletRecoverySetupWebhookPayload struct {
	// The recovery method that was set up.
	//
	// Any of "user_passcode_derived_recovery_key",
	// "privy_passcode_derived_recovery_key", "privy_generated_recovery_key",
	// "google_drive_recovery_secret", "icloud_recovery_secret",
	// "recovery_encryption_key".
	Method WalletRecoverySetupWebhookPayloadMethod `json:"method" api:"required"`
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

// The recovery method that was set up.
type WalletRecoverySetupWebhookPayloadMethod string

const (
	WalletRecoverySetupWebhookPayloadMethodUserPasscodeDerivedRecoveryKey  WalletRecoverySetupWebhookPayloadMethod = "user_passcode_derived_recovery_key"
	WalletRecoverySetupWebhookPayloadMethodPrivyPasscodeDerivedRecoveryKey WalletRecoverySetupWebhookPayloadMethod = "privy_passcode_derived_recovery_key"
	WalletRecoverySetupWebhookPayloadMethodPrivyGeneratedRecoveryKey       WalletRecoverySetupWebhookPayloadMethod = "privy_generated_recovery_key"
	WalletRecoverySetupWebhookPayloadMethodGoogleDriveRecoverySecret       WalletRecoverySetupWebhookPayloadMethod = "google_drive_recovery_secret"
	WalletRecoverySetupWebhookPayloadMethodICloudRecoverySecret            WalletRecoverySetupWebhookPayloadMethod = "icloud_recovery_secret"
	WalletRecoverySetupWebhookPayloadMethodRecoveryEncryptionKey           WalletRecoverySetupWebhookPayloadMethod = "recovery_encryption_key"
)

// The type of webhook event.
type WalletRecoverySetupWebhookPayloadType string

const (
	WalletRecoverySetupWebhookPayloadTypeWalletRecoverySetup WalletRecoverySetupWebhookPayloadType = "wallet.recovery_setup"
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

// Payload for the yield.claim.confirmed webhook event.
type YieldClaimConfirmedWebhookPayload struct {
	Caip2         string                                    `json:"caip2" api:"required"`
	Rewards       []YieldClaimConfirmedWebhookPayloadReward `json:"rewards" api:"required"`
	TransactionID string                                    `json:"transaction_id" api:"required"`
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

type YieldClaimConfirmedWebhookPayloadReward struct {
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
func (r YieldClaimConfirmedWebhookPayloadReward) RawJSON() string { return r.JSON.raw }
func (r *YieldClaimConfirmedWebhookPayloadReward) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of webhook event.
type YieldClaimConfirmedWebhookPayloadType string

const (
	YieldClaimConfirmedWebhookPayloadTypeYieldClaimConfirmed YieldClaimConfirmedWebhookPayloadType = "yield.claim.confirmed"
)

// UnsafeUnwrapWebhookEventUnion contains all possible properties and values from
// [UserCreatedWebhookPayload], [UserAuthenticatedWebhookPayload],
// [UserLinkedAccountWebhookPayload], [UserUnlinkedAccountWebhookPayload],
// [UserUpdatedAccountWebhookPayload], [UserTransferredAccountWebhookPayload],
// [UserWalletCreatedWebhookPayload], [TransactionBroadcastedWebhookPayload],
// [TransactionConfirmedWebhookPayload],
// [TransactionExecutionRevertedWebhookPayload],
// [TransactionStillPendingWebhookPayload], [TransactionFailedWebhookPayload],
// [TransactionReplacedWebhookPayload], [TransactionProviderErrorWebhookPayload],
// [UserOperationCompletedWebhookPayload], [FundsDepositedWebhookPayload],
// [FundsWithdrawnWebhookPayload], [PrivateKeyExportWebhookPayload],
// [WalletRecoverySetupWebhookPayload], [WalletRecoveredWebhookPayload],
// [WalletActionSwapCreatedWebhookPayload],
// [WalletActionSwapSucceededWebhookPayload],
// [WalletActionSwapRejectedWebhookPayload],
// [WalletActionSwapFailedWebhookPayload],
// [WalletActionTransferCreatedWebhookPayload],
// [WalletActionTransferSucceededWebhookPayload],
// [WalletActionTransferRejectedWebhookPayload],
// [WalletActionTransferFailedWebhookPayload],
// [WalletActionEarnDepositCreatedWebhookPayload],
// [WalletActionEarnDepositSucceededWebhookPayload],
// [WalletActionEarnDepositRejectedWebhookPayload],
// [WalletActionEarnDepositFailedWebhookPayload],
// [WalletActionEarnWithdrawCreatedWebhookPayload],
// [WalletActionEarnWithdrawSucceededWebhookPayload],
// [WalletActionEarnWithdrawRejectedWebhookPayload],
// [WalletActionEarnWithdrawFailedWebhookPayload],
// [WalletActionEarnIncentiveClaimCreatedWebhookPayload],
// [WalletActionEarnIncentiveClaimSucceededWebhookPayload],
// [WalletActionEarnIncentiveClaimRejectedWebhookPayload],
// [WalletActionEarnIncentiveClaimFailedWebhookPayload],
// [MfaEnabledWebhookPayload], [MfaDisabledWebhookPayload],
// [IntentCreatedWebhookPayload], [IntentAuthorizedWebhookPayload],
// [IntentExecutedWebhookPayload], [IntentFailedWebhookPayload],
// [YieldDepositConfirmedWebhookPayload], [YieldWithdrawConfirmedWebhookPayload],
// [YieldClaimConfirmedWebhookPayload].
//
// Use the [UnsafeUnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	// Any of "user.created", "user.authenticated", "user.linked_account",
	// "user.unlinked_account", "user.updated_account", "user.transferred_account",
	// "user.wallet_created", "transaction.broadcasted", "transaction.confirmed",
	// "transaction.execution_reverted", "transaction.still_pending",
	// "transaction.failed", "transaction.replaced", "transaction.provider_error",
	// "user_operation.completed", "wallet.funds_deposited", "wallet.funds_withdrawn",
	// "wallet.private_key_export", "wallet.recovery_setup", "wallet.recovered",
	// "wallet_action.swap.created", "wallet_action.swap.succeeded",
	// "wallet_action.swap.rejected", "wallet_action.swap.failed",
	// "wallet_action.transfer.created", "wallet_action.transfer.succeeded",
	// "wallet_action.transfer.rejected", "wallet_action.transfer.failed",
	// "wallet_action.earn_deposit.created", "wallet_action.earn_deposit.succeeded",
	// "wallet_action.earn_deposit.rejected", "wallet_action.earn_deposit.failed",
	// "wallet_action.earn_withdraw.created", "wallet_action.earn_withdraw.succeeded",
	// "wallet_action.earn_withdraw.rejected", "wallet_action.earn_withdraw.failed",
	// "wallet_action.earn_incentive_claim.created",
	// "wallet_action.earn_incentive_claim.succeeded",
	// "wallet_action.earn_incentive_claim.rejected",
	// "wallet_action.earn_incentive_claim.failed", "mfa.enabled", "mfa.disabled",
	// "intent.created", "intent.authorized", "intent.executed", "intent.failed",
	// "yield.deposit.confirmed", "yield.withdraw.confirmed", "yield.claim.confirmed".
	Type string `json:"type"`
	// This field is from variant [UserCreatedWebhookPayload].
	User User `json:"user"`
	// This field is from variant [UserAuthenticatedWebhookPayload].
	Account LinkedAccountUnion `json:"account"`
	// This field is from variant [UserTransferredAccountWebhookPayload].
	DeletedUser bool `json:"deletedUser"`
	// This field is from variant [UserTransferredAccountWebhookPayload].
	FromUser UserTransferredAccountWebhookPayloadFromUser `json:"fromUser"`
	// This field is from variant [UserTransferredAccountWebhookPayload].
	ToUser User `json:"toUser"`
	// This field is from variant [UserWalletCreatedWebhookPayload].
	Wallet          LinkedAccountBaseWallet `json:"wallet"`
	Caip2           string                  `json:"caip2"`
	TransactionHash string                  `json:"transaction_hash"`
	TransactionID   string                  `json:"transaction_id"`
	WalletID        string                  `json:"wallet_id"`
	ReferenceID     string                  `json:"reference_id"`
	// This field is from variant [TransactionStillPendingWebhookPayload].
	TransactionRequest UnsignedStandardEthereumTransactionResp `json:"transaction_request"`
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
	Amount     string `json:"amount"`
	// This field is a union of [WalletFundsAssetUnion], [string], [string], [string],
	// [string], [string], [string], [string], [string]
	Asset UnsafeUnwrapWebhookEventUnionAsset `json:"asset"`
	// This field is a union of [FundsDepositedWebhookPayloadBlock],
	// [FundsWithdrawnWebhookPayloadBlock]
	Block          UnsafeUnwrapWebhookEventUnionBlock `json:"block"`
	IdempotencyKey string                             `json:"idempotency_key"`
	Recipient      string                             `json:"recipient"`
	// This field is from variant [FundsDepositedWebhookPayload].
	BridgeMetadata BridgeMetadataUnion `json:"bridge_metadata"`
	TransactionFee string              `json:"transaction_fee"`
	UserID         string              `json:"user_id"`
	WalletAddress  string              `json:"wallet_address"`
	// This field is from variant [PrivateKeyExportWebhookPayload].
	ExportSource PrivateKeyExportWebhookPayloadExportSource `json:"export_source"`
	Method       string                                     `json:"method"`
	// This field is from variant [WalletActionSwapCreatedWebhookPayload].
	ActionType     WalletActionType `json:"action_type"`
	InputAmount    string           `json:"input_amount"`
	InputToken     string           `json:"input_token"`
	OutputToken    string           `json:"output_token"`
	Status         string           `json:"status"`
	WalletActionID string           `json:"wallet_action_id"`
	// This field is from variant [WalletActionSwapSucceededWebhookPayload].
	OutputAmount string                  `json:"output_amount"`
	Steps        []WalletActionStepUnion `json:"steps"`
	// This field is from variant [WalletActionSwapRejectedWebhookPayload].
	FailureReason       FailureReason `json:"failure_reason"`
	DestinationAddress  string        `json:"destination_address"`
	SourceChain         string        `json:"source_chain"`
	SourceAmount        string        `json:"source_amount"`
	SourceAsset         string        `json:"source_asset"`
	SourceAssetAddress  string        `json:"source_asset_address"`
	SourceAssetDecimals int64         `json:"source_asset_decimals"`
	AssetAddress        string        `json:"asset_address"`
	RawAmount           string        `json:"raw_amount"`
	VaultAddress        string        `json:"vault_address"`
	VaultID             string        `json:"vault_id"`
	Decimals            int64         `json:"decimals"`
	ShareAmount         string        `json:"share_amount"`
	Chain               string        `json:"chain"`
	// This field is a union of [[]EarnIncetiveClaimRewardEntry],
	// [[]EarnIncetiveClaimRewardEntry], [[]EarnIncetiveClaimRewardEntry],
	// [[]EarnIncetiveClaimRewardEntry], [[]YieldClaimConfirmedWebhookPayloadReward]
	Rewards   UnsafeUnwrapWebhookEventUnionRewards `json:"rewards"`
	CreatedAt float64                              `json:"created_at"`
	ExpiresAt float64                              `json:"expires_at"`
	IntentID  string                               `json:"intent_id"`
	// This field is from variant [IntentCreatedWebhookPayload].
	IntentType IntentType `json:"intent_type"`
	// This field is from variant [IntentCreatedWebhookPayload].
	AuthorizationDetails []IntentAuthorization `json:"authorization_details"`
	CreatedByDisplayName string                `json:"created_by_display_name"`
	CreatedByID          string                `json:"created_by_id"`
	// This field is from variant [IntentAuthorizedWebhookPayload].
	AuthorizedAt float64 `json:"authorized_at"`
	// This field is from variant [IntentAuthorizedWebhookPayload].
	Member IntentAuthorizationKeyQuorumMemberUnion `json:"member"`
	// This field is from variant [IntentExecutedWebhookPayload].
	ActionResult BaseActionResult `json:"action_result"`
	Assets       string           `json:"assets"`
	Owner        string           `json:"owner"`
	Shares       string           `json:"shares"`
	// This field is from variant [YieldWithdrawConfirmedWebhookPayload].
	Receiver string `json:"receiver"`
	JSON     struct {
		Type                 respjson.Field
		User                 respjson.Field
		Account              respjson.Field
		DeletedUser          respjson.Field
		FromUser             respjson.Field
		ToUser               respjson.Field
		Wallet               respjson.Field
		Caip2                respjson.Field
		TransactionHash      respjson.Field
		TransactionID        respjson.Field
		WalletID             respjson.Field
		ReferenceID          respjson.Field
		TransactionRequest   respjson.Field
		ActualGasCost        respjson.Field
		ActualGasUsed        respjson.Field
		BlockNumber          respjson.Field
		LogIndex             respjson.Field
		Nonce                respjson.Field
		Paymaster            respjson.Field
		Sender               respjson.Field
		Success              respjson.Field
		UserOpHash           respjson.Field
		Amount               respjson.Field
		Asset                respjson.Field
		Block                respjson.Field
		IdempotencyKey       respjson.Field
		Recipient            respjson.Field
		BridgeMetadata       respjson.Field
		TransactionFee       respjson.Field
		UserID               respjson.Field
		WalletAddress        respjson.Field
		ExportSource         respjson.Field
		Method               respjson.Field
		ActionType           respjson.Field
		InputAmount          respjson.Field
		InputToken           respjson.Field
		OutputToken          respjson.Field
		Status               respjson.Field
		WalletActionID       respjson.Field
		OutputAmount         respjson.Field
		Steps                respjson.Field
		FailureReason        respjson.Field
		DestinationAddress   respjson.Field
		SourceChain          respjson.Field
		SourceAmount         respjson.Field
		SourceAsset          respjson.Field
		SourceAssetAddress   respjson.Field
		SourceAssetDecimals  respjson.Field
		AssetAddress         respjson.Field
		RawAmount            respjson.Field
		VaultAddress         respjson.Field
		VaultID              respjson.Field
		Decimals             respjson.Field
		ShareAmount          respjson.Field
		Chain                respjson.Field
		Rewards              respjson.Field
		CreatedAt            respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		AuthorizationDetails respjson.Field
		CreatedByDisplayName respjson.Field
		CreatedByID          respjson.Field
		AuthorizedAt         respjson.Field
		Member               respjson.Field
		ActionResult         respjson.Field
		Assets               respjson.Field
		Owner                respjson.Field
		Shares               respjson.Field
		Receiver             respjson.Field
		raw                  string
	} `json:"-"`
}

// anyUnsafeUnwrapWebhookEvent is implemented by each variant of
// [UnsafeUnwrapWebhookEventUnion] to add type safety for the return type of
// [UnsafeUnwrapWebhookEventUnion.AsAny]
type anyUnsafeUnwrapWebhookEvent interface {
	implUnsafeUnwrapWebhookEventUnion()
}

func (UserCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                             {}
func (UserAuthenticatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (UserLinkedAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (UserUnlinkedAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (UserUpdatedAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                      {}
func (UserTransferredAccountWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (UserWalletCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (TransactionBroadcastedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (TransactionConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                    {}
func (TransactionExecutionRevertedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()            {}
func (TransactionStillPendingWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                 {}
func (TransactionFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                       {}
func (TransactionReplacedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (TransactionProviderErrorWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                {}
func (UserOperationCompletedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (FundsDepositedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (FundsWithdrawnWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (PrivateKeyExportWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                        {}
func (WalletRecoverySetupWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}
func (WalletRecoveredWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                         {}
func (WalletActionSwapCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                 {}
func (WalletActionSwapSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()               {}
func (WalletActionSwapRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                {}
func (WalletActionSwapFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (WalletActionTransferCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()             {}
func (WalletActionTransferSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()           {}
func (WalletActionTransferRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()            {}
func (WalletActionTransferFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()              {}
func (WalletActionEarnDepositCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()          {}
func (WalletActionEarnDepositSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()        {}
func (WalletActionEarnDepositRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()         {}
func (WalletActionEarnDepositFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()           {}
func (WalletActionEarnWithdrawCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()         {}
func (WalletActionEarnWithdrawSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion()       {}
func (WalletActionEarnWithdrawRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()        {}
func (WalletActionEarnWithdrawFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()          {}
func (WalletActionEarnIncentiveClaimCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()   {}
func (WalletActionEarnIncentiveClaimSucceededWebhookPayload) implUnsafeUnwrapWebhookEventUnion() {}
func (WalletActionEarnIncentiveClaimRejectedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()  {}
func (WalletActionEarnIncentiveClaimFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()    {}
func (MfaEnabledWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                              {}
func (MfaDisabledWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                             {}
func (IntentCreatedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                           {}
func (IntentAuthorizedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                        {}
func (IntentExecutedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                          {}
func (IntentFailedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                            {}
func (YieldDepositConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                   {}
func (YieldWithdrawConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                  {}
func (YieldClaimConfirmedWebhookPayload) implUnsafeUnwrapWebhookEventUnion()                     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UnsafeUnwrapWebhookEventUnion.AsAny().(type) {
//	case privyclient.UserCreatedWebhookPayload:
//	case privyclient.UserAuthenticatedWebhookPayload:
//	case privyclient.UserLinkedAccountWebhookPayload:
//	case privyclient.UserUnlinkedAccountWebhookPayload:
//	case privyclient.UserUpdatedAccountWebhookPayload:
//	case privyclient.UserTransferredAccountWebhookPayload:
//	case privyclient.UserWalletCreatedWebhookPayload:
//	case privyclient.TransactionBroadcastedWebhookPayload:
//	case privyclient.TransactionConfirmedWebhookPayload:
//	case privyclient.TransactionExecutionRevertedWebhookPayload:
//	case privyclient.TransactionStillPendingWebhookPayload:
//	case privyclient.TransactionFailedWebhookPayload:
//	case privyclient.TransactionReplacedWebhookPayload:
//	case privyclient.TransactionProviderErrorWebhookPayload:
//	case privyclient.UserOperationCompletedWebhookPayload:
//	case privyclient.FundsDepositedWebhookPayload:
//	case privyclient.FundsWithdrawnWebhookPayload:
//	case privyclient.PrivateKeyExportWebhookPayload:
//	case privyclient.WalletRecoverySetupWebhookPayload:
//	case privyclient.WalletRecoveredWebhookPayload:
//	case privyclient.WalletActionSwapCreatedWebhookPayload:
//	case privyclient.WalletActionSwapSucceededWebhookPayload:
//	case privyclient.WalletActionSwapRejectedWebhookPayload:
//	case privyclient.WalletActionSwapFailedWebhookPayload:
//	case privyclient.WalletActionTransferCreatedWebhookPayload:
//	case privyclient.WalletActionTransferSucceededWebhookPayload:
//	case privyclient.WalletActionTransferRejectedWebhookPayload:
//	case privyclient.WalletActionTransferFailedWebhookPayload:
//	case privyclient.WalletActionEarnDepositCreatedWebhookPayload:
//	case privyclient.WalletActionEarnDepositSucceededWebhookPayload:
//	case privyclient.WalletActionEarnDepositRejectedWebhookPayload:
//	case privyclient.WalletActionEarnDepositFailedWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawCreatedWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawSucceededWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawRejectedWebhookPayload:
//	case privyclient.WalletActionEarnWithdrawFailedWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimCreatedWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimSucceededWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimRejectedWebhookPayload:
//	case privyclient.WalletActionEarnIncentiveClaimFailedWebhookPayload:
//	case privyclient.MfaEnabledWebhookPayload:
//	case privyclient.MfaDisabledWebhookPayload:
//	case privyclient.IntentCreatedWebhookPayload:
//	case privyclient.IntentAuthorizedWebhookPayload:
//	case privyclient.IntentExecutedWebhookPayload:
//	case privyclient.IntentFailedWebhookPayload:
//	case privyclient.YieldDepositConfirmedWebhookPayload:
//	case privyclient.YieldWithdrawConfirmedWebhookPayload:
//	case privyclient.YieldClaimConfirmedWebhookPayload:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UnsafeUnwrapWebhookEventUnion) AsAny() anyUnsafeUnwrapWebhookEvent {
	switch u.Type {
	case "user.created":
		return u.AsUserCreated()
	case "user.authenticated":
		return u.AsUserAuthenticated()
	case "user.linked_account":
		return u.AsUserLinkedAccount()
	case "user.unlinked_account":
		return u.AsUserUnlinkedAccount()
	case "user.updated_account":
		return u.AsUserUpdatedAccount()
	case "user.transferred_account":
		return u.AsUserTransferredAccount()
	case "user.wallet_created":
		return u.AsUserWalletCreated()
	case "transaction.broadcasted":
		return u.AsTransactionBroadcasted()
	case "transaction.confirmed":
		return u.AsTransactionConfirmed()
	case "transaction.execution_reverted":
		return u.AsTransactionExecutionReverted()
	case "transaction.still_pending":
		return u.AsTransactionStillPending()
	case "transaction.failed":
		return u.AsTransactionFailed()
	case "transaction.replaced":
		return u.AsTransactionReplaced()
	case "transaction.provider_error":
		return u.AsTransactionProviderError()
	case "user_operation.completed":
		return u.AsUserOperationCompleted()
	case "wallet.funds_deposited":
		return u.AsWalletFundsDeposited()
	case "wallet.funds_withdrawn":
		return u.AsWalletFundsWithdrawn()
	case "wallet.private_key_export":
		return u.AsWalletPrivateKeyExport()
	case "wallet.recovery_setup":
		return u.AsWalletRecoverySetup()
	case "wallet.recovered":
		return u.AsWalletRecovered()
	case "wallet_action.swap.created":
		return u.AsWalletActionSwapCreated()
	case "wallet_action.swap.succeeded":
		return u.AsWalletActionSwapSucceeded()
	case "wallet_action.swap.rejected":
		return u.AsWalletActionSwapRejected()
	case "wallet_action.swap.failed":
		return u.AsWalletActionSwapFailed()
	case "wallet_action.transfer.created":
		return u.AsWalletActionTransferCreated()
	case "wallet_action.transfer.succeeded":
		return u.AsWalletActionTransferSucceeded()
	case "wallet_action.transfer.rejected":
		return u.AsWalletActionTransferRejected()
	case "wallet_action.transfer.failed":
		return u.AsWalletActionTransferFailed()
	case "wallet_action.earn_deposit.created":
		return u.AsWalletActionEarnDepositCreated()
	case "wallet_action.earn_deposit.succeeded":
		return u.AsWalletActionEarnDepositSucceeded()
	case "wallet_action.earn_deposit.rejected":
		return u.AsWalletActionEarnDepositRejected()
	case "wallet_action.earn_deposit.failed":
		return u.AsWalletActionEarnDepositFailed()
	case "wallet_action.earn_withdraw.created":
		return u.AsWalletActionEarnWithdrawCreated()
	case "wallet_action.earn_withdraw.succeeded":
		return u.AsWalletActionEarnWithdrawSucceeded()
	case "wallet_action.earn_withdraw.rejected":
		return u.AsWalletActionEarnWithdrawRejected()
	case "wallet_action.earn_withdraw.failed":
		return u.AsWalletActionEarnWithdrawFailed()
	case "wallet_action.earn_incentive_claim.created":
		return u.AsWalletActionEarnIncentiveClaimCreated()
	case "wallet_action.earn_incentive_claim.succeeded":
		return u.AsWalletActionEarnIncentiveClaimSucceeded()
	case "wallet_action.earn_incentive_claim.rejected":
		return u.AsWalletActionEarnIncentiveClaimRejected()
	case "wallet_action.earn_incentive_claim.failed":
		return u.AsWalletActionEarnIncentiveClaimFailed()
	case "mfa.enabled":
		return u.AsMfaEnabled()
	case "mfa.disabled":
		return u.AsMfaDisabled()
	case "intent.created":
		return u.AsIntentCreated()
	case "intent.authorized":
		return u.AsIntentAuthorized()
	case "intent.executed":
		return u.AsIntentExecuted()
	case "intent.failed":
		return u.AsIntentFailed()
	case "yield.deposit.confirmed":
		return u.AsYieldDepositConfirmed()
	case "yield.withdraw.confirmed":
		return u.AsYieldWithdrawConfirmed()
	case "yield.claim.confirmed":
		return u.AsYieldClaimConfirmed()
	}
	return nil
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserCreated() (v UserCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserAuthenticated() (v UserAuthenticatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserLinkedAccount() (v UserLinkedAccountWebhookPayload) {
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

func (u UnsafeUnwrapWebhookEventUnion) AsUserTransferredAccount() (v UserTransferredAccountWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserWalletCreated() (v UserWalletCreatedWebhookPayload) {
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

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionStillPending() (v TransactionStillPendingWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionFailed() (v TransactionFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionReplaced() (v TransactionReplacedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTransactionProviderError() (v TransactionProviderErrorWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUserOperationCompleted() (v UserOperationCompletedWebhookPayload) {
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

func (u UnsafeUnwrapWebhookEventUnion) AsWalletRecoverySetup() (v WalletRecoverySetupWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletRecovered() (v WalletRecoveredWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapCreated() (v WalletActionSwapCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapSucceeded() (v WalletActionSwapSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapRejected() (v WalletActionSwapRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionSwapFailed() (v WalletActionSwapFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferCreated() (v WalletActionTransferCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferSucceeded() (v WalletActionTransferSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferRejected() (v WalletActionTransferRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionTransferFailed() (v WalletActionTransferFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositCreated() (v WalletActionEarnDepositCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositSucceeded() (v WalletActionEarnDepositSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositRejected() (v WalletActionEarnDepositRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnDepositFailed() (v WalletActionEarnDepositFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawCreated() (v WalletActionEarnWithdrawCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawSucceeded() (v WalletActionEarnWithdrawSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawRejected() (v WalletActionEarnWithdrawRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnWithdrawFailed() (v WalletActionEarnWithdrawFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimCreated() (v WalletActionEarnIncentiveClaimCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimSucceeded() (v WalletActionEarnIncentiveClaimSucceededWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimRejected() (v WalletActionEarnIncentiveClaimRejectedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsWalletActionEarnIncentiveClaimFailed() (v WalletActionEarnIncentiveClaimFailedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsMfaEnabled() (v MfaEnabledWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsMfaDisabled() (v MfaDisabledWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentCreated() (v IntentCreatedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsIntentAuthorized() (v IntentAuthorizedWebhookPayload) {
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

func (u UnsafeUnwrapWebhookEventUnion) AsYieldDepositConfirmed() (v YieldDepositConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsYieldWithdrawConfirmed() (v YieldWithdrawConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsYieldClaimConfirmed() (v YieldClaimConfirmedWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnsafeUnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnsafeUnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
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
	// This field is a union of [any], [string], [string]
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

// UnsafeUnwrapWebhookEventUnionBlock is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionBlock provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionBlock struct {
	Number    float64 `json:"number"`
	Timestamp float64 `json:"timestamp"`
	JSON      struct {
		Number    respjson.Field
		Timestamp respjson.Field
		raw       string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionBlock) UnmarshalJSON(data []byte) error {
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
// will be valid: OfEarnIncetiveClaimRewardEntryArray
// OfYieldClaimConfirmedWebhookPayloadRewards]
type UnsafeUnwrapWebhookEventUnionRewards struct {
	// This field will be present if the value is a [[]EarnIncetiveClaimRewardEntry]
	// instead of an object.
	OfEarnIncetiveClaimRewardEntryArray []EarnIncetiveClaimRewardEntry `json:",inline"`
	// This field will be present if the value is a
	// [[]YieldClaimConfirmedWebhookPayloadReward] instead of an object.
	OfYieldClaimConfirmedWebhookPayloadRewards []YieldClaimConfirmedWebhookPayloadReward `json:",inline"`
	JSON                                       struct {
		OfEarnIncetiveClaimRewardEntryArray        respjson.Field
		OfYieldClaimConfirmedWebhookPayloadRewards respjson.Field
		raw                                        string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionRewards) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
