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

type AppleOAuth string                     // Always "apple_oauth"
type CustomAuth string                     // Always "custom_auth"
type DiscordOAuth string                   // Always "discord_oauth"
type Email string                          // Always "email"
type EthSendTransaction string             // Always "eth_sendTransaction"
type EthSign7702Authorization string       // Always "eth_sign7702Authorization"
type EthSignTransaction string             // Always "eth_signTransaction"
type EthSignTypedDataV4 string             // Always "eth_signTypedData_v4"
type Ethereum7702Authorization string      // Always "ethereum_7702_authorization"
type EthereumCalldata string               // Always "ethereum_calldata"
type EthereumTransaction string            // Always "ethereum_transaction"
type EthereumTypedDataDomain string        // Always "ethereum_typed_data_domain"
type EthereumTypedDataMessage string       // Always "ethereum_typed_data_message"
type Farcaster string                      // Always "farcaster"
type GitHubOAuth string                    // Always "github_oauth"
type GoogleOAuth string                    // Always "google_oauth"
type HD string                             // Always "hd"
type InstagramOAuth string                 // Always "instagram_oauth"
type LineOAuth string                      // Always "line_oauth"
type LinkedinOAuth string                  // Always "linkedin_oauth"
type Passkey string                        // Always "passkey"
type PersonalSign string                   // Always "personal_sign"
type Phone string                          // Always "phone"
type PrivateKey string                     // Always "private-key"
type Secp256k1Sign string                  // Always "secp256k1_sign"
type SignAndSendTransaction string         // Always "signAndSendTransaction"
type SignMessage string                    // Always "signMessage"
type SignTransaction string                // Always "signTransaction"
type SMS string                            // Always "sms"
type SolanaProgramInstruction string       // Always "solana_program_instruction"
type SolanaSystemProgramInstruction string // Always "solana_system_program_instruction"
type SolanaTokenProgramInstruction string  // Always "solana_token_program_instruction"
type SpotifyOAuth string                   // Always "spotify_oauth"
type System string                         // Always "system"
type Telegram string                       // Always "telegram"
type TiktokOAuth string                    // Always "tiktok_oauth"
type Totp string                           // Always "totp"
type TwitchOAuth string                    // Always "twitch_oauth"
type TwitterOAuth string                   // Always "twitter_oauth"
type Wallet string                         // Always "wallet"

func (c AppleOAuth) Default() AppleOAuth                 { return "apple_oauth" }
func (c CustomAuth) Default() CustomAuth                 { return "custom_auth" }
func (c DiscordOAuth) Default() DiscordOAuth             { return "discord_oauth" }
func (c Email) Default() Email                           { return "email" }
func (c EthSendTransaction) Default() EthSendTransaction { return "eth_sendTransaction" }
func (c EthSign7702Authorization) Default() EthSign7702Authorization {
	return "eth_sign7702Authorization"
}
func (c EthSignTransaction) Default() EthSignTransaction { return "eth_signTransaction" }
func (c EthSignTypedDataV4) Default() EthSignTypedDataV4 { return "eth_signTypedData_v4" }
func (c Ethereum7702Authorization) Default() Ethereum7702Authorization {
	return "ethereum_7702_authorization"
}
func (c EthereumCalldata) Default() EthereumCalldata       { return "ethereum_calldata" }
func (c EthereumTransaction) Default() EthereumTransaction { return "ethereum_transaction" }
func (c EthereumTypedDataDomain) Default() EthereumTypedDataDomain {
	return "ethereum_typed_data_domain"
}
func (c EthereumTypedDataMessage) Default() EthereumTypedDataMessage {
	return "ethereum_typed_data_message"
}
func (c Farcaster) Default() Farcaster                           { return "farcaster" }
func (c GitHubOAuth) Default() GitHubOAuth                       { return "github_oauth" }
func (c GoogleOAuth) Default() GoogleOAuth                       { return "google_oauth" }
func (c HD) Default() HD                                         { return "hd" }
func (c InstagramOAuth) Default() InstagramOAuth                 { return "instagram_oauth" }
func (c LineOAuth) Default() LineOAuth                           { return "line_oauth" }
func (c LinkedinOAuth) Default() LinkedinOAuth                   { return "linkedin_oauth" }
func (c Passkey) Default() Passkey                               { return "passkey" }
func (c PersonalSign) Default() PersonalSign                     { return "personal_sign" }
func (c Phone) Default() Phone                                   { return "phone" }
func (c PrivateKey) Default() PrivateKey                         { return "private-key" }
func (c Secp256k1Sign) Default() Secp256k1Sign                   { return "secp256k1_sign" }
func (c SignAndSendTransaction) Default() SignAndSendTransaction { return "signAndSendTransaction" }
func (c SignMessage) Default() SignMessage                       { return "signMessage" }
func (c SignTransaction) Default() SignTransaction               { return "signTransaction" }
func (c SMS) Default() SMS                                       { return "sms" }
func (c SolanaProgramInstruction) Default() SolanaProgramInstruction {
	return "solana_program_instruction"
}
func (c SolanaSystemProgramInstruction) Default() SolanaSystemProgramInstruction {
	return "solana_system_program_instruction"
}
func (c SolanaTokenProgramInstruction) Default() SolanaTokenProgramInstruction {
	return "solana_token_program_instruction"
}
func (c SpotifyOAuth) Default() SpotifyOAuth { return "spotify_oauth" }
func (c System) Default() System             { return "system" }
func (c Telegram) Default() Telegram         { return "telegram" }
func (c TiktokOAuth) Default() TiktokOAuth   { return "tiktok_oauth" }
func (c Totp) Default() Totp                 { return "totp" }
func (c TwitchOAuth) Default() TwitchOAuth   { return "twitch_oauth" }
func (c TwitterOAuth) Default() TwitterOAuth { return "twitter_oauth" }
func (c Wallet) Default() Wallet             { return "wallet" }

func (c AppleOAuth) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c CustomAuth) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c DiscordOAuth) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Email) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c EthSendTransaction) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c EthSign7702Authorization) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c EthSignTransaction) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c EthSignTypedDataV4) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Ethereum7702Authorization) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c EthereumCalldata) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c EthereumTransaction) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c EthereumTypedDataDomain) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c EthereumTypedDataMessage) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Farcaster) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c GitHubOAuth) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c GoogleOAuth) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c HD) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c InstagramOAuth) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c LineOAuth) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c LinkedinOAuth) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Passkey) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c PersonalSign) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Phone) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c PrivateKey) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Secp256k1Sign) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c SignAndSendTransaction) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c SignMessage) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c SignTransaction) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c SMS) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c SolanaProgramInstruction) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c SolanaSystemProgramInstruction) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c SolanaTokenProgramInstruction) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c SpotifyOAuth) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Telegram) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c TiktokOAuth) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Totp) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c TwitchOAuth) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c TwitterOAuth) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Wallet) MarshalJSON() ([]byte, error)                         { return marshalString(c) }

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
