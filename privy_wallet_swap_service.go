package privyclient

import (
	"context"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyWalletSwapService wraps the generated WalletSwapService with automatic
// authorization signature generation and privy-request-expiry handling.
type PrivyWalletSwapService struct {
	WalletSwapService
	wallets *PrivyWalletService
}

func newPrivyWalletSwapService(service WalletSwapService, wallets *PrivyWalletService) *PrivyWalletSwapService {
	return &PrivyWalletSwapService{WalletSwapService: service, wallets: wallets}
}

// Execute executes a token swap within a wallet with automatic authorization signature generation.
//
// This method wraps the generated WalletSwapService.Execute and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key and request expiry headers
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature, PrivyIdempotencyKey, and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext, WithIdempotencyKey, and WithRequestExpiry
func (s *PrivyWalletSwapService) Execute(
	ctx context.Context,
	walletID string,
	params WalletSwapExecuteParams,
	opts ...RequestOption,
) (*SwapActionResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		idempotencyKey:       options.IdempotencyKey,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/swap",
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

	return s.WalletSwapService.Execute(ctx, walletID, params, options.RequestOptions...)
}

// Quote gets a price quote for a token swap with automatic authorization signature generation.
//
// This method wraps the generated WalletSwapService.Quote and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the request expiry header
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext and WithRequestExpiry
func (s *PrivyWalletSwapService) Quote(
	ctx context.Context,
	walletID string,
	params WalletSwapQuoteParams,
	opts ...RequestOption,
) (*SwapQuoteResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/swap/quote",
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

	return s.WalletSwapService.Quote(ctx, walletID, params, options.RequestOptions...)
}
