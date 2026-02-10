package privyclient

import (
	"github.com/privy-io/go-sdk/authorization"
)

// RpcOption configures optional parameters for RPC calls.
type RpcOption func(*rpcOptions)

type rpcOptions struct {
	AuthorizationContext *authorization.AuthorizationContext
	IdempotencyKey       string
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
