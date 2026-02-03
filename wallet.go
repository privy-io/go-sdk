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

	"github.com/stainless-sdks/privy-api-client-go/internal/apijson"
	"github.com/stainless-sdks/privy-api-client-go/internal/apiquery"
	"github.com/stainless-sdks/privy-api-client-go/internal/requestconfig"
	"github.com/stainless-sdks/privy-api-client-go/option"
	"github.com/stainless-sdks/privy-api-client-go/packages/pagination"
	"github.com/stainless-sdks/privy-api-client-go/packages/param"
	"github.com/stainless-sdks/privy-api-client-go/packages/respjson"
	"github.com/stainless-sdks/privy-api-client-go/shared/constant"
)

// WalletService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletService] method instead.
type WalletService struct {
	Options      []option.RequestOption
	Transactions WalletTransactionService
	Balance      WalletBalanceService
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
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%s", params.PrivyIdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a wallet's policies or authorization key configuration.
func (r *WalletService) Update(ctx context.Context, walletID string, params WalletUpdateParams, opts ...option.RequestOption) (res *Wallet, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return
	}
	path := fmt.Sprintf("v1/wallets/%s", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
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
func (r *WalletService) _InitImport(ctx context.Context, body Wallet_InitImportParams, opts ...option.RequestOption) (res *WalletInitImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/import/init"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submit a wallet import request.
func (r *WalletService) _SubmitImport(ctx context.Context, body Wallet_SubmitImportParams, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/import/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Obtain a session key to enable wallet access.
func (r *WalletService) AuthenticateWithJwt(ctx context.Context, body WalletAuthenticateWithJwtParams, opts ...option.RequestOption) (res *WalletAuthenticateWithJwtResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/wallets/authenticate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Export a wallet's private key
func (r *WalletService) Export(ctx context.Context, walletID string, params WalletExportParams, opts ...option.RequestOption) (res *WalletExportResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return
	}
	path := fmt.Sprintf("v1/wallets/%s/export", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a wallet by wallet ID.
func (r *WalletService) Get(ctx context.Context, walletID string, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return
	}
	path := fmt.Sprintf("v1/wallets/%s", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sign a message with a wallet by wallet ID.
func (r *WalletService) RawSign(ctx context.Context, walletID string, params WalletRawSignParams, opts ...option.RequestOption) (res *WalletRawSignResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%s", params.PrivyIdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return
	}
	path := fmt.Sprintf("v1/wallets/%s/raw_sign", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Sign a message or transaction with a wallet by wallet ID.
func (r *WalletService) Rpc(ctx context.Context, walletID string, params WalletRpcParams, opts ...option.RequestOption) (res *WalletRpcResponseUnion, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%s", params.PrivyAuthorizationSignature.Value)))
	}
	if !param.IsOmitted(params.PrivyIdempotencyKey) {
		opts = append(opts, option.WithHeader("privy-idempotency-key", fmt.Sprintf("%s", params.PrivyIdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return
	}
	path := fmt.Sprintf("v1/wallets/%s/rpc", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// A wallet managed by Privy's wallet infrastructure.
type Wallet struct {
	// Unique ID of the wallet. This will be the primary identifier when using the
	// wallet in the future.
	ID string `json:"id,required"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletAdditionalSigner `json:"additional_signers,required"`
	// Address of the wallet.
	Address string `json:"address,required"`
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,required"`
	// Unix timestamp of when the wallet was created in milliseconds.
	CreatedAt float64 `json:"created_at,required"`
	// Unix timestamp of when the wallet was exported in milliseconds, if the wallet
	// was exported.
	ExportedAt float64 `json:"exported_at,required"`
	// Unix timestamp of when the wallet was imported in milliseconds, if the wallet
	// was imported.
	ImportedAt float64 `json:"imported_at,required"`
	// The key quorum ID of the owner of the wallet.
	OwnerID string `json:"owner_id,required"`
	// List of policy IDs for policies that are enforced on the wallet.
	PolicyIDs []string `json:"policy_ids,required"`
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
	SignerID string `json:"signer_id,required" format:"cuid2"`
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
type HpkeImportConfigParam struct {
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

func (r HpkeImportConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow HpkeImportConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HpkeImportConfigParam) UnmarshalJSON(data []byte) error {
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

// Executes the EVM `personal_sign` RPC (EIP-191) to sign a message.
//
// The properties Method, Params are required.
type EthereumPersonalSignRpcInputParam struct {
	// Any of "personal_sign".
	Method  EthereumPersonalSignRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumPersonalSignRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                       `json:"address,omitzero"`
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

type EthereumPersonalSignRpcInputMethod string

const (
	EthereumPersonalSignRpcInputMethodPersonalSign EthereumPersonalSignRpcInputMethod = "personal_sign"
)

// The properties Encoding, Message are required.
type EthereumPersonalSignRpcInputParamsParam struct {
	// Any of "utf-8", "hex".
	Encoding EthereumPersonalSignRpcInputParamsEncoding `json:"encoding,omitzero,required"`
	Message  string                                     `json:"message,required"`
	paramObj
}

func (r EthereumPersonalSignRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumPersonalSignRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumPersonalSignRpcInputParamsParam) UnmarshalJSON(data []byte) error {
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
type EthereumSignTransactionRpcInputParam struct {
	// Any of "eth_signTransaction".
	Method  EthereumSignTransactionRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumSignTransactionRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                          `json:"address,omitzero"`
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

type EthereumSignTransactionRpcInputMethod string

const (
	EthereumSignTransactionRpcInputMethodEthSignTransaction EthereumSignTransactionRpcInputMethod = "eth_signTransaction"
)

// The property Transaction is required.
type EthereumSignTransactionRpcInputParamsParam struct {
	Transaction EthereumSignTransactionRpcInputParamsTransactionParam `json:"transaction,omitzero,required"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTransactionRpcInputParamsTransactionParam struct {
	Data                 param.Opt[string]                                                              `json:"data,omitzero"`
	From                 param.Opt[string]                                                              `json:"from,omitzero"`
	To                   param.Opt[string]                                                              `json:"to,omitzero"`
	AuthorizationList    []EthereumSignTransactionRpcInputParamsTransactionAuthorizationListParam       `json:"authorization_list,omitzero"`
	ChainID              EthereumSignTransactionRpcInputParamsTransactionChainIDUnionParam              `json:"chain_id,omitzero"`
	GasLimit             EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionParam             `json:"gas_limit,omitzero"`
	GasPrice             EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionParam             `json:"gas_price,omitzero"`
	MaxFeePerGas         EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam         `json:"max_fee_per_gas,omitzero"`
	MaxPriorityFeePerGas EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam `json:"max_priority_fee_per_gas,omitzero"`
	Nonce                EthereumSignTransactionRpcInputParamsTransactionNonceUnionParam                `json:"nonce,omitzero"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                         `json:"type,omitzero"`
	Value EthereumSignTransactionRpcInputParamsTransactionValueUnionParam `json:"value,omitzero"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParamsTransactionParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParamsTransactionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParamsTransactionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EthereumSignTransactionRpcInputParamsTransactionParam](
		"type", 0, 1, 2, 4,
	)
}

// The properties ChainID, Contract, Nonce, R, S, YParity are required.
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListParam struct {
	ChainID  EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam `json:"chain_id,omitzero,required"`
	Contract string                                                                             `json:"contract,required"`
	Nonce    EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam   `json:"nonce,omitzero,required"`
	R        string                                                                             `json:"r,required"`
	S        string                                                                             `json:"s,required"`
	YParity  float64                                                                            `json:"y_parity,required"`
	paramObj
}

func (r EthereumSignTransactionRpcInputParamsTransactionAuthorizationListParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTransactionRpcInputParamsTransactionAuthorizationListParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionChainIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionChainIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionChainIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionChainIDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionGasPriceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionNonceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionNonceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionNonceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionNonceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignTransactionRpcInputParamsTransactionValueUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignTransactionRpcInputParamsTransactionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignTransactionRpcInputParamsTransactionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignTransactionRpcInputParamsTransactionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type EthereumSignTransactionRpcInputChainType string

const (
	EthereumSignTransactionRpcInputChainTypeEthereum EthereumSignTransactionRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_sendTransaction` RPC to sign and broadcast a transaction.
//
// The properties Caip2, Method, Params are required.
type EthereumSendTransactionRpcInputParam struct {
	Caip2 string `json:"caip2,required"`
	// Any of "eth_sendTransaction".
	Method  EthereumSendTransactionRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumSendTransactionRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                          `json:"address,omitzero"`
	Sponsor param.Opt[bool]                            `json:"sponsor,omitzero"`
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

type EthereumSendTransactionRpcInputMethod string

const (
	EthereumSendTransactionRpcInputMethodEthSendTransaction EthereumSendTransactionRpcInputMethod = "eth_sendTransaction"
)

// The property Transaction is required.
type EthereumSendTransactionRpcInputParamsParam struct {
	Transaction EthereumSendTransactionRpcInputParamsTransactionParam `json:"transaction,omitzero,required"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSendTransactionRpcInputParamsTransactionParam struct {
	Data                 param.Opt[string]                                                              `json:"data,omitzero"`
	From                 param.Opt[string]                                                              `json:"from,omitzero"`
	To                   param.Opt[string]                                                              `json:"to,omitzero"`
	AuthorizationList    []EthereumSendTransactionRpcInputParamsTransactionAuthorizationListParam       `json:"authorization_list,omitzero"`
	ChainID              EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam              `json:"chain_id,omitzero"`
	GasLimit             EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionParam             `json:"gas_limit,omitzero"`
	GasPrice             EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionParam             `json:"gas_price,omitzero"`
	MaxFeePerGas         EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam         `json:"max_fee_per_gas,omitzero"`
	MaxPriorityFeePerGas EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam `json:"max_priority_fee_per_gas,omitzero"`
	Nonce                EthereumSendTransactionRpcInputParamsTransactionNonceUnionParam                `json:"nonce,omitzero"`
	// Any of 0, 1, 2, 4.
	Type  float64                                                         `json:"type,omitzero"`
	Value EthereumSendTransactionRpcInputParamsTransactionValueUnionParam `json:"value,omitzero"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParamsTransactionParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParamsTransactionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParamsTransactionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EthereumSendTransactionRpcInputParamsTransactionParam](
		"type", 0, 1, 2, 4,
	)
}

// The properties ChainID, Contract, Nonce, R, S, YParity are required.
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListParam struct {
	ChainID  EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam `json:"chain_id,omitzero,required"`
	Contract string                                                                             `json:"contract,required"`
	Nonce    EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam   `json:"nonce,omitzero,required"`
	R        string                                                                             `json:"r,required"`
	S        string                                                                             `json:"s,required"`
	YParity  float64                                                                            `json:"y_parity,required"`
	paramObj
}

func (r EthereumSendTransactionRpcInputParamsTransactionAuthorizationListParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSendTransactionRpcInputParamsTransactionAuthorizationListParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionNonceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionNonceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionNonceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionNonceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSendTransactionRpcInputParamsTransactionValueUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSendTransactionRpcInputParamsTransactionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSendTransactionRpcInputParamsTransactionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSendTransactionRpcInputParamsTransactionValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type EthereumSendTransactionRpcInputChainType string

const (
	EthereumSendTransactionRpcInputChainTypeEthereum EthereumSendTransactionRpcInputChainType = "ethereum"
)

// Executes the EVM `eth_signTypedData_v4` RPC (EIP-712) to sign a typed data
// object.
//
// The properties Method, Params are required.
type EthereumSignTypedDataRpcInputParam struct {
	// Any of "eth_signTypedData_v4".
	Method  EthereumSignTypedDataRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumSignTypedDataRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                        `json:"address,omitzero"`
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

type EthereumSignTypedDataRpcInputMethod string

const (
	EthereumSignTypedDataRpcInputMethodEthSignTypedDataV4 EthereumSignTypedDataRpcInputMethod = "eth_signTypedData_v4"
)

// The property TypedData is required.
type EthereumSignTypedDataRpcInputParamsParam struct {
	TypedData EthereumSignTypedDataRpcInputParamsTypedDataParam `json:"typed_data,omitzero,required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domain, Message, PrimaryType, Types are required.
type EthereumSignTypedDataRpcInputParamsTypedDataParam struct {
	Domain      map[string]any                                                     `json:"domain,omitzero,required"`
	Message     map[string]any                                                     `json:"message,omitzero,required"`
	PrimaryType string                                                             `json:"primary_type,required"`
	Types       map[string][]EthereumSignTypedDataRpcInputParamsTypedDataTypeParam `json:"types,omitzero,required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParamsTypedDataParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParamsTypedDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParamsTypedDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type EthereumSignTypedDataRpcInputParamsTypedDataTypeParam struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	paramObj
}

func (r EthereumSignTypedDataRpcInputParamsTypedDataTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignTypedDataRpcInputParamsTypedDataTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignTypedDataRpcInputParamsTypedDataTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignTypedDataRpcInputChainType string

const (
	EthereumSignTypedDataRpcInputChainTypeEthereum EthereumSignTypedDataRpcInputChainType = "ethereum"
)

// Executes an RPC method to hash and sign a UserOperation.
//
// The properties Method, Params are required.
type EthereumSignUserOperationRpcInputParam struct {
	// Any of "eth_signUserOperation".
	Method  EthereumSignUserOperationRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumSignUserOperationRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                            `json:"address,omitzero"`
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

type EthereumSignUserOperationRpcInputMethod string

const (
	EthereumSignUserOperationRpcInputMethodEthSignUserOperation EthereumSignUserOperationRpcInputMethod = "eth_signUserOperation"
)

// The properties ChainID, Contract, UserOperation are required.
type EthereumSignUserOperationRpcInputParamsParam struct {
	ChainID       EthereumSignUserOperationRpcInputParamsChainIDUnionParam  `json:"chain_id,omitzero,required"`
	Contract      string                                                    `json:"contract,required"`
	UserOperation EthereumSignUserOperationRpcInputParamsUserOperationParam `json:"user_operation,omitzero,required"`
	paramObj
}

func (r EthereumSignUserOperationRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSignUserOperationRpcInputParamsChainIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSignUserOperationRpcInputParamsChainIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSignUserOperationRpcInputParamsChainIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSignUserOperationRpcInputParamsChainIDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The properties CallData, CallGasLimit, MaxFeePerGas, MaxPriorityFeePerGas,
// Nonce, Paymaster, PaymasterData, PaymasterPostOpGasLimit,
// PaymasterVerificationGasLimit, PreVerificationGas, Sender, VerificationGasLimit
// are required.
type EthereumSignUserOperationRpcInputParamsUserOperationParam struct {
	CallData                      string `json:"call_data,required"`
	CallGasLimit                  string `json:"call_gas_limit,required"`
	MaxFeePerGas                  string `json:"max_fee_per_gas,required"`
	MaxPriorityFeePerGas          string `json:"max_priority_fee_per_gas,required"`
	Nonce                         string `json:"nonce,required"`
	Paymaster                     string `json:"paymaster,required"`
	PaymasterData                 string `json:"paymaster_data,required"`
	PaymasterPostOpGasLimit       string `json:"paymaster_post_op_gas_limit,required"`
	PaymasterVerificationGasLimit string `json:"paymaster_verification_gas_limit,required"`
	PreVerificationGas            string `json:"pre_verification_gas,required"`
	Sender                        string `json:"sender,required"`
	VerificationGasLimit          string `json:"verification_gas_limit,required"`
	paramObj
}

func (r EthereumSignUserOperationRpcInputParamsUserOperationParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSignUserOperationRpcInputParamsUserOperationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSignUserOperationRpcInputParamsUserOperationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSignUserOperationRpcInputChainType string

const (
	EthereumSignUserOperationRpcInputChainTypeEthereum EthereumSignUserOperationRpcInputChainType = "ethereum"
)

// Signs an EIP-7702 authorization.
//
// The properties Method, Params are required.
type EthereumSign7702AuthorizationRpcInputParam struct {
	// Any of "eth_sign7702Authorization".
	Method  EthereumSign7702AuthorizationRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumSign7702AuthorizationRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                                `json:"address,omitzero"`
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

type EthereumSign7702AuthorizationRpcInputMethod string

const (
	EthereumSign7702AuthorizationRpcInputMethodEthSign7702Authorization EthereumSign7702AuthorizationRpcInputMethod = "eth_sign7702Authorization"
)

// The properties ChainID, Contract are required.
type EthereumSign7702AuthorizationRpcInputParamsParam struct {
	ChainID  EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam `json:"chain_id,omitzero,required"`
	Contract string                                                       `json:"contract,required"`
	Nonce    EthereumSign7702AuthorizationRpcInputParamsNonceUnionParam   `json:"nonce,omitzero"`
	paramObj
}

func (r EthereumSign7702AuthorizationRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSign7702AuthorizationRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSign7702AuthorizationRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EthereumSign7702AuthorizationRpcInputParamsNonceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EthereumSign7702AuthorizationRpcInputParamsNonceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EthereumSign7702AuthorizationRpcInputParamsNonceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EthereumSign7702AuthorizationRpcInputParamsNonceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type EthereumSign7702AuthorizationRpcInputChainType string

const (
	EthereumSign7702AuthorizationRpcInputChainTypeEthereum EthereumSign7702AuthorizationRpcInputChainType = "ethereum"
)

// Signs a raw hash on the secp256k1 curve.
//
// The properties Method, Params are required.
type EthereumSecp256k1SignRpcInputParam struct {
	// Any of "secp256k1_sign".
	Method  EthereumSecp256k1SignRpcInputMethod      `json:"method,omitzero,required"`
	Params  EthereumSecp256k1SignRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                        `json:"address,omitzero"`
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

type EthereumSecp256k1SignRpcInputMethod string

const (
	EthereumSecp256k1SignRpcInputMethodSecp256k1Sign EthereumSecp256k1SignRpcInputMethod = "secp256k1_sign"
)

// The property Hash is required.
type EthereumSecp256k1SignRpcInputParamsParam struct {
	Hash string `json:"hash,required"`
	paramObj
}

func (r EthereumSecp256k1SignRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow EthereumSecp256k1SignRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EthereumSecp256k1SignRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EthereumSecp256k1SignRpcInputChainType string

const (
	EthereumSecp256k1SignRpcInputChainTypeEthereum EthereumSecp256k1SignRpcInputChainType = "ethereum"
)

// Executes the SVM `signTransaction` RPC to sign a transaction.
//
// The properties Method, Params are required.
type SolanaSignTransactionRpcInputParam struct {
	// Any of "signTransaction".
	Method  SolanaSignTransactionRpcInputMethod      `json:"method,omitzero,required"`
	Params  SolanaSignTransactionRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                        `json:"address,omitzero"`
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

type SolanaSignTransactionRpcInputMethod string

const (
	SolanaSignTransactionRpcInputMethodSignTransaction SolanaSignTransactionRpcInputMethod = "signTransaction"
)

// The properties Encoding, Transaction are required.
type SolanaSignTransactionRpcInputParamsParam struct {
	// Any of "base64".
	Encoding    string `json:"encoding,omitzero,required"`
	Transaction string `json:"transaction,required"`
	paramObj
}

func (r SolanaSignTransactionRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignTransactionRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignTransactionRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SolanaSignTransactionRpcInputParamsParam](
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
type SolanaSignAndSendTransactionRpcInputParam struct {
	Caip2 string `json:"caip2,required"`
	// Any of "signAndSendTransaction".
	Method  SolanaSignAndSendTransactionRpcInputMethod      `json:"method,omitzero,required"`
	Params  SolanaSignAndSendTransactionRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                               `json:"address,omitzero"`
	Sponsor param.Opt[bool]                                 `json:"sponsor,omitzero"`
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

type SolanaSignAndSendTransactionRpcInputMethod string

const (
	SolanaSignAndSendTransactionRpcInputMethodSignAndSendTransaction SolanaSignAndSendTransactionRpcInputMethod = "signAndSendTransaction"
)

// The properties Encoding, Transaction are required.
type SolanaSignAndSendTransactionRpcInputParamsParam struct {
	// Any of "base64".
	Encoding    string `json:"encoding,omitzero,required"`
	Transaction string `json:"transaction,required"`
	paramObj
}

func (r SolanaSignAndSendTransactionRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignAndSendTransactionRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignAndSendTransactionRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SolanaSignAndSendTransactionRpcInputParamsParam](
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
type SolanaSignMessageRpcInputParam struct {
	// Any of "signMessage".
	Method  SolanaSignMessageRpcInputMethod      `json:"method,omitzero,required"`
	Params  SolanaSignMessageRpcInputParamsParam `json:"params,omitzero,required"`
	Address param.Opt[string]                    `json:"address,omitzero"`
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

type SolanaSignMessageRpcInputMethod string

const (
	SolanaSignMessageRpcInputMethodSignMessage SolanaSignMessageRpcInputMethod = "signMessage"
)

// The properties Encoding, Message are required.
type SolanaSignMessageRpcInputParamsParam struct {
	// Any of "base64".
	Encoding string `json:"encoding,omitzero,required"`
	Message  string `json:"message,required"`
	paramObj
}

func (r SolanaSignMessageRpcInputParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow SolanaSignMessageRpcInputParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolanaSignMessageRpcInputParamsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SolanaSignMessageRpcInputParamsParam](
		"encoding", "base64",
	)
}

type SolanaSignMessageRpcInputChainType string

const (
	SolanaSignMessageRpcInputChainTypeSolana SolanaSignMessageRpcInputChainType = "solana"
)

// Response to the EVM `eth_signTransaction` RPC.
type EthereumSignTransactionRpcResponse struct {
	Data EthereumSignTransactionRpcResponseData `json:"data,required"`
	// Any of "eth_signTransaction".
	Method EthereumSignTransactionRpcResponseMethod `json:"method,required"`
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
	Encoding          string `json:"encoding,required"`
	SignedTransaction string `json:"signed_transaction,required"`
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
	Data EthereumSendTransactionRpcResponseData `json:"data,required"`
	// Any of "eth_sendTransaction".
	Method EthereumSendTransactionRpcResponseMethod `json:"method,required"`
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
	Caip2              string                                                   `json:"caip2,required"`
	Hash               string                                                   `json:"hash,required"`
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
	ChainID  EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListChainIDUnion `json:"chain_id,required"`
	Contract string                                                                                `json:"contract,required"`
	Nonce    EthereumSendTransactionRpcResponseDataTransactionRequestAuthorizationListNonceUnion   `json:"nonce,required"`
	R        string                                                                                `json:"r,required"`
	S        string                                                                                `json:"s,required"`
	YParity  float64                                                                               `json:"y_parity,required"`
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
	Data EthereumPersonalSignRpcResponseData `json:"data,required"`
	// Any of "personal_sign".
	Method EthereumPersonalSignRpcResponseMethod `json:"method,required"`
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
	Encoding  string `json:"encoding,required"`
	Signature string `json:"signature,required"`
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
	Data EthereumSignTypedDataRpcResponseData `json:"data,required"`
	// Any of "eth_signTypedData_v4".
	Method EthereumSignTypedDataRpcResponseMethod `json:"method,required"`
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
	Encoding  string `json:"encoding,required"`
	Signature string `json:"signature,required"`
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
	Data EthereumSignUserOperationRpcResponseData `json:"data,required"`
	// Any of "eth_signUserOperation".
	Method EthereumSignUserOperationRpcResponseMethod `json:"method,required"`
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
	Encoding  string `json:"encoding,required"`
	Signature string `json:"signature,required"`
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
	Data EthereumSign7702AuthorizationRpcResponseData `json:"data,required"`
	// Any of "eth_sign7702Authorization".
	Method EthereumSign7702AuthorizationRpcResponseMethod `json:"method,required"`
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
	Authorization EthereumSign7702AuthorizationRpcResponseDataAuthorization `json:"authorization,required"`
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
	ChainID  EthereumSign7702AuthorizationRpcResponseDataAuthorizationChainIDUnion `json:"chain_id,required"`
	Contract string                                                                `json:"contract,required"`
	Nonce    EthereumSign7702AuthorizationRpcResponseDataAuthorizationNonceUnion   `json:"nonce,required"`
	R        string                                                                `json:"r,required"`
	S        string                                                                `json:"s,required"`
	YParity  float64                                                               `json:"y_parity,required"`
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
	Data EthereumSecp256k1SignRpcResponseData `json:"data,required"`
	// Any of "secp256k1_sign".
	Method EthereumSecp256k1SignRpcResponseMethod `json:"method,required"`
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
	Encoding  string `json:"encoding,required"`
	Signature string `json:"signature,required"`
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
	Data SolanaSignTransactionRpcResponseData `json:"data,required"`
	// Any of "signTransaction".
	Method SolanaSignTransactionRpcResponseMethod `json:"method,required"`
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
	Encoding          string `json:"encoding,required"`
	SignedTransaction string `json:"signed_transaction,required"`
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
	Data SolanaSignAndSendTransactionRpcResponseData `json:"data,required"`
	// Any of "signAndSendTransaction".
	Method SolanaSignAndSendTransactionRpcResponseMethod `json:"method,required"`
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
	Caip2         string `json:"caip2,required"`
	Hash          string `json:"hash,required"`
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
	Data SolanaSignMessageRpcResponseData `json:"data,required"`
	// Any of "signMessage".
	Method SolanaSignMessageRpcResponseMethod `json:"method,required"`
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
	Encoding  string `json:"encoding,required"`
	Signature string `json:"signature,required"`
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

type WalletInitImportResponse struct {
	// The base64-encoded encryption public key to encrypt the wallet entropy with.
	EncryptionPublicKey string `json:"encryption_public_key,required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType WalletInitImportResponseEncryptionType `json:"encryption_type,required"`
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
	EncryptedAuthorizationKey WalletAuthenticateWithJwtResponseWithEncryptionEncryptedAuthorizationKey `json:"encrypted_authorization_key,required"`
	// The expiration time of the authorization key in seconds since the epoch.
	ExpiresAt float64  `json:"expires_at,required"`
	Wallets   []Wallet `json:"wallets,required"`
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
	Ciphertext string `json:"ciphertext,required"`
	// Base64-encoded ephemeral public key used in the HPKE encryption process.
	// Required for decryption.
	EncapsulatedKey string `json:"encapsulated_key,required"`
	// The encryption type used. Currently only supports HPKE.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,required"`
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
	AuthorizationKey string `json:"authorization_key,required"`
	// The expiration time of the authorization key in seconds since the epoch.
	ExpiresAt float64  `json:"expires_at,required"`
	Wallets   []Wallet `json:"wallets,required"`
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
	Ciphertext string `json:"ciphertext,required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption.
	EncapsulatedKey string `json:"encapsulated_key,required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType WalletExportResponseEncryptionType `json:"encryption_type,required"`
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
	Data WalletRawSignResponseData `json:"data,required"`
	// Any of "raw_sign".
	Method WalletRawSignResponseMethod `json:"method,required"`
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
	Encoding  string `json:"encoding,required"`
	Signature string `json:"signature,required"`
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

// WalletRpcResponseUnion contains all possible properties and values from
// [EthereumPersonalSignRpcResponse], [EthereumSignTypedDataRpcResponse],
// [EthereumSignTransactionRpcResponse], [EthereumSendTransactionRpcResponse],
// [EthereumSignUserOperationRpcResponse],
// [EthereumSign7702AuthorizationRpcResponse], [EthereumSecp256k1SignRpcResponse],
// [SolanaSignMessageRpcResponse], [SolanaSignTransactionRpcResponse],
// [SolanaSignAndSendTransactionRpcResponse].
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
	// [SolanaSignAndSendTransactionRpcResponseData]
	Data WalletRpcResponseUnionData `json:"data"`
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

// anyWalletRpcResponse is implemented by each variant of [WalletRpcResponseUnion]
// to add type safety for the return type of [WalletRpcResponseUnion.AsAny]
type anyWalletRpcResponse interface {
	implWalletRpcResponseUnion()
}

func (EthereumPersonalSignRpcResponse) implWalletRpcResponseUnion()          {}
func (EthereumSignTypedDataRpcResponse) implWalletRpcResponseUnion()         {}
func (EthereumSignTransactionRpcResponse) implWalletRpcResponseUnion()       {}
func (EthereumSendTransactionRpcResponse) implWalletRpcResponseUnion()       {}
func (EthereumSignUserOperationRpcResponse) implWalletRpcResponseUnion()     {}
func (EthereumSign7702AuthorizationRpcResponse) implWalletRpcResponseUnion() {}
func (EthereumSecp256k1SignRpcResponse) implWalletRpcResponseUnion()         {}
func (SolanaSignMessageRpcResponse) implWalletRpcResponseUnion()             {}
func (SolanaSignTransactionRpcResponse) implWalletRpcResponseUnion()         {}
func (SolanaSignAndSendTransactionRpcResponse) implWalletRpcResponseUnion()  {}

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

func (r *WalletRpcResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletNewParams struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero,required"`
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
	SignerID string `json:"signer_id,required" format:"cuid2"`
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

func (u *WalletNewParamsOwnerUnion) asAny() any {
	if !param.IsOmitted(u.OfPublicKeyOwner) {
		return u.OfPublicKeyOwner
	} else if !param.IsOmitted(u.OfUserOwner) {
		return u.OfUserOwner
	}
	return nil
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type WalletNewParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key,required"`
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
	UserID string `json:"user_id,required"`
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
	OwnerID param.Opt[string] `json:"owner_id,omitzero"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// The owner of the resource. If you provide this, do not specify an owner_id as it
	// will be generated automatically. When updating a wallet, you can set the owner
	// to null to remove the owner.
	Owner WalletUpdateParamsOwnerUnion `json:"owner,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletUpdateParamsAdditionalSigner `json:"additional_signers,omitzero"`
	// New policy IDs to enforce on the wallet. Currently, only one policy is supported
	// per wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r WalletUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SignerID is required.
type WalletUpdateParamsAdditionalSigner struct {
	SignerID string `json:"signer_id,required" format:"cuid2"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet.
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r WalletUpdateParamsAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateParamsAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateParamsAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletUpdateParamsOwnerUnion struct {
	OfPublicKeyOwner *WalletUpdateParamsOwnerPublicKeyOwner `json:",omitzero,inline"`
	OfUserOwner      *WalletUpdateParamsOwnerUserOwner      `json:",omitzero,inline"`
	paramUnion
}

func (u WalletUpdateParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicKeyOwner, u.OfUserOwner)
}
func (u *WalletUpdateParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletUpdateParamsOwnerUnion) asAny() any {
	if !param.IsOmitted(u.OfPublicKeyOwner) {
		return u.OfPublicKeyOwner
	} else if !param.IsOmitted(u.OfUserOwner) {
		return u.OfUserOwner
	}
	return nil
}

// The P-256 public key of the owner of the resource, in base64-encoded DER format.
// If you provide this, do not specify an owner_id as it will be generated
// automatically.
//
// The property PublicKey is required.
type WalletUpdateParamsOwnerPublicKeyOwner struct {
	PublicKey string `json:"public_key,required"`
	paramObj
}

func (r WalletUpdateParamsOwnerPublicKeyOwner) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateParamsOwnerPublicKeyOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateParamsOwnerPublicKeyOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user ID of the owner of the resource. The user must already exist, and this
// value must start with "did:privy:". If you provide this, do not specify an
// owner_id as it will be generated automatically.
//
// The property UserID is required.
type WalletUpdateParamsOwnerUserOwner struct {
	UserID string `json:"user_id,required"`
	paramObj
}

func (r WalletUpdateParamsOwnerUserOwner) MarshalJSON() (data []byte, err error) {
	type shadow WalletUpdateParamsOwnerUserOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletUpdateParamsOwnerUserOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type Wallet_InitImportParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. The
	// input for HD wallets.
	OfHD *Wallet_InitImportParamsBodyHD `json:",inline"`
	// This field is a request body variant, only one variant field can be set. The
	// input for private key wallets.
	OfPrivateKey *Wallet_InitImportParamsBodyPrivateKey `json:",inline"`

	paramObj
}

func (u Wallet_InitImportParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHD, u.OfPrivateKey)
}
func (r *Wallet_InitImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The input for HD wallets.
//
// The properties Address, ChainType, EncryptionType, EntropyType, Index are
// required.
type Wallet_InitImportParamsBodyHD struct {
	// The address of the wallet to import.
	Address string `json:"address,required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero,required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero,required"`
	// The index of the wallet to import.
	Index int64 `json:"index,required"`
	// The entropy type of the wallet to import.
	//
	// This field can be elided, and will marshal its zero value as "hd".
	EntropyType constant.HD `json:"entropy_type,required"`
	paramObj
}

func (r Wallet_InitImportParamsBodyHD) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_InitImportParamsBodyHD
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_InitImportParamsBodyHD) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[Wallet_InitImportParamsBodyHD](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[Wallet_InitImportParamsBodyHD](
		"encryption_type", "HPKE",
	)
}

// The input for private key wallets.
//
// The properties Address, ChainType, EncryptionType, EntropyType are required.
type Wallet_InitImportParamsBodyPrivateKey struct {
	// The address of the wallet to import.
	Address string `json:"address,required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero,required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "private-key".
	EntropyType constant.PrivateKey `json:"entropy_type,required"`
	paramObj
}

func (r Wallet_InitImportParamsBodyPrivateKey) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_InitImportParamsBodyPrivateKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_InitImportParamsBodyPrivateKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[Wallet_InitImportParamsBodyPrivateKey](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[Wallet_InitImportParamsBodyPrivateKey](
		"encryption_type", "HPKE",
	)
}

type Wallet_SubmitImportParams struct {
	Wallet            Wallet_SubmitImportParamsWalletUnion        `json:"wallet,omitzero,required"`
	OwnerID           param.Opt[string]                           `json:"owner_id,omitzero" format:"cuid2"`
	Owner             Wallet_SubmitImportParamsOwnerUnion         `json:"owner,omitzero"`
	AdditionalSigners []Wallet_SubmitImportParamsAdditionalSigner `json:"additional_signers,omitzero"`
	PolicyIDs         []string                                    `json:"policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r Wallet_SubmitImportParams) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_SubmitImportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_SubmitImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type Wallet_SubmitImportParamsWalletUnion struct {
	OfHD         *Wallet_SubmitImportParamsWalletHD         `json:",omitzero,inline"`
	OfPrivateKey *Wallet_SubmitImportParamsWalletPrivateKey `json:",omitzero,inline"`
	paramUnion
}

func (u Wallet_SubmitImportParamsWalletUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHD, u.OfPrivateKey)
}
func (u *Wallet_SubmitImportParamsWalletUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *Wallet_SubmitImportParamsWalletUnion) asAny() any {
	if !param.IsOmitted(u.OfHD) {
		return u.OfHD
	} else if !param.IsOmitted(u.OfPrivateKey) {
		return u.OfPrivateKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetIndex() *int64 {
	if vt := u.OfHD; vt != nil {
		return &vt.Index
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetAddress() *string {
	if vt := u.OfHD; vt != nil {
		return (*string)(&vt.Address)
	} else if vt := u.OfPrivateKey; vt != nil {
		return (*string)(&vt.Address)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetChainType() *string {
	if vt := u.OfHD; vt != nil {
		return (*string)(&vt.ChainType)
	} else if vt := u.OfPrivateKey; vt != nil {
		return (*string)(&vt.ChainType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetCiphertext() *string {
	if vt := u.OfHD; vt != nil {
		return (*string)(&vt.Ciphertext)
	} else if vt := u.OfPrivateKey; vt != nil {
		return (*string)(&vt.Ciphertext)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetEncapsulatedKey() *string {
	if vt := u.OfHD; vt != nil {
		return (*string)(&vt.EncapsulatedKey)
	} else if vt := u.OfPrivateKey; vt != nil {
		return (*string)(&vt.EncapsulatedKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetEncryptionType() *string {
	if vt := u.OfHD; vt != nil {
		return (*string)(&vt.EncryptionType)
	} else if vt := u.OfPrivateKey; vt != nil {
		return (*string)(&vt.EncryptionType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetEntropyType() *string {
	if vt := u.OfHD; vt != nil {
		return (*string)(&vt.EntropyType)
	} else if vt := u.OfPrivateKey; vt != nil {
		return (*string)(&vt.EntropyType)
	}
	return nil
}

// Returns a pointer to the underlying variant's HpkeConfig property, if present.
func (u Wallet_SubmitImportParamsWalletUnion) GetHpkeConfig() *HpkeImportConfigParam {
	if vt := u.OfHD; vt != nil {
		return &vt.HpkeConfig
	} else if vt := u.OfPrivateKey; vt != nil {
		return &vt.HpkeConfig
	}
	return nil
}

func init() {
	apijson.RegisterUnion[Wallet_SubmitImportParamsWalletUnion](
		"entropy_type",
		apijson.Discriminator[Wallet_SubmitImportParamsWalletHD]("hd"),
		apijson.Discriminator[Wallet_SubmitImportParamsWalletPrivateKey]("private-key"),
	)
}

// The properties Address, ChainType, Ciphertext, EncapsulatedKey, EncryptionType,
// EntropyType, Index are required.
type Wallet_SubmitImportParamsWalletHD struct {
	// The address of the wallet to import.
	Address string `json:"address,required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero,required"`
	// The encrypted entropy of the wallet to import.
	Ciphertext string `json:"ciphertext,required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption inside the TEE.
	EncapsulatedKey string `json:"encapsulated_key,required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero,required"`
	// The index of the wallet to import.
	Index int64 `json:"index,required"`
	// Optional HPKE configuration for wallet import decryption. These parameters allow
	// importing wallets encrypted by external providers that use different HPKE
	// configurations.
	HpkeConfig HpkeImportConfigParam `json:"hpke_config,omitzero"`
	// The entropy type of the wallet to import.
	//
	// This field can be elided, and will marshal its zero value as "hd".
	EntropyType constant.HD `json:"entropy_type,required"`
	paramObj
}

func (r Wallet_SubmitImportParamsWalletHD) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_SubmitImportParamsWalletHD
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_SubmitImportParamsWalletHD) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[Wallet_SubmitImportParamsWalletHD](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[Wallet_SubmitImportParamsWalletHD](
		"encryption_type", "HPKE",
	)
}

// The properties Address, ChainType, Ciphertext, EncapsulatedKey, EncryptionType,
// EntropyType are required.
type Wallet_SubmitImportParamsWalletPrivateKey struct {
	// The address of the wallet to import.
	Address string `json:"address,required"`
	// The chain type of the wallet to import. Currently supports `ethereum` and
	// `solana`.
	//
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero,required"`
	// The encrypted entropy of the wallet to import.
	Ciphertext string `json:"ciphertext,required"`
	// The base64-encoded encapsulated key that was generated during encryption, for
	// use during decryption inside the TEE.
	EncapsulatedKey string `json:"encapsulated_key,required"`
	// The encryption type of the wallet to import. Currently only supports `HPKE`.
	//
	// Any of "HPKE".
	EncryptionType string `json:"encryption_type,omitzero,required"`
	// Optional HPKE configuration for wallet import decryption. These parameters allow
	// importing wallets encrypted by external providers that use different HPKE
	// configurations.
	HpkeConfig HpkeImportConfigParam `json:"hpke_config,omitzero"`
	// This field can be elided, and will marshal its zero value as "private-key".
	EntropyType constant.PrivateKey `json:"entropy_type,required"`
	paramObj
}

func (r Wallet_SubmitImportParamsWalletPrivateKey) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_SubmitImportParamsWalletPrivateKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_SubmitImportParamsWalletPrivateKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[Wallet_SubmitImportParamsWalletPrivateKey](
		"chain_type", "ethereum", "solana",
	)
	apijson.RegisterFieldValidator[Wallet_SubmitImportParamsWalletPrivateKey](
		"encryption_type", "HPKE",
	)
}

// The property SignerID is required.
type Wallet_SubmitImportParamsAdditionalSigner struct {
	SignerID          string   `json:"signer_id,required" format:"cuid2"`
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r Wallet_SubmitImportParamsAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_SubmitImportParamsAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_SubmitImportParamsAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type Wallet_SubmitImportParamsOwnerUnion struct {
	OfWallet_SubmitImportsOwnerUserID    *Wallet_SubmitImportParamsOwnerUserID    `json:",omitzero,inline"`
	OfWallet_SubmitImportsOwnerPublicKey *Wallet_SubmitImportParamsOwnerPublicKey `json:",omitzero,inline"`
	paramUnion
}

func (u Wallet_SubmitImportParamsOwnerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWallet_SubmitImportsOwnerUserID, u.OfWallet_SubmitImportsOwnerPublicKey)
}
func (u *Wallet_SubmitImportParamsOwnerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *Wallet_SubmitImportParamsOwnerUnion) asAny() any {
	if !param.IsOmitted(u.OfWallet_SubmitImportsOwnerUserID) {
		return u.OfWallet_SubmitImportsOwnerUserID
	} else if !param.IsOmitted(u.OfWallet_SubmitImportsOwnerPublicKey) {
		return u.OfWallet_SubmitImportsOwnerPublicKey
	}
	return nil
}

// The property UserID is required.
type Wallet_SubmitImportParamsOwnerUserID struct {
	UserID string `json:"user_id,required"`
	paramObj
}

func (r Wallet_SubmitImportParamsOwnerUserID) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_SubmitImportParamsOwnerUserID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_SubmitImportParamsOwnerUserID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PublicKey is required.
type Wallet_SubmitImportParamsOwnerPublicKey struct {
	PublicKey string `json:"public_key,required"`
	paramObj
}

func (r Wallet_SubmitImportParamsOwnerPublicKey) MarshalJSON() (data []byte, err error) {
	type shadow Wallet_SubmitImportParamsOwnerPublicKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Wallet_SubmitImportParamsOwnerPublicKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletAuthenticateWithJwtParams struct {
	// The user's JWT, to be used to authenticate the user.
	UserJwt string `json:"user_jwt,required"`
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
	EncryptionType WalletExportParamsEncryptionType `json:"encryption_type,omitzero,required"`
	// The base64-encoded encryption public key to encrypt the wallet private key with.
	RecipientPublicKey string `json:"recipient_public_key,required"`
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
	Params WalletRawSignParamsParamsUnion `json:"params,omitzero,required"`
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

func (u *WalletRawSignParamsParamsUnion) asAny() any {
	if !param.IsOmitted(u.OfHash) {
		return u.OfHash
	} else if !param.IsOmitted(u.OfBytes) {
		return u.OfBytes
	}
	return nil
}

// Sign a pre-computed hash
//
// The property Hash is required.
type WalletRawSignParamsParamsHash struct {
	// The hash to sign. Must start with `0x`.
	Hash string `json:"hash,required"`
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
	Bytes string `json:"bytes,required"`
	// The encoding scheme for the bytes.
	//
	// Any of "utf-8", "hex", "base64".
	Encoding string `json:"encoding,omitzero,required"`
	// The hash function to hash the bytes.
	//
	// Any of "keccak256", "sha256", "blake2b256".
	HashFunction string `json:"hash_function,omitzero,required"`
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

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Executes the EVM `personal_sign` RPC (EIP-191) to sign a message.
	OfPersonalSign *EthereumPersonalSignRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes the EVM `eth_signTypedData_v4` RPC (EIP-712) to sign a typed data
	// object.
	OfEthSignTypedDataV4 *EthereumSignTypedDataRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes the EVM `eth_signTransaction` RPC to sign a transaction.
	OfEthSignTransaction *EthereumSignTransactionRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes an RPC method to hash and sign a UserOperation.
	OfEthSignUserOperation *EthereumSignUserOperationRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes the EVM `eth_sendTransaction` RPC to sign and broadcast a transaction.
	OfEthSendTransaction *EthereumSendTransactionRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Signs
	// an EIP-7702 authorization.
	OfEthSign7702Authorization *EthereumSign7702AuthorizationRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Signs a
	// raw hash on the secp256k1 curve.
	OfSecp256k1Sign *EthereumSecp256k1SignRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes the SVM `signMessage` RPC to sign a message.
	OfSignMessage *SolanaSignMessageRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes the SVM `signTransaction` RPC to sign a transaction.
	OfSignTransaction *SolanaSignTransactionRpcInputParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Executes the SVM `signAndSendTransaction` RPC to sign and broadcast a
	// transaction.
	OfSignAndSendTransaction *SolanaSignAndSendTransactionRpcInputParam `json:",inline"`

	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	paramObj
}

func (u WalletRpcParams) MarshalJSON() ([]byte, error) {
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
func (r *WalletRpcParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
