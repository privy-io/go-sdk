// Package authorization provides types and functions for signing API requests
// that require owner authorization. This is used when performing operations on
// resources across the API.
package authorization

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cyberphone/json-canonicalization/go/src/webpki.org/jsoncanonicalizer"
	"github.com/privy-io/go-sdk/internal/jwtexchange"
)

// AuthorizationContext contains credentials used for signing authorization requests.
type AuthorizationContext struct {
	// PrivateKeys is an array of base64-encoded PKCS8-formatted P-256 private keys.
	// Keys must not include PEM headers.
	PrivateKeys []string

	// UserJwts contains JWTs for users that should sign the request authorization.
	// These should be valid JWTs for the user. Each JWT will be exchanged for
	// a short-lived authorization private key via the JwtExchanger.
	UserJwts []string

	// Signatures contains pre-computed base64-encoded DER-format signatures.
	// These are included directly in the authorization header without modification.
	Signatures []string

	// Signers contains external signing implementations.
	// Each signer's Sign method is called with the formatted payload.
	Signers []AuthorizationSigner
}

// WalletApiRequestSignatureInput defines the structure of a request payload
// that gets signed for authorization.
type WalletApiRequestSignatureInput struct {
	// Version is the signature version. Currently, 1 is the only valid version.
	Version int `json:"version"`
	// Method is the HTTP request method: "POST", "PUT", "PATCH", or "DELETE".
	// Signatures are not required on GET requests.
	Method string `json:"method"`
	// URL is the request URL. Should not contain a trailing slash.
	URL string `json:"url"`
	// Body is the request body (JSON-serializable). Omitted when nil.
	Body any `json:"body,omitempty"`
	// Headers contains Privy-specific headers to include in signature.
	// Required: "privy-app-id". Optional: "privy-idempotency-key".
	Headers map[string]string `json:"headers"`
}

// FormatRequestForAuthorizationSignature formats the request payload into bytes
// ready for signing. It creates a canonical representation of the request using
// RFC 8785 JSON Canonicalization Scheme (JCS).
//
// Parameters:
//   - input: WalletApiRequestSignatureInput - The request payload to format
//
// Returns:
//   - []byte: UTF-8 encoded canonicalized JSON
//   - error: Non-nil if JSON marshaling or canonicalization fails
func FormatRequestForAuthorizationSignature(input WalletApiRequestSignatureInput) ([]byte, error) {
	// Handle special case: empty body {} should become ""
	body := input.Body
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		if string(bodyBytes) == "{}" {
			body = ""
		}
	}

	// Create a copy with the potentially modified body
	inputCopy := WalletApiRequestSignatureInput{
		Version: input.Version,
		Method:  input.Method,
		URL:     input.URL,
		Body:    body,
		Headers: input.Headers,
	}

	// Marshal the input to JSON
	jsonBytes, err := json.Marshal(inputCopy)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input to JSON: %w", err)
	}

	// Canonicalize the JSON using RFC 8785
	canonicalized, err := jsoncanonicalizer.Transform(jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to canonicalize JSON: %w", err)
	}

	return canonicalized, nil
}

// GenerateAuthorizationSignature signs a payload with a P-256 private key.
//
// Parameters:
//   - privateKey: Base64-encoded PKCS8-formatted private key, without PEM headers
//   - payload: The arbitrary byte array to sign
//
// Returns:
//   - A base64-encoded DER-format signature
//   - An error if decoding, parsing, or signing fails
func GenerateAuthorizationSignature(privateKey string, payload []byte) (string, error) {
	// Decode the base64-encoded private key
	keyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("invalid base64 encoding in private key: %w", err)
	}

	// Parse the PKCS8-formatted private key
	parsedKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("invalid PKCS8 format: %w", err)
	}

	// Verify the key is an ECDSA key on the P-256 curve
	ecdsaKey, ok := parsedKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", errors.New("key is not an ECDSA key")
	}

	if ecdsaKey.Curve.Params().Name != "P-256" {
		return "", fmt.Errorf("key is not on the P-256 curve (got %s)", ecdsaKey.Curve.Params().Name)
	}

	// Sign the SHA-256 hash of the payload
	hash := sha256.Sum256(payload)

	signature, err := ecdsa.SignASN1(rand.Reader, ecdsaKey, hash[:])
	if err != nil {
		return "", fmt.Errorf("signing operation failed: %w", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// GenerateAuthorizationSignatures generates signatures for all credentials
// in an AuthorizationContext.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - auth: AuthorizationContext containing the credentials
//   - payload: The arbitrary byte array to sign
//   - exchanger: JwtExchanger for exchanging JWTs for private keys (may be nil if no JWTs in context)
//
// Returns:
//   - An array of base64-encoded DER-format signatures
//   - An error if any signing operation fails
func GenerateAuthorizationSignatures(ctx context.Context, auth AuthorizationContext, payload []byte, exchanger jwtexchange.JwtExchanger) ([]string, error) {
	// Check if JWTs are present but no exchanger provided
	if len(auth.UserJwts) > 0 && exchanger == nil {
		return nil, errors.New("JWTs present but no exchanger provided")
	}

	// Initialize slice for collecting signatures
	signatures := make([]string, 0, len(auth.Signatures)+len(auth.PrivateKeys)+len(auth.UserJwts)+len(auth.Signers))

	// Append pre-computed signatures directly (no validation required)
	signatures = append(signatures, auth.Signatures...)

	// Exchange JWTs for private keys
	jwtDerivedKeys, err := jwtexchange.ExchangeJwtsForKeys(ctx, exchanger, auth.UserJwts)
	if err != nil {
		return nil, err
	}

	// Combine all private keys and sign with them
	allKeys := append(auth.PrivateKeys, jwtDerivedKeys...)

	for i, key := range allKeys {
		sig, err := GenerateAuthorizationSignature(key, payload)
		if err != nil {
			return nil, fmt.Errorf("failed to sign with key at index %d: %w", i, err)
		}
		signatures = append(signatures, sig)
	}

	// Call each external signer
	for i, signer := range auth.Signers {
		sig, err := signer.Sign(ctx, payload)
		if err != nil {
			return nil, fmt.Errorf("signer at index %d failed: %w", i, err)
		}
		signatures = append(signatures, sig)
	}

	return signatures, nil
}

// GenerateAuthorizationSignaturesForRequest formats a request and generates
// signatures for all credentials in an AuthorizationContext.
//
// This is a convenience function that combines FormatRequestForAuthorizationSignature
// and GenerateAuthorizationSignatures.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - auth: AuthorizationContext containing the credentials
//   - input: WalletApiRequestSignatureInput describing the request to sign
//   - exchanger: JwtExchanger for exchanging JWTs for private keys (may be nil if no JWTs in context)
//
// Returns:
//   - An array of base64-encoded DER-format signatures
//   - An error if formatting or any signing operation fails
func GenerateAuthorizationSignaturesForRequest(ctx context.Context, auth AuthorizationContext, input WalletApiRequestSignatureInput, exchanger jwtexchange.JwtExchanger) ([]string, error) {
	payload, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		return nil, fmt.Errorf("failed to format request: %w", err)
	}

	return GenerateAuthorizationSignatures(ctx, auth, payload, exchanger)
}
