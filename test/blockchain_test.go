package test

import (
	"crypto/rsa"
	"crypto/sha256"
	"testing"
	"time"

	"../config"
	"../core"
)

func TestGensisBlock(t *testing.T) {
	user := createTestUser(t)
	chain := createTestBlockchain(&user.PublicKey)

	if chain.BalanceOf(&user.PublicKey) != config.MinerRewardBase {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase, chain.BalanceOf(&user.PublicKey))
	}

}

func TestBlockchainSimple(t *testing.T) {
	user := createTestUser(t)
	chain := createTestBlockchain(&user.PublicKey)
	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+1), &user.PublicKey)
	err := chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	if chain.BalanceOf(&user.PublicKey) != config.MinerRewardBase*2 {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase*2, chain.BalanceOf(&user.PublicKey))
	}
}

func TestBlockchainOneTransaction(t *testing.T) {
	user0 := createTestUser(t)
	user1 := createTestUser(t)
	chain := createTestBlockchain(&user0.PublicKey)
	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase, chain.BalanceOf(&user0.PublicKey))
	}

	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+1), &user1.PublicKey)

	/* Create transaction to transfer all coins from user0 to user1 */
	tx := core.CreateTransaction(1, 1)
	tx.Inputs[0].OutputIndex = 0
	tx.Inputs[0].PrevtxMap = sha256.Sum256(chain.GetLatestBlock().Transactions[0].GetRawDataToHashForTest())
	tx.Outputs[0].Address = user1.PublicKey
	tx.Outputs[0].Value = chain.GetLatestBlock().Transactions[0].Outputs[0].Value
	tx.SignTransaction([]*rsa.PrivateKey{user0})

	nextBlock.AddTransaction(&tx)
	err := chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	if chain.BalanceOf(&user1.PublicKey) != config.MinerRewardBase*2 {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase*2, chain.BalanceOf(&user1.PublicKey))
	}

	if chain.BalanceOf(&user0.PublicKey) != 0 {
		t.Errorf("User balance is incorrect: expected %d, actual %d", 0, chain.BalanceOf(&user0.PublicKey))
	}
}

func TestBlockchainWithTransactionUnsigned(t *testing.T) {
	user0 := createTestUser(t)
	user1 := createTestUser(t)
	chain := createTestBlockchain(&user0.PublicKey)
	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000), &user1.PublicKey)

	/* Create transaction to transfer all coins from user0 to user1 */
	tx := core.CreateTransaction(1, 1)
	tx.Inputs[0].OutputIndex = 0
	//tx.Inputs[0].PrevtxMap = sha256.Sum256(chain.GetLatestBlock().Transactions[0].getRawDataToHash())
	tx.Outputs[0].Address = user1.PublicKey
	tx.Outputs[0].Value = chain.GetLatestBlock().Transactions[0].Outputs[0].Value

	nextBlock.AddTransaction(&tx)
	err := chain.AddBlock(nextBlock)
	if err == nil {
		t.Errorf("Failed to verify a invalid block: %s", err)
	}
}

func TestBlockchainWithTransactionInvalidAmount(t *testing.T) {
	user0 := createTestUser(t)
	user1 := createTestUser(t)
	chain := createTestBlockchain(&user0.PublicKey)
	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000), &user1.PublicKey)

	/* Create transaction to transfer all coins from user0 to user1 */
	tx := core.CreateTransaction(1, 1)
	tx.Inputs[0].OutputIndex = 0
	//tx.Inputs[0].PrevtxMap = sha256.Sum256(chain.GetLatestBlock().Transactions[0].getRawDataToHash())
	tx.Outputs[0].Address = user1.PublicKey
	tx.Outputs[0].Value = chain.GetLatestBlock().Transactions[0].Outputs[0].Value + 1
	tx.SignTransaction([]*rsa.PrivateKey{user0})

	nextBlock.AddTransaction(&tx)
	err := chain.AddBlock(nextBlock)
	if err == nil {
		t.Errorf("Failed to verify a invalid block: %s", err)
	}
}

func TestBlockchainTransfer(t *testing.T) {
	user0 := createTestUser(t)
	user1 := createTestUser(t)
	chain := createTestBlockchain(&user0.PublicKey)
	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase, chain.BalanceOf(&user0.PublicKey))
	}

	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+1), &user1.PublicKey)

	/* Create transaction to transfer all coins from user0 to user1 */
	tx, _ := chain.TransferCoin(&user0.PublicKey, &user1.PublicKey, config.MinerRewardBase/2, 0)
	tx.SignTransaction([]*rsa.PrivateKey{user0})

	nextBlock.AddTransaction(tx)
	err := chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	if chain.BalanceOf(&user1.PublicKey) != config.MinerRewardBase*1.5 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*1.5, chain.BalanceOf(&user1.PublicKey))
	}

	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase*0.5 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*0.5, chain.BalanceOf(&user0.PublicKey))
	}

	nextBlock = core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+2), &user1.PublicKey)
	tx, _ = chain.TransferCoin(&user1.PublicKey, &user0.PublicKey, config.MinerRewardBase*1.2, 0)
	tx.SignTransaction([]*rsa.PrivateKey{user1, user1})
	nextBlock.AddTransaction(tx)
	err = chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	if chain.BalanceOf(&user1.PublicKey) != config.MinerRewardBase*1.3 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*1.3, chain.BalanceOf(&user1.PublicKey))
	}

	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase*1.7 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*1.7, chain.BalanceOf(&user0.PublicKey))
	}
}

func TestBlockchainFee(t *testing.T) {
	user0 := createTestUser(t)
	user1 := createTestUser(t)
	user2 := createTestUser(t)
	chain := createTestBlockchain(&user0.PublicKey)
	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase, chain.BalanceOf(&user0.PublicKey))
	}

	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+1), &user1.PublicKey)

	/* Create transaction to transfer all coins from user0 to user1 */
	tx, _ := chain.TransferCoin(&user0.PublicKey, &user2.PublicKey, config.MinerRewardBase/2, 1000)
	tx.SignTransaction([]*rsa.PrivateKey{user0})

	nextBlock.AddTransaction(tx)
	nextBlock.Transactions[0].Outputs[0].Value += 1000
	err := chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	if chain.BalanceOf(&user1.PublicKey) != config.MinerRewardBase+1000 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*0.5, chain.BalanceOf(&user1.PublicKey))
	}

	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase*0.5-1000 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*0.5-1000, chain.BalanceOf(&user0.PublicKey))
	}

	if chain.BalanceOf(&user2.PublicKey) != config.MinerRewardBase*0.5 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*0.5, chain.BalanceOf(&user0.PublicKey))
	}

}

func TestBlockchainMulitipleFee(t *testing.T) {
	user0 := createTestUser(t)
	user1 := createTestUser(t)
	user2 := createTestUser(t)
	user3 := createTestUser(t)
	chain := createTestBlockchain(&user0.PublicKey)
	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase, chain.BalanceOf(&user0.PublicKey))
	}

	nextBlock := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+1), &user1.PublicKey)
	err := chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	nextBlock = core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000+2), &user2.PublicKey)
	tx0, _ := chain.TransferCoin(&user0.PublicKey, &user2.PublicKey, config.MinerRewardBase/2, 1000)
	tx0.SignTransaction([]*rsa.PrivateKey{user0})
	nextBlock.AddTransaction(tx0)
	nextBlock.Transactions[0].Outputs[0].Value += 1000

	tx1, _ := chain.TransferCoin(&user1.PublicKey, &user3.PublicKey, config.MinerRewardBase/4, 500)
	tx1.SignTransaction([]*rsa.PrivateKey{user1})
	nextBlock.AddTransaction(tx1)
	nextBlock.Transactions[0].Outputs[0].Value += 500

	err = chain.AddBlock(nextBlock)
	if err != nil {
		t.Errorf("Failed to add a valid block: %s", err)
	}

	if chain.BalanceOf(&user0.PublicKey) != config.MinerRewardBase/2-1000 {
		t.Errorf("User balance is incorrect: expected %f, actual %d, %x", config.MinerRewardBase*0.5-1000, chain.BalanceOf(&user0.PublicKey), user0.PublicKey)
	}

	if chain.BalanceOf(&user1.PublicKey) != config.MinerRewardBase*3/4-500 {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase/4-500, chain.BalanceOf(&user1.PublicKey))
	}

	if chain.BalanceOf(&user2.PublicKey) != config.MinerRewardBase*1.5+1500 {
		t.Errorf("User balance is incorrect: expected %f, actual %d", config.MinerRewardBase*1.5+1500, chain.BalanceOf(&user2.PublicKey))
	}

	if chain.BalanceOf(&user3.PublicKey) != config.MinerRewardBase/4 {
		t.Errorf("User balance is incorrect: expected %d, actual %d", config.MinerRewardBase/4, chain.BalanceOf(&user3.PublicKey))
	}
}
