// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"os"
	"testing"

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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.New(context.TODO(), WalletNewParams{
		ChainType: WalletChainTypeEthereum,
		AdditionalSigners: []WalletNewParamsAdditionalSigner{{
			SignerID:          "signer_id",
			OverridePolicyIDs: []string{"string"},
		}},
		Owner: WalletNewParamsOwnerUnion{
			OfPublicKeyOwner: &WalletNewParamsOwnerPublicKeyOwner{
				PublicKey: "public_key",
			},
		},
		OwnerID:             String("owner_id"),
		PolicyIDs:           []string{"xxxxxxxxxxxxxxxxxxxxxxxx"},
		PrivyIdempotencyKey: String("privy-idempotency-key"),
	})
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Update(
		context.TODO(),
		"wallet_id",
		WalletUpdateParams{
			AdditionalSigners: []WalletUpdateParamsAdditionalSigner{{
				SignerID:          "signer_id",
				OverridePolicyIDs: []string{"string"},
			}},
			Owner: WalletUpdateParamsOwnerUnion{
				OfPublicKeyOwner: &WalletUpdateParamsOwnerPublicKeyOwner{
					PublicKey: "public_key",
				},
			},
			OwnerID:                     String("owner_id"),
			PolicyIDs:                   []string{"tb54eps4z44ed0jepousxi4n"},
			PrivyAuthorizationSignature: String("privy-authorization-signature"),
		},
	)
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.List(context.TODO(), WalletListParams{
		AuthorizationKey: String("s=-/fw-L-+N\n"),
		ChainType:        WalletChainTypeEthereum,
		Cursor:           String("x"),
		Limit:            Float(100),
		UserID:           String("user_id"),
	})
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets._InitImport(context.TODO(), Wallet_InitImportParams{
		OfPrivateKey: &Wallet_InitImportParamsBodyPrivateKey{
			Address:        "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
			ChainType:      "ethereum",
			EncryptionType: "HPKE",
		},
	})
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets._SubmitImport(context.TODO(), Wallet_SubmitImportParams{
		Wallet: Wallet_SubmitImportParamsWalletUnion{
			OfPrivateKey: &Wallet_SubmitImportParamsWalletPrivateKey{
				Address:         "0xF1DBff66C993EE895C8cb176c30b07A559d76496",
				ChainType:       "ethereum",
				Ciphertext:      "PRoRXygG+YYSDBXjCopNYZmx8Z6nvdl1D0lpePTYZdZI2VGfK+LkFt+GlEJqdoi9",
				EncapsulatedKey: "BOhR6xITDt5THJawHHJKrKdI9CBr2M/SDWzZZAaOW4gCMsSpC65U007WyKiwuuOVAo1BNm4YgcBBROuMmyIZXZk=",
				EncryptionType:  "HPKE",
				HpkeConfig: HpkeImportConfigParam{
					Aad:           String("aad"),
					AeadAlgorithm: HpkeImportConfigAeadAlgorithmChacha20Poly1305,
					Info:          String("info"),
				},
			},
		},
		AdditionalSigners: []Wallet_SubmitImportParamsAdditionalSigner{{
			SignerID:          "signer_id",
			OverridePolicyIDs: []string{"string"},
		}},
		Owner: Wallet_SubmitImportParamsOwnerUnion{
			OfWallet_SubmitImportsOwnerUserID: &Wallet_SubmitImportParamsOwnerUserID{
				UserID: "user_id",
			},
		},
		OwnerID:   String("rkiz0ivz254drv1xw982v3jq"),
		PolicyIDs: []string{"string"},
	})
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.AuthenticateWithJwt(context.TODO(), WalletAuthenticateWithJwtParams{
		UserJwt:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30",
		EncryptionType:     WalletAuthenticateWithJwtParamsEncryptionTypeHpke,
		RecipientPublicKey: String("DAQcDQgAEx4aoeD72yykviK+fckqE2CItVIGn1rCnvCXZ1HgpOcMEMialRmTrqIK4oZlYd1"),
	})
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Export(
		context.TODO(),
		"wallet_id",
		WalletExportParams{
			EncryptionType:              WalletExportParamsEncryptionTypeHpke,
			RecipientPublicKey:          "BDAZLOIdTaPycEYkgG0MvCzbIKJLli/yWkAV5yCa9yOsZ4JsrLweA5MnP8YIiY4k/RRzC+APhhO+P+Hoz/rt7Go=",
			PrivyAuthorizationSignature: String("privy-authorization-signature"),
		},
	)
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Get(context.TODO(), "wallet_id")
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.RawSign(
		context.TODO(),
		"wallet_id",
		WalletRawSignParams{
			Params: WalletRawSignParamsParamsUnion{
				OfHash: &WalletRawSignParamsParamsHash{
					Hash: "0x0775aeed9c9ce6e0fbc4db25c5e4e6368029651c905c286f813126a09025a21e",
				},
			},
			PrivyAuthorizationSignature: String("privy-authorization-signature"),
			PrivyIdempotencyKey:         String("privy-idempotency-key"),
		},
	)
	if err != nil {
		var apierr *Error
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
	client := NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Wallets.Rpc(
		context.TODO(),
		"wallet_id",
		WalletRpcParams{
			OfEthSendTransaction: &EthereumSendTransactionRpcInputParam{
				Caip2:  "eip155:8453",
				Method: EthereumSendTransactionRpcInputMethodEthSendTransaction,
				Params: EthereumSendTransactionRpcInputParamsParam{
					Transaction: EthereumSendTransactionRpcInputParamsTransactionParam{
						AuthorizationList: []EthereumSendTransactionRpcInputParamsTransactionAuthorizationListParam{{
							ChainID: EthereumSendTransactionRpcInputParamsTransactionAuthorizationListChainIDUnionParam{
								OfString: String("string"),
							},
							Contract: "contract",
							Nonce: EthereumSendTransactionRpcInputParamsTransactionAuthorizationListNonceUnionParam{
								OfString: String("string"),
							},
							R:       "r",
							S:       "s",
							YParity: 0,
						}},
						ChainID: EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam{
							OfString: String("string"),
						},
						Data: String("data"),
						From: String("from"),
						GasLimit: EthereumSendTransactionRpcInputParamsTransactionGasLimitUnionParam{
							OfString: String("string"),
						},
						GasPrice: EthereumSendTransactionRpcInputParamsTransactionGasPriceUnionParam{
							OfString: String("string"),
						},
						MaxFeePerGas: EthereumSendTransactionRpcInputParamsTransactionMaxFeePerGasUnionParam{
							OfString: String("string"),
						},
						MaxPriorityFeePerGas: EthereumSendTransactionRpcInputParamsTransactionMaxPriorityFeePerGasUnionParam{
							OfString: String("string"),
						},
						Nonce: EthereumSendTransactionRpcInputParamsTransactionNonceUnionParam{
							OfString: String("string"),
						},
						To:   String("0x0000000000000000000000000000000000000000"),
						Type: 0,
						Value: EthereumSendTransactionRpcInputParamsTransactionValueUnionParam{
							OfInt: Int(1),
						},
					},
				},
				Address:   String("address"),
				ChainType: EthereumSendTransactionRpcInputChainTypeEthereum,
				Sponsor:   Bool(true),
			},
			PrivyAuthorizationSignature: String("privy-authorization-signature"),
			PrivyIdempotencyKey:         String("privy-idempotency-key"),
		},
	)
	if err != nil {
		var apierr *Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
