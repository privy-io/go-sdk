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

// Export a wallet's private key
func (r *WalletService) Export(ctx context.Context, walletID string, params WalletExportParams, opts ...option.RequestOption) (res *WalletExportResponse, err error) {
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

// SUI transaction commands allowlist for raw_sign endpoint policy evaluation
type SuiCommandName string

const (
	SuiCommandNameTransferObjects SuiCommandName = "TransferObjects"
	SuiCommandNameSplitCoins      SuiCommandName = "SplitCoins"
	SuiCommandNameMergeCoins      SuiCommandName = "MergeCoins"
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

// Exports the private key of the wallet.
type ExportPrivateKeyRpcInput struct {
	Address string `json:"address" api:"required"`
	// Any of "exportPrivateKey".
	Method ExportPrivateKeyRpcInputMethod `json:"method" api:"required"`
	Params ExportPrivateKeyRpcInputParams `json:"params" api:"required"`
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
func (r ExportPrivateKeyRpcInput) RawJSON() string { return r.JSON.raw }
func (r *ExportPrivateKeyRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExportPrivateKeyRpcInput to a
// ExportPrivateKeyRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExportPrivateKeyRpcInputParam.Overrides()
func (r ExportPrivateKeyRpcInput) ToParam() ExportPrivateKeyRpcInputParam {
	return param.Override[ExportPrivateKeyRpcInputParam](json.RawMessage(r.RawJSON()))
}

type ExportPrivateKeyRpcInputMethod string

const (
	ExportPrivateKeyRpcInputMethodExportPrivateKey ExportPrivateKeyRpcInputMethod = "exportPrivateKey"
)

type ExportPrivateKeyRpcInputParams struct {
	// Any of "HPKE".
	EncryptionType     string `json:"encryption_type" api:"required"`
	RecipientPublicKey string `json:"recipient_public_key" api:"required"`
	// Any of "display", "client".
	ExportType string `json:"export_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptionType     respjson.Field
		RecipientPublicKey respjson.Field
		ExportType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPrivateKeyRpcInputParams) RawJSON() string { return r.JSON.raw }
func (r *ExportPrivateKeyRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exports the private key of the wallet.
//
// The properties Address, Method, Params are required.
type ExportPrivateKeyRpcInputParam struct {
	Address string `json:"address" api:"required"`
	// Any of "exportPrivateKey".
	Method ExportPrivateKeyRpcInputMethod      `json:"method,omitzero" api:"required"`
	Params ExportPrivateKeyRpcInputParamsParam `json:"params,omitzero" api:"required"`
	paramObj
}

func (r ExportPrivateKeyRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPrivateKeyRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPrivateKeyRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EncryptionType, RecipientPublicKey are required.
type ExportPrivateKeyRpcInputParamsParam struct {
	// Any of "HPKE".
	EncryptionType     string `json:"encryption_type,omitzero" api:"required"`
	RecipientPublicKey string `json:"recipient_public_key" api:"required"`
	// Any of "display", "client".
	ExportType string `json:"export_type,omitzero"`
	paramObj
}

func (r ExportPrivateKeyRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPrivateKeyRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPrivateKeyRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExportPrivateKeyRpcInputParamsParam](
		"encryption_type", "HPKE",
	)
	apijson.RegisterFieldValidator[ExportPrivateKeyRpcInputParamsParam](
		"export_type", "display", "client",
	)
}

// Response to the `exportPrivateKey` RPC.
type ExportPrivateKeyRpcResponse struct {
	Data ExportPrivateKeyRpcResponseData `json:"data" api:"required"`
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

type ExportPrivateKeyRpcResponseData struct {
	// Any of "HPKE".
	EncryptionType     string `json:"encryption_type" api:"required"`
	RecipientPublicKey string `json:"recipient_public_key" api:"required"`
	// Any of "display", "client".
	ExportType string `json:"export_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptionType     respjson.Field
		RecipientPublicKey respjson.Field
		ExportType         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPrivateKeyRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExportPrivateKeyRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExportPrivateKeyRpcResponseMethod string

const (
	ExportPrivateKeyRpcResponseMethodExportPrivateKey ExportPrivateKeyRpcResponseMethod = "exportPrivateKey"
)

// Parameters for signing a pre-computed hash with the `raw_sign` RPC.
//
// The property Hash is required.
type RawSignHashParams struct {
	// The hash to sign.
	Hash string `json:"hash" api:"required"`
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

func RawSignInputParamsOfRawSignHashs(hash string) RawSignInputParamsUnion {
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
	Encoding  RawSignResponseDataEncoding `json:"encoding" api:"required"`
	Signature string                      `json:"signature" api:"required"`
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
type EthereumPersonalSignRpcInput struct {
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
func (r EthereumPersonalSignRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumPersonalSignRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumPersonalSignRpcInput to a
// EthereumPersonalSignRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumPersonalSignRpcInputParam.Overrides()
func (r EthereumPersonalSignRpcInput) ToParam() EthereumPersonalSignRpcInputParam {
	return param.Override[EthereumPersonalSignRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumPersonalSignRpcInputParam struct {
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

func (r EthereumPersonalSignRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumPersonalSignRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumPersonalSignRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signTransaction` RPC.
type EthereumSignTransactionRpcInputParamsResp struct {
	Transaction EthereumSignTransactionRpcInputParamsTransactionResp `json:"transaction" api:"required"`
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

type EthereumSignTransactionRpcInputParamsTransactionResp struct {
	AuthorizationList    []EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp       `json:"authorization_list"`
	ChainID              EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp              `json:"chain_id"`
	Data                 string                                                                        `json:"data"`
	From                 string                                                                        `json:"from"`
	GasLimit             EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp             `json:"gas_limit"`
	GasPrice             EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp             `json:"gas_price"`
	MaxFeePerGas         EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp         `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp `json:"max_priority_fee_per_gas"`
	Nonce                EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp                `json:"nonce"`
	To                   string                                                                        `json:"to"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                        `json:"type"`
	Value EthereumSignTransactionRpcInputParamsTransactionValueUnionResp `json:"value"`
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
func (r EthereumSignTransactionRpcInputParamsTransactionResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTransactionRpcInputParamsTransactionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp struct {
	ChainID  EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp `json:"chain_id" api:"required"`
	Contract string                                                                            `json:"contract" api:"required"`
	Nonce    EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp   `json:"nonce" api:"required"`
	R        string                                                                            `json:"r" api:"required"`
	S        string                                                                            `json:"s" api:"required"`
	YParity  float64                                                                           `json:"y_parity" api:"required"`
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
func (r EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp) RawJSON() string {
	return r.JSON.raw
}
func (r *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSignTransactionRpcInputParamsTransactionValueUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignTransactionRpcInputParamsTransactionValueUnionResp struct {
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

func (u EthereumSignTransactionRpcInputParamsTransactionValueUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignTransactionRpcInputParamsTransactionValueUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignTransactionRpcInputParamsTransactionValueUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSignTransactionRpcInputParamsTransactionValueUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signTransaction` RPC.
//
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

// Executes the EVM `eth_signTransaction` RPC to sign a transaction.
type EthereumSignTransactionRpcInput struct {
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
func (r EthereumSignTransactionRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignTransactionRpcInput to a
// EthereumSignTransactionRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignTransactionRpcInputParam.Overrides()
func (r EthereumSignTransactionRpcInput) ToParam() EthereumSignTransactionRpcInputParam {
	return param.Override[EthereumSignTransactionRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumSignTransactionRpcInputParam struct {
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

func (r EthereumSignTransactionRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_sendTransaction` RPC.
type EthereumSendTransactionRpcInputParamsResp struct {
	Transaction EthereumSendTransactionRpcInputParamsTransactionResp `json:"transaction" api:"required"`
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

type EthereumSendTransactionRpcInputParamsTransactionResp struct {
	AuthorizationList    []EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp       `json:"authorization_list"`
	ChainID              EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp              `json:"chain_id"`
	Data                 string                                                                        `json:"data"`
	From                 string                                                                        `json:"from"`
	GasLimit             EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp             `json:"gas_limit"`
	GasPrice             EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp             `json:"gas_price"`
	MaxFeePerGas         EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp         `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp `json:"max_priority_fee_per_gas"`
	Nonce                EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp                `json:"nonce"`
	To                   string                                                                        `json:"to"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                        `json:"type"`
	Value EthereumSendTransactionRpcInputParamsTransactionValueUnionResp `json:"value"`
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
func (r EthereumSendTransactionRpcInputParamsTransactionResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcInputParamsTransactionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp struct {
	ChainID  EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp `json:"chain_id" api:"required"`
	Contract string                                                                            `json:"contract" api:"required"`
	Nonce    EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp   `json:"nonce" api:"required"`
	R        string                                                                            `json:"r" api:"required"`
	S        string                                                                            `json:"s" api:"required"`
	YParity  float64                                                                           `json:"y_parity" api:"required"`
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
func (r EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp) RawJSON() string {
	return r.JSON.raw
}
func (r *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EthereumSendTransactionRpcInputParamsTransactionValueUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSendTransactionRpcInputParamsTransactionValueUnionResp struct {
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

func (u EthereumSendTransactionRpcInputParamsTransactionValueUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSendTransactionRpcInputParamsTransactionValueUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSendTransactionRpcInputParamsTransactionValueUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSendTransactionRpcInputParamsTransactionValueUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_sendTransaction` RPC.
//
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

// Executes the EVM `eth_sendTransaction` RPC to sign and broadcast a transaction.
type EthereumSendTransactionRpcInput struct {
	Caip2 string `json:"caip2" api:"required"`
	// Any of "eth_sendTransaction".
	Method EthereumSendTransactionRpcInputMethod `json:"method" api:"required"`
	// Parameters for the EVM `eth_sendTransaction` RPC.
	Params  EthereumSendTransactionRpcInputParamsResp `json:"params" api:"required"`
	Address string                                    `json:"address"`
	// Any of "ethereum".
	ChainType EthereumSendTransactionRpcInputChainType `json:"chain_type"`
	Sponsor   bool                                     `json:"sponsor"`
	WalletID  string                                   `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2       respjson.Field
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		Sponsor     respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSendTransactionRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumSendTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSendTransactionRpcInput to a
// EthereumSendTransactionRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSendTransactionRpcInputParam.Overrides()
func (r EthereumSendTransactionRpcInput) ToParam() EthereumSendTransactionRpcInputParam {
	return param.Override[EthereumSendTransactionRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumSendTransactionRpcInputParam struct {
	Caip2 string `json:"caip2" api:"required"`
	// Any of "eth_sendTransaction".
	Method EthereumSendTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the EVM `eth_sendTransaction` RPC.
	Params   EthereumSendTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                     `json:"address,omitzero"`
	Sponsor  param.Opt[bool]                       `json:"sponsor,omitzero"`
	WalletID param.Opt[string]                     `json:"wallet_id,omitzero"`
	// Any of "ethereum".
	ChainType EthereumSendTransactionRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EIP-712 typed data object.
type EthereumTypedDataInput struct {
	Domain      map[string]any                          `json:"domain" api:"required"`
	Message     map[string]any                          `json:"message" api:"required"`
	PrimaryType string                                  `json:"primary_type" api:"required"`
	Types       map[string][]EthereumTypedDataInputType `json:"types" api:"required"`
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
func (r EthereumTypedDataInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumTypedDataInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumTypedDataInput to a EthereumTypedDataInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumTypedDataInputParam.Overrides()
func (r EthereumTypedDataInput) ToParam() EthereumTypedDataInputParam {
	return param.Override[EthereumTypedDataInputParam](json.RawMessage(r.RawJSON()))
}

type EthereumTypedDataInputType struct {
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
func (r EthereumTypedDataInputType) RawJSON() string { return r.JSON.raw }
func (r *EthereumTypedDataInputType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EIP-712 typed data object.
//
// The properties Domain, Message, PrimaryType, Types are required.
type EthereumTypedDataInputParam struct {
	Domain      map[string]any                               `json:"domain,omitzero" api:"required"`
	Message     map[string]any                               `json:"message,omitzero" api:"required"`
	PrimaryType string                                       `json:"primary_type" api:"required"`
	Types       map[string][]EthereumTypedDataInputTypeParam `json:"types,omitzero" api:"required"`
	paramObj
}

func (r EthereumTypedDataInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTypedDataInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTypedDataInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type EthereumTypedDataInputTypeParam struct {
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	paramObj
}

func (r EthereumTypedDataInputTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumTypedDataInputTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumTypedDataInputTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signTypedData_v4` RPC.
type EthereumSignTypedDataRpcInputParamsResp struct {
	// EIP-712 typed data object.
	TypedData EthereumTypedDataInput `json:"typed_data" api:"required"`
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
	TypedData EthereumTypedDataInputParam `json:"typed_data,omitzero" api:"required"`
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
type EthereumSignTypedDataRpcInput struct {
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
func (r EthereumSignTypedDataRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignTypedDataRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignTypedDataRpcInput to a
// EthereumSignTypedDataRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignTypedDataRpcInputParam.Overrides()
func (r EthereumSignTypedDataRpcInput) ToParam() EthereumSignTypedDataRpcInputParam {
	return param.Override[EthereumSignTypedDataRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumSignTypedDataRpcInputParam struct {
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

func (r EthereumSignTypedDataRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `secp256k1_sign` RPC.
type EthereumSecp256k1SignRpcInputParamsResp struct {
	Hash string `json:"hash" api:"required"`
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

// Signs a raw hash on the secp256k1 curve.
type EthereumSecp256k1SignRpcInput struct {
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
func (r EthereumSecp256k1SignRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumSecp256k1SignRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSecp256k1SignRpcInput to a
// EthereumSecp256k1SignRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSecp256k1SignRpcInputParam.Overrides()
func (r EthereumSecp256k1SignRpcInput) ToParam() EthereumSecp256k1SignRpcInputParam {
	return param.Override[EthereumSecp256k1SignRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumSecp256k1SignRpcInputParam struct {
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

func (r EthereumSecp256k1SignRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSecp256k1SignRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSecp256k1SignRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_sign7702Authorization` RPC.
type EthereumSign7702AuthorizationRpcInputParamsResp struct {
	ChainID  EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp `json:"chain_id" api:"required"`
	Contract string                                                      `json:"contract" api:"required"`
	// Any of "self".
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor       `json:"executor"`
	Nonce    EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp `json:"nonce"`
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

// EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp struct {
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

func (u EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSign7702AuthorizationRpcInputParamsExecutor string

const (
	EthereumSign7702AuthorizationRpcInputParamsExecutorSelf EthereumSign7702AuthorizationRpcInputParamsExecutor = "self"
)

// EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp struct {
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

func (u EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_sign7702Authorization` RPC.
//
// The properties ChainID, Contract are required.
type EthereumSign7702AuthorizationRpcInputParams struct {
	ChainID  EthereumSign7702AuthorizationRpcInputParamsChainIDUnion `json:"chain_id,omitzero" api:"required"`
	Contract string                                                  `json:"contract" api:"required"`
	// Any of "self".
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor   `json:"executor,omitzero"`
	Nonce    EthereumSign7702AuthorizationRpcInputParamsNonceUnion `json:"nonce,omitzero"`
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

// Signs an EIP-7702 authorization.
type EthereumSign7702AuthorizationRpcInput struct {
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
func (r EthereumSign7702AuthorizationRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumSign7702AuthorizationRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSign7702AuthorizationRpcInput to a
// EthereumSign7702AuthorizationRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSign7702AuthorizationRpcInputParam.Overrides()
func (r EthereumSign7702AuthorizationRpcInput) ToParam() EthereumSign7702AuthorizationRpcInputParam {
	return param.Override[EthereumSign7702AuthorizationRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumSign7702AuthorizationRpcInputParam struct {
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

func (r EthereumSign7702AuthorizationRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSign7702AuthorizationRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSign7702AuthorizationRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signUserOperation` RPC.
type EthereumSignUserOperationRpcInputParamsResp struct {
	ChainID       EthereumSignUserOperationRpcInputParamsChainIDUnionResp  `json:"chain_id" api:"required"`
	Contract      string                                                   `json:"contract" api:"required"`
	UserOperation EthereumSignUserOperationRpcInputParamsUserOperationResp `json:"user_operation" api:"required"`
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

// EthereumSignUserOperationRpcInputParamsChainIDUnionResp contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type EthereumSignUserOperationRpcInputParamsChainIDUnionResp struct {
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

func (u EthereumSignUserOperationRpcInputParamsChainIDUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EthereumSignUserOperationRpcInputParamsChainIDUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EthereumSignUserOperationRpcInputParamsChainIDUnionResp) RawJSON() string { return u.JSON.raw }

func (r *EthereumSignUserOperationRpcInputParamsChainIDUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignUserOperationRpcInputParamsUserOperationResp struct {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallData                      respjson.Field
		CallGasLimit                  respjson.Field
		MaxFeePerGas                  respjson.Field
		MaxPriorityFeePerGas          respjson.Field
		Nonce                         respjson.Field
		Paymaster                     respjson.Field
		PaymasterData                 respjson.Field
		PaymasterPostOpGasLimit       respjson.Field
		PaymasterVerificationGasLimit respjson.Field
		PreVerificationGas            respjson.Field
		Sender                        respjson.Field
		VerificationGasLimit          respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EthereumSignUserOperationRpcInputParamsUserOperationResp) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignUserOperationRpcInputParamsUserOperationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the EVM `eth_signUserOperation` RPC.
//
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

// Executes an RPC method to hash and sign a UserOperation.
type EthereumSignUserOperationRpcInput struct {
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
func (r EthereumSignUserOperationRpcInput) RawJSON() string { return r.JSON.raw }
func (r *EthereumSignUserOperationRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EthereumSignUserOperationRpcInput to a
// EthereumSignUserOperationRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EthereumSignUserOperationRpcInputParam.Overrides()
func (r EthereumSignUserOperationRpcInput) ToParam() EthereumSignUserOperationRpcInputParam {
	return param.Override[EthereumSignUserOperationRpcInputParam](json.RawMessage(r.RawJSON()))
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
type EthereumSignUserOperationRpcInputParam struct {
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

func (r EthereumSignUserOperationRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInputParam) UnmarshalJSON(data []byte) error {
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
	Encoding  EthereumSecp256k1SignRpcResponseDataEncoding `json:"encoding" api:"required"`
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
type SolanaSignTransactionRpcInput struct {
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
func (r SolanaSignTransactionRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignTransactionRpcInput to a
// SolanaSignTransactionRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignTransactionRpcInputParam.Overrides()
func (r SolanaSignTransactionRpcInput) ToParam() SolanaSignTransactionRpcInputParam {
	return param.Override[SolanaSignTransactionRpcInputParam](json.RawMessage(r.RawJSON()))
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
type SolanaSignTransactionRpcInputParam struct {
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

func (r SolanaSignTransactionRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignTransactionRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignTransactionRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SolanaSignAndSendTransactionRpcInput struct {
	Caip2 string `json:"caip2" api:"required"`
	// Any of "signAndSendTransaction".
	Method SolanaSignAndSendTransactionRpcInputMethod `json:"method" api:"required"`
	// Parameters for the SVM `signAndSendTransaction` RPC.
	Params  SolanaSignAndSendTransactionRpcInputParamsResp `json:"params" api:"required"`
	Address string                                         `json:"address"`
	// Any of "solana".
	ChainType SolanaSignAndSendTransactionRpcInputChainType `json:"chain_type"`
	Sponsor   bool                                          `json:"sponsor"`
	WalletID  string                                        `json:"wallet_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2       respjson.Field
		Method      respjson.Field
		Params      respjson.Field
		Address     respjson.Field
		ChainType   respjson.Field
		Sponsor     respjson.Field
		WalletID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolanaSignAndSendTransactionRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignAndSendTransactionRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignAndSendTransactionRpcInput to a
// SolanaSignAndSendTransactionRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignAndSendTransactionRpcInputParam.Overrides()
func (r SolanaSignAndSendTransactionRpcInput) ToParam() SolanaSignAndSendTransactionRpcInputParam {
	return param.Override[SolanaSignAndSendTransactionRpcInputParam](json.RawMessage(r.RawJSON()))
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
type SolanaSignAndSendTransactionRpcInputParam struct {
	Caip2 string `json:"caip2" api:"required"`
	// Any of "signAndSendTransaction".
	Method SolanaSignAndSendTransactionRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the SVM `signAndSendTransaction` RPC.
	Params   SolanaSignAndSendTransactionRpcInputParams `json:"params,omitzero" api:"required"`
	Address  param.Opt[string]                          `json:"address,omitzero"`
	Sponsor  param.Opt[bool]                            `json:"sponsor,omitzero"`
	WalletID param.Opt[string]                          `json:"wallet_id,omitzero"`
	// Any of "solana".
	ChainType SolanaSignAndSendTransactionRpcInputChainType `json:"chain_type,omitzero"`
	paramObj
}

func (r SolanaSignAndSendTransactionRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignAndSendTransactionRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignAndSendTransactionRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SolanaSignMessageRpcInput struct {
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
func (r SolanaSignMessageRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SolanaSignMessageRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SolanaSignMessageRpcInput to a
// SolanaSignMessageRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SolanaSignMessageRpcInputParam.Overrides()
func (r SolanaSignMessageRpcInput) ToParam() SolanaSignMessageRpcInputParam {
	return param.Override[SolanaSignMessageRpcInputParam](json.RawMessage(r.RawJSON()))
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
type SolanaSignMessageRpcInputParam struct {
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

func (r SolanaSignMessageRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignMessageRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignMessageRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SparkTransferRpcInput struct {
	// Any of "transfer".
	Method SparkTransferRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `transfer` RPC.
	Params SparkTransferRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkTransferRpcInputNetwork `json:"network"`
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
func (r SparkTransferRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkTransferRpcInput to a SparkTransferRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkTransferRpcInputParam.Overrides()
func (r SparkTransferRpcInput) ToParam() SparkTransferRpcInputParam {
	return param.Override[SparkTransferRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkTransferRpcInputMethod string

const (
	SparkTransferRpcInputMethodTransfer SparkTransferRpcInputMethod = "transfer"
)

type SparkTransferRpcInputNetwork string

const (
	SparkTransferRpcInputNetworkMainnet SparkTransferRpcInputNetwork = "MAINNET"
	SparkTransferRpcInputNetworkRegtest SparkTransferRpcInputNetwork = "REGTEST"
)

// Transfers satoshis to a Spark address.
//
// The properties Method, Params are required.
type SparkTransferRpcInputParam struct {
	// Any of "transfer".
	Method SparkTransferRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `transfer` RPC.
	Params SparkTransferRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkTransferRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkTransferRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gets the balance of the Spark wallet.
type SparkGetBalanceRpcInput struct {
	// Any of "getBalance".
	Method SparkGetBalanceRpcInputMethod `json:"method" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkGetBalanceRpcInputNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetBalanceRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkGetBalanceRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetBalanceRpcInput to a SparkGetBalanceRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetBalanceRpcInputParam.Overrides()
func (r SparkGetBalanceRpcInput) ToParam() SparkGetBalanceRpcInputParam {
	return param.Override[SparkGetBalanceRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkGetBalanceRpcInputMethod string

const (
	SparkGetBalanceRpcInputMethodGetBalance SparkGetBalanceRpcInputMethod = "getBalance"
)

type SparkGetBalanceRpcInputNetwork string

const (
	SparkGetBalanceRpcInputNetworkMainnet SparkGetBalanceRpcInputNetwork = "MAINNET"
	SparkGetBalanceRpcInputNetworkRegtest SparkGetBalanceRpcInputNetwork = "REGTEST"
)

// Gets the balance of the Spark wallet.
//
// The property Method is required.
type SparkGetBalanceRpcInputParam struct {
	// Any of "getBalance".
	Method SparkGetBalanceRpcInputMethod `json:"method,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkGetBalanceRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkGetBalanceRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetBalanceRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetBalanceRpcInputParam) UnmarshalJSON(data []byte) error {
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
	OutputSelectionStrategy SparkOutputSelectionStrategy                          `json:"output_selection_strategy"`
	SelectedOutputs         []SparkTransferTokensRpcInputParamsSelectedOutputResp `json:"selected_outputs"`
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

type SparkTransferTokensRpcInputParamsSelectedOutputResp struct {
	PreviousTransactionHash string                                                    `json:"previous_transaction_hash" api:"required"`
	PreviousTransactionVout float64                                                   `json:"previous_transaction_vout" api:"required"`
	Output                  SparkTransferTokensRpcInputParamsSelectedOutputOutputResp `json:"output"`
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
func (r SparkTransferTokensRpcInputParamsSelectedOutputResp) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferTokensRpcInputParamsSelectedOutputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkTransferTokensRpcInputParamsSelectedOutputOutputResp struct {
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
func (r SparkTransferTokensRpcInputParamsSelectedOutputOutputResp) RawJSON() string {
	return r.JSON.raw
}
func (r *SparkTransferTokensRpcInputParamsSelectedOutputOutputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	OutputSelectionStrategy SparkOutputSelectionStrategy                      `json:"output_selection_strategy,omitzero"`
	SelectedOutputs         []SparkTransferTokensRpcInputParamsSelectedOutput `json:"selected_outputs,omitzero"`
	paramObj
}

func (r SparkTransferTokensRpcInputParams) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferTokensRpcInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferTokensRpcInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PreviousTransactionHash, PreviousTransactionVout are required.
type SparkTransferTokensRpcInputParamsSelectedOutput struct {
	PreviousTransactionHash string                                                `json:"previous_transaction_hash" api:"required"`
	PreviousTransactionVout float64                                               `json:"previous_transaction_vout" api:"required"`
	Output                  SparkTransferTokensRpcInputParamsSelectedOutputOutput `json:"output,omitzero"`
	paramObj
}

func (r SparkTransferTokensRpcInputParamsSelectedOutput) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferTokensRpcInputParamsSelectedOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferTokensRpcInputParamsSelectedOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties OwnerPublicKey, TokenAmount are required.
type SparkTransferTokensRpcInputParamsSelectedOutputOutput struct {
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

func (r SparkTransferTokensRpcInputParamsSelectedOutputOutput) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferTokensRpcInputParamsSelectedOutputOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferTokensRpcInputParamsSelectedOutputOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transfers tokens to a Spark address.
type SparkTransferTokensRpcInput struct {
	// Any of "transferTokens".
	Method SparkTransferTokensRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `transferTokens` RPC.
	Params SparkTransferTokensRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkTransferTokensRpcInputNetwork `json:"network"`
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
func (r SparkTransferTokensRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferTokensRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkTransferTokensRpcInput to a
// SparkTransferTokensRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkTransferTokensRpcInputParam.Overrides()
func (r SparkTransferTokensRpcInput) ToParam() SparkTransferTokensRpcInputParam {
	return param.Override[SparkTransferTokensRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkTransferTokensRpcInputMethod string

const (
	SparkTransferTokensRpcInputMethodTransferTokens SparkTransferTokensRpcInputMethod = "transferTokens"
)

type SparkTransferTokensRpcInputNetwork string

const (
	SparkTransferTokensRpcInputNetworkMainnet SparkTransferTokensRpcInputNetwork = "MAINNET"
	SparkTransferTokensRpcInputNetworkRegtest SparkTransferTokensRpcInputNetwork = "REGTEST"
)

// Transfers tokens to a Spark address.
//
// The properties Method, Params are required.
type SparkTransferTokensRpcInputParam struct {
	// Any of "transferTokens".
	Method SparkTransferTokensRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `transferTokens` RPC.
	Params SparkTransferTokensRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkTransferTokensRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkTransferTokensRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkTransferTokensRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkTransferTokensRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gets a static deposit address for the Spark wallet.
type SparkGetStaticDepositAddressRpcInput struct {
	// Any of "getStaticDepositAddress".
	Method SparkGetStaticDepositAddressRpcInputMethod `json:"method" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkGetStaticDepositAddressRpcInputNetwork `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetStaticDepositAddressRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkGetStaticDepositAddressRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetStaticDepositAddressRpcInput to a
// SparkGetStaticDepositAddressRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetStaticDepositAddressRpcInputParam.Overrides()
func (r SparkGetStaticDepositAddressRpcInput) ToParam() SparkGetStaticDepositAddressRpcInputParam {
	return param.Override[SparkGetStaticDepositAddressRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkGetStaticDepositAddressRpcInputMethod string

const (
	SparkGetStaticDepositAddressRpcInputMethodGetStaticDepositAddress SparkGetStaticDepositAddressRpcInputMethod = "getStaticDepositAddress"
)

type SparkGetStaticDepositAddressRpcInputNetwork string

const (
	SparkGetStaticDepositAddressRpcInputNetworkMainnet SparkGetStaticDepositAddressRpcInputNetwork = "MAINNET"
	SparkGetStaticDepositAddressRpcInputNetworkRegtest SparkGetStaticDepositAddressRpcInputNetwork = "REGTEST"
)

// Gets a static deposit address for the Spark wallet.
//
// The property Method is required.
type SparkGetStaticDepositAddressRpcInputParam struct {
	// Any of "getStaticDepositAddress".
	Method SparkGetStaticDepositAddressRpcInputMethod `json:"method,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkGetStaticDepositAddressRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkGetStaticDepositAddressRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetStaticDepositAddressRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetStaticDepositAddressRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SparkGetClaimStaticDepositQuoteRpcInput struct {
	// Any of "getClaimStaticDepositQuote".
	Method SparkGetClaimStaticDepositQuoteRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `getClaimStaticDepositQuote` RPC.
	Params SparkGetClaimStaticDepositQuoteRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkGetClaimStaticDepositQuoteRpcInputNetwork `json:"network"`
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
func (r SparkGetClaimStaticDepositQuoteRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkGetClaimStaticDepositQuoteRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkGetClaimStaticDepositQuoteRpcInput to a
// SparkGetClaimStaticDepositQuoteRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkGetClaimStaticDepositQuoteRpcInputParam.Overrides()
func (r SparkGetClaimStaticDepositQuoteRpcInput) ToParam() SparkGetClaimStaticDepositQuoteRpcInputParam {
	return param.Override[SparkGetClaimStaticDepositQuoteRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkGetClaimStaticDepositQuoteRpcInputMethod string

const (
	SparkGetClaimStaticDepositQuoteRpcInputMethodGetClaimStaticDepositQuote SparkGetClaimStaticDepositQuoteRpcInputMethod = "getClaimStaticDepositQuote"
)

type SparkGetClaimStaticDepositQuoteRpcInputNetwork string

const (
	SparkGetClaimStaticDepositQuoteRpcInputNetworkMainnet SparkGetClaimStaticDepositQuoteRpcInputNetwork = "MAINNET"
	SparkGetClaimStaticDepositQuoteRpcInputNetworkRegtest SparkGetClaimStaticDepositQuoteRpcInputNetwork = "REGTEST"
)

// Gets a quote for claiming a static deposit.
//
// The properties Method, Params are required.
type SparkGetClaimStaticDepositQuoteRpcInputParam struct {
	// Any of "getClaimStaticDepositQuote".
	Method SparkGetClaimStaticDepositQuoteRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `getClaimStaticDepositQuote` RPC.
	Params SparkGetClaimStaticDepositQuoteRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkGetClaimStaticDepositQuoteRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkGetClaimStaticDepositQuoteRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkGetClaimStaticDepositQuoteRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkGetClaimStaticDepositQuoteRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SparkClaimStaticDepositRpcInput struct {
	// Any of "claimStaticDeposit".
	Method SparkClaimStaticDepositRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `claimStaticDeposit` RPC.
	Params SparkClaimStaticDepositRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkClaimStaticDepositRpcInputNetwork `json:"network"`
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
func (r SparkClaimStaticDepositRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkClaimStaticDepositRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkClaimStaticDepositRpcInput to a
// SparkClaimStaticDepositRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkClaimStaticDepositRpcInputParam.Overrides()
func (r SparkClaimStaticDepositRpcInput) ToParam() SparkClaimStaticDepositRpcInputParam {
	return param.Override[SparkClaimStaticDepositRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkClaimStaticDepositRpcInputMethod string

const (
	SparkClaimStaticDepositRpcInputMethodClaimStaticDeposit SparkClaimStaticDepositRpcInputMethod = "claimStaticDeposit"
)

type SparkClaimStaticDepositRpcInputNetwork string

const (
	SparkClaimStaticDepositRpcInputNetworkMainnet SparkClaimStaticDepositRpcInputNetwork = "MAINNET"
	SparkClaimStaticDepositRpcInputNetworkRegtest SparkClaimStaticDepositRpcInputNetwork = "REGTEST"
)

// Claims a static deposit into the Spark wallet.
//
// The properties Method, Params are required.
type SparkClaimStaticDepositRpcInputParam struct {
	// Any of "claimStaticDeposit".
	Method SparkClaimStaticDepositRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `claimStaticDeposit` RPC.
	Params SparkClaimStaticDepositRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkClaimStaticDepositRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkClaimStaticDepositRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkClaimStaticDepositRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkClaimStaticDepositRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SparkCreateLightningInvoiceRpcInput struct {
	// Any of "createLightningInvoice".
	Method SparkCreateLightningInvoiceRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `createLightningInvoice` RPC.
	Params SparkCreateLightningInvoiceRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkCreateLightningInvoiceRpcInputNetwork `json:"network"`
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
func (r SparkCreateLightningInvoiceRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkCreateLightningInvoiceRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkCreateLightningInvoiceRpcInput to a
// SparkCreateLightningInvoiceRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkCreateLightningInvoiceRpcInputParam.Overrides()
func (r SparkCreateLightningInvoiceRpcInput) ToParam() SparkCreateLightningInvoiceRpcInputParam {
	return param.Override[SparkCreateLightningInvoiceRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkCreateLightningInvoiceRpcInputMethod string

const (
	SparkCreateLightningInvoiceRpcInputMethodCreateLightningInvoice SparkCreateLightningInvoiceRpcInputMethod = "createLightningInvoice"
)

type SparkCreateLightningInvoiceRpcInputNetwork string

const (
	SparkCreateLightningInvoiceRpcInputNetworkMainnet SparkCreateLightningInvoiceRpcInputNetwork = "MAINNET"
	SparkCreateLightningInvoiceRpcInputNetworkRegtest SparkCreateLightningInvoiceRpcInputNetwork = "REGTEST"
)

// Creates a Lightning invoice for the Spark wallet.
//
// The properties Method, Params are required.
type SparkCreateLightningInvoiceRpcInputParam struct {
	// Any of "createLightningInvoice".
	Method SparkCreateLightningInvoiceRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `createLightningInvoice` RPC.
	Params SparkCreateLightningInvoiceRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkCreateLightningInvoiceRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkCreateLightningInvoiceRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkCreateLightningInvoiceRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkCreateLightningInvoiceRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SparkPayLightningInvoiceRpcInput struct {
	// Any of "payLightningInvoice".
	Method SparkPayLightningInvoiceRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `payLightningInvoice` RPC.
	Params SparkPayLightningInvoiceRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkPayLightningInvoiceRpcInputNetwork `json:"network"`
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
func (r SparkPayLightningInvoiceRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkPayLightningInvoiceRpcInput to a
// SparkPayLightningInvoiceRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkPayLightningInvoiceRpcInputParam.Overrides()
func (r SparkPayLightningInvoiceRpcInput) ToParam() SparkPayLightningInvoiceRpcInputParam {
	return param.Override[SparkPayLightningInvoiceRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkPayLightningInvoiceRpcInputMethod string

const (
	SparkPayLightningInvoiceRpcInputMethodPayLightningInvoice SparkPayLightningInvoiceRpcInputMethod = "payLightningInvoice"
)

type SparkPayLightningInvoiceRpcInputNetwork string

const (
	SparkPayLightningInvoiceRpcInputNetworkMainnet SparkPayLightningInvoiceRpcInputNetwork = "MAINNET"
	SparkPayLightningInvoiceRpcInputNetworkRegtest SparkPayLightningInvoiceRpcInputNetwork = "REGTEST"
)

// Pays a Lightning invoice from the Spark wallet.
//
// The properties Method, Params are required.
type SparkPayLightningInvoiceRpcInputParam struct {
	// Any of "payLightningInvoice".
	Method SparkPayLightningInvoiceRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `payLightningInvoice` RPC.
	Params SparkPayLightningInvoiceRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkPayLightningInvoiceRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkPayLightningInvoiceRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkPayLightningInvoiceRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkPayLightningInvoiceRpcInputParam) UnmarshalJSON(data []byte) error {
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
type SparkSignMessageWithIdentityKeyRpcInput struct {
	// Any of "signMessageWithIdentityKey".
	Method SparkSignMessageWithIdentityKeyRpcInputMethod `json:"method" api:"required"`
	// Parameters for the Spark `signMessageWithIdentityKey` RPC.
	Params SparkSignMessageWithIdentityKeyRpcInputParamsResp `json:"params" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkSignMessageWithIdentityKeyRpcInputNetwork `json:"network"`
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
func (r SparkSignMessageWithIdentityKeyRpcInput) RawJSON() string { return r.JSON.raw }
func (r *SparkSignMessageWithIdentityKeyRpcInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparkSignMessageWithIdentityKeyRpcInput to a
// SparkSignMessageWithIdentityKeyRpcInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparkSignMessageWithIdentityKeyRpcInputParam.Overrides()
func (r SparkSignMessageWithIdentityKeyRpcInput) ToParam() SparkSignMessageWithIdentityKeyRpcInputParam {
	return param.Override[SparkSignMessageWithIdentityKeyRpcInputParam](json.RawMessage(r.RawJSON()))
}

type SparkSignMessageWithIdentityKeyRpcInputMethod string

const (
	SparkSignMessageWithIdentityKeyRpcInputMethodSignMessageWithIdentityKey SparkSignMessageWithIdentityKeyRpcInputMethod = "signMessageWithIdentityKey"
)

type SparkSignMessageWithIdentityKeyRpcInputNetwork string

const (
	SparkSignMessageWithIdentityKeyRpcInputNetworkMainnet SparkSignMessageWithIdentityKeyRpcInputNetwork = "MAINNET"
	SparkSignMessageWithIdentityKeyRpcInputNetworkRegtest SparkSignMessageWithIdentityKeyRpcInputNetwork = "REGTEST"
)

// Signs a message with the Spark identity key.
//
// The properties Method, Params are required.
type SparkSignMessageWithIdentityKeyRpcInputParam struct {
	// Any of "signMessageWithIdentityKey".
	Method SparkSignMessageWithIdentityKeyRpcInputMethod `json:"method,omitzero" api:"required"`
	// Parameters for the Spark `signMessageWithIdentityKey` RPC.
	Params SparkSignMessageWithIdentityKeyRpcInputParams `json:"params,omitzero" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network SparkSignMessageWithIdentityKeyRpcInputNetwork `json:"network,omitzero"`
	paramObj
}

func (r SparkSignMessageWithIdentityKeyRpcInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SparkSignMessageWithIdentityKeyRpcInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparkSignMessageWithIdentityKeyRpcInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `transfer` RPC.
type SparkTransferRpcResponse struct {
	// Any of "transfer".
	Method SparkTransferRpcResponseMethod `json:"method" api:"required"`
	Data   SparkTransferRpcResponseData   `json:"data"`
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

type SparkTransferRpcResponseData struct {
	ID                        string                             `json:"id" api:"required"`
	Leaves                    []SparkTransferRpcResponseDataLeaf `json:"leaves" api:"required"`
	ReceiverIdentityPublicKey string                             `json:"receiver_identity_public_key" api:"required"`
	SenderIdentityPublicKey   string                             `json:"sender_identity_public_key" api:"required"`
	Status                    string                             `json:"status" api:"required"`
	TotalValue                float64                            `json:"total_value" api:"required"`
	TransferDirection         string                             `json:"transfer_direction" api:"required"`
	Type                      string                             `json:"type" api:"required"`
	CreatedTime               string                             `json:"created_time"`
	ExpiryTime                string                             `json:"expiry_time"`
	UpdatedTime               string                             `json:"updated_time"`
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
func (r SparkTransferRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkTransferRpcResponseDataLeaf struct {
	IntermediateRefundTx string                               `json:"intermediate_refund_tx" api:"required"`
	SecretCipher         string                               `json:"secret_cipher" api:"required"`
	Signature            string                               `json:"signature" api:"required"`
	Leaf                 SparkTransferRpcResponseDataLeafLeaf `json:"leaf"`
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
func (r SparkTransferRpcResponseDataLeaf) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcResponseDataLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkTransferRpcResponseDataLeafLeaf struct {
	ID string `json:"id" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network                string                                              `json:"network" api:"required"`
	NodeTx                 string                                              `json:"node_tx" api:"required"`
	OwnerIdentityPublicKey string                                              `json:"owner_identity_public_key" api:"required"`
	RefundTx               string                                              `json:"refund_tx" api:"required"`
	Status                 string                                              `json:"status" api:"required"`
	TreeID                 string                                              `json:"tree_id" api:"required"`
	Value                  float64                                             `json:"value" api:"required"`
	VerifyingPublicKey     string                                              `json:"verifying_public_key" api:"required"`
	Vout                   float64                                             `json:"vout" api:"required"`
	ParentNodeID           string                                              `json:"parent_node_id"`
	SigningKeyshare        SparkTransferRpcResponseDataLeafLeafSigningKeyshare `json:"signing_keyshare"`
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
func (r SparkTransferRpcResponseDataLeafLeaf) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcResponseDataLeafLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkTransferRpcResponseDataLeafLeafSigningKeyshare struct {
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
func (r SparkTransferRpcResponseDataLeafLeafSigningKeyshare) RawJSON() string { return r.JSON.raw }
func (r *SparkTransferRpcResponseDataLeafLeafSigningKeyshare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `getBalance` RPC.
type SparkGetBalanceRpcResponse struct {
	// Any of "getBalance".
	Method SparkGetBalanceRpcResponseMethod `json:"method" api:"required"`
	Data   SparkGetBalanceRpcResponseData   `json:"data"`
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

type SparkGetBalanceRpcResponseData struct {
	Balance       string                                                `json:"balance" api:"required"`
	TokenBalances map[string]SparkGetBalanceRpcResponseDataTokenBalance `json:"token_balances" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance       respjson.Field
		TokenBalances respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetBalanceRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkGetBalanceRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkGetBalanceRpcResponseDataTokenBalance struct {
	Balance       string                                                  `json:"balance" api:"required"`
	TokenMetadata SparkGetBalanceRpcResponseDataTokenBalanceTokenMetadata `json:"token_metadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance       respjson.Field
		TokenMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparkGetBalanceRpcResponseDataTokenBalance) RawJSON() string { return r.JSON.raw }
func (r *SparkGetBalanceRpcResponseDataTokenBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkGetBalanceRpcResponseDataTokenBalanceTokenMetadata struct {
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
func (r SparkGetBalanceRpcResponseDataTokenBalanceTokenMetadata) RawJSON() string { return r.JSON.raw }
func (r *SparkGetBalanceRpcResponseDataTokenBalanceTokenMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	Data   SparkCreateLightningInvoiceRpcResponseData   `json:"data"`
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

type SparkCreateLightningInvoiceRpcResponseData struct {
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
func (r SparkCreateLightningInvoiceRpcResponseData) RawJSON() string { return r.JSON.raw }
func (r *SparkCreateLightningInvoiceRpcResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response to the Spark `payLightningInvoice` RPC.
type SparkPayLightningInvoiceRpcResponse struct {
	// Any of "payLightningInvoice".
	Method SparkPayLightningInvoiceRpcResponseMethod    `json:"method" api:"required"`
	Data   SparkPayLightningInvoiceRpcResponseDataUnion `json:"data"`
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
// and values from [SparkPayLightningInvoiceRpcResponseDataObject],
// [SparkPayLightningInvoiceRpcResponseDataObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SparkPayLightningInvoiceRpcResponseDataUnion struct {
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	ID string `json:"id"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	Leaves []SparkPayLightningInvoiceRpcResponseDataObjectLeaf `json:"leaves"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	ReceiverIdentityPublicKey string `json:"receiver_identity_public_key"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	SenderIdentityPublicKey string `json:"sender_identity_public_key"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	Status string `json:"status"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	TotalValue float64 `json:"total_value"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	TransferDirection string `json:"transfer_direction"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	Type string `json:"type"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	CreatedTime string `json:"created_time"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	ExpiryTime string `json:"expiry_time"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	UpdatedTime string `json:"updated_time"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	CreatedAt string `json:"created_at"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	EncodedInvoice string `json:"encoded_invoice"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	Fee SparkPayLightningInvoiceRpcResponseDataObjectFee `json:"fee"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	IdempotencyKey string `json:"idempotency_key"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	Network string `json:"network"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	Typename string `json:"typename"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	UpdatedAt string `json:"updated_at"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
	PaymentPreimage string `json:"payment_preimage"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataObject].
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

func (u SparkPayLightningInvoiceRpcResponseDataUnion) AsSparkPayLightningInvoiceRpcResponseDataObject() (v SparkPayLightningInvoiceRpcResponseDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SparkPayLightningInvoiceRpcResponseDataUnion) AsVariant2() (v SparkPayLightningInvoiceRpcResponseDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SparkPayLightningInvoiceRpcResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *SparkPayLightningInvoiceRpcResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkPayLightningInvoiceRpcResponseDataObject struct {
	ID                        string                                              `json:"id" api:"required"`
	Leaves                    []SparkPayLightningInvoiceRpcResponseDataObjectLeaf `json:"leaves" api:"required"`
	ReceiverIdentityPublicKey string                                              `json:"receiver_identity_public_key" api:"required"`
	SenderIdentityPublicKey   string                                              `json:"sender_identity_public_key" api:"required"`
	Status                    string                                              `json:"status" api:"required"`
	TotalValue                float64                                             `json:"total_value" api:"required"`
	TransferDirection         string                                              `json:"transfer_direction" api:"required"`
	Type                      string                                              `json:"type" api:"required"`
	CreatedTime               string                                              `json:"created_time"`
	ExpiryTime                string                                              `json:"expiry_time"`
	UpdatedTime               string                                              `json:"updated_time"`
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
func (r SparkPayLightningInvoiceRpcResponseDataObject) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcResponseDataObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkPayLightningInvoiceRpcResponseDataObjectLeaf struct {
	IntermediateRefundTx string                                                `json:"intermediate_refund_tx" api:"required"`
	SecretCipher         string                                                `json:"secret_cipher" api:"required"`
	Signature            string                                                `json:"signature" api:"required"`
	Leaf                 SparkPayLightningInvoiceRpcResponseDataObjectLeafLeaf `json:"leaf"`
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
func (r SparkPayLightningInvoiceRpcResponseDataObjectLeaf) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcResponseDataObjectLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkPayLightningInvoiceRpcResponseDataObjectLeafLeaf struct {
	ID string `json:"id" api:"required"`
	// Any of "MAINNET", "REGTEST".
	Network                string                                                               `json:"network" api:"required"`
	NodeTx                 string                                                               `json:"node_tx" api:"required"`
	OwnerIdentityPublicKey string                                                               `json:"owner_identity_public_key" api:"required"`
	RefundTx               string                                                               `json:"refund_tx" api:"required"`
	Status                 string                                                               `json:"status" api:"required"`
	TreeID                 string                                                               `json:"tree_id" api:"required"`
	Value                  float64                                                              `json:"value" api:"required"`
	VerifyingPublicKey     string                                                               `json:"verifying_public_key" api:"required"`
	Vout                   float64                                                              `json:"vout" api:"required"`
	ParentNodeID           string                                                               `json:"parent_node_id"`
	SigningKeyshare        SparkPayLightningInvoiceRpcResponseDataObjectLeafLeafSigningKeyshare `json:"signing_keyshare"`
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
func (r SparkPayLightningInvoiceRpcResponseDataObjectLeafLeaf) RawJSON() string { return r.JSON.raw }
func (r *SparkPayLightningInvoiceRpcResponseDataObjectLeafLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SparkPayLightningInvoiceRpcResponseDataObjectLeafLeafSigningKeyshare struct {
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
func (r SparkPayLightningInvoiceRpcResponseDataObjectLeafLeafSigningKeyshare) RawJSON() string {
	return r.JSON.raw
}
func (r *SparkPayLightningInvoiceRpcResponseDataObjectLeafLeafSigningKeyshare) UnmarshalJSON(data []byte) error {
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
	OwnerID string `json:"owner_id" api:"required" format:"cuid2"`
	// List of policy IDs for policies that are enforced on the wallet.
	PolicyIDs []string `json:"policy_ids" api:"required"`
	// The number of keys that must sign for an action to be valid.
	AuthorizationThreshold float64 `json:"authorization_threshold"`
	// Information about the custodian managing this wallet.
	Custody WalletCustodian `json:"custody"`
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

// Request body for updating a wallet.
type WalletUpdateRequestBody struct {
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner WalletUpdateRequestBodyOwnerUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletUpdateRequestBodyAdditionalSigner `json:"additional_signers,omitzero"`
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

// The property SignerID is required.
type WalletUpdateRequestBodyAdditionalSigner struct {
	SignerID string `json:"signer_id" api:"required" format:"cuid2"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet.
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r WalletUpdateRequestBodyAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateRequestBodyAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateRequestBodyAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletUpdateRequestBodyOwnerUnion struct {
	OfPublicKeyOwner *WalletUpdateRequestBodyOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *WalletUpdateRequestBodyOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u WalletUpdateRequestBodyOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *WalletUpdateRequestBodyOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type WalletUpdateRequestBodyOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r WalletUpdateRequestBodyOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateRequestBodyOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateRequestBodyOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type WalletUpdateRequestBodyOwnerUserOwner struct {
	UserID string `json:"user_id" api:"required"`
	paramObj
}

func (r WalletUpdateRequestBodyOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateRequestBodyOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateRequestBodyOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnion contains all possible properties and values from
// [EthereumSignTransactionRpcInput], [EthereumSendTransactionRpcInput],
// [EthereumPersonalSignRpcInput], [EthereumSignTypedDataRpcInput],
// [EthereumSecp256k1SignRpcInput], [EthereumSign7702AuthorizationRpcInput],
// [EthereumSignUserOperationRpcInput], [SolanaSignTransactionRpcInput],
// [SolanaSignAndSendTransactionRpcInput], [SolanaSignMessageRpcInput],
// [SparkTransferRpcInput], [SparkGetBalanceRpcInput],
// [SparkTransferTokensRpcInput], [SparkGetStaticDepositAddressRpcInput],
// [SparkGetClaimStaticDepositQuoteRpcInput], [SparkClaimStaticDepositRpcInput],
// [SparkCreateLightningInvoiceRpcInput], [SparkPayLightningInvoiceRpcInput],
// [SparkSignMessageWithIdentityKeyRpcInput], [ExportPrivateKeyRpcInput].
//
// Use the [WalletRpcRequestBodyUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletRpcRequestBodyUnion struct {
	// Any of "eth_signTransaction", "eth_sendTransaction", "personal_sign",
	// "eth_signTypedData_v4", "secp256k1_sign", "eth_sign7702Authorization",
	// "eth_signUserOperation", "signTransaction", "signAndSendTransaction",
	// "signMessage", "transfer", "getBalance", "transferTokens",
	// "getStaticDepositAddress", "getClaimStaticDepositQuote", "claimStaticDeposit",
	// "createLightningInvoice", "payLightningInvoice", "signMessageWithIdentityKey",
	// "exportPrivateKey".
	Method string `json:"method"`
	// This field is a union of [EthereumSignTransactionRpcInputParamsResp],
	// [EthereumSendTransactionRpcInputParamsResp],
	// [EthereumPersonalSignRpcInputParamsResp],
	// [EthereumSignTypedDataRpcInputParamsResp],
	// [EthereumSecp256k1SignRpcInputParamsResp],
	// [EthereumSign7702AuthorizationRpcInputParamsResp],
	// [EthereumSignUserOperationRpcInputParamsResp],
	// [SolanaSignTransactionRpcInputParamsResp],
	// [SolanaSignAndSendTransactionRpcInputParamsResp],
	// [SolanaSignMessageRpcInputParamsResp], [SparkTransferRpcInputParamsResp],
	// [SparkTransferTokensRpcInputParamsResp],
	// [SparkGetClaimStaticDepositQuoteRpcInputParamsResp],
	// [SparkClaimStaticDepositRpcInputParamsResp],
	// [SparkCreateLightningInvoiceRpcInputParamsResp],
	// [SparkPayLightningInvoiceRpcInputParamsResp],
	// [SparkSignMessageWithIdentityKeyRpcInputParamsResp],
	// [ExportPrivateKeyRpcInputParams]
	Params    WalletRpcRequestBodyUnionParams `json:"params"`
	Address   string                          `json:"address"`
	ChainType string                          `json:"chain_type"`
	WalletID  string                          `json:"wallet_id"`
	Caip2     string                          `json:"caip2"`
	Sponsor   bool                            `json:"sponsor"`
	Network   string                          `json:"network"`
	JSON      struct {
		Method    respjson.Field
		Params    respjson.Field
		Address   respjson.Field
		ChainType respjson.Field
		WalletID  respjson.Field
		Caip2     respjson.Field
		Sponsor   respjson.Field
		Network   respjson.Field
		raw       string
	} `json:"-"`
}

// anyWalletRpcRequestBody is implemented by each variant of
// [WalletRpcRequestBodyUnion] to add type safety for the return type of
// [WalletRpcRequestBodyUnion.AsAny]
type anyWalletRpcRequestBody interface {
	implWalletRpcRequestBodyUnion()
}

func (EthereumSignTransactionRpcInput) implWalletRpcRequestBodyUnion()         {}
func (EthereumSendTransactionRpcInput) implWalletRpcRequestBodyUnion()         {}
func (EthereumPersonalSignRpcInput) implWalletRpcRequestBodyUnion()            {}
func (EthereumSignTypedDataRpcInput) implWalletRpcRequestBodyUnion()           {}
func (EthereumSecp256k1SignRpcInput) implWalletRpcRequestBodyUnion()           {}
func (EthereumSign7702AuthorizationRpcInput) implWalletRpcRequestBodyUnion()   {}
func (EthereumSignUserOperationRpcInput) implWalletRpcRequestBodyUnion()       {}
func (SolanaSignTransactionRpcInput) implWalletRpcRequestBodyUnion()           {}
func (SolanaSignAndSendTransactionRpcInput) implWalletRpcRequestBodyUnion()    {}
func (SolanaSignMessageRpcInput) implWalletRpcRequestBodyUnion()               {}
func (SparkTransferRpcInput) implWalletRpcRequestBodyUnion()                   {}
func (SparkGetBalanceRpcInput) implWalletRpcRequestBodyUnion()                 {}
func (SparkTransferTokensRpcInput) implWalletRpcRequestBodyUnion()             {}
func (SparkGetStaticDepositAddressRpcInput) implWalletRpcRequestBodyUnion()    {}
func (SparkGetClaimStaticDepositQuoteRpcInput) implWalletRpcRequestBodyUnion() {}
func (SparkClaimStaticDepositRpcInput) implWalletRpcRequestBodyUnion()         {}
func (SparkCreateLightningInvoiceRpcInput) implWalletRpcRequestBodyUnion()     {}
func (SparkPayLightningInvoiceRpcInput) implWalletRpcRequestBodyUnion()        {}
func (SparkSignMessageWithIdentityKeyRpcInput) implWalletRpcRequestBodyUnion() {}
func (ExportPrivateKeyRpcInput) implWalletRpcRequestBodyUnion()                {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletRpcRequestBodyUnion.AsAny().(type) {
//	case privyclient.EthereumSignTransactionRpcInput:
//	case privyclient.EthereumSendTransactionRpcInput:
//	case privyclient.EthereumPersonalSignRpcInput:
//	case privyclient.EthereumSignTypedDataRpcInput:
//	case privyclient.EthereumSecp256k1SignRpcInput:
//	case privyclient.EthereumSign7702AuthorizationRpcInput:
//	case privyclient.EthereumSignUserOperationRpcInput:
//	case privyclient.SolanaSignTransactionRpcInput:
//	case privyclient.SolanaSignAndSendTransactionRpcInput:
//	case privyclient.SolanaSignMessageRpcInput:
//	case privyclient.SparkTransferRpcInput:
//	case privyclient.SparkGetBalanceRpcInput:
//	case privyclient.SparkTransferTokensRpcInput:
//	case privyclient.SparkGetStaticDepositAddressRpcInput:
//	case privyclient.SparkGetClaimStaticDepositQuoteRpcInput:
//	case privyclient.SparkClaimStaticDepositRpcInput:
//	case privyclient.SparkCreateLightningInvoiceRpcInput:
//	case privyclient.SparkPayLightningInvoiceRpcInput:
//	case privyclient.SparkSignMessageWithIdentityKeyRpcInput:
//	case privyclient.ExportPrivateKeyRpcInput:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletRpcRequestBodyUnion) AsAny() anyWalletRpcRequestBody {
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
	}
	return nil
}

func (u WalletRpcRequestBodyUnion) AsEthSignTransaction() (v EthereumSignTransactionRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsEthSendTransaction() (v EthereumSendTransactionRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsPersonalSign() (v EthereumPersonalSignRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsEthSignTypedDataV4() (v EthereumSignTypedDataRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsSecp256k1Sign() (v EthereumSecp256k1SignRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsEthSign7702Authorization() (v EthereumSign7702AuthorizationRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsEthSignUserOperation() (v EthereumSignUserOperationRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsSignTransaction() (v SolanaSignTransactionRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsSignAndSendTransaction() (v SolanaSignAndSendTransactionRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsSignMessage() (v SolanaSignMessageRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsTransfer() (v SparkTransferRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsGetBalance() (v SparkGetBalanceRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsTransferTokens() (v SparkTransferTokensRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsGetStaticDepositAddress() (v SparkGetStaticDepositAddressRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsGetClaimStaticDepositQuote() (v SparkGetClaimStaticDepositQuoteRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsClaimStaticDeposit() (v SparkClaimStaticDepositRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsCreateLightningInvoice() (v SparkCreateLightningInvoiceRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsPayLightningInvoice() (v SparkPayLightningInvoiceRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsSignMessageWithIdentityKey() (v SparkSignMessageWithIdentityKeyRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcRequestBodyUnion) AsExportPrivateKey() (v ExportPrivateKeyRpcInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcRequestBodyUnion) RawJSON() string { return u.JSON.raw }

func (r *WalletRpcRequestBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParams is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParams provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
type WalletRpcRequestBodyUnionParams struct {
	// This field is a union of [EthereumSignTransactionRpcInputParamsTransactionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionResp], [string], [string]
	Transaction WalletRpcRequestBodyUnionParamsTransaction `json:"transaction"`
	Encoding    string                                     `json:"encoding"`
	Message     string                                     `json:"message"`
	// This field is from variant [EthereumSignTypedDataRpcInputParamsResp].
	TypedData EthereumTypedDataInput `json:"typed_data"`
	// This field is from variant [EthereumSecp256k1SignRpcInputParamsResp].
	Hash string `json:"hash"`
	// This field is a union of
	// [EthereumSign7702AuthorizationRpcInputParamsChainIDUnionResp],
	// [EthereumSignUserOperationRpcInputParamsChainIDUnionResp]
	ChainID  WalletRpcRequestBodyUnionParamsChainID `json:"chain_id"`
	Contract string                                 `json:"contract"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	Executor EthereumSign7702AuthorizationRpcInputParamsExecutor `json:"executor"`
	// This field is from variant [EthereumSign7702AuthorizationRpcInputParamsResp].
	Nonce EthereumSign7702AuthorizationRpcInputParamsNonceUnionResp `json:"nonce"`
	// This field is from variant [EthereumSignUserOperationRpcInputParamsResp].
	UserOperation        EthereumSignUserOperationRpcInputParamsUserOperationResp `json:"user_operation"`
	AmountSats           float64                                                  `json:"amount_sats"`
	ReceiverSparkAddress string                                                   `json:"receiver_spark_address"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	TokenAmount float64 `json:"token_amount"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	TokenIdentifier string `json:"token_identifier"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	OutputSelectionStrategy SparkOutputSelectionStrategy `json:"output_selection_strategy"`
	// This field is from variant [SparkTransferTokensRpcInputParamsResp].
	SelectedOutputs []SparkTransferTokensRpcInputParamsSelectedOutputResp `json:"selected_outputs"`
	TransactionID   string                                                `json:"transaction_id"`
	OutputIndex     float64                                               `json:"output_index"`
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
	// This field is from variant [ExportPrivateKeyRpcInputParams].
	EncryptionType string `json:"encryption_type"`
	// This field is from variant [ExportPrivateKeyRpcInputParams].
	RecipientPublicKey string `json:"recipient_public_key"`
	// This field is from variant [ExportPrivateKeyRpcInputParams].
	ExportType string `json:"export_type"`
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
		ExportType              respjson.Field
		raw                     string
	} `json:"-"`
}

func (r *WalletRpcRequestBodyUnionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransaction is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsTransaction provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type WalletRpcRequestBodyUnionParamsTransaction struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is a union of
	// [[]EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp],
	// [[]EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp]
	AuthorizationList WalletRpcRequestBodyUnionParamsTransactionAuthorizationList `json:"authorization_list"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionChainIDUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionChainIDUnionResp]
	ChainID WalletRpcRequestBodyUnionParamsTransactionChainID `json:"chain_id"`
	Data    string                                            `json:"data"`
	From    string                                            `json:"from"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionResp]
	GasLimit WalletRpcRequestBodyUnionParamsTransactionGasLimit `json:"gas_limit"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionResp]
	GasPrice WalletRpcRequestBodyUnionParamsTransactionGasPrice `json:"gas_price"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionResp]
	MaxFeePerGas WalletRpcRequestBodyUnionParamsTransactionMaxFeePerGas `json:"max_fee_per_gas"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionResp]
	MaxPriorityFeePerGas WalletRpcRequestBodyUnionParamsTransactionMaxPriorityFeePerGas `json:"max_priority_fee_per_gas"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionNonceUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionNonceUnionResp]
	Nonce WalletRpcRequestBodyUnionParamsTransactionNonce `json:"nonce"`
	To    string                                          `json:"to"`
	Type  float64                                         `json:"type"`
	// This field is a union of
	// [EthereumSignTransactionRpcInputParamsTransactionValueUnionResp],
	// [EthereumSendTransactionRpcInputParamsTransactionValueUnionResp]
	Value WalletRpcRequestBodyUnionParamsTransactionValue `json:"value"`
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

func (r *WalletRpcRequestBodyUnionParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionAuthorizationList is an implicit
// subunion of [WalletRpcRequestBodyUnion].
// WalletRpcRequestBodyUnionParamsTransactionAuthorizationList provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfEthereumSignTransactionRpcInputsTransactionAuthorizationListArray
// OfEthereumSendTransactionRpcInputsTransactionAuthorizationListArray]
type WalletRpcRequestBodyUnionParamsTransactionAuthorizationList struct {
	// This field will be present if the value is a
	// [[]EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp]
	// instead of an object.
	OfEthereumSignTransactionRpcInputsTransactionAuthorizationListArray []EthereumSignTransactionRpcInputParamsTransactionAuthorizationListResp `json:",inline"`
	// This field will be present if the value is a
	// [[]EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp]
	// instead of an object.
	OfEthereumSendTransactionRpcInputsTransactionAuthorizationListArray []EthereumSendTransactionRpcInputParamsTransactionAuthorizationListResp `json:",inline"`
	JSON                                                                struct {
		OfEthereumSignTransactionRpcInputsTransactionAuthorizationListArray respjson.Field
		OfEthereumSendTransactionRpcInputsTransactionAuthorizationListArray respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (r *WalletRpcRequestBodyUnionParamsTransactionAuthorizationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionChainID is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsTransactionChainID
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionChainID struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionChainID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionGasLimit is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsTransactionGasLimit
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionGasLimit struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionGasLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionGasPrice is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsTransactionGasPrice
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionGasPrice struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionGasPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionMaxFeePerGas is an implicit subunion
// of [WalletRpcRequestBodyUnion].
// WalletRpcRequestBodyUnionParamsTransactionMaxFeePerGas provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionMaxFeePerGas struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionMaxFeePerGas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionMaxPriorityFeePerGas is an implicit
// subunion of [WalletRpcRequestBodyUnion].
// WalletRpcRequestBodyUnionParamsTransactionMaxPriorityFeePerGas provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionMaxPriorityFeePerGas struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionMaxPriorityFeePerGas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionNonce is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsTransactionNonce
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionNonce struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionNonce) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsTransactionValue is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsTransactionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsTransactionValue struct {
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

func (r *WalletRpcRequestBodyUnionParamsTransactionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcRequestBodyUnionParamsChainID is an implicit subunion of
// [WalletRpcRequestBodyUnion]. WalletRpcRequestBodyUnionParamsChainID provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcRequestBodyUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcRequestBodyUnionParamsChainID struct {
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

func (r *WalletRpcRequestBodyUnionParamsChainID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WalletRpcRequestBodyUnion to a
// WalletRpcRequestBodyUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WalletRpcRequestBodyUnionParam.Overrides()
func (r WalletRpcRequestBodyUnion) ToParam() WalletRpcRequestBodyUnionParam {
	return param.Override[WalletRpcRequestBodyUnionParam](json.RawMessage(r.RawJSON()))
}

func WalletRpcRequestBodyParamOfEthSignTransaction(params EthereumSignTransactionRpcInputParams) WalletRpcRequestBodyUnionParam {
	var ethSignTransaction EthereumSignTransactionRpcInputParam
	ethSignTransaction.Params = params
	return WalletRpcRequestBodyUnionParam{OfEthSignTransaction: &ethSignTransaction}
}

func WalletRpcRequestBodyParamOfEthSendTransaction(caip2 string, method EthereumSendTransactionRpcInputMethod, params EthereumSendTransactionRpcInputParams) WalletRpcRequestBodyUnionParam {
	var ethSendTransaction EthereumSendTransactionRpcInputParam
	ethSendTransaction.Caip2 = caip2
	ethSendTransaction.Method = method
	ethSendTransaction.Params = params
	return WalletRpcRequestBodyUnionParam{OfEthSendTransaction: &ethSendTransaction}
}

func WalletRpcRequestBodyParamOfPersonalSign(params EthereumPersonalSignRpcInputParams) WalletRpcRequestBodyUnionParam {
	var personalSign EthereumPersonalSignRpcInputParam
	personalSign.Params = params
	return WalletRpcRequestBodyUnionParam{OfPersonalSign: &personalSign}
}

func WalletRpcRequestBodyParamOfEthSignTypedDataV4(params EthereumSignTypedDataRpcInputParams) WalletRpcRequestBodyUnionParam {
	var ethSignTypedDataV4 EthereumSignTypedDataRpcInputParam
	ethSignTypedDataV4.Params = params
	return WalletRpcRequestBodyUnionParam{OfEthSignTypedDataV4: &ethSignTypedDataV4}
}

func WalletRpcRequestBodyParamOfSecp256k1Sign(params EthereumSecp256k1SignRpcInputParams) WalletRpcRequestBodyUnionParam {
	var secp256k1Sign EthereumSecp256k1SignRpcInputParam
	secp256k1Sign.Params = params
	return WalletRpcRequestBodyUnionParam{OfSecp256k1Sign: &secp256k1Sign}
}

func WalletRpcRequestBodyParamOfEthSign7702Authorization(params EthereumSign7702AuthorizationRpcInputParams) WalletRpcRequestBodyUnionParam {
	var ethSign7702Authorization EthereumSign7702AuthorizationRpcInputParam
	ethSign7702Authorization.Params = params
	return WalletRpcRequestBodyUnionParam{OfEthSign7702Authorization: &ethSign7702Authorization}
}

func WalletRpcRequestBodyParamOfEthSignUserOperation(params EthereumSignUserOperationRpcInputParams) WalletRpcRequestBodyUnionParam {
	var ethSignUserOperation EthereumSignUserOperationRpcInputParam
	ethSignUserOperation.Params = params
	return WalletRpcRequestBodyUnionParam{OfEthSignUserOperation: &ethSignUserOperation}
}

func WalletRpcRequestBodyParamOfSignTransaction(params SolanaSignTransactionRpcInputParams) WalletRpcRequestBodyUnionParam {
	var signTransaction SolanaSignTransactionRpcInputParam
	signTransaction.Params = params
	return WalletRpcRequestBodyUnionParam{OfSignTransaction: &signTransaction}
}

func WalletRpcRequestBodyParamOfSignAndSendTransaction(caip2 string, method SolanaSignAndSendTransactionRpcInputMethod, params SolanaSignAndSendTransactionRpcInputParams) WalletRpcRequestBodyUnionParam {
	var signAndSendTransaction SolanaSignAndSendTransactionRpcInputParam
	signAndSendTransaction.Caip2 = caip2
	signAndSendTransaction.Method = method
	signAndSendTransaction.Params = params
	return WalletRpcRequestBodyUnionParam{OfSignAndSendTransaction: &signAndSendTransaction}
}

func WalletRpcRequestBodyParamOfSignMessage(params SolanaSignMessageRpcInputParams) WalletRpcRequestBodyUnionParam {
	var signMessage SolanaSignMessageRpcInputParam
	signMessage.Params = params
	return WalletRpcRequestBodyUnionParam{OfSignMessage: &signMessage}
}

func WalletRpcRequestBodyParamOfTransfer(params SparkTransferRpcInputParams) WalletRpcRequestBodyUnionParam {
	var transfer SparkTransferRpcInputParam
	transfer.Params = params
	return WalletRpcRequestBodyUnionParam{OfTransfer: &transfer}
}

func WalletRpcRequestBodyParamOfGetBalance(method SparkGetBalanceRpcInputMethod) WalletRpcRequestBodyUnionParam {
	var getBalance SparkGetBalanceRpcInputParam
	getBalance.Method = method
	return WalletRpcRequestBodyUnionParam{OfGetBalance: &getBalance}
}

func WalletRpcRequestBodyParamOfTransferTokens(params SparkTransferTokensRpcInputParams) WalletRpcRequestBodyUnionParam {
	var transferTokens SparkTransferTokensRpcInputParam
	transferTokens.Params = params
	return WalletRpcRequestBodyUnionParam{OfTransferTokens: &transferTokens}
}

func WalletRpcRequestBodyParamOfGetStaticDepositAddress(method SparkGetStaticDepositAddressRpcInputMethod) WalletRpcRequestBodyUnionParam {
	var getStaticDepositAddress SparkGetStaticDepositAddressRpcInputParam
	getStaticDepositAddress.Method = method
	return WalletRpcRequestBodyUnionParam{OfGetStaticDepositAddress: &getStaticDepositAddress}
}

func WalletRpcRequestBodyParamOfGetClaimStaticDepositQuote(params SparkGetClaimStaticDepositQuoteRpcInputParams) WalletRpcRequestBodyUnionParam {
	var getClaimStaticDepositQuote SparkGetClaimStaticDepositQuoteRpcInputParam
	getClaimStaticDepositQuote.Params = params
	return WalletRpcRequestBodyUnionParam{OfGetClaimStaticDepositQuote: &getClaimStaticDepositQuote}
}

func WalletRpcRequestBodyParamOfClaimStaticDeposit(params SparkClaimStaticDepositRpcInputParams) WalletRpcRequestBodyUnionParam {
	var claimStaticDeposit SparkClaimStaticDepositRpcInputParam
	claimStaticDeposit.Params = params
	return WalletRpcRequestBodyUnionParam{OfClaimStaticDeposit: &claimStaticDeposit}
}

func WalletRpcRequestBodyParamOfCreateLightningInvoice(params SparkCreateLightningInvoiceRpcInputParams) WalletRpcRequestBodyUnionParam {
	var createLightningInvoice SparkCreateLightningInvoiceRpcInputParam
	createLightningInvoice.Params = params
	return WalletRpcRequestBodyUnionParam{OfCreateLightningInvoice: &createLightningInvoice}
}

func WalletRpcRequestBodyParamOfPayLightningInvoice(params SparkPayLightningInvoiceRpcInputParams) WalletRpcRequestBodyUnionParam {
	var payLightningInvoice SparkPayLightningInvoiceRpcInputParam
	payLightningInvoice.Params = params
	return WalletRpcRequestBodyUnionParam{OfPayLightningInvoice: &payLightningInvoice}
}

func WalletRpcRequestBodyParamOfSignMessageWithIdentityKey(params SparkSignMessageWithIdentityKeyRpcInputParams) WalletRpcRequestBodyUnionParam {
	var signMessageWithIdentityKey SparkSignMessageWithIdentityKeyRpcInputParam
	signMessageWithIdentityKey.Params = params
	return WalletRpcRequestBodyUnionParam{OfSignMessageWithIdentityKey: &signMessageWithIdentityKey}
}

func WalletRpcRequestBodyParamOfExportPrivateKey(address string, method ExportPrivateKeyRpcInputMethod, params ExportPrivateKeyRpcInputParamsParam) WalletRpcRequestBodyUnionParam {
	var exportPrivateKey ExportPrivateKeyRpcInputParam
	exportPrivateKey.Address = address
	exportPrivateKey.Method = method
	exportPrivateKey.Params = params
	return WalletRpcRequestBodyUnionParam{OfExportPrivateKey: &exportPrivateKey}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRpcRequestBodyUnionParam struct {
	OfEthSignTransaction         *EthereumSignTransactionRpcInputParam         `json:",omitzero,inline"`
	OfEthSendTransaction         *EthereumSendTransactionRpcInputParam         `json:",omitzero,inline"`
	OfPersonalSign               *EthereumPersonalSignRpcInputParam            `json:",omitzero,inline"`
	OfEthSignTypedDataV4         *EthereumSignTypedDataRpcInputParam           `json:",omitzero,inline"`
	OfSecp256k1Sign              *EthereumSecp256k1SignRpcInputParam           `json:",omitzero,inline"`
	OfEthSign7702Authorization   *EthereumSign7702AuthorizationRpcInputParam   `json:",omitzero,inline"`
	OfEthSignUserOperation       *EthereumSignUserOperationRpcInputParam       `json:",omitzero,inline"`
	OfSignTransaction            *SolanaSignTransactionRpcInputParam           `json:",omitzero,inline"`
	OfSignAndSendTransaction     *SolanaSignAndSendTransactionRpcInputParam    `json:",omitzero,inline"`
	OfSignMessage                *SolanaSignMessageRpcInputParam               `json:",omitzero,inline"`
	OfTransfer                   *SparkTransferRpcInputParam                   `json:",omitzero,inline"`
	OfGetBalance                 *SparkGetBalanceRpcInputParam                 `json:",omitzero,inline"`
	OfTransferTokens             *SparkTransferTokensRpcInputParam             `json:",omitzero,inline"`
	OfGetStaticDepositAddress    *SparkGetStaticDepositAddressRpcInputParam    `json:",omitzero,inline"`
	OfGetClaimStaticDepositQuote *SparkGetClaimStaticDepositQuoteRpcInputParam `json:",omitzero,inline"`
	OfClaimStaticDeposit         *SparkClaimStaticDepositRpcInputParam         `json:",omitzero,inline"`
	OfCreateLightningInvoice     *SparkCreateLightningInvoiceRpcInputParam     `json:",omitzero,inline"`
	OfPayLightningInvoice        *SparkPayLightningInvoiceRpcInputParam        `json:",omitzero,inline"`
	OfSignMessageWithIdentityKey *SparkSignMessageWithIdentityKeyRpcInputParam `json:",omitzero,inline"`
	OfExportPrivateKey           *ExportPrivateKeyRpcInputParam                `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcRequestBodyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEthSignTransaction,
		u.OfEthSendTransaction,
		u.OfPersonalSign,
		u.OfEthSignTypedDataV4,
		u.OfSecp256k1Sign,
		u.OfEthSign7702Authorization,
		u.OfEthSignUserOperation,
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
		u.OfExportPrivateKey)
}
func (u *WalletRpcRequestBodyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WalletRpcRequestBodyUnionParam](
		"method",
		apijson.Discriminator[EthereumSignTransactionRpcInputParam]("eth_signTransaction"),
		apijson.Discriminator[EthereumSendTransactionRpcInputParam]("eth_sendTransaction"),
		apijson.Discriminator[EthereumPersonalSignRpcInputParam]("personal_sign"),
		apijson.Discriminator[EthereumSignTypedDataRpcInputParam]("eth_signTypedData_v4"),
		apijson.Discriminator[EthereumSecp256k1SignRpcInputParam]("secp256k1_sign"),
		apijson.Discriminator[EthereumSign7702AuthorizationRpcInputParam]("eth_sign7702Authorization"),
		apijson.Discriminator[EthereumSignUserOperationRpcInputParam]("eth_signUserOperation"),
		apijson.Discriminator[SolanaSignTransactionRpcInputParam]("signTransaction"),
		apijson.Discriminator[SolanaSignAndSendTransactionRpcInputParam]("signAndSendTransaction"),
		apijson.Discriminator[SolanaSignMessageRpcInputParam]("signMessage"),
		apijson.Discriminator[SparkTransferRpcInputParam]("transfer"),
		apijson.Discriminator[SparkGetBalanceRpcInputParam]("getBalance"),
		apijson.Discriminator[SparkTransferTokensRpcInputParam]("transferTokens"),
		apijson.Discriminator[SparkGetStaticDepositAddressRpcInputParam]("getStaticDepositAddress"),
		apijson.Discriminator[SparkGetClaimStaticDepositQuoteRpcInputParam]("getClaimStaticDepositQuote"),
		apijson.Discriminator[SparkClaimStaticDepositRpcInputParam]("claimStaticDeposit"),
		apijson.Discriminator[SparkCreateLightningInvoiceRpcInputParam]("createLightningInvoice"),
		apijson.Discriminator[SparkPayLightningInvoiceRpcInputParam]("payLightningInvoice"),
		apijson.Discriminator[SparkSignMessageWithIdentityKeyRpcInputParam]("signMessageWithIdentityKey"),
		apijson.Discriminator[ExportPrivateKeyRpcInputParam]("exportPrivateKey"),
	)
}

// WalletRpcResponseUnion contains all possible properties and values from
// [EthereumPersonalSignRpcResponse], [EthereumSignTypedDataRpcResponse],
// [EthereumSignTransactionRpcResponse], [EthereumSendTransactionRpcResponse],
// [EthereumSignUserOperationRpcResponse],
// [EthereumSign7702AuthorizationRpcResponse], [EthereumSecp256k1SignRpcResponse],
// [SolanaSignMessageRpcResponse], [SolanaSignTransactionRpcResponse],
// [SolanaSignAndSendTransactionRpcResponse], [SparkTransferRpcResponse],
// [SparkGetBalanceRpcResponse], [SparkTransferTokensRpcResponse],
// [SparkGetStaticDepositAddressRpcResponse],
// [SparkGetClaimStaticDepositQuoteRpcResponse],
// [SparkClaimStaticDepositRpcResponse], [SparkCreateLightningInvoiceRpcResponse],
// [SparkPayLightningInvoiceRpcResponse],
// [SparkSignMessageWithIdentityKeyRpcResponse], [ExportPrivateKeyRpcResponse].
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
	// [EthereumSecp256k1SignRpcResponseData], [SolanaSignMessageRpcResponseData],
	// [SolanaSignTransactionRpcResponseData],
	// [SolanaSignAndSendTransactionRpcResponseData], [SparkTransferRpcResponseData],
	// [SparkGetBalanceRpcResponseData], [SparkTransferTokensRpcResponseData],
	// [SparkGetStaticDepositAddressRpcResponseData],
	// [SparkGetClaimStaticDepositQuoteRpcResponseData],
	// [SparkClaimStaticDepositRpcResponseData],
	// [SparkCreateLightningInvoiceRpcResponseData],
	// [SparkPayLightningInvoiceRpcResponseDataUnion],
	// [SparkSignMessageWithIdentityKeyRpcResponseData],
	// [ExportPrivateKeyRpcResponseData]
	Data WalletRpcResponseUnionData `json:"data"`
	// Any of "personal_sign", "eth_signTypedData_v4", "eth_signTransaction",
	// "eth_sendTransaction", "eth_signUserOperation", "eth_sign7702Authorization",
	// "secp256k1_sign", "signMessage", "signTransaction", "signAndSendTransaction",
	// "transfer", "getBalance", "transferTokens", "getStaticDepositAddress",
	// "getClaimStaticDepositQuote", "claimStaticDeposit", "createLightningInvoice",
	// "payLightningInvoice", "signMessageWithIdentityKey", "exportPrivateKey".
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
	Caip2             string `json:"caip2"`
	Hash              string `json:"hash"`
	TransactionID     string `json:"transaction_id"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	TransactionRequest EthereumSendTransactionRpcResponseDataTransactionRequest `json:"transaction_request"`
	// This field is from variant [EthereumSendTransactionRpcResponseData].
	UserOperationHash string `json:"user_operation_hash"`
	// This field is from variant [EthereumSign7702AuthorizationRpcResponseData].
	Authorization EthereumSign7702AuthorizationRpcResponseDataAuthorization `json:"authorization"`
	ID            string                                                    `json:"id"`
	// This field is a union of [[]SparkTransferRpcResponseDataLeaf],
	// [[]SparkPayLightningInvoiceRpcResponseDataObjectLeaf]
	Leaves                    WalletRpcResponseUnionDataLeaves `json:"leaves"`
	ReceiverIdentityPublicKey string                           `json:"receiver_identity_public_key"`
	SenderIdentityPublicKey   string                           `json:"sender_identity_public_key"`
	Status                    string                           `json:"status"`
	TotalValue                float64                          `json:"total_value"`
	TransferDirection         string                           `json:"transfer_direction"`
	Type                      string                           `json:"type"`
	CreatedTime               string                           `json:"created_time"`
	ExpiryTime                string                           `json:"expiry_time"`
	UpdatedTime               string                           `json:"updated_time"`
	// This field is from variant [SparkGetBalanceRpcResponseData].
	Balance string `json:"balance"`
	// This field is from variant [SparkGetBalanceRpcResponseData].
	TokenBalances map[string]SparkGetBalanceRpcResponseDataTokenBalance `json:"token_balances"`
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
	// This field is from variant [SparkCreateLightningInvoiceRpcResponseData].
	Invoice         any    `json:"invoice"`
	PaymentPreimage string `json:"payment_preimage"`
	Transfer        any    `json:"transfer"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataUnion].
	EncodedInvoice string `json:"encoded_invoice"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataUnion].
	Fee SparkPayLightningInvoiceRpcResponseDataObjectFee `json:"fee"`
	// This field is from variant [SparkPayLightningInvoiceRpcResponseDataUnion].
	IdempotencyKey string `json:"idempotency_key"`
	// This field is from variant [ExportPrivateKeyRpcResponseData].
	EncryptionType string `json:"encryption_type"`
	// This field is from variant [ExportPrivateKeyRpcResponseData].
	RecipientPublicKey string `json:"recipient_public_key"`
	// This field is from variant [ExportPrivateKeyRpcResponseData].
	ExportType string `json:"export_type"`
	JSON       struct {
		Encoding                  respjson.Field
		Signature                 respjson.Field
		SignedTransaction         respjson.Field
		Caip2                     respjson.Field
		Hash                      respjson.Field
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
		ExportType                respjson.Field
		raw                       string
	} `json:"-"`
}

func (r *WalletRpcResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseUnionDataLeaves is an implicit subunion of
// [WalletRpcResponseUnion]. WalletRpcResponseUnionDataLeaves provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WalletRpcResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSparkTransferRpcResponseDataLeaves
// OfSparkPayLightningInvoiceRpcResponseDataObjectLeaves]
type WalletRpcResponseUnionDataLeaves struct {
	// This field will be present if the value is a
	// [[]SparkTransferRpcResponseDataLeaf] instead of an object.
	OfSparkTransferRpcResponseDataLeaves []SparkTransferRpcResponseDataLeaf `json:",inline"`
	// This field will be present if the value is a
	// [[]SparkPayLightningInvoiceRpcResponseDataObjectLeaf] instead of an object.
	OfSparkPayLightningInvoiceRpcResponseDataObjectLeaves []SparkPayLightningInvoiceRpcResponseDataObjectLeaf `json:",inline"`
	JSON                                                  struct {
		OfSparkTransferRpcResponseDataLeaves                  respjson.Field
		OfSparkPayLightningInvoiceRpcResponseDataObjectLeaves respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (r *WalletRpcResponseUnionDataLeaves) UnmarshalJSON(data []byte) error {
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
	return json.Unmarshal(data, &r.WalletUpdateRequestBody)
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
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
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
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
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
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
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
	ChainType WalletImportSupportedChains `json:"chain_type,omitzero" api:"required"`
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
	// Request body for wallet authentication with HPKE-encrypted response.
	WalletAuthenticateRequestBody WalletAuthenticateRequestBody
	paramObj
}

func (r WalletAuthenticateWithJwtParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WalletAuthenticateRequestBody)
}
func (r *WalletAuthenticateWithJwtParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WalletAuthenticateRequestBody)
}

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
	// Request expiry. Value is a Unix timestamp in milliseconds representing the
	// deadline by which the request must be processed.
	PrivyRequestExpiry param.Opt[string] `header:"privy-request-expiry,omitzero" json:"-"`
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
	return json.Unmarshal(data, &r.RawSignInput)
}

type WalletRpcParams struct {
	// Request body for wallet RPC operations, discriminated by method.
	WalletRpcRequestBody WalletRpcRequestBodyUnionParam
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
	return json.Unmarshal(data, &r.WalletRpcRequestBody)
}
