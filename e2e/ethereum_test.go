package e2e_test

import (
	"context"
	"os"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/packages/param"
)

func TestWallets_Ethereum(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()
	jwt := generateTestJWT(t)
	sk := os.Getenv("P256_PRIVATE_KEY")

	wallets := []struct {
		name    string
		id      string
		address string
		authCtx *authorization.AuthorizationContext
	}{
		{
			name:    "Ownerless",
			id:      os.Getenv("OWNERLESS_ETHEREUM_WALLET_ID"),
			address: os.Getenv("OWNERLESS_ETHEREUM_WALLET_ADDRESS"),
			authCtx: nil, // no authorization context for ownerless
		},
		{
			name:    "KeyOwned",
			id:      os.Getenv("P256_OWNED_ETHEREUM_WALLET_ID"),
			address: os.Getenv("P256_OWNED_ETHEREUM_WALLET_ADDRESS"),
			authCtx: &authorization.AuthorizationContext{PrivateKeys: []string{sk}},
		},
		{
			name:    "UserOwned",
			id:      os.Getenv("USER_OWNED_ETHEREUM_WALLET_ID"),
			address: os.Getenv("USER_OWNED_ETHEREUM_WALLET_ADDRESS"),
			authCtx: &authorization.AuthorizationContext{UserJwts: []string{jwt}},
		},
	}

	t.Run("Sign7702Authorization", func(t *testing.T) {
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Ethereum.Sign7702Authorization(ctx, wallet.id,
					EthereumSign7702AuthorizationRpcInputParam{
						Method: EthereumSign7702AuthorizationRpcInputMethodEthSign7702Authorization,
						Params: EthereumSign7702AuthorizationRpcInputParamsParam{
							ChainID: EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam{
								OfInt: param.NewOpt[int64](11155111), // Sepolia
							},
							Contract: "0x1234567890123456789012345678901234567890",
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign 7702 authorization: %v", err)
				}

				auth := data.Authorization
				if auth.Contract == "" {
					t.Error("expected authorization contract to be defined")
				}
				if auth.R == "" {
					t.Error("expected authorization R value to be defined")
				}
				if auth.S == "" {
					t.Error("expected authorization S value to be defined")
				}
				if auth.YParity != 0 && auth.YParity != 1 {
					t.Errorf("expected y_parity to be 0 or 1, got %f", auth.YParity)
				}
			})
		}
	})

	t.Run("SignUserOperation", func(t *testing.T) {
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Ethereum.SignUserOperation(ctx, wallet.id,
					EthereumSignUserOperationRpcInputParam{
						Method: EthereumSignUserOperationRpcInputMethodEthSignUserOperation,
						Params: EthereumSignUserOperationRpcInputParamsParam{
							ChainID: EthereumSignUserOperationRpcInputParamsChainIDUnionParam{
								OfString: param.NewOpt("0x66eee"), // Arbitrum Sepolia
							},
							Contract: "0x69007702764179f14F51cdce752f4f775d74E139",
							UserOperation: EthereumSignUserOperationRpcInputParamsUserOperationParam{
								Sender:                        "0xdf1Bff521006396b2dd11725681ebA6998DB37e3",
								Nonce:                         "0x1000000000000000a",
								CallData:                      "0x34fcd5be000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000036cbd53842c5426634e7929541ec2318f3dcf7e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000044a9059cbb000000000000000000000000cc9c3d98163f4f6af884e259132e15d6d27a5c57000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000",
								CallGasLimit:                  "0x877f",
								VerificationGasLimit:          "0x8550",
								PreVerificationGas:            "0xbcc4",
								MaxFeePerGas:                  "0x18700",
								MaxPriorityFeePerGas:          "0x186a0",
								Paymaster:                     "0x2cc0c7981D846b9F2a16276556f6e8cb52BfB633",
								PaymasterData:                 "0x000000000000000069174750fbd97f1583efc0158107838d694bb88594fc428f431892960194089ef4e15c8b330b402626cd01ce11057354528807581e7400e4edf33e15e3b55bcc6ef63fcf1c",
								PaymasterVerificationGasLimit: "0x7680",
								PaymasterPostOpGasLimit:       "0x00",
							},
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign user operation: %v", err)
				}

				if data.Signature == "" {
					t.Error("expected signature to be defined")
				}
				if data.Encoding != "hex" {
					t.Errorf("expected encoding to be hex, got %s", data.Encoding)
				}
			})
		}
	})

	t.Run("SignSecp256k1", func(t *testing.T) {
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				hash := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
				data, err := client.Wallets.Ethereum.SignSecp256k1(ctx, wallet.id,
					EthereumSecp256k1SignRpcInputParam{
						Method: EthereumSecp256k1SignRpcInputMethodSecp256k1Sign,
						Params: EthereumSecp256k1SignRpcInputParamsParam{
							Hash: hash,
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign secp256k1: %v", err)
				}

				if data.Signature == "" {
					t.Error("expected signature to be defined")
				}
				if data.Encoding != "hex" {
					t.Errorf("expected encoding to be hex, got %s", data.Encoding)
				}
			})
		}
	})

	t.Run("SignMessage", func(t *testing.T) {
		messages := []struct {
			name    string
			message string
		}{
			{name: "UTF8", message: "Hello, world!"},
			{name: "Hex", message: "0x48656c6c6f2c20776f726c6421"}, // "Hello, world!" in hex
		}

		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				for _, msg := range messages {
					t.Run(msg.name, func(t *testing.T) {
						data, err := client.Wallets.Ethereum.SignMessage(ctx, wallet.id,
							msg.message,
							WithAuthorizationContext(wallet.authCtx),
						)
						if err != nil {
							t.Fatalf("failed to sign message: %v", err)
						}

						if data.Signature == "" {
							t.Error("expected signature to be defined")
						}
						if data.Encoding != "hex" {
							t.Errorf("expected encoding to be hex, got %s", data.Encoding)
						}
					})
				}
			})
		}
	})

	t.Run("SignMessageBytes", func(t *testing.T) {
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Ethereum.SignMessageBytes(ctx, wallet.id,
					[]byte("Hello, world!"),
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign message bytes: %v", err)
				}

				if data.Signature == "" {
					t.Error("expected signature to be defined")
				}
				if data.Encoding != "hex" {
					t.Errorf("expected encoding to be hex, got %s", data.Encoding)
				}
			})
		}
	})

	t.Run("SignTypedData", func(t *testing.T) {
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Ethereum.SignTypedData(ctx, wallet.id,
					EthereumSignTypedDataRpcInputParam{
						Method: EthereumSignTypedDataRpcInputMethodEthSignTypedDataV4,
						Params: EthereumSignTypedDataRpcInputParamsParam{
							TypedData: EthereumSignTypedDataRpcInputParamsTypedDataParam{
								Domain: map[string]any{
									"name":              "Test",
									"version":           "1",
									"chainId":           1,
									"verifyingContract": "0x1234567890123456789012345678901234567890",
								},
								PrimaryType: "Message",
								Types: map[string][]EthereumSignTypedDataRpcInputParamsTypedDataTypeParam{
									"Message": {
										{Name: "content", Type: "string"},
									},
								},
								Message: map[string]any{
									"content": "Hello world",
								},
							},
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign typed data: %v", err)
				}

				if data.Signature == "" {
					t.Error("expected signature to be defined")
				}
				if data.Encoding != "hex" {
					t.Errorf("expected encoding to be hex, got %s", data.Encoding)
				}
			})
		}
	})

	t.Run("SignTransaction", func(t *testing.T) {
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Ethereum.SignTransaction(ctx, wallet.id,
					EthereumSignTransactionRpcInputParam{
						Method: EthereumSignTransactionRpcInputMethodEthSignTransaction,
						Params: EthereumSignTransactionRpcInputParamsParam{
							Transaction: EthereumSignTransactionRpcInputParamsTransactionParam{
								Type: 2,
								ChainID: EthereumSignTransactionRpcInputParamsTransactionChainIDUnionParam{
									OfInt: param.NewOpt[int64](1),
								},
								To: param.NewOpt("0x742d35Cc6634C0532925a3b8D1A8a9ff1e7a7A4C"),
								Value: EthereumSignTransactionRpcInputParamsTransactionValueUnionParam{
									OfString: param.NewOpt("0x1"),
								},
								GasLimit: EthereumSignTransactionRpcInputParamsTransactionGasLimitUnionParam{
									OfString: param.NewOpt("0x5208"),
								},
								Data: param.NewOpt("0x"),
							},
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to sign transaction: %v", err)
				}

				if data.SignedTransaction == "" {
					t.Error("expected signed_transaction to be defined")
				}
				if data.Encoding != "rlp" {
					t.Errorf("expected encoding to be rlp, got %s", data.Encoding)
				}
			})
		}
	})

	// SendTransaction is skipped to not waste funds. Logic is shared with signing transactions so safe to not frequently test.
	t.Run("SendTransaction", func(t *testing.T) {
		t.Skip("skipped to not waste funds")
		for _, wallet := range wallets {
			t.Run(wallet.name, func(t *testing.T) {
				data, err := client.Wallets.Ethereum.SendTransaction(ctx, wallet.id,
					EthereumSendTransactionRpcInputParam{
						Method: EthereumSendTransactionRpcInputMethodEthSendTransaction,
						Caip2:  "eip155:11155111", // Sepolia
						Params: EthereumSendTransactionRpcInputParamsParam{
							Transaction: EthereumSendTransactionRpcInputParamsTransactionParam{
								To: param.NewOpt("0x429c8e85D3A18F9F0a64a7A851777e24D591485C"),
								Value: EthereumSendTransactionRpcInputParamsTransactionValueUnionParam{
									OfString: param.NewOpt("0x1"), // 1 wei
								},
								ChainID: EthereumSendTransactionRpcInputParamsTransactionChainIDUnionParam{
									OfInt: param.NewOpt[int64](11155111), // Sepolia
								},
							},
						},
					},
					WithAuthorizationContext(wallet.authCtx),
				)
				if err != nil {
					t.Fatalf("failed to send transaction: %v", err)
				}

				if data.Hash == "" {
					t.Error("expected hash to be defined")
				}
				if data.Caip2 != "eip155:11155111" {
					t.Errorf("expected caip2 to be eip155:11155111, got %s", data.Caip2)
				}
			})
		}
	})
}
