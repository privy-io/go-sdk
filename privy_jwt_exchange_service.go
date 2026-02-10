package privyclient

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/privy-io/go-sdk/internal/hpke"
	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyJwtExchangeService provides JWT-to-authorization-key exchange functionality.
// It implements the internal/jwtexchange.JwtExchanger interface.
// The service is safe for concurrent use.
type PrivyJwtExchangeService struct {
	wallet *WalletService
	logger logger
}

// newPrivyJwtExchangeService creates a new JWT exchange service.
// This is unexported; the service is created automatically by NewPrivyClient.
func newPrivyJwtExchangeService(wallet *WalletService, logger logger) *PrivyJwtExchangeService {
	return &PrivyJwtExchangeService{
		wallet: wallet,
		logger: logger,
	}
}

// ExchangeJwtForAuthorizationKey exchanges a user JWT for a short-lived authorization private key.
// The returned string is a base64-encoded PKCS8-formatted P-256 private key.
func (s *PrivyJwtExchangeService) ExchangeJwtForAuthorizationKey(ctx context.Context, jwt string) (string, error) {
	recipient, err := hpke.NewHpkeRecipient()
	if err != nil {
		return "", fmt.Errorf("failed to create HPKE recipient: %w", err)
	}

	recipientPubKey := base64.StdEncoding.EncodeToString(recipient.PublicKeySpki)

	resp, err := s.wallet.AuthenticateWithJwt(ctx, WalletAuthenticateWithJwtParams{
		UserJwt:            jwt,
		RecipientPublicKey: param.NewOpt(recipientPubKey),
		EncryptionType:     WalletAuthenticateWithJwtParamsEncryptionTypeHpke,
	})
	if err != nil {
		return "", fmt.Errorf("failed to authenticate with JWT: %w", err)
	}

	encResp := resp.AsWithEncryption()
	encKey := encResp.EncryptedAuthorizationKey

	if encKey.EncapsulatedKey == "" || encKey.Ciphertext == "" {
		return "", fmt.Errorf("response missing required encryption fields")
	}

	encapKey, err := base64.StdEncoding.DecodeString(encKey.EncapsulatedKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode encapsulated key: %w", err)
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encKey.Ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	plaintext, err := recipient.Decrypt(encapKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt authorization key: %w", err)
	}

	return string(plaintext), nil
}
