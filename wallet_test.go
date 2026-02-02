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

func TestWalletNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.New(context.TODO(), privyapiclient.WalletNewParams{
		ChainType: privyapiclient.WalletNewParamsChainTypeEthereum,
		AdditionalSigners: []privyapiclient.WalletNewParamsAdditionalSigner{{
			OverridePolicyIDs: []string{"string"},
			SignerID:          "signer_id",
		}},
		Owner: privyapiclient.WalletNewParamsOwnerUnion{
			OfPublicKeyOwner: &privyapiclient.WalletNewParamsOwnerPublicKeyOwner{
				PublicKey: "public_key",
			},
		},
		OwnerID:             privyapiclient.String("owner_id"),
		PolicyIDs:           []string{"xxxxxxxxxxxxxxxxxxxxxxxx"},
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

func TestWalletUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Update(
		context.TODO(),
		"wallet_id",
		privyapiclient.WalletUpdateParams{
			AdditionalSigners: []privyapiclient.WalletUpdateParamsAdditionalSigner{{
				OverridePolicyIDs: []string{"string"},
				SignerID:          "signer_id",
			}},
			Owner: privyapiclient.WalletUpdateParamsOwnerUnion{
				OfPublicKeyOwner: &privyapiclient.WalletUpdateParamsOwnerPublicKeyOwner{
					PublicKey: "public_key",
				},
			},
			OwnerID:                     privyapiclient.String("owner_id"),
			PolicyIDs:                   []string{"tb54eps4z44ed0jepousxi4n"},
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

func TestWalletListWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.List(context.TODO(), privyapiclient.WalletListParams{
		ChainType: privyapiclient.WalletListParamsChainTypeCosmos,
		Cursor:    privyapiclient.String("x"),
		Limit:     privyapiclient.Float(100),
		UserID:    privyapiclient.String("user_id"),
	})
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWallet_InitImport(t *testing.T) {
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
	_, err := client.Wallets._InitImport(context.TODO(), privyapiclient.Wallet_InitImportParams{
		OfPrivateKey: &privyapiclient.Wallet_InitImportParamsBodyPrivateKey{
			Address:        "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
			ChainType:      "ethereum",
			EncryptionType: "HPKE",
		},
	})
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWallet_SubmitImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets._SubmitImport(context.TODO(), privyapiclient.Wallet_SubmitImportParams{
		Wallet: privyapiclient.Wallet_SubmitImportParamsWalletUnion{
			OfPrivateKey: &privyapiclient.Wallet_SubmitImportParamsWalletPrivateKey{
				Address:         "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
				ChainType:       "ethereum",
				Ciphertext:      "PRoRXygG+YYSDBXjCopNYZmx8Z6nvdl1D0lpePTYZdZI2VGfK+LkFt+GlEJqdoi9",
				EncapsulatedKey: "BOhR6xITDt5THJawHHJKrKdI9CBr2M/SDWzZZAaOW4gCMsSpC65U007WyKiwuuOVAo1BNm4YgcBBROuMmyIZXZk=",
				EncryptionType:  "HPKE",
			},
		},
		AdditionalSigners: []privyapiclient.Wallet_SubmitImportParamsAdditionalSigner{{
			SignerID:          "signer_id",
			OverridePolicyIDs: []string{"string"},
		}},
		Owner: privyapiclient.Wallet_SubmitImportParamsOwnerUnion{
			OfWallet_SubmitImportsOwnerUserID: &privyapiclient.Wallet_SubmitImportParamsOwnerUserID{
				UserID: "user_id",
			},
		},
		OwnerID:   privyapiclient.String("rkiz0ivz254drv1xw982v3jq"),
		PolicyIDs: []string{"string"},
	})
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletAuthenticateWithJwtWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.AuthenticateWithJwt(context.TODO(), privyapiclient.WalletAuthenticateWithJwtParams{
		UserJwt:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30",
		EncryptionType:     privyapiclient.WalletAuthenticateWithJwtParamsEncryptionTypeHpke,
		RecipientPublicKey: privyapiclient.String("DAQcDQgAEx4aoeD72yykviK+fckqE2CItVIGn1rCnvCXZ1HgpOcMEMialRmTrqIK4oZlYd1"),
	})
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletExportWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Export(
		context.TODO(),
		"wallet_id",
		privyapiclient.WalletExportParams{
			EncryptionType:              privyapiclient.WalletExportParamsEncryptionTypeHpke,
			RecipientPublicKey:          "BDAZLOIdTaPycEYkgG0MvCzbIKJLli/yWkAV5yCa9yOsZ4JsrLweA5MnP8YIiY4k/RRzC+APhhO+P+Hoz/rt7Go=",
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

func TestWalletGet(t *testing.T) {
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
	_, err := client.Wallets.Get(context.TODO(), "wallet_id")
	if err != nil {
		var apierr *privyapiclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletRawSignWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.RawSign(
		context.TODO(),
		"wallet_id",
		privyapiclient.WalletRawSignParams{
			Params: privyapiclient.WalletRawSignParamsParams{
				Hash: privyapiclient.String("0x0775aeed9c9ce6e0fbc4db25c5e4e6368029651c905c286f813126a09025a21e"),
			},
			PrivyAuthorizationSignature: privyapiclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyapiclient.String("privy-idempotency-key"),
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

func TestWalletRpcWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Rpc(
		context.TODO(),
		"wallet_id",
		privyapiclient.WalletRpcParams{
			OfEthSendTransaction: &privyapiclient.WalletRpcParamsBodyEthSendTransaction{
				Caip2: "eip155:8453",
				Params: privyapiclient.WalletRpcParamsBodyEthSendTransactionParams{
					Transaction: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransaction{
						ChainID: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion{
							OfString: privyapiclient.String("string"),
						},
						Data: privyapiclient.String("data"),
						From: privyapiclient.String("from"),
						GasLimit: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion{
							OfString: privyapiclient.String("string"),
						},
						GasPrice: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion{
							OfString: privyapiclient.String("string"),
						},
						MaxFeePerGas: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion{
							OfString: privyapiclient.String("string"),
						},
						MaxPriorityFeePerGas: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion{
							OfString: privyapiclient.String("string"),
						},
						Nonce: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion{
							OfString: privyapiclient.String("string"),
						},
						To:   privyapiclient.String("0x0000000000000000000000000000000000000000"),
						Type: 0,
						Value: privyapiclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion{
							OfInt: privyapiclient.Int(1),
						},
					},
				},
				Address:   privyapiclient.String("address"),
				ChainType: "ethereum",
				Sponsor:   privyapiclient.Bool(true),
			},
			PrivyAuthorizationSignature: privyapiclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyapiclient.String("privy-idempotency-key"),
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
