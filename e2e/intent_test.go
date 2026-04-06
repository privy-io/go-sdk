package e2e_test

import (
	"context"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
)

func TestIntents_UpdateWallet(t *testing.T) {
	client := newTestClient(t)
	res := setupTestWalletResources(t, client)
	ctx := res.ctx

	wallets := res.createTestWallets(t, WalletChainTypeEthereum)
	keyOwnedWallet := wallets[1] // KeyOwned wallet

	newPair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate new P256 key pair: %v", err)
	}

	intent, err := client.Intents.UpdateWallet(ctx, keyOwnedWallet.id, IntentUpdateWalletParams{
		WalletUpdateRequestBody: WalletUpdateRequestBody{
			Owner: OwnerInputUnionParam{
				OfOwnerInputPublicKey: &OwnerInputPublicKeyParam{
					PublicKey: newPair.PublicKey,
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create update wallet intent: %v", err)
	}

	if intent.IntentID == "" {
		t.Error("expected intent ID to be defined")
	}
	if intent.Status != IntentStatusPending {
		t.Errorf("expected status to be pending, got %s", intent.Status)
	}
	if intent.ResourceID != keyOwnedWallet.id {
		t.Errorf("expected resource_id %s, got %s", keyOwnedWallet.id, intent.ResourceID)
	}

	t.Logf("Created wallet update intent %s (status=%s)", intent.IntentID, intent.Status)

	t.Run("Get", func(t *testing.T) {
		result, err := client.Intents.Get(ctx, intent.IntentID)
		if err != nil {
			t.Fatalf("failed to get intent: %v", err)
		}

		if result.IntentID != intent.IntentID {
			t.Errorf("expected intent ID %s, got %s", intent.IntentID, result.IntentID)
		}
	})

	t.Run("List", func(t *testing.T) {
		page, err := client.Intents.List(ctx, IntentListParams{
			IntentType: IntentTypeWallet,
			ResourceID: String(keyOwnedWallet.id),
		})
		if err != nil {
			t.Fatalf("failed to list intents: %v", err)
		}

		intents := page.Data
		if len(intents) == 0 {
			t.Fatal("expected at least one intent in list")
		}

		found := false
		for _, item := range intents {
			if item.IntentID == intent.IntentID {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected to find intent %s in list", intent.IntentID)
		}
	})
}

func TestIntents_UpdatePolicy(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	pair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate key pair: %v", err)
	}
	authCtx := authorization.AuthorizationContext{
		PrivateKeys: []string{pair.PrivateKey},
	}

	policy := createPolicy(t, ctx, client, PolicyNewParams{
		Name:      "go-sdk-test-intent-policy",
		ChainType: WalletChainTypeEthereum,
		Version:   PolicyNewParamsVersion1_0,
		Rules:     []PolicyNewParamsRule{},
		Owner: OwnerInputUnionParam{
			OfOwnerInputPublicKey: &OwnerInputPublicKeyParam{
				PublicKey: pair.PublicKey,
			},
		},
	}, authCtx)

	intent, err := client.Intents.UpdatePolicy(ctx, policy.ID, IntentUpdatePolicyParams{
		Name: String("go-sdk-test-intent-policy-updated"),
	})
	if err != nil {
		t.Fatalf("failed to create update policy intent: %v", err)
	}

	if intent.IntentID == "" {
		t.Error("expected intent ID to be defined")
	}
	if intent.Status != IntentStatusPending {
		t.Errorf("expected status to be pending, got %s", intent.Status)
	}
	if intent.ResourceID != policy.ID {
		t.Errorf("expected resource_id %s, got %s", policy.ID, intent.ResourceID)
	}

	t.Logf("Created policy update intent %s (status=%s)", intent.IntentID, intent.Status)

	t.Run("Get", func(t *testing.T) {
		result, err := client.Intents.Get(ctx, intent.IntentID)
		if err != nil {
			t.Fatalf("failed to get intent: %v", err)
		}

		if result.IntentID != intent.IntentID {
			t.Errorf("expected intent ID %s, got %s", intent.IntentID, result.IntentID)
		}
	})
}

func TestIntents_PolicyRules(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	pair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate key pair: %v", err)
	}
	authCtx := authorization.AuthorizationContext{
		PrivateKeys: []string{pair.PrivateKey},
	}

	policy := createPolicy(t, ctx, client, PolicyNewParams{
		Name:      "go-sdk-test-intent-policy-rules",
		ChainType: WalletChainTypeEthereum,
		Version:   PolicyNewParamsVersion1_0,
		Rules:     []PolicyNewParamsRule{},
		Owner: OwnerInputUnionParam{
			OfOwnerInputPublicKey: &OwnerInputPublicKeyParam{
				PublicKey: pair.PublicKey,
			},
		},
	}, authCtx)

	t.Run("NewPolicyRule", func(t *testing.T) {
		intent, err := client.Intents.NewPolicyRule(ctx, policy.ID, IntentNewPolicyRuleParams{
			PolicyRuleRequestBody: PolicyRuleRequestBodyParam{
				Name:       "go-sdk-test-intent-rule",
				Action:     PolicyActionAllow,
				Method:     PolicyMethodStar,
				Conditions: []PolicyConditionUnionParam{},
			},
		})
		if err != nil {
			t.Fatalf("failed to create new policy rule intent: %v", err)
		}

		if intent.IntentID == "" {
			t.Error("expected intent ID to be defined")
		}
		if intent.Status != IntentStatusPending {
			t.Errorf("expected status to be pending, got %s", intent.Status)
		}

		t.Logf("Created new policy rule intent %s (status=%s)", intent.IntentID, intent.Status)
	})
}

func TestIntents_UpdateKeyQuorum(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	pair1, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate key pair 1: %v", err)
	}
	pair2, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate key pair 2: %v", err)
	}

	kq := createKeyQuorum(t, ctx, client, KeyQuorumNewParams{
		KeyQuorumCreateRequestBody: KeyQuorumCreateRequestBody{
			DisplayName: String("test-key-quorum"),
			PublicKeys:  []string{pair1.PublicKey, pair2.PublicKey},
		},
	}, authorization.AuthorizationContext{
		PrivateKeys: []string{pair1.PrivateKey, pair2.PrivateKey},
	})

	intent, err := client.Intents.UpdateKeyQuorum(ctx, kq.ID, IntentUpdateKeyQuorumParams{
		KeyQuorumUpdateRequestBody: KeyQuorumUpdateRequestBodyParam{
			DisplayName: String("go-sdk-test-intent-key-quorum-updated"),
		},
	})
	if err != nil {
		t.Fatalf("failed to create update key quorum intent: %v", err)
	}

	if intent.IntentID == "" {
		t.Error("expected intent ID to be defined")
	}
	if intent.Status != IntentStatusPending {
		t.Errorf("expected status to be pending, got %s", intent.Status)
	}
	if intent.ResourceID != kq.ID {
		t.Errorf("expected resource_id %s, got %s", kq.ID, intent.ResourceID)
	}

	t.Logf("Created key quorum update intent %s (status=%s)", intent.IntentID, intent.Status)
}

func TestIntents_List(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("ListAll", func(t *testing.T) {
		page, err := client.Intents.List(ctx, IntentListParams{
			Limit: Float(10),
		})
		if err != nil {
			t.Fatalf("failed to list intents: %v", err)
		}

		t.Logf("Found %d intents", len(page.Data))
	})

	t.Run("FilterByType", func(t *testing.T) {
		for _, intentType := range []IntentType{
			IntentTypeWallet,
			IntentTypePolicy,
			IntentTypeRule,
			IntentTypeRpc,
			IntentTypeKeyQuorum,
		} {
			t.Run(string(intentType), func(t *testing.T) {
				page, err := client.Intents.List(ctx, IntentListParams{
					IntentType: intentType,
					Limit:      Float(5),
				})
				if err != nil {
					t.Fatalf("failed to list %s intents: %v", intentType, err)
				}

				for _, intent := range page.Data {
					if intent.IntentType != string(intentType) {
						t.Errorf("expected intent_type %s, got %s", intentType, intent.IntentType)
					}
				}
			})
		}
	})

	t.Run("FilterByStatus", func(t *testing.T) {
		page, err := client.Intents.List(ctx, IntentListParams{
			Status: IntentStatusPending,
			Limit:  Float(5),
		})
		if err != nil {
			t.Fatalf("failed to list pending intents: %v", err)
		}

		for _, intent := range page.Data {
			if intent.Status != IntentStatusPending {
				t.Errorf("expected status pending, got %s", intent.Status)
			}
		}
	})
}

func TestIntents_Rpc(t *testing.T) {
	client := newTestClient(t)
	res := setupTestWalletResources(t, client)
	ctx := res.ctx

	wallets := res.createTestWallets(t, WalletChainTypeEthereum)
	keyOwnedWallet := wallets[1]

	intent, err := client.Intents.Rpc(ctx, keyOwnedWallet.id, IntentRpcParams{
		WalletRpcRequestBody: WalletRpcRequestBodyUnionParam{
			OfEthSignTransaction: &EthereumSignTransactionRpcInputParam{
				Method: EthereumSignTransactionRpcInputMethodEthSignTransaction,
				Params: EthereumSignTransactionRpcInputParams{
					Transaction: UnsignedEthereumTransactionParam{
						Type: 2,
						ChainID: QuantityUnionParam{
							OfInt: Int(1),
						},
						To: String("0x742d35Cc6634C0532925a3b8D1A8a9ff1e7a7A4C"),
						Value: QuantityUnionParam{
							OfString: String("0x1"),
						},
						GasLimit: QuantityUnionParam{
							OfString: String("0x5208"),
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create RPC intent: %v", err)
	}

	if intent.IntentID == "" {
		t.Error("expected intent ID to be defined")
	}
	if intent.Status != IntentStatusPending {
		t.Errorf("expected status to be pending, got %s", intent.Status)
	}
	if intent.ResourceID != keyOwnedWallet.id {
		t.Errorf("expected resource_id %s, got %s", keyOwnedWallet.id, intent.ResourceID)
	}

	t.Logf("Created RPC intent %s (status=%s)", intent.IntentID, intent.Status)

	t.Run("Get", func(t *testing.T) {
		result, err := client.Intents.Get(ctx, intent.IntentID)
		if err != nil {
			t.Fatalf("failed to get intent: %v", err)
		}

		if result.IntentID != intent.IntentID {
			t.Errorf("expected intent ID %s, got %s", intent.IntentID, result.IntentID)
		}
	})
}
