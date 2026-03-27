package privyclient

type PrivyYieldService struct {
	// Directly embed the generated YieldService to expose all its methods through PrivyYieldService
	YieldService
	logger logger
}

// newPrivyYieldService creates a new wrapped yield service.
// This is unexported so only PrivyClient can construct it.
func newPrivyYieldService(service YieldService, logger logger) *PrivyYieldService {
	return &PrivyYieldService{
		YieldService: service,
		logger:       logger,
	}
}
