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

func TestOrganizationExternalFiatAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.ExternalFiatAccounts.New(
		context.TODO(),
		"organization_id",
		privyclient.OrganizationExternalFiatAccountNewParams{
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

func TestOrganizationExternalFiatAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.ExternalFiatAccounts.List(
		context.TODO(),
		"organization_id",
		privyclient.OrganizationExternalFiatAccountListParams{
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

func TestOrganizationExternalFiatAccountDelete(t *testing.T) {
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
	_, err := client.Organizations.ExternalFiatAccounts.Delete(
		context.TODO(),
		"account_id",
		privyclient.OrganizationExternalFiatAccountDeleteParams{
			OrganizationID: "organization_id",
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

func TestOrganizationExternalFiatAccountGet(t *testing.T) {
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
	_, err := client.Organizations.ExternalFiatAccounts.Get(
		context.TODO(),
		"account_id",
		privyclient.OrganizationExternalFiatAccountGetParams{
			OrganizationID: "organization_id",
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
