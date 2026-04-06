package e2e_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
)

func TestMain(m *testing.M) {
	// Load .env.secret file from project root
	_ = godotenv.Load("../.env.secret")
	os.Exit(m.Run())
}

func newTestClient(t *testing.T) *PrivyClient {
	t.Helper()
	appId := os.Getenv("TEST_APP_ID")
	appSecret := os.Getenv("TEST_APP_SECRET")
	if appSecret == "" {
		t.Fatal("TEST_APP_SECRET environment variable is required")
	}
	opts := PrivyClientOptions{
		AppID:     appId,
		AppSecret: appSecret,
		APIUrl:    "https://api.staging.privy.io",
	}
	if os.Getenv("TEST_DEBUG_LOGS") != "" {
		opts.LogLevel = LogLevelVerbose
	}
	client := NewPrivyClient(opts)
	return client
}

// testWalletResources holds shared resources created for wallet tests.
type testWalletResources struct {
	client        *PrivyClient
	ctx           context.Context
	p256KeyPair   authorization.P256KeyPair
	quorumKeyPair authorization.P256KeyPair
	userID        string
	customUserID  string
	quorumID      string
}

// testWallet holds a wallet's details and authorization context for testing.
type testWallet struct {
	name      string
	id        string
	address   string
	publicKey string
	authCtx   *authorization.AuthorizationContext
}

// setupTestWalletResources creates shared test resources on demand:
// a P256 key pair, a quorum key pair, a user, and a key quorum.
// A cleanup function is registered to delete the user.
func setupTestWalletResources(t *testing.T, client *PrivyClient) *testWalletResources {
	t.Helper()

	ctx := context.Background()

	// Generate P256 key pair for key-owned wallets
	p256KeyPair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate P256 key pair: %v", err)
	}

	// Generate separate P256 key pair for key quorum
	quorumKeyPair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate quorum P256 key pair: %v", err)
	}

	// Create user with custom auth
	customUserID := fmt.Sprintf("test-user-%d", time.Now().UnixNano())
	user, err := client.Users.New(ctx, UserNewParams{
		LinkedAccounts: []LinkedAccountInputUnion{
			{OfCustomAuth: &LinkedAccountCustomJwtInput{CustomUserID: customUserID}},
		},
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Register user cleanup
	t.Cleanup(func() {
		if err := client.Users.Delete(ctx, user.ID); err != nil {
			t.Logf("warning: failed to delete user %s: %v", user.ID, err)
		}
	})

	// Create key quorum with quorum key + user
	quorum, err := client.KeyQuorums.New(ctx, KeyQuorumNewParams{
		KeyQuorumCreateRequestBody: KeyQuorumCreateRequestBody{
			PublicKeys: []string{quorumKeyPair.PublicKey},
			UserIDs:    []string{user.ID},
		},
	})
	if err != nil {
		t.Fatalf("failed to create key quorum: %v", err)
	}

	return &testWalletResources{
		client:        client,
		ctx:           ctx,
		p256KeyPair:   p256KeyPair,
		quorumKeyPair: quorumKeyPair,
		userID:        user.ID,
		customUserID:  customUserID,
		quorumID:      quorum.ID,
	}
}

// createTestWallets creates 4 wallets (ownerless, key-owned, user-owned, quorum-owned)
// for the given chain type, each with the correct authorization context.
func (r *testWalletResources) createTestWallets(t *testing.T, chainType WalletChainType) []testWallet {
	t.Helper()

	jwt := generateTestJWT(t, r.customUserID)

	configs := []struct {
		name    string
		params  WalletNewParams
		authCtx *authorization.AuthorizationContext
	}{
		{
			name: "Ownerless",
			params: WalletNewParams{
				ChainType: chainType,
			},
			authCtx: nil,
		},
		{
			name: "KeyOwned",
			params: WalletNewParams{
				ChainType: chainType,
				Owner: OwnerInputUnionParam{
					OfOwnerInputPublicKey: &OwnerInputPublicKeyParam{PublicKey: r.p256KeyPair.PublicKey},
				},
			},
			authCtx: &authorization.AuthorizationContext{PrivateKeys: []string{r.p256KeyPair.PrivateKey}},
		},
		{
			name: "UserOwned",
			params: WalletNewParams{
				ChainType: chainType,
				Owner: OwnerInputUnionParam{
					OfOwnerInputUser: &OwnerInputUserParam{UserID: r.userID},
				},
			},
			authCtx: &authorization.AuthorizationContext{UserJwts: []string{jwt}},
		},
		{
			name: "KeyAndUserQuorumOwned",
			params: WalletNewParams{
				ChainType: chainType,
				OwnerID:   String(r.quorumID),
			},
			authCtx: &authorization.AuthorizationContext{
				PrivateKeys: []string{r.quorumKeyPair.PrivateKey},
				UserJwts:    []string{jwt},
			},
		},
	}

	wallets := make([]testWallet, 0, len(configs))
	for _, cfg := range configs {
		wallet, err := r.client.Wallets.New(r.ctx, cfg.params)
		if err != nil {
			t.Fatalf("failed to create %s %s wallet: %v", cfg.name, chainType, err)
		}

		wallets = append(wallets, testWallet{
			name:      cfg.name,
			id:        wallet.ID,
			address:   wallet.Address,
			publicKey: wallet.PublicKey,
			authCtx:   cfg.authCtx,
		})
	}

	return wallets
}
