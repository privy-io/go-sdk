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

func TestUserKYCList(t *testing.T) {
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
	_, err := client.Users.KYC.List(context.TODO(), "user_id")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserKYCInitiateLinksWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.KYC.InitiateLinks(
		context.TODO(),
		"user_id",
		privyclient.UserKYCInitiateLinksParams{
			KYCLinksRequestBody: privyclient.KYCLinksRequestBody{
				Provider:          privyclient.KyxProviderBridge,
				ClientAgreementID: privyclient.String("x"),
				Email:             privyclient.String("dev@stainless.com"),
				Endorsements:      []privyclient.KyxEndorsementName{"sepa"},
				Environment:       privyclient.KyxEnvironmentProduction,
				RedirectUri:       privyclient.String("https://example.com"),
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

func TestUserKYCInitiateTosWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.KYC.InitiateTos(
		context.TODO(),
		"user_id",
		privyclient.UserKYCInitiateTosParams{
			KyxTosRequestBody: privyclient.KyxTosRequestBody{
				Provider:    privyclient.KyxProviderBridge,
				Email:       privyclient.String("dev@stainless.com"),
				Environment: privyclient.KyxEnvironmentProduction,
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

func TestUserKYCSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.KYC.Submit(
		context.TODO(),
		"user_id",
		privyclient.UserKYCSubmitParams{
			KYCSubmitRequestBody: privyclient.KYCSubmitRequestBody{
				Data: privyclient.KYCSubmitData{
					DateOfBirth: privyclient.String("7321-69-10"),
					Email:       privyclient.String("dev@stainless.com"),
					FirstName:   privyclient.String("x"),
					IdentifyingInformation: []privyclient.VerificationDocument{{
						IssuingCountry: "xxx",
						Type:           "type",
						Description:    privyclient.String("description"),
						Expiration:     privyclient.String("expiration"),
						ImageBack:      privyclient.String("image_back"),
						ImageFront:     privyclient.String("image_front"),
						Number:         privyclient.String("number"),
					}},
					LastName: privyclient.String("x"),
					Phone:    privyclient.String("phone"),
					ResidentialAddress: privyclient.VerificationAddress{
						City:        "x",
						Country:     "xxx",
						StreetLine1: "xxxx",
						PostalCode:  privyclient.String("x"),
						StreetLine2: privyclient.String("x"),
						Subdivision: privyclient.String("x"),
					},
				},
				Provider:          privyclient.KyxProviderBridge,
				ClientAgreementID: privyclient.String("x"),
				Endorsements:      []privyclient.KyxEndorsementName{"sepa"},
				Environment:       privyclient.KyxEnvironmentProduction,
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
