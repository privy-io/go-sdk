// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
	"github.com/privy-io/go-sdk/shared/constant"
)

// PolicyService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r PolicyService) {
	r = PolicyService{}
	r.Options = opts
	return
}

// Create a new policy.
func (r *PolicyService) New(ctx context.Context, params PolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%s", params.PrivyIdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a policy by policy ID.
func (r *PolicyService) Update(ctx context.Context, policyID string, params PolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s", policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Delete a policy by policy ID.
func (r *PolicyService) Delete(ctx context.Context, policyID string, body PolicyDeleteParams, opts ...option.RequestOption) (res *PolicyDeleteResponse, err error) {
	if !param.IsOmitted(body.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", body.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s", policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new rule for a policy.
func (r *PolicyService) NewRule(ctx context.Context, policyID string, params PolicyNewRuleParams, opts ...option.RequestOption) (res *PolicyNewRuleResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s/rules", policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a rule by policy ID and rule ID.
func (r *PolicyService) DeleteRule(ctx context.Context, ruleID string, params PolicyDeleteRuleParams, opts ...option.RequestOption) (res *PolicyDeleteRuleResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s/rules/%s", params.PolicyID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a policy by policy ID.
func (r *PolicyService) Get(ctx context.Context, policyID string, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s", policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a rule by policy ID and rule ID.
func (r *PolicyService) GetRule(ctx context.Context, ruleID string, query PolicyGetRuleParams, opts ...option.RequestOption) (res *PolicyGetRuleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s/rules/%s", query.PolicyID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a rule by policy ID and rule ID.
func (r *PolicyService) UpdateRule(ctx context.Context, ruleID string, params PolicyUpdateRuleParams, opts ...option.RequestOption) (res *PolicyUpdateRuleResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("v1/policies/%s/rules/%s", params.PolicyID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// A policy for controlling wallet operations.
type Policy struct {
	// Unique ID of the created policy. This will be the primary identifier when using
	// the policy in the future.
	ID string `json:"id,required"`
	// The chain type the policy applies to.
	//
	// Any of "ethereum", "solana", "tron", "sui".
	ChainType PolicyChainType `json:"chain_type,required"`
	// Unix timestamp of when the policy was created in milliseconds.
	CreatedAt float64 `json:"created_at,required"`
	// Name to assign to policy.
	Name string `json:"name,required"`
	// The key quorum ID of the owner of the policy.
	OwnerID string       `json:"owner_id,required" format:"cuid2"`
	Rules   []PolicyRule `json:"rules,required"`
	// Version of the policy. Currently, 1.0 is the only version.
	//
	// Any of "1.0".
	Version PolicyVersion `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ChainType   respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		OwnerID     respjson.Field
		Rules       respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Policy) RawJSON() string { return r.JSON.raw }
func (r *Policy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The chain type the policy applies to.
type PolicyChainType string

const (
	PolicyChainTypeEthereum PolicyChainType = "ethereum"
	PolicyChainTypeSolana   PolicyChainType = "solana"
	PolicyChainTypeTron     PolicyChainType = "tron"
	PolicyChainTypeSui      PolicyChainType = "sui"
)

// A rule that defines the conditions and action to take if the conditions are
// true.
type PolicyRule struct {
	ID string `json:"id,required"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                     `json:"action,required"`
	Conditions []PolicyRuleConditionUnion `json:"conditions,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method,required"`
	Name   string `json:"name,required"`
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
func (r PolicyRule) RawJSON() string { return r.JSON.raw }
func (r *PolicyRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionUnion contains all possible properties and values from
// [PolicyRuleConditionEthereumTransaction], [PolicyRuleConditionEthereumCalldata],
// [PolicyRuleConditionEthereumTypedDataDomain],
// [PolicyRuleConditionEthereumTypedDataMessage],
// [PolicyRuleConditionEthereum7702Authorization],
// [PolicyRuleConditionSolanaProgramInstruction],
// [PolicyRuleConditionSolanaSystemProgramInstruction],
// [PolicyRuleConditionSolanaTokenProgramInstruction], [PolicyRuleConditionSystem],
// [TronTransactionCondition], [SuiTransactionCommandCondition],
// [SuiTransferObjectsCommandCondition].
//
// Use the [PolicyRuleConditionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyRuleConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "sui_transaction_command",
	// "sui_transfer_objects_command".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of [PolicyRuleConditionEthereumTransactionValueUnion],
	// [PolicyRuleConditionEthereumCalldataValueUnion],
	// [PolicyRuleConditionEthereumTypedDataDomainValueUnion],
	// [PolicyRuleConditionEthereumTypedDataMessageValueUnion],
	// [PolicyRuleConditionEthereum7702AuthorizationValueUnion],
	// [PolicyRuleConditionSolanaProgramInstructionValueUnion],
	// [PolicyRuleConditionSolanaSystemProgramInstructionValueUnion],
	// [PolicyRuleConditionSolanaTokenProgramInstructionValueUnion],
	// [PolicyRuleConditionSystemValueUnion], [TronTransactionConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion]
	Value PolicyRuleConditionUnionValue `json:"value"`
	// This field is from variant [PolicyRuleConditionEthereumCalldata].
	Abi any `json:"abi"`
	// This field is from variant [PolicyRuleConditionEthereumTypedDataMessage].
	TypedData PolicyRuleConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
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

// anyPolicyRuleCondition is implemented by each variant of
// [PolicyRuleConditionUnion] to add type safety for the return type of
// [PolicyRuleConditionUnion.AsAny]
type anyPolicyRuleCondition interface {
	implPolicyRuleConditionUnion()
}

func (PolicyRuleConditionEthereumTransaction) implPolicyRuleConditionUnion()            {}
func (PolicyRuleConditionEthereumCalldata) implPolicyRuleConditionUnion()               {}
func (PolicyRuleConditionEthereumTypedDataDomain) implPolicyRuleConditionUnion()        {}
func (PolicyRuleConditionEthereumTypedDataMessage) implPolicyRuleConditionUnion()       {}
func (PolicyRuleConditionEthereum7702Authorization) implPolicyRuleConditionUnion()      {}
func (PolicyRuleConditionSolanaProgramInstruction) implPolicyRuleConditionUnion()       {}
func (PolicyRuleConditionSolanaSystemProgramInstruction) implPolicyRuleConditionUnion() {}
func (PolicyRuleConditionSolanaTokenProgramInstruction) implPolicyRuleConditionUnion()  {}
func (PolicyRuleConditionSystem) implPolicyRuleConditionUnion()                         {}
func (TronTransactionCondition) implPolicyRuleConditionUnion()                          {}
func (SuiTransactionCommandCondition) implPolicyRuleConditionUnion()                    {}
func (SuiTransferObjectsCommandCondition) implPolicyRuleConditionUnion()                {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyRuleConditionUnion.AsAny().(type) {
//	case privyclient.PolicyRuleConditionEthereumTransaction:
//	case privyclient.PolicyRuleConditionEthereumCalldata:
//	case privyclient.PolicyRuleConditionEthereumTypedDataDomain:
//	case privyclient.PolicyRuleConditionEthereumTypedDataMessage:
//	case privyclient.PolicyRuleConditionEthereum7702Authorization:
//	case privyclient.PolicyRuleConditionSolanaProgramInstruction:
//	case privyclient.PolicyRuleConditionSolanaSystemProgramInstruction:
//	case privyclient.PolicyRuleConditionSolanaTokenProgramInstruction:
//	case privyclient.PolicyRuleConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyRuleConditionUnion) AsAny() anyPolicyRuleCondition {
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
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	}
	return nil
}

func (u PolicyRuleConditionUnion) AsEthereumTransaction() (v PolicyRuleConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsEthereumCalldata() (v PolicyRuleConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsEthereumTypedDataDomain() (v PolicyRuleConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsEthereumTypedDataMessage() (v PolicyRuleConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsEthereum7702Authorization() (v PolicyRuleConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsSolanaProgramInstruction() (v PolicyRuleConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsSolanaSystemProgramInstruction() (v PolicyRuleConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsSolanaTokenProgramInstruction() (v PolicyRuleConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsSystem() (v PolicyRuleConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionUnionValue is an implicit subunion of
// [PolicyRuleConditionUnion]. PolicyRuleConditionUnionValue provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyRuleConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyRuleConditionUnionValue struct {
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

func (r *PolicyRuleConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type PolicyRuleConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field,required"`
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                           `json:"operator,required"`
	Value    PolicyRuleConditionEthereumTransactionValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionEthereumTransaction) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionEthereumTransactionValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionEthereumTransactionValueUnion struct {
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

func (u PolicyRuleConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionEthereumTransactionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type PolicyRuleConditionEthereumCalldata struct {
	Abi         any                       `json:"abi,required"`
	Field       string                    `json:"field,required"`
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                        `json:"operator,required"`
	Value    PolicyRuleConditionEthereumCalldataValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionEthereumCalldata) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionEthereumCalldataValueUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionEthereumCalldataValueUnion struct {
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

func (u PolicyRuleConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionEthereumCalldataValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type PolicyRuleConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field       string                           `json:"field,required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                               `json:"operator,required"`
	Value    PolicyRuleConditionEthereumTypedDataDomainValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionEthereumTypedDataDomain) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionEthereumTypedDataDomainValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionEthereumTypedDataDomainValueUnion struct {
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

func (u PolicyRuleConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionEthereumTypedDataDomainValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type PolicyRuleConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field,required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                `json:"operator,required"`
	TypedData PolicyRuleConditionEthereumTypedDataMessageTypedData  `json:"typed_data,required"`
	Value     PolicyRuleConditionEthereumTypedDataMessageValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionEthereumTypedDataMessage) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyRuleConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                `json:"primary_type,required"`
	Types       map[string][]PolicyRuleConditionEthereumTypedDataMessageTypedDataType `json:"types,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyRuleConditionEthereumTypedDataMessageTypedData) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyRuleConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyRuleConditionEthereumTypedDataMessageTypedDataType) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionEthereumTypedDataMessageValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionEthereumTypedDataMessageValueUnion struct {
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

func (u PolicyRuleConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionEthereumTypedDataMessageValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_signAuthorization requests.
type PolicyRuleConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field,required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                 `json:"operator,required"`
	Value    PolicyRuleConditionEthereum7702AuthorizationValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionEthereum7702Authorization) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionEthereum7702AuthorizationValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionEthereum7702AuthorizationValueUnion struct {
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

func (u PolicyRuleConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionEthereum7702AuthorizationValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type PolicyRuleConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field,required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                `json:"operator,required"`
	Value    PolicyRuleConditionSolanaProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionSolanaProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionSolanaProgramInstructionValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionSolanaProgramInstructionValueUnion struct {
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

func (u PolicyRuleConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionSolanaProgramInstructionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type PolicyRuleConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field,required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                      `json:"operator,required"`
	Value    PolicyRuleConditionSolanaSystemProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionSolanaSystemProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionSolanaSystemProgramInstructionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionSolanaSystemProgramInstructionValueUnion struct {
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

func (u PolicyRuleConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyRuleConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type PolicyRuleConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field       string                                 `json:"field,required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                     `json:"operator,required"`
	Value    PolicyRuleConditionSolanaTokenProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionSolanaTokenProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionSolanaTokenProgramInstructionValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionSolanaTokenProgramInstructionValueUnion struct {
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

func (u PolicyRuleConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyRuleConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type PolicyRuleConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field,required"`
	FieldSource constant.System `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                              `json:"operator,required"`
	Value    PolicyRuleConditionSystemValueUnion `json:"value,required"`
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
func (r PolicyRuleConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyRuleConditionSystemValueUnion contains all possible properties and values
// from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyRuleConditionSystemValueUnion struct {
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

func (u PolicyRuleConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyRuleConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyRuleConditionSystemValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyRuleConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version of the policy. Currently, 1.0 is the only version.
type PolicyVersion string

const (
	PolicyVersion1_0 PolicyVersion = "1.0"
)

// Operator to use for SUI transaction command conditions. Only 'eq' and 'in' are
// supported for command names.
type SuiTransactionCommandOperator string

const (
	SuiTransactionCommandOperatorEq SuiTransactionCommandOperator = "eq"
	SuiTransactionCommandOperatorIn SuiTransactionCommandOperator = "in"
)

// Supported fields for SUI TransferObjects command conditions. Only 'recipient'
// and 'amount' are supported.
type SuiTransferObjectsCommandField string

const (
	SuiTransferObjectsCommandFieldRecipient SuiTransferObjectsCommandField = "recipient"
	SuiTransferObjectsCommandFieldAmount    SuiTransferObjectsCommandField = "amount"
)

// TRON transaction fields for TransferContract and TriggerSmartContract
// transaction types.
type TronTransactionCondition struct {
	// Supported TRON transaction fields in format "TransactionType.field_name"
	//
	// Any of "TransferContract.to_address", "TransferContract.amount",
	// "TriggerSmartContract.contract_address", "TriggerSmartContract.call_value",
	// "TriggerSmartContract.token_id", "TriggerSmartContract.call_token_value".
	Field TronTransactionConditionField `json:"field,required"`
	// Any of "tron_transaction".
	FieldSource TronTransactionConditionFieldSource `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator TronTransactionConditionOperator   `json:"operator,required"`
	Value    TronTransactionConditionValueUnion `json:"value,required"`
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
func (r TronTransactionCondition) RawJSON() string { return r.JSON.raw }
func (r *TronTransactionCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TronTransactionCondition to a
// TronTransactionConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TronTransactionConditionParam.Overrides()
func (r TronTransactionCondition) ToParam() TronTransactionConditionParam {
	return param.Override[TronTransactionConditionParam](json.RawMessage(r.RawJSON()))
}

// Supported TRON transaction fields in format "TransactionType.field_name"
type TronTransactionConditionField string

const (
	TronTransactionConditionFieldTransferContractToAddress           TronTransactionConditionField = "TransferContract.to_address"
	TronTransactionConditionFieldTransferContractAmount              TronTransactionConditionField = "TransferContract.amount"
	TronTransactionConditionFieldTriggerSmartContractContractAddress TronTransactionConditionField = "TriggerSmartContract.contract_address"
	TronTransactionConditionFieldTriggerSmartContractCallValue       TronTransactionConditionField = "TriggerSmartContract.call_value"
	TronTransactionConditionFieldTriggerSmartContractTokenID         TronTransactionConditionField = "TriggerSmartContract.token_id"
	TronTransactionConditionFieldTriggerSmartContractCallTokenValue  TronTransactionConditionField = "TriggerSmartContract.call_token_value"
)

type TronTransactionConditionFieldSource string

const (
	TronTransactionConditionFieldSourceTronTransaction TronTransactionConditionFieldSource = "tron_transaction"
)

type TronTransactionConditionOperator string

const (
	TronTransactionConditionOperatorEq             TronTransactionConditionOperator = "eq"
	TronTransactionConditionOperatorGt             TronTransactionConditionOperator = "gt"
	TronTransactionConditionOperatorGte            TronTransactionConditionOperator = "gte"
	TronTransactionConditionOperatorLt             TronTransactionConditionOperator = "lt"
	TronTransactionConditionOperatorLte            TronTransactionConditionOperator = "lte"
	TronTransactionConditionOperatorIn             TronTransactionConditionOperator = "in"
	TronTransactionConditionOperatorInConditionSet TronTransactionConditionOperator = "in_condition_set"
)

// TronTransactionConditionValueUnion contains all possible properties and values
// from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type TronTransactionConditionValueUnion struct {
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

func (u TronTransactionConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TronTransactionConditionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TronTransactionConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TronTransactionConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TRON transaction fields for TransferContract and TriggerSmartContract
// transaction types.
//
// The properties Field, FieldSource, Operator, Value are required.
type TronTransactionConditionParam struct {
	// Supported TRON transaction fields in format "TransactionType.field_name"
	//
	// Any of "TransferContract.to_address", "TransferContract.amount",
	// "TriggerSmartContract.contract_address", "TriggerSmartContract.call_value",
	// "TriggerSmartContract.token_id", "TriggerSmartContract.call_token_value".
	Field TronTransactionConditionField `json:"field,omitzero,required"`
	// Any of "tron_transaction".
	FieldSource TronTransactionConditionFieldSource `json:"field_source,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator TronTransactionConditionOperator        `json:"operator,omitzero,required"`
	Value    TronTransactionConditionValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r TronTransactionConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow TronTransactionConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TronTransactionConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TronTransactionConditionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u TronTransactionConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *TronTransactionConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TronTransactionConditionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// SUI transaction command attributes, enables allowlisting specific command types.
// Allowed commands: 'TransferObjects', 'SplitCoins', 'MergeCoins'. Only 'eq' and
// 'in' operators are supported.
type SuiTransactionCommandCondition struct {
	// Any of "commandName".
	Field SuiTransactionCommandConditionField `json:"field,required"`
	// Any of "sui_transaction_command".
	FieldSource SuiTransactionCommandConditionFieldSource `json:"field_source,required"`
	// Operator to use for SUI transaction command conditions. Only 'eq' and 'in' are
	// supported for command names.
	//
	// Any of "eq", "in".
	Operator SuiTransactionCommandOperator `json:"operator,required"`
	// Command name(s) to match. Must be one of: 'TransferObjects', 'SplitCoins',
	// 'MergeCoins'
	Value SuiTransactionCommandConditionValueUnion `json:"value,required"`
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
func (r SuiTransactionCommandCondition) RawJSON() string { return r.JSON.raw }
func (r *SuiTransactionCommandCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SuiTransactionCommandCondition to a
// SuiTransactionCommandConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SuiTransactionCommandConditionParam.Overrides()
func (r SuiTransactionCommandCondition) ToParam() SuiTransactionCommandConditionParam {
	return param.Override[SuiTransactionCommandConditionParam](json.RawMessage(r.RawJSON()))
}

type SuiTransactionCommandConditionField string

const (
	SuiTransactionCommandConditionFieldCommandName SuiTransactionCommandConditionField = "commandName"
)

type SuiTransactionCommandConditionFieldSource string

const (
	SuiTransactionCommandConditionFieldSourceSuiTransactionCommand SuiTransactionCommandConditionFieldSource = "sui_transaction_command"
)

// SuiTransactionCommandConditionValueUnion contains all possible properties and
// values from [SuiCommandName], [[]SuiCommandName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSuiCommandName OfSuiCommandNameArray]
type SuiTransactionCommandConditionValueUnion struct {
	// This field will be present if the value is a [SuiCommandName] instead of an
	// object.
	OfSuiCommandName SuiCommandName `json:",inline"`
	// This field will be present if the value is a [[]SuiCommandName] instead of an
	// object.
	OfSuiCommandNameArray []SuiCommandName `json:",inline"`
	JSON                  struct {
		OfSuiCommandName      respjson.Field
		OfSuiCommandNameArray respjson.Field
		raw                   string
	} `json:"-"`
}

func (u SuiTransactionCommandConditionValueUnion) AsSuiCommandName() (v SuiCommandName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SuiTransactionCommandConditionValueUnion) AsSuiCommandNameArray() (v []SuiCommandName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SuiTransactionCommandConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *SuiTransactionCommandConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SUI transaction command attributes, enables allowlisting specific command types.
// Allowed commands: 'TransferObjects', 'SplitCoins', 'MergeCoins'. Only 'eq' and
// 'in' operators are supported.
//
// The properties Field, FieldSource, Operator, Value are required.
type SuiTransactionCommandConditionParam struct {
	// Any of "commandName".
	Field SuiTransactionCommandConditionField `json:"field,omitzero,required"`
	// Any of "sui_transaction_command".
	FieldSource SuiTransactionCommandConditionFieldSource `json:"field_source,omitzero,required"`
	// Operator to use for SUI transaction command conditions. Only 'eq' and 'in' are
	// supported for command names.
	//
	// Any of "eq", "in".
	Operator SuiTransactionCommandOperator `json:"operator,omitzero,required"`
	// Command name(s) to match. Must be one of: 'TransferObjects', 'SplitCoins',
	// 'MergeCoins'
	Value SuiTransactionCommandConditionValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r SuiTransactionCommandConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow SuiTransactionCommandConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuiTransactionCommandConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SuiTransactionCommandConditionValueUnionParam struct {
	// Check if union is this variant with !param.IsOmitted(union.OfSuiCommandName)
	OfSuiCommandName      param.Opt[SuiCommandName] `json:",omitzero,inline"`
	OfSuiCommandNameArray []SuiCommandName          `json:",omitzero,inline"`
	paramUnion
}

func (u SuiTransactionCommandConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSuiCommandName, u.OfSuiCommandNameArray)
}
func (u *SuiTransactionCommandConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SuiTransactionCommandConditionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSuiCommandName) {
		return &u.OfSuiCommandName
	} else if !param.IsOmitted(u.OfSuiCommandNameArray) {
		return &u.OfSuiCommandNameArray
	}
	return nil
}

// SUI TransferObjects command attributes, including recipient and amount fields.
type SuiTransferObjectsCommandCondition struct {
	// Supported fields for SUI TransferObjects command conditions. Only 'recipient'
	// and 'amount' are supported.
	//
	// Any of "recipient", "amount".
	Field SuiTransferObjectsCommandField `json:"field,required"`
	// Any of "sui_transfer_objects_command".
	FieldSource SuiTransferObjectsCommandConditionFieldSource `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator SuiTransferObjectsCommandConditionOperator   `json:"operator,required"`
	Value    SuiTransferObjectsCommandConditionValueUnion `json:"value,required"`
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
func (r SuiTransferObjectsCommandCondition) RawJSON() string { return r.JSON.raw }
func (r *SuiTransferObjectsCommandCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SuiTransferObjectsCommandCondition to a
// SuiTransferObjectsCommandConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SuiTransferObjectsCommandConditionParam.Overrides()
func (r SuiTransferObjectsCommandCondition) ToParam() SuiTransferObjectsCommandConditionParam {
	return param.Override[SuiTransferObjectsCommandConditionParam](json.RawMessage(r.RawJSON()))
}

type SuiTransferObjectsCommandConditionFieldSource string

const (
	SuiTransferObjectsCommandConditionFieldSourceSuiTransferObjectsCommand SuiTransferObjectsCommandConditionFieldSource = "sui_transfer_objects_command"
)

type SuiTransferObjectsCommandConditionOperator string

const (
	SuiTransferObjectsCommandConditionOperatorEq             SuiTransferObjectsCommandConditionOperator = "eq"
	SuiTransferObjectsCommandConditionOperatorGt             SuiTransferObjectsCommandConditionOperator = "gt"
	SuiTransferObjectsCommandConditionOperatorGte            SuiTransferObjectsCommandConditionOperator = "gte"
	SuiTransferObjectsCommandConditionOperatorLt             SuiTransferObjectsCommandConditionOperator = "lt"
	SuiTransferObjectsCommandConditionOperatorLte            SuiTransferObjectsCommandConditionOperator = "lte"
	SuiTransferObjectsCommandConditionOperatorIn             SuiTransferObjectsCommandConditionOperator = "in"
	SuiTransferObjectsCommandConditionOperatorInConditionSet SuiTransferObjectsCommandConditionOperator = "in_condition_set"
)

// SuiTransferObjectsCommandConditionValueUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type SuiTransferObjectsCommandConditionValueUnion struct {
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

func (u SuiTransferObjectsCommandConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SuiTransferObjectsCommandConditionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SuiTransferObjectsCommandConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *SuiTransferObjectsCommandConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SUI TransferObjects command attributes, including recipient and amount fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type SuiTransferObjectsCommandConditionParam struct {
	// Supported fields for SUI TransferObjects command conditions. Only 'recipient'
	// and 'amount' are supported.
	//
	// Any of "recipient", "amount".
	Field SuiTransferObjectsCommandField `json:"field,omitzero,required"`
	// Any of "sui_transfer_objects_command".
	FieldSource SuiTransferObjectsCommandConditionFieldSource `json:"field_source,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator SuiTransferObjectsCommandConditionOperator        `json:"operator,omitzero,required"`
	Value    SuiTransferObjectsCommandConditionValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r SuiTransferObjectsCommandConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow SuiTransferObjectsCommandConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuiTransferObjectsCommandConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SuiTransferObjectsCommandConditionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u SuiTransferObjectsCommandConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *SuiTransferObjectsCommandConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SuiTransferObjectsCommandConditionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type PolicyDeleteResponse struct {
	// Whether the policy was deleted successfully.
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rule that defines the conditions and action to take if the conditions are
// true.
type PolicyNewRuleResponse struct {
	ID string `json:"id,required"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyNewRuleResponseAction           `json:"action,required"`
	Conditions []PolicyNewRuleResponseConditionUnion `json:"conditions,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method PolicyNewRuleResponseMethod `json:"method,required"`
	Name   string                      `json:"name,required"`
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
func (r PolicyNewRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type PolicyNewRuleResponseAction string

const (
	PolicyNewRuleResponseActionAllow PolicyNewRuleResponseAction = "ALLOW"
	PolicyNewRuleResponseActionDeny  PolicyNewRuleResponseAction = "DENY"
)

// PolicyNewRuleResponseConditionUnion contains all possible properties and values
// from [PolicyNewRuleResponseConditionEthereumTransaction],
// [PolicyNewRuleResponseConditionEthereumCalldata],
// [PolicyNewRuleResponseConditionEthereumTypedDataDomain],
// [PolicyNewRuleResponseConditionEthereumTypedDataMessage],
// [PolicyNewRuleResponseConditionEthereum7702Authorization],
// [PolicyNewRuleResponseConditionSolanaProgramInstruction],
// [PolicyNewRuleResponseConditionSolanaSystemProgramInstruction],
// [PolicyNewRuleResponseConditionSolanaTokenProgramInstruction],
// [PolicyNewRuleResponseConditionSystem], [TronTransactionCondition],
// [SuiTransactionCommandCondition], [SuiTransferObjectsCommandCondition].
//
// Use the [PolicyNewRuleResponseConditionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyNewRuleResponseConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "sui_transaction_command",
	// "sui_transfer_objects_command".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [PolicyNewRuleResponseConditionEthereumTransactionValueUnion],
	// [PolicyNewRuleResponseConditionEthereumCalldataValueUnion],
	// [PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion],
	// [PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion],
	// [PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion],
	// [PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion],
	// [PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion],
	// [PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion],
	// [PolicyNewRuleResponseConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion]
	Value PolicyNewRuleResponseConditionUnionValue `json:"value"`
	// This field is from variant [PolicyNewRuleResponseConditionEthereumCalldata].
	Abi any `json:"abi"`
	// This field is from variant
	// [PolicyNewRuleResponseConditionEthereumTypedDataMessage].
	TypedData PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
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

// anyPolicyNewRuleResponseCondition is implemented by each variant of
// [PolicyNewRuleResponseConditionUnion] to add type safety for the return type of
// [PolicyNewRuleResponseConditionUnion.AsAny]
type anyPolicyNewRuleResponseCondition interface {
	implPolicyNewRuleResponseConditionUnion()
}

func (PolicyNewRuleResponseConditionEthereumTransaction) implPolicyNewRuleResponseConditionUnion() {}
func (PolicyNewRuleResponseConditionEthereumCalldata) implPolicyNewRuleResponseConditionUnion()    {}
func (PolicyNewRuleResponseConditionEthereumTypedDataDomain) implPolicyNewRuleResponseConditionUnion() {
}
func (PolicyNewRuleResponseConditionEthereumTypedDataMessage) implPolicyNewRuleResponseConditionUnion() {
}
func (PolicyNewRuleResponseConditionEthereum7702Authorization) implPolicyNewRuleResponseConditionUnion() {
}
func (PolicyNewRuleResponseConditionSolanaProgramInstruction) implPolicyNewRuleResponseConditionUnion() {
}
func (PolicyNewRuleResponseConditionSolanaSystemProgramInstruction) implPolicyNewRuleResponseConditionUnion() {
}
func (PolicyNewRuleResponseConditionSolanaTokenProgramInstruction) implPolicyNewRuleResponseConditionUnion() {
}
func (PolicyNewRuleResponseConditionSystem) implPolicyNewRuleResponseConditionUnion() {}
func (TronTransactionCondition) implPolicyNewRuleResponseConditionUnion()             {}
func (SuiTransactionCommandCondition) implPolicyNewRuleResponseConditionUnion()       {}
func (SuiTransferObjectsCommandCondition) implPolicyNewRuleResponseConditionUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyNewRuleResponseConditionUnion.AsAny().(type) {
//	case privyclient.PolicyNewRuleResponseConditionEthereumTransaction:
//	case privyclient.PolicyNewRuleResponseConditionEthereumCalldata:
//	case privyclient.PolicyNewRuleResponseConditionEthereumTypedDataDomain:
//	case privyclient.PolicyNewRuleResponseConditionEthereumTypedDataMessage:
//	case privyclient.PolicyNewRuleResponseConditionEthereum7702Authorization:
//	case privyclient.PolicyNewRuleResponseConditionSolanaProgramInstruction:
//	case privyclient.PolicyNewRuleResponseConditionSolanaSystemProgramInstruction:
//	case privyclient.PolicyNewRuleResponseConditionSolanaTokenProgramInstruction:
//	case privyclient.PolicyNewRuleResponseConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyNewRuleResponseConditionUnion) AsAny() anyPolicyNewRuleResponseCondition {
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
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	}
	return nil
}

func (u PolicyNewRuleResponseConditionUnion) AsEthereumTransaction() (v PolicyNewRuleResponseConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsEthereumCalldata() (v PolicyNewRuleResponseConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsEthereumTypedDataDomain() (v PolicyNewRuleResponseConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsEthereumTypedDataMessage() (v PolicyNewRuleResponseConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsEthereum7702Authorization() (v PolicyNewRuleResponseConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsSolanaProgramInstruction() (v PolicyNewRuleResponseConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsSolanaSystemProgramInstruction() (v PolicyNewRuleResponseConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsSolanaTokenProgramInstruction() (v PolicyNewRuleResponseConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsSystem() (v PolicyNewRuleResponseConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyNewRuleResponseConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionUnionValue is an implicit subunion of
// [PolicyNewRuleResponseConditionUnion]. PolicyNewRuleResponseConditionUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyNewRuleResponseConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyNewRuleResponseConditionUnionValue struct {
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

func (r *PolicyNewRuleResponseConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type PolicyNewRuleResponseConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field,required"`
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                      `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionEthereumTransactionValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionEthereumTransaction) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionEthereumTransactionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionEthereumTransactionValueUnion struct {
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

func (u PolicyNewRuleResponseConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type PolicyNewRuleResponseConditionEthereumCalldata struct {
	Abi         any                       `json:"abi,required"`
	Field       string                    `json:"field,required"`
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                   `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionEthereumCalldataValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionEthereumCalldata) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionEthereumCalldataValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionEthereumCalldataValueUnion struct {
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

func (u PolicyNewRuleResponseConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionEthereumCalldataValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyNewRuleResponseConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type PolicyNewRuleResponseConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field       string                           `json:"field,required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionEthereumTypedDataDomain) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion struct {
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

func (u PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type PolicyNewRuleResponseConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field,required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                           `json:"operator,required"`
	TypedData PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedData  `json:"typed_data,required"`
	Value     PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionEthereumTypedDataMessage) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                           `json:"primary_type,required"`
	Types       map[string][]PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedDataType `json:"types,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedDataType) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyNewRuleResponseConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion struct {
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

func (u PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_signAuthorization requests.
type PolicyNewRuleResponseConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field,required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                            `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionEthereum7702Authorization) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion struct {
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

func (u PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type PolicyNewRuleResponseConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field,required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                           `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionSolanaProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion struct {
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

func (u PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type PolicyNewRuleResponseConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field,required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyNewRuleResponseConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion struct {
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

func (u PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type PolicyNewRuleResponseConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field       string                                 `json:"field,required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyNewRuleResponseConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion struct {
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

func (u PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyNewRuleResponseConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type PolicyNewRuleResponseConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field,required"`
	FieldSource constant.System `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                         `json:"operator,required"`
	Value    PolicyNewRuleResponseConditionSystemValueUnion `json:"value,required"`
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
func (r PolicyNewRuleResponseConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *PolicyNewRuleResponseConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNewRuleResponseConditionSystemValueUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyNewRuleResponseConditionSystemValueUnion struct {
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

func (u PolicyNewRuleResponseConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNewRuleResponseConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNewRuleResponseConditionSystemValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyNewRuleResponseConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Method the rule applies to.
type PolicyNewRuleResponseMethod string

const (
	PolicyNewRuleResponseMethodEthSendTransaction       PolicyNewRuleResponseMethod = "eth_sendTransaction"
	PolicyNewRuleResponseMethodEthSignTransaction       PolicyNewRuleResponseMethod = "eth_signTransaction"
	PolicyNewRuleResponseMethodEthSignUserOperation     PolicyNewRuleResponseMethod = "eth_signUserOperation"
	PolicyNewRuleResponseMethodEthSignTypedDataV4       PolicyNewRuleResponseMethod = "eth_signTypedData_v4"
	PolicyNewRuleResponseMethodEthSign7702Authorization PolicyNewRuleResponseMethod = "eth_sign7702Authorization"
	PolicyNewRuleResponseMethodSignTransaction          PolicyNewRuleResponseMethod = "signTransaction"
	PolicyNewRuleResponseMethodSignAndSendTransaction   PolicyNewRuleResponseMethod = "signAndSendTransaction"
	PolicyNewRuleResponseMethodExportPrivateKey         PolicyNewRuleResponseMethod = "exportPrivateKey"
	PolicyNewRuleResponseMethodSignTransactionBytes     PolicyNewRuleResponseMethod = "signTransactionBytes"
	PolicyNewRuleResponseMethodStar                     PolicyNewRuleResponseMethod = "*"
)

type PolicyDeleteRuleResponse struct {
	// Whether the rule was deleted successfully.
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyDeleteRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyDeleteRuleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rule that defines the conditions and action to take if the conditions are
// true.
type PolicyGetRuleResponse struct {
	ID string `json:"id,required"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyGetRuleResponseAction           `json:"action,required"`
	Conditions []PolicyGetRuleResponseConditionUnion `json:"conditions,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method PolicyGetRuleResponseMethod `json:"method,required"`
	Name   string                      `json:"name,required"`
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
func (r PolicyGetRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type PolicyGetRuleResponseAction string

const (
	PolicyGetRuleResponseActionAllow PolicyGetRuleResponseAction = "ALLOW"
	PolicyGetRuleResponseActionDeny  PolicyGetRuleResponseAction = "DENY"
)

// PolicyGetRuleResponseConditionUnion contains all possible properties and values
// from [PolicyGetRuleResponseConditionEthereumTransaction],
// [PolicyGetRuleResponseConditionEthereumCalldata],
// [PolicyGetRuleResponseConditionEthereumTypedDataDomain],
// [PolicyGetRuleResponseConditionEthereumTypedDataMessage],
// [PolicyGetRuleResponseConditionEthereum7702Authorization],
// [PolicyGetRuleResponseConditionSolanaProgramInstruction],
// [PolicyGetRuleResponseConditionSolanaSystemProgramInstruction],
// [PolicyGetRuleResponseConditionSolanaTokenProgramInstruction],
// [PolicyGetRuleResponseConditionSystem], [TronTransactionCondition],
// [SuiTransactionCommandCondition], [SuiTransferObjectsCommandCondition].
//
// Use the [PolicyGetRuleResponseConditionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyGetRuleResponseConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "sui_transaction_command",
	// "sui_transfer_objects_command".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [PolicyGetRuleResponseConditionEthereumTransactionValueUnion],
	// [PolicyGetRuleResponseConditionEthereumCalldataValueUnion],
	// [PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion],
	// [PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion],
	// [PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion],
	// [PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion],
	// [PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion],
	// [PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion],
	// [PolicyGetRuleResponseConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion]
	Value PolicyGetRuleResponseConditionUnionValue `json:"value"`
	// This field is from variant [PolicyGetRuleResponseConditionEthereumCalldata].
	Abi any `json:"abi"`
	// This field is from variant
	// [PolicyGetRuleResponseConditionEthereumTypedDataMessage].
	TypedData PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
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

// anyPolicyGetRuleResponseCondition is implemented by each variant of
// [PolicyGetRuleResponseConditionUnion] to add type safety for the return type of
// [PolicyGetRuleResponseConditionUnion.AsAny]
type anyPolicyGetRuleResponseCondition interface {
	implPolicyGetRuleResponseConditionUnion()
}

func (PolicyGetRuleResponseConditionEthereumTransaction) implPolicyGetRuleResponseConditionUnion() {}
func (PolicyGetRuleResponseConditionEthereumCalldata) implPolicyGetRuleResponseConditionUnion()    {}
func (PolicyGetRuleResponseConditionEthereumTypedDataDomain) implPolicyGetRuleResponseConditionUnion() {
}
func (PolicyGetRuleResponseConditionEthereumTypedDataMessage) implPolicyGetRuleResponseConditionUnion() {
}
func (PolicyGetRuleResponseConditionEthereum7702Authorization) implPolicyGetRuleResponseConditionUnion() {
}
func (PolicyGetRuleResponseConditionSolanaProgramInstruction) implPolicyGetRuleResponseConditionUnion() {
}
func (PolicyGetRuleResponseConditionSolanaSystemProgramInstruction) implPolicyGetRuleResponseConditionUnion() {
}
func (PolicyGetRuleResponseConditionSolanaTokenProgramInstruction) implPolicyGetRuleResponseConditionUnion() {
}
func (PolicyGetRuleResponseConditionSystem) implPolicyGetRuleResponseConditionUnion() {}
func (TronTransactionCondition) implPolicyGetRuleResponseConditionUnion()             {}
func (SuiTransactionCommandCondition) implPolicyGetRuleResponseConditionUnion()       {}
func (SuiTransferObjectsCommandCondition) implPolicyGetRuleResponseConditionUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyGetRuleResponseConditionUnion.AsAny().(type) {
//	case privyclient.PolicyGetRuleResponseConditionEthereumTransaction:
//	case privyclient.PolicyGetRuleResponseConditionEthereumCalldata:
//	case privyclient.PolicyGetRuleResponseConditionEthereumTypedDataDomain:
//	case privyclient.PolicyGetRuleResponseConditionEthereumTypedDataMessage:
//	case privyclient.PolicyGetRuleResponseConditionEthereum7702Authorization:
//	case privyclient.PolicyGetRuleResponseConditionSolanaProgramInstruction:
//	case privyclient.PolicyGetRuleResponseConditionSolanaSystemProgramInstruction:
//	case privyclient.PolicyGetRuleResponseConditionSolanaTokenProgramInstruction:
//	case privyclient.PolicyGetRuleResponseConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyGetRuleResponseConditionUnion) AsAny() anyPolicyGetRuleResponseCondition {
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
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	}
	return nil
}

func (u PolicyGetRuleResponseConditionUnion) AsEthereumTransaction() (v PolicyGetRuleResponseConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsEthereumCalldata() (v PolicyGetRuleResponseConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsEthereumTypedDataDomain() (v PolicyGetRuleResponseConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsEthereumTypedDataMessage() (v PolicyGetRuleResponseConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsEthereum7702Authorization() (v PolicyGetRuleResponseConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsSolanaProgramInstruction() (v PolicyGetRuleResponseConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsSolanaSystemProgramInstruction() (v PolicyGetRuleResponseConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsSolanaTokenProgramInstruction() (v PolicyGetRuleResponseConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsSystem() (v PolicyGetRuleResponseConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyGetRuleResponseConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionUnionValue is an implicit subunion of
// [PolicyGetRuleResponseConditionUnion]. PolicyGetRuleResponseConditionUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyGetRuleResponseConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyGetRuleResponseConditionUnionValue struct {
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

func (r *PolicyGetRuleResponseConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type PolicyGetRuleResponseConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field,required"`
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                      `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionEthereumTransactionValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionEthereumTransaction) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionEthereumTransactionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionEthereumTransactionValueUnion struct {
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

func (u PolicyGetRuleResponseConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type PolicyGetRuleResponseConditionEthereumCalldata struct {
	Abi         any                       `json:"abi,required"`
	Field       string                    `json:"field,required"`
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                   `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionEthereumCalldataValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionEthereumCalldata) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionEthereumCalldataValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionEthereumCalldataValueUnion struct {
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

func (u PolicyGetRuleResponseConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionEthereumCalldataValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyGetRuleResponseConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type PolicyGetRuleResponseConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field       string                           `json:"field,required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionEthereumTypedDataDomain) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion struct {
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

func (u PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type PolicyGetRuleResponseConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field,required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                           `json:"operator,required"`
	TypedData PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedData  `json:"typed_data,required"`
	Value     PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionEthereumTypedDataMessage) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                           `json:"primary_type,required"`
	Types       map[string][]PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedDataType `json:"types,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedDataType) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyGetRuleResponseConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion struct {
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

func (u PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_signAuthorization requests.
type PolicyGetRuleResponseConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field,required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                            `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionEthereum7702Authorization) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion struct {
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

func (u PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type PolicyGetRuleResponseConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field,required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                           `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionSolanaProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion struct {
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

func (u PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type PolicyGetRuleResponseConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field,required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyGetRuleResponseConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion struct {
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

func (u PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type PolicyGetRuleResponseConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field       string                                 `json:"field,required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyGetRuleResponseConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion struct {
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

func (u PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyGetRuleResponseConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type PolicyGetRuleResponseConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field,required"`
	FieldSource constant.System `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                         `json:"operator,required"`
	Value    PolicyGetRuleResponseConditionSystemValueUnion `json:"value,required"`
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
func (r PolicyGetRuleResponseConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *PolicyGetRuleResponseConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyGetRuleResponseConditionSystemValueUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyGetRuleResponseConditionSystemValueUnion struct {
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

func (u PolicyGetRuleResponseConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyGetRuleResponseConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyGetRuleResponseConditionSystemValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyGetRuleResponseConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Method the rule applies to.
type PolicyGetRuleResponseMethod string

const (
	PolicyGetRuleResponseMethodEthSendTransaction       PolicyGetRuleResponseMethod = "eth_sendTransaction"
	PolicyGetRuleResponseMethodEthSignTransaction       PolicyGetRuleResponseMethod = "eth_signTransaction"
	PolicyGetRuleResponseMethodEthSignUserOperation     PolicyGetRuleResponseMethod = "eth_signUserOperation"
	PolicyGetRuleResponseMethodEthSignTypedDataV4       PolicyGetRuleResponseMethod = "eth_signTypedData_v4"
	PolicyGetRuleResponseMethodEthSign7702Authorization PolicyGetRuleResponseMethod = "eth_sign7702Authorization"
	PolicyGetRuleResponseMethodSignTransaction          PolicyGetRuleResponseMethod = "signTransaction"
	PolicyGetRuleResponseMethodSignAndSendTransaction   PolicyGetRuleResponseMethod = "signAndSendTransaction"
	PolicyGetRuleResponseMethodExportPrivateKey         PolicyGetRuleResponseMethod = "exportPrivateKey"
	PolicyGetRuleResponseMethodSignTransactionBytes     PolicyGetRuleResponseMethod = "signTransactionBytes"
	PolicyGetRuleResponseMethodStar                     PolicyGetRuleResponseMethod = "*"
)

type PolicyUpdateRuleResponse struct {
	ID string `json:"id,required"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyUpdateRuleResponseAction           `json:"action,required"`
	Conditions []PolicyUpdateRuleResponseConditionUnion `json:"conditions,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method PolicyUpdateRuleResponseMethod `json:"method,required"`
	Name   string                         `json:"name,required"`
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
func (r PolicyUpdateRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyUpdateRuleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type PolicyUpdateRuleResponseAction string

const (
	PolicyUpdateRuleResponseActionAllow PolicyUpdateRuleResponseAction = "ALLOW"
	PolicyUpdateRuleResponseActionDeny  PolicyUpdateRuleResponseAction = "DENY"
)

// PolicyUpdateRuleResponseConditionUnion contains all possible properties and
// values from [PolicyUpdateRuleResponseConditionEthereumTransaction],
// [PolicyUpdateRuleResponseConditionEthereumCalldata],
// [PolicyUpdateRuleResponseConditionEthereumTypedDataDomain],
// [PolicyUpdateRuleResponseConditionEthereumTypedDataMessage],
// [PolicyUpdateRuleResponseConditionEthereum7702Authorization],
// [PolicyUpdateRuleResponseConditionSolanaProgramInstruction],
// [PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction],
// [PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction],
// [PolicyUpdateRuleResponseConditionSystem], [TronTransactionCondition],
// [SuiTransactionCommandCondition], [SuiTransferObjectsCommandCondition].
//
// Use the [PolicyUpdateRuleResponseConditionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyUpdateRuleResponseConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "sui_transaction_command",
	// "sui_transfer_objects_command".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of
	// [PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion],
	// [PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion],
	// [PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion],
	// [PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion],
	// [PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion],
	// [PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion],
	// [PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion],
	// [PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion],
	// [PolicyUpdateRuleResponseConditionSystemValueUnion],
	// [TronTransactionConditionValueUnion],
	// [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion]
	Value PolicyUpdateRuleResponseConditionUnionValue `json:"value"`
	// This field is from variant [PolicyUpdateRuleResponseConditionEthereumCalldata].
	Abi any `json:"abi"`
	// This field is from variant
	// [PolicyUpdateRuleResponseConditionEthereumTypedDataMessage].
	TypedData PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
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

// anyPolicyUpdateRuleResponseCondition is implemented by each variant of
// [PolicyUpdateRuleResponseConditionUnion] to add type safety for the return type
// of [PolicyUpdateRuleResponseConditionUnion.AsAny]
type anyPolicyUpdateRuleResponseCondition interface {
	implPolicyUpdateRuleResponseConditionUnion()
}

func (PolicyUpdateRuleResponseConditionEthereumTransaction) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionEthereumCalldata) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionEthereumTypedDataDomain) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionEthereumTypedDataMessage) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionEthereum7702Authorization) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionSolanaProgramInstruction) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction) implPolicyUpdateRuleResponseConditionUnion() {
}
func (PolicyUpdateRuleResponseConditionSystem) implPolicyUpdateRuleResponseConditionUnion() {}
func (TronTransactionCondition) implPolicyUpdateRuleResponseConditionUnion()                {}
func (SuiTransactionCommandCondition) implPolicyUpdateRuleResponseConditionUnion()          {}
func (SuiTransferObjectsCommandCondition) implPolicyUpdateRuleResponseConditionUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyUpdateRuleResponseConditionUnion.AsAny().(type) {
//	case privyclient.PolicyUpdateRuleResponseConditionEthereumTransaction:
//	case privyclient.PolicyUpdateRuleResponseConditionEthereumCalldata:
//	case privyclient.PolicyUpdateRuleResponseConditionEthereumTypedDataDomain:
//	case privyclient.PolicyUpdateRuleResponseConditionEthereumTypedDataMessage:
//	case privyclient.PolicyUpdateRuleResponseConditionEthereum7702Authorization:
//	case privyclient.PolicyUpdateRuleResponseConditionSolanaProgramInstruction:
//	case privyclient.PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction:
//	case privyclient.PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction:
//	case privyclient.PolicyUpdateRuleResponseConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyUpdateRuleResponseConditionUnion) AsAny() anyPolicyUpdateRuleResponseCondition {
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
	case "sui_transaction_command":
		return u.AsSuiTransactionCommand()
	case "sui_transfer_objects_command":
		return u.AsSuiTransferObjectsCommand()
	}
	return nil
}

func (u PolicyUpdateRuleResponseConditionUnion) AsEthereumTransaction() (v PolicyUpdateRuleResponseConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsEthereumCalldata() (v PolicyUpdateRuleResponseConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsEthereumTypedDataDomain() (v PolicyUpdateRuleResponseConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsEthereumTypedDataMessage() (v PolicyUpdateRuleResponseConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsEthereum7702Authorization() (v PolicyUpdateRuleResponseConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsSolanaProgramInstruction() (v PolicyUpdateRuleResponseConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsSolanaSystemProgramInstruction() (v PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsSolanaTokenProgramInstruction() (v PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsSystem() (v PolicyUpdateRuleResponseConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyUpdateRuleResponseConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionUnionValue is an implicit subunion of
// [PolicyUpdateRuleResponseConditionUnion].
// PolicyUpdateRuleResponseConditionUnionValue provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyUpdateRuleResponseConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyUpdateRuleResponseConditionUnionValue struct {
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

func (r *PolicyUpdateRuleResponseConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type PolicyUpdateRuleResponseConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field,required"`
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                         `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionEthereumTransaction) RawJSON() string { return r.JSON.raw }
func (r *PolicyUpdateRuleResponseConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type PolicyUpdateRuleResponseConditionEthereumCalldata struct {
	Abi         any                       `json:"abi,required"`
	Field       string                    `json:"field,required"`
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                      `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionEthereumCalldata) RawJSON() string { return r.JSON.raw }
func (r *PolicyUpdateRuleResponseConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type PolicyUpdateRuleResponseConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field       string                           `json:"field,required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                             `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionEthereumTypedDataDomain) RawJSON() string { return r.JSON.raw }
func (r *PolicyUpdateRuleResponseConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type PolicyUpdateRuleResponseConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field,required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                              `json:"operator,required"`
	TypedData PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedData  `json:"typed_data,required"`
	Value     PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionEthereumTypedDataMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                              `json:"primary_type,required"`
	Types       map[string][]PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedDataType `json:"types,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedData) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedDataType) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_signAuthorization requests.
type PolicyUpdateRuleResponseConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field,required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                               `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionEthereum7702Authorization) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type PolicyUpdateRuleResponseConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field,required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                              `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionSolanaProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field,required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                    `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field       string                                 `json:"field,required"`
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                   `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyUpdateRuleResponseConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion
// contains all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyUpdateRuleResponseConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type PolicyUpdateRuleResponseConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field,required"`
	FieldSource constant.System `json:"field_source,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                            `json:"operator,required"`
	Value    PolicyUpdateRuleResponseConditionSystemValueUnion `json:"value,required"`
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
func (r PolicyUpdateRuleResponseConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *PolicyUpdateRuleResponseConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyUpdateRuleResponseConditionSystemValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyUpdateRuleResponseConditionSystemValueUnion struct {
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

func (u PolicyUpdateRuleResponseConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyUpdateRuleResponseConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyUpdateRuleResponseConditionSystemValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyUpdateRuleResponseConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Method the rule applies to.
type PolicyUpdateRuleResponseMethod string

const (
	PolicyUpdateRuleResponseMethodEthSendTransaction       PolicyUpdateRuleResponseMethod = "eth_sendTransaction"
	PolicyUpdateRuleResponseMethodEthSignTransaction       PolicyUpdateRuleResponseMethod = "eth_signTransaction"
	PolicyUpdateRuleResponseMethodEthSignUserOperation     PolicyUpdateRuleResponseMethod = "eth_signUserOperation"
	PolicyUpdateRuleResponseMethodEthSignTypedDataV4       PolicyUpdateRuleResponseMethod = "eth_signTypedData_v4"
	PolicyUpdateRuleResponseMethodEthSign7702Authorization PolicyUpdateRuleResponseMethod = "eth_sign7702Authorization"
	PolicyUpdateRuleResponseMethodSignTransaction          PolicyUpdateRuleResponseMethod = "signTransaction"
	PolicyUpdateRuleResponseMethodSignAndSendTransaction   PolicyUpdateRuleResponseMethod = "signAndSendTransaction"
	PolicyUpdateRuleResponseMethodExportPrivateKey         PolicyUpdateRuleResponseMethod = "exportPrivateKey"
	PolicyUpdateRuleResponseMethodSignTransactionBytes     PolicyUpdateRuleResponseMethod = "signTransactionBytes"
	PolicyUpdateRuleResponseMethodStar                     PolicyUpdateRuleResponseMethod = "*"
)

type PolicyNewParams struct {
	// The chain type the policy applies to.
	//
	// Any of "ethereum", "solana", "tron", "sui".
	ChainType PolicyNewParamsChainType `json:"chain_type,omitzero,required"`
	// Name to assign to policy.
	Name  string                `json:"name,required"`
	Rules []PolicyNewParamsRule `json:"rules,omitzero,required"`
	// Version of the policy. Currently, 1.0 is the only version.
	//
	// Any of "1.0".
	Version PolicyNewParamsVersion `json:"version,omitzero,required"`
	OwnerID param.Opt[string]      `json:"owner_id,omitzero"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner PolicyNewParamsOwnerUnion `json:"owner,omitzero"`
	paramObj
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The chain type the policy applies to.
type PolicyNewParamsChainType string

const (
	PolicyNewParamsChainTypeEthereum PolicyNewParamsChainType = "ethereum"
	PolicyNewParamsChainTypeSolana   PolicyNewParamsChainType = "solana"
	PolicyNewParamsChainTypeTron     PolicyNewParamsChainType = "tron"
	PolicyNewParamsChainTypeSui      PolicyNewParamsChainType = "sui"
)

// The rules that apply to each method the policy covers.
//
// The properties Action, Conditions, Method, Name are required.
type PolicyNewParamsRule struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                              `json:"action,omitzero,required"`
	Conditions []PolicyNewParamsRuleConditionUnion `json:"conditions,omitzero,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method,omitzero,required"`
	Name   string `json:"name,required"`
	paramObj
}

func (r PolicyNewParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRule](
		"action", "ALLOW", "DENY",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRule](
		"method", "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation", "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction", "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "*",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionUnion struct {
	OfEthereumTransaction            *PolicyNewParamsRuleConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *PolicyNewParamsRuleConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *PolicyNewParamsRuleConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *PolicyNewParamsRuleConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *PolicyNewParamsRuleConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *PolicyNewParamsRuleConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *PolicyNewParamsRuleConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *PolicyNewParamsRuleConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *PolicyNewParamsRuleConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                              `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                        `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                    `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionUnion) MarshalJSON() ([]byte, error) {
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
}
func (u *PolicyNewParamsRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionUnion) asAny() any {
	if !param.IsOmitted(u.OfEthereumTransaction) {
		return u.OfEthereumTransaction
	} else if !param.IsOmitted(u.OfEthereumCalldata) {
		return u.OfEthereumCalldata
	} else if !param.IsOmitted(u.OfEthereumTypedDataDomain) {
		return u.OfEthereumTypedDataDomain
	} else if !param.IsOmitted(u.OfEthereumTypedDataMessage) {
		return u.OfEthereumTypedDataMessage
	} else if !param.IsOmitted(u.OfEthereum7702Authorization) {
		return u.OfEthereum7702Authorization
	} else if !param.IsOmitted(u.OfSolanaProgramInstruction) {
		return u.OfSolanaProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaSystemProgramInstruction) {
		return u.OfSolanaSystemProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaTokenProgramInstruction) {
		return u.OfSolanaTokenProgramInstruction
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTronTransaction) {
		return u.OfTronTransaction
	} else if !param.IsOmitted(u.OfSuiTransactionCommand) {
		return u.OfSuiTransactionCommand
	} else if !param.IsOmitted(u.OfSuiTransferObjectsCommand) {
		return u.OfSuiTransferObjectsCommand
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewParamsRuleConditionUnion) GetAbi() *any {
	if vt := u.OfEthereumCalldata; vt != nil {
		return &vt.Abi
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewParamsRuleConditionUnion) GetTypedData() *PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedData {
	if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return &vt.TypedData
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewParamsRuleConditionUnion) GetField() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Field)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewParamsRuleConditionUnion) GetFieldSource() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewParamsRuleConditionUnion) GetOperator() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PolicyNewParamsRuleConditionUnion) GetValue() (res policyNewParamsRuleConditionUnionValue) {
	if vt := u.OfEthereumTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumCalldata; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfTronTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		res.any = vt.Value.asAny()
	}
	return
}

// Can have the runtime types [*string], [_[]string], [_[]SuiCommandName]
type policyNewParamsRuleConditionUnionValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	case *[]privyclient.SuiCommandName:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u policyNewParamsRuleConditionUnionValue) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[PolicyNewParamsRuleConditionUnion](
		"field_source",
		apijson.Discriminator[PolicyNewParamsRuleConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[PolicyNewParamsRuleConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[PolicyNewParamsRuleConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[PolicyNewParamsRuleConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[PolicyNewParamsRuleConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[PolicyNewParamsRuleConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[PolicyNewParamsRuleConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[PolicyNewParamsRuleConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[PolicyNewParamsRuleConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                    `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionEthereumTransactionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionEthereumTransactionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero,required"`
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                 `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionEthereumCalldataValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionEthereumCalldataValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                        `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionEthereumTypedDataDomainValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type PolicyNewParamsRuleConditionEthereumTypedDataMessage struct {
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                         `json:"operator,omitzero,required"`
	TypedData PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero,required"`
	Value     PolicyNewParamsRuleConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                         `json:"primary_type,required"`
	Types       map[string][]PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionEthereumTypedDataMessageValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Allowed contract addresses for eth_signAuthorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionEthereum7702AuthorizationValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                         `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionSolanaProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionSolanaProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                               `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionSolanaSystemProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                              `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionSolanaTokenProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewParamsRuleConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                       `json:"operator,omitzero,required"`
	Value    PolicyNewParamsRuleConditionSystemValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source,required"`
	paramObj
}

func (r PolicyNewParamsRuleConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRuleConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRuleConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[PolicyNewParamsRuleConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsRuleConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsRuleConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewParamsRuleConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsRuleConditionSystemValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Version of the policy. Currently, 1.0 is the only version.
type PolicyNewParamsVersion string

const (
	PolicyNewParamsVersion1_0 PolicyNewParamsVersion = "1.0"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewParamsOwnerUnion struct {
	OfPublicKeyOwner *PolicyNewParamsOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *PolicyNewParamsOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *PolicyNewParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewParamsOwnerUnion) asAny() any {
	if !param.IsOmitted(u.OfPublicKeyOwner) {
		return u.OfPublicKeyOwner
	} else if !param.IsOmitted(u.OfUserOwner) {
		return u.OfUserOwner
	}
	return nil
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type PolicyNewParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key,required"`
	paramObj
}

func (r PolicyNewParamsOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type PolicyNewParamsOwnerUserOwner struct {
	UserID string `json:"user_id,required"`
	paramObj
}

func (r PolicyNewParamsOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyUpdateParams struct {
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// Name to assign to policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner PolicyUpdateParamsOwnerUnion `json:"owner,omitzero"`
	Rules []PolicyUpdateParamsRule     `json:"rules,omitzero"`
	paramObj
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsOwnerUnion struct {
	OfPublicKeyOwner *PolicyUpdateParamsOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *PolicyUpdateParamsOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *PolicyUpdateParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsOwnerUnion) asAny() any {
	if !param.IsOmitted(u.OfPublicKeyOwner) {
		return u.OfPublicKeyOwner
	} else if !param.IsOmitted(u.OfUserOwner) {
		return u.OfUserOwner
	}
	return nil
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type PolicyUpdateParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key,required"`
	paramObj
}

func (r PolicyUpdateParamsOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type PolicyUpdateParamsOwnerUserOwner struct {
	UserID string `json:"user_id,required"`
	paramObj
}

func (r PolicyUpdateParamsOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rules that apply to each method the policy covers.
//
// The properties Action, Conditions, Method, Name are required.
type PolicyUpdateParamsRule struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     string                                 `json:"action,omitzero,required"`
	Conditions []PolicyUpdateParamsRuleConditionUnion `json:"conditions,omitzero,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method string `json:"method,omitzero,required"`
	Name   string `json:"name,required"`
	paramObj
}

func (r PolicyUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRule](
		"action", "ALLOW", "DENY",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRule](
		"method", "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation", "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction", "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "*",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionUnion struct {
	OfEthereumTransaction            *PolicyUpdateParamsRuleConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *PolicyUpdateParamsRuleConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *PolicyUpdateParamsRuleConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *PolicyUpdateParamsRuleConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *PolicyUpdateParamsRuleConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *PolicyUpdateParamsRuleConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *PolicyUpdateParamsRuleConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                                 `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                           `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                       `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionUnion) MarshalJSON() ([]byte, error) {
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
}
func (u *PolicyUpdateParamsRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionUnion) asAny() any {
	if !param.IsOmitted(u.OfEthereumTransaction) {
		return u.OfEthereumTransaction
	} else if !param.IsOmitted(u.OfEthereumCalldata) {
		return u.OfEthereumCalldata
	} else if !param.IsOmitted(u.OfEthereumTypedDataDomain) {
		return u.OfEthereumTypedDataDomain
	} else if !param.IsOmitted(u.OfEthereumTypedDataMessage) {
		return u.OfEthereumTypedDataMessage
	} else if !param.IsOmitted(u.OfEthereum7702Authorization) {
		return u.OfEthereum7702Authorization
	} else if !param.IsOmitted(u.OfSolanaProgramInstruction) {
		return u.OfSolanaProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaSystemProgramInstruction) {
		return u.OfSolanaSystemProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaTokenProgramInstruction) {
		return u.OfSolanaTokenProgramInstruction
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTronTransaction) {
		return u.OfTronTransaction
	} else if !param.IsOmitted(u.OfSuiTransactionCommand) {
		return u.OfSuiTransactionCommand
	} else if !param.IsOmitted(u.OfSuiTransferObjectsCommand) {
		return u.OfSuiTransferObjectsCommand
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateParamsRuleConditionUnion) GetAbi() *any {
	if vt := u.OfEthereumCalldata; vt != nil {
		return &vt.Abi
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateParamsRuleConditionUnion) GetTypedData() *PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedData {
	if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return &vt.TypedData
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateParamsRuleConditionUnion) GetField() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Field)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateParamsRuleConditionUnion) GetFieldSource() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateParamsRuleConditionUnion) GetOperator() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PolicyUpdateParamsRuleConditionUnion) GetValue() (res policyUpdateParamsRuleConditionUnionValue) {
	if vt := u.OfEthereumTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumCalldata; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfTronTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		res.any = vt.Value.asAny()
	}
	return
}

// Can have the runtime types [*string], [_[]string], [_[]SuiCommandName]
type policyUpdateParamsRuleConditionUnionValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	case *[]privyclient.SuiCommandName:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u policyUpdateParamsRuleConditionUnionValue) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[PolicyUpdateParamsRuleConditionUnion](
		"field_source",
		apijson.Discriminator[PolicyUpdateParamsRuleConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[PolicyUpdateParamsRuleConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                       `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero,required"`
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                    `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionEthereumCalldataValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionEthereumCalldataValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                           `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionEthereumTypedDataDomainValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type PolicyUpdateParamsRuleConditionEthereumTypedDataMessage struct {
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                            `json:"operator,omitzero,required"`
	TypedData PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero,required"`
	Value     PolicyUpdateParamsRuleConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                            `json:"primary_type,required"`
	Types       map[string][]PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionEthereumTypedDataMessageValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Allowed contract addresses for eth_signAuthorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                             `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionEthereum7702AuthorizationValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                            `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionSolanaProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionSolanaProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                  `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionSolanaSystemProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionSolanaTokenProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateParamsRuleConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                          `json:"operator,omitzero,required"`
	Value    PolicyUpdateParamsRuleConditionSystemValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateParamsRuleConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsRuleConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsRuleConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[PolicyUpdateParamsRuleConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateParamsRuleConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateParamsRuleConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateParamsRuleConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateParamsRuleConditionSystemValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type PolicyDeleteParams struct {
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

type PolicyNewRuleParams struct {
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyNewRuleParamsAction           `json:"action,omitzero,required"`
	Conditions []PolicyNewRuleParamsConditionUnion `json:"conditions,omitzero,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method PolicyNewRuleParamsMethod `json:"method,omitzero,required"`
	Name   string                    `json:"name,required"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r PolicyNewRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type PolicyNewRuleParamsAction string

const (
	PolicyNewRuleParamsActionAllow PolicyNewRuleParamsAction = "ALLOW"
	PolicyNewRuleParamsActionDeny  PolicyNewRuleParamsAction = "DENY"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionUnion struct {
	OfEthereumTransaction            *PolicyNewRuleParamsConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *PolicyNewRuleParamsConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *PolicyNewRuleParamsConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *PolicyNewRuleParamsConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *PolicyNewRuleParamsConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *PolicyNewRuleParamsConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *PolicyNewRuleParamsConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *PolicyNewRuleParamsConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *PolicyNewRuleParamsConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                              `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                        `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                    `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionUnion) MarshalJSON() ([]byte, error) {
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
}
func (u *PolicyNewRuleParamsConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionUnion) asAny() any {
	if !param.IsOmitted(u.OfEthereumTransaction) {
		return u.OfEthereumTransaction
	} else if !param.IsOmitted(u.OfEthereumCalldata) {
		return u.OfEthereumCalldata
	} else if !param.IsOmitted(u.OfEthereumTypedDataDomain) {
		return u.OfEthereumTypedDataDomain
	} else if !param.IsOmitted(u.OfEthereumTypedDataMessage) {
		return u.OfEthereumTypedDataMessage
	} else if !param.IsOmitted(u.OfEthereum7702Authorization) {
		return u.OfEthereum7702Authorization
	} else if !param.IsOmitted(u.OfSolanaProgramInstruction) {
		return u.OfSolanaProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaSystemProgramInstruction) {
		return u.OfSolanaSystemProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaTokenProgramInstruction) {
		return u.OfSolanaTokenProgramInstruction
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTronTransaction) {
		return u.OfTronTransaction
	} else if !param.IsOmitted(u.OfSuiTransactionCommand) {
		return u.OfSuiTransactionCommand
	} else if !param.IsOmitted(u.OfSuiTransferObjectsCommand) {
		return u.OfSuiTransferObjectsCommand
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewRuleParamsConditionUnion) GetAbi() *any {
	if vt := u.OfEthereumCalldata; vt != nil {
		return &vt.Abi
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewRuleParamsConditionUnion) GetTypedData() *PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedData {
	if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return &vt.TypedData
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewRuleParamsConditionUnion) GetField() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Field)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewRuleParamsConditionUnion) GetFieldSource() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyNewRuleParamsConditionUnion) GetOperator() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PolicyNewRuleParamsConditionUnion) GetValue() (res policyNewRuleParamsConditionUnionValue) {
	if vt := u.OfEthereumTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumCalldata; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfTronTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		res.any = vt.Value.asAny()
	}
	return
}

// Can have the runtime types [*string], [_[]string], [_[]SuiCommandName]
type policyNewRuleParamsConditionUnionValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	case *[]privyclient.SuiCommandName:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u policyNewRuleParamsConditionUnionValue) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[PolicyNewRuleParamsConditionUnion](
		"field_source",
		apijson.Discriminator[PolicyNewRuleParamsConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[PolicyNewRuleParamsConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[PolicyNewRuleParamsConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[PolicyNewRuleParamsConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[PolicyNewRuleParamsConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[PolicyNewRuleParamsConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[PolicyNewRuleParamsConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[PolicyNewRuleParamsConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[PolicyNewRuleParamsConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                    `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionEthereumTransactionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionEthereumTransactionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero,required"`
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                 `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionEthereumCalldataValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionEthereumCalldataValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                        `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionEthereumTypedDataDomainValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type PolicyNewRuleParamsConditionEthereumTypedDataMessage struct {
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                         `json:"operator,omitzero,required"`
	TypedData PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero,required"`
	Value     PolicyNewRuleParamsConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                         `json:"primary_type,required"`
	Types       map[string][]PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionEthereumTypedDataMessageValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Allowed contract addresses for eth_signAuthorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                          `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionEthereum7702AuthorizationValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                         `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionSolanaProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionSolanaProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                               `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionSolanaSystemProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                              `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionSolanaTokenProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyNewRuleParamsConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                       `json:"operator,omitzero,required"`
	Value    PolicyNewRuleParamsConditionSystemValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source,required"`
	paramObj
}

func (r PolicyNewRuleParamsConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewRuleParamsConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewRuleParamsConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[PolicyNewRuleParamsConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyNewRuleParamsConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyNewRuleParamsConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyNewRuleParamsConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyNewRuleParamsConditionSystemValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Method the rule applies to.
type PolicyNewRuleParamsMethod string

const (
	PolicyNewRuleParamsMethodEthSendTransaction       PolicyNewRuleParamsMethod = "eth_sendTransaction"
	PolicyNewRuleParamsMethodEthSignTransaction       PolicyNewRuleParamsMethod = "eth_signTransaction"
	PolicyNewRuleParamsMethodEthSignUserOperation     PolicyNewRuleParamsMethod = "eth_signUserOperation"
	PolicyNewRuleParamsMethodEthSignTypedDataV4       PolicyNewRuleParamsMethod = "eth_signTypedData_v4"
	PolicyNewRuleParamsMethodEthSign7702Authorization PolicyNewRuleParamsMethod = "eth_sign7702Authorization"
	PolicyNewRuleParamsMethodSignTransaction          PolicyNewRuleParamsMethod = "signTransaction"
	PolicyNewRuleParamsMethodSignAndSendTransaction   PolicyNewRuleParamsMethod = "signAndSendTransaction"
	PolicyNewRuleParamsMethodExportPrivateKey         PolicyNewRuleParamsMethod = "exportPrivateKey"
	PolicyNewRuleParamsMethodSignTransactionBytes     PolicyNewRuleParamsMethod = "signTransactionBytes"
	PolicyNewRuleParamsMethodStar                     PolicyNewRuleParamsMethod = "*"
)

type PolicyDeleteRuleParams struct {
	PolicyID string `path:"policy_id,required" json:"-"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

type PolicyGetRuleParams struct {
	PolicyID string `path:"policy_id,required" json:"-"`
	paramObj
}

type PolicyUpdateRuleParams struct {
	PolicyID string `path:"policy_id,required" json:"-"`
	// Action to take if the conditions are true.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyUpdateRuleParamsAction           `json:"action,omitzero,required"`
	Conditions []PolicyUpdateRuleParamsConditionUnion `json:"conditions,omitzero,required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "signTransactionBytes", "\*".
	Method PolicyUpdateRuleParamsMethod `json:"method,omitzero,required"`
	Name   string                       `json:"name,required"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r PolicyUpdateRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take if the conditions are true.
type PolicyUpdateRuleParamsAction string

const (
	PolicyUpdateRuleParamsActionAllow PolicyUpdateRuleParamsAction = "ALLOW"
	PolicyUpdateRuleParamsActionDeny  PolicyUpdateRuleParamsAction = "DENY"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionUnion struct {
	OfEthereumTransaction            *PolicyUpdateRuleParamsConditionEthereumTransaction            `json:",omitzero,inline"`
	OfEthereumCalldata               *PolicyUpdateRuleParamsConditionEthereumCalldata               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *PolicyUpdateRuleParamsConditionEthereumTypedDataDomain        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *PolicyUpdateRuleParamsConditionEthereumTypedDataMessage       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *PolicyUpdateRuleParamsConditionEthereum7702Authorization      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *PolicyUpdateRuleParamsConditionSolanaProgramInstruction       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction  `json:",omitzero,inline"`
	OfSystem                         *PolicyUpdateRuleParamsConditionSystem                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                                 `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                           `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                       `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionUnion) MarshalJSON() ([]byte, error) {
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
}
func (u *PolicyUpdateRuleParamsConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionUnion) asAny() any {
	if !param.IsOmitted(u.OfEthereumTransaction) {
		return u.OfEthereumTransaction
	} else if !param.IsOmitted(u.OfEthereumCalldata) {
		return u.OfEthereumCalldata
	} else if !param.IsOmitted(u.OfEthereumTypedDataDomain) {
		return u.OfEthereumTypedDataDomain
	} else if !param.IsOmitted(u.OfEthereumTypedDataMessage) {
		return u.OfEthereumTypedDataMessage
	} else if !param.IsOmitted(u.OfEthereum7702Authorization) {
		return u.OfEthereum7702Authorization
	} else if !param.IsOmitted(u.OfSolanaProgramInstruction) {
		return u.OfSolanaProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaSystemProgramInstruction) {
		return u.OfSolanaSystemProgramInstruction
	} else if !param.IsOmitted(u.OfSolanaTokenProgramInstruction) {
		return u.OfSolanaTokenProgramInstruction
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTronTransaction) {
		return u.OfTronTransaction
	} else if !param.IsOmitted(u.OfSuiTransactionCommand) {
		return u.OfSuiTransactionCommand
	} else if !param.IsOmitted(u.OfSuiTransferObjectsCommand) {
		return u.OfSuiTransferObjectsCommand
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateRuleParamsConditionUnion) GetAbi() *any {
	if vt := u.OfEthereumCalldata; vt != nil {
		return &vt.Abi
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateRuleParamsConditionUnion) GetTypedData() *PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedData {
	if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return &vt.TypedData
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateRuleParamsConditionUnion) GetField() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Field)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateRuleParamsConditionUnion) GetFieldSource() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.FieldSource)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PolicyUpdateRuleParamsConditionUnion) GetOperator() *string {
	if vt := u.OfEthereumTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumCalldata; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfTronTransaction; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PolicyUpdateRuleParamsConditionUnion) GetValue() (res policyUpdateRuleParamsConditionUnionValue) {
	if vt := u.OfEthereumTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumCalldata; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataDomain; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereumTypedDataMessage; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfEthereum7702Authorization; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaSystemProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSolanaTokenProgramInstruction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfTronTransaction; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransactionCommand; vt != nil {
		res.any = vt.Value.asAny()
	} else if vt := u.OfSuiTransferObjectsCommand; vt != nil {
		res.any = vt.Value.asAny()
	}
	return
}

// Can have the runtime types [*string], [_[]string], [_[]SuiCommandName]
type policyUpdateRuleParamsConditionUnionValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	case *[]privyclient.SuiCommandName:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u policyUpdateRuleParamsConditionUnionValue) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[PolicyUpdateRuleParamsConditionUnion](
		"field_source",
		apijson.Discriminator[PolicyUpdateRuleParamsConditionEthereumTransaction]("ethereum_transaction"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionEthereumCalldata]("ethereum_calldata"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionEthereumTypedDataDomain]("ethereum_typed_data_domain"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionEthereumTypedDataMessage]("ethereum_typed_data_message"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionEthereum7702Authorization]("ethereum_7702_authorization"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionSolanaProgramInstruction]("solana_program_instruction"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction]("solana_system_program_instruction"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction]("solana_token_program_instruction"),
		apijson.Discriminator[PolicyUpdateRuleParamsConditionSystem]("system"),
		apijson.Discriminator[TronTransactionConditionParam]("tron_transaction"),
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
	)
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                       `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereumTransaction](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereumTransaction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionEthereumCalldata struct {
	Abi   any    `json:"abi,omitzero,required"`
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                    `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionEthereumCalldataValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereumCalldata) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereumCalldata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereumCalldata](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionEthereumCalldataValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionEthereumCalldataValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionEthereumCalldataValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                           `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereumTypedDataDomain) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereumTypedDataDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereumTypedDataDomain](
		"field", "chainId", "verifyingContract",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereumTypedDataDomain](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionEthereumTypedDataDomainValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionEthereumTypedDataDomainValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionEthereumTypedDataDomainValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type PolicyUpdateRuleParamsConditionEthereumTypedDataMessage struct {
	Field string `json:"field,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                            `json:"operator,omitzero,required"`
	TypedData PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedData  `json:"typed_data,omitzero,required"`
	Value     PolicyUpdateRuleParamsConditionEthereumTypedDataMessageValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereumTypedDataMessage) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereumTypedDataMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereumTypedDataMessage](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedData struct {
	PrimaryType string                                                                            `json:"primary_type,required"`
	Types       map[string][]PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionEthereumTypedDataMessageValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionEthereumTypedDataMessageValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionEthereumTypedDataMessageValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Allowed contract addresses for eth_signAuthorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                             `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionEthereum7702AuthorizationValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionEthereum7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionEthereum7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereum7702Authorization](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionEthereum7702Authorization](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionEthereum7702AuthorizationValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionEthereum7702AuthorizationValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionEthereum7702AuthorizationValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                            `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionSolanaProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionSolanaProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionSolanaProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSolanaProgramInstruction](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSolanaProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionSolanaProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionSolanaProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionSolanaProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                  `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionSolanaSystemProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSolanaSystemProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionSolanaSystemProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionSolanaSystemProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionSolanaSystemProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction struct {
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction](
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSolanaTokenProgramInstruction](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionSolanaTokenProgramInstructionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionSolanaTokenProgramInstructionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionSolanaTokenProgramInstructionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyUpdateRuleParamsConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero,required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                          `json:"operator,omitzero,required"`
	Value    PolicyUpdateRuleParamsConditionSystemValueUnion `json:"value,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source,required"`
	paramObj
}

func (r PolicyUpdateRuleParamsConditionSystem) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateRuleParamsConditionSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateRuleParamsConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSystem](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[PolicyUpdateRuleParamsConditionSystem](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyUpdateRuleParamsConditionSystemValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyUpdateRuleParamsConditionSystemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyUpdateRuleParamsConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PolicyUpdateRuleParamsConditionSystemValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Method the rule applies to.
type PolicyUpdateRuleParamsMethod string

const (
	PolicyUpdateRuleParamsMethodEthSendTransaction       PolicyUpdateRuleParamsMethod = "eth_sendTransaction"
	PolicyUpdateRuleParamsMethodEthSignTransaction       PolicyUpdateRuleParamsMethod = "eth_signTransaction"
	PolicyUpdateRuleParamsMethodEthSignUserOperation     PolicyUpdateRuleParamsMethod = "eth_signUserOperation"
	PolicyUpdateRuleParamsMethodEthSignTypedDataV4       PolicyUpdateRuleParamsMethod = "eth_signTypedData_v4"
	PolicyUpdateRuleParamsMethodEthSign7702Authorization PolicyUpdateRuleParamsMethod = "eth_sign7702Authorization"
	PolicyUpdateRuleParamsMethodSignTransaction          PolicyUpdateRuleParamsMethod = "signTransaction"
	PolicyUpdateRuleParamsMethodSignAndSendTransaction   PolicyUpdateRuleParamsMethod = "signAndSendTransaction"
	PolicyUpdateRuleParamsMethodExportPrivateKey         PolicyUpdateRuleParamsMethod = "exportPrivateKey"
	PolicyUpdateRuleParamsMethodSignTransactionBytes     PolicyUpdateRuleParamsMethod = "signTransactionBytes"
	PolicyUpdateRuleParamsMethodStar                     PolicyUpdateRuleParamsMethod = "*"
)
