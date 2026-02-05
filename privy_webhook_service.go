package privyclient

type PrivyWebhookService struct {
	// Directly embed the generated WebhookService to expose all its methods through PrivyWebhookService
	WebhookService
	logger Logger
}

// newPrivyWebhookService creates a new wrapped webhook service.
// This is unexported so only PrivyClient can construct it.
func newPrivyWebhookService(service WebhookService, logger Logger) *PrivyWebhookService {
	return &PrivyWebhookService{
		WebhookService: service,
		logger:         logger,
	}
}
