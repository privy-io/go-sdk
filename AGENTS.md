# Go SDK

Go SDK (`github.com/privy-io/go-sdk`). Stainless generates the API core; Privy-owned files add the public client, authorization, and service behavior on top.

## Commands

```sh
./scripts/bootstrap                    # install dependencies
./scripts/format                       # gofmt the repository
./scripts/lint                         # build and compile all tests
./scripts/test                         # all tests, including live e2e tests
go test -run TestName ./...             # focused test
go test $(go list ./... | grep -v /e2e) # offline tests only
go test -v ./e2e/...                    # live staging tests; requires .env
```

Run `./scripts/format`, `./scripts/lint`, and the smallest relevant test set before finishing. `./scripts/test` includes `e2e/` and therefore requires staging credentials.

## Directory Structure

```
client.go, wallet.go, ...              # GENERATED API client, resources, and models
internal/ option/ packages/ shared/    # GENERATED support packages
privy_client.go                        # CUSTOM public PrivyClient
privy_*_service.go                     # CUSTOM service wrappers
authorization/                         # CUSTOM signing and AuthorizationContext
e2e/                                   # live staging tests
```

## Key Patterns

### Keep custom behavior out of generated files

Generated files contain this notice:

```go
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
```

Top-level `.go` files without a `privy_` prefix are generally generated. Add public behavior in `privy_*.go` or `authorization/`. Generated edits can survive temporarily, but create regeneration conflicts and separate custom behavior from the extension layer.

### Embed generated services and delegate to them

`privy_client.go` exposes services as exported pointer fields. A custom service embeds its generated counterpart so unchanged methods remain available.

```go
type PrivyWalletService struct {
	WalletService
	Ethereum *PrivyEthereumWalletService
}
```

When replacing a generated sub-service:

1. Define a `privy_*_service.go` type that embeds the generated service.
2. Add a pointer field with the same name as the promoted generated field so it shadows that field.
3. Initialize it in the parent service constructor.

Callers use field access: `client.Wallets.Get(ctx, walletID)`. Do not duplicate the generated HTTP request when embedding or delegation can preserve its parameter and response handling.

### Reuse the authorization package

Use `authorization.AuthorizationContext` and the existing request preparation/signing path for authorized mutations. Do not independently canonicalize bodies, sign payloads, or assemble authorization headers inside a service.

### Separate offline and live tests

Top-level and internal package tests should remain offline. Live tests belong in `e2e/` as package `e2e_test` and use the dot import `. "github.com/privy-io/go-sdk"`.

Key helpers in `e2e/setup_test.go`:

- `newTestClient(t)` reads staging credentials.
- `setupTestWalletResources(t, client)` creates a user, key pairs, and quorum and registers cleanup.
- `resources.createTestWallets(t, chainType)` creates ownerless, key-owned, user-owned, and quorum-owned wallets.

Use these helpers for wallet behavior instead of relying on pre-existing resources. Never point e2e tests at production.

## Pull Requests

Use a Conventional Commit title. Keep generated API updates separate from handwritten service changes when possible.
