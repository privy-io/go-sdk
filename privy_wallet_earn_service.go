package privyclient

import (
	"context"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyWalletEarnService wraps the generated WalletEarnService. Ethereum is
// replaced with a signing wrapper.
type PrivyWalletEarnService struct {
	WalletEarnService

	// Ethereum shadows WalletEarnService.Ethereum with signing support.
	Ethereum *PrivyWalletEarnEthereumService
}

func newPrivyWalletEarnService(service WalletEarnService, wallets *PrivyWalletService) *PrivyWalletEarnService {
	return &PrivyWalletEarnService{
		WalletEarnService: service,
		Ethereum:          newPrivyWalletEarnEthereumService(service.Ethereum, wallets),
	}
}

// PrivyWalletEarnEthereumService wraps the generated WalletEarnEthereumService
// with automatic authorization signature generation and privy-request-expiry
// handling for deposits and withdrawals.
type PrivyWalletEarnEthereumService struct {
	WalletEarnEthereumService
	wallets *PrivyWalletService

	// Incentive shadows WalletEarnEthereumService.Incentive with signing support.
	Incentive *PrivyWalletEarnEthereumIncentiveService
}

func newPrivyWalletEarnEthereumService(service WalletEarnEthereumService, wallets *PrivyWalletService) *PrivyWalletEarnEthereumService {
	return &PrivyWalletEarnEthereumService{
		WalletEarnEthereumService: service,
		wallets:                   wallets,
		Incentive:                 newPrivyWalletEarnEthereumIncentiveService(service.Incentive, wallets),
	}
}

// Deposit deposits assets into an ERC-4626 vault with automatic authorization signature generation.
//
// This method wraps the generated WalletEarnEthereumService.Deposit and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key and request expiry headers
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature, PrivyIdempotencyKey, and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext, WithIdempotencyKey, and WithRequestExpiry
func (s *PrivyWalletEarnEthereumService) Deposit(
	ctx context.Context,
	walletID string,
	params WalletEarnEthereumDepositParams,
	opts ...RequestOption,
) (*EarnDepositActionResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		idempotencyKey:       options.IdempotencyKey,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/earn/ethereum/deposit",
		body:                 params,
	})
	if err != nil {
		return nil, err
	}

	if prepared.privyAuthorizationSignature != nil {
		params.PrivyAuthorizationSignature = param.NewOpt(*prepared.privyAuthorizationSignature)
	}
	if prepared.privyIdempotencyKey != nil {
		params.PrivyIdempotencyKey = param.NewOpt(*prepared.privyIdempotencyKey)
	}
	if prepared.privyRequestExpiry != nil {
		params.PrivyRequestExpiry = param.NewOpt(*prepared.privyRequestExpiry)
	}

	return s.WalletEarnEthereumService.Deposit(ctx, walletID, params, options.RequestOptions...)
}

// Withdraw withdraws assets from an ERC-4626 vault with automatic authorization signature generation.
//
// This method wraps the generated WalletEarnEthereumService.Withdraw and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key and request expiry headers
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature, PrivyIdempotencyKey, and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext, WithIdempotencyKey, and WithRequestExpiry
func (s *PrivyWalletEarnEthereumService) Withdraw(
	ctx context.Context,
	walletID string,
	params WalletEarnEthereumWithdrawParams,
	opts ...RequestOption,
) (*EarnWithdrawActionResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		idempotencyKey:       options.IdempotencyKey,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/earn/ethereum/withdraw",
		body:                 params,
	})
	if err != nil {
		return nil, err
	}

	if prepared.privyAuthorizationSignature != nil {
		params.PrivyAuthorizationSignature = param.NewOpt(*prepared.privyAuthorizationSignature)
	}
	if prepared.privyIdempotencyKey != nil {
		params.PrivyIdempotencyKey = param.NewOpt(*prepared.privyIdempotencyKey)
	}
	if prepared.privyRequestExpiry != nil {
		params.PrivyRequestExpiry = param.NewOpt(*prepared.privyRequestExpiry)
	}

	return s.WalletEarnEthereumService.Withdraw(ctx, walletID, params, options.RequestOptions...)
}

// PrivyWalletEarnEthereumIncentiveService wraps the generated
// WalletEarnEthereumIncentiveService with automatic authorization signature
// generation and privy-request-expiry handling for claims.
type PrivyWalletEarnEthereumIncentiveService struct {
	WalletEarnEthereumIncentiveService
	wallets *PrivyWalletService
}

func newPrivyWalletEarnEthereumIncentiveService(service WalletEarnEthereumIncentiveService, wallets *PrivyWalletService) *PrivyWalletEarnEthereumIncentiveService {
	return &PrivyWalletEarnEthereumIncentiveService{WalletEarnEthereumIncentiveService: service, wallets: wallets}
}

// Claim claims incentive rewards for a wallet with automatic authorization signature generation.
//
// This method wraps the generated WalletEarnEthereumIncentiveService.Claim and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key and request expiry headers
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature, PrivyIdempotencyKey, and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext, WithIdempotencyKey, and WithRequestExpiry
func (s *PrivyWalletEarnEthereumIncentiveService) Claim(
	ctx context.Context,
	walletID string,
	params WalletEarnEthereumIncentiveClaimParams,
	opts ...RequestOption,
) (*EarnIncentiveClaimActionResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		idempotencyKey:       options.IdempotencyKey,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/earn/ethereum/incentive/claim",
		body:                 params,
	})
	if err != nil {
		return nil, err
	}

	if prepared.privyAuthorizationSignature != nil {
		params.PrivyAuthorizationSignature = param.NewOpt(*prepared.privyAuthorizationSignature)
	}
	if prepared.privyIdempotencyKey != nil {
		params.PrivyIdempotencyKey = param.NewOpt(*prepared.privyIdempotencyKey)
	}
	if prepared.privyRequestExpiry != nil {
		params.PrivyRequestExpiry = param.NewOpt(*prepared.privyRequestExpiry)
	}

	return s.WalletEarnEthereumIncentiveService.Claim(ctx, walletID, params, options.RequestOptions...)
}
