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

	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
)

// Operations related to app settings and allowlist management
//
// AppAllowlistService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppAllowlistService] method instead.
type AppAllowlistService struct {
	Options []option.RequestOption
}

// NewAppAllowlistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppAllowlistService(opts ...option.RequestOption) (r AppAllowlistService) {
	r = AppAllowlistService{}
	r.Options = opts
	return
}

// Add a new entry to the allowlist for an app. The allowlist must be enabled.
func (r *AppAllowlistService) New(ctx context.Context, appID string, body AppAllowlistNewParams, opts ...option.RequestOption) (res *AllowlistEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("v1/apps/%s/allowlist", url.PathEscape(appID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all allowlist entries for an app. Returns the list of users allowed to
// access the app when the allowlist is enabled.
func (r *AppAllowlistService) List(ctx context.Context, appID string, opts ...option.RequestOption) (res *[]AllowlistEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("v1/apps/%s/allowlist", url.PathEscape(appID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove an entry from the allowlist for an app. The allowlist must be enabled.
func (r *AppAllowlistService) Delete(ctx context.Context, appID string, body AppAllowlistDeleteParams, opts ...option.RequestOption) (res *AllowlistDeletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("v1/apps/%s/allowlist", url.PathEscape(appID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AppAllowlistNewParams struct {
	// Input for adding or removing an allowlist entry. Discriminated by type.
	UserInviteInput UserInviteInputUnion
	paramObj
}

func (r AppAllowlistNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UserInviteInput)
}
func (r *AppAllowlistNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UserInviteInput)
}

type AppAllowlistDeleteParams struct {
	// Input for adding or removing an allowlist entry. Discriminated by type.
	UserInviteInput UserInviteInputUnion
	paramObj
}

func (r AppAllowlistDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UserInviteInput)
}
func (r *AppAllowlistDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UserInviteInput)
}
