// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient_test

import (
	"context"
	"fmt"

	privyclient "github.com/privy-io/go-sdk"
)

// ExampleNewPrivyClient demonstrates creating a PrivyClient with the options struct.
func ExampleNewPrivyClient() {
	// Production environment (default)
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:     "my-app-id",
		AppSecret: "my-app-secret",
	})

	// Use the client
	wallet, err := client.Wallets.Get(context.Background(), "wallet_123")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Wallet ID: %s\n", wallet.ID)
}

// ExampleNewPrivyClient_staging demonstrates using the staging environment.
func ExampleNewPrivyClient_staging() {
	// Staging environment
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:     "my-app-id",
		AppSecret: "my-app-secret",
		APIUrl:    "https://auth.staging.privy.io",
	})

	// Use the client
	wallet, err := client.Wallets.Get(context.Background(), "wallet_123")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Wallet ID: %s\n", wallet.ID)
}
