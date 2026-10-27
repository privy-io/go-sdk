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
// UserMfaSMSEnrollmentService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaSMSEnrollmentService] method instead.
type UserMfaSMSEnrollmentService struct {
	Options []option.RequestOption
}

// NewUserMfaSMSEnrollmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserMfaSMSEnrollmentService(opts ...option.RequestOption) (r UserMfaSMSEnrollmentService) {
	r = UserMfaSMSEnrollmentService{}
	r.Options = opts
	return
}

// Sends an SMS MFA enrollment code to a phone number.
func (r *UserMfaSMSEnrollmentService) Init(ctx context.Context, userID string, params UserMfaSMSEnrollmentInitParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
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
	path := fmt.Sprintf("v1/users/%s/mfa/sms/enrollment/init", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Enrolls the phone number from init using the code sent to it.
func (r *UserMfaSMSEnrollmentService) Verify(ctx context.Context, userID string, params UserMfaSMSEnrollmentVerifyParams, opts ...option.RequestOption) (res *MfaVerificationResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/sms/enrollment/verify", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UserMfaSMSEnrollmentInitParams struct {
	// Starts SMS MFA enrollment by sending a code to the phone number.
	SMSMfaEnrollmentInitRequestBody SMSMfaEnrollmentInitRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	// MFA token from a completed MFA verification, required on MFA-protected requests
	// for users with MFA enrolled.
	PrivyMfaToken param.Opt[string] `header:"privy-mfa-token,omitzero" json:"-"`
	paramObj
}

func (r UserMfaSMSEnrollmentInitParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SMSMfaEnrollmentInitRequestBody)
}
func (r *UserMfaSMSEnrollmentInitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMfaSMSEnrollmentVerifyParams struct {
	// Completes an SMS MFA ceremony with the code sent at init.
	SMSMfaVerifyRequestBody SMSMfaVerifyRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaSMSEnrollmentVerifyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SMSMfaVerifyRequestBody)
}
func (r *UserMfaSMSEnrollmentVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
