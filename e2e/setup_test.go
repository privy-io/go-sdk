//go:build e2e

package e2e_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/option"
)

func TestMain(m *testing.M) {
	// Load .env file from project root
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../.env.secret")
	os.Exit(m.Run())
}

func newTestClient(t *testing.T) *Client {
	t.Helper()
	appId := os.Getenv("TEST_APP_ID")
	appSecret := os.Getenv("TEST_APP_SECRET")
	if appSecret == "" {
		t.Fatal("TEST_APP_SECRET environment variable is required")
	}
	client := NewClient(
		option.WithAppID(appId),
		option.WithAppSecret(appSecret),
		option.WithEnvironmentStaging(),
		option.WithHeader("privy-app-id", appId),
	)
	return &client
}
