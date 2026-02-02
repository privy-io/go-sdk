// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/privy-api-client-go"
	"github.com/stainless-sdks/privy-api-client-go/internal/testutil"
	"github.com/stainless-sdks/privy-api-client-go/option"
)

func TestKeyQuorumNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.KeyQuorums.New(context.TODO(), privyapiclient.KeyQuorumNewParams{
		AuthorizationThreshold: privyapiclient.Float(1),
		DisplayName:            privyapiclient.String("Prod key quorum"),
		PublicKeys:             []string{"-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEx4aoeD72yykviK+f/ckqE2CItVIG\n1rCnvC3/XZ1HgpOcMEMialRmTrqIK4oZlYd1RfxU3za/C9yjhboIuoPD3g==\n-----END PUBLIC KEY-----", "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErzZtQr/bMIh3Y8f9ZqseB9i/AfjQ\nhu+agbNqXcJy/TfoNqvc/Y3Mh7gIZ8ZLXQEykycx4mYSpqrxp1lBKqsZDQ==\n-----END PUBLIC KEY-----\","},
		UserIDs:                []string{"string"},
	})
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyQuorumUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.KeyQuorums.Update(
		context.TODO(),
		"key_quorum_id",
		privyapiclient.KeyQuorumUpdateParams{
			AuthorizationThreshold:      privyapiclient.Float(1),
			DisplayName:                 privyapiclient.String("Prod key quorum"),
			PublicKeys:                  []string{"-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEx4aoeD72yykviK+f/ckqE2CItVIG\n1rCnvC3/XZ1HgpOcMEMialRmTrqIK4oZlYd1RfxU3za/C9yjhboIuoPD3g==\n-----END PUBLIC KEY-----"},
			UserIDs:                     []string{"string"},
			PrivyAuthorizationSignature: privyapiclient.String("privy-authorization-signature"),
		},
	)
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyQuorumDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.KeyQuorums.Delete(
		context.TODO(),
		"key_quorum_id",
		privyapiclient.KeyQuorumDeleteParams{
			PrivyAuthorizationSignature: privyapiclient.String("privy-authorization-signature"),
		},
	)
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyQuorumGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.KeyQuorums.Get(context.TODO(), "key_quorum_id")
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
