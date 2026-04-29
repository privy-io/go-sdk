package privyclient_test

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	svix "github.com/svix/svix-webhooks/go"
	privyclient "github.com/privy-io/go-sdk"
)

func TestPrivyWebhookService_Verify(t *testing.T) {
	secret := generateWebhookTestSecret(t)

	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:                "test-app-id",
		AppSecret:            "test-app-secret",
		WebhookSigningSecret: secret,
	})

	t.Run("verifies valid webhook and returns typed event", func(t *testing.T) {
		payload := []byte(`{"type":"user.created","user":{"id":"did:privy:test123"}}`)
		headers := signWebhookPayload(t, secret, payload)

		event, err := client.Webhooks.Verify(privyclient.VerifyInput{
			Payload: payload,
			Headers: headers,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := event.AsUserCreated()
		if got.User.ID != "did:privy:test123" {
			t.Errorf("got user ID %q, want %q", got.User.ID, "did:privy:test123")
		}
	})

	t.Run("allows per-call signing secret override", func(t *testing.T) {
		perCallSecret := generateWebhookTestSecret(t)
		payload := []byte(`{"type":"user.created","user":{"id":"did:privy:override"}}`)
		headers := signWebhookPayload(t, perCallSecret, payload)

		event, err := client.Webhooks.Verify(privyclient.VerifyInput{
			Payload:      payload,
			Headers:      headers,
			SigningSecret: perCallSecret,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := event.AsUserCreated()
		if got.User.ID != "did:privy:override" {
			t.Errorf("got user ID %q, want %q", got.User.ID, "did:privy:override")
		}
	})

	t.Run("returns error on invalid signature", func(t *testing.T) {
		payload := []byte(`{"type":"user.created"}`)
		headers := http.Header{}
		headers.Set("svix-id", "msg_fake")
		headers.Set("svix-timestamp", fmt.Sprintf("%d", time.Now().Unix()))
		headers.Set("svix-signature", "v1,aW52YWxpZFNpZ25hdHVyZQ==")

		_, err := client.Webhooks.Verify(privyclient.VerifyInput{
			Payload: payload,
			Headers: headers,
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		var webhookErr *privyclient.InvalidWebhookError
		if !errors.As(err, &webhookErr) {
			t.Errorf("expected InvalidWebhookError, got %T", err)
		}
	})

	t.Run("returns error when no signing secret provided", func(t *testing.T) {
		noSecretClient := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
			AppID:     "test-app-id",
			AppSecret: "test-app-secret",
		})

		_, err := noSecretClient.Webhooks.Verify(privyclient.VerifyInput{
			Payload: []byte(`{}`),
			Headers: http.Header{},
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		var webhookErr *privyclient.InvalidWebhookError
		if !errors.As(err, &webhookErr) {
			t.Errorf("expected InvalidWebhookError, got %T", err)
		}
	})
}

func generateWebhookTestSecret(t *testing.T) string {
	t.Helper()
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		t.Fatalf("failed to generate secret: %v", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func signWebhookPayload(t *testing.T, secret string, payload []byte) http.Header {
	t.Helper()
	wh, err := svix.NewWebhook(secret)
	if err != nil {
		t.Fatalf("failed to create svix webhook: %v", err)
	}

	msgID := "msg_test123"
	timestamp := time.Now()

	signature, err := wh.Sign(msgID, timestamp, payload)
	if err != nil {
		t.Fatalf("failed to sign: %v", err)
	}

	headers := http.Header{}
	headers.Set("svix-id", msgID)
	headers.Set("svix-timestamp", fmt.Sprintf("%d", timestamp.Unix()))
	headers.Set("svix-signature", signature)
	return headers
}
