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
	"github.com/privy-io/go-sdk/shared/constant"
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

// Export a wallet's private key
func (r *WalletService) Export(ctx context.Context, walletID string, params WalletExportParams, opts ...option.RequestOption) (res *WalletExportResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
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

// Sign a message with a wallet by wallet ID.
func (r *WalletService) RawSign(ctx context.Context, walletID string, params WalletRawSignParams, opts ...option.RequestOption) (res *WalletRawSignResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
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
func (r *WalletService) Rpc(ctx context.Context, walletID string, params WalletRpcParams, opts ...option.RequestOption) (res *WalletRpcResponseBodyUnion, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%v", params.PrivyIdempotencyKey.Value)))
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

// A wallet managed by Privy's wallet infrastructure.
type Wallet struct {
	// Unique ID of the wallet. This will be the primary identifier when using the
	// wallet in the future.
	ID string `json:"id" api:"required"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletAdditionalSigner `json:"additional_signers" api:"required"`
	// Address of the wallet.
	Address string `json:"address" api:"required"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
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
	OwnerID string `json:"owner_id" api:"required"`
	// List of policy IDs for policies that are enforced on the wallet.
	PolicyIDs []string `json:"policy_ids" api:"required"`
	// The compressed, raw public key for the wallet along the chain cryptographic
	// curve.
	PublicKey string `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AdditionalSigners respjson.Field
		Address           respjson.Field
		ChainType         respjson.Field
		CreatedAt         respjson.Field
		ExportedAt        respjson.Field
		ImportedAt        respjson.Field
		OwnerID           respjson.Field
		PolicyIDs         respjson.Field
		PublicKey         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Wallet) RawJSON() string { return r.JSON.raw }
func (r *Wallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAdditionalSigner struct {
	SignerID string `json:"signer_id" api:"required" format:"cuid2"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet.
	OverridePolicyIDs []string `json:"override_policy_ids" format:"cuid2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignerID          respjson.Field
		OverridePolicyIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAdditionalSigner) RawJSON() string { return r.JSON.raw }
func (r *WalletAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The wallet chain types that support curve-based signing.
type CurveSigningChainType string

const (
	CurveSigningChainTypeCosmos        CurveSigningChainType = "cosmos"
	CurveSigningChainTypeStellar       CurveSigningChainType = "stellar"
	CurveSigningChainTypeSui           CurveSigningChainType = "sui"
	CurveSigningChainTypeAptos         CurveSigningChainType = "aptos"
	CurveSigningChainTypeMovement      CurveSigningChainType = "movement"
	CurveSigningChainTypeTron          CurveSigningChainType = "tron"
	CurveSigningChainTypeBitcoinSegwit CurveSigningChainType = "bitcoin-segwit"
	CurveSigningChainTypeNear          CurveSigningChainType = "near"
	CurveSigningChainTypeTon           CurveSigningChainType = "ton"
	CurveSigningChainTypeStarknet      CurveSigningChainType = "starknet"
)

// The wallet chain types.
type WalletChainType string

const (
	WalletChainTypeEthereum      WalletChainType = "ethereum"
	WalletChainTypeSolana        WalletChainType = "solana"
	WalletChainTypeCosmos        WalletChainType = "cosmos"
	WalletChainTypeStellar       WalletChainType = "stellar"
	WalletChainTypeSui           WalletChainType = "sui"
	WalletChainTypeAptos         WalletChainType = "aptos"
	WalletChainTypeMovement      WalletChainType = "movement"
	WalletChainTypeTron          WalletChainType = "tron"
	WalletChainTypeBitcoinSegwit WalletChainType = "bitcoin-segwit"
	WalletChainTypeNear          WalletChainType = "near"
	WalletChainTypeTon           WalletChainType = "ton"
	WalletChainTypeStarknet      WalletChainType = "starknet"
	WalletChainTypeSpark         WalletChainType = "spark"
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
	// The AEAD algorithm used for encryption. Defaults to CHACHA20_POLY1305 if not
	// specified.
	//
	// Any of "CHACHA20_POLY1305", "AES_GCM256".
	AeadAlgorithm HpkeImportConfigAeadAlgorithm `json:"aead_algorithm,omitzero"`
	paramObj
}

func (r HpkeImportConfig) MarshalJSON() (data []byte, err error) {
	type shadow HpkeImportConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HpkeImportConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The AEAD algorithm used for encryption. Defaults to CHACHA20_POLY1305 if not
// specified.
type HpkeImportConfigAeadAlgorithm string

const (
	HpkeImportConfigAeadAlgorithmChacha20Poly1305 HpkeImportConfigAeadAlgorithm = "CHACHA20_POLY1305"
	HpkeImportConfigAeadAlgorithmAesGcm256        HpkeImportConfigAeadAlgorithm = "AES_GCM256"
)

// SUI transaction commands allowlist for raw_sign endpoint policy evaluation
type SuiCommandName string

const (
	SuiCommandNameTransferObjects SuiCommandName = "TransferObjects"
	SuiCommandNameSplitCoins      SuiCommandName = "SplitCoins"
	SuiCommandNameMergeCoins      SuiCommandName = "MergeCoins"
)

// Request body for updating a wallet.
type PatchWalletRequestBody struct {
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner PatchWalletRequestBodyOwnerUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners []PatchWalletRequestBodyAdditionalSigner `json:"additional_signers,omitzero"`
	// New policy IDs to enforce on the wallet. Currently, only one policy is supported
	// per wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r PatchWalletRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow PatchWalletRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PatchWalletRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SignerID is required.
type PatchWalletRequestBodyAdditionalSigner struct {
	SignerID string `json:"signer_id" api:"required" format:"cuid2"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet.
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r PatchWalletRequestBodyAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow PatchWalletRequestBodyAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PatchWalletRequestBodyAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PatchWalletRequestBodyOwnerUnion struct {
	OfPublicKeyOwner *PatchWalletRequestBodyOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *PatchWalletRequestBodyOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u PatchWalletRequestBodyOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *PatchWalletRequestBodyOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type PatchWalletRequestBodyOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r PatchWalletRequestBodyOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow PatchWalletRequestBodyOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PatchWalletRequestBodyOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type PatchWalletRequestBodyOwnerUserOwner struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r PatchWalletRequestBodyOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow PatchWalletRequestBodyOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PatchWalletRequestBodyOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Executes the EVM `personal_sign` RPC (EIP-191) to sign a message.
//
// The properties Method, Params are required.
type EthereumPersonalSignRpcInput struct {
	// Any of "personal_sign".
	Method  EthereumPersonalSignRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumPersonalSignRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                  `json:"address,omitzero"`
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

type EthereumPersonalSignRpcInputMethod string

const (
	EthereumPersonalSignRpcInputMethodPersonalSign EthereumPersonalSignRpcInputMethod = "personal_sign"
)

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

type EthereumPersonalSignRpcInputParamsEncoding string

const (
	EthereumPersonalSignRpcInputParamsEncodingUtf8 EthereumPersonalSignRpcInputParamsEncoding = "utf-8"
	EthereumPersonalSignRpcInputParamsEncodingHex  EthereumPersonalSignRpcInputParamsEncoding = "hex"
)

type EthereumPersonalSignRpcInputChainType string

const (
	EthereumPersonalSignRpcInputChainTypeEthereum EthereumPersonalSignRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_signTransaction` RPC to sign a transaction.
//
// The properties Method, Params are required.
type EthereumSignTransactionRpcInput struct {
	// Any of "eth_signTransaction".
	Method  EthereumSignTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumSignTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                     `json:"address,omitzero"`
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

type EthereumSignTransactionRpcInputMethod string

const (
	EthereumSignTransactionRpcInputMethodEthSignTransaction EthereumSignTransactionRpcInputMethod = "eth_signTransaction"
)

// The property Transaction is required.
type EthereumSignTransactionRpcInputParams struct {
	Transaction EthereumSignTransactionRpcInputParamsTransaction `json:"transaction,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTransactionRpcInputParamsTransaction struct {
	Data                 param.Opt[string]                                                         `json:"data,omitzero"`
	From                 param.Opt[string]                                                         `json:"from,omitzero"`
	To                   param.Opt[string]                                                         `json:"to,omitzero"`
	AuthorizationList    []EthereumSignTransactionRpcInputParamsTransactionAuthorizationList       `json:"authorization_list,omitzero"`
	ChainID              EthereumSignTransactionRpcInputParamsTransactionChainIDUnion              `json:"chain_id,omitzero"`
	GasLimit             EthereumSignTransactionRpcInputParamsTransactionGasLimitUnion             `json:"gas_limit,omitzero"`
	GasPrice             EthereumSignTransactionRpcInputParamsTransactionGasPriceUnion             `json:"gas_price,omitzero"`
	MaxFeePerGas         EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnion         `json:"max_fee_per_gas,omitzero"`
	MaxPriorityFeePerGas EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas,omitzero"`
	Nonce                EthereumSignTransactionRpcInputParamsTransactionNonceUnion                `json:"nonce,omitzero"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                    `json:"type,omitzero"`
	Value EthereumSignTransactionRpcInputParamsTransactionValueUnion `json:"value,omitzero"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EthereumSignTransactionRpcInputParamsTransaction](
		"type", 0, 1, 2, 4,
	)
}

// The properties ChainID, Contract, Nonce, R, S, YParity are required.
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationList struct {
	ChainID  EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion `json:"chain_id,omitzero" api:"required"`
	Contract string                                                                        `json:"contract" api:"required"`
	Nonce    EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnion   `json:"nonce,omitzero" api:"required"`
	R        string                                                                        `json:"r" api:"required"`
	S        string                                                                        `json:"s" api:"required"`
	YParity  float64                                                                       `json:"y_parity" api:"required"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParamsTransactionAuthorizationList) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParamsTransactionAuthorizationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParamsTransactionAuthorizationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionGasLimitUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionGasLimitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionGasPriceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionGasPriceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionValueUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type EthereumSignTransactionRpcInputChainType string

const (
	EthereumSignTransactionRpcInputChainTypeEthereum EthereumSignTransactionRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_sendTransaction` RPC to sign and broadcast a transaction.
//
// The properties Caip2, Method, Params are required.
type EthereumSendTransactionRpcInput struct {
	Caip2 string `json:"caip2" api:"required"`
	// Any of "eth_sendTransaction".
	Method  EthereumSendTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumSendTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                     `json:"address,omitzero"`
	Sponsor param.Opt[bool]                       `json:"sponsor,omitzero"`
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

type EthereumSendTransactionRpcInputMethod string

const (
	EthereumSendTransactionRpcInputMethodEthSendTransaction EthereumSendTransactionRpcInputMethod = "eth_sendTransaction"
)

// The property Transaction is required.
type EthereumSendTransactionRpcInputParams struct {
	Transaction EthereumSendTransactionRpcInputParamsTransaction `json:"transaction,omitzero" api:"required"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendTransactionRpcInputParamsTransaction struct {
	Data                 param.Opt[string]                                                         `json:"data,omitzero"`
	From                 param.Opt[string]                                                         `json:"from,omitzero"`
	To                   param.Opt[string]                                                         `json:"to,omitzero"`
	AuthorizationList    []EthereumSendTransactionRpcInputParamsTransactionAuthorizationList       `json:"authorization_list,omitzero"`
	ChainID              EthereumSendTransactionRpcInputParamsTransactionChainIDUnion              `json:"chain_id,omitzero"`
	GasLimit             EthereumSendTransactionRpcInputParamsTransactionGasLimitUnion             `json:"gas_limit,omitzero"`
	GasPrice             EthereumSendTransactionRpcInputParamsTransactionGasPriceUnion             `json:"gas_price,omitzero"`
	MaxFeePerGas         EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnion         `json:"max_fee_per_gas,omitzero"`
	MaxPriorityFeePerGas EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas,omitzero"`
	Nonce                EthereumSendTransactionRpcInputParamsTransactionNonceUnion                `json:"nonce,omitzero"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                    `json:"type,omitzero"`
	Value EthereumSendTransactionRpcInputParamsTransactionValueUnion `json:"value,omitzero"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EthereumSendTransactionRpcInputParamsTransaction](
		"type", 0, 1, 2, 4,
	)
}

// The properties ChainID, Contract, Nonce, R, S, YParity are required.
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationList struct {
	ChainID  EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion `json:"chain_id,omitzero" api:"required"`
	Contract string                                                                        `json:"contract" api:"required"`
	Nonce    EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnion   `json:"nonce,omitzero" api:"required"`
	R        string                                                                        `json:"r" api:"required"`
	S        string                                                                        `json:"s" api:"required"`
	YParity  float64                                                                       `json:"y_parity" api:"required"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParamsTransactionAuthorizationList) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParamsTransactionAuthorizationList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParamsTransactionAuthorizationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionGasLimitUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionGasLimitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionGasPriceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionGasPriceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionValueUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type EthereumSendTransactionRpcInputChainType string

const (
	EthereumSendTransactionRpcInputChainTypeEthereum EthereumSendTransactionRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_signTypedData_v4` RPC (EIP-712) to sign a typed data
// object.
//
// The properties Method, Params are required.
type EthereumSignTypedDataRpcInput struct {
	// Any of "eth_signTypedData_v4".
	Method  EthereumSignTypedDataRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumSignTypedDataRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                   `json:"address,omitzero"`
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

type EthereumSignTypedDataRpcInputMethod string

const (
	EthereumSignTypedDataRpcInputMethodEthSignTypedDataV4 EthereumSignTypedDataRpcInputMethod = "eth_signTypedData_v4"
)

// The property TypedData is required.
type EthereumSignTypedDataRpcInputParams struct {
	TypedData EthereumSignTypedDataRpcInputParamsTypedData `json:"typed_data,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domain, Message, PrimaryType, Types are required.
type EthereumSignTypedDataRpcInputParamsTypedData struct {
	Domain      map[string]any                                                `json:"domain,omitzero" api:"required"`
	Message     map[string]any                                                `json:"message,omitzero" api:"required"`
	PrimaryType string                                                        `json:"primary_type" api:"required"`
	Types       map[string][]EthereumSignTypedDataRpcInputParamsTypedDataType `json:"types,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParamsTypedData) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParamsTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParamsTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type EthereumSignTypedDataRpcInputParamsTypedDataType struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParamsTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParamsTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParamsTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTypedDataRpcInputChainType string

const (
	EthereumSignTypedDataRpcInputChainTypeEthereum EthereumSignTypedDataRpcInputChainType = "ethereum"
)

// Executes an RPC method to hash and sign a UserOperation.
//
// The properties Method, Params are required.
type EthereumSignUserOperationRpcInput struct {
	// Any of "eth_signUserOperation".
	Method  EthereumSignUserOperationRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumSignUserOperationRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                       `json:"address,omitzero"`
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

type EthereumSignUserOperationRpcInputMethod string

const (
	EthereumSignUserOperationRpcInputMethodEthSignUserOperation EthereumSignUserOperationRpcInputMethod = "eth_signUserOperation"
)

// The properties ChainID, Contract, UserOperation are required.
type EthereumSignUserOperationRpcInputParams struct {
	ChainID       EthereumSignUserOperationRpcInputParamsChainIDUnion  `json:"chain_id,omitzero" api:"required"`
	Contract      string                                               `json:"contract" api:"required"`
	UserOperation EthereumSignUserOperationRpcInputParamsUserOperation `json:"user_operation,omitzero" api:"required"`
	paramObj
}

func (r EthereumSignUserOperationRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignUserOperationRpcInputParamsChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignUserOperationRpcInputParamsChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignUserOperationRpcInputParamsChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties CallData, CallGasLimit, MaxFeePerGas, MaxPriorityFeePerGas,
// Nonce, Paymaster, PaymasterData, PaymasterPostOpGasLimit,
// PaymasterVerificationGasLimit, PreVerificationGas, Sender, VerificationGasLimit
// are required.
type EthereumSignUserOperationRpcInputParamsUserOperation struct {
	CallData                      string `json:"call_data" api:"required"`
	CallGasLimit                  string `json:"call_gas_limit" api:"required"`
	MaxFeePerGas                  string `json:"max_fee_per_gas" api:"required"`
	MaxPriorityFeePerGas          string `json:"max_priority_fee_per_gas" api:"required"`
	Nonce                         string `json:"nonce" api:"required"`
	Paymaster                     string `json:"paymaster" api:"required"`
	PaymasterData                 string `json:"paymaster_data" api:"required"`
	PaymasterPostOpGasLimit       string `json:"paymaster_post_op_gas_limit" api:"required"`
	PaymasterVerificationGasLimit string `json:"paymaster_verification_gas_limit" api:"required"`
	PreVerificationGas            string `json:"pre_verification_gas" api:"required"`
	Sender                        string `json:"sender" api:"required"`
	VerificationGasLimit          string `json:"verification_gas_limit" api:"required"`
	paramObj
}

func (r EthereumSignUserOperationRpcInputParamsUserOperation) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInputParamsUserOperation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInputParamsUserOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignUserOperationRpcInputChainType string

const (
	EthereumSignUserOperationRpcInputChainTypeEthereum EthereumSignUserOperationRpcInputChainType = "ethereum"
)

// Signs an EIP-7702 authorization.
//
// The properties Method, Params are required.
type EthereumSign7702AuthorizationRpcInput struct {
	// Any of "eth_sign7702Authorization".
	Method  EthereumSign7702AuthorizationRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumSign7702AuthorizationRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                           `json:"address,omitzero"`
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

type EthereumSign7702AuthorizationRpcInputMethod string

const (
	EthereumSign7702AuthorizationRpcInputMethodEthSign7702Authorization EthereumSign7702AuthorizationRpcInputMethod = "eth_sign7702Authorization"
)

// The properties ChainID, Contract are required.
type EthereumSign7702AuthorizationRpcInputParams struct {
	ChainID  EthereumSign7702AuthorizationRpcInputParamsChainIDUnion `json:"chain_id,omitzero" api:"required"`
	Contract string                                                  `json:"contract" api:"required"`
	Nonce    EthereumSign7702AuthorizationRpcInputParamsNonceUnion   `json:"nonce,omitzero"`
	paramObj
}

func (r EthereumSign7702AuthorizationRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSign7702AuthorizationRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSign7702AuthorizationRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSign7702AuthorizationRpcInputParamsChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSign7702AuthorizationRpcInputParamsChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSign7702AuthorizationRpcInputParamsChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSign7702AuthorizationRpcInputParamsNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSign7702AuthorizationRpcInputParamsNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSign7702AuthorizationRpcInputParamsNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type EthereumSign7702AuthorizationRpcInputChainType string

const (
	EthereumSign7702AuthorizationRpcInputChainTypeEthereum EthereumSign7702AuthorizationRpcInputChainType = "ethereum"
)

// Signs a raw hash on the secp256k1 curve.
//
// The properties Method, Params are required.
type EthereumSecp256k1SignRpcInput struct {
	// Any of "secp256k1_sign".
	Method  EthereumSecp256k1SignRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  EthereumSecp256k1SignRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                   `json:"address,omitzero"`
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

type EthereumSecp256k1SignRpcInputMethod string

const (
	EthereumSecp256k1SignRpcInputMethodSecp256k1Sign EthereumSecp256k1SignRpcInputMethod = "secp256k1_sign"
)

// The property Hash is required.
type EthereumSecp256k1SignRpcInputParams struct {
	Hash string `json:"hash" api:"required"`
	paramObj
}

func (r EthereumSecp256k1SignRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSecp256k1SignRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSecp256k1SignRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSecp256k1SignRpcInputChainType string

const (
	EthereumSecp256k1SignRpcInputChainTypeEthereum EthereumSecp256k1SignRpcInputChainType = "ethereum"
)

// Executes the SVM `signTransaction` RPC to sign a transaction.
//
// The properties Method, Params are required.
type SolanaSignTransactionRpcInput struct {
	// Any of "signTransaction".
	Method  SolanaSignTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  SolanaSignTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                   `json:"address,omitzero"`
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

type SolanaSignTransactionRpcInputMethod string

const (
	SolanaSignTransactionRpcInputMethodSignTransaction SolanaSignTransactionRpcInputMethod = "signTransaction"
)

// The properties Encoding, Transaction are required.
type SolanaSignTransactionRpcInputParams struct {
	// Any of "base64".
	Encoding    string `json:"encoding,omitzero" api:"required"`
	Transaction string `json:"transaction" api:"required"`
	paramObj
}

func (r SolanaSignTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SolanaSignTransactionRpcInputParams](
		"encoding", "base64",
	)
}

type SolanaSignTransactionRpcInputChainType string

const (
	SolanaSignTransactionRpcInputChainTypeSolana SolanaSignTransactionRpcInputChainType = "solana"
)

// Executes the SVM `signAndSendTransaction` RPC to sign and broadcast a
// transaction.
//
// The properties Caip2, Method, Params are required.
type SolanaSignAndSendTransactionRpcInput struct {
	Caip2 string `json:"caip2" api:"required"`
	// Any of "signAndSendTransaction".
	Method  SolanaSignAndSendTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  SolanaSignAndSendTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]                          `json:"address,omitzero"`
	Sponsor param.Opt[bool]                            `json:"sponsor,omitzero"`
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

type SolanaSignAndSendTransactionRpcInputMethod string

const (
	SolanaSignAndSendTransactionRpcInputMethodSignAndSendTransaction SolanaSignAndSendTransactionRpcInputMethod = "signAndSendTransaction"
)

// The properties Encoding, Transaction are required.
type SolanaSignAndSendTransactionRpcInputParams struct {
	// Any of "base64".
	Encoding    string `json:"encoding,omitzero" api:"required"`
	Transaction string `json:"transaction" api:"required"`
	paramObj
}

func (r SolanaSignAndSendTransactionRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignAndSendTransactionRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignAndSendTransactionRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SolanaSignAndSendTransactionRpcInputParams](
		"encoding", "base64",
	)
}

type SolanaSignAndSendTransactionRpcInputChainType string

const (
	SolanaSignAndSendTransactionRpcInputChainTypeSolana SolanaSignAndSendTransactionRpcInputChainType = "solana"
)

// Executes the SVM `signMessage` RPC to sign a message.
//
// The properties Method, Params are required.
type SolanaSignMessageRpcInput struct {
	// Any of "signMessage".
	Method  SolanaSignMessageRpcInputMethod `json:"method,omitzero" api:"required"`
	Params  SolanaSignMessageRpcInputParams `json:"params,omitzero" api:"required"`
	Address param.Opt[string]               `json:"address,omitzero"`
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

type SolanaSignMessageRpcInputMethod string

const (
	SolanaSignMessageRpcInputMethodSignMessage SolanaSignMessageRpcInputMethod = "signMessage"
)

// The properties Encoding, Message are required.
type SolanaSignMessageRpcInputParams struct {
	// Any of "base64".
	Encoding string `json:"encoding,omitzero" api:"required"`
	Message  string `json:"message" api:"required"`
	paramObj
}

func (r SolanaSignMessageRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignMessageRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignMessageRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SolanaSignMessageRpcInputParams](
		"encoding", "base64",
	)
}

type SolanaSignMessageRpcInputChainType string

const (
	SolanaSignMessageRpcInputChainTypeSolana SolanaSignMessageRpcInputChainType = "solana"
)

// Response to the EVM `eth_signTransaction` RPC.
type EthereumSignTransactionRpcResponse struct {
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

type EthereumSignTransactionRpcResponseData struct {
	// Any of "rlp".
	Encoding          string `json:"encoding" api:"required"`
	SignedTransaction string `json:"signed_transaction" api:"required"`
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

type EthereumSignTransactionRpcResponseMethod string

const (
	EthereumSignTransactionRpcResponseMethodEthSignTransaction EthereumSignTransactionRpcResponseMethod = "eth_signTransaction"
)

// Response to the EVM `eth_sendTransaction` RPC.
type EthereumSendTransactionRpcResponse struct {
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

type EthereumSendTransactionRpcResponseData struct {
	Caip2              string                                                   `json:"caip2" api:"required"`
	Hash               string                                                   `json:"hash" api:"required"`
	TransactionID      string                                                   `json:"transaction_id"`
	TransactionRequest EthereumSendTransactionRpcResponseDataTransactionRequest `json:"transaction_request"`
	UserOperationHash  string                                                   `json:"user_operation_hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2              respjson.Field
		Hash               respjson.Field
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

type EthereumSendTransactionRpcResponseDataTransactionRequest struct {
	AuthorizationList    []EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationList       `json:"authorization_list"`
	ChainID              EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion              `json:"chain_id"`
	Data                 string                                                                            `json:"data"`
	From                 string                                                                            `json:"from"`
	GasLimit             EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion             `json:"gas_limit"`
	GasPrice             EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion             `json:"gas_price"`
	MaxFeePerGas         EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion         `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas"`
	Nonce                EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion                `json:"nonce"`
	To                   string                                                                            `json:"to"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                            `json:"type"`
	Value EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion `json:"value"`
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
func (r EthereumSendTransactionRpcResponseDataTransactionRequest) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcResponseDataTransactionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationList struct {
	ChainID  EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion `json:"chain_id" api:"required"`
	Contract string                                                                                `json:"contract" api:"required"`
	Nonce    EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion   `json:"nonce" api:"required"`
	R        string                                                                                `json:"r" api:"required"`
	S        string                                                                                `json:"s" api:"required"`
	YParity  float64                                                                               `json:"y_parity" api:"required"`
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
func (r EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationList) RawJSON() string {
	return r.JSON.raw
}
func (r *EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcResponseDataTransactionRequestValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendTransactionRpcResponseMethod string

const (
	EthereumSendTransactionRpcResponseMethodEthSendTransaction EthereumSendTransactionRpcResponseMethod = "eth_sendTransaction"
)

// Response to the EVM `personal_sign` RPC.
type EthereumPersonalSignRpcResponse struct {
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

type EthereumPersonalSignRpcResponseData struct {
	// Any of "hex".
	Encoding  string `json:"encoding" api:"required"`
	Signature string `json:"signature" api:"required"`
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

type EthereumPersonalSignRpcResponseMethod string

const (
	EthereumPersonalSignRpcResponseMethodPersonalSign EthereumPersonalSignRpcResponseMethod = "personal_sign"
)

// Response to the EVM `eth_signTypedData_v4` RPC.
type EthereumSignTypedDataRpcResponse struct {
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

type EthereumSignTypedDataRpcResponseData struct {
	// Any of "hex".
	Encoding  string `json:"encoding" api:"required"`
	Signature string `json:"signature" api:"required"`
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

type EthereumSignTypedDataRpcResponseMethod string

const (
	EthereumSignTypedDataRpcResponseMethodEthSignTypedDataV4 EthereumSignTypedDataRpcResponseMethod = "eth_signTypedData_v4"
)

// Response to the EVM `eth_signUserOperation` RPC.
type EthereumSignUserOperationRpcResponse struct {
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

type EthereumSignUserOperationRpcResponseData struct {
	// Any of "hex".
	Encoding  string `json:"encoding" api:"required"`
	Signature string `json:"signature" api:"required"`
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

type EthereumSignUserOperationRpcResponseMethod string

const (
	EthereumSignUserOperationRpcResponseMethodEthSignUserOperation EthereumSignUserOperationRpcResponseMethod = "eth_signUserOperation"
)

// Response to the EVM `eth_sign7702Authorization` RPC.
type EthereumSign7702AuthorizationRpcResponse struct {
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

type EthereumSign7702AuthorizationRpcResponseData struct {
	Authorization EthereumSign7702AuthorizationRpcResponseDataAuthorization `json:"authorization" api:"required"`
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

type EthereumSign7702AuthorizationRpcResponseDataAuthorization struct {
	ChainID  EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion `json:"chain_id" api:"required"`
	Contract string                                                                `json:"contract" api:"required"`
	Nonce    EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion   `json:"nonce" api:"required"`
	R        string                                                                `json:"r" api:"required"`
	S        string                                                                `json:"s" api:"required"`
	YParity  float64                                                               `json:"y_parity" api:"required"`
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
func (r EthereumSign7702AuthorizationRpcResponseDataAuthorization) RawJSON() string {
	return r.JSON.raw
}
func (r *EthereumSign7702AuthorizationRpcResponseDataAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSign7702AuthorizationRpcResponseMethod string

const (
	EthereumSign7702AuthorizationRpcResponseMethodEthSign7702Authorization EthereumSign7702AuthorizationRpcResponseMethod = "eth_sign7702Authorization"
)

// Response to the EVM `secp256k1_sign` RPC.
type EthereumSecp256k1SignRpcResponse struct {
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

type EthereumSecp256k1SignRpcResponseData struct {
	// Any of "hex".
	Encoding  string `json:"encoding" api:"required"`
	Signature string `json:"signature" api:"required"`
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

type EthereumSecp256k1SignRpcResponseMethod string

const (
	EthereumSecp256k1SignRpcResponseMethodSecp256k1Sign EthereumSecp256k1SignRpcResponseMethod = "secp256k1_sign"
)

// Response to the SVM `signTransaction` RPC.
type SolanaSignTransactionRpcResponse struct {
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

type SolanaSignTransactionRpcResponseData struct {
	// Any of "base64".
	Encoding          string `json:"encoding" api:"required"`
	SignedTransaction string `json:"signed_transaction" api:"required"`
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

type SolanaSignTransactionRpcResponseMethod string

const (
	SolanaSignTransactionRpcResponseMethodSignTransaction SolanaSignTransactionRpcResponseMethod = "signTransaction"
)

// Response to the SVM `signAndSendTransaction` RPC.
type SolanaSignAndSendTransactionRpcResponse struct {
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

type SolanaSignAndSendTransactionRpcResponseData struct {
	Caip2         string `json:"caip2" api:"required"`
	Hash          string `json:"hash" api:"required"`
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2         respjson.Field
		Hash          respjson.Field
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

type SolanaSignAndSendTransactionRpcResponseMethod string

const (
	SolanaSignAndSendTransactionRpcResponseMethodSignAndSendTransaction SolanaSignAndSendTransactionRpcResponseMethod = "signAndSendTransaction"
)

// Response to the SVM `signMessage` RPC.
type SolanaSignMessageRpcResponse struct {
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

type SolanaSignMessageRpcResponseData struct {
	// Any of "base64".
	Encoding  string `json:"encoding" api:"required"`
	Signature string `json:"signature" api:"required"`
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

type SolanaSignMessageRpcResponseMethod string

const (
	SolanaSignMessageRpcResponseMethodSignMessage SolanaSignMessageRpcResponseMethod = "signMessage"
)

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

func WalletRpcRequestBodyOfEthSignTransaction(params EthereumSignTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSignTransaction EthereumSignTransactionRpcInput
	ethSignTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfEthSignTransaction: &ethSignTransaction}
}

func WalletRpcRequestBodyOfEthSignUserOperation(params EthereumSignUserOperationRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSignUserOperation EthereumSignUserOperationRpcInput
	ethSignUserOperation.Params = params
	return WalletRpcRequestBodyUnion{OfEthSignUserOperation: &ethSignUserOperation}
}

func WalletRpcRequestBodyOfEthSendTransaction(caip2 string, method EthereumSendTransactionRpcInputMethod, params EthereumSendTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSendTransaction EthereumSendTransactionRpcInput
	ethSendTransaction.Caip2 = caip2
	ethSendTransaction.Method = method
	ethSendTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfEthSendTransaction: &ethSendTransaction}
}

func WalletRpcRequestBodyOfEthSign7702Authorization(params EthereumSign7702AuthorizationRpcInputParams) WalletRpcRequestBodyUnion {
	var ethSign7702Authorization EthereumSign7702AuthorizationRpcInput
	ethSign7702Authorization.Params = params
	return WalletRpcRequestBodyUnion{OfEthSign7702Authorization: &ethSign7702Authorization}
}

func WalletRpcRequestBodyOfSecp256k1Sign(params EthereumSecp256k1SignRpcInputParams) WalletRpcRequestBodyUnion {
	var secp256k1Sign EthereumSecp256k1SignRpcInput
	secp256k1Sign.Params = params
	return WalletRpcRequestBodyUnion{OfSecp256k1Sign: &secp256k1Sign}
}

func WalletRpcRequestBodyOfSignMessage(params SolanaSignMessageRpcInputParams) WalletRpcRequestBodyUnion {
	var signMessage SolanaSignMessageRpcInput
	signMessage.Params = params
	return WalletRpcRequestBodyUnion{OfSignMessage: &signMessage}
}

func WalletRpcRequestBodyOfSignTransaction(params SolanaSignTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var signTransaction SolanaSignTransactionRpcInput
	signTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfSignTransaction: &signTransaction}
}

func WalletRpcRequestBodyOfSignAndSendTransaction(caip2 string, method SolanaSignAndSendTransactionRpcInputMethod, params SolanaSignAndSendTransactionRpcInputParams) WalletRpcRequestBodyUnion {
	var signAndSendTransaction SolanaSignAndSendTransactionRpcInput
	signAndSendTransaction.Caip2 = caip2
	signAndSendTransaction.Method = method
	signAndSendTransaction.Params = params
	return WalletRpcRequestBodyUnion{OfSignAndSendTransaction: &signAndSendTransaction}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRpcRequestBodyUnion struct {
	OfPersonalSign             *EthereumPersonalSignRpcInput          `json:",omitzero,inline"`
	OfEthSignTypedDataV4       *EthereumSignTypedDataRpcInput         `json:",omitzero,inline"`
	OfEthSignTransaction       *EthereumSignTransactionRpcInput       `json:",omitzero,inline"`
	OfEthSignUserOperation     *EthereumSignUserOperationRpcInput     `json:",omitzero,inline"`
	OfEthSendTransaction       *EthereumSendTransactionRpcInput       `json:",omitzero,inline"`
	OfEthSign7702Authorization *EthereumSign7702AuthorizationRpcInput `json:",omitzero,inline"`
	OfSecp256k1Sign            *EthereumSecp256k1SignRpcInput         `json:",omitzero,inline"`
	OfSignMessage              *SolanaSignMessageRpcInput             `json:",omitzero,inline"`
	OfSignTransaction          *SolanaSignTransactionRpcInput         `json:",omitzero,inline"`
	OfSignAndSendTransaction   *SolanaSignAndSendTransactionRpcInput  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcRequestBodyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPersonalSign,
		u.OfEthSignTypedDataV4,
		u.OfEthSignTransaction,
		u.OfEthSignUserOperation,
		u.OfEthSendTransaction,
		u.OfEthSign7702Authorization,
		u.OfSecp256k1Sign,
		u.OfSignMessage,
		u.OfSignTransaction,
		u.OfSignAndSendTransaction)
}
func (u *WalletRpcRequestBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WalletRpcRequestBodyUnion](
		"method",
		apijson.Discriminator[EthereumPersonalSignRpcInput]("personal_sign"),
		apijson.Discriminator[EthereumSignTypedDataRpcInput]("eth_signTypedData_v4"),
		apijson.Discriminator[EthereumSignTransactionRpcInput]("eth_signTransaction"),
		apijson.Discriminator[EthereumSignUserOperationRpcInput]("eth_signUserOperation"),
		apijson.Discriminator[EthereumSendTransactionRpcInput]("eth_sendTransaction"),
		apijson.Discriminator[EthereumSign7702AuthorizationRpcInput]("eth_sign7702Authorization"),
		apijson.Discriminator[EthereumSecp256k1SignRpcInput]("secp256k1_sign"),
		apijson.Discriminator[SolanaSignMessageRpcInput]("signMessage"),
		apijson.Discriminator[SolanaSignTransactionRpcInput]("signTransaction"),
		apijson.Discriminator[SolanaSignAndSendTransactionRpcInput]("signAndSendTransaction"),
	)
}

// WalletRpcResponseBodyUnion contains all possible properties and values from
// [EthereumPersonalSignRpcResponse], [EthereumSignTypedDataRpcResponse],
// [EthereumSignTransactionRpcResponse], [EthereumSendTransactionRpcResponse],
// [EthereumSignUserOperationRpcResponse],
// [EthereumSign7702AuthorizationRpcResponse], [EthereumSecp256k1SignRpcResponse],
// [SolanaSignMessageRpcResponse], [SolanaSignTransactionRpcResponse],
// [SolanaSignAndSendTransactionRpcResponse].
//
// Use the [WalletRpcResponseBodyUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletRpcResponseBodyUnion struct {
	// This field is a union of [EthereumPersonalSignRpcResponseData],
	// [EthereumSignTypedDataRpcResponseData],
	// [EthereumSignTransactionRpcResponseData],
	// [EthereumSendTransactionRpcResponseData],
	// [EthereumSignUserOperationRpcResponseData],
	// [EthereumSign7702AuthorizationRpcResponseData],
	// [EthereumSecp256k1SignRpcResponseData], [SolanaSignMessageRpcResponseData],
	// [SolanaSignTransactionRpcResponseData],
	// [SolanaSignAndSendTransactionRpcResponseData]
	Data WalletRpcResponseBodyUnionData `json:"data"`
	// Any of "personal_sign", "eth_signTypedData_v4", "eth_signTransaction",
	// "eth_sendTransaction", "eth_signUserOperation", "eth_sign7702Authorization",
	// "secp256k1_sign", "signMessage", "signTransaction", "signAndSendTransaction".
	Method string `json:"method"`
	JSON   struct {
		Data   respjson.Field
		Method respjson.Field
		raw    string
	} `json:"-"`
}

// anyWalletRpcResponseBody is implemented by each variant of
// [WalletRpcResponseBodyUnion] to add type safety for the return type of
// [WalletRpcResponseBodyUnion.AsAny]
type anyWalletRpcResponseBody interface {
	implWalletRpcResponseBodyUnion()
}

func (EthereumPersonalSignRpcResponse) implWalletRpcResponseBodyUnion()          {}
func (EthereumSignTypedDataRpcResponse) implWalletRpcResponseBodyUnion()         {}
func (EthereumSignTransactionRpcResponse) implWalletRpcResponseBodyUnion()       {}
func (EthereumSendTransactionRpcResponse) implWalletRpcResponseBodyUnion()       {}
func (EthereumSignUserOperationRpcResponse) implWalletRpcResponseBodyUnion()     {}
func (EthereumSign7702AuthorizationRpcResponse) implWalletRpcResponseBodyUnion() {}
func (EthereumSecp256k1SignRpcResponse) implWalletRpcResponseBodyUnion()         {}
func (SolanaSignMessageRpcResponse) implWalletRpcResponseBodyUnion()             {}
func (SolanaSignTransactionRpcResponse) implWalletRpcResponseBodyUnion()         {}
func (SolanaSignAndSendTransactionRpcResponse) implWalletRpcResponseBodyUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletRpcResponseBodyUnion.AsAny().(type) {
//	case privyclient.EthereumPersonalSignRpcResponse:
//	case privyclient.EthereumSignTypedDataRpcResponse:
//	case privyclient.EthereumSignTransactionRpcResponse:
//	case privyclient.EthereumSendTransactionRpcResponse:
//	case privyclient.EthereumSignUserOperationRpcResponse:
//	case privyclient.EthereumSign7702AuthorizationRpcResponse:
//	case privyclient.EthereumSecp256k1SignRpcResponse:
//	case privyclient.SolanaSignMessageRpcResponse:
//	case privyclient.SolanaSignTransactionRpcResponse:
//	case privyclient.SolanaSignAndSendTransactionRpcResponse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletRpcResponseBodyUnion) AsAny() anyWalletRpcResponseBody {
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
	case "signMessage":
		return u.AsSignMessage()
	case "signTransaction":
		return u.AsSignTransaction()
	case "signAndSendTransaction":
		return u.AsSignAndSendTransaction()
	}
	return nil
}

func (u WalletRpcResponseBodyUnion) AsPersonalSign() (v EthereumPersonalSignRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsEthSignTypedDataV4() (v EthereumSignTypedDataRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsEthSignTransaction() (v EthereumSignTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsEthSendTransaction() (v EthereumSendTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsEthSignUserOperation() (v EthereumSignUserOperationRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsEthSign7702Authorization() (v EthereumSign7702AuthorizationRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsSecp256k1Sign() (v EthereumSecp256k1SignRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsSignMessage() (v SolanaSignMessageRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsSignTransaction() (v SolanaSignTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseBodyUnion) AsSignAndSendTransaction() (v SolanaSignAndSendTransactionRpcResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseBodyUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletRpcResponseBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseBodyUnionData is an implicit subunion of
// [WalletRpcResponseBodyUnion]. WalletRpcResponseBodyUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcResponseBodyUnion].
type WalletRpcResponseBodyUnionData struct {
	Encoding          string `json:"encoding"`
	Signature         string `json:"signature"`
	SignedTransaction string `json:"signed_transaction"`
	Caip2             string `json:"caip2"`
	Hash              string `json:"hash"`
	TransactionID     string `json:"transaction_id"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	TransactionRequest EthereumSendTransactionRpcResponseDataTransactionRequest `json:"transaction_request"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	UserOperationHash string `json:"user_operation_hash"`
	// This field is from variant [EthereumSign7702AuthorizationRpcResponseData].
	Authorization EthereumSign7702AuthorizationRpcResponseDataAuthorization `json:"authorization"`
	JSON          struct {
		Encoding           respjson.Field
		Signature          respjson.Field
		SignedTransaction  respjson.Field
		Caip2              respjson.Field
		Hash               respjson.Field
		TransactionID      respjson.Field
		TransactionRequest respjson.Field
		UserOperationHash  respjson.Field
		Authorization      respjson.Field
		raw                string
	} `json:"-"`
}

func (r *WalletRpcResponseBodyUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletInitImportResponse struct {
	// The base64-encoded encryption public key to encrypt the wallet entropy with.
	EncryptionPublicKey string `json:"encryption_public_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType WalletInitImportResponseEncryptionType `json:"encryption_type" api:"required"`
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

// The encryption type of the wallet to import. Currently only supports `HPKE`.
type WalletInitImportResponseEncryptionType string

const (
	WalletInitImportResponseEncryptionTypeHpke WalletInitImportResponseEncryptionType = "HPKE"
)

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
	// The expiration time of the authorization key in seconds since the epoch.
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
	// The expiration time of the authorization key in seconds since the epoch.
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

type WalletExportResponse struct {
	// The encrypted private key.
	Ciphertext string `json:"ciphertext" api:"required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType WalletExportResponseEncryptionType `json:"encryption_type" api:"required"`
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
func (r WalletExportResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encryption type of the wallet to import. Currently only supports `HPKE`.
type WalletExportResponseEncryptionType string

const (
	WalletExportResponseEncryptionTypeHpke WalletExportResponseEncryptionType = "HPKE"
)

type WalletRawSignResponse struct {
	Data WalletRawSignResponseData `json:"data" api:"required"`
	// Any of "raw_sign".
	Method WalletRawSignResponseMethod `json:"method" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRawSignResponse) RawJSON() string { return r.JSON.raw }
func (r *WalletRawSignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRawSignResponseData struct {
	// Any of "hex".
	Encoding  string `json:"encoding" api:"required"`
	Signature string `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Encoding    respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRawSignResponseData) RawJSON() string { return r.JSON.raw }
func (r *WalletRawSignResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRawSignResponseMethod string

const (
	WalletRawSignResponseMethodRawSign WalletRawSignResponseMethod = "raw_sign"
)

type WalletNewParams struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero" api:"required"`
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner WalletNewParamsOwnerUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletNewParamsAdditionalSigner `json:"additional_signers,omitzero"`
	// List of policy IDs for policies that should be enforced on the wallet.
	// Currently, only one policy is supported per wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r WalletNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SignerID is required.
type WalletNewParamsAdditionalSigner struct {
	SignerID string `json:"signer_id" api:"required" format:"cuid2"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet.
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r WalletNewParamsAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow WalletNewParamsAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletNewParamsAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletNewParamsOwnerUnion struct {
	OfPublicKeyOwner *WalletNewParamsOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *WalletNewParamsOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u WalletNewParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *WalletNewParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type WalletNewParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r WalletNewParamsOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow WalletNewParamsOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletNewParamsOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type WalletNewParamsOwnerUserOwner struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r WalletNewParamsOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow WalletNewParamsOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletNewParamsOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletUpdateParams struct {
	// Request body for updating a wallet.
	PatchWalletRequestBody PatchWalletRequestBody
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r WalletUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PatchWalletRequestBody)
}
func (r *WalletUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PatchWalletRequestBody)
}

type WalletListParams struct {
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Filter wallets by authorization public key. Returns wallets owned by key quorums
	// that include the specified P-256 public key (base64-encoded DER format). Cannot
	// be used together with user_id.
	AuthorizationKey param.Opt[string] `query:"authorization_key,omitzero" json:"-"`
	Cursor           param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter wallets by user ID. Cannot be used together with authorization_key.
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
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
	OfHD *WalletInitImportParamsBodyHD `json:",inline"`
	// This field is a request body variant, only one variant field can be set. The
	// input for private key wallets.
	OfPrivateKey *WalletInitImportParamsBodyPrivateKey `json:",inline"`

	paramObj
}

func (u WalletInitImportParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHD, u.OfPrivateKey)
}
func (r *WalletInitImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The input for HD wallets.
//
// The properties Address, ChainType, EncryptionType, EntropyType, Index are
// required.
type WalletInitImportParamsBodyHD struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero" api:"required"`
	// The index of the wallet to import.
	Index int64 `json:"index" api:"required"`
	// The entropy type of the wallet to import.
	//
	// This field can be elided, and will marshal its zero value as "hd".
	EntropyType constant.HD `json:"entropy_type" api:"required"`
	paramObj
}

func (r WalletInitImportParamsBodyHD) MarshalJSON() (data []byte, err error) {
	type shadow WalletInitImportParamsBodyHD
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletInitImportParamsBodyHD) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletInitImportParamsBodyHD](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[WalletInitImportParamsBodyHD](
		"encryption_type", "HPKE",
	)
}

// The input for private key wallets.
//
// The properties Address, ChainType, EncryptionType, EntropyType are required.
type WalletInitImportParamsBodyPrivateKey struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "private-key".
	EntropyType constant.PrivateKey `json:"entropy_type" api:"required"`
	paramObj
}

func (r WalletInitImportParamsBodyPrivateKey) MarshalJSON() (data []byte, err error) {
	type shadow WalletInitImportParamsBodyPrivateKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletInitImportParamsBodyPrivateKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletInitImportParamsBodyPrivateKey](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[WalletInitImportParamsBodyPrivateKey](
		"encryption_type", "HPKE",
	)
}

type WalletSubmitImportParams struct {
	Wallet            WalletSubmitImportParamsWalletUnion        `json:"wallet,omitzero" api:"required"`
	OwnerID           param.Opt[string]                          `json:"owner_id,omitzero" format:"cuid2"`
	Owner             WalletSubmitImportParamsOwnerUnion         `json:"owner,omitzero"`
	AdditionalSigners []WalletSubmitImportParamsAdditionalSigner `json:"additional_signers,omitzero"`
	PolicyIDs         []string                                   `json:"policy_ids,omitzero" format:"cuid2"`
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
	OfHD         *WalletSubmitImportParamsWalletHD         `json:",omitzero,inline"`
	OfPrivateKey *WalletSubmitImportParamsWalletPrivateKey `json:",omitzero,inline"`
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
		apijson.Discriminator[WalletSubmitImportParamsWalletHD]("hd"),
		apijson.Discriminator[WalletSubmitImportParamsWalletPrivateKey]("private-key"),
	)
}

// The properties Address, ChainType, Ciphertext, EncapsulatedKey, EncryptionType,
// EntropyType, Index are required.
type WalletSubmitImportParamsWalletHD struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero" api:"required"`
	// The encrypted entropy of the wallet to import.
	Ciphertext string `json:"ciphertext" api:"required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption inside the TEE.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero" api:"required"`
	// The index of the wallet to import.
	Index int64 `json:"index" api:"required"`
	// Optional HPKE configuration for wallet import decryption. These parameters allow
	// importing wallets encrypted by external providers that use different HPKE
	// configurations.
	HpkeConfig HpkeImportConfig `json:"hpke_config,omitzero"`
	// The entropy type of the wallet to import.
	//
	// This field can be elided, and will marshal its zero value as "hd".
	EntropyType constant.HD `json:"entropy_type" api:"required"`
	paramObj
}

func (r WalletSubmitImportParamsWalletHD) MarshalJSON() (data []byte, err error) {
	type shadow WalletSubmitImportParamsWalletHD
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletSubmitImportParamsWalletHD) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletSubmitImportParamsWalletHD](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[WalletSubmitImportParamsWalletHD](
		"encryption_type", "HPKE",
	)
}

// The properties Address, ChainType, Ciphertext, EncapsulatedKey, EncryptionType,
// EntropyType are required.
type WalletSubmitImportParamsWalletPrivateKey struct {
	// The address of the wallet to import.
	Address string `json:"address" api:"required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero" api:"required"`
	// The encrypted entropy of the wallet to import.
	Ciphertext string `json:"ciphertext" api:"required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption inside the TEE.
	EncapsulatedKey string `json:"encapsulated_key" api:"required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero" api:"required"`
	// Optional HPKE configuration for wallet import decryption. These parameters allow
	// importing wallets encrypted by external providers that use different HPKE
	// configurations.
	HpkeConfig HpkeImportConfig `json:"hpke_config,omitzero"`
	// This field can be elided, and will marshal its zero value as "private-key".
	EntropyType constant.PrivateKey `json:"entropy_type" api:"required"`
	paramObj
}

func (r WalletSubmitImportParamsWalletPrivateKey) MarshalJSON() (data []byte, err error) {
	type shadow WalletSubmitImportParamsWalletPrivateKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletSubmitImportParamsWalletPrivateKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletSubmitImportParamsWalletPrivateKey](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[WalletSubmitImportParamsWalletPrivateKey](
		"encryption_type", "HPKE",
	)
}

// The property SignerID is required.
type WalletSubmitImportParamsAdditionalSigner struct {
	SignerID          string   `json:"signer_id" api:"required" format:"cuid2"`
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r WalletSubmitImportParamsAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow WalletSubmitImportParamsAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletSubmitImportParamsAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletSubmitImportParamsOwnerUnion struct {
	OfWalletSubmitImportsOwnerUserID    *WalletSubmitImportParamsOwnerUserID    `json:",omitzero,inline"`
	OfWalletSubmitImportsOwnerPublicKey *WalletSubmitImportParamsOwnerPublicKey `json:",omitzero,inline"`
	paramUnion
}

func (u WalletSubmitImportParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWalletSubmitImportsOwnerUserID, u.OfWalletSubmitImportsOwnerPublicKey)
}
func (u *WalletSubmitImportParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property UserID is required.
type WalletSubmitImportParamsOwnerUserID struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r WalletSubmitImportParamsOwnerUserID) MarshalJSON() (data []byte, err error) {
	type shadow WalletSubmitImportParamsOwnerUserID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletSubmitImportParamsOwnerUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PublicKey is required.
type WalletSubmitImportParamsOwnerPublicKey struct {
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r WalletSubmitImportParamsOwnerPublicKey) MarshalJSON() (data []byte, err error) {
	type shadow WalletSubmitImportParamsOwnerPublicKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletSubmitImportParamsOwnerPublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAuthenticateWithJwtParams struct {
	// The user's JWT, to be used to authenticate the user.
	UserJwt string `json:"user_jwt" api:"required"`
	// The public key of your ECDH keypair, in base64-encoded, SPKI-format, whose
	// private key will be able to decrypt the session key.
	RecipientPublicKey param.Opt[string] `json:"recipient_public_key,omitzero"`
	// The encryption type for the authentication response. Currently only supports
	// HPKE.
	//
	// Any of "HPKE".
	EncryptionType WalletAuthenticateWithJwtParamsEncryptionType `json:"encryption_type,omitzero"`
	paramObj
}

func (r WalletAuthenticateWithJwtParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletAuthenticateWithJwtParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletAuthenticateWithJwtParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encryption type for the authentication response. Currently only supports
// HPKE.
type WalletAuthenticateWithJwtParamsEncryptionType string

const (
	WalletAuthenticateWithJwtParamsEncryptionTypeHpke WalletAuthenticateWithJwtParamsEncryptionType = "HPKE"
)

type WalletExportParams struct {
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType WalletExportParamsEncryptionType `json:"encryption_type,omitzero" api:"required"`
	// The base64-encoded encryption public key to encrypt the wallet private key with.
	RecipientPublicKey string `json:"recipient_public_key" api:"required"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r WalletExportParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletExportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletExportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encryption type of the wallet to import. Currently only supports `HPKE`.
type WalletExportParamsEncryptionType string

const (
	WalletExportParamsEncryptionTypeHpke WalletExportParamsEncryptionType = "HPKE"
)

type WalletRawSignParams struct {
	// Sign a pre-computed hash
	Params WalletRawSignParamsParamsUnion `json:"params,omitzero" api:"required"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	paramObj
}

func (r WalletRawSignParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRawSignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRawSignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRawSignParamsParamsUnion struct {
	OfHash  *WalletRawSignParamsParamsHash  `json:",omitzero,inline"`
	OfBytes *WalletRawSignParamsParamsBytes `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRawSignParamsParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHash, u.OfBytes)
}
func (u *WalletRawSignParamsParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Sign a pre-computed hash
//
// The property Hash is required.
type WalletRawSignParamsParamsHash struct {
	// The hash to sign. Must start with `0x`.
	Hash string `json:"hash" api:"required"`
	paramObj
}

func (r WalletRawSignParamsParamsHash) MarshalJSON() (data []byte, err error) {
	type shadow WalletRawSignParamsParamsHash
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRawSignParamsParamsHash) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hash and sign bytes using the specified encoding and hash function.
//
// The properties Bytes, Encoding, HashFunction are required.
type WalletRawSignParamsParamsBytes struct {
	// The bytes to hash and sign.
	Bytes string `json:"bytes" api:"required"`
	// The encoding scheme for the bytes.
	//
	// Any of "utf-8", "hex", "base64".
	Encoding string `json:"encoding,omitzero" api:"required"`
	// The hash function to hash the bytes.
	//
	// Any of "keccak256", "sha256", "blake2b256".
	HashFunction string `json:"hash_function,omitzero" api:"required"`
	paramObj
}

func (r WalletRawSignParamsParamsBytes) MarshalJSON() (data []byte, err error) {
	type shadow WalletRawSignParamsParamsBytes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRawSignParamsParamsBytes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRawSignParamsParamsBytes](
		"encoding", "utf-8", "hex", "base64",
	)
	apijson.RegisterFieldValidator[WalletRawSignParamsParamsBytes](
		"hash_function", "keccak256", "sha256", "blake2b256",
	)
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
	paramObj
}

func (r WalletRpcParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletRpcRequestBody)
}
func (r *WalletRpcParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WalletRpcRequestBody)
}
