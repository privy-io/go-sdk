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
)

// Operations related to fiat onramping and offramping
//
// OrganizationExternalFiatAccountService contains methods and other services that
// help with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationExternalFiatAccountService] method instead.
type OrganizationExternalFiatAccountService struct {
	Options []option.RequestOption
}

// NewOrganizationExternalFiatAccountService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationExternalFiatAccountService(opts ...option.RequestOption) (r OrganizationExternalFiatAccountService) {
	r = OrganizationExternalFiatAccountService{}
	r.Options = opts
	return
}

// Creates an external fiat account linked to an organization for use in offramp
// transfers.
func (r *OrganizationExternalFiatAccountService) New(ctx context.Context, organizationID string, body OrganizationExternalFiatAccountNewParams, opts ...option.RequestOption) (res *OrganizationExternalFiatAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/external_fiat_accounts", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a list of external fiat accounts linked to an organization.
func (r *OrganizationExternalFiatAccountService) List(ctx context.Context, organizationID string, query OrganizationExternalFiatAccountListParams, opts ...option.RequestOption) (res *ListOrganizationExternalFiatAccountsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/external_fiat_accounts", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an external fiat account linked to an organization.
func (r *OrganizationExternalFiatAccountService) Delete(ctx context.Context, accountID string, body OrganizationExternalFiatAccountDeleteParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/external_fiat_accounts/%s", url.PathEscape(body.OrganizationID), url.PathEscape(accountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns a single external fiat account linked to an organization.
func (r *OrganizationExternalFiatAccountService) Get(ctx context.Context, accountID string, query OrganizationExternalFiatAccountGetParams, opts ...option.RequestOption) (res *OrganizationExternalFiatAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/external_fiat_accounts/%s", url.PathEscape(query.OrganizationID), url.PathEscape(accountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type OrganizationExternalFiatAccountNewParams struct {
	// Request body for creating a Bridge external fiat account.
	CreateExternalFiatAccountRequestBody CreateExternalFiatAccountRequestBody
	paramObj
}

func (r OrganizationExternalFiatAccountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateExternalFiatAccountRequestBody)
}
func (r *OrganizationExternalFiatAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationExternalFiatAccountListParams struct {
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `query:"provider,omitzero" api:"required" json:"-"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationExternalFiatAccountListParams]'s query
// parameters as `url.Values`.
func (r OrganizationExternalFiatAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationExternalFiatAccountDeleteParams struct {
	// The ID of the organization.
	OrganizationID string `path:"organization_id" api:"required" json:"-"`
	paramObj
}

type OrganizationExternalFiatAccountGetParams struct {
	// The ID of the organization.
	OrganizationID string `path:"organization_id" api:"required" json:"-"`
	paramObj
}
