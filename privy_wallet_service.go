package privyclient

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/internal/hpke"
	"github.com/privy-io/go-sdk/internal/jwtexchange"
	"github.com/privy-io/go-sdk/packages/param"
)

// applyRpcOptions applies the given options and returns the resulting rpcOptions.
func applyRpcOptions(opts []RpcOption) *rpcOptions {
	options := &rpcOptions{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// PrivyWalletService wraps the generated WalletService with automatic
// authorization signature generation for RPC calls.
type PrivyWalletService struct {
	// Directly embed the generated WalletService to expose all its methods through PrivyWalletService
	WalletService

	jwtExchanger jwtexchange.JwtExchanger // For exchanging JWTs to auth keys
	baseURL      string                   // API base URL
	appID        string                   // App ID for headers
	logger       logger

	// Ethereum provides convenience methods for Ethereum wallet operations.
	Ethereum *PrivyEthereumWalletService

	// Solana provides convenience methods for Solana wallet operations.
	Solana *PrivySolanaWalletService
}

// newPrivyWalletService creates a new wrapped wallet service.
// This is unexported so only PrivyClient can construct it.
func newPrivyWalletService(
	service WalletService,
	jwtExchanger jwtexchange.JwtExchanger,
	baseURL string,
	appID string,
	logger logger,
) *PrivyWalletService {
	s := &PrivyWalletService{
		WalletService: service,
		jwtExchanger:  jwtExchanger,
		baseURL:       baseURL,
		appID:         appID,
		logger:        logger,
	}
	s.Ethereum = newPrivyEthereumWalletService(s)
	s.Solana = newPrivySolanaWalletService(s)
	return s
}

// Rpc executes an RPC method on a wallet with automatic authorization signature generation.
//
// This method wraps the generated WalletService.Rpc and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID to execute the RPC on
//   - params: The RPC parameters (callers can skip PrivyAuthorizationSignature and PrivyIdempotencyKey)
//   - opts: use WithAuthorizationContext and WithIdempotencyKey
func (s *PrivyWalletService) Rpc(
	ctx context.Context,
	walletID string,
	params WalletRpcParams,
	opts ...RpcOption,
) (*WalletRpcResponseUnion, error) {
	options := applyRpcOptions(opts)

	// Set idempotency key if provided
	if options.IdempotencyKey != "" {
		params.PrivyIdempotencyKey = param.NewOpt(options.IdempotencyKey)
	}

	// Generate authorization signature if context is provided
	if options.AuthorizationContext != nil {
		// Build headers map
		headers := map[string]string{
			"privy-app-id": s.appID,
		}
		if options.IdempotencyKey != "" {
			headers["privy-idempotency-key"] = options.IdempotencyKey
		}

		// Build signature input
		sigInput := authorization.WalletApiRequestSignatureInput{
			Version: 1,
			Method:  "POST",
			URL:     s.baseURL + "/v1/wallets/" + walletID + "/rpc",
			Body:    params,
			Headers: headers,
		}

		// Generate signatures
		signatures, err := authorization.GenerateAuthorizationSignaturesForRequest(
			ctx,
			*options.AuthorizationContext,
			sigInput,
			s.jwtExchanger,
		)
		if err != nil {
			return nil, err
		}

		// Set the authorization header
		if len(signatures) > 0 {
			params.PrivyAuthorizationSignature = param.NewOpt(strings.Join(signatures, ","))
		}
	}

	// Call the underlying service
	return s.WalletService.Rpc(ctx, walletID, params)
}

// Update modifies a wallet with automatic authorization signature generation.
//
// This method wraps the generated WalletService.Update and handles:
//   - Building the authorization signature from an AuthorizationContext
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyWalletService) Update(
	ctx context.Context,
	walletID string,
	params WalletUpdateParams,
	opts ...RpcOption,
) (*Wallet, error) {
	options := applyRpcOptions(opts)

	// Generate authorization signature if context is provided
	if options.AuthorizationContext != nil {
		// Build headers map
		headers := map[string]string{
			"privy-app-id": s.appID,
		}

		// Build signature input
		sigInput := authorization.WalletApiRequestSignatureInput{
			Version: 1,
			Method:  "PATCH",
			URL:     s.baseURL + "/v1/wallets/" + walletID,
			Body:    params,
			Headers: headers,
		}

		// Generate signatures
		signatures, err := authorization.GenerateAuthorizationSignaturesForRequest(
			ctx,
			*options.AuthorizationContext,
			sigInput,
			s.jwtExchanger,
		)
		if err != nil {
			return nil, err
		}

		// Set the authorization header
		if len(signatures) > 0 {
			params.PrivyAuthorizationSignature = param.NewOpt(strings.Join(signatures, ","))
		}
	}

	// Call the underlying service
	return s.WalletService.Update(ctx, walletID, params)
}

// RawSign signs a hash or bytes with a wallet, with automatic authorization signature generation.
//
// This method wraps the generated WalletService.RawSign and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID to sign with
//   - params: The raw sign parameters (callers can skip PrivyAuthorizationSignature and PrivyIdempotencyKey)
//   - opts: use WithAuthorizationContext and WithIdempotencyKey
func (s *PrivyWalletService) RawSign(
	ctx context.Context,
	walletID string,
	params WalletRawSignParams,
	opts ...RpcOption,
) (*WalletRawSignResponse, error) {
	options := applyRpcOptions(opts)

	// Set idempotency key if provided
	if options.IdempotencyKey != "" {
		params.PrivyIdempotencyKey = param.NewOpt(options.IdempotencyKey)
	}

	// Generate authorization signature if context is provided
	if options.AuthorizationContext != nil {
		// Build headers map
		headers := map[string]string{
			"privy-app-id": s.appID,
		}
		if options.IdempotencyKey != "" {
			headers["privy-idempotency-key"] = options.IdempotencyKey
		}

		// Build signature input
		sigInput := authorization.WalletApiRequestSignatureInput{
			Version: 1,
			Method:  "POST",
			URL:     s.baseURL + "/v1/wallets/" + walletID + "/raw_sign",
			Body:    params,
			Headers: headers,
		}

		// Generate signatures
		signatures, err := authorization.GenerateAuthorizationSignaturesForRequest(
			ctx,
			*options.AuthorizationContext,
			sigInput,
			s.jwtExchanger,
		)
		if err != nil {
			return nil, err
		}

		// Set the authorization header
		if len(signatures) > 0 {
			params.PrivyAuthorizationSignature = param.NewOpt(strings.Join(signatures, ","))
		}
	}

	// Call the underlying service
	return s.WalletService.RawSign(ctx, walletID, params)
}

// WalletImportParams contains the parameters for importing a wallet.
// The caller provides wallet metadata and the raw private key; the SDK handles
// HPKE encryption and the two-step InitImport/SubmitImport flow internally.
type WalletImportParams struct {
	// Wallet contains the wallet details and private key to import.
	Wallet WalletImportParamsWalletUnion
	// Owner optionally sets the owner of the imported wallet.
	Owner WalletSubmitImportParamsOwnerUnion
	// OwnerID optionally sets the owner ID (key quorum) of the imported wallet.
	OwnerID param.Opt[string]
	// AdditionalSigners optionally sets additional signers for the imported wallet.
	AdditionalSigners []WalletSubmitImportParamsAdditionalSigner
	// PolicyIDs optionally sets policy IDs for the imported wallet.
	PolicyIDs []string
}

// WalletImportParamsWalletUnion is a union of wallet import variants.
// Exactly one field must be set.
type WalletImportParamsWalletUnion struct {
	OfHD         *WalletImportParamsWalletHD
	OfPrivateKey *WalletImportParamsWalletPrivateKey
}

// WalletImportParamsWalletHD contains details for importing an HD wallet.
type WalletImportParamsWalletHD struct {
	Address    string
	ChainType  string
	Index      int64
	PrivateKey []byte // raw entropy bytes (e.g. BIP39 seed phrase as UTF-8)
}

// WalletImportParamsWalletPrivateKey contains details for importing a private key wallet.
type WalletImportParamsWalletPrivateKey struct {
	Address    string
	ChainType  string
	PrivateKey []byte // raw private key bytes
}

// Import imports a wallet by orchestrating the two-step InitImport/SubmitImport
// flow with automatic HPKE encryption of the private key material.
func (s *PrivyWalletService) Import(ctx context.Context, params WalletImportParams) (*Wallet, error) {
	sender := hpke.NewHpkeSender()

	// Determine wallet variant and build InitImport params
	var initParams WalletInitImportParams
	var privateKeyBytes []byte

	switch {
	case params.Wallet.OfHD != nil:
		hd := params.Wallet.OfHD
		privateKeyBytes = hd.PrivateKey
		initParams = WalletInitImportParams{
			OfHD: &WalletInitImportParamsBodyHD{
				Address:        hd.Address,
				ChainType:      hd.ChainType,
				EncryptionType: "HPKE",
				Index:          hd.Index,
			},
		}
	case params.Wallet.OfPrivateKey != nil:
		pk := params.Wallet.OfPrivateKey
		privateKeyBytes = pk.PrivateKey
		initParams = WalletInitImportParams{
			OfPrivateKey: &WalletInitImportParamsBodyPrivateKey{
				Address:        pk.Address,
				ChainType:      pk.ChainType,
				EncryptionType: "HPKE",
			},
		}
	default:
		return nil, fmt.Errorf("wallet import params must have either OfHD or OfPrivateKey set")
	}

	if len(privateKeyBytes) == 0 {
		return nil, fmt.Errorf("private key must not be empty")
	}

	// Step 1: InitImport to get the server's encryption public key
	initResp, err := s.WalletService.InitImport(ctx, initParams)
	if err != nil {
		return nil, fmt.Errorf("failed to init import: %w", err)
	}

	if initResp.EncryptionType != WalletInitImportResponseEncryptionTypeHpke {
		return nil, fmt.Errorf("unexpected encryption type: %s", initResp.EncryptionType)
	}

	// Step 2: Decode the server's public key and encrypt the private key.
	// The server returns the raw uncompressed EC point bytes (not SPKI), base64-encoded.
	recipientPubKey, err := base64.StdEncoding.DecodeString(initResp.EncryptionPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode encryption public key: %w", err)
	}

	encapsulatedKey, ciphertext, err := sender.Encrypt(recipientPubKey, privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt private key: %w", err)
	}

	encKeyB64 := base64.StdEncoding.EncodeToString(encapsulatedKey)
	ctB64 := base64.StdEncoding.EncodeToString(ciphertext)

	// Step 3: Build and submit the import
	var submitParams WalletSubmitImportParams
	switch {
	case params.Wallet.OfHD != nil:
		hd := params.Wallet.OfHD
		submitParams.Wallet = WalletSubmitImportParamsWalletUnion{
			OfHD: &WalletSubmitImportParamsWalletHD{
				Address:         hd.Address,
				ChainType:       hd.ChainType,
				Ciphertext:      ctB64,
				EncapsulatedKey: encKeyB64,
				EncryptionType:  "HPKE",
				Index:           hd.Index,
			},
		}
	case params.Wallet.OfPrivateKey != nil:
		pk := params.Wallet.OfPrivateKey
		submitParams.Wallet = WalletSubmitImportParamsWalletUnion{
			OfPrivateKey: &WalletSubmitImportParamsWalletPrivateKey{
				Address:         pk.Address,
				ChainType:       pk.ChainType,
				Ciphertext:      ctB64,
				EncapsulatedKey: encKeyB64,
				EncryptionType:  "HPKE",
			},
		}
	}

	submitParams.Owner = params.Owner
	submitParams.OwnerID = params.OwnerID
	submitParams.AdditionalSigners = params.AdditionalSigners
	submitParams.PolicyIDs = params.PolicyIDs

	return s.WalletService.SubmitImport(ctx, submitParams)
}

// WalletExportResult contains the decrypted private key from a wallet export operation.
type WalletExportResult struct {
	// PrivateKey is the decrypted wallet private key.
	PrivateKey string
}

// Export exports a wallet's private key, handling HPKE key exchange automatically
// for an end-to-end encrypted flow.
//
// This method wraps the generated WalletService.Export and handles:
//   - Generating an ephemeral HPKE keypair for encryption
//   - Building the authorization signature from an AuthorizationContext
//   - Decrypting the response to extract the plaintext private key
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID to export
//   - opts: use WithAuthorizationContext for owned wallets
func (s *PrivyWalletService) Export(
	ctx context.Context,
	walletID string,
	opts ...RpcOption,
) (*WalletExportResult, error) {
	options := applyRpcOptions(opts)

	recipient, err := hpke.NewHpkeRecipient()
	if err != nil {
		return nil, fmt.Errorf("failed to generate HPKE keypair: %w", err)
	}

	params := WalletExportParams{
		EncryptionType:     WalletExportParamsEncryptionTypeHpke,
		RecipientPublicKey: base64.StdEncoding.EncodeToString(recipient.PublicKeySpki),
	}

	if options.AuthorizationContext != nil {
		headers := map[string]string{
			"privy-app-id": s.appID,
		}

		sigInput := authorization.WalletApiRequestSignatureInput{
			Version: 1,
			Method:  "POST",
			URL:     s.baseURL + "/v1/wallets/" + walletID + "/export",
			Body:    params,
			Headers: headers,
		}

		signatures, err := authorization.GenerateAuthorizationSignaturesForRequest(
			ctx,
			*options.AuthorizationContext,
			sigInput,
			s.jwtExchanger,
		)
		if err != nil {
			return nil, err
		}

		if len(signatures) > 0 {
			params.PrivyAuthorizationSignature = param.NewOpt(strings.Join(signatures, ","))
		}
	}

	response, err := s.WalletService.Export(ctx, walletID, params)
	if err != nil {
		return nil, err
	}

	encapKey, err := base64.StdEncoding.DecodeString(response.EncapsulatedKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode encapsulated key: %w", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(response.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	plaintext, err := recipient.Decrypt(encapKey, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt exported key: %w", err)
	}

	return &WalletExportResult{PrivateKey: string(plaintext)}, nil
}
