package privyclient

type PrivyPolicyService struct {
	// Directly embed the generated PolicyService to expose all its methods through PrivyPolicyService
	PolicyService
	logger logger
}

// newPrivyPolicyService creates a new wrapped policy service.
// This is unexported so only PrivyClient can construct it.
func newPrivyPolicyService(service PolicyService, logger logger) *PrivyPolicyService {
	return &PrivyPolicyService{
		PolicyService: service,
		logger:        logger,
	}
}
