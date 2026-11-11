package privyclient_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	privyclient "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
)

// Wrapped clients must contain each identifier in one path segment.
func TestKeyQuorumIdentifierContainment(t *testing.T) {
	for _, id := range []string{"kq_legit", "kq?limit=100", "kq#fragment", "kq%2Fother"} {
		t.Run(id, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				want := "/v1/key_quorums/" + url.PathEscape(id)
				if r.URL.RequestURI() != want {
					t.Errorf("got %q, want %q", r.URL.RequestURI(), want)
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{}`)
			}))
			defer server.Close()
			wrapped := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{AppID: "test", AppSecret: "test", APIUrl: server.URL})
			for _, call := range []func() error{
				func() error { _, e := wrapped.KeyQuorums.Get(context.Background(), id); return e },
				func() error {
					_, e := wrapped.KeyQuorums.Update(context.Background(), id, privyclient.KeyQuorumUpdateParams{})
					return e
				},
				func() error {
					_, e := wrapped.KeyQuorums.Delete(context.Background(), id, privyclient.KeyQuorumDeleteParams{})
					return e
				},
			} {
				if err := call(); err != nil {
					t.Fatal(err)
				}
			}
			if calls != 3 {
				t.Fatalf("got %d requests, want 3", calls)
			}
		})
	}
}

func TestKeyQuorumDotSegmentsRejected(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++; fmt.Fprint(w, `{}`) }))
	defer server.Close()
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{AppID: "test", AppSecret: "test", APIUrl: server.URL})
	for _, id := range []string{".", "..", "../users/victim", "%2e%2e/users/victim"} {
		if _, err := client.KeyQuorums.Get(context.Background(), id); (id == "." || id == ".." || id == "../users/victim") && err == nil {
			t.Errorf("accepted traversal %q", id)
		}
	}
	// A literally percent-encoded identifier is escaped once, not decoded into traversal.
	if calls != 1 {
		t.Errorf("got %d requests, want only the percent-literal identifier", calls)
	}
}

func TestWalletSubresourceIdentifierContainment(t *testing.T) {
	id := "id?limit=100#fragment"
	var got []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = append(got, r.URL.RequestURI())
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{AppID: "test", AppSecret: "test", APIUrl: server.URL})
	client.Wallets.Actions.Get(context.Background(), id, privyclient.WalletActionGetParams{WalletID: "w"})
	client.Wallets.DepositAccounts.Crypto.Orders.Get(context.Background(), id, privyclient.WalletDepositAccountCryptoOrderGetParams{WalletID: "w"})
	escaped := url.PathEscape(id)
	want := []string{"/v1/wallets/w/actions/" + escaped, "/v1/wallets/w/deposit_accounts/crypto/orders/" + escaped}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

type keyQuorumURLSigner struct{ url string }

func (s *keyQuorumURLSigner) Sign(_ context.Context, payload []byte) (string, error) {
	var request struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(payload, &request); err != nil {
		return "", err
	}
	s.url = request.URL
	return "test-signature", nil
}

func TestKeyQuorumSignedURLMatchesWire(t *testing.T) {
	signer := &keyQuorumURLSigner{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if signer.url != "http://"+r.Host+r.URL.RequestURI() {
			t.Errorf("signed %q, sent %q", signer.url, "http://"+r.Host+r.URL.RequestURI())
		}
		if r.Header.Get("privy-authorization-signature") != "test-signature" {
			t.Error("missing authorization signature")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{AppID: "test", AppSecret: "test", APIUrl: server.URL})
	auth := privyclient.WithAuthorizationContext(&authorization.AuthorizationContext{Signers: []authorization.AuthorizationSigner{signer}})
	for _, id := range []string{"kq_legit", "quorum/child", "kq?query#fragment", "kq%2Fother"} {
		if _, err := client.KeyQuorums.Update(context.Background(), id, privyclient.KeyQuorumUpdateParams{}, auth); err != nil {
			t.Fatal(err)
		}
		if _, err := client.KeyQuorums.Delete(context.Background(), id, privyclient.KeyQuorumDeleteParams{}, auth); err != nil {
			t.Fatal(err)
		}
	}
}
