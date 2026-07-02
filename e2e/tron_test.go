package e2e_test

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	. "github.com/privy-io/go-sdk"
)

// verifyTronSignature verifies a Tron secp256k1 signature.
//
// signedTransactionHex is the hex-encoded signed transaction returned by
// tron_signTransaction: raw_data_hex followed by a 65-byte (130 hex char)
// signature in r(32) || s(32) || v(1) format, where v is 27 or 28.
//
// expectedPublicKeyHex is the compressed secp256k1 public key returned by
// the wallet API (wallet.PublicKey).
func verifyTronSignature(t *testing.T, signedTransactionHex, expectedPublicKeyHex string) {
	t.Helper()

	if len(signedTransactionHex) < 130 {
		t.Fatalf("signed transaction too short (%d hex chars, need at least 130)", len(signedTransactionHex))
	}

	// Last 130 hex chars = 65 bytes: r(32) + s(32) + v(1)
	splitAt := len(signedTransactionHex) - 130
	rawDataHex := signedTransactionHex[:splitAt]
	sigHex := signedTransactionHex[splitAt:]

	sigBytes, err := hex.DecodeString(sigHex)
	if err != nil {
		t.Fatalf("failed to decode signature hex: %v", err)
	}
	rawDataBytes, err := hex.DecodeString(rawDataHex)
	if err != nil {
		t.Fatalf("failed to decode raw_data hex: %v", err)
	}

	// txID = SHA256(raw_data_bytes)
	txID := sha256.Sum256(rawDataBytes)

	// Tron signature format: r(32) || s(32) || v(1)
	// btcec RecoverCompact expects: [recovery_code, r(32), s(32)]
	// where recovery_code = 27 + 4 (compressed) + (v - 27)
	v := sigBytes[64]
	recoveryID := int(v)
	if recoveryID >= 27 {
		recoveryID -= 27
	}
	// Build compact signature: [27 + 4 + recoveryID, r(32), s(32)]
	compact := make([]byte, 65)
	compact[0] = byte(27 + 4 + recoveryID)
	copy(compact[1:33], sigBytes[0:32])   // r
	copy(compact[33:65], sigBytes[32:64]) // s

	recoveredPubKey, _, err := ecdsa.RecoverCompact(compact, txID[:])
	if err != nil {
		t.Fatalf("failed to recover public key from signature: %v", err)
	}

	recoveredHex := hex.EncodeToString(recoveredPubKey.SerializeCompressed())
	if recoveredHex != expectedPublicKeyHex {
		t.Errorf("recovered public key %s does not match expected %s", recoveredHex, expectedPublicKeyHex)
	}
}

// tronTransferRawData builds a minimal TronRawDataForSign for a TRX transfer.
func tronTransferRawData(ownerAddressHex string) TronRawDataForSign {
	now := time.Now().UnixMilli()
	return TronRawDataForSign{
		Contract: []TronContractUnion{
			{
				OfTransferContract: &TronTransferContract{
					Type:         TronTransferContractTypeTransferContract,
					OwnerAddress: ownerAddressHex,
					ToAddress:    "410000000000000000000000000000000000000000",
					Amount:       1,
				},
			},
		},
		RefBlockBytes: "1a2b",
		RefBlockHash:  "abc1234567890def",
		Expiration:    now + 60000,
	}
}

func TestWallets_Tron(t *testing.T) {
	client := newTestClient(t)
	res := setupTestWalletResources(t, client)
	ctx := res.ctx
	wallets := res.createTestWallets(t, WalletChainTypeTron)

	t.Run("SignTransaction", func(t *testing.T) {
		for _, wallet := range wallets {
			wallet := wallet
			t.Run(wallet.name, func(t *testing.T) {
				// wallet.Address is base58check (T...), but TronTransferContract
				// requires hex (41...). The wallet's PublicKey is the compressed
				// secp256k1 public key in hex; the server derives the address from
				// it. For the owner_address we use a known placeholder so the
				// raw_data is structurally valid — the server signs whatever we
				// provide.
				rawData := tronTransferRawData("410000000000000000000000000000000000000001")

				data, err := client.Wallets.Tron.SignTransaction(ctx, wallet.id,
					TronSignTransactionRpcInputParams{
						RawData: rawData,
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign transaction: %v", err)
				}

				if data.SignedTransaction == "" {
					t.Error("expected signed_transaction to be non-empty")
				}
				if data.Encoding != "hex" {
					t.Errorf("expected encoding to be hex, got %s", data.Encoding)
				}

				// Verify the signature cryptographically
				verifyTronSignature(t, data.SignedTransaction, wallet.publicKey)
			})
		}
	})

	// SendTransaction is skipped to not waste funds.
	t.Run("SendTransaction", func(t *testing.T) {
		t.Skip("skipped to not waste funds")

		for _, wallet := range wallets {
			wallet := wallet
			t.Run(wallet.name, func(t *testing.T) {
				caip2 := "tron:0xcd8690dc" // Nile testnet
				rawData := TronRawDataForSend{
					Contract: []TronContractUnion{
						{
							OfTransferContract: &TronTransferContract{
								Type:         TronTransferContractTypeTransferContract,
								OwnerAddress: "410000000000000000000000000000000000000001",
								ToAddress:    "410000000000000000000000000000000000000000",
								Amount:       1,
							},
						},
					},
				}

				data, err := client.Wallets.Tron.SendTransaction(ctx, wallet.id, caip2,
					TronSendTransactionRpcInputParams{
						RawData: rawData,
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to send transaction: %v", err)
				}

				if data.Hash == "" {
					t.Error("expected hash to be non-empty")
				}
				if data.TransactionID == "" {
					t.Error("expected transaction_id to be non-empty")
				}
				if data.Caip2 != caip2 {
					t.Errorf("expected caip2 to be %s, got %s", caip2, data.Caip2)
				}
			})
		}
	})
}
