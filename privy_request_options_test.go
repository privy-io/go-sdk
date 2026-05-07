package privyclient_test

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	privyclient "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/option"
)

func TestRequestExpiryDisabled(t *testing.T) {
	var capturedReq *http.Request
	customClient := &http.Client{
		Transport: &closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				capturedReq = req
				return &http.Response{StatusCode: http.StatusOK}, nil
			},
		},
	}

	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:                "test-app-id",
		AppSecret:            "test-app-secret",
		DisableRequestExpiry: true,
		HTTPClient:           customClient,
	})

	_, _ = client.Wallets.Rpc(
		context.Background(),
		"wallet-id",
		privyclient.WalletRpcParams{},
	)

	if capturedReq == nil {
		t.Fatal("expected request to be captured")
	}
	if got := capturedReq.Header.Get("Privy-Request-Expiry"); got != "" {
		t.Errorf("expected no privy-request-expiry header when disabled, got %q", got)
	}
}

func TestRequestExpiryDisabledSupersedesDefaultRequestExpiryMs(t *testing.T) {
	var capturedReq *http.Request
	customClient := &http.Client{
		Transport: &closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				capturedReq = req
				return &http.Response{StatusCode: http.StatusOK}, nil
			},
		},
	}

	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:                  "test-app-id",
		AppSecret:              "test-app-secret",
		DefaultRequestExpiryMs: 30 * 60 * 1000,
		DisableRequestExpiry:   true,
		HTTPClient:             customClient,
	})

	_, _ = client.Wallets.Rpc(
		context.Background(),
		"wallet-id",
		privyclient.WalletRpcParams{},
	)

	if capturedReq == nil {
		t.Fatal("expected request to be captured")
	}
	if got := capturedReq.Header.Get("Privy-Request-Expiry"); got != "" {
		t.Errorf("expected DisableRequestExpiry to supersede DefaultRequestExpiryMs, got header %q", got)
	}
}

func TestRequestExpiryEnabledByDefault(t *testing.T) {
	var capturedReq *http.Request
	customClient := &http.Client{
		Transport: &closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				capturedReq = req
				return &http.Response{StatusCode: http.StatusOK}, nil
			},
		},
	}

	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:      "test-app-id",
		AppSecret:  "test-app-secret",
		HTTPClient: customClient,
	})

	_, _ = client.Wallets.Rpc(
		context.Background(),
		"wallet-id",
		privyclient.WalletRpcParams{},
	)

	if capturedReq == nil {
		t.Fatal("expected request to be captured")
	}
	if got := capturedReq.Header.Get("Privy-Request-Expiry"); got == "" {
		t.Error("expected privy-request-expiry header to be set by default")
	}
}

func TestRequestExpiryUsesDefaultMsFromNested(t *testing.T) {
	var capturedReq *http.Request
	customClient := &http.Client{
		Transport: &closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				capturedReq = req
				return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
			},
		},
	}

	tenMinutes := int64(10 * 60 * 1000)
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:     "test-app-id",
		AppSecret: "test-app-secret",
		RequestExpiry: privyclient.PrivyRequestExpiryOptions{
			DefaultMs: tenMinutes,
		},
		HTTPClient: customClient,
	})

	_, _ = client.Wallets.Rpc(context.Background(), "wallet-id", privyclient.WalletRpcParams{})

	if capturedReq == nil {
		t.Fatal("expected request to be captured")
	}
	raw := capturedReq.Header.Get("Privy-Request-Expiry")
	if raw == "" {
		t.Fatal("expected privy-request-expiry header to be set")
	}
	expiry, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		t.Fatalf("invalid privy-request-expiry header %q: %v", raw, err)
	}
	now := time.Now().UnixMilli()
	if expiry < now+tenMinutes-2000 || expiry > now+tenMinutes+2000 {
		t.Errorf("expected expiry ≈ now+10min, got %d (now=%d)", expiry, now)
	}
}

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
