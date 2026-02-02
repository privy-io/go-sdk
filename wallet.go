// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

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

// Create a new wallet.
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

type Wallet struct {
	// Unique ID of the wallet. This will be the primary identifier when using the
	// wallet in the future.
	ID string `json:"id,required"`
	// Additional signers for the wallet.
	AdditionalSigners []WalletAdditionalSigner `json:"additional_signers,required"`
	// Address of the wallet.
	Address string `json:"address,required"`
	// Chain type of the wallet
	//
	// Any of "solana", "ethereum", "cosmos", "stellar", "sui", "tron",
	// "bitcoin-segwit", "near", "spark", "ton", "starknet", "movement".
	ChainType WalletChainType `json:"chain_type,required"`
	// Unix timestamp of when the wallet was created in milliseconds.
	CreatedAt float64 `json:"created_at,required"`
	// Unix timestamp of when the wallet was exported in milliseconds, if the wallet
	// was exported.
	ExportedAt float64 `json:"exported_at,required"`
	// Unix timestamp of when the wallet was imported in milliseconds, if the wallet
	// was imported.
	ImportedAt float64 `json:"imported_at,required"`
	// List of policy IDs for policies that are enforced on the wallet.
	PolicyIDs []string `json:"policy_ids,required"`
	// The key quorum ID of the owner of the wallet.
	OwnerID string `json:"owner_id"`
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
		PolicyIDs         respjson.Field
		OwnerID           respjson.Field
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
	OverridePolicyIDs []string `json:"override_policy_ids,required"`
	SignerID          string   `json:"signer_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OverridePolicyIDs respjson.Field
		SignerID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletAdditionalSigner) RawJSON() string { return r.JSON.raw }
func (r *WalletAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chain type of the wallet
type WalletChainType string

const (
	WalletChainTypeSolana        WalletChainType = "solana"
	WalletChainTypeEthereum      WalletChainType = "ethereum"
	WalletChainTypeCosmos        WalletChainType = "cosmos"
	WalletChainTypeStellar       WalletChainType = "stellar"
	WalletChainTypeSui           WalletChainType = "sui"
	WalletChainTypeTron          WalletChainType = "tron"
	WalletChainTypeBitcoinSegwit WalletChainType = "bitcoin-segwit"
	WalletChainTypeNear          WalletChainType = "near"
	WalletChainTypeSpark         WalletChainType = "spark"
	WalletChainTypeTon           WalletChainType = "ton"
	WalletChainTypeStarknet      WalletChainType = "starknet"
	WalletChainTypeMovement      WalletChainType = "movement"
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
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

// WalletRpcResponseUnion contains all possible properties and values from
// [WalletRpcResponseSignTransaction], [WalletRpcResponseSignAndSendTransaction],
// [WalletRpcResponseSignMessage], [WalletRpcResponseEthSignTransaction],
// [WalletRpcResponseEthSendTransaction], [WalletRpcResponsePersonalSign],
// [WalletRpcResponseEthSignTypedDataV4],
// [WalletRpcResponseEthSign7702Authorization], [WalletRpcResponseSecp256k1Sign].
//
// Use the [WalletRpcResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WalletRpcResponseUnion struct {
	// This field is a union of [WalletRpcResponseSignTransactionData],
	// [WalletRpcResponseSignAndSendTransactionData],
	// [WalletRpcResponseSignMessageData], [WalletRpcResponseEthSignTransactionData],
	// [WalletRpcResponseEthSendTransactionData], [WalletRpcResponsePersonalSignData],
	// [WalletRpcResponseEthSignTypedDataV4Data],
	// [WalletRpcResponseEthSign7702AuthorizationData],
	// [WalletRpcResponseSecp256k1SignData]
	Data WalletRpcResponseUnionData `json:"data"`
	// Any of "signTransaction", "signAndSendTransaction", "signMessage",
	// "eth_signTransaction", "eth_sendTransaction", "personal_sign",
	// "eth_signTypedData_v4", "eth_sign7702Authorization", "secp256k1_sign".
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

func (WalletRpcResponseSignTransaction) implWalletRpcResponseUnion()          {}
func (WalletRpcResponseSignAndSendTransaction) implWalletRpcResponseUnion()   {}
func (WalletRpcResponseSignMessage) implWalletRpcResponseUnion()              {}
func (WalletRpcResponseEthSignTransaction) implWalletRpcResponseUnion()       {}
func (WalletRpcResponseEthSendTransaction) implWalletRpcResponseUnion()       {}
func (WalletRpcResponsePersonalSign) implWalletRpcResponseUnion()             {}
func (WalletRpcResponseEthSignTypedDataV4) implWalletRpcResponseUnion()       {}
func (WalletRpcResponseEthSign7702Authorization) implWalletRpcResponseUnion() {}
func (WalletRpcResponseSecp256k1Sign) implWalletRpcResponseUnion()            {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WalletRpcResponseUnion.AsAny().(type) {
//	case privyapiclient.WalletRpcResponseSignTransaction:
//	case privyapiclient.WalletRpcResponseSignAndSendTransaction:
//	case privyapiclient.WalletRpcResponseSignMessage:
//	case privyapiclient.WalletRpcResponseEthSignTransaction:
//	case privyapiclient.WalletRpcResponseEthSendTransaction:
//	case privyapiclient.WalletRpcResponsePersonalSign:
//	case privyapiclient.WalletRpcResponseEthSignTypedDataV4:
//	case privyapiclient.WalletRpcResponseEthSign7702Authorization:
//	case privyapiclient.WalletRpcResponseSecp256k1Sign:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WalletRpcResponseUnion) AsAny() anyWalletRpcResponse {
	switch u.Method {
	case "signTransaction":
		return u.AsSignTransaction()
	case "signAndSendTransaction":
		return u.AsSignAndSendTransaction()
	case "signMessage":
		return u.AsSignMessage()
	case "eth_signTransaction":
		return u.AsEthSignTransaction()
	case "eth_sendTransaction":
		return u.AsEthSendTransaction()
	case "personal_sign":
		return u.AsPersonalSign()
	case "eth_signTypedData_v4":
		return u.AsEthSignTypedDataV4()
	case "eth_sign7702Authorization":
		return u.AsEthSign7702Authorization()
	case "secp256k1_sign":
		return u.AsSecp256k1Sign()
	}
	return nil
}

func (u WalletRpcResponseUnion) AsSignTransaction() (v WalletRpcResponseSignTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSignAndSendTransaction() (v WalletRpcResponseSignAndSendTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSignMessage() (v WalletRpcResponseSignMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSignTransaction() (v WalletRpcResponseEthSignTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSendTransaction() (v WalletRpcResponseEthSendTransaction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsPersonalSign() (v WalletRpcResponsePersonalSign) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSignTypedDataV4() (v WalletRpcResponseEthSignTypedDataV4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsEthSign7702Authorization() (v WalletRpcResponseEthSign7702Authorization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseUnion) AsSecp256k1Sign() (v WalletRpcResponseSecp256k1Sign) {
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
	SignedTransaction string `json:"signed_transaction"`
	Caip2             string `json:"caip2"`
	Hash              string `json:"hash"`
	TransactionID     string `json:"transaction_id"`
	Signature         string `json:"signature"`
	// This field is from variant [WalletRpcResponseEthSendTransactionData].
	TransactionRequest WalletRpcResponseEthSendTransactionDataTransactionRequest `json:"transaction_request"`
	// This field is from variant [WalletRpcResponseEthSign7702AuthorizationData].
	Authorization WalletRpcResponseEthSign7702AuthorizationDataAuthorization `json:"authorization"`
	JSON          struct {
		Encoding           respjson.Field
		SignedTransaction  respjson.Field
		Caip2              respjson.Field
		Hash               respjson.Field
		TransactionID      respjson.Field
		Signature          respjson.Field
		TransactionRequest respjson.Field
		Authorization      respjson.Field
		raw                string
	} `json:"-"`
}

func (r *WalletRpcResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSignTransaction struct {
	Data   WalletRpcResponseSignTransactionData `json:"data,required"`
	Method constant.SignTransaction             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseSignTransaction) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSignTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSignTransactionData struct {
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
func (r WalletRpcResponseSignTransactionData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSignTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSignAndSendTransaction struct {
	Data   WalletRpcResponseSignAndSendTransactionData `json:"data,required"`
	Method constant.SignAndSendTransaction             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseSignAndSendTransaction) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSignAndSendTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSignAndSendTransactionData struct {
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
func (r WalletRpcResponseSignAndSendTransactionData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSignAndSendTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSignMessage struct {
	Data   WalletRpcResponseSignMessageData `json:"data,required"`
	Method constant.SignMessage             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseSignMessage) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSignMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSignMessageData struct {
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
func (r WalletRpcResponseSignMessageData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSignMessageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSignTransaction struct {
	Data   WalletRpcResponseEthSignTransactionData `json:"data,required"`
	Method constant.EthSignTransaction             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseEthSignTransaction) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSignTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSignTransactionData struct {
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
func (r WalletRpcResponseEthSignTransactionData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSignTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSendTransaction struct {
	Data   WalletRpcResponseEthSendTransactionData `json:"data,required"`
	Method constant.EthSendTransaction             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseEthSendTransaction) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSendTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSendTransactionData struct {
	Caip2              string                                                    `json:"caip2,required"`
	Hash               string                                                    `json:"hash,required"`
	TransactionID      string                                                    `json:"transaction_id"`
	TransactionRequest WalletRpcResponseEthSendTransactionDataTransactionRequest `json:"transaction_request"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caip2              respjson.Field
		Hash               respjson.Field
		TransactionID      respjson.Field
		TransactionRequest respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseEthSendTransactionData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSendTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSendTransactionDataTransactionRequest struct {
	ChainID              WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion              `json:"chain_id"`
	Data                 string                                                                             `json:"data"`
	From                 string                                                                             `json:"from"`
	GasLimit             WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion             `json:"gas_limit"`
	GasPrice             WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion             `json:"gas_price"`
	MaxFeePerGas         WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion         `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas"`
	Nonce                WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion                `json:"nonce"`
	To                   string                                                                             `json:"to"`
	// Any of 0, 1, 2.
	Type  float64                                                             `json:"type"`
	Value WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
func (r WalletRpcResponseEthSendTransactionDataTransactionRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *WalletRpcResponseEthSendTransactionDataTransactionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion
// contains all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion struct {
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

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSendTransactionDataTransactionRequestValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponsePersonalSign struct {
	Data   WalletRpcResponsePersonalSignData `json:"data,required"`
	Method constant.PersonalSign             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponsePersonalSign) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponsePersonalSign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponsePersonalSignData struct {
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
func (r WalletRpcResponsePersonalSignData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponsePersonalSignData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSignTypedDataV4 struct {
	Data   WalletRpcResponseEthSignTypedDataV4Data `json:"data,required"`
	Method constant.EthSignTypedDataV4             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseEthSignTypedDataV4) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSignTypedDataV4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSignTypedDataV4Data struct {
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
func (r WalletRpcResponseEthSignTypedDataV4Data) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSignTypedDataV4Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSign7702Authorization struct {
	Data   WalletRpcResponseEthSign7702AuthorizationData `json:"data,required"`
	Method constant.EthSign7702Authorization             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseEthSign7702Authorization) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSign7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSign7702AuthorizationData struct {
	Authorization WalletRpcResponseEthSign7702AuthorizationDataAuthorization `json:"authorization,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authorization respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseEthSign7702AuthorizationData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseEthSign7702AuthorizationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseEthSign7702AuthorizationDataAuthorization struct {
	ChainID  WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion `json:"chain_id,required"`
	Contract string                                                                 `json:"contract,required"`
	Nonce    WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion   `json:"nonce,required"`
	R        string                                                                 `json:"r,required"`
	S        string                                                                 `json:"s,required"`
	YParity  float64                                                                `json:"y_parity,required"`
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
func (r WalletRpcResponseEthSign7702AuthorizationDataAuthorization) RawJSON() string {
	return r.JSON.raw
}
func (r *WalletRpcResponseEthSign7702AuthorizationDataAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion struct {
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

func (u WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSign7702AuthorizationDataAuthorizationChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion struct {
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

func (u WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WalletRpcResponseEthSign7702AuthorizationDataAuthorizationNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSecp256k1Sign struct {
	Data   WalletRpcResponseSecp256k1SignData `json:"data,required"`
	Method constant.Secp256k1Sign             `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WalletRpcResponseSecp256k1Sign) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSecp256k1Sign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcResponseSecp256k1SignData struct {
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
func (r WalletRpcResponseSecp256k1SignData) RawJSON() string { return r.JSON.raw }
func (r *WalletRpcResponseSecp256k1SignData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletNewParams struct {
	// Chain type of the wallet
	//
	// Any of "solana", "ethereum", "cosmos", "stellar", "sui", "tron",
	// "bitcoin-segwit", "near", "spark", "ton", "starknet", "movement".
	ChainType WalletNewParamsChainType `json:"chain_type,omitzero,required"`
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

// Chain type of the wallet
type WalletNewParamsChainType string

const (
	WalletNewParamsChainTypeSolana        WalletNewParamsChainType = "solana"
	WalletNewParamsChainTypeEthereum      WalletNewParamsChainType = "ethereum"
	WalletNewParamsChainTypeCosmos        WalletNewParamsChainType = "cosmos"
	WalletNewParamsChainTypeStellar       WalletNewParamsChainType = "stellar"
	WalletNewParamsChainTypeSui           WalletNewParamsChainType = "sui"
	WalletNewParamsChainTypeTron          WalletNewParamsChainType = "tron"
	WalletNewParamsChainTypeBitcoinSegwit WalletNewParamsChainType = "bitcoin-segwit"
	WalletNewParamsChainTypeNear          WalletNewParamsChainType = "near"
	WalletNewParamsChainTypeSpark         WalletNewParamsChainType = "spark"
	WalletNewParamsChainTypeTon           WalletNewParamsChainType = "ton"
	WalletNewParamsChainTypeStarknet      WalletNewParamsChainType = "starknet"
	WalletNewParamsChainTypeMovement      WalletNewParamsChainType = "movement"
)

// The properties OverridePolicyIDs, SignerID are required.
type WalletNewParamsAdditionalSigner struct {
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero,required"`
	SignerID          string   `json:"signer_id,required"`
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

// The P-256 public key of the owner of the resource. If you provide this, do not
// specify an owner_id as it will be generated automatically.
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
	// The key quorum ID to set as the owner of the resource. If you provide this, do
	// not specify an owner.
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

// The properties OverridePolicyIDs, SignerID are required.
type WalletUpdateParamsAdditionalSigner struct {
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero,required"`
	SignerID          string   `json:"signer_id,required"`
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

// The P-256 public key of the owner of the resource. If you provide this, do not
// specify an owner_id as it will be generated automatically.
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
	Limit  param.Opt[float64] `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]  `query:"cursor,omitzero" json:"-"`
	UserID param.Opt[string]  `query:"user_id,omitzero" json:"-"`
	// Any of "cosmos", "stellar", "sui", "aptos", "movement", "tron",
	// "bitcoin-segwit", "near", "ton", "starknet", "spark", "solana", "ethereum".
	ChainType WalletListParamsChainType `query:"chain_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletListParams]'s query parameters as `url.Values`.
func (r WalletListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletListParamsChainType string

const (
	WalletListParamsChainTypeCosmos        WalletListParamsChainType = "cosmos"
	WalletListParamsChainTypeStellar       WalletListParamsChainType = "stellar"
	WalletListParamsChainTypeSui           WalletListParamsChainType = "sui"
	WalletListParamsChainTypeAptos         WalletListParamsChainType = "aptos"
	WalletListParamsChainTypeMovement      WalletListParamsChainType = "movement"
	WalletListParamsChainTypeTron          WalletListParamsChainType = "tron"
	WalletListParamsChainTypeBitcoinSegwit WalletListParamsChainType = "bitcoin-segwit"
	WalletListParamsChainTypeNear          WalletListParamsChainType = "near"
	WalletListParamsChainTypeTon           WalletListParamsChainType = "ton"
	WalletListParamsChainTypeStarknet      WalletListParamsChainType = "starknet"
	WalletListParamsChainTypeSpark         WalletListParamsChainType = "spark"
	WalletListParamsChainTypeSolana        WalletListParamsChainType = "solana"
	WalletListParamsChainTypeEthereum      WalletListParamsChainType = "ethereum"
)

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
	OwnerID           param.Opt[string]                           `json:"owner_id,omitzero"`
	Owner             Wallet_SubmitImportParamsOwnerUnion         `json:"owner,omitzero"`
	AdditionalSigners []Wallet_SubmitImportParamsAdditionalSigner `json:"additional_signers,omitzero"`
	PolicyIDs         []string                                    `json:"policy_ids,omitzero"`
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
	SignerID          string   `json:"signer_id,required"`
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero"`
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
	Params WalletRawSignParamsParams `json:"params,omitzero,required"`
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

type WalletRawSignParamsParams struct {
	// The hash to sign. Must start with `0x`.
	Hash param.Opt[string] `json:"hash,omitzero"`
	paramObj
}

func (r WalletRawSignParamsParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRawSignParamsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRawSignParamsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfEthSignTransaction *WalletRpcParamsBodyEthSignTransaction `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfEthSendTransaction *WalletRpcParamsBodyEthSendTransaction `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPersonalSign *WalletRpcParamsBodyPersonalSign `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfEthSignTypedDataV4 *WalletRpcParamsBodyEthSignTypedDataV4 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfEthSign7702Authorization *WalletRpcParamsBodyEthSign7702Authorization `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSecp256k1Sign *WalletRpcParamsBodySecp256k1Sign `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSignTransaction *WalletRpcParamsBodySignTransaction `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSignAndSendTransaction *WalletRpcParamsBodySignAndSendTransaction `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSignMessage *WalletRpcParamsBodySignMessage `json:",inline"`

	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Idempotency keys ensure API requests are executed only once within a 24-hour
	// window.
	PrivyIdempotencyKey param.Opt[string] `header:"privy-idempotency-key,omitzero" json:"-"`
	paramObj
}

func (u WalletRpcParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEthSignTransaction,
		u.OfEthSendTransaction,
		u.OfPersonalSign,
		u.OfEthSignTypedDataV4,
		u.OfEthSign7702Authorization,
		u.OfSecp256k1Sign,
		u.OfSignTransaction,
		u.OfSignAndSendTransaction,
		u.OfSignMessage)
}
func (r *WalletRpcParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Method, Params are required.
type WalletRpcParamsBodyEthSignTransaction struct {
	Params  WalletRpcParamsBodyEthSignTransactionParams `json:"params,omitzero,required"`
	Address param.Opt[string]                           `json:"address,omitzero"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "eth_signTransaction".
	Method constant.EthSignTransaction `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTransaction) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyEthSignTransaction](
		"chain_type", "ethereum",
	)
}

// The property Transaction is required.
type WalletRpcParamsBodyEthSignTransactionParams struct {
	Transaction WalletRpcParamsBodyEthSignTransactionParamsTransaction `json:"transaction,omitzero,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcParamsBodyEthSignTransactionParamsTransaction struct {
	Data                 param.Opt[string]                                                               `json:"data,omitzero"`
	From                 param.Opt[string]                                                               `json:"from,omitzero"`
	To                   param.Opt[string]                                                               `json:"to,omitzero"`
	ChainID              WalletRpcParamsBodyEthSignTransactionParamsTransactionChainIDUnion              `json:"chain_id,omitzero"`
	GasLimit             WalletRpcParamsBodyEthSignTransactionParamsTransactionGasLimitUnion             `json:"gas_limit,omitzero"`
	GasPrice             WalletRpcParamsBodyEthSignTransactionParamsTransactionGasPriceUnion             `json:"gas_price,omitzero"`
	MaxFeePerGas         WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxFeePerGasUnion         `json:"max_fee_per_gas,omitzero"`
	MaxPriorityFeePerGas WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas,omitzero"`
	Nonce                WalletRpcParamsBodyEthSignTransactionParamsTransactionNonceUnion                `json:"nonce,omitzero"`
	// Any of 0, 1, 2.
	Type  float64                                                          `json:"type,omitzero"`
	Value WalletRpcParamsBodyEthSignTransactionParamsTransactionValueUnion `json:"value,omitzero"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTransactionParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTransactionParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTransactionParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyEthSignTransactionParamsTransaction](
		"type", 0, 1, 2,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRpcParamsBodyEthSignTransactionParamsTransactionChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionChainIDUnion) asAny() any {
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
type WalletRpcParamsBodyEthSignTransactionParamsTransactionGasLimitUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionGasLimitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionGasLimitUnion) asAny() any {
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
type WalletRpcParamsBodyEthSignTransactionParamsTransactionGasPriceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionGasPriceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionGasPriceUnion) asAny() any {
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
type WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxFeePerGasUnion) asAny() any {
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
type WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxPriorityFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxPriorityFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionMaxPriorityFeePerGasUnion) asAny() any {
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
type WalletRpcParamsBodyEthSignTransactionParamsTransactionNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionNonceUnion) asAny() any {
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
type WalletRpcParamsBodyEthSignTransactionParamsTransactionValueUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSignTransactionParamsTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSignTransactionParamsTransactionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The properties Caip2, Method, Params are required.
type WalletRpcParamsBodyEthSendTransaction struct {
	Caip2   string                                      `json:"caip2,required"`
	Params  WalletRpcParamsBodyEthSendTransactionParams `json:"params,omitzero,required"`
	Address param.Opt[string]                           `json:"address,omitzero"`
	Sponsor param.Opt[bool]                             `json:"sponsor,omitzero"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "eth_sendTransaction".
	Method constant.EthSendTransaction `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSendTransaction) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSendTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSendTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyEthSendTransaction](
		"chain_type", "ethereum",
	)
}

// The property Transaction is required.
type WalletRpcParamsBodyEthSendTransactionParams struct {
	Transaction WalletRpcParamsBodyEthSendTransactionParamsTransaction `json:"transaction,omitzero,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSendTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSendTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSendTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcParamsBodyEthSendTransactionParamsTransaction struct {
	Data                 param.Opt[string]                                                               `json:"data,omitzero"`
	From                 param.Opt[string]                                                               `json:"from,omitzero"`
	To                   param.Opt[string]                                                               `json:"to,omitzero"`
	ChainID              WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion              `json:"chain_id,omitzero"`
	GasLimit             WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion             `json:"gas_limit,omitzero"`
	GasPrice             WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion             `json:"gas_price,omitzero"`
	MaxFeePerGas         WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion         `json:"max_fee_per_gas,omitzero"`
	MaxPriorityFeePerGas WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion `json:"max_priority_fee_per_gas,omitzero"`
	Nonce                WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion                `json:"nonce,omitzero"`
	// Any of 0, 1, 2.
	Type  float64                                                          `json:"type,omitzero"`
	Value WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion `json:"value,omitzero"`
	paramObj
}

func (r WalletRpcParamsBodyEthSendTransactionParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSendTransactionParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSendTransactionParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyEthSendTransactionParamsTransaction](
		"type", 0, 1, 2,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion) asAny() any {
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
type WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion) asAny() any {
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
type WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion) asAny() any {
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
type WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion) asAny() any {
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
type WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion) asAny() any {
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
type WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion) asAny() any {
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
type WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The properties Method, Params are required.
type WalletRpcParamsBodyPersonalSign struct {
	Params  WalletRpcParamsBodyPersonalSignParams `json:"params,omitzero,required"`
	Address param.Opt[string]                     `json:"address,omitzero"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as "personal_sign".
	Method constant.PersonalSign `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodyPersonalSign) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyPersonalSign
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyPersonalSign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyPersonalSign](
		"chain_type", "ethereum",
	)
}

// The properties Encoding, Message are required.
type WalletRpcParamsBodyPersonalSignParams struct {
	// Any of "utf-8", "hex".
	Encoding WalletRpcParamsBodyPersonalSignParamsEncoding `json:"encoding,omitzero,required"`
	Message  string                                        `json:"message,required"`
	paramObj
}

func (r WalletRpcParamsBodyPersonalSignParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyPersonalSignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyPersonalSignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletRpcParamsBodyPersonalSignParamsEncoding string

const (
	WalletRpcParamsBodyPersonalSignParamsEncodingUtf8 WalletRpcParamsBodyPersonalSignParamsEncoding = "utf-8"
	WalletRpcParamsBodyPersonalSignParamsEncodingHex  WalletRpcParamsBodyPersonalSignParamsEncoding = "hex"
)

// The properties Method, Params are required.
type WalletRpcParamsBodyEthSignTypedDataV4 struct {
	Params  WalletRpcParamsBodyEthSignTypedDataV4Params `json:"params,omitzero,required"`
	Address param.Opt[string]                           `json:"address,omitzero"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "eth_signTypedData_v4".
	Method constant.EthSignTypedDataV4 `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTypedDataV4) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTypedDataV4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTypedDataV4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyEthSignTypedDataV4](
		"chain_type", "ethereum",
	)
}

// The property TypedData is required.
type WalletRpcParamsBodyEthSignTypedDataV4Params struct {
	TypedData WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedData `json:"typed_data,omitzero,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTypedDataV4Params) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTypedDataV4Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTypedDataV4Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domain, Message, PrimaryType, Types are required.
type WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedData struct {
	Domain      map[string]any                                                        `json:"domain,omitzero,required"`
	Message     map[string]any                                                        `json:"message,omitzero,required"`
	PrimaryType string                                                                `json:"primary_type,required"`
	Types       map[string][]WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedDataType `json:"types,omitzero,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedData) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedDataType struct {
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedDataType) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedDataType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSignTypedDataV4ParamsTypedDataType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Method, Params are required.
type WalletRpcParamsBodyEthSign7702Authorization struct {
	Params  WalletRpcParamsBodyEthSign7702AuthorizationParams `json:"params,omitzero,required"`
	Address param.Opt[string]                                 `json:"address,omitzero"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "eth_sign7702Authorization".
	Method constant.EthSign7702Authorization `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodyEthSign7702Authorization) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSign7702Authorization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSign7702Authorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodyEthSign7702Authorization](
		"chain_type", "ethereum",
	)
}

// The properties ChainID, Contract are required.
type WalletRpcParamsBodyEthSign7702AuthorizationParams struct {
	ChainID  WalletRpcParamsBodyEthSign7702AuthorizationParamsChainIDUnion `json:"chain_id,omitzero,required"`
	Contract string                                                        `json:"contract,required"`
	Nonce    WalletRpcParamsBodyEthSign7702AuthorizationParamsNonceUnion   `json:"nonce,omitzero"`
	paramObj
}

func (r WalletRpcParamsBodyEthSign7702AuthorizationParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodyEthSign7702AuthorizationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodyEthSign7702AuthorizationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WalletRpcParamsBodyEthSign7702AuthorizationParamsChainIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSign7702AuthorizationParamsChainIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSign7702AuthorizationParamsChainIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSign7702AuthorizationParamsChainIDUnion) asAny() any {
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
type WalletRpcParamsBodyEthSign7702AuthorizationParamsNonceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WalletRpcParamsBodyEthSign7702AuthorizationParamsNonceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WalletRpcParamsBodyEthSign7702AuthorizationParamsNonceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WalletRpcParamsBodyEthSign7702AuthorizationParamsNonceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The properties Method, Params are required.
type WalletRpcParamsBodySecp256k1Sign struct {
	Params  WalletRpcParamsBodySecp256k1SignParams `json:"params,omitzero,required"`
	Address param.Opt[string]                      `json:"address,omitzero"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as "secp256k1_sign".
	Method constant.Secp256k1Sign `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodySecp256k1Sign) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySecp256k1Sign
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySecp256k1Sign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySecp256k1Sign](
		"chain_type", "ethereum",
	)
}

// The property Hash is required.
type WalletRpcParamsBodySecp256k1SignParams struct {
	Hash string `json:"hash,required"`
	paramObj
}

func (r WalletRpcParamsBodySecp256k1SignParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySecp256k1SignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySecp256k1SignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Method, Params are required.
type WalletRpcParamsBodySignTransaction struct {
	Params  WalletRpcParamsBodySignTransactionParams `json:"params,omitzero,required"`
	Address param.Opt[string]                        `json:"address,omitzero"`
	// Any of "solana".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as "signTransaction".
	Method constant.SignTransaction `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodySignTransaction) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySignTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySignTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySignTransaction](
		"chain_type", "solana",
	)
}

// The properties Encoding, Transaction are required.
type WalletRpcParamsBodySignTransactionParams struct {
	// Any of "base64".
	Encoding    string `json:"encoding,omitzero,required"`
	Transaction string `json:"transaction,required"`
	paramObj
}

func (r WalletRpcParamsBodySignTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySignTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySignTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySignTransactionParams](
		"encoding", "base64",
	)
}

// The properties Caip2, Method, Params are required.
type WalletRpcParamsBodySignAndSendTransaction struct {
	Caip2   string                                          `json:"caip2,required"`
	Params  WalletRpcParamsBodySignAndSendTransactionParams `json:"params,omitzero,required"`
	Address param.Opt[string]                               `json:"address,omitzero"`
	Sponsor param.Opt[bool]                                 `json:"sponsor,omitzero"`
	// Any of "solana".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "signAndSendTransaction".
	Method constant.SignAndSendTransaction `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodySignAndSendTransaction) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySignAndSendTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySignAndSendTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySignAndSendTransaction](
		"chain_type", "solana",
	)
}

// The properties Encoding, Transaction are required.
type WalletRpcParamsBodySignAndSendTransactionParams struct {
	// Any of "base64".
	Encoding    string `json:"encoding,omitzero,required"`
	Transaction string `json:"transaction,required"`
	paramObj
}

func (r WalletRpcParamsBodySignAndSendTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySignAndSendTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySignAndSendTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySignAndSendTransactionParams](
		"encoding", "base64",
	)
}

// The properties Method, Params are required.
type WalletRpcParamsBodySignMessage struct {
	Params  WalletRpcParamsBodySignMessageParams `json:"params,omitzero,required"`
	Address param.Opt[string]                    `json:"address,omitzero"`
	// Any of "solana".
	ChainType string `json:"chain_type,omitzero"`
	// This field can be elided, and will marshal its zero value as "signMessage".
	Method constant.SignMessage `json:"method,required"`
	paramObj
}

func (r WalletRpcParamsBodySignMessage) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySignMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySignMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySignMessage](
		"chain_type", "solana",
	)
}

// The properties Encoding, Message are required.
type WalletRpcParamsBodySignMessageParams struct {
	// Any of "base64".
	Encoding string `json:"encoding,omitzero,required"`
	Message  string `json:"message,required"`
	paramObj
}

func (r WalletRpcParamsBodySignMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow WalletRpcParamsBodySignMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletRpcParamsBodySignMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WalletRpcParamsBodySignMessageParams](
		"encoding", "base64",
	)
}
