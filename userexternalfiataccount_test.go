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

func TestUserExternalFiatAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.ExternalFiatAccounts.New(
		context.TODO(),
		"user_id",
		privyclient.UserExternalFiatAccountNewParams{
			CreateExternalFiatAccountRequestBody: privyclient.CreateExternalFiatAccountRequestBody{
				Account: privyclient.ExternalFiatAccountDataUnion{
					OfUs: &privyclient.ExternalFiatAccountUsData{
						AccountNumber:     "x",
						RoutingNumber:     "xxxxxxxxx",
						Type:              privyclient.ExternalFiatAccountUsDataTypeUs,
						CheckingOrSavings: privyclient.String("checking_or_savings"),
					},
				},
				AccountOwnerName: "xxx",
				Currency:         "currency",
				Provider:         privyclient.CreateExternalFiatAccountRequestBodyProviderBridge,
				Address: privyclient.ExternalFiatAccountAddress{
					City:        "x",
					Country:     "xxx",
					StreetLine1: "x",
					PostalCode:  privyclient.String("x"),
					State:       privyclient.String("x"),
					StreetLine2: privyclient.String("x"),
				},
				BankName:    privyclient.String("x"),
				Environment: privyclient.EnvironmentSandbox,
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

func TestUserExternalFiatAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.ExternalFiatAccounts.List(
		context.TODO(),
		"user_id",
		privyclient.UserExternalFiatAccountListParams{
			Provider:    privyclient.OrchestrationProviderBridge,
			Environment: privyclient.EnvironmentSandbox,
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

func TestUserExternalFiatAccountDelete(t *testing.T) {
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
	_, err := client.Users.ExternalFiatAccounts.Delete(
		context.TODO(),
		"account_id",
		privyclient.UserExternalFiatAccountDeleteParams{
			UserID: "user_id",
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

func TestUserExternalFiatAccountGet(t *testing.T) {
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
	_, err := client.Users.ExternalFiatAccounts.Get(
		context.TODO(),
		"account_id",
		privyclient.UserExternalFiatAccountGetParams{
			UserID: "user_id",
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
