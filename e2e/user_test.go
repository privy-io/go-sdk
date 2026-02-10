package e2e_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/privy-io/go-sdk"
)

func TestUsers(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("New", func(t *testing.T) {
		t.Run("WithCustomAuthAndWallet", func(t *testing.T) {
			t.Skip("skipped to avoid creating test resources")

			customUserID := fmt.Sprintf("test-user-%d", time.Now().UnixNano())

			user, err := client.Users.New(ctx, UserNewParams{
				LinkedAccounts: []LinkedAccountInputUnionParam{
					{
						OfCustomAuth: &LinkedAccountCustomJwtInputParam{
							Type:         LinkedAccountCustomJwtInputTypeCustomAuth,
							CustomUserID: customUserID,
						},
					},
				},
				Wallets: []UserNewParamsWallet{
					{ChainType: WalletChainTypeEthereum},
				},
			})
			if err != nil {
				t.Fatalf("failed to create user: %v", err)
			}

			if user.ID == "" {
				t.Error("expected user ID to be defined")
			}

			var foundCustomAuth bool
			var foundWallet bool

			for _, account := range user.LinkedAccounts {
				switch account.Type {
				case "custom_auth":
					foundCustomAuth = true
					customJwt := account.AsLinkedAccountCustomJwt()
					if customJwt.CustomUserID != customUserID {
						t.Errorf("expected custom_user_id to be %s, got %s", customUserID, customJwt.CustomUserID)
					}
				case "wallet":
					foundWallet = true
					if account.ChainType != "ethereum" {
						t.Errorf("expected wallet chain_type to be ethereum, got %s", account.ChainType)
					}
					if account.Address == "" {
						t.Error("expected wallet address to be defined")
					}
				}
			}

			if !foundCustomAuth {
				t.Error("expected to find custom_auth linked account")
			}
			if !foundWallet {
				t.Error("expected to find wallet linked account")
			}
		})
	})
}
