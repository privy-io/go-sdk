//go:build e2e

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
	client := NewPrivyClient(PrivyClientOptions{
		AppID:     appId,
		AppSecret: appSecret,
		APIUrl:    "https://auth.staging.privy.io",
		LogLevel:  LogLevelVerbose,
	})
	return client
}
