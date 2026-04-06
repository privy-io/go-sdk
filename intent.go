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
	Body PolicyRuleRequestBody `json:"body" api:"required"`
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

type RuleIntentCreateRequestDetailsMethod string

const (
	RuleIntentCreateRequestDetailsMethodPost RuleIntentCreateRequestDetailsMethod = "POST"
)

// Request details for updating a rule via intent.
type RuleIntentUpdateRequestDetails struct {
	// The rules that apply to each method the policy covers.
	Body PolicyRuleRequestBody `json:"body" api:"required"`
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
	// This field is a union of [PolicyRuleRequestBody], [any]
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
	OfRuleIntentDeleteRequestDetailsBody any `json:",inline"`
	// This field is from variant [PolicyRuleRequestBody].
	Action PolicyAction `json:"action"`
	// This field is from variant [PolicyRuleRequestBody].
	Conditions []PolicyConditionUnion `json:"conditions"`
	// This field is from variant [PolicyRuleRequestBody].
	Method PolicyMethod `json:"method"`
	// This field is from variant [PolicyRuleRequestBody].
	Name string `json:"name"`
	JSON struct {
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
	DisplayName            string                `json:"display_name" api:"nullable"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner" api:"nullable"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID OwnerIDInput `json:"owner_id" api:"nullable" format:"cuid2"`
	// An optional list of up to one policy ID to enforce on the wallet.
	PolicyIDs PolicyInput `json:"policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalSigners      respjson.Field
		AuthorizationKeyIDs    respjson.Field
		AuthorizationThreshold respjson.Field
		DisplayName            respjson.Field
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
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner" api:"nullable"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID OwnerIDInput            `json:"owner_id" api:"nullable" format:"cuid2"`
	Rules   []PolicyRuleRequestBody `json:"rules"`
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
	// A rule that defines the conditions and action to take if the conditions are
	// true.
	CurrentResourceData PolicyRuleResponse `json:"current_resource_data"`
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
	// This field is a union of [Wallet], [Policy], [PolicyRuleResponse], [KeyQuorum]
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
	// [PolicyIntentResponseRequestDetailsBody], [PolicyRuleRequestBody], [any],
	// [KeyQuorumUpdateRequestBody]
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
	// [SparkSignMessageWithIdentityKeyRpcInputParamsResp], [PrivateKeyExportInput],
	// [SeedPhraseExportInput]
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
	DisplayName            string   `json:"display_name"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	Owner OwnerInputUnion `json:"owner"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	OwnerID OwnerIDInput `json:"owner_id"`
	// This field is from variant [WalletIntentResponseRequestDetailsBody].
	PolicyIDs PolicyInput `json:"policy_ids"`
	Name      string      `json:"name"`
	// This field is from variant [PolicyIntentResponseRequestDetailsBody].
	Rules []PolicyRuleRequestBody `json:"rules"`
	// This field is from variant [PolicyRuleRequestBody].
	Action PolicyAction `json:"action"`
	// This field is from variant [PolicyRuleRequestBody].
	Conditions []PolicyConditionUnion `json:"conditions"`
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
		DisplayName                          respjson.Field
		Owner                                respjson.Field
		OwnerID                              respjson.Field
		PolicyIDs                            respjson.Field
		Name                                 respjson.Field
		Rules                                respjson.Field
		Action                               respjson.Field
		Conditions                           respjson.Field
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
	ExportSeedPhrase   bool               `json:"export_seed_phrase"`
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
		ExportSeedPhrase        respjson.Field
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

// IntentResponseUnionCurrentResourceData is an implicit subunion of
// [IntentResponseUnion]. IntentResponseUnionCurrentResourceData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [IntentResponseUnion].
type IntentResponseUnionCurrentResourceData struct {
	ID string `json:"id"`
	// This field is from variant [Wallet].
	AdditionalSigners WalletAdditionalSigner `json:"additional_signers"`
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
	Custody     WalletCustodian `json:"custody"`
	DisplayName string          `json:"display_name"`
	// This field is from variant [Wallet].
	ExternalID string `json:"external_id"`
	// This field is from variant [Wallet].
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
	// This field is from variant [Policy].
	Rules []PolicyRuleResponse `json:"rules"`
	// This field is from variant [Policy].
	Version PolicyVersion `json:"version"`
	// This field is from variant [PolicyRuleResponse].
	Action PolicyAction `json:"action"`
	// This field is from variant [PolicyRuleResponse].
	Conditions []PolicyConditionUnion `json:"conditions"`
	// This field is from variant [PolicyRuleResponse].
	Method PolicyMethod `json:"method"`
	// This field is from variant [KeyQuorum].
	AuthorizationKeys []KeyQuorumAuthorizationKey `json:"authorization_keys"`
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
		DisplayName            respjson.Field
		ExternalID             respjson.Field
		PublicKey              respjson.Field
		Name                   respjson.Field
		Rules                  respjson.Field
		Version                respjson.Field
		Action                 respjson.Field
		Conditions             respjson.Field
		Method                 respjson.Field
		AuthorizationKeys      respjson.Field
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
	// The rules that apply to each method the policy covers.
	PolicyRuleRequestBody PolicyRuleRequestBodyParam
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentNewPolicyRuleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PolicyRuleRequestBody)
}
func (r *IntentNewPolicyRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// Name to assign to policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnionParam         `json:"owner,omitzero"`
	Rules []PolicyRuleRequestBodyParam `json:"rules,omitzero"`
	paramObj
}

func (r IntentUpdatePolicyParams) MarshalJSON() (data []byte, err error) {
	type shadow IntentUpdatePolicyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntentUpdatePolicyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntentUpdatePolicyRuleParams struct {
	// ID of the policy.
	PolicyID string `path:"policy_id" api:"required" json:"-"`
	// The rules that apply to each method the policy covers.
	PolicyRuleRequestBody PolicyRuleRequestBodyParam
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r IntentUpdatePolicyRuleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PolicyRuleRequestBody)
}
func (r *IntentUpdatePolicyRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
