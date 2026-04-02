package privyclient

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/privy-io/go-sdk/authorization"
	"github.com/privy-io/go-sdk/internal/jwtexchange"
)

// RequestExpiry computes a request expiry timestamp (Unix milliseconds)
// from a duration offset. For example, RequestExpiry(20 * 60 * 1000)
// returns a timestamp 20 minutes from now.
func RequestExpiry(durationMsFromNow int64) int64 {
	return time.Now().UnixMilli() + durationMsFromNow
}

// applyRequestOptions applies the given options and returns the resulting requestOptions.
func applyRequestOptions(opts []RequestOption) *requestOptions {
	options := &requestOptions{}
	for _, opt := range opts {
		opt.applyOption(options)
	}
	return options
}

// prepareRequestInput contains the request details for prepareRequest.
type prepareRequestInput struct {
	authorizationContext *authorization.AuthorizationContext
	idempotencyKey       *string
	requestExpiry        *int64
	method               string // HTTP method: "POST", "PATCH", etc.
	url                  string // Full request URL
	body                 any    // Request body (JSON-serializable)
}

// preparedRequest contains the headers computed by prepareRequest that should
// be applied to the underlying API call.
type preparedRequest struct {
	// privyAuthorizationSignature is the computed request signature header.
	// Nil if no authorization signature was generated.
	privyAuthorizationSignature *string

	// privyIdempotencyKey is the idempotency key header, if set.
	privyIdempotencyKey *string

	// privyRequestExpiry is the request expiry header, if set.
	privyRequestExpiry *string
}

// prepareRequest computes the authorization signature and assembles all
// Privy-specific headers for an API request.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - appID: The Privy app ID
//   - jwtExchanger: For exchanging user JWTs for authorization keys
//   - input: The request details including authorization context, method, URL, body,
//     and optional idempotency key and request expiry
func prepareRequest(
	ctx context.Context,
	appID string,
	jwtExchanger jwtexchange.JwtExchanger,
	input prepareRequestInput,
) (*preparedRequest, error) {
	result := &preparedRequest{}

	if input.idempotencyKey != nil {
		result.privyIdempotencyKey = input.idempotencyKey
	}

	if input.requestExpiry != nil {
		result.privyRequestExpiry = stringPtr(strconv.FormatInt(*input.requestExpiry, 10))
	}

	// Generate authorization signature if context is provided
	if input.authorizationContext != nil {
		// Build headers for signature computation
		sigHeaders := map[string]string{
			"privy-app-id": appID,
		}
		if result.privyIdempotencyKey != nil {
			sigHeaders["privy-idempotency-key"] = *result.privyIdempotencyKey
		}
		if result.privyRequestExpiry != nil {
			sigHeaders["privy-request-expiry"] = *result.privyRequestExpiry
		}

		sigInput := authorization.WalletApiRequestSignatureInput{
			Version: 1,
			Method:  input.method,
			URL:     input.url,
			Body:    input.body,
			Headers: sigHeaders,
		}

		signatures, err := authorization.GenerateAuthorizationSignaturesForRequest(
			ctx,
			*input.authorizationContext,
			sigInput,
			jwtExchanger,
		)

		if err != nil {
			return nil, err
		}

		if len(signatures) > 0 {
			result.privyAuthorizationSignature = stringPtr(strings.Join(signatures, ","))
		}
	}

	return result, nil
}

func stringPtr(v string) *string {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}
