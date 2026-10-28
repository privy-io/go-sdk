// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"

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
	Passkey UserMfaPasskeyService
	SMS     UserMfaSMSService
	Totp    UserMfaTotpService
}

// NewUserMfaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserMfaService(opts ...option.RequestOption) (r UserMfaService) {
	r = UserMfaService{}
	r.Options = opts
	r.Passkey = NewUserMfaPasskeyService(opts...)
	r.SMS = NewUserMfaSMSService(opts...)
	r.Totp = NewUserMfaTotpService(opts...)
	return
}

// The authenticator assertion response from a WebAuthn authentication ceremony.
//
// The properties AuthenticatorData, ClientDataJson, Signature are required.
type MfaPasskeyAssertionResponse struct {
	AuthenticatorData string            `json:"authenticator_data" api:"required"`
	ClientDataJson    string            `json:"client_data_json" api:"required"`
	Signature         string            `json:"signature" api:"required"`
	UserHandle        param.Opt[string] `json:"user_handle,omitzero"`
	paramObj
}

func (r MfaPasskeyAssertionResponse) MarshalJSON() (data []byte, err error) {
	type shadow MfaPasskeyAssertionResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MfaPasskeyAssertionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The authenticator attestation response from a WebAuthn registration ceremony.
//
// The properties AttestationObject, ClientDataJson are required.
type MfaPasskeyAttestationResponse struct {
	AttestationObject  string             `json:"attestation_object" api:"required"`
	ClientDataJson     string             `json:"client_data_json" api:"required"`
	AuthenticatorData  param.Opt[string]  `json:"authenticator_data,omitzero"`
	PublicKey          param.Opt[string]  `json:"public_key,omitzero"`
	PublicKeyAlgorithm param.Opt[float64] `json:"public_key_algorithm,omitzero"`
	Transports         []string           `json:"transports,omitzero"`
	paramObj
}

func (r MfaPasskeyAttestationResponse) MarshalJSON() (data []byte, err error) {
	type shadow MfaPasskeyAttestationResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MfaPasskeyAttestationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebAuthn registration options as defined by the Web Authentication
// specification.
type MfaPasskeyAuthenticatorEnrollmentOptions struct {
	Challenge        string                          `json:"challenge" api:"required"`
	PubKeyCredParams []MfaPasskeyPubKeyCredParamResp `json:"pub_key_cred_params" api:"required"`
	// Relying party information for a WebAuthn ceremony.
	Rp MfaPasskeyRelyingParty `json:"rp" api:"required"`
	// User entity for a WebAuthn registration ceremony.
	User        MfaPasskeyUser `json:"user" api:"required"`
	Attestation string         `json:"attestation"`
	// Authenticator selection criteria for a WebAuthn registration ceremony.
	AuthenticatorSelection MfaPasskeyAuthenticatorSelection `json:"authenticator_selection"`
	ExcludeCredentials     []MfaPasskeyCredentialDescriptor `json:"exclude_credentials"`
	// Extensions for a WebAuthn registration ceremony.
	Extensions MfaPasskeyEnrollmentExtensions `json:"extensions"`
	Timeout    float64                        `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Challenge              respjson.Field
		PubKeyCredParams       respjson.Field
		Rp                     respjson.Field
		User                   respjson.Field
		Attestation            respjson.Field
		AuthenticatorSelection respjson.Field
		ExcludeCredentials     respjson.Field
		Extensions             respjson.Field
		Timeout                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyAuthenticatorEnrollmentOptions) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyAuthenticatorEnrollmentOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebAuthn registration response as defined by the Web Authentication
// specification.
//
// The properties ID, ClientExtensionResults, RawID, Response, Type are required.
type MfaPasskeyAuthenticatorEnrollmentResponse struct {
	ID string `json:"id" api:"required"`
	// Client extension results returned by the WebAuthn authenticator.
	ClientExtensionResults MfaPasskeyClientExtensionResults `json:"client_extension_results,omitzero" api:"required"`
	RawID                  string                           `json:"raw_id" api:"required"`
	// The authenticator attestation response from a WebAuthn registration ceremony.
	Response MfaPasskeyAttestationResponse `json:"response,omitzero" api:"required"`
	// Any of "public-key".
	Type                    MfaPasskeyAuthenticatorEnrollmentResponseType `json:"type,omitzero" api:"required"`
	AuthenticatorAttachment param.Opt[string]                             `json:"authenticator_attachment,omitzero"`
	paramObj
}

func (r MfaPasskeyAuthenticatorEnrollmentResponse) MarshalJSON() (data []byte, err error) {
	type shadow MfaPasskeyAuthenticatorEnrollmentResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MfaPasskeyAuthenticatorEnrollmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MfaPasskeyAuthenticatorEnrollmentResponseType string

const (
	MfaPasskeyAuthenticatorEnrollmentResponseTypePublicKey MfaPasskeyAuthenticatorEnrollmentResponseType = "public-key"
)

// Authenticator selection criteria for a WebAuthn registration ceremony.
type MfaPasskeyAuthenticatorSelection struct {
	AuthenticatorAttachment string `json:"authenticator_attachment"`
	RequireResidentKey      bool   `json:"require_resident_key"`
	ResidentKey             string `json:"resident_key"`
	UserVerification        string `json:"user_verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticatorAttachment respjson.Field
		RequireResidentKey      respjson.Field
		ResidentKey             respjson.Field
		UserVerification        respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyAuthenticatorSelection) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyAuthenticatorSelection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebAuthn authentication options as defined by the Web Authentication
// specification.
type MfaPasskeyAuthenticatorVerifyOptions struct {
	Challenge        string                           `json:"challenge" api:"required"`
	AllowCredentials []MfaPasskeyCredentialDescriptor `json:"allow_credentials"`
	// Extensions for a WebAuthn authentication ceremony.
	Extensions       MfaPasskeyVerifyExtensions `json:"extensions"`
	RpID             string                     `json:"rp_id"`
	Timeout          float64                    `json:"timeout"`
	UserVerification string                     `json:"user_verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Challenge        respjson.Field
		AllowCredentials respjson.Field
		Extensions       respjson.Field
		RpID             respjson.Field
		Timeout          respjson.Field
		UserVerification respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyAuthenticatorVerifyOptions) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyAuthenticatorVerifyOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebAuthn authentication response as defined by the Web Authentication
// specification.
//
// The properties ID, ClientExtensionResults, RawID, Response, Type are required.
type MfaPasskeyAuthenticatorVerifyResponse struct {
	ID string `json:"id" api:"required"`
	// Client extension results returned by the WebAuthn authenticator.
	ClientExtensionResults MfaPasskeyClientExtensionResults `json:"client_extension_results,omitzero" api:"required"`
	RawID                  string                           `json:"raw_id" api:"required"`
	// The authenticator assertion response from a WebAuthn authentication ceremony.
	Response MfaPasskeyAssertionResponse `json:"response,omitzero" api:"required"`
	// Any of "public-key".
	Type                    MfaPasskeyAuthenticatorVerifyResponseType `json:"type,omitzero" api:"required"`
	AuthenticatorAttachment param.Opt[string]                         `json:"authenticator_attachment,omitzero"`
	paramObj
}

func (r MfaPasskeyAuthenticatorVerifyResponse) MarshalJSON() (data []byte, err error) {
	type shadow MfaPasskeyAuthenticatorVerifyResponse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MfaPasskeyAuthenticatorVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MfaPasskeyAuthenticatorVerifyResponseType string

const (
	MfaPasskeyAuthenticatorVerifyResponseTypePublicKey MfaPasskeyAuthenticatorVerifyResponseType = "public-key"
)

// Client extension results returned by the WebAuthn authenticator.
type MfaPasskeyClientExtensionResults struct {
	AppID            param.Opt[bool] `json:"app_id,omitzero"`
	HmacCreateSecret param.Opt[bool] `json:"hmac_create_secret,omitzero"`
	// The result of the WebAuthn credProps extension.
	CredProps MfaPasskeyCredPropsResult `json:"cred_props,omitzero"`
	paramObj
}

func (r MfaPasskeyClientExtensionResults) MarshalJSON() (data []byte, err error) {
	type shadow MfaPasskeyClientExtensionResults
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MfaPasskeyClientExtensionResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of the WebAuthn credProps extension.
type MfaPasskeyCredPropsResultResp struct {
	Rk bool `json:"rk"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rk          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyCredPropsResultResp) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyCredPropsResultResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MfaPasskeyCredPropsResultResp to a
// MfaPasskeyCredPropsResult.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MfaPasskeyCredPropsResult.Overrides()
func (r MfaPasskeyCredPropsResultResp) ToParam() MfaPasskeyCredPropsResult {
	return param.Override[MfaPasskeyCredPropsResult](json.RawMessage(r.RawJSON()))
}

// The result of the WebAuthn credProps extension.
type MfaPasskeyCredPropsResult struct {
	Rk param.Opt[bool] `json:"rk,omitzero"`
	paramObj
}

func (r MfaPasskeyCredPropsResult) MarshalJSON() (data []byte, err error) {
	type shadow MfaPasskeyCredPropsResult
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MfaPasskeyCredPropsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A WebAuthn credential descriptor identifying a specific public key credential.
type MfaPasskeyCredentialDescriptor struct {
	ID         string   `json:"id" api:"required"`
	Type       string   `json:"type" api:"required"`
	Transports []string `json:"transports"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Transports  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyCredentialDescriptor) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyCredentialDescriptor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extensions for a WebAuthn registration ceremony.
type MfaPasskeyEnrollmentExtensions struct {
	AppID string `json:"app_id"`
	// The result of the WebAuthn credProps extension.
	CredProps        MfaPasskeyCredPropsResultResp `json:"cred_props"`
	HmacCreateSecret bool                          `json:"hmac_create_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID            respjson.Field
		CredProps        respjson.Field
		HmacCreateSecret respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyEnrollmentExtensions) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyEnrollmentExtensions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A public key credential parameter specifying the algorithm and credential type.
type MfaPasskeyPubKeyCredParamResp struct {
	Alg float64 `json:"alg" api:"required"`
	// Any of "public-key".
	Type MfaPasskeyPubKeyCredParamType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alg         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyPubKeyCredParamResp) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyPubKeyCredParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MfaPasskeyPubKeyCredParamType string

const (
	MfaPasskeyPubKeyCredParamTypePublicKey MfaPasskeyPubKeyCredParamType = "public-key"
)

// Relying party information for a WebAuthn ceremony.
type MfaPasskeyRelyingParty struct {
	Name string `json:"name" api:"required"`
	ID   string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyRelyingParty) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyRelyingParty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User entity for a WebAuthn registration ceremony.
type MfaPasskeyUser struct {
	ID          string `json:"id" api:"required"`
	DisplayName string `json:"display_name" api:"required"`
	Name        string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyUser) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extensions for a WebAuthn authentication ceremony.
type MfaPasskeyVerifyExtensions struct {
	AppID            string `json:"app_id"`
	CredProps        bool   `json:"cred_props"`
	HmacCreateSecret bool   `json:"hmac_create_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID            respjson.Field
		CredProps        respjson.Field
		HmacCreateSecret respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MfaPasskeyVerifyExtensions) RawJSON() string { return r.JSON.raw }
func (r *MfaPasskeyVerifyExtensions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// WebAuthn registration options for the started passkey MFA enrollment.
type PasskeyMfaEnrollmentInitResponse struct {
	// WebAuthn registration options as defined by the Web Authentication
	// specification.
	Options MfaPasskeyAuthenticatorEnrollmentOptions `json:"options" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PasskeyMfaEnrollmentInitResponse) RawJSON() string { return r.JSON.raw }
func (r *PasskeyMfaEnrollmentInitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Completes passkey MFA enrollment with a WebAuthn registration response.
//
// The property Credential is required.
type PasskeyMfaEnrollmentVerifyRequestBody struct {
	// WebAuthn registration response as defined by the Web Authentication
	// specification.
	Credential MfaPasskeyAuthenticatorEnrollmentResponse `json:"credential,omitzero" api:"required"`
	paramObj
}

func (r PasskeyMfaEnrollmentVerifyRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow PasskeyMfaEnrollmentVerifyRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PasskeyMfaEnrollmentVerifyRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Starts a passkey MFA ceremony for the specified relying party origin.
//
// The property RelyingParty is required.
type PasskeyMfaInitRequestBody struct {
	// Origin where the authenticator runs, such as https://app.example.com.
	RelyingParty string `json:"relying_party" api:"required" format:"uri"`
	paramObj
}

func (r PasskeyMfaInitRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow PasskeyMfaInitRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PasskeyMfaInitRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebAuthn authentication options for the started passkey MFA verification.
type PasskeyMfaVerificationInitResponse struct {
	// WebAuthn authentication options as defined by the Web Authentication
	// specification.
	Options MfaPasskeyAuthenticatorVerifyOptions `json:"options" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PasskeyMfaVerificationInitResponse) RawJSON() string { return r.JSON.raw }
func (r *PasskeyMfaVerificationInitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Completes passkey MFA verification with a WebAuthn authentication response.
//
// The property Credential is required.
type PasskeyMfaVerificationVerifyRequestBody struct {
	// WebAuthn authentication response as defined by the Web Authentication
	// specification.
	Credential MfaPasskeyAuthenticatorVerifyResponse `json:"credential,omitzero" api:"required"`
	paramObj
}

func (r PasskeyMfaVerificationVerifyRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow PasskeyMfaVerificationVerifyRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PasskeyMfaVerificationVerifyRequestBody) UnmarshalJSON(data []byte) error {
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
