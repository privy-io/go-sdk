package e2e_test

import (
	"context"
	"testing"

	. "github.com/privy-io/go-sdk"
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
}
