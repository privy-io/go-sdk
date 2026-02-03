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

	"github.com/stainless-sdks/privy-api-client-go/internal/apijson"
	"github.com/stainless-sdks/privy-api-client-go/internal/apiquery"
	"github.com/stainless-sdks/privy-api-client-go/internal/requestconfig"
	"github.com/stainless-sdks/privy-api-client-go/option"
	"github.com/stainless-sdks/privy-api-client-go/packages/pagination"
	"github.com/stainless-sdks/privy-api-client-go/packages/param"
	"github.com/stainless-sdks/privy-api-client-go/packages/respjson"
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

func (u *CustomMetadataItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The payload for importing a wallet account.
//
// The properties Address, ChainType, Type are required.
type LinkedAccountWalletInputParam struct {
	Address string `json:"address,required"`
	// Any of "ethereum", "solana".
	ChainType LinkedAccountWalletInputChainType `json:"chain_type,omitzero,required"`
	// Any of "wallet".
	Type LinkedAccountWalletInputType `json:"type,omitzero,required"`
	paramObj
}

func (r LinkedAccountWalletInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountWalletInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountWalletInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountWalletInputChainType string

const (
	LinkedAccountWalletInputChainTypeEthereum LinkedAccountWalletInputChainType = "ethereum"
	LinkedAccountWalletInputChainTypeSolana   LinkedAccountWalletInputChainType = "solana"
)

type LinkedAccountWalletInputType string

const (
	LinkedAccountWalletInputTypeWallet LinkedAccountWalletInputType = "wallet"
)

// The payload for importing an email account.
//
// The properties Address, Type are required.
type LinkedAccountEmailInputParam struct {
	Address string `json:"address,required" format:"email"`
	// Any of "email".
	Type LinkedAccountEmailInputType `json:"type,omitzero,required"`
	paramObj
}

func (r LinkedAccountEmailInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountEmailInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountEmailInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountEmailInputType string

const (
	LinkedAccountEmailInputTypeEmail LinkedAccountEmailInputType = "email"
)

// The payload for importing a phone account.
//
// The properties Number, Type are required.
type LinkedAccountPhoneInputParam struct {
	Number string `json:"number,required"`
	// Any of "phone".
	Type LinkedAccountPhoneInputType `json:"type,omitzero,required"`
	paramObj
}

func (r LinkedAccountPhoneInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountPhoneInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountPhoneInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountPhoneInputType string

const (
	LinkedAccountPhoneInputTypePhone LinkedAccountPhoneInputType = "phone"
)

// The payload for importing a Google account.
//
// The properties Email, Name, Subject, Type are required.
type LinkedAccountGoogleInputParam struct {
	Email   string `json:"email,required" format:"email"`
	Name    string `json:"name,required"`
	Subject string `json:"subject,required"`
	// Any of "google_oauth".
	Type LinkedAccountGoogleInputType `json:"type,omitzero,required"`
	paramObj
}

func (r LinkedAccountGoogleInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountGoogleInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountGoogleInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountGoogleInputType string

const (
	LinkedAccountGoogleInputTypeGoogleOAuth LinkedAccountGoogleInputType = "google_oauth"
)

// The payload for importing a Twitter account.
//
// The properties Name, Subject, Type, Username are required.
type LinkedAccountTwitterInputParam struct {
	Name    string `json:"name,required"`
	Subject string `json:"subject,required"`
	// Any of "twitter_oauth".
	Type              LinkedAccountTwitterInputType `json:"type,omitzero,required"`
	Username          string                        `json:"username,required"`
	ProfilePictureURL param.Opt[string]             `json:"profile_picture_url,omitzero" format:"uri"`
	paramObj
}

func (r LinkedAccountTwitterInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTwitterInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTwitterInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTwitterInputType string

const (
	LinkedAccountTwitterInputTypeTwitterOAuth LinkedAccountTwitterInputType = "twitter_oauth"
)

// The payload for importing a Discord account.
//
// The properties Subject, Type, Username are required.
type LinkedAccountDiscordInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "discord_oauth".
	Type     LinkedAccountDiscordInputType `json:"type,omitzero,required"`
	Username string                        `json:"username,required"`
	Email    param.Opt[string]             `json:"email,omitzero" format:"email"`
	paramObj
}

func (r LinkedAccountDiscordInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountDiscordInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountDiscordInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountDiscordInputType string

const (
	LinkedAccountDiscordInputTypeDiscordOAuth LinkedAccountDiscordInputType = "discord_oauth"
)

// The payload for importing a Github account.
//
// The properties Subject, Type, Username are required.
type LinkedAccountGitHubInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "github_oauth".
	Type     LinkedAccountGitHubInputType `json:"type,omitzero,required"`
	Username string                       `json:"username,required"`
	Email    param.Opt[string]            `json:"email,omitzero" format:"email"`
	Name     param.Opt[string]            `json:"name,omitzero"`
	paramObj
}

func (r LinkedAccountGitHubInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountGitHubInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountGitHubInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountGitHubInputType string

const (
	LinkedAccountGitHubInputTypeGitHubOAuth LinkedAccountGitHubInputType = "github_oauth"
)

// The payload for importing a Spotify account.
//
// The properties Subject, Type are required.
type LinkedAccountSpotifyInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "spotify_oauth".
	Type  LinkedAccountSpotifyInputType `json:"type,omitzero,required"`
	Email param.Opt[string]             `json:"email,omitzero" format:"email"`
	Name  param.Opt[string]             `json:"name,omitzero"`
	paramObj
}

func (r LinkedAccountSpotifyInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountSpotifyInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountSpotifyInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountSpotifyInputType string

const (
	LinkedAccountSpotifyInputTypeSpotifyOAuth LinkedAccountSpotifyInputType = "spotify_oauth"
)

// The payload for importing an Instagram account.
//
// The properties Subject, Type, Username are required.
type LinkedAccountInstagramInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "instagram_oauth".
	Type     LinkedAccountInstagramInputType `json:"type,omitzero,required"`
	Username string                          `json:"username,required"`
	paramObj
}

func (r LinkedAccountInstagramInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountInstagramInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountInstagramInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountInstagramInputType string

const (
	LinkedAccountInstagramInputTypeInstagramOAuth LinkedAccountInstagramInputType = "instagram_oauth"
)

// The payload for importing a Tiktok account.
//
// The properties Name, Subject, Type, Username are required.
type LinkedAccountTiktokInputParam struct {
	Name    param.Opt[string] `json:"name,omitzero,required"`
	Subject string            `json:"subject,required"`
	// Any of "tiktok_oauth".
	Type     LinkedAccountTiktokInputType `json:"type,omitzero,required"`
	Username string                       `json:"username,required"`
	paramObj
}

func (r LinkedAccountTiktokInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTiktokInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTiktokInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTiktokInputType string

const (
	LinkedAccountTiktokInputTypeTiktokOAuth LinkedAccountTiktokInputType = "tiktok_oauth"
)

// The payload for importing a LINE account.
//
// The properties Subject, Type are required.
type LinkedAccountLineInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "line_oauth".
	Type              LinkedAccountLineInputType `json:"type,omitzero,required"`
	Email             param.Opt[string]          `json:"email,omitzero" format:"email"`
	Name              param.Opt[string]          `json:"name,omitzero"`
	ProfilePictureURL param.Opt[string]          `json:"profile_picture_url,omitzero" format:"uri"`
	paramObj
}

func (r LinkedAccountLineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountLineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountLineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountLineInputType string

const (
	LinkedAccountLineInputTypeLineOAuth LinkedAccountLineInputType = "line_oauth"
)

// The payload for importing a Twitch account.
//
// The properties Subject, Type are required.
type LinkedAccountTwitchInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "twitch_oauth".
	Type     LinkedAccountTwitchInputType `json:"type,omitzero,required"`
	Username param.Opt[string]            `json:"username,omitzero"`
	paramObj
}

func (r LinkedAccountTwitchInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTwitchInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTwitchInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTwitchInputType string

const (
	LinkedAccountTwitchInputTypeTwitchOAuth LinkedAccountTwitchInputType = "twitch_oauth"
)

// The payload for importing an Apple account.
//
// The properties Subject, Type are required.
type LinkedAccountAppleInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "apple_oauth".
	Type  LinkedAccountAppleInputType `json:"type,omitzero,required"`
	Email param.Opt[string]           `json:"email,omitzero" format:"email"`
	paramObj
}

func (r LinkedAccountAppleInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountAppleInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountAppleInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountAppleInputType string

const (
	LinkedAccountAppleInputTypeAppleOAuth LinkedAccountAppleInputType = "apple_oauth"
)

// The payload for importing a LinkedIn account.
//
// The properties Subject, Type are required.
type LinkedAccountLinkedInInputParam struct {
	Subject string `json:"subject,required"`
	// Any of "linkedin_oauth".
	Type       LinkedAccountLinkedInInputType `json:"type,omitzero,required"`
	Email      param.Opt[string]              `json:"email,omitzero" format:"email"`
	Name       param.Opt[string]              `json:"name,omitzero"`
	VanityName param.Opt[string]              `json:"vanityName,omitzero"`
	paramObj
}

func (r LinkedAccountLinkedInInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountLinkedInInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountLinkedInInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountLinkedInInputType string

const (
	LinkedAccountLinkedInInputTypeLinkedinOAuth LinkedAccountLinkedInInputType = "linkedin_oauth"
)

// The payload for importing a Farcaster account.
//
// The properties Fid, OwnerAddress, Type are required.
type LinkedAccountFarcasterInputParam struct {
	Fid          int64  `json:"fid,required"`
	OwnerAddress string `json:"owner_address,required"`
	// Any of "farcaster".
	Type              LinkedAccountFarcasterInputType `json:"type,omitzero,required"`
	Bio               param.Opt[string]               `json:"bio,omitzero"`
	DisplayName       param.Opt[string]               `json:"display_name,omitzero"`
	HomepageURL       param.Opt[string]               `json:"homepage_url,omitzero"`
	ProfilePictureURL param.Opt[string]               `json:"profile_picture_url,omitzero"`
	Username          param.Opt[string]               `json:"username,omitzero"`
	paramObj
}

func (r LinkedAccountFarcasterInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountFarcasterInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountFarcasterInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountFarcasterInputType string

const (
	LinkedAccountFarcasterInputTypeFarcaster LinkedAccountFarcasterInputType = "farcaster"
)

// The payload for importing a Telegram account.
//
// The properties TelegramUserID, Type are required.
type LinkedAccountTelegramInputParam struct {
	TelegramUserID string `json:"telegram_user_id,required"`
	// Any of "telegram".
	Type      LinkedAccountTelegramInputType `json:"type,omitzero,required"`
	FirstName param.Opt[string]              `json:"first_name,omitzero"`
	LastName  param.Opt[string]              `json:"last_name,omitzero"`
	PhotoURL  param.Opt[string]              `json:"photo_url,omitzero"`
	Username  param.Opt[string]              `json:"username,omitzero"`
	paramObj
}

func (r LinkedAccountTelegramInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountTelegramInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountTelegramInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountTelegramInputType string

const (
	LinkedAccountTelegramInputTypeTelegram LinkedAccountTelegramInputType = "telegram"
)

// The payload for importing a Custom JWT account.
//
// The properties CustomUserID, Type are required.
type LinkedAccountCustomJwtInputParam struct {
	CustomUserID string `json:"custom_user_id,required"`
	// Any of "custom_auth".
	Type LinkedAccountCustomJwtInputType `json:"type,omitzero,required"`
	paramObj
}

func (r LinkedAccountCustomJwtInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedAccountCustomJwtInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedAccountCustomJwtInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedAccountCustomJwtInputType string

const (
	LinkedAccountCustomJwtInputTypeCustomAuth LinkedAccountCustomJwtInputType = "custom_auth"
)

func LinkedAccountInputParamOfWallet(address string, chainType LinkedAccountWalletInputChainType, type_ LinkedAccountWalletInputType) LinkedAccountInputUnionParam {
	var wallet LinkedAccountWalletInputParam
	wallet.Address = address
	wallet.ChainType = chainType
	wallet.Type = type_
	return LinkedAccountInputUnionParam{OfWallet: &wallet}
}

func LinkedAccountInputParamOfEmail(address string) LinkedAccountInputUnionParam {
	var email LinkedAccountEmailInputParam
	email.Address = address
	return LinkedAccountInputUnionParam{OfEmail: &email}
}

func LinkedAccountInputParamOfPhone(number string) LinkedAccountInputUnionParam {
	var phone LinkedAccountPhoneInputParam
	phone.Number = number
	return LinkedAccountInputUnionParam{OfPhone: &phone}
}

func LinkedAccountInputParamOfDiscordOAuth(subject string, type_ LinkedAccountDiscordInputType, username string) LinkedAccountInputUnionParam {
	var discordOAuth LinkedAccountDiscordInputParam
	discordOAuth.Subject = subject
	discordOAuth.Type = type_
	discordOAuth.Username = username
	return LinkedAccountInputUnionParam{OfDiscordOAuth: &discordOAuth}
}

func LinkedAccountInputParamOfGitHubOAuth(subject string, type_ LinkedAccountGitHubInputType, username string) LinkedAccountInputUnionParam {
	var githubOAuth LinkedAccountGitHubInputParam
	githubOAuth.Subject = subject
	githubOAuth.Type = type_
	githubOAuth.Username = username
	return LinkedAccountInputUnionParam{OfGitHubOAuth: &githubOAuth}
}

func LinkedAccountInputParamOfSpotifyOAuth(subject string) LinkedAccountInputUnionParam {
	var spotifyOAuth LinkedAccountSpotifyInputParam
	spotifyOAuth.Subject = subject
	return LinkedAccountInputUnionParam{OfSpotifyOAuth: &spotifyOAuth}
}

func LinkedAccountInputParamOfInstagramOAuth(subject string, type_ LinkedAccountInstagramInputType, username string) LinkedAccountInputUnionParam {
	var instagramOAuth LinkedAccountInstagramInputParam
	instagramOAuth.Subject = subject
	instagramOAuth.Type = type_
	instagramOAuth.Username = username
	return LinkedAccountInputUnionParam{OfInstagramOAuth: &instagramOAuth}
}

func LinkedAccountInputParamOfLineOAuth(subject string) LinkedAccountInputUnionParam {
	var lineOAuth LinkedAccountLineInputParam
	lineOAuth.Subject = subject
	return LinkedAccountInputUnionParam{OfLineOAuth: &lineOAuth}
}

func LinkedAccountInputParamOfTwitchOAuth(subject string) LinkedAccountInputUnionParam {
	var twitchOAuth LinkedAccountTwitchInputParam
	twitchOAuth.Subject = subject
	return LinkedAccountInputUnionParam{OfTwitchOAuth: &twitchOAuth}
}

func LinkedAccountInputParamOfAppleOAuth(subject string) LinkedAccountInputUnionParam {
	var appleOAuth LinkedAccountAppleInputParam
	appleOAuth.Subject = subject
	return LinkedAccountInputUnionParam{OfAppleOAuth: &appleOAuth}
}

func LinkedAccountInputParamOfLinkedinOAuth(subject string) LinkedAccountInputUnionParam {
	var linkedinOAuth LinkedAccountLinkedInInputParam
	linkedinOAuth.Subject = subject
	return LinkedAccountInputUnionParam{OfLinkedinOAuth: &linkedinOAuth}
}

func LinkedAccountInputParamOfFarcaster(fid int64, ownerAddress string, type_ LinkedAccountFarcasterInputType) LinkedAccountInputUnionParam {
	var farcaster LinkedAccountFarcasterInputParam
	farcaster.Fid = fid
	farcaster.OwnerAddress = ownerAddress
	farcaster.Type = type_
	return LinkedAccountInputUnionParam{OfFarcaster: &farcaster}
}

func LinkedAccountInputParamOfTelegram(telegramUserID string) LinkedAccountInputUnionParam {
	var telegram LinkedAccountTelegramInputParam
	telegram.TelegramUserID = telegramUserID
	return LinkedAccountInputUnionParam{OfTelegram: &telegram}
}

func LinkedAccountInputParamOfCustomAuth(customUserID string) LinkedAccountInputUnionParam {
	var customAuth LinkedAccountCustomJwtInputParam
	customAuth.CustomUserID = customUserID
	return LinkedAccountInputUnionParam{OfCustomAuth: &customAuth}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LinkedAccountInputUnionParam struct {
	OfWallet         *LinkedAccountWalletInputParam    `json:",omitzero,inline"`
	OfEmail          *LinkedAccountEmailInputParam     `json:",omitzero,inline"`
	OfPhone          *LinkedAccountPhoneInputParam     `json:",omitzero,inline"`
	OfGoogleOAuth    *LinkedAccountGoogleInputParam    `json:",omitzero,inline"`
	OfTwitterOAuth   *LinkedAccountTwitterInputParam   `json:",omitzero,inline"`
	OfDiscordOAuth   *LinkedAccountDiscordInputParam   `json:",omitzero,inline"`
	OfGitHubOAuth    *LinkedAccountGitHubInputParam    `json:",omitzero,inline"`
	OfSpotifyOAuth   *LinkedAccountSpotifyInputParam   `json:",omitzero,inline"`
	OfInstagramOAuth *LinkedAccountInstagramInputParam `json:",omitzero,inline"`
	OfTiktokOAuth    *LinkedAccountTiktokInputParam    `json:",omitzero,inline"`
	OfLineOAuth      *LinkedAccountLineInputParam      `json:",omitzero,inline"`
	OfTwitchOAuth    *LinkedAccountTwitchInputParam    `json:",omitzero,inline"`
	OfAppleOAuth     *LinkedAccountAppleInputParam     `json:",omitzero,inline"`
	OfLinkedinOAuth  *LinkedAccountLinkedInInputParam  `json:",omitzero,inline"`
	OfFarcaster      *LinkedAccountFarcasterInputParam `json:",omitzero,inline"`
	OfTelegram       *LinkedAccountTelegramInputParam  `json:",omitzero,inline"`
	OfCustomAuth     *LinkedAccountCustomJwtInputParam `json:",omitzero,inline"`
	paramUnion
}

func (u LinkedAccountInputUnionParam) MarshalJSON() ([]byte, error) {
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
func (u *LinkedAccountInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LinkedAccountInputUnionParam) asAny() any {
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
func (u LinkedAccountInputUnionParam) GetChainType() *string {
	if vt := u.OfWallet; vt != nil {
		return (*string)(&vt.ChainType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetNumber() *string {
	if vt := u.OfPhone; vt != nil {
		return &vt.Number
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetVanityName() *string {
	if vt := u.OfLinkedinOAuth; vt != nil && vt.VanityName.Valid() {
		return &vt.VanityName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetFid() *int64 {
	if vt := u.OfFarcaster; vt != nil {
		return &vt.Fid
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetOwnerAddress() *string {
	if vt := u.OfFarcaster; vt != nil {
		return &vt.OwnerAddress
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetBio() *string {
	if vt := u.OfFarcaster; vt != nil && vt.Bio.Valid() {
		return &vt.Bio.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetDisplayName() *string {
	if vt := u.OfFarcaster; vt != nil && vt.DisplayName.Valid() {
		return &vt.DisplayName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetHomepageURL() *string {
	if vt := u.OfFarcaster; vt != nil && vt.HomepageURL.Valid() {
		return &vt.HomepageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetTelegramUserID() *string {
	if vt := u.OfTelegram; vt != nil {
		return &vt.TelegramUserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetFirstName() *string {
	if vt := u.OfTelegram; vt != nil && vt.FirstName.Valid() {
		return &vt.FirstName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetLastName() *string {
	if vt := u.OfTelegram; vt != nil && vt.LastName.Valid() {
		return &vt.LastName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetPhotoURL() *string {
	if vt := u.OfTelegram; vt != nil && vt.PhotoURL.Valid() {
		return &vt.PhotoURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetCustomUserID() *string {
	if vt := u.OfCustomAuth; vt != nil {
		return &vt.CustomUserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetAddress() *string {
	if vt := u.OfWallet; vt != nil {
		return (*string)(&vt.Address)
	} else if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Address)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LinkedAccountInputUnionParam) GetType() *string {
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
func (u LinkedAccountInputUnionParam) GetEmail() *string {
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
func (u LinkedAccountInputUnionParam) GetName() *string {
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
func (u LinkedAccountInputUnionParam) GetSubject() *string {
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
func (u LinkedAccountInputUnionParam) GetUsername() *string {
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
func (u LinkedAccountInputUnionParam) GetProfilePictureURL() *string {
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
	apijson.RegisterUnion[LinkedAccountInputUnionParam](
		"type",
		apijson.Discriminator[LinkedAccountWalletInputParam]("wallet"),
		apijson.Discriminator[LinkedAccountEmailInputParam]("email"),
		apijson.Discriminator[LinkedAccountPhoneInputParam]("phone"),
		apijson.Discriminator[LinkedAccountGoogleInputParam]("google_oauth"),
		apijson.Discriminator[LinkedAccountTwitterInputParam]("twitter_oauth"),
		apijson.Discriminator[LinkedAccountDiscordInputParam]("discord_oauth"),
		apijson.Discriminator[LinkedAccountGitHubInputParam]("github_oauth"),
		apijson.Discriminator[LinkedAccountSpotifyInputParam]("spotify_oauth"),
		apijson.Discriminator[LinkedAccountInstagramInputParam]("instagram_oauth"),
		apijson.Discriminator[LinkedAccountTiktokInputParam]("tiktok_oauth"),
		apijson.Discriminator[LinkedAccountLineInputParam]("line_oauth"),
		apijson.Discriminator[LinkedAccountTwitchInputParam]("twitch_oauth"),
		apijson.Discriminator[LinkedAccountAppleInputParam]("apple_oauth"),
		apijson.Discriminator[LinkedAccountLinkedInInputParam]("linkedin_oauth"),
		apijson.Discriminator[LinkedAccountFarcasterInputParam]("farcaster"),
		apijson.Discriminator[LinkedAccountTelegramInputParam]("telegram"),
		apijson.Discriminator[LinkedAccountCustomJwtInputParam]("custom_auth"),
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
	LinkedAccounts []LinkedAccountInputUnionParam `json:"linked_accounts,omitzero,required"`
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
