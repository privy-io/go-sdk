package privyclient

type PrivyWalletService struct {
	// Directly embed the generated WalletService to expose all its methods through PrivyWalletService
	WalletService
	logger Logger
}

// newPrivyWalletService creates a new wrapped wallet service.
// This is unexported so only PrivyClient can construct it.
func newPrivyWalletService(service WalletService, logger Logger) *PrivyWalletService {
	return &PrivyWalletService{
		WalletService: service,
		logger:        logger,
	}
}
