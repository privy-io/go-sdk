// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/internal/testutil"
	"github.com/privy-io/go-sdk/option"
)

func TestWalletSwapExecuteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Swap.Execute(
		context.TODO(),
		"wallet_id",
		privyclient.WalletSwapExecuteParams{
			SwapRequestBody: privyclient.SwapRequestBody{
				BaseAmount: "1000000000000000000",
				Destination: privyclient.SwapDestination{
					AssetAddress:       "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					Caip2:              privyclient.String("eip155:1"),
					DestinationAddress: privyclient.String("destination_address"),
				},
				Source: privyclient.SwapSource{
					AssetAddress: "native",
					Caip2:        "eip155:1",
				},
				AmountType: privyclient.AmountTypeExactInput,
				FeeConfiguration: privyclient.FeeConfiguration{
					Type:  privyclient.FeeConfigurationTypeTotalFeeBps,
					Value: 50,
				},
				SlippageBps: privyclient.Float(50),
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletSwapQuoteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Swap.Quote(
		context.TODO(),
		"wallet_id",
		privyclient.WalletSwapQuoteParams{
			SwapQuoteRequestBody: privyclient.SwapQuoteRequestBody{
				BaseAmount: "1000000000000000000",
				Destination: privyclient.SwapQuoteDestination{
					AssetAddress:       "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					Caip2:              privyclient.String("eip155:1"),
					DestinationAddress: privyclient.String("destination_address"),
				},
				Source: privyclient.SwapSource{
					AssetAddress: "native",
					Caip2:        "eip155:1",
				},
				AmountType: privyclient.AmountTypeExactInput,
				FeeConfiguration: privyclient.FeeConfiguration{
					Type:  privyclient.FeeConfigurationTypeTotalFeeBps,
					Value: 50,
				},
				SlippageBps: privyclient.Float(0),
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
