package privyclient

import (
	"context"
	"fmt"
)

// PrivyEthereumWalletService provides convenience methods for Ethereum wallet operations.
// Each method wraps PrivyWalletService.Rpc with the appropriate RPC input.
type PrivyEthereumWalletService struct {
	walletService *PrivyWalletService
}

// newPrivyEthereumWalletService creates a new Ethereum wallet service.
func newPrivyEthereumWalletService(walletService *PrivyWalletService) *PrivyEthereumWalletService {
	return &PrivyEthereumWalletService{walletService: walletService}
}

// Sign7702Authorization calls eth_sign7702authorization with the given wallet
func (s *PrivyEthereumWalletService) Sign7702Authorization(
	ctx context.Context,
	walletID string,
	input EthereumSign7702AuthorizationRpcInputParam,
	opts ...RpcOption,
) (*EthereumSign7702AuthorizationRpcResponseData, error) {
	params := WalletRpcParams{
		OfEthSign7702Authorization: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "eth_sign7702Authorization" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "eth_sign7702Authorization", response.Method)
	}

	data := response.AsEthSign7702Authorization().Data
	return &data, nil
}

// SignUserOperation calls eth_signUserOperation with the given wallet
func (s *PrivyEthereumWalletService) SignUserOperation(
	ctx context.Context,
	walletID string,
	input EthereumSignUserOperationRpcInputParam,
	opts ...RpcOption,
) (*EthereumSignUserOperationRpcResponseData, error) {
	params := WalletRpcParams{
		OfEthSignUserOperation: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "eth_signUserOperation" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "eth_signUserOperation", response.Method)
	}

	data := response.AsEthSignUserOperation().Data
	return &data, nil
}
