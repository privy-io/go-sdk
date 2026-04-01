package privyclient

import (
	"github.com/privy-io/go-sdk/authorization"
)

// RequestOption configures optional parameters for API requests.
//
// Not all options apply to every API method. Callers should only use the
// options that are relevant to the method being called — unsupported options
// are ignored, similar to sending an HTTP header the server doesn't read.
//
// Available options:
//   - WithAuthorizationContext: sets the authorization context for user-owned wallet operations.
//   - WithIdempotencyKey: sets an idempotency key (applicable to Rpc, RawSign, and Ethereum/Solana convenience methods).
//   - WithRequestExpiry: sets the request expiry timestamp in milliseconds.
type RequestOption interface {
	applyOption(*requestOptions)
}

// RpcOption is a backwards-compatible alias for RequestOption.
type RpcOption = RequestOption

type requestOptions struct {
	AuthorizationContext *authorization.AuthorizationContext
	IdempotencyKey       *string
	RequestExpiry        *int64
}

type requestOptionFunc struct {
	fn func(*requestOptions)
}

func (f requestOptionFunc) applyOption(o *requestOptions) { f.fn(o) }

// WithAuthorizationContext sets the authorization context for user-owned wallet operations.
func WithAuthorizationContext(ctx *authorization.AuthorizationContext) RequestOption {
	return requestOptionFunc{fn: func(o *requestOptions) {
		o.AuthorizationContext = ctx
	}}
}

// WithIdempotencyKey sets the idempotency key for the request.
// This is applicable to Rpc, RawSign, and Ethereum/Solana convenience methods.
func WithIdempotencyKey(key string) RequestOption {
	return requestOptionFunc{fn: func(o *requestOptions) {
		o.IdempotencyKey = &key
	}}
}

// WithRequestExpiry sets the request expiry for the request.
// The value should be a Unix timestamp in milliseconds.
// If not set, the client's DefaultRequestExpiryMs is used, or 15 minutes from now.
func WithRequestExpiry(expiry int64) RequestOption {
	return requestOptionFunc{fn: func(o *requestOptions) {
		o.RequestExpiry = &expiry
	}}
}
