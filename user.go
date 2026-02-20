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
	path := fmt.Sprintf("v1/users/%s", url.PathEscape(userID))
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
	path := fmt.Sprintf("v1/users/%s", url.PathEscape(userID))
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
	path := fmt.Sprintf("v1/users/%s/wallets", url.PathEscape(userID))
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
	path := fmt.Sprintf("v1/users/%s/custom_metadata", url.PathEscape(userID))
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
	path := fmt.Sprintf("v1/users/%s/accounts/unlink", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A Privy user object.
type User struct {
	ID string `json:"id,required"`
	// Unix timestamp of when the user was created in milliseconds.
	CreatedAt float64 `json:"created_at,required"`
	// Indicates if the user has accepted the terms of service.
	HasAcceptedTerms bool `json:"has_accepted_terms,required"`
	// Indicates if the user is a guest account user.
	IsGuest        bool                   `json:"is_guest,required"`
	LinkedAccounts []LinkedAccountUnion   `json:"linked_accounts,required"`
	MfaMethods     []LinkedMfaMethodUnion `json:"mfa_methods,required"`
	// Custom metadata associated with the user.
	CustomMetadata CustomMetadata `json:"custom_metadata"`
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

// An email account linked to the user.
type LinkedAccountEmail struct {
	Address          string  `json:"address,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "email".
	Type       LinkedAccountEmailType `json:"type,required"`
	VerifiedAt float64                `json:"verified_at,required"`
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
func (r LinkedAccountEmail) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountEmailType string

const (
	LinkedAccountEmailTypeEmail LinkedAccountEmailType = "email"
)

// A phone number account linked to the user.
type LinkedAccountPhone struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PhoneNumber      string  `json:"phoneNumber,required"`
	// Any of "phone".
	Type       LinkedAccountPhoneType `json:"type,required"`
	VerifiedAt float64                `json:"verified_at,required"`
	Number     string                 `json:"number"`
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
func (r LinkedAccountPhone) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountPhoneType string

const (
	LinkedAccountPhoneTypePhone LinkedAccountPhoneType = "phone"
)

// An Ethereum wallet account linked to the user.
type LinkedAccountEthereum struct {
	Address string `json:"address,required"`
	// Any of "ethereum".
	ChainType        LinkedAccountEthereumChainType `json:"chain_type,required"`
	FirstVerifiedAt  float64                        `json:"first_verified_at,required"`
	LatestVerifiedAt float64                        `json:"latest_verified_at,required"`
	// Any of "wallet".
	Type       LinkedAccountEthereumType `json:"type,required"`
	VerifiedAt float64                   `json:"verified_at,required"`
	// Any of "unknown".
	WalletClient     LinkedAccountEthereumWalletClient `json:"wallet_client,required"`
	ChainID          string                            `json:"chain_id"`
	ConnectorType    string                            `json:"connector_type"`
	WalletClientType string                            `json:"wallet_client_type"`
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
func (r LinkedAccountEthereum) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountEthereum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountEthereumChainType string

const (
	LinkedAccountEthereumChainTypeEthereum LinkedAccountEthereumChainType = "ethereum"
)

type LinkedAccountEthereumType string

const (
	LinkedAccountEthereumTypeWallet LinkedAccountEthereumType = "wallet"
)

type LinkedAccountEthereumWalletClient string

const (
	LinkedAccountEthereumWalletClientUnknown LinkedAccountEthereumWalletClient = "unknown"
)

// The provider for a smart wallet.
type SmartWalletType string

const (
	SmartWalletTypeSafe                SmartWalletType = "safe"
	SmartWalletTypeKernel              SmartWalletType = "kernel"
	SmartWalletTypeLightAccount        SmartWalletType = "light_account"
	SmartWalletTypeBiconomy            SmartWalletType = "biconomy"
	SmartWalletTypeCoinbaseSmartWallet SmartWalletType = "coinbase_smart_wallet"
	SmartWalletTypeThirdweb            SmartWalletType = "thirdweb"
)

// A smart wallet account linked to the user.
type LinkedAccountSmartWallet struct {
	Address          string  `json:"address,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// The provider for a smart wallet.
	//
	// Any of "safe", "kernel", "light_account", "biconomy", "coinbase_smart_wallet",
	// "thirdweb".
	SmartWalletType SmartWalletType `json:"smart_wallet_type,required"`
	// Any of "smart_wallet".
	Type               LinkedAccountSmartWalletType `json:"type,required"`
	VerifiedAt         float64                      `json:"verified_at,required"`
	SmartWalletVersion string                       `json:"smart_wallet_version"`
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
func (r LinkedAccountSmartWallet) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountSmartWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountSmartWalletType string

const (
	LinkedAccountSmartWalletTypeSmartWallet LinkedAccountSmartWalletType = "smart_wallet"
)

// A Solana wallet account linked to the user.
type LinkedAccountSolana struct {
	Address string `json:"address,required"`
	// Any of "solana".
	ChainType        LinkedAccountSolanaChainType `json:"chain_type,required"`
	FirstVerifiedAt  float64                      `json:"first_verified_at,required"`
	LatestVerifiedAt float64                      `json:"latest_verified_at,required"`
	// Any of "wallet".
	Type       LinkedAccountSolanaType `json:"type,required"`
	VerifiedAt float64                 `json:"verified_at,required"`
	// Any of "unknown".
	WalletClient     LinkedAccountSolanaWalletClient `json:"wallet_client,required"`
	ConnectorType    string                          `json:"connector_type"`
	WalletClientType string                          `json:"wallet_client_type"`
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
func (r LinkedAccountSolana) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountSolana) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountSolanaChainType string

const (
	LinkedAccountSolanaChainTypeSolana LinkedAccountSolanaChainType = "solana"
)

type LinkedAccountSolanaType string

const (
	LinkedAccountSolanaTypeWallet LinkedAccountSolanaType = "wallet"
)

type LinkedAccountSolanaWalletClient string

const (
	LinkedAccountSolanaWalletClientUnknown LinkedAccountSolanaWalletClient = "unknown"
)

// A Farcaster account linked to the user.
type LinkedAccountFarcaster struct {
	Fid              float64 `json:"fid,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	OwnerAddress     string  `json:"owner_address,required"`
	// Any of "farcaster".
	Type              LinkedAccountFarcasterType `json:"type,required"`
	VerifiedAt        float64                    `json:"verified_at,required"`
	Bio               string                     `json:"bio"`
	DisplayName       string                     `json:"display_name"`
	HomepageURL       string                     `json:"homepage_url"`
	ProfilePicture    string                     `json:"profile_picture"`
	ProfilePictureURL string                     `json:"profile_picture_url"`
	SignerPublicKey   string                     `json:"signer_public_key"`
	Username          string                     `json:"username"`
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
func (r LinkedAccountFarcaster) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountFarcaster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountFarcasterType string

const (
	LinkedAccountFarcasterTypeFarcaster LinkedAccountFarcasterType = "farcaster"
)

// A passkey account linked to the user.
type LinkedAccountPasskey struct {
	CredentialID     string  `json:"credential_id,required"`
	EnrolledInMfa    bool    `json:"enrolled_in_mfa,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "passkey".
	Type               LinkedAccountPasskeyType `json:"type,required"`
	VerifiedAt         float64                  `json:"verified_at,required"`
	AuthenticatorName  string                   `json:"authenticator_name"`
	CreatedWithBrowser string                   `json:"created_with_browser"`
	CreatedWithDevice  string                   `json:"created_with_device"`
	CreatedWithOs      string                   `json:"created_with_os"`
	PublicKey          string                   `json:"public_key"`
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
func (r LinkedAccountPasskey) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountPasskey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountPasskeyType string

const (
	LinkedAccountPasskeyTypePasskey LinkedAccountPasskeyType = "passkey"
)

// A Telegram account linked to the user.
type LinkedAccountTelegram struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	TelegramUserID   string  `json:"telegram_user_id,required"`
	// Any of "telegram".
	Type       LinkedAccountTelegramType `json:"type,required"`
	VerifiedAt float64                   `json:"verified_at,required"`
	FirstName  string                    `json:"first_name,nullable"`
	LastName   string                    `json:"last_name,nullable"`
	PhotoURL   string                    `json:"photo_url,nullable"`
	Username   string                    `json:"username,nullable"`
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
func (r LinkedAccountTelegram) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountTelegram) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTelegramType string

const (
	LinkedAccountTelegramTypeTelegram LinkedAccountTelegramType = "telegram"
)

// The method used to recover an embedded wallet account.
type EmbeddedWalletRecoveryMethod string

const (
	EmbeddedWalletRecoveryMethodPrivy                 EmbeddedWalletRecoveryMethod = "privy"
	EmbeddedWalletRecoveryMethodUserPasscode          EmbeddedWalletRecoveryMethod = "user-passcode"
	EmbeddedWalletRecoveryMethodGoogleDrive           EmbeddedWalletRecoveryMethod = "google-drive"
	EmbeddedWalletRecoveryMethodIcloud                EmbeddedWalletRecoveryMethod = "icloud"
	EmbeddedWalletRecoveryMethodRecoveryEncryptionKey EmbeddedWalletRecoveryMethod = "recovery-encryption-key"
	EmbeddedWalletRecoveryMethodPrivyV2               EmbeddedWalletRecoveryMethod = "privy-v2"
)

// An Ethereum embedded wallet account linked to the user.
type LinkedAccountEthereumEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "ethereum".
	ChainType LinkedAccountEthereumEmbeddedWalletChainType `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    LinkedAccountEthereumEmbeddedWalletConnectorType `json:"connector_type,required"`
	Delegated        bool                                             `json:"delegated,required"`
	FirstVerifiedAt  float64                                          `json:"first_verified_at,required"`
	Imported         bool                                             `json:"imported,required"`
	LatestVerifiedAt float64                                          `json:"latest_verified_at,required"`
	// The method used to recover an embedded wallet account.
	//
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod EmbeddedWalletRecoveryMethod `json:"recovery_method,required"`
	// Any of "wallet".
	Type       LinkedAccountEthereumEmbeddedWalletType `json:"type,required"`
	VerifiedAt float64                                 `json:"verified_at,required"`
	// Any of "privy".
	WalletClient LinkedAccountEthereumEmbeddedWalletWalletClient `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType LinkedAccountEthereumEmbeddedWalletWalletClientType `json:"wallet_client_type,required"`
	WalletIndex      float64                                             `json:"wallet_index,required"`
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
func (r LinkedAccountEthereumEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountEthereumEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountEthereumEmbeddedWalletChainType string

const (
	LinkedAccountEthereumEmbeddedWalletChainTypeEthereum LinkedAccountEthereumEmbeddedWalletChainType = "ethereum"
)

type LinkedAccountEthereumEmbeddedWalletConnectorType string

const (
	LinkedAccountEthereumEmbeddedWalletConnectorTypeEmbedded LinkedAccountEthereumEmbeddedWalletConnectorType = "embedded"
)

type LinkedAccountEthereumEmbeddedWalletType string

const (
	LinkedAccountEthereumEmbeddedWalletTypeWallet LinkedAccountEthereumEmbeddedWalletType = "wallet"
)

type LinkedAccountEthereumEmbeddedWalletWalletClient string

const (
	LinkedAccountEthereumEmbeddedWalletWalletClientPrivy LinkedAccountEthereumEmbeddedWalletWalletClient = "privy"
)

type LinkedAccountEthereumEmbeddedWalletWalletClientType string

const (
	LinkedAccountEthereumEmbeddedWalletWalletClientTypePrivy LinkedAccountEthereumEmbeddedWalletWalletClientType = "privy"
)

// A Solana embedded wallet account linked to the user.
type LinkedAccountSolanaEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "solana".
	ChainType LinkedAccountSolanaEmbeddedWalletChainType `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    LinkedAccountSolanaEmbeddedWalletConnectorType `json:"connector_type,required"`
	Delegated        bool                                           `json:"delegated,required"`
	FirstVerifiedAt  float64                                        `json:"first_verified_at,required"`
	Imported         bool                                           `json:"imported,required"`
	LatestVerifiedAt float64                                        `json:"latest_verified_at,required"`
	PublicKey        string                                         `json:"public_key,required"`
	// The method used to recover an embedded wallet account.
	//
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod EmbeddedWalletRecoveryMethod `json:"recovery_method,required"`
	// Any of "wallet".
	Type       LinkedAccountSolanaEmbeddedWalletType `json:"type,required"`
	VerifiedAt float64                               `json:"verified_at,required"`
	// Any of "privy".
	WalletClient LinkedAccountSolanaEmbeddedWalletWalletClient `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType LinkedAccountSolanaEmbeddedWalletWalletClientType `json:"wallet_client_type,required"`
	WalletIndex      float64                                           `json:"wallet_index,required"`
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
func (r LinkedAccountSolanaEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountSolanaEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountSolanaEmbeddedWalletChainType string

const (
	LinkedAccountSolanaEmbeddedWalletChainTypeSolana LinkedAccountSolanaEmbeddedWalletChainType = "solana"
)

type LinkedAccountSolanaEmbeddedWalletConnectorType string

const (
	LinkedAccountSolanaEmbeddedWalletConnectorTypeEmbedded LinkedAccountSolanaEmbeddedWalletConnectorType = "embedded"
)

type LinkedAccountSolanaEmbeddedWalletType string

const (
	LinkedAccountSolanaEmbeddedWalletTypeWallet LinkedAccountSolanaEmbeddedWalletType = "wallet"
)

type LinkedAccountSolanaEmbeddedWalletWalletClient string

const (
	LinkedAccountSolanaEmbeddedWalletWalletClientPrivy LinkedAccountSolanaEmbeddedWalletWalletClient = "privy"
)

type LinkedAccountSolanaEmbeddedWalletWalletClientType string

const (
	LinkedAccountSolanaEmbeddedWalletWalletClientTypePrivy LinkedAccountSolanaEmbeddedWalletWalletClientType = "privy"
)

// A Bitcoin SegWit embedded wallet account linked to the user.
type LinkedAccountBitcoinSegwitEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "bitcoin-segwit".
	ChainType LinkedAccountBitcoinSegwitEmbeddedWalletChainType `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    LinkedAccountBitcoinSegwitEmbeddedWalletConnectorType `json:"connector_type,required"`
	Delegated        bool                                                  `json:"delegated,required"`
	FirstVerifiedAt  float64                                               `json:"first_verified_at,required"`
	Imported         bool                                                  `json:"imported,required"`
	LatestVerifiedAt float64                                               `json:"latest_verified_at,required"`
	PublicKey        string                                                `json:"public_key,required"`
	// The method used to recover an embedded wallet account.
	//
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod EmbeddedWalletRecoveryMethod `json:"recovery_method,required"`
	// Any of "wallet".
	Type       LinkedAccountBitcoinSegwitEmbeddedWalletType `json:"type,required"`
	VerifiedAt float64                                      `json:"verified_at,required"`
	// Any of "privy".
	WalletClient LinkedAccountBitcoinSegwitEmbeddedWalletWalletClient `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType LinkedAccountBitcoinSegwitEmbeddedWalletWalletClientType `json:"wallet_client_type,required"`
	WalletIndex      float64                                                  `json:"wallet_index,required"`
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
func (r LinkedAccountBitcoinSegwitEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountBitcoinSegwitEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountBitcoinSegwitEmbeddedWalletChainType string

const (
	LinkedAccountBitcoinSegwitEmbeddedWalletChainTypeBitcoinSegwit LinkedAccountBitcoinSegwitEmbeddedWalletChainType = "bitcoin-segwit"
)

type LinkedAccountBitcoinSegwitEmbeddedWalletConnectorType string

const (
	LinkedAccountBitcoinSegwitEmbeddedWalletConnectorTypeEmbedded LinkedAccountBitcoinSegwitEmbeddedWalletConnectorType = "embedded"
)

type LinkedAccountBitcoinSegwitEmbeddedWalletType string

const (
	LinkedAccountBitcoinSegwitEmbeddedWalletTypeWallet LinkedAccountBitcoinSegwitEmbeddedWalletType = "wallet"
)

type LinkedAccountBitcoinSegwitEmbeddedWalletWalletClient string

const (
	LinkedAccountBitcoinSegwitEmbeddedWalletWalletClientPrivy LinkedAccountBitcoinSegwitEmbeddedWalletWalletClient = "privy"
)

type LinkedAccountBitcoinSegwitEmbeddedWalletWalletClientType string

const (
	LinkedAccountBitcoinSegwitEmbeddedWalletWalletClientTypePrivy LinkedAccountBitcoinSegwitEmbeddedWalletWalletClientType = "privy"
)

// A Bitcoin Taproot embedded wallet account linked to the user.
type LinkedAccountBitcoinTaprootEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// Any of "bitcoin-taproot".
	ChainType LinkedAccountBitcoinTaprootEmbeddedWalletChainType `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    LinkedAccountBitcoinTaprootEmbeddedWalletConnectorType `json:"connector_type,required"`
	Delegated        bool                                                   `json:"delegated,required"`
	FirstVerifiedAt  float64                                                `json:"first_verified_at,required"`
	Imported         bool                                                   `json:"imported,required"`
	LatestVerifiedAt float64                                                `json:"latest_verified_at,required"`
	PublicKey        string                                                 `json:"public_key,required"`
	// The method used to recover an embedded wallet account.
	//
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod EmbeddedWalletRecoveryMethod `json:"recovery_method,required"`
	// Any of "wallet".
	Type       LinkedAccountBitcoinTaprootEmbeddedWalletType `json:"type,required"`
	VerifiedAt float64                                       `json:"verified_at,required"`
	// Any of "privy".
	WalletClient LinkedAccountBitcoinTaprootEmbeddedWalletWalletClient `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType LinkedAccountBitcoinTaprootEmbeddedWalletWalletClientType `json:"wallet_client_type,required"`
	WalletIndex      float64                                                   `json:"wallet_index,required"`
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
func (r LinkedAccountBitcoinTaprootEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountBitcoinTaprootEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountBitcoinTaprootEmbeddedWalletChainType string

const (
	LinkedAccountBitcoinTaprootEmbeddedWalletChainTypeBitcoinTaproot LinkedAccountBitcoinTaprootEmbeddedWalletChainType = "bitcoin-taproot"
)

type LinkedAccountBitcoinTaprootEmbeddedWalletConnectorType string

const (
	LinkedAccountBitcoinTaprootEmbeddedWalletConnectorTypeEmbedded LinkedAccountBitcoinTaprootEmbeddedWalletConnectorType = "embedded"
)

type LinkedAccountBitcoinTaprootEmbeddedWalletType string

const (
	LinkedAccountBitcoinTaprootEmbeddedWalletTypeWallet LinkedAccountBitcoinTaprootEmbeddedWalletType = "wallet"
)

type LinkedAccountBitcoinTaprootEmbeddedWalletWalletClient string

const (
	LinkedAccountBitcoinTaprootEmbeddedWalletWalletClientPrivy LinkedAccountBitcoinTaprootEmbeddedWalletWalletClient = "privy"
)

type LinkedAccountBitcoinTaprootEmbeddedWalletWalletClientType string

const (
	LinkedAccountBitcoinTaprootEmbeddedWalletWalletClientTypePrivy LinkedAccountBitcoinTaprootEmbeddedWalletWalletClientType = "privy"
)

// A curve signing embedded wallet account linked to the user.
type LinkedAccountCurveSigningEmbeddedWallet struct {
	ID      string `json:"id,required"`
	Address string `json:"address,required"`
	ChainID string `json:"chain_id,required"`
	// The wallet chain types that support curve-based signing.
	//
	// Any of "cosmos", "stellar", "sui", "aptos", "movement", "tron",
	// "bitcoin-segwit", "near", "ton", "starknet".
	ChainType CurveSigningChainType `json:"chain_type,required"`
	// Any of "embedded".
	ConnectorType    LinkedAccountCurveSigningEmbeddedWalletConnectorType `json:"connector_type,required"`
	Delegated        bool                                                 `json:"delegated,required"`
	FirstVerifiedAt  float64                                              `json:"first_verified_at,required"`
	Imported         bool                                                 `json:"imported,required"`
	LatestVerifiedAt float64                                              `json:"latest_verified_at,required"`
	PublicKey        string                                               `json:"public_key,required"`
	// The method used to recover an embedded wallet account.
	//
	// Any of "privy", "user-passcode", "google-drive", "icloud",
	// "recovery-encryption-key", "privy-v2".
	RecoveryMethod EmbeddedWalletRecoveryMethod `json:"recovery_method,required"`
	// Any of "wallet".
	Type       LinkedAccountCurveSigningEmbeddedWalletType `json:"type,required"`
	VerifiedAt float64                                     `json:"verified_at,required"`
	// Any of "privy".
	WalletClient LinkedAccountCurveSigningEmbeddedWalletWalletClient `json:"wallet_client,required"`
	// Any of "privy".
	WalletClientType LinkedAccountCurveSigningEmbeddedWalletWalletClientType `json:"wallet_client_type,required"`
	WalletIndex      float64                                                 `json:"wallet_index,required"`
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
func (r LinkedAccountCurveSigningEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountCurveSigningEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountCurveSigningEmbeddedWalletConnectorType string

const (
	LinkedAccountCurveSigningEmbeddedWalletConnectorTypeEmbedded LinkedAccountCurveSigningEmbeddedWalletConnectorType = "embedded"
)

type LinkedAccountCurveSigningEmbeddedWalletType string

const (
	LinkedAccountCurveSigningEmbeddedWalletTypeWallet LinkedAccountCurveSigningEmbeddedWalletType = "wallet"
)

type LinkedAccountCurveSigningEmbeddedWalletWalletClient string

const (
	LinkedAccountCurveSigningEmbeddedWalletWalletClientPrivy LinkedAccountCurveSigningEmbeddedWalletWalletClient = "privy"
)

type LinkedAccountCurveSigningEmbeddedWalletWalletClientType string

const (
	LinkedAccountCurveSigningEmbeddedWalletWalletClientTypePrivy LinkedAccountCurveSigningEmbeddedWalletWalletClientType = "privy"
)

// A Google OAuth account linked to the user.
type LinkedAccountGoogleOAuth struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "google_oauth".
	Type       LinkedAccountGoogleOAuthType `json:"type,required"`
	VerifiedAt float64                      `json:"verified_at,required"`
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
func (r LinkedAccountGoogleOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountGoogleOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountGoogleOAuthType string

const (
	LinkedAccountGoogleOAuthTypeGoogleOAuth LinkedAccountGoogleOAuthType = "google_oauth"
)

// A Twitter OAuth account linked to the user.
type LinkedAccountTwitterOAuth struct {
	FirstVerifiedAt   float64 `json:"first_verified_at,required"`
	LatestVerifiedAt  float64 `json:"latest_verified_at,required"`
	Name              string  `json:"name,required"`
	ProfilePictureURL string  `json:"profile_picture_url,required"`
	Subject           string  `json:"subject,required"`
	// Any of "twitter_oauth".
	Type       LinkedAccountTwitterOAuthType `json:"type,required"`
	Username   string                        `json:"username,required"`
	VerifiedAt float64                       `json:"verified_at,required"`
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
func (r LinkedAccountTwitterOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountTwitterOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTwitterOAuthType string

const (
	LinkedAccountTwitterOAuthTypeTwitterOAuth LinkedAccountTwitterOAuthType = "twitter_oauth"
)

// A Discord OAuth account linked to the user.
type LinkedAccountDiscordOAuth struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "discord_oauth".
	Type       LinkedAccountDiscordOAuthType `json:"type,required"`
	Username   string                        `json:"username,required"`
	VerifiedAt float64                       `json:"verified_at,required"`
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
func (r LinkedAccountDiscordOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountDiscordOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountDiscordOAuthType string

const (
	LinkedAccountDiscordOAuthTypeDiscordOAuth LinkedAccountDiscordOAuthType = "discord_oauth"
)

// A GitHub OAuth account linked to the user.
type LinkedAccountGitHubOAuth struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "github_oauth".
	Type       LinkedAccountGitHubOAuthType `json:"type,required"`
	Username   string                       `json:"username,required"`
	VerifiedAt float64                      `json:"verified_at,required"`
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
func (r LinkedAccountGitHubOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountGitHubOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountGitHubOAuthType string

const (
	LinkedAccountGitHubOAuthTypeGitHubOAuth LinkedAccountGitHubOAuthType = "github_oauth"
)

// A LinkedIn OAuth account linked to the user.
type LinkedAccountLinkedInOAuth struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "linkedin_oauth".
	Type       LinkedAccountLinkedInOAuthType `json:"type,required"`
	VerifiedAt float64                        `json:"verified_at,required"`
	Name       string                         `json:"name"`
	VanityName string                         `json:"vanity_name"`
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
func (r LinkedAccountLinkedInOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountLinkedInOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountLinkedInOAuthType string

const (
	LinkedAccountLinkedInOAuthTypeLinkedinOAuth LinkedAccountLinkedInOAuthType = "linkedin_oauth"
)

// A Spotify OAuth account linked to the user.
type LinkedAccountSpotifyOAuth struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "spotify_oauth".
	Type       LinkedAccountSpotifyOAuthType `json:"type,required"`
	VerifiedAt float64                       `json:"verified_at,required"`
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
func (r LinkedAccountSpotifyOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountSpotifyOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountSpotifyOAuthType string

const (
	LinkedAccountSpotifyOAuthTypeSpotifyOAuth LinkedAccountSpotifyOAuthType = "spotify_oauth"
)

// An Instagram OAuth account linked to the user.
type LinkedAccountInstagramOAuth struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "instagram_oauth".
	Type       LinkedAccountInstagramOAuthType `json:"type,required"`
	Username   string                          `json:"username,required"`
	VerifiedAt float64                         `json:"verified_at,required"`
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
func (r LinkedAccountInstagramOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountInstagramOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountInstagramOAuthType string

const (
	LinkedAccountInstagramOAuthTypeInstagramOAuth LinkedAccountInstagramOAuthType = "instagram_oauth"
)

// A TikTok OAuth account linked to the user.
type LinkedAccountTiktokOAuth struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Name             string  `json:"name,required"`
	Subject          string  `json:"subject,required"`
	// Any of "tiktok_oauth".
	Type       LinkedAccountTiktokOAuthType `json:"type,required"`
	Username   string                       `json:"username,required"`
	VerifiedAt float64                      `json:"verified_at,required"`
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
func (r LinkedAccountTiktokOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountTiktokOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTiktokOAuthType string

const (
	LinkedAccountTiktokOAuthTypeTiktokOAuth LinkedAccountTiktokOAuthType = "tiktok_oauth"
)

// A LINE OAuth account linked to the user.
type LinkedAccountLineOAuth struct {
	Email             string  `json:"email,required"`
	FirstVerifiedAt   float64 `json:"first_verified_at,required"`
	LatestVerifiedAt  float64 `json:"latest_verified_at,required"`
	Name              string  `json:"name,required"`
	ProfilePictureURL string  `json:"profile_picture_url,required"`
	Subject           string  `json:"subject,required"`
	// Any of "line_oauth".
	Type       LinkedAccountLineOAuthType `json:"type,required"`
	VerifiedAt float64                    `json:"verified_at,required"`
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
func (r LinkedAccountLineOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountLineOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountLineOAuthType string

const (
	LinkedAccountLineOAuthTypeLineOAuth LinkedAccountLineOAuthType = "line_oauth"
)

// A Twitch OAuth account linked to the user.
type LinkedAccountTwitchOAuth struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "twitch_oauth".
	Type       LinkedAccountTwitchOAuthType `json:"type,required"`
	Username   string                       `json:"username,required"`
	VerifiedAt float64                      `json:"verified_at,required"`
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
func (r LinkedAccountTwitchOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountTwitchOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTwitchOAuthType string

const (
	LinkedAccountTwitchOAuthTypeTwitchOAuth LinkedAccountTwitchOAuthType = "twitch_oauth"
)

// An Apple OAuth account linked to the user.
type LinkedAccountAppleOAuth struct {
	Email            string  `json:"email,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// Any of "apple_oauth".
	Type       LinkedAccountAppleOAuthType `json:"type,required"`
	VerifiedAt float64                     `json:"verified_at,required"`
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
func (r LinkedAccountAppleOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountAppleOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountAppleOAuthType string

const (
	LinkedAccountAppleOAuthTypeAppleOAuth LinkedAccountAppleOAuthType = "apple_oauth"
)

// A custom OAuth account linked to the user.
type LinkedAccountCustomOAuth struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	Subject          string  `json:"subject,required"`
	// The ID of a custom OAuth provider, set up for this app. Must start with
	// "custom:".
	Type              CustomOAuthProviderID `json:"type,required"`
	VerifiedAt        float64               `json:"verified_at,required"`
	Email             string                `json:"email"`
	Name              string                `json:"name"`
	ProfilePictureURL string                `json:"profile_picture_url"`
	Username          string                `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstVerifiedAt   respjson.Field
		LatestVerifiedAt  respjson.Field
		Subject           respjson.Field
		Type              respjson.Field
		VerifiedAt        respjson.Field
		Email             respjson.Field
		Name              respjson.Field
		ProfilePictureURL respjson.Field
		Username          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedAccountCustomOAuth) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountCustomOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom JWT account linked to the user.
type LinkedAccountCustomJwt struct {
	CustomUserID     string  `json:"custom_user_id,required"`
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	// Any of "custom_auth".
	Type       LinkedAccountCustomJwtType `json:"type,required"`
	VerifiedAt float64                    `json:"verified_at,required"`
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
func (r LinkedAccountCustomJwt) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountCustomJwt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountCustomJwtType string

const (
	LinkedAccountCustomJwtTypeCustomAuth LinkedAccountCustomJwtType = "custom_auth"
)

// An embedded wallet associated with a cross-app account.
type CrossAppEmbeddedWallet struct {
	Address string `json:"address,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrossAppEmbeddedWallet) RawJSON() string { return r.JSON.raw }
func (r *CrossAppEmbeddedWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A smart wallet associated with a cross-app account.
type CrossAppSmartWallet struct {
	Address string `json:"address,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrossAppSmartWallet) RawJSON() string { return r.JSON.raw }
func (r *CrossAppSmartWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A cross-app account linked to the user.
type LinkedAccountCrossApp struct {
	EmbeddedWallets  []CrossAppEmbeddedWallet `json:"embedded_wallets,required"`
	FirstVerifiedAt  float64                  `json:"first_verified_at,required"`
	LatestVerifiedAt float64                  `json:"latest_verified_at,required"`
	ProviderAppID    string                   `json:"provider_app_id,required"`
	SmartWallets     []CrossAppSmartWallet    `json:"smart_wallets,required"`
	Subject          string                   `json:"subject,required"`
	// Any of "cross_app".
	Type       LinkedAccountCrossAppType `json:"type,required"`
	VerifiedAt float64                   `json:"verified_at,required"`
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
func (r LinkedAccountCrossApp) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountCrossApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountCrossAppType string

const (
	LinkedAccountCrossAppTypeCrossApp LinkedAccountCrossAppType = "cross_app"
)

// An authorization key linked to the user.
type LinkedAccountAuthorizationKey struct {
	FirstVerifiedAt  float64 `json:"first_verified_at,required"`
	LatestVerifiedAt float64 `json:"latest_verified_at,required"`
	PublicKey        string  `json:"public_key,required"`
	// Any of "authorization_key".
	Type       LinkedAccountAuthorizationKeyType `json:"type,required"`
	VerifiedAt float64                           `json:"verified_at,required"`
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
func (r LinkedAccountAuthorizationKey) RawJSON() string { return r.JSON.raw }
func (r *LinkedAccountAuthorizationKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountAuthorizationKeyType string

const (
	LinkedAccountAuthorizationKeyTypeAuthorizationKey LinkedAccountAuthorizationKeyType = "authorization_key"
)

// LinkedAccountUnion contains all possible properties and values from
// [LinkedAccountEmail], [LinkedAccountPhone], [LinkedAccountEthereum],
// [LinkedAccountSolana], [LinkedAccountSmartWallet],
// [LinkedAccountEthereumEmbeddedWallet], [LinkedAccountSolanaEmbeddedWallet],
// [LinkedAccountBitcoinSegwitEmbeddedWallet],
// [LinkedAccountBitcoinTaprootEmbeddedWallet],
// [LinkedAccountCurveSigningEmbeddedWallet], [LinkedAccountGoogleOAuth],
// [LinkedAccountTwitterOAuth], [LinkedAccountDiscordOAuth],
// [LinkedAccountGitHubOAuth], [LinkedAccountSpotifyOAuth],
// [LinkedAccountInstagramOAuth], [LinkedAccountTiktokOAuth],
// [LinkedAccountLineOAuth], [LinkedAccountTwitchOAuth],
// [LinkedAccountLinkedInOAuth], [LinkedAccountAppleOAuth],
// [LinkedAccountCustomOAuth], [LinkedAccountCustomJwt], [LinkedAccountFarcaster],
// [LinkedAccountPasskey], [LinkedAccountTelegram], [LinkedAccountCrossApp],
// [LinkedAccountAuthorizationKey].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LinkedAccountUnion struct {
	Address          string  `json:"address"`
	FirstVerifiedAt  float64 `json:"first_verified_at"`
	LatestVerifiedAt float64 `json:"latest_verified_at"`
	Type             string  `json:"type"`
	VerifiedAt       float64 `json:"verified_at"`
	// This field is from variant [LinkedAccountPhone].
	PhoneNumber string `json:"phoneNumber"`
	// This field is from variant [LinkedAccountPhone].
	Number           string `json:"number"`
	ChainType        string `json:"chain_type"`
	WalletClient     string `json:"wallet_client"`
	ChainID          string `json:"chain_id"`
	ConnectorType    string `json:"connector_type"`
	WalletClientType string `json:"wallet_client_type"`
	// This field is from variant [LinkedAccountSmartWallet].
	SmartWalletType SmartWalletType `json:"smart_wallet_type"`
	// This field is from variant [LinkedAccountSmartWallet].
	SmartWalletVersion string `json:"smart_wallet_version"`
	ID                 string `json:"id"`
	Delegated          bool   `json:"delegated"`
	Imported           bool   `json:"imported"`
	// This field is from variant [LinkedAccountEthereumEmbeddedWallet].
	RecoveryMethod    EmbeddedWalletRecoveryMethod `json:"recovery_method"`
	WalletIndex       float64                      `json:"wallet_index"`
	PublicKey         string                       `json:"public_key"`
	Email             string                       `json:"email"`
	Name              string                       `json:"name"`
	Subject           string                       `json:"subject"`
	ProfilePictureURL string                       `json:"profile_picture_url"`
	Username          string                       `json:"username"`
	// This field is from variant [LinkedAccountLinkedInOAuth].
	VanityName string `json:"vanity_name"`
	// This field is from variant [LinkedAccountCustomJwt].
	CustomUserID string `json:"custom_user_id"`
	// This field is from variant [LinkedAccountFarcaster].
	Fid float64 `json:"fid"`
	// This field is from variant [LinkedAccountFarcaster].
	OwnerAddress string `json:"owner_address"`
	// This field is from variant [LinkedAccountFarcaster].
	Bio string `json:"bio"`
	// This field is from variant [LinkedAccountFarcaster].
	DisplayName string `json:"display_name"`
	// This field is from variant [LinkedAccountFarcaster].
	HomepageURL string `json:"homepage_url"`
	// This field is from variant [LinkedAccountFarcaster].
	ProfilePicture string `json:"profile_picture"`
	// This field is from variant [LinkedAccountFarcaster].
	SignerPublicKey string `json:"signer_public_key"`
	// This field is from variant [LinkedAccountPasskey].
	CredentialID string `json:"credential_id"`
	// This field is from variant [LinkedAccountPasskey].
	EnrolledInMfa bool `json:"enrolled_in_mfa"`
	// This field is from variant [LinkedAccountPasskey].
	AuthenticatorName string `json:"authenticator_name"`
	// This field is from variant [LinkedAccountPasskey].
	CreatedWithBrowser string `json:"created_with_browser"`
	// This field is from variant [LinkedAccountPasskey].
	CreatedWithDevice string `json:"created_with_device"`
	// This field is from variant [LinkedAccountPasskey].
	CreatedWithOs string `json:"created_with_os"`
	// This field is from variant [LinkedAccountTelegram].
	TelegramUserID string `json:"telegram_user_id"`
	// This field is from variant [LinkedAccountTelegram].
	FirstName string `json:"first_name"`
	// This field is from variant [LinkedAccountTelegram].
	LastName string `json:"last_name"`
	// This field is from variant [LinkedAccountTelegram].
	PhotoURL string `json:"photo_url"`
	// This field is from variant [LinkedAccountCrossApp].
	EmbeddedWallets []CrossAppEmbeddedWallet `json:"embedded_wallets"`
	// This field is from variant [LinkedAccountCrossApp].
	ProviderAppID string `json:"provider_app_id"`
	// This field is from variant [LinkedAccountCrossApp].
	SmartWallets []CrossAppSmartWallet `json:"smart_wallets"`
	JSON         struct {
		Address            respjson.Field
		FirstVerifiedAt    respjson.Field
		LatestVerifiedAt   respjson.Field
		Type               respjson.Field
		VerifiedAt         respjson.Field
		PhoneNumber        respjson.Field
		Number             respjson.Field
		ChainType          respjson.Field
		WalletClient       respjson.Field
		ChainID            respjson.Field
		ConnectorType      respjson.Field
		WalletClientType   respjson.Field
		SmartWalletType    respjson.Field
		SmartWalletVersion respjson.Field
		ID                 respjson.Field
		Delegated          respjson.Field
		Imported           respjson.Field
		RecoveryMethod     respjson.Field
		WalletIndex        respjson.Field
		PublicKey          respjson.Field
		Email              respjson.Field
		Name               respjson.Field
		Subject            respjson.Field
		ProfilePictureURL  respjson.Field
		Username           respjson.Field
		VanityName         respjson.Field
		CustomUserID       respjson.Field
		Fid                respjson.Field
		OwnerAddress       respjson.Field
		Bio                respjson.Field
		DisplayName        respjson.Field
		HomepageURL        respjson.Field
		ProfilePicture     respjson.Field
		SignerPublicKey    respjson.Field
		CredentialID       respjson.Field
		EnrolledInMfa      respjson.Field
		AuthenticatorName  respjson.Field
		CreatedWithBrowser respjson.Field
		CreatedWithDevice  respjson.Field
		CreatedWithOs      respjson.Field
		TelegramUserID     respjson.Field
		FirstName          respjson.Field
		LastName           respjson.Field
		PhotoURL           respjson.Field
		EmbeddedWallets    respjson.Field
		ProviderAppID      respjson.Field
		SmartWallets       respjson.Field
		raw                string
	} `json:"-"`
}

func (u LinkedAccountUnion) AsLinkedAccountEmail() (v LinkedAccountEmail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountPhone() (v LinkedAccountPhone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountEthereum() (v LinkedAccountEthereum) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountSolana() (v LinkedAccountSolana) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountSmartWallet() (v LinkedAccountSmartWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountEthereumEmbeddedWallet() (v LinkedAccountEthereumEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountSolanaEmbeddedWallet() (v LinkedAccountSolanaEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountBitcoinSegwitEmbeddedWallet() (v LinkedAccountBitcoinSegwitEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountBitcoinTaprootEmbeddedWallet() (v LinkedAccountBitcoinTaprootEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountCurveSigningEmbeddedWallet() (v LinkedAccountCurveSigningEmbeddedWallet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountGoogleOAuth() (v LinkedAccountGoogleOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountTwitterOAuth() (v LinkedAccountTwitterOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountDiscordOAuth() (v LinkedAccountDiscordOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountGitHubOAuth() (v LinkedAccountGitHubOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountSpotifyOAuth() (v LinkedAccountSpotifyOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountInstagramOAuth() (v LinkedAccountInstagramOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountTiktokOAuth() (v LinkedAccountTiktokOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountLineOAuth() (v LinkedAccountLineOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountTwitchOAuth() (v LinkedAccountTwitchOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountLinkedInOAuth() (v LinkedAccountLinkedInOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountAppleOAuth() (v LinkedAccountAppleOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountCustomOAuth() (v LinkedAccountCustomOAuth) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountCustomJwt() (v LinkedAccountCustomJwt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountFarcaster() (v LinkedAccountFarcaster) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountPasskey() (v LinkedAccountPasskey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountTelegram() (v LinkedAccountTelegram) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountCrossApp() (v LinkedAccountCrossApp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedAccountUnion) AsLinkedAccountAuthorizationKey() (v LinkedAccountAuthorizationKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LinkedAccountUnion) RawJSON() string { return u.JSON.raw }

func (r *LinkedAccountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The possible types of linked accounts.
type LinkedAccountType string

const (
	LinkedAccountTypeEmail            LinkedAccountType = "email"
	LinkedAccountTypePhone            LinkedAccountType = "phone"
	LinkedAccountTypeWallet           LinkedAccountType = "wallet"
	LinkedAccountTypeSmartWallet      LinkedAccountType = "smart_wallet"
	LinkedAccountTypeGoogleOAuth      LinkedAccountType = "google_oauth"
	LinkedAccountTypeTwitterOAuth     LinkedAccountType = "twitter_oauth"
	LinkedAccountTypeDiscordOAuth     LinkedAccountType = "discord_oauth"
	LinkedAccountTypeGitHubOAuth      LinkedAccountType = "github_oauth"
	LinkedAccountTypeSpotifyOAuth     LinkedAccountType = "spotify_oauth"
	LinkedAccountTypeInstagramOAuth   LinkedAccountType = "instagram_oauth"
	LinkedAccountTypeTiktokOAuth      LinkedAccountType = "tiktok_oauth"
	LinkedAccountTypeLineOAuth        LinkedAccountType = "line_oauth"
	LinkedAccountTypeTwitchOAuth      LinkedAccountType = "twitch_oauth"
	LinkedAccountTypeLinkedinOAuth    LinkedAccountType = "linkedin_oauth"
	LinkedAccountTypeAppleOAuth       LinkedAccountType = "apple_oauth"
	LinkedAccountTypeCustomAuth       LinkedAccountType = "custom_auth"
	LinkedAccountTypeFarcaster        LinkedAccountType = "farcaster"
	LinkedAccountTypePasskey          LinkedAccountType = "passkey"
	LinkedAccountTypeTelegram         LinkedAccountType = "telegram"
	LinkedAccountTypeCrossApp         LinkedAccountType = "cross_app"
	LinkedAccountTypeAuthorizationKey LinkedAccountType = "authorization_key"
)

type CustomMetadata map[string]CustomMetadataItemUnion

// CustomMetadataItemUnion contains all possible properties and values from
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataItemUnion struct {
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

func (u CustomMetadataItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataItemUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomMetadataParam map[string]CustomMetadataItemUnionParam

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The payload for importing a wallet account.
//
// The properties Address, ChainType, Type are required.
type LinkedAccountWalletInput struct {
	Address string `json:"address,required"`
	// Any of "ethereum", "solana".
	ChainType LinkedAccountWalletInputChainType `json:"chain_type,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "wallet".
	Type constant.Wallet `json:"type,required"`
	paramObj
}

func (r LinkedAccountWalletInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountWalletInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountWalletInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountWalletInputChainType string

const (
	LinkedAccountWalletInputChainTypeEthereum LinkedAccountWalletInputChainType = "ethereum"
	LinkedAccountWalletInputChainTypeSolana   LinkedAccountWalletInputChainType = "solana"
)

// The payload for importing an email account.
//
// The properties Address, Type are required.
type LinkedAccountEmailInput struct {
	Address string `json:"address,required" format:"email"`
	// This field can be elided, and will marshal its zero value as "email".
	Type constant.Email `json:"type,required"`
	paramObj
}

func (r LinkedAccountEmailInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountEmailInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountEmailInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a phone account.
//
// The properties Number, Type are required.
type LinkedAccountPhoneInput struct {
	Number string `json:"number,required"`
	// This field can be elided, and will marshal its zero value as "phone".
	Type constant.Phone `json:"type,required"`
	paramObj
}

func (r LinkedAccountPhoneInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountPhoneInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountPhoneInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Google account.
//
// The properties Email, Name, Subject, Type are required.
type LinkedAccountGoogleInput struct {
	Email   string `json:"email,required" format:"email"`
	Name    string `json:"name,required"`
	Subject string `json:"subject,required"`
	// This field can be elided, and will marshal its zero value as "google_oauth".
	Type constant.GoogleOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountGoogleInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountGoogleInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountGoogleInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Twitter account.
//
// The properties Name, Subject, Type, Username are required.
type LinkedAccountTwitterInput struct {
	Name              string            `json:"name,required"`
	Subject           string            `json:"subject,required"`
	Username          string            `json:"username,required"`
	ProfilePictureURL param.Opt[string] `json:"profile_picture_url,omitzero" format:"uri"`
	// This field can be elided, and will marshal its zero value as "twitter_oauth".
	Type constant.TwitterOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountTwitterInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTwitterInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTwitterInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Discord account.
//
// The properties Subject, Type, Username are required.
type LinkedAccountDiscordInput struct {
	Subject  string            `json:"subject,required"`
	Username string            `json:"username,required"`
	Email    param.Opt[string] `json:"email,omitzero" format:"email"`
	// This field can be elided, and will marshal its zero value as "discord_oauth".
	Type constant.DiscordOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountDiscordInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountDiscordInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountDiscordInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Github account.
//
// The properties Subject, Type, Username are required.
type LinkedAccountGitHubInput struct {
	Subject  string            `json:"subject,required"`
	Username string            `json:"username,required"`
	Email    param.Opt[string] `json:"email,omitzero" format:"email"`
	Name     param.Opt[string] `json:"name,omitzero"`
	// This field can be elided, and will marshal its zero value as "github_oauth".
	Type constant.GitHubOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountGitHubInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountGitHubInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountGitHubInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Spotify account.
//
// The properties Subject, Type are required.
type LinkedAccountSpotifyInput struct {
	Subject string            `json:"subject,required"`
	Email   param.Opt[string] `json:"email,omitzero" format:"email"`
	Name    param.Opt[string] `json:"name,omitzero"`
	// This field can be elided, and will marshal its zero value as "spotify_oauth".
	Type constant.SpotifyOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountSpotifyInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountSpotifyInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountSpotifyInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing an Instagram account.
//
// The properties Subject, Type, Username are required.
type LinkedAccountInstagramInput struct {
	Subject  string `json:"subject,required"`
	Username string `json:"username,required"`
	// This field can be elided, and will marshal its zero value as "instagram_oauth".
	Type constant.InstagramOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountInstagramInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountInstagramInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountInstagramInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Tiktok account.
//
// The properties Name, Subject, Type, Username are required.
type LinkedAccountTiktokInput struct {
	Name     param.Opt[string] `json:"name,omitzero,required"`
	Subject  string            `json:"subject,required"`
	Username string            `json:"username,required"`
	// This field can be elided, and will marshal its zero value as "tiktok_oauth".
	Type constant.TiktokOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountTiktokInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTiktokInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTiktokInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a LINE account.
//
// The properties Subject, Type are required.
type LinkedAccountLineInput struct {
	Subject           string            `json:"subject,required"`
	Email             param.Opt[string] `json:"email,omitzero" format:"email"`
	Name              param.Opt[string] `json:"name,omitzero"`
	ProfilePictureURL param.Opt[string] `json:"profile_picture_url,omitzero" format:"uri"`
	// This field can be elided, and will marshal its zero value as "line_oauth".
	Type constant.LineOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountLineInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountLineInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountLineInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Twitch account.
//
// The properties Subject, Type are required.
type LinkedAccountTwitchInput struct {
	Subject  string            `json:"subject,required"`
	Username param.Opt[string] `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "twitch_oauth".
	Type constant.TwitchOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountTwitchInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTwitchInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTwitchInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing an Apple account.
//
// The properties Subject, Type are required.
type LinkedAccountAppleInput struct {
	Subject string            `json:"subject,required"`
	Email   param.Opt[string] `json:"email,omitzero" format:"email"`
	// This field can be elided, and will marshal its zero value as "apple_oauth".
	Type constant.AppleOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountAppleInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountAppleInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountAppleInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a LinkedIn account.
//
// The properties Subject, Type are required.
type LinkedAccountLinkedInInput struct {
	Subject    string            `json:"subject,required"`
	Email      param.Opt[string] `json:"email,omitzero" format:"email"`
	Name       param.Opt[string] `json:"name,omitzero"`
	VanityName param.Opt[string] `json:"vanityName,omitzero"`
	// This field can be elided, and will marshal its zero value as "linkedin_oauth".
	Type constant.LinkedinOAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountLinkedInInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountLinkedInInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountLinkedInInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Farcaster account.
//
// The properties Fid, OwnerAddress, Type are required.
type LinkedAccountFarcasterInput struct {
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

func (r LinkedAccountFarcasterInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountFarcasterInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountFarcasterInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Telegram account.
//
// The properties TelegramUserID, Type are required.
type LinkedAccountTelegramInput struct {
	TelegramUserID string            `json:"telegram_user_id,required"`
	FirstName      param.Opt[string] `json:"first_name,omitzero"`
	LastName       param.Opt[string] `json:"last_name,omitzero"`
	PhotoURL       param.Opt[string] `json:"photo_url,omitzero"`
	Username       param.Opt[string] `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "telegram".
	Type constant.Telegram `json:"type,required"`
	paramObj
}

func (r LinkedAccountTelegramInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTelegramInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTelegramInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a Custom JWT account.
//
// The properties CustomUserID, Type are required.
type LinkedAccountCustomJwtInput struct {
	CustomUserID string `json:"custom_user_id,required"`
	// This field can be elided, and will marshal its zero value as "custom_auth".
	Type constant.CustomAuth `json:"type,required"`
	paramObj
}

func (r LinkedAccountCustomJwtInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountCustomJwtInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountCustomJwtInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload for importing a passkey account.
//
// The properties CredentialDeviceType, CredentialID, CredentialPublicKey,
// CredentialUsername, Type are required.
type LinkedAccountPasskeyInput struct {
	// Any of "singleDevice", "multiDevice".
	CredentialDeviceType LinkedAccountPasskeyInputCredentialDeviceType `json:"credential_device_type,omitzero,required"`
	CredentialID         string                                        `json:"credential_id,required"`
	CredentialPublicKey  string                                        `json:"credential_public_key,required"`
	CredentialUsername   string                                        `json:"credential_username,required"`
	// This field can be elided, and will marshal its zero value as "passkey".
	Type constant.Passkey `json:"type,required"`
	paramObj
}

func (r LinkedAccountPasskeyInput) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountPasskeyInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountPasskeyInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountPasskeyInputCredentialDeviceType string

const (
	LinkedAccountPasskeyInputCredentialDeviceTypeSingleDevice LinkedAccountPasskeyInputCredentialDeviceType = "singleDevice"
	LinkedAccountPasskeyInputCredentialDeviceTypeMultiDevice  LinkedAccountPasskeyInputCredentialDeviceType = "multiDevice"
)

func LinkedAccountInputOfWallet(address string, chainType LinkedAccountWalletInputChainType) LinkedAccountInputUnion {
	var wallet LinkedAccountWalletInput
	wallet.Address = address
	wallet.ChainType = chainType
	return LinkedAccountInputUnion{OfWallet: &wallet}
}

func LinkedAccountInputOfEmail(address string) LinkedAccountInputUnion {
	var email LinkedAccountEmailInput
	email.Address = address
	return LinkedAccountInputUnion{OfEmail: &email}
}

func LinkedAccountInputOfPhone(number string) LinkedAccountInputUnion {
	var phone LinkedAccountPhoneInput
	phone.Number = number
	return LinkedAccountInputUnion{OfPhone: &phone}
}

func LinkedAccountInputOfGoogleOAuth(email string, name string, subject string) LinkedAccountInputUnion {
	var googleOAuth LinkedAccountGoogleInput
	googleOAuth.Email = email
	googleOAuth.Name = name
	googleOAuth.Subject = subject
	return LinkedAccountInputUnion{OfGoogleOAuth: &googleOAuth}
}

func LinkedAccountInputOfTwitterOAuth(name string, subject string, username string) LinkedAccountInputUnion {
	var twitterOAuth LinkedAccountTwitterInput
	twitterOAuth.Name = name
	twitterOAuth.Subject = subject
	twitterOAuth.Username = username
	return LinkedAccountInputUnion{OfTwitterOAuth: &twitterOAuth}
}

func LinkedAccountInputOfDiscordOAuth(subject string, username string) LinkedAccountInputUnion {
	var discordOAuth LinkedAccountDiscordInput
	discordOAuth.Subject = subject
	discordOAuth.Username = username
	return LinkedAccountInputUnion{OfDiscordOAuth: &discordOAuth}
}

func LinkedAccountInputOfGitHubOAuth(subject string, username string) LinkedAccountInputUnion {
	var githubOAuth LinkedAccountGitHubInput
	githubOAuth.Subject = subject
	githubOAuth.Username = username
	return LinkedAccountInputUnion{OfGitHubOAuth: &githubOAuth}
}

func LinkedAccountInputOfSpotifyOAuth(subject string) LinkedAccountInputUnion {
	var spotifyOAuth LinkedAccountSpotifyInput
	spotifyOAuth.Subject = subject
	return LinkedAccountInputUnion{OfSpotifyOAuth: &spotifyOAuth}
}

func LinkedAccountInputOfInstagramOAuth(subject string, username string) LinkedAccountInputUnion {
	var instagramOAuth LinkedAccountInstagramInput
	instagramOAuth.Subject = subject
	instagramOAuth.Username = username
	return LinkedAccountInputUnion{OfInstagramOAuth: &instagramOAuth}
}

func LinkedAccountInputOfTiktokOAuth(name string, subject string, username string) LinkedAccountInputUnion {
	var tiktokOAuth LinkedAccountTiktokInput
	tiktokOAuth.Name = param.NewOpt(name)
	tiktokOAuth.Subject = subject
	tiktokOAuth.Username = username
	return LinkedAccountInputUnion{OfTiktokOAuth: &tiktokOAuth}
}

func LinkedAccountInputOfLineOAuth(subject string) LinkedAccountInputUnion {
	var lineOAuth LinkedAccountLineInput
	lineOAuth.Subject = subject
	return LinkedAccountInputUnion{OfLineOAuth: &lineOAuth}
}

func LinkedAccountInputOfTwitchOAuth(subject string) LinkedAccountInputUnion {
	var twitchOAuth LinkedAccountTwitchInput
	twitchOAuth.Subject = subject
	return LinkedAccountInputUnion{OfTwitchOAuth: &twitchOAuth}
}

func LinkedAccountInputOfAppleOAuth(subject string) LinkedAccountInputUnion {
	var appleOAuth LinkedAccountAppleInput
	appleOAuth.Subject = subject
	return LinkedAccountInputUnion{OfAppleOAuth: &appleOAuth}
}

func LinkedAccountInputOfLinkedinOAuth(subject string) LinkedAccountInputUnion {
	var linkedinOAuth LinkedAccountLinkedInInput
	linkedinOAuth.Subject = subject
	return LinkedAccountInputUnion{OfLinkedinOAuth: &linkedinOAuth}
}

func LinkedAccountInputOfFarcaster(fid int64, ownerAddress string) LinkedAccountInputUnion {
	var farcaster LinkedAccountFarcasterInput
	farcaster.Fid = fid
	farcaster.OwnerAddress = ownerAddress
	return LinkedAccountInputUnion{OfFarcaster: &farcaster}
}

func LinkedAccountInputOfTelegram(telegramUserID string) LinkedAccountInputUnion {
	var telegram LinkedAccountTelegramInput
	telegram.TelegramUserID = telegramUserID
	return LinkedAccountInputUnion{OfTelegram: &telegram}
}

func LinkedAccountInputOfCustomAuth(customUserID string) LinkedAccountInputUnion {
	var customAuth LinkedAccountCustomJwtInput
	customAuth.CustomUserID = customUserID
	return LinkedAccountInputUnion{OfCustomAuth: &customAuth}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LinkedAccountInputUnion struct {
	OfWallet         *LinkedAccountWalletInput    `json:",omitzero,inline"`
	OfEmail          *LinkedAccountEmailInput     `json:",omitzero,inline"`
	OfPhone          *LinkedAccountPhoneInput     `json:",omitzero,inline"`
	OfGoogleOAuth    *LinkedAccountGoogleInput    `json:",omitzero,inline"`
	OfTwitterOAuth   *LinkedAccountTwitterInput   `json:",omitzero,inline"`
	OfDiscordOAuth   *LinkedAccountDiscordInput   `json:",omitzero,inline"`
	OfGitHubOAuth    *LinkedAccountGitHubInput    `json:",omitzero,inline"`
	OfSpotifyOAuth   *LinkedAccountSpotifyInput   `json:",omitzero,inline"`
	OfInstagramOAuth *LinkedAccountInstagramInput `json:",omitzero,inline"`
	OfTiktokOAuth    *LinkedAccountTiktokInput    `json:",omitzero,inline"`
	OfLineOAuth      *LinkedAccountLineInput      `json:",omitzero,inline"`
	OfTwitchOAuth    *LinkedAccountTwitchInput    `json:",omitzero,inline"`
	OfAppleOAuth     *LinkedAccountAppleInput     `json:",omitzero,inline"`
	OfLinkedinOAuth  *LinkedAccountLinkedInInput  `json:",omitzero,inline"`
	OfFarcaster      *LinkedAccountFarcasterInput `json:",omitzero,inline"`
	OfTelegram       *LinkedAccountTelegramInput  `json:",omitzero,inline"`
	OfCustomAuth     *LinkedAccountCustomJwtInput `json:",omitzero,inline"`
	OfPasskey        *LinkedAccountPasskeyInput   `json:",omitzero,inline"`
	paramUnion
}

func (u LinkedAccountInputUnion) MarshalJSON() ([]byte, error) {
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
		u.OfCustomAuth,
		u.OfPasskey)
}
func (u *LinkedAccountInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[LinkedAccountInputUnion](
		"type",
		apijson.Discriminator[LinkedAccountWalletInput]("wallet"),
		apijson.Discriminator[LinkedAccountEmailInput]("email"),
		apijson.Discriminator[LinkedAccountPhoneInput]("phone"),
		apijson.Discriminator[LinkedAccountGoogleInput]("google_oauth"),
		apijson.Discriminator[LinkedAccountTwitterInput]("twitter_oauth"),
		apijson.Discriminator[LinkedAccountDiscordInput]("discord_oauth"),
		apijson.Discriminator[LinkedAccountGitHubInput]("github_oauth"),
		apijson.Discriminator[LinkedAccountSpotifyInput]("spotify_oauth"),
		apijson.Discriminator[LinkedAccountInstagramInput]("instagram_oauth"),
		apijson.Discriminator[LinkedAccountTiktokInput]("tiktok_oauth"),
		apijson.Discriminator[LinkedAccountLineInput]("line_oauth"),
		apijson.Discriminator[LinkedAccountTwitchInput]("twitch_oauth"),
		apijson.Discriminator[LinkedAccountAppleInput]("apple_oauth"),
		apijson.Discriminator[LinkedAccountLinkedInInput]("linkedin_oauth"),
		apijson.Discriminator[LinkedAccountFarcasterInput]("farcaster"),
		apijson.Discriminator[LinkedAccountTelegramInput]("telegram"),
		apijson.Discriminator[LinkedAccountCustomJwtInput]("custom_auth"),
		apijson.Discriminator[LinkedAccountPasskeyInput]("passkey"),
	)
}

// A SMS MFA method.
type SMSMfaMethod struct {
	// Any of "sms".
	Type       SMSMfaMethodType `json:"type,required"`
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
func (r SMSMfaMethod) RawJSON() string { return r.JSON.raw }
func (r *SMSMfaMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SMSMfaMethodType string

const (
	SMSMfaMethodTypeSMS SMSMfaMethodType = "sms"
)

// A TOTP MFA method.
type TotpMfaMethod struct {
	// Any of "totp".
	Type       TotpMfaMethodType `json:"type,required"`
	VerifiedAt float64           `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TotpMfaMethod) RawJSON() string { return r.JSON.raw }
func (r *TotpMfaMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TotpMfaMethodType string

const (
	TotpMfaMethodTypeTotp TotpMfaMethodType = "totp"
)

// A Passkey MFA method.
type PasskeyMfaMethod struct {
	// Any of "passkey".
	Type       PasskeyMfaMethodType `json:"type,required"`
	VerifiedAt float64              `json:"verified_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PasskeyMfaMethod) RawJSON() string { return r.JSON.raw }
func (r *PasskeyMfaMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PasskeyMfaMethodType string

const (
	PasskeyMfaMethodTypePasskey PasskeyMfaMethodType = "passkey"
)

// LinkedMfaMethodUnion contains all possible properties and values from
// [SMSMfaMethod], [TotpMfaMethod], [PasskeyMfaMethod].
//
// Use the [LinkedMfaMethodUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LinkedMfaMethodUnion struct {
	// Any of "sms", "totp", "passkey".
	Type       string  `json:"type"`
	VerifiedAt float64 `json:"verified_at"`
	JSON       struct {
		Type       respjson.Field
		VerifiedAt respjson.Field
		raw        string
	} `json:"-"`
}

// anyLinkedMfaMethod is implemented by each variant of [LinkedMfaMethodUnion] to
// add type safety for the return type of [LinkedMfaMethodUnion.AsAny]
type anyLinkedMfaMethod interface {
	implLinkedMfaMethodUnion()
}

func (SMSMfaMethod) implLinkedMfaMethodUnion()     {}
func (TotpMfaMethod) implLinkedMfaMethodUnion()    {}
func (PasskeyMfaMethod) implLinkedMfaMethodUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := LinkedMfaMethodUnion.AsAny().(type) {
//	case privyclient.SMSMfaMethod:
//	case privyclient.TotpMfaMethod:
//	case privyclient.PasskeyMfaMethod:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LinkedMfaMethodUnion) AsAny() anyLinkedMfaMethod {
	switch u.Type {
	case "sms":
		return u.AsSMS()
	case "totp":
		return u.AsTotp()
	case "passkey":
		return u.AsPasskey()
	}
	return nil
}

func (u LinkedMfaMethodUnion) AsSMS() (v SMSMfaMethod) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedMfaMethodUnion) AsTotp() (v TotpMfaMethod) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LinkedMfaMethodUnion) AsPasskey() (v PasskeyMfaMethod) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LinkedMfaMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *LinkedMfaMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	LinkedAccounts []LinkedAccountInputUnion `json:"linked_accounts,omitzero,required"`
	// Custom metadata associated with the user.
	CustomMetadata CustomMetadataParam `json:"custom_metadata,omitzero"`
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

// The property ChainType is required.
type UserNewParamsWallet struct {
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType WalletChainType `json:"chain_type,omitzero,required"`
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

// The property SignerID is required.
type UserNewParamsWalletAdditionalSigner struct {
	// The key quorum ID for the signer.
	SignerID string `json:"signer_id,required" format:"cuid2"`
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
	// The wallet chain types.
	//
	// Any of "ethereum", "solana", "cosmos", "stellar", "sui", "aptos", "movement",
	// "tron", "bitcoin-segwit", "near", "ton", "starknet", "spark".
	ChainType         WalletChainType                                      `json:"chain_type,omitzero,required"`
	CreateSmartWallet param.Opt[bool]                                      `json:"create_smart_wallet,omitzero"`
	AdditionalSigners []UserPregenerateWalletsParamsWalletAdditionalSigner `json:"additional_signers,omitzero"`
	PolicyIDs         []string                                             `json:"policy_ids,omitzero" format:"cuid2"`
	paramObj
}

func (r UserPregenerateWalletsParamsWallet) MarshalJSON() (data []byte, err error) {
	type shadow UserPregenerateWalletsParamsWallet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPregenerateWalletsParamsWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SignerID is required.
type UserPregenerateWalletsParamsWalletAdditionalSigner struct {
	SignerID          string   `json:"signer_id,required" format:"cuid2"`
	OverridePolicyIDs []string `json:"override_policy_ids,omitzero" format:"cuid2"`
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
	SearchTerm string `json:"searchTerm,required"`
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
	CustomMetadata CustomMetadataParam `json:"custom_metadata,omitzero,required"`
	paramObj
}

func (r UserSetCustomMetadataParams) MarshalJSON() (data []byte, err error) {
	type shadow UserSetCustomMetadataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserSetCustomMetadataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUnlinkLinkedAccountParams struct {
	Handle string `json:"handle,required"`
	// The possible types of linked accounts.
	Type     LinkedAccountType `json:"type,omitzero,required"`
	Provider param.Opt[string] `json:"provider,omitzero"`
	paramObj
}

func (r UserUnlinkLinkedAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUnlinkLinkedAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUnlinkLinkedAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
