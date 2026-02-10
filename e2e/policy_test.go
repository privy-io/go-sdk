package e2e_test

import (
	"context"
	"testing"

	. "github.com/privy-io/go-sdk"
)

func TestPolicies(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("New", func(t *testing.T) {
		t.Skip("skipped to avoid creating test resources")

		policy, err := client.Policies.New(ctx, PolicyNewParams{
			Name:      "test-policy",
			ChainType: PolicyNewParamsChainTypeEthereum,
			Version:   PolicyNewParamsVersion1_0,
			Rules:     []PolicyNewParamsRule{},
		})
		if err != nil {
			t.Fatalf("failed to create policy: %v", err)
		}

		if policy.ID == "" {
			t.Error("expected policy ID to be defined")
		}
		if policy.Name != "test-policy" {
			t.Errorf("expected name to be test-policy, got %s", policy.Name)
		}
		if policy.ChainType != PolicyChainTypeEthereum {
			t.Errorf("expected chain_type to be ethereum, got %s", policy.ChainType)
		}
	})
}
