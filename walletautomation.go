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

// Operations related to wallet automations
//
// WalletAutomationService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletAutomationService] method instead.
type WalletAutomationService struct {
	Options []option.RequestOption
}

// NewWalletAutomationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWalletAutomationService(opts ...option.RequestOption) (r WalletAutomationService) {
	r = WalletAutomationService{}
	r.Options = opts
	return
}

// Create a new wallet automation that triggers actions on deposit events.
func (r *WalletAutomationService) New(ctx context.Context, body WalletAutomationNewParams, opts ...option.RequestOption) (res *WalletAutomationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallet_automations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a wallet automation by ID.
func (r *WalletAutomationService) Update(ctx context.Context, automationID string, body WalletAutomationUpdateParams, opts ...option.RequestOption) (res *WalletAutomationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if automationID == "" {
		err = errors.New("missing required automation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallet_automations/%s", url.PathEscape(automationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all wallet automations for your app, with optional filtering by wallet.
func (r *WalletAutomationService) List(ctx context.Context, query WalletAutomationListParams, opts ...option.RequestOption) (res *pagination.Cursor[WalletAutomationResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/wallet_automations"
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

// List all wallet automations for your app, with optional filtering by wallet.
func (r *WalletAutomationService) ListAutoPaging(ctx context.Context, query WalletAutomationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[WalletAutomationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a wallet automation by ID.
func (r *WalletAutomationService) Delete(ctx context.Context, automationID string, opts ...option.RequestOption) (res *WalletAutomationSuccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if automationID == "" {
		err = errors.New("missing required automation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallet_automations/%s", url.PathEscape(automationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get a wallet automation by ID.
func (r *WalletAutomationService) Get(ctx context.Context, automationID string, opts ...option.RequestOption) (res *WalletAutomationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if automationID == "" {
		err = errors.New("missing required automation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallet_automations/%s", url.PathEscape(automationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all wallet automation execution records, with optional filtering by wallet.
func (r *WalletAutomationService) ListExecutions(ctx context.Context, query WalletAutomationListExecutionsParams, opts ...option.RequestOption) (res *pagination.Cursor[WalletAutomationExecutionResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/wallet_automations/executions"
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

// List all wallet automation execution records, with optional filtering by wallet.
func (r *WalletAutomationService) ListExecutionsAutoPaging(ctx context.Context, query WalletAutomationListExecutionsParams, opts ...option.RequestOption) *pagination.CursorAutoPager[WalletAutomationExecutionResponse] {
	return pagination.NewCursorAutoPager(r.ListExecutions(ctx, query, opts...))
}

// Re-checks a wallet (identified by wallet_id or deposit_address) for funds
// matching its wallet automation configs and triggers an automation run if a match
// is found. Use this to recover a deposit whose automation trigger was missed or
// failed.
func (r *WalletAutomationService) Reindex(ctx context.Context, body WalletAutomationReindexParams, opts ...option.RequestOption) (res *WalletAutomationReindexResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallet_automations/reindex"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// AutomationActionConfigUnion contains all possible properties and values from
// [AutomationSwapActionConfig], [AutomationEarnDepositActionConfig].
//
// Use the [AutomationActionConfigUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationActionConfigUnion struct {
	// This field is from variant [AutomationSwapActionConfig].
	DestinationChainAsset AutomationDestinationAssetResp `json:"destination_chain_asset"`
	// Any of "swap", "earn_deposit".
	Type string `json:"type"`
	// This field is from variant [AutomationEarnDepositActionConfig].
	VaultID string `json:"vault_id"`
	JSON    struct {
		DestinationChainAsset respjson.Field
		Type                  respjson.Field
		VaultID               respjson.Field
		raw                   string
	} `json:"-"`
}

// anyAutomationActionConfig is implemented by each variant of
// [AutomationActionConfigUnion] to add type safety for the return type of
// [AutomationActionConfigUnion.AsAny]
type anyAutomationActionConfig interface {
	implAutomationActionConfigUnion()
}

func (AutomationSwapActionConfig) implAutomationActionConfigUnion()        {}
func (AutomationEarnDepositActionConfig) implAutomationActionConfigUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AutomationActionConfigUnion.AsAny().(type) {
//	case privyclient.AutomationSwapActionConfig:
//	case privyclient.AutomationEarnDepositActionConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AutomationActionConfigUnion) AsAny() anyAutomationActionConfig {
	switch u.Type {
	case "swap":
		return u.AsSwap()
	case "earn_deposit":
		return u.AsEarnDeposit()
	}
	return nil
}

func (u AutomationActionConfigUnion) AsSwap() (v AutomationSwapActionConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationActionConfigUnion) AsEarnDeposit() (v AutomationEarnDepositActionConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationActionConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationActionConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func AutomationActionConfigInputOfSwap(destinationChainAsset AutomationDestinationAssetInput) AutomationActionConfigInputUnion {
	var swap AutomationSwapActionConfigInput
	swap.DestinationChainAsset = destinationChainAsset
	return AutomationActionConfigInputUnion{OfSwap: &swap}
}

func AutomationActionConfigInputOfEarnDeposit(vaultID string) AutomationActionConfigInputUnion {
	var earnDeposit AutomationEarnDepositActionConfigInput
	earnDeposit.VaultID = vaultID
	return AutomationActionConfigInputUnion{OfEarnDeposit: &earnDeposit}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationActionConfigInputUnion struct {
	OfSwap        *AutomationSwapActionConfigInput        `json:",omitzero,inline"`
	OfEarnDeposit *AutomationEarnDepositActionConfigInput `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationActionConfigInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSwap, u.OfEarnDeposit)
}
func (u *AutomationActionConfigInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[AutomationActionConfigInputUnion](
		"type",
		apijson.Discriminator[AutomationSwapActionConfigInput]("swap"),
		apijson.Discriminator[AutomationEarnDepositActionConfigInput]("earn_deposit"),
	)
}

// AutomationAssetFilterUnion contains all possible properties and values from
// [AutomationAssetFilterAllResp], [AutomationAssetFilterInclude],
// [AutomationAssetFilterExclude].
//
// Use the [AutomationAssetFilterUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomationAssetFilterUnion struct {
	// Any of "all", "include", "exclude".
	Mode   string                    `json:"mode"`
	Values []AutomationAssetSpecResp `json:"values"`
	JSON   struct {
		Mode   respjson.Field
		Values respjson.Field
		raw    string
	} `json:"-"`
}

// anyAutomationAssetFilter is implemented by each variant of
// [AutomationAssetFilterUnion] to add type safety for the return type of
// [AutomationAssetFilterUnion.AsAny]
type anyAutomationAssetFilter interface {
	implAutomationAssetFilterUnion()
}

func (AutomationAssetFilterAllResp) implAutomationAssetFilterUnion() {}
func (AutomationAssetFilterInclude) implAutomationAssetFilterUnion() {}
func (AutomationAssetFilterExclude) implAutomationAssetFilterUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AutomationAssetFilterUnion.AsAny().(type) {
//	case privyclient.AutomationAssetFilterAllResp:
//	case privyclient.AutomationAssetFilterInclude:
//	case privyclient.AutomationAssetFilterExclude:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AutomationAssetFilterUnion) AsAny() anyAutomationAssetFilter {
	switch u.Mode {
	case "all":
		return u.AsAll()
	case "include":
		return u.AsInclude()
	case "exclude":
		return u.AsExclude()
	}
	return nil
}

func (u AutomationAssetFilterUnion) AsAll() (v AutomationAssetFilterAllResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationAssetFilterUnion) AsInclude() (v AutomationAssetFilterInclude) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomationAssetFilterUnion) AsExclude() (v AutomationAssetFilterExclude) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomationAssetFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomationAssetFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match all assets.
type AutomationAssetFilterAllResp struct {
	// Any of "all".
	Mode AutomationAssetFilterAllMode `json:"mode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationAssetFilterAllResp) RawJSON() string { return r.JSON.raw }
func (r *AutomationAssetFilterAllResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AutomationAssetFilterAllResp to a
// AutomationAssetFilterAll.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutomationAssetFilterAll.Overrides()
func (r AutomationAssetFilterAllResp) ToParam() AutomationAssetFilterAll {
	return param.Override[AutomationAssetFilterAll](json.RawMessage(r.RawJSON()))
}

type AutomationAssetFilterAllMode string

const (
	AutomationAssetFilterAllModeAll AutomationAssetFilterAllMode = "all"
)

// Match all assets.
//
// The property Mode is required.
type AutomationAssetFilterAll struct {
	// Any of "all".
	Mode AutomationAssetFilterAllMode `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r AutomationAssetFilterAll) MarshalJSON() (data []byte, err error) {
	type shadow AutomationAssetFilterAll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationAssetFilterAll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match all assets except the specified ones.
type AutomationAssetFilterExclude struct {
	// Any of "exclude".
	Mode   AutomationAssetFilterExcludeMode `json:"mode" api:"required"`
	Values []AutomationAssetSpecResp        `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationAssetFilterExclude) RawJSON() string { return r.JSON.raw }
func (r *AutomationAssetFilterExclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationAssetFilterExcludeMode string

const (
	AutomationAssetFilterExcludeModeExclude AutomationAssetFilterExcludeMode = "exclude"
)

// Match only the specified assets.
type AutomationAssetFilterInclude struct {
	// Any of "include".
	Mode   AutomationAssetFilterIncludeMode `json:"mode" api:"required"`
	Values []AutomationAssetSpecResp        `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationAssetFilterInclude) RawJSON() string { return r.JSON.raw }
func (r *AutomationAssetFilterInclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationAssetFilterIncludeMode string

const (
	AutomationAssetFilterIncludeModeInclude AutomationAssetFilterIncludeMode = "include"
)

func AutomationAssetFilterInputOfAll(mode AutomationAssetFilterAllMode) AutomationAssetFilterInputUnion {
	var all AutomationAssetFilterAll
	all.Mode = mode
	return AutomationAssetFilterInputUnion{OfAll: &all}
}

func AutomationAssetFilterInputOfInclude(values []AutomationAssetSpecInput) AutomationAssetFilterInputUnion {
	var include AutomationAssetFilterInputInclude
	include.Values = values
	return AutomationAssetFilterInputUnion{OfInclude: &include}
}

func AutomationAssetFilterInputOfExclude(values []AutomationAssetSpecInput) AutomationAssetFilterInputUnion {
	var exclude AutomationAssetFilterInputExclude
	exclude.Values = values
	return AutomationAssetFilterInputUnion{OfExclude: &exclude}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationAssetFilterInputUnion struct {
	OfAll     *AutomationAssetFilterAll          `json:",omitzero,inline"`
	OfInclude *AutomationAssetFilterInputInclude `json:",omitzero,inline"`
	OfExclude *AutomationAssetFilterInputExclude `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationAssetFilterInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAll, u.OfInclude, u.OfExclude)
}
func (u *AutomationAssetFilterInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[AutomationAssetFilterInputUnion](
		"mode",
		apijson.Discriminator[AutomationAssetFilterAll]("all"),
		apijson.Discriminator[AutomationAssetFilterInputInclude]("include"),
		apijson.Discriminator[AutomationAssetFilterInputExclude]("exclude"),
	)
}

// Match all assets except the specified ones (input form with alias support).
//
// The properties Mode, Values are required.
type AutomationAssetFilterInputExclude struct {
	// Any of "exclude".
	Mode   AutomationAssetFilterInputExcludeMode `json:"mode,omitzero" api:"required"`
	Values []AutomationAssetSpecInput            `json:"values,omitzero" api:"required"`
	paramObj
}

func (r AutomationAssetFilterInputExclude) MarshalJSON() (data []byte, err error) {
	type shadow AutomationAssetFilterInputExclude
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationAssetFilterInputExclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationAssetFilterInputExcludeMode string

const (
	AutomationAssetFilterInputExcludeModeExclude AutomationAssetFilterInputExcludeMode = "exclude"
)

// Match only the specified assets (input form with alias support).
//
// The properties Mode, Values are required.
type AutomationAssetFilterInputInclude struct {
	// Any of "include".
	Mode   AutomationAssetFilterInputIncludeMode `json:"mode,omitzero" api:"required"`
	Values []AutomationAssetSpecInput            `json:"values,omitzero" api:"required"`
	paramObj
}

func (r AutomationAssetFilterInputInclude) MarshalJSON() (data []byte, err error) {
	type shadow AutomationAssetFilterInputInclude
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationAssetFilterInputInclude) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationAssetFilterInputIncludeMode string

const (
	AutomationAssetFilterInputIncludeModeInclude AutomationAssetFilterInputIncludeMode = "include"
)

// An asset identified by contract address on a specific chain (CAIP-2). Either
// field may be "_": asset_address: "_" matches any asset on the chain; caip2: "\*"
// matches the asset on any chain (in this case asset_address holds the asset
// symbol id, e.g. "usdc" or "eth", not a contract address).
type AutomationAssetSpecResp struct {
	AssetAddress string `json:"asset_address" api:"required"`
	Caip2        string `json:"caip2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetAddress respjson.Field
		Caip2        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationAssetSpecResp) RawJSON() string { return r.JSON.raw }
func (r *AutomationAssetSpecResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AutomationAssetSpecResp to a AutomationAssetSpec.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutomationAssetSpec.Overrides()
func (r AutomationAssetSpecResp) ToParam() AutomationAssetSpec {
	return param.Override[AutomationAssetSpec](json.RawMessage(r.RawJSON()))
}

// An asset identified by contract address on a specific chain (CAIP-2). Either
// field may be "_": asset_address: "_" matches any asset on the chain; caip2: "\*"
// matches the asset on any chain (in this case asset_address holds the asset
// symbol id, e.g. "usdc" or "eth", not a contract address).
//
// The properties AssetAddress, Caip2 are required.
type AutomationAssetSpec struct {
	AssetAddress string `json:"asset_address" api:"required"`
	Caip2        string `json:"caip2" api:"required"`
	paramObj
}

func (r AutomationAssetSpec) MarshalJSON() (data []byte, err error) {
	type shadow AutomationAssetSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationAssetSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An asset spec accepting either raw identifiers (asset_address, caip2) or
// human-readable aliases (asset, chain). Exactly one of asset_address or asset
// must be provided; at most one of caip2 or chain may be provided. Use "_" for
// asset_address or asset to match any asset on a given chain (chain is then
// required). Omitting chain/caip2 (or passing "_" for either) matches the
// specified asset on any chain.
type AutomationAssetSpecInput struct {
	Asset param.Opt[string] `json:"asset,omitzero"`
	Chain param.Opt[string] `json:"chain,omitzero"`
	AutomationAssetSpec
}

func (r AutomationAssetSpecInput) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*AutomationAssetSpecInput
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Full configuration for a wallet automation (trigger + action).
type AutomationConfig struct {
	// Configuration for an automation action.
	Action AutomationActionConfigUnion `json:"action" api:"required"`
	// Trigger configuration for deposit events.
	Trigger AutomationTriggerConfig `json:"trigger" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Trigger     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationConfig) RawJSON() string { return r.JSON.raw }
func (r *AutomationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full configuration for a wallet automation (trigger + action) accepting
// human-readable aliases.
//
// The properties Action, Trigger are required.
type AutomationConfigInput struct {
	// Configuration for an automation action (input form with alias support).
	Action AutomationActionConfigInputUnion `json:"action,omitzero" api:"required"`
	// Trigger configuration for deposit events (input form with alias support).
	Trigger AutomationTriggerConfigInput `json:"trigger,omitzero" api:"required"`
	paramObj
}

func (r AutomationConfigInput) MarshalJSON() (data []byte, err error) {
	type shadow AutomationConfigInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationConfigInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Destination asset identified by contract address on a specific chain (CAIP-2).
type AutomationDestinationAssetResp struct {
	AssetAddress string `json:"asset_address" api:"required"`
	Caip2        string `json:"caip2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetAddress respjson.Field
		Caip2        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationDestinationAssetResp) RawJSON() string { return r.JSON.raw }
func (r *AutomationDestinationAssetResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AutomationDestinationAssetResp to a
// AutomationDestinationAsset.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutomationDestinationAsset.Overrides()
func (r AutomationDestinationAssetResp) ToParam() AutomationDestinationAsset {
	return param.Override[AutomationDestinationAsset](json.RawMessage(r.RawJSON()))
}

// Destination asset identified by contract address on a specific chain (CAIP-2).
//
// The properties AssetAddress, Caip2 are required.
type AutomationDestinationAsset struct {
	AssetAddress string `json:"asset_address" api:"required"`
	Caip2        string `json:"caip2" api:"required"`
	paramObj
}

func (r AutomationDestinationAsset) MarshalJSON() (data []byte, err error) {
	type shadow AutomationDestinationAsset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationDestinationAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A destination asset spec accepting either raw identifiers (asset_address, caip2)
// or human-readable aliases (asset, chain). Exactly one of asset_address or asset
// must be provided; exactly one of caip2 or chain must be provided.
type AutomationDestinationAssetInput struct {
	Asset param.Opt[string] `json:"asset,omitzero"`
	Chain param.Opt[string] `json:"chain,omitzero"`
	AutomationDestinationAsset
}

func (r AutomationDestinationAssetInput) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*AutomationDestinationAssetInput
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Action configuration for depositing into an Earn vault.
type AutomationEarnDepositActionConfig struct {
	// Any of "earn_deposit".
	Type    AutomationEarnDepositActionConfigType `json:"type" api:"required"`
	VaultID string                                `json:"vault_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		VaultID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationEarnDepositActionConfig) RawJSON() string { return r.JSON.raw }
func (r *AutomationEarnDepositActionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationEarnDepositActionConfigType string

const (
	AutomationEarnDepositActionConfigTypeEarnDeposit AutomationEarnDepositActionConfigType = "earn_deposit"
)

// Action configuration for depositing into an Earn vault (input form).
//
// The properties Type, VaultID are required.
type AutomationEarnDepositActionConfigInput struct {
	// Any of "earn_deposit".
	Type    AutomationEarnDepositActionConfigInputType `json:"type,omitzero" api:"required"`
	VaultID string                                     `json:"vault_id" api:"required"`
	paramObj
}

func (r AutomationEarnDepositActionConfigInput) MarshalJSON() (data []byte, err error) {
	type shadow AutomationEarnDepositActionConfigInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationEarnDepositActionConfigInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationEarnDepositActionConfigInputType string

const (
	AutomationEarnDepositActionConfigInputTypeEarnDeposit AutomationEarnDepositActionConfigInputType = "earn_deposit"
)

// Action configuration for swap operations.
type AutomationSwapActionConfig struct {
	// Destination asset identified by contract address on a specific chain (CAIP-2).
	DestinationChainAsset AutomationDestinationAssetResp `json:"destination_chain_asset" api:"required"`
	// Any of "swap".
	Type AutomationSwapActionConfigType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationChainAsset respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationSwapActionConfig) RawJSON() string { return r.JSON.raw }
func (r *AutomationSwapActionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationSwapActionConfigType string

const (
	AutomationSwapActionConfigTypeSwap AutomationSwapActionConfigType = "swap"
)

// Action configuration for swap operations (input form with alias support).
//
// The properties DestinationChainAsset, Type are required.
type AutomationSwapActionConfigInput struct {
	// A destination asset spec accepting either raw identifiers (asset_address, caip2)
	// or human-readable aliases (asset, chain). Exactly one of asset_address or asset
	// must be provided; exactly one of caip2 or chain must be provided.
	DestinationChainAsset AutomationDestinationAssetInput `json:"destination_chain_asset,omitzero" api:"required"`
	// Any of "swap".
	Type AutomationSwapActionConfigInputType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationSwapActionConfigInput) MarshalJSON() (data []byte, err error) {
	type shadow AutomationSwapActionConfigInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationSwapActionConfigInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationSwapActionConfigInputType string

const (
	AutomationSwapActionConfigInputTypeSwap AutomationSwapActionConfigInputType = "swap"
)

// Trigger configuration for deposit events.
type AutomationTriggerConfig struct {
	// Which assets to include/exclude for an automation trigger.
	Assets AutomationAssetFilterUnion `json:"assets" api:"required"`
	// Any of "deposit".
	Type AutomationTriggerConfigType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationTriggerConfig) RawJSON() string { return r.JSON.raw }
func (r *AutomationTriggerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationTriggerConfigType string

const (
	AutomationTriggerConfigTypeDeposit AutomationTriggerConfigType = "deposit"
)

// Trigger configuration for deposit events (input form with alias support).
//
// The properties Assets, Type are required.
type AutomationTriggerConfigInput struct {
	// Which assets to include/exclude for an automation trigger (input form with alias
	// support).
	Assets AutomationAssetFilterInputUnion `json:"assets,omitzero" api:"required"`
	// Any of "deposit".
	Type AutomationTriggerConfigInputType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AutomationTriggerConfigInput) MarshalJSON() (data []byte, err error) {
	type shadow AutomationTriggerConfigInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationTriggerConfigInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationTriggerConfigInputType string

const (
	AutomationTriggerConfigInputTypeDeposit AutomationTriggerConfigInputType = "deposit"
)

// Request body for creating a wallet automation.
//
// The properties Config, OwnerID are required.
type CreateAutomationRequestBody struct {
	OwnerID param.Opt[string] `json:"owner_id,omitzero" api:"required"`
	// Full configuration for a wallet automation (trigger + action) accepting
	// human-readable aliases.
	Config AutomationConfigInput `json:"config,omitzero" api:"required"`
	Name   param.Opt[string]     `json:"name,omitzero"`
	paramObj
}

func (r CreateAutomationRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CreateAutomationRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAutomationRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-attachment parameters for swap automations.
type SwapAttachmentParamsResp struct {
	DestinationAddress string `json:"destination_address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationAddress respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwapAttachmentParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SwapAttachmentParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SwapAttachmentParamsResp to a SwapAttachmentParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SwapAttachmentParams.Overrides()
func (r SwapAttachmentParamsResp) ToParam() SwapAttachmentParams {
	return param.Override[SwapAttachmentParams](json.RawMessage(r.RawJSON()))
}

// Per-attachment parameters for swap automations.
//
// The property DestinationAddress is required.
type SwapAttachmentParams struct {
	DestinationAddress string `json:"destination_address" api:"required"`
	paramObj
}

func (r SwapAttachmentParams) MarshalJSON() (data []byte, err error) {
	type shadow SwapAttachmentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwapAttachmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for updating a wallet automation.
type UpdateAutomationRequestBody struct {
	Name    param.Opt[string] `json:"name,omitzero"`
	Enabled param.Opt[bool]   `json:"enabled,omitzero"`
	// A unique identifier for a key quorum.
	OwnerID param.Opt[KeyQuorumID] `json:"owner_id,omitzero" format:"cuid2"`
	// Full configuration for a wallet automation (trigger + action) accepting
	// human-readable aliases.
	Config AutomationConfigInput `json:"config,omitzero"`
	paramObj
}

func (r UpdateAutomationRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAutomationRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAutomationRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of wallet automation executions.
type WalletAutomationExecutionListResponse struct {
	Data       []WalletAutomationExecutionResponse `json:"data" api:"required"`
	NextCursor string                              `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationExecutionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationExecutionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A record of a single automation execution created by a deposit.
type WalletAutomationExecutionResponse struct {
	ID                     string `json:"id" api:"required"`
	AutomationAttachmentID string `json:"automation_attachment_id" api:"required"`
	CompletedAt            string `json:"completed_at" api:"required"`
	CreatedAt              string `json:"created_at" api:"required"`
	FailedAt               string `json:"failed_at" api:"required"`
	FailureReason          string `json:"failure_reason" api:"required"`
	// Execution lifecycle status.
	//
	// Any of "pending", "submitted", "completed", "failed", "skipped".
	Status              WalletAutomationExecutionStatus `json:"status" api:"required"`
	SubmittedAt         string                          `json:"submitted_at" api:"required"`
	TriggerAssetAddress string                          `json:"trigger_asset_address" api:"required"`
	TriggerBlockNumber  string                          `json:"trigger_block_number" api:"required"`
	TriggerCaip2        string                          `json:"trigger_caip2" api:"required"`
	TriggerTxHash       string                          `json:"trigger_tx_hash" api:"required"`
	UpdatedAt           string                          `json:"updated_at" api:"required"`
	WalletActionID      string                          `json:"wallet_action_id" api:"required"`
	WalletID            string                          `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AutomationAttachmentID respjson.Field
		CompletedAt            respjson.Field
		CreatedAt              respjson.Field
		FailedAt               respjson.Field
		FailureReason          respjson.Field
		Status                 respjson.Field
		SubmittedAt            respjson.Field
		TriggerAssetAddress    respjson.Field
		TriggerBlockNumber     respjson.Field
		TriggerCaip2           respjson.Field
		TriggerTxHash          respjson.Field
		UpdatedAt              respjson.Field
		WalletActionID         respjson.Field
		WalletID               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationExecutionResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationExecutionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Execution lifecycle status.
type WalletAutomationExecutionStatus string

const (
	WalletAutomationExecutionStatusPending   WalletAutomationExecutionStatus = "pending"
	WalletAutomationExecutionStatusSubmitted WalletAutomationExecutionStatus = "submitted"
	WalletAutomationExecutionStatusCompleted WalletAutomationExecutionStatus = "completed"
	WalletAutomationExecutionStatusFailed    WalletAutomationExecutionStatus = "failed"
	WalletAutomationExecutionStatusSkipped   WalletAutomationExecutionStatus = "skipped"
)

// Paginated list of wallet automations.
type WalletAutomationListResponse struct {
	Data       []WalletAutomationResponse `json:"data" api:"required"`
	NextCursor string                     `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationListResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The outcome of checking one asset on the requested chain during a reindex.
type WalletAutomationReindexAssetResult struct {
	// Asset contract address; the native asset uses `native`.
	AssetAddress string `json:"asset_address" api:"required"`
	// An EVM, Solana, or Tron CAIP-2 chain identifier supported by wallet automation
	// reindex.
	Caip2 TronCaip2 `json:"caip2" api:"required"`
	// ID of the in-flight execution blocking a re-trigger. Populated only when
	// `status` is `skipped_existing_execution`; `null` otherwise.
	ExistingExecutionID string `json:"existing_execution_id" api:"required"`
	// On-chain balance in base units. Populated when `status` is `submitted` or
	// `skipped_zero_balance`; `null` otherwise. For example, 1 OUSD is `1000000`.
	RawBalance string `json:"raw_balance" api:"required"`
	// Outcome of checking a single asset during a wallet automation reindex. One of
	// `submitted`, `skipped_zero_balance`, `skipped_no_match`,
	// `skipped_existing_execution`, or `failed`. `submitted` confirms that an
	// execution was enqueued.
	Status WalletAutomationReindexAssetStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetAddress        respjson.Field
		Caip2               respjson.Field
		ExistingExecutionID respjson.Field
		RawBalance          respjson.Field
		Status              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationReindexAssetResult) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationReindexAssetResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAutomationReindexAssetStatus = string

// Request body for re-checking a wallet against its wallet automations. Identify
// the wallet by wallet_id or deposit_address (at least one is required). If both
// are provided, wallet_id takes precedence and deposit_address must match that
// wallet's address. Specify exactly one of caip2 or chain, and the asset_address
// to check. Useful for recovering a deposit that was missed or failed to trigger
// its automation.
//
// The property AssetAddress is required.
type WalletAutomationReindexRequestBody struct {
	// Asset contract address to check; the native asset uses `native`.
	AssetAddress string `json:"asset_address" api:"required"`
	// Human-readable chain name to check. Specify exactly one of `caip2` or `chain`.
	Chain param.Opt[string] `json:"chain,omitzero"`
	// On-chain deposit address of the wallet to reindex. Must match the resolved
	// wallet's address if `wallet_id` is also provided.
	DepositAddress param.Opt[string] `json:"deposit_address,omitzero"`
	// Privy wallet ID to reindex. Takes precedence over `deposit_address` when both
	// are supplied.
	WalletID param.Opt[string] `json:"wallet_id,omitzero"`
	// An EVM, Solana, or Tron CAIP-2 chain identifier supported by wallet automation
	// reindex.
	Caip2 TronCaip2 `json:"caip2,omitzero"`
	paramObj
}

func (r WalletAutomationReindexRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow WalletAutomationReindexRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletAutomationReindexRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of re-checking a wallet against its wallet automations.
type WalletAutomationReindexResponse struct {
	Results  []WalletAutomationReindexAssetResult `json:"results" api:"required"`
	WalletID string                               `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationReindexResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationReindexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A wallet automation.
type WalletAutomationResponse struct {
	ID    string `json:"id" api:"required"`
	AppID string `json:"app_id" api:"required"`
	// Full configuration for a wallet automation (trigger + action).
	Config    AutomationConfig `json:"config" api:"required"`
	CreatedAt string           `json:"created_at" api:"required"`
	Name      string           `json:"name" api:"required"`
	OwnerID   string           `json:"owner_id" api:"required"`
	// Automation lifecycle state: 'enabled' = running, 'disabled' = not running.
	//
	// Any of "enabled", "disabled".
	Status    WalletAutomationStatus `json:"status" api:"required"`
	UpdatedAt string                 `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		Config      respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		OwnerID     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automation lifecycle state: 'enabled' = running, 'disabled' = not running.
type WalletAutomationStatus string

const (
	WalletAutomationStatusEnabled  WalletAutomationStatus = "enabled"
	WalletAutomationStatusDisabled WalletAutomationStatus = "disabled"
)

// Confirmation of a successful automation operation.
type WalletAutomationSuccessResponse struct {
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAutomationSuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletAutomationSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAutomationNewParams struct {
	// Request body for creating a wallet automation.
	CreateAutomationRequestBody CreateAutomationRequestBody
	paramObj
}

func (r WalletAutomationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAutomationRequestBody)
}
func (r *WalletAutomationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAutomationUpdateParams struct {
	// Request body for updating a wallet automation.
	UpdateAutomationRequestBody UpdateAutomationRequestBody
	paramObj
}

func (r WalletAutomationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAutomationRequestBody)
}
func (r *WalletAutomationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAutomationListParams struct {
	Cursor   param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	WalletID param.Opt[string] `query:"wallet_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletAutomationListParams]'s query parameters as
// `url.Values`.
func (r WalletAutomationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletAutomationListExecutionsParams struct {
	Cursor   param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	WalletID param.Opt[string] `query:"wallet_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletAutomationListExecutionsParams]'s query parameters as
// `url.Values`.
func (r WalletAutomationListExecutionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletAutomationReindexParams struct {
	// Request body for re-checking a wallet against its wallet automations. Identify
	// the wallet by wallet_id or deposit_address (at least one is required). If both
	// are provided, wallet_id takes precedence and deposit_address must match that
	// wallet's address. Specify exactly one of caip2 or chain, and the asset_address
	// to check. Useful for recovering a deposit that was missed or failed to trigger
	// its automation.
	WalletAutomationReindexRequestBody WalletAutomationReindexRequestBody
	paramObj
}

func (r WalletAutomationReindexParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletAutomationReindexRequestBody)
}
func (r *WalletAutomationReindexParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
