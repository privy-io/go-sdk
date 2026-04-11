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
func (r *PolicyService) NewRule(ctx context.Context, policyID string, params PolicyNewRuleParams, opts ...option.RequestOption) (res *PolicyRuleResponse, err error) {
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
func (r *PolicyService) GetRule(ctx context.Context, ruleID string, query PolicyGetRuleParams, opts ...option.RequestOption) (res *PolicyRuleResponse, err error) {
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
func (r *PolicyService) UpdateRule(ctx context.Context, ruleID string, params PolicyUpdateRuleParams, opts ...option.RequestOption) (res *PolicyRuleResponse, err error) {
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

type AbiSchemaResp []AbiSchemaItemResp

type AbiSchemaItemResp struct {
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
func (r AbiSchemaItemResp) RawJSON() string { return r.JSON.raw }
func (r *AbiSchemaItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AbiSchema []AbiSchemaItem

// The property Type is required.
type AbiSchemaItem struct {
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

func (r AbiSchemaItem) MarshalJSON() (data []byte, err error) {
	type shadow AbiSchemaItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbiSchemaItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AbiSchemaItem](
		"type", "function", "constructor", "event", "fallback", "receive",
	)
	apijson.RegisterFieldValidator[AbiSchemaItem](
		"stateMutability", "pure", "view", "nonpayable", "payable",
	)
}

// Operator to use for policy conditions.
type ConditionOperator string

const (
	ConditionOperatorEq             ConditionOperator = "eq"
	ConditionOperatorGt             ConditionOperator = "gt"
	ConditionOperatorGte            ConditionOperator = "gte"
	ConditionOperatorLt             ConditionOperator = "lt"
	ConditionOperatorLte            ConditionOperator = "lte"
	ConditionOperatorIn             ConditionOperator = "in"
	ConditionOperatorInConditionSet ConditionOperator = "in_condition_set"
)

// ConditionValueUnionResp contains all possible properties and values from
// [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type ConditionValueUnionResp struct {
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

func (u ConditionValueUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionValueUnionResp) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConditionValueUnionResp) RawJSON() string { return u.JSON.raw }

func (r *ConditionValueUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConditionValueUnionResp to a ConditionValueUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConditionValueUnion.Overrides()
func (r ConditionValueUnionResp) ToParam() ConditionValueUnion {
	return param.Override[ConditionValueUnion](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConditionValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ConditionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
type EthereumTransactionConditionResp struct {
	// Any of "to", "value", "chain_id".
	Field EthereumTransactionConditionField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource EthereumTransactionConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r EthereumTransactionConditionResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumTransactionConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumTransactionConditionResp to a
// EthereumTransactionCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumTransactionCondition.Overrides()
func (r EthereumTransactionConditionResp) ToParam() EthereumTransactionCondition {
	return param.Override[EthereumTransactionCondition](json.RawMessage(r.RawJSON()))
}

type EthereumTransactionConditionField string

const (
	EthereumTransactionConditionFieldTo      EthereumTransactionConditionField = "to"
	EthereumTransactionConditionFieldValue   EthereumTransactionConditionField = "value"
	EthereumTransactionConditionFieldChainID EthereumTransactionConditionField = "chain_id"
)

type EthereumTransactionConditionFieldSource string

const (
	EthereumTransactionConditionFieldSourceEthereumTransaction EthereumTransactionConditionFieldSource = "ethereum_transaction"
)

// The verbatim Ethereum transaction object in an eth_signTransaction or
// eth_sendTransaction request.
//
// The properties Field, FieldSource, Operator, Value are required.
type EthereumTransactionCondition struct {
	// Any of "to", "value", "chain_id".
	Field EthereumTransactionConditionField `json:"field,omitzero" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource EthereumTransactionConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r EthereumTransactionCondition) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTransactionCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTransactionCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
type EthereumCalldataConditionResp struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchemaResp `json:"abi" api:"required"`
	Field string        `json:"field" api:"required"`
	// Any of "ethereum_calldata".
	FieldSource EthereumCalldataConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r EthereumCalldataConditionResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumCalldataConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumCalldataConditionResp to a
// EthereumCalldataCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumCalldataCondition.Overrides()
func (r EthereumCalldataConditionResp) ToParam() EthereumCalldataCondition {
	return param.Override[EthereumCalldataCondition](json.RawMessage(r.RawJSON()))
}

type EthereumCalldataConditionFieldSource string

const (
	EthereumCalldataConditionFieldSourceEthereumCalldata EthereumCalldataConditionFieldSource = "ethereum_calldata"
)

// The decoded calldata in a smart contract interaction as the smart contract
// method's parameters. Note that that 'ethereum_calldata' conditions must contain
// an abi parameter with the JSON ABI of the smart contract.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type EthereumCalldataCondition struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchema `json:"abi,omitzero" api:"required"`
	Field string    `json:"field" api:"required"`
	// Any of "ethereum_calldata".
	FieldSource EthereumCalldataConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r EthereumCalldataCondition) MarshalJSON() (data []byte, err error) {
	type shadow EthereumCalldataCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumCalldataCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes from the signing domain that will verify the signature.
type EthereumTypedDataDomainConditionResp struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field EthereumTypedDataDomainConditionField `json:"field" api:"required"`
	// Any of "ethereum_typed_data_domain".
	FieldSource EthereumTypedDataDomainConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r EthereumTypedDataDomainConditionResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumTypedDataDomainConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumTypedDataDomainConditionResp to a
// EthereumTypedDataDomainCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumTypedDataDomainCondition.Overrides()
func (r EthereumTypedDataDomainConditionResp) ToParam() EthereumTypedDataDomainCondition {
	return param.Override[EthereumTypedDataDomainCondition](json.RawMessage(r.RawJSON()))
}

type EthereumTypedDataDomainConditionField string

const (
	EthereumTypedDataDomainConditionFieldChainIDMixedCase           EthereumTypedDataDomainConditionField = "chainId"
	EthereumTypedDataDomainConditionFieldVerifyingContractCamelCase EthereumTypedDataDomainConditionField = "verifyingContract"
	EthereumTypedDataDomainConditionFieldChainID                    EthereumTypedDataDomainConditionField = "chain_id"
	EthereumTypedDataDomainConditionFieldVerifyingContract          EthereumTypedDataDomainConditionField = "verifying_contract"
)

type EthereumTypedDataDomainConditionFieldSource string

const (
	EthereumTypedDataDomainConditionFieldSourceEthereumTypedDataDomain EthereumTypedDataDomainConditionFieldSource = "ethereum_typed_data_domain"
)

// Attributes from the signing domain that will verify the signature.
//
// The properties Field, FieldSource, Operator, Value are required.
type EthereumTypedDataDomainCondition struct {
	// Any of "chainId", "verifyingContract", "chain_id", "verifying_contract".
	Field EthereumTypedDataDomainConditionField `json:"field,omitzero" api:"required"`
	// Any of "ethereum_typed_data_domain".
	FieldSource EthereumTypedDataDomainConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r EthereumTypedDataDomainCondition) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTypedDataDomainCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTypedDataDomainCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
type EthereumTypedDataMessageConditionResp struct {
	Field string `json:"field" api:"required"`
	// Any of "ethereum_typed_data_message".
	FieldSource EthereumTypedDataMessageConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  ConditionOperator                              `json:"operator" api:"required"`
	TypedData EthereumTypedDataMessageConditionTypedDataResp `json:"typed_data" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r EthereumTypedDataMessageConditionResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumTypedDataMessageConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumTypedDataMessageConditionResp to a
// EthereumTypedDataMessageCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumTypedDataMessageCondition.Overrides()
func (r EthereumTypedDataMessageConditionResp) ToParam() EthereumTypedDataMessageCondition {
	return param.Override[EthereumTypedDataMessageCondition](json.RawMessage(r.RawJSON()))
}

type EthereumTypedDataMessageConditionFieldSource string

const (
	EthereumTypedDataMessageConditionFieldSourceEthereumTypedDataMessage EthereumTypedDataMessageConditionFieldSource = "ethereum_typed_data_message"
)

type EthereumTypedDataMessageConditionTypedDataResp struct {
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
func (r EthereumTypedDataMessageConditionTypedDataResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumTypedDataMessageConditionTypedDataResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'types' and 'primary_type' attributes of the TypedData JSON object defined in
// EIP-712.
//
// The properties Field, FieldSource, Operator, TypedData, Value are required.
type EthereumTypedDataMessageCondition struct {
	Field string `json:"field" api:"required"`
	// Any of "ethereum_typed_data_message".
	FieldSource EthereumTypedDataMessageConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator  ConditionOperator                          `json:"operator,omitzero" api:"required"`
	TypedData EthereumTypedDataMessageConditionTypedData `json:"typed_data,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r EthereumTypedDataMessageCondition) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTypedDataMessageCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTypedDataMessageCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PrimaryType, Types are required.
type EthereumTypedDataMessageConditionTypedData struct {
	PrimaryType string `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParams `json:"types,omitzero" api:"required"`
	paramObj
}

func (r EthereumTypedDataMessageConditionTypedData) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTypedDataMessageConditionTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTypedDataMessageConditionTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed contract addresses for eth_sign7702Authorization requests.
type Ethereum7702AuthorizationConditionResp struct {
	// Any of "contract".
	Field Ethereum7702AuthorizationConditionField `json:"field" api:"required"`
	// Any of "ethereum_7702_authorization".
	FieldSource Ethereum7702AuthorizationConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r Ethereum7702AuthorizationConditionResp) RawJSON() string { return r.JSON.raw }
func (r *Ethereum7702AuthorizationConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Ethereum7702AuthorizationConditionResp to a
// Ethereum7702AuthorizationCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// Ethereum7702AuthorizationCondition.Overrides()
func (r Ethereum7702AuthorizationConditionResp) ToParam() Ethereum7702AuthorizationCondition {
	return param.Override[Ethereum7702AuthorizationCondition](json.RawMessage(r.RawJSON()))
}

type Ethereum7702AuthorizationConditionField string

const (
	Ethereum7702AuthorizationConditionFieldContract Ethereum7702AuthorizationConditionField = "contract"
)

type Ethereum7702AuthorizationConditionFieldSource string

const (
	Ethereum7702AuthorizationConditionFieldSourceEthereum7702Authorization Ethereum7702AuthorizationConditionFieldSource = "ethereum_7702_authorization"
)

// Allowed contract addresses for eth_sign7702Authorization requests.
//
// The properties Field, FieldSource, Operator, Value are required.
type Ethereum7702AuthorizationCondition struct {
	// Any of "contract".
	Field Ethereum7702AuthorizationConditionField `json:"field,omitzero" api:"required"`
	// Any of "ethereum_7702_authorization".
	FieldSource Ethereum7702AuthorizationConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r Ethereum7702AuthorizationCondition) MarshalJSON() (data []byte, err error) {
	type shadow Ethereum7702AuthorizationCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Ethereum7702AuthorizationCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Program attributes, enables allowlisting Solana Programs.
type SolanaProgramInstructionConditionResp struct {
	// Any of "programId".
	Field SolanaProgramInstructionConditionField `json:"field" api:"required"`
	// Any of "solana_program_instruction".
	FieldSource SolanaProgramInstructionConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r SolanaProgramInstructionConditionResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaProgramInstructionConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaProgramInstructionConditionResp to a
// SolanaProgramInstructionCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaProgramInstructionCondition.Overrides()
func (r SolanaProgramInstructionConditionResp) ToParam() SolanaProgramInstructionCondition {
	return param.Override[SolanaProgramInstructionCondition](json.RawMessage(r.RawJSON()))
}

type SolanaProgramInstructionConditionField string

const (
	SolanaProgramInstructionConditionFieldProgramID SolanaProgramInstructionConditionField = "programId"
)

type SolanaProgramInstructionConditionFieldSource string

const (
	SolanaProgramInstructionConditionFieldSourceSolanaProgramInstruction SolanaProgramInstructionConditionFieldSource = "solana_program_instruction"
)

// Solana Program attributes, enables allowlisting Solana Programs.
//
// The properties Field, FieldSource, Operator, Value are required.
type SolanaProgramInstructionCondition struct {
	// Any of "programId".
	Field SolanaProgramInstructionConditionField `json:"field,omitzero" api:"required"`
	// Any of "solana_program_instruction".
	FieldSource SolanaProgramInstructionConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r SolanaProgramInstructionCondition) MarshalJSON() (data []byte, err error) {
	type shadow SolanaProgramInstructionCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaProgramInstructionCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana System Program attributes, including more granular Transfer instruction
// fields.
type SolanaSystemProgramInstructionConditionResp struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field SolanaSystemProgramInstructionConditionField `json:"field" api:"required"`
	// Any of "solana_system_program_instruction".
	FieldSource SolanaSystemProgramInstructionConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r SolanaSystemProgramInstructionConditionResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSystemProgramInstructionConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSystemProgramInstructionConditionResp to a
// SolanaSystemProgramInstructionCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSystemProgramInstructionCondition.Overrides()
func (r SolanaSystemProgramInstructionConditionResp) ToParam() SolanaSystemProgramInstructionCondition {
	return param.Override[SolanaSystemProgramInstructionCondition](json.RawMessage(r.RawJSON()))
}

type SolanaSystemProgramInstructionConditionField string

const (
	SolanaSystemProgramInstructionConditionFieldInstructionName  SolanaSystemProgramInstructionConditionField = "instructionName"
	SolanaSystemProgramInstructionConditionFieldTransferFrom     SolanaSystemProgramInstructionConditionField = "Transfer.from"
	SolanaSystemProgramInstructionConditionFieldTransferTo       SolanaSystemProgramInstructionConditionField = "Transfer.to"
	SolanaSystemProgramInstructionConditionFieldTransferLamports SolanaSystemProgramInstructionConditionField = "Transfer.lamports"
)

type SolanaSystemProgramInstructionConditionFieldSource string

const (
	SolanaSystemProgramInstructionConditionFieldSourceSolanaSystemProgramInstruction SolanaSystemProgramInstructionConditionFieldSource = "solana_system_program_instruction"
)

// Solana System Program attributes, including more granular Transfer instruction
// fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type SolanaSystemProgramInstructionCondition struct {
	// Any of "instructionName", "Transfer.from", "Transfer.to", "Transfer.lamports".
	Field SolanaSystemProgramInstructionConditionField `json:"field,omitzero" api:"required"`
	// Any of "solana_system_program_instruction".
	FieldSource SolanaSystemProgramInstructionConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r SolanaSystemProgramInstructionCondition) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSystemProgramInstructionCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSystemProgramInstructionCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
type SolanaTokenProgramInstructionConditionResp struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field SolanaTokenProgramInstructionConditionField `json:"field" api:"required"`
	// Any of "solana_token_program_instruction".
	FieldSource SolanaTokenProgramInstructionConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r SolanaTokenProgramInstructionConditionResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaTokenProgramInstructionConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaTokenProgramInstructionConditionResp to a
// SolanaTokenProgramInstructionCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaTokenProgramInstructionCondition.Overrides()
func (r SolanaTokenProgramInstructionConditionResp) ToParam() SolanaTokenProgramInstructionCondition {
	return param.Override[SolanaTokenProgramInstructionCondition](json.RawMessage(r.RawJSON()))
}

type SolanaTokenProgramInstructionConditionField string

const (
	SolanaTokenProgramInstructionConditionFieldInstructionName            SolanaTokenProgramInstructionConditionField = "instructionName"
	SolanaTokenProgramInstructionConditionFieldTransferSource             SolanaTokenProgramInstructionConditionField = "Transfer.source"
	SolanaTokenProgramInstructionConditionFieldTransferDestination        SolanaTokenProgramInstructionConditionField = "Transfer.destination"
	SolanaTokenProgramInstructionConditionFieldTransferAuthority          SolanaTokenProgramInstructionConditionField = "Transfer.authority"
	SolanaTokenProgramInstructionConditionFieldTransferAmount             SolanaTokenProgramInstructionConditionField = "Transfer.amount"
	SolanaTokenProgramInstructionConditionFieldTransferCheckedSource      SolanaTokenProgramInstructionConditionField = "TransferChecked.source"
	SolanaTokenProgramInstructionConditionFieldTransferCheckedDestination SolanaTokenProgramInstructionConditionField = "TransferChecked.destination"
	SolanaTokenProgramInstructionConditionFieldTransferCheckedAuthority   SolanaTokenProgramInstructionConditionField = "TransferChecked.authority"
	SolanaTokenProgramInstructionConditionFieldTransferCheckedAmount      SolanaTokenProgramInstructionConditionField = "TransferChecked.amount"
	SolanaTokenProgramInstructionConditionFieldTransferCheckedMint        SolanaTokenProgramInstructionConditionField = "TransferChecked.mint"
	SolanaTokenProgramInstructionConditionFieldBurnAccount                SolanaTokenProgramInstructionConditionField = "Burn.account"
	SolanaTokenProgramInstructionConditionFieldBurnMint                   SolanaTokenProgramInstructionConditionField = "Burn.mint"
	SolanaTokenProgramInstructionConditionFieldBurnAuthority              SolanaTokenProgramInstructionConditionField = "Burn.authority"
	SolanaTokenProgramInstructionConditionFieldBurnAmount                 SolanaTokenProgramInstructionConditionField = "Burn.amount"
	SolanaTokenProgramInstructionConditionFieldMintToMint                 SolanaTokenProgramInstructionConditionField = "MintTo.mint"
	SolanaTokenProgramInstructionConditionFieldMintToAccount              SolanaTokenProgramInstructionConditionField = "MintTo.account"
	SolanaTokenProgramInstructionConditionFieldMintToAuthority            SolanaTokenProgramInstructionConditionField = "MintTo.authority"
	SolanaTokenProgramInstructionConditionFieldMintToAmount               SolanaTokenProgramInstructionConditionField = "MintTo.amount"
	SolanaTokenProgramInstructionConditionFieldCloseAccountAccount        SolanaTokenProgramInstructionConditionField = "CloseAccount.account"
	SolanaTokenProgramInstructionConditionFieldCloseAccountDestination    SolanaTokenProgramInstructionConditionField = "CloseAccount.destination"
	SolanaTokenProgramInstructionConditionFieldCloseAccountAuthority      SolanaTokenProgramInstructionConditionField = "CloseAccount.authority"
	SolanaTokenProgramInstructionConditionFieldInitializeAccount3Account  SolanaTokenProgramInstructionConditionField = "InitializeAccount3.account"
	SolanaTokenProgramInstructionConditionFieldInitializeAccount3Mint     SolanaTokenProgramInstructionConditionField = "InitializeAccount3.mint"
	SolanaTokenProgramInstructionConditionFieldInitializeAccount3Owner    SolanaTokenProgramInstructionConditionField = "InitializeAccount3.owner"
)

type SolanaTokenProgramInstructionConditionFieldSource string

const (
	SolanaTokenProgramInstructionConditionFieldSourceSolanaTokenProgramInstruction SolanaTokenProgramInstructionConditionFieldSource = "solana_token_program_instruction"
)

// Solana Token Program attributes, including more granular TransferChecked
// instruction fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type SolanaTokenProgramInstructionCondition struct {
	// Any of "instructionName", "Transfer.source", "Transfer.destination",
	// "Transfer.authority", "Transfer.amount", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint", "Burn.account", "Burn.mint",
	// "Burn.authority", "Burn.amount", "MintTo.mint", "MintTo.account",
	// "MintTo.authority", "MintTo.amount", "CloseAccount.account",
	// "CloseAccount.destination", "CloseAccount.authority",
	// "InitializeAccount3.account", "InitializeAccount3.mint",
	// "InitializeAccount3.owner".
	Field SolanaTokenProgramInstructionConditionField `json:"field,omitzero" api:"required"`
	// Any of "solana_token_program_instruction".
	FieldSource SolanaTokenProgramInstructionConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r SolanaTokenProgramInstructionCondition) MarshalJSON() (data []byte, err error) {
	type shadow SolanaTokenProgramInstructionCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaTokenProgramInstructionCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System attributes, including current unix timestamp (in seconds).
type SystemConditionResp struct {
	// Any of "current_unix_timestamp".
	Field SystemConditionField `json:"field" api:"required"`
	// Any of "system".
	FieldSource SystemConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r SystemConditionResp) RawJSON() string { return r.JSON.raw }
func (r *SystemConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SystemConditionResp to a SystemCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SystemCondition.Overrides()
func (r SystemConditionResp) ToParam() SystemCondition {
	return param.Override[SystemCondition](json.RawMessage(r.RawJSON()))
}

type SystemConditionField string

const (
	SystemConditionFieldCurrentUnixTimestamp SystemConditionField = "current_unix_timestamp"
)

type SystemConditionFieldSource string

const (
	SystemConditionFieldSourceSystem SystemConditionFieldSource = "system"
)

// System attributes, including current unix timestamp (in seconds).
//
// The properties Field, FieldSource, Operator, Value are required.
type SystemCondition struct {
	// Any of "current_unix_timestamp".
	Field SystemConditionField `json:"field,omitzero" api:"required"`
	// Any of "system".
	FieldSource SystemConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r SystemCondition) MarshalJSON() (data []byte, err error) {
	type shadow SystemCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TRON transaction fields for TransferContract and TriggerSmartContract
// transaction types.
type TronTransactionConditionResp struct {
	// Supported TRON transaction fields in format "TransactionType.field_name"
	//
	// Any of "TransferContract.to_address", "TransferContract.amount",
	// "TriggerSmartContract.contract_address", "TriggerSmartContract.call_value",
	// "TriggerSmartContract.token_id", "TriggerSmartContract.call_token_value".
	Field TronTransactionConditionField `json:"field" api:"required"`
	// Any of "tron_transaction".
	FieldSource TronTransactionConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r TronTransactionConditionResp) RawJSON() string { return r.JSON.raw }
func (r *TronTransactionConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TronTransactionConditionResp to a
// TronTransactionCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TronTransactionCondition.Overrides()
func (r TronTransactionConditionResp) ToParam() TronTransactionCondition {
	return param.Override[TronTransactionCondition](json.RawMessage(r.RawJSON()))
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

// TRON transaction fields for TransferContract and TriggerSmartContract
// transaction types.
//
// The properties Field, FieldSource, Operator, Value are required.
type TronTransactionCondition struct {
	// Supported TRON transaction fields in format "TransactionType.field_name"
	//
	// Any of "TransferContract.to_address", "TransferContract.amount",
	// "TriggerSmartContract.contract_address", "TriggerSmartContract.call_value",
	// "TriggerSmartContract.token_id", "TriggerSmartContract.call_token_value".
	Field TronTransactionConditionField `json:"field,omitzero" api:"required"`
	// Any of "tron_transaction".
	FieldSource TronTransactionConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r TronTransactionCondition) MarshalJSON() (data []byte, err error) {
	type shadow TronTransactionCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TronTransactionCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Decoded calldata from a TRON TriggerSmartContract interaction.
type TronCalldataConditionResp struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchemaResp `json:"abi" api:"required"`
	Field string        `json:"field" api:"required"`
	// Any of "tron_trigger_smart_contract_data".
	FieldSource TronCalldataConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r TronCalldataConditionResp) RawJSON() string { return r.JSON.raw }
func (r *TronCalldataConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TronCalldataConditionResp to a TronCalldataCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TronCalldataCondition.Overrides()
func (r TronCalldataConditionResp) ToParam() TronCalldataCondition {
	return param.Override[TronCalldataCondition](json.RawMessage(r.RawJSON()))
}

type TronCalldataConditionFieldSource string

const (
	TronCalldataConditionFieldSourceTronTriggerSmartContractData TronCalldataConditionFieldSource = "tron_trigger_smart_contract_data"
)

// Decoded calldata from a TRON TriggerSmartContract interaction.
//
// The properties Abi, Field, FieldSource, Operator, Value are required.
type TronCalldataCondition struct {
	// A Solidity ABI definition for decoding smart contract calldata.
	Abi   AbiSchema `json:"abi,omitzero" api:"required"`
	Field string    `json:"field" api:"required"`
	// Any of "tron_trigger_smart_contract_data".
	FieldSource TronCalldataConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r TronCalldataCondition) MarshalJSON() (data []byte, err error) {
	type shadow TronCalldataCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TronCalldataCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SUI transaction command attributes, enables allowlisting specific command types.
// Allowed commands: 'TransferObjects', 'SplitCoins', 'MergeCoins'. Only 'eq' and
// 'in' operators are supported.
type SuiTransactionCommandConditionResp struct {
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
	Value SuiTransactionCommandConditionValueUnionResp `json:"value" api:"required"`
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
func (r SuiTransactionCommandConditionResp) RawJSON() string { return r.JSON.raw }
func (r *SuiTransactionCommandConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SuiTransactionCommandConditionResp to a
// SuiTransactionCommandCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SuiTransactionCommandCondition.Overrides()
func (r SuiTransactionCommandConditionResp) ToParam() SuiTransactionCommandCondition {
	return param.Override[SuiTransactionCommandCondition](json.RawMessage(r.RawJSON()))
}

type SuiTransactionCommandConditionField string

const (
	SuiTransactionCommandConditionFieldCommandName SuiTransactionCommandConditionField = "commandName"
)

type SuiTransactionCommandConditionFieldSource string

const (
	SuiTransactionCommandConditionFieldSourceSuiTransactionCommand SuiTransactionCommandConditionFieldSource = "sui_transaction_command"
)

// SuiTransactionCommandConditionValueUnionResp contains all possible properties
// and values from [SuiCommandName], [[]SuiCommandName].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSuiCommandName OfSuiCommandNameArray]
type SuiTransactionCommandConditionValueUnionResp struct {
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

func (u SuiTransactionCommandConditionValueUnionResp) AsSuiCommandName() (v SuiCommandName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SuiTransactionCommandConditionValueUnionResp) AsSuiCommandNameArray() (v []SuiCommandName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SuiTransactionCommandConditionValueUnionResp) RawJSON() string { return u.JSON.raw }

func (r *SuiTransactionCommandConditionValueUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SUI transaction command attributes, enables allowlisting specific command types.
// Allowed commands: 'TransferObjects', 'SplitCoins', 'MergeCoins'. Only 'eq' and
// 'in' operators are supported.
//
// The properties Field, FieldSource, Operator, Value are required.
type SuiTransactionCommandCondition struct {
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
	Value SuiTransactionCommandConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r SuiTransactionCommandCondition) MarshalJSON() (data []byte, err error) {
	type shadow SuiTransactionCommandCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuiTransactionCommandCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SuiTransactionCommandConditionValueUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfSuiCommandName)
	OfSuiCommandName      param.Opt[SuiCommandName] `json:",omitzero,inline"`
	OfSuiCommandNameArray []SuiCommandName          `json:",omitzero,inline"`
	paramUnion
}

func (u SuiTransactionCommandConditionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSuiCommandName, u.OfSuiCommandNameArray)
}
func (u *SuiTransactionCommandConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// SUI TransferObjects command attributes, including recipient and amount fields.
type SuiTransferObjectsCommandConditionResp struct {
	// Supported fields for SUI TransferObjects command conditions. Only 'recipient'
	// and 'amount' are supported.
	//
	// Any of "recipient", "amount".
	Field SuiTransferObjectsCommandField `json:"field" api:"required"`
	// Any of "sui_transfer_objects_command".
	FieldSource SuiTransferObjectsCommandConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r SuiTransferObjectsCommandConditionResp) RawJSON() string { return r.JSON.raw }
func (r *SuiTransferObjectsCommandConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SuiTransferObjectsCommandConditionResp to a
// SuiTransferObjectsCommandCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SuiTransferObjectsCommandCondition.Overrides()
func (r SuiTransferObjectsCommandConditionResp) ToParam() SuiTransferObjectsCommandCondition {
	return param.Override[SuiTransferObjectsCommandCondition](json.RawMessage(r.RawJSON()))
}

type SuiTransferObjectsCommandConditionFieldSource string

const (
	SuiTransferObjectsCommandConditionFieldSourceSuiTransferObjectsCommand SuiTransferObjectsCommandConditionFieldSource = "sui_transfer_objects_command"
)

// SUI TransferObjects command attributes, including recipient and amount fields.
//
// The properties Field, FieldSource, Operator, Value are required.
type SuiTransferObjectsCommandCondition struct {
	// Supported fields for SUI TransferObjects command conditions. Only 'recipient'
	// and 'amount' are supported.
	//
	// Any of "recipient", "amount".
	Field SuiTransferObjectsCommandField `json:"field,omitzero" api:"required"`
	// Any of "sui_transfer_objects_command".
	FieldSource SuiTransferObjectsCommandConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r SuiTransferObjectsCommandCondition) MarshalJSON() (data []byte, err error) {
	type shadow SuiTransferObjectsCommandCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuiTransferObjectsCommandCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition referencing an aggregation value. The field must start with
// "aggregation." followed by the aggregation ID.
type AggregationConditionResp struct {
	Field string `json:"field" api:"required"`
	// Any of "reference".
	FieldSource AggregationConditionFieldSource `json:"field_source" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnionResp `json:"value" api:"required"`
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
func (r AggregationConditionResp) RawJSON() string { return r.JSON.raw }
func (r *AggregationConditionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AggregationConditionResp to a AggregationCondition.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AggregationCondition.Overrides()
func (r AggregationConditionResp) ToParam() AggregationCondition {
	return param.Override[AggregationCondition](json.RawMessage(r.RawJSON()))
}

type AggregationConditionFieldSource string

const (
	AggregationConditionFieldSourceReference AggregationConditionFieldSource = "reference"
)

// Condition referencing an aggregation value. The field must start with
// "aggregation." followed by the aggregation ID.
//
// The properties Field, FieldSource, Operator, Value are required.
type AggregationCondition struct {
	Field string `json:"field" api:"required"`
	// Any of "reference".
	FieldSource AggregationConditionFieldSource `json:"field_source,omitzero" api:"required"`
	// Operator to use for policy conditions.
	//
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// Value to compare against in a policy condition. Can be a single string or an
	// array of strings.
	Value ConditionValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r AggregationCondition) MarshalJSON() (data []byte, err error) {
	type shadow AggregationCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AggregationCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionUnionResp contains all possible properties and values from
// [EthereumTransactionConditionResp], [EthereumCalldataConditionResp],
// [EthereumTypedDataDomainConditionResp], [EthereumTypedDataMessageConditionResp],
// [Ethereum7702AuthorizationConditionResp],
// [SolanaProgramInstructionConditionResp],
// [SolanaSystemProgramInstructionConditionResp],
// [SolanaTokenProgramInstructionConditionResp], [SystemConditionResp],
// [TronTransactionConditionResp], [TronCalldataConditionResp],
// [SuiTransactionCommandConditionResp], [SuiTransferObjectsCommandConditionResp],
// [AggregationConditionResp].
//
// Use the [PolicyConditionUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyConditionUnionResp struct {
	Field string `json:"field"`
	// Any of "ethereum_transaction", "ethereum_calldata",
	// "ethereum_typed_data_domain", "ethereum_typed_data_message",
	// "ethereum_7702_authorization", "solana_program_instruction",
	// "solana_system_program_instruction", "solana_token_program_instruction",
	// "system", "tron_transaction", "tron_trigger_smart_contract_data",
	// "sui_transaction_command", "sui_transfer_objects_command", "reference".
	FieldSource string `json:"field_source"`
	Operator    string `json:"operator"`
	// This field is a union of [ConditionValueUnionResp],
	// [SuiTransactionCommandConditionValueUnionResp]
	Value PolicyConditionUnionRespValue `json:"value"`
	// This field is from variant [EthereumCalldataConditionResp].
	Abi AbiSchemaResp `json:"abi"`
	// This field is from variant [EthereumTypedDataMessageConditionResp].
	TypedData EthereumTypedDataMessageConditionTypedDataResp `json:"typed_data"`
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

// anyPolicyConditionResp is implemented by each variant of
// [PolicyConditionUnionResp] to add type safety for the return type of
// [PolicyConditionUnionResp.AsAny]
type anyPolicyConditionResp interface {
	implPolicyConditionUnionResp()
}

func (EthereumTransactionConditionResp) implPolicyConditionUnionResp()            {}
func (EthereumCalldataConditionResp) implPolicyConditionUnionResp()               {}
func (EthereumTypedDataDomainConditionResp) implPolicyConditionUnionResp()        {}
func (EthereumTypedDataMessageConditionResp) implPolicyConditionUnionResp()       {}
func (Ethereum7702AuthorizationConditionResp) implPolicyConditionUnionResp()      {}
func (SolanaProgramInstructionConditionResp) implPolicyConditionUnionResp()       {}
func (SolanaSystemProgramInstructionConditionResp) implPolicyConditionUnionResp() {}
func (SolanaTokenProgramInstructionConditionResp) implPolicyConditionUnionResp()  {}
func (SystemConditionResp) implPolicyConditionUnionResp()                         {}
func (TronTransactionConditionResp) implPolicyConditionUnionResp()                {}
func (TronCalldataConditionResp) implPolicyConditionUnionResp()                   {}
func (SuiTransactionCommandConditionResp) implPolicyConditionUnionResp()          {}
func (SuiTransferObjectsCommandConditionResp) implPolicyConditionUnionResp()      {}
func (AggregationConditionResp) implPolicyConditionUnionResp()                    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PolicyConditionUnionResp.AsAny().(type) {
//	case privyclient.EthereumTransactionConditionResp:
//	case privyclient.EthereumCalldataConditionResp:
//	case privyclient.EthereumTypedDataDomainConditionResp:
//	case privyclient.EthereumTypedDataMessageConditionResp:
//	case privyclient.Ethereum7702AuthorizationConditionResp:
//	case privyclient.SolanaProgramInstructionConditionResp:
//	case privyclient.SolanaSystemProgramInstructionConditionResp:
//	case privyclient.SolanaTokenProgramInstructionConditionResp:
//	case privyclient.SystemConditionResp:
//	case privyclient.TronTransactionConditionResp:
//	case privyclient.TronCalldataConditionResp:
//	case privyclient.SuiTransactionCommandConditionResp:
//	case privyclient.SuiTransferObjectsCommandConditionResp:
//	case privyclient.AggregationConditionResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PolicyConditionUnionResp) AsAny() anyPolicyConditionResp {
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

func (u PolicyConditionUnionResp) AsEthereumTransaction() (v EthereumTransactionConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsEthereumCalldata() (v EthereumCalldataConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsEthereumTypedDataDomain() (v EthereumTypedDataDomainConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsEthereumTypedDataMessage() (v EthereumTypedDataMessageConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsEthereum7702Authorization() (v Ethereum7702AuthorizationConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsSolanaProgramInstruction() (v SolanaProgramInstructionConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsSolanaSystemProgramInstruction() (v SolanaSystemProgramInstructionConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsSolanaTokenProgramInstruction() (v SolanaTokenProgramInstructionConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsSystem() (v SystemConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsTronTransaction() (v TronTransactionConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsTronTriggerSmartContractData() (v TronCalldataConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsSuiTransactionCommand() (v SuiTransactionCommandConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsSuiTransferObjectsCommand() (v SuiTransferObjectsCommandConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyConditionUnionResp) AsReference() (v AggregationConditionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyConditionUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PolicyConditionUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyConditionUnionRespValue is an implicit subunion of
// [PolicyConditionUnionResp]. PolicyConditionUnionRespValue provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PolicyConditionUnionResp].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray OfSuiCommandName OfSuiCommandNameArray]
type PolicyConditionUnionRespValue struct {
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

func (r *PolicyConditionUnionRespValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicyConditionUnionResp to a PolicyConditionUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyConditionUnion.Overrides()
func (r PolicyConditionUnionResp) ToParam() PolicyConditionUnion {
	return param.Override[PolicyConditionUnion](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolicyConditionUnion struct {
	OfEthereumTransaction            *EthereumTransactionCondition            `json:",omitzero,inline"`
	OfEthereumCalldata               *EthereumCalldataCondition               `json:",omitzero,inline"`
	OfEthereumTypedDataDomain        *EthereumTypedDataDomainCondition        `json:",omitzero,inline"`
	OfEthereumTypedDataMessage       *EthereumTypedDataMessageCondition       `json:",omitzero,inline"`
	OfEthereum7702Authorization      *Ethereum7702AuthorizationCondition      `json:",omitzero,inline"`
	OfSolanaProgramInstruction       *SolanaProgramInstructionCondition       `json:",omitzero,inline"`
	OfSolanaSystemProgramInstruction *SolanaSystemProgramInstructionCondition `json:",omitzero,inline"`
	OfSolanaTokenProgramInstruction  *SolanaTokenProgramInstructionCondition  `json:",omitzero,inline"`
	OfSystem                         *SystemCondition                         `json:",omitzero,inline"`
	OfTronTransaction                *TronTransactionCondition                `json:",omitzero,inline"`
	OfTronTriggerSmartContractData   *TronCalldataCondition                   `json:",omitzero,inline"`
	OfSuiTransactionCommand          *SuiTransactionCommandCondition          `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandCondition      `json:",omitzero,inline"`
	OfReference                      *AggregationCondition                    `json:",omitzero,inline"`
	paramUnion
}

func (u PolicyConditionUnion) MarshalJSON() ([]byte, error) {
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
func (u *PolicyConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PolicyConditionUnion](
		"field_source",
		apijson.Discriminator[EthereumTransactionCondition]("ethereum_transaction"),
		apijson.Discriminator[EthereumCalldataCondition]("ethereum_calldata"),
		apijson.Discriminator[EthereumTypedDataDomainCondition]("ethereum_typed_data_domain"),
		apijson.Discriminator[EthereumTypedDataMessageCondition]("ethereum_typed_data_message"),
		apijson.Discriminator[Ethereum7702AuthorizationCondition]("ethereum_7702_authorization"),
		apijson.Discriminator[SolanaProgramInstructionCondition]("solana_program_instruction"),
		apijson.Discriminator[SolanaSystemProgramInstructionCondition]("solana_system_program_instruction"),
		apijson.Discriminator[SolanaTokenProgramInstructionCondition]("solana_token_program_instruction"),
		apijson.Discriminator[SystemCondition]("system"),
		apijson.Discriminator[TronTransactionCondition]("tron_transaction"),
		apijson.Discriminator[TronCalldataCondition]("tron_trigger_smart_contract_data"),
		apijson.Discriminator[SuiTransactionCommandCondition]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandCondition]("sui_transfer_objects_command"),
		apijson.Discriminator[AggregationCondition]("reference"),
	)
}

// Method the rule applies to.
type PolicyMethod string

const (
	PolicyMethodEthSendTransaction       PolicyMethod = "eth_sendTransaction"
	PolicyMethodEthSignTransaction       PolicyMethod = "eth_signTransaction"
	PolicyMethodEthSignUserOperation     PolicyMethod = "eth_signUserOperation"
	PolicyMethodEthSignTypedDataV4       PolicyMethod = "eth_signTypedData_v4"
	PolicyMethodEthSign7702Authorization PolicyMethod = "eth_sign7702Authorization"
	PolicyMethodWalletSendCalls          PolicyMethod = "wallet_sendCalls"
	PolicyMethodSignTransaction          PolicyMethod = "signTransaction"
	PolicyMethodSignAndSendTransaction   PolicyMethod = "signAndSendTransaction"
	PolicyMethodExportPrivateKey         PolicyMethod = "exportPrivateKey"
	PolicyMethodExportSeedPhrase         PolicyMethod = "exportSeedPhrase"
	PolicyMethodSignTransactionBytes     PolicyMethod = "signTransactionBytes"
	PolicyMethodStar                     PolicyMethod = "*"
)

// The rules that apply to each method the policy covers.
type PolicyRuleRequestBodyResp struct {
	// The action to take when a policy rule matches.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyAction               `json:"action" api:"required"`
	Conditions []PolicyConditionUnionResp `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "wallet_sendCalls",
	// "signTransaction", "signAndSendTransaction", "exportPrivateKey",
	// "exportSeedPhrase", "signTransactionBytes", "\*".
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
func (r PolicyRuleRequestBodyResp) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleRequestBodyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicyRuleRequestBodyResp to a PolicyRuleRequestBody.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyRuleRequestBody.Overrides()
func (r PolicyRuleRequestBodyResp) ToParam() PolicyRuleRequestBody {
	return param.Override[PolicyRuleRequestBody](json.RawMessage(r.RawJSON()))
}

// The rules that apply to each method the policy covers.
//
// The properties Action, Conditions, Method, Name are required.
type PolicyRuleRequestBody struct {
	// The action to take when a policy rule matches.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyAction           `json:"action,omitzero" api:"required"`
	Conditions []PolicyConditionUnion `json:"conditions,omitzero" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "wallet_sendCalls",
	// "signTransaction", "signAndSendTransaction", "exportPrivateKey",
	// "exportSeedPhrase", "signTransactionBytes", "\*".
	Method PolicyMethod `json:"method,omitzero" api:"required"`
	Name   string       `json:"name" api:"required"`
	paramObj
}

func (r PolicyRuleRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow PolicyRuleRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyRuleRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rule that defines the conditions and action to take if the conditions are
// true.
type PolicyRuleResponse struct {
	ID string `json:"id" api:"required"`
	// The action to take when a policy rule matches.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyAction               `json:"action" api:"required"`
	Conditions []PolicyConditionUnionResp `json:"conditions" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "wallet_sendCalls",
	// "signTransaction", "signAndSendTransaction", "exportPrivateKey",
	// "exportSeedPhrase", "signTransactionBytes", "\*".
	Method PolicyMethod `json:"method" api:"required"`
	Name   string       `json:"name" api:"required"`
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
func (r PolicyRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyRuleResponse) UnmarshalJSON(data []byte) error {
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
	// "tron", "bitcoin-segwit", "bitcoin-taproot", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type" api:"required"`
	// Unix timestamp of when the policy was created in milliseconds.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Name to assign to policy.
	Name string `json:"name" api:"required"`
	// The key quorum ID of the owner of the policy.
	OwnerID string               `json:"owner_id" api:"required"`
	Rules   []PolicyRuleResponse `json:"rules" api:"required"`
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
	// "tron", "bitcoin-segwit", "bitcoin-taproot", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero" api:"required"`
	// Name to assign to policy.
	Name  string                `json:"name" api:"required"`
	Rules []PolicyNewParamsRule `json:"rules,omitzero" api:"required"`
	// Version of the policy. Currently, 1.0 is the only version.
	//
	// Any of "1.0".
	Version PolicyNewParamsVersion `json:"version,omitzero" api:"required"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner,omitzero"`
	paramObj
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, Conditions, Method, Name are required.
type PolicyNewParamsRule struct {
	// The action to take when a policy rule matches.
	//
	// Any of "ALLOW", "DENY".
	Action     PolicyAction           `json:"action,omitzero" api:"required"`
	Conditions []PolicyConditionUnion `json:"conditions,omitzero" api:"required"`
	// Method the rule applies to.
	//
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signUserOperation",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "wallet_sendCalls",
	// "signTransaction", "signAndSendTransaction", "exportPrivateKey",
	// "exportSeedPhrase", "signTransactionBytes", "\*".
	Method PolicyMethod      `json:"method,omitzero" api:"required"`
	Name   string            `json:"name" api:"required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r PolicyNewParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow PolicyNewParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyNewParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version of the policy. Currently, 1.0 is the only version.
type PolicyNewParamsVersion string

const (
	PolicyNewParamsVersion1_0 PolicyNewParamsVersion = "1.0"
)

type PolicyUpdateParams struct {
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// Name to assign to policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion         `json:"owner,omitzero"`
	Rules []PolicyRuleRequestBody `json:"rules,omitzero"`
	paramObj
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyUpdateParams) UnmarshalJSON(data []byte) error {
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
	PolicyRuleRequestBody PolicyRuleRequestBody
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
	PolicyRuleRequestBody PolicyRuleRequestBody
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
