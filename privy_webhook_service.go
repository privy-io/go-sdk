// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
package privyclient

type PrivyWebhookService struct {
	// Directly embed the generated WebhookService to expose all its methods through PrivyWebhookService
	WebhookService
}

// newPrivyWebhookService creates a new wrapped webhook service.
// This is unexported so only PrivyClient can construct it.
func newPrivyWebhookService(service WebhookService) *PrivyWebhookService {
	return &PrivyWebhookService{
		WebhookService: service,
	}
}
