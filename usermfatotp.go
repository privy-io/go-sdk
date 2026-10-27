// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"github.com/privy-io/go-sdk/option"
)

// UserMfaTotpService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMfaTotpService] method instead.
type UserMfaTotpService struct {
	Options []option.RequestOption
	// Operations related to user MFA
	Enrollment UserMfaTotpEnrollmentService
	// Operations related to user MFA
	Verification UserMfaTotpVerificationService
}

// NewUserMfaTotpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMfaTotpService(opts ...option.RequestOption) (r UserMfaTotpService) {
	r = UserMfaTotpService{}
	r.Options = opts
	r.Enrollment = NewUserMfaTotpEnrollmentService(opts...)
	r.Verification = NewUserMfaTotpVerificationService(opts...)
	return
}
