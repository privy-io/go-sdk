package privyclient_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	privyclient "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/internal/hpke"
)

// authenticateRequest is a request received by the fake /v1/wallets/authenticate endpoint.
type authenticateRequest struct {
	userJwt          string
	mfaTokenHeaders  []string
	authorizationKey *ecdsa.PrivateKey
}

// newAuthenticateTestClient returns a PrivyClient whose API is a fake
// /v1/wallets/authenticate endpoint, which HPKE-encrypts a fresh P-256
// authorization key to each request's recipient public key, and the requests
// that endpoint receives.
func newAuthenticateTestClient(t *testing.T) (*privyclient.PrivyClient, *[]authenticateRequest) {
	t.Helper()
	var requests []authenticateRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/wallets/authenticate" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var body struct {
			UserJwt            string `json:"user_jwt"`
			RecipientPublicKey string `json:"recipient_public_key"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		spki, err := base64.StdEncoding.DecodeString(body.RecipientPublicKey)
		if err != nil {
			t.Errorf("failed to decode recipient public key: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		recipientPublicKey, err := x509.ParsePKIXPublicKey(spki)
		if err != nil {
			t.Errorf("failed to parse recipient public key: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		recipientEcdhKey, err := recipientPublicKey.(*ecdsa.PublicKey).ECDH()
		if err != nil {
			t.Errorf("failed to convert recipient public key: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		authorizationKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("failed to generate authorization key: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		pkcs8, err := x509.MarshalPKCS8PrivateKey(authorizationKey)
		if err != nil {
			t.Errorf("failed to marshal authorization key: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		encapsulatedKey, ciphertext, err := hpke.NewHpkeSender().Encrypt(recipientEcdhKey.Bytes(), []byte(base64.StdEncoding.EncodeToString(pkcs8)))
		if err != nil {
			t.Errorf("failed to encrypt authorization key: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		requests = append(requests, authenticateRequest{
			userJwt:          body.UserJwt,
			mfaTokenHeaders:  r.Header.Values("privy-mfa-token"),
			authorizationKey: authorizationKey,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"encrypted_authorization_key": map[string]any{
				"encryption_type":  "HPKE",
				"encapsulated_key": base64.StdEncoding.EncodeToString(encapsulatedKey),
				"ciphertext":       base64.StdEncoding.EncodeToString(ciphertext),
			},
			"expires_at": 0,
			"wallets":    []any{},
		})
	}))
	t.Cleanup(server.Close)

	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
		AppID:     "test-app-id",
		AppSecret: "test-app-secret",
		APIUrl:    server.URL,
	})
	return client, &requests
}

func TestJwtExchangeSendsMfaToken(t *testing.T) {
	client, requests := newAuthenticateTestClient(t)

	key, err := client.JwtExchange.ExchangeJwtForAuthorizationKeyWithMfaToken(context.Background(), "user-jwt", "mfa-token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(*requests) != 1 {
		t.Fatalf("expected 1 authenticate request, got %d", len(*requests))
	}
	request := (*requests)[0]
	if request.userJwt != "user-jwt" {
		t.Errorf("expected user_jwt %q, got %q", "user-jwt", request.userJwt)
	}
	if !slices.Equal(request.mfaTokenHeaders, []string{"mfa-token"}) {
		t.Errorf("expected privy-mfa-token header [mfa-token], got %v", request.mfaTokenHeaders)
	}

	pkcs8, err := x509.MarshalPKCS8PrivateKey(request.authorizationKey)
	if err != nil {
		t.Fatalf("failed to marshal authorization key: %v", err)
	}
	if key != base64.StdEncoding.EncodeToString(pkcs8) {
		t.Error("expected the decrypted authorization key")
	}
}

func TestJwtExchangeWithoutMfaTokenSendsNoHeader(t *testing.T) {
	client, requests := newAuthenticateTestClient(t)

	if _, err := client.JwtExchange.ExchangeJwtForAuthorizationKeyWithMfaToken(context.Background(), "user-jwt", ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := client.JwtExchange.ExchangeJwtForAuthorizationKey(context.Background(), "user-jwt"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(*requests) != 2 {
		t.Fatalf("expected 2 authenticate requests, got %d", len(*requests))
	}
	for i, request := range *requests {
		if len(request.mfaTokenHeaders) != 0 {
			t.Errorf("request %d: expected no privy-mfa-token header, got %v", i, request.mfaTokenHeaders)
		}
	}
}

func TestGenerateAuthorizationSignaturesForRequestSendsUserAuthMfaToken(t *testing.T) {
	client, requests := newAuthenticateTestClient(t)

	input := authorization.WalletApiRequestSignatureInput{
		Version: 1,
		Method:  "POST",
		URL:     "https://api.privy.io/v1/wallets/wallet-id/rpc",
		Body:    map[string]any{"method": "personal_sign"},
		Headers: map[string]string{"privy-app-id": "test-app-id"},
	}
	signatures, err := client.GenerateAuthorizationSignaturesForRequest(
		context.Background(),
		authorization.AuthorizationContext{
			UserAuth: []authorization.UserAuth{{UserJwt: "user-jwt", MfaToken: "mfa-token"}},
		},
		input,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(*requests) != 1 {
		t.Fatalf("expected 1 authenticate request, got %d", len(*requests))
	}
	request := (*requests)[0]
	if !slices.Equal(request.mfaTokenHeaders, []string{"mfa-token"}) {
		t.Errorf("expected privy-mfa-token header [mfa-token], got %v", request.mfaTokenHeaders)
	}

	if len(signatures) != 1 {
		t.Fatalf("expected 1 signature, got %d", len(signatures))
	}
	payload, err := authorization.FormatRequestForAuthorizationSignature(input)
	if err != nil {
		t.Fatalf("failed to format request: %v", err)
	}
	hash := sha256.Sum256(payload)
	signature, err := base64.StdEncoding.DecodeString(signatures[0])
	if err != nil {
		t.Fatalf("failed to decode signature: %v", err)
	}
	if !ecdsa.VerifyASN1(&request.authorizationKey.PublicKey, hash[:], signature) {
		t.Error("expected a signature from the exchanged authorization key")
	}
}
