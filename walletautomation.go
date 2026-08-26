// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
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

type AutomationAssetFilterAllMode string

const (
	AutomationAssetFilterAllModeAll AutomationAssetFilterAllMode = "all"
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
