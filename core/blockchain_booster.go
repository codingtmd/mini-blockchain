package core

import (
	"crypto/rsa"
	"time"

	"../config"
)

//InitializeBlockchainWithDiff creates a blockchain from scratch
func InitializeBlockchainWithDiff(gensisAddress *rsa.PublicKey, diff Difficulty) Blockchain {
	var chain Blockchain
	chain.txMap = make(map[[config.HashSize]byte]*Transaction)
	chain.utxoMap = make(map[UTXO]bool)
	chain.blockMap = make(map[[config.HashSize]byte]*Block)
	chain.difficulty = diff
	chain.AddressMap = make(map[rsa.PublicKey]map[UTXO]bool)
	chain.TransactionPool = make(map[string]*Transaction)

	gensisBlock := CreateFirstBlock(uint64(time.Now().UnixNano()/1000000), gensisAddress)
	chain.performMinerTransactionAndAddBlock(gensisBlock)

	return chain
}

/*
 * CreateICOTransaction vests amount of coins to speific users
 * currently I use miner to vest the coin, but need a better thought
 */
/* func (chain *Blockchain) PopulateICOTransaction(from_address rsa.PublicKey, from_key *rsa.PrivateKey, to rsa.PublicKey, amount uint64) {
	tx, err := chain.TransferCoin(&from_address, &to, config.MinerRewardBase/4, 500)
	if err != nil {
		util.GetBoosterLogger().Errorf("%v\n", err)
		return
	}
	tx.SignTransaction([]*rsa.PrivateKey{from_key})

	util.GetBoosterLogger().Debugf("%s\n", tx.Print())
	chain.AcceptBroadcastedTransaction(tx)
} */
