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

// Operations related to organizations
//
// OrganizationService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	// Operations related to fiat onramping and offramping
	ExternalFiatAccounts OrganizationExternalFiatAccountService
	// Operations related to fiat onramping and offramping
	KYB OrganizationKYBService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	r.ExternalFiatAccounts = NewOrganizationExternalFiatAccountService(opts...)
	r.KYB = NewOrganizationKYBService(opts...)
	return
}

// Create an organization in an app.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an organization by ID.
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List organizations in an app.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *pagination.Cursor[Organization], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/organizations"
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

// List organizations in an app.
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Organization] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an organization by ID.
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get an organization by ID.
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A Privy organization object.
type Organization struct {
	// Unique organization identifier
	ID string `json:"id" api:"required"`
	// Unix timestamp when the organization was created
	CreatedAt float64 `json:"created_at" api:"required"`
	// A unique identifier for a key quorum.
	DefaultKeyQuorumID KeyQuorumID `json:"default_key_quorum_id" api:"required" format:"cuid2"`
	// Defaults applied to wallets created with the organization as their entity, keyed
	// by chain type.
	DefaultWalletConfigurations OrganizationDefaultWalletConfigurationsResp `json:"default_wallet_configurations" api:"required"`
	// Organization display name
	DisplayName string `json:"display_name" api:"required"`
	// Unix timestamp when the organization was last updated
	UpdatedAt float64 `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		CreatedAt                   respjson.Field
		DefaultKeyQuorumID          respjson.Field
		DefaultWalletConfigurations respjson.Field
		DisplayName                 respjson.Field
		UpdatedAt                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organization) RawJSON() string { return r.JSON.raw }
func (r *Organization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating an organization.
//
// The properties DefaultKeyQuorumID, DisplayName are required.
type OrganizationCreateRequestBody struct {
	// A unique identifier for a key quorum.
	DefaultKeyQuorumID KeyQuorumID `json:"default_key_quorum_id" api:"required" format:"cuid2"`
	DisplayName        string      `json:"display_name" api:"required"`
	// Defaults applied to wallets created with the organization as their entity, keyed
	// by chain type.
	DefaultWalletConfigurations OrganizationDefaultWalletConfigurations `json:"default_wallet_configurations,omitzero"`
	paramObj
}

func (r OrganizationCreateRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationCreateRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationCreateRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults applied to wallets created with the organization as their entity, keyed
// by chain type.
type OrganizationDefaultWalletConfigurationsResp struct {
	// Default policies and additional signers for wallets of one chain type.
	Ethereum OrganizationWalletConfigurationResp `json:"ethereum"`
	// Default policies and additional signers for wallets of one chain type.
	Solana OrganizationWalletConfigurationResp `json:"solana"`
	// Default policies and additional signers for wallets of one chain type.
	Tron        OrganizationWalletConfigurationResp            `json:"tron"`
	ExtraFields map[string]OrganizationWalletConfigurationResp `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ethereum    respjson.Field
		Solana      respjson.Field
		Tron        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationDefaultWalletConfigurationsResp) RawJSON() string { return r.JSON.raw }
func (r *OrganizationDefaultWalletConfigurationsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OrganizationDefaultWalletConfigurationsResp to a
// OrganizationDefaultWalletConfigurations.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OrganizationDefaultWalletConfigurations.Overrides()
func (r OrganizationDefaultWalletConfigurationsResp) ToParam() OrganizationDefaultWalletConfigurations {
	return param.Override[OrganizationDefaultWalletConfigurations](json.RawMessage(r.RawJSON()))
}

// Defaults applied to wallets created with the organization as their entity, keyed
// by chain type.
type OrganizationDefaultWalletConfigurations struct {
	// Default policies and additional signers for wallets of one chain type.
	Ethereum OrganizationWalletConfiguration `json:"ethereum,omitzero"`
	// Default policies and additional signers for wallets of one chain type.
	Solana OrganizationWalletConfiguration `json:"solana,omitzero"`
	// Default policies and additional signers for wallets of one chain type.
	Tron        OrganizationWalletConfiguration            `json:"tron,omitzero"`
	ExtraFields map[string]OrganizationWalletConfiguration `json:"-"`
	paramObj
}

func (r OrganizationDefaultWalletConfigurations) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationDefaultWalletConfigurations
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OrganizationDefaultWalletConfigurations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Changes to the defaults, keyed by chain type. Omitted chains and fields are
// unchanged, and `null` removes a chain.
type OrganizationDefaultWalletConfigurationsUpdate struct {
	// Default policies and additional signers for wallets of one chain type.
	Ethereum OrganizationWalletConfiguration `json:"ethereum,omitzero"`
	// Default policies and additional signers for wallets of one chain type.
	Solana OrganizationWalletConfiguration `json:"solana,omitzero"`
	// Default policies and additional signers for wallets of one chain type.
	Tron        OrganizationWalletConfiguration            `json:"tron,omitzero"`
	ExtraFields map[string]OrganizationWalletConfiguration `json:"-"`
	paramObj
}

func (r OrganizationDefaultWalletConfigurationsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationDefaultWalletConfigurationsUpdate
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *OrganizationDefaultWalletConfigurationsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for updating an organization.
type OrganizationUpdateRequestBody struct {
	// A unique identifier for a key quorum.
	DefaultKeyQuorumID param.Opt[KeyQuorumID] `json:"default_key_quorum_id,omitzero" format:"cuid2"`
	DisplayName        param.Opt[string]      `json:"display_name,omitzero"`
	// Changes to the defaults, keyed by chain type. Omitted chains and fields are
	// unchanged, and `null` removes a chain.
	DefaultWalletConfigurations OrganizationDefaultWalletConfigurationsUpdate `json:"default_wallet_configurations,omitzero"`
	paramObj
}

func (r OrganizationUpdateRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUpdateRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUpdateRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default policies and additional signers for wallets of one chain type.
type OrganizationWalletConfigurationResp struct {
	// Additional signers for the wallet.
	AdditionalSigners AdditionalSignerInputResp `json:"additional_signers"`
	// An optional list of up to one policy ID to enforce on the wallet.
	PolicyIDs PolicyInput `json:"policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalSigners respjson.Field
		PolicyIDs         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationWalletConfigurationResp) RawJSON() string { return r.JSON.raw }
func (r *OrganizationWalletConfigurationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OrganizationWalletConfigurationResp to a
// OrganizationWalletConfiguration.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OrganizationWalletConfiguration.Overrides()
func (r OrganizationWalletConfigurationResp) ToParam() OrganizationWalletConfiguration {
	return param.Override[OrganizationWalletConfiguration](json.RawMessage(r.RawJSON()))
}

// Default policies and additional signers for wallets of one chain type.
type OrganizationWalletConfiguration struct {
	// Additional signers for the wallet.
	AdditionalSigners AdditionalSignerInput `json:"additional_signers,omitzero"`
	// An optional list of up to one policy ID to enforce on the wallet.
	PolicyIDs PolicyInput `json:"policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r OrganizationWalletConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationWalletConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationWalletConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned when listing organizations for an app.
type OrganizationsListResponse struct {
	Data       []Organization `json:"data" api:"required"`
	NextCursor string         `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationsListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationNewParams struct {
	// Request body for creating an organization.
	OrganizationCreateRequestBody OrganizationCreateRequestBody
	paramObj
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.OrganizationCreateRequestBody)
}
func (r *OrganizationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUpdateParams struct {
	// Request body for updating an organization.
	OrganizationUpdateRequestBody OrganizationUpdateRequestBody
	paramObj
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.OrganizationUpdateRequestBody)
}
func (r *OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListParams struct {
	// Cursor returned by the previous page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
