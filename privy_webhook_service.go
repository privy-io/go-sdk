package privyclient

import (
	"errors"
	"net/http"

	svix "github.com/svix/svix-webhooks/go"
)

// PrivyWebhookService wraps the generated WebhookService with signature verification.
type PrivyWebhookService struct {
	WebhookService
	signingSecret string
	logger        logger
}

// newPrivyWebhookService creates a new wrapped webhook service.
func newPrivyWebhookService(service WebhookService, signingSecret string, logger logger) *PrivyWebhookService {
	return &PrivyWebhookService{
		WebhookService: service,
		signingSecret:  signingSecret,
		logger:         logger,
	}
}

// VerifyInput contains the parameters for webhook verification.
type VerifyInput struct {
	// Payload is the raw request body bytes. Must not be modified.
	Payload []byte

	// Headers are the HTTP request headers. Must include svix-id, svix-timestamp, svix-signature.
	Headers http.Header

	// SigningSecret is an optional per-call override for the webhook signing secret.
	// Falls back to the client-level secret if empty.
	SigningSecret string
}

// Verify checks the svix webhook signature and returns the typed event payload.
// Returns an InvalidWebhookError if verification fails.
func (s *PrivyWebhookService) Verify(input VerifyInput) (*UnsafeUnwrapWebhookEventUnion, error) {
	secret := input.SigningSecret
	if secret == "" {
		secret = s.signingSecret
	}
	if secret == "" {
		return nil, &InvalidWebhookError{Err: errors.New("webhook signing secret is required: pass it to PrivyClientOptions or VerifyInput")}
	}

	wh, err := svix.NewWebhook(secret)
	if err != nil {
		return nil, &InvalidWebhookError{Err: err}
	}

	if err := wh.Verify(input.Payload, input.Headers); err != nil {
		return nil, &InvalidWebhookError{Err: err}
	}

	return s.UnsafeUnwrap(input.Payload)
}

// InvalidWebhookError is returned when webhook signature verification fails.
type InvalidWebhookError struct {
	Err error
}

func (e *InvalidWebhookError) Error() string {
	return "invalid webhook: " + e.Err.Error()
}

func (e *InvalidWebhookError) Unwrap() error {
	return e.Err
}
