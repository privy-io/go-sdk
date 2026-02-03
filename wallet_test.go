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

func TestWalletNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.New(context.TODO(), privyclient.WalletNewParams{
		ChainType: privyclient.WalletNewParamsChainTypeEthereum,
		AdditionalSigners: []privyclient.WalletNewParamsAdditionalSigner{{
			OverridePolicyIDs: []string{"string"},
			SignerID:          "signer_id",
		}},
		Owner: privyclient.WalletNewParamsOwnerUnion{
			OfPublicKeyOwner: &privyclient.WalletNewParamsOwnerPublicKeyOwner{
				PublicKey: "public_key",
			},
		},
		OwnerID:             privyclient.String("owner_id"),
		PolicyIDs:           []string{"xxxxxxxxxxxxxxxxxxxxxxxx"},
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

func TestWalletUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Update(
		context.TODO(),
		"wallet_id",
		privyclient.WalletUpdateParams{
			AdditionalSigners: []privyclient.WalletUpdateParamsAdditionalSigner{{
				OverridePolicyIDs: []string{"string"},
				SignerID:          "signer_id",
			}},
			Owner: privyclient.WalletUpdateParamsOwnerUnion{
				OfPublicKeyOwner: &privyclient.WalletUpdateParamsOwnerPublicKeyOwner{
					PublicKey: "public_key",
				},
			},
			OwnerID:                     privyclient.String("owner_id"),
			PolicyIDs:                   []string{"tb54eps4z44ed0jepousxi4n"},
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

func TestWalletListWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.List(context.TODO(), privyclient.WalletListParams{
		ChainType: privyclient.WalletListParamsChainTypeCosmos,
		Cursor:    privyclient.String("x"),
		Limit:     privyclient.Float(100),
		UserID:    privyclient.String("user_id"),
	})
	if err != nil {
		var apierr *privyclient.Error
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
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets._InitImport(context.TODO(), privyclient.Wallet_InitImportParams{
		OfPrivateKey: &privyclient.Wallet_InitImportParamsBodyPrivateKey{
			Address:        "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
			ChainType:      "ethereum",
			EncryptionType: "HPKE",
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

func TestWallet_SubmitImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets._SubmitImport(context.TODO(), privyclient.Wallet_SubmitImportParams{
		Wallet: privyclient.Wallet_SubmitImportParamsWalletUnion{
			OfPrivateKey: &privyclient.Wallet_SubmitImportParamsWalletPrivateKey{
				Address:         "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
				ChainType:       "ethereum",
				Ciphertext:      "PRoRXygG+YYSDBXjCopNYZmx8Z6nvdl1D0lpePTYZdZI2VGfK+LkFt+GlEJqdoi9",
				EncapsulatedKey: "BOhR6xITDt5THJawHHJKrKdI9CBr2M/SDWzZZAaOW4gCMsSpC65U007WyKiwuuOVAo1BNm4YgcBBROuMmyIZXZk=",
				EncryptionType:  "HPKE",
			},
		},
		AdditionalSigners: []privyclient.Wallet_SubmitImportParamsAdditionalSigner{{
			SignerID:          "signer_id",
			OverridePolicyIDs: []string{"string"},
		}},
		Owner: privyclient.Wallet_SubmitImportParamsOwnerUnion{
			OfWallet_SubmitImportsOwnerUserID: &privyclient.Wallet_SubmitImportParamsOwnerUserID{
				UserID: "user_id",
			},
		},
		OwnerID:   privyclient.String("rkiz0ivz254drv1xw982v3jq"),
		PolicyIDs: []string{"string"},
	})
	if err != nil {
		var apierr *privyclient.Error
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
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.AuthenticateWithJwt(context.TODO(), privyclient.WalletAuthenticateWithJwtParams{
		UserJwt:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30",
		EncryptionType:     privyclient.WalletAuthenticateWithJwtParamsEncryptionTypeHpke,
		RecipientPublicKey: privyclient.String("DAQcDQgAEx4aoeD72yykviK+fckqE2CItVIGn1rCnvCXZ1HgpOcMEMialRmTrqIK4oZlYd1"),
	})
	if err != nil {
		var apierr *privyclient.Error
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
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Export(
		context.TODO(),
		"wallet_id",
		privyclient.WalletExportParams{
			EncryptionType:              privyclient.WalletExportParamsEncryptionTypeHpke,
			RecipientPublicKey:          "BDAZLOIdTaPycEYkgG0MvCzbIKJLli/yWkAV5yCa9yOsZ4JsrLweA5MnP8YIiY4k/RRzC+APhhO+P+Hoz/rt7Go=",
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

func TestWalletGet(t *testing.T) {
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
	_, err := client.Wallets.Get(context.TODO(), "wallet_id")
	if err != nil {
		var apierr *privyclient.Error
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
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.RawSign(
		context.TODO(),
		"wallet_id",
		privyclient.WalletRawSignParams{
			Params: privyclient.WalletRawSignParamsParams{
				Hash: privyclient.String("0x0775aeed9c9ce6e0fbc4db25c5e4e6368029651c905c286f813126a09025a21e"),
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
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

func TestWalletRpcWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Rpc(
		context.TODO(),
		"wallet_id",
		privyclient.WalletRpcParams{
			OfEthSendTransaction: &privyclient.WalletRpcParamsBodyEthSendTransaction{
				Caip2: "eip155:8453",
				Params: privyclient.WalletRpcParamsBodyEthSendTransactionParams{
					Transaction: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransaction{
						ChainID: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionChainIDUnion{
							OfString: privyclient.String("string"),
						},
						Data: privyclient.String("data"),
						From: privyclient.String("from"),
						GasLimit: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionGasLimitUnion{
							OfString: privyclient.String("string"),
						},
						GasPrice: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionGasPriceUnion{
							OfString: privyclient.String("string"),
						},
						MaxFeePerGas: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxFeePerGasUnion{
							OfString: privyclient.String("string"),
						},
						MaxPriorityFeePerGas: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionMaxPriorityFeePerGasUnion{
							OfString: privyclient.String("string"),
						},
						Nonce: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionNonceUnion{
							OfString: privyclient.String("string"),
						},
						To:   privyclient.String("0x0000000000000000000000000000000000000000"),
						Type: 0,
						Value: privyclient.WalletRpcParamsBodyEthSendTransactionParamsTransactionValueUnion{
							OfInt: privyclient.Int(1),
						},
					},
				},
				Address:   privyclient.String("address"),
				ChainType: "ethereum",
				Sponsor:   privyclient.Bool(true),
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
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
