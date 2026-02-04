//go:build e2e

package e2e_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/packages/param"
)

func TestWallets_Rpc_EthSign7702Authorization(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	response, err := client.Wallets.Rpc(ctx, os.Getenv("OWNERLESS_ETHEREUM_WALLET_ID"), WalletRpcParams{
		OfEthSign7702Authorization: &EthereumSign7702AuthorizationRpcInputParam{
			Method: EthereumSign7702AuthorizationRpcInputMethodEthSign7702Authorization,
			Params: EthereumSign7702AuthorizationRpcInputParamsParam{
				ChainID: EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam{
					OfInt: param.NewOpt[int64](11155111), // Sepolia
				},
				Contract: "0x1234567890123456789012345678901234567890",
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to sign 7702 authorization: %v", err)
	}

	if response.Method != "eth_sign7702Authorization" {
		t.Errorf("expected method to be eth_sign7702Authorization, got %s", response.Method)
	}

	authResponse := response.AsEthSign7702Authorization()
	auth := authResponse.Data.Authorization
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
}

func TestWallets_Rpc_EthSignUserOperation(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	response, err := client.Wallets.Rpc(ctx, os.Getenv("OWNERLESS_ETHEREUM_WALLET_ID"), WalletRpcParams{
		OfEthSignUserOperation: &EthereumSignUserOperationRpcInputParam{
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
	})
	if err != nil {
		t.Fatalf("failed to sign user operation: %v", err)
	}

	if response.Method != "eth_signUserOperation" {
		t.Errorf("expected method to be eth_signUserOperation, got %s", response.Method)
	}

	userOpResponse := response.AsEthSignUserOperation()
	if userOpResponse.Data.Signature == "" {
		t.Error("expected signature to be defined")
	}
	if userOpResponse.Data.Encoding != "hex" {
		t.Errorf("expected encoding to be hex, got %s", userOpResponse.Data.Encoding)
	}
}

func TestWallets_Rpc_EthSign7702Authorization_UserOwned(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	walletID := os.Getenv("USER_OWNED_ETHEREUM_WALLET_ID")
	if walletID == "" {
		t.Fatal("USER_OWNED_ETHEREUM_WALLET_ID environment variable is required")
	}
	appID := os.Getenv("TEST_APP_ID")

	jwt := generateTestJWT(t)

	// Build RPC request body
	rpcBody := &EthereumSign7702AuthorizationRpcInputParam{
		Method: EthereumSign7702AuthorizationRpcInputMethodEthSign7702Authorization,
		Params: EthereumSign7702AuthorizationRpcInputParamsParam{
			ChainID: EthereumSign7702AuthorizationRpcInputParamsChainIDUnionParam{
				OfInt: param.NewOpt[int64](11155111), // Sepolia
			},
			Contract: "0x1234567890123456789012345678901234567890",
		},
	}

	// Build authorization signature
	input := authorization.WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     fmt.Sprintf("https://auth.staging.privy.io/v1/wallets/%s/rpc", walletID),
		Body:    rpcBody,
		Headers: map[string]string{"privy-app-id": appID},
	}
	payload, err := authorization.FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("failed to format request for signature: %v", err)
	}

	auth := authorization.AuthorizationContext{
		UserJwts: []string{jwt},
	}
	sigs, err := authorization.GenerateAuthorizationSignatures(ctx, auth, payload, client.JwtExchange)
	if err != nil {
		t.Fatalf("failed to generate authorization signatures: %v", err)
	}

	// Call RPC with authorization
	response, err := client.Wallets.Rpc(ctx, walletID, WalletRpcParams{
		OfEthSign7702Authorization:  rpcBody,
		PrivyAuthorizationSignature: param.NewOpt(strings.Join(sigs, ",")),
	})
	if err != nil {
		t.Fatalf("failed to sign 7702 authorization: %v", err)
	}

	if response.Method != "eth_sign7702Authorization" {
		t.Errorf("expected method to be eth_sign7702Authorization, got %s", response.Method)
	}

	authResponse := response.AsEthSign7702Authorization()
	authData := authResponse.Data.Authorization
	if authData.Contract == "" {
		t.Error("expected authorization contract to be defined")
	}
	if authData.R == "" {
		t.Error("expected authorization R value to be defined")
	}
	if authData.S == "" {
		t.Error("expected authorization S value to be defined")
	}
	if authData.YParity != 0 && authData.YParity != 1 {
		t.Errorf("expected y_parity to be 0 or 1, got %f", authData.YParity)
	}
}
