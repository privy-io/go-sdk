package hpke

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"testing"
)

// recipientRawPubKey extracts the uncompressed EC point from a recipient's SPKI public key.
func recipientRawPubKey(t *testing.T, r *HpkeRecipient) []byte {
	t.Helper()
	pubKeyIface, err := x509.ParsePKIXPublicKey(r.PublicKeySpki)
	if err != nil {
		t.Fatalf("failed to parse SPKI public key: %v", err)
	}
	ecdsaPubKey := pubKeyIface.(*ecdsa.PublicKey)
	return elliptic.Marshal(ecdsaPubKey.Curve, ecdsaPubKey.X, ecdsaPubKey.Y)
}

// --- Sender tests ---

func TestNewHpkeSender(t *testing.T) {
	sender := NewHpkeSender()
	if sender == nil {
		t.Fatal("NewHpkeSender returned nil")
	}
}

func TestSenderEncryptInvalidPublicKey(t *testing.T) {
	sender := NewHpkeSender()

	_, _, err := sender.Encrypt([]byte{0x04, 0x01, 0x02}, []byte("test"))
	if err == nil {
		t.Fatal("expected error when encrypting with invalid public key")
	}
}

// --- Recipient tests ---

func TestNewHpkeRecipient(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("NewHpkeRecipient returned error: %v", err)
	}
	if recipient == nil {
		t.Fatal("NewHpkeRecipient returned nil")
	}
	if len(recipient.PublicKeySpki) == 0 {
		t.Error("expected non-empty PublicKeySpki")
	}
}

func TestNewHpkeRecipientPublicKeyIsValidP256Spki(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}

	pubKeyIface, err := x509.ParsePKIXPublicKey(recipient.PublicKeySpki)
	if err != nil {
		t.Fatalf("PublicKeySpki is not valid SPKI: %v", err)
	}
	ecdsaPubKey, ok := pubKeyIface.(*ecdsa.PublicKey)
	if !ok {
		t.Fatalf("expected *ecdsa.PublicKey, got %T", pubKeyIface)
	}
	if ecdsaPubKey.Curve != elliptic.P256() {
		t.Errorf("expected P-256 curve, got %v", ecdsaPubKey.Curve.Params().Name)
	}
}

func TestNewHpkeRecipientGeneratesUniqueKeys(t *testing.T) {
	r1, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient 1: %v", err)
	}
	r2, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient 2: %v", err)
	}
	if bytes.Equal(r1.PublicKeySpki, r2.PublicKeySpki) {
		t.Error("expected two recipients to have different public keys")
	}
}

func TestRecipientDecryptInvalidEncapsulatedKey(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}

	_, err = recipient.Decrypt([]byte{0x00, 0x01, 0x02}, []byte("ciphertext"))
	if err == nil {
		t.Fatal("expected error when decrypting with invalid encapsulated key")
	}
}

func TestRecipientDecryptTamperedEncapsulatedKey(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}
	rawPubKey := recipientRawPubKey(t, recipient)

	sender := NewHpkeSender()
	enc, ct, err := sender.Encrypt(rawPubKey, []byte("secret"))
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	tampered := make([]byte, len(enc))
	copy(tampered, enc)
	tampered[0] ^= 0xff

	_, err = recipient.Decrypt(tampered, ct)
	if err == nil {
		t.Fatal("expected decryption to fail with tampered encapsulated key")
	}
}

func TestRecipientDecryptTamperedCiphertext(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}
	rawPubKey := recipientRawPubKey(t, recipient)

	sender := NewHpkeSender()
	enc, ct, err := sender.Encrypt(rawPubKey, []byte("secret"))
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	tampered := make([]byte, len(ct))
	copy(tampered, ct)
	tampered[0] ^= 0xff

	_, err = recipient.Decrypt(enc, tampered)
	if err == nil {
		t.Fatal("expected decryption to fail with tampered ciphertext")
	}
}

func TestRecipientDecryptWrongKey(t *testing.T) {
	recipient1, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient1: %v", err)
	}
	recipient2, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient2: %v", err)
	}

	rawPubKey := recipientRawPubKey(t, recipient1)

	sender := NewHpkeSender()
	enc, ct, err := sender.Encrypt(rawPubKey, []byte("secret"))
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	// Attempt to decrypt with a different recipient's key
	_, err = recipient2.Decrypt(enc, ct)
	if err == nil {
		t.Fatal("expected decryption to fail with wrong recipient key")
	}
}

// --- Round-trip (sender + recipient) tests ---

func TestEncryptDecryptRoundTrip(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}
	rawPubKey := recipientRawPubKey(t, recipient)

	sender := NewHpkeSender()
	plaintext := []byte("hello, world")

	enc, ct, err := sender.Encrypt(rawPubKey, plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if len(enc) == 0 {
		t.Error("expected non-empty encapsulated key")
	}
	if len(ct) == 0 {
		t.Error("expected non-empty ciphertext")
	}

	decrypted, err := recipient.Decrypt(enc, ct)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}
	if !bytes.Equal(decrypted, plaintext) {
		t.Errorf("round-trip mismatch: got %q, want %q", decrypted, plaintext)
	}
}

func TestEncryptDecryptEmptyPlaintext(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}
	rawPubKey := recipientRawPubKey(t, recipient)

	sender := NewHpkeSender()
	enc, ct, err := sender.Encrypt(rawPubKey, []byte{})
	if err != nil {
		t.Fatalf("Encrypt of empty plaintext failed: %v", err)
	}

	decrypted, err := recipient.Decrypt(enc, ct)
	if err != nil {
		t.Fatalf("Decrypt of empty plaintext failed: %v", err)
	}
	if len(decrypted) != 0 {
		t.Errorf("expected empty plaintext, got %d bytes", len(decrypted))
	}
}

func TestEncryptDecryptLargePlaintext(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}
	rawPubKey := recipientRawPubKey(t, recipient)

	sender := NewHpkeSender()
	plaintext := make([]byte, 64*1024) // 64 KB
	if _, err := rand.Read(plaintext); err != nil {
		t.Fatalf("failed to generate random plaintext: %v", err)
	}

	enc, ct, err := sender.Encrypt(rawPubKey, plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := recipient.Decrypt(enc, ct)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}
	if !bytes.Equal(decrypted, plaintext) {
		t.Error("round-trip mismatch for large plaintext")
	}
}

func TestEncryptProducesDifferentCiphertexts(t *testing.T) {
	recipient, err := NewHpkeRecipient()
	if err != nil {
		t.Fatalf("failed to create recipient: %v", err)
	}
	rawPubKey := recipientRawPubKey(t, recipient)

	sender := NewHpkeSender()
	plaintext := []byte("same plaintext")

	enc1, ct1, err := sender.Encrypt(rawPubKey, plaintext)
	if err != nil {
		t.Fatalf("first Encrypt failed: %v", err)
	}
	enc2, ct2, err := sender.Encrypt(rawPubKey, plaintext)
	if err != nil {
		t.Fatalf("second Encrypt failed: %v", err)
	}

	if bytes.Equal(enc1, enc2) {
		t.Error("expected different encapsulated keys for two encryptions")
	}
	if bytes.Equal(ct1, ct2) {
		t.Error("expected different ciphertexts for two encryptions")
	}

	// Both should still decrypt correctly
	dec1, err := recipient.Decrypt(enc1, ct1)
	if err != nil {
		t.Fatalf("first Decrypt failed: %v", err)
	}
	dec2, err := recipient.Decrypt(enc2, ct2)
	if err != nil {
		t.Fatalf("second Decrypt failed: %v", err)
	}
	if !bytes.Equal(dec1, plaintext) || !bytes.Equal(dec2, plaintext) {
		t.Error("decrypted plaintexts do not match original")
	}
}
