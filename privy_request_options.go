package privyclient

import (
	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/option"
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
//   - WithRequestExpiry: sets the request expiry as an absolute Unix timestamp in milliseconds.
type RequestOption interface {
	applyOption(*requestOptions)
}

// RpcOption is a backwards-compatible alias for RequestOption.
type RpcOption = RequestOption

type requestOptions struct {
	AuthorizationContext *authorization.AuthorizationContext
	IdempotencyKey       *string
	RequestExpiry        *int64
	RequestOptions       []option.RequestOption
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

// WithRequestExpiry sets the "privy-request-expiry" header for a single call.
// The value is an absolute Unix timestamp in milliseconds and is sent verbatim;
// use RequestExpiry(durationMs) to compute one relative to now.
//
// A per-call value takes precedence over everything else, and is sent even when
// request expiry is disabled on the client. Without it, the SDK uses the client
// default: PrivyRequestExpiryOptions.DefaultMs (DefaultRequestExpiryMs, 15
// minutes, if unset) for most endpoints, and
// PrivyRequestExpiryOptions.DefaultIntentMs (DefaultIntentRequestExpiryMs, 72
// hours, if unset) for client.Intents methods.
func WithRequestExpiry(expiry int64) RequestOption {
	return requestOptionFunc{fn: func(o *requestOptions) {
		o.RequestExpiry = &expiry
	}}
}

// WithRequestOptions passes through [option.RequestOption] values to the
// underlying API call. This allows per-call overrides of transport-level
// settings such as [option.WithHTTPClient] or [option.WithRequestTimeout].
func WithRequestOptions(opts ...option.RequestOption) RequestOption {
	return requestOptionFunc{fn: func(o *requestOptions) {
		o.RequestOptions = append(o.RequestOptions, opts...)
	}}
}
