// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to app settings and allowlist management
//
// AppService contains methods and other services that help with interacting with
// the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
	// Operations related to app settings and allowlist management
	Allowlist AppAllowlistService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	r.Allowlist = NewAppAllowlistService(opts...)
	return
}

// Get the settings and configuration for an app.
func (r *AppService) Get(ctx context.Context, appID string, opts ...option.RequestOption) (res *AppResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/apps/%s", url.PathEscape(appID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The response for getting an app.
type AppResponse struct {
	ID                         string                           `json:"id" api:"required"`
	AccentColor                string                           `json:"accent_color" api:"required"`
	AllowedDomains             []string                         `json:"allowed_domains" api:"required"`
	AllowedNativeAppIDs        []string                         `json:"allowed_native_app_ids" api:"required"`
	AllowedNativeAppURLSchemes []string                         `json:"allowed_native_app_url_schemes" api:"required"`
	AllowlistConfig            AppResponseAllowlistConfig       `json:"allowlist_config" api:"required"`
	AllowlistEnabled           bool                             `json:"allowlist_enabled" api:"required"`
	AppleOAuth                 bool                             `json:"apple_oauth" api:"required"`
	CaptchaEnabled             bool                             `json:"captcha_enabled" api:"required"`
	CustomAPIURL               string                           `json:"custom_api_url" api:"required"`
	CustomJwtAuth              bool                             `json:"custom_jwt_auth" api:"required"`
	CustomOAuthProviders       []AppResponseCustomOAuthProvider `json:"custom_oauth_providers" api:"required"`
	DisablePlusEmails          bool                             `json:"disable_plus_emails" api:"required"`
	DiscordOAuth               bool                             `json:"discord_oauth" api:"required"`
	EmailAuth                  bool                             `json:"email_auth" api:"required"`
	EmbeddedWalletConfig       AppResponseEmbeddedWalletConfig  `json:"embedded_wallet_config" api:"required"`
	// Any of "turnstile", "hcaptcha".
	EnabledCaptchaProvider      AppResponseEnabledCaptchaProvider `json:"enabled_captcha_provider" api:"required"`
	EnforceWalletUis            bool                              `json:"enforce_wallet_uis" api:"required"`
	FarcasterAuth               bool                              `json:"farcaster_auth" api:"required"`
	FarcasterLinkWalletsEnabled bool                              `json:"farcaster_link_wallets_enabled" api:"required"`
	FiatOnRampEnabled           bool                              `json:"fiat_on_ramp_enabled" api:"required"`
	GitHubOAuth                 bool                              `json:"github_oauth" api:"required"`
	GoogleOAuth                 bool                              `json:"google_oauth" api:"required"`
	GuestAuth                   bool                              `json:"guest_auth" api:"required"`
	IconURL                     string                            `json:"icon_url" api:"required"`
	InstagramOAuth              bool                              `json:"instagram_oauth" api:"required"`
	LegacyWalletUiConfig        bool                              `json:"legacy_wallet_ui_config" api:"required"`
	LineOAuth                   bool                              `json:"line_oauth" api:"required"`
	LinkedinOAuth               bool                              `json:"linkedin_oauth" api:"required"`
	LogoURL                     string                            `json:"logo_url" api:"required"`
	MaxLinkedWalletsPerUser     float64                           `json:"max_linked_wallets_per_user" api:"required"`
	// Any of "sms", "totp", "passkey".
	MfaMethods                  []string                          `json:"mfa_methods" api:"required"`
	Name                        string                            `json:"name" api:"required"`
	PasskeyAuth                 bool                              `json:"passkey_auth" api:"required"`
	PasskeysForSignupEnabled    bool                              `json:"passkeys_for_signup_enabled" api:"required"`
	PrivacyPolicyURL            string                            `json:"privacy_policy_url" api:"required"`
	RequireUsersAcceptTerms     bool                              `json:"require_users_accept_terms" api:"required"`
	ShowWalletLoginFirst        bool                              `json:"show_wallet_login_first" api:"required"`
	SmartWalletConfig           AppResponseSmartWalletConfigUnion `json:"smart_wallet_config" api:"required"`
	SMSAuth                     bool                              `json:"sms_auth" api:"required"`
	SolanaWalletAuth            bool                              `json:"solana_wallet_auth" api:"required"`
	SpotifyOAuth                bool                              `json:"spotify_oauth" api:"required"`
	TelegramAuth                bool                              `json:"telegram_auth" api:"required"`
	TermsAndConditionsURL       string                            `json:"terms_and_conditions_url" api:"required"`
	Theme                       string                            `json:"theme" api:"required"`
	TiktokOAuth                 bool                              `json:"tiktok_oauth" api:"required"`
	TwitchOAuth                 bool                              `json:"twitch_oauth" api:"required"`
	TwitterOAuth                bool                              `json:"twitter_oauth" api:"required"`
	TwitterOAuthOnMobileEnabled bool                              `json:"twitter_oauth_on_mobile_enabled" api:"required"`
	VerificationKey             string                            `json:"verification_key" api:"required"`
	WalletAuth                  bool                              `json:"wallet_auth" api:"required"`
	WalletConnectCloudProjectID string                            `json:"wallet_connect_cloud_project_id" api:"required"`
	WhatsappEnabled             bool                              `json:"whatsapp_enabled" api:"required"`
	CaptchaSiteKey              string                            `json:"captcha_site_key"`
	FundingConfig               AppResponseFundingConfig          `json:"funding_config"`
	TelegramAuthConfig          AppResponseTelegramAuthConfig     `json:"telegram_auth_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccentColor                 respjson.Field
		AllowedDomains              respjson.Field
		AllowedNativeAppIDs         respjson.Field
		AllowedNativeAppURLSchemes  respjson.Field
		AllowlistConfig             respjson.Field
		AllowlistEnabled            respjson.Field
		AppleOAuth                  respjson.Field
		CaptchaEnabled              respjson.Field
		CustomAPIURL                respjson.Field
		CustomJwtAuth               respjson.Field
		CustomOAuthProviders        respjson.Field
		DisablePlusEmails           respjson.Field
		DiscordOAuth                respjson.Field
		EmailAuth                   respjson.Field
		EmbeddedWalletConfig        respjson.Field
		EnabledCaptchaProvider      respjson.Field
		EnforceWalletUis            respjson.Field
		FarcasterAuth               respjson.Field
		FarcasterLinkWalletsEnabled respjson.Field
		FiatOnRampEnabled           respjson.Field
		GitHubOAuth                 respjson.Field
		GoogleOAuth                 respjson.Field
		GuestAuth                   respjson.Field
		IconURL                     respjson.Field
		InstagramOAuth              respjson.Field
		LegacyWalletUiConfig        respjson.Field
		LineOAuth                   respjson.Field
		LinkedinOAuth               respjson.Field
		LogoURL                     respjson.Field
		MaxLinkedWalletsPerUser     respjson.Field
		MfaMethods                  respjson.Field
		Name                        respjson.Field
		PasskeyAuth                 respjson.Field
		PasskeysForSignupEnabled    respjson.Field
		PrivacyPolicyURL            respjson.Field
		RequireUsersAcceptTerms     respjson.Field
		ShowWalletLoginFirst        respjson.Field
		SmartWalletConfig           respjson.Field
		SMSAuth                     respjson.Field
		SolanaWalletAuth            respjson.Field
		SpotifyOAuth                respjson.Field
		TelegramAuth                respjson.Field
		TermsAndConditionsURL       respjson.Field
		Theme                       respjson.Field
		TiktokOAuth                 respjson.Field
		TwitchOAuth                 respjson.Field
		TwitterOAuth                respjson.Field
		TwitterOAuthOnMobileEnabled respjson.Field
		VerificationKey             respjson.Field
		WalletAuth                  respjson.Field
		WalletConnectCloudProjectID respjson.Field
		WhatsappEnabled             respjson.Field
		CaptchaSiteKey              respjson.Field
		FundingConfig               respjson.Field
		TelegramAuthConfig          respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponse) RawJSON() string { return r.JSON.raw }
func (r *AppResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseAllowlistConfig struct {
	CtaLink     string `json:"cta_link" api:"required"`
	CtaText     string `json:"cta_text" api:"required"`
	ErrorDetail string `json:"error_detail" api:"required"`
	ErrorTitle  string `json:"error_title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CtaLink     respjson.Field
		CtaText     respjson.Field
		ErrorDetail respjson.Field
		ErrorTitle  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseAllowlistConfig) RawJSON() string { return r.JSON.raw }
func (r *AppResponseAllowlistConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseCustomOAuthProvider struct {
	Enabled bool `json:"enabled" api:"required"`
	// The ID of a custom OAuth provider, set up for this app. Must start with
	// "custom:".
	Provider            CustomOAuthProviderID `json:"provider" api:"required"`
	ProviderDisplayName string                `json:"provider_display_name" api:"required"`
	ProviderIconURL     string                `json:"provider_icon_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled             respjson.Field
		Provider            respjson.Field
		ProviderDisplayName respjson.Field
		ProviderIconURL     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseCustomOAuthProvider) RawJSON() string { return r.JSON.raw }
func (r *AppResponseCustomOAuthProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseEmbeddedWalletConfig struct {
	// Any of "users-without-wallets", "all-users", "off".
	CreateOnLogin string                                  `json:"create_on_login" api:"required"`
	Ethereum      AppResponseEmbeddedWalletConfigEthereum `json:"ethereum" api:"required"`
	// Any of "legacy-embedded-wallets-only", "user-controlled-server-wallets-only".
	Mode   string                                `json:"mode" api:"required"`
	Solana AppResponseEmbeddedWalletConfigSolana `json:"solana" api:"required"`
	// Any of "user-passcode", "google-drive", "icloud".
	UserOwnedRecoveryOptions         []string `json:"user_owned_recovery_options" api:"required"`
	RequireUserOwnedRecoveryOnCreate bool     `json:"require_user_owned_recovery_on_create"`
	RequireUserPasswordOnCreate      bool     `json:"require_user_password_on_create"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateOnLogin                    respjson.Field
		Ethereum                         respjson.Field
		Mode                             respjson.Field
		Solana                           respjson.Field
		UserOwnedRecoveryOptions         respjson.Field
		RequireUserOwnedRecoveryOnCreate respjson.Field
		RequireUserPasswordOnCreate      respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseEmbeddedWalletConfig) RawJSON() string { return r.JSON.raw }
func (r *AppResponseEmbeddedWalletConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseEmbeddedWalletConfigEthereum struct {
	// Any of "users-without-wallets", "all-users", "off".
	CreateOnLogin string `json:"create_on_login" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateOnLogin respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseEmbeddedWalletConfigEthereum) RawJSON() string { return r.JSON.raw }
func (r *AppResponseEmbeddedWalletConfigEthereum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseEmbeddedWalletConfigSolana struct {
	// Any of "users-without-wallets", "all-users", "off".
	CreateOnLogin string `json:"create_on_login" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateOnLogin respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseEmbeddedWalletConfigSolana) RawJSON() string { return r.JSON.raw }
func (r *AppResponseEmbeddedWalletConfigSolana) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseEnabledCaptchaProvider string

const (
	AppResponseEnabledCaptchaProviderTurnstile AppResponseEnabledCaptchaProvider = "turnstile"
	AppResponseEnabledCaptchaProviderHcaptcha  AppResponseEnabledCaptchaProvider = "hcaptcha"
)

// AppResponseSmartWalletConfigUnion contains all possible properties and values
// from [AppResponseSmartWalletConfigEnabled],
// [AppResponseSmartWalletConfigObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppResponseSmartWalletConfigUnion struct {
	Enabled bool `json:"enabled"`
	// This field is from variant [AppResponseSmartWalletConfigObject].
	ConfiguredNetworks []AppResponseSmartWalletConfigObjectConfiguredNetwork `json:"configured_networks"`
	// This field is from variant [AppResponseSmartWalletConfigObject].
	SmartWalletType string `json:"smart_wallet_type"`
	// This field is from variant [AppResponseSmartWalletConfigObject].
	SmartWalletVersion string `json:"smart_wallet_version"`
	JSON               struct {
		Enabled            respjson.Field
		ConfiguredNetworks respjson.Field
		SmartWalletType    respjson.Field
		SmartWalletVersion respjson.Field
		raw                string
	} `json:"-"`
}

func (u AppResponseSmartWalletConfigUnion) AsAppResponseSmartWalletConfigEnabled() (v AppResponseSmartWalletConfigEnabled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResponseSmartWalletConfigUnion) AsAppResponseSmartWalletConfigObject() (v AppResponseSmartWalletConfigObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppResponseSmartWalletConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *AppResponseSmartWalletConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseSmartWalletConfigEnabled struct {
	// Any of false.
	Enabled bool `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseSmartWalletConfigEnabled) RawJSON() string { return r.JSON.raw }
func (r *AppResponseSmartWalletConfigEnabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseSmartWalletConfigObject struct {
	ConfiguredNetworks []AppResponseSmartWalletConfigObjectConfiguredNetwork `json:"configured_networks" api:"required"`
	// Any of true.
	Enabled bool `json:"enabled" api:"required"`
	// Any of "safe", "kernel", "light_account", "biconomy", "coinbase_smart_wallet",
	// "thirdweb".
	SmartWalletType    string `json:"smart_wallet_type" api:"required"`
	SmartWalletVersion string `json:"smart_wallet_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfiguredNetworks respjson.Field
		Enabled            respjson.Field
		SmartWalletType    respjson.Field
		SmartWalletVersion respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseSmartWalletConfigObject) RawJSON() string { return r.JSON.raw }
func (r *AppResponseSmartWalletConfigObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseSmartWalletConfigObjectConfiguredNetwork struct {
	BundlerURL       string                                                              `json:"bundler_url" api:"required"`
	ChainID          string                                                              `json:"chain_id" api:"required"`
	ChainName        string                                                              `json:"chain_name"`
	PaymasterContext AppResponseSmartWalletConfigObjectConfiguredNetworkPaymasterContext `json:"paymaster_context"`
	PaymasterURL     string                                                              `json:"paymaster_url"`
	RpcURL           string                                                              `json:"rpc_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BundlerURL       respjson.Field
		ChainID          respjson.Field
		ChainName        respjson.Field
		PaymasterContext respjson.Field
		PaymasterURL     respjson.Field
		RpcURL           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseSmartWalletConfigObjectConfiguredNetwork) RawJSON() string { return r.JSON.raw }
func (r *AppResponseSmartWalletConfigObjectConfiguredNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseSmartWalletConfigObjectConfiguredNetworkPaymasterContext struct {
	PolicyID string `json:"policy_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PolicyID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseSmartWalletConfigObjectConfiguredNetworkPaymasterContext) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResponseSmartWalletConfigObjectConfiguredNetworkPaymasterContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseFundingConfig struct {
	CrossChainBridgingEnabled  bool                                               `json:"cross_chain_bridging_enabled" api:"required"`
	DefaultRecommendedAmount   string                                             `json:"default_recommended_amount" api:"required"`
	DefaultRecommendedCurrency AppResponseFundingConfigDefaultRecommendedCurrency `json:"default_recommended_currency" api:"required"`
	// Any of "moonpay", "coinbase-onramp", "external".
	Methods                       []string                         `json:"methods" api:"required"`
	Options                       []AppResponseFundingConfigOption `json:"options" api:"required"`
	PromptFundingOnWalletCreation bool                             `json:"prompt_funding_on_wallet_creation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossChainBridgingEnabled     respjson.Field
		DefaultRecommendedAmount      respjson.Field
		DefaultRecommendedCurrency    respjson.Field
		Methods                       respjson.Field
		Options                       respjson.Field
		PromptFundingOnWalletCreation respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseFundingConfig) RawJSON() string { return r.JSON.raw }
func (r *AppResponseFundingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseFundingConfigDefaultRecommendedCurrency struct {
	Chain string `json:"chain" api:"required"`
	// Any of "native-currency", "USDC".
	Asset string `json:"asset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chain       respjson.Field
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseFundingConfigDefaultRecommendedCurrency) RawJSON() string { return r.JSON.raw }
func (r *AppResponseFundingConfigDefaultRecommendedCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseFundingConfigOption struct {
	Method   string `json:"method" api:"required"`
	Provider string `json:"provider" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseFundingConfigOption) RawJSON() string { return r.JSON.raw }
func (r *AppResponseFundingConfigOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResponseTelegramAuthConfig struct {
	BotID               string `json:"bot_id" api:"required"`
	BotName             string `json:"bot_name" api:"required"`
	LinkEnabled         bool   `json:"link_enabled" api:"required"`
	SeamlessAuthEnabled bool   `json:"seamless_auth_enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BotID               respjson.Field
		BotName             respjson.Field
		LinkEnabled         respjson.Field
		SeamlessAuthEnabled respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResponseTelegramAuthConfig) RawJSON() string { return r.JSON.raw }
func (r *AppResponseTelegramAuthConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowlist invite input for an email address.
//
// The properties Type, Value are required.
type EmailInviteInput struct {
	// Any of "email".
	Type  EmailInviteInputType `json:"type,omitzero" api:"required"`
	Value string               `json:"value" api:"required" format:"email"`
	paramObj
}

func (r EmailInviteInput) MarshalJSON() (data []byte, err error) {
	type shadow EmailInviteInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailInviteInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInviteInputType string

const (
	EmailInviteInputTypeEmail EmailInviteInputType = "email"
)

// Allowlist invite input for a wallet address.
//
// The properties Type, Value are required.
type WalletInviteInput struct {
	// Any of "wallet".
	Type  WalletInviteInputType `json:"type,omitzero" api:"required"`
	Value string                `json:"value" api:"required"`
	paramObj
}

func (r WalletInviteInput) MarshalJSON() (data []byte, err error) {
	type shadow WalletInviteInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WalletInviteInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletInviteInputType string

const (
	WalletInviteInputTypeWallet WalletInviteInputType = "wallet"
)

// Allowlist invite input for a phone number.
//
// The properties Type, Value are required.
type PhoneInviteInput struct {
	// Any of "phone".
	Type  PhoneInviteInputType `json:"type,omitzero" api:"required"`
	Value string               `json:"value" api:"required"`
	paramObj
}

func (r PhoneInviteInput) MarshalJSON() (data []byte, err error) {
	type shadow PhoneInviteInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneInviteInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneInviteInputType string

const (
	PhoneInviteInputTypePhone PhoneInviteInputType = "phone"
)

func UserInviteInputOfEmail(value string) UserInviteInputUnion {
	var email EmailInviteInput
	email.Value = value
	return UserInviteInputUnion{OfEmail: &email}
}

func UserInviteInputOfWallet(value string) UserInviteInputUnion {
	var wallet WalletInviteInput
	wallet.Value = value
	return UserInviteInputUnion{OfWallet: &wallet}
}

func UserInviteInputOfPhone(value string) UserInviteInputUnion {
	var phone PhoneInviteInput
	phone.Value = value
	return UserInviteInputUnion{OfPhone: &phone}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserInviteInputUnion struct {
	OfEmail  *EmailInviteInput  `json:",omitzero,inline"`
	OfWallet *WalletInviteInput `json:",omitzero,inline"`
	OfPhone  *PhoneInviteInput  `json:",omitzero,inline"`
	paramUnion
}

func (u UserInviteInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmail, u.OfWallet, u.OfPhone)
}
func (u *UserInviteInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[UserInviteInputUnion](
		"type",
		apijson.Discriminator[EmailInviteInput]("email"),
		apijson.Discriminator[WalletInviteInput]("wallet"),
		apijson.Discriminator[PhoneInviteInput]("phone"),
	)
}

// An allowlist entry for an app.
type AllowlistEntry struct {
	ID         string  `json:"id" api:"required"`
	AcceptedAt float64 `json:"acceptedAt" api:"required"`
	AppID      string  `json:"appId" api:"required"`
	Type       string  `json:"type" api:"required"`
	Value      string  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AcceptedAt  respjson.Field
		AppID       respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowlistEntry) RawJSON() string { return r.JSON.raw }
func (r *AllowlistEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confirmation response for deleting an allowlist entry.
type AllowlistDeletionResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowlistDeletionResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowlistDeletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
