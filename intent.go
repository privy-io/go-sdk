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
func (r *IntentService) NewPolicyRule(ctx context.Context, policyID string, body IntentNewPolicyRuleParams, opts ...option.RequestOption) (res *RuleIntentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s/rules", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create an intent to delete a rule from a policy. The intent must be authorized
// by the policy owner before it can be executed.
func (r *IntentService) DeletePolicyRule(ctx context.Context, ruleID string, body IntentDeletePolicyRuleParams, opts ...option.RequestOption) (res *RuleIntentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PolicyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s/rules/%s", url.PathEscape(body.PolicyID), url.PathEscape(ruleID))
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
func (r *IntentService) Rpc(ctx context.Context, walletID string, body IntentRpcParams, opts ...option.RequestOption) (res *RpcIntentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/wallets/%s/rpc", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create an intent to update a key quorum. The intent must be authorized by the
// key quorum members before it can be executed.
func (r *IntentService) UpdateKeyQuorum(ctx context.Context, keyQuorumID string, body IntentUpdateKeyQuorumParams, opts ...option.RequestOption) (res *KeyQuorumIntentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/key_quorums/%s", url.PathEscape(keyQuorumID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Create an intent to update a policy. The intent must be authorized by the policy
// owner before it can be executed.
func (r *IntentService) UpdatePolicy(ctx context.Context, policyID string, body IntentUpdatePolicyParams, opts ...option.RequestOption) (res *PolicyIntentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Create an intent to update a rule on a policy. The intent must be authorized by
// the policy owner before it can be executed.
func (r *IntentService) UpdatePolicyRule(ctx context.Context, ruleID string, params IntentUpdatePolicyRuleParams, opts ...option.RequestOption) (res *RuleIntentResponse, err error) {
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
func (r *IntentService) UpdateWallet(ctx context.Context, walletID string, body IntentUpdateWalletParams, opts ...option.RequestOption) (res *WalletIntentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/intents/wallets/%s", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Type of intent.
type IntentType string

const (
	IntentTypeKeyQuorum IntentType = "KEY_QUORUM"
	IntentTypePolicy    IntentType = "POLICY"
	IntentTypeRule      IntentType = "RULE"
	IntentTypeRpc       IntentType = "RPC"
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

type RuleIntentCreateRequestDetailsBody struct {
	// Any of "ALLOW", "DENY".
	Action     string                                             `json:"action" api:"required"`
	Conditions []RuleIntentCreateRequestDetailsBodyConditionUnion `json:"conditions" api:"required"`
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signTypedData_v4",
	// "eth_signUserOperation", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "signTransactionBytes", "exportPrivateKey", "\*".
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
// properties and values from [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject],
// [RuleIntentCreateRequestDetailsBodyConditionObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentCreateRequestDetailsBodyConditionUnion struct {
	// This field is from variant [RuleIntentCreateRequestDetailsBodyConditionObject].
	Field RuleIntentCreateRequestDetailsBodyConditionObjectField `json:"field"`
	// This field is from variant [RuleIntentCreateRequestDetailsBodyConditionObject].
	FieldSource string `json:"field_source"`
	// This field is from variant [RuleIntentCreateRequestDetailsBodyConditionObject].
	Operator RuleIntentCreateRequestDetailsBodyConditionObjectOperator `json:"operator"`
	// This field is from variant [RuleIntentCreateRequestDetailsBodyConditionObject].
	Value RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion `json:"value"`
	// This field is from variant [RuleIntentCreateRequestDetailsBodyConditionObject].
	Abi []RuleIntentCreateRequestDetailsBodyConditionObjectAbi `json:"abi"`
	// This field is from variant [RuleIntentCreateRequestDetailsBodyConditionObject].
	TypedData RuleIntentCreateRequestDetailsBodyConditionObjectTypedData `json:"typed_data"`
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

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsRuleIntentCreateRequestDetailsBodyConditionObject() (v RuleIntentCreateRequestDetailsBodyConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsVariant2() (v RuleIntentCreateRequestDetailsBodyConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionUnion) AsVariant3() (v RuleIntentCreateRequestDetailsBodyConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentCreateRequestDetailsBodyConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentCreateRequestDetailsBodyConditionObject struct {
	// Any of "to", "value", "chain_id".
	Field RuleIntentCreateRequestDetailsBodyConditionObjectField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource string `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator RuleIntentCreateRequestDetailsBodyConditionObjectOperator   `json:"operator" api:"required"`
	Value    RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion `json:"value" api:"required"`
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
func (r RuleIntentCreateRequestDetailsBodyConditionObject) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentCreateRequestDetailsBodyConditionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentCreateRequestDetailsBodyConditionObjectField string

const (
	RuleIntentCreateRequestDetailsBodyConditionObjectFieldTo      RuleIntentCreateRequestDetailsBodyConditionObjectField = "to"
	RuleIntentCreateRequestDetailsBodyConditionObjectFieldValue   RuleIntentCreateRequestDetailsBodyConditionObjectField = "value"
	RuleIntentCreateRequestDetailsBodyConditionObjectFieldChainID RuleIntentCreateRequestDetailsBodyConditionObjectField = "chain_id"
)

type RuleIntentCreateRequestDetailsBodyConditionObjectOperator string

const (
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorEq             RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "eq"
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorGt             RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "gt"
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorGte            RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "gte"
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorLt             RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "lt"
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorLte            RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "lte"
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorIn             RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "in"
	RuleIntentCreateRequestDetailsBodyConditionObjectOperatorInConditionSet RuleIntentCreateRequestDetailsBodyConditionObjectOperator = "in_condition_set"
)

// RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion struct {
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

func (u RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentCreateRequestDetailsBodyConditionObjectValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentCreateRequestDetailsBodyConditionOperator string

const (
	RuleIntentCreateRequestDetailsBodyConditionOperatorEq             RuleIntentCreateRequestDetailsBodyConditionOperator = "eq"
	RuleIntentCreateRequestDetailsBodyConditionOperatorGt             RuleIntentCreateRequestDetailsBodyConditionOperator = "gt"
	RuleIntentCreateRequestDetailsBodyConditionOperatorGte            RuleIntentCreateRequestDetailsBodyConditionOperator = "gte"
	RuleIntentCreateRequestDetailsBodyConditionOperatorLt             RuleIntentCreateRequestDetailsBodyConditionOperator = "lt"
	RuleIntentCreateRequestDetailsBodyConditionOperatorLte            RuleIntentCreateRequestDetailsBodyConditionOperator = "lte"
	RuleIntentCreateRequestDetailsBodyConditionOperatorIn             RuleIntentCreateRequestDetailsBodyConditionOperator = "in"
	RuleIntentCreateRequestDetailsBodyConditionOperatorInConditionSet RuleIntentCreateRequestDetailsBodyConditionOperator = "in_condition_set"
)

type RuleIntentCreateRequestDetailsMethod string

const (
	RuleIntentCreateRequestDetailsMethodPost RuleIntentCreateRequestDetailsMethod = "POST"
)

// Request details for updating a rule via intent.
type RuleIntentUpdateRequestDetails struct {
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

type RuleIntentUpdateRequestDetailsBody struct {
	// Any of "ALLOW", "DENY".
	Action     string                                             `json:"action" api:"required"`
	Conditions []RuleIntentUpdateRequestDetailsBodyConditionUnion `json:"conditions" api:"required"`
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signTypedData_v4",
	// "eth_signUserOperation", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "signTransactionBytes", "exportPrivateKey", "\*".
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
// properties and values from [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject],
// [RuleIntentUpdateRequestDetailsBodyConditionObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentUpdateRequestDetailsBodyConditionUnion struct {
	// This field is from variant [RuleIntentUpdateRequestDetailsBodyConditionObject].
	Field RuleIntentUpdateRequestDetailsBodyConditionObjectField `json:"field"`
	// This field is from variant [RuleIntentUpdateRequestDetailsBodyConditionObject].
	FieldSource string `json:"field_source"`
	// This field is from variant [RuleIntentUpdateRequestDetailsBodyConditionObject].
	Operator RuleIntentUpdateRequestDetailsBodyConditionObjectOperator `json:"operator"`
	// This field is from variant [RuleIntentUpdateRequestDetailsBodyConditionObject].
	Value RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion `json:"value"`
	// This field is from variant [RuleIntentUpdateRequestDetailsBodyConditionObject].
	Abi []RuleIntentUpdateRequestDetailsBodyConditionObjectAbi `json:"abi"`
	// This field is from variant [RuleIntentUpdateRequestDetailsBodyConditionObject].
	TypedData RuleIntentUpdateRequestDetailsBodyConditionObjectTypedData `json:"typed_data"`
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

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsRuleIntentUpdateRequestDetailsBodyConditionObject() (v RuleIntentUpdateRequestDetailsBodyConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsVariant2() (v RuleIntentUpdateRequestDetailsBodyConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) AsVariant3() (v RuleIntentUpdateRequestDetailsBodyConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentUpdateRequestDetailsBodyConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentUpdateRequestDetailsBodyConditionObject struct {
	// Any of "to", "value", "chain_id".
	Field RuleIntentUpdateRequestDetailsBodyConditionObjectField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource string `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator RuleIntentUpdateRequestDetailsBodyConditionObjectOperator   `json:"operator" api:"required"`
	Value    RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion `json:"value" api:"required"`
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
func (r RuleIntentUpdateRequestDetailsBodyConditionObject) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentUpdateRequestDetailsBodyConditionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentUpdateRequestDetailsBodyConditionObjectField string

const (
	RuleIntentUpdateRequestDetailsBodyConditionObjectFieldTo      RuleIntentUpdateRequestDetailsBodyConditionObjectField = "to"
	RuleIntentUpdateRequestDetailsBodyConditionObjectFieldValue   RuleIntentUpdateRequestDetailsBodyConditionObjectField = "value"
	RuleIntentUpdateRequestDetailsBodyConditionObjectFieldChainID RuleIntentUpdateRequestDetailsBodyConditionObjectField = "chain_id"
)

type RuleIntentUpdateRequestDetailsBodyConditionObjectOperator string

const (
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorEq             RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "eq"
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorGt             RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "gt"
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorGte            RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "gte"
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorLt             RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "lt"
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorLte            RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "lte"
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorIn             RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "in"
	RuleIntentUpdateRequestDetailsBodyConditionObjectOperatorInConditionSet RuleIntentUpdateRequestDetailsBodyConditionObjectOperator = "in_condition_set"
)

// RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion struct {
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

func (u RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentUpdateRequestDetailsBodyConditionObjectValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentUpdateRequestDetailsBodyConditionOperator string

const (
	RuleIntentUpdateRequestDetailsBodyConditionOperatorEq             RuleIntentUpdateRequestDetailsBodyConditionOperator = "eq"
	RuleIntentUpdateRequestDetailsBodyConditionOperatorGt             RuleIntentUpdateRequestDetailsBodyConditionOperator = "gt"
	RuleIntentUpdateRequestDetailsBodyConditionOperatorGte            RuleIntentUpdateRequestDetailsBodyConditionOperator = "gte"
	RuleIntentUpdateRequestDetailsBodyConditionOperatorLt             RuleIntentUpdateRequestDetailsBodyConditionOperator = "lt"
	RuleIntentUpdateRequestDetailsBodyConditionOperatorLte            RuleIntentUpdateRequestDetailsBodyConditionOperator = "lte"
	RuleIntentUpdateRequestDetailsBodyConditionOperatorIn             RuleIntentUpdateRequestDetailsBodyConditionOperator = "in"
	RuleIntentUpdateRequestDetailsBodyConditionOperatorInConditionSet RuleIntentUpdateRequestDetailsBodyConditionOperator = "in_condition_set"
)

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
	// Detailed authorization information including key quorum members, thresholds, and
	// signature status
	AuthorizationDetails []IntentAuthorization `json:"authorization_details" api:"required"`
	// Unix timestamp when the intent was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// Display name of the user who created the intent
	CreatedByDisplayName string `json:"created_by_display_name" api:"required"`
	// Unix timestamp when the intent expires
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// Unique ID for the intent
	IntentID string `json:"intent_id" api:"required"`
	// Any of "RPC".
	IntentType RpcIntentResponseIntentType `json:"intent_type" api:"required"`
	// The original RPC request that would be sent to the wallet endpoint
	RequestDetails RpcIntentResponseRequestDetails `json:"request_details" api:"required"`
	// ID of the resource being modified (wallet_id, policy_id, etc)
	ResourceID string `json:"resource_id" api:"required"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `json:"status" api:"required"`
	// Result of RPC execution (only present if status is 'executed' or 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// ID of the user who created the intent. If undefined, the intent was created
	// using the app secret
	CreatedByID string `json:"created_by_id"`
	// A wallet managed by Privy's wallet infrastructure.
	CurrentResourceData Wallet `json:"current_resource_data"`
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
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		ActionResult         respjson.Field
		CreatedByID          respjson.Field
		CurrentResourceData  respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *RpcIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RpcIntentResponseIntentType string

const (
	RpcIntentResponseIntentTypeRpc RpcIntentResponseIntentType = "RPC"
)

// The original RPC request that would be sent to the wallet endpoint
type RpcIntentResponseRequestDetails struct {
	Body RpcIntentResponseRequestDetailsBodyUnion `json:"body" api:"required"`
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

// RpcIntentResponseRequestDetailsBodyUnion contains all possible properties and
// values from [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject],
// [RpcIntentResponseRequestDetailsBodyObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RpcIntentResponseRequestDetailsBodyUnion struct {
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	Method string `json:"method"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	Params RpcIntentResponseRequestDetailsBodyObjectParams `json:"params"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	Address string `json:"address"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	ChainType string `json:"chain_type"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	WalletID string `json:"wallet_id"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	Caip2 string `json:"caip2"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	Sponsor bool `json:"sponsor"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyObject].
	Network string `json:"network"`
	JSON    struct {
		Method    respjson.Field
		Params    respjson.Field
		Address   respjson.Field
		ChainType respjson.Field
		WalletID  respjson.Field
		Caip2     respjson.Field
		Sponsor   respjson.Field
		Network   respjson.Field
		raw       string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyUnion) AsRpcIntentResponseRequestDetailsBodyObject() (v RpcIntentResponseRequestDetailsBodyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyUnion) AsVariant2() (v RpcIntentResponseRequestDetailsBodyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyUnion) AsVariant3() (v RpcIntentResponseRequestDetailsBodyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyUnion) RawJSON() string { return u.JSON.raw }

func (r *RpcIntentResponseRequestDetailsBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RpcIntentResponseRequestDetailsBodyObject struct {
	// Any of "eth_signTransaction".
	Method  string                                          `json:"method" api:"required"`
	Params  RpcIntentResponseRequestDetailsBodyObjectParams `json:"params" api:"required"`
	Address string                                          `json:"address"`
	// Any of "ethereum".
	ChainType string `json:"chain_type"`
	WalletID  string `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponseRequestDetailsBodyObject) RawJSON() string { return r.JSON.raw }
func (r *RpcIntentResponseRequestDetailsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RpcIntentResponseRequestDetailsBodyObjectParams struct {
	Transaction RpcIntentResponseRequestDetailsBodyObjectParamsTransaction `json:"transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponseRequestDetailsBodyObjectParams) RawJSON() string { return r.JSON.raw }
func (r *RpcIntentResponseRequestDetailsBodyObjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RpcIntentResponseRequestDetailsBodyObjectParamsTransaction struct {
	AuthorizationList    []RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationList       `json:"authorization_list"`
	ChainID              RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion              `json:"chain_id"`
	Data                 string                                                                              `json:"data"`
	From                 string                                                                              `json:"from"`
	GasLimit             RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion             `json:"gas_limit"`
	GasPrice             RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion             `json:"gas_price"`
	MaxFeePerGas         RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion         `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas"`
	Nonce                RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion                `json:"nonce"`
	To                   string                                                                              `json:"to"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                              `json:"type"`
	Value RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponseRequestDetailsBodyObjectParamsTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationList struct {
	ChainID  RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion `json:"chain_id" api:"required"`
	Contract string                                                                                  `json:"contract" api:"required"`
	Nonce    RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion   `json:"nonce" api:"required"`
	R        string                                                                                  `json:"r" api:"required"`
	S        string                                                                                  `json:"s" api:"required"`
	YParity  float64                                                                                 `json:"y_parity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Contract    respjson.Field
		Nonce       respjson.Field
		R           respjson.Field
		S           respjson.Field
		YParity     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationList) RawJSON() string {
	return r.JSON.raw
}
func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionAuthorizationListNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RpcIntentResponseRequestDetailsBodyObjectParamsTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a wallet intent
type WalletIntentResponse struct {
	// Detailed authorization information including key quorum members, thresholds, and
	// signature status
	AuthorizationDetails []IntentAuthorization `json:"authorization_details" api:"required"`
	// Unix timestamp when the intent was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// Display name of the user who created the intent
	CreatedByDisplayName string `json:"created_by_display_name" api:"required"`
	// Unix timestamp when the intent expires
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// Unique ID for the intent
	IntentID string `json:"intent_id" api:"required"`
	// Any of "WALLET".
	IntentType WalletIntentResponseIntentType `json:"intent_type" api:"required"`
	// The original wallet update request that would be sent to the wallet endpoint
	RequestDetails WalletIntentResponseRequestDetails `json:"request_details" api:"required"`
	// ID of the resource being modified (wallet_id, policy_id, etc)
	ResourceID string `json:"resource_id" api:"required"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `json:"status" api:"required"`
	// Result of wallet update execution (only present if status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// ID of the user who created the intent. If undefined, the intent was created
	// using the app secret
	CreatedByID string `json:"created_by_id"`
	// A wallet managed by Privy's wallet infrastructure.
	CurrentResourceData Wallet `json:"current_resource_data"`
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
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		ActionResult         respjson.Field
		CreatedByID          respjson.Field
		CurrentResourceData  respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletIntentResponseIntentType string

const (
	WalletIntentResponseIntentTypeWallet WalletIntentResponseIntentType = "WALLET"
)

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
	AdditionalSigners      []WalletIntentResponseRequestDetailsBodyAdditionalSigner `json:"additional_signers"`
	AuthorizationKeyIDs    []string                                                 `json:"authorization_key_ids"`
	AuthorizationThreshold float64                                                  `json:"authorization_threshold"`
	Owner                  WalletIntentResponseRequestDetailsBodyOwnerUnion         `json:"owner" api:"nullable"`
	OwnerID                string                                                   `json:"owner_id" api:"nullable" format:"cuid2"`
	PolicyIDs              []string                                                 `json:"policy_ids" format:"cuid2"`
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

type WalletIntentResponseRequestDetailsBodyAdditionalSigner struct {
	SignerID          string   `json:"signer_id" api:"required" format:"cuid2"`
	OverridePolicyIDs []string `json:"override_policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignerID          respjson.Field
		OverridePolicyIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponseRequestDetailsBodyAdditionalSigner) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponseRequestDetailsBodyAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletIntentResponseRequestDetailsBodyOwnerUnion contains all possible
// properties and values from [WalletIntentResponseRequestDetailsBodyOwnerUserID],
// [WalletIntentResponseRequestDetailsBodyOwnerPublicKey].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletIntentResponseRequestDetailsBodyOwnerUnion struct {
	// This field is from variant [WalletIntentResponseRequestDetailsBodyOwnerUserID].
	UserID string `json:"user_id"`
	// This field is from variant
	// [WalletIntentResponseRequestDetailsBodyOwnerPublicKey].
	PublicKey string `json:"public_key"`
	JSON      struct {
		UserID    respjson.Field
		PublicKey respjson.Field
		raw       string
	} `json:"-"`
}

func (u WalletIntentResponseRequestDetailsBodyOwnerUnion) AsWalletIntentResponseRequestDetailsBodyOwnerUserID() (v WalletIntentResponseRequestDetailsBodyOwnerUserID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletIntentResponseRequestDetailsBodyOwnerUnion) AsWalletIntentResponseRequestDetailsBodyOwnerPublicKey() (v WalletIntentResponseRequestDetailsBodyOwnerPublicKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletIntentResponseRequestDetailsBodyOwnerUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletIntentResponseRequestDetailsBodyOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletIntentResponseRequestDetailsBodyOwnerUserID struct {
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponseRequestDetailsBodyOwnerUserID) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponseRequestDetailsBodyOwnerUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletIntentResponseRequestDetailsBodyOwnerPublicKey struct {
	PublicKey string `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletIntentResponseRequestDetailsBodyOwnerPublicKey) RawJSON() string { return r.JSON.raw }
func (r *WalletIntentResponseRequestDetailsBodyOwnerPublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a policy intent
type PolicyIntentResponse struct {
	// Detailed authorization information including key quorum members, thresholds, and
	// signature status
	AuthorizationDetails []IntentAuthorization `json:"authorization_details" api:"required"`
	// Unix timestamp when the intent was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// Display name of the user who created the intent
	CreatedByDisplayName string `json:"created_by_display_name" api:"required"`
	// Unix timestamp when the intent expires
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// Unique ID for the intent
	IntentID string `json:"intent_id" api:"required"`
	// Any of "POLICY".
	IntentType PolicyIntentResponseIntentType `json:"intent_type" api:"required"`
	// The original policy update request that would be sent to the policy endpoint
	RequestDetails PolicyIntentResponseRequestDetails `json:"request_details" api:"required"`
	// ID of the resource being modified (wallet_id, policy_id, etc)
	ResourceID string `json:"resource_id" api:"required"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `json:"status" api:"required"`
	// Result of policy update execution (only present if status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// ID of the user who created the intent. If undefined, the intent was created
	// using the app secret
	CreatedByID string `json:"created_by_id"`
	// Current state of the policy before any changes. If undefined, the resource was
	// deleted and no longer exists
	CurrentResourceData PolicyIntentResponseCurrentResourceData `json:"current_resource_data"`
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
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		ActionResult         respjson.Field
		CreatedByID          respjson.Field
		CurrentResourceData  respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseIntentType string

const (
	PolicyIntentResponseIntentTypePolicy PolicyIntentResponseIntentType = "POLICY"
)

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
	Name    string                                           `json:"name"`
	Owner   PolicyIntentResponseRequestDetailsBodyOwnerUnion `json:"owner" api:"nullable"`
	OwnerID string                                           `json:"owner_id" api:"nullable" format:"cuid2"`
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
// properties and values from [PolicyIntentResponseRequestDetailsBodyOwnerUserID],
// [PolicyIntentResponseRequestDetailsBodyOwnerPublicKey].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyIntentResponseRequestDetailsBodyOwnerUnion struct {
	// This field is from variant [PolicyIntentResponseRequestDetailsBodyOwnerUserID].
	UserID string `json:"user_id"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyOwnerPublicKey].
	PublicKey string `json:"public_key"`
	JSON      struct {
		UserID    respjson.Field
		PublicKey respjson.Field
		raw       string
	} `json:"-"`
}

func (u PolicyIntentResponseRequestDetailsBodyOwnerUnion) AsPolicyIntentResponseRequestDetailsBodyOwnerUserID() (v PolicyIntentResponseRequestDetailsBodyOwnerUserID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyOwnerUnion) AsPolicyIntentResponseRequestDetailsBodyOwnerPublicKey() (v PolicyIntentResponseRequestDetailsBodyOwnerPublicKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyOwnerUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyIntentResponseRequestDetailsBodyOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyOwnerUserID struct {
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyOwnerUserID) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseRequestDetailsBodyOwnerUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyOwnerPublicKey struct {
	PublicKey string `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyIntentResponseRequestDetailsBodyOwnerPublicKey) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseRequestDetailsBodyOwnerPublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyRule struct {
	// Any of "ALLOW", "DENY".
	Action     string                                                     `json:"action" api:"required"`
	Conditions []PolicyIntentResponseRequestDetailsBodyRuleConditionUnion `json:"conditions" api:"required"`
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signTypedData_v4",
	// "eth_signUserOperation", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "signTransactionBytes", "exportPrivateKey", "\*".
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
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject],
// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyIntentResponseRequestDetailsBodyRuleConditionUnion struct {
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
	Field PolicyIntentResponseRequestDetailsBodyRuleConditionObjectField `json:"field"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
	FieldSource string `json:"field_source"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
	Operator PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator `json:"operator"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
	Value PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion `json:"value"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
	Abi []PolicyIntentResponseRequestDetailsBodyRuleConditionObjectAbi `json:"abi"`
	// This field is from variant
	// [PolicyIntentResponseRequestDetailsBodyRuleConditionObject].
	TypedData PolicyIntentResponseRequestDetailsBodyRuleConditionObjectTypedData `json:"typed_data"`
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

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsPolicyIntentResponseRequestDetailsBodyRuleConditionObject() (v PolicyIntentResponseRequestDetailsBodyRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsVariant2() (v PolicyIntentResponseRequestDetailsBodyRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) AsVariant3() (v PolicyIntentResponseRequestDetailsBodyRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyRuleConditionObject struct {
	// Any of "to", "value", "chain_id".
	Field PolicyIntentResponseRequestDetailsBodyRuleConditionObjectField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource string `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator   `json:"operator" api:"required"`
	Value    PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion `json:"value" api:"required"`
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
func (r PolicyIntentResponseRequestDetailsBodyRuleConditionObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyRuleConditionObjectField string

const (
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectFieldTo      PolicyIntentResponseRequestDetailsBodyRuleConditionObjectField = "to"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectFieldValue   PolicyIntentResponseRequestDetailsBodyRuleConditionObjectField = "value"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectFieldChainID PolicyIntentResponseRequestDetailsBodyRuleConditionObjectField = "chain_id"
)

type PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator string

const (
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorEq             PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "eq"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorGt             PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "gt"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorGte            PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "gte"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorLt             PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "lt"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorLte            PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "lte"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorIn             PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "in"
	PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperatorInConditionSet PolicyIntentResponseRequestDetailsBodyRuleConditionObjectOperator = "in_condition_set"
)

// PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion struct {
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

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseRequestDetailsBodyRuleConditionObjectValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseRequestDetailsBodyRuleConditionOperator string

const (
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorEq             PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "eq"
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorGt             PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "gt"
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorGte            PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "gte"
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorLt             PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "lt"
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorLte            PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "lte"
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorIn             PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "in"
	PolicyIntentResponseRequestDetailsBodyRuleConditionOperatorInConditionSet PolicyIntentResponseRequestDetailsBodyRuleConditionOperator = "in_condition_set"
)

// Current state of the policy before any changes. If undefined, the resource was
// deleted and no longer exists
type PolicyIntentResponseCurrentResourceData struct {
	ID string `json:"id" api:"required"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType                               `json:"chain_type" api:"required"`
	CreatedAt float64                                       `json:"created_at" api:"required"`
	Name      string                                        `json:"name" api:"required"`
	OwnerID   string                                        `json:"owner_id" api:"required" format:"cuid2"`
	Rules     []PolicyIntentResponseCurrentResourceDataRule `json:"rules" api:"required"`
	// Any of "1.0".
	Version string `json:"version" api:"required"`
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
func (r PolicyIntentResponseCurrentResourceData) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseCurrentResourceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseCurrentResourceDataRule struct {
	ID string `json:"id" api:"required"`
	// Any of "ALLOW", "DENY".
	Action     string                                                      `json:"action" api:"required"`
	Conditions []PolicyIntentResponseCurrentResourceDataRuleConditionUnion `json:"conditions" api:"required"`
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signTypedData_v4",
	// "eth_signUserOperation", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "signTransactionBytes", "exportPrivateKey", "\*".
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
func (r PolicyIntentResponseCurrentResourceDataRule) RawJSON() string { return r.JSON.raw }
func (r *PolicyIntentResponseCurrentResourceDataRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyIntentResponseCurrentResourceDataRuleConditionUnion contains all possible
// properties and values from
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject],
// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PolicyIntentResponseCurrentResourceDataRuleConditionUnion struct {
	// This field is from variant
	// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
	Field PolicyIntentResponseCurrentResourceDataRuleConditionObjectField `json:"field"`
	// This field is from variant
	// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
	FieldSource string `json:"field_source"`
	// This field is from variant
	// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
	Operator PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator `json:"operator"`
	// This field is from variant
	// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
	Value PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion `json:"value"`
	// This field is from variant
	// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
	Abi []PolicyIntentResponseCurrentResourceDataRuleConditionObjectAbi `json:"abi"`
	// This field is from variant
	// [PolicyIntentResponseCurrentResourceDataRuleConditionObject].
	TypedData PolicyIntentResponseCurrentResourceDataRuleConditionObjectTypedData `json:"typed_data"`
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

func (u PolicyIntentResponseCurrentResourceDataRuleConditionUnion) AsPolicyIntentResponseCurrentResourceDataRuleConditionObject() (v PolicyIntentResponseCurrentResourceDataRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseCurrentResourceDataRuleConditionUnion) AsVariant2() (v PolicyIntentResponseCurrentResourceDataRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseCurrentResourceDataRuleConditionUnion) AsVariant3() (v PolicyIntentResponseCurrentResourceDataRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseCurrentResourceDataRuleConditionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseCurrentResourceDataRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseCurrentResourceDataRuleConditionObject struct {
	// Any of "to", "value", "chain_id".
	Field PolicyIntentResponseCurrentResourceDataRuleConditionObjectField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource string `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator   `json:"operator" api:"required"`
	Value    PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion `json:"value" api:"required"`
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
func (r PolicyIntentResponseCurrentResourceDataRuleConditionObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PolicyIntentResponseCurrentResourceDataRuleConditionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseCurrentResourceDataRuleConditionObjectField string

const (
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectFieldTo      PolicyIntentResponseCurrentResourceDataRuleConditionObjectField = "to"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectFieldValue   PolicyIntentResponseCurrentResourceDataRuleConditionObjectField = "value"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectFieldChainID PolicyIntentResponseCurrentResourceDataRuleConditionObjectField = "chain_id"
)

type PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator string

const (
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorEq             PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "eq"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorGt             PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "gt"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorGte            PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "gte"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorLt             PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "lt"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorLte            PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "lte"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorIn             PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "in"
	PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperatorInConditionSet PolicyIntentResponseCurrentResourceDataRuleConditionObjectOperator = "in_condition_set"
)

// PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion contains
// all possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion struct {
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

func (u PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PolicyIntentResponseCurrentResourceDataRuleConditionObjectValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyIntentResponseCurrentResourceDataRuleConditionOperator string

const (
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorEq             PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "eq"
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorGt             PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "gt"
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorGte            PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "gte"
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorLt             PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "lt"
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorLte            PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "lte"
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorIn             PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "in"
	PolicyIntentResponseCurrentResourceDataRuleConditionOperatorInConditionSet PolicyIntentResponseCurrentResourceDataRuleConditionOperator = "in_condition_set"
)

// Response for a key quorum intent
type KeyQuorumIntentResponse struct {
	// Detailed authorization information including key quorum members, thresholds, and
	// signature status
	AuthorizationDetails []IntentAuthorization `json:"authorization_details" api:"required"`
	// Unix timestamp when the intent was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// Display name of the user who created the intent
	CreatedByDisplayName string `json:"created_by_display_name" api:"required"`
	// Unix timestamp when the intent expires
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// Unique ID for the intent
	IntentID string `json:"intent_id" api:"required"`
	// Any of "KEY_QUORUM".
	IntentType KeyQuorumIntentResponseIntentType `json:"intent_type" api:"required"`
	// The original key quorum update request that would be sent to the key quorum
	// endpoint
	RequestDetails KeyQuorumIntentResponseRequestDetails `json:"request_details" api:"required"`
	// ID of the resource being modified (wallet_id, policy_id, etc)
	ResourceID string `json:"resource_id" api:"required"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `json:"status" api:"required"`
	// Result of key quorum update execution (only present if status is 'executed' or
	// 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// ID of the user who created the intent. If undefined, the intent was created
	// using the app secret
	CreatedByID string `json:"created_by_id"`
	// Current state of the key quorum before any changes. If undefined, the resource
	// was deleted and no longer exists
	CurrentResourceData KeyQuorumIntentResponseCurrentResourceData `json:"current_resource_data"`
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
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		ActionResult         respjson.Field
		CreatedByID          respjson.Field
		CurrentResourceData  respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumIntentResponseIntentType string

const (
	KeyQuorumIntentResponseIntentTypeKeyQuorum KeyQuorumIntentResponseIntentType = "KEY_QUORUM"
)

// The original key quorum update request that would be sent to the key quorum
// endpoint
type KeyQuorumIntentResponseRequestDetails struct {
	Body KeyQuorumIntentResponseRequestDetailsBody `json:"body" api:"required"`
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

type KeyQuorumIntentResponseRequestDetailsBody struct {
	AuthorizationThreshold float64  `json:"authorization_threshold"`
	DisplayName            string   `json:"display_name"`
	KeyQuorumIDs           []string `json:"key_quorum_ids"`
	PublicKeys             []string `json:"public_keys"`
	UserIDs                []string `json:"user_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationThreshold respjson.Field
		DisplayName            respjson.Field
		KeyQuorumIDs           respjson.Field
		PublicKeys             respjson.Field
		UserIDs                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumIntentResponseRequestDetailsBody) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumIntentResponseRequestDetailsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the key quorum before any changes. If undefined, the resource
// was deleted and no longer exists
type KeyQuorumIntentResponseCurrentResourceData struct {
	ID                     string                                                       `json:"id" api:"required"`
	AuthorizationKeys      []KeyQuorumIntentResponseCurrentResourceDataAuthorizationKey `json:"authorization_keys" api:"required"`
	AuthorizationThreshold float64                                                      `json:"authorization_threshold" api:"required"`
	DisplayName            string                                                       `json:"display_name" api:"required"`
	UserIDs                []string                                                     `json:"user_ids" api:"required"`
	KeyQuorumIDs           []string                                                     `json:"key_quorum_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AuthorizationKeys      respjson.Field
		AuthorizationThreshold respjson.Field
		DisplayName            respjson.Field
		UserIDs                respjson.Field
		KeyQuorumIDs           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumIntentResponseCurrentResourceData) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumIntentResponseCurrentResourceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumIntentResponseCurrentResourceDataAuthorizationKey struct {
	DisplayName string `json:"display_name" api:"required"`
	PublicKey   string `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumIntentResponseCurrentResourceDataAuthorizationKey) RawJSON() string {
	return r.JSON.raw
}
func (r *KeyQuorumIntentResponseCurrentResourceDataAuthorizationKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a rule intent
type RuleIntentResponse struct {
	// Detailed authorization information including key quorum members, thresholds, and
	// signature status
	AuthorizationDetails []IntentAuthorization `json:"authorization_details" api:"required"`
	// Unix timestamp when the intent was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// Display name of the user who created the intent
	CreatedByDisplayName string `json:"created_by_display_name" api:"required"`
	// Unix timestamp when the intent expires
	ExpiresAt float64 `json:"expires_at" api:"required"`
	// Unique ID for the intent
	IntentID string `json:"intent_id" api:"required"`
	// Any of "RULE".
	IntentType RuleIntentResponseIntentType `json:"intent_type" api:"required"`
	// The original rule request. Method is POST (create), PATCH (update), or DELETE
	// (delete)
	RequestDetails RuleIntentRequestDetailsUnion `json:"request_details" api:"required"`
	// ID of the resource being modified (wallet_id, policy_id, etc)
	ResourceID string `json:"resource_id" api:"required"`
	// Current status of an intent.
	//
	// Any of "pending", "executed", "failed", "expired", "rejected", "dismissed".
	Status IntentStatus `json:"status" api:"required"`
	// Result of rule execution (only present if status is 'executed' or 'failed')
	ActionResult BaseActionResult `json:"action_result"`
	// ID of the user who created the intent. If undefined, the intent was created
	// using the app secret
	CreatedByID string `json:"created_by_id"`
	// Current state of the rule before any changes. Undefined for create intents or if
	// the rule was deleted
	CurrentResourceData RuleIntentResponseCurrentResourceData `json:"current_resource_data"`
	// Human-readable reason for dismissal, present when status is 'dismissed'
	DismissalReason string `json:"dismissal_reason"`
	// Unix timestamp when the intent was dismissed, present when status is 'dismissed'
	DismissedAt float64 `json:"dismissed_at"`
	// Parent policy containing this rule, including sibling rules for contextual
	// display
	Policy RuleIntentResponsePolicy `json:"policy"`
	// Unix timestamp when the intent was rejected, present when status is 'rejected'
	RejectedAt float64 `json:"rejected_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationDetails respjson.Field
		CreatedAt            respjson.Field
		CreatedByDisplayName respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		ActionResult         respjson.Field
		CreatedByID          respjson.Field
		CurrentResourceData  respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		Policy               respjson.Field
		RejectedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleIntentResponse) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponseIntentType string

const (
	RuleIntentResponseIntentTypeRule RuleIntentResponseIntentType = "RULE"
)

// Current state of the rule before any changes. Undefined for create intents or if
// the rule was deleted
type RuleIntentResponseCurrentResourceData struct {
	ID string `json:"id" api:"required"`
	// Any of "ALLOW", "DENY".
	Action     string                                                `json:"action" api:"required"`
	Conditions []RuleIntentResponseCurrentResourceDataConditionUnion `json:"conditions" api:"required"`
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signTypedData_v4",
	// "eth_signUserOperation", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "signTransactionBytes", "exportPrivateKey", "\*".
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
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject],
// [RuleIntentResponseCurrentResourceDataConditionObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentResponseCurrentResourceDataConditionUnion struct {
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionObject].
	Field RuleIntentResponseCurrentResourceDataConditionObjectField `json:"field"`
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionObject].
	FieldSource string `json:"field_source"`
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionObject].
	Operator RuleIntentResponseCurrentResourceDataConditionObjectOperator `json:"operator"`
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionObject].
	Value RuleIntentResponseCurrentResourceDataConditionObjectValueUnion `json:"value"`
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionObject].
	Abi []RuleIntentResponseCurrentResourceDataConditionObjectAbi `json:"abi"`
	// This field is from variant
	// [RuleIntentResponseCurrentResourceDataConditionObject].
	TypedData RuleIntentResponseCurrentResourceDataConditionObjectTypedData `json:"typed_data"`
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

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsRuleIntentResponseCurrentResourceDataConditionObject() (v RuleIntentResponseCurrentResourceDataConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsVariant2() (v RuleIntentResponseCurrentResourceDataConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionUnion) AsVariant3() (v RuleIntentResponseCurrentResourceDataConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentResponseCurrentResourceDataConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponseCurrentResourceDataConditionObject struct {
	// Any of "to", "value", "chain_id".
	Field RuleIntentResponseCurrentResourceDataConditionObjectField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource string `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator RuleIntentResponseCurrentResourceDataConditionObjectOperator   `json:"operator" api:"required"`
	Value    RuleIntentResponseCurrentResourceDataConditionObjectValueUnion `json:"value" api:"required"`
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
func (r RuleIntentResponseCurrentResourceDataConditionObject) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponseCurrentResourceDataConditionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponseCurrentResourceDataConditionObjectField string

const (
	RuleIntentResponseCurrentResourceDataConditionObjectFieldTo      RuleIntentResponseCurrentResourceDataConditionObjectField = "to"
	RuleIntentResponseCurrentResourceDataConditionObjectFieldValue   RuleIntentResponseCurrentResourceDataConditionObjectField = "value"
	RuleIntentResponseCurrentResourceDataConditionObjectFieldChainID RuleIntentResponseCurrentResourceDataConditionObjectField = "chain_id"
)

type RuleIntentResponseCurrentResourceDataConditionObjectOperator string

const (
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorEq             RuleIntentResponseCurrentResourceDataConditionObjectOperator = "eq"
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorGt             RuleIntentResponseCurrentResourceDataConditionObjectOperator = "gt"
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorGte            RuleIntentResponseCurrentResourceDataConditionObjectOperator = "gte"
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorLt             RuleIntentResponseCurrentResourceDataConditionObjectOperator = "lt"
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorLte            RuleIntentResponseCurrentResourceDataConditionObjectOperator = "lte"
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorIn             RuleIntentResponseCurrentResourceDataConditionObjectOperator = "in"
	RuleIntentResponseCurrentResourceDataConditionObjectOperatorInConditionSet RuleIntentResponseCurrentResourceDataConditionObjectOperator = "in_condition_set"
)

// RuleIntentResponseCurrentResourceDataConditionObjectValueUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponseCurrentResourceDataConditionObjectValueUnion struct {
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

func (u RuleIntentResponseCurrentResourceDataConditionObjectValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponseCurrentResourceDataConditionObjectValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponseCurrentResourceDataConditionObjectValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *RuleIntentResponseCurrentResourceDataConditionObjectValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponseCurrentResourceDataConditionOperator string

const (
	RuleIntentResponseCurrentResourceDataConditionOperatorEq             RuleIntentResponseCurrentResourceDataConditionOperator = "eq"
	RuleIntentResponseCurrentResourceDataConditionOperatorGt             RuleIntentResponseCurrentResourceDataConditionOperator = "gt"
	RuleIntentResponseCurrentResourceDataConditionOperatorGte            RuleIntentResponseCurrentResourceDataConditionOperator = "gte"
	RuleIntentResponseCurrentResourceDataConditionOperatorLt             RuleIntentResponseCurrentResourceDataConditionOperator = "lt"
	RuleIntentResponseCurrentResourceDataConditionOperatorLte            RuleIntentResponseCurrentResourceDataConditionOperator = "lte"
	RuleIntentResponseCurrentResourceDataConditionOperatorIn             RuleIntentResponseCurrentResourceDataConditionOperator = "in"
	RuleIntentResponseCurrentResourceDataConditionOperatorInConditionSet RuleIntentResponseCurrentResourceDataConditionOperator = "in_condition_set"
)

// Parent policy containing this rule, including sibling rules for contextual
// display
type RuleIntentResponsePolicy struct {
	ID string `json:"id" api:"required"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType                `json:"chain_type" api:"required"`
	CreatedAt float64                        `json:"created_at" api:"required"`
	Name      string                         `json:"name" api:"required"`
	OwnerID   string                         `json:"owner_id" api:"required" format:"cuid2"`
	Rules     []RuleIntentResponsePolicyRule `json:"rules" api:"required"`
	// Any of "1.0".
	Version string `json:"version" api:"required"`
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
func (r RuleIntentResponsePolicy) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponsePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponsePolicyRule struct {
	ID string `json:"id" api:"required"`
	// Any of "ALLOW", "DENY".
	Action     string                                       `json:"action" api:"required"`
	Conditions []RuleIntentResponsePolicyRuleConditionUnion `json:"conditions" api:"required"`
	// Any of "eth_sendTransaction", "eth_signTransaction", "eth_signTypedData_v4",
	// "eth_signUserOperation", "eth_sign7702Authorization", "signTransaction",
	// "signAndSendTransaction", "signTransactionBytes", "exportPrivateKey", "\*".
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
func (r RuleIntentResponsePolicyRule) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponsePolicyRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RuleIntentResponsePolicyRuleConditionUnion contains all possible properties and
// values from [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject],
// [RuleIntentResponsePolicyRuleConditionObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RuleIntentResponsePolicyRuleConditionUnion struct {
	// This field is from variant [RuleIntentResponsePolicyRuleConditionObject].
	Field RuleIntentResponsePolicyRuleConditionObjectField `json:"field"`
	// This field is from variant [RuleIntentResponsePolicyRuleConditionObject].
	FieldSource string `json:"field_source"`
	// This field is from variant [RuleIntentResponsePolicyRuleConditionObject].
	Operator RuleIntentResponsePolicyRuleConditionObjectOperator `json:"operator"`
	// This field is from variant [RuleIntentResponsePolicyRuleConditionObject].
	Value RuleIntentResponsePolicyRuleConditionObjectValueUnion `json:"value"`
	// This field is from variant [RuleIntentResponsePolicyRuleConditionObject].
	Abi []RuleIntentResponsePolicyRuleConditionObjectAbi `json:"abi"`
	// This field is from variant [RuleIntentResponsePolicyRuleConditionObject].
	TypedData RuleIntentResponsePolicyRuleConditionObjectTypedData `json:"typed_data"`
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

func (u RuleIntentResponsePolicyRuleConditionUnion) AsRuleIntentResponsePolicyRuleConditionObject() (v RuleIntentResponsePolicyRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponsePolicyRuleConditionUnion) AsVariant2() (v RuleIntentResponsePolicyRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponsePolicyRuleConditionUnion) AsVariant3() (v RuleIntentResponsePolicyRuleConditionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponsePolicyRuleConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentResponsePolicyRuleConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponsePolicyRuleConditionObject struct {
	// Any of "to", "value", "chain_id".
	Field RuleIntentResponsePolicyRuleConditionObjectField `json:"field" api:"required"`
	// Any of "ethereum_transaction".
	FieldSource string `json:"field_source" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator RuleIntentResponsePolicyRuleConditionObjectOperator   `json:"operator" api:"required"`
	Value    RuleIntentResponsePolicyRuleConditionObjectValueUnion `json:"value" api:"required"`
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
func (r RuleIntentResponsePolicyRuleConditionObject) RawJSON() string { return r.JSON.raw }
func (r *RuleIntentResponsePolicyRuleConditionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponsePolicyRuleConditionObjectField string

const (
	RuleIntentResponsePolicyRuleConditionObjectFieldTo      RuleIntentResponsePolicyRuleConditionObjectField = "to"
	RuleIntentResponsePolicyRuleConditionObjectFieldValue   RuleIntentResponsePolicyRuleConditionObjectField = "value"
	RuleIntentResponsePolicyRuleConditionObjectFieldChainID RuleIntentResponsePolicyRuleConditionObjectField = "chain_id"
)

type RuleIntentResponsePolicyRuleConditionObjectOperator string

const (
	RuleIntentResponsePolicyRuleConditionObjectOperatorEq             RuleIntentResponsePolicyRuleConditionObjectOperator = "eq"
	RuleIntentResponsePolicyRuleConditionObjectOperatorGt             RuleIntentResponsePolicyRuleConditionObjectOperator = "gt"
	RuleIntentResponsePolicyRuleConditionObjectOperatorGte            RuleIntentResponsePolicyRuleConditionObjectOperator = "gte"
	RuleIntentResponsePolicyRuleConditionObjectOperatorLt             RuleIntentResponsePolicyRuleConditionObjectOperator = "lt"
	RuleIntentResponsePolicyRuleConditionObjectOperatorLte            RuleIntentResponsePolicyRuleConditionObjectOperator = "lte"
	RuleIntentResponsePolicyRuleConditionObjectOperatorIn             RuleIntentResponsePolicyRuleConditionObjectOperator = "in"
	RuleIntentResponsePolicyRuleConditionObjectOperatorInConditionSet RuleIntentResponsePolicyRuleConditionObjectOperator = "in_condition_set"
)

// RuleIntentResponsePolicyRuleConditionObjectValueUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type RuleIntentResponsePolicyRuleConditionObjectValueUnion struct {
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

func (u RuleIntentResponsePolicyRuleConditionObjectValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RuleIntentResponsePolicyRuleConditionObjectValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RuleIntentResponsePolicyRuleConditionObjectValueUnion) RawJSON() string { return u.JSON.raw }

func (r *RuleIntentResponsePolicyRuleConditionObjectValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleIntentResponsePolicyRuleConditionOperator string

const (
	RuleIntentResponsePolicyRuleConditionOperatorEq             RuleIntentResponsePolicyRuleConditionOperator = "eq"
	RuleIntentResponsePolicyRuleConditionOperatorGt             RuleIntentResponsePolicyRuleConditionOperator = "gt"
	RuleIntentResponsePolicyRuleConditionOperatorGte            RuleIntentResponsePolicyRuleConditionOperator = "gte"
	RuleIntentResponsePolicyRuleConditionOperatorLt             RuleIntentResponsePolicyRuleConditionOperator = "lt"
	RuleIntentResponsePolicyRuleConditionOperatorLte            RuleIntentResponsePolicyRuleConditionOperator = "lte"
	RuleIntentResponsePolicyRuleConditionOperatorIn             RuleIntentResponsePolicyRuleConditionOperator = "in"
	RuleIntentResponsePolicyRuleConditionOperatorInConditionSet RuleIntentResponsePolicyRuleConditionOperator = "in_condition_set"
)

// IntentResponseUnion contains all possible properties and values from
// [RpcIntentResponse], [WalletIntentResponse], [PolicyIntentResponse],
// [RuleIntentResponse], [KeyQuorumIntentResponse].
//
// Use the [IntentResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IntentResponseUnion struct {
	AuthorizationDetails []IntentAuthorization `json:"authorization_details"`
	CreatedAt            float64               `json:"created_at"`
	CreatedByDisplayName string                `json:"created_by_display_name"`
	ExpiresAt            float64               `json:"expires_at"`
	IntentID             string                `json:"intent_id"`
	// Any of "RPC", "WALLET", "POLICY", "RULE", "KEY_QUORUM".
	IntentType string `json:"intent_type"`
	// This field is a union of [RpcIntentResponseRequestDetails],
	// [WalletIntentResponseRequestDetails], [PolicyIntentResponseRequestDetails],
	// [RuleIntentRequestDetailsUnion], [KeyQuorumIntentResponseRequestDetails]
	RequestDetails IntentResponseUnionRequestDetails `json:"request_details"`
	ResourceID     string                            `json:"resource_id"`
	// This field is from variant [RpcIntentResponse].
	Status IntentStatus `json:"status"`
	// This field is from variant [RpcIntentResponse].
	ActionResult BaseActionResult `json:"action_result"`
	CreatedByID  string           `json:"created_by_id"`
	// This field is a union of [Wallet], [PolicyIntentResponseCurrentResourceData],
	// [RuleIntentResponseCurrentResourceData],
	// [KeyQuorumIntentResponseCurrentResourceData]
	CurrentResourceData IntentResponseUnionCurrentResourceData `json:"current_resource_data"`
	DismissalReason     string                                 `json:"dismissal_reason"`
	DismissedAt         float64                                `json:"dismissed_at"`
	RejectedAt          float64                                `json:"rejected_at"`
	// This field is from variant [RuleIntentResponse].
	Policy RuleIntentResponsePolicy `json:"policy"`
	JSON   struct {
		AuthorizationDetails respjson.Field
		CreatedAt            respjson.Field
		CreatedByDisplayName respjson.Field
		ExpiresAt            respjson.Field
		IntentID             respjson.Field
		IntentType           respjson.Field
		RequestDetails       respjson.Field
		ResourceID           respjson.Field
		Status               respjson.Field
		ActionResult         respjson.Field
		CreatedByID          respjson.Field
		CurrentResourceData  respjson.Field
		DismissalReason      respjson.Field
		DismissedAt          respjson.Field
		RejectedAt           respjson.Field
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
func (WalletIntentResponse) implIntentResponseUnion()    {}
func (PolicyIntentResponse) implIntentResponseUnion()    {}
func (RuleIntentResponse) implIntentResponseUnion()      {}
func (KeyQuorumIntentResponse) implIntentResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := IntentResponseUnion.AsAny().(type) {
//	case privyclient.RpcIntentResponse:
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
	// This field is a union of [RpcIntentResponseRequestDetailsBodyUnion],
	// [WalletIntentResponseRequestDetailsBody],
	// [PolicyIntentResponseRequestDetailsBody], [RuleIntentCreateRequestDetailsBody],
	// [RuleIntentUpdateRequestDetailsBody], [any],
	// [KeyQuorumIntentResponseRequestDetailsBody]
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
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	Params RpcIntentResponseRequestDetailsBodyObjectParams `json:"params"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	Address string `json:"address"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	ChainType string `json:"chain_type"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	WalletID string `json:"wallet_id"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	Caip2 string `json:"caip2"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	Sponsor bool `json:"sponsor"`
	// This field is from variant [RpcIntentResponseRequestDetailsBodyUnion].
	Network string `json:"network"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	AdditionalSigners []WalletIntentResponseRequestDetailsBodyAdditionalSigner `json:"additional_signers"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	AuthorizationKeyIDs    []string `json:"authorization_key_ids"`
	AuthorizationThreshold float64  `json:"authorization_threshold"`
	// This field is a union of [WalletIntentResponseRequestDetailsBodyOwnerUnion],
	// [PolicyIntentResponseRequestDetailsBodyOwnerUnion]
	Owner   IntentResponseUnionRequestDetailsBodyOwner `json:"owner"`
	OwnerID string                                     `json:"owner_id"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	PolicyIDs []string `json:"policy_ids"`
	Name      string   `json:"name"`
	// This field is from variant [PolicyIntentResponseRequestDetailsBody].
	Rules  []PolicyIntentResponseRequestDetailsBodyRule `json:"rules"`
	Action string                                       `json:"action"`
	// This field is a union of [[]RuleIntentCreateRequestDetailsBodyConditionUnion],
	// [[]RuleIntentUpdateRequestDetailsBodyConditionUnion]
	Conditions IntentResponseUnionRequestDetailsBodyConditions `json:"conditions"`
	// This field is from variant [KeyQuorumIntentResponseRequestDetailsBody].
	DisplayName string `json:"display_name"`
	// This field is from variant [KeyQuorumIntentResponseRequestDetailsBody].
	KeyQuorumIDs []string `json:"key_quorum_ids"`
	// This field is from variant [KeyQuorumIntentResponseRequestDetailsBody].
	PublicKeys []string `json:"public_keys"`
	// This field is from variant [KeyQuorumIntentResponseRequestDetailsBody].
	UserIDs []string `json:"user_ids"`
	JSON    struct {
		OfRuleIntentDeleteRequestDetailsBody respjson.Field
		Method                               respjson.Field
		Params                               respjson.Field
		Address                              respjson.Field
		ChainType                            respjson.Field
		WalletID                             respjson.Field
		Caip2                                respjson.Field
		Sponsor                              respjson.Field
		Network                              respjson.Field
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

// IntentResponseUnionRequestDetailsBodyOwner is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionRequestDetailsBodyOwner provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
type IntentResponseUnionRequestDetailsBodyOwner struct {
	UserID    string `json:"user_id"`
	PublicKey string `json:"public_key"`
	JSON      struct {
		UserID    respjson.Field
		PublicKey respjson.Field
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
	// This field is from variant [PolicyIntentResponseCurrentResourceData].
	Rules []PolicyIntentResponseCurrentResourceDataRule `json:"rules"`
	// This field is from variant [PolicyIntentResponseCurrentResourceData].
	Version string `json:"version"`
	// This field is from variant [RuleIntentResponseCurrentResourceData].
	Action string `json:"action"`
	// This field is from variant [RuleIntentResponseCurrentResourceData].
	Conditions []RuleIntentResponseCurrentResourceDataConditionUnion `json:"conditions"`
	// This field is from variant [RuleIntentResponseCurrentResourceData].
	Method string `json:"method"`
	// This field is from variant [KeyQuorumIntentResponseCurrentResourceData].
	AuthorizationKeys []KeyQuorumIntentResponseCurrentResourceDataAuthorizationKey `json:"authorization_keys"`
	// This field is from variant [KeyQuorumIntentResponseCurrentResourceData].
	DisplayName string `json:"display_name"`
	// This field is from variant [KeyQuorumIntentResponseCurrentResourceData].
	UserIDs []string `json:"user_ids"`
	// This field is from variant [KeyQuorumIntentResponseCurrentResourceData].
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
	// Any of "KEY_QUORUM", "POLICY", "RULE", "RPC", "WALLET".
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
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                              `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                          `json:",omitzero,inline"`
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
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
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
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
	FieldSource constant.EthereumTransaction `json:"field_source" api:"required"`
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
	FieldSource constant.EthereumCalldata `json:"field_source" api:"required"`
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
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                              `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" api:"required"`
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
		"field", "chainId", "verifyingContract",
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
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" api:"required"`
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
	PrimaryType string                                                                               `json:"primary_type" api:"required"`
	Types       map[string][]IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero" api:"required"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	paramObj
}

func (r IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentNewPolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
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
	FieldSource constant.Ethereum7702Authorization `json:"field_source" api:"required"`
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
	FieldSource constant.SolanaProgramInstruction `json:"field_source" api:"required"`
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
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" api:"required"`
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
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                    `json:"operator,omitzero" api:"required"`
	Value    IntentNewPolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" api:"required"`
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
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
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
	FieldSource constant.System `json:"field_source" api:"required"`
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
	paramObj
}

type IntentRpcParams struct {
	// Request body for wallet RPC operations, discriminated by method.
	WalletRpcRequestBody WalletRpcRequestBodyUnion
	paramObj
}

func (r IntentRpcParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletRpcRequestBody)
}
func (r *IntentRpcParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WalletRpcRequestBody)
}

type IntentUpdateKeyQuorumParams struct {
	// Request input for creating or updating a key quorum.
	KeyQuorumCreateParams KeyQuorumCreateParams
	paramObj
}

func (r IntentUpdateKeyQuorumParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KeyQuorumCreateParams)
}
func (r *IntentUpdateKeyQuorumParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.KeyQuorumCreateParams)
}

type IntentUpdatePolicyParams struct {
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// Name to assign to policy.
	Name param.Opt[string] `json:"name,omitzero"`
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
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                                 `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                             `json:",omitzero,inline"`
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
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
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
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
	FieldSource constant.EthereumTransaction `json:"field_source" api:"required"`
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
	FieldSource constant.EthereumCalldata `json:"field_source" api:"required"`
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
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" api:"required"`
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
		"field", "chainId", "verifyingContract",
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
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" api:"required"`
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
	PrimaryType string                                                                                  `json:"primary_type" api:"required"`
	Types       map[string][]IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParamsRuleConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
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
	FieldSource constant.Ethereum7702Authorization `json:"field_source" api:"required"`
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
	FieldSource constant.SolanaProgramInstruction `json:"field_source" api:"required"`
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
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" api:"required"`
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
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                       `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyParamsRuleConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" api:"required"`
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
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
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
	FieldSource constant.System `json:"field_source" api:"required"`
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
	OfSuiTransactionCommand          *SuiTransactionCommandConditionParam                                 `json:",omitzero,inline"`
	OfSuiTransferObjectsCommand      *SuiTransferObjectsCommandConditionParam                             `json:",omitzero,inline"`
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
		u.OfSuiTransactionCommand,
		u.OfSuiTransferObjectsCommand)
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
		apijson.Discriminator[SuiTransactionCommandConditionParam]("sui_transaction_command"),
		apijson.Discriminator[SuiTransferObjectsCommandConditionParam]("sui_transfer_objects_command"),
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
	FieldSource constant.EthereumTransaction `json:"field_source" api:"required"`
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
	FieldSource constant.EthereumCalldata `json:"field_source" api:"required"`
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
	// Any of "chainId", "verifyingContract".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                 `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionEthereumTypedDataDomainValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ethereum_typed_data_domain".
	FieldSource constant.EthereumTypedDataDomain `json:"field_source" api:"required"`
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
		"field", "chainId", "verifyingContract",
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
	FieldSource constant.EthereumTypedDataMessage `json:"field_source" api:"required"`
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
	PrimaryType string                                                                                  `json:"primary_type" api:"required"`
	Types       map[string][]IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType `json:"types,omitzero" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	paramObj
}

func (r IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyRuleParamsConditionEthereumTypedDataMessageTypedDataType) UnmarshalJSON(data []byte) error {
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
	FieldSource constant.Ethereum7702Authorization `json:"field_source" api:"required"`
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
	FieldSource constant.SolanaProgramInstruction `json:"field_source" api:"required"`
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
	FieldSource constant.SolanaSystemProgramInstruction `json:"field_source" api:"required"`
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
	// Any of "instructionName", "TransferChecked.source",
	// "TransferChecked.destination", "TransferChecked.authority",
	// "TransferChecked.amount", "TransferChecked.mint".
	Field string `json:"field,omitzero" api:"required"`
	// Any of "eq", "gt", "gte", "lt", "lte", "in", "in_condition_set".
	Operator string                                                                       `json:"operator,omitzero" api:"required"`
	Value    IntentUpdatePolicyRuleParamsConditionSolanaTokenProgramInstructionValueUnion `json:"value,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "solana_token_program_instruction".
	FieldSource constant.SolanaTokenProgramInstruction `json:"field_source" api:"required"`
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
		"field", "instructionName", "TransferChecked.source", "TransferChecked.destination", "TransferChecked.authority", "TransferChecked.amount", "TransferChecked.mint",
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
	FieldSource constant.System `json:"field_source" api:"required"`
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
	paramObj
}

func (r IntentUpdateWalletParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletUpdateRequestBody)
}
func (r *IntentUpdateWalletParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WalletUpdateRequestBody)
}
