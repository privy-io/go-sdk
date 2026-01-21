// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/privy-api-client-go"
	"github.com/stainless-sdks/privy-api-client-go/internal/testutil"
	"github.com/stainless-sdks/privy-api-client-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Wallets.ListAutoPaging(context.TODO(), privyapiclient.WalletListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		wallet := iter.Current()
		t.Logf("%+v\n", wallet.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
