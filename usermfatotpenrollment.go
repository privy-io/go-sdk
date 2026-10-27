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
// UserMfaTotpEnrollmentService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaTotpEnrollmentService] method instead.
type UserMfaTotpEnrollmentService struct {
	Options []option.RequestOption
}

// NewUserMfaTotpEnrollmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserMfaTotpEnrollmentService(opts ...option.RequestOption) (r UserMfaTotpEnrollmentService) {
	r = UserMfaTotpEnrollmentService{}
	r.Options = opts
	return
}

// Creates a TOTP secret for the user to add to an authenticator app.
func (r *UserMfaTotpEnrollmentService) Init(ctx context.Context, userID string, body UserMfaTotpEnrollmentInitParams, opts ...option.RequestOption) (res *TotpMfaEnrollmentInitResponse, err error) {
	if !param.IsOmitted(body.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", body.PrivyUserToken)))
	}
	if !param.IsOmitted(body.PrivyMfaToken) {
		opts = append(opts, option.WithHeader("privy-mfa-token", fmt.Sprintf("%v", body.PrivyMfaToken.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/totp/enrollment/init", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Enrolls the authenticator from init using a code it generated.
func (r *UserMfaTotpEnrollmentService) Verify(ctx context.Context, userID string, params UserMfaTotpEnrollmentVerifyParams, opts ...option.RequestOption) (res *MfaVerificationResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/totp/enrollment/verify", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UserMfaTotpEnrollmentInitParams struct {
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	// MFA token from a completed MFA verification, required on MFA-protected requests
	// for users with MFA enrolled.
	PrivyMfaToken param.Opt[string] `header:"privy-mfa-token,omitzero" json:"-"`
	paramObj
}

type UserMfaTotpEnrollmentVerifyParams struct {
	// Completes a TOTP MFA ceremony with a code from the authenticator.
	TotpMfaVerifyRequestBody TotpMfaVerifyRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaTotpEnrollmentVerifyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TotpMfaVerifyRequestBody)
}
func (r *UserMfaTotpEnrollmentVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
