package privyclient_test

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	privyclient "github.com/privy-io/go-sdk"
)

const intentTestToleranceMs int64 = 2000

// captureIntentExpiryHeader runs client.Intents.Rpc with the provided call
// options and returns the value of the Privy-Request-Expiry header on the
// outgoing request, along with the timestamp (unix ms) recorded just before
// the call. The HTTP transport short-circuits with a 200 response so no real
// network traffic leaves the process.
func captureIntentExpiryHeader(
	t *testing.T,
	clientOpts privyclient.PrivyClientOptions,
	callOpts ...privyclient.RequestOption,
) (headerValue string, beforeMs int64) {
	t.Helper()

	req, beforeMs := captureIntentRpcRequest(t, clientOpts, callOpts...)
	return req.Header.Get("Privy-Request-Expiry"), beforeMs
}

// captureIntentRpcRequest runs client.Intents.Rpc with the provided call
// options and returns the full outgoing *http.Request, along with the
// timestamp (unix ms) recorded just before the call. The HTTP transport
// short-circuits with a 200 response so no real network traffic leaves the
// process.
func captureIntentRpcRequest(
	t *testing.T,
	clientOpts privyclient.PrivyClientOptions,
	callOpts ...privyclient.RequestOption,
) (req *http.Request, beforeMs int64) {
	t.Helper()

	var capturedReq *http.Request
	customClient := &http.Client{
		Transport: &closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				capturedReq = req
				return &http.Response{StatusCode: http.StatusOK}, nil
			},
		},
	}

	clientOpts.AppID = "test-app-id"
	clientOpts.AppSecret = "test-app-secret"
	clientOpts.HTTPClient = customClient

	client := privyclient.NewPrivyClient(clientOpts)

	beforeMs = time.Now().UnixMilli()
	_, _ = client.Intents.Rpc(
		context.Background(),
		"wallet-id",
		privyclient.IntentRpcParams{},
		callOpts...,
	)

	if capturedReq == nil {
		t.Fatal("expected request to be captured")
	}
	return capturedReq, beforeMs
}

func assertExpiryWithin(t *testing.T, headerValue string, beforeMs, expectedOffsetMs int64) {
	t.Helper()
	if headerValue == "" {
		t.Fatal("expected privy-request-expiry header to be set")
	}
	gotMs, err := strconv.ParseInt(headerValue, 10, 64)
	if err != nil {
		t.Fatalf("expected header to be an integer, got %q: %v", headerValue, err)
	}

	afterMs := time.Now().UnixMilli()
	earliest := beforeMs + expectedOffsetMs
	latest := afterMs + expectedOffsetMs + intentTestToleranceMs

	if gotMs < earliest-intentTestToleranceMs || gotMs > latest {
		t.Errorf(
			"expected expiry header ≈ now + %d ms (in [%d, %d]), got %d",
			expectedOffsetMs, earliest-intentTestToleranceMs, latest, gotMs,
		)
	}
}

func TestIntentRequestExpiryDefaults72h(t *testing.T) {
	header, before := captureIntentExpiryHeader(t, privyclient.PrivyClientOptions{})
	assertExpiryWithin(t, header, before, 72*60*60*1000)
}

func TestIntentRequestExpiryUsesDefaultIntentOption(t *testing.T) {
	const customMs = 6 * 60 * 60 * 1000 // 6 hours
	header, before := captureIntentExpiryHeader(t, privyclient.PrivyClientOptions{
		DefaultIntentRequestExpiryMs: customMs,
	})
	assertExpiryWithin(t, header, before, customMs)
}

func TestIntentRequestExpiryPerCallOverrides(t *testing.T) {
	// Set an explicit absolute expiry timestamp via WithRequestExpiry. The
	// per-call value should be written into the header verbatim, overriding
	// both DefaultIntentRequestExpiryMs and the 72h fallback.
	const perCallOffsetMs int64 = 45 * 60 * 1000 // 45 minutes
	expected := time.Now().UnixMilli() + perCallOffsetMs

	header, _ := captureIntentExpiryHeader(t,
		privyclient.PrivyClientOptions{
			DefaultIntentRequestExpiryMs: 6 * 60 * 60 * 1000,
		},
		privyclient.WithRequestExpiry(expected),
	)

	if header == "" {
		t.Fatal("expected privy-request-expiry header to be set")
	}
	gotMs, err := strconv.ParseInt(header, 10, 64)
	if err != nil {
		t.Fatalf("expected header to be an integer, got %q: %v", header, err)
	}
	if gotMs != expected {
		t.Errorf("expected per-call expiry %d to be written verbatim, got %d", expected, gotMs)
	}
}

func TestIntentRequestExpiryIndependentFromDefaultRequestExpiryMs(t *testing.T) {
	// Setting only DefaultRequestExpiryMs must NOT influence intent calls:
	// they should still resolve to the 72h fallback.
	header, before := captureIntentExpiryHeader(t, privyclient.PrivyClientOptions{
		DefaultRequestExpiryMs: 30 * 60 * 1000, // 30 minutes — must be ignored for intents
	})
	assertExpiryWithin(t, header, before, 72*60*60*1000)
}

func TestIntentRequestExpiryHonorsDisableFlag(t *testing.T) {
	req, _ := captureIntentRpcRequest(t, privyclient.PrivyClientOptions{
		DisableRequestExpiry: true,
	})
	if got := req.Header.Get("Privy-Request-Expiry"); got != "" {
		t.Errorf("expected no privy-request-expiry header when DisableRequestExpiry=true, got %q", got)
	}
}
