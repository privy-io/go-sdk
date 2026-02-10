package privyclient

import (
	"context"
	"strings"

	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/internal/jwtexchange"
	"github.com/privy-io/go-sdk/packages/param"
)

type PrivyKeyQuorumService struct {
	// Directly embed the generated KeyQuorumService to expose all its methods through PrivyKeyQuorumService
	KeyQuorumService

	jwtExchanger jwtexchange.JwtExchanger
	baseURL      string
	appID        string
	logger       logger
}

// newPrivyKeyQuorumService creates a new wrapped key quorum service.
// This is unexported so only PrivyClient can construct it.
func newPrivyKeyQuorumService(
	service KeyQuorumService,
	jwtExchanger jwtexchange.JwtExchanger,
	baseURL string,
	appID string,
	logger logger,
) *PrivyKeyQuorumService {
	return &PrivyKeyQuorumService{
		KeyQuorumService: service,
		jwtExchanger:     jwtExchanger,
		baseURL:          baseURL,
		appID:            appID,
		logger:           logger,
	}
}

// Update modifies a key quorum
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - keyQuorumID: The key quorum ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyKeyQuorumService) Update(
	ctx context.Context,
	keyQuorumID string,
	params KeyQuorumUpdateParams,
	opts ...RpcOption,
) (*KeyQuorum, error) {
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
			URL:     s.baseURL + "/v1/key_quorums/" + keyQuorumID,
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
	return s.KeyQuorumService.Update(ctx, keyQuorumID, params)
}

// Delete removes a key quorum
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - keyQuorumID: The key quorum ID to delete
//   - params: The delete parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyKeyQuorumService) Delete(
	ctx context.Context,
	keyQuorumID string,
	params KeyQuorumDeleteParams,
	opts ...RpcOption,
) (*KeyQuorumDeleteResponse, error) {
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
			Method:  "DELETE",
			URL:     s.baseURL + "/v1/key_quorums/" + keyQuorumID,
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
	return s.KeyQuorumService.Delete(ctx, keyQuorumID, params)
}
