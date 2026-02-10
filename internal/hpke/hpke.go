// Package hpke provides HPKE encryption/decryption functionality.
package hpke

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"

	"github.com/cloudflare/circl/hpke"
	"github.com/cloudflare/circl/kem"
)

// HpkeRecipient handles HPKE decryption operations using P-256/HKDF-SHA256/ChaCha20-Poly1305.
// It is safe for concurrent use as it only holds immutable cryptographic material.
type HpkeRecipient struct {
	// PublicKeySpki is the SPKI-encoded public key bytes (for sharing with senders)
	PublicKeySpki []byte
	// unexported fields
	suite hpke.Suite
	sk    kem.PrivateKey
}

// NewHpkeRecipient generates a new HPKE recipient with a fresh P-256 keypair.
// Returns error if key generation fails.
func NewHpkeRecipient() (*HpkeRecipient, error) {
	// 1. Generate ECDSA P-256 keypair
	ecdsaKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDSA keypair: %w", err)
	}

	// 2. Initialize HPKE suite with specified parameters
	kemID := hpke.KEM_P256_HKDF_SHA256
	kdfID := hpke.KDF_HKDF_SHA256
	aeadID := hpke.AEAD_ChaCha20Poly1305
	suite := hpke.NewSuite(kemID, kdfID, aeadID)

	// 3. Convert ECDSA private key to circl KEM format via ECDH
	ecdhKey, err := ecdsaKey.ECDH()
	if err != nil {
		return nil, fmt.Errorf("failed to convert ECDSA key to ECDH: %w", err)
	}
	sk, err := kemID.Scheme().UnmarshalBinaryPrivateKey(ecdhKey.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to convert to circl KEM format: %w", err)
	}

	// 4. Marshal public key to SPKI format
	spkiBytes, err := x509.MarshalPKIXPublicKey(&ecdsaKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SPKI public key: %w", err)
	}

	return &HpkeRecipient{
		PublicKeySpki: spkiBytes,
		suite:         suite,
		sk:            sk,
	}, nil
}

// Decrypt decrypts HPKE-encrypted data using the recipient's private key.
// Parameters:
//   - encapsulatedKey: The HPKE encapsulated key from the sender
//   - ciphertext: The encrypted data
//
// Returns:
//   - Decrypted plaintext bytes
//   - Error if decryption fails (invalid key, corrupted data, auth failure)
func (r *HpkeRecipient) Decrypt(encapsulatedKey []byte, ciphertext []byte) ([]byte, error) {
	// Create HPKE receiver with nil info (no additional context)
	receiver, err := r.suite.NewReceiver(r.sk, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HPKE receiver: %w", err)
	}

	// Setup with encapsulated key
	opener, err := receiver.Setup(encapsulatedKey)
	if err != nil {
		return nil, fmt.Errorf("failed to setup HPKE receiver: %w", err)
	}

	// Decrypt with nil AAD (no additional authenticated data)
	plaintext, err := opener.Open(ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}
