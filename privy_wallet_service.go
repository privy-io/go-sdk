// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

// PrivyWalletService wraps the generated WalletService with embedding.
// All methods from WalletService are directly available on PrivyWalletService,
// allowing you to add custom functionality while maintaining full compatibility
// with the generated SDK.
//
// Since WalletService is embedded, you can call any WalletService method directly:
//
//	wallet, err := client.Wallets.Get(ctx, "wallet_id")
//	wallets, err := client.Wallets.List(ctx, params)
//
// You can also add custom methods to this service without affecting the generated code.
type PrivyWalletService struct {
	WalletService
}

// NewPrivyWalletService creates a new wrapped wallet service.
func NewPrivyWalletService(service WalletService) *PrivyWalletService {
	return &PrivyWalletService{
		WalletService: service,
	}
}

// Add your custom methods below this line.
// Example:
//
// func (w *PrivyWalletService) CustomMethod(ctx context.Context) error {
//     // Your custom logic here
//     return nil
// }
