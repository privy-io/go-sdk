package privyclient_test

import (
	"context"
	"net/http"
	"testing"

	privyclient "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/option"
)

func TestWithRequestOptionsForwardsHTTPClient(t *testing.T) {
	called := false
	customClient := &http.Client{
		Transport: &closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				called = true
				return &http.Response{StatusCode: http.StatusOK}, nil
			},
		},
	}

	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:     "test-app-id",
		AppSecret: "test-app-secret",
	})

	// Call a wrapper method with WithRequestOptions to override the HTTP client.
	// The request will fail with a JSON decode error, but that's fine —
	// we only care that the custom transport was invoked.
	_, _ = client.Wallets.Rpc(
		context.Background(),
		"wallet-id",
		privyclient.WalletRpcParams{},
		privyclient.WithRequestOptions(option.WithHTTPClient(customClient)),
	)

	if !called {
		t.Error("expected custom HTTP client to be used via WithRequestOptions, but it was not called")
	}
}
