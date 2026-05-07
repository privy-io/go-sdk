package privyclient

import (
	"context"

	"github.com/privy-io/go-sdk/internal/jwtexchange"
	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyPolicyService wraps the generated PolicyService with automatic
// authorization signature generation for policy operations.
type PrivyPolicyService struct {
	// Directly embed the generated PolicyService to expose all its methods through PrivyPolicyService
	PolicyService

	jwtExchanger           jwtexchange.JwtExchanger
	baseURL                string
	appID                  string
	defaultRequestExpiryMs int64
	requestExpiryEnabled   bool
	logger                 logger
}

// newPrivyPolicyService creates a new wrapped policy service.
// This is unexported so only PrivyClient can construct it.
func newPrivyPolicyService(
	service PolicyService,
	jwtExchanger jwtexchange.JwtExchanger,
	baseURL string,
	appID string,
	defaultRequestExpiryMs int64,
	requestExpiryEnabled bool,
	logger logger,
) *PrivyPolicyService {
	return &PrivyPolicyService{
		PolicyService:          service,
		jwtExchanger:           jwtExchanger,
		baseURL:                baseURL,
		appID:                  appID,
		defaultRequestExpiryMs: defaultRequestExpiryMs,
		requestExpiryEnabled:   requestExpiryEnabled,
		logger:                 logger,
	}
}

// Update modifies a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.Update and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the request expiry header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - policyID: The policy ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyPolicyService) Update(
	ctx context.Context,
	policyID string,
	params PolicyUpdateParams,
	opts ...RequestOption,
) (*Policy, error) {
	options := applyRequestOptions(opts)

	requestExpiry := options.RequestExpiry
	if requestExpiry == nil && s.requestExpiryEnabled {
		requestExpiry = int64Ptr(RequestExpiry(s.defaultRequestExpiryMs))
	}

	prepared, err := prepareRequest(ctx, s.appID, s.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		requestExpiry:        requestExpiry,
		method:               "PATCH",
		url:                  s.baseURL + "/v1/policies/" + policyID,
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

	return s.PolicyService.Update(ctx, policyID, params, options.RequestOptions...)
}

// Delete removes a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.Delete and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the request expiry header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - policyID: The policy ID to delete
//   - params: The delete parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyPolicyService) Delete(
	ctx context.Context,
	policyID string,
	params PolicyDeleteParams,
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
		url:                  s.baseURL + "/v1/policies/" + policyID,
		body:                 "",
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

	return s.PolicyService.Delete(ctx, policyID, params, options.RequestOptions...)
}

// NewRule creates a new rule on a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.NewRule and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the request expiry header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - policyID: The policy ID to add a rule to
//   - params: The new rule parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyPolicyService) NewRule(
	ctx context.Context,
	policyID string,
	params PolicyNewRuleParams,
	opts ...RequestOption,
) (*PolicyRuleResponse, error) {
	options := applyRequestOptions(opts)

	requestExpiry := options.RequestExpiry
	if requestExpiry == nil && s.requestExpiryEnabled {
		requestExpiry = int64Ptr(RequestExpiry(s.defaultRequestExpiryMs))
	}

	prepared, err := prepareRequest(ctx, s.appID, s.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		requestExpiry:        requestExpiry,
		method:               "POST",
		url:                  s.baseURL + "/v1/policies/" + policyID + "/rules",
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

	return s.PolicyService.NewRule(ctx, policyID, params, options.RequestOptions...)
}

// DeleteRule removes a rule from a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.DeleteRule and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the request expiry header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - ruleID: The rule ID to delete
//   - params: The delete parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyPolicyService) DeleteRule(
	ctx context.Context,
	ruleID string,
	params PolicyDeleteRuleParams,
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
		url:                  s.baseURL + "/v1/policies/" + params.PolicyID + "/rules/" + ruleID,
		body:                 "",
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

	return s.PolicyService.DeleteRule(ctx, ruleID, params, options.RequestOptions...)
}

// UpdateRule modifies a rule on a policy with automatic authorization signature generation.
//
// This method wraps the generated PolicyService.UpdateRule and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the request expiry header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - ruleID: The rule ID to update
//   - params: The update parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyPolicyService) UpdateRule(
	ctx context.Context,
	ruleID string,
	params PolicyUpdateRuleParams,
	opts ...RequestOption,
) (*PolicyRuleResponse, error) {
	options := applyRequestOptions(opts)

	requestExpiry := options.RequestExpiry
	if requestExpiry == nil && s.requestExpiryEnabled {
		requestExpiry = int64Ptr(RequestExpiry(s.defaultRequestExpiryMs))
	}

	prepared, err := prepareRequest(ctx, s.appID, s.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		requestExpiry:        requestExpiry,
		method:               "PATCH",
		url:                  s.baseURL + "/v1/policies/" + params.PolicyID + "/rules/" + ruleID,
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

	return s.PolicyService.UpdateRule(ctx, ruleID, params, options.RequestOptions...)
}
