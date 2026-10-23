package privyclient

import (
	"context"
	"strconv"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyIntentService wraps the generated IntentService.
// Intents represent pending operations that require dashboard signing
// before they can be executed. Intent requests are not signed by the SDK.
//
// The wrapper sets the "privy-request-expiry" header on every mutating intent
// call. Resolution order, highest priority first:
//  1. Per-call WithRequestExpiry(...), sent verbatim even when request expiry
//     is disabled on the client.
//  2. params.PrivyRequestExpiry, if the caller set it.
//  3. No header, if the client was built with RequestExpiry.Disabled (or the
//     deprecated DisableRequestExpiry).
//  4. Now + RequestExpiry.DefaultIntentMs, or DefaultIntentRequestExpiryMs
//     (72 hours) when unset. The generic RequestExpiry.DefaultMs and the
//     deprecated DefaultRequestExpiryMs do not apply to intents.
type PrivyIntentService struct {
	IntentService
	defaultIntentRequestExpiryMs int64
	requestExpiryEnabled         bool
	logger                       logger
}

// newPrivyIntentService creates a new wrapped intent service.
// This is unexported so only PrivyClient can construct it.
func newPrivyIntentService(service IntentService, defaultIntentRequestExpiryMs int64, requestExpiryEnabled bool, logger logger) *PrivyIntentService {
	return &PrivyIntentService{
		IntentService:                service,
		defaultIntentRequestExpiryMs: defaultIntentRequestExpiryMs,
		requestExpiryEnabled:         requestExpiryEnabled,
		logger:                       logger,
	}
}

// setRequestExpiry sets a params PrivyRequestExpiry field following the
// precedence documented on PrivyIntentService.
func (s *PrivyIntentService) setRequestExpiry(options *requestOptions, field *param.Opt[string]) {
	expiry := options.RequestExpiry
	if expiry == nil {
		if !s.requestExpiryEnabled || !param.IsOmitted(*field) {
			return
		}
		expiry = int64Ptr(RequestExpiry(s.defaultIntentRequestExpiryMs))
	}
	*field = param.NewOpt(strconv.FormatInt(*expiry, 10))
}

// Rpc executes a wallet RPC intent.
func (s *PrivyIntentService) Rpc(
	ctx context.Context,
	walletID string,
	params IntentRpcParams,
	opts ...RequestOption,
) (*RpcIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
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
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
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
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.UpdateWallet(ctx, walletID, params, options.RequestOptions...)
}

// NewPolicyRule executes a new-policy-rule intent.
func (s *PrivyIntentService) NewPolicyRule(
	ctx context.Context,
	policyID string,
	params IntentNewPolicyRuleParams,
	opts ...RequestOption,
) (*RuleMutateIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.NewPolicyRule(ctx, policyID, params, options.RequestOptions...)
}

// DeletePolicyRule executes a delete-policy-rule intent.
func (s *PrivyIntentService) DeletePolicyRule(
	ctx context.Context,
	ruleID string,
	params IntentDeletePolicyRuleParams,
	opts ...RequestOption,
) (*RuleDeleteIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.DeletePolicyRule(ctx, ruleID, params, options.RequestOptions...)
}

// UpdatePolicy executes an update-policy intent.
func (s *PrivyIntentService) UpdatePolicy(
	ctx context.Context,
	policyID string,
	params IntentUpdatePolicyParams,
	opts ...RequestOption,
) (*PolicyUpdateIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.UpdatePolicy(ctx, policyID, params, options.RequestOptions...)
}

// DeletePolicy executes a delete-policy intent.
func (s *PrivyIntentService) DeletePolicy(
	ctx context.Context,
	policyID string,
	params IntentDeletePolicyParams,
	opts ...RequestOption,
) (*PolicyDeleteIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.DeletePolicy(ctx, policyID, params, options.RequestOptions...)
}

// UpdatePolicyRule executes an update-policy-rule intent.
func (s *PrivyIntentService) UpdatePolicyRule(
	ctx context.Context,
	ruleID string,
	params IntentUpdatePolicyRuleParams,
	opts ...RequestOption,
) (*RuleMutateIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.UpdatePolicyRule(ctx, ruleID, params, options.RequestOptions...)
}

// UpdateKeyQuorum executes an update-key-quorum intent.
func (s *PrivyIntentService) UpdateKeyQuorum(
	ctx context.Context,
	keyQuorumID string,
	params IntentUpdateKeyQuorumParams,
	opts ...RequestOption,
) (*KeyQuorumUpdateIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.UpdateKeyQuorum(ctx, keyQuorumID, params, options.RequestOptions...)
}

// DeleteKeyQuorum executes a delete-key-quorum intent.
func (s *PrivyIntentService) DeleteKeyQuorum(
	ctx context.Context,
	keyQuorumID string,
	params IntentDeleteKeyQuorumParams,
	opts ...RequestOption,
) (*KeyQuorumDeleteIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.DeleteKeyQuorum(ctx, keyQuorumID, params, options.RequestOptions...)
}

// RawSign executes a wallet raw-sign intent.
func (s *PrivyIntentService) RawSign(
	ctx context.Context,
	walletID string,
	params IntentRawSignParams,
	opts ...RequestOption,
) (*RawSignIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.RawSign(ctx, walletID, params, options.RequestOptions...)
}

// NewConditionSetItems executes an add-condition-set-items intent.
func (s *PrivyIntentService) NewConditionSetItems(
	ctx context.Context,
	conditionSetID string,
	params IntentNewConditionSetItemsParams,
	opts ...RequestOption,
) (*ConditionSetItemMutateIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.NewConditionSetItems(ctx, conditionSetID, params, options.RequestOptions...)
}

// UpdateConditionSetItems executes a replace-condition-set-items intent.
func (s *PrivyIntentService) UpdateConditionSetItems(
	ctx context.Context,
	conditionSetID string,
	params IntentUpdateConditionSetItemsParams,
	opts ...RequestOption,
) (*ConditionSetItemMutateIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.UpdateConditionSetItems(ctx, conditionSetID, params, options.RequestOptions...)
}

// DeleteConditionSet executes a delete-condition-set intent.
func (s *PrivyIntentService) DeleteConditionSet(
	ctx context.Context,
	conditionSetID string,
	params IntentDeleteConditionSetParams,
	opts ...RequestOption,
) (*ConditionSetDeleteIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.DeleteConditionSet(ctx, conditionSetID, params, options.RequestOptions...)
}

// DeleteConditionSetItem executes a delete-condition-set-item intent.
func (s *PrivyIntentService) DeleteConditionSetItem(
	ctx context.Context,
	conditionSetItemID string,
	params IntentDeleteConditionSetItemParams,
	opts ...RequestOption,
) (*ConditionSetItemDeleteIntentResponse, error) {
	options := applyRequestOptions(opts)
	s.setRequestExpiry(options, &params.PrivyRequestExpiry)
	return s.IntentService.DeleteConditionSetItem(ctx, conditionSetItemID, params, options.RequestOptions...)
}
