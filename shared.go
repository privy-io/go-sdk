// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// SharedService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSharedService] method instead.
type SharedService struct {
	Options []option.RequestOption
}

// NewSharedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSharedService(opts ...option.RequestOption) (r SharedService) {
	r = SharedService{}
	r.Options = opts
	return
}

// A simple success response.
type SuccessResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *SuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type P256PublicKey = string

type KeyQuorumID = string

// Owner input specifying a Privy user ID.
type OwnerInputUser struct {
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnerInputUser) RawJSON() string { return r.JSON.raw }
func (r *OwnerInputUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerInputUser to a OwnerInputUserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerInputUserParam.Overrides()
func (r OwnerInputUser) ToParam() OwnerInputUserParam {
	return param.Override[OwnerInputUserParam](json.RawMessage(r.RawJSON()))
}

// Owner input specifying a Privy user ID.
//
// The property UserID is required.
type OwnerInputUserParam struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r OwnerInputUserParam) MarshalJSON() (data []byte, err error) {
	type shadow OwnerInputUserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OwnerInputUserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Owner input specifying a P-256 public key.
type OwnerInputPublicKey struct {
	// A P-256 (secp256r1) public key.
	PublicKey P256PublicKey `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnerInputPublicKey) RawJSON() string { return r.JSON.raw }
func (r *OwnerInputPublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerInputPublicKey to a OwnerInputPublicKeyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerInputPublicKeyParam.Overrides()
func (r OwnerInputPublicKey) ToParam() OwnerInputPublicKeyParam {
	return param.Override[OwnerInputPublicKeyParam](json.RawMessage(r.RawJSON()))
}

// Owner input specifying a P-256 public key.
//
// The property PublicKey is required.
type OwnerInputPublicKeyParam struct {
	// A P-256 (secp256r1) public key.
	PublicKey P256PublicKey `json:"public_key" api:"required"`
	paramObj
}

func (r OwnerInputPublicKeyParam) MarshalJSON() (data []byte, err error) {
	type shadow OwnerInputPublicKeyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OwnerInputPublicKeyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OwnerInputUnion contains all possible properties and values from
// [OwnerInputUser], [OwnerInputPublicKey].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OwnerInputUnion struct {
	// This field is from variant [OwnerInputUser].
	UserID string `json:"user_id"`
	// This field is from variant [OwnerInputPublicKey].
	PublicKey P256PublicKey `json:"public_key"`
	JSON      struct {
		UserID    respjson.Field
		PublicKey respjson.Field
		raw       string
	} `json:"-"`
}

func (u OwnerInputUnion) AsOwnerInputUser() (v OwnerInputUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OwnerInputUnion) AsOwnerInputPublicKey() (v OwnerInputPublicKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OwnerInputUnion) RawJSON() string { return u.JSON.raw }

func (r *OwnerInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerInputUnion to a OwnerInputUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerInputUnionParam.Overrides()
func (r OwnerInputUnion) ToParam() OwnerInputUnionParam {
	return param.Override[OwnerInputUnionParam](json.RawMessage(r.RawJSON()))
}

func OwnerInputParamOfOwnerInputUser(userID string) OwnerInputUnionParam {
	var variant OwnerInputUserParam
	variant.UserID = userID
	return OwnerInputUnionParam{OfOwnerInputUser: &variant}
}

func OwnerInputParamOfOwnerInputPublicKey(publicKey P256PublicKey) OwnerInputUnionParam {
	var variant OwnerInputPublicKeyParam
	variant.PublicKey = publicKey
	return OwnerInputUnionParam{OfOwnerInputPublicKey: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OwnerInputUnionParam struct {
	OfOwnerInputUser      *OwnerInputUserParam      `json:",omitzero,inline"`
	OfOwnerInputPublicKey *OwnerInputPublicKeyParam `json:",omitzero,inline"`
	paramUnion
}

func (u OwnerInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOwnerInputUser, u.OfOwnerInputPublicKey)
}
func (u *OwnerInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OwnerIDInput = string
