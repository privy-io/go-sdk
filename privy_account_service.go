package privyclient

type PrivyAccountService struct {
	// Directly embed the generated AccountService to expose all its methods through PrivyAccountService
	AccountService
	logger logger
}

// newPrivyAccountService creates a new wrapped account service.
// This is unexported so only PrivyClient can construct it.
func newPrivyAccountService(service AccountService, logger logger) *PrivyAccountService {
	return &PrivyAccountService{
		AccountService: service,
		logger:         logger,
	}
}
