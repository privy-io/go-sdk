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

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/apiquery"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/pagination"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
	"github.com/privy-io/go-sdk/shared/constant"
)

// IntentService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntentService] method instead.
type IntentService struct {
	Options []option.RequestOption
}

// NewIntentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIntentService(opts ...option.RequestOption) (r IntentService) {
	r = IntentService{}
	r.Options = opts
	return
}

// List intents for an app. Returns a paginated list of intents with their current
// status and details.
func (r *IntentService) List(ctx context.Context, query IntentListParams, opts ...option.RequestOption) (res *pagination.Cursor[IntentResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/intents"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List intents for an app. Returns a paginated list of intents with their current
// status and details.
func (r *IntentService) ListAutoPaging(ctx context.Context, query IntentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[IntentResponseUnion] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create an intent to add a rule to a policy. The intent must be authorized by the
// policy owner before it can be executed.
func (r *IntentService) NewPolicyRule(ctx context.Context, policyID string, params IntentNewPolicyRuleParams, opts ...option.RequestOption) (res *RuleIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s/rules", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create an intent to delete a rule from a policy. The intent must be authorized
// by the policy owner before it can be executed.
func (r *IntentService) DeletePolicyRule(ctx context.Context, ruleID string, params IntentDeletePolicyRuleParams, opts ...option.RequestOption) (res *RuleIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s/rules/%s", url.PathEscape(params.PolicyID), url.PathEscape(ruleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve an intent by ID. Returns the intent details including its current
// status, authorization details, and execution result if applicable.
func (r *IntentService) Get(ctx context.Context, intentID string, opts ...option.RequestOption) (res *IntentResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if intentID == "" {
		err = errors.New("missing required intent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/%s", url.PathEscape(intentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create an intent to execute an RPC method on a wallet. The intent must be
// authorized by either the wallet owner or signers before it can be executed.
func (r *IntentService) Rpc(ctx context.Context, walletID string, params IntentRpcParams, opts ...option.RequestOption) (res *RpcIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/wallets/%s/rpc", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create an intent to update a key quorum. The intent must be authorized by the
// key quorum members before it can be executed.
func (r *IntentService) UpdateKeyQuorum(ctx context.Context, keyQuorumID string, params IntentUpdateKeyQuorumParams, opts ...option.RequestOption) (res *KeyQuorumIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/key_quorums/%s", url.PathEscape(keyQuorumID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Create an intent to update a policy. The intent must be authorized by the policy
// owner before it can be executed.
func (r *IntentService) UpdatePolicy(ctx context.Context, policyID string, params IntentUpdatePolicyParams, opts ...option.RequestOption) (res *PolicyIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Create an intent to update a rule on a policy. The intent must be authorized by
// the policy owner before it can be executed.
func (r *IntentService) UpdatePolicyRule(ctx context.Context, ruleID string, params IntentUpdatePolicyRuleParams, opts ...option.RequestOption) (res *RuleIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s/rules/%s", url.PathEscape(params.PolicyID), url.PathEscape(ruleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Create an intent to update a wallet. The intent must be authorized by the wallet
// owner before it can be executed.
func (r *IntentService) UpdateWallet(ctx context.Context, walletID string, params IntentUpdateWalletParams, opts ...option.RequestOption) (res *WalletIntentResponse, err error) {
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/wallets/%s", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Type of intent.
type IntentType string

const (
	IntentTypeKeyQuorum IntentType = "KEY_QUORUM"
	IntentTypePolicy    IntentType = "POLICY"
	IntentTypeRule      IntentType = "RULE"
	IntentTypeRpc       IntentType = "RPC"
	IntentTypeTransfer  IntentType = "TRANSFER"
	IntentTypeWallet    IntentType = "WALLET"
)

// Current status of an intent.
type IntentStatus string

const (
	IntentStatusPending   IntentStatus = "pending"
	IntentStatusExecuted  IntentStatus = "executed"
	IntentStatusFailed    IntentStatus = "failed"
	IntentStatusExpired   IntentStatus = "expired"
	IntentStatusRejected  IntentStatus = "rejected"
	IntentStatusDismissed IntentStatus = "dismissed"
)

// Request details for creating a rule via intent.
type RuleIntentCreateRequestDetails struct {
	// The rules that apply to each method the policy covers.
	Body RuleIntentCreateRequestDetailsBody `json:"body" api:"required"`
	// Any of "POST".
	Method RuleIntentCreateRequestDetailsMethod `json:"method" api:"required"`
	URL    string                               `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentCreateRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rules that apply to each method the policy covers.
type RuleIntentCreateRequestDetailsBody struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                                             `json:"action" api:"required"`
	Conditions []RuleIntentCreateRequestDetailsBodyConditionUnion `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Conditions  respjson.Field
		Method      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBody) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentCreateRequestDetailsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionUnion contains all possible
// properties and values from
// [RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction],
// [RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata],
// [RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain],
// [RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage],
// [RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization],
// [RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction],
// [RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction],
// [RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction],
// [RuleIntentCreateRequestDetailsBodyConditionSystem], [TronTransactionCondition],
// [TronCalldataCondition], [SuiTransactionCommandCondition],
// [SuiTransferObjectsCommandCondition], [AggregationCondition].
//
// Use the [RuleIntentCreateRequestDetailsBodyConditionUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentCreateRequestDetailsBodyConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "tron_trigger_smart_contract_data",
	// "sui_transaction_command", "sui_transfer_objects_command", "reference".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion],
	// [RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion], [TronCalldataConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion], [AggregationConditionValueUnion]
	Value RuleIntentCreateRequestDetailsBodyConditionUnionValue `json:"value"`
	Abi   any                                                   `json:"abi"`
	// This field is from variant
	// [RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage].
	TypedData RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
	JSON      struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		Abi         respjson.Field
		TypedData   respjson.Field
		raw         string
	} `json:"-"`
}

// anyRuleIntentCreateRequestDetailsBodyCondition is implemented by each variant of
// [RuleIntentCreateRequestDetailsBodyConditionUnion] to add type safety for the
// return type of [RuleIntentCreateRequestDetailsBodyConditionUnion.AsAny]
type anyRuleIntentCreateRequestDetailsBodyCondition interface {
	implRuleIntentCreateRequestDetailsBodyConditionUnion()
}

func (RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentCreateRequestDetailsBodyConditionSystem) implRuleIntentCreateRequestDetailsBodyConditionUnion() {
}
func (TronTransactionCondition) implRuleIntentCreateRequestDetailsBodyConditionUnion()           {}
func (TronCalldataCondition) implRuleIntentCreateRequestDetailsBodyConditionUnion()              {}
func (SuiTransactionCommandCondition) implRuleIntentCreateRequestDetailsBodyConditionUnion()     {}
func (SuiTransferObjectsCommandCondition) implRuleIntentCreateRequestDetailsBodyConditionUnion() {}
func (AggregationCondition) implRuleIntentCreateRequestDetailsBodyConditionUnion()               {}

// Use the following switch statement to find the correct variant
//
//	switch variant := RuleIntentCreateRequestDetailsBodyConditionUnion.AsAny().(type) {
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction:
//	case privyclient.RuleIntentCreateRequestDetailsBodyConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.TronCalldataCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	case privyclient.AggregationCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsAny() anyRuleIntentCreateRequestDetailsBodyCondition {
	switch u.FieldSource {
	case "ethereum_transaction":
		return u.AsEthereumTransaction()
	case "ethereum_calldata":
		return u.AsEthereumCalldata()
	case "ethereum_typed_data_domain":
		return u.AsEthereumTypedDataDomain()
	case "ethereum_typed_data_message":
		return u.AsEthereumTypedDataMessage()
	case "ethereum_7702_authorization":
		return u.AsEthereum7702Authorization()
	case "solana_program_instruction":
		return u.AsSolanaProgramInstruction()
	case "solana_system_program_instruction":
		return u.AsSolanaSystemProgramInstruction()
	case "solana_token_program_instruction":
		return u.AsSolanaTokenProgramInstruction()
	case "system":
		return u.AsSystem()
	case "tron_transaction":
		return u.AsTronTransaction()
	case "tron_trigger_smart_contract_data":
		return u.AsTronTriggerSmartContractData()
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	case "reference":
		return u.AsReference()
	}
	return nil
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsEthereumTransaction() (v RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsEthereumCalldata() (v RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsEthereumTypedDataDomain() (v RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsEthereumTypedDataMessage() (v RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsEthereum7702Authorization() (v RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsSolanaProgramInstruction() (v RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsSolanaSystemProgramInstruction() (v RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsSolanaTokenProgramInstruction() (v RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsSystem() (v RuleIntentCreateRequestDetailsBodyConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsTronTriggerSmartContractData() (v TronCalldataCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsReference() (v AggregationCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentCreateRequestDetailsBodyConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionUnionValue is an implicit subunion of
// [RuleIntentCreateRequestDetailsBodyConditionUnion].
// RuleIntentCreateRequestDetailsBodyConditionUnionValue provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [RuleIntentCreateRequestDetailsBodyConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type RuleIntentCreateRequestDetailsBodyConditionUnionValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a [SuiCommandName] instead of an
	// object.
	OfSuiCommandName SuiCommandName `json:",inline"`
	// This field will be present if the value is a [[]SuiCommandName] instead of an
	// object.
	OfSuiCommandNameArray []SuiCommandName `json:",inline"`
	JSON                  struct {
		OfString              respjson.Field
		OfStringArray         respjson.Field
		OfSuiCommandName      respjson.Field
		OfSuiCommandNameArray respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *RuleIntentCreateRequestDetailsBodyConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field" api:"required"`
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                   `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata struct {
	Abi         any                       `json:"abi" api:"required"`
	Field       string                    `json:"field" api:"required"`
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abi         respjson.Field
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field       string                           `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                       `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                                        `json:"operator" api:"required"`
	TypedData RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData  `json:"typed_data" api:"required"`
	Value     RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		TypedData   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParamsResp `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
type RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field" api:"required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                         `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                        `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field" api:"required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                              `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field       string                                 `json:"field" api:"required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                             `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type RuleIntentCreateRequestDetailsBodyConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field" api:"required"`
	FieldSource constant.System `json:"field_source" default:"system"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                      `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentCreateRequestDetailsBodyConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentCreateRequestDetailsBodyConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentCreateRequestDetailsMethod string

const (
	RuleIntentCreateRequestDetailsMethodPost RuleIntentCreateRequestDetailsMethod = "POST"
)

// Request details for updating a rule via intent.
type RuleIntentUpdateRequestDetails struct {
	// The rules that apply to each method the policy covers.
	Body RuleIntentUpdateRequestDetailsBody `json:"body" api:"required"`
	// Any of "PATCH".
	Method RuleIntentUpdateRequestDetailsMethod `json:"method" api:"required"`
	URL    string                               `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentUpdateRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rules that apply to each method the policy covers.
type RuleIntentUpdateRequestDetailsBody struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                                             `json:"action" api:"required"`
	Conditions []RuleIntentUpdateRequestDetailsBodyConditionUnion `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Conditions  respjson.Field
		Method      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBody) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentUpdateRequestDetailsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionUnion contains all possible
// properties and values from
// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction],
// [RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata],
// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain],
// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage],
// [RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization],
// [RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction],
// [RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction],
// [RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction],
// [RuleIntentUpdateRequestDetailsBodyConditionSystem], [TronTransactionCondition],
// [TronCalldataCondition], [SuiTransactionCommandCondition],
// [SuiTransferObjectsCommandCondition], [AggregationCondition].
//
// Use the [RuleIntentUpdateRequestDetailsBodyConditionUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentUpdateRequestDetailsBodyConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "tron_trigger_smart_contract_data",
	// "sui_transaction_command", "sui_transfer_objects_command", "reference".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion],
	// [RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion], [TronCalldataConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion], [AggregationConditionValueUnion]
	Value RuleIntentUpdateRequestDetailsBodyConditionUnionValue `json:"value"`
	Abi   any                                                   `json:"abi"`
	// This field is from variant
	// [RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage].
	TypedData RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
	JSON      struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		Abi         respjson.Field
		TypedData   respjson.Field
		raw         string
	} `json:"-"`
}

// anyRuleIntentUpdateRequestDetailsBodyCondition is implemented by each variant of
// [RuleIntentUpdateRequestDetailsBodyConditionUnion] to add type safety for the
// return type of [RuleIntentUpdateRequestDetailsBodyConditionUnion.AsAny]
type anyRuleIntentUpdateRequestDetailsBodyCondition interface {
	implRuleIntentUpdateRequestDetailsBodyConditionUnion()
}

func (RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (RuleIntentUpdateRequestDetailsBodyConditionSystem) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {
}
func (TronTransactionCondition) implRuleIntentUpdateRequestDetailsBodyConditionUnion()           {}
func (TronCalldataCondition) implRuleIntentUpdateRequestDetailsBodyConditionUnion()              {}
func (SuiTransactionCommandCondition) implRuleIntentUpdateRequestDetailsBodyConditionUnion()     {}
func (SuiTransferObjectsCommandCondition) implRuleIntentUpdateRequestDetailsBodyConditionUnion() {}
func (AggregationCondition) implRuleIntentUpdateRequestDetailsBodyConditionUnion()               {}

// Use the following switch statement to find the correct variant
//
//	switch variant := RuleIntentUpdateRequestDetailsBodyConditionUnion.AsAny().(type) {
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction:
//	case privyclient.RuleIntentUpdateRequestDetailsBodyConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.TronCalldataCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	case privyclient.AggregationCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsAny() anyRuleIntentUpdateRequestDetailsBodyCondition {
	switch u.FieldSource {
	case "ethereum_transaction":
		return u.AsEthereumTransaction()
	case "ethereum_calldata":
		return u.AsEthereumCalldata()
	case "ethereum_typed_data_domain":
		return u.AsEthereumTypedDataDomain()
	case "ethereum_typed_data_message":
		return u.AsEthereumTypedDataMessage()
	case "ethereum_7702_authorization":
		return u.AsEthereum7702Authorization()
	case "solana_program_instruction":
		return u.AsSolanaProgramInstruction()
	case "solana_system_program_instruction":
		return u.AsSolanaSystemProgramInstruction()
	case "solana_token_program_instruction":
		return u.AsSolanaTokenProgramInstruction()
	case "system":
		return u.AsSystem()
	case "tron_transaction":
		return u.AsTronTransaction()
	case "tron_trigger_smart_contract_data":
		return u.AsTronTriggerSmartContractData()
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	case "reference":
		return u.AsReference()
	}
	return nil
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsEthereumTransaction() (v RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsEthereumCalldata() (v RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsEthereumTypedDataDomain() (v RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsEthereumTypedDataMessage() (v RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsEthereum7702Authorization() (v RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsSolanaProgramInstruction() (v RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsSolanaSystemProgramInstruction() (v RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsSolanaTokenProgramInstruction() (v RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsSystem() (v RuleIntentUpdateRequestDetailsBodyConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsTronTriggerSmartContractData() (v TronCalldataCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsReference() (v AggregationCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentUpdateRequestDetailsBodyConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionUnionValue is an implicit subunion of
// [RuleIntentUpdateRequestDetailsBodyConditionUnion].
// RuleIntentUpdateRequestDetailsBodyConditionUnionValue provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [RuleIntentUpdateRequestDetailsBodyConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type RuleIntentUpdateRequestDetailsBodyConditionUnionValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a [SuiCommandName] instead of an
	// object.
	OfSuiCommandName SuiCommandName `json:",inline"`
	// This field will be present if the value is a [[]SuiCommandName] instead of an
	// object.
	OfSuiCommandNameArray []SuiCommandName `json:",inline"`
	JSON                  struct {
		OfString              respjson.Field
		OfStringArray         respjson.Field
		OfSuiCommandName      respjson.Field
		OfSuiCommandNameArray respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field" api:"required"`
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                   `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata struct {
	Abi         any                       `json:"abi" api:"required"`
	Field       string                    `json:"field" api:"required"`
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abi         respjson.Field
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field       string                           `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                       `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                                        `json:"operator" api:"required"`
	TypedData RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData  `json:"typed_data" api:"required"`
	Value     RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		TypedData   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParamsResp `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
type RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field" api:"required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                         `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                        `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field" api:"required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                              `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field       string                                 `json:"field" api:"required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                             `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type RuleIntentUpdateRequestDetailsBodyConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field" api:"required"`
	FieldSource constant.System `json:"field_source" default:"system"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                      `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentUpdateRequestDetailsBodyConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentUpdateRequestDetailsBodyConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentUpdateRequestDetailsMethod string

const (
	RuleIntentUpdateRequestDetailsMethodPatch RuleIntentUpdateRequestDetailsMethod = "PATCH"
)

// Request details for deleting a rule via intent.
type RuleIntentDeleteRequestDetails struct {
	// Any of "DELETE".
	Method RuleIntentDeleteRequestDetailsMethod `json:"method" api:"required"`
	URL    string                               `json:"url" api:"required"`
	Body   any                                  `json:"body"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		URL         respjson.Field
		Body        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentDeleteRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentDeleteRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentDeleteRequestDetailsMethod string

const (
	RuleIntentDeleteRequestDetailsMethodDelete RuleIntentDeleteRequestDetailsMethod = "DELETE"
)

// RuleIntentRequestDetailsUnion contains all possible properties and values from
// [RuleIntentCreateRequestDetails], [RuleIntentUpdateRequestDetails],
// [RuleIntentDeleteRequestDetails].
//
// Use the [RuleIntentRequestDetailsUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentRequestDetailsUnion struct {
	// This field is a union of [RuleIntentCreateRequestDetailsBody],
	// [RuleIntentUpdateRequestDetailsBody], [any]
	Body RuleIntentRequestDetailsUnionBody `json:"body"`
	// Any of "POST", "PATCH", "DELETE".
	Method string `json:"method"`
	URL    string `json:"url"`
	JSON   struct {
		Body   respjson.Field
		Method respjson.Field
		URL    respjson.Field
		raw    string
	} `json:"-"`
}

// anyRuleIntentRequestDetails is implemented by each variant of
// [RuleIntentRequestDetailsUnion] to add type safety for the return type of
// [RuleIntentRequestDetailsUnion.AsAny]
type anyRuleIntentRequestDetails interface {
	implRuleIntentRequestDetailsUnion()
}

func (RuleIntentCreateRequestDetails) implRuleIntentRequestDetailsUnion() {}
func (RuleIntentUpdateRequestDetails) implRuleIntentRequestDetailsUnion() {}
func (RuleIntentDeleteRequestDetails) implRuleIntentRequestDetailsUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := RuleIntentRequestDetailsUnion.AsAny().(type) {
//	case privyclient.RuleIntentCreateRequestDetails:
//	case privyclient.RuleIntentUpdateRequestDetails:
//	case privyclient.RuleIntentDeleteRequestDetails:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u RuleIntentRequestDetailsUnion) AsAny() anyRuleIntentRequestDetails {
	switch u.Method {
	case "POST":
		return u.AsPost()
	case "PATCH":
		return u.AsPatch()
	case "DELETE":
		return u.AsDelete()
	}
	return nil
}

func (u RuleIntentRequestDetailsUnion) AsPost() (v RuleIntentCreateRequestDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentRequestDetailsUnion) AsPatch() (v RuleIntentUpdateRequestDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentRequestDetailsUnion) AsDelete() (v RuleIntentDeleteRequestDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentRequestDetailsUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentRequestDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentRequestDetailsUnionBody is an implicit subunion of
// [RuleIntentRequestDetailsUnion]. RuleIntentRequestDetailsUnionBody provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [RuleIntentRequestDetailsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfRuleIntentDeleteRequestDetailsBody]
type RuleIntentRequestDetailsUnionBody struct {
	// This field will be present if the value is a [any] instead of an object.
	OfRuleIntentDeleteRequestDetailsBody any    `json:",inline"`
	Action                               string `json:"action"`
	// This field is a union of [[]RuleIntentCreateRequestDetailsBodyConditionUnion],
	// [[]RuleIntentUpdateRequestDetailsBodyConditionUnion]
	Conditions RuleIntentRequestDetailsUnionBodyConditions `json:"conditions"`
	Method     string                                      `json:"method"`
	Name       string                                      `json:"name"`
	JSON       struct {
		OfRuleIntentDeleteRequestDetailsBody respjson.Field
		Action                               respjson.Field
		Conditions                           respjson.Field
		Method                               respjson.Field
		Name                                 respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *RuleIntentRequestDetailsUnionBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentRequestDetailsUnionBodyConditions is an implicit subunion of
// [RuleIntentRequestDetailsUnion]. RuleIntentRequestDetailsUnionBodyConditions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [RuleIntentRequestDetailsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfRuleIntentCreateRequestDetailsBodyConditions
// OfRuleIntentUpdateRequestDetailsBodyConditions]
type RuleIntentRequestDetailsUnionBodyConditions struct {
	// This field will be present if the value is a
	// [[]RuleIntentCreateRequestDetailsBodyConditionUnion] instead of an object.
	OfRuleIntentCreateRequestDetailsBodyConditions []RuleIntentCreateRequestDetailsBodyConditionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]RuleIntentUpdateRequestDetailsBodyConditionUnion] instead of an object.
	OfRuleIntentUpdateRequestDetailsBodyConditions []RuleIntentUpdateRequestDetailsBodyConditionUnion `json:",inline"`
	JSON                                           struct {
		OfRuleIntentCreateRequestDetailsBodyConditions respjson.Field
		OfRuleIntentUpdateRequestDetailsBodyConditions respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *RuleIntentRequestDetailsUnionBodyConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentAuthorizationKeyQuorumMemberUnion contains all possible properties and
// values from [IntentAuthorizationKeyQuorumMemberUserMember],
// [IntentAuthorizationKeyQuorumMemberKeyMember].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IntentAuthorizationKeyQuorumMemberUnion struct {
	SignedAt float64 `json:"signed_at"`
	Type     string  `json:"type"`
	// This field is from variant [IntentAuthorizationKeyQuorumMemberUserMember].
	UserID string `json:"user_id"`
	// This field is from variant [IntentAuthorizationKeyQuorumMemberKeyMember].
	PublicKey string `json:"public_key"`
	JSON      struct {
		SignedAt  respjson.Field
		Type      respjson.Field
		UserID    respjson.Field
		PublicKey respjson.Field
		raw       string
	} `json:"-"`
}

func (u IntentAuthorizationKeyQuorumMemberUnion) AsUserMember() (v IntentAuthorizationKeyQuorumMemberUserMember) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentAuthorizationKeyQuorumMemberUnion) AsKeyMember() (v IntentAuthorizationKeyQuorumMemberKeyMember) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IntentAuthorizationKeyQuorumMemberUnion) RawJSON() string { return u.JSON.raw }

func (r *IntentAuthorizationKeyQuorumMemberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentAuthorizationKeyQuorumMemberUserMember struct {
	// Unix timestamp when this member signed, or null if not yet signed.
	SignedAt float64 `json:"signed_at" api:"required"`
	// Any of "user".
	Type string `json:"type" api:"required"`
	// User ID of the key quorum member
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignedAt    respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorizationKeyQuorumMemberUserMember) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorizationKeyQuorumMemberUserMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentAuthorizationKeyQuorumMemberKeyMember struct {
	// Public key of the key quorum member
	PublicKey string `json:"public_key" api:"required"`
	// Unix timestamp when this member signed, or null if not yet signed.
	SignedAt float64 `json:"signed_at" api:"required"`
	// Any of "key".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		SignedAt    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorizationKeyQuorumMemberKeyMember) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorizationKeyQuorumMemberKeyMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentAuthorizationMemberUnion contains all possible properties and values from
// [IntentAuthorizationMemberUserMember], [IntentAuthorizationMemberKeyMember],
// [IntentAuthorizationMemberKeyQuorumMember].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IntentAuthorizationMemberUnion struct {
	SignedAt float64 `json:"signed_at"`
	Type     string  `json:"type"`
	// This field is from variant [IntentAuthorizationMemberUserMember].
	UserID string `json:"user_id"`
	// This field is from variant [IntentAuthorizationMemberKeyMember].
	PublicKey string `json:"public_key"`
	// This field is from variant [IntentAuthorizationMemberKeyQuorumMember].
	KeyQuorumID string `json:"key_quorum_id"`
	// This field is from variant [IntentAuthorizationMemberKeyQuorumMember].
	Members []IntentAuthorizationKeyQuorumMemberUnion `json:"members"`
	// This field is from variant [IntentAuthorizationMemberKeyQuorumMember].
	Threshold float64 `json:"threshold"`
	// This field is from variant [IntentAuthorizationMemberKeyQuorumMember].
	ThresholdMet bool `json:"threshold_met"`
	// This field is from variant [IntentAuthorizationMemberKeyQuorumMember].
	DisplayName string `json:"display_name"`
	JSON        struct {
		SignedAt     respjson.Field
		Type         respjson.Field
		UserID       respjson.Field
		PublicKey    respjson.Field
		KeyQuorumID  respjson.Field
		Members      respjson.Field
		Threshold    respjson.Field
		ThresholdMet respjson.Field
		DisplayName  respjson.Field
		raw          string
	} `json:"-"`
}

func (u IntentAuthorizationMemberUnion) AsUserMember() (v IntentAuthorizationMemberUserMember) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentAuthorizationMemberUnion) AsKeyMember() (v IntentAuthorizationMemberKeyMember) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentAuthorizationMemberUnion) AsKeyQuorumMember() (v IntentAuthorizationMemberKeyQuorumMember) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IntentAuthorizationMemberUnion) RawJSON() string { return u.JSON.raw }

func (r *IntentAuthorizationMemberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentAuthorizationMemberUserMember struct {
	// Unix timestamp when this member signed, or null if not yet signed.
	SignedAt float64 `json:"signed_at" api:"required"`
	// Any of "user".
	Type string `json:"type" api:"required"`
	// User ID of the key quorum member
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignedAt    respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorizationMemberUserMember) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorizationMemberUserMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentAuthorizationMemberKeyMember struct {
	// Public key of the key quorum member
	PublicKey string `json:"public_key" api:"required"`
	// Unix timestamp when this member signed, or null if not yet signed.
	SignedAt float64 `json:"signed_at" api:"required"`
	// Any of "key".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		SignedAt    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorizationMemberKeyMember) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorizationMemberKeyMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentAuthorizationMemberKeyQuorumMember struct {
	// ID of the child key quorum member
	KeyQuorumID string `json:"key_quorum_id" api:"required"`
	// Members of this child quorum
	Members []IntentAuthorizationKeyQuorumMemberUnion `json:"members" api:"required"`
	// Number of signatures required from this child quorum
	Threshold float64 `json:"threshold" api:"required"`
	// Whether this child key quorum has met its signature threshold
	ThresholdMet bool `json:"threshold_met" api:"required"`
	// Any of "key_quorum".
	Type string `json:"type" api:"required"`
	// Display name for the child key quorum (if any)
	DisplayName string `json:"display_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyQuorumID  respjson.Field
		Members      respjson.Field
		Threshold    respjson.Field
		ThresholdMet respjson.Field
		Type         respjson.Field
		DisplayName  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorizationMemberKeyQuorumMember) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorizationMemberKeyQuorumMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authorization quorum for an intent
type IntentAuthorization struct {
	// Members in this authorization quorum
	Members []IntentAuthorizationMemberUnion `json:"members" api:"required"`
	// Number of signatures required to satisfy this quorum
	Threshold float64 `json:"threshold" api:"required"`
	// Display name of the key quorum
	DisplayName string `json:"display_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Members     respjson.Field
		Threshold   respjson.Field
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntentAuthorization) RawJSON() string { return r.JSON.raw }
func (r *IntentAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Common fields shared by all intent response types.
type BaseIntentResponse struct {
	// Detailed authorization information including key quorum members, thresholds, and
	// signature status
	AuthorizationDetails []IntentAuthorization `json:"authorization_details" api:"required"`
	// Unix timestamp when the intent was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// Display name of the user who created the intent
	CreatedByDisplayName string `json:"created_by_display_name" api:"required"`
	// Whether this intent has a custom expiry time set by the client. If false, the
	// intent expires after a default duration.
	CustomExpiry bool `json:"custom_expiry" api:"required"`
	// Unix timestamp when the intent expires
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// Unique ID for the intent
	IntentID string `json:"intent_id" api:"required"`
	// ID of the resource being modified (wallet_id, policy_id, etc)
	ResourceID string `json:"resource_id" api:"required"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `json:"status" api:"required"`
	// ID of the user who created the intent. If undefined, the intent was created
	// using the app secret
	CreatedByID string `json:"created_by_id"`
	// Human-readable reason for dismissal, present when status is 'dismissed'
	DismissalReason string `json:"dismissal_reason"`
	// Unix timestamp when the intent was dismissed, present when status is 'dismissed'
	DismissedAt float64 `json:"dismissed_at"`
	// Unix timestamp when the intent was rejected, present when status is 'rejected'
	RejectedAt float64 `json:"rejected_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationDetails respjson.Field
		CreatedAt            respjson.Field
		CreatedByDisplayName respjson.Field
		CustomExpiry         respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		CreatedByID          respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *BaseIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Common fields for intent action execution results.
type BaseActionResult struct {
	// Unix timestamp when the action was executed
	ExecutedAt float64 `json:"executed_at" api:"required"`
	// HTTP status code from the action execution
	StatusCode float64 `json:"status_code" api:"required"`
	// Display name of the key quorum that authorized execution
	AuthorizedByDisplayName string `json:"authorized_by_display_name"`
	// ID of the key quorum that authorized execution
	AuthorizedByID string `json:"authorized_by_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExecutedAt              respjson.Field
		StatusCode              respjson.Field
		AuthorizedByDisplayName respjson.Field
		AuthorizedByID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseActionResult) RawJSON() string { return r.JSON.raw }
func (r *BaseActionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for an RPC intent
type RpcIntentResponse struct {
	// Any of "RPC".
	IntentType string `json:"intent_type" api:"required"`
	// The original RPC request that would be sent to the wallet endpoint
	RequestDetails RpcIntentResponseRequestDetails `json:"request_details" api:"required"`
	// Result of RPC execution (only present if status is 'executed' or 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// A wallet managed by Privy's wallet infrastructure.
	CurrentResourceData Wallet `json:"current_resource_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntentType          respjson.Field
		RequestDetails      respjson.Field
		ActionResult        respjson.Field
		CurrentResourceData respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
	BaseIntentResponse
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *RpcIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original RPC request that would be sent to the wallet endpoint
type RpcIntentResponseRequestDetails struct {
	// Request body for wallet RPC operations, discriminated by method.
	Body WalletRpcRequestBodyUnion `json:"body" api:"required"`
	// Any of "POST".
	Method string `json:"method" api:"required"`
	URL    string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponseRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *RpcIntentResponseRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a transfer intent
type TransferIntentResponse struct {
	// Any of "TRANSFER".
	IntentType string `json:"intent_type" api:"required"`
	// The original transfer request that would be sent to the wallet transfer endpoint
	RequestDetails TransferIntentResponseRequestDetails `json:"request_details" api:"required"`
	// Result of transfer execution (only present if intent status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// A wallet managed by Privy's wallet infrastructure.
	CurrentResourceData Wallet `json:"current_resource_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntentType          respjson.Field
		RequestDetails      respjson.Field
		ActionResult        respjson.Field
		CurrentResourceData respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
	BaseIntentResponse
}

// Returns the unmodified JSON received from the API
func (r TransferIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original transfer request that would be sent to the wallet transfer endpoint
type TransferIntentResponseRequestDetails struct {
	// Request body for initiating a sponsored token transfer from an embedded wallet.
	Body CreateTokenTransferRequest `json:"body" api:"required"`
	// Any of "POST".
	Method string `json:"method" api:"required"`
	URL    string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferIntentResponseRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *TransferIntentResponseRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a wallet intent
type WalletIntentResponse struct {
	// Any of "WALLET".
	IntentType string `json:"intent_type" api:"required"`
	// The original wallet update request that would be sent to the wallet endpoint
	RequestDetails WalletIntentResponseRequestDetails `json:"request_details" api:"required"`
	// Result of wallet update execution (only present if status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// A wallet managed by Privy's wallet infrastructure.
	CurrentResourceData Wallet `json:"current_resource_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntentType          respjson.Field
		RequestDetails      respjson.Field
		ActionResult        respjson.Field
		CurrentResourceData respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
	BaseIntentResponse
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original wallet update request that would be sent to the wallet endpoint
type WalletIntentResponseRequestDetails struct {
	Body WalletIntentResponseRequestDetailsBody `json:"body" api:"required"`
	// Any of "PATCH".
	Method string `json:"method" api:"required"`
	URL    string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponseRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponseRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletIntentResponseRequestDetailsBody struct {
	// Additional signers for the wallet.
	AdditionalSigners      AdditionalSignerInput `json:"additional_signers"`
	AuthorizationKeyIDs    []string              `json:"authorization_key_ids"`
	AuthorizationThreshold float64               `json:"authorization_threshold"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner   WalletIntentResponseRequestDetailsBodyOwner `json:"owner"`
	OwnerID string                                      `json:"owner_id" api:"nullable"`
	// An optional list of up to one policy ID to enforce on the wallet.
	PolicyIDs PolicyInput `json:"policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalSigners      respjson.Field
		AuthorizationKeyIDs    respjson.Field
		AuthorizationThreshold respjson.Field
		Owner                  respjson.Field
		OwnerID                respjson.Field
		PolicyIDs              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponseRequestDetailsBody) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponseRequestDetailsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the resource. If you provide this, do not specify an owner_id as it
// will be generated automatically. When updating a wallet, you can set the owner
// to null to remove the owner.
type WalletIntentResponseRequestDetailsBodyOwner struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	OwnerInputUnion
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponseRequestDetailsBodyOwner) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponseRequestDetailsBodyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a policy intent
type PolicyIntentResponse struct {
	// Any of "POLICY".
	IntentType string `json:"intent_type" api:"required"`
	// The original policy update request that would be sent to the policy endpoint
	RequestDetails PolicyIntentResponseRequestDetails `json:"request_details" api:"required"`
	// Result of policy update execution (only present if status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// A policy for controlling wallet operations.
	CurrentResourceData Policy `json:"current_resource_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntentType          respjson.Field
		RequestDetails      respjson.Field
		ActionResult        respjson.Field
		CurrentResourceData respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
	BaseIntentResponse
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original policy update request that would be sent to the policy endpoint
type PolicyIntentResponseRequestDetails struct {
	Body PolicyIntentResponseRequestDetailsBody `json:"body" api:"required"`
	// Any of "PATCH".
	Method string `json:"method" api:"required"`
	URL    string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBody struct {
	// Name to assign to policy.
	Name string `json:"name"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner   PolicyIntentResponseRequestDetailsBodyOwnerUnion `json:"owner" api:"nullable"`
	OwnerID string                                           `json:"owner_id" api:"nullable"`
	Rules   []PolicyIntentResponseRequestDetailsBodyRule     `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Owner       respjson.Field
		OwnerID     respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBody) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseRequestDetailsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyOwnerUnion contains all possible
// properties and values from
// [PolicyIntentResponseRequestDetailsBodyOwnerPublicKeyOwner],
// [PolicyIntentResponseRequestDetailsBodyOwnerUserOwner].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyIntentResponseRequestDetailsBodyOwnerUnion struct {
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyOwnerPublicKeyOwner].
	PublicKey string `json:"public_key"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyOwnerUserOwner].
	UserID string `json:"user_id"`
	JSON   struct {
		PublicKey respjson.Field
		UserID    respjson.Field
		raw       string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyOwnerUnion) AsPublicKeyOwner() (v PolicyIntentResponseRequestDetailsBodyOwnerPublicKeyOwner) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyOwnerUnion) AsUserOwner() (v PolicyIntentResponseRequestDetailsBodyOwnerUserOwner) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyOwnerUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyIntentResponseRequestDetailsBodyOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
type PolicyIntentResponseRequestDetailsBodyOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyOwnerPublicKeyOwner) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
type PolicyIntentResponseRequestDetailsBodyOwnerUserOwner struct {
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyOwnerUserOwner) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseRequestDetailsBodyOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rules that apply to each method the policy covers.
type PolicyIntentResponseRequestDetailsBodyRule struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                                                     `json:"action" api:"required"`
	Conditions []PolicyIntentResponseRequestDetailsBodyRuleConditionUnion `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Conditions  respjson.Field
		Method      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRule) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseRequestDetailsBodyRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionUnion contains all possible
// properties and values from
// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionSystem],
// [TronTransactionCondition], [TronCalldataCondition],
// [SuiTransactionCommandCondition], [SuiTransferObjectsCommandCondition],
// [AggregationCondition].
//
// Use the [PolicyIntentResponseRequestDetailsBodyRuleConditionUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyIntentResponseRequestDetailsBodyRuleConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "tron_trigger_smart_contract_data",
	// "sui_transaction_command", "sui_transfer_objects_command", "reference".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion],
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion], [TronCalldataConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion], [AggregationConditionValueUnion]
	Value PolicyIntentResponseRequestDetailsBodyRuleConditionUnionValue `json:"value"`
	Abi   any                                                           `json:"abi"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage].
	TypedData PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
	JSON      struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		Abi         respjson.Field
		TypedData   respjson.Field
		raw         string
	} `json:"-"`
}

// anyPolicyIntentResponseRequestDetailsBodyRuleCondition is implemented by each
// variant of [PolicyIntentResponseRequestDetailsBodyRuleConditionUnion] to add
// type safety for the return type of
// [PolicyIntentResponseRequestDetailsBodyRuleConditionUnion.AsAny]
type anyPolicyIntentResponseRequestDetailsBodyRuleCondition interface {
	implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion()
}

func (PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (PolicyIntentResponseRequestDetailsBodyRuleConditionSystem) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (TronTransactionCondition) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {}
func (TronCalldataCondition) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion()    {}
func (SuiTransactionCommandCondition) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (SuiTransferObjectsCommandCondition) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {
}
func (AggregationCondition) implPolicyIntentResponseRequestDetailsBodyRuleConditionUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyIntentResponseRequestDetailsBodyRuleConditionUnion.AsAny().(type) {
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction:
//	case privyclient.PolicyIntentResponseRequestDetailsBodyRuleConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.TronCalldataCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	case privyclient.AggregationCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsAny() anyPolicyIntentResponseRequestDetailsBodyRuleCondition {
	switch u.FieldSource {
	case "ethereum_transaction":
		return u.AsEthereumTransaction()
	case "ethereum_calldata":
		return u.AsEthereumCalldata()
	case "ethereum_typed_data_domain":
		return u.AsEthereumTypedDataDomain()
	case "ethereum_typed_data_message":
		return u.AsEthereumTypedDataMessage()
	case "ethereum_7702_authorization":
		return u.AsEthereum7702Authorization()
	case "solana_program_instruction":
		return u.AsSolanaProgramInstruction()
	case "solana_system_program_instruction":
		return u.AsSolanaSystemProgramInstruction()
	case "solana_token_program_instruction":
		return u.AsSolanaTokenProgramInstruction()
	case "system":
		return u.AsSystem()
	case "tron_transaction":
		return u.AsTronTransaction()
	case "tron_trigger_smart_contract_data":
		return u.AsTronTriggerSmartContractData()
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	case "reference":
		return u.AsReference()
	}
	return nil
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsEthereumTransaction() (v PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsEthereumCalldata() (v PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsEthereumTypedDataDomain() (v PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsEthereumTypedDataMessage() (v PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsEthereum7702Authorization() (v PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsSolanaProgramInstruction() (v PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsSolanaSystemProgramInstruction() (v PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsSolanaTokenProgramInstruction() (v PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsSystem() (v PolicyIntentResponseRequestDetailsBodyRuleConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsTronTriggerSmartContractData() (v TronCalldataCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsReference() (v AggregationCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionUnionValue is an implicit
// subunion of [PolicyIntentResponseRequestDetailsBodyRuleConditionUnion].
// PolicyIntentResponseRequestDetailsBodyRuleConditionUnionValue provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyIntentResponseRequestDetailsBodyRuleConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionUnionValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a [SuiCommandName] instead of an
	// object.
	OfSuiCommandName SuiCommandName `json:",inline"`
	// This field will be present if the value is a [[]SuiCommandName] instead of an
	// object.
	OfSuiCommandNameArray []SuiCommandName `json:",inline"`
	JSON                  struct {
		OfString              respjson.Field
		OfStringArray         respjson.Field
		OfSuiCommandName      respjson.Field
		OfSuiCommandNameArray respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field" api:"required"`
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                           `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata struct {
	Abi         any                       `json:"abi" api:"required"`
	Field       string                    `json:"field" api:"required"`
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                        `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abi         respjson.Field
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field       string                           `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                               `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                                                `json:"operator" api:"required"`
	TypedData PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageTypedData  `json:"typed_data" api:"required"`
	Value     PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		TypedData   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParamsResp `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field" api:"required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                                 `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                                `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field" api:"required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                                      `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field       string                                 `json:"field" api:"required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                                     `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type PolicyIntentResponseRequestDetailsBodyRuleConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field" api:"required"`
	FieldSource constant.System `json:"field_source" default:"system"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                              `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionSystem) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a key quorum intent
type KeyQuorumIntentResponse struct {
	// Any of "KEY_QUORUM".
	IntentType string `json:"intent_type" api:"required"`
	// The original key quorum update request that would be sent to the key quorum
	// endpoint
	RequestDetails KeyQuorumIntentResponseRequestDetails `json:"request_details" api:"required"`
	// Result of key quorum update execution (only present if status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// A key quorum for authorizing wallet operations.
	CurrentResourceData KeyQuorum `json:"current_resource_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntentType          respjson.Field
		RequestDetails      respjson.Field
		ActionResult        respjson.Field
		CurrentResourceData respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
	BaseIntentResponse
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The original key quorum update request that would be sent to the key quorum
// endpoint
type KeyQuorumIntentResponseRequestDetails struct {
	// Request input for updating an existing key quorum.
	Body KeyQuorumUpdateRequestBody `json:"body" api:"required"`
	// Any of "PATCH".
	Method string `json:"method" api:"required"`
	URL    string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumIntentResponseRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumIntentResponseRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a rule intent
type RuleIntentResponse struct {
	// Any of "RULE".
	IntentType string `json:"intent_type" api:"required"`
	// The original rule request. Method is POST (create), PATCH (update), or DELETE
	// (delete)
	RequestDetails RuleIntentRequestDetailsUnion `json:"request_details" api:"required"`
	// Result of rule execution (only present if status is 'executed' or 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// Current state of the rule before any changes. Undefined for create intents or if
	// the rule was deleted
	CurrentResourceData RuleIntentResponseCurrentResourceData `json:"current_resource_data"`
	// A policy for controlling wallet operations.
	Policy Policy `json:"policy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntentType          respjson.Field
		RequestDetails      respjson.Field
		ActionResult        respjson.Field
		CurrentResourceData respjson.Field
		Policy              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
	BaseIntentResponse
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the rule before any changes. Undefined for create intents or if
// the rule was deleted
type RuleIntentResponseCurrentResourceData struct {
	ID string `json:"id" api:"required"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                                                `json:"action" api:"required"`
	Conditions []RuleIntentResponseCurrentResourceDataConditionUnion `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Conditions  respjson.Field
		Method      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceData) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponseCurrentResourceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionUnion contains all possible
// properties and values from
// [RuleIntentResponseCurrentResourceDataConditionEthereumTransaction],
// [RuleIntentResponseCurrentResourceDataConditionEthereumCalldata],
// [RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain],
// [RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage],
// [RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization],
// [RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction],
// [RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction],
// [RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction],
// [RuleIntentResponseCurrentResourceDataConditionSystem],
// [TronTransactionCondition], [TronCalldataCondition],
// [SuiTransactionCommandCondition], [SuiTransferObjectsCommandCondition],
// [AggregationCondition].
//
// Use the [RuleIntentResponseCurrentResourceDataConditionUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentResponseCurrentResourceDataConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "tron_trigger_smart_contract_data",
	// "sui_transaction_command", "sui_transfer_objects_command", "reference".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion],
	// [RuleIntentResponseCurrentResourceDataConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion], [TronCalldataConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion], [AggregationConditionValueUnion]
	Value RuleIntentResponseCurrentResourceDataConditionUnionValue `json:"value"`
	Abi   any                                                      `json:"abi"`
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage].
	TypedData RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
	JSON      struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		Abi         respjson.Field
		TypedData   respjson.Field
		raw         string
	} `json:"-"`
}

// anyRuleIntentResponseCurrentResourceDataCondition is implemented by each variant
// of [RuleIntentResponseCurrentResourceDataConditionUnion] to add type safety for
// the return type of [RuleIntentResponseCurrentResourceDataConditionUnion.AsAny]
type anyRuleIntentResponseCurrentResourceDataCondition interface {
	implRuleIntentResponseCurrentResourceDataConditionUnion()
}

func (RuleIntentResponseCurrentResourceDataConditionEthereumTransaction) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionEthereumCalldata) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (RuleIntentResponseCurrentResourceDataConditionSystem) implRuleIntentResponseCurrentResourceDataConditionUnion() {
}
func (TronTransactionCondition) implRuleIntentResponseCurrentResourceDataConditionUnion()           {}
func (TronCalldataCondition) implRuleIntentResponseCurrentResourceDataConditionUnion()              {}
func (SuiTransactionCommandCondition) implRuleIntentResponseCurrentResourceDataConditionUnion()     {}
func (SuiTransferObjectsCommandCondition) implRuleIntentResponseCurrentResourceDataConditionUnion() {}
func (AggregationCondition) implRuleIntentResponseCurrentResourceDataConditionUnion()               {}

// Use the following switch statement to find the correct variant
//
//	switch variant := RuleIntentResponseCurrentResourceDataConditionUnion.AsAny().(type) {
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionEthereumTransaction:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionEthereumCalldata:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction:
//	case privyclient.RuleIntentResponseCurrentResourceDataConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.TronCalldataCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	case privyclient.AggregationCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsAny() anyRuleIntentResponseCurrentResourceDataCondition {
	switch u.FieldSource {
	case "ethereum_transaction":
		return u.AsEthereumTransaction()
	case "ethereum_calldata":
		return u.AsEthereumCalldata()
	case "ethereum_typed_data_domain":
		return u.AsEthereumTypedDataDomain()
	case "ethereum_typed_data_message":
		return u.AsEthereumTypedDataMessage()
	case "ethereum_7702_authorization":
		return u.AsEthereum7702Authorization()
	case "solana_program_instruction":
		return u.AsSolanaProgramInstruction()
	case "solana_system_program_instruction":
		return u.AsSolanaSystemProgramInstruction()
	case "solana_token_program_instruction":
		return u.AsSolanaTokenProgramInstruction()
	case "system":
		return u.AsSystem()
	case "tron_transaction":
		return u.AsTronTransaction()
	case "tron_trigger_smart_contract_data":
		return u.AsTronTriggerSmartContractData()
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	case "reference":
		return u.AsReference()
	}
	return nil
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsEthereumTransaction() (v RuleIntentResponseCurrentResourceDataConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsEthereumCalldata() (v RuleIntentResponseCurrentResourceDataConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsEthereumTypedDataDomain() (v RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsEthereumTypedDataMessage() (v RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsEthereum7702Authorization() (v RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsSolanaProgramInstruction() (v RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsSolanaSystemProgramInstruction() (v RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsSolanaTokenProgramInstruction() (v RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsSystem() (v RuleIntentResponseCurrentResourceDataConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsTronTriggerSmartContractData() (v TronCalldataCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsReference() (v AggregationCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentResponseCurrentResourceDataConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionUnionValue is an implicit subunion
// of [RuleIntentResponseCurrentResourceDataConditionUnion].
// RuleIntentResponseCurrentResourceDataConditionUnionValue provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [RuleIntentResponseCurrentResourceDataConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type RuleIntentResponseCurrentResourceDataConditionUnionValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a [SuiCommandName] instead of an
	// object.
	OfSuiCommandName SuiCommandName `json:",inline"`
	// This field will be present if the value is a [[]SuiCommandName] instead of an
	// object.
	OfSuiCommandNameArray []SuiCommandName `json:",inline"`
	JSON                  struct {
		OfString              respjson.Field
		OfStringArray         respjson.Field
		OfSuiCommandName      respjson.Field
		OfSuiCommandNameArray respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *RuleIntentResponseCurrentResourceDataConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type RuleIntentResponseCurrentResourceDataConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field" api:"required"`
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                      `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionEthereumTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type RuleIntentResponseCurrentResourceDataConditionEthereumCalldata struct {
	Abi         any                       `json:"abi" api:"required"`
	Field       string                    `json:"field" api:"required"`
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                   `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abi         respjson.Field
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionEthereumCalldata) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field       string                           `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                          `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                                           `json:"operator" api:"required"`
	TypedData RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageTypedData  `json:"typed_data" api:"required"`
	Value     RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		TypedData   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParamsResp `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
type RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field" api:"required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                            `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                           `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field" api:"required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                                 `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field       string                                 `json:"field" api:"required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                                `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type RuleIntentResponseCurrentResourceDataConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field" api:"required"`
	FieldSource constant.System `json:"field_source" default:"system"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                         `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionSystemValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		FieldSource respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponseCurrentResourceDataConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponseCurrentResourceDataConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponseCurrentResourceDataConditionSystemValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionSystemValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RuleIntentResponseCurrentResourceDataConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionSystemValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnion contains all possible properties and values from
// [RpcIntentResponse], [TransferIntentResponse], [WalletIntentResponse],
// [PolicyIntentResponse], [RuleIntentResponse], [KeyQuorumIntentResponse].
//
// Use the [IntentResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IntentResponseUnion struct {
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	AuthorizationDetails []IntentAuthorization `json:"authorization_details"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	CreatedAt float64 `json:"created_at"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	CreatedByDisplayName string `json:"created_by_display_name"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	CustomExpiry bool `json:"custom_expiry"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	ExpiresAt float64 `json:"expires_at"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	IntentID string `json:"intent_id"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	ResourceID string `json:"resource_id"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	Status IntentStatus `json:"status"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	CreatedByID string `json:"created_by_id"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	DismissalReason string `json:"dismissal_reason"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	DismissedAt float64 `json:"dismissed_at"`
	// This field is from variant [RpcIntentResponse], [TransferIntentResponse],
	// [WalletIntentResponse], [PolicyIntentResponse], [RuleIntentResponse],
	// [KeyQuorumIntentResponse].
	RejectedAt float64 `json:"rejected_at"`
	// Any of "RPC", "TRANSFER", "WALLET", "POLICY", "RULE", "KEY_QUORUM".
	IntentType string `json:"intent_type"`
	// This field is a union of [RpcIntentResponseRequestDetails],
	// [TransferIntentResponseRequestDetails], [WalletIntentResponseRequestDetails],
	// [PolicyIntentResponseRequestDetails], [RuleIntentRequestDetailsUnion],
	// [KeyQuorumIntentResponseRequestDetails]
	RequestDetails IntentResponseUnionRequestDetails `json:"request_details"`
	// This field is from variant [RpcIntentResponse].
	ActionResult BaseActionResult `json:"action_result"`
	// This field is a union of [Wallet], [Policy],
	// [RuleIntentResponseCurrentResourceData], [KeyQuorum]
	CurrentResourceData IntentResponseUnionCurrentResourceData `json:"current_resource_data"`
	// This field is from variant [RuleIntentResponse].
	Policy Policy `json:"policy"`
	JSON   struct {
		AuthorizationDetails respjson.Field
		CreatedAt            respjson.Field
		CreatedByDisplayName respjson.Field
		CustomExpiry         respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		CreatedByID          respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ActionResult         respjson.Field
		CurrentResourceData  respjson.Field
		Policy               respjson.Field
		raw                  string
	} `json:"-"`
}

// anyIntentResponse is implemented by each variant of [IntentResponseUnion] to add
// type safety for the return type of [IntentResponseUnion.AsAny]
type anyIntentResponse interface {
	implIntentResponseUnion()
}

func (RpcIntentResponse) implIntentResponseUnion()       {}
func (TransferIntentResponse) implIntentResponseUnion()  {}
func (WalletIntentResponse) implIntentResponseUnion()    {}
func (PolicyIntentResponse) implIntentResponseUnion()    {}
func (RuleIntentResponse) implIntentResponseUnion()      {}
func (KeyQuorumIntentResponse) implIntentResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := IntentResponseUnion.AsAny().(type) {
//	case privyclient.RpcIntentResponse:
//	case privyclient.TransferIntentResponse:
//	case privyclient.WalletIntentResponse:
//	case privyclient.PolicyIntentResponse:
//	case privyclient.RuleIntentResponse:
//	case privyclient.KeyQuorumIntentResponse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u IntentResponseUnion) AsAny() anyIntentResponse {
	switch u.IntentType {
	case "RPC":
		return u.AsRpc()
	case "TRANSFER":
		return u.AsTransfer()
	case "WALLET":
		return u.AsWallet()
	case "POLICY":
		return u.AsPolicy()
	case "RULE":
		return u.AsRule()
	case "KEY_QUORUM":
		return u.AsKeyQuorum()
	}
	return nil
}

func (u IntentResponseUnion) AsRpc() (v RpcIntentResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentResponseUnion) AsTransfer() (v TransferIntentResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentResponseUnion) AsWallet() (v WalletIntentResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentResponseUnion) AsPolicy() (v PolicyIntentResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentResponseUnion) AsRule() (v RuleIntentResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntentResponseUnion) AsKeyQuorum() (v KeyQuorumIntentResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IntentResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *IntentResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionRequestDetails is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionRequestDetails provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
type IntentResponseUnionRequestDetails struct {
	// This field is a union of [WalletRpcRequestBodyUnion],
	// [CreateTokenTransferRequest], [WalletIntentResponseRequestDetailsBody],
	// [PolicyIntentResponseRequestDetailsBody], [RuleIntentCreateRequestDetailsBody],
	// [RuleIntentUpdateRequestDetailsBody], [any], [KeyQuorumUpdateRequestBody]
	Body   IntentResponseUnionRequestDetailsBody `json:"body"`
	Method string                                `json:"method"`
	URL    string                                `json:"url"`
	JSON   struct {
		Body   respjson.Field
		Method respjson.Field
		URL    respjson.Field
		raw    string
	} `json:"-"`
}

func (r *IntentResponseUnionRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionRequestDetailsBody is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionRequestDetailsBody provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfRuleIntentDeleteRequestDetailsBody]
type IntentResponseUnionRequestDetailsBody struct {
	// This field will be present if the value is a [any] instead of an object.
	OfRuleIntentDeleteRequestDetailsBody any    `json:",inline"`
	Method                               string `json:"method"`
	// This field is a union of [EthereumSignTransactionRpcInputParamsResp],
	// [EthereumSendTransactionRpcInputParamsResp],
	// [EthereumPersonalSignRpcInputParamsResp],
	// [EthereumSignTypedDataRpcInputParamsResp],
	// [EthereumSecp256k1SignRpcInputParamsResp],
	// [EthereumSign7702AuthorizationRpcInputParamsResp],
	// [EthereumSignUserOperationRpcInputParamsResp],
	// [EthereumSendCallsRpcInputParamsResp],
	// [SolanaSignTransactionRpcInputParamsResp],
	// [SolanaSignAndSendTransactionRpcInputParamsResp],
	// [SolanaSignMessageRpcInputParamsResp], [SparkTransferRpcInputParamsResp],
	// [SparkTransferTokensRpcInputParamsResp],
	// [SparkGetClaimStaticDepositQuoteRpcInputParamsResp],
	// [SparkClaimStaticDepositRpcInputParamsResp],
	// [SparkCreateLightningInvoiceRpcInputParamsResp],
	// [SparkPayLightningInvoiceRpcInputParamsResp],
	// [SparkSignMessageWithIdentityKeyRpcInputParamsResp], [PrivateKeyExportInput]
	Params    IntentResponseUnionRequestDetailsBodyParams `json:"params"`
	Address   string                                      `json:"address"`
	ChainType string                                      `json:"chain_type"`
	WalletID  string                                      `json:"wallet_id"`
	// This field is from variant [WalletRpcRequestBodyUnion].
	Caip2       Caip2  `json:"caip2"`
	ReferenceID string `json:"reference_id"`
	Sponsor     bool   `json:"sponsor"`
	// This field is from variant [WalletRpcRequestBodyUnion].
	Network SparkNetwork `json:"network"`
	// This field is from variant [CreateTokenTransferRequest].
	Destination TokenTransferDestination `json:"destination"`
	// This field is from variant [CreateTokenTransferRequest].
	Source TokenTransferSource `json:"source"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	AdditionalSigners AdditionalSignerInput `json:"additional_signers"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	AuthorizationKeyIDs    []string `json:"authorization_key_ids"`
	AuthorizationThreshold float64  `json:"authorization_threshold"`
	// This field is a union of [WalletIntentResponseRequestDetailsBodyOwner],
	// [PolicyIntentResponseRequestDetailsBodyOwnerUnion]
	Owner   IntentResponseUnionRequestDetailsBodyOwner `json:"owner"`
	OwnerID string                                     `json:"owner_id"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	PolicyIDs PolicyInput `json:"policy_ids"`
	Name      string      `json:"name"`
	// This field is from variant [PolicyIntentResponseRequestDetailsBody].
	Rules  []PolicyIntentResponseRequestDetailsBodyRule `json:"rules"`
	Action string                                       `json:"action"`
	// This field is a union of [[]RuleIntentCreateRequestDetailsBodyConditionUnion],
	// [[]RuleIntentUpdateRequestDetailsBodyConditionUnion]
	Conditions IntentResponseUnionRequestDetailsBodyConditions `json:"conditions"`
	// This field is from variant [KeyQuorumUpdateRequestBody].
	DisplayName string `json:"display_name"`
	// This field is from variant [KeyQuorumUpdateRequestBody].
	KeyQuorumIDs []string `json:"key_quorum_ids"`
	// This field is from variant [KeyQuorumUpdateRequestBody].
	PublicKeys []string `json:"public_keys"`
	// This field is from variant [KeyQuorumUpdateRequestBody].
	UserIDs []string `json:"user_ids"`
	JSON    struct {
		OfRuleIntentDeleteRequestDetailsBody respjson.Field
		Method                               respjson.Field
		Params                               respjson.Field
		Address                              respjson.Field
		ChainType                            respjson.Field
		WalletID                             respjson.Field
		Caip2                                respjson.Field
		ReferenceID                          respjson.Field
		Sponsor                              respjson.Field
		Network                              respjson.Field
		Destination                          respjson.Field
		Source                               respjson.Field
		AdditionalSigners                    respjson.Field
		AuthorizationKeyIDs                  respjson.Field
		AuthorizationThreshold               respjson.Field
		Owner                                respjson.Field
		OwnerID                              respjson.Field
		PolicyIDs                            respjson.Field
		Name                                 respjson.Field
		Rules                                respjson.Field
		Action                               respjson.Field
		Conditions                           respjson.Field
		DisplayName                          respjson.Field
		KeyQuorumIDs                         respjson.Field
		PublicKeys                           respjson.Field
		UserIDs                              respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *IntentResponseUnionRequestDetailsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionRequestDetailsBodyParams is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionRequestDetailsBodyParams provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
type IntentResponseUnionRequestDetailsBodyParams struct {
	// This field is a union of [UnsignedEthereumTransaction], [string], [string]
	Transaction IntentResponseUnionRequestDetailsBodyParamsTransaction `json:"transaction"`
	Encoding    string                                                 `json:"encoding"`
	Message     string                                                 `json:"message"`
	// This field is from variant [EthereumSignTypedDataRpcInputParamsResp].
	TypedData EthereumTypedDataInput `json:"typed_data"`
	// This field is from variant [EthereumSecp256k1SignRpcInputParamsResp].
	Hash Hex `json:"hash"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	ChainID  QuantityUnion `json:"chain_id"`
	Contract string        `json:"contract"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor `json:"executor"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	Nonce QuantityUnion `json:"nonce"`
	// This field is from variant [EthereumSignUserOperationRpcInputParamsResp].
	UserOperation UserOperationInput `json:"user_operation"`
	// This field is from variant [EthereumSendCallsRpcInputParamsResp].
	Calls                []EthereumSendCallsCall `json:"calls"`
	AmountSats           float64                 `json:"amount_sats"`
	ReceiverSparkAddress string                  `json:"receiver_spark_address"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	TokenAmount float64 `json:"token_amount"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	TokenIdentifier string `json:"token_identifier"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	OutputSelectionStrategy SparkOutputSelectionStrategy `json:"output_selection_strategy"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	SelectedOutputs []OutputWithPreviousTransactionData `json:"selected_outputs"`
	TransactionID   string                              `json:"transaction_id"`
	OutputIndex     float64                             `json:"output_index"`
	// This field is from variant [SparkClaimStaticDepositRpcInputParamsResp].
	CreditAmountSats float64 `json:"credit_amount_sats"`
	// This field is from variant [SparkClaimStaticDepositRpcInputParamsResp].
	Signature string `json:"signature"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	DescriptionHash string `json:"description_hash"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	ExpirySeconds float64 `json:"expiry_seconds"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	IncludeSparkAddress bool `json:"include_spark_address"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	Memo string `json:"memo"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	ReceiverIdentityPubkey string `json:"receiver_identity_pubkey"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	Invoice string `json:"invoice"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	MaxFeeSats float64 `json:"max_fee_sats"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	AmountSatsToSend float64 `json:"amount_sats_to_send"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	PreferSpark bool `json:"prefer_spark"`
	// This field is from variant [SparkSignMessageWithIdentityKeyRpcInputParamsResp].
	Compact bool `json:"compact"`
	// This field is from variant [PrivateKeyExportInput].
	EncryptionType HpkeEncryption `json:"encryption_type"`
	// This field is from variant [PrivateKeyExportInput].
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key"`
	// This field is from variant [PrivateKeyExportInput].
	ExportType ExportType `json:"export_type"`
	JSON       struct {
		Transaction             respjson.Field
		Encoding                respjson.Field
		Message                 respjson.Field
		TypedData               respjson.Field
		Hash                    respjson.Field
		ChainID                 respjson.Field
		Contract                respjson.Field
		Executor                respjson.Field
		Nonce                   respjson.Field
		UserOperation           respjson.Field
		Calls                   respjson.Field
		AmountSats              respjson.Field
		ReceiverSparkAddress    respjson.Field
		TokenAmount             respjson.Field
		TokenIdentifier         respjson.Field
		OutputSelectionStrategy respjson.Field
		SelectedOutputs         respjson.Field
		TransactionID           respjson.Field
		OutputIndex             respjson.Field
		CreditAmountSats        respjson.Field
		Signature               respjson.Field
		DescriptionHash         respjson.Field
		ExpirySeconds           respjson.Field
		IncludeSparkAddress     respjson.Field
		Memo                    respjson.Field
		ReceiverIdentityPubkey  respjson.Field
		Invoice                 respjson.Field
		MaxFeeSats              respjson.Field
		AmountSatsToSend        respjson.Field
		PreferSpark             respjson.Field
		Compact                 respjson.Field
		EncryptionType          respjson.Field
		RecipientPublicKey      respjson.Field
		ExportType              respjson.Field
		raw                     string
	} `json:"-"`
}

func (r *IntentResponseUnionRequestDetailsBodyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionRequestDetailsBodyParamsTransaction is an implicit subunion
// of [IntentResponseUnion]. IntentResponseUnionRequestDetailsBodyParamsTransaction
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type IntentResponseUnionRequestDetailsBodyParamsTransaction struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [UnsignedEthereumTransaction].
	AuthorizationList []EthereumSign7702Authorization `json:"authorization_list"`
	// This field is from variant [UnsignedEthereumTransaction].
	ChainID QuantityUnion `json:"chain_id"`
	// This field is from variant [UnsignedEthereumTransaction].
	Data Hex `json:"data"`
	// This field is from variant [UnsignedEthereumTransaction].
	From string `json:"from"`
	// This field is from variant [UnsignedEthereumTransaction].
	GasLimit QuantityUnion `json:"gas_limit"`
	// This field is from variant [UnsignedEthereumTransaction].
	GasPrice QuantityUnion `json:"gas_price"`
	// This field is from variant [UnsignedEthereumTransaction].
	MaxFeePerGas QuantityUnion `json:"max_fee_per_gas"`
	// This field is from variant [UnsignedEthereumTransaction].
	MaxPriorityFeePerGas QuantityUnion `json:"max_priority_fee_per_gas"`
	// This field is from variant [UnsignedEthereumTransaction].
	Nonce QuantityUnion `json:"nonce"`
	// This field is from variant [UnsignedEthereumTransaction].
	To string `json:"to"`
	// This field is from variant [UnsignedEthereumTransaction].
	Type float64 `json:"type"`
	// This field is from variant [UnsignedEthereumTransaction].
	Value QuantityUnion `json:"value"`
	JSON  struct {
		OfString             respjson.Field
		AuthorizationList    respjson.Field
		ChainID              respjson.Field
		Data                 respjson.Field
		From                 respjson.Field
		GasLimit             respjson.Field
		GasPrice             respjson.Field
		MaxFeePerGas         respjson.Field
		MaxPriorityFeePerGas respjson.Field
		Nonce                respjson.Field
		To                   respjson.Field
		Type                 respjson.Field
		Value                respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *IntentResponseUnionRequestDetailsBodyParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionRequestDetailsBodyOwner is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionRequestDetailsBodyOwner provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
type IntentResponseUnionRequestDetailsBodyOwner struct {
	PublicKey string `json:"public_key"`
	UserID    string `json:"user_id"`
	JSON      struct {
		PublicKey respjson.Field
		UserID    respjson.Field
		raw       string
	} `json:"-"`
}

func (r *IntentResponseUnionRequestDetailsBodyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionRequestDetailsBodyConditions is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionRequestDetailsBodyConditions provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfRuleIntentCreateRequestDetailsBodyConditions
// OfRuleIntentUpdateRequestDetailsBodyConditions]
type IntentResponseUnionRequestDetailsBodyConditions struct {
	// This field will be present if the value is a
	// [[]RuleIntentCreateRequestDetailsBodyConditionUnion] instead of an object.
	OfRuleIntentCreateRequestDetailsBodyConditions []RuleIntentCreateRequestDetailsBodyConditionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]RuleIntentUpdateRequestDetailsBodyConditionUnion] instead of an object.
	OfRuleIntentUpdateRequestDetailsBodyConditions []RuleIntentUpdateRequestDetailsBodyConditionUnion `json:",inline"`
	JSON                                           struct {
		OfRuleIntentCreateRequestDetailsBodyConditions respjson.Field
		OfRuleIntentUpdateRequestDetailsBodyConditions respjson.Field
		raw                                            string
	} `json:"-"`
}

func (r *IntentResponseUnionRequestDetailsBodyConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IntentResponseUnionCurrentResourceData is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionCurrentResourceData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
type IntentResponseUnionCurrentResourceData struct {
	ID string `json:"id"`
	// This field is from variant [Wallet].
	AdditionalSigners []WalletAdditionalSigner `json:"additional_signers"`
	// This field is from variant [Wallet].
	Address string `json:"address"`
	// This field is from variant [Wallet].
	ChainType WalletChainType `json:"chain_type"`
	CreatedAt float64         `json:"created_at"`
	// This field is from variant [Wallet].
	ExportedAt float64 `json:"exported_at"`
	// This field is from variant [Wallet].
	ImportedAt float64 `json:"imported_at"`
	OwnerID    string  `json:"owner_id"`
	// This field is from variant [Wallet].
	PolicyIDs              []string `json:"policy_ids"`
	AuthorizationThreshold float64  `json:"authorization_threshold"`
	// This field is from variant [Wallet].
	Custody WalletCustodian `json:"custody"`
	// This field is from variant [Wallet].
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
	// This field is from variant [Policy].
	Rules []PolicyRule `json:"rules"`
	// This field is from variant [Policy].
	Version PolicyVersion `json:"version"`
	// This field is from variant [RuleIntentResponseCurrentResourceData].
	Action string `json:"action"`
	// This field is from variant [RuleIntentResponseCurrentResourceData].
	Conditions []RuleIntentResponseCurrentResourceDataConditionUnion `json:"conditions"`
	// This field is from variant [RuleIntentResponseCurrentResourceData].
	Method string `json:"method"`
	// This field is from variant [KeyQuorum].
	AuthorizationKeys []KeyQuorumAuthorizationKey `json:"authorization_keys"`
	// This field is from variant [KeyQuorum].
	DisplayName string `json:"display_name"`
	// This field is from variant [KeyQuorum].
	UserIDs []string `json:"user_ids"`
	// This field is from variant [KeyQuorum].
	KeyQuorumIDs []string `json:"key_quorum_ids"`
	JSON         struct {
		ID                     respjson.Field
		AdditionalSigners      respjson.Field
		Address                respjson.Field
		ChainType              respjson.Field
		CreatedAt              respjson.Field
		ExportedAt             respjson.Field
		ImportedAt             respjson.Field
		OwnerID                respjson.Field
		PolicyIDs              respjson.Field
		AuthorizationThreshold respjson.Field
		Custody                respjson.Field
		PublicKey              respjson.Field
		Name                   respjson.Field
		Rules                  respjson.Field
		Version                respjson.Field
		Action                 respjson.Field
		Conditions             respjson.Field
		Method                 respjson.Field
		AuthorizationKeys      respjson.Field
		DisplayName            respjson.Field
		UserIDs                respjson.Field
		KeyQuorumIDs           respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *IntentResponseUnionCurrentResourceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentListParams struct {
	Limit           param.Opt[float64] `query:"limit,omitzero" json:"-"`
	CreatedByID     param.Opt[string]  `query:"created_by_id,omitzero" json:"-"`
	Cursor          param.Opt[string]  `query:"cursor,omitzero" json:"-"`
	PendingMemberID param.Opt[string]  `query:"pending_member_id,omitzero" json:"-"`
	ResourceID      param.Opt[string]  `query:"resource_id,omitzero" json:"-"`
	// Any of "true", "false".
	CurrentUserHasSigned IntentListParamsCurrentUserHasSigned `query:"current_user_has_signed,omitzero" json:"-"`
	// Type of intent.
	//
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "TRANSFER", "WALLET".
	IntentType IntentType `query:"intent_type,omitzero" json:"-"`
	// Any of "created_at_desc", "expires_at_asc", "updated_at_desc".
	SortBy IntentListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntentListParams]'s query parameters as `url.Values`.
func (r IntentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntentListParamsCurrentUserHasSigned string

const (
	IntentListParamsCurrentUserHasSignedTrue  IntentListParamsCurrentUserHasSigned = "true"
	IntentListParamsCurrentUserHasSignedFalse IntentListParamsCurrentUserHasSigned = "false"
)

type IntentListParamsSortBy string

const (
	IntentListParamsSortByCreatedAtDesc IntentListParamsSortBy = "created_at_desc"
	IntentListParamsSortByExpiresAtAsc  IntentListParamsSortBy = "expires_at_asc"
	IntentListParamsSortByUpdatedAtDesc IntentListParamsSortBy = "updated_at_desc"
)

type IntentNewPolicyRuleParams struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     IntentNewPolicyRuleParamsAction           `json:"action,omitzero" api:"required"`
	Conditions []IntentNewPolicyRuleParamsConditionUnion `json:"conditions,omitzero" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method IntentNewPolicyRuleParamsMethod `json:"method,omitzero" api:"required"`
	Name   string                          `json:"name" api:"required"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentNewPolicyRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type IntentNewPolicyRuleParamsAction string

const (
	IntentNewPolicyRuleParamsActionAllow IntentNewPolicyRuleParamsAction = "ALLOW"
	IntentNewPolicyRuleParamsActionDeny  IntentNewPolicyRuleParamsAction = "DENY"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionUnion struct {
	OfEthereumTransaction            *IntentNewPolicyRuleParamsConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *IntentNewPolicyRuleParamsConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *IntentNewPolicyRuleParamsConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *IntentNewPolicyRuleParamsConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *IntentNewPolicyRuleParamsConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                                    `json:",omitzero,inline"`
	OfTronTriggerSmartContractData   *TronCalldataConditionParam                                       `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                              `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                          `json:",omitzero,inline"`
	OfReference                      *AggregationConditionParam                                        `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEthereumTransaction,
		u.OfEthereumCalldata,
		u.OfEthereumTypedDataDomain,
		u.OfEthereumTypedDataMessage,
		u.OfEthereum7702Authorization,
		u.OfSolanaProgramInstruction,
		u.OfSolanaSystemProgramInstruction,
		u.OfSolanaTokenProgramInstruction,
		u.OfSystem,
		u.OfTronTransaction,
		u.OfTronTriggerSmartContractData,
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand,
		u.OfReference)
}
func (u *IntentNewPolicyRuleParamsConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[IntentNewPolicyRuleParamsConditionUnion](
		"field_source",
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[IntentNewPolicyRuleParamsConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[TronCalldataConditionParam]("tron_trigger_smart_contract_data"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
		apijson.Discriminator[AggregationConditionParam]("reference"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionEthereumTransactionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero" api:"required"`
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                       `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionEthereumCalldataValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                              `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract", "chain_id", "verifying_contract",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage struct {
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                               `json:"operator,omitzero" api:"required"`
	TypedData IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero" api:"required"`
	Value     IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParams `json:"types,omitzero" api:"required"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                               `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionSolanaProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                     `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                    `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "Transfer.source", "Transfer.destination", "Transfer.authority", "Transfer.amount", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint", "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account", "MintTo.authority", "MintTo.amount", "CloseAccount.account", "CloseAccount.destination", "CloseAccount.authority", "InitializeAccount3.account", "InitializeAccount3.mint", "InitializeAccount3.owner",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentNewPolicyRuleParamsConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                             `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionSystemValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source" default:"system"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[IntentNewPolicyRuleParamsConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentNewPolicyRuleParamsConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentNewPolicyRuleParamsConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentNewPolicyRuleParamsConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Method the rule applies to.
type IntentNewPolicyRuleParamsMethod string

const (
	IntentNewPolicyRuleParamsMethodEthSendTransaction       IntentNewPolicyRuleParamsMethod = "eth_sendTransaction"
	IntentNewPolicyRuleParamsMethodEthSignTransaction       IntentNewPolicyRuleParamsMethod = "eth_signTransaction"
	IntentNewPolicyRuleParamsMethodEthSignUserOperation     IntentNewPolicyRuleParamsMethod = "eth_signUserOperation"
	IntentNewPolicyRuleParamsMethodEthSignTypedDataV4       IntentNewPolicyRuleParamsMethod = "eth_signTypedData_v4"
	IntentNewPolicyRuleParamsMethodEthSign7702Authorization IntentNewPolicyRuleParamsMethod = "eth_sign7702Authorization"
	IntentNewPolicyRuleParamsMethodSignTransaction          IntentNewPolicyRuleParamsMethod = "signTransaction"
	IntentNewPolicyRuleParamsMethodSignAndSendTransaction   IntentNewPolicyRuleParamsMethod = "signAndSendTransaction"
	IntentNewPolicyRuleParamsMethodExportPrivateKey         IntentNewPolicyRuleParamsMethod = "exportPrivateKey"
	IntentNewPolicyRuleParamsMethodSignTransactionBytes     IntentNewPolicyRuleParamsMethod = "signTransactionBytes"
	IntentNewPolicyRuleParamsMethodStar                     IntentNewPolicyRuleParamsMethod = "*"
)

type IntentDeletePolicyRuleParams struct {
	// ID of the policy.
	PolicyID string `path:"policy_id" api:"required" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

type IntentRpcParams struct {
	// Request body for wallet RPC operations, discriminated by method.
	WalletRpcRequestBody WalletRpcRequestBodyUnionParam
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentRpcParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletRpcRequestBody)
}
func (r *IntentRpcParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentUpdateKeyQuorumParams struct {
	// Request input for updating an existing key quorum.
	KeyQuorumUpdateRequestBody KeyQuorumUpdateRequestBodyParam
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentUpdateKeyQuorumParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KeyQuorumUpdateRequestBody)
}
func (r *IntentUpdateKeyQuorumParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentUpdatePolicyParams struct {
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// Name to assign to policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner IntentUpdatePolicyParamsOwnerUnion `json:"owner,omitzero"`
	Rules []IntentUpdatePolicyParamsRule     `json:"rules,omitzero"`
	paramObj
}

func (r IntentUpdatePolicyParams) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsOwnerUnion struct {
	OfPublicKeyOwner *IntentUpdatePolicyParamsOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *IntentUpdatePolicyParamsOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *IntentUpdatePolicyParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type IntentUpdatePolicyParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyParamsOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type IntentUpdatePolicyParamsOwnerUserOwner struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyParamsOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rules that apply to each method the policy covers.
//
// The properties Action, Conditions, Method, Name are required.
type IntentUpdatePolicyParamsRule struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                                       `json:"action,omitzero" api:"required"`
	Conditions []IntentUpdatePolicyParamsRuleConditionUnion `json:"conditions,omitzero" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method,omitzero" api:"required"`
	Name   string `json:"name" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRule](
		"action", "ALLOW", "DENY",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRule](
		"method", "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation", "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction", "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "*",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionUnion struct {
	OfEthereumTransaction            *IntentUpdatePolicyParamsRuleConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *IntentUpdatePolicyParamsRuleConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *IntentUpdatePolicyParamsRuleConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                                       `json:",omitzero,inline"`
	OfTronTriggerSmartContractData   *TronCalldataConditionParam                                          `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                                 `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                             `json:",omitzero,inline"`
	OfReference                      *AggregationConditionParam                                           `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEthereumTransaction,
		u.OfEthereumCalldata,
		u.OfEthereumTypedDataDomain,
		u.OfEthereumTypedDataMessage,
		u.OfEthereum7702Authorization,
		u.OfSolanaProgramInstruction,
		u.OfSolanaSystemProgramInstruction,
		u.OfSolanaTokenProgramInstruction,
		u.OfSystem,
		u.OfTronTransaction,
		u.OfTronTriggerSmartContractData,
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand,
		u.OfReference)
}
func (u *IntentUpdatePolicyParamsRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[IntentUpdatePolicyParamsRuleConditionUnion](
		"field_source",
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[IntentUpdatePolicyParamsRuleConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[TronCalldataConditionParam]("tron_trigger_smart_contract_data"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
		apijson.Discriminator[AggregationConditionParam]("reference"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                             `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionEthereumTransactionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero" api:"required"`
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionEthereumCalldataValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract", "chain_id", "verifying_contract",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage struct {
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                                  `json:"operator,omitzero" api:"required"`
	TypedData IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero" api:"required"`
	Value     IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParams `json:"types,omitzero" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                   `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                  `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionSolanaProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                        `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                       `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "Transfer.source", "Transfer.destination", "Transfer.authority", "Transfer.amount", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint", "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account", "MintTo.authority", "MintTo.amount", "CloseAccount.account", "CloseAccount.destination", "CloseAccount.authority", "InitializeAccount3.account", "InitializeAccount3.mint", "InitializeAccount3.owner",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyParamsRuleConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionSystemValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source" default:"system"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyParamsRuleConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyParamsRuleConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyParamsRuleConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyParamsRuleConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type IntentUpdatePolicyRuleParams struct {
	// ID of the policy.
	PolicyID string `path:"policy_id" api:"required" json:"-"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     IntentUpdatePolicyRuleParamsAction           `json:"action,omitzero" api:"required"`
	Conditions []IntentUpdatePolicyRuleParamsConditionUnion `json:"conditions,omitzero" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method IntentUpdatePolicyRuleParamsMethod `json:"method,omitzero" api:"required"`
	Name   string                             `json:"name" api:"required"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentUpdatePolicyRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type IntentUpdatePolicyRuleParamsAction string

const (
	IntentUpdatePolicyRuleParamsActionAllow IntentUpdatePolicyRuleParamsAction = "ALLOW"
	IntentUpdatePolicyRuleParamsActionDeny  IntentUpdatePolicyRuleParamsAction = "DENY"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionUnion struct {
	OfEthereumTransaction            *IntentUpdatePolicyRuleParamsConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *IntentUpdatePolicyRuleParamsConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *IntentUpdatePolicyRuleParamsConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                                       `json:",omitzero,inline"`
	OfTronTriggerSmartContractData   *TronCalldataConditionParam                                          `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                                 `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                             `json:",omitzero,inline"`
	OfReference                      *AggregationConditionParam                                           `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEthereumTransaction,
		u.OfEthereumCalldata,
		u.OfEthereumTypedDataDomain,
		u.OfEthereumTypedDataMessage,
		u.OfEthereum7702Authorization,
		u.OfSolanaProgramInstruction,
		u.OfSolanaSystemProgramInstruction,
		u.OfSolanaTokenProgramInstruction,
		u.OfSystem,
		u.OfTronTransaction,
		u.OfTronTriggerSmartContractData,
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand,
		u.OfReference)
}
func (u *IntentUpdatePolicyRuleParamsConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[IntentUpdatePolicyRuleParamsConditionUnion](
		"field_source",
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[IntentUpdatePolicyRuleParamsConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[TronCalldataConditionParam]("tron_trigger_smart_contract_data"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
		apijson.Discriminator[AggregationConditionParam]("reference"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                             `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionEthereumTransactionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero" api:"required"`
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionEthereumCalldataValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract", "chain_id", "verifying_contract",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage struct {
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                                  `json:"operator,omitzero" api:"required"`
	TypedData IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero" api:"required"`
	Value     IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParams `json:"types,omitzero" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                   `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                  `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionSolanaProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                        `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                       `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "Transfer.source", "Transfer.destination", "Transfer.authority", "Transfer.amount", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint", "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account", "MintTo.authority", "MintTo.amount", "CloseAccount.account", "CloseAccount.destination", "CloseAccount.authority", "InitializeAccount3.account", "InitializeAccount3.mint", "InitializeAccount3.owner",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type IntentUpdatePolicyRuleParamsConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionSystemValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source" default:"system"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[IntentUpdatePolicyRuleParamsConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type IntentUpdatePolicyRuleParamsConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u IntentUpdatePolicyRuleParamsConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *IntentUpdatePolicyRuleParamsConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Method the rule applies to.
type IntentUpdatePolicyRuleParamsMethod string

const (
	IntentUpdatePolicyRuleParamsMethodEthSendTransaction       IntentUpdatePolicyRuleParamsMethod = "eth_sendTransaction"
	IntentUpdatePolicyRuleParamsMethodEthSignTransaction       IntentUpdatePolicyRuleParamsMethod = "eth_signTransaction"
	IntentUpdatePolicyRuleParamsMethodEthSignUserOperation     IntentUpdatePolicyRuleParamsMethod = "eth_signUserOperation"
	IntentUpdatePolicyRuleParamsMethodEthSignTypedDataV4       IntentUpdatePolicyRuleParamsMethod = "eth_signTypedData_v4"
	IntentUpdatePolicyRuleParamsMethodEthSign7702Authorization IntentUpdatePolicyRuleParamsMethod = "eth_sign7702Authorization"
	IntentUpdatePolicyRuleParamsMethodSignTransaction          IntentUpdatePolicyRuleParamsMethod = "signTransaction"
	IntentUpdatePolicyRuleParamsMethodSignAndSendTransaction   IntentUpdatePolicyRuleParamsMethod = "signAndSendTransaction"
	IntentUpdatePolicyRuleParamsMethodExportPrivateKey         IntentUpdatePolicyRuleParamsMethod = "exportPrivateKey"
	IntentUpdatePolicyRuleParamsMethodSignTransactionBytes     IntentUpdatePolicyRuleParamsMethod = "signTransactionBytes"
	IntentUpdatePolicyRuleParamsMethodStar                     IntentUpdatePolicyRuleParamsMethod = "*"
)

type IntentUpdateWalletParams struct {
	// Request body for updating a wallet.
	WalletUpdateRequestBody WalletUpdateRequestBody
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentUpdateWalletParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletUpdateRequestBody)
}
func (r *IntentUpdateWalletParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
