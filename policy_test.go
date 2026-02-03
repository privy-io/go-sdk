// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/privy-api-client-go"
	"github.com/stainless-sdks/privy-api-client-go/internal/testutil"
	"github.com/stainless-sdks/privy-api-client-go/option"
)

func TestPolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.New(context.TODO(), privyclient.PolicyNewParams{
		ChainType: privyclient.PolicyNewParamsChainTypeEthereum,
		Name:      "x",
		Rules: []privyclient.PolicyNewParamsRule{{
			Action: "ALLOW",
			Conditions: []privyclient.PolicyNewParamsRuleConditionUnion{{
				OfEthereumTransaction: &privyclient.PolicyNewParamsRuleConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyclient.PolicyNewParamsRuleConditionEthereumTransactionValueUnion{
						OfString: privyclient.String("string"),
					},
				},
			}},
			Method: "eth_sendTransaction",
			Name:   "x",
		}},
		Version: privyclient.PolicyNewParamsVersion1_0,
		Owner: privyclient.PolicyNewParamsOwnerUnion{
			OfPublicKeyOwner: &privyclient.PolicyNewParamsOwnerPublicKeyOwner{
				PublicKey: "public_key",
			},
		},
		OwnerID:             privyclient.String("owner_id"),
		PrivyIdempotencyKey: privyclient.String("privy-idempotency-key"),
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.Update(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyUpdateParams{
			Name: privyclient.String("x"),
			Owner: privyclient.PolicyUpdateParamsOwnerUnion{
				OfPublicKeyOwner: &privyclient.PolicyUpdateParamsOwnerPublicKeyOwner{
					PublicKey: "public_key",
				},
			},
			OwnerID: privyclient.String("owner_id"),
			Rules: []privyclient.PolicyUpdateParamsRule{{
				Action: "ALLOW",
				Conditions: []privyclient.PolicyUpdateParamsRuleConditionUnion{{
					OfEthereumTransaction: &privyclient.PolicyUpdateParamsRuleConditionEthereumTransaction{
						Field:    "to",
						Operator: "eq",
						Value: privyclient.PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion{
							OfString: privyclient.String("string"),
						},
					},
				}},
				Method: "eth_sendTransaction",
				Name:   "x",
			}},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.Delete(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyDeleteParams{
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyNewRuleWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.NewRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyNewRuleParams{
			Action: privyclient.PolicyNewRuleParamsActionAllow,
			Conditions: []privyclient.PolicyNewRuleParamsConditionUnion{{
				OfEthereumTransaction: &privyclient.PolicyNewRuleParamsConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyclient.PolicyNewRuleParamsConditionEthereumTransactionValueUnion{
						OfString: privyclient.String("string"),
					},
				},
			}},
			Method:                      privyclient.PolicyNewRuleParamsMethodEthSendTransaction,
			Name:                        "x",
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyDeleteRuleWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.DeleteRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyDeleteRuleParams{
			PolicyID:                    "xxxxxxxxxxxxxxxxxxxxxxxx",
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.Get(context.TODO(), "xxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyGetRule(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.GetRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyGetRuleParams{
			PolicyID: "xxxxxxxxxxxxxxxxxxxxxxxx",
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

func TestPolicyUpdateRuleWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Policies.UpdateRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyUpdateRuleParams{
			PolicyID: "xxxxxxxxxxxxxxxxxxxxxxxx",
			Action:   privyclient.PolicyUpdateRuleParamsActionAllow,
			Conditions: []privyclient.PolicyUpdateRuleParamsConditionUnion{{
				OfEthereumTransaction: &privyclient.PolicyUpdateRuleParamsConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyclient.PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion{
						OfString: privyclient.String("string"),
					},
				},
			}},
			Method:                      privyclient.PolicyUpdateRuleParamsMethodEthSendTransaction,
			Name:                        "x",
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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
