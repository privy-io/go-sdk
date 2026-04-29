package privyclient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/internal"
	"github.com/privy-io/go-sdk/option"
)

// PrivyClientOptions contains configuration options for creating a PrivyClient.
type PrivyClientOptions struct {
	// AppID is your Privy application ID (required).
	AppID string

	// AppSecret is your Privy application secret (required).
	AppSecret string

	// APIUrl is the base URL for the Privy API (optional).
	// If not provided, defaults to the production environment.
	// Use "https://api.staging.privy.io" for staging.
	APIUrl string

	// LogLevel sets the verbosity of SDK logging (optional).
	// If not provided, defaults to LogLevelNone (no logging).
	// Available levels: LogLevelNone, LogLevelError, LogLevelInfo, LogLevelDebug, LogLevelVerbose
	LogLevel LogLevel

	// DefaultRequestExpiryMs sets the default request expiry duration in milliseconds (optional).
	// This is used as the offset from the current time to compute the "privy-request-expiry" header.
	// If not provided, defaults to 15 minutes (900000 ms).
	// Can be overridden per-request, where applicable, using WithRequestExpiry.
	DefaultRequestExpiryMs int64

	// HTTPClient sets the default *http.Client used across all requests (optional).
	// If not provided, defaults to http.DefaultClient.
	// Can be overridden per-request using WithHTTPClient.
	HTTPClient *http.Client

	// WebhookSigningSecret is used to verify incoming webhook signatures (optional).
	// Can be overridden per-call via VerifyInput.SigningSecret.
	WebhookSigningSecret string
}

// PrivyClient is the main entrypoint for the Privy API Go SDK.
//
// Example:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	    APIUrl:    "https://api.staging.privy.io", // optional
//	})
type PrivyClient struct {
	client Client
	logger logger

	Wallets      *PrivyWalletService
	Users        *PrivyUserService
	Policies     *PrivyPolicyService
	Transactions *PrivyTransactionService
	KeyQuorums   *PrivyKeyQuorumService
	Intents      *PrivyIntentService
	Analytics    *PrivyAnalyticsService
	Apps         *PrivyAppService
	Aggregations *PrivyAggregationService
	Webhooks     *PrivyWebhookService
	JwtExchange  *PrivyJwtExchangeService
}

// NewPrivyClient creates a new enhanced Privy client.
// This is the recommended way to create a client for most use cases.
//
// Example:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	})
//
// For staging environment:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	    APIUrl:    "https://api.staging.privy.io",
//	})
//
// With logging enabled:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	    LogLevel:  privyclient.LogLevelDebug,
//	})
func NewPrivyClient(opts PrivyClientOptions) *PrivyClient {
	// Create logger internally based on log level
	logger := newPrivyLogger(opts.LogLevel)

	// Build option.RequestOption slice from PrivyClientOptions
	requestOpts := []option.RequestOption{
		option.WithAppID(opts.AppID),
		option.WithAppSecret(opts.AppSecret),
		option.WithHeader("privy-app-id", opts.AppID),
		option.WithHeader("privy-client", fmt.Sprintf("go:%s", internal.PackageVersion)),
	}

	// Add API URL if provided, otherwise use production default
	if opts.APIUrl != "" {
		requestOpts = append(requestOpts, option.WithBaseURL(opts.APIUrl))
	} else {
		requestOpts = append(requestOpts, option.WithEnvironmentProduction())
	}

	// Enable HTTP debug logging for Debug and Verbose log levels
	if opts.LogLevel >= LogLevelDebug {
		debugLogger := log.New(os.Stdout, "[Privy][HTTP] ", 0)
		requestOpts = append(requestOpts, option.WithDebugLog(debugLogger))
	}

	if opts.HTTPClient != nil {
		requestOpts = append(requestOpts, option.WithHTTPClient(opts.HTTPClient))
	}

	// Compute base URL
	baseURL := opts.APIUrl
	if baseURL == "" {
		baseURL = "https://api.privy.io"
	}

	// Resolve default request expiry (fallback to 15 minutes)
	defaultRequestExpiryMs := opts.DefaultRequestExpiryMs
	if defaultRequestExpiryMs == 0 {
		defaultRequestExpiryMs = 15 * 60 * 1000
	}

	client := NewClient(requestOpts...)

	// Create JWT exchange service (uses generated WalletService for AuthenticateWithJwt)
	jwtExchange := newPrivyJwtExchangeService(&client.Wallets, logger)

	// Create wallet service with jwtExchanger for authorization
	wallets := newPrivyWalletService(client.Wallets, jwtExchange, baseURL, opts.AppID, defaultRequestExpiryMs, logger)

	return &PrivyClient{
		client:       client,
		logger:       logger,
		Wallets:      wallets,
		Users:        newPrivyUserService(client.Users, logger),
		Policies:     newPrivyPolicyService(client.Policies, jwtExchange, baseURL, opts.AppID, defaultRequestExpiryMs, logger),
		Transactions: newPrivyTransactionService(client.Transactions, logger),
		KeyQuorums:   newPrivyKeyQuorumService(client.KeyQuorums, jwtExchange, baseURL, opts.AppID, defaultRequestExpiryMs, logger),
		Intents:      newPrivyIntentService(client.Intents, logger),
		Analytics:    newPrivyAnalyticsService(client.Analytics, logger),
		Apps:         newPrivyAppService(client.Apps, logger),
		Aggregations: newPrivyAggregationService(client.Aggregations, logger),
		Webhooks:     newPrivyWebhookService(client.Webhooks, opts.WebhookSigningSecret, logger),
		JwtExchange:  jwtExchange,
	}
}

// GenerateAuthorizationSignaturesForRequest formats a request and generates
// signatures for all credentials in an AuthorizationContext, using the client's
// built-in JWT exchanger for any JWTs in the authorization context.
//
// This is a convenience method that delegates to
// [authorization.GenerateAuthorizationSignaturesForRequest] with the client's
// JWT exchange service, so callers don't need to pass the exchanger explicitly.
//
// Example:
//
//	signatures, err := client.GenerateAuthorizationSignaturesForRequest(ctx,
//	    authorization.AuthorizationContext{
//	        UserJwts: []string{userJWT},
//	    },
//	    authorization.WalletApiRequestSignatureInput{
//	        Version: 1,
//	        Method:  "POST",
//	        URL:     "https://api.privy.io/v1/wallets/my-wallet/rpc",
//	        Body:    requestBody,
//	        Headers: map[string]string{"privy-app-id": "my-app-id"},
//	    },
//	)
func (c *PrivyClient) GenerateAuthorizationSignaturesForRequest(
	ctx context.Context,
	auth authorization.AuthorizationContext,
	input authorization.WalletApiRequestSignatureInput,
) ([]string, error) {
	return authorization.GenerateAuthorizationSignaturesForRequest(ctx, auth, input, c.JwtExchange)
}
