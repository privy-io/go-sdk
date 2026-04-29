package privyclient

import (
	"context"

	"github.com/privy-io/go-sdk/internal/jwtexchange"
	"github.com/privy-io/go-sdk/packages/param"
)

type PrivyKeyQuorumService struct {
	// Directly embed the generated KeyQuorumService to expose all its methods through PrivyKeyQuorumService
	KeyQuorumService

	jwtExchanger           jwtexchange.JwtExchanger
	baseURL                string
	appID                  string
	defaultRequestExpiryMs int64
	requestExpiryEnabled   bool
	logger                 logger
}

// newPrivyKeyQuorumService creates a new wrapped key quorum service.
// This is unexported so only PrivyClient can construct it.
func newPrivyKeyQuorumService(
	service KeyQuorumService,
	jwtExchanger jwtexchange.JwtExchanger,
	baseURL string,
	appID string,
	defaultRequestExpiryMs int64,
	requestExpiryEnabled bool,
	logger logger,
) *PrivyKeyQuorumService {
	return &PrivyKeyQuorumService{
		KeyQuorumService:       service,
		jwtExchanger:           jwtExchanger,
		baseURL:                baseURL,
		appID:                  appID,
		defaultRequestExpiryMs: defaultRequestExpiryMs,
		requestExpiryEnabled:   requestExpiryEnabled,
		logger:                 logger,
	}
}

// Update modifies a key quorum
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - keyQuorumID: The key quorum ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyKeyQuorumService) Update(
	ctx context.Context,
	keyQuorumID string,
	params KeyQuorumUpdateParams,
	opts ...RequestOption,
) (*KeyQuorum, error) {
	options := applyRequestOptions(opts)

	requestExpiry := options.RequestExpiry
	if requestExpiry == nil && s.requestExpiryEnabled {
		requestExpiry = int64Ptr(RequestExpiry(s.defaultRequestExpiryMs))
	}

	prepared, err := prepareRequest(ctx, s.appID, s.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		requestExpiry:        requestExpiry,
		method:               "PATCH",
		url:                  s.baseURL + "/v1/key_quorums/" + keyQuorumID,
		body:                 params,
	})
	if err != nil {
		return nil, err
	}

	if prepared.privyAuthorizationSignature != nil {
		params.PrivyAuthorizationSignature = param.NewOpt(*prepared.privyAuthorizationSignature)
	}
	if prepared.privyRequestExpiry != nil {
		params.PrivyRequestExpiry = param.NewOpt(*prepared.privyRequestExpiry)
	}

	return s.KeyQuorumService.Update(ctx, keyQuorumID, params, options.RequestOptions...)
}

// Delete removes a key quorum
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - keyQuorumID: The key quorum ID to delete
//   - params: The delete parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyKeyQuorumService) Delete(
	ctx context.Context,
	keyQuorumID string,
	params KeyQuorumDeleteParams,
	opts ...RequestOption,
) (*SuccessResponse, error) {
	options := applyRequestOptions(opts)

	requestExpiry := options.RequestExpiry
	if requestExpiry == nil && s.requestExpiryEnabled {
		requestExpiry = int64Ptr(RequestExpiry(s.defaultRequestExpiryMs))
	}

	prepared, err := prepareRequest(ctx, s.appID, s.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		requestExpiry:        requestExpiry,
		method:               "DELETE",
		url:                  s.baseURL + "/v1/key_quorums/" + keyQuorumID,
		body:                 params,
	})
	if err != nil {
		return nil, err
	}

	if prepared.privyAuthorizationSignature != nil {
		params.PrivyAuthorizationSignature = param.NewOpt(*prepared.privyAuthorizationSignature)
	}
	if prepared.privyRequestExpiry != nil {
		params.PrivyRequestExpiry = param.NewOpt(*prepared.privyRequestExpiry)
	}

	return s.KeyQuorumService.Delete(ctx, keyQuorumID, params, options.RequestOptions...)
}
