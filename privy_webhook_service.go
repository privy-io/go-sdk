package privyclient

type PrivyWebhookService struct {
	// Directly embed the generated WebhookService to expose all its methods through PrivyWebhookService
	WebhookService
	logger logger
}

// newPrivyWebhookService creates a new wrapped webhook service.
// This is unexported so only PrivyClient can construct it.
func newPrivyWebhookService(service WebhookService, logger logger) *PrivyWebhookService {
	return &PrivyWebhookService{
		WebhookService: service,
		logger:         logger,
	}
}
