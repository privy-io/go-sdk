// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/apiquery"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/pagination"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
	"github.com/privy-io/go-sdk/shared/constant"
)

// UserService contains methods and other services that help with interacting with
// the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Create a new user with linked accounts. Optionally pre-generate embedded wallets
// for the user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all users in your app.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.Cursor[User], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/users"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all users in your app.
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[User] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a user by user ID.
func (r *UserService) Delete(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a user by user ID.
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Looks up a user by their custom auth ID.
func (r *UserService) GetByCustomAuthID(ctx context.Context, body UserGetByCustomAuthIDParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/custom_auth/id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Discord username.
func (r *UserService) GetByDiscordUsername(ctx context.Context, body UserGetByDiscordUsernameParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/discord/username"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their email address.
func (r *UserService) GetByEmailAddress(ctx context.Context, body UserGetByEmailAddressParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/email/address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Farcaster ID.
func (r *UserService) GetByFarcasterID(ctx context.Context, body UserGetByFarcasterIDParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/farcaster/fid"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Github username.
func (r *UserService) GetByGitHubUsername(ctx context.Context, body UserGetByGitHubUsernameParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/github/username"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their phone number.
func (r *UserService) GetByPhoneNumber(ctx context.Context, body UserGetByPhoneNumberParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/phone/number"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their smart wallet address.
func (r *UserService) GetBySmartWalletAddress(ctx context.Context, body UserGetBySmartWalletAddressParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/smart_wallet/address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Telegram user ID.
func (r *UserService) GetByTelegramUserID(ctx context.Context, body UserGetByTelegramUserIDParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/telegram/telegram_user_id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Telegram username.
func (r *UserService) GetByTelegramUsername(ctx context.Context, body UserGetByTelegramUsernameParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/telegram/username"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Twitter subject.
func (r *UserService) GetByTwitterSubject(ctx context.Context, body UserGetByTwitterSubjectParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/twitter/subject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their Twitter username.
func (r *UserService) GetByTwitterUsername(ctx context.Context, body UserGetByTwitterUsernameParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/twitter/username"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Looks up a user by their wallet address.
func (r *UserService) GetByWalletAddress(ctx context.Context, body UserGetByWalletAddressParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/wallet/address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates an embedded wallet for an existing user.
func (r *UserService) PregenerateWallets(ctx context.Context, userID string, body UserPregenerateWalletsParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/wallets", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search users by search term, emails, phone numbers, or wallet addresses.
func (r *UserService) Search(ctx context.Context, body UserSearchParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Adds custom metadata to a user by user ID.
func (r *UserService) SetCustomMetadata(ctx context.Context, userID string, body UserSetCustomMetadataParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/custom_metadata", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unlinks a user linked account.
func (r *UserService) UnlinkLinkedAccount(ctx context.Context, userID string, body UserUnlinkLinkedAccountParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/accounts/unlink", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type User struct {
	ID string `json:"id,required"`
	// Unix timestamp of when the user was created in milliseconds.
	CreatedAt float64 `json:"created_at,required"`
	// Indicates if the user has accepted the terms of service.
	HasAcceptedTerms bool `json:"has_accepted_terms,required"`
	// Indicates if the user is a guest account user.
	IsGuest        bool                     `json:"is_guest,required"`
	LinkedAccounts []UserLinkedAccountUnion `json:"linked_accounts,required"`
	MfaMethods     []UserMfaMethodUnion     `json:"mfa_methods,required"`
	// Custom metadata associated with the user.
	CustomMetadata map[string]UserCustomMetadataUnion `json:"custom_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		HasAcceptedTerms respjson.Field
		IsGuest          respjson.Field
		LinkedAccounts   respjson.Field
		MfaMethods       respjson.Field
		CustomMetadata   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserLinkedAccountUnion contains all possible properties and values from
// [UserLinkedAccountEmail], [UserLinkedAccountPhone], [UserLinkedAccountCrossApp],
// [UserLinkedAccountAuthorizationKey], [UserLinkedAccountCustomJwt],
// [UserLinkedAccountApple], [UserLinkedAccountDiscord], [UserLinkedAccountGitHub],
// [UserLinkedAccountGoogle], [UserLinkedAccountInstagram],
// [UserLinkedAccountLinkedIn], [UserLinkedAccountSpotify],
// [UserLinkedAccountTiktok], [UserLinkedAccountLine], [UserLinkedAccountTwitch],
// [UserLinkedAccountTwitter], [UserLinkedAccountSmartWallet],
// [UserLinkedAccountPasskey], [UserLinkedAccountFarcaster],
// [UserLinkedAccountTelegram], [UserLinkedAccountEthereum],
// [UserLinkedAccountEthereumEmbeddedWallet], [UserLinkedAccountSolana],
// [UserLinkedAccountSolanaEmbeddedWallet],
// [UserLinkedAccountBitcoinSegwitEmbeddedWallet],
// [UserLinkedAccountBitcoinTaprootEmbeddedWallet].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserLinkedAccountUnion struct {
	Address          string  `json:"address"`
	FirstVerifiedAt  float64 `json:"first_verified_at"`
	LatestVerifiedAt float64 `json:"latest_verified_at"`
	Type             string  `json:"type"`
	VerifiedAt       float64 `json:"verified_at"`
	// This field is from variant [UserLinkedAccountPhone].
	PhoneNumber string `json:"phoneNumber"`
	// This field is from variant [UserLinkedAccountPhone].
	Number string `json:"number"`
	// This field is from variant [UserLinkedAccountCrossApp].
	EmbeddedWallets []UserLinkedAccountCrossAppEmbeddedWallet `json:"embedded_wallets"`
	// This field is from variant [UserLinkedAccountCrossApp].
	ProviderAppID string `json:"provider_app_id"`
	// This field is from variant [UserLinkedAccountCrossApp].
	SmartWallets []UserLinkedAccountCrossAppSmartWallet `json:"smart_wallets"`
	Subject      string                                 `json:"subject"`
	PublicKey    string                                 `json:"public_key"`
	// This field is from variant [UserLinkedAccountCustomJwt].
	CustomUserID string `json:"custom_user_id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	// This field is from variant [UserLinkedAccountLinkedIn].
	VanityName        string `json:"vanity_name"`
	ProfilePictureURL string `json:"profile_picture_url"`
	// This field is from variant [UserLinkedAccountSmartWallet].
	SmartWalletType string `json:"smart_wallet_type"`
	// This field is from variant [UserLinkedAccountSmartWallet].
	SmartWalletVersion string `json:"smart_wallet_version"`
	// This field is from variant [UserLinkedAccountPasskey].
	CredentialID string `json:"credential_id"`
	// This field is from variant [UserLinkedAccountPasskey].
	EnrolledInMfa bool `json:"enrolled_in_mfa"`
	// This field is from variant [UserLinkedAccountPasskey].
	AuthenticatorName string `json:"authenticator_name"`
	// This field is from variant [UserLinkedAccountPasskey].
	CreatedWithBrowser string `json:"created_with_browser"`
	// This field is from variant [UserLinkedAccountPasskey].
	CreatedWithDevice string `json:"created_with_device"`
	// This field is from variant [UserLinkedAccountPasskey].
	CreatedWithOs string `json:"created_with_os"`
	// This field is from variant [UserLinkedAccountFarcaster].
	Fid float64 `json:"fid"`
	// This field is from variant [UserLinkedAccountFarcaster].
	OwnerAddress string `json:"owner_address"`
	// This field is from variant [UserLinkedAccountFarcaster].
	Bio string `json:"bio"`
	// This field is from variant [UserLinkedAccountFarcaster].
	DisplayName string `json:"display_name"`
	// This field is from variant [UserLinkedAccountFarcaster].
	HomepageURL string `json:"homepage_url"`
	// This field is from variant [UserLinkedAccountFarcaster].
	ProfilePicture string `json:"profile_picture"`
	// This field is from variant [UserLinkedAccountFarcaster].
	SignerPublicKey string `json:"signer_public_key"`
	// This field is from variant [UserLinkedAccountTelegram].
	TelegramUserID string `json:"telegram_user_id"`
	// This field is from variant [UserLinkedAccountTelegram].
	FirstName string `json:"first_name"`
	// This field is from variant [UserLinkedAccountTelegram].
	LastName string `json:"last_name"`
	// This field is from variant [UserLinkedAccountTelegram].
	PhotoURL         string  `json:"photo_url"`
	ChainType        string  `json:"chain_type"`
	WalletClient     string  `json:"wallet_client"`
	ChainID          string  `json:"chain_id"`
	ConnectorType    string  `json:"connector_type"`
	WalletClientType string  `json:"wallet_client_type"`
	ID               string  `json:"id"`
	Delegated        bool    `json:"delegated"`
	Imported         bool    `json:"imported"`
	RecoveryMethod   string  `json:"recovery_method"`
	WalletIndex      float64 `json:"wallet_index"`
	JSON             struct {
		Address            respjson.Field
		FirstVerifiedAt    respjson.Field
		LatestVerifiedAt   respjson.Field
		Type               respjson.Field
		VerifiedAt         respjson.Field
		PhoneNumber        respjson.Field
		Number             respjson.Field
		EmbeddedWallets    respjson.Field
		ProviderAppID      respjson.Field
		SmartWallets       respjson.Field
		Subject            respjson.Field
		PublicKey          respjson.Field
		CustomUserID       respjson.Field
		Email              respjson.Field
		Username           respjson.Field
		Name               respjson.Field
		VanityName         respjson.Field
		ProfilePictureURL  respjson.Field
		SmartWalletType    respjson.Field
		SmartWalletVersion respjson.Field
		CredentialID       respjson.Field
		EnrolledInMfa      respjson.Field
		AuthenticatorName  respjson.Field
		CreatedWithBrowser respjson.Field
		CreatedWithDevice  respjson.Field
		CreatedWithOs      respjson.Field
		Fid                respjson.Field
		OwnerAddress       respjson.Field
		Bio                respjson.Field
		DisplayName        respjson.Field
		HomepageURL        respjson.Field
		ProfilePicture     respjson.Field
		SignerPublicKey    respjson.Field
		TelegramUserID     respjson.Field
		FirstName          respjson.Field
		LastName           respjson.Field
		PhotoURL           respjson.Field
		ChainType          respjson.Field
		WalletClient       respjson.Field
		ChainID            respjson.Field
		ConnectorType      respjson.Field
		WalletClientType   respjson.Field
		ID                 respjson.Field
		Delegated          respjson.Field
		Imported           respjson.Field
		RecoveryMethod     respjson.Field
		WalletIndex        respjson.Field
		raw                string
	} `json:"-"`
}

func (u UserLinkedAccountUnion) AsEmail() (v UserLinkedAccountEmail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsPhone() (v UserLinkedAccountPhone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsCrossApp() (v UserLinkedAccountCrossApp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsAuthorizationKey() (v UserLinkedAccountAuthorizationKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsCustomJwt() (v UserLinkedAccountCustomJwt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsApple() (v UserLinkedAccountApple) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsDiscord() (v UserLinkedAccountDiscord) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsGitHub() (v UserLinkedAccountGitHub) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsGoogle() (v UserLinkedAccountGoogle) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsInstagram() (v UserLinkedAccountInstagram) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsLinkedIn() (v UserLinkedAccountLinkedIn) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsSpotify() (v UserLinkedAccountSpotify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsTiktok() (v UserLinkedAccountTiktok) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsLine() (v UserLinkedAccountLine) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsTwitch() (v UserLinkedAccountTwitch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsTwitter() (v UserLinkedAccountTwitter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsSmartWallet() (v UserLinkedAccountSmartWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsPasskey() (v UserLinkedAccountPasskey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsFarcaster() (v UserLinkedAccountFarcaster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsTelegram() (v UserLinkedAccountTelegram) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsEthereum() (v UserLinkedAccountEthereum) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsEthereumEmbeddedWallet() (v UserLinkedAccountEthereumEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsSolana() (v UserLinkedAccountSolana) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsSolanaEmbeddedWallet() (v UserLinkedAccountSolanaEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsBitcoinSegwitEmbeddedWallet() (v UserLinkedAccountBitcoinSegwitEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserLinkedAccountUnion) AsBitcoinTaprootEmbeddedWallet() (v UserLinkedAccountBitcoinTaprootEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserLinkedAccountUnion) RawJSON() string { return u.JSON.raw }

func (r *UserLinkedAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountEmail struct {
	Address          string  `json:"address,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "email".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address          respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountEmail) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountPhone struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PhoneNumber      string  `json:"phoneNumber,required"`
	// Any of "phone".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	Number     string  `json:"number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		PhoneNumber      respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		Number           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountPhone) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountCrossApp struct {
	EmbeddedWallets  []UserLinkedAccountCrossAppEmbeddedWallet `json:"embedded_wallets,required"`
	FirstVerifiedAt  float64                                   `json:"first_verified_at,required"`
	LatestVerifiedAt float64                                   `json:"latest_verified_at,required"`
	ProviderAppID    string                                    `json:"provider_app_id,required"`
	SmartWallets     []UserLinkedAccountCrossAppSmartWallet    `json:"smart_wallets,required"`
	Subject          string                                    `json:"subject,required"`
	// Any of "cross_app".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbeddedWallets  respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		ProviderAppID    respjson.Field
		SmartWallets     respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountCrossApp) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountCrossApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountCrossAppEmbeddedWallet struct {
	Address string `json:"address,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountCrossAppEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountCrossAppEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountCrossAppSmartWallet struct {
	Address string `json:"address,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountCrossAppSmartWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountCrossAppSmartWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountAuthorizationKey struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PublicKey        string  `json:"public_key,required"`
	// Any of "authorization_key".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		PublicKey        respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountAuthorizationKey) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountAuthorizationKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountCustomJwt struct {
	CustomUserID     string  `json:"custom_user_id,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "custom_auth".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomUserID     respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountCustomJwt) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountCustomJwt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountApple struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "apple_oauth".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountApple) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountApple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountDiscord struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "discord_oauth".
	Type       string  `json:"type,required"`
	Username   string  `json:"username,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		Username         respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountDiscord) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountDiscord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountGitHub struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "github_oauth".
	Type       string  `json:"type,required"`
	Username   string  `json:"username,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Name             respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		Username         respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountGitHub) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountGitHub) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountGoogle struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "google_oauth".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Name             respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountGoogle) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountGoogle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountInstagram struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "instagram_oauth".
	Type       string  `json:"type,required"`
	Username   string  `json:"username,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		Username         respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountInstagram) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountInstagram) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountLinkedIn struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "linkedin_oauth".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	Name       string  `json:"name"`
	VanityName string  `json:"vanity_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		Name             respjson.Field
		VanityName       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountLinkedIn) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountLinkedIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountSpotify struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "spotify_oauth".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Name             respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountSpotify) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountSpotify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountTiktok struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "tiktok_oauth".
	Type       string  `json:"type,required"`
	Username   string  `json:"username,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Name             respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		Username         respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountTiktok) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountTiktok) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountLine struct {
	Email             string  `json:"email,required"`
	FirstVerifiedAt   float64 `json:"first_verified_at,required"`
	LatestVerifiedAt  float64 `json:"latest_verified_at,required"`
	Name              string  `json:"name,required"`
	ProfilePictureURL string  `json:"profile_picture_url,required"`
	Subject           string  `json:"subject,required"`
	// Any of "line_oauth".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email             respjson.Field
		FirstVerifiedAt   respjson.Field
		LatestVerifiedAt  respjson.Field
		Name              respjson.Field
		ProfilePictureURL respjson.Field
		Subject           respjson.Field
		Type              respjson.Field
		VerifiedAt        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountLine) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountTwitch struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "twitch_oauth".
	Type       string  `json:"type,required"`
	Username   string  `json:"username,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Subject          respjson.Field
		Type             respjson.Field
		Username         respjson.Field
		VerifiedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountTwitch) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountTwitch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountTwitter struct {
	FirstVerifiedAt   float64 `json:"first_verified_at,required"`
	LatestVerifiedAt  float64 `json:"latest_verified_at,required"`
	Name              string  `json:"name,required"`
	ProfilePictureURL string  `json:"profile_picture_url,required"`
	Subject           string  `json:"subject,required"`
	// Any of "twitter_oauth".
	Type       string  `json:"type,required"`
	Username   string  `json:"username,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt   respjson.Field
		LatestVerifiedAt  respjson.Field
		Name              respjson.Field
		ProfilePictureURL respjson.Field
		Subject           respjson.Field
		Type              respjson.Field
		Username          respjson.Field
		VerifiedAt        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountTwitter) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountTwitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountSmartWallet struct {
	Address          string  `json:"address,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "safe", "kernel", "biconomy", "light_account", "coinbase_smart_wallet",
	// "thirdweb".
	SmartWalletType string `json:"smart_wallet_type,required"`
	// Any of "smart_wallet".
	Type               string  `json:"type,required"`
	VerifiedAt         float64 `json:"verified_at,required"`
	SmartWalletVersion string  `json:"smart_wallet_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address            respjson.Field
		FirstVerifiedAt    respjson.Field
		LatestVerifiedAt   respjson.Field
		SmartWalletType    respjson.Field
		Type               respjson.Field
		VerifiedAt         respjson.Field
		SmartWalletVersion respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountSmartWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountSmartWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountPasskey struct {
	CredentialID     string  `json:"credential_id,required"`
	EnrolledInMfa    bool    `json:"enrolled_in_mfa,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "passkey".
	Type               string  `json:"type,required"`
	VerifiedAt         float64 `json:"verified_at,required"`
	AuthenticatorName  string  `json:"authenticator_name"`
	CreatedWithBrowser string  `json:"created_with_browser"`
	CreatedWithDevice  string  `json:"created_with_device"`
	CreatedWithOs      string  `json:"created_with_os"`
	PublicKey          string  `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CredentialID       respjson.Field
		EnrolledInMfa      respjson.Field
		FirstVerifiedAt    respjson.Field
		LatestVerifiedAt   respjson.Field
		Type               respjson.Field
		VerifiedAt         respjson.Field
		AuthenticatorName  respjson.Field
		CreatedWithBrowser respjson.Field
		CreatedWithDevice  respjson.Field
		CreatedWithOs      respjson.Field
		PublicKey          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountPasskey) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountPasskey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountFarcaster struct {
	Fid              float64 `json:"fid,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	OwnerAddress     string  `json:"owner_address,required"`
	// Any of "farcaster".
	Type              string  `json:"type,required"`
	VerifiedAt        float64 `json:"verified_at,required"`
	Bio               string  `json:"bio"`
	DisplayName       string  `json:"display_name"`
	HomepageURL       string  `json:"homepage_url"`
	ProfilePicture    string  `json:"profile_picture"`
	ProfilePictureURL string  `json:"profile_picture_url"`
	SignerPublicKey   string  `json:"signer_public_key"`
	Username          string  `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fid               respjson.Field
		FirstVerifiedAt   respjson.Field
		LatestVerifiedAt  respjson.Field
		OwnerAddress      respjson.Field
		Type              respjson.Field
		VerifiedAt        respjson.Field
		Bio               respjson.Field
		DisplayName       respjson.Field
		HomepageURL       respjson.Field
		ProfilePicture    respjson.Field
		ProfilePictureURL respjson.Field
		SignerPublicKey   respjson.Field
		Username          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountFarcaster) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountFarcaster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountTelegram struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	TelegramUserID   string  `json:"telegram_user_id,required"`
	// Any of "telegram".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	FirstName  string  `json:"first_name,nullable"`
	LastName   string  `json:"last_name,nullable"`
	PhotoURL   string  `json:"photo_url,nullable"`
	Username   string  `json:"username,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		TelegramUserID   respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		FirstName        respjson.Field
		LastName         respjson.Field
		PhotoURL         respjson.Field
		Username         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountTelegram) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountTelegram) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountEthereum struct {
	Address string `json:"address,required"`
	// Any of "ethereum".
	ChainType        string  `json:"chain_type,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "wallet".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// Any of "unknown".
	WalletClient     string `json:"wallet_client,required"`
	ChainID          string `json:"chain_id"`
	ConnectorType    string `json:"connector_type"`
	WalletClientType string `json:"wallet_client_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address          respjson.Field
		ChainType        respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		WalletClient     respjson.Field
		ChainID          respjson.Field
		ConnectorType    respjson.Field
		WalletClientType respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountEthereum) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountEthereum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountEthereumEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "ethereum".
	ChainType string `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    string  `json:"connector_type,required"`
	Delegated        bool    `json:"delegated,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	Imported         bool    `json:"imported,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod string `json:"recovery_method,required"`
	// Any of "wallet".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// Any of "privy".
	WalletClient string `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType string  `json:"wallet_client_type,required"`
	WalletIndex      float64 `json:"wallet_index,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Address          respjson.Field
		ChainID          respjson.Field
		ChainType        respjson.Field
		ConnectorType    respjson.Field
		Delegated        respjson.Field
		FirstVerifiedAt  respjson.Field
		Imported         respjson.Field
		LatestVerifiedAt respjson.Field
		RecoveryMethod   respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		WalletClient     respjson.Field
		WalletClientType respjson.Field
		WalletIndex      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountEthereumEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountEthereumEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountSolana struct {
	Address string `json:"address,required"`
	// Any of "solana".
	ChainType        string  `json:"chain_type,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "wallet".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// Any of "unknown".
	WalletClient     string `json:"wallet_client,required"`
	ConnectorType    string `json:"connector_type"`
	WalletClientType string `json:"wallet_client_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address          respjson.Field
		ChainType        respjson.Field
		FirstVerifiedAt  respjson.Field
		LatestVerifiedAt respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		WalletClient     respjson.Field
		ConnectorType    respjson.Field
		WalletClientType respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountSolana) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountSolana) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountSolanaEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "solana".
	ChainType string `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    string  `json:"connector_type,required"`
	Delegated        bool    `json:"delegated,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	Imported         bool    `json:"imported,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PublicKey        string  `json:"public_key,required"`
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod string `json:"recovery_method,required"`
	// Any of "wallet".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// Any of "privy".
	WalletClient string `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType string  `json:"wallet_client_type,required"`
	WalletIndex      float64 `json:"wallet_index,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Address          respjson.Field
		ChainID          respjson.Field
		ChainType        respjson.Field
		ConnectorType    respjson.Field
		Delegated        respjson.Field
		FirstVerifiedAt  respjson.Field
		Imported         respjson.Field
		LatestVerifiedAt respjson.Field
		PublicKey        respjson.Field
		RecoveryMethod   respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		WalletClient     respjson.Field
		WalletClientType respjson.Field
		WalletIndex      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountSolanaEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountSolanaEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountBitcoinSegwitEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "bitcoin-segwit".
	ChainType string `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    string  `json:"connector_type,required"`
	Delegated        bool    `json:"delegated,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	Imported         bool    `json:"imported,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PublicKey        string  `json:"public_key,required"`
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod string `json:"recovery_method,required"`
	// Any of "wallet".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// Any of "privy".
	WalletClient string `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType string  `json:"wallet_client_type,required"`
	WalletIndex      float64 `json:"wallet_index,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Address          respjson.Field
		ChainID          respjson.Field
		ChainType        respjson.Field
		ConnectorType    respjson.Field
		Delegated        respjson.Field
		FirstVerifiedAt  respjson.Field
		Imported         respjson.Field
		LatestVerifiedAt respjson.Field
		PublicKey        respjson.Field
		RecoveryMethod   respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		WalletClient     respjson.Field
		WalletClientType respjson.Field
		WalletIndex      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountBitcoinSegwitEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountBitcoinSegwitEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLinkedAccountBitcoinTaprootEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "bitcoin-taproot".
	ChainType string `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    string  `json:"connector_type,required"`
	Delegated        bool    `json:"delegated,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	Imported         bool    `json:"imported,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PublicKey        string  `json:"public_key,required"`
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod string `json:"recovery_method,required"`
	// Any of "wallet".
	Type       string  `json:"type,required"`
	VerifiedAt float64 `json:"verified_at,required"`
	// Any of "privy".
	WalletClient string `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType string  `json:"wallet_client_type,required"`
	WalletIndex      float64 `json:"wallet_index,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Address          respjson.Field
		ChainID          respjson.Field
		ChainType        respjson.Field
		ConnectorType    respjson.Field
		Delegated        respjson.Field
		FirstVerifiedAt  respjson.Field
		Imported         respjson.Field
		LatestVerifiedAt respjson.Field
		PublicKey        respjson.Field
		RecoveryMethod   respjson.Field
		Type             respjson.Field
		VerifiedAt       respjson.Field
		WalletClient     respjson.Field
		WalletClientType respjson.Field
		WalletIndex      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserLinkedAccountBitcoinTaprootEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *UserLinkedAccountBitcoinTaprootEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserMfaMethodUnion contains all possible properties and values from
// [UserMfaMethodPasskey], [UserMfaMethodSMS], [UserMfaMethodTotp].
//
// Use the [UserMfaMethodUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserMfaMethodUnion struct {
	// Any of "passkey", "sms", "totp".
	Type       string  `json:"type"`
	VerifiedAt float64 `json:"verified_at"`
	JSON       struct {
		Type       respjson.Field
		VerifiedAt respjson.Field
		raw        string
	} `json:"-"`
}

// anyUserMfaMethod is implemented by each variant of [UserMfaMethodUnion] to add
// type safety for the return type of [UserMfaMethodUnion.AsAny]
type anyUserMfaMethod interface {
	implUserMfaMethodUnion()
}

func (UserMfaMethodPasskey) implUserMfaMethodUnion() {}
func (UserMfaMethodSMS) implUserMfaMethodUnion()     {}
func (UserMfaMethodTotp) implUserMfaMethodUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UserMfaMethodUnion.AsAny().(type) {
//	case privyapiclient.UserMfaMethodPasskey:
//	case privyapiclient.UserMfaMethodSMS:
//	case privyapiclient.UserMfaMethodTotp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UserMfaMethodUnion) AsAny() anyUserMfaMethod {
	switch u.Type {
	case "passkey":
		return u.AsPasskey()
	case "sms":
		return u.AsSMS()
	case "totp":
		return u.AsTotp()
	}
	return nil
}

func (u UserMfaMethodUnion) AsPasskey() (v UserMfaMethodPasskey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserMfaMethodUnion) AsSMS() (v UserMfaMethodSMS) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserMfaMethodUnion) AsTotp() (v UserMfaMethodTotp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserMfaMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *UserMfaMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMfaMethodPasskey struct {
	Type       constant.Passkey `json:"type,required"`
	VerifiedAt float64          `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMfaMethodPasskey) RawJSON() string { return r.JSON.raw }
func (r *UserMfaMethodPasskey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMfaMethodSMS struct {
	Type       constant.SMS `json:"type,required"`
	VerifiedAt float64      `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMfaMethodSMS) RawJSON() string { return r.JSON.raw }
func (r *UserMfaMethodSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMfaMethodTotp struct {
	Type       constant.Totp `json:"type,required"`
	VerifiedAt float64       `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMfaMethodTotp) RawJSON() string { return r.JSON.raw }
func (r *UserMfaMethodTotp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserCustomMetadataUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type UserCustomMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u UserCustomMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserCustomMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserCustomMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserCustomMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *UserCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	LinkedAccounts []UserNewParamsLinkedAccountUnion `json:"linked_accounts,omitzero,required"`
	// Custom metadata associated with the user.
	CustomMetadata map[string]UserNewParamsCustomMetadataUnion `json:"custom_metadata,omitzero"`
	// Wallets to create for the user.
	Wallets []UserNewParamsWallet `json:"wallets,omitzero"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserNewParamsLinkedAccountUnion struct {
	OfWallet         *UserNewParamsLinkedAccountWallet         `json:",omitzero,inline"`
	OfEmail          *UserNewParamsLinkedAccountEmail          `json:",omitzero,inline"`
	OfPhone          *UserNewParamsLinkedAccountPhone          `json:",omitzero,inline"`
	OfGoogleOAuth    *UserNewParamsLinkedAccountGoogleOAuth    `json:",omitzero,inline"`
	OfTwitterOAuth   *UserNewParamsLinkedAccountTwitterOAuth   `json:",omitzero,inline"`
	OfDiscordOAuth   *UserNewParamsLinkedAccountDiscordOAuth   `json:",omitzero,inline"`
	OfGitHubOAuth    *UserNewParamsLinkedAccountGitHubOAuth    `json:",omitzero,inline"`
	OfSpotifyOAuth   *UserNewParamsLinkedAccountSpotifyOAuth   `json:",omitzero,inline"`
	OfInstagramOAuth *UserNewParamsLinkedAccountInstagramOAuth `json:",omitzero,inline"`
	OfTiktokOAuth    *UserNewParamsLinkedAccountTiktokOAuth    `json:",omitzero,inline"`
	OfLineOAuth      *UserNewParamsLinkedAccountLineOAuth      `json:",omitzero,inline"`
	OfTwitchOAuth    *UserNewParamsLinkedAccountTwitchOAuth    `json:",omitzero,inline"`
	OfAppleOAuth     *UserNewParamsLinkedAccountAppleOAuth     `json:",omitzero,inline"`
	OfLinkedinOAuth  *UserNewParamsLinkedAccountLinkedinOAuth  `json:",omitzero,inline"`
	OfFarcaster      *UserNewParamsLinkedAccountFarcaster      `json:",omitzero,inline"`
	OfTelegram       *UserNewParamsLinkedAccountTelegram       `json:",omitzero,inline"`
	OfCustomAuth     *UserNewParamsLinkedAccountCustomAuth     `json:",omitzero,inline"`
	paramUnion
}

func (u UserNewParamsLinkedAccountUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWallet,
		u.OfEmail,
		u.OfPhone,
		u.OfGoogleOAuth,
		u.OfTwitterOAuth,
		u.OfDiscordOAuth,
		u.OfGitHubOAuth,
		u.OfSpotifyOAuth,
		u.OfInstagramOAuth,
		u.OfTiktokOAuth,
		u.OfLineOAuth,
		u.OfTwitchOAuth,
		u.OfAppleOAuth,
		u.OfLinkedinOAuth,
		u.OfFarcaster,
		u.OfTelegram,
		u.OfCustomAuth)
}
func (u *UserNewParamsLinkedAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UserNewParamsLinkedAccountUnion) asAny() any {
	if !param.IsOmitted(u.OfWallet) {
		return u.OfWallet
	} else if !param.IsOmitted(u.OfEmail) {
		return u.OfEmail
	} else if !param.IsOmitted(u.OfPhone) {
		return u.OfPhone
	} else if !param.IsOmitted(u.OfGoogleOAuth) {
		return u.OfGoogleOAuth
	} else if !param.IsOmitted(u.OfTwitterOAuth) {
		return u.OfTwitterOAuth
	} else if !param.IsOmitted(u.OfDiscordOAuth) {
		return u.OfDiscordOAuth
	} else if !param.IsOmitted(u.OfGitHubOAuth) {
		return u.OfGitHubOAuth
	} else if !param.IsOmitted(u.OfSpotifyOAuth) {
		return u.OfSpotifyOAuth
	} else if !param.IsOmitted(u.OfInstagramOAuth) {
		return u.OfInstagramOAuth
	} else if !param.IsOmitted(u.OfTiktokOAuth) {
		return u.OfTiktokOAuth
	} else if !param.IsOmitted(u.OfLineOAuth) {
		return u.OfLineOAuth
	} else if !param.IsOmitted(u.OfTwitchOAuth) {
		return u.OfTwitchOAuth
	} else if !param.IsOmitted(u.OfAppleOAuth) {
		return u.OfAppleOAuth
	} else if !param.IsOmitted(u.OfLinkedinOAuth) {
		return u.OfLinkedinOAuth
	} else if !param.IsOmitted(u.OfFarcaster) {
		return u.OfFarcaster
	} else if !param.IsOmitted(u.OfTelegram) {
		return u.OfTelegram
	} else if !param.IsOmitted(u.OfCustomAuth) {
		return u.OfCustomAuth
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetChainType() *string {
	if vt := u.OfWallet; vt != nil {
		return &vt.ChainType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetNumber() *string {
	if vt := u.OfPhone; vt != nil {
		return &vt.Number
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetVanityName() *string {
	if vt := u.OfLinkedinOAuth; vt != nil && vt.VanityName.Valid() {
		return &vt.VanityName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetFid() *int64 {
	if vt := u.OfFarcaster; vt != nil {
		return &vt.Fid
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetOwnerAddress() *string {
	if vt := u.OfFarcaster; vt != nil {
		return &vt.OwnerAddress
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetBio() *string {
	if vt := u.OfFarcaster; vt != nil && vt.Bio.Valid() {
		return &vt.Bio.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetDisplayName() *string {
	if vt := u.OfFarcaster; vt != nil && vt.DisplayName.Valid() {
		return &vt.DisplayName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetHomepageURL() *string {
	if vt := u.OfFarcaster; vt != nil && vt.HomepageURL.Valid() {
		return &vt.HomepageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetTelegramUserID() *string {
	if vt := u.OfTelegram; vt != nil {
		return &vt.TelegramUserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetFirstName() *string {
	if vt := u.OfTelegram; vt != nil && vt.FirstName.Valid() {
		return &vt.FirstName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetLastName() *string {
	if vt := u.OfTelegram; vt != nil && vt.LastName.Valid() {
		return &vt.LastName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetPhotoURL() *string {
	if vt := u.OfTelegram; vt != nil && vt.PhotoURL.Valid() {
		return &vt.PhotoURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetCustomUserID() *string {
	if vt := u.OfCustomAuth; vt != nil {
		return &vt.CustomUserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetAddress() *string {
	if vt := u.OfWallet; vt != nil {
		return (*string)(&vt.Address)
	} else if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Address)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetType() *string {
	if vt := u.OfWallet; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGoogleOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTwitterOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDiscordOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGitHubOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSpotifyOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInstagramOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTiktokOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLineOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTwitchOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAppleOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLinkedinOAuth; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFarcaster; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTelegram; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCustomAuth; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetEmail() *string {
	if vt := u.OfGoogleOAuth; vt != nil {
		return (*string)(&vt.Email)
	} else if vt := u.OfDiscordOAuth; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	} else if vt := u.OfGitHubOAuth; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	} else if vt := u.OfSpotifyOAuth; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	} else if vt := u.OfLineOAuth; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	} else if vt := u.OfAppleOAuth; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	} else if vt := u.OfLinkedinOAuth; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetName() *string {
	if vt := u.OfGoogleOAuth; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfTwitterOAuth; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfGitHubOAuth; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSpotifyOAuth; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfTiktokOAuth; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfLineOAuth; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfLinkedinOAuth; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetSubject() *string {
	if vt := u.OfGoogleOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfTwitterOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfDiscordOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfGitHubOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfSpotifyOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfInstagramOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfTiktokOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfLineOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfTwitchOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfAppleOAuth; vt != nil {
		return (*string)(&vt.Subject)
	} else if vt := u.OfLinkedinOAuth; vt != nil {
		return (*string)(&vt.Subject)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetUsername() *string {
	if vt := u.OfTwitterOAuth; vt != nil {
		return (*string)(&vt.Username)
	} else if vt := u.OfDiscordOAuth; vt != nil {
		return (*string)(&vt.Username)
	} else if vt := u.OfGitHubOAuth; vt != nil {
		return (*string)(&vt.Username)
	} else if vt := u.OfInstagramOAuth; vt != nil {
		return (*string)(&vt.Username)
	} else if vt := u.OfTiktokOAuth; vt != nil {
		return (*string)(&vt.Username)
	} else if vt := u.OfTwitchOAuth; vt != nil && vt.Username.Valid() {
		return &vt.Username.Value
	} else if vt := u.OfFarcaster; vt != nil && vt.Username.Valid() {
		return &vt.Username.Value
	} else if vt := u.OfTelegram; vt != nil && vt.Username.Valid() {
		return &vt.Username.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UserNewParamsLinkedAccountUnion) GetProfilePictureURL() *string {
	if vt := u.OfTwitterOAuth; vt != nil && vt.ProfilePictureURL.Valid() {
		return &vt.ProfilePictureURL.Value
	} else if vt := u.OfLineOAuth; vt != nil && vt.ProfilePictureURL.Valid() {
		return &vt.ProfilePictureURL.Value
	} else if vt := u.OfFarcaster; vt != nil && vt.ProfilePictureURL.Valid() {
		return &vt.ProfilePictureURL.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UserNewParamsLinkedAccountUnion](
		"type",
		apijson.Discriminator[UserNewParamsLinkedAccountWallet]("wallet"),
		apijson.Discriminator[UserNewParamsLinkedAccountEmail]("email"),
		apijson.Discriminator[UserNewParamsLinkedAccountPhone]("phone"),
		apijson.Discriminator[UserNewParamsLinkedAccountGoogleOAuth]("google_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountTwitterOAuth]("twitter_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountDiscordOAuth]("discord_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountGitHubOAuth]("github_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountSpotifyOAuth]("spotify_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountInstagramOAuth]("instagram_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountTiktokOAuth]("tiktok_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountLineOAuth]("line_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountTwitchOAuth]("twitch_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountAppleOAuth]("apple_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountLinkedinOAuth]("linkedin_oauth"),
		apijson.Discriminator[UserNewParamsLinkedAccountFarcaster]("farcaster"),
		apijson.Discriminator[UserNewParamsLinkedAccountTelegram]("telegram"),
		apijson.Discriminator[UserNewParamsLinkedAccountCustomAuth]("custom_auth"),
	)
}

// The properties Address, ChainType, Type are required.
type UserNewParamsLinkedAccountWallet struct {
	Address string `json:"address,required"`
	// Any of "ethereum", "solana".
	ChainType string `json:"chain_type,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "wallet".
	Type constant.Wallet `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountWallet) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountWallet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserNewParamsLinkedAccountWallet](
		"chain_type", "ethereum", "solana",
	)
}

// The properties Address, Type are required.
type UserNewParamsLinkedAccountEmail struct {
	Address string `json:"address,required" format:"email"`
	// This field can be elided, and will marshal its zero value as "email".
	Type constant.Email `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountEmail) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Number, Type are required.
type UserNewParamsLinkedAccountPhone struct {
	Number string `json:"number,required"`
	// This field can be elided, and will marshal its zero value as "phone".
	Type constant.Phone `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountPhone) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Email, Name, Subject, Type are required.
type UserNewParamsLinkedAccountGoogleOAuth struct {
	Email   string `json:"email,required" format:"email"`
	Name    string `json:"name,required"`
	Subject string `json:"subject,required"`
	// This field can be elided, and will marshal its zero value as "google_oauth".
	Type constant.GoogleOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountGoogleOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountGoogleOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountGoogleOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Subject, Type, Username are required.
type UserNewParamsLinkedAccountTwitterOAuth struct {
	Name              string            `json:"name,required"`
	Subject           string            `json:"subject,required"`
	Username          string            `json:"username,required"`
	ProfilePictureURL param.Opt[string] `json:"profile_picture_url,omitzero" format:"uri"`
	// This field can be elided, and will marshal its zero value as "twitter_oauth".
	Type constant.TwitterOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountTwitterOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountTwitterOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountTwitterOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type, Username are required.
type UserNewParamsLinkedAccountDiscordOAuth struct {
	Subject  string            `json:"subject,required"`
	Username string            `json:"username,required"`
	Email    param.Opt[string] `json:"email,omitzero" format:"email"`
	// This field can be elided, and will marshal its zero value as "discord_oauth".
	Type constant.DiscordOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountDiscordOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountDiscordOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountDiscordOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type, Username are required.
type UserNewParamsLinkedAccountGitHubOAuth struct {
	Subject  string            `json:"subject,required"`
	Username string            `json:"username,required"`
	Email    param.Opt[string] `json:"email,omitzero" format:"email"`
	Name     param.Opt[string] `json:"name,omitzero"`
	// This field can be elided, and will marshal its zero value as "github_oauth".
	Type constant.GitHubOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountGitHubOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountGitHubOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountGitHubOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type are required.
type UserNewParamsLinkedAccountSpotifyOAuth struct {
	Subject string            `json:"subject,required"`
	Email   param.Opt[string] `json:"email,omitzero" format:"email"`
	Name    param.Opt[string] `json:"name,omitzero"`
	// This field can be elided, and will marshal its zero value as "spotify_oauth".
	Type constant.SpotifyOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountSpotifyOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountSpotifyOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountSpotifyOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type, Username are required.
type UserNewParamsLinkedAccountInstagramOAuth struct {
	Subject  string `json:"subject,required"`
	Username string `json:"username,required"`
	// This field can be elided, and will marshal its zero value as "instagram_oauth".
	Type constant.InstagramOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountInstagramOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountInstagramOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountInstagramOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Subject, Type, Username are required.
type UserNewParamsLinkedAccountTiktokOAuth struct {
	Name     param.Opt[string] `json:"name,omitzero,required"`
	Subject  string            `json:"subject,required"`
	Username string            `json:"username,required"`
	// This field can be elided, and will marshal its zero value as "tiktok_oauth".
	Type constant.TiktokOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountTiktokOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountTiktokOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountTiktokOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type are required.
type UserNewParamsLinkedAccountLineOAuth struct {
	Subject           string            `json:"subject,required"`
	Email             param.Opt[string] `json:"email,omitzero" format:"email"`
	Name              param.Opt[string] `json:"name,omitzero"`
	ProfilePictureURL param.Opt[string] `json:"profile_picture_url,omitzero" format:"uri"`
	// This field can be elided, and will marshal its zero value as "line_oauth".
	Type constant.LineOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountLineOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountLineOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountLineOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type are required.
type UserNewParamsLinkedAccountTwitchOAuth struct {
	Subject  string            `json:"subject,required"`
	Username param.Opt[string] `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "twitch_oauth".
	Type constant.TwitchOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountTwitchOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountTwitchOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountTwitchOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type are required.
type UserNewParamsLinkedAccountAppleOAuth struct {
	Subject string            `json:"subject,required"`
	Email   param.Opt[string] `json:"email,omitzero" format:"email"`
	// This field can be elided, and will marshal its zero value as "apple_oauth".
	Type constant.AppleOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountAppleOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountAppleOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountAppleOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, Type are required.
type UserNewParamsLinkedAccountLinkedinOAuth struct {
	Subject    string            `json:"subject,required"`
	Email      param.Opt[string] `json:"email,omitzero" format:"email"`
	Name       param.Opt[string] `json:"name,omitzero"`
	VanityName param.Opt[string] `json:"vanityName,omitzero"`
	// This field can be elided, and will marshal its zero value as "linkedin_oauth".
	Type constant.LinkedinOAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountLinkedinOAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountLinkedinOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountLinkedinOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Fid, OwnerAddress, Type are required.
type UserNewParamsLinkedAccountFarcaster struct {
	Fid               int64             `json:"fid,required"`
	OwnerAddress      string            `json:"owner_address,required"`
	Bio               param.Opt[string] `json:"bio,omitzero"`
	DisplayName       param.Opt[string] `json:"display_name,omitzero"`
	HomepageURL       param.Opt[string] `json:"homepage_url,omitzero"`
	ProfilePictureURL param.Opt[string] `json:"profile_picture_url,omitzero"`
	Username          param.Opt[string] `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "farcaster".
	Type constant.Farcaster `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountFarcaster) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountFarcaster
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountFarcaster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties TelegramUserID, Type are required.
type UserNewParamsLinkedAccountTelegram struct {
	TelegramUserID string            `json:"telegram_user_id,required"`
	FirstName      param.Opt[string] `json:"first_name,omitzero"`
	LastName       param.Opt[string] `json:"last_name,omitzero"`
	PhotoURL       param.Opt[string] `json:"photo_url,omitzero"`
	Username       param.Opt[string] `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "telegram".
	Type constant.Telegram `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountTelegram) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountTelegram
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountTelegram) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomUserID, Type are required.
type UserNewParamsLinkedAccountCustomAuth struct {
	CustomUserID string `json:"custom_user_id,required"`
	// This field can be elided, and will marshal its zero value as "custom_auth".
	Type constant.CustomAuth `json:"type,required"`
	paramObj
}

func (r UserNewParamsLinkedAccountCustomAuth) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsLinkedAccountCustomAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsLinkedAccountCustomAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserNewParamsCustomMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u UserNewParamsCustomMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *UserNewParamsCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UserNewParamsCustomMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property ChainType is required.
type UserNewParamsWallet struct {
	// Chain type of the wallet
	//
	// Any of "solana", "ethereum", "cosmos", "stellar", "sui", "tron",
	// "bitcoin-segwit", "near", "spark", "ton", "starknet", "movement".
	ChainType string `json:"chain_type,omitzero,required"`
	// Create a smart wallet with this wallet as the signer. Only supported for wallets
	// with `chain_type: "ethereum"`.
	CreateSmartWallet param.Opt[bool] `json:"create_smart_wallet,omitzero"`
	// Additional signers for the wallet.
	AdditionalSigners []UserNewParamsWalletAdditionalSigner `json:"additional_signers,omitzero"`
	// Policy IDs to enforce on the wallet. Currently, only one policy is supported per
	// wallet.
	PolicyIDs []string `json:"policy_ids,omitzero"`
	paramObj
}

func (r UserNewParamsWallet) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsWallet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserNewParamsWallet](
		"chain_type", "solana", "ethereum", "cosmos", "stellar", "sui", "tron", "bitcoin-segwit", "near", "spark", "ton", "starknet", "movement",
	)
}

// The property SignerID is required.
type UserNewParamsWalletAdditionalSigner struct {
	// The key quorum ID for the signer.
	SignerID string `json:"signer_id,required"`
	// The array of policy IDs that will be applied to wallet requests. If specified,
	// this will override the base policy IDs set on the wallet. Currently, only one
	// policy is supported per signer.
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero"`
	paramObj
}

func (r UserNewParamsWalletAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParamsWalletAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParamsWalletAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	Limit  param.Opt[float64] `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]  `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGetByCustomAuthIDParams struct {
	CustomUserID string `json:"custom_user_id,required"`
	paramObj
}

func (r UserGetByCustomAuthIDParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByCustomAuthIDParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByCustomAuthIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByDiscordUsernameParams struct {
	Username string `json:"username,required"`
	paramObj
}

func (r UserGetByDiscordUsernameParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByDiscordUsernameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByDiscordUsernameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByEmailAddressParams struct {
	Address string `json:"address,required" format:"email"`
	paramObj
}

func (r UserGetByEmailAddressParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByEmailAddressParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByEmailAddressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByFarcasterIDParams struct {
	Fid float64 `json:"fid,required"`
	paramObj
}

func (r UserGetByFarcasterIDParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByFarcasterIDParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByFarcasterIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByGitHubUsernameParams struct {
	Username string `json:"username,required"`
	paramObj
}

func (r UserGetByGitHubUsernameParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByGitHubUsernameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByGitHubUsernameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByPhoneNumberParams struct {
	Number string `json:"number,required"`
	paramObj
}

func (r UserGetByPhoneNumberParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByPhoneNumberParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByPhoneNumberParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetBySmartWalletAddressParams struct {
	Address string `json:"address,required"`
	paramObj
}

func (r UserGetBySmartWalletAddressParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetBySmartWalletAddressParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetBySmartWalletAddressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByTelegramUserIDParams struct {
	TelegramUserID string `json:"telegram_user_id,required"`
	paramObj
}

func (r UserGetByTelegramUserIDParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByTelegramUserIDParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByTelegramUserIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByTelegramUsernameParams struct {
	Username string `json:"username,required"`
	paramObj
}

func (r UserGetByTelegramUsernameParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByTelegramUsernameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByTelegramUsernameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByTwitterSubjectParams struct {
	Subject string `json:"subject,required"`
	paramObj
}

func (r UserGetByTwitterSubjectParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByTwitterSubjectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByTwitterSubjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByTwitterUsernameParams struct {
	Username string `json:"username,required"`
	paramObj
}

func (r UserGetByTwitterUsernameParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByTwitterUsernameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByTwitterUsernameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetByWalletAddressParams struct {
	Address string `json:"address,required"`
	paramObj
}

func (r UserGetByWalletAddressParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGetByWalletAddressParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGetByWalletAddressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPregenerateWalletsParams struct {
	Wallets []UserPregenerateWalletsParamsWallet `json:"wallets,omitzero,required"`
	paramObj
}

func (r UserPregenerateWalletsParams) MarshalJSON() (data []byte, err error) {
	type shadow UserPregenerateWalletsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPregenerateWalletsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ChainType is required.
type UserPregenerateWalletsParamsWallet struct {
	// Any of "cosmos", "stellar", "sui", "aptos", "movement", "tron",
	// "bitcoin-segwit", "near", "ton", "starknet", "spark", "solana", "ethereum".
	ChainType         UserPregenerateWalletsParamsWalletChainType          `json:"chain_type,omitzero,required"`
	CreateSmartWallet param.Opt[bool]                                      `json:"create_smart_wallet,omitzero"`
	AdditionalSigners []UserPregenerateWalletsParamsWalletAdditionalSigner `json:"additional_signers,omitzero"`
	PolicyIDs         []string                                             `json:"policy_ids,omitzero"`
	paramObj
}

func (r UserPregenerateWalletsParamsWallet) MarshalJSON() (data []byte, err error) {
	type shadow UserPregenerateWalletsParamsWallet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPregenerateWalletsParamsWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPregenerateWalletsParamsWalletChainType string

const (
	UserPregenerateWalletsParamsWalletChainTypeCosmos        UserPregenerateWalletsParamsWalletChainType = "cosmos"
	UserPregenerateWalletsParamsWalletChainTypeStellar       UserPregenerateWalletsParamsWalletChainType = "stellar"
	UserPregenerateWalletsParamsWalletChainTypeSui           UserPregenerateWalletsParamsWalletChainType = "sui"
	UserPregenerateWalletsParamsWalletChainTypeAptos         UserPregenerateWalletsParamsWalletChainType = "aptos"
	UserPregenerateWalletsParamsWalletChainTypeMovement      UserPregenerateWalletsParamsWalletChainType = "movement"
	UserPregenerateWalletsParamsWalletChainTypeTron          UserPregenerateWalletsParamsWalletChainType = "tron"
	UserPregenerateWalletsParamsWalletChainTypeBitcoinSegwit UserPregenerateWalletsParamsWalletChainType = "bitcoin-segwit"
	UserPregenerateWalletsParamsWalletChainTypeNear          UserPregenerateWalletsParamsWalletChainType = "near"
	UserPregenerateWalletsParamsWalletChainTypeTon           UserPregenerateWalletsParamsWalletChainType = "ton"
	UserPregenerateWalletsParamsWalletChainTypeStarknet      UserPregenerateWalletsParamsWalletChainType = "starknet"
	UserPregenerateWalletsParamsWalletChainTypeSpark         UserPregenerateWalletsParamsWalletChainType = "spark"
	UserPregenerateWalletsParamsWalletChainTypeSolana        UserPregenerateWalletsParamsWalletChainType = "solana"
	UserPregenerateWalletsParamsWalletChainTypeEthereum      UserPregenerateWalletsParamsWalletChainType = "ethereum"
)

// The property SignerID is required.
type UserPregenerateWalletsParamsWalletAdditionalSigner struct {
	SignerID          string   `json:"signer_id,required"`
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero"`
	paramObj
}

func (r UserPregenerateWalletsParamsWalletAdditionalSigner) MarshalJSON() (data []byte, err error) {
	type shadow UserPregenerateWalletsParamsWalletAdditionalSigner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPregenerateWalletsParamsWalletAdditionalSigner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserSearchParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfSearchTerm *UserSearchParamsBodySearchTerm `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfObject *UserSearchParamsBodyObject `json:",inline"`

	paramObj
}

func (u UserSearchParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSearchTerm, u.OfObject)
}
func (r *UserSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SearchTerm is required.
type UserSearchParamsBodySearchTerm struct {
	SearchTerm string `json:"search_term,required"`
	paramObj
}

func (r UserSearchParamsBodySearchTerm) MarshalJSON() (data []byte, err error) {
	type shadow UserSearchParamsBodySearchTerm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserSearchParamsBodySearchTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Emails, PhoneNumbers, WalletAddresses are required.
type UserSearchParamsBodyObject struct {
	Emails          []string `json:"emails,omitzero,required" format:"email"`
	PhoneNumbers    []string `json:"phoneNumbers,omitzero,required"`
	WalletAddresses []string `json:"walletAddresses,omitzero,required"`
	paramObj
}

func (r UserSearchParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow UserSearchParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserSearchParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserSetCustomMetadataParams struct {
	// Custom metadata associated with the user.
	CustomMetadata map[string]UserSetCustomMetadataParamsCustomMetadataUnion `json:"custom_metadata,omitzero,required"`
	paramObj
}

func (r UserSetCustomMetadataParams) MarshalJSON() (data []byte, err error) {
	type shadow UserSetCustomMetadataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserSetCustomMetadataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserSetCustomMetadataParamsCustomMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u UserSetCustomMetadataParamsCustomMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *UserSetCustomMetadataParamsCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UserSetCustomMetadataParamsCustomMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type UserUnlinkLinkedAccountParams struct {
	Handle   string                            `json:"handle,required"`
	Type     UserUnlinkLinkedAccountParamsType `json:"type,omitzero,required"`
	Provider param.Opt[string]                 `json:"provider,omitzero"`
	paramObj
}

func (r UserUnlinkLinkedAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUnlinkLinkedAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUnlinkLinkedAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUnlinkLinkedAccountParamsType string

const (
	UserUnlinkLinkedAccountParamsTypeEmail          UserUnlinkLinkedAccountParamsType = "email"
	UserUnlinkLinkedAccountParamsTypeWallet         UserUnlinkLinkedAccountParamsType = "wallet"
	UserUnlinkLinkedAccountParamsTypeSmartWallet    UserUnlinkLinkedAccountParamsType = "smart_wallet"
	UserUnlinkLinkedAccountParamsTypeFarcaster      UserUnlinkLinkedAccountParamsType = "farcaster"
	UserUnlinkLinkedAccountParamsTypePasskey        UserUnlinkLinkedAccountParamsType = "passkey"
	UserUnlinkLinkedAccountParamsTypePhone          UserUnlinkLinkedAccountParamsType = "phone"
	UserUnlinkLinkedAccountParamsTypeGoogleOAuth    UserUnlinkLinkedAccountParamsType = "google_oauth"
	UserUnlinkLinkedAccountParamsTypeDiscordOAuth   UserUnlinkLinkedAccountParamsType = "discord_oauth"
	UserUnlinkLinkedAccountParamsTypeTwitterOAuth   UserUnlinkLinkedAccountParamsType = "twitter_oauth"
	UserUnlinkLinkedAccountParamsTypeGitHubOAuth    UserUnlinkLinkedAccountParamsType = "github_oauth"
	UserUnlinkLinkedAccountParamsTypeLinkedinOAuth  UserUnlinkLinkedAccountParamsType = "linkedin_oauth"
	UserUnlinkLinkedAccountParamsTypeAppleOAuth     UserUnlinkLinkedAccountParamsType = "apple_oauth"
	UserUnlinkLinkedAccountParamsTypeSpotifyOAuth   UserUnlinkLinkedAccountParamsType = "spotify_oauth"
	UserUnlinkLinkedAccountParamsTypeInstagramOAuth UserUnlinkLinkedAccountParamsType = "instagram_oauth"
	UserUnlinkLinkedAccountParamsTypeTiktokOAuth    UserUnlinkLinkedAccountParamsType = "tiktok_oauth"
	UserUnlinkLinkedAccountParamsTypeLineOAuth      UserUnlinkLinkedAccountParamsType = "line_oauth"
	UserUnlinkLinkedAccountParamsTypeTwitchOAuth    UserUnlinkLinkedAccountParamsType = "twitch_oauth"
	UserUnlinkLinkedAccountParamsTypeCustomAuth     UserUnlinkLinkedAccountParamsType = "custom_auth"
	UserUnlinkLinkedAccountParamsTypeTelegram       UserUnlinkLinkedAccountParamsType = "telegram"
	UserUnlinkLinkedAccountParamsTypeCrossApp       UserUnlinkLinkedAccountParamsType = "cross_app"
	UserUnlinkLinkedAccountParamsTypeGuest          UserUnlinkLinkedAccountParamsType = "guest"
)
