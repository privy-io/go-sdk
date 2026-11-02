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

func TestUserMfaPasskeyVerificationInit(t *testing.T) {
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
	_, err := client.Users.Mfa.Passkey.Verification.Init(
		context.TODO(),
		"user_id",
		privyclient.UserMfaPasskeyVerificationInitParams{
			PasskeyMfaInitRequestBody: privyclient.PasskeyMfaInitRequestBody{
				RelyingParty: "https://example.com",
			},
			UserJwt: "x",
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

func TestUserMfaPasskeyVerificationVerify(t *testing.T) {
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
	_, err := client.Users.Mfa.Passkey.Verification.Verify(
		context.TODO(),
		"user_id",
		privyclient.UserMfaPasskeyVerificationVerifyParams{
			PasskeyMfaVerificationVerifyRequestBody: privyclient.PasskeyMfaVerificationVerifyRequestBody{
				Credential: privyclient.MfaPasskeyAuthenticatorVerifyResponse{
					ID: "id",
					ClientExtensionResults: privyclient.MfaPasskeyClientExtensionResults{
						AppID: privyclient.Bool(true),
						CredProps: privyclient.MfaPasskeyCredPropsResult{
							Rk: privyclient.Bool(true),
						},
						HmacCreateSecret: privyclient.Bool(true),
					},
					RawID: "raw_id",
					Response: privyclient.MfaPasskeyAssertionResponse{
						AuthenticatorData: "authenticator_data",
						ClientDataJson:    "client_data_json",
						Signature:         "signature",
						UserHandle:        privyclient.String("user_handle"),
					},
					Type:                    privyclient.MfaPasskeyAuthenticatorVerifyResponseTypePublicKey,
					AuthenticatorAttachment: privyclient.String("authenticator_attachment"),
				},
			},
			UserJwt: "x",
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
