package test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"../util"
)

func TestSign(t *testing.T) {
	message := []byte("Hello world!")

	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Errorf("Failed to generate keys %s\n", err)
	}

	signature, err := util.Sign(message, priv)

	if err != nil {
		t.Errorf("Failed to sign message %s\n", err)
	}

	if util.VerifySignature(message, signature, &priv.PublicKey) != nil {
		t.Errorf("Signature expected true, actual false")
	}

	priv1, err := rsa.GenerateKey(rand.Reader, 1024)
	signature1, err := util.Sign(message, priv1)

	if util.VerifySignature(message, signature1, &priv.PublicKey) == nil {
		t.Error("Signature expect false, actual true")
	}
}
