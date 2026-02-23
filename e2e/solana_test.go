package e2e_test

import (
	"crypto/ed25519"
	"encoding/base64"
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	. "github.com/privy-io/go-sdk"
)

// verifySolanaSignature verifies an ed25519 signature against a Solana wallet's
// base58-encoded public key address.
func verifySolanaSignature(t *testing.T, address string, message, signature []byte) {
	t.Helper()
	pubKey := base58.Decode(address)
	if len(pubKey) != ed25519.PublicKeySize {
		t.Fatalf("unexpected public key length %d (expected %d)", len(pubKey), ed25519.PublicKeySize)
	}
	if !ed25519.Verify(pubKey, message, signature) {
		t.Error("signature verification failed")
	}
}

func TestWallets_Solana(t *testing.T) {
	client := newTestClient(t)
	res := setupTestWalletResources(t, client)
	ctx := res.ctx
	wallets := res.createTestWallets(t, WalletChainTypeSolana)

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

				sig, err := base64.StdEncoding.DecodeString(data.Signature)
				if err != nil {
					t.Fatalf("failed to decode signature: %v", err)
				}
				verifySolanaSignature(t, wallet.address, messageBytes, sig)
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

				sig, err := base64.StdEncoding.DecodeString(data.Signature)
				if err != nil {
					t.Fatalf("failed to decode signature: %v", err)
				}
				verifySolanaSignature(t, wallet.address, messageBytes, sig)
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
					SolanaSignTransactionRpcInputParams{
						Transaction: transaction,
						Encoding:    "base64",
					},
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

	// SignAndSendTransaction is skipped to not waste funds
	t.Run("SignAndSendTransaction", func(t *testing.T) {
		t.Skip("skipped to not waste funds")
		transaction := "AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAECBXJmGwQcj3yBWbVmZjP7TJdCdPRYRVyxV3BCR/j1AxwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQECAAAMAgAAAADh9QUAAAAAAA=="

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Solana.SignAndSendTransaction(ctx, wallet.id,
					"solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1", // Solana devnet
					SolanaSignAndSendTransactionRpcInputParams{
						Transaction: transaction,
						Encoding:    "base64",
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign and send transaction: %v", err)
				}

				if data.Hash == "" {
					t.Error("expected hash to be defined")
				}
				if data.Caip2 != "solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1" {
					t.Errorf("expected caip2 to be solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1, got %s", data.Caip2)
				}
			})
		}
	})

	// SignAndSendTransactionBytes is skipped to not waste funds
	t.Run("SignAndSendTransactionBytes", func(t *testing.T) {
		t.Skip("skipped to not waste funds")
		transactionB64 := "AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAECBXJmGwQcj3yBWbVmZjP7TJdCdPRYRVyxV3BCR/j1AxwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQECAAAMAgAAAADh9QUAAAAAAA=="
		transaction, err := base64.StdEncoding.DecodeString(transactionB64)
		if err != nil {
			t.Fatalf("failed to decode test transaction: %v", err)
		}

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Solana.SignAndSendTransactionBytes(ctx, wallet.id,
					"solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1", // Solana devnet
					transaction,
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign and send transaction bytes: %v", err)
				}

				if data.Hash == "" {
					t.Error("expected hash to be defined")
				}
				if data.Caip2 != "solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1" {
					t.Errorf("expected caip2 to be solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1, got %s", data.Caip2)
				}
			})
		}
	})
}
