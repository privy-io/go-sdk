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

func TestPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.New(context.TODO(), privyclient.PolicyNewParams{
		ChainType: privyclient.WalletChainTypeEthereum,
		Name:      "x",
		Rules: []privyclient.PolicyNewParamsRule{{
			Action: privyclient.PolicyActionAllow,
			Conditions: []privyclient.PolicyConditionUnion{{
				OfEthereumTransaction: &privyclient.EthereumTransactionCondition{
					Field:       privyclient.EthereumTransactionConditionFieldTo,
					FieldSource: privyclient.EthereumTransactionConditionFieldSourceEthereumTransaction,
					Operator:    privyclient.ConditionOperatorEq,
					Value: privyclient.ConditionValueUnion{
						OfString: privyclient.String("string"),
					},
				},
			}},
			Method: privyclient.PolicyMethodEthSendTransaction,
			Name:   "x",
			ID:     privyclient.String("id"),
		}},
		Version: privyclient.PolicyNewParamsVersion1_0,
		Owner: privyclient.OwnerInputUnion{
			OfOwnerInputUser: &privyclient.OwnerInputUser{
				UserID: "user_id",
			},
		},
		OwnerID:             privyclient.String("string"),
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
	_, err := client.Policies.Update(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyUpdateParams{
			Name: privyclient.String("x"),
			Owner: privyclient.OwnerInputUnion{
				OfOwnerInputUser: &privyclient.OwnerInputUser{
					UserID: "user_id",
				},
			},
			OwnerID: privyclient.String("string"),
			Rules: []privyclient.PolicyRuleRequestBody{{
				Action: privyclient.PolicyActionAllow,
				Conditions: []privyclient.PolicyConditionUnion{{
					OfEthereumTransaction: &privyclient.EthereumTransactionCondition{
						Field:       privyclient.EthereumTransactionConditionFieldTo,
						FieldSource: privyclient.EthereumTransactionConditionFieldSourceEthereumTransaction,
						Operator:    privyclient.ConditionOperatorEq,
						Value: privyclient.ConditionValueUnion{
							OfString: privyclient.String("string"),
						},
					},
				}},
				Method: privyclient.PolicyMethodEthSendTransaction,
				Name:   "x",
			}},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.Delete(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyDeleteParams{
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyNewRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.NewRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyNewRuleParams{
			PolicyRuleRequestBody: privyclient.PolicyRuleRequestBody{
				Action: privyclient.PolicyActionAllow,
				Conditions: []privyclient.PolicyConditionUnion{{
					OfEthereumTransaction: &privyclient.EthereumTransactionCondition{
						Field:       privyclient.EthereumTransactionConditionFieldTo,
						FieldSource: privyclient.EthereumTransactionConditionFieldSourceEthereumTransaction,
						Operator:    privyclient.ConditionOperatorEq,
						Value: privyclient.ConditionValueUnion{
							OfString: privyclient.String("string"),
						},
					},
				}},
				Method: privyclient.PolicyMethodEthSendTransaction,
				Name:   "x",
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyDeleteRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Policies.DeleteRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyDeleteRuleParams{
			PolicyID:                    "xxxxxxxxxxxxxxxxxxxxxxxx",
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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

func TestPolicyGet(t *testing.T) {
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
	_, err := client.Policies.UpdateRule(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxx",
		privyclient.PolicyUpdateRuleParams{
			PolicyID: "xxxxxxxxxxxxxxxxxxxxxxxx",
			PolicyRuleRequestBody: privyclient.PolicyRuleRequestBody{
				Action: privyclient.PolicyActionAllow,
				Conditions: []privyclient.PolicyConditionUnion{{
					OfEthereumTransaction: &privyclient.EthereumTransactionCondition{
						Field:       privyclient.EthereumTransactionConditionFieldTo,
						FieldSource: privyclient.EthereumTransactionConditionFieldSourceEthereumTransaction,
						Operator:    privyclient.ConditionOperatorEq,
						Value: privyclient.ConditionValueUnion{
							OfString: privyclient.String("string"),
						},
					},
				}},
				Method: privyclient.PolicyMethodEthSendTransaction,
				Name:   "x",
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
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
