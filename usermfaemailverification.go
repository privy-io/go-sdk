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
// UserMfaEmailVerificationService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaEmailVerificationService] method instead.
type UserMfaEmailVerificationService struct {
	Options []option.RequestOption
}

// NewUserMfaEmailVerificationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserMfaEmailVerificationService(opts ...option.RequestOption) (r UserMfaEmailVerificationService) {
	r = UserMfaEmailVerificationService{}
	r.Options = opts
	return
}

// Sends an email MFA verification code to the enrolled email address.
func (r *UserMfaEmailVerificationService) Init(ctx context.Context, userID string, body UserMfaEmailVerificationInitParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	if !param.IsOmitted(body.UserJwt) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", body.UserJwt)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/email/verification/init", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Verifies the code sent to the enrolled email address.
func (r *UserMfaEmailVerificationService) Verify(ctx context.Context, userID string, params UserMfaEmailVerificationVerifyParams, opts ...option.RequestOption) (res *MfaVerificationResponse, err error) {
	if !param.IsOmitted(params.UserJwt) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.UserJwt)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/email/verification/verify", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UserMfaEmailVerificationInitParams struct {
	// Access token or custom authentication JWT for the user the request acts on.
	UserJwt string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

type UserMfaEmailVerificationVerifyParams struct {
	// Completes an email MFA ceremony with the code sent at init.
	EmailMfaVerifyRequestBody EmailMfaVerifyRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	UserJwt string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaEmailVerificationVerifyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailMfaVerifyRequestBody)
}
func (r *UserMfaEmailVerificationVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
