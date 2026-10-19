package privyclient

import (
	"context"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyWalletPayoutService wraps the generated WalletPayoutService. Fiat is
// replaced with a signing wrapper.
type PrivyWalletPayoutService struct {
	WalletPayoutService

	// Fiat shadows WalletPayoutService.Fiat with signing support.
	Fiat *PrivyWalletPayoutFiatService
}

func newPrivyWalletPayoutService(service WalletPayoutService, wallets *PrivyWalletService) *PrivyWalletPayoutService {
	return &PrivyWalletPayoutService{
		WalletPayoutService: service,
		Fiat:                newPrivyWalletPayoutFiatService(service.Fiat, wallets),
	}
}

// PrivyWalletPayoutFiatService wraps the generated WalletPayoutFiatService
// with automatic authorization signature generation and privy-request-expiry
// handling.
type PrivyWalletPayoutFiatService struct {
	WalletPayoutFiatService
	wallets *PrivyWalletService
}

func newPrivyWalletPayoutFiatService(service WalletPayoutFiatService, wallets *PrivyWalletService) *PrivyWalletPayoutFiatService {
	return &PrivyWalletPayoutFiatService{WalletPayoutFiatService: service, wallets: wallets}
}

// New initiates a fiat payout from the wallet with automatic authorization signature generation.
//
// This method wraps the generated WalletPayoutFiatService.New and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key and request expiry headers
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature, PrivyIdempotencyKey, and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext, WithIdempotencyKey, and WithRequestExpiry
func (s *PrivyWalletPayoutFiatService) New(
	ctx context.Context,
	walletID string,
	params WalletPayoutFiatNewParams,
	opts ...RequestOption,
) (*PayoutResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		idempotencyKey:       options.IdempotencyKey,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/payout/fiat",
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

	return s.WalletPayoutFiatService.New(ctx, walletID, params, options.RequestOptions...)
}
