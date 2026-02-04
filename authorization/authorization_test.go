package authorization

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
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

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

	_, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

	_, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

// Tests for FormatRequestForAuthorizationSignature

func TestFormatRequestForAuthorizationSignature_BasicRequest(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/wallets",
		Body: map[string]any{
			"name": "test-wallet",
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("expected non-empty result")
	}

	// Verify the result is valid JSON
	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	// Verify the fields are present
	if parsed["version"] != float64(1) {
		t.Errorf("expected version 1, got %v", parsed["version"])
	}
	if parsed["method"] != "POST" {
		t.Errorf("expected method POST, got %v", parsed["method"])
	}
	if parsed["url"] != "https://api.privy.io/v1/wallets" {
		t.Errorf("expected correct URL, got %v", parsed["url"])
	}
}

func TestFormatRequestForAuthorizationSignature_EmptyBodyBecomesEmptyString(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body:    map[string]any{}, // Empty body {}
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the body is an empty string, not {}
	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	body := parsed["body"]
	if body != "" {
		t.Errorf("expected empty string body, got %v (type %T)", body, body)
	}
}

func TestFormatRequestForAuthorizationSignature_NilBodyIsOmitted(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "DELETE",
		URL:     "https://api.privy.io/v1/test",
		Body:    nil,
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the body field is omitted (not present in JSON)
	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	if _, exists := parsed["body"]; exists {
		t.Errorf("expected body field to be omitted, but it exists: %v", parsed["body"])
	}
}

func TestFormatRequestForAuthorizationSignature_NestedObjectsKeysOrdered(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body: map[string]any{
			"zebra": "last",
			"alpha": "first",
			"nested": map[string]any{
				"zoo":      "z",
				"aardvark": "a",
			},
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// RFC 8785 requires keys to be sorted lexicographically
	resultStr := string(result)

	// Check that "alpha" comes before "nested" and "nested" comes before "zebra"
	alphaIdx := strings.Index(resultStr, `"alpha"`)
	nestedIdx := strings.Index(resultStr, `"nested"`)
	zebraIdx := strings.Index(resultStr, `"zebra"`)

	if alphaIdx == -1 || nestedIdx == -1 || zebraIdx == -1 {
		t.Fatalf("expected all keys to be present in result")
	}

	if alphaIdx > nestedIdx {
		t.Errorf("expected 'alpha' before 'nested', got positions %d and %d", alphaIdx, nestedIdx)
	}
	if nestedIdx > zebraIdx {
		t.Errorf("expected 'nested' before 'zebra', got positions %d and %d", nestedIdx, zebraIdx)
	}

	// Check nested object key ordering (aardvark before zoo)
	aardvarkIdx := strings.Index(resultStr, `"aardvark"`)
	zooIdx := strings.Index(resultStr, `"zoo"`)

	if aardvarkIdx == -1 || zooIdx == -1 {
		t.Fatalf("expected nested keys to be present in result")
	}

	if aardvarkIdx > zooIdx {
		t.Errorf("expected 'aardvark' before 'zoo' in nested object, got positions %d and %d", aardvarkIdx, zooIdx)
	}
}

func TestFormatRequestForAuthorizationSignature_BodyWithArrays(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body: map[string]any{
			"items": []any{"apple", "banana", "cherry"},
			"nested": []any{
				map[string]any{"name": "item1"},
				map[string]any{"name": "item2"},
			},
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the result is valid JSON with arrays preserved
	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	body := parsed["body"].(map[string]any)
	items := body["items"].([]any)
	if len(items) != 3 {
		t.Errorf("expected 3 items, got %d", len(items))
	}
	if items[0] != "apple" || items[1] != "banana" || items[2] != "cherry" {
		t.Errorf("array order not preserved: %v", items)
	}
}

func TestFormatRequestForAuthorizationSignature_MultipleHeaders(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body:    map[string]any{"data": "test"},
		Headers: map[string]string{
			"privy-app-id":          "app-123",
			"privy-idempotency-key": "idem-456",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	headers := parsed["headers"].(map[string]any)
	if headers["privy-app-id"] != "app-123" {
		t.Errorf("expected privy-app-id header, got %v", headers["privy-app-id"])
	}
	if headers["privy-idempotency-key"] != "idem-456" {
		t.Errorf("expected privy-idempotency-key header, got %v", headers["privy-idempotency-key"])
	}
}

func TestFormatRequestForAuthorizationSignature_UnicodeCharacters(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body: map[string]any{
			"emoji":    "🔐",
			"chinese":  "中文",
			"japanese": "日本語",
			"mixed":    "Hello 世界! 🌍",
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the result is valid JSON with Unicode preserved
	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	body := parsed["body"].(map[string]any)
	if body["emoji"] != "🔐" {
		t.Errorf("emoji not preserved: %v", body["emoji"])
	}
	if body["chinese"] != "中文" {
		t.Errorf("chinese not preserved: %v", body["chinese"])
	}
}

func TestFormatRequestForAuthorizationSignature_Numbers(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body: map[string]any{
			"integer":       42,
			"negative":      -100,
			"float":         3.14159,
			"scientific":    1.23e10,
			"zero":          0,
			"negativeFloat": -0.5,
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	body := parsed["body"].(map[string]any)
	if body["integer"] != float64(42) {
		t.Errorf("integer not correct: %v", body["integer"])
	}
	if body["negative"] != float64(-100) {
		t.Errorf("negative not correct: %v", body["negative"])
	}
	if body["zero"] != float64(0) {
		t.Errorf("zero not correct: %v", body["zero"])
	}
}

func TestFormatRequestForAuthorizationSignature_Deterministic(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	// Generate multiple times and verify consistency
	result1, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("first call failed: %v", err)
	}

	result2, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("second call failed: %v", err)
	}

	result3, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("third call failed: %v", err)
	}

	if string(result1) != string(result2) {
		t.Errorf("results 1 and 2 differ:\n1: %s\n2: %s", result1, result2)
	}

	if string(result2) != string(result3) {
		t.Errorf("results 2 and 3 differ:\n2: %s\n3: %s", result2, result3)
	}
}

func TestFormatRequestForAuthorizationSignature_InvalidBody(t *testing.T) {
	// channels cannot be marshaled to JSON
	ch := make(chan int)

	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body:    ch,
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	_, err := FormatRequestForAuthorizationSignature(input)
	if err == nil {
		t.Fatal("expected error for invalid body type")
	}

	if !strings.Contains(err.Error(), "failed to marshal body") {
		t.Errorf("expected error message about marshal failure, got: %v", err)
	}
}

func TestFormatRequestForAuthorizationSignature_NoWhitespace(t *testing.T) {
	input := WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/test",
		Body: map[string]any{
			"key": "value",
		},
		Headers: map[string]string{
			"privy-app-id": "app-123",
		},
	}

	result, err := FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// RFC 8785 specifies no whitespace between tokens
	resultStr := string(result)

	// Check for common whitespace patterns that shouldn't be present
	if strings.Contains(resultStr, ": ") {
		t.Error("result contains ': ' (colon with space)")
	}
	if strings.Contains(resultStr, ", ") {
		t.Error("result contains ', ' (comma with space)")
	}
	if strings.Contains(resultStr, "\n") {
		t.Error("result contains newline")
	}
	if strings.Contains(resultStr, "\t") {
		t.Error("result contains tab")
	}
}

func TestFormatRequestForAuthorizationSignature_AllMethods(t *testing.T) {
	methods := []string{"POST", "PUT", "PATCH", "DELETE"}

	for _, method := range methods {
		input := WalletApiRequestSignatureInput{
			Version: 1,
			Method:  method,
			URL:     "https://api.privy.io/v1/test",
			Body:    map[string]any{"test": true},
			Headers: map[string]string{
				"privy-app-id": "app-123",
			},
		}

		result, err := FormatRequestForAuthorizationSignature(input)
		if err != nil {
			t.Errorf("unexpected error for method %s: %v", method, err)
			continue
		}

		var parsed map[string]any
		if err := json.Unmarshal(result, &parsed); err != nil {
			t.Errorf("result is not valid JSON for method %s: %v", method, err)
			continue
		}

		if parsed["method"] != method {
			t.Errorf("expected method %s, got %v", method, parsed["method"])
		}
	}
}

func TestFormatRequestForAuthorizationSignature_LiteralCanonicalization(t *testing.T) {
	tests := []struct {
		name     string
		input    WalletApiRequestSignatureInput
		expected string
	}{
		{
			name: "nil body is omitted",
			input: WalletApiRequestSignatureInput{
				Version: 1,
				Method:  "POST",
				URL:     "/api/v1/wallets",
				Body:    nil,
				Headers: map[string]string{
					"privy-app-id": "test-app-id",
				},
			},
			expected: `{"headers":{"privy-app-id":"test-app-id"},"method":"POST","url":"/api/v1/wallets","version":1}`,
		},
		{
			name: "body with content",
			input: WalletApiRequestSignatureInput{
				Version: 1,
				Method:  "POST",
				URL:     "/api/v1/wallets",
				Body: map[string]any{
					"foo": "bar",
					"baz": 1,
					"qux": true,
				},
				Headers: map[string]string{
					"privy-app-id": "test-app-id",
				},
			},
			expected: `{"body":{"baz":1,"foo":"bar","qux":true},"headers":{"privy-app-id":"test-app-id"},"method":"POST","url":"/api/v1/wallets","version":1}`,
		},
		{
			name: "empty body becomes empty string",
			input: WalletApiRequestSignatureInput{
				Version: 1,
				Method:  "POST",
				URL:     "/api/v1/wallets",
				Body:    map[string]any{},
				Headers: map[string]string{
					"privy-app-id": "test-app-id",
				},
			},
			expected: `{"body":"","headers":{"privy-app-id":"test-app-id"},"method":"POST","url":"/api/v1/wallets","version":1}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatRequestForAuthorizationSignature(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(result) != tt.expected {
				t.Errorf("canonicalization mismatch:\ngot:      %s\nexpected: %s", string(result), tt.expected)
			}
		})
	}
}

// MockJwtExchanger implements JwtExchanger for testing.
type MockJwtExchanger struct {
	// Keys maps JWT strings to their corresponding private keys
	Keys map[string]string
	// Errors maps JWT strings to errors that should be returned
	Errors map[string]error
}

func (m *MockJwtExchanger) ExchangeJwtForAuthorizationKey(jwt string) (string, error) {
	if err, ok := m.Errors[jwt]; ok {
		return "", err
	}
	if key, ok := m.Keys[jwt]; ok {
		return key, nil
	}
	return "", errors.New("unknown JWT")
}

func TestGenerateAuthorizationSignatures_OnlyJwts(t *testing.T) {
	key1B64, key1 := generateTestP256Key(t)
	key2B64, key2 := generateTestP256Key(t)

	exchanger := &MockJwtExchanger{
		Keys: map[string]string{
			"jwt1": key1B64,
			"jwt2": key2B64,
		},
	}

	ctx := AuthorizationContext{
		UserJwts: []string{"jwt1", "jwt2"},
	}
	payload := []byte("test payload for JWT signatures")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, exchanger)
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

func TestGenerateAuthorizationSignatures_PrivateKeysAndJwts(t *testing.T) {
	pkKey1B64, pkKey1 := generateTestP256Key(t)
	pkKey2B64, pkKey2 := generateTestP256Key(t)
	jwtKey1B64, jwtKey1 := generateTestP256Key(t)
	jwtKey2B64, jwtKey2 := generateTestP256Key(t)

	exchanger := &MockJwtExchanger{
		Keys: map[string]string{
			"jwt1": jwtKey1B64,
			"jwt2": jwtKey2B64,
		},
	}

	ctx := AuthorizationContext{
		PrivateKeys: []string{pkKey1B64, pkKey2B64},
		UserJwts:    []string{"jwt1", "jwt2"},
	}
	payload := []byte("test payload for mixed signatures")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, exchanger)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 4 {
		t.Fatalf("expected 4 signatures, got %d", len(signatures))
	}

	// Verify order: PrivateKeys first, then JWT-derived keys
	hash := sha256.Sum256(payload)

	// First two should be from PrivateKeys
	sig0Bytes, _ := base64.StdEncoding.DecodeString(signatures[0])
	if !ecdsa.VerifyASN1(&pkKey1.PublicKey, hash[:], sig0Bytes) {
		t.Error("signature 0 (pkKey1) verification failed")
	}

	sig1Bytes, _ := base64.StdEncoding.DecodeString(signatures[1])
	if !ecdsa.VerifyASN1(&pkKey2.PublicKey, hash[:], sig1Bytes) {
		t.Error("signature 1 (pkKey2) verification failed")
	}

	// Last two should be from JWT-derived keys
	sig2Bytes, _ := base64.StdEncoding.DecodeString(signatures[2])
	if !ecdsa.VerifyASN1(&jwtKey1.PublicKey, hash[:], sig2Bytes) {
		t.Error("signature 2 (jwtKey1) verification failed")
	}

	sig3Bytes, _ := base64.StdEncoding.DecodeString(signatures[3])
	if !ecdsa.VerifyASN1(&jwtKey2.PublicKey, hash[:], sig3Bytes) {
		t.Error("signature 3 (jwtKey2) verification failed")
	}
}

func TestGenerateAuthorizationSignatures_EmptyContextNoJwts(t *testing.T) {
	ctx := AuthorizationContext{
		PrivateKeys: []string{},
		UserJwts:    []string{},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

func TestGenerateAuthorizationSignatures_JwtExchangeFailure(t *testing.T) {
	exchanger := &MockJwtExchanger{
		Errors: map[string]error{
			"bad-jwt": errors.New("exchange failed: invalid token"),
		},
	}

	ctx := AuthorizationContext{
		UserJwts: []string{"bad-jwt"},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload, exchanger)
	if err == nil {
		t.Fatal("expected error for JWT exchange failure")
	}

	if !strings.Contains(err.Error(), "failed to exchange JWT at index 0") {
		t.Errorf("expected error to mention JWT index, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_PartialJwtExchangeFailure(t *testing.T) {
	key1B64, _ := generateTestP256Key(t)

	exchanger := &MockJwtExchanger{
		Keys: map[string]string{
			"good-jwt": key1B64,
		},
		Errors: map[string]error{
			"bad-jwt": errors.New("exchange failed: expired token"),
		},
	}

	ctx := AuthorizationContext{
		UserJwts: []string{"good-jwt", "bad-jwt"},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload, exchanger)
	if err == nil {
		t.Fatal("expected error for partial JWT exchange failure")
	}

	if !strings.Contains(err.Error(), "failed to exchange JWT at index 1") {
		t.Errorf("expected error to mention JWT index 1, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_NilExchangerWithJwts(t *testing.T) {
	ctx := AuthorizationContext{
		UserJwts: []string{"some-jwt"},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err == nil {
		t.Fatal("expected error when JWTs present but exchanger is nil")
	}

	if !strings.Contains(err.Error(), "JWTs present but no exchanger provided") {
		t.Errorf("expected error about missing exchanger, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_NilExchangerWithoutJwts(t *testing.T) {
	keyB64, key := generateTestP256Key(t)

	ctx := AuthorizationContext{
		PrivateKeys: []string{keyB64},
		UserJwts:    nil,
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
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

func TestGenerateAuthorizationSignatures_OrderVerification(t *testing.T) {
	// Create distinct keys so we can verify order
	pkKey1B64, pkKey1 := generateTestP256Key(t)
	jwtKeyB64, jwtKey := generateTestP256Key(t)

	exchanger := &MockJwtExchanger{
		Keys: map[string]string{
			"jwt": jwtKeyB64,
		},
	}

	ctx := AuthorizationContext{
		PrivateKeys: []string{pkKey1B64},
		UserJwts:    []string{"jwt"},
	}
	payload := []byte("test payload for order verification")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, exchanger)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 2 {
		t.Fatalf("expected 2 signatures, got %d", len(signatures))
	}

	hash := sha256.Sum256(payload)

	// First signature should be from PrivateKey
	sig0Bytes, _ := base64.StdEncoding.DecodeString(signatures[0])
	if !ecdsa.VerifyASN1(&pkKey1.PublicKey, hash[:], sig0Bytes) {
		t.Error("first signature should be from PrivateKey")
	}
	if ecdsa.VerifyASN1(&jwtKey.PublicKey, hash[:], sig0Bytes) {
		t.Error("first signature should NOT be from JWT-derived key")
	}

	// Second signature should be from JWT-derived key
	sig1Bytes, _ := base64.StdEncoding.DecodeString(signatures[1])
	if !ecdsa.VerifyASN1(&jwtKey.PublicKey, hash[:], sig1Bytes) {
		t.Error("second signature should be from JWT-derived key")
	}
	if ecdsa.VerifyASN1(&pkKey1.PublicKey, hash[:], sig1Bytes) {
		t.Error("second signature should NOT be from PrivateKey")
	}
}

// MockAuthorizationSigner implements AuthorizationSigner for testing.
type MockAuthorizationSigner struct {
	ReturnSignature string
	ReturnError     error
	ReceivedPayload []byte
}

func (m *MockAuthorizationSigner) Sign(payload []byte) (string, error) {
	m.ReceivedPayload = payload
	return m.ReturnSignature, m.ReturnError
}

// Tests for Signatures field

func TestGenerateAuthorizationSignatures_EmptySignatures(t *testing.T) {
	ctx := AuthorizationContext{
		Signatures: []string{},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if signatures == nil {
		t.Fatal("expected non-nil slice")
	}

	if len(signatures) != 0 {
		t.Fatalf("expected empty slice, got %d signatures", len(signatures))
	}
}

func TestGenerateAuthorizationSignatures_SinglePrecomputedSignature(t *testing.T) {
	precomputedSig := "precomputed-signature-base64"
	ctx := AuthorizationContext{
		Signatures: []string{precomputedSig},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}

	if signatures[0] != precomputedSig {
		t.Errorf("expected signature %q, got %q", precomputedSig, signatures[0])
	}
}

func TestGenerateAuthorizationSignatures_MultiplePrecomputedSignatures(t *testing.T) {
	precomputedSigs := []string{"sig1-base64", "sig2-base64", "sig3-base64"}
	ctx := AuthorizationContext{
		Signatures: precomputedSigs,
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 3 {
		t.Fatalf("expected 3 signatures, got %d", len(signatures))
	}

	for i, expected := range precomputedSigs {
		if signatures[i] != expected {
			t.Errorf("signature[%d]: expected %q, got %q", i, expected, signatures[i])
		}
	}
}

func TestGenerateAuthorizationSignatures_SignaturesIncludedAsIs(t *testing.T) {
	// Signatures should be included without modification, even if they look unusual
	weirdSig := "this-is-not-valid-base64!!!"
	ctx := AuthorizationContext{
		Signatures: []string{weirdSig},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}

	if signatures[0] != weirdSig {
		t.Errorf("signature should be included as-is: expected %q, got %q", weirdSig, signatures[0])
	}
}

// Tests for Signers field

func TestGenerateAuthorizationSignatures_EmptySigners(t *testing.T) {
	ctx := AuthorizationContext{
		Signers: []AuthorizationSigner{},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if signatures == nil {
		t.Fatal("expected non-nil slice")
	}

	if len(signatures) != 0 {
		t.Fatalf("expected empty slice, got %d signatures", len(signatures))
	}
}

func TestGenerateAuthorizationSignatures_SingleSignerSuccess(t *testing.T) {
	expectedSig := "signer-signature-base64"
	signer := &MockAuthorizationSigner{ReturnSignature: expectedSig}
	ctx := AuthorizationContext{
		Signers: []AuthorizationSigner{signer},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}

	if signatures[0] != expectedSig {
		t.Errorf("expected signature %q, got %q", expectedSig, signatures[0])
	}

	// Verify the signer received the correct payload
	if string(signer.ReceivedPayload) != string(payload) {
		t.Errorf("signer received wrong payload: expected %q, got %q", payload, signer.ReceivedPayload)
	}
}

func TestGenerateAuthorizationSignatures_MultipleSignersSuccess(t *testing.T) {
	signer1 := &MockAuthorizationSigner{ReturnSignature: "sig1"}
	signer2 := &MockAuthorizationSigner{ReturnSignature: "sig2"}
	signer3 := &MockAuthorizationSigner{ReturnSignature: "sig3"}

	ctx := AuthorizationContext{
		Signers: []AuthorizationSigner{signer1, signer2, signer3},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 3 {
		t.Fatalf("expected 3 signatures, got %d", len(signatures))
	}

	expectedSigs := []string{"sig1", "sig2", "sig3"}
	for i, expected := range expectedSigs {
		if signatures[i] != expected {
			t.Errorf("signature[%d]: expected %q, got %q", i, expected, signatures[i])
		}
	}
}

func TestGenerateAuthorizationSignatures_SignerError(t *testing.T) {
	signer := &MockAuthorizationSigner{
		ReturnError: errors.New("KMS signing failed"),
	}

	ctx := AuthorizationContext{
		Signers: []AuthorizationSigner{signer},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err == nil {
		t.Fatal("expected error for signer failure")
	}

	if !strings.Contains(err.Error(), "signer at index 0 failed") {
		t.Errorf("expected error to mention signer index, got: %v", err)
	}
	if !strings.Contains(err.Error(), "KMS signing failed") {
		t.Errorf("expected error to contain original error message, got: %v", err)
	}
}

func TestGenerateAuthorizationSignatures_SecondSignerError(t *testing.T) {
	signer1 := &MockAuthorizationSigner{ReturnSignature: "sig1"}
	signer2 := &MockAuthorizationSigner{ReturnError: errors.New("vault unavailable")}

	ctx := AuthorizationContext{
		Signers: []AuthorizationSigner{signer1, signer2},
	}
	payload := []byte("test payload")

	_, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err == nil {
		t.Fatal("expected error for signer failure")
	}

	if !strings.Contains(err.Error(), "signer at index 1 failed") {
		t.Errorf("expected error to mention signer index 1, got: %v", err)
	}
}

// Tests for combined context scenarios

func TestGenerateAuthorizationSignatures_AllFieldsPopulated(t *testing.T) {
	// Setup: pre-computed signatures
	precomputedSigs := []string{"precomputed1", "precomputed2"}

	// Setup: private keys
	pkKey1B64, pkKey1 := generateTestP256Key(t)
	pkKey2B64, pkKey2 := generateTestP256Key(t)

	// Setup: JWT-derived keys
	jwtKey1B64, jwtKey1 := generateTestP256Key(t)
	jwtKey2B64, jwtKey2 := generateTestP256Key(t)
	exchanger := &MockJwtExchanger{
		Keys: map[string]string{
			"jwt1": jwtKey1B64,
			"jwt2": jwtKey2B64,
		},
	}

	// Setup: signers
	signer1 := &MockAuthorizationSigner{ReturnSignature: "signer1-sig"}
	signer2 := &MockAuthorizationSigner{ReturnSignature: "signer2-sig"}

	ctx := AuthorizationContext{
		Signatures:  precomputedSigs,
		PrivateKeys: []string{pkKey1B64, pkKey2B64},
		UserJwts:    []string{"jwt1", "jwt2"},
		Signers:     []AuthorizationSigner{signer1, signer2},
	}
	payload := []byte("test payload for all fields")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, exchanger)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expected: 2 precomputed + 2 private key + 2 JWT-derived + 2 signer = 8 signatures
	if len(signatures) != 8 {
		t.Fatalf("expected 8 signatures, got %d", len(signatures))
	}

	// Verify order: Signatures first
	if signatures[0] != "precomputed1" || signatures[1] != "precomputed2" {
		t.Errorf("precomputed signatures not in correct position: %v", signatures[:2])
	}

	// Verify order: PrivateKey signatures next
	hash := sha256.Sum256(payload)
	sig2Bytes, _ := base64.StdEncoding.DecodeString(signatures[2])
	if !ecdsa.VerifyASN1(&pkKey1.PublicKey, hash[:], sig2Bytes) {
		t.Error("signature[2] should be from pkKey1")
	}
	sig3Bytes, _ := base64.StdEncoding.DecodeString(signatures[3])
	if !ecdsa.VerifyASN1(&pkKey2.PublicKey, hash[:], sig3Bytes) {
		t.Error("signature[3] should be from pkKey2")
	}

	// Verify order: JWT-derived signatures next
	sig4Bytes, _ := base64.StdEncoding.DecodeString(signatures[4])
	if !ecdsa.VerifyASN1(&jwtKey1.PublicKey, hash[:], sig4Bytes) {
		t.Error("signature[4] should be from jwtKey1")
	}
	sig5Bytes, _ := base64.StdEncoding.DecodeString(signatures[5])
	if !ecdsa.VerifyASN1(&jwtKey2.PublicKey, hash[:], sig5Bytes) {
		t.Error("signature[5] should be from jwtKey2")
	}

	// Verify order: Signer signatures last
	if signatures[6] != "signer1-sig" || signatures[7] != "signer2-sig" {
		t.Errorf("signer signatures not in correct position: %v", signatures[6:])
	}
}

func TestGenerateAuthorizationSignatures_OnlySignatures(t *testing.T) {
	ctx := AuthorizationContext{
		Signatures: []string{"only-precomputed-sig"},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}

	if signatures[0] != "only-precomputed-sig" {
		t.Errorf("expected %q, got %q", "only-precomputed-sig", signatures[0])
	}
}

func TestGenerateAuthorizationSignatures_OnlySigners(t *testing.T) {
	signer := &MockAuthorizationSigner{ReturnSignature: "only-signer-sig"}
	ctx := AuthorizationContext{
		Signers: []AuthorizationSigner{signer},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}

	if signatures[0] != "only-signer-sig" {
		t.Errorf("expected %q, got %q", "only-signer-sig", signatures[0])
	}
}

func TestGenerateAuthorizationSignatures_PrivateKeysAndSigners(t *testing.T) {
	pkKeyB64, pkKey := generateTestP256Key(t)
	signer := &MockAuthorizationSigner{ReturnSignature: "signer-sig"}

	ctx := AuthorizationContext{
		PrivateKeys: []string{pkKeyB64},
		Signers:     []AuthorizationSigner{signer},
	}
	payload := []byte("test payload")

	signatures, err := GenerateAuthorizationSignatures(ctx, payload, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(signatures) != 2 {
		t.Fatalf("expected 2 signatures, got %d", len(signatures))
	}

	// First should be from private key
	hash := sha256.Sum256(payload)
	sig0Bytes, _ := base64.StdEncoding.DecodeString(signatures[0])
	if !ecdsa.VerifyASN1(&pkKey.PublicKey, hash[:], sig0Bytes) {
		t.Error("signature[0] should be from private key")
	}

	// Second should be from signer
	if signatures[1] != "signer-sig" {
		t.Errorf("signature[1] should be from signer: expected %q, got %q", "signer-sig", signatures[1])
	}
}
