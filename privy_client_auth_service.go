package privyclient

type PrivyClientAuthService struct {
	// Directly embed the generated ClientAuthService to expose all its methods through PrivyClientAuthService
	ClientAuthService
	logger logger
}

// newPrivyClientAuthService creates a new wrapped client auth service.
// This is unexported so only PrivyClient can construct it.
func newPrivyClientAuthService(service ClientAuthService, logger logger) *PrivyClientAuthService {
	return &PrivyClientAuthService{
		ClientAuthService: service,
		logger:            logger,
	}
}
