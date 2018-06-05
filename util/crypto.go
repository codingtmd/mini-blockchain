package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Sign(message []byte, priv *rsa.PrivateKey) ([]byte, error) {
	rng := rand.Reader

	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(rng, priv, crypto.SHA256, hashed[:])

	return signature, err
}

func VerifySignature(message []byte, signature []byte, pub *rsa.PublicKey) error {
	hashed := sha256.Sum256(message)

	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

func GetShortIdentity(address rsa.PublicKey) string {
	full_identity := address.N.String()
	return full_identity[len(full_identity)-5:]
}
