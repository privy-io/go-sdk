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
				OfEthSignTransaction: &privyclient.EthereumSignTransactionRpcInput{
					Method: privyclient.EthereumSignTransactionRpcInputMethodEthSignTransaction,
					Params: privyclient.EthereumSignTransactionRpcInputParams{
						Transaction: privyclient.UnsignedStandardEthereumTransaction{
							AuthorizationList: []privyclient.EthereumSign7702Authorization{{
								ChainID: privyclient.QuantityUnion{
									OfString: privyclient.String("string"),
								},
								Contract: "contract",
								Nonce: privyclient.QuantityUnion{
									OfString: privyclient.String("string"),
								},
								R:       "string",
								S:       "string",
								YParity: 0,
							}},
							ChainID: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
							Data: privyclient.String("string"),
							From: privyclient.String("from"),
							GasLimit: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
							GasPrice: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
							MaxFeePerGas: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
							MaxPriorityFeePerGas: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
							Nonce: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
							To:   privyclient.String("to"),
							Type: 0,
							Value: privyclient.QuantityUnion{
								OfString: privyclient.String("string"),
							},
						},
					},
					Address:   privyclient.String("address"),
					ChainType: privyclient.EthereumSignTransactionRpcInputChainTypeEthereum,
					WalletID:  privyclient.String("wallet_id"),
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

func TestIntentTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Intents.Transfer(
		context.TODO(),
		"wallet_id",
		privyclient.IntentTransferParams{
			TransferRequestBody: privyclient.TransferRequestBody{
				Destination: privyclient.TokenTransferDestination{
					Address: "0xB00F0759DbeeF5E543Cc3E3B07A6442F5f3928a2",
					Asset:   privyclient.String("asset"),
					Chain:   privyclient.String("chain"),
				},
				Source: privyclient.TokenTransferSource{
					Amount: "10.5",
					Asset:  "usdc",
					Chain:  "base",
				},
				AmountType:  privyclient.AmountTypeExactInput,
				SlippageBps: privyclient.Int(0),
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
			KeyQuorumUpdateRequestBody: privyclient.KeyQuorumUpdateRequestBody{
				AuthorizationThreshold: privyclient.Float(0),
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
				AdditionalSigners: privyclient.AdditionalSignerInput{privyclient.AdditionalSignerItemInput{
					SignerID:          "string",
					OverridePolicyIDs: privyclient.PolicyInput{"xxxxxxxxxxxxxxxxxxxxxxxx"},
				}},
				DisplayName: privyclient.String("display_name"),
				Owner: privyclient.OwnerInputUnion{
					OfOwnerInputUser: &privyclient.OwnerInputUser{
						UserID: "user_id",
					},
				},
				OwnerID:   privyclient.String("string"),
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
