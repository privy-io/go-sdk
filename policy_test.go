// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/internal/testutil"
	"github.com/privy-io/go-sdk/option"
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
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Policies.New(context.TODO(), privyapiclient.PolicyNewParams{
		ChainType: privyapiclient.PolicyNewParamsChainTypeEthereum,
		Name:      "name",
		Rules: []privyapiclient.PolicyNewParamsRule{{
			Action: "ALLOW",
			Conditions: []privyapiclient.PolicyNewParamsRuleConditionUnion{{
				OfEthereumTransaction: &privyapiclient.PolicyNewParamsRuleConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyapiclient.PolicyNewParamsRuleConditionEthereumTransactionValueUnion{
						OfString: privyapiclient.String("string"),
					},
				},
			}},
			Method: "eth_sendTransaction",
			Name:   "name",
		}},
		Version: privyapiclient.PolicyNewParamsVersion1_0,
		Owner: privyapiclient.PolicyNewParamsOwnerUnion{
			OfPublicKeyOwner: &privyapiclient.PolicyNewParamsOwnerPublicKeyOwner{
				PublicKey: "public_key",
			},
		},
		OwnerID:             privyapiclient.String("owner_id"),
		PrivyIdempotencyKey: privyapiclient.String("privy-idempotency-key"),
	})
	if err != nil {
		var apierr *privyapiclient.Error
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
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Policies.Update(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyapiclient.PolicyUpdateParams{
			Name: privyapiclient.String("x"),
			Owner: privyapiclient.PolicyUpdateParamsOwnerUnion{
				OfPublicKeyOwner: &privyapiclient.PolicyUpdateParamsOwnerPublicKeyOwner{
					PublicKey: "public_key",
				},
			},
			OwnerID: privyapiclient.String("owner_id"),
			Rules: []privyapiclient.PolicyUpdateParamsRule{{
				Action: "ALLOW",
				Conditions: []privyapiclient.PolicyUpdateParamsRuleConditionUnion{{
					OfEthereumTransaction: &privyapiclient.PolicyUpdateParamsRuleConditionEthereumTransaction{
						Field:    "to",
						Operator: "eq",
						Value: privyapiclient.PolicyUpdateParamsRuleConditionEthereumTransactionValueUnion{
							OfString: privyapiclient.String("string"),
						},
					},
				}},
				Method: "eth_sendTransaction",
				Name:   "name",
			}},
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

func TestPolicyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.Delete(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyapiclient.PolicyDeleteParams{
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

func TestPolicyNewRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.NewRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyapiclient.PolicyNewRuleParams{
			Action: privyapiclient.PolicyNewRuleParamsActionAllow,
			Conditions: []privyapiclient.PolicyNewRuleParamsConditionUnion{{
				OfEthereumTransaction: &privyapiclient.PolicyNewRuleParamsConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyapiclient.PolicyNewRuleParamsConditionEthereumTransactionValueUnion{
						OfString: privyapiclient.String("string"),
					},
				},
			}},
			Method:                      privyapiclient.PolicyNewRuleParamsMethodEthSendTransaction,
			Name:                        "name",
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

func TestPolicyDeleteRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.DeleteRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyapiclient.PolicyDeleteRuleParams{
			PolicyID:                    "xxxxxxxxxxxxxxxxxxxxxxxx",
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

func TestPolicyGet(t *testing.T) {
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
	_, err := client.Policies.Get(context.TODO(), "xxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		var apierr *privyapiclient.Error
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
	client := privyapiclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Policies.GetRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyapiclient.PolicyGetRuleParams{
			PolicyID: "xxxxxxxxxxxxxxxxxxxxxxxx",
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

func TestPolicyUpdateRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.UpdateRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyapiclient.PolicyUpdateRuleParams{
			PolicyID: "xxxxxxxxxxxxxxxxxxxxxxxx",
			Action:   privyapiclient.PolicyUpdateRuleParamsActionAllow,
			Conditions: []privyapiclient.PolicyUpdateRuleParamsConditionUnion{{
				OfEthereumTransaction: &privyapiclient.PolicyUpdateRuleParamsConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyapiclient.PolicyUpdateRuleParamsConditionEthereumTransactionValueUnion{
						OfString: privyapiclient.String("string"),
					},
				},
			}},
			Method:                      privyapiclient.PolicyUpdateRuleParamsMethodEthSendTransaction,
			Name:                        "name",
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
