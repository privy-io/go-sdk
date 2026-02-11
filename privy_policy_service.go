package privyclient

import (
	"context"
	"strings"

	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/internal/jwtexchange"
	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyPolicyService wraps the generated PolicyService with automatic
// authorization signature generation for policy operations.
type PrivyPolicyService struct {
	// Directly embed the generated PolicyService to expose all its methods through PrivyPolicyService
	PolicyService

	jwtExchanger jwtexchange.JwtExchanger
	baseURL      string
	appID        string
	logger       logger
}

// newPrivyPolicyService creates a new wrapped policy service.
// This is unexported so only PrivyClient can construct it.
func newPrivyPolicyService(
	service PolicyService,
	jwtExchanger jwtexchange.JwtExchanger,
	baseURL string,
	appID string,
	logger logger,
) *PrivyPolicyService {
	return &PrivyPolicyService{
		PolicyService: service,
		jwtExchanger:  jwtExchanger,
		baseURL:       baseURL,
		appID:         appID,
		logger:        logger,
	}
}

// Update modifies a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.Update and handles:
//   - Building the authorization signature from an AuthorizationContext
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - policyID: The policy ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyPolicyService) Update(
	ctx context.Context,
	policyID string,
	params PolicyUpdateParams,
	opts ...RpcOption,
) (*Policy, error) {
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
			URL:     s.baseURL + "/v1/policies/" + policyID,
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
	return s.PolicyService.Update(ctx, policyID, params)
}

// Delete removes a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.Delete and handles:
//   - Building the authorization signature from an AuthorizationContext
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - policyID: The policy ID to delete
//   - params: The delete parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyPolicyService) Delete(
	ctx context.Context,
	policyID string,
	params PolicyDeleteParams,
	opts ...RpcOption,
) (*PolicyDeleteResponse, error) {
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
			URL:     s.baseURL + "/v1/policies/" + policyID,
			Body:    "",
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
	return s.PolicyService.Delete(ctx, policyID, params)
}

// NewRule creates a new rule on a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.NewRule and handles:
//   - Building the authorization signature from an AuthorizationContext
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - policyID: The policy ID to add a rule to
//   - params: The new rule parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyPolicyService) NewRule(
	ctx context.Context,
	policyID string,
	params PolicyNewRuleParams,
	opts ...RpcOption,
) (*PolicyNewRuleResponse, error) {
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
			Method:  "POST",
			URL:     s.baseURL + "/v1/policies/" + policyID + "/rules",
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
	return s.PolicyService.NewRule(ctx, policyID, params)
}

// DeleteRule removes a rule from a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.DeleteRule and handles:
//   - Building the authorization signature from an AuthorizationContext
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - ruleID: The rule ID to delete
//   - params: The delete parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyPolicyService) DeleteRule(
	ctx context.Context,
	ruleID string,
	params PolicyDeleteRuleParams,
	opts ...RpcOption,
) (*PolicyDeleteRuleResponse, error) {
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
			URL:     s.baseURL + "/v1/policies/" + params.PolicyID + "/rules/" + ruleID,
			Body:    "",
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
	return s.PolicyService.DeleteRule(ctx, ruleID, params)
}

// UpdateRule modifies a rule on a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.UpdateRule and handles:
//   - Building the authorization signature from an AuthorizationContext
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - ruleID: The rule ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature)
//   - opts: use WithAuthorizationContext
func (s *PrivyPolicyService) UpdateRule(
	ctx context.Context,
	ruleID string,
	params PolicyUpdateRuleParams,
	opts ...RpcOption,
) (*PolicyUpdateRuleResponse, error) {
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
			URL:     s.baseURL + "/v1/policies/" + params.PolicyID + "/rules/" + ruleID,
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
	return s.PolicyService.UpdateRule(ctx, ruleID, params)
}
