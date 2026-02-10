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
	"github.com/privy-io/go-sdk/authorization"
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

func createKeyQuorum(t *testing.T, ctx context.Context, client *PrivyClient, params KeyQuorumNewParams, auth authorization.AuthorizationContext) *KeyQuorum {
	keyQuorum, err := client.KeyQuorums.New(ctx, params)
	if err != nil {
		t.Fatalf("failed to create key quorum: %v", err)
	}

	t.Logf("Created Key Quorum %s", keyQuorum.ID)

	if keyQuorum.ID == "" {
		t.Error("expected key quorum ID to be defined")
	}

	if keyQuorum.DisplayName != "test-key-quorum" {
		t.Errorf("expected display_name to be test-key-quorum, got %s", keyQuorum.DisplayName)
	}

	t.Cleanup(func() {
		t.Logf("Deleting Key Quorum %s", keyQuorum.ID)
		result, err := client.KeyQuorums.Delete(
			ctx,
			keyQuorum.ID,
			KeyQuorumDeleteParams{},
			WithAuthorizationContext(&auth),
		)
		if err != nil {
			t.Fatalf("failed to delete key quorum: %v", err)
		}

		if !result.Success {
			t.Error("expected delete to succeed")
		}
	})

	return keyQuorum
}

func TestKeyQuorums(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	pk1, sk1 := generateKeyPair(t)
	pk2, sk2 := generateKeyPair(t)

	kq1 := createKeyQuorum(t, ctx, client, KeyQuorumNewParams{
		DisplayName: String("test-key-quorum"),
		PublicKeys:  []string{pk1, pk2},
	}, authorization.AuthorizationContext{
		PrivateKeys: []string{sk1, sk2},
	})

	t.Run("Update", func(t *testing.T) {
		authCtx := &authorization.AuthorizationContext{
			PrivateKeys: []string{sk1, sk2},
		}

		// Call Update with authorization
		result, err := client.KeyQuorums.Update(
			ctx,
			kq1.ID,
			KeyQuorumUpdateParams{
				DisplayName: String("go-sdk-test-key-quorum-updated"),
			},
			WithAuthorizationContext(authCtx),
		)
		if err != nil {
			t.Fatalf("failed to update key quorum: %v", err)
		}

		if result.ID != kq1.ID {
			t.Errorf("expected key quorum ID %s, got %s", kq1.ID, result.ID)
		}
		if result.DisplayName != "go-sdk-test-key-quorum-updated" {
			t.Errorf("expected display name 'go-sdk-test-key-quorum-updated', got %s", result.DisplayName)
		}
	})
}
