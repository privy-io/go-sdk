package e2e_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/privy-io/go-sdk"
)

func TestMain(m *testing.M) {
	// Load .env file from project root
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../.env.secret")
	os.Exit(m.Run())
}

func newTestClient(t *testing.T) *PrivyClient {
	t.Helper()
	appId := os.Getenv("TEST_APP_ID")
	appSecret := os.Getenv("TEST_APP_SECRET")
	if appSecret == "" {
		t.Fatal("TEST_APP_SECRET environment variable is required")
	}
	opts := PrivyClientOptions{
		AppID:     appId,
		AppSecret: appSecret,
		APIUrl:    "https://auth.staging.privy.io",
	}
	if os.Getenv("TEST_DEBUG_LOGS") != "" {
		opts.LogLevel = LogLevelVerbose
	}
	client := NewPrivyClient(opts)
	return client
}
