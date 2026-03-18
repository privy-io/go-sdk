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

func TestIntentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.List(context.TODO(), privyclient.IntentListParams{
		CreatedByID:          privyclient.String("created_by_id"),
		CurrentUserHasSigned: privyclient.IntentListParamsCurrentUserHasSignedTrue,
		Cursor:               privyclient.String("x"),
		IntentType:           privyclient.IntentTypeKeyQuorum,
		Limit:                privyclient.Float(100),
		PendingMemberID:      privyclient.String("pending_member_id"),
		ResourceID:           privyclient.String("resource_id"),
		SortBy:               privyclient.IntentListParamsSortByCreatedAtDesc,
		Status:               privyclient.IntentStatusPending,
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntentNewPolicyRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.NewPolicyRule(
		context.TODO(),
		"policy_id",
		privyclient.IntentNewPolicyRuleParams{
			Action: privyclient.IntentNewPolicyRuleParamsActionAllow,
			Conditions: []privyclient.IntentNewPolicyRuleParamsConditionUnion{{
				OfEthereumTransaction: &privyclient.IntentNewPolicyRuleParamsConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyclient.IntentNewPolicyRuleParamsConditionEthereumTransactionValueUnion{
						OfString: privyclient.String("string"),
					},
				},
			}},
			Method:             privyclient.IntentNewPolicyRuleParamsMethodEthSendTransaction,
			Name:               "x",
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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

func TestIntentDeletePolicyRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.DeletePolicyRule(
		context.TODO(),
		"rule_id",
		privyclient.IntentDeletePolicyRuleParams{
			PolicyID:           "policy_id",
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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

func TestIntentGet(t *testing.T) {
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
	_, err := client.Intents.Get(context.TODO(), "intent_id")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntentRpcWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.Rpc(
		context.TODO(),
		"wallet_id",
		privyclient.IntentRpcParams{
			WalletRpcRequestBody: privyclient.WalletRpcRequestBodyUnion{
				OfPersonalSign: &privyclient.EthereumPersonalSignRpcInput{
					Method: privyclient.EthereumPersonalSignRpcInputMethodPersonalSign,
					Params: privyclient.EthereumPersonalSignRpcInputParams{
						Encoding: privyclient.EthereumPersonalSignRpcInputParamsEncodingUtf8,
						Message:  "message",
					},
					Address:   privyclient.String("address"),
					ChainType: privyclient.EthereumPersonalSignRpcInputChainTypeEthereum,
				},
			},
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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

func TestIntentUpdateKeyQuorumWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.UpdateKeyQuorum(
		context.TODO(),
		"key_quorum_id",
		privyclient.IntentUpdateKeyQuorumParams{
			KeyQuorumCreateParams: privyclient.KeyQuorumCreateParams{
				AuthorizationThreshold: privyclient.Float(1),
				DisplayName:            privyclient.String("display_name"),
				KeyQuorumIDs:           []string{"string"},
				PublicKeys:             []string{"string"},
				UserIDs:                []string{"string"},
			},
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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

func TestIntentUpdatePolicyWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.UpdatePolicy(
		context.TODO(),
		"policy_id",
		privyclient.IntentUpdatePolicyParams{
			Name: privyclient.String("x"),
			Owner: privyclient.IntentUpdatePolicyParamsOwnerUnion{
				OfPublicKeyOwner: &privyclient.IntentUpdatePolicyParamsOwnerPublicKeyOwner{
					PublicKey: "public_key",
				},
			},
			OwnerID: privyclient.String("owner_id"),
			Rules: []privyclient.IntentUpdatePolicyParamsRule{{
				Action: "ALLOW",
				Conditions: []privyclient.IntentUpdatePolicyParamsRuleConditionUnion{{
					OfEthereumTransaction: &privyclient.IntentUpdatePolicyParamsRuleConditionEthereumTransaction{
						Field:    "to",
						Operator: "eq",
						Value: privyclient.IntentUpdatePolicyParamsRuleConditionEthereumTransactionValueUnion{
							OfString: privyclient.String("string"),
						},
					},
				}},
				Method: "eth_sendTransaction",
				Name:   "x",
			}},
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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

func TestIntentUpdatePolicyRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.UpdatePolicyRule(
		context.TODO(),
		"rule_id",
		privyclient.IntentUpdatePolicyRuleParams{
			PolicyID: "policy_id",
			Action:   privyclient.IntentUpdatePolicyRuleParamsActionAllow,
			Conditions: []privyclient.IntentUpdatePolicyRuleParamsConditionUnion{{
				OfEthereumTransaction: &privyclient.IntentUpdatePolicyRuleParamsConditionEthereumTransaction{
					Field:    "to",
					Operator: "eq",
					Value: privyclient.IntentUpdatePolicyRuleParamsConditionEthereumTransactionValueUnion{
						OfString: privyclient.String("string"),
					},
				},
			}},
			Method:             privyclient.IntentUpdatePolicyRuleParamsMethodEthSendTransaction,
			Name:               "x",
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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

func TestIntentUpdateWalletWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.UpdateWallet(
		context.TODO(),
		"wallet_id",
		privyclient.IntentUpdateWalletParams{
			WalletUpdateRequestBody: privyclient.WalletUpdateRequestBody{
				AdditionalSigners: []privyclient.WalletUpdateRequestBodyAdditionalSigner{{
					SignerID:          "signer_id",
					OverridePolicyIDs: []string{"string"},
				}},
				Owner: privyclient.WalletUpdateRequestBodyOwnerUnion{
					OfPublicKeyOwner: &privyclient.WalletUpdateRequestBodyOwnerPublicKeyOwner{
						PublicKey: "public_key",
					},
				},
				OwnerID:   privyclient.String("owner_id"),
				PolicyIDs: []string{"xxxxxxxxxxxxxxxxxxxxxxxx"},
			},
			PrivyRequestExpiry: privyclient.String("privy-request-expiry"),
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
