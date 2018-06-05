package test

import (
	"testing"

	"../core"
)

func TestBlock(t *testing.T) {
	users, tran, err := createTestTransaction()
	if err != nil {
		t.Error("Fail to create test transaction")
	}

	block := core.CreateFirstBlock(0, &users[0].PublicKey)
	block.AddTransaction(tran)
	block.FinalizeBlockAt(0, 0)

	/* Forge block */
	Value := block.Transactions[0].Outputs[0].Value
	block.Transactions[0].Outputs[0].Value = Value * 100
	if block.VerifyBlockHash() {
		t.Error("Failed to verify block hash")
	}
	block.Transactions[0].Outputs[0].Value = Value

	outputIndex := block.Transactions[1].Inputs[1].OutputIndex
	block.Transactions[1].Inputs[1].OutputIndex = 3
	if block.VerifyBlockHash() {
		t.Error("Failed to verify block hash")
	}
	block.Transactions[1].Inputs[1].OutputIndex = outputIndex
}
