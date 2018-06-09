package test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"../config"
	"../core"
)

type NoDifficulty struct {
}

func (d NoDifficulty) ReachDifficulty(hash [config.HashSize]byte) bool {
	return true
}

func (d NoDifficulty) UpdateDifficulty(usedTimeMs uint64) error {
	return nil
}

func (d NoDifficulty) Print() string {
	return ""
}

func createTestUser(t *testing.T) *rsa.PrivateKey {
	user, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Errorf("Failed to create a user %s", err)
		return nil
	}
	return user
}

func createTestTransaction() ([]*rsa.PrivateKey, *core.Transaction, error) {
	/* Create 4 users */
	var users []*rsa.PrivateKey
	for i := 0; i < 4; i++ {
		user, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, user)
	}

	/* Create a transation with 2 Inputs and 3 Outputs */
	tran := core.CreateTransaction(2, 3)
	rand.Read(tran.Inputs[0].PrevtxMap[:])
	rand.Read(tran.Inputs[1].PrevtxMap[:])

	tran.Outputs[0].Value = 1000
	tran.Outputs[0].Address = users[2].PublicKey
	tran.Outputs[1].Value = 2000
	tran.Outputs[1].Address = users[3].PublicKey
	tran.Outputs[2].Value = 3000
	tran.Outputs[2].Address = users[1].PublicKey

	/* Sign the transaction */
	tran.SignTransaction([]*rsa.PrivateKey{users[0], users[1]})
	return users, &tran, nil
}

/*
 * Create a blockchain with gensis block created by an Address
 */
func createTestBlockchain(gensisAddress *rsa.PublicKey) core.Blockchain {
	var diff NoDifficulty
	return core.InitializeBlockchainWithDiff(gensisAddress, diff)
}
