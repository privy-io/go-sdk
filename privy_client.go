package privyclient

import (
	"fmt"
	"log"
	"os"

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
	// Use "https://auth.staging.privy.io" for staging.
	APIUrl string

	// LogLevel sets the verbosity of SDK logging (optional).
	// If not provided, defaults to LogLevelNone (no logging).
	// Available levels: LogLevelNone, LogLevelError, LogLevelInfo, LogLevelDebug, LogLevelVerbose
	LogLevel LogLevel
}

// PrivyClient is the main entrypoint for the Privy API Go SDK.
//
// Example:
//
//	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
//	    AppID:     "my-app-id",
//	    AppSecret: "my-app-secret",
//	    APIUrl:    "https://auth.staging.privy.io", // optional
//	})
type PrivyClient struct {
	client Client
	logger logger

	Wallets      *PrivyWalletService
	Users        *PrivyUserService
	Policies     *PrivyPolicyService
	Transactions *PrivyTransactionService
	KeyQuorums   *PrivyKeyQuorumService
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
//	    APIUrl:    "https://auth.staging.privy.io",
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

	// Compute base URL
	baseURL := opts.APIUrl
	if baseURL == "" {
		baseURL = "https://api.privy.io"
	}

	client := NewClient(requestOpts...)

	// Create JWT exchange service (uses generated WalletService for AuthenticateWithJwt)
	jwtExchange := newPrivyJwtExchangeService(&client.Wallets, logger)

	// Create wallet service with jwtExchanger for authorization
	wallets := newPrivyWalletService(client.Wallets, jwtExchange, baseURL, opts.AppID, logger)

	return &PrivyClient{
		client:       client,
		logger:       logger,
		Wallets:      wallets,
		Users:        newPrivyUserService(client.Users, logger),
		Policies:     newPrivyPolicyService(client.Policies, jwtExchange, baseURL, opts.AppID, logger),
		Transactions: newPrivyTransactionService(client.Transactions, logger),
		KeyQuorums:   newPrivyKeyQuorumService(client.KeyQuorums, jwtExchange, baseURL, opts.AppID, logger),
		Analytics:    newPrivyAnalyticsService(client.Analytics, logger),
		Apps:         newPrivyAppService(client.Apps, logger),
		Aggregations: newPrivyAggregationService(client.Aggregations, logger),
		Webhooks:     newPrivyWebhookService(client.Webhooks, logger),
		JwtExchange:  jwtExchange,
	}
}
