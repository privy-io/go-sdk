package e2e_test

import (
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/packages/param"
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
					OfPublicKeyOwner: &WalletUpdateParamsOwnerPublicKeyOwner{PublicKey: res.p256KeyPair.PublicKey},
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
						Params: WalletRawSignParamsParamsUnion{
							OfHash: &WalletRawSignParamsParamsHash{
								Hash: hash,
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
				if result.Method != WalletRawSignResponseMethodRawSign {
					t.Errorf("expected method to be raw_sign, got %s", result.Method)
				}
			})
		}
	})
}
