//go:build e2e

package e2e_test

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generateTestJWT creates a JWT token for testing purposes.
// It reads the signing key from JWT_AUTH_SK and subject from JWT_AUTH_SUBJECT environment variables.
func generateTestJWT(t *testing.T) string {
	t.Helper()

	secretKey := os.Getenv("JWT_AUTH_SK")
	if secretKey == "" {
		t.Fatal("JWT_AUTH_SK environment variable is required")
	}

	subject := os.Getenv("JWT_AUTH_SUBJECT")
	if subject == "" {
		t.Fatal("JWT_AUTH_SUBJECT environment variable is required")
	}

	block, _ := pem.Decode([]byte(secretKey))
	if block == nil {
		t.Fatal("failed to decode PEM block from JWT_AUTH_SK")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("failed to parse PKCS8 private key: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	token.Header["typ"] = "JWT"

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		t.Fatalf("failed to sign JWT: %v", err)
	}

	return signedToken
}
