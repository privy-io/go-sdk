//go:build e2e

package e2e_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/option"
)

const testAppID = "cm8osegv00004r13y7500o2yz"

func TestMain(m *testing.M) {
	// Load .env file from project root
	_ = godotenv.Load("../.env")
	os.Exit(m.Run())
}

func newTestClient(t *testing.T) *Client {
	t.Helper()
	appSecret := os.Getenv("TEST_APP_SECRET")
	if appSecret == "" {
		t.Fatal("TEST_APP_SECRET environment variable is required")
	}
	client := NewClient(
		option.WithAppID(testAppID),
		option.WithAppSecret(appSecret),
		option.WithEnvironmentStaging(),
		option.WithHeader("privy-app-id", testAppID),
	)
	return &client
}
