package privyclient

type PrivyIntentService struct {
	// Directly embed the generated IntentService to expose all its methods through PrivyIntentService
	IntentService
	logger logger
}

// newPrivyIntentService creates a new wrapped intent service.
// This is unexported so only PrivyClient can construct it.
func newPrivyIntentService(service IntentService, logger logger) *PrivyIntentService {
	return &PrivyIntentService{
		IntentService: service,
		logger:        logger,
	}
}
