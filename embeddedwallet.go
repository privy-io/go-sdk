// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
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
