package e2e_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"testing"

	. "github.com/privy-io/go-sdk"
)

func generateKeyPair(t *testing.T) (string, string) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate P-256 key: %v", err)
	}

	pkcs8, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("failed to marshal key to PKCS8: %v", err)
	}

	spki, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatalf("failed to marshal public key to SPKI: %v", err)
	}

	pk := base64.StdEncoding.EncodeToString(spki)
	sk := base64.StdEncoding.EncodeToString(pkcs8)
	return pk, sk
}

func TestKeyQuorums(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	pk1, _ := generateKeyPair(t)
	pk2, _ := generateKeyPair(t)

	t.Run("New", func(t *testing.T) {
		t.Skip("skipped to avoid creating test resources")
		t.Run("PublicKeys", func(t *testing.T) {
			quorum, err := client.KeyQuorums.New(ctx, KeyQuorumNewParams{
				DisplayName: String("test-key-quorum"),
				PublicKeys:  []string{pk1, pk2},
			})
			if err != nil {
				t.Fatalf("failed to create key quorum: %v", err)
			}

			if quorum.ID == "" {
				t.Error("expected key quorum ID to be defined")
			}
			if quorum.DisplayName != "test-key-quorum" {
				t.Errorf("expected display_name to be test-key-quorum, got %s", quorum.DisplayName)
			}
		})
	})
}
