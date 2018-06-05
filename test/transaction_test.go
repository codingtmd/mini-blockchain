package test

import (
	"crypto/rsa"
	"testing"
)

func TestTransaction(t *testing.T) {
	/* Create 4 users */
	users, tran, err := createTestTransaction()
	if err != nil {
		t.Errorf("Failed to generate keys %s", err)
	}

	/* Sign the transaction */
	tran.SignTransaction([]*rsa.PrivateKey{users[0], users[1]})

	/* Verify the transaction */
	if tran.VerifyTransaction([]*rsa.PublicKey{&users[0].PublicKey, &users[1].PublicKey}) != nil {
		t.Error("Failed to verify transaction")
	}

	/* Forge the transaction and verify */
	tran.Outputs[1].Value = 20000
	if tran.VerifyTransaction([]*rsa.PublicKey{&users[0].PublicKey, &users[1].PublicKey}) == nil {
		t.Error("Verified a forged transaction")
	}

	tran.Outputs[1].Value = 2000
	tran.Outputs[2].Address = users[0].PublicKey
	if tran.VerifyTransaction([]*rsa.PublicKey{&users[0].PublicKey, &users[1].PublicKey}) == nil {
		t.Error("Verified a forged transaction")
	}
}
