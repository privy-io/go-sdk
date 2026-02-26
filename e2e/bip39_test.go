package e2e_test

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

// secp256k1N is the order of the secp256k1 curve.
var secp256k1N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)

// generateMnemonic generates a random 12-word BIP39 mnemonic from 128 bits of entropy.
func generateMnemonic(t *testing.T) string {
	t.Helper()

	// Generate 128 bits (16 bytes) of entropy
	entropy := make([]byte, 16)
	if _, err := rand.Read(entropy); err != nil {
		t.Fatalf("failed to generate entropy: %v", err)
	}

	// Compute SHA-256 checksum
	checksum := sha256.Sum256(entropy)

	// Combine entropy (128 bits) + first 4 bits of checksum = 132 bits
	// Split into 12 groups of 11 bits, each mapping to a wordlist index
	// We work with a bit buffer built from entropy || checksum[0]
	var bits []byte
	for _, b := range entropy {
		for i := 7; i >= 0; i-- {
			bits = append(bits, (b>>uint(i))&1)
		}
	}
	// Append first 4 bits of the checksum byte
	for i := 7; i >= 4; i-- {
		bits = append(bits, (checksum[0]>>uint(i))&1)
	}

	words := make([]string, 12)
	for i := 0; i < 12; i++ {
		var idx uint32
		for j := 0; j < 11; j++ {
			idx = (idx << 1) | uint32(bits[i*11+j])
		}
		words[i] = bip39EnglishWordlist[idx]
	}

	return strings.Join(words, " ")
}

// mnemonicToSeed converts a BIP39 mnemonic to a 64-byte seed using PBKDF2-SHA512.
func mnemonicToSeed(mnemonic string) []byte {
	return pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"), 2048, 64, sha512.New)
}

// deriveBIP32Child derives a BIP32 child key from a parent key and chain code.
func deriveBIP32Child(key, chainCode []byte, index uint32) ([]byte, []byte) {
	mac := hmac.New(sha512.New, chainCode)
	if index >= 0x80000000 {
		// Hardened child: 0x00 || parent_key || index
		mac.Write([]byte{0x00})
		mac.Write(key)
	} else {
		// Normal child: compressed_pubkey || index
		privKey, _ := btcec.PrivKeyFromBytes(key)
		mac.Write(privKey.PubKey().SerializeCompressed())
	}
	var indexBytes [4]byte
	binary.BigEndian.PutUint32(indexBytes[:], index)
	mac.Write(indexBytes[:])

	data := mac.Sum(nil)

	// child_key = (IL + parent_key) mod N
	il := new(big.Int).SetBytes(data[:32])
	parentKeyInt := new(big.Int).SetBytes(key)
	childKeyInt := new(big.Int).Add(il, parentKeyInt)
	childKeyInt.Mod(childKeyInt, secp256k1N)

	childKey := childKeyInt.Bytes()
	// Left-pad to 32 bytes
	if len(childKey) < 32 {
		padded := make([]byte, 32)
		copy(padded[32-len(childKey):], childKey)
		childKey = padded
	}

	return childKey, data[32:]
}

// deriveSLIP10Child derives a SLIP-0010 ed25519 child key (hardened only).
func deriveSLIP10Child(key, chainCode []byte, index uint32) ([]byte, []byte) {
	mac := hmac.New(sha512.New, chainCode)
	mac.Write([]byte{0x00})
	mac.Write(key)
	var indexBytes [4]byte
	binary.BigEndian.PutUint32(indexBytes[:], index)
	mac.Write(indexBytes[:])
	data := mac.Sum(nil)
	return data[:32], data[32:]
}

// deriveSolAddressFromMnemonic derives a Solana address from a BIP39 mnemonic
// using SLIP-0010 ed25519 derivation at the path m/44'/501'/{index}'/0'.
func deriveSolAddressFromMnemonic(t *testing.T, mnemonic string, index int) string {
	t.Helper()

	seed := mnemonicToSeed(mnemonic)

	// SLIP-0010 master key for ed25519
	mac := hmac.New(sha512.New, []byte("ed25519 seed"))
	mac.Write(seed)
	masterData := mac.Sum(nil)
	key := masterData[:32]
	chainCode := masterData[32:]

	// Derive m/44'/501'/{index}'/0'
	path := []uint32{
		0x80000000 + 44,           // 44' (purpose)
		0x80000000 + 501,          // 501' (Solana coin type)
		0x80000000 + uint32(index), // {index}' (account)
		0x80000000 + 0,            // 0' (change)
	}

	for _, childIndex := range path {
		key, chainCode = deriveSLIP10Child(key, chainCode, childIndex)
	}
	_ = chainCode

	// ed25519 public key from the derived 32-byte seed
	privKey := ed25519.NewKeyFromSeed(key)
	pubKey := privKey.Public().(ed25519.PublicKey)

	return base58.Encode(pubKey)
}

// deriveEthAddressFromMnemonic derives an Ethereum address from a BIP39 mnemonic
// at the standard BIP44 path m/44'/60'/0'/0/{index}.
func deriveEthAddressFromMnemonic(t *testing.T, mnemonic string, index int) string {
	t.Helper()

	seed := mnemonicToSeed(mnemonic)

	// BIP32 master key from seed
	mac := hmac.New(sha512.New, []byte("Bitcoin seed"))
	mac.Write(seed)
	masterData := mac.Sum(nil)
	key := masterData[:32]
	chainCode := masterData[32:]

	// Derive m/44'/60'/0'/0/{index}
	path := []uint32{
		0x80000000 + 44, // 44' (purpose)
		0x80000000 + 60, // 60' (Ethereum coin type)
		0x80000000 + 0,  // 0' (account)
		0,               // 0 (external chain)
		uint32(index),   // address index
	}

	for _, childIndex := range path {
		key, chainCode = deriveBIP32Child(key, chainCode, childIndex)
	}
	_ = chainCode

	// Derive Ethereum address from private key
	privKey, _ := btcec.PrivKeyFromBytes(key)
	pubKeyBytes := privKey.PubKey().SerializeUncompressed()

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(pubKeyBytes[1:]) // skip 0x04 prefix
	hash := hasher.Sum(nil)

	return "0x" + hex.EncodeToString(hash[len(hash)-20:])
}
