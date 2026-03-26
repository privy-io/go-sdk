package privyclient

import (
	"github.com/privy-io/go-sdk/authorization"
)

// RpcOption configures optional parameters for RPC calls.
type RpcOption func(*rpcOptions)

type rpcOptions struct {
	AuthorizationContext *authorization.AuthorizationContext
	IdempotencyKey       string
	RequestExpiry        int64
}

// WithAuthorizationContext sets the authorization context for user-owned wallet operations.
func WithAuthorizationContext(ctx *authorization.AuthorizationContext) RpcOption {
	return func(o *rpcOptions) {
		o.AuthorizationContext = ctx
	}
}

// WithIdempotencyKey sets the idempotency key for the request.
func WithIdempotencyKey(key string) RpcOption {
	return func(o *rpcOptions) {
		o.IdempotencyKey = key
	}
}

// WithRequestExpiry sets the request expiry for the request.
// The value should be a Unix timestamp in milliseconds.
// If not set, the client's DefaultRequestExpiryMs is used, or 15 minutes from now.
func WithRequestExpiry(expiry int64) RpcOption {
	return func(o *rpcOptions) {
		o.RequestExpiry = expiry
	}
}
