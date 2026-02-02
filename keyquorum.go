// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/privy-api-client-go/internal/apijson"
	"github.com/stainless-sdks/privy-api-client-go/internal/requestconfig"
	"github.com/stainless-sdks/privy-api-client-go/option"
	"github.com/stainless-sdks/privy-api-client-go/packages/param"
	"github.com/stainless-sdks/privy-api-client-go/packages/respjson"
)

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
	return
}

// Update a key quorum by key quorum ID.
func (r *KeyQuorumService) Update(ctx context.Context, keyQuorumID string, params KeyQuorumUpdateParams, opts ...option.RequestOption) (res *KeyQuorum, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return
	}
	path := fmt.Sprintf("v1/key_quorums/%s", keyQuorumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Delete a key quorum by key quorum ID.
func (r *KeyQuorumService) Delete(ctx context.Context, keyQuorumID string, body KeyQuorumDeleteParams, opts ...option.RequestOption) (res *KeyQuorumDeleteResponse, err error) {
	if !param.IsOmitted(body.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", body.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return
	}
	path := fmt.Sprintf("v1/key_quorums/%s", keyQuorumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a key quorum by ID.
func (r *KeyQuorumService) Get(ctx context.Context, keyQuorumID string, opts ...option.RequestOption) (res *KeyQuorum, err error) {
	opts = slices.Concat(r.Options, opts)
	if keyQuorumID == "" {
		err = errors.New("missing required key_quorum_id parameter")
		return
	}
	path := fmt.Sprintf("v1/key_quorums/%s", keyQuorumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type KeyQuorum struct {
	ID                     string                      `json:"id,required"`
	AuthorizationKeys      []KeyQuorumAuthorizationKey `json:"authorization_keys,required"`
	AuthorizationThreshold float64                     `json:"authorization_threshold"`
	DisplayName            string                      `json:"display_name"`
	UserIDs                []string                    `json:"user_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AuthorizationKeys      respjson.Field
		AuthorizationThreshold respjson.Field
		DisplayName            respjson.Field
		UserIDs                respjson.Field
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
	DisplayName string `json:"display_name,required"`
	PublicKey   string `json:"public_key,required"`
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

type KeyQuorumDeleteResponse struct {
	// Whether the key quorum was deleted successfully.
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyQuorumDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *KeyQuorumDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumNewParams struct {
	AuthorizationThreshold param.Opt[float64] `json:"authorization_threshold,omitzero"`
	DisplayName            param.Opt[string]  `json:"display_name,omitzero"`
	PublicKeys             []string           `json:"public_keys,omitzero"`
	UserIDs                []string           `json:"user_ids,omitzero"`
	paramObj
}

func (r KeyQuorumNewParams) MarshalJSON() (data []byte, err error) {
	type shadow KeyQuorumNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyQuorumNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumUpdateParams struct {
	AuthorizationThreshold param.Opt[float64] `json:"authorization_threshold,omitzero"`
	DisplayName            param.Opt[string]  `json:"display_name,omitzero"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	PublicKeys                  []string          `json:"public_keys,omitzero"`
	UserIDs                     []string          `json:"user_ids,omitzero"`
	paramObj
}

func (r KeyQuorumUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow KeyQuorumUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyQuorumUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyQuorumDeleteParams struct {
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}
