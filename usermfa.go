// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// UserMfaService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaService] method instead.
type UserMfaService struct {
	Options []option.RequestOption
	SMS     UserMfaSMSService
	Totp    UserMfaTotpService
}

// NewUserMfaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserMfaService(opts ...option.RequestOption) (r UserMfaService) {
	r = UserMfaService{}
	r.Options = opts
	r.SMS = NewUserMfaSMSService(opts...)
	r.Totp = NewUserMfaTotpService(opts...)
	return
}

// An MFA token and the current user after completing an MFA ceremony.
type MfaVerificationResponse struct {
	// JWT proving the user completed MFA.
	Token string `json:"token" api:"required"`
	// Unix timestamp in seconds when the MFA token expires.
	ExpiresAt int64 `json:"expires_at" api:"required"`
	// A Privy user object.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaVerificationResponse) RawJSON() string { return r.JSON.raw }
func (r *MfaVerificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Starts SMS MFA enrollment by sending a code to the phone number.
//
// The property PhoneNumber is required.
type SMSMfaEnrollmentInitRequestBody struct {
	PhoneNumber string `json:"phone_number" api:"required"`
	paramObj
}

func (r SMSMfaEnrollmentInitRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow SMSMfaEnrollmentInitRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SMSMfaEnrollmentInitRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Completes an SMS MFA ceremony with the code sent at init.
//
// The property Code is required.
type SMSMfaVerifyRequestBody struct {
	Code string `json:"code" api:"required"`
	paramObj
}

func (r SMSMfaVerifyRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow SMSMfaVerifyRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SMSMfaVerifyRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authenticator setup information for the started TOTP MFA enrollment.
type TotpMfaEnrollmentInitResponse struct {
	// Authenticator setup information for TOTP MFA enrollment.
	Options TotpMfaEnrollmentOptions `json:"options" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TotpMfaEnrollmentInitResponse) RawJSON() string { return r.JSON.raw }
func (r *TotpMfaEnrollmentInitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authenticator setup information for TOTP MFA enrollment.
type TotpMfaEnrollmentOptions struct {
	// Base32-encoded TOTP secret.
	Secret string `json:"secret" api:"required"`
	// Authenticator setup URI using the otpauth scheme.
	Uri string `json:"uri" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TotpMfaEnrollmentOptions) RawJSON() string { return r.JSON.raw }
func (r *TotpMfaEnrollmentOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Completes a TOTP MFA ceremony with a code from the authenticator.
//
// The property Code is required.
type TotpMfaVerifyRequestBody struct {
	Code string `json:"code" api:"required"`
	paramObj
}

func (r TotpMfaVerifyRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow TotpMfaVerifyRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TotpMfaVerifyRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
