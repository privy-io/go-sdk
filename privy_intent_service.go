package privyclient

// PrivyIntentService wraps the generated IntentService.
// Intents represent pending operations that require multi-party authorization
// before they can be executed.
type PrivyIntentService struct {
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
