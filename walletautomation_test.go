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

func TestWalletAutomationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WalletAutomations.New(context.TODO(), privyclient.WalletAutomationNewParams{
		CreateAutomationRequestBody: privyclient.CreateAutomationRequestBody{
			Config: privyclient.AutomationConfigInput{
				Action: privyclient.AutomationActionConfigInputUnion{
					OfSwap: &privyclient.AutomationSwapActionConfigInput{
						DestinationChainAsset: privyclient.AutomationDestinationAssetInput{
							AutomationDestinationAsset: privyclient.AutomationDestinationAsset{
								AssetAddress: "x",
								Caip2:        "x",
							},
							Asset: privyclient.String("x"),
							Chain: privyclient.String("x"),
						},
						Type: privyclient.AutomationSwapActionConfigInputTypeSwap,
					},
				},
				Trigger: privyclient.AutomationTriggerConfigInput{
					Assets: privyclient.AutomationAssetFilterInputUnion{
						OfAll: &privyclient.AutomationAssetFilterAll{
							Mode: privyclient.AutomationAssetFilterAllModeAll,
						},
					},
					Type: privyclient.AutomationTriggerConfigInputTypeDeposit,
				},
			},
			OwnerID: privyclient.String("x"),
			Name:    privyclient.String("x"),
		},
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletAutomationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WalletAutomations.Update(
		context.TODO(),
		"automation_id",
		privyclient.WalletAutomationUpdateParams{
			UpdateAutomationRequestBody: privyclient.UpdateAutomationRequestBody{
				Config: privyclient.AutomationConfigInput{
					Action: privyclient.AutomationActionConfigInputUnion{
						OfSwap: &privyclient.AutomationSwapActionConfigInput{
							DestinationChainAsset: privyclient.AutomationDestinationAssetInput{
								AutomationDestinationAsset: privyclient.AutomationDestinationAsset{
									AssetAddress: "x",
									Caip2:        "x",
								},
								Asset: privyclient.String("x"),
								Chain: privyclient.String("x"),
							},
							Type: privyclient.AutomationSwapActionConfigInputTypeSwap,
						},
					},
					Trigger: privyclient.AutomationTriggerConfigInput{
						Assets: privyclient.AutomationAssetFilterInputUnion{
							OfAll: &privyclient.AutomationAssetFilterAll{
								Mode: privyclient.AutomationAssetFilterAllModeAll,
							},
						},
						Type: privyclient.AutomationTriggerConfigInputTypeDeposit,
					},
				},
				Enabled: privyclient.Bool(true),
				Name:    privyclient.String("x"),
				OwnerID: privyclient.String("string"),
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

func TestWalletAutomationListWithOptionalParams(t *testing.T) {
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
	_, err := client.WalletAutomations.List(context.TODO(), privyclient.WalletAutomationListParams{
		Cursor:   privyclient.String("cursor"),
		Limit:    privyclient.Int(1),
		WalletID: privyclient.String("x"),
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletAutomationDelete(t *testing.T) {
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
	_, err := client.WalletAutomations.Delete(context.TODO(), "automation_id")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletAutomationGet(t *testing.T) {
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
	_, err := client.WalletAutomations.Get(context.TODO(), "automation_id")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletAutomationListExecutionsWithOptionalParams(t *testing.T) {
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
	_, err := client.WalletAutomations.ListExecutions(context.TODO(), privyclient.WalletAutomationListExecutionsParams{
		Cursor:   privyclient.String("cursor"),
		Limit:    privyclient.Int(1),
		WalletID: privyclient.String("x"),
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletAutomationReindexWithOptionalParams(t *testing.T) {
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
	_, err := client.WalletAutomations.Reindex(context.TODO(), privyclient.WalletAutomationReindexParams{
		WalletAutomationReindexRequestBody: privyclient.WalletAutomationReindexRequestBody{
			AssetAddress:   "x",
			Caip2:          privyclient.TronCaip2TronMainnet,
			Chain:          privyclient.String("x"),
			DepositAddress: privyclient.String("x"),
			WalletID:       privyclient.String("x"),
		},
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
