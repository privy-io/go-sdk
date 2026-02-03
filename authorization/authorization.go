// Package authorization provides types and functions for signing API requests
// that require owner authorization. This is used when performing operations on
// resources across the API.
package authorization

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cyberphone/json-canonicalization/go/src/webpki.org/jsoncanonicalizer"
)

// AuthorizationContext contains credentials used for signing authorization requests.
type AuthorizationContext struct {
	// PrivateKeys is an array of base64-encoded PKCS8-formatted P-256 private keys.
	// Keys must not include PEM headers.
	PrivateKeys []string
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
//   - ctx: AuthorizationContext containing the credentials
//   - payload: The arbitrary byte array to sign
//
// Returns:
//   - An array of base64-encoded DER-format signatures, one per key
//   - An error if any key fails to sign, indicating which key index failed
func GenerateAuthorizationSignatures(ctx AuthorizationContext, payload []byte) ([]string, error) {
	if len(ctx.PrivateKeys) == 0 {
		return []string{}, nil
	}

	signatures := make([]string, len(ctx.PrivateKeys))
	for i, key := range ctx.PrivateKeys {
		sig, err := GenerateAuthorizationSignature(key, payload)
		if err != nil {
			return nil, fmt.Errorf("failed to sign with key at index %d: %w", i, err)
		}
		signatures[i] = sig
	}

	return signatures, nil
}
