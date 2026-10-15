package privyclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	privyclient "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/option"
)

func TestResourcePathsRejectDotSegments(t *testing.T) {
	for _, id := range []string{".", "..", "../policies/target", "%2e%2e/policies/target", ".%2E/policies/target", "%2e/policies/target", "safe/../target", "safe%2f..%2ftarget"} {
		t.Run(id, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++; w.Write([]byte(`{}`)) }))
			defer server.Close()
			client := privyclient.NewClient(option.WithBaseURL(server.URL), option.WithAppID("test"), option.WithAppSecret("test"))
			operations := map[string]func() error{
				"organization_get":    func() error { _, err := client.Organizations.Get(context.Background(), id); return err },
				"organization_delete": func() error { return client.Organizations.Delete(context.Background(), id) },
				"key_quorum_get":      func() error { _, err := client.KeyQuorums.Get(context.Background(), id); return err },
			}
			// Escaped resource methods preserve literal percent sequences in IDs.
			// Literal dot segments must still be rejected after path escaping.
			if !strings.Contains(id, "%") {
				operations["wallet_get"] = func() error {
					_, err := client.Wallets.Get(context.Background(), id, privyclient.WalletGetParams{})
					return err
				}
			}
			for name, operation := range operations {
				t.Run(name, func(t *testing.T) {
					if err := operation(); err == nil || !strings.Contains(err.Error(), "dot segments") {
						t.Fatalf("expected dot-segment rejection, got %v", err)
					}
				})
			}
			if calls != 0 {
				t.Fatalf("sent %d requests for rejected identifiers", calls)
			}
		})
	}
}

func TestResourcePathsPreserveValidIdentifiers(t *testing.T) {
	for _, id := range []string{"organization_id", "name.with.dots", ".well-known", "...", "did:privy:target"} {
		t.Run(id, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				if r.URL.Path != "/prefix/v1/organizations/"+id {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}
				if r.URL.Query().Get("filter") != "../allowed-in-query" {
					t.Errorf("query changed: %s", r.URL.RawQuery)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{}`))
			}))
			defer server.Close()
			client := privyclient.NewClient(option.WithBaseURL(server.URL + "/prefix/"))
			_, err := client.Organizations.Get(context.Background(), id, option.WithQuery("filter", "../allowed-in-query"))
			if err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Fatalf("expected one request, got %d", calls)
			}
		})
	}
}
