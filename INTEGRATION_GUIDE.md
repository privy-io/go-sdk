# Privy Go SDK Integration Guide

## Getting Started

### Installation
Install the Privy Go SDK using go get:
```bash
go get github.com/privy-io/go-sdk
```

### Requirements
- Go 1.23 or higher
- Privy App ID and App Secret (available in your Privy dashboard)

### Quick Setup
Initialize the Privy client with your credentials:
```go
import privyclient "github.com/privy-io/go-sdk"

client := privyclient.NewPrivyClient(privyclient.PrivyClientOptions{
    AppID:     "your-app-id",
    AppSecret: "your-app-secret",
    LogLevel:  privyclient.LogLevelInfo, // Optional: LogLevelNone, LogLevelError, LogLevelInfo, LogLevelDebug, LogLevelVerbose
})
```

---

## Overview

The Privy client is the main entry point for the SDK. Once initialized, you can access multiple services that represent different parts of the Privy API:

- **Users** - Manage user accounts and linked identities
- **Wallets** - Create and manage embedded wallets across multiple chains
- **Policies** - Define authorization rules for wallet operations
- **KeyQuorums** - Manage multi-signature wallet configurations
- **JwtExchange** - Exchange user JWTs for authorization keys
- **Transactions** - Access transaction-related functionality

Each service provides methods for interacting with its respective API endpoints.

---

## Authorization Context & Signatures

When updating resources like wallets, policies, or key quorums in the Privy API, requests [must be signed](https://docs.privy.io/controls/authorization-keys/using-owners/sign/overview) 
by the resource owner in order to be authorized. Privy's Go SDK exposes utilities to simplify the authorization flow.


### Authorization Context

AuthorizationContex is a struct that contains credentials used for signing authorization requests, which can be passed into methods that
require owner authorization.

```go
import "github.com/privy-io/go-sdk/authorization"

authCtx := authorization.AuthorizationContext{
    // Option 1: Use private keys directly
    PrivateKeys: []string{"base64-encoded-pkcs8-p256-key"},

    // Option 2: Use user JWTs (automatically exchanged for auth keys)
    UserJwts: []string{"user-jwt-token"},

    // Option 3: Use pre-computed signatures
    Signatures: []string{"base64-signature"},

    // Option 4: Use external signers (e.g., KMS, hardware wallets)
    Signers: []authorization.AuthorizationSigner{customSigner},
}
```

### SDK Convenience Functions

The SDK provides some convenience functions that accept an `AuthorizationContext` and handle all authorization steps automatically under the hood. These functions:
- Build the signature input from request parameters
- Format the request payload for signing
- Generate signatures from all credentials in the authorization context
- Set the authorization signature header on the request

This simplifies authorized operations to a single function call.

```go
// Create an authorization context with your credentials
authCtx := authorization.AuthorizationContext{
    PrivateKeys: []string{"MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg..."}, // Your private key
    UserJwts:    []string{"eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9..."},              // Your user JWT
}

// Execute an RPC call with automatic authorization
// The SDK handles signature generation and header injection automatically
result, err := client.Wallets.Rpc(
    context.Background(),
    "wallet-id",
    privyclient.WalletRpcParams{
        Method:  "eth_signTypedData_v4",
        Params:  privyclient.WalletRpcParamsParamsUnion{OfTypedDataSign: &privyclient.TypedDataSignParams{...}},
        ChainID: "1",
    },
    WithAuthorizationContext(&authorization.AuthorizationContext{
        UserJwts: []string{jwt},
    }),
)

if err != nil {
    panic(err)
}
```

### Generating Signatures

If the SDK doesn't have a convenience function for a particular action, you can build the signature 
input and generate the authorization signature.

```go
// Create an authorization context with both a private key and a JWT token
authCtx := authorization.AuthorizationContext{
    // Add a base64-encoded PKCS8-formatted P-256 private key
    PrivateKeys: []string{
        "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg...", // Your base64-encoded private key
    },
    // Add a user JWT token (will be automatically exchanged for authorization keys under the hood)
    UserJwts: []string{
        "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyXzEyMzQ1Njc4OTAi...", // Your user JWT
    },
}

input := authorization.WalletApiRequestSignatureInput{
    Version: 1,
    Method:  "POST",
    URL:     "https://auth.privy.io/v1/wallets/{wallet_ID}/rpc",
    Body:    params,
    Headers: headers,
}

# // TODO: LUCAS - UPDATE THIS EXAMPLE WHEN GenerateAuthorizationSignature TAKES IN WalletApiRequestSignatureInput directly


// Single signature
signature, err := authorization.GenerateAuthorizationSignature(privateKey, payload)

// Multiple signatures (all credentials in context)
signatures, err := authorization.GenerateAuthorizationSignatures(
    ctx,
    authCtx,
    payload,
    client.JwtExchange, // For JWT exchange
)
```

### Formatting Requests for Signing

If you prefer to sign a request yourself through an external service (like a KMS), use the 
`FormatRequestForAuthorizationSignature` helper to to generate your signature payload. 
You can then take the returned serialized payload and call out to a signing service to generate a P256 signature over the payload.

```go
input := authorization.WalletApiRequestSignatureInput{
    Version: 1,
    Method:  "POST",
    URL:     "https://auth.privy.io/v1/wallets/{wallet_ID}/rpc",
    Body:    params,
    Headers: headers,
}

payload, err := authorization.FormatRequestForAuthorizationSignature(input)
if err != nil {
    panic(err)
}
```

**Key Requirements:**
- Private keys must be base64-encoded PKCS8-formatted P-256 keys
- Payloads are hashed with SHA-256 before signing
- Signatures use ECDSA with DER encoding

---

## User Management

### Creating Users
```go
user, err := client.Users.New(context.Background(), privyclient.UserNewParams{
    LinkedAccounts: []privyclient.LinkedAccountInputUnionParam{{
        OfEmail: &privyclient.LinkedAccountEmailInputParam{
            Address: "user@example.com",
            Type:    privyclient.LinkedAccountEmailInputTypeEmail,
        },
    }},
})
```

### Looking Up Users
Find users by various identifiers:
```go
// By email
user, err := client.Users.GetByEmailAddress(ctx, privyclient.UserGetByEmailAddressParams{
    Address: "user@example.com",
})

// By user ID
user, err := client.Users.Get(ctx, "user_id")
```

---

## Wallet Operations

### Creating Wallets
```go
wallet, err := client.Wallets.New(context.Background(), privyclient.WalletNewParams{
    ChainType: privyclient.WalletNewParamsChainTypeEthereum,
    OwnerID:   "user_id_or_key_quorum_id",
})
```

### Signing Operations
```go
// Sign a message
signature, err := client.Wallets.SignMessage(ctx, "wallet_id",
    privyclient.WalletSignMessageParams{
        Message: "Hello, blockchain!",
    })

// Execute RPC call
result, err := client.Wallets.Rpc(ctx, "wallet_id",
    privyclient.WalletRpcParams{
        Method: "eth_sendTransaction",
        Params: []interface{}{...},
    })
```

### TODO: Showcase 7702 and sign user operation convenience functions