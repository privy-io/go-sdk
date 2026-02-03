//go:build e2e

package e2e_test

import (
	"context"
	"testing"

	. "github.com/privy-io/go-sdk"
)

func TestWallets_Create_Ethereum(t *testing.T) {
	client := newTestClient(t)

	wallet, err := client.Wallets.New(context.Background(), WalletNewParams{
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
}
