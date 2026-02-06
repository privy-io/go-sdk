package privyclient

// PrivySolanaWalletService provides convenience methods for Solana wallet operations.
// Each method wraps PrivyWalletService.Rpc with the appropriate RPC input.
type PrivySolanaWalletService struct {
	walletService *PrivyWalletService
}

// newPrivySolanaWalletService creates a new Solana wallet service.
func newPrivySolanaWalletService(walletService *PrivyWalletService) *PrivySolanaWalletService {
	return &PrivySolanaWalletService{walletService: walletService}
}
