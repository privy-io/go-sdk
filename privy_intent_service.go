package privyclient

import (
	"context"
	"strconv"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyIntentService wraps the generated IntentService.
// Intents represent pending operations that require dashboard signing
// before they can be executed.
//
// The wrapper auto-populates the "privy-request-expiry" header on every
// mutating intent call. Resolution order, highest priority first:
//  1. Per-call WithRequestExpiry(...).
//  2. params.PrivyRequestExpiry already set explicitly by the caller.
//  3. DefaultIntentRequestExpiryMs from client options.
//  4. Hardcoded 72 hours.
//
// Intents always resolve to a header value — there is no "disable" flag
// for intent expiry.
type PrivyIntentService struct {
	IntentService
	defaultIntentRequestExpiryMs int64
	logger                       logger
}

// newPrivyIntentService creates a new wrapped intent service.
// This is unexported so only PrivyClient can construct it.
func newPrivyIntentService(service IntentService, defaultIntentRequestExpiryMs int64, logger logger) *PrivyIntentService {
	return &PrivyIntentService{
		IntentService:                service,
		defaultIntentRequestExpiryMs: defaultIntentRequestExpiryMs,
		logger:                       logger,
	}
}

// Rpc executes a wallet RPC intent.
func (s *PrivyIntentService) Rpc(
	ctx context.Context,
	walletID string,
	params IntentRpcParams,
	opts ...RequestOption,
) (*RpcIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.Rpc(ctx, walletID, params, options.RequestOptions...)
}

// Transfer executes a wallet transfer intent.
func (s *PrivyIntentService) Transfer(
	ctx context.Context,
	walletID string,
	params IntentTransferParams,
	opts ...RequestOption,
) (*TransferIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.Transfer(ctx, walletID, params, options.RequestOptions...)
}

// UpdateWallet executes a wallet update intent.
func (s *PrivyIntentService) UpdateWallet(
	ctx context.Context,
	walletID string,
	params IntentUpdateWalletParams,
	opts ...RequestOption,
) (*WalletIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.UpdateWallet(ctx, walletID, params, options.RequestOptions...)
}

// NewPolicyRule executes a new-policy-rule intent.
func (s *PrivyIntentService) NewPolicyRule(
	ctx context.Context,
	policyID string,
	params IntentNewPolicyRuleParams,
	opts ...RequestOption,
) (*RuleIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.NewPolicyRule(ctx, policyID, params, options.RequestOptions...)
}

// DeletePolicyRule executes a delete-policy-rule intent.
func (s *PrivyIntentService) DeletePolicyRule(
	ctx context.Context,
	ruleID string,
	params IntentDeletePolicyRuleParams,
	opts ...RequestOption,
) (*RuleIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.DeletePolicyRule(ctx, ruleID, params, options.RequestOptions...)
}

// UpdatePolicy executes an update-policy intent.
func (s *PrivyIntentService) UpdatePolicy(
	ctx context.Context,
	policyID string,
	params IntentUpdatePolicyParams,
	opts ...RequestOption,
) (*PolicyIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.UpdatePolicy(ctx, policyID, params, options.RequestOptions...)
}

// UpdatePolicyRule executes an update-policy-rule intent.
func (s *PrivyIntentService) UpdatePolicyRule(
	ctx context.Context,
	ruleID string,
	params IntentUpdatePolicyRuleParams,
	opts ...RequestOption,
) (*RuleIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.UpdatePolicyRule(ctx, ruleID, params, options.RequestOptions...)
}

// UpdateKeyQuorum executes an update-key-quorum intent.
func (s *PrivyIntentService) UpdateKeyQuorum(
	ctx context.Context,
	keyQuorumID string,
	params IntentUpdateKeyQuorumParams,
	opts ...RequestOption,
) (*KeyQuorumIntentResponse, error) {
	options := applyRequestOptions(opts)
	if param.IsOmitted(params.PrivyRequestExpiry) {
		expiry := options.RequestExpiry
		if expiry == nil {
			expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
		}
		params.PrivyRequestExpiry = param.NewOpt(strconv.FormatInt(*expiry, 10))
	}
	return s.IntentService.UpdateKeyQuorum(ctx, keyQuorumID, params, options.RequestOptions...)
}
