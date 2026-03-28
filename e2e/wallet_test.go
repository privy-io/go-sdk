package e2e_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	randv2 "math/rand/v2"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/base58"
	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/packages/param"
	"golang.org/x/crypto/sha3"
)

func TestWallets(t *testing.T) {
	client := newTestClient(t)
	res := setupTestWalletResources(t, client)
	ctx := res.ctx

	t.Run("New", func(t *testing.T) {
		chainTypes := []struct {
			name  string
			value WalletChainType
		}{
			{name: "Ethereum", value: WalletChainTypeEthereum},
			{name: "Solana", value: WalletChainTypeSolana},
			{name: "Tron", value: WalletChainTypeTron},
		}

		owners := []struct {
			name  string
			value WalletNewParamsOwnerUnion
			ID    param.Opt[string]
		}{
			{name: "Ownerless", value: param.NullStruct[WalletNewParamsOwnerUnion]()},
			{name: "P256Owned", value: WalletNewParamsOwnerUnion{
				OfPublicKeyOwner: &WalletNewParamsOwnerPublicKeyOwner{PublicKey: res.p256KeyPair.PublicKey},
			}},
			{name: "UserOwned", value: WalletNewParamsOwnerUnion{
				OfUserOwner: &WalletNewParamsOwnerUserOwner{UserID: res.userID},
			}},
			{name: "KeyAndUserQuorumOwned", value: param.NullStruct[WalletNewParamsOwnerUnion](), ID: String(res.quorumID)},
		}
		for _, chainType := range chainTypes {
			t.Run(chainType.name, func(t *testing.T) {
				for _, owner := range owners {
					t.Run(owner.name, func(t *testing.T) {
						t.Skip("skipped to avoid creating test resources")

						wallet, err := client.Wallets.New(ctx, WalletNewParams{
							ChainType: chainType.value,
							Owner:     owner.value,
							OwnerID:   owner.ID,
						})
						if err != nil {
							t.Fatalf("failed to create wallet: %v", err)
						}

						if wallet.ID == "" {
							t.Error("expected wallet ID to be defined")
						}
						t.Logf("Created %s_%s wallet", owner.name, chainType.name)
						t.Logf("%s_%s_WALLET_ID=%s", owner.name, chainType.name, wallet.ID)
						t.Logf("%s_%s_WALLET_ADDRESS=%s", owner.name, chainType.name, wallet.Address)
						t.Logf("%s_%s_WALLET_PUBLIC_KEY=%s", owner.name, chainType.name, wallet.PublicKey)
					})
				}
			})
		}

		t.Run("Ethereum", func(t *testing.T) {
			t.Skip("skipped to avoid creating test resources")

			wallet, err := client.Wallets.New(ctx, WalletNewParams{
				ChainType: WalletChainTypeEthereum,
			})
			if err != nil {
				t.Fatalf("failed to create ethereum wallet: %v", err)
			}

			if wallet.ID == "" {
				t.Error("expected wallet ID to be defined")
			}
			if wallet.Address == "" {
				t.Error("expected wallet address to be defined")
			}
			if wallet.ChainType != WalletChainTypeEthereum {
				t.Errorf("expected chain_type to be ethereum, got %s", wallet.ChainType)
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		// Create a P256-owned wallet on demand
		wallet, err := client.Wallets.New(ctx, WalletNewParams{
			ChainType: WalletChainTypeEthereum,
			Owner: WalletNewParamsOwnerUnion{
				OfPublicKeyOwner: &WalletNewParamsOwnerPublicKeyOwner{PublicKey: res.p256KeyPair.PublicKey},
			},
		})
		if err != nil {
			t.Fatalf("failed to create wallet: %v", err)
		}
		walletID := wallet.ID
		authCtx := &authorization.AuthorizationContext{
			PrivateKeys: []string{res.p256KeyPair.PrivateKey},
		}

		// Update the wallet to be ownerless
		result, err := client.Wallets.Update(
			ctx,
			walletID,
			WalletUpdateParams{
				WalletUpdateRequestBody: WalletUpdateRequestBody{
					Owner: param.NullStruct[WalletUpdateRequestBodyOwnerUnion](),
				},
			},
			WithAuthorizationContext(authCtx),
		)
		if err != nil {
			t.Fatalf("failed to update wallet: %v", err)
		}

		if result.ID != walletID {
			t.Fatalf("expected wallet ID %s, got %s", walletID, result.ID)
		}

		// Check no OwnerID (empty string)
		if result.OwnerID != "" {
			t.Fatalf("expected wallet to have no owner, got %s", result.OwnerID)
		}

		// Update the wallet back to be owned by the P256 key
		result2, err := client.Wallets.Update(
			ctx,
			walletID,
			WalletUpdateParams{
				WalletUpdateRequestBody: WalletUpdateRequestBody{
					Owner: WalletUpdateRequestBodyOwnerUnion{
						OfPublicKeyOwner: &WalletUpdateRequestBodyOwnerPublicKeyOwner{PublicKey: res.p256KeyPair.PublicKey},
					},
				},
			},
			WithAuthorizationContext(authCtx),
		)
		if err != nil {
			t.Fatalf("failed to update wallet: %v", err)
		}

		if result2.ID != walletID {
			t.Errorf("expected wallet ID %s, got %s", walletID, result2.ID)
		}
	})

	t.Run("Export", func(t *testing.T) {
		chainTypes := []struct {
			name  string
			value WalletChainType
		}{
			{name: "Ethereum", value: WalletChainTypeEthereum},
			{name: "Solana", value: WalletChainTypeSolana},
			{name: "Tron", value: WalletChainTypeTron},
		}

		for _, chainType := range chainTypes {
			t.Run(chainType.name, func(t *testing.T) {
				wallets := res.createTestWallets(t, chainType.value)
				for _, wallet := range wallets {
					if wallet.authCtx == nil {
						continue // ownerless wallets cannot be exported
					}
					t.Run(wallet.name, func(t *testing.T) {
						result, err := client.Wallets.Export(
							ctx,
							wallet.id,
							WithAuthorizationContext(wallet.authCtx),
						)
						if err != nil {
							t.Fatalf("failed to export wallet: %v", err)
						}

						if result.PrivateKey == "" {
							t.Error("expected private key to be non-empty")
						}

						if chainType.value == WalletChainTypeEthereum {
							hexKey := strings.TrimPrefix(result.PrivateKey, "0x")
							privKeyBytes, err := hex.DecodeString(hexKey)
							if err != nil {
								t.Fatalf("failed to decode private key hex: %v", err)
							}
							_, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
							pubKeyBytes := pubKey.SerializeUncompressed()
							hasher := sha3.NewLegacyKeccak256()
							hasher.Write(pubKeyBytes[1:])
							hash := hasher.Sum(nil)
							derivedAddress := "0x" + hex.EncodeToString(hash[len(hash)-20:])
							if !strings.EqualFold(derivedAddress, wallet.address) {
								t.Errorf("expected derived address %s to match wallet address %s", derivedAddress, wallet.address)
							}
						}

						if chainType.value == WalletChainTypeSolana {
							privKeyBytes := base58.Decode(result.PrivateKey)
							privKey := ed25519.PrivateKey(privKeyBytes)
							pubKey := privKey.Public().(ed25519.PublicKey)
							derivedAddress := base58.Encode(pubKey)
							if derivedAddress != wallet.address {
								t.Errorf("expected derived address %s to match wallet address %s", derivedAddress, wallet.address)
							}
						}

						if chainType.value == WalletChainTypeTron {
							hexKey := strings.TrimPrefix(result.PrivateKey, "0x")
							privKeyBytes, err := hex.DecodeString(hexKey)
							if err != nil {
								t.Fatalf("failed to decode private key hex: %v", err)
							}
							_, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
							derivedPubKey := hex.EncodeToString(pubKey.SerializeCompressed())
							if derivedPubKey != wallet.publicKey {
								t.Errorf("expected derived public key %s to match wallet public key %s", derivedPubKey, wallet.publicKey)
							}
						}
					})
				}
			})
		}
	})

	t.Run("Import", func(t *testing.T) {
		t.Run("PrivateKey", func(t *testing.T) {
			t.Run("Ethereum", func(t *testing.T) {
				// Generate a fresh secp256k1 keypair
				privKey, err := btcec.NewPrivateKey()
				if err != nil {
					t.Fatalf("failed to generate private key: %v", err)
				}
				privKeyBytes := privKey.Serialize()
				pubKeyBytes := privKey.PubKey().SerializeUncompressed()
				hasher := sha3.NewLegacyKeccak256()
				hasher.Write(pubKeyBytes[1:])
				hash := hasher.Sum(nil)
				address := "0x" + hex.EncodeToString(hash[len(hash)-20:])

				imported, err := client.Wallets.Import(ctx, WalletImportParams{
					Wallet: WalletImportParamsWalletUnion{
						OfPrivateKey: &WalletImportParamsWalletPrivateKey{
							Address:    address,
							ChainType:  string(WalletChainTypeEthereum),
							PrivateKey: privKeyBytes,
						},
					},
				})
				if err != nil {
					t.Fatalf("failed to import wallet: %v", err)
				}

				if imported.ID == "" {
					t.Error("expected imported wallet ID to be non-empty")
				}
				if imported.Address == "" {
					t.Error("expected imported wallet address to be non-empty")
				}
				if !strings.EqualFold(imported.Address, address) {
					t.Errorf("expected imported address %s to match %s", imported.Address, address)
				}
				if imported.ChainType != WalletChainTypeEthereum {
					t.Errorf("expected chain type ethereum, got %s", imported.ChainType)
				}
			})

			t.Run("Solana", func(t *testing.T) {
				// Generate a fresh ed25519 keypair
				pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
				if err != nil {
					t.Fatalf("failed to generate keypair: %v", err)
				}
				address := base58.Encode(pubKey)

				imported, err := client.Wallets.Import(ctx, WalletImportParams{
					Wallet: WalletImportParamsWalletUnion{
						OfPrivateKey: &WalletImportParamsWalletPrivateKey{
							Address:    address,
							ChainType:  string(WalletChainTypeSolana),
							PrivateKey: []byte(privKey),
						},
					},
				})
				if err != nil {
					t.Fatalf("failed to import wallet: %v", err)
				}

				if imported.ID == "" {
					t.Error("expected imported wallet ID to be non-empty")
				}
				if imported.Address == "" {
					t.Error("expected imported wallet address to be non-empty")
				}
				if imported.Address != address {
					t.Errorf("expected imported address %s to match %s", imported.Address, address)
				}
				if imported.ChainType != WalletChainTypeSolana {
					t.Errorf("expected chain type solana, got %s", imported.ChainType)
				}
			})
		})

		t.Run("MnemonicHD", func(t *testing.T) {
			mnemonic := generateMnemonic(t)
			n := randv2.IntN(98) + 2 // random index >= 2

			indices := []struct {
				name  string
				index int
			}{
				{"Index=0", 0},
				{"Index=1", 1},
				{fmt.Sprintf("Index=%d", n), n},
			}

			chains := []struct {
				name            string
				chainType       WalletChainType
				deriveAddr      func(*testing.T, string, int) string
				caseInsensitive bool
			}{
				{"Ethereum", WalletChainTypeEthereum, deriveEthAddressFromMnemonic, true},
				{"Solana", WalletChainTypeSolana, deriveSolAddressFromMnemonic, false},
			}

			for _, idx := range indices {
				t.Run(idx.name, func(t *testing.T) {
					for _, chain := range chains {
						t.Run(chain.name, func(t *testing.T) {
							address := chain.deriveAddr(t, mnemonic, idx.index)

							imported, err := client.Wallets.Import(ctx, WalletImportParams{
								Wallet: WalletImportParamsWalletUnion{
									OfHD: &WalletImportParamsWalletHD{
										Address:    address,
										ChainType:  string(chain.chainType),
										Index:      int64(idx.index),
										PrivateKey: []byte(mnemonic),
									},
								},
							})
							if err != nil {
								t.Fatalf("failed to import HD wallet: %v", err)
							}

							if imported.ID == "" {
								t.Error("expected imported wallet ID to be non-empty")
							}
							if imported.Address == "" {
								t.Error("expected imported wallet address to be non-empty")
							}
							if chain.caseInsensitive {
								if !strings.EqualFold(imported.Address, address) {
									t.Errorf("expected imported address %s to match %s", imported.Address, address)
								}
							} else {
								if imported.Address != address {
									t.Errorf("expected imported address %s to match %s", imported.Address, address)
								}
							}
							if imported.ChainType != chain.chainType {
								t.Errorf("expected chain type %s, got %s", chain.chainType, imported.ChainType)
							}
						})
					}
				})
			}
		})
	})

	t.Run("RawSign", func(t *testing.T) {
		wallets := res.createTestWallets(t, WalletChainTypeTron)
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				// A 32-byte hash (keccak256 of "hello") in hex, prefixed with 0x
				hash := "0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8"

				// Call RawSign with authorization
				result, err := client.Wallets.RawSign(
					ctx,
					wallet.id,
					WalletRawSignParams{
						RawSignInput: RawSignInput{
							Params: RawSignInputParamsUnion{
								OfRawSignHashs: &RawSignHashParams{
									Hash: hash,
								},
							},
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to raw sign: %v", err)
				}

				if result.Data.Signature == "" {
					t.Error("expected signature to be non-empty")
				}
				if result.Method != RawSignResponseMethodRawSign {
					t.Errorf("expected method to be raw_sign, got %s", result.Method)
				}
			})
		}
	})
}
