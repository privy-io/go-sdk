// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// UserMfaEmailService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaEmailService] method instead.
type UserMfaEmailService struct {
	Options []option.RequestOption
	// Operations related to user MFA
	Enrollment UserMfaEmailEnrollmentService
	// Operations related to user MFA
	Verification UserMfaEmailVerificationService
}

// NewUserMfaEmailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMfaEmailService(opts ...option.RequestOption) (r UserMfaEmailService) {
	r = UserMfaEmailService{}
	r.Options = opts
	r.Enrollment = NewUserMfaEmailEnrollmentService(opts...)
	r.Verification = NewUserMfaEmailVerificationService(opts...)
	return
}
