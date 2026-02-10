package privyclient

import (
	"context"
	"encoding/base64"
	"fmt"
)

// PrivySolanaWalletService provides convenience methods for Solana wallet operations.
// Each method wraps PrivyWalletService.Rpc with the appropriate RPC input.
type PrivySolanaWalletService struct {
	walletService *PrivyWalletService
}

// newPrivySolanaWalletService creates a new Solana wallet service.
func newPrivySolanaWalletService(walletService *PrivyWalletService) *PrivySolanaWalletService {
	return &PrivySolanaWalletService{walletService: walletService}
}

// SignMessage calls signMessage with the given wallet.
// The message should be a base64-encoded string.
func (s *PrivySolanaWalletService) SignMessage(
	ctx context.Context,
	walletID string,
	message string,
	opts ...RpcOption,
) (*SolanaSignMessageRpcResponseData, error) {
	input := SolanaSignMessageRpcInputParam{
		Method: SolanaSignMessageRpcInputMethodSignMessage,
		Params: SolanaSignMessageRpcInputParamsParam{
			Message:  message,
			Encoding: "base64",
		},
	}

	params := WalletRpcParams{
		OfSignMessage: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "signMessage" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "signMessage", response.Method)
	}

	data := response.AsSignMessage().Data
	return &data, nil
}

// SignMessageBytes calls signMessage with raw bytes.
// The bytes are base64-encoded for transmission.
func (s *PrivySolanaWalletService) SignMessageBytes(
	ctx context.Context,
	walletID string,
	message []byte,
	opts ...RpcOption,
) (*SolanaSignMessageRpcResponseData, error) {
	return s.SignMessage(ctx, walletID, base64.StdEncoding.EncodeToString(message), opts...)
}

// SignTransaction calls signTransaction with the given wallet.
// The transaction should be a base64-encoded string.
func (s *PrivySolanaWalletService) SignTransaction(
	ctx context.Context,
	walletID string,
	transaction string,
	opts ...RpcOption,
) (*SolanaSignTransactionRpcResponseData, error) {
	input := SolanaSignTransactionRpcInputParam{
		Method: SolanaSignTransactionRpcInputMethodSignTransaction,
		Params: SolanaSignTransactionRpcInputParamsParam{
			Transaction: transaction,
			Encoding:    "base64",
		},
	}

	params := WalletRpcParams{
		OfSignTransaction: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "signTransaction" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "signTransaction", response.Method)
	}

	data := response.AsSignTransaction().Data
	return &data, nil
}

// SignTransactionBytes calls signTransaction with raw bytes.
// The bytes are base64-encoded for transmission.
func (s *PrivySolanaWalletService) SignTransactionBytes(
	ctx context.Context,
	walletID string,
	transaction []byte,
	opts ...RpcOption,
) (*SolanaSignTransactionRpcResponseData, error) {
	return s.SignTransaction(ctx, walletID, base64.StdEncoding.EncodeToString(transaction), opts...)
}
