// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/apiquery"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/pagination"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to policies
//
// PolicyConditionSetService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyConditionSetService] method instead.
type PolicyConditionSetService struct {
	Options []option.RequestOption
}

// NewPolicyConditionSetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPolicyConditionSetService(opts ...option.RequestOption) (r PolicyConditionSetService) {
	r = PolicyConditionSetService{}
	r.Options = opts
	return
}

// List condition sets in an app.
func (r *PolicyConditionSetService) List(ctx context.Context, query PolicyConditionSetListParams, opts ...option.RequestOption) (res *pagination.Cursor[ConditionSet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/condition_sets"
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

// List condition sets in an app.
func (r *PolicyConditionSetService) ListAutoPaging(ctx context.Context, query PolicyConditionSetListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ConditionSet] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Paginated list of condition sets in an app.
type ConditionSetsResponse struct {
	// Condition sets in this page.
	Data []ConditionSet `json:"data" api:"required"`
	// Cursor for the next page. Null when there are no further pages.
	NextCursor string `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConditionSetsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConditionSetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyConditionSetListParams struct {
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Cursor returned by the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PolicyConditionSetListParams]'s query parameters as
// `url.Values`.
func (r PolicyConditionSetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
