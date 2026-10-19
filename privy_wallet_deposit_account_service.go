package privyclient

import (
	"context"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyWalletDepositAccountService wraps the generated
// WalletDepositAccountService. Crypto is replaced with a signing wrapper; the
// other sub-services are promoted unchanged.
type PrivyWalletDepositAccountService struct {
	WalletDepositAccountService

	// Crypto shadows WalletDepositAccountService.Crypto with signing support.
	Crypto *PrivyWalletDepositAccountCryptoService
}

func newPrivyWalletDepositAccountService(service WalletDepositAccountService, wallets *PrivyWalletService) *PrivyWalletDepositAccountService {
	return &PrivyWalletDepositAccountService{
		WalletDepositAccountService: service,
		Crypto:                      newPrivyWalletDepositAccountCryptoService(service.Crypto, wallets),
	}
}

// PrivyWalletDepositAccountCryptoService wraps the generated
// WalletDepositAccountCryptoService with automatic authorization signature
// generation and privy-request-expiry handling.
type PrivyWalletDepositAccountCryptoService struct {
	WalletDepositAccountCryptoService
	wallets *PrivyWalletService
}

func newPrivyWalletDepositAccountCryptoService(service WalletDepositAccountCryptoService, wallets *PrivyWalletService) *PrivyWalletDepositAccountCryptoService {
	return &PrivyWalletDepositAccountCryptoService{WalletDepositAccountCryptoService: service, wallets: wallets}
}

// New creates crypto deposit accounts that sweep into the wallet with automatic authorization signature generation.
//
// This method wraps the generated WalletDepositAccountCryptoService.New and handles:
//   - Building the authorization signature from an AuthorizationContext
//   - Setting the idempotency key and request expiry headers
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - walletID: The wallet ID
//   - params: The request parameters (callers can skip PrivyAuthorizationSignature, PrivyIdempotencyKey, and PrivyRequestExpiry)
//   - opts: use WithAuthorizationContext, WithIdempotencyKey, and WithRequestExpiry
func (s *PrivyWalletDepositAccountCryptoService) New(
	ctx context.Context,
	walletID string,
	params WalletDepositAccountCryptoNewParams,
	opts ...RequestOption,
) (*CreateCryptoDepositAccountResponse, error) {
	options := applyRequestOptions(opts)
	w := s.wallets

	prepared, err := prepareRequest(ctx, w.appID, w.jwtExchanger, prepareRequestInput{
		authorizationContext: options.AuthorizationContext,
		idempotencyKey:       options.IdempotencyKey,
		requestExpiry:        w.requestExpiryOrDefault(options),
		method:               "POST",
		url:                  w.baseURL + "/v1/wallets/" + walletID + "/deposit_accounts/crypto",
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

	return s.WalletDepositAccountCryptoService.New(ctx, walletID, params, options.RequestOptions...)
}
