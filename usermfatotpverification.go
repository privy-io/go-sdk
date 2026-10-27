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
// UserMfaTotpVerificationService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaTotpVerificationService] method instead.
type UserMfaTotpVerificationService struct {
	Options []option.RequestOption
}

// NewUserMfaTotpVerificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserMfaTotpVerificationService(opts ...option.RequestOption) (r UserMfaTotpVerificationService) {
	r = UserMfaTotpVerificationService{}
	r.Options = opts
	return
}

// Verifies a code from the enrolled authenticator.
func (r *UserMfaTotpVerificationService) Verify(ctx context.Context, userID string, params UserMfaTotpVerificationVerifyParams, opts ...option.RequestOption) (res *MfaVerificationResponse, err error) {
	if !param.IsOmitted(params.PrivyUserToken) {
		opts = append(opts, option.WithHeader("privy-user-token", fmt.Sprintf("%v", params.PrivyUserToken)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/mfa/totp/verification/verify", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UserMfaTotpVerificationVerifyParams struct {
	// Completes a TOTP MFA ceremony with a code from the authenticator.
	TotpMfaVerifyRequestBody TotpMfaVerifyRequestBody
	// Access token or custom authentication JWT for the user the request acts on.
	PrivyUserToken string `header:"privy-user-token" api:"required" json:"-"`
	paramObj
}

func (r UserMfaTotpVerificationVerifyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TotpMfaVerifyRequestBody)
}
func (r *UserMfaTotpVerificationVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
