package hpke

import (
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/hpke"
)

// HpkeSender handles HPKE encryption operations using P-256/HKDF-SHA256/ChaCha20-Poly1305.
type HpkeSender struct {
	suite hpke.Suite
	kemID hpke.KEM
}

// NewHpkeSender creates a new HPKE sender configured with the same cipher suite
// as the recipient (P-256/HKDF-SHA256/ChaCha20-Poly1305).
func NewHpkeSender() *HpkeSender {
	kemID := hpke.KEM_P256_HKDF_SHA256
	kdfID := hpke.KDF_HKDF_SHA256
	aeadID := hpke.AEAD_ChaCha20Poly1305
	return &HpkeSender{
		suite: hpke.NewSuite(kemID, kdfID, aeadID),
		kemID: kemID,
	}
}

// Encrypt encrypts plaintext using the recipient's public key.
// Returns the HPKE encapsulated key and ciphertext.
func (s *HpkeSender) Encrypt(recipientPubKey []byte, plaintext []byte) (encapsulatedKey []byte, ciphertext []byte, err error) {
	// 1. Convert raw EC point bytes to circl KEM public key
	kemPubKey, err := s.kemID.Scheme().UnmarshalBinaryPublicKey(recipientPubKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal KEM public key: %w", err)
	}

	// 2. Create HPKE sender and setup encryption
	sender, err := s.suite.NewSender(kemPubKey, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create HPKE sender: %w", err)
	}
	enc, sealer, err := sender.Setup(rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup HPKE sender: %w", err)
	}

	// 3. Encrypt plaintext with nil AAD
	ct, err := sealer.Seal(plaintext, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encrypt: %w", err)
	}

	return enc, ct, nil
}
