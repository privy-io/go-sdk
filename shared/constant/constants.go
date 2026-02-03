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

type Ethereum7702Authorization string      // Always "ethereum_7702_authorization"
type EthereumCalldata string               // Always "ethereum_calldata"
type EthereumTransaction string            // Always "ethereum_transaction"
type EthereumTypedDataDomain string        // Always "ethereum_typed_data_domain"
type EthereumTypedDataMessage string       // Always "ethereum_typed_data_message"
type HD string                             // Always "hd"
type PrivateKey string                     // Always "private-key"
type SolanaProgramInstruction string       // Always "solana_program_instruction"
type SolanaSystemProgramInstruction string // Always "solana_system_program_instruction"
type SolanaTokenProgramInstruction string  // Always "solana_token_program_instruction"
type System string                         // Always "system"

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
func (c HD) Default() HD                 { return "hd" }
func (c PrivateKey) Default() PrivateKey { return "private-key" }
func (c SolanaProgramInstruction) Default() SolanaProgramInstruction {
	return "solana_program_instruction"
}
func (c SolanaSystemProgramInstruction) Default() SolanaSystemProgramInstruction {
	return "solana_system_program_instruction"
}
func (c SolanaTokenProgramInstruction) Default() SolanaTokenProgramInstruction {
	return "solana_token_program_instruction"
}
func (c System) Default() System { return "system" }

func (c Ethereum7702Authorization) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c EthereumCalldata) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c EthereumTransaction) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c EthereumTypedDataDomain) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c EthereumTypedDataMessage) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c HD) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c PrivateKey) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c SolanaProgramInstruction) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c SolanaSystemProgramInstruction) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c SolanaTokenProgramInstruction) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                         { return marshalString(c) }

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
