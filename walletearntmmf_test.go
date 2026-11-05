// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/internal/testutil"
	"github.com/privy-io/go-sdk/option"
)

func TestWalletEarnTmmfDepositWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Earn.Tmmf.Deposit(
		context.TODO(),
		"wallet_id",
		privyclient.WalletEarnTmmfDepositParams{
			TmmfDepositRequestBody: privyclient.TmmfDepositRequestBody{
				Amount:       "1000.00",
				Chain:        "ethereum_sepolia",
				DepositAsset: "USDC",
				Fund:         "WTGXX",
				OrderType:    privyclient.TmmfOrderTypeSecondary,
				Nonce:        privyclient.String("xxxxxxxxxxxxxxxxxxxxxxxx"),
				ReferenceID:  privyclient.String("x"),
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
			PrivyRequestExpiry:          privyclient.String("privy-request-expiry"),
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletEarnTmmfListPositionsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Earn.Tmmf.ListPositions(
		context.TODO(),
		"wallet_id",
		privyclient.WalletEarnTmmfListPositionsParams{
			Chain:    privyclient.String("ethereum_sepolia"),
			Cursor:   privyclient.String("x"),
			Fund:     privyclient.String("WTGXX"),
			Limit:    privyclient.Int(1),
			Provider: privyclient.TmmfProviderWisdomtree,
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletEarnTmmfRegister(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Earn.Tmmf.Register(
		context.TODO(),
		"wallet_id",
		privyclient.WalletEarnTmmfRegisterParams{
			TmmfRegisterRequestBody: privyclient.TmmfRegisterRequestBody{
				Chain:    "ethereum_sepolia",
				Provider: privyclient.TmmfProviderWisdomtree,
			},
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletEarnTmmfWithdrawWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Earn.Tmmf.Withdraw(
		context.TODO(),
		"wallet_id",
		privyclient.WalletEarnTmmfWithdrawParams{
			TmmfWithdrawRequestBody: privyclient.TmmfWithdrawRequestBody{
				Chain:         "ethereum_sepolia",
				Fund:          "WTGXX",
				OrderType:     privyclient.TmmfOrderTypePrimary,
				Shares:        "1000.000000",
				WithdrawAsset: "USDC",
				Nonce:         privyclient.String("xxxxxxxxxxxxxxxxxxxxxxxx"),
				ReferenceID:   privyclient.String("x"),
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
			PrivyRequestExpiry:          privyclient.String("privy-request-expiry"),
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
