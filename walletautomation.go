// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

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

// An asset identified by contract address, scoped to a chain via CAIP-2.
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

// An asset identified by contract address, scoped to a chain via CAIP-2.
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
// must be provided; at most one of caip2 or chain may be provided.
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
