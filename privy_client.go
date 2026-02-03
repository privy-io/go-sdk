// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

import (
	"github.com/privy-io/go-sdk/option"
)

// PrivyClient is an enhanced client with custom functionality.
// It wraps the generated Client and provides access to enhanced services.
//
// This is the recommended client type for most use cases as it provides
// wrapped services with additional functionality while maintaining full
// compatibility with the base Client.
//
// Example:
//
//	client := privyapiclient.NewPrivyClient(
//	    option.WithAppID("my-app-id"),
//	    option.WithAppSecret("my-app-secret"),
//	    option.WithEnvironmentStaging(),
//	)
//	wallet, err := client.Wallets.Get(ctx, "wallet_id")
type PrivyClient struct {
	// Client is the underlying generated client.
	// You can use this to access any services not yet wrapped by PrivyClient.
	Client Client

	// Wallets provides access to wallet operations with enhanced functionality.
	Wallets *PrivyWalletService
}

// NewPrivyClient creates a new enhanced Privy client.
// This is the recommended way to create a client for most use cases.
//
// It accepts the same options as the generated client:
//
//	client := privyapiclient.NewPrivyClient(
//	    option.WithAppID("my-app-id"),
//	    option.WithAppSecret("my-app-secret"),
//	    option.WithEnvironmentStaging(),
//	)
func NewPrivyClient(opts ...option.RequestOption) *PrivyClient {
	client := NewClient(opts...)
	return &PrivyClient{
		Client:  client,
		Wallets: NewPrivyWalletService(client.Wallets),
	}
}

// Add your custom client methods below this line.
// Example:
//
// func (c *PrivyClient) CustomMethod() error {
//     // Your custom logic here
//     return nil
// }
