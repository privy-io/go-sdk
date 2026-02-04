package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// PrivyClientOptions contains configuration options for creating a PrivyClient.
type PrivyClientOptions struct {
	// AppID is your Privy application ID (required).
	AppID string

	// AppSecret is your Privy application secret (required).
	AppSecret string

	// APIUrl is the base URL for the Privy API (optional).
	// If not provided, defaults to the production environment.
	// Use "https://auth.staging.privy.io" for staging.
	APIUrl string
}

// PrivyClient is the main entrypoint for the Privy API Go SDK.
//
// Example:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	    APIUrl:    "https://auth.staging.privy.io", // optional
//	})
type PrivyClient struct {
	client Client

	Wallets      *PrivyWalletService
	Users        *PrivyUserService
	Policies     *PrivyPolicyService
	Transactions *PrivyTransactionService
	KeyQuorums   *PrivyKeyQuorumService
	Analytics    *PrivyAnalyticsService
	Apps         *PrivyAppService
	Aggregations *PrivyAggregationService
	Webhooks     *PrivyWebhookService
}

// NewPrivyClient creates a new enhanced Privy client.
// This is the recommended way to create a client for most use cases.
//
// Example:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	})
//
// For staging environment:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	    APIUrl:    "https://auth.staging.privy.io",
//	})
func NewPrivyClient(opts PrivyClientOptions) *PrivyClient {
	// Build option.RequestOption slice from PrivyClientOptions
	requestOpts := []option.RequestOption{
		option.WithAppID(opts.AppID),
		option.WithAppSecret(opts.AppSecret),
	}

	// Add API URL if provided, otherwise use production default
	if opts.APIUrl != "" {
		requestOpts = append(requestOpts, option.WithBaseURL(opts.APIUrl))
	} else {
		requestOpts = append(requestOpts, option.WithEnvironmentProduction())
	}

	client := NewClient(requestOpts...)
	return &PrivyClient{
		client:       client,
		Wallets:      newPrivyWalletService(client.Wallets),
		Users:        newPrivyUserService(client.Users),
		Policies:     newPrivyPolicyService(client.Policies),
		Transactions: newPrivyTransactionService(client.Transactions),
		KeyQuorums:   newPrivyKeyQuorumService(client.KeyQuorums),
		Analytics:    newPrivyAnalyticsService(client.Analytics),
		Apps:         newPrivyAppService(client.Apps),
		Aggregations: newPrivyAggregationService(client.Aggregations),
		Webhooks:     newPrivyWebhookService(client.Webhooks),
	}
}
