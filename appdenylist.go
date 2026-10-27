// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
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
)

// Operations related to app settings and allowlist management
//
// AppDenylistService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppDenylistService] method instead.
type AppDenylistService struct {
	Options []option.RequestOption
}

// NewAppDenylistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppDenylistService(opts ...option.RequestOption) (r AppDenylistService) {
	r = AppDenylistService{}
	r.Options = opts
	return
}

// Add a new entry to the denylist for an app. Adding an entry that already exists
// returns the existing entry.
func (r *AppDenylistService) New(ctx context.Context, appID string, body AppDenylistNewParams, opts ...option.RequestOption) (res *DenylistEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/apps/%s/denylist", url.PathEscape(appID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a page of denylist entries for an app, which block matching users from
// logging in or linking the account.
func (r *AppDenylistService) List(ctx context.Context, appID string, query AppDenylistListParams, opts ...option.RequestOption) (res *pagination.Cursor[DenylistEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/apps/%s/denylist", url.PathEscape(appID))
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

// Get a page of denylist entries for an app, which block matching users from
// logging in or linking the account.
func (r *AppDenylistService) ListAutoPaging(ctx context.Context, appID string, query AppDenylistListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DenylistEntry] {
	return pagination.NewCursorAutoPager(r.List(ctx, appID, query, opts...))
}

// Remove an entry from the denylist for an app by its ID. Get entry IDs from the
// list denylist entries endpoint.
func (r *AppDenylistService) Delete(ctx context.Context, appID string, body AppDenylistDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/apps/%s/denylist", url.PathEscape(appID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

type AppDenylistNewParams struct {
	// Input for adding a denylist entry. Discriminated by type.
	DenylistEntryCreateRequestBody DenylistEntryCreateRequestBodyUnion
	paramObj
}

func (r AppDenylistNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DenylistEntryCreateRequestBody)
}
func (r *AppDenylistNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDenylistListParams struct {
	// Cursor returned by the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppDenylistListParams]'s query parameters as `url.Values`.
func (r AppDenylistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AppDenylistDeleteParams struct {
	// Input for removing a denylist entry.
	DenylistEntryDeleteRequestBody DenylistEntryDeleteRequestBody
	paramObj
}

func (r AppDenylistDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DenylistEntryDeleteRequestBody)
}
func (r *AppDenylistDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
