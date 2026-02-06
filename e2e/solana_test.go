package e2e_test

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"os"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
)

func TestWallets_Solana(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()
	jwt := generateTestJWT(t)
	sk := os.Getenv("P256_PRIVATE_KEY")

	wallets := []struct {
		name      string
		id        string
		publicKey string // base64-encoded ed25519 public key
		authCtx   *authorization.AuthorizationContext
	}{
		{
			name:      "Ownerless",
			id:        os.Getenv("OWNERLESS_SOLANA_WALLET_ID"),
			publicKey: os.Getenv("OWNERLESS_SOLANA_WALLET_PUBLIC_KEY"),
			authCtx:   nil, // no authorization context for ownerless
		},
		{
			name:      "KeyOwned",
			id:        os.Getenv("P256_OWNED_SOLANA_WALLET_ID"),
			publicKey: os.Getenv("P256_OWNED_SOLANA_WALLET_PUBLIC_KEY"),
			authCtx:   &authorization.AuthorizationContext{PrivateKeys: []string{sk}},
		},
		{
			name:      "UserOwned",
			id:        os.Getenv("USER_OWNED_SOLANA_WALLET_ID"),
			publicKey: os.Getenv("USER_OWNED_SOLANA_WALLET_PUBLIC_KEY"),
			authCtx:   &authorization.AuthorizationContext{UserJwts: []string{jwt}},
		},
	}

	t.Run("SignMessage", func(t *testing.T) {
		messageBytes := []byte("Hello, world!")
		message := base64.StdEncoding.EncodeToString(messageBytes)

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Solana.SignMessage(ctx, wallet.id,
					message,
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign message: %v", err)
				}

				if data.Signature == "" {
					t.Error("expected signature to be defined")
				}
				if data.Encoding != "base64" {
					t.Errorf("expected encoding to be base64, got %s", data.Encoding)
				}

				// Verify the signature
				pubKey, err := base64.StdEncoding.DecodeString(wallet.publicKey)
				if err != nil {
					t.Fatalf("failed to decode public key: %v", err)
				}
				sig, err := base64.StdEncoding.DecodeString(data.Signature)
				if err != nil {
					t.Fatalf("failed to decode signature: %v", err)
				}
				if !ed25519.Verify(pubKey, messageBytes, sig) {
					t.Error("signature verification failed")
				}
			})
		}
	})

	t.Run("SignMessageBytes", func(t *testing.T) {
		messageBytes := []byte("Hello, world!")

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Solana.SignMessageBytes(ctx, wallet.id,
					messageBytes,
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign message bytes: %v", err)
				}

				if data.Signature == "" {
					t.Error("expected signature to be defined")
				}
				if data.Encoding != "base64" {
					t.Errorf("expected encoding to be base64, got %s", data.Encoding)
				}

				// Verify the signature
				pubKey, err := base64.StdEncoding.DecodeString(wallet.publicKey)
				if err != nil {
					t.Fatalf("failed to decode public key: %v", err)
				}
				sig, err := base64.StdEncoding.DecodeString(data.Signature)
				if err != nil {
					t.Fatalf("failed to decode signature: %v", err)
				}
				if !ed25519.Verify(pubKey, messageBytes, sig) {
					t.Error("signature verification failed")
				}
			})
		}
	})

	t.Run("SignTransaction", func(t *testing.T) {
		// A placeholder base64-encoded transaction for testing
		// This is a minimal valid Solana transaction structure
		transaction := "AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAQABA6Sih224FoBZ3LfpkolHQKFK6YSbidfv1FW9YACkTdazfHrf9hlOV6sJkws1eGUPR+0+wKPsm78llwnS4EhArhoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQICAAEMAgAAAGQAAAAAAAAAAA=="

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Solana.SignTransaction(ctx, wallet.id,
					transaction,
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign transaction: %v", err)
				}

				if data.SignedTransaction == "" {
					t.Error("expected signed_transaction to be defined")
				}
				if data.Encoding != "base64" {
					t.Errorf("expected encoding to be base64, got %s", data.Encoding)
				}
			})
		}
	})

	t.Run("SignTransactionBytes", func(t *testing.T) {
		// Same transaction as above, decoded from base64 to raw bytes
		// The service will base64-encode it for transmission
		transactionB64 := "AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAQABA6Sih224FoBZ3LfpkolHQKFK6YSbidfv1FW9YACkTdazfHrf9hlOV6sJkws1eGUPR+0+wKPsm78llwnS4EhArhoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQICAAEMAgAAAGQAAAAAAAAAAA=="

		transaction, err := base64.StdEncoding.DecodeString(transactionB64)
		if err != nil {
			t.Fatalf("failed to decode test transaction: %v", err)
		}

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Solana.SignTransactionBytes(ctx, wallet.id,
					transaction,
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign transaction bytes: %v", err)
				}

				if data.SignedTransaction == "" {
					t.Error("expected signed_transaction to be defined")
				}
				if data.Encoding != "base64" {
					t.Errorf("expected encoding to be base64, got %s", data.Encoding)
				}
			})
		}
	})
}
