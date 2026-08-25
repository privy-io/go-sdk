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
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
)

// Operations related to fiat onramping and offramping
//
// OrganizationKYBService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationKYBService] method instead.
type OrganizationKYBService struct {
	Options []option.RequestOption
}

// NewOrganizationKYBService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationKYBService(opts ...option.RequestOption) (r OrganizationKYBService) {
	r = OrganizationKYBService{}
	r.Options = opts
	return
}

// Returns KYB status for all providers the organization has initiated KYB with.
func (r *OrganizationKYBService) List(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *KYBStatusListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/kyb", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Generates a hosted KYB link for the organization and returns the current KYB
// status snapshot.
func (r *OrganizationKYBService) InitiateLinks(ctx context.Context, organizationID string, body OrganizationKYBInitiateLinksParams, opts ...option.RequestOption) (res *KYBStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/kyb/links", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generates a Bridge terms-of-service acceptance link for the organization.
func (r *OrganizationKYBService) InitiateTos(ctx context.Context, organizationID string, body OrganizationKYBInitiateTosParams, opts ...option.RequestOption) (res *KyxTosResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/kyb/tos", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type OrganizationKYBInitiateLinksParams struct {
	// Request body for initiating a hosted KYB flow for an organization.
	KYBLinksRequestBody KYBLinksRequestBody
	paramObj
}

func (r OrganizationKYBInitiateLinksParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KYBLinksRequestBody)
}
func (r *OrganizationKYBInitiateLinksParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationKYBInitiateTosParams struct {
	// Request body for initiating Terms of Service acceptance for an organization.
	KYBTosRequestBody KYBTosRequestBody
	paramObj
}

func (r OrganizationKYBInitiateTosParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KYBTosRequestBody)
}
func (r *OrganizationKYBInitiateTosParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
