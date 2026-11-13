// Package jwtexchange provides internal types for JWT exchange handling.
package jwtexchange

import (
	"context"
	"fmt"
)

// JwtExchanger is an interface for exchanging user JWTs for authorization private keys.
type JwtExchanger interface {
	// ExchangeJwtForAuthorizationKeyWithMfaToken exchanges a user JWT for a
	// short-lived authorization private key, sending an MFA token from the
	// user's completed MFA verification with the exchange.
	//
	// Input:
	//   - ctx: Context for cancellation and timeouts
	//   - jwt: A valid user JWT
	//   - mfaToken: An MFA token for the user. Empty sends no MFA token.
	//
	// Output:
	//   - string: Base64-encoded PKCS8-formatted P-256 private key
	//   - error: Non-nil if the exchange fails
	ExchangeJwtForAuthorizationKeyWithMfaToken(ctx context.Context, jwt string, mfaToken string) (string, error)
}

// ExchangeJwtsForKeys exchanges a list of JWTs for authorization private keys.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - exchanger: JwtExchanger for exchanging JWTs for private keys
//   - jwts: List of JWTs to exchange
//   - mfaTokens: MFA token to send with each JWT's exchange, keyed by JWT. A JWT with no entry sends no MFA token.
//
// Returns:
//   - []string: List of base64-encoded PKCS8-formatted P-256 private keys
//   - error: Non-nil if any exchange fails, indicating which JWT index failed
func ExchangeJwtsForKeys(ctx context.Context, exchanger JwtExchanger, jwts []string, mfaTokens map[string]string) ([]string, error) {
	keys := make([]string, 0, len(jwts))
	for i, jwt := range jwts {
		key, err := exchanger.ExchangeJwtForAuthorizationKeyWithMfaToken(ctx, jwt, mfaTokens[jwt])
		if err != nil {
			return nil, fmt.Errorf("failed to exchange JWT at index %d: %w", i, err)
		}
		keys = append(keys, key)
	}
	return keys, nil
}
