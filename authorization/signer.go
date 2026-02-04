package authorization

import "context"

// AuthorizationSigner is an interface for external signing implementations.
// This enables the use of key management systems like AWS KMS
// without having to pull the key or share it with Privy directly.
type AuthorizationSigner interface {
	// Sign signs the payload and returns a base64-encoded DER-format ECDSA P-256 signature.
	// Implementations are responsible for hashing the payload with SHA-256 before signing.
	//
	// Input:
	//   - ctx: Context for cancellation and timeouts
	//   - payload: The raw byte array to sign (must be hashed with SHA-256 before signing)
	//
	// Output:
	//   - string: Base64-encoded DER-format signature
	//   - error: Non-nil if signing fails
	Sign(ctx context.Context, payload []byte) (string, error)
}
