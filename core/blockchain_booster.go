package core

import (
	"crypto/rsa"
	"crypto/sha256"
	"time"

	"../util"
)

/*
 *  InitializeBlockchainWithDiff creates a blockchain from scratch
 */
func InitializeBlockchainWithDiff(gensisAddress *rsa.PublicKey, diff Difficulty) Blockchain {
	var chain Blockchain
	chain.txMap = make(map[[HashSize]byte]*Transaction)
	chain.utxoMap = make(map[UTXO]bool)
	chain.blockMap = make(map[[HashSize]byte]*Block)
	chain.difficulty = diff
	chain.AddressMap = make(map[rsa.PublicKey]map[UTXO]bool)
	chain.TransactionPool = map[*Transaction]bool{}

	gensisBlock := CreateFirstBlock(uint64(time.Now().UnixNano()/1000000), gensisAddress)
	chain.performMinerTransactionAndAddBlock(gensisBlock)

	return chain
}

/*
 * CreateICOTransaction vests amount of coins to speific users
 * currently I use miner to vest the coin, but need a better thought
 */
func (chain *Blockchain) PopulateICOTransaction(from *rsa.PrivateKey, to rsa.PublicKey, amount uint64) {
	tx := CreateTransaction(1, 1)
	tx.Inputs[0].OutputIndex = 0
	tx.Inputs[0].PrevtxMap = sha256.Sum256(chain.GetLatestBlock().Transactions[0].GetRawDataToHash())
	tx.Outputs[0].Address = to
	tx.Outputs[0].Value = amount
	tx.SignTransaction([]*rsa.PrivateKey{from})

	util.GetBoosterLogger().Debugf("%s\n", tx.Print())
	chain.AcceptBroadcastedTransaction(&tx)
	util.GetBoosterLogger().Infof("Vest user %v %d coins\n", util.GetShortIdentity(to), amount)
}
