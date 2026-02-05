package privyclient

type PrivyAppService struct {
	// Directly embed the generated AppService to expose all its methods through PrivyAppService
	AppService
	logger Logger
}

// newPrivyAppService creates a new wrapped app service.
// This is unexported so only PrivyClient can construct it.
func newPrivyAppService(service AppService, logger Logger) *PrivyAppService {
	return &PrivyAppService{
		AppService: service,
		logger:     logger,
	}
}
