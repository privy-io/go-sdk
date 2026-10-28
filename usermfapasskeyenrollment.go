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
// UserMfaPasskeyEnrollmentService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaPasskeyEnrollmentService] method instead.
type UserMfaPasskeyEnrollmentService struct {
	Options []option.RequestOption
}

// NewUserMfaPasskeyEnrollmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserMfaPasskeyEnrollmentService(opts ...option.RequestOption) (r UserMfaPasskeyEnrollmentService) {
	r = UserMfaPasskeyEnrollmentService{}
	r.Options = opts
	return
}

// Returns WebAuthn registration options for a new MFA passkey.
func (r *UserMfaPasskeyEnrollmentService) Init(ctx context.Context, userID string, params UserMfaPasskeyEnrollmentInitParams, opts ...option.RequestOption) (res *PasskeyMfaEnrollmentInitResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	if !param.IsOmitted(params.PrivyMfaToken) {
		opts = append(opts, option.WithHeader("privy-mfa-token", fmt.Sprintf("%v", params.PrivyMfaToken.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/passkey/enrollment/init", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Enrolls a passkey in MFA using the WebAuthn registration response.
func (r *UserMfaPasskeyEnrollmentService) Verify(ctx context.Context, userID string, params UserMfaPasskeyEnrollmentVerifyParams, opts ...option.RequestOption) (res *MfaVerificationResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/passkey/enrollment/verify", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UserMfaPasskeyEnrollmentInitParams struct {
	// Starts a passkey MFA ceremony for the specified relying party origin.
	PasskeyMfaInitRequestBody PasskeyMfaInitRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	// MFA token from a completed MFA verification, required on MFA-protected requests
	// for users with MFA enrolled.
	PrivyMfaToken param.Opt[string] `header:"privy-mfa-token,omitzero" json:"-"`
	paramObj
}

func (r UserMfaPasskeyEnrollmentInitParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PasskeyMfaInitRequestBody)
}
func (r *UserMfaPasskeyEnrollmentInitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMfaPasskeyEnrollmentVerifyParams struct {
	// Completes passkey MFA enrollment with a WebAuthn registration response.
	PasskeyMfaEnrollmentVerifyRequestBody PasskeyMfaEnrollmentVerifyRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaPasskeyEnrollmentVerifyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PasskeyMfaEnrollmentVerifyRequestBody)
}
func (r *UserMfaPasskeyEnrollmentVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
