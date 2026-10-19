package e2e_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	. "github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/authorization"
)

// requestExpiryHash is a 32-byte hash (keccak256 of "hello") used for raw sign requests.
const requestExpiryHash = "0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8"

// Rejections must come from the request expiry check, not from another validation error.
var expiryErrorPattern = regexp.MustCompile(`(?i)expir`)

// intentExpiryToleranceMs absorbs clock skew and request latency in intent expiry checks.
const intentExpiryToleranceMs = int64(5 * 60 * 1000)

func requireExpiryError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected the request to be rejected, got no error")
	}
	if !expiryErrorPattern.MatchString(err.Error()) {
		t.Fatalf("expected a request expiry error, got: %v", err)
	}
}

// newKeyOwnedWallet creates a wallet owned by a fresh P256 key. Raw sign does
// not support ethereum or solana wallets on staging, so raw sign tests use tron.
func newKeyOwnedWallet(t *testing.T, ctx context.Context, client *PrivyClient, chainType WalletChainType) (string, *authorization.AuthorizationContext) {
	t.Helper()
	keyPair, err := authorization.GenerateP256KeyPair()
	if err != nil {
		t.Fatalf("failed to generate P256 key pair: %v", err)
	}
	wallet, err := client.Wallets.New(ctx, WalletNewParams{
		ChainType: chainType,
		Owner: OwnerInputUnion{
			OfOwnerInputPublicKey: &OwnerInputPublicKey{PublicKey: keyPair.PublicKey},
		},
	})
	if err != nil {
		t.Fatalf("failed to create %s wallet: %v", chainType, err)
	}
	return wallet.ID, &authorization.AuthorizationContext{PrivateKeys: []string{keyPair.PrivateKey}}
}

func rawSignParams() WalletRawSignParams {
	return WalletRawSignParams{
		RawSignInput: RawSignInput{
			Params: RawSignInputParamsUnion{
				OfRawSignHashs: &RawSignHashParams{Hash: requestExpiryHash},
			},
		},
	}
}

// intentExpiresAtMs converts an intent expires_at Unix timestamp to milliseconds,
// accepting either seconds or milliseconds.
func intentExpiresAtMs(expiresAt float64) int64 {
	if expiresAt < 1e12 {
		return int64(expiresAt * 1000)
	}
	return int64(expiresAt)
}

func expectDefaultIntentExpiry(t *testing.T, intent BaseIntentResponse, before, after int64) {
	t.Helper()
	expiresAt := intentExpiresAtMs(intent.ExpiresAt)
	if expiresAt < before+DefaultIntentRequestExpiryMs-intentExpiryToleranceMs ||
		expiresAt > after+DefaultIntentRequestExpiryMs+intentExpiryToleranceMs {
		t.Errorf("expected expires_at about 72h from now, got %d (now %d)", expiresAt, after)
	}
}

func expectCustomIntentExpiry(t *testing.T, intent BaseIntentResponse, expiry int64) {
	t.Helper()
	if !intent.CustomExpiry {
		t.Error("expected custom_expiry to be true")
	}
	expiresAt := intentExpiresAtMs(intent.ExpiresAt)
	if diff := expiresAt - expiry; diff > intentExpiryToleranceMs || diff < -intentExpiryToleranceMs {
		t.Errorf("expected expires_at about %d, got %d", expiry, expiresAt)
	}
}

func TestRequestExpiry_SignedWrappers(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()
	walletID, authCtx := newKeyOwnedWallet(t, ctx, client, WalletChainTypeTron)

	t.Run("RawSign", func(t *testing.T) {
		t.Run("AcceptsDefault", func(t *testing.T) {
			result, err := client.Wallets.RawSign(ctx, walletID, rawSignParams(), WithAuthorizationContext(authCtx))
			if err != nil {
				t.Fatalf("failed to raw sign: %v", err)
			}
			if result.Data.Signature == "" {
				t.Error("expected signature to be non-empty")
			}
		})

		t.Run("AcceptsPerCall", func(t *testing.T) {
			result, err := client.Wallets.RawSign(ctx, walletID, rawSignParams(),
				WithAuthorizationContext(authCtx),
				WithRequestExpiry(RequestExpiry(20*60*1000)),
			)
			if err != nil {
				t.Fatalf("failed to raw sign: %v", err)
			}
			if result.Data.Signature == "" {
				t.Error("expected signature to be non-empty")
			}
		})

		t.Run("RejectsExpired", func(t *testing.T) {
			_, err := client.Wallets.RawSign(ctx, walletID, rawSignParams(),
				WithAuthorizationContext(authCtx),
				WithRequestExpiry(RequestExpiry(-60*1000)),
			)
			requireExpiryError(t, err)
		})

		t.Run("RejectsZero", func(t *testing.T) {
			_, err := client.Wallets.RawSign(ctx, walletID, rawSignParams(),
				WithAuthorizationContext(authCtx),
				WithRequestExpiry(0),
			)
			requireExpiryError(t, err)
		})
	})

	t.Run("Update", func(t *testing.T) {
		update := func(name string, opts ...RequestOption) (*Wallet, error) {
			return client.Wallets.Update(ctx, walletID, WalletUpdateParams{
				WalletUpdateRequestBody: WalletUpdateRequestBody{DisplayName: String(name)},
			}, append([]RequestOption{WithAuthorizationContext(authCtx)}, opts...)...)
		}

		t.Run("AcceptsDefault", func(t *testing.T) {
			wallet, err := update("request expiry default")
			if err != nil {
				t.Fatalf("failed to update wallet: %v", err)
			}
			if wallet.ID != walletID {
				t.Errorf("expected wallet ID %s, got %s", walletID, wallet.ID)
			}
		})

		t.Run("AcceptsPerCall", func(t *testing.T) {
			wallet, err := update("request expiry per-call", WithRequestExpiry(RequestExpiry(20*60*1000)))
			if err != nil {
				t.Fatalf("failed to update wallet: %v", err)
			}
			if wallet.ID != walletID {
				t.Errorf("expected wallet ID %s, got %s", walletID, wallet.ID)
			}
		})

		t.Run("RejectsExpired", func(t *testing.T) {
			_, err := update("request expiry expired", WithRequestExpiry(RequestExpiry(-60*1000)))
			requireExpiryError(t, err)
		})
	})

	t.Run("Disabled", func(t *testing.T) {
		disabled := newTestClientWithOptions(t, PrivyClientOptions{
			RequestExpiry: PrivyRequestExpiryOptions{Disabled: true},
		})

		t.Run("AcceptsNoExpiry", func(t *testing.T) {
			result, err := disabled.Wallets.RawSign(ctx, walletID, rawSignParams(), WithAuthorizationContext(authCtx))
			if err != nil {
				t.Fatalf("failed to raw sign: %v", err)
			}
			if result.Data.Signature == "" {
				t.Error("expected signature to be non-empty")
			}
		})

		// Without a client default, the request only fails if the per-call 0 is sent as the header.
		t.Run("SendsPerCallZero", func(t *testing.T) {
			_, err := disabled.Wallets.RawSign(ctx, walletID, rawSignParams(),
				WithAuthorizationContext(authCtx),
				WithRequestExpiry(0),
			)
			requireExpiryError(t, err)
		})
	})
}

func TestRequestExpiry_Intents(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Each intent gets its own wallet so no two intents are pending on the same wallet.
	updateWallet := func(t *testing.T, c *PrivyClient, opts ...RequestOption) *WalletIntentResponse {
		t.Helper()
		walletID, _ := newKeyOwnedWallet(t, ctx, client, WalletChainTypeTron)
		newOwner, err := authorization.GenerateP256KeyPair()
		if err != nil {
			t.Fatalf("failed to generate P256 key pair: %v", err)
		}
		intent, err := c.Intents.UpdateWallet(ctx, walletID, IntentUpdateWalletParams{
			WalletUpdateRequestBody: WalletUpdateRequestBody{
				Owner: OwnerInputUnion{
					OfOwnerInputPublicKey: &OwnerInputPublicKey{PublicKey: newOwner.PublicKey},
				},
			},
		}, opts...)
		if err != nil {
			t.Fatalf("failed to create update wallet intent: %v", err)
		}
		return intent
	}

	t.Run("DefaultsTo72Hours", func(t *testing.T) {
		before := time.Now().UnixMilli()
		intent := updateWallet(t, client)
		after := time.Now().UnixMilli()
		expectDefaultIntentExpiry(t, intent.BaseIntentResponse, before, after)
	})

	t.Run("UsesPerCall", func(t *testing.T) {
		expiry := RequestExpiry(60 * 60 * 1000)
		intent := updateWallet(t, client, WithRequestExpiry(expiry))
		expectCustomIntentExpiry(t, intent.BaseIntentResponse, expiry)
	})

	t.Run("SendsPerCallWhenDisabled", func(t *testing.T) {
		disabled := newTestClientWithOptions(t, PrivyClientOptions{
			RequestExpiry: PrivyRequestExpiryOptions{Disabled: true},
		})
		expiry := RequestExpiry(60 * 60 * 1000)
		intent := updateWallet(t, disabled, WithRequestExpiry(expiry))
		expectCustomIntentExpiry(t, intent.BaseIntentResponse, expiry)
	})
}

func TestRequestExpiry_Transfer(t *testing.T) {
	// The server checks the balance before the expiry, so even the expired case fails this way.
	t.Skip("staging wallets are unfunded, so transfers fail with insufficient balance")
	client := newTestClient(t)
	ctx := context.Background()
	walletID, authCtx := newKeyOwnedWallet(t, ctx, client, WalletChainTypeEthereum)

	transfer := func(opts ...RequestOption) (*TransferActionResponse, error) {
		return client.Wallets.Transfer(ctx, walletID, WalletTransferParams{
			TransferRequestBody: TransferRequestBody{
				Source: TokenTransferSourceUnion{
					OfNamedTokenTransferSource: &NamedTokenTransferSource{Asset: "usdc", Chain: "base", Amount: String("0.01")},
				},
				Destination: TokenTransferDestination{Address: "0xB00F0759DbeeF5E543Cc3E3B07A6442F5f3928a2"},
			},
		}, append([]RequestOption{WithAuthorizationContext(authCtx)}, opts...)...)
	}

	t.Run("AcceptsPerCall", func(t *testing.T) {
		if _, err := transfer(WithRequestExpiry(RequestExpiry(20 * 60 * 1000))); err != nil {
			t.Fatalf("failed to transfer: %v", err)
		}
	})

	t.Run("RejectsExpired", func(t *testing.T) {
		_, err := transfer(WithRequestExpiry(RequestExpiry(-60 * 1000)))
		requireExpiryError(t, err)
	})
}

func TestRequestExpiry_Automations(t *testing.T) {
	t.Skip("attaching needs an existing automation, and PrivyClient does not expose automation creation")
	client := newTestClient(t)
	ctx := context.Background()
	walletID, authCtx := newKeyOwnedWallet(t, ctx, client, WalletChainTypeEthereum)
	automationIDs := []string{"automation-id"}

	t.Run("Attach", func(t *testing.T) {
		_, err := client.Wallets.AttachAutomations(ctx, walletID, WalletAttachAutomationsParams{
			AttachWalletAutomationRequestBody: AttachWalletAutomationRequestBody{AutomationIDs: automationIDs},
		}, WithAuthorizationContext(authCtx), WithRequestExpiry(RequestExpiry(20*60*1000)))
		if err != nil {
			t.Fatalf("failed to attach automations: %v", err)
		}
	})

	t.Run("Detach", func(t *testing.T) {
		_, err := client.Wallets.DetachAutomations(ctx, walletID, WalletDetachAutomationsParams{
			DetachWalletAutomationRequestBody: DetachWalletAutomationRequestBody{AutomationIDs: automationIDs},
		}, WithAuthorizationContext(authCtx), WithRequestExpiry(RequestExpiry(20*60*1000)))
		if err != nil {
			t.Fatalf("failed to detach automations: %v", err)
		}
	})
}

func TestRequestExpiry_DepositAccountsAndPayouts(t *testing.T) {
	// Without these, staging rejects the requests before checking the expiry.
	t.Skip("deposit accounts need a deposit config and a wallet owned 1-of-1 by an authenticated user, and payouts need a fiat account; the staging test app has neither")
	client := newTestClient(t)
	ctx := context.Background()
	walletID, authCtx := newKeyOwnedWallet(t, ctx, client, WalletChainTypeEthereum)
	expired := WithRequestExpiry(RequestExpiry(-60 * 1000))

	t.Run("DepositAccountRejectsExpired", func(t *testing.T) {
		_, err := client.Wallets.DepositAccounts.Crypto.New(ctx, walletID, WalletDepositAccountCryptoNewParams{
			CreateCryptoDepositAccountRequestBody: CreateCryptoDepositAccountRequestBodyUnion{
				OfDepositConfig: &CreateCryptoDepositAccountWithConfigRequestBody{
					DepositConfigID: "deposit-config-id",
					Type:            CreateCryptoDepositAccountWithConfigRequestBodyTypeDepositConfig,
				},
			},
		}, WithAuthorizationContext(authCtx), expired)
		requireExpiryError(t, err)
	})

	t.Run("PayoutRejectsExpired", func(t *testing.T) {
		_, err := client.Wallets.Payout.Fiat.New(ctx, walletID, WalletPayoutFiatNewParams{
			CreatePayoutRequestBody: CreatePayoutRequestBody{
				Destination: PayoutDestination{FiatAccountID: "fiat-account-id"},
				Source:      PayoutSource{Amount: "1.00", Asset: "usdc", Chain: "base"},
			},
		}, WithAuthorizationContext(authCtx), expired)
		requireExpiryError(t, err)
	})
}

// Skipped like the Node SDK's swap and earn tests: they would move funds, and
// staging wallets are unfunded. They still type-check the per-call expiry calls.
func TestRequestExpiry_SwapsAndEarn(t *testing.T) {
	t.Skip("swaps and earn move funds, and staging wallets are unfunded")
	client := newTestClient(t)
	ctx := context.Background()
	walletID, authCtx := newKeyOwnedWallet(t, ctx, client, WalletChainTypeEthereum)
	perCall := WithRequestExpiry(RequestExpiry(20 * 60 * 1000))
	auth := WithAuthorizationContext(authCtx)
	source := SwapSource{AssetAddress: "native", Caip2: "eip155:8453"}
	const vaultID = "cm7oxq1el000e11o8iwp7d0d0"

	if _, err := client.Wallets.Swap.Quote(ctx, walletID, WalletSwapQuoteParams{
		SwapQuoteRequestBody: SwapQuoteRequestBody{
			BaseAmount:  "1000000000000000000",
			Source:      source,
			Destination: SwapQuoteDestination{AssetAddress: "0x0000000000000000000000000000000000000001"},
		},
	}, auth, perCall); err != nil {
		t.Errorf("failed to get swap quote: %v", err)
	}

	if _, err := client.Wallets.Swap.Execute(ctx, walletID, WalletSwapExecuteParams{
		SwapRequestBody: SwapRequestBody{
			BaseAmount:  "1000000000000000000",
			Source:      source,
			Destination: SwapDestination{AssetAddress: "0x0000000000000000000000000000000000000001"},
		},
	}, auth, perCall); err != nil {
		t.Errorf("failed to execute swap: %v", err)
	}

	if _, err := client.Wallets.Earn.Ethereum.Deposit(ctx, walletID, WalletEarnEthereumDepositParams{
		EarnDepositRequestBody: EarnDepositRequestBody{VaultID: vaultID, Amount: String("1.0")},
	}, auth, perCall); err != nil {
		t.Errorf("failed to deposit: %v", err)
	}

	if _, err := client.Wallets.Earn.Ethereum.Withdraw(ctx, walletID, WalletEarnEthereumWithdrawParams{
		EarnWithdrawRequestBody: EarnWithdrawRequestBody{VaultID: vaultID, Amount: String("1.0")},
	}, auth, perCall); err != nil {
		t.Errorf("failed to withdraw: %v", err)
	}

	if _, err := client.Wallets.Earn.Ethereum.Incentive.Claim(ctx, walletID, WalletEarnEthereumIncentiveClaimParams{
		EarnIncentiveClaimRequestBody: EarnIncentiveClaimRequestBody{Chain: "base"},
	}, auth, perCall); err != nil {
		t.Errorf("failed to claim incentives: %v", err)
	}
}
