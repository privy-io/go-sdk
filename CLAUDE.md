# Go SDK

Go SDK (`github.com/privy-io/go-sdk`).

## Generated Files — DO NOT MODIFY

Top-level `.go` files **without** a `privy_` prefix are auto-generated from the OpenAPI spec:

```
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
```

Add new behaviour in `privy_*.go` files instead.

## Service Layer Pattern

`privy_client.go` exposes services as exported pointer fields (`Wallets`, `Users`, etc.).
Each `privy_*_service.go` embeds its generated counterpart to inherit all methods.

To add a sub-service on an existing service struct:

1. Create a new `privy_*_service.go` that embeds the generated service struct.
2. Add a pointer field of the new type to the parent service struct, using the **same name** as
   the promoted field from the embedded generated struct — this shadows it.
3. Initialize the field in the parent service's constructor function.

Callers use field access: `client.Wallets.Get(ctx, walletID)`.

## E2E Tests

Tests live in `e2e/` as package `e2e_test`. Use dot import `. "github.com/privy-io/go-sdk"`.

Key helpers in `e2e/setup_test.go`:

- `newTestClient(t)` — reads credentials from env
- `setupTestWalletResources(t, client)` — creates a user, key pair, and quorum; registers cleanup
- `createTestWallets(t, chainType)` — returns ownerless / key-owned / user-owned / quorum-owned wallets

Use `setupTestWalletResources` + `createTestWallets` for all tests.

## Lint

```sh
just go::lint
```
