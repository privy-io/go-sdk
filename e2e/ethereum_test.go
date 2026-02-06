//go:build e2e

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
}
