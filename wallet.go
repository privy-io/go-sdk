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

// WalletService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletService] method instead.
type WalletService struct {
	Options []option.RequestOption
	// Operations related to wallets
	Transactions WalletTransactionService
	// Operations related to wallets
	Balance WalletBalanceService
}

// NewWalletService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWalletService(opts ...option.RequestOption) (r WalletService) {
	r = WalletService{}
	r.Options = opts
	r.Transactions = NewWalletTransactionService(opts...)
	r.Balance = NewWalletBalanceService(opts...)
	return
}

// Creates a new wallet on the requested chain and for the requested owner.
func (r *WalletService) New(ctx context.Context, params WalletNewParams, opts ...option.RequestOption) (res *Wallet, err error) {
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a wallet's policies or authorization key configuration.
func (r *WalletService) Update(ctx context.Context, walletID string, params WalletUpdateParams, opts ...option.RequestOption) (res *Wallet, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Get all wallets in your app.
func (r *WalletService) List(ctx context.Context, query WalletListParams, opts ...option.RequestOption) (res *pagination.Cursor[Wallet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/wallets"
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

// Get all wallets in your app.
func (r *WalletService) ListAutoPaging(ctx context.Context, query WalletListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Wallet] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Initialize a wallet import. Complete by submitting the import.
func (r *WalletService) InitImport(ctx context.Context, body WalletInitImportParams, opts ...option.RequestOption) (res *WalletInitImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/import/init"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submit a wallet import request.
func (r *WalletService) SubmitImport(ctx context.Context, body WalletSubmitImportParams, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/import/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Obtain a session key to enable wallet access.
func (r *WalletService) AuthenticateWithJwt(ctx context.Context, body WalletAuthenticateWithJwtParams, opts ...option.RequestOption) (res *WalletAuthenticateWithJwtResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/authenticate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Creates multiple wallets in a single request. Each wallet creation is
// independent; failures for one wallet do not affect others. Maximum batch size is
// 100 wallets.
func (r *WalletService) Batch(ctx context.Context, body WalletBatchParams, opts ...option.RequestOption) (res *WalletBatchCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Export a wallet's private key
func (r *WalletService) Export(ctx context.Context, walletID string, params WalletExportParams, opts ...option.RequestOption) (res *WalletExportResponseBody, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/export", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a wallet by wallet ID.
func (r *WalletService) Get(ctx context.Context, walletID string, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Look up a wallet by its blockchain address. Returns the wallet object if found.
func (r *WalletService) GetWalletByAddress(ctx context.Context, body WalletGetWalletByAddressParams, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sign a message with a wallet by wallet ID.
func (r *WalletService) RawSign(ctx context.Context, walletID string, params WalletRawSignParams, opts ...option.RequestOption) (res *RawSignResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/raw_sign", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sign a message or transaction with a wallet by wallet ID.
func (r *WalletService) Rpc(ctx context.Context, walletID string, params WalletRpcParams, opts ...option.RequestOption) (res *WalletRpcResponseUnion, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.PrivyRequestExpiry) {
		opts = append(opts, option.WithHeader("privy-request-expiry", fmt.Sprintf("%v", params.PrivyRequestExpiry.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/rpc", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// The wallet chain types that support curve-based signing.
type CurveSigningChainType string

const (
	CurveSigningChainTypeCosmos         CurveSigningChainType = "cosmos"
	CurveSigningChainTypeStellar        CurveSigningChainType = "stellar"
	CurveSigningChainTypeSui            CurveSigningChainType = "sui"
	CurveSigningChainTypeAptos          CurveSigningChainType = "aptos"
	CurveSigningChainTypeMovement       CurveSigningChainType = "movement"
	CurveSigningChainTypeTron           CurveSigningChainType = "tron"
	CurveSigningChainTypeBitcoinSegwit  CurveSigningChainType = "bitcoin-segwit"
	CurveSigningChainTypeBitcoinTaproot CurveSigningChainType = "bitcoin-taproot"
	CurveSigningChainTypeNear           CurveSigningChainType = "near"
	CurveSigningChainTypeTon            CurveSigningChainType = "ton"
	CurveSigningChainTypeStarknet       CurveSigningChainType = "starknet"
)

// The wallet chain types.
type WalletChainType string

const (
	WalletChainTypeEthereum       WalletChainType = "ethereum"
	WalletChainTypeSolana         WalletChainType = "solana"
	WalletChainTypeCosmos         WalletChainType = "cosmos"
	WalletChainTypeStellar        WalletChainType = "stellar"
	WalletChainTypeSui            WalletChainType = "sui"
	WalletChainTypeAptos          WalletChainType = "aptos"
	WalletChainTypeMovement       WalletChainType = "movement"
	WalletChainTypeTron           WalletChainType = "tron"
	WalletChainTypeBitcoinSegwit  WalletChainType = "bitcoin-segwit"
	WalletChainTypeBitcoinTaproot WalletChainType = "bitcoin-taproot"
	WalletChainTypeNear           WalletChainType = "near"
	WalletChainTypeTon            WalletChainType = "ton"
	WalletChainTypeStarknet       WalletChainType = "starknet"
	WalletChainTypeSpark          WalletChainType = "spark"
)

type PolicyInput []string

// A single additional signer for a wallet, with an optional policy override.
type AdditionalSignerItemInputResp struct {
	// A unique identifier for a key quorum.
	SignerID KeyQuorumID `json:"signer_id" api:"required" format:"cuid2"`
	// An optional list of up to one policy ID to enforce on the wallet.
	OverridePolicyIDs PolicyInput `json:"override_policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignerID          respjson.Field
		OverridePolicyIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdditionalSignerItemInputResp) RawJSON() string { return r.JSON.raw }
func (r *AdditionalSignerItemInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AdditionalSignerItemInputResp to a
// AdditionalSignerItemInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AdditionalSignerItemInput.Overrides()
func (r AdditionalSignerItemInputResp) ToParam() AdditionalSignerItemInput {
	return param.Override[AdditionalSignerItemInput](json.RawMessage(r.RawJSON()))
}

// A single additional signer for a wallet, with an optional policy override.
//
// The property SignerID is required.
type AdditionalSignerItemInput struct {
	// A unique identifier for a key quorum.
	SignerID KeyQuorumID `json:"signer_id" api:"required" format:"cuid2"`
	// An optional list of up to one policy ID to enforce on the wallet.
	OverridePolicyIDs PolicyInput `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r AdditionalSignerItemInput) MarshalJSON() (data []byte, err error) {
	type shadow AdditionalSignerItemInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdditionalSignerItemInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdditionalSignerInputResp []AdditionalSignerItemInputResp

type AdditionalSignerInput []AdditionalSignerItemInput

// A single additional signer on a wallet, with an optional policy override.
type WalletAdditionalSignerItem struct {
	// A unique identifier for a key quorum.
	SignerID KeyQuorumID `json:"signer_id" api:"required" format:"cuid2"`
	// An optional list of up to one policy ID to enforce on the wallet.
	OverridePolicyIDs PolicyInput `json:"override_policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignerID          respjson.Field
		OverridePolicyIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAdditionalSignerItem) RawJSON() string { return r.JSON.raw }
func (r *WalletAdditionalSignerItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAdditionalSigner []WalletAdditionalSignerItem

type Address = string

// A named asset supported across all chains.
type WalletAsset string

const (
	WalletAssetUsdc  WalletAsset = "usdc"
	WalletAssetUsdcE WalletAsset = "usdc.e"
	WalletAssetEth   WalletAsset = "eth"
	WalletAssetPol   WalletAsset = "pol"
	WalletAssetUsdt  WalletAsset = "usdt"
	WalletAssetEurc  WalletAsset = "eurc"
	WalletAssetUsdb  WalletAsset = "usdb"
	WalletAssetSol   WalletAsset = "sol"
)

// Information about the custodian managing this wallet.
type WalletCustodian struct {
	// The custodian responsible for the wallet.
	Provider string `json:"provider" api:"required"`
	// The resource ID of the beneficiary of the custodial wallet.
	ProviderUserID string `json:"provider_user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Provider       respjson.Field
		ProviderUserID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletCustodian) RawJSON() string { return r.JSON.raw }
func (r *WalletCustodian) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encryption type of the wallet to import. Currently only supports `HPKE`.
type HpkeEncryption string

const (
	HpkeEncryptionHpke HpkeEncryption = "HPKE"
)

type RecipientPublicKey = string

// The export type. 'display' is for showing the key to the user in the UI,
// 'client' is for exporting to the client application.
type ExportType string

const (
	ExportTypeDisplay ExportType = "display"
	ExportTypeClient  ExportType = "client"
)

// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
type PrivateKeyExportInputResp struct {
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type" api:"required"`
	// The recipient public key for HPKE encryption, in PEM or DER (base64-encoded)
	// format.
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key" api:"required"`
	ExportSeedPhrase   bool               `json:"export_seed_phrase"`
	// The export type. 'display' is for showing the key to the user in the UI,
	// 'client' is for exporting to the client application.
	//
	// Any of "display", "client".
	ExportType ExportType `json:"export_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptionType     respjson.Field
		RecipientPublicKey respjson.Field
		ExportSeedPhrase   respjson.Field
		ExportType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivateKeyExportInputResp) RawJSON() string { return r.JSON.raw }
func (r *PrivateKeyExportInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PrivateKeyExportInputResp to a PrivateKeyExportInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PrivateKeyExportInput.Overrides()
func (r PrivateKeyExportInputResp) ToParam() PrivateKeyExportInput {
	return param.Override[PrivateKeyExportInput](json.RawMessage(r.RawJSON()))
}

// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
//
// The properties EncryptionType, RecipientPublicKey are required.
type PrivateKeyExportInput struct {
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// The recipient public key for HPKE encryption, in PEM or DER (base64-encoded)
	// format.
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key" api:"required"`
	ExportSeedPhrase   param.Opt[bool]    `json:"export_seed_phrase,omitzero"`
	// The export type. 'display' is for showing the key to the user in the UI,
	// 'client' is for exporting to the client application.
	//
	// Any of "display", "client".
	ExportType ExportType `json:"export_type,omitzero"`
	paramObj
}

func (r PrivateKeyExportInput) MarshalJSON() (data []byte, err error) {
	type shadow PrivateKeyExportInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrivateKeyExportInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
type SeedPhraseExportInputResp struct {
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type" api:"required"`
	// The recipient public key for HPKE encryption, in PEM or DER (base64-encoded)
	// format.
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key" api:"required"`
	ExportSeedPhrase   bool               `json:"export_seed_phrase"`
	// The export type. 'display' is for showing the key to the user in the UI,
	// 'client' is for exporting to the client application.
	//
	// Any of "display", "client".
	ExportType ExportType `json:"export_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptionType     respjson.Field
		RecipientPublicKey respjson.Field
		ExportSeedPhrase   respjson.Field
		ExportType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeedPhraseExportInputResp) RawJSON() string { return r.JSON.raw }
func (r *SeedPhraseExportInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SeedPhraseExportInputResp to a SeedPhraseExportInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SeedPhraseExportInput.Overrides()
func (r SeedPhraseExportInputResp) ToParam() SeedPhraseExportInput {
	return param.Override[SeedPhraseExportInput](json.RawMessage(r.RawJSON()))
}

// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
//
// The properties EncryptionType, RecipientPublicKey are required.
type SeedPhraseExportInput struct {
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// The recipient public key for HPKE encryption, in PEM or DER (base64-encoded)
	// format.
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key" api:"required"`
	ExportSeedPhrase   param.Opt[bool]    `json:"export_seed_phrase,omitzero"`
	// The export type. 'display' is for showing the key to the user in the UI,
	// 'client' is for exporting to the client application.
	//
	// Any of "display", "client".
	ExportType ExportType `json:"export_type,omitzero"`
	paramObj
}

func (r SeedPhraseExportInput) MarshalJSON() (data []byte, err error) {
	type shadow SeedPhraseExportInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SeedPhraseExportInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing HPKE-encrypted wallet data (private key or seed phrase).
type SeedPhraseExportResponse struct {
	Ciphertext      string `json:"ciphertext" api:"required"`
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ciphertext      respjson.Field
		EncapsulatedKey respjson.Field
		EncryptionType  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SeedPhraseExportResponse) RawJSON() string { return r.JSON.raw }
func (r *SeedPhraseExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the amount refers to the input token or output token.
type AmountType string

const (
	AmountTypeExactInput  AmountType = "exact_input"
	AmountTypeExactOutput AmountType = "exact_output"
)

// The chain type of the wallet to import. Currently supports `ethereum` and
// `solana`.
type WalletImportSupportedChains string

const (
	WalletImportSupportedChainsEthereum WalletImportSupportedChains = "ethereum"
	WalletImportSupportedChainsSolana   WalletImportSupportedChains = "solana"
)

// The AEAD algorithm used for HPKE encryption.
type HpkeAeadAlgorithm string

const (
	HpkeAeadAlgorithmChacha20Poly1305 HpkeAeadAlgorithm = "CHACHA20_POLY1305"
	HpkeAeadAlgorithmAesGcm256        HpkeAeadAlgorithm = "AES_GCM256"
)

// Optional HPKE configuration for wallet import decryption. These parameters allow
// importing wallets encrypted by external providers that use different HPKE
// configurations.
type HpkeImportConfig struct {
	// Additional Authenticated Data (AAD) used during encryption. Should be
	// base64-encoded bytes.
	Aad param.Opt[string] `json:"aad,omitzero"`
	// Application-specific context information (INFO) used during HPKE encryption.
	// Should be base64-encoded bytes.
	Info param.Opt[string] `json:"info,omitzero"`
	// The AEAD algorithm used for HPKE encryption.
	//
	// Any of "CHACHA20_POLY1305", "AES_GCM256".
	AeadAlgorithm HpkeAeadAlgorithm `json:"aead_algorithm,omitzero"`
	paramObj
}

func (r HpkeImportConfig) MarshalJSON() (data []byte, err error) {
	type shadow HpkeImportConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HpkeImportConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Hex = string

// QuantityUnionResp contains all possible properties and values from [Hex],
// [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type QuantityUnionResp struct {
	// This field will be present if the value is a [Hex] instead of an object.
	OfString Hex `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u QuantityUnionResp) AsString() (v Hex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuantityUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QuantityUnionResp) RawJSON() string { return u.JSON.raw }

func (r *QuantityUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this QuantityUnionResp to a QuantityUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// QuantityUnion.Overrides()
func (r QuantityUnionResp) ToParam() QuantityUnion {
	return param.Override[QuantityUnion](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QuantityUnion struct {
	OfString param.Opt[Hex]   `json:",omitzero,inline"`
	OfInt    param.Opt[int64] `json:",omitzero,inline"`
	paramUnion
}

func (u QuantityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *QuantityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Exports the private key of the wallet.
type ExportPrivateKeyRpcInputResp struct {
	Address string `json:"address" api:"required"`
	// Any of "exportPrivateKey".
	Method ExportPrivateKeyRpcInputMethod `json:"method" api:"required"`
	// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
	Params PrivateKeyExportInputResp `json:"params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Method      respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPrivateKeyRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *ExportPrivateKeyRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExportPrivateKeyRpcInputResp to a
// ExportPrivateKeyRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExportPrivateKeyRpcInput.Overrides()
func (r ExportPrivateKeyRpcInputResp) ToParam() ExportPrivateKeyRpcInput {
	return param.Override[ExportPrivateKeyRpcInput](json.RawMessage(r.RawJSON()))
}

type ExportPrivateKeyRpcInputMethod string

const (
	ExportPrivateKeyRpcInputMethodExportPrivateKey ExportPrivateKeyRpcInputMethod = "exportPrivateKey"
)

// Exports the private key of the wallet.
//
// The properties Address, Method, Params are required.
type ExportPrivateKeyRpcInput struct {
	Address string `json:"address" api:"required"`
	// Any of "exportPrivateKey".
	Method ExportPrivateKeyRpcInputMethod `json:"method,omitzero" api:"required"`
	// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
	Params PrivateKeyExportInput `json:"params,omitzero" api:"required"`
	paramObj
}

func (r ExportPrivateKeyRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow ExportPrivateKeyRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPrivateKeyRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the `exportPrivateKey` RPC.
type ExportPrivateKeyRpcResponse struct {
	// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
	Data PrivateKeyExportInputResp `json:"data" api:"required"`
	// Any of "exportPrivateKey".
	Method ExportPrivateKeyRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPrivateKeyRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *ExportPrivateKeyRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportPrivateKeyRpcResponseMethod string

const (
	ExportPrivateKeyRpcResponseMethodExportPrivateKey ExportPrivateKeyRpcResponseMethod = "exportPrivateKey"
)

// Exports the seed phrase of the wallet.
type ExportSeedPhraseRpcInputResp struct {
	Address string `json:"address" api:"required"`
	// Any of "exportSeedPhrase".
	Method ExportSeedPhraseRpcInputMethod `json:"method" api:"required"`
	// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
	Params SeedPhraseExportInputResp `json:"params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Method      respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportSeedPhraseRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *ExportSeedPhraseRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExportSeedPhraseRpcInputResp to a
// ExportSeedPhraseRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExportSeedPhraseRpcInput.Overrides()
func (r ExportSeedPhraseRpcInputResp) ToParam() ExportSeedPhraseRpcInput {
	return param.Override[ExportSeedPhraseRpcInput](json.RawMessage(r.RawJSON()))
}

type ExportSeedPhraseRpcInputMethod string

const (
	ExportSeedPhraseRpcInputMethodExportSeedPhrase ExportSeedPhraseRpcInputMethod = "exportSeedPhrase"
)

// Exports the seed phrase of the wallet.
//
// The properties Address, Method, Params are required.
type ExportSeedPhraseRpcInput struct {
	Address string `json:"address" api:"required"`
	// Any of "exportSeedPhrase".
	Method ExportSeedPhraseRpcInputMethod `json:"method,omitzero" api:"required"`
	// Input for exporting a wallet (private key or seed phrase) with HPKE encryption.
	Params SeedPhraseExportInput `json:"params,omitzero" api:"required"`
	paramObj
}

func (r ExportSeedPhraseRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow ExportSeedPhraseRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportSeedPhraseRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the `exportSeedPhrase` RPC.
type ExportSeedPhraseRpcResponse struct {
	// Response containing HPKE-encrypted wallet data (private key or seed phrase).
	Data SeedPhraseExportResponse `json:"data" api:"required"`
	// Any of "exportSeedPhrase".
	Method ExportSeedPhraseRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportSeedPhraseRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *ExportSeedPhraseRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportSeedPhraseRpcResponseMethod string

const (
	ExportSeedPhraseRpcResponseMethodExportSeedPhrase ExportSeedPhraseRpcResponseMethod = "exportSeedPhrase"
)

// Parameters for signing a pre-computed hash with the `raw_sign` RPC.
//
// The property Hash is required.
type RawSignHashParams struct {
	// A hex-encoded string prefixed with '0x'.
	Hash Hex `json:"hash" api:"required"`
	paramObj
}

func (r RawSignHashParams) MarshalJSON() (data []byte, err error) {
	type shadow RawSignHashParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RawSignHashParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encoding scheme for bytes in the `raw_sign` RPC.
type RawSignBytesEncoding string

const (
	RawSignBytesEncodingUtf8   RawSignBytesEncoding = "utf-8"
	RawSignBytesEncodingHex    RawSignBytesEncoding = "hex"
	RawSignBytesEncodingBase64 RawSignBytesEncoding = "base64"
)

// Hash function for bytes in the `raw_sign` RPC.
type RawSignBytesHashFunction string

const (
	RawSignBytesHashFunctionKeccak256  RawSignBytesHashFunction = "keccak256"
	RawSignBytesHashFunctionSha256     RawSignBytesHashFunction = "sha256"
	RawSignBytesHashFunctionBlake2b256 RawSignBytesHashFunction = "blake2b256"
)

// Parameters for hashing and signing bytes with the `raw_sign` RPC.
//
// The properties Bytes, Encoding, HashFunction are required.
type RawSignBytesParams struct {
	// The bytes to hash and sign.
	Bytes string `json:"bytes" api:"required"`
	// Encoding scheme for bytes in the `raw_sign` RPC.
	//
	// Any of "utf-8", "hex", "base64".
	Encoding RawSignBytesEncoding `json:"encoding,omitzero" api:"required"`
	// Hash function for bytes in the `raw_sign` RPC.
	//
	// Any of "keccak256", "sha256", "blake2b256".
	HashFunction RawSignBytesHashFunction `json:"hash_function,omitzero" api:"required"`
	paramObj
}

func (r RawSignBytesParams) MarshalJSON() (data []byte, err error) {
	type shadow RawSignBytesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RawSignBytesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func RawSignInputParamsOfRawSignHashs(hash Hex) RawSignInputParamsUnion {
	var variant RawSignHashParams
	variant.Hash = hash
	return RawSignInputParamsUnion{OfRawSignHashs: &variant}
}

func RawSignInputParamsOfRawSignBytess(bytes string, encoding RawSignBytesEncoding, hashFunction RawSignBytesHashFunction) RawSignInputParamsUnion {
	var variant RawSignBytesParams
	variant.Bytes = bytes
	variant.Encoding = encoding
	variant.HashFunction = hashFunction
	return RawSignInputParamsUnion{OfRawSignBytess: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RawSignInputParamsUnion struct {
	OfRawSignHashs  *RawSignHashParams  `json:",omitzero,inline"`
	OfRawSignBytess *RawSignBytesParams `json:",omitzero,inline"`
	paramUnion
}

func (u RawSignInputParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRawSignHashs, u.OfRawSignBytess)
}
func (u *RawSignInputParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Provide either `hash` (to sign a pre-computed hash) OR `bytes`, `encoding`, and
// `hash_function` (to hash and then sign). These options are mutually exclusive.
//
// The property Params is required.
type RawSignInput struct {
	// Parameters for the `raw_sign` RPC.
	Params RawSignInputParamsUnion `json:"params,omitzero" api:"required"`
	paramObj
}

func (r RawSignInput) MarshalJSON() (data []byte, err error) {
	type shadow RawSignInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RawSignInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data returned by the `raw_sign` RPC.
type RawSignResponseData struct {
	// Any of "hex".
	Encoding RawSignResponseDataEncoding `json:"encoding" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	Signature Hex `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawSignResponseData) RawJSON() string { return r.JSON.raw }
func (r *RawSignResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawSignResponseDataEncoding string

const (
	RawSignResponseDataEncodingHex RawSignResponseDataEncoding = "hex"
)

// Response to the `raw_sign` RPC.
type RawSignResponse struct {
	// Data returned by the `raw_sign` RPC.
	Data RawSignResponseData `json:"data" api:"required"`
	// Any of "raw_sign".
	Method RawSignResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawSignResponse) RawJSON() string { return r.JSON.raw }
func (r *RawSignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawSignResponseMethod string

const (
	RawSignResponseMethodRawSign RawSignResponseMethod = "raw_sign"
)

// A signed EIP-7702 authorization that delegates code execution to a contract
// address.
type EthereumSign7702AuthorizationResp struct {
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID  QuantityUnionResp `json:"chain_id" api:"required"`
	Contract string            `json:"contract" api:"required"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Nonce QuantityUnionResp `json:"nonce" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	R Hex `json:"r" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	S       Hex     `json:"s" api:"required"`
	YParity float64 `json:"y_parity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Contract    respjson.Field
		Nonce       respjson.Field
		R           respjson.Field
		S           respjson.Field
		YParity     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSign7702AuthorizationResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSign7702AuthorizationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSign7702AuthorizationResp to a
// EthereumSign7702Authorization.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSign7702Authorization.Overrides()
func (r EthereumSign7702AuthorizationResp) ToParam() EthereumSign7702Authorization {
	return param.Override[EthereumSign7702Authorization](json.RawMessage(r.RawJSON()))
}

// A signed EIP-7702 authorization that delegates code execution to a contract
// address.
//
// The properties ChainID, Contract, Nonce, R, S, YParity are required.
type EthereumSign7702Authorization struct {
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID  QuantityUnion `json:"chain_id,omitzero" api:"required"`
	Contract string        `json:"contract" api:"required"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Nonce QuantityUnion `json:"nonce,omitzero" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	R Hex `json:"r" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	S       Hex     `json:"s" api:"required"`
	YParity float64 `json:"y_parity" api:"required"`
	paramObj
}

func (r EthereumSign7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSign7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSign7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An unsigned standard Ethereum transaction object. Supports EVM transaction types
// 0, 1, 2, and 4.
type UnsignedStandardEthereumTransactionResp struct {
	AuthorizationList []EthereumSign7702AuthorizationResp `json:"authorization_list"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID QuantityUnionResp `json:"chain_id"`
	// A hex-encoded string prefixed with '0x'.
	Data Hex    `json:"data"`
	From string `json:"from"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	GasLimit QuantityUnionResp `json:"gas_limit"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	GasPrice QuantityUnionResp `json:"gas_price"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	MaxFeePerGas QuantityUnionResp `json:"max_fee_per_gas"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	MaxPriorityFeePerGas QuantityUnionResp `json:"max_priority_fee_per_gas"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Nonce QuantityUnionResp `json:"nonce"`
	To    string            `json:"to"`
	// Any of 0, 1, 2, 4.
	Type float64 `json:"type"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Value QuantityUnionResp `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationList    respjson.Field
		ChainID              respjson.Field
		Data                 respjson.Field
		From                 respjson.Field
		GasLimit             respjson.Field
		GasPrice             respjson.Field
		MaxFeePerGas         respjson.Field
		MaxPriorityFeePerGas respjson.Field
		Nonce                respjson.Field
		To                   respjson.Field
		Type                 respjson.Field
		Value                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnsignedStandardEthereumTransactionResp) RawJSON() string { return r.JSON.raw }
func (r *UnsignedStandardEthereumTransactionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UnsignedStandardEthereumTransactionResp to a
// UnsignedStandardEthereumTransaction.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UnsignedStandardEthereumTransaction.Overrides()
func (r UnsignedStandardEthereumTransactionResp) ToParam() UnsignedStandardEthereumTransaction {
	return param.Override[UnsignedStandardEthereumTransaction](json.RawMessage(r.RawJSON()))
}

// An unsigned standard Ethereum transaction object. Supports EVM transaction types
// 0, 1, 2, and 4.
type UnsignedStandardEthereumTransaction struct {
	// A hex-encoded string prefixed with '0x'.
	Data              param.Opt[Hex]                  `json:"data,omitzero"`
	From              param.Opt[string]               `json:"from,omitzero"`
	To                param.Opt[string]               `json:"to,omitzero"`
	AuthorizationList []EthereumSign7702Authorization `json:"authorization_list,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID QuantityUnion `json:"chain_id,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	GasLimit QuantityUnion `json:"gas_limit,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	GasPrice QuantityUnion `json:"gas_price,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	MaxFeePerGas QuantityUnion `json:"max_fee_per_gas,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	MaxPriorityFeePerGas QuantityUnion `json:"max_priority_fee_per_gas,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Nonce QuantityUnion `json:"nonce,omitzero"`
	// Any of 0, 1, 2, 4.
	Type float64 `json:"type,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Value QuantityUnion `json:"value,omitzero"`
	paramObj
}

func (r UnsignedStandardEthereumTransaction) MarshalJSON() (data []byte, err error) {
	type shadow UnsignedStandardEthereumTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnsignedStandardEthereumTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UnsignedStandardEthereumTransaction](
		"type", 0, 1, 2, 4,
	)
}

// An ERC-4337 user operation.
type UserOperationInputResp struct {
	// A hex-encoded string prefixed with '0x'.
	CallData Hex `json:"call_data" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	CallGasLimit Hex `json:"call_gas_limit" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	MaxFeePerGas Hex `json:"max_fee_per_gas" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	MaxPriorityFeePerGas Hex `json:"max_priority_fee_per_gas" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	Nonce Hex `json:"nonce" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	PreVerificationGas Hex    `json:"pre_verification_gas" api:"required"`
	Sender             string `json:"sender" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	VerificationGasLimit Hex    `json:"verification_gas_limit" api:"required"`
	Paymaster            string `json:"paymaster"`
	// A hex-encoded string prefixed with '0x'.
	PaymasterData Hex `json:"paymaster_data"`
	// A hex-encoded string prefixed with '0x'.
	PaymasterPostOpGasLimit Hex `json:"paymaster_post_op_gas_limit"`
	// A hex-encoded string prefixed with '0x'.
	PaymasterVerificationGasLimit Hex `json:"paymaster_verification_gas_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallData                      respjson.Field
		CallGasLimit                  respjson.Field
		MaxFeePerGas                  respjson.Field
		MaxPriorityFeePerGas          respjson.Field
		Nonce                         respjson.Field
		PreVerificationGas            respjson.Field
		Sender                        respjson.Field
		VerificationGasLimit          respjson.Field
		Paymaster                     respjson.Field
		PaymasterData                 respjson.Field
		PaymasterPostOpGasLimit       respjson.Field
		PaymasterVerificationGasLimit respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserOperationInputResp) RawJSON() string { return r.JSON.raw }
func (r *UserOperationInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserOperationInputResp to a UserOperationInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserOperationInput.Overrides()
func (r UserOperationInputResp) ToParam() UserOperationInput {
	return param.Override[UserOperationInput](json.RawMessage(r.RawJSON()))
}

// An ERC-4337 user operation.
//
// The properties CallData, CallGasLimit, MaxFeePerGas, MaxPriorityFeePerGas,
// Nonce, PreVerificationGas, Sender, VerificationGasLimit are required.
type UserOperationInput struct {
	// A hex-encoded string prefixed with '0x'.
	CallData Hex `json:"call_data" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	CallGasLimit Hex `json:"call_gas_limit" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	MaxFeePerGas Hex `json:"max_fee_per_gas" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	MaxPriorityFeePerGas Hex `json:"max_priority_fee_per_gas" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	Nonce Hex `json:"nonce" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	PreVerificationGas Hex    `json:"pre_verification_gas" api:"required"`
	Sender             string `json:"sender" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	VerificationGasLimit Hex               `json:"verification_gas_limit" api:"required"`
	Paymaster            param.Opt[string] `json:"paymaster,omitzero"`
	// A hex-encoded string prefixed with '0x'.
	PaymasterData param.Opt[Hex] `json:"paymaster_data,omitzero"`
	// A hex-encoded string prefixed with '0x'.
	PaymasterPostOpGasLimit param.Opt[Hex] `json:"paymaster_post_op_gas_limit,omitzero"`
	// A hex-encoded string prefixed with '0x'.
	PaymasterVerificationGasLimit param.Opt[Hex] `json:"paymaster_verification_gas_limit,omitzero"`
	paramObj
}

func (r UserOperationInput) MarshalJSON() (data []byte, err error) {
	type shadow UserOperationInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserOperationInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TypedDataDomainInputParams map[string]any

// A single field definition in an EIP-712 typed data type.
type TypedDataTypeFieldInputResp struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TypedDataTypeFieldInputResp) RawJSON() string { return r.JSON.raw }
func (r *TypedDataTypeFieldInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TypedDataTypeFieldInputResp to a TypedDataTypeFieldInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TypedDataTypeFieldInput.Overrides()
func (r TypedDataTypeFieldInputResp) ToParam() TypedDataTypeFieldInput {
	return param.Override[TypedDataTypeFieldInput](json.RawMessage(r.RawJSON()))
}

// A single field definition in an EIP-712 typed data type.
//
// The properties Name, Type are required.
type TypedDataTypeFieldInput struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	paramObj
}

func (r TypedDataTypeFieldInput) MarshalJSON() (data []byte, err error) {
	type shadow TypedDataTypeFieldInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TypedDataTypeFieldInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TypedDataTypesInputParamsResp map[string][]TypedDataTypeFieldInputResp

type TypedDataTypesInputParams map[string][]TypedDataTypeFieldInput

// Parameters for the EVM `personal_sign` RPC.
type EthereumPersonalSignRpcInputParamsResp struct {
	// Any of "utf-8", "hex".
	Encoding EthereumPersonalSignRpcInputParamsEncoding `json:"encoding" api:"required"`
	Message  string                                     `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumPersonalSignRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumPersonalSignRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumPersonalSignRpcInputParamsResp to a
// EthereumPersonalSignRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumPersonalSignRpcInputParams.Overrides()
func (r EthereumPersonalSignRpcInputParamsResp) ToParam() EthereumPersonalSignRpcInputParams {
	return param.Override[EthereumPersonalSignRpcInputParams](json.RawMessage(r.RawJSON()))
}

type EthereumPersonalSignRpcInputParamsEncoding string

const (
	EthereumPersonalSignRpcInputParamsEncodingUtf8 EthereumPersonalSignRpcInputParamsEncoding = "utf-8"
	EthereumPersonalSignRpcInputParamsEncodingHex  EthereumPersonalSignRpcInputParamsEncoding = "hex"
)

// Parameters for the EVM `personal_sign` RPC.
//
// The properties Encoding, Message are required.
type EthereumPersonalSignRpcInputParams struct {
	// Any of "utf-8", "hex".
	Encoding EthereumPersonalSignRpcInputParamsEncoding `json:"encoding,omitzero" api:"required"`
	Message  string                                     `json:"message" api:"required"`
	paramObj
}

func (r EthereumPersonalSignRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumPersonalSignRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumPersonalSignRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the EVM `personal_sign` RPC (EIP-191) to sign a message.
type EthereumPersonalSignRpcInputResp struct {
	// Any of "personal_sign".
	Method EthereumPersonalSignRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `personal_sign` RPC.
	Params  EthereumPersonalSignRpcInputParamsResp `json:"params" api:"required"`
	Address string                                 `json:"address"`
	// Any of "ethereum".
	ChainType EthereumPersonalSignRpcInputChainType `json:"chain_type"`
	WalletID  string                                `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumPersonalSignRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumPersonalSignRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumPersonalSignRpcInputResp to a
// EthereumPersonalSignRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumPersonalSignRpcInput.Overrides()
func (r EthereumPersonalSignRpcInputResp) ToParam() EthereumPersonalSignRpcInput {
	return param.Override[EthereumPersonalSignRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumPersonalSignRpcInputMethod string

const (
	EthereumPersonalSignRpcInputMethodPersonalSign EthereumPersonalSignRpcInputMethod = "personal_sign"
)

type EthereumPersonalSignRpcInputChainType string

const (
	EthereumPersonalSignRpcInputChainTypeEthereum EthereumPersonalSignRpcInputChainType = "ethereum"
)

// Executes the EVM `personal_sign` RPC (EIP-191) to sign a message.
//
// The properties Method, Params are required.
type EthereumPersonalSignRpcInput struct {
	// Any of "personal_sign".
	Method EthereumPersonalSignRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `personal_sign` RPC.
	Params   EthereumPersonalSignRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                  `json:"address,omitzero"`
	WalletID param.Opt[string]                  `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumPersonalSignRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumPersonalSignRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumPersonalSignRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumPersonalSignRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signTransaction` RPC.
type EthereumSignTransactionRpcInputParamsResp struct {
	// An unsigned standard Ethereum transaction object. Supports EVM transaction types
	// 0, 1, 2, and 4.
	Transaction UnsignedStandardEthereumTransactionResp `json:"transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTransactionRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTransactionRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignTransactionRpcInputParamsResp to a
// EthereumSignTransactionRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignTransactionRpcInputParams.Overrides()
func (r EthereumSignTransactionRpcInputParamsResp) ToParam() EthereumSignTransactionRpcInputParams {
	return param.Override[EthereumSignTransactionRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the EVM `eth_signTransaction` RPC.
//
// The property Transaction is required.
type EthereumSignTransactionRpcInputParams struct {
	// An unsigned standard Ethereum transaction object. Supports EVM transaction types
	// 0, 1, 2, and 4.
	Transaction UnsignedStandardEthereumTransaction `json:"transaction,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the EVM `eth_signTransaction` RPC to sign a transaction.
type EthereumSignTransactionRpcInputResp struct {
	// Any of "eth_signTransaction".
	Method EthereumSignTransactionRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `eth_signTransaction` RPC.
	Params  EthereumSignTransactionRpcInputParamsResp `json:"params" api:"required"`
	Address string                                    `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSignTransactionRpcInputChainType `json:"chain_type"`
	WalletID  string                                   `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTransactionRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTransactionRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignTransactionRpcInputResp to a
// EthereumSignTransactionRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignTransactionRpcInput.Overrides()
func (r EthereumSignTransactionRpcInputResp) ToParam() EthereumSignTransactionRpcInput {
	return param.Override[EthereumSignTransactionRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSignTransactionRpcInputMethod string

const (
	EthereumSignTransactionRpcInputMethodEthSignTransaction EthereumSignTransactionRpcInputMethod = "eth_signTransaction"
)

type EthereumSignTransactionRpcInputChainType string

const (
	EthereumSignTransactionRpcInputChainTypeEthereum EthereumSignTransactionRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_signTransaction` RPC to sign a transaction.
//
// The properties Method, Params are required.
type EthereumSignTransactionRpcInput struct {
	// Any of "eth_signTransaction".
	Method EthereumSignTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `eth_signTransaction` RPC.
	Params   EthereumSignTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                     `json:"address,omitzero"`
	WalletID param.Opt[string]                     `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSignTransactionRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSignTransactionRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_sendTransaction` RPC.
type EthereumSendTransactionRpcInputParamsResp struct {
	// An unsigned standard Ethereum transaction object. Supports EVM transaction types
	// 0, 1, 2, and 4.
	Transaction UnsignedStandardEthereumTransactionResp `json:"transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendTransactionRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSendTransactionRpcInputParamsResp to a
// EthereumSendTransactionRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSendTransactionRpcInputParams.Overrides()
func (r EthereumSendTransactionRpcInputParamsResp) ToParam() EthereumSendTransactionRpcInputParams {
	return param.Override[EthereumSendTransactionRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the EVM `eth_sendTransaction` RPC.
//
// The property Transaction is required.
type EthereumSendTransactionRpcInputParams struct {
	// An unsigned standard Ethereum transaction object. Supports EVM transaction types
	// 0, 1, 2, and 4.
	Transaction UnsignedStandardEthereumTransaction `json:"transaction,omitzero" api:"required"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the EVM `eth_sendTransaction` RPC to sign and broadcast a transaction.
type EthereumSendTransactionRpcInputResp struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2 Caip2 `json:"caip2" api:"required"`
	// Any of "eth_sendTransaction".
	Method EthereumSendTransactionRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `eth_sendTransaction` RPC.
	Params  EthereumSendTransactionRpcInputParamsResp `json:"params" api:"required"`
	Address string                                    `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSendTransactionRpcInputChainType `json:"chain_type"`
	// A hex-encoded string prefixed with '0x'.
	ExperimentalDataSuffix Hex    `json:"experimental_data_suffix"`
	ReferenceID            string `json:"reference_id"`
	Sponsor                bool   `json:"sponsor"`
	WalletID               string `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2                  respjson.Field
		Method                 respjson.Field
		Params                 respjson.Field
		Address                respjson.Field
		ChainType              respjson.Field
		ExperimentalDataSuffix respjson.Field
		ReferenceID            respjson.Field
		Sponsor                respjson.Field
		WalletID               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendTransactionRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSendTransactionRpcInputResp to a
// EthereumSendTransactionRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSendTransactionRpcInput.Overrides()
func (r EthereumSendTransactionRpcInputResp) ToParam() EthereumSendTransactionRpcInput {
	return param.Override[EthereumSendTransactionRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSendTransactionRpcInputMethod string

const (
	EthereumSendTransactionRpcInputMethodEthSendTransaction EthereumSendTransactionRpcInputMethod = "eth_sendTransaction"
)

type EthereumSendTransactionRpcInputChainType string

const (
	EthereumSendTransactionRpcInputChainTypeEthereum EthereumSendTransactionRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_sendTransaction` RPC to sign and broadcast a transaction.
//
// The properties Caip2, Method, Params are required.
type EthereumSendTransactionRpcInput struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2 Caip2 `json:"caip2" api:"required"`
	// Any of "eth_sendTransaction".
	Method EthereumSendTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `eth_sendTransaction` RPC.
	Params  EthereumSendTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                     `json:"address,omitzero"`
	// A hex-encoded string prefixed with '0x'.
	ExperimentalDataSuffix param.Opt[Hex]    `json:"experimental_data_suffix,omitzero"`
	ReferenceID            param.Opt[string] `json:"reference_id,omitzero"`
	Sponsor                param.Opt[bool]   `json:"sponsor,omitzero"`
	WalletID               param.Opt[string] `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSendTransactionRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSendTransactionRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EIP-712 typed data object.
type EthereumTypedDataInputResp struct {
	// The domain parameters for EIP-712 typed data signing.
	Domain      TypedDataDomainInputParams `json:"domain" api:"required"`
	Message     map[string]any             `json:"message" api:"required"`
	PrimaryType string                     `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParamsResp `json:"types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Message     respjson.Field
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumTypedDataInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumTypedDataInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumTypedDataInputResp to a EthereumTypedDataInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumTypedDataInput.Overrides()
func (r EthereumTypedDataInputResp) ToParam() EthereumTypedDataInput {
	return param.Override[EthereumTypedDataInput](json.RawMessage(r.RawJSON()))
}

// EIP-712 typed data object.
//
// The properties Domain, Message, PrimaryType, Types are required.
type EthereumTypedDataInput struct {
	// The domain parameters for EIP-712 typed data signing.
	Domain      TypedDataDomainInputParams `json:"domain,omitzero" api:"required"`
	Message     map[string]any             `json:"message,omitzero" api:"required"`
	PrimaryType string                     `json:"primary_type" api:"required"`
	// The type definitions for EIP-712 typed data signing.
	Types TypedDataTypesInputParams `json:"types,omitzero" api:"required"`
	paramObj
}

func (r EthereumTypedDataInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTypedDataInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTypedDataInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signTypedData_v4` RPC.
type EthereumSignTypedDataRpcInputParamsResp struct {
	// EIP-712 typed data object.
	TypedData EthereumTypedDataInputResp `json:"typed_data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TypedData   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTypedDataRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTypedDataRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignTypedDataRpcInputParamsResp to a
// EthereumSignTypedDataRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignTypedDataRpcInputParams.Overrides()
func (r EthereumSignTypedDataRpcInputParamsResp) ToParam() EthereumSignTypedDataRpcInputParams {
	return param.Override[EthereumSignTypedDataRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the EVM `eth_signTypedData_v4` RPC.
//
// The property TypedData is required.
type EthereumSignTypedDataRpcInputParams struct {
	// EIP-712 typed data object.
	TypedData EthereumTypedDataInput `json:"typed_data,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the EVM `eth_signTypedData_v4` RPC (EIP-712) to sign a typed data
// object.
type EthereumSignTypedDataRpcInputResp struct {
	// Any of "eth_signTypedData_v4".
	Method EthereumSignTypedDataRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `eth_signTypedData_v4` RPC.
	Params  EthereumSignTypedDataRpcInputParamsResp `json:"params" api:"required"`
	Address string                                  `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSignTypedDataRpcInputChainType `json:"chain_type"`
	WalletID  string                                 `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTypedDataRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTypedDataRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignTypedDataRpcInputResp to a
// EthereumSignTypedDataRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignTypedDataRpcInput.Overrides()
func (r EthereumSignTypedDataRpcInputResp) ToParam() EthereumSignTypedDataRpcInput {
	return param.Override[EthereumSignTypedDataRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSignTypedDataRpcInputMethod string

const (
	EthereumSignTypedDataRpcInputMethodEthSignTypedDataV4 EthereumSignTypedDataRpcInputMethod = "eth_signTypedData_v4"
)

type EthereumSignTypedDataRpcInputChainType string

const (
	EthereumSignTypedDataRpcInputChainTypeEthereum EthereumSignTypedDataRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_signTypedData_v4` RPC (EIP-712) to sign a typed data
// object.
//
// The properties Method, Params are required.
type EthereumSignTypedDataRpcInput struct {
	// Any of "eth_signTypedData_v4".
	Method EthereumSignTypedDataRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `eth_signTypedData_v4` RPC.
	Params   EthereumSignTypedDataRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                   `json:"address,omitzero"`
	WalletID param.Opt[string]                   `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSignTypedDataRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSignTypedDataRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `secp256k1_sign` RPC.
type EthereumSecp256k1SignRpcInputParamsResp struct {
	// A hex-encoded string prefixed with '0x'.
	Hash Hex `json:"hash" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSecp256k1SignRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSecp256k1SignRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSecp256k1SignRpcInputParamsResp to a
// EthereumSecp256k1SignRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSecp256k1SignRpcInputParams.Overrides()
func (r EthereumSecp256k1SignRpcInputParamsResp) ToParam() EthereumSecp256k1SignRpcInputParams {
	return param.Override[EthereumSecp256k1SignRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the EVM `secp256k1_sign` RPC.
//
// The property Hash is required.
type EthereumSecp256k1SignRpcInputParams struct {
	// A hex-encoded string prefixed with '0x'.
	Hash Hex `json:"hash" api:"required"`
	paramObj
}

func (r EthereumSecp256k1SignRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSecp256k1SignRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSecp256k1SignRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signs a raw hash on the secp256k1 curve.
type EthereumSecp256k1SignRpcInputResp struct {
	// Any of "secp256k1_sign".
	Method EthereumSecp256k1SignRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `secp256k1_sign` RPC.
	Params  EthereumSecp256k1SignRpcInputParamsResp `json:"params" api:"required"`
	Address string                                  `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSecp256k1SignRpcInputChainType `json:"chain_type"`
	WalletID  string                                 `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSecp256k1SignRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSecp256k1SignRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSecp256k1SignRpcInputResp to a
// EthereumSecp256k1SignRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSecp256k1SignRpcInput.Overrides()
func (r EthereumSecp256k1SignRpcInputResp) ToParam() EthereumSecp256k1SignRpcInput {
	return param.Override[EthereumSecp256k1SignRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSecp256k1SignRpcInputMethod string

const (
	EthereumSecp256k1SignRpcInputMethodSecp256k1Sign EthereumSecp256k1SignRpcInputMethod = "secp256k1_sign"
)

type EthereumSecp256k1SignRpcInputChainType string

const (
	EthereumSecp256k1SignRpcInputChainTypeEthereum EthereumSecp256k1SignRpcInputChainType = "ethereum"
)

// Signs a raw hash on the secp256k1 curve.
//
// The properties Method, Params are required.
type EthereumSecp256k1SignRpcInput struct {
	// Any of "secp256k1_sign".
	Method EthereumSecp256k1SignRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `secp256k1_sign` RPC.
	Params   EthereumSecp256k1SignRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                   `json:"address,omitzero"`
	WalletID param.Opt[string]                   `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSecp256k1SignRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSecp256k1SignRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSecp256k1SignRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSecp256k1SignRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_sign7702Authorization` RPC.
type EthereumSign7702AuthorizationRpcInputParamsResp struct {
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID  QuantityUnionResp `json:"chain_id" api:"required"`
	Contract string            `json:"contract" api:"required"`
	// Any of "self".
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor `json:"executor"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Nonce QuantityUnionResp `json:"nonce"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Contract    respjson.Field
		Executor    respjson.Field
		Nonce       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSign7702AuthorizationRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSign7702AuthorizationRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSign7702AuthorizationRpcInputParamsResp to a
// EthereumSign7702AuthorizationRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSign7702AuthorizationRpcInputParams.Overrides()
func (r EthereumSign7702AuthorizationRpcInputParamsResp) ToParam() EthereumSign7702AuthorizationRpcInputParams {
	return param.Override[EthereumSign7702AuthorizationRpcInputParams](json.RawMessage(r.RawJSON()))
}

type EthereumSign7702AuthorizationRpcInputParamsExecutor string

const (
	EthereumSign7702AuthorizationRpcInputParamsExecutorSelf EthereumSign7702AuthorizationRpcInputParamsExecutor = "self"
)

// Parameters for the EVM `eth_sign7702Authorization` RPC.
//
// The properties ChainID, Contract are required.
type EthereumSign7702AuthorizationRpcInputParams struct {
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID  QuantityUnion `json:"chain_id,omitzero" api:"required"`
	Contract string        `json:"contract" api:"required"`
	// Any of "self".
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor `json:"executor,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Nonce QuantityUnion `json:"nonce,omitzero"`
	paramObj
}

func (r EthereumSign7702AuthorizationRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSign7702AuthorizationRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSign7702AuthorizationRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signs an EIP-7702 authorization.
type EthereumSign7702AuthorizationRpcInputResp struct {
	// Any of "eth_sign7702Authorization".
	Method EthereumSign7702AuthorizationRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `eth_sign7702Authorization` RPC.
	Params  EthereumSign7702AuthorizationRpcInputParamsResp `json:"params" api:"required"`
	Address string                                          `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSign7702AuthorizationRpcInputChainType `json:"chain_type"`
	WalletID  string                                         `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSign7702AuthorizationRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSign7702AuthorizationRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSign7702AuthorizationRpcInputResp to a
// EthereumSign7702AuthorizationRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSign7702AuthorizationRpcInput.Overrides()
func (r EthereumSign7702AuthorizationRpcInputResp) ToParam() EthereumSign7702AuthorizationRpcInput {
	return param.Override[EthereumSign7702AuthorizationRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSign7702AuthorizationRpcInputMethod string

const (
	EthereumSign7702AuthorizationRpcInputMethodEthSign7702Authorization EthereumSign7702AuthorizationRpcInputMethod = "eth_sign7702Authorization"
)

type EthereumSign7702AuthorizationRpcInputChainType string

const (
	EthereumSign7702AuthorizationRpcInputChainTypeEthereum EthereumSign7702AuthorizationRpcInputChainType = "ethereum"
)

// Signs an EIP-7702 authorization.
//
// The properties Method, Params are required.
type EthereumSign7702AuthorizationRpcInput struct {
	// Any of "eth_sign7702Authorization".
	Method EthereumSign7702AuthorizationRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `eth_sign7702Authorization` RPC.
	Params   EthereumSign7702AuthorizationRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                           `json:"address,omitzero"`
	WalletID param.Opt[string]                           `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSign7702AuthorizationRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSign7702AuthorizationRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSign7702AuthorizationRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSign7702AuthorizationRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signUserOperation` RPC.
type EthereumSignUserOperationRpcInputParamsResp struct {
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID  QuantityUnionResp `json:"chain_id" api:"required"`
	Contract string            `json:"contract" api:"required"`
	// An ERC-4337 user operation.
	UserOperation UserOperationInputResp `json:"user_operation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID       respjson.Field
		Contract      respjson.Field
		UserOperation respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignUserOperationRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignUserOperationRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignUserOperationRpcInputParamsResp to a
// EthereumSignUserOperationRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignUserOperationRpcInputParams.Overrides()
func (r EthereumSignUserOperationRpcInputParamsResp) ToParam() EthereumSignUserOperationRpcInputParams {
	return param.Override[EthereumSignUserOperationRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the EVM `eth_signUserOperation` RPC.
//
// The properties ChainID, Contract, UserOperation are required.
type EthereumSignUserOperationRpcInputParams struct {
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	ChainID  QuantityUnion `json:"chain_id,omitzero" api:"required"`
	Contract string        `json:"contract" api:"required"`
	// An ERC-4337 user operation.
	UserOperation UserOperationInput `json:"user_operation,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignUserOperationRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes an RPC method to hash and sign a UserOperation.
type EthereumSignUserOperationRpcInputResp struct {
	// Any of "eth_signUserOperation".
	Method EthereumSignUserOperationRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `eth_signUserOperation` RPC.
	Params  EthereumSignUserOperationRpcInputParamsResp `json:"params" api:"required"`
	Address string                                      `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSignUserOperationRpcInputChainType `json:"chain_type"`
	WalletID  string                                     `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignUserOperationRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignUserOperationRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignUserOperationRpcInputResp to a
// EthereumSignUserOperationRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignUserOperationRpcInput.Overrides()
func (r EthereumSignUserOperationRpcInputResp) ToParam() EthereumSignUserOperationRpcInput {
	return param.Override[EthereumSignUserOperationRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSignUserOperationRpcInputMethod string

const (
	EthereumSignUserOperationRpcInputMethodEthSignUserOperation EthereumSignUserOperationRpcInputMethod = "eth_signUserOperation"
)

type EthereumSignUserOperationRpcInputChainType string

const (
	EthereumSignUserOperationRpcInputChainTypeEthereum EthereumSignUserOperationRpcInputChainType = "ethereum"
)

// Executes an RPC method to hash and sign a UserOperation.
//
// The properties Method, Params are required.
type EthereumSignUserOperationRpcInput struct {
	// Any of "eth_signUserOperation".
	Method EthereumSignUserOperationRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `eth_signUserOperation` RPC.
	Params   EthereumSignUserOperationRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                       `json:"address,omitzero"`
	WalletID param.Opt[string]                       `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSignUserOperationRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSignUserOperationRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single call within a batched wallet_sendCalls request.
type EthereumSendCallsCallResp struct {
	To string `json:"to" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	Data Hex `json:"data"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Value QuantityUnionResp `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		To          respjson.Field
		Data        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendCallsCallResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendCallsCallResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSendCallsCallResp to a EthereumSendCallsCall.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSendCallsCall.Overrides()
func (r EthereumSendCallsCallResp) ToParam() EthereumSendCallsCall {
	return param.Override[EthereumSendCallsCall](json.RawMessage(r.RawJSON()))
}

// A single call within a batched wallet_sendCalls request.
//
// The property To is required.
type EthereumSendCallsCall struct {
	To string `json:"to" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	Data param.Opt[Hex] `json:"data,omitzero"`
	// A quantity value that can be either a hex string starting with '0x' or a
	// non-negative integer.
	Value QuantityUnion `json:"value,omitzero"`
	paramObj
}

func (r EthereumSendCallsCall) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendCallsCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendCallsCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the `wallet_sendCalls` RPC.
type EthereumSendCallsRpcInputParamsResp struct {
	Calls []EthereumSendCallsCallResp `json:"calls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Calls       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendCallsRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendCallsRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSendCallsRpcInputParamsResp to a
// EthereumSendCallsRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSendCallsRpcInputParams.Overrides()
func (r EthereumSendCallsRpcInputParamsResp) ToParam() EthereumSendCallsRpcInputParams {
	return param.Override[EthereumSendCallsRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the `wallet_sendCalls` RPC.
//
// The property Calls is required.
type EthereumSendCallsRpcInputParams struct {
	Calls []EthereumSendCallsCall `json:"calls,omitzero" api:"required"`
	paramObj
}

func (r EthereumSendCallsRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendCallsRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendCallsRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the `wallet_sendCalls` RPC (EIP-5792) to batch multiple calls into a
// single atomic transaction.
type EthereumSendCallsRpcInputResp struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2 Caip2 `json:"caip2" api:"required"`
	// Any of "wallet_sendCalls".
	Method EthereumSendCallsRpcInputMethod `json:"method" api:"required"`
	// Parameters for the `wallet_sendCalls` RPC.
	Params  EthereumSendCallsRpcInputParamsResp `json:"params" api:"required"`
	Address string                              `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSendCallsRpcInputChainType `json:"chain_type"`
	// A hex-encoded string prefixed with '0x'.
	ExperimentalDataSuffix Hex    `json:"experimental_data_suffix"`
	Sponsor                bool   `json:"sponsor"`
	WalletID               string `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2                  respjson.Field
		Method                 respjson.Field
		Params                 respjson.Field
		Address                respjson.Field
		ChainType              respjson.Field
		ExperimentalDataSuffix respjson.Field
		Sponsor                respjson.Field
		WalletID               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendCallsRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendCallsRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSendCallsRpcInputResp to a
// EthereumSendCallsRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSendCallsRpcInput.Overrides()
func (r EthereumSendCallsRpcInputResp) ToParam() EthereumSendCallsRpcInput {
	return param.Override[EthereumSendCallsRpcInput](json.RawMessage(r.RawJSON()))
}

type EthereumSendCallsRpcInputMethod string

const (
	EthereumSendCallsRpcInputMethodWalletSendCalls EthereumSendCallsRpcInputMethod = "wallet_sendCalls"
)

type EthereumSendCallsRpcInputChainType string

const (
	EthereumSendCallsRpcInputChainTypeEthereum EthereumSendCallsRpcInputChainType = "ethereum"
)

// Executes the `wallet_sendCalls` RPC (EIP-5792) to batch multiple calls into a
// single atomic transaction.
//
// The properties Caip2, Method, Params are required.
type EthereumSendCallsRpcInput struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2 Caip2 `json:"caip2" api:"required"`
	// Any of "wallet_sendCalls".
	Method EthereumSendCallsRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the `wallet_sendCalls` RPC.
	Params  EthereumSendCallsRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]               `json:"address,omitzero"`
	// A hex-encoded string prefixed with '0x'.
	ExperimentalDataSuffix param.Opt[Hex]    `json:"experimental_data_suffix,omitzero"`
	Sponsor                param.Opt[bool]   `json:"sponsor,omitzero"`
	WalletID               param.Opt[string] `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSendCallsRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSendCallsRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendCallsRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendCallsRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data returned by the EVM `personal_sign` RPC.
type EthereumPersonalSignRpcResponseData struct {
	// Any of "hex".
	Encoding  EthereumPersonalSignRpcResponseDataEncoding `json:"encoding" api:"required"`
	Signature string                                      `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumPersonalSignRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumPersonalSignRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumPersonalSignRpcResponseDataEncoding string

const (
	EthereumPersonalSignRpcResponseDataEncodingHex EthereumPersonalSignRpcResponseDataEncoding = "hex"
)

// Response to the EVM `personal_sign` RPC.
type EthereumPersonalSignRpcResponse struct {
	// Data returned by the EVM `personal_sign` RPC.
	Data EthereumPersonalSignRpcResponseData `json:"data" api:"required"`
	// Any of "personal_sign".
	Method EthereumPersonalSignRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumPersonalSignRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumPersonalSignRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumPersonalSignRpcResponseMethod string

const (
	EthereumPersonalSignRpcResponseMethodPersonalSign EthereumPersonalSignRpcResponseMethod = "personal_sign"
)

// Data returned by the EVM `eth_signTransaction` RPC.
type EthereumSignTransactionRpcResponseData struct {
	// Any of "rlp".
	Encoding          EthereumSignTransactionRpcResponseDataEncoding `json:"encoding" api:"required"`
	SignedTransaction string                                         `json:"signed_transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding          respjson.Field
		SignedTransaction respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTransactionRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTransactionRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTransactionRpcResponseDataEncoding string

const (
	EthereumSignTransactionRpcResponseDataEncodingRlp EthereumSignTransactionRpcResponseDataEncoding = "rlp"
)

// Response to the EVM `eth_signTransaction` RPC.
type EthereumSignTransactionRpcResponse struct {
	// Data returned by the EVM `eth_signTransaction` RPC.
	Data EthereumSignTransactionRpcResponseData `json:"data" api:"required"`
	// Any of "eth_signTransaction".
	Method EthereumSignTransactionRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTransactionRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTransactionRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTransactionRpcResponseMethod string

const (
	EthereumSignTransactionRpcResponseMethodEthSignTransaction EthereumSignTransactionRpcResponseMethod = "eth_signTransaction"
)

// Data returned by the EVM `eth_sendTransaction` RPC.
type EthereumSendTransactionRpcResponseData struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2         Caip2  `json:"caip2" api:"required"`
	Hash          string `json:"hash" api:"required"`
	ReferenceID   string `json:"reference_id" api:"nullable"`
	TransactionID string `json:"transaction_id"`
	// An unsigned standard Ethereum transaction object. Supports EVM transaction types
	// 0, 1, 2, and 4.
	TransactionRequest UnsignedStandardEthereumTransactionResp `json:"transaction_request"`
	UserOperationHash  string                                  `json:"user_operation_hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2              respjson.Field
		Hash               respjson.Field
		ReferenceID        respjson.Field
		TransactionID      respjson.Field
		TransactionRequest respjson.Field
		UserOperationHash  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendTransactionRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the EVM `eth_sendTransaction` RPC.
type EthereumSendTransactionRpcResponse struct {
	// Data returned by the EVM `eth_sendTransaction` RPC.
	Data EthereumSendTransactionRpcResponseData `json:"data" api:"required"`
	// Any of "eth_sendTransaction".
	Method EthereumSendTransactionRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendTransactionRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendTransactionRpcResponseMethod string

const (
	EthereumSendTransactionRpcResponseMethodEthSendTransaction EthereumSendTransactionRpcResponseMethod = "eth_sendTransaction"
)

// Data returned by the EVM `eth_signTypedData_v4` RPC.
type EthereumSignTypedDataRpcResponseData struct {
	// Any of "hex".
	Encoding  EthereumSignTypedDataRpcResponseDataEncoding `json:"encoding" api:"required"`
	Signature string                                       `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTypedDataRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTypedDataRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTypedDataRpcResponseDataEncoding string

const (
	EthereumSignTypedDataRpcResponseDataEncodingHex EthereumSignTypedDataRpcResponseDataEncoding = "hex"
)

// Response to the EVM `eth_signTypedData_v4` RPC.
type EthereumSignTypedDataRpcResponse struct {
	// Data returned by the EVM `eth_signTypedData_v4` RPC.
	Data EthereumSignTypedDataRpcResponseData `json:"data" api:"required"`
	// Any of "eth_signTypedData_v4".
	Method EthereumSignTypedDataRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignTypedDataRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTypedDataRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTypedDataRpcResponseMethod string

const (
	EthereumSignTypedDataRpcResponseMethodEthSignTypedDataV4 EthereumSignTypedDataRpcResponseMethod = "eth_signTypedData_v4"
)

// Data returned by the EVM `secp256k1_sign` RPC.
type EthereumSecp256k1SignRpcResponseData struct {
	// Any of "hex".
	Encoding EthereumSecp256k1SignRpcResponseDataEncoding `json:"encoding" api:"required"`
	// A hex-encoded string prefixed with '0x'.
	Signature Hex `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSecp256k1SignRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSecp256k1SignRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSecp256k1SignRpcResponseDataEncoding string

const (
	EthereumSecp256k1SignRpcResponseDataEncodingHex EthereumSecp256k1SignRpcResponseDataEncoding = "hex"
)

// Response to the EVM `secp256k1_sign` RPC.
type EthereumSecp256k1SignRpcResponse struct {
	// Data returned by the EVM `secp256k1_sign` RPC.
	Data EthereumSecp256k1SignRpcResponseData `json:"data" api:"required"`
	// Any of "secp256k1_sign".
	Method EthereumSecp256k1SignRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSecp256k1SignRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSecp256k1SignRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSecp256k1SignRpcResponseMethod string

const (
	EthereumSecp256k1SignRpcResponseMethodSecp256k1Sign EthereumSecp256k1SignRpcResponseMethod = "secp256k1_sign"
)

// Data returned by the EVM `eth_sign7702Authorization` RPC.
type EthereumSign7702AuthorizationRpcResponseData struct {
	// A signed EIP-7702 authorization that delegates code execution to a contract
	// address.
	Authorization EthereumSign7702AuthorizationResp `json:"authorization" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authorization respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSign7702AuthorizationRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSign7702AuthorizationRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the EVM `eth_sign7702Authorization` RPC.
type EthereumSign7702AuthorizationRpcResponse struct {
	// Data returned by the EVM `eth_sign7702Authorization` RPC.
	Data EthereumSign7702AuthorizationRpcResponseData `json:"data" api:"required"`
	// Any of "eth_sign7702Authorization".
	Method EthereumSign7702AuthorizationRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSign7702AuthorizationRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSign7702AuthorizationRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSign7702AuthorizationRpcResponseMethod string

const (
	EthereumSign7702AuthorizationRpcResponseMethodEthSign7702Authorization EthereumSign7702AuthorizationRpcResponseMethod = "eth_sign7702Authorization"
)

// Data returned by the EVM `eth_signUserOperation` RPC.
type EthereumSignUserOperationRpcResponseData struct {
	// Any of "hex".
	Encoding  EthereumSignUserOperationRpcResponseDataEncoding `json:"encoding" api:"required"`
	Signature string                                           `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignUserOperationRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignUserOperationRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignUserOperationRpcResponseDataEncoding string

const (
	EthereumSignUserOperationRpcResponseDataEncodingHex EthereumSignUserOperationRpcResponseDataEncoding = "hex"
)

// Response to the EVM `eth_signUserOperation` RPC.
type EthereumSignUserOperationRpcResponse struct {
	// Data returned by the EVM `eth_signUserOperation` RPC.
	Data EthereumSignUserOperationRpcResponseData `json:"data" api:"required"`
	// Any of "eth_signUserOperation".
	Method EthereumSignUserOperationRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignUserOperationRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignUserOperationRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignUserOperationRpcResponseMethod string

const (
	EthereumSignUserOperationRpcResponseMethodEthSignUserOperation EthereumSignUserOperationRpcResponseMethod = "eth_signUserOperation"
)

// Data returned by the `wallet_sendCalls` RPC.
type EthereumSendCallsRpcResponseData struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2         Caip2  `json:"caip2" api:"required"`
	TransactionID string `json:"transaction_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2         respjson.Field
		TransactionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendCallsRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendCallsRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the `wallet_sendCalls` RPC.
type EthereumSendCallsRpcResponse struct {
	// Data returned by the `wallet_sendCalls` RPC.
	Data EthereumSendCallsRpcResponseData `json:"data" api:"required"`
	// Any of "wallet_sendCalls".
	Method EthereumSendCallsRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendCallsRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendCallsRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendCallsRpcResponseMethod string

const (
	EthereumSendCallsRpcResponseMethodWalletSendCalls EthereumSendCallsRpcResponseMethod = "wallet_sendCalls"
)

// Parameters for the SVM `signTransaction` RPC.
type SolanaSignTransactionRpcInputParamsResp struct {
	// Any of "base64".
	Encoding    SolanaSignTransactionRpcInputParamsEncoding `json:"encoding" api:"required"`
	Transaction string                                      `json:"transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignTransactionRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignTransactionRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignTransactionRpcInputParamsResp to a
// SolanaSignTransactionRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignTransactionRpcInputParams.Overrides()
func (r SolanaSignTransactionRpcInputParamsResp) ToParam() SolanaSignTransactionRpcInputParams {
	return param.Override[SolanaSignTransactionRpcInputParams](json.RawMessage(r.RawJSON()))
}

type SolanaSignTransactionRpcInputParamsEncoding string

const (
	SolanaSignTransactionRpcInputParamsEncodingBase64 SolanaSignTransactionRpcInputParamsEncoding = "base64"
)

// Parameters for the SVM `signTransaction` RPC.
//
// The properties Encoding, Transaction are required.
type SolanaSignTransactionRpcInputParams struct {
	// Any of "base64".
	Encoding    SolanaSignTransactionRpcInputParamsEncoding `json:"encoding,omitzero" api:"required"`
	Transaction string                                      `json:"transaction" api:"required"`
	paramObj
}

func (r SolanaSignTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the SVM `signTransaction` RPC to sign a transaction.
type SolanaSignTransactionRpcInputResp struct {
	// Any of "signTransaction".
	Method SolanaSignTransactionRpcInputMethod `json:"method" api:"required"`
	// Parameters for the SVM `signTransaction` RPC.
	Params  SolanaSignTransactionRpcInputParamsResp `json:"params" api:"required"`
	Address string                                  `json:"address"`
	// Any of "solana".
	ChainType SolanaSignTransactionRpcInputChainType `json:"chain_type"`
	WalletID  string                                 `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignTransactionRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignTransactionRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignTransactionRpcInputResp to a
// SolanaSignTransactionRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignTransactionRpcInput.Overrides()
func (r SolanaSignTransactionRpcInputResp) ToParam() SolanaSignTransactionRpcInput {
	return param.Override[SolanaSignTransactionRpcInput](json.RawMessage(r.RawJSON()))
}

type SolanaSignTransactionRpcInputMethod string

const (
	SolanaSignTransactionRpcInputMethodSignTransaction SolanaSignTransactionRpcInputMethod = "signTransaction"
)

type SolanaSignTransactionRpcInputChainType string

const (
	SolanaSignTransactionRpcInputChainTypeSolana SolanaSignTransactionRpcInputChainType = "solana"
)

// Executes the SVM `signTransaction` RPC to sign a transaction.
//
// The properties Method, Params are required.
type SolanaSignTransactionRpcInput struct {
	// Any of "signTransaction".
	Method SolanaSignTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the SVM `signTransaction` RPC.
	Params   SolanaSignTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                   `json:"address,omitzero"`
	WalletID param.Opt[string]                   `json:"wallet_id,omitzero"`
	// Any of "solana".
	ChainType SolanaSignTransactionRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r SolanaSignTransactionRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignTransactionRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the SVM `signAndSendTransaction` RPC.
type SolanaSignAndSendTransactionRpcInputParamsResp struct {
	// Any of "base64".
	Encoding    SolanaSignAndSendTransactionRpcInputParamsEncoding `json:"encoding" api:"required"`
	Transaction string                                             `json:"transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Transaction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignAndSendTransactionRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignAndSendTransactionRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignAndSendTransactionRpcInputParamsResp to a
// SolanaSignAndSendTransactionRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignAndSendTransactionRpcInputParams.Overrides()
func (r SolanaSignAndSendTransactionRpcInputParamsResp) ToParam() SolanaSignAndSendTransactionRpcInputParams {
	return param.Override[SolanaSignAndSendTransactionRpcInputParams](json.RawMessage(r.RawJSON()))
}

type SolanaSignAndSendTransactionRpcInputParamsEncoding string

const (
	SolanaSignAndSendTransactionRpcInputParamsEncodingBase64 SolanaSignAndSendTransactionRpcInputParamsEncoding = "base64"
)

// Parameters for the SVM `signAndSendTransaction` RPC.
//
// The properties Encoding, Transaction are required.
type SolanaSignAndSendTransactionRpcInputParams struct {
	// Any of "base64".
	Encoding    SolanaSignAndSendTransactionRpcInputParamsEncoding `json:"encoding,omitzero" api:"required"`
	Transaction string                                             `json:"transaction" api:"required"`
	paramObj
}

func (r SolanaSignAndSendTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignAndSendTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignAndSendTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the SVM `signAndSendTransaction` RPC to sign and broadcast a
// transaction.
type SolanaSignAndSendTransactionRpcInputResp struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2 Caip2 `json:"caip2" api:"required"`
	// Any of "signAndSendTransaction".
	Method SolanaSignAndSendTransactionRpcInputMethod `json:"method" api:"required"`
	// Parameters for the SVM `signAndSendTransaction` RPC.
	Params  SolanaSignAndSendTransactionRpcInputParamsResp `json:"params" api:"required"`
	Address string                                         `json:"address"`
	// Any of "solana".
	ChainType   SolanaSignAndSendTransactionRpcInputChainType `json:"chain_type"`
	ReferenceID string                                        `json:"reference_id"`
	Sponsor     bool                                          `json:"sponsor"`
	WalletID    string                                        `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2       respjson.Field
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		ReferenceID respjson.Field
		Sponsor     respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignAndSendTransactionRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignAndSendTransactionRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignAndSendTransactionRpcInputResp to a
// SolanaSignAndSendTransactionRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignAndSendTransactionRpcInput.Overrides()
func (r SolanaSignAndSendTransactionRpcInputResp) ToParam() SolanaSignAndSendTransactionRpcInput {
	return param.Override[SolanaSignAndSendTransactionRpcInput](json.RawMessage(r.RawJSON()))
}

type SolanaSignAndSendTransactionRpcInputMethod string

const (
	SolanaSignAndSendTransactionRpcInputMethodSignAndSendTransaction SolanaSignAndSendTransactionRpcInputMethod = "signAndSendTransaction"
)

type SolanaSignAndSendTransactionRpcInputChainType string

const (
	SolanaSignAndSendTransactionRpcInputChainTypeSolana SolanaSignAndSendTransactionRpcInputChainType = "solana"
)

// Executes the SVM `signAndSendTransaction` RPC to sign and broadcast a
// transaction.
//
// The properties Caip2, Method, Params are required.
type SolanaSignAndSendTransactionRpcInput struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2 Caip2 `json:"caip2" api:"required"`
	// Any of "signAndSendTransaction".
	Method SolanaSignAndSendTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the SVM `signAndSendTransaction` RPC.
	Params      SolanaSignAndSendTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address     param.Opt[string]                          `json:"address,omitzero"`
	ReferenceID param.Opt[string]                          `json:"reference_id,omitzero"`
	Sponsor     param.Opt[bool]                            `json:"sponsor,omitzero"`
	WalletID    param.Opt[string]                          `json:"wallet_id,omitzero"`
	// Any of "solana".
	ChainType SolanaSignAndSendTransactionRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r SolanaSignAndSendTransactionRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignAndSendTransactionRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignAndSendTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the SVM `signMessage` RPC.
type SolanaSignMessageRpcInputParamsResp struct {
	// Any of "base64".
	Encoding SolanaSignMessageRpcInputParamsEncoding `json:"encoding" api:"required"`
	Message  string                                  `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignMessageRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignMessageRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignMessageRpcInputParamsResp to a
// SolanaSignMessageRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignMessageRpcInputParams.Overrides()
func (r SolanaSignMessageRpcInputParamsResp) ToParam() SolanaSignMessageRpcInputParams {
	return param.Override[SolanaSignMessageRpcInputParams](json.RawMessage(r.RawJSON()))
}

type SolanaSignMessageRpcInputParamsEncoding string

const (
	SolanaSignMessageRpcInputParamsEncodingBase64 SolanaSignMessageRpcInputParamsEncoding = "base64"
)

// Parameters for the SVM `signMessage` RPC.
//
// The properties Encoding, Message are required.
type SolanaSignMessageRpcInputParams struct {
	// Any of "base64".
	Encoding SolanaSignMessageRpcInputParamsEncoding `json:"encoding,omitzero" api:"required"`
	Message  string                                  `json:"message" api:"required"`
	paramObj
}

func (r SolanaSignMessageRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignMessageRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignMessageRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the SVM `signMessage` RPC to sign a message.
type SolanaSignMessageRpcInputResp struct {
	// Any of "signMessage".
	Method SolanaSignMessageRpcInputMethod `json:"method" api:"required"`
	// Parameters for the SVM `signMessage` RPC.
	Params  SolanaSignMessageRpcInputParamsResp `json:"params" api:"required"`
	Address string                              `json:"address"`
	// Any of "solana".
	ChainType SolanaSignMessageRpcInputChainType `json:"chain_type"`
	WalletID  string                             `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignMessageRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignMessageRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignMessageRpcInputResp to a
// SolanaSignMessageRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignMessageRpcInput.Overrides()
func (r SolanaSignMessageRpcInputResp) ToParam() SolanaSignMessageRpcInput {
	return param.Override[SolanaSignMessageRpcInput](json.RawMessage(r.RawJSON()))
}

type SolanaSignMessageRpcInputMethod string

const (
	SolanaSignMessageRpcInputMethodSignMessage SolanaSignMessageRpcInputMethod = "signMessage"
)

type SolanaSignMessageRpcInputChainType string

const (
	SolanaSignMessageRpcInputChainTypeSolana SolanaSignMessageRpcInputChainType = "solana"
)

// Executes the SVM `signMessage` RPC to sign a message.
//
// The properties Method, Params are required.
type SolanaSignMessageRpcInput struct {
	// Any of "signMessage".
	Method SolanaSignMessageRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the SVM `signMessage` RPC.
	Params   SolanaSignMessageRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]               `json:"address,omitzero"`
	WalletID param.Opt[string]               `json:"wallet_id,omitzero"`
	// Any of "solana".
	ChainType SolanaSignMessageRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r SolanaSignMessageRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignMessageRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignMessageRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data returned by the SVM `signTransaction` RPC.
type SolanaSignTransactionRpcResponseData struct {
	// Any of "base64".
	Encoding          SolanaSignTransactionRpcResponseDataEncoding `json:"encoding" api:"required"`
	SignedTransaction string                                       `json:"signed_transaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding          respjson.Field
		SignedTransaction respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignTransactionRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignTransactionRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolanaSignTransactionRpcResponseDataEncoding string

const (
	SolanaSignTransactionRpcResponseDataEncodingBase64 SolanaSignTransactionRpcResponseDataEncoding = "base64"
)

// Response to the SVM `signTransaction` RPC.
type SolanaSignTransactionRpcResponse struct {
	// Data returned by the SVM `signTransaction` RPC.
	Data SolanaSignTransactionRpcResponseData `json:"data" api:"required"`
	// Any of "signTransaction".
	Method SolanaSignTransactionRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignTransactionRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignTransactionRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolanaSignTransactionRpcResponseMethod string

const (
	SolanaSignTransactionRpcResponseMethodSignTransaction SolanaSignTransactionRpcResponseMethod = "signTransaction"
)

// Data returned by the SVM `signAndSendTransaction` RPC.
type SolanaSignAndSendTransactionRpcResponseData struct {
	// A valid CAIP-2 chain ID (e.g. 'eip155:1').
	Caip2         Caip2  `json:"caip2" api:"required"`
	Hash          string `json:"hash" api:"required"`
	ReferenceID   string `json:"reference_id" api:"nullable"`
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2         respjson.Field
		Hash          respjson.Field
		ReferenceID   respjson.Field
		TransactionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignAndSendTransactionRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignAndSendTransactionRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the SVM `signAndSendTransaction` RPC.
type SolanaSignAndSendTransactionRpcResponse struct {
	// Data returned by the SVM `signAndSendTransaction` RPC.
	Data SolanaSignAndSendTransactionRpcResponseData `json:"data" api:"required"`
	// Any of "signAndSendTransaction".
	Method SolanaSignAndSendTransactionRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignAndSendTransactionRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignAndSendTransactionRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolanaSignAndSendTransactionRpcResponseMethod string

const (
	SolanaSignAndSendTransactionRpcResponseMethodSignAndSendTransaction SolanaSignAndSendTransactionRpcResponseMethod = "signAndSendTransaction"
)

// Data returned by the SVM `signMessage` RPC.
type SolanaSignMessageRpcResponseData struct {
	// Any of "base64".
	Encoding  SolanaSignMessageRpcResponseDataEncoding `json:"encoding" api:"required"`
	Signature string                                   `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignMessageRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignMessageRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolanaSignMessageRpcResponseDataEncoding string

const (
	SolanaSignMessageRpcResponseDataEncodingBase64 SolanaSignMessageRpcResponseDataEncoding = "base64"
)

// Response to the SVM `signMessage` RPC.
type SolanaSignMessageRpcResponse struct {
	// Data returned by the SVM `signMessage` RPC.
	Data SolanaSignMessageRpcResponseData `json:"data" api:"required"`
	// Any of "signMessage".
	Method SolanaSignMessageRpcResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignMessageRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignMessageRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolanaSignMessageRpcResponseMethod string

const (
	SolanaSignMessageRpcResponseMethodSignMessage SolanaSignMessageRpcResponseMethod = "signMessage"
)

// The Spark network.
type SparkNetwork string

const (
	SparkNetworkMainnet SparkNetwork = "MAINNET"
	SparkNetworkRegtest SparkNetwork = "REGTEST"
)

// A Spark signing keyshare.
type SparkSigningKeyshare struct {
	OwnerIdentifiers []string          `json:"owner_identifiers" api:"required"`
	PublicKey        string            `json:"public_key" api:"required"`
	PublicShares     map[string]string `json:"public_shares" api:"required"`
	Threshold        float64           `json:"threshold" api:"required"`
	UpdatedTime      string            `json:"updated_time" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OwnerIdentifiers respjson.Field
		PublicKey        respjson.Field
		PublicShares     respjson.Field
		Threshold        respjson.Field
		UpdatedTime      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkSigningKeyshare) RawJSON() string { return r.JSON.raw }
func (r *SparkSigningKeyshare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark wallet leaf node.
type SparkWalletLeaf struct {
	ID string `json:"id" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network                SparkNetwork `json:"network" api:"required"`
	NodeTx                 string       `json:"node_tx" api:"required"`
	OwnerIdentityPublicKey string       `json:"owner_identity_public_key" api:"required"`
	RefundTx               string       `json:"refund_tx" api:"required"`
	Status                 string       `json:"status" api:"required"`
	TreeID                 string       `json:"tree_id" api:"required"`
	Value                  float64      `json:"value" api:"required"`
	VerifyingPublicKey     string       `json:"verifying_public_key" api:"required"`
	Vout                   float64      `json:"vout" api:"required"`
	ParentNodeID           string       `json:"parent_node_id"`
	// A Spark signing keyshare.
	SigningKeyshare SparkSigningKeyshare `json:"signing_keyshare"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Network                respjson.Field
		NodeTx                 respjson.Field
		OwnerIdentityPublicKey respjson.Field
		RefundTx               respjson.Field
		Status                 respjson.Field
		TreeID                 respjson.Field
		Value                  respjson.Field
		VerifyingPublicKey     respjson.Field
		Vout                   respjson.Field
		ParentNodeID           respjson.Field
		SigningKeyshare        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkWalletLeaf) RawJSON() string { return r.JSON.raw }
func (r *SparkWalletLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark transfer leaf.
type SparkTransferLeaf struct {
	IntermediateRefundTx string `json:"intermediate_refund_tx" api:"required"`
	SecretCipher         string `json:"secret_cipher" api:"required"`
	Signature            string `json:"signature" api:"required"`
	// A Spark wallet leaf node.
	Leaf SparkWalletLeaf `json:"leaf"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntermediateRefundTx respjson.Field
		SecretCipher         respjson.Field
		Signature            respjson.Field
		Leaf                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferLeaf) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark transfer.
type SparkTransfer struct {
	ID                        string              `json:"id" api:"required"`
	Leaves                    []SparkTransferLeaf `json:"leaves" api:"required"`
	ReceiverIdentityPublicKey string              `json:"receiver_identity_public_key" api:"required"`
	SenderIdentityPublicKey   string              `json:"sender_identity_public_key" api:"required"`
	Status                    string              `json:"status" api:"required"`
	TotalValue                float64             `json:"total_value" api:"required"`
	TransferDirection         string              `json:"transfer_direction" api:"required"`
	Type                      string              `json:"type" api:"required"`
	CreatedTime               string              `json:"created_time"`
	ExpiryTime                string              `json:"expiry_time"`
	UpdatedTime               string              `json:"updated_time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Leaves                    respjson.Field
		ReceiverIdentityPublicKey respjson.Field
		SenderIdentityPublicKey   respjson.Field
		Status                    respjson.Field
		TotalValue                respjson.Field
		TransferDirection         respjson.Field
		Type                      respjson.Field
		CreatedTime               respjson.Field
		ExpiryTime                respjson.Field
		UpdatedTime               respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransfer) RawJSON() string { return r.JSON.raw }
func (r *SparkTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for a Spark user token.
type SparkUserTokenMetadata struct {
	Decimals           float64 `json:"decimals" api:"required"`
	MaxSupply          string  `json:"max_supply" api:"required"`
	RawTokenIdentifier string  `json:"raw_token_identifier" api:"required"`
	TokenName          string  `json:"token_name" api:"required"`
	TokenPublicKey     string  `json:"token_public_key" api:"required"`
	TokenTicker        string  `json:"token_ticker" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decimals           respjson.Field
		MaxSupply          respjson.Field
		RawTokenIdentifier respjson.Field
		TokenName          respjson.Field
		TokenPublicKey     respjson.Field
		TokenTicker        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkUserTokenMetadata) RawJSON() string { return r.JSON.raw }
func (r *SparkUserTokenMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Balance of a Spark token.
type SparkTokenBalance struct {
	Balance string `json:"balance" api:"required"`
	// Metadata for a Spark user token.
	TokenMetadata SparkUserTokenMetadata `json:"token_metadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance       respjson.Field
		TokenMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTokenBalance) RawJSON() string { return r.JSON.raw }
func (r *SparkTokenBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The balance of a Spark wallet.
type SparkBalance struct {
	Balance       string                       `json:"balance" api:"required"`
	TokenBalances map[string]SparkTokenBalance `json:"token_balances" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance       respjson.Field
		TokenBalances respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkBalance) RawJSON() string { return r.JSON.raw }
func (r *SparkBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark token output.
type TokenOutputResp struct {
	OwnerPublicKey                string  `json:"owner_public_key" api:"required"`
	TokenAmount                   string  `json:"token_amount" api:"required"`
	ID                            string  `json:"id"`
	RevocationCommitment          string  `json:"revocation_commitment"`
	TokenIdentifier               string  `json:"token_identifier"`
	TokenPublicKey                string  `json:"token_public_key"`
	WithdrawBondSats              float64 `json:"withdraw_bond_sats"`
	WithdrawRelativeBlockLocktime float64 `json:"withdraw_relative_block_locktime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OwnerPublicKey                respjson.Field
		TokenAmount                   respjson.Field
		ID                            respjson.Field
		RevocationCommitment          respjson.Field
		TokenIdentifier               respjson.Field
		TokenPublicKey                respjson.Field
		WithdrawBondSats              respjson.Field
		WithdrawRelativeBlockLocktime respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenOutputResp) RawJSON() string { return r.JSON.raw }
func (r *TokenOutputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TokenOutputResp to a TokenOutput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TokenOutput.Overrides()
func (r TokenOutputResp) ToParam() TokenOutput {
	return param.Override[TokenOutput](json.RawMessage(r.RawJSON()))
}

// A Spark token output.
//
// The properties OwnerPublicKey, TokenAmount are required.
type TokenOutput struct {
	OwnerPublicKey                string             `json:"owner_public_key" api:"required"`
	TokenAmount                   string             `json:"token_amount" api:"required"`
	ID                            param.Opt[string]  `json:"id,omitzero"`
	RevocationCommitment          param.Opt[string]  `json:"revocation_commitment,omitzero"`
	TokenIdentifier               param.Opt[string]  `json:"token_identifier,omitzero"`
	TokenPublicKey                param.Opt[string]  `json:"token_public_key,omitzero"`
	WithdrawBondSats              param.Opt[float64] `json:"withdraw_bond_sats,omitzero"`
	WithdrawRelativeBlockLocktime param.Opt[float64] `json:"withdraw_relative_block_locktime,omitzero"`
	paramObj
}

func (r TokenOutput) MarshalJSON() (data []byte, err error) {
	type shadow TokenOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TokenOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark token output with its previous transaction data.
type OutputWithPreviousTransactionDataResp struct {
	PreviousTransactionHash string  `json:"previous_transaction_hash" api:"required"`
	PreviousTransactionVout float64 `json:"previous_transaction_vout" api:"required"`
	// A Spark token output.
	Output TokenOutputResp `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PreviousTransactionHash respjson.Field
		PreviousTransactionVout respjson.Field
		Output                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputWithPreviousTransactionDataResp) RawJSON() string { return r.JSON.raw }
func (r *OutputWithPreviousTransactionDataResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutputWithPreviousTransactionDataResp to a
// OutputWithPreviousTransactionData.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutputWithPreviousTransactionData.Overrides()
func (r OutputWithPreviousTransactionDataResp) ToParam() OutputWithPreviousTransactionData {
	return param.Override[OutputWithPreviousTransactionData](json.RawMessage(r.RawJSON()))
}

// A Spark token output with its previous transaction data.
//
// The properties PreviousTransactionHash, PreviousTransactionVout are required.
type OutputWithPreviousTransactionData struct {
	PreviousTransactionHash string  `json:"previous_transaction_hash" api:"required"`
	PreviousTransactionVout float64 `json:"previous_transaction_vout" api:"required"`
	// A Spark token output.
	Output TokenOutput `json:"output,omitzero"`
	paramObj
}

func (r OutputWithPreviousTransactionData) MarshalJSON() (data []byte, err error) {
	type shadow OutputWithPreviousTransactionData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutputWithPreviousTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The fee for a Spark Lightning payment.
type SparkLightningFee struct {
	OriginalUnit  string  `json:"original_unit" api:"required"`
	OriginalValue float64 `json:"original_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OriginalUnit  respjson.Field
		OriginalValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkLightningFee) RawJSON() string { return r.JSON.raw }
func (r *SparkLightningFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark Lightning receive request.
type SparkLightningReceiveRequest struct {
	ID                        string `json:"id" api:"required"`
	CreatedAt                 string `json:"created_at" api:"required"`
	Network                   string `json:"network" api:"required"`
	Status                    string `json:"status" api:"required"`
	Typename                  string `json:"typename" api:"required"`
	UpdatedAt                 string `json:"updated_at" api:"required"`
	Invoice                   any    `json:"invoice"`
	PaymentPreimage           string `json:"payment_preimage"`
	ReceiverIdentityPublicKey string `json:"receiver_identity_public_key"`
	Transfer                  any    `json:"transfer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CreatedAt                 respjson.Field
		Network                   respjson.Field
		Status                    respjson.Field
		Typename                  respjson.Field
		UpdatedAt                 respjson.Field
		Invoice                   respjson.Field
		PaymentPreimage           respjson.Field
		ReceiverIdentityPublicKey respjson.Field
		Transfer                  respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkLightningReceiveRequest) RawJSON() string { return r.JSON.raw }
func (r *SparkLightningReceiveRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Spark Lightning send request.
type SparkLightningSendRequest struct {
	ID             string `json:"id" api:"required"`
	CreatedAt      string `json:"created_at" api:"required"`
	EncodedInvoice string `json:"encoded_invoice" api:"required"`
	// The fee for a Spark Lightning payment.
	Fee             SparkLightningFee `json:"fee" api:"required"`
	IdempotencyKey  string            `json:"idempotency_key" api:"required"`
	Network         string            `json:"network" api:"required"`
	Status          string            `json:"status" api:"required"`
	Typename        string            `json:"typename" api:"required"`
	UpdatedAt       string            `json:"updated_at" api:"required"`
	PaymentPreimage string            `json:"payment_preimage"`
	Transfer        any               `json:"transfer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		EncodedInvoice  respjson.Field
		Fee             respjson.Field
		IdempotencyKey  respjson.Field
		Network         respjson.Field
		Status          respjson.Field
		Typename        respjson.Field
		UpdatedAt       respjson.Field
		PaymentPreimage respjson.Field
		Transfer        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkLightningSendRequest) RawJSON() string { return r.JSON.raw }
func (r *SparkLightningSendRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the Spark `transfer` RPC.
type SparkTransferRpcInputParamsResp struct {
	AmountSats           float64 `json:"amount_sats" api:"required"`
	ReceiverSparkAddress string  `json:"receiver_spark_address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountSats           respjson.Field
		ReceiverSparkAddress respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkTransferRpcInputParamsResp to a
// SparkTransferRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkTransferRpcInputParams.Overrides()
func (r SparkTransferRpcInputParamsResp) ToParam() SparkTransferRpcInputParams {
	return param.Override[SparkTransferRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `transfer` RPC.
//
// The properties AmountSats, ReceiverSparkAddress are required.
type SparkTransferRpcInputParams struct {
	AmountSats           float64 `json:"amount_sats" api:"required"`
	ReceiverSparkAddress string  `json:"receiver_spark_address" api:"required"`
	paramObj
}

func (r SparkTransferRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transfers satoshis to a Spark address.
type SparkTransferRpcInputResp struct {
	// Any of "transfer".
	Method SparkTransferRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `transfer` RPC.
	Params SparkTransferRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkTransferRpcInputResp to a SparkTransferRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkTransferRpcInput.Overrides()
func (r SparkTransferRpcInputResp) ToParam() SparkTransferRpcInput {
	return param.Override[SparkTransferRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkTransferRpcInputMethod string

const (
	SparkTransferRpcInputMethodTransfer SparkTransferRpcInputMethod = "transfer"
)

// Transfers satoshis to a Spark address.
//
// The properties Method, Params are required.
type SparkTransferRpcInput struct {
	// Any of "transfer".
	Method SparkTransferRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `transfer` RPC.
	Params SparkTransferRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkTransferRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gets the balance of the Spark wallet.
type SparkGetBalanceRpcInputResp struct {
	// Any of "getBalance".
	Method SparkGetBalanceRpcInputMethod `json:"method" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetBalanceRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkGetBalanceRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetBalanceRpcInputResp to a SparkGetBalanceRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetBalanceRpcInput.Overrides()
func (r SparkGetBalanceRpcInputResp) ToParam() SparkGetBalanceRpcInput {
	return param.Override[SparkGetBalanceRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkGetBalanceRpcInputMethod string

const (
	SparkGetBalanceRpcInputMethodGetBalance SparkGetBalanceRpcInputMethod = "getBalance"
)

// Gets the balance of the Spark wallet.
//
// The property Method is required.
type SparkGetBalanceRpcInput struct {
	// Any of "getBalance".
	Method SparkGetBalanceRpcInputMethod `json:"method,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkGetBalanceRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetBalanceRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetBalanceRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strategy for selecting outputs in a Spark token transfer.
type SparkOutputSelectionStrategy string

const (
	SparkOutputSelectionStrategySmallFirst SparkOutputSelectionStrategy = "SMALL_FIRST"
	SparkOutputSelectionStrategyLargeFirst SparkOutputSelectionStrategy = "LARGE_FIRST"
)

// Parameters for the Spark `transferTokens` RPC.
type SparkTransferTokensRpcInputParamsResp struct {
	ReceiverSparkAddress string  `json:"receiver_spark_address" api:"required"`
	TokenAmount          float64 `json:"token_amount" api:"required"`
	TokenIdentifier      string  `json:"token_identifier" api:"required"`
	// Strategy for selecting outputs in a Spark token transfer.
	//
	// Any of "SMALL_FIRST", "LARGE_FIRST".
	OutputSelectionStrategy SparkOutputSelectionStrategy            `json:"output_selection_strategy"`
	SelectedOutputs         []OutputWithPreviousTransactionDataResp `json:"selected_outputs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReceiverSparkAddress    respjson.Field
		TokenAmount             respjson.Field
		TokenIdentifier         respjson.Field
		OutputSelectionStrategy respjson.Field
		SelectedOutputs         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferTokensRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferTokensRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkTransferTokensRpcInputParamsResp to a
// SparkTransferTokensRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkTransferTokensRpcInputParams.Overrides()
func (r SparkTransferTokensRpcInputParamsResp) ToParam() SparkTransferTokensRpcInputParams {
	return param.Override[SparkTransferTokensRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `transferTokens` RPC.
//
// The properties ReceiverSparkAddress, TokenAmount, TokenIdentifier are required.
type SparkTransferTokensRpcInputParams struct {
	ReceiverSparkAddress string  `json:"receiver_spark_address" api:"required"`
	TokenAmount          float64 `json:"token_amount" api:"required"`
	TokenIdentifier      string  `json:"token_identifier" api:"required"`
	// Strategy for selecting outputs in a Spark token transfer.
	//
	// Any of "SMALL_FIRST", "LARGE_FIRST".
	OutputSelectionStrategy SparkOutputSelectionStrategy        `json:"output_selection_strategy,omitzero"`
	SelectedOutputs         []OutputWithPreviousTransactionData `json:"selected_outputs,omitzero"`
	paramObj
}

func (r SparkTransferTokensRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferTokensRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferTokensRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transfers tokens to a Spark address.
type SparkTransferTokensRpcInputResp struct {
	// Any of "transferTokens".
	Method SparkTransferTokensRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `transferTokens` RPC.
	Params SparkTransferTokensRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferTokensRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferTokensRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkTransferTokensRpcInputResp to a
// SparkTransferTokensRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkTransferTokensRpcInput.Overrides()
func (r SparkTransferTokensRpcInputResp) ToParam() SparkTransferTokensRpcInput {
	return param.Override[SparkTransferTokensRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkTransferTokensRpcInputMethod string

const (
	SparkTransferTokensRpcInputMethodTransferTokens SparkTransferTokensRpcInputMethod = "transferTokens"
)

// Transfers tokens to a Spark address.
//
// The properties Method, Params are required.
type SparkTransferTokensRpcInput struct {
	// Any of "transferTokens".
	Method SparkTransferTokensRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `transferTokens` RPC.
	Params SparkTransferTokensRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkTransferTokensRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferTokensRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferTokensRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gets a static deposit address for the Spark wallet.
type SparkGetStaticDepositAddressRpcInputResp struct {
	// Any of "getStaticDepositAddress".
	Method SparkGetStaticDepositAddressRpcInputMethod `json:"method" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetStaticDepositAddressRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkGetStaticDepositAddressRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetStaticDepositAddressRpcInputResp to a
// SparkGetStaticDepositAddressRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetStaticDepositAddressRpcInput.Overrides()
func (r SparkGetStaticDepositAddressRpcInputResp) ToParam() SparkGetStaticDepositAddressRpcInput {
	return param.Override[SparkGetStaticDepositAddressRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkGetStaticDepositAddressRpcInputMethod string

const (
	SparkGetStaticDepositAddressRpcInputMethodGetStaticDepositAddress SparkGetStaticDepositAddressRpcInputMethod = "getStaticDepositAddress"
)

// Gets a static deposit address for the Spark wallet.
//
// The property Method is required.
type SparkGetStaticDepositAddressRpcInput struct {
	// Any of "getStaticDepositAddress".
	Method SparkGetStaticDepositAddressRpcInputMethod `json:"method,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkGetStaticDepositAddressRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetStaticDepositAddressRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetStaticDepositAddressRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the Spark `getClaimStaticDepositQuote` RPC.
type SparkGetClaimStaticDepositQuoteRpcInputParamsResp struct {
	TransactionID string  `json:"transaction_id" api:"required"`
	OutputIndex   float64 `json:"output_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TransactionID respjson.Field
		OutputIndex   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetClaimStaticDepositQuoteRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkGetClaimStaticDepositQuoteRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetClaimStaticDepositQuoteRpcInputParamsResp to a
// SparkGetClaimStaticDepositQuoteRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetClaimStaticDepositQuoteRpcInputParams.Overrides()
func (r SparkGetClaimStaticDepositQuoteRpcInputParamsResp) ToParam() SparkGetClaimStaticDepositQuoteRpcInputParams {
	return param.Override[SparkGetClaimStaticDepositQuoteRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `getClaimStaticDepositQuote` RPC.
//
// The property TransactionID is required.
type SparkGetClaimStaticDepositQuoteRpcInputParams struct {
	TransactionID string             `json:"transaction_id" api:"required"`
	OutputIndex   param.Opt[float64] `json:"output_index,omitzero"`
	paramObj
}

func (r SparkGetClaimStaticDepositQuoteRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetClaimStaticDepositQuoteRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetClaimStaticDepositQuoteRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gets a quote for claiming a static deposit.
type SparkGetClaimStaticDepositQuoteRpcInputResp struct {
	// Any of "getClaimStaticDepositQuote".
	Method SparkGetClaimStaticDepositQuoteRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `getClaimStaticDepositQuote` RPC.
	Params SparkGetClaimStaticDepositQuoteRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetClaimStaticDepositQuoteRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkGetClaimStaticDepositQuoteRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetClaimStaticDepositQuoteRpcInputResp to a
// SparkGetClaimStaticDepositQuoteRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetClaimStaticDepositQuoteRpcInput.Overrides()
func (r SparkGetClaimStaticDepositQuoteRpcInputResp) ToParam() SparkGetClaimStaticDepositQuoteRpcInput {
	return param.Override[SparkGetClaimStaticDepositQuoteRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkGetClaimStaticDepositQuoteRpcInputMethod string

const (
	SparkGetClaimStaticDepositQuoteRpcInputMethodGetClaimStaticDepositQuote SparkGetClaimStaticDepositQuoteRpcInputMethod = "getClaimStaticDepositQuote"
)

// Gets a quote for claiming a static deposit.
//
// The properties Method, Params are required.
type SparkGetClaimStaticDepositQuoteRpcInput struct {
	// Any of "getClaimStaticDepositQuote".
	Method SparkGetClaimStaticDepositQuoteRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `getClaimStaticDepositQuote` RPC.
	Params SparkGetClaimStaticDepositQuoteRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkGetClaimStaticDepositQuoteRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetClaimStaticDepositQuoteRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetClaimStaticDepositQuoteRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the Spark `claimStaticDeposit` RPC.
type SparkClaimStaticDepositRpcInputParamsResp struct {
	CreditAmountSats float64 `json:"credit_amount_sats" api:"required"`
	Signature        string  `json:"signature" api:"required"`
	TransactionID    string  `json:"transaction_id" api:"required"`
	OutputIndex      float64 `json:"output_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditAmountSats respjson.Field
		Signature        respjson.Field
		TransactionID    respjson.Field
		OutputIndex      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkClaimStaticDepositRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkClaimStaticDepositRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkClaimStaticDepositRpcInputParamsResp to a
// SparkClaimStaticDepositRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkClaimStaticDepositRpcInputParams.Overrides()
func (r SparkClaimStaticDepositRpcInputParamsResp) ToParam() SparkClaimStaticDepositRpcInputParams {
	return param.Override[SparkClaimStaticDepositRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `claimStaticDeposit` RPC.
//
// The properties CreditAmountSats, Signature, TransactionID are required.
type SparkClaimStaticDepositRpcInputParams struct {
	CreditAmountSats float64            `json:"credit_amount_sats" api:"required"`
	Signature        string             `json:"signature" api:"required"`
	TransactionID    string             `json:"transaction_id" api:"required"`
	OutputIndex      param.Opt[float64] `json:"output_index,omitzero"`
	paramObj
}

func (r SparkClaimStaticDepositRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkClaimStaticDepositRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkClaimStaticDepositRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Claims a static deposit into the Spark wallet.
type SparkClaimStaticDepositRpcInputResp struct {
	// Any of "claimStaticDeposit".
	Method SparkClaimStaticDepositRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `claimStaticDeposit` RPC.
	Params SparkClaimStaticDepositRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkClaimStaticDepositRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkClaimStaticDepositRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkClaimStaticDepositRpcInputResp to a
// SparkClaimStaticDepositRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkClaimStaticDepositRpcInput.Overrides()
func (r SparkClaimStaticDepositRpcInputResp) ToParam() SparkClaimStaticDepositRpcInput {
	return param.Override[SparkClaimStaticDepositRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkClaimStaticDepositRpcInputMethod string

const (
	SparkClaimStaticDepositRpcInputMethodClaimStaticDeposit SparkClaimStaticDepositRpcInputMethod = "claimStaticDeposit"
)

// Claims a static deposit into the Spark wallet.
//
// The properties Method, Params are required.
type SparkClaimStaticDepositRpcInput struct {
	// Any of "claimStaticDeposit".
	Method SparkClaimStaticDepositRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `claimStaticDeposit` RPC.
	Params SparkClaimStaticDepositRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkClaimStaticDepositRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkClaimStaticDepositRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkClaimStaticDepositRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the Spark `createLightningInvoice` RPC.
type SparkCreateLightningInvoiceRpcInputParamsResp struct {
	AmountSats             float64 `json:"amount_sats" api:"required"`
	DescriptionHash        string  `json:"description_hash"`
	ExpirySeconds          float64 `json:"expiry_seconds"`
	IncludeSparkAddress    bool    `json:"include_spark_address"`
	Memo                   string  `json:"memo"`
	ReceiverIdentityPubkey string  `json:"receiver_identity_pubkey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountSats             respjson.Field
		DescriptionHash        respjson.Field
		ExpirySeconds          respjson.Field
		IncludeSparkAddress    respjson.Field
		Memo                   respjson.Field
		ReceiverIdentityPubkey respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkCreateLightningInvoiceRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkCreateLightningInvoiceRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkCreateLightningInvoiceRpcInputParamsResp to a
// SparkCreateLightningInvoiceRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkCreateLightningInvoiceRpcInputParams.Overrides()
func (r SparkCreateLightningInvoiceRpcInputParamsResp) ToParam() SparkCreateLightningInvoiceRpcInputParams {
	return param.Override[SparkCreateLightningInvoiceRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `createLightningInvoice` RPC.
//
// The property AmountSats is required.
type SparkCreateLightningInvoiceRpcInputParams struct {
	AmountSats             float64            `json:"amount_sats" api:"required"`
	DescriptionHash        param.Opt[string]  `json:"description_hash,omitzero"`
	ExpirySeconds          param.Opt[float64] `json:"expiry_seconds,omitzero"`
	IncludeSparkAddress    param.Opt[bool]    `json:"include_spark_address,omitzero"`
	Memo                   param.Opt[string]  `json:"memo,omitzero"`
	ReceiverIdentityPubkey param.Opt[string]  `json:"receiver_identity_pubkey,omitzero"`
	paramObj
}

func (r SparkCreateLightningInvoiceRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkCreateLightningInvoiceRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkCreateLightningInvoiceRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates a Lightning invoice for the Spark wallet.
type SparkCreateLightningInvoiceRpcInputResp struct {
	// Any of "createLightningInvoice".
	Method SparkCreateLightningInvoiceRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `createLightningInvoice` RPC.
	Params SparkCreateLightningInvoiceRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkCreateLightningInvoiceRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkCreateLightningInvoiceRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkCreateLightningInvoiceRpcInputResp to a
// SparkCreateLightningInvoiceRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkCreateLightningInvoiceRpcInput.Overrides()
func (r SparkCreateLightningInvoiceRpcInputResp) ToParam() SparkCreateLightningInvoiceRpcInput {
	return param.Override[SparkCreateLightningInvoiceRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkCreateLightningInvoiceRpcInputMethod string

const (
	SparkCreateLightningInvoiceRpcInputMethodCreateLightningInvoice SparkCreateLightningInvoiceRpcInputMethod = "createLightningInvoice"
)

// Creates a Lightning invoice for the Spark wallet.
//
// The properties Method, Params are required.
type SparkCreateLightningInvoiceRpcInput struct {
	// Any of "createLightningInvoice".
	Method SparkCreateLightningInvoiceRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `createLightningInvoice` RPC.
	Params SparkCreateLightningInvoiceRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkCreateLightningInvoiceRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkCreateLightningInvoiceRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkCreateLightningInvoiceRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the Spark `payLightningInvoice` RPC.
type SparkPayLightningInvoiceRpcInputParamsResp struct {
	Invoice          string  `json:"invoice" api:"required"`
	MaxFeeSats       float64 `json:"max_fee_sats" api:"required"`
	AmountSatsToSend float64 `json:"amount_sats_to_send"`
	PreferSpark      bool    `json:"prefer_spark"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoice          respjson.Field
		MaxFeeSats       respjson.Field
		AmountSatsToSend respjson.Field
		PreferSpark      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkPayLightningInvoiceRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkPayLightningInvoiceRpcInputParamsResp to a
// SparkPayLightningInvoiceRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkPayLightningInvoiceRpcInputParams.Overrides()
func (r SparkPayLightningInvoiceRpcInputParamsResp) ToParam() SparkPayLightningInvoiceRpcInputParams {
	return param.Override[SparkPayLightningInvoiceRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `payLightningInvoice` RPC.
//
// The properties Invoice, MaxFeeSats are required.
type SparkPayLightningInvoiceRpcInputParams struct {
	Invoice          string             `json:"invoice" api:"required"`
	MaxFeeSats       float64            `json:"max_fee_sats" api:"required"`
	AmountSatsToSend param.Opt[float64] `json:"amount_sats_to_send,omitzero"`
	PreferSpark      param.Opt[bool]    `json:"prefer_spark,omitzero"`
	paramObj
}

func (r SparkPayLightningInvoiceRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkPayLightningInvoiceRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkPayLightningInvoiceRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pays a Lightning invoice from the Spark wallet.
type SparkPayLightningInvoiceRpcInputResp struct {
	// Any of "payLightningInvoice".
	Method SparkPayLightningInvoiceRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `payLightningInvoice` RPC.
	Params SparkPayLightningInvoiceRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkPayLightningInvoiceRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkPayLightningInvoiceRpcInputResp to a
// SparkPayLightningInvoiceRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkPayLightningInvoiceRpcInput.Overrides()
func (r SparkPayLightningInvoiceRpcInputResp) ToParam() SparkPayLightningInvoiceRpcInput {
	return param.Override[SparkPayLightningInvoiceRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkPayLightningInvoiceRpcInputMethod string

const (
	SparkPayLightningInvoiceRpcInputMethodPayLightningInvoice SparkPayLightningInvoiceRpcInputMethod = "payLightningInvoice"
)

// Pays a Lightning invoice from the Spark wallet.
//
// The properties Method, Params are required.
type SparkPayLightningInvoiceRpcInput struct {
	// Any of "payLightningInvoice".
	Method SparkPayLightningInvoiceRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `payLightningInvoice` RPC.
	Params SparkPayLightningInvoiceRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkPayLightningInvoiceRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkPayLightningInvoiceRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkPayLightningInvoiceRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the Spark `signMessageWithIdentityKey` RPC.
type SparkSignMessageWithIdentityKeyRpcInputParamsResp struct {
	Message string `json:"message" api:"required"`
	Compact bool   `json:"compact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Compact     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkSignMessageWithIdentityKeyRpcInputParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SparkSignMessageWithIdentityKeyRpcInputParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkSignMessageWithIdentityKeyRpcInputParamsResp to a
// SparkSignMessageWithIdentityKeyRpcInputParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkSignMessageWithIdentityKeyRpcInputParams.Overrides()
func (r SparkSignMessageWithIdentityKeyRpcInputParamsResp) ToParam() SparkSignMessageWithIdentityKeyRpcInputParams {
	return param.Override[SparkSignMessageWithIdentityKeyRpcInputParams](json.RawMessage(r.RawJSON()))
}

// Parameters for the Spark `signMessageWithIdentityKey` RPC.
//
// The property Message is required.
type SparkSignMessageWithIdentityKeyRpcInputParams struct {
	Message string          `json:"message" api:"required"`
	Compact param.Opt[bool] `json:"compact,omitzero"`
	paramObj
}

func (r SparkSignMessageWithIdentityKeyRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkSignMessageWithIdentityKeyRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkSignMessageWithIdentityKeyRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signs a message with the Spark identity key.
type SparkSignMessageWithIdentityKeyRpcInputResp struct {
	// Any of "signMessageWithIdentityKey".
	Method SparkSignMessageWithIdentityKeyRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `signMessageWithIdentityKey` RPC.
	Params SparkSignMessageWithIdentityKeyRpcInputParamsResp `json:"params" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkSignMessageWithIdentityKeyRpcInputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkSignMessageWithIdentityKeyRpcInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkSignMessageWithIdentityKeyRpcInputResp to a
// SparkSignMessageWithIdentityKeyRpcInput.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkSignMessageWithIdentityKeyRpcInput.Overrides()
func (r SparkSignMessageWithIdentityKeyRpcInputResp) ToParam() SparkSignMessageWithIdentityKeyRpcInput {
	return param.Override[SparkSignMessageWithIdentityKeyRpcInput](json.RawMessage(r.RawJSON()))
}

type SparkSignMessageWithIdentityKeyRpcInputMethod string

const (
	SparkSignMessageWithIdentityKeyRpcInputMethodSignMessageWithIdentityKey SparkSignMessageWithIdentityKeyRpcInputMethod = "signMessageWithIdentityKey"
)

// Signs a message with the Spark identity key.
//
// The properties Method, Params are required.
type SparkSignMessageWithIdentityKeyRpcInput struct {
	// Any of "signMessageWithIdentityKey".
	Method SparkSignMessageWithIdentityKeyRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `signMessageWithIdentityKey` RPC.
	Params SparkSignMessageWithIdentityKeyRpcInputParams `json:"params,omitzero" api:"required"`
	// The Spark network.
	//
	// Any of "MAINNET", "REGTEST".
	Network SparkNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkSignMessageWithIdentityKeyRpcInput) MarshalJSON() (data []byte, err error) {
	type shadow SparkSignMessageWithIdentityKeyRpcInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkSignMessageWithIdentityKeyRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `transfer` RPC.
type SparkTransferRpcResponse struct {
	// Any of "transfer".
	Method SparkTransferRpcResponseMethod `json:"method" api:"required"`
	// A Spark transfer.
	Data SparkTransfer `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkTransferRpcResponseMethod string

const (
	SparkTransferRpcResponseMethodTransfer SparkTransferRpcResponseMethod = "transfer"
)

// Response to the Spark `getBalance` RPC.
type SparkGetBalanceRpcResponse struct {
	// Any of "getBalance".
	Method SparkGetBalanceRpcResponseMethod `json:"method" api:"required"`
	// The balance of a Spark wallet.
	Data SparkBalance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetBalanceRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkGetBalanceRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkGetBalanceRpcResponseMethod string

const (
	SparkGetBalanceRpcResponseMethodGetBalance SparkGetBalanceRpcResponseMethod = "getBalance"
)

// Data returned by the Spark `transferTokens` RPC.
type SparkTransferTokensRpcResponseData struct {
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferTokensRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferTokensRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `transferTokens` RPC.
type SparkTransferTokensRpcResponse struct {
	// Any of "transferTokens".
	Method SparkTransferTokensRpcResponseMethod `json:"method" api:"required"`
	// Data returned by the Spark `transferTokens` RPC.
	Data SparkTransferTokensRpcResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkTransferTokensRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferTokensRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkTransferTokensRpcResponseMethod string

const (
	SparkTransferTokensRpcResponseMethodTransferTokens SparkTransferTokensRpcResponseMethod = "transferTokens"
)

// Data returned by the Spark `getStaticDepositAddress` RPC.
type SparkGetStaticDepositAddressRpcResponseData struct {
	Address string `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetStaticDepositAddressRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkGetStaticDepositAddressRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `getStaticDepositAddress` RPC.
type SparkGetStaticDepositAddressRpcResponse struct {
	// Any of "getStaticDepositAddress".
	Method SparkGetStaticDepositAddressRpcResponseMethod `json:"method" api:"required"`
	// Data returned by the Spark `getStaticDepositAddress` RPC.
	Data SparkGetStaticDepositAddressRpcResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetStaticDepositAddressRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkGetStaticDepositAddressRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkGetStaticDepositAddressRpcResponseMethod string

const (
	SparkGetStaticDepositAddressRpcResponseMethodGetStaticDepositAddress SparkGetStaticDepositAddressRpcResponseMethod = "getStaticDepositAddress"
)

// Data returned by the Spark `getClaimStaticDepositQuote` RPC.
type SparkGetClaimStaticDepositQuoteRpcResponseData struct {
	CreditAmountSats float64 `json:"credit_amount_sats" api:"required"`
	Network          string  `json:"network" api:"required"`
	OutputIndex      float64 `json:"output_index" api:"required"`
	Signature        string  `json:"signature" api:"required"`
	TransactionID    string  `json:"transaction_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditAmountSats respjson.Field
		Network          respjson.Field
		OutputIndex      respjson.Field
		Signature        respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetClaimStaticDepositQuoteRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkGetClaimStaticDepositQuoteRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `getClaimStaticDepositQuote` RPC.
type SparkGetClaimStaticDepositQuoteRpcResponse struct {
	// Any of "getClaimStaticDepositQuote".
	Method SparkGetClaimStaticDepositQuoteRpcResponseMethod `json:"method" api:"required"`
	// Data returned by the Spark `getClaimStaticDepositQuote` RPC.
	Data SparkGetClaimStaticDepositQuoteRpcResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetClaimStaticDepositQuoteRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkGetClaimStaticDepositQuoteRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkGetClaimStaticDepositQuoteRpcResponseMethod string

const (
	SparkGetClaimStaticDepositQuoteRpcResponseMethodGetClaimStaticDepositQuote SparkGetClaimStaticDepositQuoteRpcResponseMethod = "getClaimStaticDepositQuote"
)

// Data returned by the Spark `claimStaticDeposit` RPC.
type SparkClaimStaticDepositRpcResponseData struct {
	TransferID string `json:"transfer_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TransferID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkClaimStaticDepositRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkClaimStaticDepositRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `claimStaticDeposit` RPC.
type SparkClaimStaticDepositRpcResponse struct {
	// Any of "claimStaticDeposit".
	Method SparkClaimStaticDepositRpcResponseMethod `json:"method" api:"required"`
	// Data returned by the Spark `claimStaticDeposit` RPC.
	Data SparkClaimStaticDepositRpcResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkClaimStaticDepositRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkClaimStaticDepositRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkClaimStaticDepositRpcResponseMethod string

const (
	SparkClaimStaticDepositRpcResponseMethodClaimStaticDeposit SparkClaimStaticDepositRpcResponseMethod = "claimStaticDeposit"
)

// Response to the Spark `createLightningInvoice` RPC.
type SparkCreateLightningInvoiceRpcResponse struct {
	// Any of "createLightningInvoice".
	Method SparkCreateLightningInvoiceRpcResponseMethod `json:"method" api:"required"`
	// A Spark Lightning receive request.
	Data SparkLightningReceiveRequest `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkCreateLightningInvoiceRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkCreateLightningInvoiceRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkCreateLightningInvoiceRpcResponseMethod string

const (
	SparkCreateLightningInvoiceRpcResponseMethodCreateLightningInvoice SparkCreateLightningInvoiceRpcResponseMethod = "createLightningInvoice"
)

// Response to the Spark `payLightningInvoice` RPC.
type SparkPayLightningInvoiceRpcResponse struct {
	// Any of "payLightningInvoice".
	Method SparkPayLightningInvoiceRpcResponseMethod `json:"method" api:"required"`
	// A Spark transfer.
	Data SparkPayLightningInvoiceRpcResponseDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkPayLightningInvoiceRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkPayLightningInvoiceRpcResponseMethod string

const (
	SparkPayLightningInvoiceRpcResponseMethodPayLightningInvoice SparkPayLightningInvoiceRpcResponseMethod = "payLightningInvoice"
)

// SparkPayLightningInvoiceRpcResponseDataUnion contains all possible properties
// and values from [SparkTransfer], [SparkLightningSendRequest].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SparkPayLightningInvoiceRpcResponseDataUnion struct {
	ID string `json:"id"`
	// This field is from variant [SparkTransfer].
	Leaves []SparkTransferLeaf `json:"leaves"`
	// This field is from variant [SparkTransfer].
	ReceiverIdentityPublicKey string `json:"receiver_identity_public_key"`
	// This field is from variant [SparkTransfer].
	SenderIdentityPublicKey string `json:"sender_identity_public_key"`
	Status                  string `json:"status"`
	// This field is from variant [SparkTransfer].
	TotalValue float64 `json:"total_value"`
	// This field is from variant [SparkTransfer].
	TransferDirection string `json:"transfer_direction"`
	// This field is from variant [SparkTransfer].
	Type string `json:"type"`
	// This field is from variant [SparkTransfer].
	CreatedTime string `json:"created_time"`
	// This field is from variant [SparkTransfer].
	ExpiryTime string `json:"expiry_time"`
	// This field is from variant [SparkTransfer].
	UpdatedTime string `json:"updated_time"`
	// This field is from variant [SparkLightningSendRequest].
	CreatedAt string `json:"created_at"`
	// This field is from variant [SparkLightningSendRequest].
	EncodedInvoice string `json:"encoded_invoice"`
	// This field is from variant [SparkLightningSendRequest].
	Fee SparkLightningFee `json:"fee"`
	// This field is from variant [SparkLightningSendRequest].
	IdempotencyKey string `json:"idempotency_key"`
	// This field is from variant [SparkLightningSendRequest].
	Network string `json:"network"`
	// This field is from variant [SparkLightningSendRequest].
	Typename string `json:"typename"`
	// This field is from variant [SparkLightningSendRequest].
	UpdatedAt string `json:"updated_at"`
	// This field is from variant [SparkLightningSendRequest].
	PaymentPreimage string `json:"payment_preimage"`
	// This field is from variant [SparkLightningSendRequest].
	Transfer any `json:"transfer"`
	JSON     struct {
		ID                        respjson.Field
		Leaves                    respjson.Field
		ReceiverIdentityPublicKey respjson.Field
		SenderIdentityPublicKey   respjson.Field
		Status                    respjson.Field
		TotalValue                respjson.Field
		TransferDirection         respjson.Field
		Type                      respjson.Field
		CreatedTime               respjson.Field
		ExpiryTime                respjson.Field
		UpdatedTime               respjson.Field
		CreatedAt                 respjson.Field
		EncodedInvoice            respjson.Field
		Fee                       respjson.Field
		IdempotencyKey            respjson.Field
		Network                   respjson.Field
		Typename                  respjson.Field
		UpdatedAt                 respjson.Field
		PaymentPreimage           respjson.Field
		Transfer                  respjson.Field
		raw                       string
	} `json:"-"`
}

func (u SparkPayLightningInvoiceRpcResponseDataUnion) AsSparkTransfer() (v SparkTransfer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SparkPayLightningInvoiceRpcResponseDataUnion) AsSparkLightningSendRequest() (v SparkLightningSendRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SparkPayLightningInvoiceRpcResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *SparkPayLightningInvoiceRpcResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data returned by the Spark `signMessageWithIdentityKey` RPC.
type SparkSignMessageWithIdentityKeyRpcResponseData struct {
	Signature string `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkSignMessageWithIdentityKeyRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkSignMessageWithIdentityKeyRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `signMessageWithIdentityKey` RPC.
type SparkSignMessageWithIdentityKeyRpcResponse struct {
	// Any of "signMessageWithIdentityKey".
	Method SparkSignMessageWithIdentityKeyRpcResponseMethod `json:"method" api:"required"`
	// Data returned by the Spark `signMessageWithIdentityKey` RPC.
	Data SparkSignMessageWithIdentityKeyRpcResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkSignMessageWithIdentityKeyRpcResponse) RawJSON() string { return r.JSON.raw }
func (r *SparkSignMessageWithIdentityKeyRpcResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkSignMessageWithIdentityKeyRpcResponseMethod string

const (
	SparkSignMessageWithIdentityKeyRpcResponseMethodSignMessageWithIdentityKey SparkSignMessageWithIdentityKeyRpcResponseMethod = "signMessageWithIdentityKey"
)

// A wallet managed by Privy's wallet infrastructure.
type Wallet struct {
	// Unique ID of the wallet. This will be the primary identifier when using the
	// wallet in the future.
	ID string `json:"id" api:"required"`
	// Additional signers for the wallet.
	AdditionalSigners WalletAdditionalSigner `json:"additional_signers" api:"required"`
	// Address of the wallet.
	Address string `json:"address" api:"required"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "bitcoin-taproot", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type" api:"required"`
	// Unix timestamp of when the wallet was created in milliseconds.
	CreatedAt float64 `json:"created_at" api:"required"`
	// Unix timestamp of when the wallet was exported in milliseconds, if the wallet
	// was exported.
	ExportedAt float64 `json:"exported_at" api:"required"`
	// Unix timestamp of when the wallet was imported in milliseconds, if the wallet
	// was imported.
	ImportedAt float64 `json:"imported_at" api:"required"`
	// The key quorum ID of the owner of the wallet.
	OwnerID string `json:"owner_id" api:"required" format:"cuid2"`
	// List of policy IDs for policies that are enforced on the wallet.
	PolicyIDs []string `json:"policy_ids" api:"required"`
	// The number of keys that must sign for an action to be valid.
	AuthorizationThreshold float64 `json:"authorization_threshold"`
	// Information about the custodian managing this wallet.
	Custody WalletCustodian `json:"custody"`
	// A human-readable label for the wallet.
	DisplayName string `json:"display_name"`
	// A customer-provided identifier for mapping to external systems. Write-once, set
	// only at creation.
	ExternalID string `json:"external_id"`
	// The compressed, raw public key for the wallet along the chain cryptographic
	// curve.
	PublicKey string `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AdditionalSigners      respjson.Field
		Address                respjson.Field
		ChainType              respjson.Field
		CreatedAt              respjson.Field
		ExportedAt             respjson.Field
		ImportedAt             respjson.Field
		OwnerID                respjson.Field
		PolicyIDs              respjson.Field
		AuthorizationThreshold respjson.Field
		Custody                respjson.Field
		DisplayName            respjson.Field
		ExternalID             respjson.Field
		PublicKey              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Wallet) RawJSON() string { return r.JSON.raw }
func (r *Wallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for looking up a wallet by its blockchain address.
//
// The property Address is required.
type GetByWalletAddressRequestBody struct {
	// A blockchain wallet address (Ethereum or Solana).
	Address Address `json:"address" api:"required"`
	paramObj
}

func (r GetByWalletAddressRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow GetByWalletAddressRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetByWalletAddressRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for updating a wallet.
type WalletUpdateRequestBody struct {
	// A human-readable label for the wallet. Set to null to clear.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners AdditionalSignerInput `json:"additional_signers,omitzero"`
	// New policy IDs to enforce on the wallet. Currently, only one policy is supported
	// per wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r WalletUpdateRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for a single wallet in a batch creation request.
//
// The property ChainType is required.
type WalletBatchItemInput struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "bitcoin-taproot", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero" api:"required"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// A human-readable label for the wallet.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// A customer-provided identifier for mapping to external systems. URL-safe
	// characters only ([a-zA-Z0-9_-]), max 64 chars. Write-once: cannot be changed
	// after creation.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners AdditionalSignerInput `json:"additional_signers,omitzero"`
	// List of policy IDs for policies that should be enforced on the wallet.
	// Currently, only one policy is supported per wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r WalletBatchItemInput) MarshalJSON() (data []byte, err error) {
	type shadow WalletBatchItemInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletBatchItemInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for batch wallet creation.
//
// The property Wallets is required.
type WalletBatchCreateInput struct {
	// Array of wallet creation requests. Minimum 1, maximum 100.
	Wallets []WalletBatchItemInput `json:"wallets,omitzero" api:"required"`
	paramObj
}

func (r WalletBatchCreateInput) MarshalJSON() (data []byte, err error) {
	type shadow WalletBatchCreateInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletBatchCreateInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletBatchCreateResultUnion contains all possible properties and values from
// [WalletBatchCreateResultWalletBatchCreateSuccess],
// [WalletBatchCreateResultWalletBatchCreateFailure].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletBatchCreateResultUnion struct {
	Index   float64 `json:"index"`
	Success bool    `json:"success"`
	// This field is from variant [WalletBatchCreateResultWalletBatchCreateSuccess].
	Wallet Wallet `json:"wallet"`
	// This field is from variant [WalletBatchCreateResultWalletBatchCreateFailure].
	Code string `json:"code"`
	// This field is from variant [WalletBatchCreateResultWalletBatchCreateFailure].
	Error string `json:"error"`
	JSON  struct {
		Index   respjson.Field
		Success respjson.Field
		Wallet  respjson.Field
		Code    respjson.Field
		Error   respjson.Field
		raw     string
	} `json:"-"`
}

func (u WalletBatchCreateResultUnion) AsWalletBatchCreateSuccess() (v WalletBatchCreateResultWalletBatchCreateSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletBatchCreateResultUnion) AsWalletBatchCreateFailure() (v WalletBatchCreateResultWalletBatchCreateFailure) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletBatchCreateResultUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletBatchCreateResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A successful wallet creation result within a batch operation.
type WalletBatchCreateResultWalletBatchCreateSuccess struct {
	// The index of the wallet in the original request array.
	Index float64 `json:"index" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// A wallet managed by Privy's wallet infrastructure.
	Wallet Wallet `json:"wallet" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		Success     respjson.Field
		Wallet      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletBatchCreateResultWalletBatchCreateSuccess) RawJSON() string { return r.JSON.raw }
func (r *WalletBatchCreateResultWalletBatchCreateSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A failed wallet creation result within a batch operation.
type WalletBatchCreateResultWalletBatchCreateFailure struct {
	// A PrivyErrorCode string identifying the error type (e.g., "invalid_data",
	// "resource_conflict").
	Code string `json:"code" api:"required"`
	// A human-readable error message with details about what went wrong.
	Error string `json:"error" api:"required"`
	// The index of the wallet in the original request array.
	Index float64 `json:"index" api:"required"`
	// Any of false.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Error       respjson.Field
		Index       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletBatchCreateResultWalletBatchCreateFailure) RawJSON() string { return r.JSON.raw }
func (r *WalletBatchCreateResultWalletBatchCreateFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a batch wallet creation request.
type WalletBatchCreateResponse struct {
	// Array of results for each wallet creation request, in the same order as input.
	Results []WalletBatchCreateResultUnion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletBatchCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletBatchCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionResp contains all possible properties and values from
// [EthereumSignTransactionRpcInputResp], [EthereumSendTransactionRpcInputResp],
// [EthereumPersonalSignRpcInputResp], [EthereumSignTypedDataRpcInputResp],
// [EthereumSecp256k1SignRpcInputResp],
// [EthereumSign7702AuthorizationRpcInputResp],
// [EthereumSignUserOperationRpcInputResp], [EthereumSendCallsRpcInputResp],
// [SolanaSignTransactionRpcInputResp], [SolanaSignAndSendTransactionRpcInputResp],
// [SolanaSignMessageRpcInputResp], [SparkTransferRpcInputResp],
// [SparkGetBalanceRpcInputResp], [SparkTransferTokensRpcInputResp],
// [SparkGetStaticDepositAddressRpcInputResp],
// [SparkGetClaimStaticDepositQuoteRpcInputResp],
// [SparkClaimStaticDepositRpcInputResp],
// [SparkCreateLightningInvoiceRpcInputResp],
// [SparkPayLightningInvoiceRpcInputResp],
// [SparkSignMessageWithIdentityKeyRpcInputResp], [ExportPrivateKeyRpcInputResp],
// [ExportSeedPhraseRpcInputResp].
//
// Use the [WalletRpcRequestBodyUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletRpcRequestBodyUnionResp struct {
	// Any of "eth_signTransaction", "eth_sendTransaction", "personal_sign",
	// "eth_signTypedData_v4", "secp256k1_sign", "eth_sign7702Authorization",
	// "eth_signUserOperation", "wallet_sendCalls", "signTransaction",
	// "signAndSendTransaction", "signMessage", "transfer", "getBalance",
	// "transferTokens", "getStaticDepositAddress", "getClaimStaticDepositQuote",
	// "claimStaticDeposit", "createLightningInvoice", "payLightningInvoice",
	// "signMessageWithIdentityKey", "exportPrivateKey", "exportSeedPhrase".
	Method string `json:"method"`
	// This field is a union of [EthereumSignTransactionRpcInputParamsResp],
	// [EthereumSendTransactionRpcInputParamsResp],
	// [EthereumPersonalSignRpcInputParamsResp],
	// [EthereumSignTypedDataRpcInputParamsResp],
	// [EthereumSecp256k1SignRpcInputParamsResp],
	// [EthereumSign7702AuthorizationRpcInputParamsResp],
	// [EthereumSignUserOperationRpcInputParamsResp],
	// [EthereumSendCallsRpcInputParamsResp],
	// [SolanaSignTransactionRpcInputParamsResp],
	// [SolanaSignAndSendTransactionRpcInputParamsResp],
	// [SolanaSignMessageRpcInputParamsResp], [SparkTransferRpcInputParamsResp],
	// [SparkTransferTokensRpcInputParamsResp],
	// [SparkGetClaimStaticDepositQuoteRpcInputParamsResp],
	// [SparkClaimStaticDepositRpcInputParamsResp],
	// [SparkCreateLightningInvoiceRpcInputParamsResp],
	// [SparkPayLightningInvoiceRpcInputParamsResp],
	// [SparkSignMessageWithIdentityKeyRpcInputParamsResp],
	// [PrivateKeyExportInputResp], [SeedPhraseExportInputResp]
	Params    WalletRpcRequestBodyUnionRespParams `json:"params"`
	Address   string                              `json:"address"`
	ChainType string                              `json:"chain_type"`
	WalletID  string                              `json:"wallet_id"`
	// This field is from variant [EthereumSendTransactionRpcInputResp].
	Caip2 Caip2 `json:"caip2"`
	// This field is from variant [EthereumSendTransactionRpcInputResp].
	ExperimentalDataSuffix Hex    `json:"experimental_data_suffix"`
	ReferenceID            string `json:"reference_id"`
	Sponsor                bool   `json:"sponsor"`
	// This field is from variant [SparkTransferRpcInputResp].
	Network SparkNetwork `json:"network"`
	JSON    struct {
		Method                 respjson.Field
		Params                 respjson.Field
		Address                respjson.Field
		ChainType              respjson.Field
		WalletID               respjson.Field
		Caip2                  respjson.Field
		ExperimentalDataSuffix respjson.Field
		ReferenceID            respjson.Field
		Sponsor                respjson.Field
		Network                respjson.Field
		raw                    string
	} `json:"-"`
}

// anyWalletRpcRequestBodyResp is implemented by each variant of
// [WalletRpcRequestBodyUnionResp] to add type safety for the return type of
// [WalletRpcRequestBodyUnionResp.AsAny]
type anyWalletRpcRequestBodyResp interface {
	implWalletRpcRequestBodyUnionResp()
}

func (EthereumSignTransactionRpcInputResp) implWalletRpcRequestBodyUnionResp()         {}
func (EthereumSendTransactionRpcInputResp) implWalletRpcRequestBodyUnionResp()         {}
func (EthereumPersonalSignRpcInputResp) implWalletRpcRequestBodyUnionResp()            {}
func (EthereumSignTypedDataRpcInputResp) implWalletRpcRequestBodyUnionResp()           {}
func (EthereumSecp256k1SignRpcInputResp) implWalletRpcRequestBodyUnionResp()           {}
func (EthereumSign7702AuthorizationRpcInputResp) implWalletRpcRequestBodyUnionResp()   {}
func (EthereumSignUserOperationRpcInputResp) implWalletRpcRequestBodyUnionResp()       {}
func (EthereumSendCallsRpcInputResp) implWalletRpcRequestBodyUnionResp()               {}
func (SolanaSignTransactionRpcInputResp) implWalletRpcRequestBodyUnionResp()           {}
func (SolanaSignAndSendTransactionRpcInputResp) implWalletRpcRequestBodyUnionResp()    {}
func (SolanaSignMessageRpcInputResp) implWalletRpcRequestBodyUnionResp()               {}
func (SparkTransferRpcInputResp) implWalletRpcRequestBodyUnionResp()                   {}
func (SparkGetBalanceRpcInputResp) implWalletRpcRequestBodyUnionResp()                 {}
func (SparkTransferTokensRpcInputResp) implWalletRpcRequestBodyUnionResp()             {}
func (SparkGetStaticDepositAddressRpcInputResp) implWalletRpcRequestBodyUnionResp()    {}
func (SparkGetClaimStaticDepositQuoteRpcInputResp) implWalletRpcRequestBodyUnionResp() {}
func (SparkClaimStaticDepositRpcInputResp) implWalletRpcRequestBodyUnionResp()         {}
func (SparkCreateLightningInvoiceRpcInputResp) implWalletRpcRequestBodyUnionResp()     {}
func (SparkPayLightningInvoiceRpcInputResp) implWalletRpcRequestBodyUnionResp()        {}
func (SparkSignMessageWithIdentityKeyRpcInputResp) implWalletRpcRequestBodyUnionResp() {}
func (ExportPrivateKeyRpcInputResp) implWalletRpcRequestBodyUnionResp()                {}
func (ExportSeedPhraseRpcInputResp) implWalletRpcRequestBodyUnionResp()                {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletRpcRequestBodyUnionResp.AsAny().(type) {
//	case privyclient.EthereumSignTransactionRpcInputResp:
//	case privyclient.EthereumSendTransactionRpcInputResp:
//	case privyclient.EthereumPersonalSignRpcInputResp:
//	case privyclient.EthereumSignTypedDataRpcInputResp:
//	case privyclient.EthereumSecp256k1SignRpcInputResp:
//	case privyclient.EthereumSign7702AuthorizationRpcInputResp:
//	case privyclient.EthereumSignUserOperationRpcInputResp:
//	case privyclient.EthereumSendCallsRpcInputResp:
//	case privyclient.SolanaSignTransactionRpcInputResp:
//	case privyclient.SolanaSignAndSendTransactionRpcInputResp:
//	case privyclient.SolanaSignMessageRpcInputResp:
//	case privyclient.SparkTransferRpcInputResp:
//	case privyclient.SparkGetBalanceRpcInputResp:
//	case privyclient.SparkTransferTokensRpcInputResp:
//	case privyclient.SparkGetStaticDepositAddressRpcInputResp:
//	case privyclient.SparkGetClaimStaticDepositQuoteRpcInputResp:
//	case privyclient.SparkClaimStaticDepositRpcInputResp:
//	case privyclient.SparkCreateLightningInvoiceRpcInputResp:
//	case privyclient.SparkPayLightningInvoiceRpcInputResp:
//	case privyclient.SparkSignMessageWithIdentityKeyRpcInputResp:
//	case privyclient.ExportPrivateKeyRpcInputResp:
//	case privyclient.ExportSeedPhraseRpcInputResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletRpcRequestBodyUnionResp) AsAny() anyWalletRpcRequestBodyResp {
	switch u.Method {
	case "eth_signTransaction":
		return u.AsEthSignTransaction()
	case "eth_sendTransaction":
		return u.AsEthSendTransaction()
	case "personal_sign":
		return u.AsPersonalSign()
	case "eth_signTypedData_v4":
		return u.AsEthSignTypedDataV4()
	case "secp256k1_sign":
		return u.AsSecp256k1Sign()
	case "eth_sign7702Authorization":
		return u.AsEthSign7702Authorization()
	case "eth_signUserOperation":
		return u.AsEthSignUserOperation()
	case "wallet_sendCalls":
		return u.AsWalletSendCalls()
	case "signTransaction":
		return u.AsSignTransaction()
	case "signAndSendTransaction":
		return u.AsSignAndSendTransaction()
	case "signMessage":
		return u.AsSignMessage()
	case "transfer":
		return u.AsTransfer()
	case "getBalance":
		return u.AsGetBalance()
	case "transferTokens":
		return u.AsTransferTokens()
	case "getStaticDepositAddress":
		return u.AsGetStaticDepositAddress()
	case "getClaimStaticDepositQuote":
		return u.AsGetClaimStaticDepositQuote()
	case "claimStaticDeposit":
		return u.AsClaimStaticDeposit()
	case "createLightningInvoice":
		return u.AsCreateLightningInvoice()
	case "payLightningInvoice":
		return u.AsPayLightningInvoice()
	case "signMessageWithIdentityKey":
		return u.AsSignMessageWithIdentityKey()
	case "exportPrivateKey":
		return u.AsExportPrivateKey()
	case "exportSeedPhrase":
		return u.AsExportSeedPhrase()
	}
	return nil
}

func (u WalletRpcRequestBodyUnionResp) AsEthSignTransaction() (v EthereumSignTransactionRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsEthSendTransaction() (v EthereumSendTransactionRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsPersonalSign() (v EthereumPersonalSignRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsEthSignTypedDataV4() (v EthereumSignTypedDataRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsSecp256k1Sign() (v EthereumSecp256k1SignRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsEthSign7702Authorization() (v EthereumSign7702AuthorizationRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsEthSignUserOperation() (v EthereumSignUserOperationRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsWalletSendCalls() (v EthereumSendCallsRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsSignTransaction() (v SolanaSignTransactionRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsSignAndSendTransaction() (v SolanaSignAndSendTransactionRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsSignMessage() (v SolanaSignMessageRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsTransfer() (v SparkTransferRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsGetBalance() (v SparkGetBalanceRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsTransferTokens() (v SparkTransferTokensRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsGetStaticDepositAddress() (v SparkGetStaticDepositAddressRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsGetClaimStaticDepositQuote() (v SparkGetClaimStaticDepositQuoteRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsClaimStaticDeposit() (v SparkClaimStaticDepositRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsCreateLightningInvoice() (v SparkCreateLightningInvoiceRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsPayLightningInvoice() (v SparkPayLightningInvoiceRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsSignMessageWithIdentityKey() (v SparkSignMessageWithIdentityKeyRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsExportPrivateKey() (v ExportPrivateKeyRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnionResp) AsExportSeedPhrase() (v ExportSeedPhraseRpcInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcRequestBodyUnionResp) RawJSON() string { return u.JSON.raw }

func (r *WalletRpcRequestBodyUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionRespParams is an implicit subunion of
// [WalletRpcRequestBodyUnionResp]. WalletRpcRequestBodyUnionRespParams provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnionResp].
type WalletRpcRequestBodyUnionRespParams struct {
	// This field is a union of [UnsignedStandardEthereumTransactionResp], [string],
	// [string]
	Transaction WalletRpcRequestBodyUnionRespParamsTransaction `json:"transaction"`
	Encoding    string                                         `json:"encoding"`
	Message     string                                         `json:"message"`
	// This field is from variant [EthereumSignTypedDataRpcInputParamsResp].
	TypedData EthereumTypedDataInputResp `json:"typed_data"`
	// This field is from variant [EthereumSecp256k1SignRpcInputParamsResp].
	Hash Hex `json:"hash"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	ChainID  QuantityUnionResp `json:"chain_id"`
	Contract string            `json:"contract"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor `json:"executor"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	Nonce QuantityUnionResp `json:"nonce"`
	// This field is from variant [EthereumSignUserOperationRpcInputParamsResp].
	UserOperation UserOperationInputResp `json:"user_operation"`
	// This field is from variant [EthereumSendCallsRpcInputParamsResp].
	Calls                []EthereumSendCallsCallResp `json:"calls"`
	AmountSats           float64                     `json:"amount_sats"`
	ReceiverSparkAddress string                      `json:"receiver_spark_address"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	TokenAmount float64 `json:"token_amount"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	TokenIdentifier string `json:"token_identifier"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	OutputSelectionStrategy SparkOutputSelectionStrategy `json:"output_selection_strategy"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	SelectedOutputs []OutputWithPreviousTransactionDataResp `json:"selected_outputs"`
	TransactionID   string                                  `json:"transaction_id"`
	OutputIndex     float64                                 `json:"output_index"`
	// This field is from variant [SparkClaimStaticDepositRpcInputParamsResp].
	CreditAmountSats float64 `json:"credit_amount_sats"`
	// This field is from variant [SparkClaimStaticDepositRpcInputParamsResp].
	Signature string `json:"signature"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	DescriptionHash string `json:"description_hash"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	ExpirySeconds float64 `json:"expiry_seconds"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	IncludeSparkAddress bool `json:"include_spark_address"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	Memo string `json:"memo"`
	// This field is from variant [SparkCreateLightningInvoiceRpcInputParamsResp].
	ReceiverIdentityPubkey string `json:"receiver_identity_pubkey"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	Invoice string `json:"invoice"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	MaxFeeSats float64 `json:"max_fee_sats"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	AmountSatsToSend float64 `json:"amount_sats_to_send"`
	// This field is from variant [SparkPayLightningInvoiceRpcInputParamsResp].
	PreferSpark bool `json:"prefer_spark"`
	// This field is from variant [SparkSignMessageWithIdentityKeyRpcInputParamsResp].
	Compact bool `json:"compact"`
	// This field is from variant [PrivateKeyExportInputResp].
	EncryptionType HpkeEncryption `json:"encryption_type"`
	// This field is from variant [PrivateKeyExportInputResp].
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key"`
	ExportSeedPhrase   bool               `json:"export_seed_phrase"`
	// This field is from variant [PrivateKeyExportInputResp].
	ExportType ExportType `json:"export_type"`
	JSON       struct {
		Transaction             respjson.Field
		Encoding                respjson.Field
		Message                 respjson.Field
		TypedData               respjson.Field
		Hash                    respjson.Field
		ChainID                 respjson.Field
		Contract                respjson.Field
		Executor                respjson.Field
		Nonce                   respjson.Field
		UserOperation           respjson.Field
		Calls                   respjson.Field
		AmountSats              respjson.Field
		ReceiverSparkAddress    respjson.Field
		TokenAmount             respjson.Field
		TokenIdentifier         respjson.Field
		OutputSelectionStrategy respjson.Field
		SelectedOutputs         respjson.Field
		TransactionID           respjson.Field
		OutputIndex             respjson.Field
		CreditAmountSats        respjson.Field
		Signature               respjson.Field
		DescriptionHash         respjson.Field
		ExpirySeconds           respjson.Field
		IncludeSparkAddress     respjson.Field
		Memo                    respjson.Field
		ReceiverIdentityPubkey  respjson.Field
		Invoice                 respjson.Field
		MaxFeeSats              respjson.Field
		AmountSatsToSend        respjson.Field
		PreferSpark             respjson.Field
		Compact                 respjson.Field
		EncryptionType          respjson.Field
		RecipientPublicKey      respjson.Field
		ExportSeedPhrase        respjson.Field
		ExportType              respjson.Field
		raw                     string
	} `json:"-"`
}

func (r *WalletRpcRequestBodyUnionRespParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionRespParamsTransaction is an implicit subunion of
// [WalletRpcRequestBodyUnionResp]. WalletRpcRequestBodyUnionRespParamsTransaction
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnionResp].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type WalletRpcRequestBodyUnionRespParamsTransaction struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	AuthorizationList []EthereumSign7702AuthorizationResp `json:"authorization_list"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	ChainID QuantityUnionResp `json:"chain_id"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	Data Hex `json:"data"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	From string `json:"from"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	GasLimit QuantityUnionResp `json:"gas_limit"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	GasPrice QuantityUnionResp `json:"gas_price"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	MaxFeePerGas QuantityUnionResp `json:"max_fee_per_gas"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	MaxPriorityFeePerGas QuantityUnionResp `json:"max_priority_fee_per_gas"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	Nonce QuantityUnionResp `json:"nonce"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	To string `json:"to"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	Type float64 `json:"type"`
	// This field is from variant [UnsignedStandardEthereumTransactionResp].
	Value QuantityUnionResp `json:"value"`
	JSON  struct {
		OfString             respjson.Field
		AuthorizationList    respjson.Field
		ChainID              respjson.Field
		Data                 respjson.Field
		From                 respjson.Field
		GasLimit             respjson.Field
		GasPrice             respjson.Field
		MaxFeePerGas         respjson.Field
		MaxPriorityFeePerGas respjson.Field
		Nonce                respjson.Field
		To                   respjson.Field
		Type                 respjson.Field
		Value                respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *WalletRpcRequestBodyUnionRespParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WalletRpcRequestBodyUnionResp to a
// WalletRpcRequestBodyUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WalletRpcRequestBodyUnion.Overrides()
func (r WalletRpcRequestBodyUnionResp) ToParam() WalletRpcRequestBodyUnion {
	return param.Override[WalletRpcRequestBodyUnion](json.RawMessage(r.RawJSON()))
}

func WalletRpcRequestBodyOfEthSignTransaction(params EthereumSignTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSignTransaction EthereumSignTransactionRpcInput
	ethSignTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfEthSignTransaction: &ethSignTransaction}
}

func WalletRpcRequestBodyOfEthSendTransaction(caip2 Caip2, method EthereumSendTransactionRpcInputMethod, params EthereumSendTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSendTransaction EthereumSendTransactionRpcInput
	ethSendTransaction.Caip2 = caip2
	ethSendTransaction.Method = method
	ethSendTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfEthSendTransaction: &ethSendTransaction}
}

func WalletRpcRequestBodyOfPersonalSign(params EthereumPersonalSignRpcInputParams) WalletRpcRequestBodyUnion {
	var personalSign EthereumPersonalSignRpcInput
	personalSign.Params = params
	return WalletRpcRequestBodyUnion{OfPersonalSign: &personalSign}
}

func WalletRpcRequestBodyOfEthSignTypedDataV4(params EthereumSignTypedDataRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSignTypedDataV4 EthereumSignTypedDataRpcInput
	ethSignTypedDataV4.Params = params
	return WalletRpcRequestBodyUnion{OfEthSignTypedDataV4: &ethSignTypedDataV4}
}

func WalletRpcRequestBodyOfSecp256k1Sign(params EthereumSecp256k1SignRpcInputParams) WalletRpcRequestBodyUnion {
	var secp256k1Sign EthereumSecp256k1SignRpcInput
	secp256k1Sign.Params = params
	return WalletRpcRequestBodyUnion{OfSecp256k1Sign: &secp256k1Sign}
}

func WalletRpcRequestBodyOfEthSign7702Authorization(params EthereumSign7702AuthorizationRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSign7702Authorization EthereumSign7702AuthorizationRpcInput
	ethSign7702Authorization.Params = params
	return WalletRpcRequestBodyUnion{OfEthSign7702Authorization: &ethSign7702Authorization}
}

func WalletRpcRequestBodyOfEthSignUserOperation(params EthereumSignUserOperationRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSignUserOperation EthereumSignUserOperationRpcInput
	ethSignUserOperation.Params = params
	return WalletRpcRequestBodyUnion{OfEthSignUserOperation: &ethSignUserOperation}
}

func WalletRpcRequestBodyOfWalletSendCalls(caip2 Caip2, method EthereumSendCallsRpcInputMethod, params EthereumSendCallsRpcInputParams) WalletRpcRequestBodyUnion {
	var walletSendCalls EthereumSendCallsRpcInput
	walletSendCalls.Caip2 = caip2
	walletSendCalls.Method = method
	walletSendCalls.Params = params
	return WalletRpcRequestBodyUnion{OfWalletSendCalls: &walletSendCalls}
}

func WalletRpcRequestBodyOfSignTransaction(params SolanaSignTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var signTransaction SolanaSignTransactionRpcInput
	signTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfSignTransaction: &signTransaction}
}

func WalletRpcRequestBodyOfSignAndSendTransaction(caip2 Caip2, method SolanaSignAndSendTransactionRpcInputMethod, params SolanaSignAndSendTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var signAndSendTransaction SolanaSignAndSendTransactionRpcInput
	signAndSendTransaction.Caip2 = caip2
	signAndSendTransaction.Method = method
	signAndSendTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfSignAndSendTransaction: &signAndSendTransaction}
}

func WalletRpcRequestBodyOfSignMessage(params SolanaSignMessageRpcInputParams) WalletRpcRequestBodyUnion {
	var signMessage SolanaSignMessageRpcInput
	signMessage.Params = params
	return WalletRpcRequestBodyUnion{OfSignMessage: &signMessage}
}

func WalletRpcRequestBodyOfTransfer(params SparkTransferRpcInputParams) WalletRpcRequestBodyUnion {
	var transfer SparkTransferRpcInput
	transfer.Params = params
	return WalletRpcRequestBodyUnion{OfTransfer: &transfer}
}

func WalletRpcRequestBodyOfGetBalance(method SparkGetBalanceRpcInputMethod) WalletRpcRequestBodyUnion {
	var getBalance SparkGetBalanceRpcInput
	getBalance.Method = method
	return WalletRpcRequestBodyUnion{OfGetBalance: &getBalance}
}

func WalletRpcRequestBodyOfTransferTokens(params SparkTransferTokensRpcInputParams) WalletRpcRequestBodyUnion {
	var transferTokens SparkTransferTokensRpcInput
	transferTokens.Params = params
	return WalletRpcRequestBodyUnion{OfTransferTokens: &transferTokens}
}

func WalletRpcRequestBodyOfGetStaticDepositAddress(method SparkGetStaticDepositAddressRpcInputMethod) WalletRpcRequestBodyUnion {
	var getStaticDepositAddress SparkGetStaticDepositAddressRpcInput
	getStaticDepositAddress.Method = method
	return WalletRpcRequestBodyUnion{OfGetStaticDepositAddress: &getStaticDepositAddress}
}

func WalletRpcRequestBodyOfGetClaimStaticDepositQuote(params SparkGetClaimStaticDepositQuoteRpcInputParams) WalletRpcRequestBodyUnion {
	var getClaimStaticDepositQuote SparkGetClaimStaticDepositQuoteRpcInput
	getClaimStaticDepositQuote.Params = params
	return WalletRpcRequestBodyUnion{OfGetClaimStaticDepositQuote: &getClaimStaticDepositQuote}
}

func WalletRpcRequestBodyOfClaimStaticDeposit(params SparkClaimStaticDepositRpcInputParams) WalletRpcRequestBodyUnion {
	var claimStaticDeposit SparkClaimStaticDepositRpcInput
	claimStaticDeposit.Params = params
	return WalletRpcRequestBodyUnion{OfClaimStaticDeposit: &claimStaticDeposit}
}

func WalletRpcRequestBodyOfCreateLightningInvoice(params SparkCreateLightningInvoiceRpcInputParams) WalletRpcRequestBodyUnion {
	var createLightningInvoice SparkCreateLightningInvoiceRpcInput
	createLightningInvoice.Params = params
	return WalletRpcRequestBodyUnion{OfCreateLightningInvoice: &createLightningInvoice}
}

func WalletRpcRequestBodyOfPayLightningInvoice(params SparkPayLightningInvoiceRpcInputParams) WalletRpcRequestBodyUnion {
	var payLightningInvoice SparkPayLightningInvoiceRpcInput
	payLightningInvoice.Params = params
	return WalletRpcRequestBodyUnion{OfPayLightningInvoice: &payLightningInvoice}
}

func WalletRpcRequestBodyOfSignMessageWithIdentityKey(params SparkSignMessageWithIdentityKeyRpcInputParams) WalletRpcRequestBodyUnion {
	var signMessageWithIdentityKey SparkSignMessageWithIdentityKeyRpcInput
	signMessageWithIdentityKey.Params = params
	return WalletRpcRequestBodyUnion{OfSignMessageWithIdentityKey: &signMessageWithIdentityKey}
}

func WalletRpcRequestBodyOfExportPrivateKey(address string, method ExportPrivateKeyRpcInputMethod, params PrivateKeyExportInput) WalletRpcRequestBodyUnion {
	var exportPrivateKey ExportPrivateKeyRpcInput
	exportPrivateKey.Address = address
	exportPrivateKey.Method = method
	exportPrivateKey.Params = params
	return WalletRpcRequestBodyUnion{OfExportPrivateKey: &exportPrivateKey}
}

func WalletRpcRequestBodyOfExportSeedPhrase(address string, method ExportSeedPhraseRpcInputMethod, params SeedPhraseExportInput) WalletRpcRequestBodyUnion {
	var exportSeedPhrase ExportSeedPhraseRpcInput
	exportSeedPhrase.Address = address
	exportSeedPhrase.Method = method
	exportSeedPhrase.Params = params
	return WalletRpcRequestBodyUnion{OfExportSeedPhrase: &exportSeedPhrase}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRpcRequestBodyUnion struct {
	OfEthSignTransaction         *EthereumSignTransactionRpcInput         `json:",omitzero,inline"`
	OfEthSendTransaction         *EthereumSendTransactionRpcInput         `json:",omitzero,inline"`
	OfPersonalSign               *EthereumPersonalSignRpcInput            `json:",omitzero,inline"`
	OfEthSignTypedDataV4         *EthereumSignTypedDataRpcInput           `json:",omitzero,inline"`
	OfSecp256k1Sign              *EthereumSecp256k1SignRpcInput           `json:",omitzero,inline"`
	OfEthSign7702Authorization   *EthereumSign7702AuthorizationRpcInput   `json:",omitzero,inline"`
	OfEthSignUserOperation       *EthereumSignUserOperationRpcInput       `json:",omitzero,inline"`
	OfWalletSendCalls            *EthereumSendCallsRpcInput               `json:",omitzero,inline"`
	OfSignTransaction            *SolanaSignTransactionRpcInput           `json:",omitzero,inline"`
	OfSignAndSendTransaction     *SolanaSignAndSendTransactionRpcInput    `json:",omitzero,inline"`
	OfSignMessage                *SolanaSignMessageRpcInput               `json:",omitzero,inline"`
	OfTransfer                   *SparkTransferRpcInput                   `json:",omitzero,inline"`
	OfGetBalance                 *SparkGetBalanceRpcInput                 `json:",omitzero,inline"`
	OfTransferTokens             *SparkTransferTokensRpcInput             `json:",omitzero,inline"`
	OfGetStaticDepositAddress    *SparkGetStaticDepositAddressRpcInput    `json:",omitzero,inline"`
	OfGetClaimStaticDepositQuote *SparkGetClaimStaticDepositQuoteRpcInput `json:",omitzero,inline"`
	OfClaimStaticDeposit         *SparkClaimStaticDepositRpcInput         `json:",omitzero,inline"`
	OfCreateLightningInvoice     *SparkCreateLightningInvoiceRpcInput     `json:",omitzero,inline"`
	OfPayLightningInvoice        *SparkPayLightningInvoiceRpcInput        `json:",omitzero,inline"`
	OfSignMessageWithIdentityKey *SparkSignMessageWithIdentityKeyRpcInput `json:",omitzero,inline"`
	OfExportPrivateKey           *ExportPrivateKeyRpcInput                `json:",omitzero,inline"`
	OfExportSeedPhrase           *ExportSeedPhraseRpcInput                `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcRequestBodyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEthSignTransaction,
		u.OfEthSendTransaction,
		u.OfPersonalSign,
		u.OfEthSignTypedDataV4,
		u.OfSecp256k1Sign,
		u.OfEthSign7702Authorization,
		u.OfEthSignUserOperation,
		u.OfWalletSendCalls,
		u.OfSignTransaction,
		u.OfSignAndSendTransaction,
		u.OfSignMessage,
		u.OfTransfer,
		u.OfGetBalance,
		u.OfTransferTokens,
		u.OfGetStaticDepositAddress,
		u.OfGetClaimStaticDepositQuote,
		u.OfClaimStaticDeposit,
		u.OfCreateLightningInvoice,
		u.OfPayLightningInvoice,
		u.OfSignMessageWithIdentityKey,
		u.OfExportPrivateKey,
		u.OfExportSeedPhrase)
}
func (u *WalletRpcRequestBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WalletRpcRequestBodyUnion](
		"method",
		apijson.Discriminator[EthereumSignTransactionRpcInput]("eth_signTransaction"),
		apijson.Discriminator[EthereumSendTransactionRpcInput]("eth_sendTransaction"),
		apijson.Discriminator[EthereumPersonalSignRpcInput]("personal_sign"),
		apijson.Discriminator[EthereumSignTypedDataRpcInput]("eth_signTypedData_v4"),
		apijson.Discriminator[EthereumSecp256k1SignRpcInput]("secp256k1_sign"),
		apijson.Discriminator[EthereumSign7702AuthorizationRpcInput]("eth_sign7702Authorization"),
		apijson.Discriminator[EthereumSignUserOperationRpcInput]("eth_signUserOperation"),
		apijson.Discriminator[EthereumSendCallsRpcInput]("wallet_sendCalls"),
		apijson.Discriminator[SolanaSignTransactionRpcInput]("signTransaction"),
		apijson.Discriminator[SolanaSignAndSendTransactionRpcInput]("signAndSendTransaction"),
		apijson.Discriminator[SolanaSignMessageRpcInput]("signMessage"),
		apijson.Discriminator[SparkTransferRpcInput]("transfer"),
		apijson.Discriminator[SparkGetBalanceRpcInput]("getBalance"),
		apijson.Discriminator[SparkTransferTokensRpcInput]("transferTokens"),
		apijson.Discriminator[SparkGetStaticDepositAddressRpcInput]("getStaticDepositAddress"),
		apijson.Discriminator[SparkGetClaimStaticDepositQuoteRpcInput]("getClaimStaticDepositQuote"),
		apijson.Discriminator[SparkClaimStaticDepositRpcInput]("claimStaticDeposit"),
		apijson.Discriminator[SparkCreateLightningInvoiceRpcInput]("createLightningInvoice"),
		apijson.Discriminator[SparkPayLightningInvoiceRpcInput]("payLightningInvoice"),
		apijson.Discriminator[SparkSignMessageWithIdentityKeyRpcInput]("signMessageWithIdentityKey"),
		apijson.Discriminator[ExportPrivateKeyRpcInput]("exportPrivateKey"),
		apijson.Discriminator[ExportSeedPhraseRpcInput]("exportSeedPhrase"),
	)
}

// WalletRpcResponseUnion contains all possible properties and values from
// [EthereumPersonalSignRpcResponse], [EthereumSignTypedDataRpcResponse],
// [EthereumSignTransactionRpcResponse], [EthereumSendTransactionRpcResponse],
// [EthereumSignUserOperationRpcResponse],
// [EthereumSign7702AuthorizationRpcResponse], [EthereumSecp256k1SignRpcResponse],
// [EthereumSendCallsRpcResponse], [SolanaSignMessageRpcResponse],
// [SolanaSignTransactionRpcResponse], [SolanaSignAndSendTransactionRpcResponse],
// [SparkTransferRpcResponse], [SparkGetBalanceRpcResponse],
// [SparkTransferTokensRpcResponse], [SparkGetStaticDepositAddressRpcResponse],
// [SparkGetClaimStaticDepositQuoteRpcResponse],
// [SparkClaimStaticDepositRpcResponse], [SparkCreateLightningInvoiceRpcResponse],
// [SparkPayLightningInvoiceRpcResponse],
// [SparkSignMessageWithIdentityKeyRpcResponse], [ExportPrivateKeyRpcResponse],
// [ExportSeedPhraseRpcResponse].
//
// Use the [WalletRpcResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletRpcResponseUnion struct {
	// This field is a union of [EthereumPersonalSignRpcResponseData],
	// [EthereumSignTypedDataRpcResponseData],
	// [EthereumSignTransactionRpcResponseData],
	// [EthereumSendTransactionRpcResponseData],
	// [EthereumSignUserOperationRpcResponseData],
	// [EthereumSign7702AuthorizationRpcResponseData],
	// [EthereumSecp256k1SignRpcResponseData], [EthereumSendCallsRpcResponseData],
	// [SolanaSignMessageRpcResponseData], [SolanaSignTransactionRpcResponseData],
	// [SolanaSignAndSendTransactionRpcResponseData], [SparkTransfer], [SparkBalance],
	// [SparkTransferTokensRpcResponseData],
	// [SparkGetStaticDepositAddressRpcResponseData],
	// [SparkGetClaimStaticDepositQuoteRpcResponseData],
	// [SparkClaimStaticDepositRpcResponseData], [SparkLightningReceiveRequest],
	// [SparkPayLightningInvoiceRpcResponseDataUnion],
	// [SparkSignMessageWithIdentityKeyRpcResponseData], [PrivateKeyExportInputResp],
	// [SeedPhraseExportResponse]
	Data WalletRpcResponseUnionData `json:"data"`
	// Any of "personal_sign", "eth_signTypedData_v4", "eth_signTransaction",
	// "eth_sendTransaction", "eth_signUserOperation", "eth_sign7702Authorization",
	// "secp256k1_sign", "wallet_sendCalls", "signMessage", "signTransaction",
	// "signAndSendTransaction", "transfer", "getBalance", "transferTokens",
	// "getStaticDepositAddress", "getClaimStaticDepositQuote", "claimStaticDeposit",
	// "createLightningInvoice", "payLightningInvoice", "signMessageWithIdentityKey",
	// "exportPrivateKey", "exportSeedPhrase".
	Method string `json:"method"`
	JSON   struct {
		Data   respjson.Field
		Method respjson.Field
		raw    string
	} `json:"-"`
}

// anyWalletRpcResponse is implemented by each variant of [WalletRpcResponseUnion]
// to add type safety for the return type of [WalletRpcResponseUnion.AsAny]
type anyWalletRpcResponse interface {
	implWalletRpcResponseUnion()
}

func (EthereumPersonalSignRpcResponse) implWalletRpcResponseUnion()            {}
func (EthereumSignTypedDataRpcResponse) implWalletRpcResponseUnion()           {}
func (EthereumSignTransactionRpcResponse) implWalletRpcResponseUnion()         {}
func (EthereumSendTransactionRpcResponse) implWalletRpcResponseUnion()         {}
func (EthereumSignUserOperationRpcResponse) implWalletRpcResponseUnion()       {}
func (EthereumSign7702AuthorizationRpcResponse) implWalletRpcResponseUnion()   {}
func (EthereumSecp256k1SignRpcResponse) implWalletRpcResponseUnion()           {}
func (EthereumSendCallsRpcResponse) implWalletRpcResponseUnion()               {}
func (SolanaSignMessageRpcResponse) implWalletRpcResponseUnion()               {}
func (SolanaSignTransactionRpcResponse) implWalletRpcResponseUnion()           {}
func (SolanaSignAndSendTransactionRpcResponse) implWalletRpcResponseUnion()    {}
func (SparkTransferRpcResponse) implWalletRpcResponseUnion()                   {}
func (SparkGetBalanceRpcResponse) implWalletRpcResponseUnion()                 {}
func (SparkTransferTokensRpcResponse) implWalletRpcResponseUnion()             {}
func (SparkGetStaticDepositAddressRpcResponse) implWalletRpcResponseUnion()    {}
func (SparkGetClaimStaticDepositQuoteRpcResponse) implWalletRpcResponseUnion() {}
func (SparkClaimStaticDepositRpcResponse) implWalletRpcResponseUnion()         {}
func (SparkCreateLightningInvoiceRpcResponse) implWalletRpcResponseUnion()     {}
func (SparkPayLightningInvoiceRpcResponse) implWalletRpcResponseUnion()        {}
func (SparkSignMessageWithIdentityKeyRpcResponse) implWalletRpcResponseUnion() {}
func (ExportPrivateKeyRpcResponse) implWalletRpcResponseUnion()                {}
func (ExportSeedPhraseRpcResponse) implWalletRpcResponseUnion()                {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletRpcResponseUnion.AsAny().(type) {
//	case privyclient.EthereumPersonalSignRpcResponse:
//	case privyclient.EthereumSignTypedDataRpcResponse:
//	case privyclient.EthereumSignTransactionRpcResponse:
//	case privyclient.EthereumSendTransactionRpcResponse:
//	case privyclient.EthereumSignUserOperationRpcResponse:
//	case privyclient.EthereumSign7702AuthorizationRpcResponse:
//	case privyclient.EthereumSecp256k1SignRpcResponse:
//	case privyclient.EthereumSendCallsRpcResponse:
//	case privyclient.SolanaSignMessageRpcResponse:
//	case privyclient.SolanaSignTransactionRpcResponse:
//	case privyclient.SolanaSignAndSendTransactionRpcResponse:
//	case privyclient.SparkTransferRpcResponse:
//	case privyclient.SparkGetBalanceRpcResponse:
//	case privyclient.SparkTransferTokensRpcResponse:
//	case privyclient.SparkGetStaticDepositAddressRpcResponse:
//	case privyclient.SparkGetClaimStaticDepositQuoteRpcResponse:
//	case privyclient.SparkClaimStaticDepositRpcResponse:
//	case privyclient.SparkCreateLightningInvoiceRpcResponse:
//	case privyclient.SparkPayLightningInvoiceRpcResponse:
//	case privyclient.SparkSignMessageWithIdentityKeyRpcResponse:
//	case privyclient.ExportPrivateKeyRpcResponse:
//	case privyclient.ExportSeedPhraseRpcResponse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletRpcResponseUnion) AsAny() anyWalletRpcResponse {
	switch u.Method {
	case "personal_sign":
		return u.AsPersonalSign()
	case "eth_signTypedData_v4":
		return u.AsEthSignTypedDataV4()
	case "eth_signTransaction":
		return u.AsEthSignTransaction()
	case "eth_sendTransaction":
		return u.AsEthSendTransaction()
	case "eth_signUserOperation":
		return u.AsEthSignUserOperation()
	case "eth_sign7702Authorization":
		return u.AsEthSign7702Authorization()
	case "secp256k1_sign":
		return u.AsSecp256k1Sign()
	case "wallet_sendCalls":
		return u.AsWalletSendCalls()
	case "signMessage":
		return u.AsSignMessage()
	case "signTransaction":
		return u.AsSignTransaction()
	case "signAndSendTransaction":
		return u.AsSignAndSendTransaction()
	case "transfer":
		return u.AsTransfer()
	case "getBalance":
		return u.AsGetBalance()
	case "transferTokens":
		return u.AsTransferTokens()
	case "getStaticDepositAddress":
		return u.AsGetStaticDepositAddress()
	case "getClaimStaticDepositQuote":
		return u.AsGetClaimStaticDepositQuote()
	case "claimStaticDeposit":
		return u.AsClaimStaticDeposit()
	case "createLightningInvoice":
		return u.AsCreateLightningInvoice()
	case "payLightningInvoice":
		return u.AsPayLightningInvoice()
	case "signMessageWithIdentityKey":
		return u.AsSignMessageWithIdentityKey()
	case "exportPrivateKey":
		return u.AsExportPrivateKey()
	case "exportSeedPhrase":
		return u.AsExportSeedPhrase()
	}
	return nil
}

func (u WalletRpcResponseUnion) AsPersonalSign() (v EthereumPersonalSignRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSignTypedDataV4() (v EthereumSignTypedDataRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSignTransaction() (v EthereumSignTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSendTransaction() (v EthereumSendTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSignUserOperation() (v EthereumSignUserOperationRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSign7702Authorization() (v EthereumSign7702AuthorizationRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSecp256k1Sign() (v EthereumSecp256k1SignRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsWalletSendCalls() (v EthereumSendCallsRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSignMessage() (v SolanaSignMessageRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSignTransaction() (v SolanaSignTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSignAndSendTransaction() (v SolanaSignAndSendTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsTransfer() (v SparkTransferRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsGetBalance() (v SparkGetBalanceRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsTransferTokens() (v SparkTransferTokensRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsGetStaticDepositAddress() (v SparkGetStaticDepositAddressRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsGetClaimStaticDepositQuote() (v SparkGetClaimStaticDepositQuoteRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsClaimStaticDeposit() (v SparkClaimStaticDepositRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsCreateLightningInvoice() (v SparkCreateLightningInvoiceRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsPayLightningInvoice() (v SparkPayLightningInvoiceRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSignMessageWithIdentityKey() (v SparkSignMessageWithIdentityKeyRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsExportPrivateKey() (v ExportPrivateKeyRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsExportSeedPhrase() (v ExportSeedPhraseRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletRpcResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseUnionData is an implicit subunion of [WalletRpcResponseUnion].
// WalletRpcResponseUnionData provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcResponseUnion].
type WalletRpcResponseUnionData struct {
	Encoding          string `json:"encoding"`
	Signature         string `json:"signature"`
	SignedTransaction string `json:"signed_transaction"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	Caip2         Caip2  `json:"caip2"`
	Hash          string `json:"hash"`
	ReferenceID   string `json:"reference_id"`
	TransactionID string `json:"transaction_id"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	TransactionRequest UnsignedStandardEthereumTransactionResp `json:"transaction_request"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	UserOperationHash string `json:"user_operation_hash"`
	// This field is from variant [EthereumSign7702AuthorizationRpcResponseData].
	Authorization EthereumSign7702AuthorizationResp `json:"authorization"`
	ID            string                            `json:"id"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	Leaves                    []SparkTransferLeaf `json:"leaves"`
	ReceiverIdentityPublicKey string              `json:"receiver_identity_public_key"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	SenderIdentityPublicKey string `json:"sender_identity_public_key"`
	Status                  string `json:"status"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	TotalValue float64 `json:"total_value"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	TransferDirection string `json:"transfer_direction"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	Type string `json:"type"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	CreatedTime string `json:"created_time"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	ExpiryTime string `json:"expiry_time"`
	// This field is from variant [SparkTransfer],
	// [SparkPayLightningInvoiceRpcResponseDataUnion].
	UpdatedTime string `json:"updated_time"`
	// This field is from variant [SparkBalance].
	Balance string `json:"balance"`
	// This field is from variant [SparkBalance].
	TokenBalances map[string]SparkTokenBalance `json:"token_balances"`
	// This field is from variant [SparkGetStaticDepositAddressRpcResponseData].
	Address string `json:"address"`
	// This field is from variant [SparkGetClaimStaticDepositQuoteRpcResponseData].
	CreditAmountSats float64 `json:"credit_amount_sats"`
	Network          string  `json:"network"`
	// This field is from variant [SparkGetClaimStaticDepositQuoteRpcResponseData].
	OutputIndex float64 `json:"output_index"`
	// This field is from variant [SparkClaimStaticDepositRpcResponseData].
	TransferID string `json:"transfer_id"`
	CreatedAt  string `json:"created_at"`
	Typename   string `json:"typename"`
	UpdatedAt  string `json:"updated_at"`
	// This field is from variant [SparkLightningReceiveRequest].
	Invoice         any    `json:"invoice"`
	PaymentPreimage string `json:"payment_preimage"`
	Transfer        any    `json:"transfer"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataUnion].
	EncodedInvoice string `json:"encoded_invoice"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataUnion].
	Fee SparkLightningFee `json:"fee"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataUnion].
	IdempotencyKey string `json:"idempotency_key"`
	// This field is from variant [PrivateKeyExportInputResp].
	EncryptionType HpkeEncryption `json:"encryption_type"`
	// This field is from variant [PrivateKeyExportInputResp].
	RecipientPublicKey RecipientPublicKey `json:"recipient_public_key"`
	// This field is from variant [PrivateKeyExportInputResp].
	ExportSeedPhrase bool `json:"export_seed_phrase"`
	// This field is from variant [PrivateKeyExportInputResp].
	ExportType ExportType `json:"export_type"`
	// This field is from variant [SeedPhraseExportResponse].
	Ciphertext string `json:"ciphertext"`
	// This field is from variant [SeedPhraseExportResponse].
	EncapsulatedKey string `json:"encapsulated_key"`
	JSON            struct {
		Encoding                  respjson.Field
		Signature                 respjson.Field
		SignedTransaction         respjson.Field
		Caip2                     respjson.Field
		Hash                      respjson.Field
		ReferenceID               respjson.Field
		TransactionID             respjson.Field
		TransactionRequest        respjson.Field
		UserOperationHash         respjson.Field
		Authorization             respjson.Field
		ID                        respjson.Field
		Leaves                    respjson.Field
		ReceiverIdentityPublicKey respjson.Field
		SenderIdentityPublicKey   respjson.Field
		Status                    respjson.Field
		TotalValue                respjson.Field
		TransferDirection         respjson.Field
		Type                      respjson.Field
		CreatedTime               respjson.Field
		ExpiryTime                respjson.Field
		UpdatedTime               respjson.Field
		Balance                   respjson.Field
		TokenBalances             respjson.Field
		Address                   respjson.Field
		CreditAmountSats          respjson.Field
		Network                   respjson.Field
		OutputIndex               respjson.Field
		TransferID                respjson.Field
		CreatedAt                 respjson.Field
		Typename                  respjson.Field
		UpdatedAt                 respjson.Field
		Invoice                   respjson.Field
		PaymentPreimage           respjson.Field
		Transfer                  respjson.Field
		EncodedInvoice            respjson.Field
		Fee                       respjson.Field
		IdempotencyKey            respjson.Field
		EncryptionType            respjson.Field
		RecipientPublicKey        respjson.Field
		ExportSeedPhrase          respjson.Field
		ExportType                respjson.Field
		Ciphertext                respjson.Field
		EncapsulatedKey           respjson.Field
		raw                       string
	} `json:"-"`
}

func (r *WalletRpcResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for wallet authentication with HPKE-encrypted response.
//
// The properties EncryptionType, RecipientPublicKey, UserJwt are required.
type WalletAuthenticateRequestBody struct {
	// The encryption type for the authentication response. Currently only supports
	// HPKE.
	//
	// Any of "HPKE".
	EncryptionType WalletAuthenticateRequestBodyEncryptionType `json:"encryption_type,omitzero" api:"required"`
	// The public key of your ECDH keypair, in base64-encoded, SPKI-format, whose
	// private key will be able to decrypt the session key.
	RecipientPublicKey string `json:"recipient_public_key" api:"required"`
	// The user's JWT, to be used to authenticate the user.
	UserJwt string `json:"user_jwt" api:"required"`
	paramObj
}

func (r WalletAuthenticateRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow WalletAuthenticateRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletAuthenticateRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encryption type for the authentication response. Currently only supports
// HPKE.
type WalletAuthenticateRequestBodyEncryptionType string

const (
	WalletAuthenticateRequestBodyEncryptionTypeHpke WalletAuthenticateRequestBodyEncryptionType = "HPKE"
)

// The input for private key wallets.
//
// The properties Address, ChainType, EncryptionType, EntropyType are required.
type PrivateKeyInitInput struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// Any of "private-key".
	EntropyType PrivateKeyInitInputEntropyType `json:"entropy_type,omitzero" api:"required"`
	paramObj
}

func (r PrivateKeyInitInput) MarshalJSON() (data []byte, err error) {
	type shadow PrivateKeyInitInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrivateKeyInitInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateKeyInitInputEntropyType string

const (
	PrivateKeyInitInputEntropyTypePrivateKey PrivateKeyInitInputEntropyType = "private-key"
)

// The input for HD wallets.
//
// The properties Address, ChainType, EncryptionType, EntropyType, Index are
// required.
type HDInitInput struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// The entropy type of the wallet to import.
	//
	// Any of "hd".
	EntropyType HDInitInputEntropyType `json:"entropy_type,omitzero" api:"required"`
	// The index of the wallet to import.
	Index int64 `json:"index" api:"required"`
	paramObj
}

func (r HDInitInput) MarshalJSON() (data []byte, err error) {
	type shadow HDInitInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HDInitInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The entropy type of the wallet to import.
type HDInitInputEntropyType string

const (
	HDInitInputEntropyTypeHD HDInitInputEntropyType = "hd"
)

// The submission input for importing a private key wallet.
//
// The properties Address, ChainType, Ciphertext, EncapsulatedKey, EncryptionType,
// EntropyType are required.
type PrivateKeySubmitInput struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
	// The encrypted entropy of the wallet to import.
	Ciphertext string `json:"ciphertext" api:"required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption inside the TEE.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// Any of "private-key".
	EntropyType PrivateKeySubmitInputEntropyType `json:"entropy_type,omitzero" api:"required"`
	// Optional HPKE configuration for wallet import decryption. These parameters allow
	// importing wallets encrypted by external providers that use different HPKE
	// configurations.
	HpkeConfig HpkeImportConfig `json:"hpke_config,omitzero"`
	paramObj
}

func (r PrivateKeySubmitInput) MarshalJSON() (data []byte, err error) {
	type shadow PrivateKeySubmitInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrivateKeySubmitInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivateKeySubmitInputEntropyType string

const (
	PrivateKeySubmitInputEntropyTypePrivateKey PrivateKeySubmitInputEntropyType = "private-key"
)

// The submission input for importing an HD wallet.
//
// The properties Address, ChainType, Ciphertext, EncapsulatedKey, EncryptionType,
// EntropyType, Index are required.
type HDSubmitInput struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
	// The encrypted entropy of the wallet to import.
	Ciphertext string `json:"ciphertext" api:"required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption inside the TEE.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// The entropy type of the wallet to import.
	//
	// Any of "hd".
	EntropyType HDSubmitInputEntropyType `json:"entropy_type,omitzero" api:"required"`
	// The index of the wallet to import.
	Index int64 `json:"index" api:"required"`
	// Optional HPKE configuration for wallet import decryption. These parameters allow
	// importing wallets encrypted by external providers that use different HPKE
	// configurations.
	HpkeConfig HpkeImportConfig `json:"hpke_config,omitzero"`
	paramObj
}

func (r HDSubmitInput) MarshalJSON() (data []byte, err error) {
	type shadow HDSubmitInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HDSubmitInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The entropy type of the wallet to import.
type HDSubmitInputEntropyType string

const (
	HDSubmitInputEntropyTypeHD HDSubmitInputEntropyType = "hd"
)

// Request body for exporting a wallet private key.
//
// The properties EncryptionType, RecipientPublicKey are required.
type WalletExportRequestBody struct {
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type,omitzero" api:"required"`
	// The base64-encoded encryption public key to encrypt the wallet private key with.
	RecipientPublicKey string          `json:"recipient_public_key" api:"required"`
	ExportSeedPhrase   param.Opt[bool] `json:"export_seed_phrase,omitzero"`
	paramObj
}

func (r WalletExportRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow WalletExportRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletExportRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response body containing the encrypted wallet private key.
type WalletExportResponseBody struct {
	// The encrypted private key.
	Ciphertext string `json:"ciphertext" api:"required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ciphertext      respjson.Field
		EncapsulatedKey respjson.Field
		EncryptionType  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletExportResponseBody) RawJSON() string { return r.JSON.raw }
func (r *WalletExportResponseBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SUI transaction commands allowlist for raw_sign endpoint policy evaluation
type SuiCommandName string

const (
	SuiCommandNameTransferObjects SuiCommandName = "TransferObjects"
	SuiCommandNameSplitCoins      SuiCommandName = "SplitCoins"
	SuiCommandNameMergeCoins      SuiCommandName = "MergeCoins"
)

// Source for a transfer identified by a named asset (e.g. "usdc", "eth"). Use this
// variant for first-class assets maintained by Privy.
type NamedTokenTransferSourceResp struct {
	// Amount as a decimal string in the token's standard unit (e.g. "1.5" for 1.5
	// USDC, "0.01" for 0.01 ETH). Not in the smallest on-chain unit (wei, lamports,
	// etc.).
	Amount string `json:"amount" api:"required"`
	// The asset to transfer. Supported: 'usdc', 'usdb', 'usdt' (stablecoins), 'eth'
	// (native Ethereum), 'sol' (native Solana).
	Asset string `json:"asset" api:"required"`
	// The blockchain network on which to perform the transfer. Supported chains
	// include: 'ethereum', 'base', 'arbitrum', 'polygon', 'solana', and their
	// respective testnets.
	Chain string `json:"chain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Asset       respjson.Field
		Chain       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamedTokenTransferSourceResp) RawJSON() string { return r.JSON.raw }
func (r *NamedTokenTransferSourceResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NamedTokenTransferSourceResp to a
// NamedTokenTransferSource.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NamedTokenTransferSource.Overrides()
func (r NamedTokenTransferSourceResp) ToParam() NamedTokenTransferSource {
	return param.Override[NamedTokenTransferSource](json.RawMessage(r.RawJSON()))
}

// Source for a transfer identified by a named asset (e.g. "usdc", "eth"). Use this
// variant for first-class assets maintained by Privy.
//
// The properties Amount, Asset, Chain are required.
type NamedTokenTransferSource struct {
	// Amount as a decimal string in the token's standard unit (e.g. "1.5" for 1.5
	// USDC, "0.01" for 0.01 ETH). Not in the smallest on-chain unit (wei, lamports,
	// etc.).
	Amount string `json:"amount" api:"required"`
	// The asset to transfer. Supported: 'usdc', 'usdb', 'usdt' (stablecoins), 'eth'
	// (native Ethereum), 'sol' (native Solana).
	Asset string `json:"asset" api:"required"`
	// The blockchain network on which to perform the transfer. Supported chains
	// include: 'ethereum', 'base', 'arbitrum', 'polygon', 'solana', and their
	// respective testnets.
	Chain string `json:"chain" api:"required"`
	paramObj
}

func (r NamedTokenTransferSource) MarshalJSON() (data []byte, err error) {
	type shadow NamedTokenTransferSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamedTokenTransferSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source for a transfer identified by a token contract address (EVM) or mint
// address (Solana). Use this variant for tokens that are not first-class assets.
type CustomTokenTransferSourceResp struct {
	// Amount as a decimal string in the token's standard unit (e.g. "1.5" for 1.5
	// USDC, "0.01" for 0.01 ETH). Not in the smallest on-chain unit (wei, lamports,
	// etc.).
	Amount string `json:"amount" api:"required"`
	// The token contract address (EVM) or mint address (Solana) of the asset to
	// transfer.
	AssetAddress string `json:"asset_address" api:"required"`
	// The blockchain network on which to perform the transfer. Supported chains
	// include: 'ethereum', 'base', 'arbitrum', 'polygon', 'solana', and their
	// respective testnets.
	Chain string `json:"chain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		AssetAddress respjson.Field
		Chain        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomTokenTransferSourceResp) RawJSON() string { return r.JSON.raw }
func (r *CustomTokenTransferSourceResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CustomTokenTransferSourceResp to a
// CustomTokenTransferSource.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CustomTokenTransferSource.Overrides()
func (r CustomTokenTransferSourceResp) ToParam() CustomTokenTransferSource {
	return param.Override[CustomTokenTransferSource](json.RawMessage(r.RawJSON()))
}

// Source for a transfer identified by a token contract address (EVM) or mint
// address (Solana). Use this variant for tokens that are not first-class assets.
//
// The properties Amount, AssetAddress, Chain are required.
type CustomTokenTransferSource struct {
	// Amount as a decimal string in the token's standard unit (e.g. "1.5" for 1.5
	// USDC, "0.01" for 0.01 ETH). Not in the smallest on-chain unit (wei, lamports,
	// etc.).
	Amount string `json:"amount" api:"required"`
	// The token contract address (EVM) or mint address (Solana) of the asset to
	// transfer.
	AssetAddress string `json:"asset_address" api:"required"`
	// The blockchain network on which to perform the transfer. Supported chains
	// include: 'ethereum', 'base', 'arbitrum', 'polygon', 'solana', and their
	// respective testnets.
	Chain string `json:"chain" api:"required"`
	paramObj
}

func (r CustomTokenTransferSource) MarshalJSON() (data []byte, err error) {
	type shadow CustomTokenTransferSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomTokenTransferSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TokenTransferSourceUnionResp contains all possible properties and values from
// [NamedTokenTransferSourceResp], [CustomTokenTransferSourceResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TokenTransferSourceUnionResp struct {
	Amount string `json:"amount"`
	// This field is from variant [NamedTokenTransferSourceResp].
	Asset string `json:"asset"`
	Chain string `json:"chain"`
	// This field is from variant [CustomTokenTransferSourceResp].
	AssetAddress string `json:"asset_address"`
	JSON         struct {
		Amount       respjson.Field
		Asset        respjson.Field
		Chain        respjson.Field
		AssetAddress respjson.Field
		raw          string
	} `json:"-"`
}

func (u TokenTransferSourceUnionResp) AsNamedTokenTransferSource() (v NamedTokenTransferSourceResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TokenTransferSourceUnionResp) AsCustomTokenTransferSource() (v CustomTokenTransferSourceResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TokenTransferSourceUnionResp) RawJSON() string { return u.JSON.raw }

func (r *TokenTransferSourceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TokenTransferSourceUnionResp to a
// TokenTransferSourceUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TokenTransferSourceUnion.Overrides()
func (r TokenTransferSourceUnionResp) ToParam() TokenTransferSourceUnion {
	return param.Override[TokenTransferSourceUnion](json.RawMessage(r.RawJSON()))
}

func TokenTransferSourceOfNamedTokenTransferSource(amount string, asset string, chain string) TokenTransferSourceUnion {
	var variant NamedTokenTransferSource
	variant.Amount = amount
	variant.Asset = asset
	variant.Chain = chain
	return TokenTransferSourceUnion{OfNamedTokenTransferSource: &variant}
}

func TokenTransferSourceOfCustomTokenTransferSource(amount string, assetAddress string, chain string) TokenTransferSourceUnion {
	var variant CustomTokenTransferSource
	variant.Amount = amount
	variant.AssetAddress = assetAddress
	variant.Chain = chain
	return TokenTransferSourceUnion{OfCustomTokenTransferSource: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TokenTransferSourceUnion struct {
	OfNamedTokenTransferSource  *NamedTokenTransferSource  `json:",omitzero,inline"`
	OfCustomTokenTransferSource *CustomTokenTransferSource `json:",omitzero,inline"`
	paramUnion
}

func (u TokenTransferSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNamedTokenTransferSource, u.OfCustomTokenTransferSource)
}
func (u *TokenTransferSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The destination address for a token transfer. Optionally specify a different
// asset or chain for cross-asset or cross-chain transfers.
type TokenTransferDestinationResp struct {
	// Recipient address (hex for EVM, base58 for Solana)
	Address string `json:"address" api:"required"`
	// The destination asset. Required for cross-asset transfers (e.g., source 'usdt'
	// to destination 'usdc').
	Asset string `json:"asset"`
	// The destination blockchain network. Required for cross-chain transfers (e.g.,
	// source 'base' to destination 'arbitrum').
	Chain string `json:"chain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Asset       respjson.Field
		Chain       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenTransferDestinationResp) RawJSON() string { return r.JSON.raw }
func (r *TokenTransferDestinationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TokenTransferDestinationResp to a
// TokenTransferDestination.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TokenTransferDestination.Overrides()
func (r TokenTransferDestinationResp) ToParam() TokenTransferDestination {
	return param.Override[TokenTransferDestination](json.RawMessage(r.RawJSON()))
}

// The destination address for a token transfer. Optionally specify a different
// asset or chain for cross-asset or cross-chain transfers.
//
// The property Address is required.
type TokenTransferDestination struct {
	// Recipient address (hex for EVM, base58 for Solana)
	Address string `json:"address" api:"required"`
	// The destination asset. Required for cross-asset transfers (e.g., source 'usdt'
	// to destination 'usdc').
	Asset param.Opt[string] `json:"asset,omitzero"`
	// The destination blockchain network. Required for cross-chain transfers (e.g.,
	// source 'base' to destination 'arbitrum').
	Chain param.Opt[string] `json:"chain,omitzero"`
	paramObj
}

func (r TokenTransferDestination) MarshalJSON() (data []byte, err error) {
	type shadow TokenTransferDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TokenTransferDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for initiating a sponsored token transfer from an embedded wallet.
type TransferRequestBodyResp struct {
	// The destination address for a token transfer. Optionally specify a different
	// asset or chain for cross-asset or cross-chain transfers.
	Destination TokenTransferDestinationResp `json:"destination" api:"required"`
	// The source asset, amount, and chain for a token transfer. Specify either `asset`
	// (named) or `asset_address` (custom), not both.
	Source TokenTransferSourceUnionResp `json:"source" api:"required"`
	// Whether the amount refers to the input token or output token.
	//
	// Any of "exact_input", "exact_output".
	AmountType AmountType `json:"amount_type"`
	// Maximum allowed slippage in basis points (1 bps = 0.01%).
	SlippageBps int64 `json:"slippage_bps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destination respjson.Field
		Source      respjson.Field
		AmountType  respjson.Field
		SlippageBps respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferRequestBodyResp) RawJSON() string { return r.JSON.raw }
func (r *TransferRequestBodyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TransferRequestBodyResp to a TransferRequestBody.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TransferRequestBody.Overrides()
func (r TransferRequestBodyResp) ToParam() TransferRequestBody {
	return param.Override[TransferRequestBody](json.RawMessage(r.RawJSON()))
}

// Request body for initiating a sponsored token transfer from an embedded wallet.
//
// The properties Destination, Source are required.
type TransferRequestBody struct {
	// The destination address for a token transfer. Optionally specify a different
	// asset or chain for cross-asset or cross-chain transfers.
	Destination TokenTransferDestination `json:"destination,omitzero" api:"required"`
	// The source asset, amount, and chain for a token transfer. Specify either `asset`
	// (named) or `asset_address` (custom), not both.
	Source TokenTransferSourceUnion `json:"source,omitzero" api:"required"`
	// Maximum allowed slippage in basis points (1 bps = 0.01%).
	SlippageBps param.Opt[int64] `json:"slippage_bps,omitzero"`
	// Whether the amount refers to the input token or output token.
	//
	// Any of "exact_input", "exact_output".
	AmountType AmountType `json:"amount_type,omitzero"`
	paramObj
}

func (r TransferRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow TransferRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletInitImportResponse struct {
	// The base64-encoded encryption public key to encrypt the wallet entropy with.
	EncryptionPublicKey string `json:"encryption_public_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType HpkeEncryption `json:"encryption_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptionPublicKey respjson.Field
		EncryptionType      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletInitImportResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletInitImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletAuthenticateWithJwtResponseUnion contains all possible properties and
// values from [WalletAuthenticateWithJwtResponseWithEncryption],
// [WalletAuthenticateWithJwtResponseWithoutEncryption].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletAuthenticateWithJwtResponseUnion struct {
	// This field is from variant [WalletAuthenticateWithJwtResponseWithEncryption].
	EncryptedAuthorizationKey WalletAuthenticateWithJwtResponseWithEncryptionEncryptedAuthorizationKey `json:"encrypted_authorization_key"`
	ExpiresAt                 float64                                                                  `json:"expires_at"`
	Wallets                   []Wallet                                                                 `json:"wallets"`
	// This field is from variant [WalletAuthenticateWithJwtResponseWithoutEncryption].
	AuthorizationKey string `json:"authorization_key"`
	JSON             struct {
		EncryptedAuthorizationKey respjson.Field
		ExpiresAt                 respjson.Field
		Wallets                   respjson.Field
		AuthorizationKey          respjson.Field
		raw                       string
	} `json:"-"`
}

func (u WalletAuthenticateWithJwtResponseUnion) AsWithEncryption() (v WalletAuthenticateWithJwtResponseWithEncryption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletAuthenticateWithJwtResponseUnion) AsWithoutEncryption() (v WalletAuthenticateWithJwtResponseWithoutEncryption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletAuthenticateWithJwtResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletAuthenticateWithJwtResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAuthenticateWithJwtResponseWithEncryption struct {
	// The encrypted authorization key data.
	EncryptedAuthorizationKey WalletAuthenticateWithJwtResponseWithEncryptionEncryptedAuthorizationKey `json:"encrypted_authorization_key" api:"required"`
	// The expiration time of the authorization key in milliseconds since the epoch.
	ExpiresAt float64  `json:"expires_at" api:"required"`
	Wallets   []Wallet `json:"wallets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptedAuthorizationKey respjson.Field
		ExpiresAt                 respjson.Field
		Wallets                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAuthenticateWithJwtResponseWithEncryption) RawJSON() string { return r.JSON.raw }
func (r *WalletAuthenticateWithJwtResponseWithEncryption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encrypted authorization key data.
type WalletAuthenticateWithJwtResponseWithEncryptionEncryptedAuthorizationKey struct {
	// The encrypted authorization key corresponding to the user's current
	// authentication session.
	Ciphertext string `json:"ciphertext" api:"required"`
	// Base64-encoded ephemeral public key used in the HPKE encryption process.
	// Required for decryption.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type used. Currently only supports HPKE.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ciphertext      respjson.Field
		EncapsulatedKey respjson.Field
		EncryptionType  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAuthenticateWithJwtResponseWithEncryptionEncryptedAuthorizationKey) RawJSON() string {
	return r.JSON.raw
}
func (r *WalletAuthenticateWithJwtResponseWithEncryptionEncryptedAuthorizationKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAuthenticateWithJwtResponseWithoutEncryption struct {
	// The raw authorization key data.
	AuthorizationKey string `json:"authorization_key" api:"required"`
	// The expiration time of the authorization key in milliseconds since the epoch.
	ExpiresAt float64  `json:"expires_at" api:"required"`
	Wallets   []Wallet `json:"wallets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationKey respjson.Field
		ExpiresAt        respjson.Field
		Wallets          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAuthenticateWithJwtResponseWithoutEncryption) RawJSON() string { return r.JSON.raw }
func (r *WalletAuthenticateWithJwtResponseWithoutEncryption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletNewParams struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "bitcoin-taproot", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero" api:"required"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// A human-readable label for the wallet.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// A customer-provided identifier for mapping to external systems. URL-safe
	// characters only ([a-zA-Z0-9_-]), max 64 chars. Write-once: cannot be changed
	// after creation.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners AdditionalSignerInput `json:"additional_signers,omitzero"`
	// An optional list of up to one policy ID to enforce on the wallet.
	PolicyIDs PolicyInput `json:"policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r WalletNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletUpdateParams struct {
	// Request body for updating a wallet.
	WalletUpdateRequestBody WalletUpdateRequestBody
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r WalletUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletUpdateRequestBody)
}
func (r *WalletUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletListParams struct {
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Filter wallets by authorization public key. Returns wallets owned by key quorums
	// that include the specified P-256 public key (base64-encoded DER format). Cannot
	// be used together with user_id.
	AuthorizationKey param.Opt[string] `query:"authorization_key,omitzero" json:"-"`
	Cursor           param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter wallets by external ID.
	ExternalID param.Opt[string] `query:"external_id,omitzero" json:"-"`
	// Filter wallets by user ID. Cannot be used together with authorization_key.
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "bitcoin-taproot", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `query:"chain_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletListParams]'s query parameters as `url.Values`.
func (r WalletListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletInitImportParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. The
	// input for HD wallets.
	OfHD *HDInitInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set. The
	// input for private key wallets.
	OfPrivateKey *PrivateKeyInitInput `json:",inline"`

	paramObj
}

func (u WalletInitImportParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHD, u.OfPrivateKey)
}
func (r *WalletInitImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletSubmitImportParams struct {
	// The submission input for importing an HD wallet.
	Wallet WalletSubmitImportParamsWalletUnion `json:"wallet,omitzero" api:"required"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[OwnerIDInput] `json:"owner_id,omitzero" format:"cuid2"`
	// A human-readable label for the wallet.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// A customer-provided identifier for mapping to external systems. URL-safe
	// characters only ([a-zA-Z0-9_-]), max 64 chars. Write-once: cannot be changed
	// after creation.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// The owner of the resource, specified as a Privy user ID, a P-256 public key, or
	// null to remove the current owner.
	Owner OwnerInputUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners AdditionalSignerInput `json:"additional_signers,omitzero"`
	// An optional list of up to one policy ID to enforce on the wallet.
	PolicyIDs PolicyInput `json:"policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r WalletSubmitImportParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletSubmitImportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletSubmitImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletSubmitImportParamsWalletUnion struct {
	OfHD         *HDSubmitInput         `json:",omitzero,inline"`
	OfPrivateKey *PrivateKeySubmitInput `json:",omitzero,inline"`
	paramUnion
}

func (u WalletSubmitImportParamsWalletUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHD, u.OfPrivateKey)
}
func (u *WalletSubmitImportParamsWalletUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WalletSubmitImportParamsWalletUnion](
		"entropy_type",
		apijson.Discriminator[HDSubmitInput]("hd"),
		apijson.Discriminator[PrivateKeySubmitInput]("private-key"),
	)
}

type WalletAuthenticateWithJwtParams struct {
	// Request body for wallet authentication with HPKE-encrypted response.
	WalletAuthenticateRequestBody WalletAuthenticateRequestBody
	paramObj
}

func (r WalletAuthenticateWithJwtParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletAuthenticateRequestBody)
}
func (r *WalletAuthenticateWithJwtParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletBatchParams struct {
	// Request body for batch wallet creation.
	WalletBatchCreateInput WalletBatchCreateInput
	paramObj
}

func (r WalletBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletBatchCreateInput)
}
func (r *WalletBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletExportParams struct {
	// Request body for exporting a wallet private key.
	WalletExportRequestBody WalletExportRequestBody
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r WalletExportParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletExportRequestBody)
}
func (r *WalletExportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletGetWalletByAddressParams struct {
	// Request body for looking up a wallet by its blockchain address.
	GetByWalletAddressRequestBody GetByWalletAddressRequestBody
	paramObj
}

func (r WalletGetWalletByAddressParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GetByWalletAddressRequestBody)
}
func (r *WalletGetWalletByAddressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRawSignParams struct {
	// Provide either `hash` (to sign a pre-computed hash) OR `bytes`, `encoding`, and
	// `hash_function` (to hash and then sign). These options are mutually exclusive.
	RawSignInput RawSignInput
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r WalletRawSignParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RawSignInput)
}
func (r *WalletRawSignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcParams struct {
	// Request body for wallet RPC operations, discriminated by method.
	WalletRpcRequestBody WalletRpcRequestBodyUnion
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
	paramObj
}

func (r WalletRpcParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletRpcRequestBody)
}
func (r *WalletRpcParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
