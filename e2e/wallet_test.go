package e2e_test

import (
	"context"
	"os"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/packages/param"
)

func TestWallets(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("New", func(t *testing.T) {
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
		walletID := os.Getenv("P256_OWNED_ETHEREUM_WALLET_ID")
		pk := os.Getenv("P256_PUBLIC_KEY")
		sk := os.Getenv("P256_PRIVATE_KEY")
		authCtx := &authorization.AuthorizationContext{
			PrivateKeys: []string{sk},
		}

		// Update the wallet to be ownerless
		result, err := client.Wallets.Update(
			ctx,
			walletID,
			WalletUpdateParams{
				Owner: param.NullStruct[WalletUpdateParamsOwnerUnion](),
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
				Owner: WalletUpdateParamsOwnerUnion{
					OfPublicKeyOwner: &WalletUpdateParamsOwnerPublicKeyOwner{PublicKey: pk},
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

	t.Run("RawSign", func(t *testing.T) {
		walletID := os.Getenv("P256_OWNED_TRON_WALLET_ID")
		sk := os.Getenv("P256_PRIVATE_KEY")
		authCtx := &authorization.AuthorizationContext{
			PrivateKeys: []string{sk},
		}

		// A 32-byte hash (keccak256 of "hello") in hex, prefixed with 0x
		hash := "0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8"

		// Call RawSign with authorization
		result, err := client.Wallets.RawSign(
			ctx,
			walletID,
			WalletRawSignParams{
				Params: WalletRawSignParamsParamsUnion{
					OfHash: &WalletRawSignParamsParamsHash{
						Hash: hash,
					},
				},
			},
			WithAuthorizationContext(authCtx),
		)
		if err != nil {
			t.Fatalf("failed to raw sign: %v", err)
		}

		if result.Data.Signature == "" {
			t.Error("expected signature to be non-empty")
		}
		if result.Method != WalletRawSignResponseMethodRawSign {
			t.Errorf("expected method to be raw_sign, got %s", result.Method)
		}
	})
}
