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

func TestWalletEarnEthereumDepositWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Earn.Ethereum.Deposit(
		context.TODO(),
		"wallet_id",
		privyclient.WalletEarnEthereumDepositParams{
			EarnDepositRequestBody: privyclient.EarnDepositRequestBody{
				VaultID:   "cm7oxq1el000e11o8iwp7d0d0",
				Amount:    privyclient.String("1.5"),
				RawAmount: privyclient.String("321669910225"),
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

func TestWalletEarnEthereumWithdrawWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Earn.Ethereum.Withdraw(
		context.TODO(),
		"wallet_id",
		privyclient.WalletEarnEthereumWithdrawParams{
			EarnWithdrawRequestBody: privyclient.EarnWithdrawRequestBody{
				VaultID:   "cm7oxq1el000e11o8iwp7d0d0",
				Amount:    privyclient.String("1.5"),
				RawAmount: privyclient.String("321669910225"),
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
