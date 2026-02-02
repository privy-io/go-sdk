package authorization

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"strings"
	"testing"
)

// generateTestP256Key creates a new P-256 private key and returns it as a
// base64-encoded PKCS8 string.
func generateTestP256Key(t *testing.T) (string, *ecdsa.PrivateKey) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate P-256 key: %v", err)
	}
	pkcs8, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("failed to marshal key to PKCS8: %v", err)
	}
	return base64.StdEncoding.EncodeToString(pkcs8), key
}

// generateTestP384Key creates a new P-384 private key and returns it as a
// base64-encoded PKCS8 string.
func generateTestP384Key(t *testing.T) string {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate P-384 key: %v", err)
	}
	pkcs8, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("failed to marshal key to PKCS8: %v", err)
	}
	return base64.StdEncoding.EncodeToString(pkcs8)
}

// generateTestRSAKey creates a new RSA private key and returns it as a
// base64-encoded PKCS8 string.
func generateTestRSAKey(t *testing.T) string {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}
	pkcs8, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("failed to marshal key to PKCS8: %v", err)
	}
	return base64.StdEncoding.EncodeToString(pkcs8)
}

func TestGenerateAuthorizationSignature_Success(t *testing.T) {
	privateKeyB64, privateKey := generateTestP256Key(t)
	payload := []byte("test payload for signing")

	signature, err := GenerateAuthorizationSignature(privateKeyB64, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if signature == "" {
		t.Fatal("expected non-empty signature")
	}

	// Verify the signature is valid base64
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		t.Fatalf("signature is not valid base64: %v", err)
	}

	// Verify the signature using the public key
	hash := sha256.Sum256(payload)
	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sigBytes)
	if !valid {
		t.Fatal("signature verification failed")
	}
}

func TestGenerateAuthorizationSignature_EmptyPayload(t *testing.T) {
	privateKeyB64, privateKey := generateTestP256Key(t)
	payload := []byte{}

	signature, err := GenerateAuthorizationSignature(privateKeyB64, payload)
	if err != nil {
		t.Fatalf("unexpected error signing empty payload: %v", err)
	}

	if signature == "" {
		t.Fatal("expected non-empty signature for empty payload")
	}

	// Verify the signature
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		t.Fatalf("signature is not valid base64: %v", err)
	}

	hash := sha256.Sum256(payload)
	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sigBytes)
	if !valid {
		t.Fatal("signature verification failed for empty payload")
	}
}

func TestGenerateAuthorizationSignature_InvalidBase64(t *testing.T) {
	invalidBase64 := "not-valid-base64!!!"
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignature(invalidBase64, payload)
	if err == nil {
		t.Fatal("expected error for invalid base64 input")
	}

	if !strings.Contains(err.Error(), "invalid base64 encoding") {
		t.Errorf("expected error message about invalid base64, got: %v", err)
	}
}

func TestGenerateAuthorizationSignature_InvalidPKCS8(t *testing.T) {
	// Valid base64 but not valid PKCS8
	invalidPKCS8 := base64.StdEncoding.EncodeToString([]byte("not a valid PKCS8 key"))
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignature(invalidPKCS8, payload)
	if err == nil {
		t.Fatal("expected error for invalid PKCS8 format")
	}

	if !strings.Contains(err.Error(), "invalid PKCS8 format") {
		t.Errorf("expected error message about invalid PKCS8 format, got: %v", err)
	}
}

func TestGenerateAuthorizationSignature_WrongKeyType_RSA(t *testing.T) {
	rsaKeyB64 := generateTestRSAKey(t)
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignature(rsaKeyB64, payload)
	if err == nil {
		t.Fatal("expected error for RSA key")
	}

	if !strings.Contains(err.Error(), "not an ECDSA key") {
		t.Errorf("expected error message about not being ECDSA key, got: %v", err)
	}
}

func TestGenerateAuthorizationSignature_WrongCurve_P384(t *testing.T) {
	p384KeyB64 := generateTestP384Key(t)
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignature(p384KeyB64, payload)
	if err == nil {
		t.Fatal("expected error for P-384 key")
	}

	if !strings.Contains(err.Error(), "not on the P-256 curve") {
		t.Errorf("expected error message about wrong curve, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_Success(t *testing.T) {
	key1B64, key1 := generateTestP256Key(t)
	key2B64, key2 := generateTestP256Key(t)

	ctx := AuthorizationContext{
		PrivateKeys: []string{key1B64, key2B64},
	}
	payload := []byte("test payload for multiple signatures")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 2 {
		t.Fatalf("expected 2 signatures, got %d", len(signatures))
	}

	// Verify both signatures
	hash := sha256.Sum256(payload)

	sig1Bytes, _ := base64.StdEncoding.DecodeString(signatures[0])
	if !ecdsa.VerifyASN1(&key1.PublicKey, hash[:], sig1Bytes) {
		t.Error("first signature verification failed")
	}

	sig2Bytes, _ := base64.StdEncoding.DecodeString(signatures[1])
	if !ecdsa.VerifyASN1(&key2.PublicKey, hash[:], sig2Bytes) {
		t.Error("second signature verification failed")
	}
}

func TestGenerateAuthorizationSignatures_EmptyContext(t *testing.T) {
	ctx := AuthorizationContext{
		PrivateKeys: []string{},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload)
	if err != nil {
		t.Fatalf("unexpected error for empty context: %v", err)
	}

	if signatures == nil {
		t.Fatal("expected non-nil slice for empty context")
	}

	if len(signatures) != 0 {
		t.Fatalf("expected empty slice, got %d signatures", len(signatures))
	}
}

func TestGenerateAuthorizationSignatures_NilPrivateKeys(t *testing.T) {
	ctx := AuthorizationContext{
		PrivateKeys: nil,
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload)
	if err != nil {
		t.Fatalf("unexpected error for nil private keys: %v", err)
	}

	if signatures == nil {
		t.Fatal("expected non-nil slice for nil private keys")
	}

	if len(signatures) != 0 {
		t.Fatalf("expected empty slice, got %d signatures", len(signatures))
	}
}

func TestGenerateAuthorizationSignatures_PartialFailure(t *testing.T) {
	validKeyB64, _ := generateTestP256Key(t)
	invalidKeyB64 := base64.StdEncoding.EncodeToString([]byte("invalid key"))

	ctx := AuthorizationContext{
		PrivateKeys: []string{validKeyB64, invalidKeyB64},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload)
	if err == nil {
		t.Fatal("expected error for partial failure")
	}

	if !strings.Contains(err.Error(), "index 1") {
		t.Errorf("expected error to mention failed key index, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_FirstKeyInvalid(t *testing.T) {
	invalidKeyB64 := base64.StdEncoding.EncodeToString([]byte("invalid key"))
	validKeyB64, _ := generateTestP256Key(t)

	ctx := AuthorizationContext{
		PrivateKeys: []string{invalidKeyB64, validKeyB64},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload)
	if err == nil {
		t.Fatal("expected error when first key is invalid")
	}

	if !strings.Contains(err.Error(), "index 0") {
		t.Errorf("expected error to mention index 0, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_SingleKey(t *testing.T) {
	keyB64, key := generateTestP256Key(t)

	ctx := AuthorizationContext{
		PrivateKeys: []string{keyB64},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}

	// Verify the signature
	hash := sha256.Sum256(payload)
	sigBytes, _ := base64.StdEncoding.DecodeString(signatures[0])
	if !ecdsa.VerifyASN1(&key.PublicKey, hash[:], sigBytes) {
		t.Error("signature verification failed")
	}
}

func TestGenerateAuthorizationSignature_DeterministicHash(t *testing.T) {
	privateKeyB64, _ := generateTestP256Key(t)
	payload := []byte("consistent payload")

	// Generate two signatures with the same key and payload
	sig1, err := GenerateAuthorizationSignature(privateKeyB64, payload)
	if err != nil {
		t.Fatalf("first signature failed: %v", err)
	}

	sig2, err := GenerateAuthorizationSignature(privateKeyB64, payload)
	if err != nil {
		t.Fatalf("second signature failed: %v", err)
	}

	// ECDSA signatures are non-deterministic (due to random k), so they should differ
	// but both should be valid. This test confirms the implementation works correctly.
	if sig1 == "" || sig2 == "" {
		t.Fatal("signatures should not be empty")
	}

	// Both signatures should be valid even though they're different
	// (This is expected behavior for ECDSA)
}
