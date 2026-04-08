// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type AppleOAuth string     // Always "apple_oauth"
type CustomAuth string     // Always "custom_auth"
type DiscordOAuth string   // Always "discord_oauth"
type Email string          // Always "email"
type Farcaster string      // Always "farcaster"
type GitHubOAuth string    // Always "github_oauth"
type GoogleOAuth string    // Always "google_oauth"
type InstagramOAuth string // Always "instagram_oauth"
type LineOAuth string      // Always "line_oauth"
type LinkedinOAuth string  // Always "linkedin_oauth"
type Passkey string        // Always "passkey"
type Phone string          // Always "phone"
type SpotifyOAuth string   // Always "spotify_oauth"
type Telegram string       // Always "telegram"
type TiktokOAuth string    // Always "tiktok_oauth"
type TwitchOAuth string    // Always "twitch_oauth"
type TwitterOAuth string   // Always "twitter_oauth"
type Wallet string         // Always "wallet"

func (c AppleOAuth) Default() AppleOAuth         { return "apple_oauth" }
func (c CustomAuth) Default() CustomAuth         { return "custom_auth" }
func (c DiscordOAuth) Default() DiscordOAuth     { return "discord_oauth" }
func (c Email) Default() Email                   { return "email" }
func (c Farcaster) Default() Farcaster           { return "farcaster" }
func (c GitHubOAuth) Default() GitHubOAuth       { return "github_oauth" }
func (c GoogleOAuth) Default() GoogleOAuth       { return "google_oauth" }
func (c InstagramOAuth) Default() InstagramOAuth { return "instagram_oauth" }
func (c LineOAuth) Default() LineOAuth           { return "line_oauth" }
func (c LinkedinOAuth) Default() LinkedinOAuth   { return "linkedin_oauth" }
func (c Passkey) Default() Passkey               { return "passkey" }
func (c Phone) Default() Phone                   { return "phone" }
func (c SpotifyOAuth) Default() SpotifyOAuth     { return "spotify_oauth" }
func (c Telegram) Default() Telegram             { return "telegram" }
func (c TiktokOAuth) Default() TiktokOAuth       { return "tiktok_oauth" }
func (c TwitchOAuth) Default() TwitchOAuth       { return "twitch_oauth" }
func (c TwitterOAuth) Default() TwitterOAuth     { return "twitter_oauth" }
func (c Wallet) Default() Wallet                 { return "wallet" }

func (c AppleOAuth) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c CustomAuth) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c DiscordOAuth) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Email) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Farcaster) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c GitHubOAuth) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c GoogleOAuth) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c InstagramOAuth) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c LineOAuth) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c LinkedinOAuth) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Passkey) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Phone) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c SpotifyOAuth) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Telegram) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c TiktokOAuth) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c TwitchOAuth) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c TwitterOAuth) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Wallet) MarshalJSON() ([]byte, error)         { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
