package privyclient

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
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

// SignSecp256k1 calls secp256k1_sign with the given wallet
func (s *PrivyEthereumWalletService) SignSecp256k1(
	ctx context.Context,
	walletID string,
	input EthereumSecp256k1SignRpcInputParam,
	opts ...RpcOption,
) (*EthereumSecp256k1SignRpcResponseData, error) {
	params := WalletRpcParams{
		OfSecp256k1Sign: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "secp256k1_sign" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "secp256k1_sign", response.Method)
	}

	data := response.AsSecp256k1Sign().Data
	return &data, nil
}

// SignMessage calls personal_sign with the given wallet.
// If the message starts with "0x", it is treated as hex-encoded data.
// Otherwise, it is treated as a UTF-8 string.
func (s *PrivyEthereumWalletService) SignMessage(
	ctx context.Context,
	walletID string,
	message string,
	opts ...RpcOption,
) (*EthereumPersonalSignRpcResponseData, error) {
	var msgContent string
	var encoding EthereumPersonalSignRpcInputParamsEncoding

	if strings.HasPrefix(message, "0x") {
		// The 0x prefix is removed as encoding: hex is sufficient
		msgContent = message[2:]
		encoding = EthereumPersonalSignRpcInputParamsEncodingHex
	} else {
		msgContent = message
		encoding = EthereumPersonalSignRpcInputParamsEncodingUtf8
	}

	return s.signMessage(ctx, walletID, msgContent, encoding, opts...)
}

// SignMessageBytes calls personal_sign with the given wallet using raw bytes.
// The bytes are hex-encoded for transmission.
func (s *PrivyEthereumWalletService) SignMessageBytes(
	ctx context.Context,
	walletID string,
	message []byte,
	opts ...RpcOption,
) (*EthereumPersonalSignRpcResponseData, error) {
	return s.signMessage(ctx, walletID, hex.EncodeToString(message), EthereumPersonalSignRpcInputParamsEncodingHex, opts...)
}

func (s *PrivyEthereumWalletService) signMessage(
	ctx context.Context,
	walletID string,
	message string,
	encoding EthereumPersonalSignRpcInputParamsEncoding,
	opts ...RpcOption,
) (*EthereumPersonalSignRpcResponseData, error) {
	input := EthereumPersonalSignRpcInputParam{
		Method: EthereumPersonalSignRpcInputMethodPersonalSign,
		Params: EthereumPersonalSignRpcInputParamsParam{
			Message:  message,
			Encoding: encoding,
		},
	}

	params := WalletRpcParams{
		OfPersonalSign: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "personal_sign" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "personal_sign", response.Method)
	}

	data := response.AsPersonalSign().Data
	return &data, nil
}

// SignTypedData calls eth_signTypedData_v4 with the given wallet
func (s *PrivyEthereumWalletService) SignTypedData(
	ctx context.Context,
	walletID string,
	input EthereumSignTypedDataRpcInputParam,
	opts ...RpcOption,
) (*EthereumSignTypedDataRpcResponseData, error) {
	params := WalletRpcParams{
		OfEthSignTypedDataV4: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "eth_signTypedData_v4" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "eth_signTypedData_v4", response.Method)
	}

	data := response.AsEthSignTypedDataV4().Data
	return &data, nil
}

// SignTransaction calls eth_signTransaction with the given wallet
func (s *PrivyEthereumWalletService) SignTransaction(
	ctx context.Context,
	walletID string,
	input EthereumSignTransactionRpcInputParam,
	opts ...RpcOption,
) (*EthereumSignTransactionRpcResponseData, error) {
	params := WalletRpcParams{
		OfEthSignTransaction: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "eth_signTransaction" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "eth_signTransaction", response.Method)
	}

	data := response.AsEthSignTransaction().Data
	return &data, nil
}

// SendTransaction calls eth_sendTransaction with the given wallet
func (s *PrivyEthereumWalletService) SendTransaction(
	ctx context.Context,
	walletID string,
	input EthereumSendTransactionRpcInputParam,
	opts ...RpcOption,
) (*EthereumSendTransactionRpcResponseData, error) {
	params := WalletRpcParams{
		OfEthSendTransaction: &input,
	}

	response, err := s.walletService.Rpc(ctx, walletID, params, opts...)
	if err != nil {
		return nil, err
	}

	if response.Method != "eth_sendTransaction" {
		return nil, fmt.Errorf("unexpected response method: expected %q, got %q", "eth_sendTransaction", response.Method)
	}

	data := response.AsEthSendTransaction().Data
	return &data, nil
}
