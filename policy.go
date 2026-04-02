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
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
	"github.com/privy-io/go-sdk/shared/constant"
)

// Operations related to policies
//
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
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a policy by policy ID.
func (r *PolicyService) Update(ctx context.Context, policyID string, params PolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete a policy by policy ID.
func (r *PolicyService) Delete(ctx context.Context, policyID string, body PolicyDeleteParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	if !param.IsOmitted(body.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", body.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(body.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", body.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Create a new rule for a policy.
func (r *PolicyService) NewRule(ctx context.Context, policyID string, params PolicyNewRuleParams, opts ...option.RequestOption) (res *PolicyRuleRequestBody, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/policies/%s/rules", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete a rule by policy ID and rule ID.
func (r *PolicyService) DeleteRule(ctx context.Context, ruleID string, params PolicyDeleteRuleParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
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
	path := fmt.Sprintf("v1/policies/%s/rules/%s", url.PathEscape(params.PolicyID), url.PathEscape(ruleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get a policy by policy ID.
func (r *PolicyService) Get(ctx context.Context, policyID string, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a rule by policy ID and rule ID.
func (r *PolicyService) GetRule(ctx context.Context, ruleID string, query PolicyGetRuleParams, opts ...option.RequestOption) (res *PolicyRuleRequestBody, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/policies/%s/rules/%s", url.PathEscape(query.PolicyID), url.PathEscape(ruleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a rule by policy ID and rule ID.
func (r *PolicyService) UpdateRule(ctx context.Context, ruleID string, params PolicyUpdateRuleParams, opts ...option.RequestOption) (res *PolicyRuleRequestBody, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
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
	path := fmt.Sprintf("v1/policies/%s/rules/%s", url.PathEscape(params.PolicyID), url.PathEscape(ruleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// The action to take when a policy rule matches.
type PolicyAction string

const (
	PolicyActionAllow PolicyAction = "ALLOW"
	PolicyActionDeny  PolicyAction = "DENY"
)

// A parameter in a Solidity ABI function or event definition.
type AbiParameterResp struct {
	Type         string `json:"type" api:"required"`
	Components   []any  `json:"components"`
	Indexed      bool   `json:"indexed"`
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		Components   respjson.Field
		Indexed      respjson.Field
		InternalType respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbiParameterResp) RawJSON() string { return r.JSON.raw }
func (r *AbiParameterResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AbiParameterResp to a AbiParameter.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AbiParameter.Overrides()
func (r AbiParameterResp) ToParam() AbiParameter {
	return param.Override[AbiParameter](json.RawMessage(r.RawJSON()))
}

// A parameter in a Solidity ABI function or event definition.
//
// The property Type is required.
type AbiParameter struct {
	Type         string            `json:"type" api:"required"`
	Indexed      param.Opt[bool]   `json:"indexed,omitzero"`
	InternalType param.Opt[string] `json:"internalType,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	Components   []any             `json:"components,omitzero"`
	paramObj
}

func (r AbiParameter) MarshalJSON() (data []byte, err error) {
	type shadow AbiParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbiParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AbiSchema []AbiSchemaItem

type AbiSchemaItem struct {
	// Any of "function", "constructor", "event", "fallback", "receive".
	Type      string             `json:"type" api:"required"`
	Anonymous bool               `json:"anonymous"`
	Inputs    []AbiParameterResp `json:"inputs"`
	Name      string             `json:"name"`
	Outputs   []AbiParameterResp `json:"outputs"`
	// Any of "pure", "view", "nonpayable", "payable".
	StateMutability string `json:"stateMutability"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Anonymous       respjson.Field
		Inputs          respjson.Field
		Name            respjson.Field
		Outputs         respjson.Field
		StateMutability respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbiSchemaItem) RawJSON() string { return r.JSON.raw }
func (r *AbiSchemaItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AbiSchemaParam []AbiSchemaItemParam

// The property Type is required.
type AbiSchemaItemParam struct {
	// Any of "function", "constructor", "event", "fallback", "receive".
	Type      string            `json:"type,omitzero" api:"required"`
	Anonymous param.Opt[bool]   `json:"anonymous,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	Inputs    []AbiParameter    `json:"inputs,omitzero"`
	Outputs   []AbiParameter    `json:"outputs,omitzero"`
	// Any of "pure", "view", "nonpayable", "payable".
	StateMutability string `json:"stateMutability,omitzero"`
	paramObj
}

func (r AbiSchemaItemParam) MarshalJSON() (data []byte, err error) {
	type shadow AbiSchemaItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbiSchemaItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AbiSchemaItemParam](
		"type", "function", "constructor", "event", "fallback", "receive",
	)
	apijson.RegisterFieldValidator[AbiSchemaItemParam](
		"stateMutability", "pure", "view", "nonpayable", "payable",
	)
}

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
	Field TronTransactionConditionField `json:"field" api:"required"`
	// Any of "tron_transaction".
	FieldSource TronTransactionConditionFieldSource `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator TronTransactionConditionOperator   `json:"operator" api:"required"`
	Value    TronTransactionConditionValueUnion `json:"value" api:"required"`
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
	Field TronTransactionConditionField `json:"field,omitzero" api:"required"`
	// Any of "tron_transaction".
	FieldSource TronTransactionConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator TronTransactionConditionOperator        `json:"operator,omitzero" api:"required"`
	Value    TronTransactionConditionValueUnionParam `json:"value,omitzero" api:"required"`
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

// Decoded calldata from a TRON TriggerSmartContract interaction.
type TronCalldataCondition struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchema `json:"abi" api:"required"`
	Field string    `json:"field" api:"required"`
	// Any of "tron_trigger_smart_contract_data".
	FieldSource TronCalldataConditionFieldSource `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator TronCalldataConditionOperator   `json:"operator" api:"required"`
	Value    TronCalldataConditionValueUnion `json:"value" api:"required"`
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
func (r TronCalldataCondition) RawJSON() string { return r.JSON.raw }
func (r *TronCalldataCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TronCalldataCondition to a TronCalldataConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TronCalldataConditionParam.Overrides()
func (r TronCalldataCondition) ToParam() TronCalldataConditionParam {
	return param.Override[TronCalldataConditionParam](json.RawMessage(r.RawJSON()))
}

type TronCalldataConditionFieldSource string

const (
	TronCalldataConditionFieldSourceTronTriggerSmartContractData TronCalldataConditionFieldSource = "tron_trigger_smart_contract_data"
)

type TronCalldataConditionOperator string

const (
	TronCalldataConditionOperatorEq             TronCalldataConditionOperator = "eq"
	TronCalldataConditionOperatorGt             TronCalldataConditionOperator = "gt"
	TronCalldataConditionOperatorGte            TronCalldataConditionOperator = "gte"
	TronCalldataConditionOperatorLt             TronCalldataConditionOperator = "lt"
	TronCalldataConditionOperatorLte            TronCalldataConditionOperator = "lte"
	TronCalldataConditionOperatorIn             TronCalldataConditionOperator = "in"
	TronCalldataConditionOperatorInConditionSet TronCalldataConditionOperator = "in_condition_set"
)

// TronCalldataConditionValueUnion contains all possible properties and values from
// [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type TronCalldataConditionValueUnion struct {
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

func (u TronCalldataConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TronCalldataConditionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TronCalldataConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TronCalldataConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Decoded calldata from a TRON TriggerSmartContract interaction.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type TronCalldataConditionParam struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchemaParam `json:"abi,omitzero" api:"required"`
	Field string         `json:"field" api:"required"`
	// Any of "tron_trigger_smart_contract_data".
	FieldSource TronCalldataConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator TronCalldataConditionOperator        `json:"operator,omitzero" api:"required"`
	Value    TronCalldataConditionValueUnionParam `json:"value,omitzero" api:"required"`
	paramObj
}

func (r TronCalldataConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow TronCalldataConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TronCalldataConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TronCalldataConditionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u TronCalldataConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *TronCalldataConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// SUI transaction command attributes, enables allowlisting specific command types.
// Allowed commands: 'TransferObjects', 'SplitCoins', 'MergeCoins'. Only 'eq' and
// 'in' operators are supported.
type SuiTransactionCommandCondition struct {
	// Any of "commandName".
	Field SuiTransactionCommandConditionField `json:"field" api:"required"`
	// Any of "sui_transaction_command".
	FieldSource SuiTransactionCommandConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for SUI transaction command conditions. Only 'eq' and 'in' are
	// supported for command names.
	//
	// Any of "eq", "in".
	Operator SuiTransactionCommandOperator `json:"operator" api:"required"`
	// Command name(s) to match. Must be one of: 'TransferObjects', 'SplitCoins',
	// 'MergeCoins'
	Value SuiTransactionCommandConditionValueUnion `json:"value" api:"required"`
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
	Field SuiTransactionCommandConditionField `json:"field,omitzero" api:"required"`
	// Any of "sui_transaction_command".
	FieldSource SuiTransactionCommandConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for SUI transaction command conditions. Only 'eq' and 'in' are
	// supported for command names.
	//
	// Any of "eq", "in".
	Operator SuiTransactionCommandOperator `json:"operator,omitzero" api:"required"`
	// Command name(s) to match. Must be one of: 'TransferObjects', 'SplitCoins',
	// 'MergeCoins'
	Value SuiTransactionCommandConditionValueUnionParam `json:"value,omitzero" api:"required"`
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

// SUI TransferObjects command attributes, including recipient and amount fields.
type SuiTransferObjectsCommandCondition struct {
	// Supported fields for SUI TransferObjects command conditions. Only 'recipient'
	// and 'amount' are supported.
	//
	// Any of "recipient", "amount".
	Field SuiTransferObjectsCommandField `json:"field" api:"required"`
	// Any of "sui_transfer_objects_command".
	FieldSource SuiTransferObjectsCommandConditionFieldSource `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator SuiTransferObjectsCommandConditionOperator   `json:"operator" api:"required"`
	Value    SuiTransferObjectsCommandConditionValueUnion `json:"value" api:"required"`
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
	Field SuiTransferObjectsCommandField `json:"field,omitzero" api:"required"`
	// Any of "sui_transfer_objects_command".
	FieldSource SuiTransferObjectsCommandConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator SuiTransferObjectsCommandConditionOperator        `json:"operator,omitzero" api:"required"`
	Value    SuiTransferObjectsCommandConditionValueUnionParam `json:"value,omitzero" api:"required"`
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

// Condition referencing an aggregation value. The field must start with
// "aggregation." followed by the aggregation ID.
type AggregationCondition struct {
	Field string `json:"field" api:"required"`
	// Any of "reference".
	FieldSource AggregationConditionFieldSource `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator AggregationConditionOperator   `json:"operator" api:"required"`
	Value    AggregationConditionValueUnion `json:"value" api:"required"`
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
func (r AggregationCondition) RawJSON() string { return r.JSON.raw }
func (r *AggregationCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AggregationCondition to a AggregationConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AggregationConditionParam.Overrides()
func (r AggregationCondition) ToParam() AggregationConditionParam {
	return param.Override[AggregationConditionParam](json.RawMessage(r.RawJSON()))
}

type AggregationConditionFieldSource string

const (
	AggregationConditionFieldSourceReference AggregationConditionFieldSource = "reference"
)

type AggregationConditionOperator string

const (
	AggregationConditionOperatorEq             AggregationConditionOperator = "eq"
	AggregationConditionOperatorGt             AggregationConditionOperator = "gt"
	AggregationConditionOperatorGte            AggregationConditionOperator = "gte"
	AggregationConditionOperatorLt             AggregationConditionOperator = "lt"
	AggregationConditionOperatorLte            AggregationConditionOperator = "lte"
	AggregationConditionOperatorIn             AggregationConditionOperator = "in"
	AggregationConditionOperatorInConditionSet AggregationConditionOperator = "in_condition_set"
)

// AggregationConditionValueUnion contains all possible properties and values from
// [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type AggregationConditionValueUnion struct {
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

func (u AggregationConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AggregationConditionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AggregationConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *AggregationConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition referencing an aggregation value. The field must start with
// "aggregation." followed by the aggregation ID.
//
// The properties Field, FieldSource, Operator, Value are required.
type AggregationConditionParam struct {
	Field string `json:"field" api:"required"`
	// Any of "reference".
	FieldSource AggregationConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator AggregationConditionOperator        `json:"operator,omitzero" api:"required"`
	Value    AggregationConditionValueUnionParam `json:"value,omitzero" api:"required"`
	paramObj
}

func (r AggregationConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow AggregationConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AggregationConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AggregationConditionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AggregationConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AggregationConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// PolicyConditionUnion contains all possible properties and values from
// [PolicyConditionEthereumTransaction], [PolicyConditionEthereumCalldata],
// [PolicyConditionEthereumTypedDataDomain],
// [PolicyConditionEthereumTypedDataMessage],
// [PolicyConditionEthereum7702Authorization],
// [PolicyConditionSolanaProgramInstruction],
// [PolicyConditionSolanaSystemProgramInstruction],
// [PolicyConditionSolanaTokenProgramInstruction], [PolicyConditionSystem],
// [TronTransactionCondition], [TronCalldataCondition],
// [SuiTransactionCommandCondition], [SuiTransferObjectsCommandCondition],
// [AggregationCondition].
//
// Use the [PolicyConditionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyConditionUnion struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "tron_trigger_smart_contract_data",
	// "sui_transaction_command", "sui_transfer_objects_command", "reference".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of [PolicyConditionEthereumTransactionValueUnion],
	// [PolicyConditionEthereumCalldataValueUnion],
	// [PolicyConditionEthereumTypedDataDomainValueUnion],
	// [PolicyConditionEthereumTypedDataMessageValueUnion],
	// [PolicyConditionEthereum7702AuthorizationValueUnion],
	// [PolicyConditionSolanaProgramInstructionValueUnion],
	// [PolicyConditionSolanaSystemProgramInstructionValueUnion],
	// [PolicyConditionSolanaTokenProgramInstructionValueUnion],
	// [PolicyConditionSystemValueUnion], [TronTransactionConditionValueUnion],
	// [TronCalldataConditionValueUnion], [SuiTransactionCommandConditionValueUnion],
	// [SuiTransferObjectsCommandConditionValueUnion], [AggregationConditionValueUnion]
	Value PolicyConditionUnionValue `json:"value"`
	// This field is from variant [PolicyConditionEthereumCalldata].
	Abi AbiSchema `json:"abi"`
	// This field is from variant [PolicyConditionEthereumTypedDataMessage].
	TypedData PolicyConditionEthereumTypedDataMessageTypedData `json:"typed_data"`
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

// anyPolicyCondition is implemented by each variant of [PolicyConditionUnion] to
// add type safety for the return type of [PolicyConditionUnion.AsAny]
type anyPolicyCondition interface {
	implPolicyConditionUnion()
}

func (PolicyConditionEthereumTransaction) implPolicyConditionUnion()            {}
func (PolicyConditionEthereumCalldata) implPolicyConditionUnion()               {}
func (PolicyConditionEthereumTypedDataDomain) implPolicyConditionUnion()        {}
func (PolicyConditionEthereumTypedDataMessage) implPolicyConditionUnion()       {}
func (PolicyConditionEthereum7702Authorization) implPolicyConditionUnion()      {}
func (PolicyConditionSolanaProgramInstruction) implPolicyConditionUnion()       {}
func (PolicyConditionSolanaSystemProgramInstruction) implPolicyConditionUnion() {}
func (PolicyConditionSolanaTokenProgramInstruction) implPolicyConditionUnion()  {}
func (PolicyConditionSystem) implPolicyConditionUnion()                         {}
func (TronTransactionCondition) implPolicyConditionUnion()                      {}
func (TronCalldataCondition) implPolicyConditionUnion()                         {}
func (SuiTransactionCommandCondition) implPolicyConditionUnion()                {}
func (SuiTransferObjectsCommandCondition) implPolicyConditionUnion()            {}
func (AggregationCondition) implPolicyConditionUnion()                          {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyConditionUnion.AsAny().(type) {
//	case privyclient.PolicyConditionEthereumTransaction:
//	case privyclient.PolicyConditionEthereumCalldata:
//	case privyclient.PolicyConditionEthereumTypedDataDomain:
//	case privyclient.PolicyConditionEthereumTypedDataMessage:
//	case privyclient.PolicyConditionEthereum7702Authorization:
//	case privyclient.PolicyConditionSolanaProgramInstruction:
//	case privyclient.PolicyConditionSolanaSystemProgramInstruction:
//	case privyclient.PolicyConditionSolanaTokenProgramInstruction:
//	case privyclient.PolicyConditionSystem:
//	case privyclient.TronTransactionCondition:
//	case privyclient.TronCalldataCondition:
//	case privyclient.SuiTransactionCommandCondition:
//	case privyclient.SuiTransferObjectsCommandCondition:
//	case privyclient.AggregationCondition:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyConditionUnion) AsAny() anyPolicyCondition {
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

func (u PolicyConditionUnion) AsEthereumTransaction() (v PolicyConditionEthereumTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsEthereumCalldata() (v PolicyConditionEthereumCalldata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsEthereumTypedDataDomain() (v PolicyConditionEthereumTypedDataDomain) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsEthereumTypedDataMessage() (v PolicyConditionEthereumTypedDataMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsEthereum7702Authorization() (v PolicyConditionEthereum7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsSolanaProgramInstruction() (v PolicyConditionSolanaProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsSolanaSystemProgramInstruction() (v PolicyConditionSolanaSystemProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsSolanaTokenProgramInstruction() (v PolicyConditionSolanaTokenProgramInstruction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsSystem() (v PolicyConditionSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsTronTransaction() (v TronTransactionCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsTronTriggerSmartContractData() (v TronCalldataCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsSuiTransactionCommand() (v SuiTransactionCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnion) AsReference() (v AggregationCondition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionUnionValue is an implicit subunion of [PolicyConditionUnion].
// PolicyConditionUnionValue provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyConditionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyConditionUnionValue struct {
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

func (r *PolicyConditionUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicyConditionUnion to a PolicyConditionUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyConditionUnionParam.Overrides()
func (r PolicyConditionUnion) ToParam() PolicyConditionUnionParam {
	return param.Override[PolicyConditionUnionParam](json.RawMessage(r.RawJSON()))
}

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type PolicyConditionEthereumTransaction struct {
	// Any of "to", "value", "chain_id".
	Field       string                       `json:"field" api:"required"`
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                       `json:"operator" api:"required"`
	Value    PolicyConditionEthereumTransactionValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionEthereumTransaction) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionEthereumTransactionValueUnion contains all possible properties
// and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionEthereumTransactionValueUnion struct {
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

func (u PolicyConditionEthereumTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionEthereumTransactionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionEthereumTransactionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionEthereumTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type PolicyConditionEthereumCalldata struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi         AbiSchema                 `json:"abi" api:"required"`
	Field       string                    `json:"field" api:"required"`
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                    `json:"operator" api:"required"`
	Value    PolicyConditionEthereumCalldataValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionEthereumCalldata) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionEthereumCalldata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionEthereumCalldataValueUnion contains all possible properties and
// values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionEthereumCalldataValueUnion struct {
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

func (u PolicyConditionEthereumCalldataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionEthereumCalldataValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionEthereumCalldataValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionEthereumCalldataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type PolicyConditionEthereumTypedDataDomain struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field       string                           `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                           `json:"operator" api:"required"`
	Value    PolicyConditionEthereumTypedDataDomainValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionEthereumTypedDataDomain) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionEthereumTypedDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionEthereumTypedDataDomainValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionEthereumTypedDataDomainValueUnion struct {
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

func (u PolicyConditionEthereumTypedDataDomainValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionEthereumTypedDataDomainValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionEthereumTypedDataDomainValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionEthereumTypedDataDomainValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type PolicyConditionEthereumTypedDataMessage struct {
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                            `json:"operator" api:"required"`
	TypedData PolicyConditionEthereumTypedDataMessageTypedData  `json:"typed_data" api:"required"`
	Value     PolicyConditionEthereumTypedDataMessageValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionEthereumTypedDataMessage) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionEthereumTypedDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyConditionEthereumTypedDataMessageTypedData struct {
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
func (r PolicyConditionEthereumTypedDataMessageTypedData) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionEthereumTypedDataMessageValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionEthereumTypedDataMessageValueUnion struct {
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

func (u PolicyConditionEthereumTypedDataMessageValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionEthereumTypedDataMessageValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionEthereumTypedDataMessageValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionEthereumTypedDataMessageValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
type PolicyConditionEthereum7702Authorization struct {
	// Any of "contract".
	Field       string                             `json:"field" api:"required"`
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                             `json:"operator" api:"required"`
	Value    PolicyConditionEthereum7702AuthorizationValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionEthereum7702Authorization) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionEthereum7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionEthereum7702AuthorizationValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionEthereum7702AuthorizationValueUnion struct {
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

func (u PolicyConditionEthereum7702AuthorizationValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionEthereum7702AuthorizationValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionEthereum7702AuthorizationValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionEthereum7702AuthorizationValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type PolicyConditionSolanaProgramInstruction struct {
	// Any of "programId".
	Field       string                            `json:"field" api:"required"`
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                            `json:"operator" api:"required"`
	Value    PolicyConditionSolanaProgramInstructionValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionSolanaProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionSolanaProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionSolanaProgramInstructionValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionSolanaProgramInstructionValueUnion struct {
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

func (u PolicyConditionSolanaProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionSolanaProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionSolanaProgramInstructionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionSolanaProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type PolicyConditionSolanaSystemProgramInstruction struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field       string                                  `json:"field" api:"required"`
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                  `json:"operator" api:"required"`
	Value    PolicyConditionSolanaSystemProgramInstructionValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionSolanaSystemProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionSolanaSystemProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionSolanaSystemProgramInstructionValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionSolanaSystemProgramInstructionValueUnion struct {
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

func (u PolicyConditionSolanaSystemProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionSolanaSystemProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionSolanaSystemProgramInstructionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionSolanaSystemProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type PolicyConditionSolanaTokenProgramInstruction struct {
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
	Operator string                                                 `json:"operator" api:"required"`
	Value    PolicyConditionSolanaTokenProgramInstructionValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionSolanaTokenProgramInstruction) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionSolanaTokenProgramInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionSolanaTokenProgramInstructionValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionSolanaTokenProgramInstructionValueUnion struct {
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

func (u PolicyConditionSolanaTokenProgramInstructionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionSolanaTokenProgramInstructionValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionSolanaTokenProgramInstructionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionSolanaTokenProgramInstructionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type PolicyConditionSystem struct {
	// Any of "current_unix_timestamp".
	Field       string          `json:"field" api:"required"`
	FieldSource constant.System `json:"field_source" default:"system"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                          `json:"operator" api:"required"`
	Value    PolicyConditionSystemValueUnion `json:"value" api:"required"`
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
func (r PolicyConditionSystem) RawJSON() string { return r.JSON.raw }
func (r *PolicyConditionSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionSystemValueUnion contains all possible properties and values from
// [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyConditionSystemValueUnion struct {
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

func (u PolicyConditionSystemValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionSystemValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionSystemValueUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionSystemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func PolicyConditionParamOfEthereumTransaction[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var ethereumTransaction PolicyConditionEthereumTransactionParam
	ethereumTransaction.Field = field
	ethereumTransaction.Operator = operator
	switch v := any(value).(type) {
	case string:
		ethereumTransaction.Value.OfString = param.NewOpt(v)
	case []string:
		ethereumTransaction.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfEthereumTransaction: &ethereumTransaction}
}

func PolicyConditionParamOfEthereumTypedDataDomain[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var ethereumTypedDataDomain PolicyConditionEthereumTypedDataDomainParam
	ethereumTypedDataDomain.Field = field
	ethereumTypedDataDomain.Operator = operator
	switch v := any(value).(type) {
	case string:
		ethereumTypedDataDomain.Value.OfString = param.NewOpt(v)
	case []string:
		ethereumTypedDataDomain.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfEthereumTypedDataDomain: &ethereumTypedDataDomain}
}

func PolicyConditionParamOfEthereum7702Authorization[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var ethereum7702Authorization PolicyConditionEthereum7702AuthorizationParam
	ethereum7702Authorization.Field = field
	ethereum7702Authorization.Operator = operator
	switch v := any(value).(type) {
	case string:
		ethereum7702Authorization.Value.OfString = param.NewOpt(v)
	case []string:
		ethereum7702Authorization.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfEthereum7702Authorization: &ethereum7702Authorization}
}

func PolicyConditionParamOfSolanaProgramInstruction[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var solanaProgramInstruction PolicyConditionSolanaProgramInstructionParam
	solanaProgramInstruction.Field = field
	solanaProgramInstruction.Operator = operator
	switch v := any(value).(type) {
	case string:
		solanaProgramInstruction.Value.OfString = param.NewOpt(v)
	case []string:
		solanaProgramInstruction.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfSolanaProgramInstruction: &solanaProgramInstruction}
}

func PolicyConditionParamOfSolanaSystemProgramInstruction[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var solanaSystemProgramInstruction PolicyConditionSolanaSystemProgramInstructionParam
	solanaSystemProgramInstruction.Field = field
	solanaSystemProgramInstruction.Operator = operator
	switch v := any(value).(type) {
	case string:
		solanaSystemProgramInstruction.Value.OfString = param.NewOpt(v)
	case []string:
		solanaSystemProgramInstruction.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfSolanaSystemProgramInstruction: &solanaSystemProgramInstruction}
}

func PolicyConditionParamOfSolanaTokenProgramInstruction[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var solanaTokenProgramInstruction PolicyConditionSolanaTokenProgramInstructionParam
	solanaTokenProgramInstruction.Field = field
	solanaTokenProgramInstruction.Operator = operator
	switch v := any(value).(type) {
	case string:
		solanaTokenProgramInstruction.Value.OfString = param.NewOpt(v)
	case []string:
		solanaTokenProgramInstruction.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfSolanaTokenProgramInstruction: &solanaTokenProgramInstruction}
}

func PolicyConditionParamOfSystem[T string | []string](field string, operator string, value T) PolicyConditionUnionParam {
	var system PolicyConditionSystemParam
	system.Field = field
	system.Operator = operator
	switch v := any(value).(type) {
	case string:
		system.Value.OfString = param.NewOpt(v)
	case []string:
		system.Value.OfStringArray = v
	}
	return PolicyConditionUnionParam{OfSystem: &system}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionUnionParam struct {
	OfEthereumTransaction            *PolicyConditionEthereumTransactionParam            `json:",omitzero,inline"`
	OfEthereumCalldata               *PolicyConditionEthereumCalldataParam               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *PolicyConditionEthereumTypedDataDomainParam        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *PolicyConditionEthereumTypedDataMessageParam       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *PolicyConditionEthereum7702AuthorizationParam      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *PolicyConditionSolanaProgramInstructionParam       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *PolicyConditionSolanaSystemProgramInstructionParam `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *PolicyConditionSolanaTokenProgramInstructionParam  `json:",omitzero,inline"`
	OfSystem                         *PolicyConditionSystemParam                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionConditionParam                      `json:",omitzero,inline"`
	OfTronTriggerSmartContractData   *TronCalldataConditionParam                         `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam            `json:",omitzero,inline"`
	OfReference                      *AggregationConditionParam                          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionUnionParam) MarshalJSON() ([]byte, error) {
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
func (u *PolicyConditionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PolicyConditionUnionParam](
		"field_source",
		apijson.Discriminator[PolicyConditionEthereumTransactionParam]("ethereum_transaction"),
		apijson.Discriminator[PolicyConditionEthereumCalldataParam]("ethereum_calldata"),
		apijson.Discriminator[PolicyConditionEthereumTypedDataDomainParam]("ethereum_typed_data_domain"),
		apijson.Discriminator[PolicyConditionEthereumTypedDataMessageParam]("ethereum_typed_data_message"),
		apijson.Discriminator[PolicyConditionEthereum7702AuthorizationParam]("ethereum_7702_authorization"),
		apijson.Discriminator[PolicyConditionSolanaProgramInstructionParam]("solana_program_instruction"),
		apijson.Discriminator[PolicyConditionSolanaSystemProgramInstructionParam]("solana_system_program_instruction"),
		apijson.Discriminator[PolicyConditionSolanaTokenProgramInstructionParam]("solana_token_program_instruction"),
		apijson.Discriminator[PolicyConditionSystemParam]("system"),
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
type PolicyConditionEthereumTransactionParam struct {
	// Any of "to", "value", "chain_id".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                            `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionEthereumTransactionValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_transaction".
	FieldSource constant.EthereumTransaction `json:"field_source" default:"ethereum_transaction"`
	paramObj
}

func (r PolicyConditionEthereumTransactionParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionEthereumTransactionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionEthereumTransactionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionEthereumTransactionParam](
		"field", "to", "value", "chain_id",
	)
	apijson.RegisterFieldValidator[PolicyConditionEthereumTransactionParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionEthereumTransactionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionEthereumTransactionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionEthereumTransactionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type PolicyConditionEthereumCalldataParam struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchemaParam `json:"abi,omitzero" api:"required"`
	Field string         `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                         `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionEthereumCalldataValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_calldata".
	FieldSource constant.EthereumCalldata `json:"field_source" default:"ethereum_calldata"`
	paramObj
}

func (r PolicyConditionEthereumCalldataParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionEthereumCalldataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionEthereumCalldataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionEthereumCalldataParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionEthereumCalldataValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionEthereumCalldataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionEthereumCalldataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyConditionEthereumTypedDataDomainParam struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionEthereumTypedDataDomainValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" default:"ethereum_typed_data_domain"`
	paramObj
}

func (r PolicyConditionEthereumTypedDataDomainParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionEthereumTypedDataDomainParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionEthereumTypedDataDomainParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionEthereumTypedDataDomainParam](
		"field", "chainId", "verifyingContract", "chain_id", "verifying_contract",
	)
	apijson.RegisterFieldValidator[PolicyConditionEthereumTypedDataDomainParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionEthereumTypedDataDomainValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionEthereumTypedDataDomainValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionEthereumTypedDataDomainValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type PolicyConditionEthereumTypedDataMessageParam struct {
	Field string `json:"field" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  string                                                 `json:"operator,omitzero" api:"required"`
	TypedData PolicyConditionEthereumTypedDataMessageTypedDataParam  `json:"typed_data,omitzero" api:"required"`
	Value     PolicyConditionEthereumTypedDataMessageValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_message".
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" default:"ethereum_typed_data_message"`
	paramObj
}

func (r PolicyConditionEthereumTypedDataMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionEthereumTypedDataMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionEthereumTypedDataMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionEthereumTypedDataMessageParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// The properties PrimaryType, Types are required.
type PolicyConditionEthereumTypedDataMessageTypedDataParam struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParams `json:"types,omitzero" api:"required"`
	paramObj
}

func (r PolicyConditionEthereumTypedDataMessageTypedDataParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionEthereumTypedDataMessageTypedDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionEthereumTypedDataMessageTypedDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionEthereumTypedDataMessageValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionEthereumTypedDataMessageValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionEthereumTypedDataMessageValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyConditionEthereum7702AuthorizationParam struct {
	// Any of "contract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                  `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionEthereum7702AuthorizationValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_7702_authorization".
	FieldSource constant.Ethereum7702Authorization `json:"field_source" default:"ethereum_7702_authorization"`
	paramObj
}

func (r PolicyConditionEthereum7702AuthorizationParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionEthereum7702AuthorizationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionEthereum7702AuthorizationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionEthereum7702AuthorizationParam](
		"field", "contract",
	)
	apijson.RegisterFieldValidator[PolicyConditionEthereum7702AuthorizationParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionEthereum7702AuthorizationValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionEthereum7702AuthorizationValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionEthereum7702AuthorizationValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyConditionSolanaProgramInstructionParam struct {
	// Any of "programId".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                 `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionSolanaProgramInstructionValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_program_instruction".
	FieldSource constant.SolanaProgramInstruction `json:"field_source" default:"solana_program_instruction"`
	paramObj
}

func (r PolicyConditionSolanaProgramInstructionParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionSolanaProgramInstructionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionSolanaProgramInstructionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionSolanaProgramInstructionParam](
		"field", "programId",
	)
	apijson.RegisterFieldValidator[PolicyConditionSolanaProgramInstructionParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionSolanaProgramInstructionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionSolanaProgramInstructionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionSolanaProgramInstructionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyConditionSolanaSystemProgramInstructionParam struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                       `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionSolanaSystemProgramInstructionValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_system_program_instruction".
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" default:"solana_system_program_instruction"`
	paramObj
}

func (r PolicyConditionSolanaSystemProgramInstructionParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionSolanaSystemProgramInstructionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionSolanaSystemProgramInstructionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionSolanaSystemProgramInstructionParam](
		"field", "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports",
	)
	apijson.RegisterFieldValidator[PolicyConditionSolanaSystemProgramInstructionParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionSolanaSystemProgramInstructionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionSolanaSystemProgramInstructionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionSolanaSystemProgramInstructionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyConditionSolanaTokenProgramInstructionParam struct {
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
	Operator string                                                      `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionSolanaTokenProgramInstructionValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" default:"solana_token_program_instruction"`
	paramObj
}

func (r PolicyConditionSolanaTokenProgramInstructionParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionSolanaTokenProgramInstructionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionSolanaTokenProgramInstructionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionSolanaTokenProgramInstructionParam](
		"field", "instructionName", "Transfer.source", "Transfer.destination", "Transfer.authority", "Transfer.amount", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint", "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account", "MintTo.authority", "MintTo.amount", "CloseAccount.account", "CloseAccount.destination", "CloseAccount.authority", "InitializeAccount3.account", "InitializeAccount3.mint", "InitializeAccount3.owner",
	)
	apijson.RegisterFieldValidator[PolicyConditionSolanaTokenProgramInstructionParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionSolanaTokenProgramInstructionValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionSolanaTokenProgramInstructionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionSolanaTokenProgramInstructionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type PolicyConditionSystemParam struct {
	// Any of "current_unix_timestamp".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                               `json:"operator,omitzero" api:"required"`
	Value    PolicyConditionSystemValueUnionParam `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "system".
	FieldSource constant.System `json:"field_source" default:"system"`
	paramObj
}

func (r PolicyConditionSystemParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyConditionSystemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyConditionSystemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicyConditionSystemParam](
		"field", "current_unix_timestamp",
	)
	apijson.RegisterFieldValidator[PolicyConditionSystemParam](
		"operator", "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionSystemValueUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionSystemValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *PolicyConditionSystemValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Method the rule applies to.
type PolicyMethod string

const (
	PolicyMethodEthSendTransaction       PolicyMethod = "eth_sendTransaction"
	PolicyMethodEthSignTransaction       PolicyMethod = "eth_signTransaction"
	PolicyMethodEthSignUserOperation     PolicyMethod = "eth_signUserOperation"
	PolicyMethodEthSignTypedDataV4       PolicyMethod = "eth_signTypedData_v4"
	PolicyMethodEthSign7702Authorization PolicyMethod = "eth_sign7702Authorization"
	PolicyMethodSignTransaction          PolicyMethod = "signTransaction"
	PolicyMethodSignAndSendTransaction   PolicyMethod = "signAndSendTransaction"
	PolicyMethodExportPrivateKey         PolicyMethod = "exportPrivateKey"
	PolicyMethodExportSeedPhrase         PolicyMethod = "exportSeedPhrase"
	PolicyMethodSignTransactionBytes     PolicyMethod = "signTransactionBytes"
	PolicyMethodStar                     PolicyMethod = "*"
)

// The rules that apply to each method the policy covers.
type PolicyRuleRequestBody struct {
	// The action to take when a policy rule matches.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyAction           `json:"action" api:"required"`
	Conditions []PolicyConditionUnion `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "exportSeedPhrase",
	// "signTransactionBytes", "\*".
	Method PolicyMethod `json:"method" api:"required"`
	Name   string       `json:"name" api:"required"`
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
func (r PolicyRuleRequestBody) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicyRuleRequestBody to a PolicyRuleRequestBodyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyRuleRequestBodyParam.Overrides()
func (r PolicyRuleRequestBody) ToParam() PolicyRuleRequestBodyParam {
	return param.Override[PolicyRuleRequestBodyParam](json.RawMessage(r.RawJSON()))
}

// The rules that apply to each method the policy covers.
//
// The properties Action, Conditions, Method, Name are required.
type PolicyRuleRequestBodyParam struct {
	// The action to take when a policy rule matches.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyAction                `json:"action,omitzero" api:"required"`
	Conditions []PolicyConditionUnionParam `json:"conditions,omitzero" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "exportPrivateKey", "exportSeedPhrase",
	// "signTransactionBytes", "\*".
	Method PolicyMethod `json:"method,omitzero" api:"required"`
	Name   string       `json:"name" api:"required"`
	paramObj
}

func (r PolicyRuleRequestBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyRuleRequestBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyRuleRequestBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A policy for controlling wallet operations.
type Policy struct {
	// Unique ID of the created policy. This will be the primary identifier when using
	// the policy in the future.
	ID string `json:"id" api:"required"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type" api:"required"`
	// Unix timestamp of when the policy was created in milliseconds.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Name to assign to policy.
	Name string `json:"name" api:"required"`
	// The key quorum ID of the owner of the policy.
	OwnerID string                  `json:"owner_id" api:"required"`
	Rules   []PolicyRuleRequestBody `json:"rules" api:"required"`
	// Version of the policy. Currently, 1.0 is the only version.
	//
	// Any of "1.0".
	Version PolicyVersion `json:"version" api:"required"`
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

// Version of the policy. Currently, 1.0 is the only version.
type PolicyVersion string

const (
	PolicyVersion1_0 PolicyVersion = "1.0"
)

type PolicyNewParams struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero" api:"required"`
	// Name to assign to policy.
	Name  string                       `json:"name" api:"required"`
	Rules []PolicyRuleRequestBodyParam `json:"rules,omitzero" api:"required"`
	// Version of the policy. Currently, 1.0 is the only version.
	//
	// Any of "1.0".
	Version PolicyNewParamsVersion `json:"version,omitzero" api:"required"`
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

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type PolicyNewParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
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
	UserID string `json:"user_id" api:"required"`
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
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner PolicyUpdateParamsOwnerUnion `json:"owner,omitzero"`
	Rules []PolicyRuleRequestBodyParam `json:"rules,omitzero"`
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

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type PolicyUpdateParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
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
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r PolicyUpdateParamsOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParamsOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParamsOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyDeleteParams struct {
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

type PolicyNewRuleParams struct {
	// The rules that apply to each method the policy covers.
	PolicyRuleRequestBody PolicyRuleRequestBodyParam
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r PolicyNewRuleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PolicyRuleRequestBody)
}
func (r *PolicyNewRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyDeleteRuleParams struct {
	PolicyID string `path:"policy_id" api:"required" json:"-"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

type PolicyGetRuleParams struct {
	PolicyID string `path:"policy_id" api:"required" json:"-"`
	paramObj
}

type PolicyUpdateRuleParams struct {
	PolicyID string `path:"policy_id" api:"required" json:"-"`
	// The rules that apply to each method the policy covers.
	PolicyRuleRequestBody PolicyRuleRequestBodyParam
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r PolicyUpdateRuleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PolicyRuleRequestBody)
}
func (r *PolicyUpdateRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
