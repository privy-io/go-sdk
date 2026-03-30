// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// EmbeddedWalletService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddedWalletService] method instead.
type EmbeddedWalletService struct {
	Options []option.RequestOption
}

// NewEmbeddedWalletService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddedWalletService(opts ...option.RequestOption) (r EmbeddedWalletService) {
	r = EmbeddedWalletService{}
	r.Options = opts
	return
}

// The supported smart wallet providers.
type SmartWalletType string

const (
	SmartWalletTypeSafe                SmartWalletType = "safe"
	SmartWalletTypeKernel              SmartWalletType = "kernel"
	SmartWalletTypeLightAccount        SmartWalletType = "light_account"
	SmartWalletTypeBiconomy            SmartWalletType = "biconomy"
	SmartWalletTypeCoinbaseSmartWallet SmartWalletType = "coinbase_smart_wallet"
	SmartWalletTypeThirdweb            SmartWalletType = "thirdweb"
)

// The Alchemy paymaster context for a smart wallet network configuration.
type AlchemyPaymasterContext struct {
	PolicyID string `json:"policy_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PolicyID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlchemyPaymasterContext) RawJSON() string { return r.JSON.raw }
func (r *AlchemyPaymasterContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Network configuration for a smart wallet.
type SmartWalletNetworkConfiguration struct {
	BundlerURL string `json:"bundler_url" api:"required"`
	ChainID    string `json:"chain_id" api:"required"`
	ChainName  string `json:"chain_name"`
	// The Alchemy paymaster context for a smart wallet network configuration.
	PaymasterContext AlchemyPaymasterContext `json:"paymaster_context"`
	PaymasterURL     string                  `json:"paymaster_url"`
	RpcURL           string                  `json:"rpc_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BundlerURL       respjson.Field
		ChainID          respjson.Field
		ChainName        respjson.Field
		PaymasterContext respjson.Field
		PaymasterURL     respjson.Field
		RpcURL           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmartWalletNetworkConfiguration) RawJSON() string { return r.JSON.raw }
func (r *SmartWalletNetworkConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A disabled smart wallet configuration.
type SmartWalletConfigurationDisabled struct {
	// Any of false.
	Enabled bool `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmartWalletConfigurationDisabled) RawJSON() string { return r.JSON.raw }
func (r *SmartWalletConfigurationDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An enabled smart wallet configuration.
type SmartWalletConfigurationEnabled struct {
	ConfiguredNetworks []SmartWalletNetworkConfiguration `json:"configured_networks" api:"required"`
	// Any of true.
	Enabled bool `json:"enabled" api:"required"`
	// The supported smart wallet providers.
	//
	// Any of "safe", "kernel", "light_account", "biconomy", "coinbase_smart_wallet",
	// "thirdweb".
	SmartWalletType    SmartWalletType `json:"smart_wallet_type" api:"required"`
	SmartWalletVersion string          `json:"smart_wallet_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfiguredNetworks respjson.Field
		Enabled            respjson.Field
		SmartWalletType    respjson.Field
		SmartWalletVersion respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmartWalletConfigurationEnabled) RawJSON() string { return r.JSON.raw }
func (r *SmartWalletConfigurationEnabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmartWalletConfigurationUnion contains all possible properties and values from
// [SmartWalletConfigurationDisabled], [SmartWalletConfigurationEnabled].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmartWalletConfigurationUnion struct {
	Enabled bool `json:"enabled"`
	// This field is from variant [SmartWalletConfigurationEnabled].
	ConfiguredNetworks []SmartWalletNetworkConfiguration `json:"configured_networks"`
	// This field is from variant [SmartWalletConfigurationEnabled].
	SmartWalletType SmartWalletType `json:"smart_wallet_type"`
	// This field is from variant [SmartWalletConfigurationEnabled].
	SmartWalletVersion string `json:"smart_wallet_version"`
	JSON               struct {
		Enabled            respjson.Field
		ConfiguredNetworks respjson.Field
		SmartWalletType    respjson.Field
		SmartWalletVersion respjson.Field
		raw                string
	} `json:"-"`
}

func (u SmartWalletConfigurationUnion) AsSmartWalletConfigurationDisabled() (v SmartWalletConfigurationDisabled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmartWalletConfigurationUnion) AsSmartWalletConfigurationEnabled() (v SmartWalletConfigurationEnabled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmartWalletConfigurationUnion) RawJSON() string { return u.JSON.raw }

func (r *SmartWalletConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An additional signer configuration for a wallet.
//
// The property SignerID is required.
type WalletCreationAdditionalSignerItem struct {
	// The key quorum ID for the signer.
	SignerID string `json:"signer_id" api:"required" format:"cuid2"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet. Currently, only one
	// policy is supported per signer.
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero"`
	paramObj
}

func (r WalletCreationAdditionalSignerItem) MarshalJSON() (data []byte, err error) {
	type shadow WalletCreationAdditionalSignerItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletCreationAdditionalSignerItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The fields on wallet creation that can be specified when creating a
// user-controlled embedded server wallet.
//
// The property ChainType is required.
type WalletCreationInput struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero" api:"required"`
	// Create a smart wallet with this wallet as the signer. Only supported for wallets
	// with `chain_type: "ethereum"`.
	CreateSmartWallet param.Opt[bool] `json:"create_smart_wallet,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletCreationAdditionalSignerItem `json:"additional_signers,omitzero"`
	// Policy IDs to enforce on the wallet. Currently, only one policy is supported per
	// wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r WalletCreationInput) MarshalJSON() (data []byte, err error) {
	type shadow WalletCreationInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletCreationInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
