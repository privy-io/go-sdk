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
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to key quorums
//
// KeyQuorumService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyQuorumService] method instead.
type KeyQuorumService struct {
	Options []option.RequestOption
}

// NewKeyQuorumService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKeyQuorumService(opts ...option.RequestOption) (r KeyQuorumService) {
	r = KeyQuorumService{}
	r.Options = opts
	return
}

// Create a new key quorum.
func (r *KeyQuorumService) New(ctx context.Context, body KeyQuorumNewParams, opts ...option.RequestOption) (res *KeyQuorum, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/key_quorums"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a key quorum by key quorum ID.
func (r *KeyQuorumService) Update(ctx context.Context, keyQuorumID KeyQuorumID, params KeyQuorumUpdateParams, opts ...option.RequestOption) (res *KeyQuorum, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/key_quorums/%s", keyQuorumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete a key quorum by key quorum ID.
func (r *KeyQuorumService) Delete(ctx context.Context, keyQuorumID KeyQuorumID, body KeyQuorumDeleteParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	if !param.IsOmitted(body.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", body.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(body.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", body.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/key_quorums/%s", keyQuorumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get a key quorum by ID.
func (r *KeyQuorumService) Get(ctx context.Context, keyQuorumID KeyQuorumID, opts ...option.RequestOption) (res *KeyQuorum, err error) {
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/key_quorums/%s", keyQuorumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A key quorum for authorizing wallet operations.
type KeyQuorum struct {
	ID                     string                      `json:"id" api:"required" format:"cuid2"`
	AuthorizationKeys      []KeyQuorumAuthorizationKey `json:"authorization_keys" api:"required"`
	AuthorizationThreshold float64                     `json:"authorization_threshold" api:"required"`
	DisplayName            string                      `json:"display_name" api:"required"`
	UserIDs                []string                    `json:"user_ids" api:"required"`
	// List of nested key quorum IDs that are members of this key quorum.
	KeyQuorumIDs []string `json:"key_quorum_ids"`
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
func (r KeyQuorum) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumAuthorizationKey struct {
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
func (r KeyQuorumAuthorizationKey) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumAuthorizationKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request input for creating a key quorum.
type KeyQuorumCreateRequestBody struct {
	// The number of keys that must sign for an action to be valid. Must be less than
	// or equal to total number of key quorum members.
	AuthorizationThreshold param.Opt[float64] `json:"authorization_threshold,omitzero"`
	DisplayName            param.Opt[string]  `json:"display_name,omitzero"`
	// List of key quorum IDs that should be members of this key quorum. Key quorums
	// can only be nested 1 level deep.
	KeyQuorumIDs []string `json:"key_quorum_ids,omitzero"`
	// List of P-256 public keys of the keys that should be authorized to sign on the
	// key quorum, in base64-encoded DER format.
	PublicKeys []string `json:"public_keys,omitzero"`
	// List of user IDs of the users that should be authorized to sign on the key
	// quorum.
	UserIDs []string `json:"user_ids,omitzero"`
	paramObj
}

func (r KeyQuorumCreateRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KeyQuorumCreateRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyQuorumCreateRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request input for updating an existing key quorum.
type KeyQuorumUpdateRequestBody struct {
	// The number of keys that must sign for an action to be valid. Must be less than
	// or equal to total number of key quorum members.
	AuthorizationThreshold float64 `json:"authorization_threshold"`
	DisplayName            string  `json:"display_name"`
	// List of key quorum IDs that should be members of this key quorum. Key quorums
	// can only be nested 1 level deep.
	KeyQuorumIDs []string `json:"key_quorum_ids"`
	// List of P-256 public keys of the keys that should be authorized to sign on the
	// key quorum, in base64-encoded DER format.
	PublicKeys []string `json:"public_keys"`
	// List of user IDs of the users that should be authorized to sign on the key
	// quorum.
	UserIDs []string `json:"user_ids"`
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
func (r KeyQuorumUpdateRequestBody) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumUpdateRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this KeyQuorumUpdateRequestBody to a
// KeyQuorumUpdateRequestBodyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// KeyQuorumUpdateRequestBodyParam.Overrides()
func (r KeyQuorumUpdateRequestBody) ToParam() KeyQuorumUpdateRequestBodyParam {
	return param.Override[KeyQuorumUpdateRequestBodyParam](json.RawMessage(r.RawJSON()))
}

// Request input for updating an existing key quorum.
type KeyQuorumUpdateRequestBodyParam struct {
	// The number of keys that must sign for an action to be valid. Must be less than
	// or equal to total number of key quorum members.
	AuthorizationThreshold param.Opt[float64] `json:"authorization_threshold,omitzero"`
	DisplayName            param.Opt[string]  `json:"display_name,omitzero"`
	// List of key quorum IDs that should be members of this key quorum. Key quorums
	// can only be nested 1 level deep.
	KeyQuorumIDs []string `json:"key_quorum_ids,omitzero"`
	// List of P-256 public keys of the keys that should be authorized to sign on the
	// key quorum, in base64-encoded DER format.
	PublicKeys []string `json:"public_keys,omitzero"`
	// List of user IDs of the users that should be authorized to sign on the key
	// quorum.
	UserIDs []string `json:"user_ids,omitzero"`
	paramObj
}

func (r KeyQuorumUpdateRequestBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow KeyQuorumUpdateRequestBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyQuorumUpdateRequestBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumNewParams struct {
	// Request input for creating a key quorum.
	KeyQuorumCreateRequestBody KeyQuorumCreateRequestBody
	paramObj
}

func (r KeyQuorumNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KeyQuorumCreateRequestBody)
}
func (r *KeyQuorumNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumUpdateParams struct {
	// Request input for updating an existing key quorum.
	KeyQuorumUpdateRequestBody KeyQuorumUpdateRequestBodyParam
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r KeyQuorumUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KeyQuorumUpdateRequestBody)
}
func (r *KeyQuorumUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumDeleteParams struct {
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}
