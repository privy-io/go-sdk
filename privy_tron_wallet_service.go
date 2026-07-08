package privyclient

import (
	"context"
	"fmt"

	"github.com/privy-io/go-sdk/packages/param"
)

// PrivyTronWalletService provides convenience methods for Tron wallet operations.
// Each method wraps PrivyWalletService.Rpc with the appropriate RPC input.
type PrivyTronWalletService struct {
	walletService *PrivyWalletService
}

// newPrivyTronWalletService creates a new Tron wallet service.
func newPrivyTronWalletService(walletService *PrivyWalletService) *PrivyTronWalletService {
	return &PrivyTronWalletService{walletService: walletService}
}

// SignTransaction calls tron_signTransaction with the given wallet.
// Returns TronSignTransactionRpcResponseData. The caller is responsible for broadcasting.
func (s *PrivyTronWalletService) SignTransaction(
	ctx context.Context,
	walletID string,
	params TronSignTransactionRpcInputParams,
	opts ...RequestOption,
) (*TronSignTransactionRpcResponseData, error) {
	input := TronSignTransactionRpcInput{
		Method: TronSignTransactionRpcInputMethodTronSignTransaction,
		Params: params,
	}
	rpcParams := WalletRpcParams{
		WalletRpcRequestBody: WalletRpcRequestBodyUnion{
			OfTronSignTransaction: &input,
		},
	}
	response, err := s.walletService.Rpc(ctx, walletID, rpcParams, opts...)
	if err != nil {
		return nil, err
	}
	if response.Method != "tron_signTransaction" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "tron_signTransaction", response.Method)
	}
	data := response.AsTronSignTransaction().Data
	return &data, nil
}

// SendTransaction calls tron_sendTransaction with the given wallet.
// Signs and broadcasts the transaction. caip2 identifies the target network
// (e.g. "tron:0xcd8690dc" for Nile testnet).
func (s *PrivyTronWalletService) SendTransaction(
	ctx context.Context,
	walletID string,
	caip2 string,
	params TronSendTransactionRpcInputParams,
	opts ...RequestOption,
) (*TronSendTransactionRpcResponseData, error) {
	input := TronSendTransactionRpcInput{
		Method: TronSendTransactionRpcInputMethodTronSendTransaction,
		Caip2:  param.NewOpt(caip2),
		Params: params,
	}
	rpcParams := WalletRpcParams{
		WalletRpcRequestBody: WalletRpcRequestBodyUnion{
			OfTronSendTransaction: &input,
		},
	}
	response, err := s.walletService.Rpc(ctx, walletID, rpcParams, opts...)
	if err != nil {
		return nil, err
	}
	if response.Method != "tron_sendTransaction" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "tron_sendTransaction", response.Method)
	}
	data := response.AsTronSendTransaction().Data
	return &data, nil
}
