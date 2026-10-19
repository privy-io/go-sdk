package privyclient_test

import (
	"reflect"
	"testing"

	privyclient "github.com/privy-io/go-sdk"
)

// TestRequestExpiryHeaderMethodsAreWrapped guards against new generated
// endpoints with a privy-request-expiry header that bypass the Privy wrappers.
// For each generated service, every method whose params struct has a
// PrivyRequestExpiry field must be shadowed by a Privy method that accepts
// privyclient.RequestOption.
func TestRequestExpiryHeaderMethodsAreWrapped(t *testing.T) {
	client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{AppID: "test-app-id", AppSecret: "test-app-secret"})
	pairs := []struct {
		name      string
		generated any
		wrapper   any
	}{
		{"Intents", &client.Intents.IntentService, client.Intents},
		{"Wallets", &client.Wallets.WalletService, client.Wallets},
		{"Policies", &client.Policies.PolicyService, client.Policies},
		{"KeyQuorums", &client.KeyQuorums.KeyQuorumService, client.KeyQuorums},
		{"Wallets.Swap", &client.Wallets.Swap.WalletSwapService, client.Wallets.Swap},
		{"Wallets.DepositAccounts.Crypto", &client.Wallets.DepositAccounts.Crypto.WalletDepositAccountCryptoService, client.Wallets.DepositAccounts.Crypto},
		{"Wallets.Payout.Fiat", &client.Wallets.Payout.Fiat.WalletPayoutFiatService, client.Wallets.Payout.Fiat},
		{"Wallets.Earn.Ethereum", &client.Wallets.Earn.Ethereum.WalletEarnEthereumService, client.Wallets.Earn.Ethereum},
		{"Wallets.Earn.Ethereum.Incentive", &client.Wallets.Earn.Ethereum.Incentive.WalletEarnEthereumIncentiveService, client.Wallets.Earn.Ethereum.Incentive},
	}
	requestOptionSlice := reflect.TypeOf([]privyclient.RequestOption{})

	for _, pair := range pairs {
		genType := reflect.TypeOf(pair.generated)
		wrapType := reflect.TypeOf(pair.wrapper)
		for i := 0; i < genType.NumMethod(); i++ {
			method := genType.Method(i)
			if !hasRequestExpiryParam(method.Type) {
				continue
			}
			wrapped, ok := wrapType.MethodByName(method.Name)
			if !ok || wrapped.Type.NumIn() == 0 || wrapped.Type.In(wrapped.Type.NumIn()-1) != requestOptionSlice {
				t.Errorf("%s.%s sends privy-request-expiry but has no Privy wrapper", pair.name, method.Name)
			}
		}
	}
}

func hasRequestExpiryParam(fn reflect.Type) bool {
	for i := 0; i < fn.NumIn(); i++ {
		in := fn.In(i)
		if in.Kind() != reflect.Struct {
			continue
		}
		if _, ok := in.FieldByName("PrivyRequestExpiry"); ok {
			return true
		}
	}
	return false
}
