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

func TestWalletTransactionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Transactions.Get(
		context.TODO(),
		"wallet_id",
		privyclient.WalletTransactionGetParams{
			Chain: privyclient.TransactionChainNameInputEthereum,
			Token: privyclient.WalletTransactionGetParamsTokenUnion{
				OfString: privyclient.String("string"),
			},
			Asset: privyclient.WalletTransactionGetParamsAssetUnion{
				OfWalletTransactionGetsAssetString: privyclient.Opt(privyclient.WalletTransactionGetParamsAssetStringUsdc),
			},
			Cursor:          privyclient.String("x"),
			IncludeArchived: privyclient.Bool(true),
			Limit:           privyclient.Float(100),
			TxHash:          privyclient.String("tx_hash"),
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
