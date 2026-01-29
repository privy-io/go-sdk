// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/privy-api-client-go"
	"github.com/stainless-sdks/privy-api-client-go/internal/testutil"
	"github.com/stainless-sdks/privy-api-client-go/option"
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
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Balance.Get(
		context.TODO(),
		"wallet_id",
		privyapiclient.WalletBalanceGetParams{
			Token: privyapiclient.WalletBalanceGetParamsTokenUnion{
				OfString: privyapiclient.String("string"),
			},
			Asset: privyapiclient.WalletBalanceGetParamsAssetUnion{
				OfWalletBalanceGetsAssetString: privyapiclient.String("usdc"),
			},
			Chain: privyapiclient.WalletBalanceGetParamsChainUnion{
				OfWalletBalanceGetsChainString: privyapiclient.String("ethereum"),
			},
			IncludeCurrency: privyapiclient.WalletBalanceGetParamsIncludeCurrencyUsd,
		},
	)
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
