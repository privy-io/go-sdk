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
	_, err := client.Wallets.New(context.TODO(), privyclient.WalletNewParams{
		ChainType: privyclient.WalletChainTypeEthereum,
		AdditionalSigners: privyclient.AdditionalSignerInput{privyclient.AdditionalSignerItemInput{
			SignerID:          "string",
			OverridePolicyIDs: privyclient.PolicyInput{"xxxxxxxxxxxxxxxxxxxxxxxx"},
		}},
		DisplayName: privyclient.String("display_name"),
		ExternalID:  privyclient.String("my-order-123"),
		Owner: privyclient.OwnerInputUnion{
			OfOwnerInputUser: &privyclient.OwnerInputUser{
				UserID: "user_id",
			},
		},
		OwnerID:             privyclient.String("string"),
		PolicyIDs:           privyclient.PolicyInput{"xxxxxxxxxxxxxxxxxxxxxxxx"},
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
	_, err := client.Wallets.Update(
		context.TODO(),
		"wallet_id",
		privyclient.WalletUpdateParams{
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
				PolicyIDs: []string{"tb54eps4z44ed0jepousxi4n"},
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

func TestWalletListWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.List(context.TODO(), privyclient.WalletListParams{
		AuthorizationKey: privyclient.String("s=-/fw-L-+N\n"),
		ChainType:        privyclient.WalletChainTypeEthereum,
		Cursor:           privyclient.String("x"),
		ExternalID:       privyclient.String("external_id"),
		Limit:            privyclient.Float(100),
		UserID:           privyclient.String("user_id"),
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletInitImport(t *testing.T) {
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
	_, err := client.Wallets.InitImport(context.TODO(), privyclient.WalletInitImportParams{
		OfPrivateKey: &privyclient.PrivateKeyInitInput{
			Address:        "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
			ChainType:      privyclient.WalletImportSupportedChainsEthereum,
			EncryptionType: privyclient.HpkeEncryptionHpke,
			EntropyType:    privyclient.PrivateKeyInitInputEntropyTypePrivateKey,
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

func TestWalletSubmitImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.SubmitImport(context.TODO(), privyclient.WalletSubmitImportParams{
		Wallet: privyclient.WalletSubmitImportParamsWalletUnion{
			OfPrivateKey: &privyclient.PrivateKeySubmitInput{
				Address:         "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
				ChainType:       privyclient.WalletImportSupportedChainsEthereum,
				Ciphertext:      "PRoRXygG+YYSDBXjCopNYZmx8Z6nvdl1D0lpePTYZdZI2VGfK+LkFt+GlEJqdoi9",
				EncapsulatedKey: "BOhR6xITDt5THJawHHJKrKdI9CBr2M/SDWzZZAaOW4gCMsSpC65U007WyKiwuuOVAo1BNm4YgcBBROuMmyIZXZk=",
				EncryptionType:  privyclient.HpkeEncryptionHpke,
				EntropyType:     privyclient.PrivateKeySubmitInputEntropyTypePrivateKey,
				HpkeConfig: privyclient.HpkeImportConfig{
					Aad:           privyclient.String("aad"),
					AeadAlgorithm: privyclient.HpkeAeadAlgorithmChacha20Poly1305,
					Info:          privyclient.String("info"),
				},
			},
		},
		AdditionalSigners: privyclient.AdditionalSignerInput{privyclient.AdditionalSignerItemInput{
			SignerID:          "string",
			OverridePolicyIDs: privyclient.PolicyInput{"xxxxxxxxxxxxxxxxxxxxxxxx"},
		}},
		DisplayName: privyclient.String("display_name"),
		ExternalID:  privyclient.String("external_id"),
		Owner: privyclient.OwnerInputUnion{
			OfOwnerInputUser: &privyclient.OwnerInputUser{
				UserID: "user_id",
			},
		},
		OwnerID:   privyclient.String("rkiz0ivz254drv1xw982v3jq"),
		PolicyIDs: privyclient.PolicyInput{"xxxxxxxxxxxxxxxxxxxxxxxx"},
	})
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Transfer(
		context.TODO(),
		"wallet_id",
		privyclient.WalletTransferParams{
			TransferRequestBody: privyclient.TransferRequestBody{
				Destination: privyclient.TokenTransferDestination{
					Address: "0xB00F0759DbeeF5E543Cc3E3B07A6442F5f3928a2",
				},
				Source: privyclient.TokenTransferSource{
					Amount: "10.5",
					Asset:  "usdc",
					Chain:  "base",
				},
			},
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

func TestWalletAuthenticateWithJwt(t *testing.T) {
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
	_, err := client.Wallets.AuthenticateWithJwt(context.TODO(), privyclient.WalletAuthenticateWithJwtParams{
		WalletAuthenticateRequestBody: privyclient.WalletAuthenticateRequestBody{
			EncryptionType:     privyclient.WalletAuthenticateRequestBodyEncryptionTypeHpke,
			RecipientPublicKey: "DAQcDQgAEx4aoeD72yykviK+fckqE2CItVIGn1rCnvCXZ1HgpOcMEMialRmTrqIK4oZlYd1",
			UserJwt:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30",
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

func TestWalletExportWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Export(
		context.TODO(),
		"wallet_id",
		privyclient.WalletExportParams{
			WalletExportRequestBody: privyclient.WalletExportRequestBody{
				EncryptionType:     privyclient.HpkeEncryptionHpke,
				RecipientPublicKey: "BDAZLOIdTaPycEYkgG0MvCzbIKJLli/yWkAV5yCa9yOsZ4JsrLweA5MnP8YIiY4k/RRzC+APhhO+P+Hoz/rt7Go=",
				ExportSeedPhrase:   privyclient.Bool(true),
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

func TestWalletGet(t *testing.T) {
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
	_, err := client.Wallets.Get(context.TODO(), "wallet_id")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWalletGetWalletByAddress(t *testing.T) {
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
	_, err := client.Wallets.GetWalletByAddress(context.TODO(), privyclient.WalletGetWalletByAddressParams{
		GetByWalletAddressRequestBody: privyclient.GetByWalletAddressRequestBody{
			Address: "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
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

func TestWalletRawSignWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.RawSign(
		context.TODO(),
		"wallet_id",
		privyclient.WalletRawSignParams{
			RawSignInput: privyclient.RawSignInput{
				Params: privyclient.RawSignInputParamsUnion{
					OfRawSignHashs: &privyclient.RawSignHashParams{
						Hash: "0x0775aeed9c9ce6e0fbc4db25c5e4e6368029651c905c286f813126a09025a21e",
					},
				},
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
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

func TestWalletRpcWithOptionalParams(t *testing.T) {
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
	_, err := client.Wallets.Rpc(
		context.TODO(),
		"wallet_id",
		privyclient.WalletRpcParams{
			WalletRpcRequestBody: privyclient.WalletRpcRequestBodyUnion{
				OfEthSendTransaction: &privyclient.EthereumSendTransactionRpcInput{
					Caip2:  "eip155:8453",
					Method: privyclient.EthereumSendTransactionRpcInputMethodEthSendTransaction,
					Params: privyclient.EthereumSendTransactionRpcInputParams{
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
							To:   privyclient.String("0x0000000000000000000000000000000000000000"),
							Type: 0,
							Value: privyclient.QuantityUnion{
								OfInt: privyclient.Int(1),
							},
						},
					},
					Address:                privyclient.String("address"),
					ChainType:              privyclient.EthereumSendTransactionRpcInputChainTypeEthereum,
					ExperimentalDataSuffix: privyclient.String("string"),
					ReferenceID:            privyclient.String("x"),
					Sponsor:                privyclient.Bool(true),
					WalletID:               privyclient.String("wallet_id"),
				},
			},
			PrivyAuthorizationSignature: privyclient.String("privy-authorization-signature"),
			PrivyIdempotencyKey:         privyclient.String("privy-idempotency-key"),
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
