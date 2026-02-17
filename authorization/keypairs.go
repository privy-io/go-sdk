package authorization

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// P256KeyPair contains a base64-encoded P-256 key pair.
type P256KeyPair struct {
	// PublicKey is the base64-encoded SPKI-formatted public key.
	PublicKey string
	// PrivateKey is the base64-encoded PKCS8-formatted private key.
	PrivateKey string
}

// GenerateP256KeyPair generates a new ECDSA P-256 key pair and returns
// the public key in SPKI format and the private key in PKCS8 format,
// both base64-encoded.
func GenerateP256KeyPair() (P256KeyPair, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return P256KeyPair{}, fmt.Errorf("failed to generate P-256 key: %w", err)
	}

	pkcs8, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return P256KeyPair{}, fmt.Errorf("failed to marshal private key to PKCS8: %w", err)
	}

	spki, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return P256KeyPair{}, fmt.Errorf("failed to marshal public key to SPKI: %w", err)
	}

	return P256KeyPair{
		PublicKey:  base64.StdEncoding.EncodeToString(spki),
		PrivateKey: base64.StdEncoding.EncodeToString(pkcs8),
	}, nil
}
