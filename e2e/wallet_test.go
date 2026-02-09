package e2e_test

import (
	"context"
	"os"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
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
		t.Skip("skipped to avoid changing test resources")
		walletID := os.Getenv("P256_OWNED_ETHEREUM_WALLET_ID")
		sk := os.Getenv("P256_PRIVATE_KEY")
		authCtx := &authorization.AuthorizationContext{
			PrivateKeys: []string{sk},
		}

		// Call Update with authorization (empty params just validates auth works)
		result, err := client.Wallets.Update(
			ctx,
			walletID,
			WalletUpdateParams{},
			WithAuthorizationContext(authCtx),
		)
		if err != nil {
			t.Fatalf("failed to update wallet: %v", err)
		}

		if result.ID != walletID {
			t.Errorf("expected wallet ID %s, got %s", walletID, result.ID)
		}
	})
}
