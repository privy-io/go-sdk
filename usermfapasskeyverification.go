// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
)

// Operations related to user MFA
//
// UserMfaPasskeyVerificationService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaPasskeyVerificationService] method instead.
type UserMfaPasskeyVerificationService struct {
	Options []option.RequestOption
}

// NewUserMfaPasskeyVerificationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserMfaPasskeyVerificationService(opts ...option.RequestOption) (r UserMfaPasskeyVerificationService) {
	r = UserMfaPasskeyVerificationService{}
	r.Options = opts
	return
}

// Returns WebAuthn authentication options for the enrolled MFA passkeys.
func (r *UserMfaPasskeyVerificationService) Init(ctx context.Context, userID string, params UserMfaPasskeyVerificationInitParams, opts ...option.RequestOption) (res *PasskeyMfaVerificationInitResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/passkey/verification/init", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Verifies the WebAuthn authentication response from an MFA passkey.
func (r *UserMfaPasskeyVerificationService) Verify(ctx context.Context, userID string, params UserMfaPasskeyVerificationVerifyParams, opts ...option.RequestOption) (res *MfaVerificationResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/passkey/verification/verify", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UserMfaPasskeyVerificationInitParams struct {
	// Starts a passkey MFA ceremony for the specified relying party origin.
	PasskeyMfaInitRequestBody PasskeyMfaInitRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaPasskeyVerificationInitParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PasskeyMfaInitRequestBody)
}
func (r *UserMfaPasskeyVerificationInitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMfaPasskeyVerificationVerifyParams struct {
	// Completes passkey MFA verification with a WebAuthn authentication response.
	PasskeyMfaVerificationVerifyRequestBody PasskeyMfaVerificationVerifyRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaPasskeyVerificationVerifyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PasskeyMfaVerificationVerifyRequestBody)
}
func (r *UserMfaPasskeyVerificationVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
