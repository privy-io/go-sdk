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

type KeyQuorumID = string

type OwnerIDInput = string

// OwnerInputUnionResp contains all possible properties and values from
// [OwnerInputUserResp], [OwnerInputPublicKeyResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OwnerInputUnionResp struct {
	// This field is from variant [OwnerInputUserResp].
	UserID string `json:"user_id"`
	// This field is from variant [OwnerInputPublicKeyResp].
	PublicKey P256PublicKey `json:"public_key"`
	JSON      struct {
		UserID    respjson.Field
		PublicKey respjson.Field
		raw       string
	} `json:"-"`
}

func (u OwnerInputUnionResp) AsOwnerInputUser() (v OwnerInputUserResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OwnerInputUnionResp) AsOwnerInputPublicKey() (v OwnerInputPublicKeyResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OwnerInputUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OwnerInputUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerInputUnionResp to a OwnerInputUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerInputUnion.Overrides()
func (r OwnerInputUnionResp) ToParam() OwnerInputUnion {
	return param.Override[OwnerInputUnion](json.RawMessage(r.RawJSON()))
}

func OwnerInputOfOwnerInputUser(userID string) OwnerInputUnion {
	var variant OwnerInputUser
	variant.UserID = userID
	return OwnerInputUnion{OfOwnerInputUser: &variant}
}

func OwnerInputOfOwnerInputPublicKey(publicKey P256PublicKey) OwnerInputUnion {
	var variant OwnerInputPublicKey
	variant.PublicKey = publicKey
	return OwnerInputUnion{OfOwnerInputPublicKey: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OwnerInputUnion struct {
	OfOwnerInputUser      *OwnerInputUser      `json:",omitzero,inline"`
	OfOwnerInputPublicKey *OwnerInputPublicKey `json:",omitzero,inline"`
	paramUnion
}

func (u OwnerInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOwnerInputUser, u.OfOwnerInputPublicKey)
}
func (u *OwnerInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Owner input specifying a P-256 public key.
type OwnerInputPublicKeyResp struct {
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
func (r OwnerInputPublicKeyResp) RawJSON() string { return r.JSON.raw }
func (r *OwnerInputPublicKeyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerInputPublicKeyResp to a OwnerInputPublicKey.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerInputPublicKey.Overrides()
func (r OwnerInputPublicKeyResp) ToParam() OwnerInputPublicKey {
	return param.Override[OwnerInputPublicKey](json.RawMessage(r.RawJSON()))
}

// Owner input specifying a P-256 public key.
//
// The property PublicKey is required.
type OwnerInputPublicKey struct {
	// A P-256 (secp256r1) public key.
	PublicKey P256PublicKey `json:"public_key" api:"required"`
	paramObj
}

func (r OwnerInputPublicKey) MarshalJSON() (data []byte, err error) {
	type shadow OwnerInputPublicKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OwnerInputPublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Owner input specifying a Privy user ID.
type OwnerInputUserResp struct {
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnerInputUserResp) RawJSON() string { return r.JSON.raw }
func (r *OwnerInputUserResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerInputUserResp to a OwnerInputUser.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerInputUser.Overrides()
func (r OwnerInputUserResp) ToParam() OwnerInputUser {
	return param.Override[OwnerInputUser](json.RawMessage(r.RawJSON()))
}

// Owner input specifying a Privy user ID.
//
// The property UserID is required.
type OwnerInputUser struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r OwnerInputUser) MarshalJSON() (data []byte, err error) {
	type shadow OwnerInputUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OwnerInputUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type P256PublicKey = string

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
