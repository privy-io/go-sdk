package privyclient

import (
	"context"
	"strings"

	"github.com/privy-io/go-sdk/authorization"
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

		// Format the request for signing
		payload, err := authorization.FormatRequestForAuthorizationSignature(sigInput)
		if err != nil {
			return nil, err
		}

		// Generate signatures
		signatures, err := authorization.GenerateAuthorizationSignatures(
			ctx,
			*options.AuthorizationContext,
			payload,
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
