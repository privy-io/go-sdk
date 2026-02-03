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

func TestWalletBalanceGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Wallets.Balance.Get(
		context.TODO(),
		"wallet_id",
		privyclient.WalletBalanceGetParams{
			Asset: privyclient.WalletBalanceGetParamsAssetUnion{
				OfWalletBalanceGetsAssetString: privyclient.String("usdc"),
			},
			Chain: privyclient.WalletBalanceGetParamsChainUnion{
				OfWalletBalanceGetsChainString: privyclient.String("ethereum"),
			},
			IncludeCurrency: privyclient.WalletBalanceGetParamsIncludeCurrencyUsd,
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
