// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// PrivyClient is the main entrypoint for the Privy API Go SDK.
//
// Example:
//
//	client := privyclient.NewPrivyClient(
//	    option.WithAppID("my-app-id"),
//	    option.WithAppSecret("my-app-secret"),
//	    option.WithEnvironmentStaging(),
//	)
type PrivyClient struct {
	client Client

	Wallets      *PrivyWalletService
	Users        *PrivyUserService
	Policies     *PrivyPolicyService
	Transactions *PrivyTransactionService
	KeyQuorums   *PrivyKeyQuorumService
	ClientAuth   *PrivyClientAuthService
	Analytics    *PrivyAnalyticsService
	Apps         *PrivyAppService
	Aggregations *PrivyAggregationService
	Webhooks     *PrivyWebhookService
}

// NewPrivyClient creates a new enhanced Privy client.
// This is the recommended way to create a client for most use cases.
//
// It accepts the same options as the generated client:
//
//	client := privyclient.NewPrivyClient(
//	    option.WithAppID("my-app-id"),
//	    option.WithAppSecret("my-app-secret"),
//	    option.WithEnvironmentStaging(),
//	)
func NewPrivyClient(opts ...option.RequestOption) *PrivyClient {
	client := NewClient(opts...)
	return &PrivyClient{
		client:       client,
		Wallets:      newPrivyWalletService(client.Wallets),
		Users:        newPrivyUserService(client.Users),
		Policies:     newPrivyPolicyService(client.Policies),
		Transactions: newPrivyTransactionService(client.Transactions),
		KeyQuorums:   newPrivyKeyQuorumService(client.KeyQuorums),
		ClientAuth:   newPrivyClientAuthService(client.ClientAuth),
		Analytics:    newPrivyAnalyticsService(client.Analytics),
		Apps:         newPrivyAppService(client.Apps),
		Aggregations: newPrivyAggregationService(client.Aggregations),
		Webhooks:     newPrivyWebhookService(client.Webhooks),
	}
}
