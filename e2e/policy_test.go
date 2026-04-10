package e2e_test

import (
	"context"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
)

func createPolicy(t *testing.T, ctx context.Context, client *PrivyClient, params PolicyNewParams, auth authorization.AuthorizationContext) *Policy {
	t.Helper()
	policy, err := client.Policies.New(ctx, params)
	if err != nil {
		t.Fatalf("failed to create policy: %v", err)
	}

	t.Logf("Created Policy %s", policy.ID)

	if policy.ID == "" {
		t.Error("expected policy ID to be defined")
	}

	if policy.Name != params.Name {
		t.Errorf("expected name to be %s, got %s", params.Name, policy.Name)
	}

	t.Cleanup(func() {
		t.Logf("Deleting Policy %s", policy.ID)
		result, err := client.Policies.Delete(
			ctx,
			policy.ID,
			PolicyDeleteParams{},
			WithAuthorizationContext(&auth),
		)
		if err != nil {
			t.Fatalf("failed to delete policy: %v", err)
		}

		if !result.Success {
			t.Error("expected delete to succeed")
		}
	})

	return policy
}

func createRule(t *testing.T, ctx context.Context, client *PrivyClient, policyID string, params PolicyNewRuleParams, auth authorization.AuthorizationContext) *PolicyRuleResponse {
	t.Helper()
	rule, err := client.Policies.NewRule(ctx, policyID, params, WithAuthorizationContext(&auth))
	if err != nil {
		t.Fatalf("failed to create rule: %v", err)
	}

	t.Logf("Created Rule %s on Policy %s", rule.ID, policyID)

	if rule.ID == "" {
		t.Error("expected rule ID to be defined")
	}

	if rule.Name != params.PolicyRuleRequestBody.Name {
		t.Errorf("expected name to be %s, got %s", params.PolicyRuleRequestBody.Name, rule.Name)
	}

	t.Cleanup(func() {
		t.Logf("Deleting Rule %s from Policy %s", rule.ID, policyID)
		result, err := client.Policies.DeleteRule(
			ctx,
			rule.ID,
			PolicyDeleteRuleParams{
				PolicyID: policyID,
			},
			WithAuthorizationContext(&auth),
		)
		if err != nil {
			t.Fatalf("failed to delete rule: %v", err)
		}

		if !result.Success {
			t.Error("expected delete to succeed")
		}
	})

	return rule
}

func TestPolicies(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	pair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate key pair: %v", err)
	}

	policy := createPolicy(t, ctx, client, PolicyNewParams{
		Name:      "go-sdk-test-policy",
		ChainType: WalletChainTypeEthereum,
		Version:   PolicyNewParamsVersion1_0,
		Rules:     []PolicyNewParamsRule{},
		Owner: OwnerInputUnion{
			OfOwnerInputPublicKey: &OwnerInputPublicKey{
				PublicKey: pair.PublicKey,
			},
		},
	}, authorization.AuthorizationContext{
		PrivateKeys: []string{pair.PrivateKey},
	})

	if policy.ChainType != WalletChainTypeEthereum {
		t.Errorf("expected chain_type to be ethereum, got %s", policy.ChainType)
	}

	t.Run("Update", func(t *testing.T) {
		authCtx := &authorization.AuthorizationContext{
			PrivateKeys: []string{pair.PrivateKey},
		}

		result, err := client.Policies.Update(
			ctx,
			policy.ID,
			PolicyUpdateParams{
				Name: String("go-sdk-test-policy-updated"),
			},
			WithAuthorizationContext(authCtx),
		)
		if err != nil {
			t.Fatalf("failed to update policy: %v", err)
		}

		if result.ID != policy.ID {
			t.Errorf("expected policy ID %s, got %s", policy.ID, result.ID)
		}
		if result.Name != "go-sdk-test-policy-updated" {
			t.Errorf("expected name 'go-sdk-test-policy-updated', got %s", result.Name)
		}
	})

	t.Run("Get", func(t *testing.T) {
		result, err := client.Policies.Get(ctx, policy.ID)
		if err != nil {
			t.Fatalf("failed to get policy: %v", err)
		}

		if result.ID != policy.ID {
			t.Errorf("expected policy ID %s, got %s", policy.ID, result.ID)
		}
	})
}

func TestPolicyRules(t *testing.T) {
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
		Name:      "go-sdk-test-policy-for-rules",
		ChainType: WalletChainTypeEthereum,
		Version:   PolicyNewParamsVersion1_0,
		Rules:     []PolicyNewParamsRule{},
		Owner: OwnerInputUnion{
			OfOwnerInputPublicKey: &OwnerInputPublicKey{
				PublicKey: pair.PublicKey,
			},
		},
	}, authCtx)

	rule := createRule(t, ctx, client, policy.ID, PolicyNewRuleParams{
		PolicyRuleRequestBody: PolicyRuleRequestBody{
			Name:       "go-sdk-test-rule",
			Action:     PolicyActionAllow,
			Method:     PolicyMethodStar,
			Conditions: []PolicyConditionUnion{},
		},
	}, authCtx)

	t.Run("UpdateRule", func(t *testing.T) {
		result, err := client.Policies.UpdateRule(
			ctx,
			rule.ID,
			PolicyUpdateRuleParams{
				PolicyID: policy.ID,
				PolicyRuleRequestBody: PolicyRuleRequestBody{
					Name:       "go-sdk-test-rule-updated",
					Action:     PolicyActionDeny,
					Method:     PolicyMethodStar,
					Conditions: []PolicyConditionUnion{},
				},
			},
			WithAuthorizationContext(&authCtx),
		)
		if err != nil {
			t.Fatalf("failed to update rule: %v", err)
		}

		if result.ID != rule.ID {
			t.Errorf("expected rule ID %s, got %s", rule.ID, result.ID)
		}
		if result.Name != "go-sdk-test-rule-updated" {
			t.Errorf("expected name 'go-sdk-test-rule-updated', got %s", result.Name)
		}
	})

	t.Run("GetRule", func(t *testing.T) {
		result, err := client.Policies.GetRule(ctx, rule.ID, PolicyGetRuleParams{
			PolicyID: policy.ID,
		})
		if err != nil {
			t.Fatalf("failed to get rule: %v", err)
		}

		if result.ID != rule.ID {
			t.Errorf("expected rule ID %s, got %s", rule.ID, result.ID)
		}
	})
}
