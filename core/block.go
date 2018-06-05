package core

import (
	"bytes"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"../util"
)

type uint256 struct {
	data [4]uint64
}

type Block struct {
	hash          [HashSize]byte
	prevBlockHash [HashSize]byte
	blockIdx      uint64

	blockValue   uint64 /* Mining Value of the block */
	timeStampMs  uint64 /* Epoch when mined in ms */
	minerAddress rsa.PublicKey
	nuance       uint256 /* Use to mine so that hash Value must reach a specifc difficulty */

	Transactions []Transaction
}

func createBlock(prevBlockHash [HashSize]byte, blockIdx uint64, timeStampMs uint64, minerAddress *rsa.PublicKey, transactions []Transaction) *Block {
	var block Block
	block.prevBlockHash = prevBlockHash
	block.blockIdx = blockIdx
	block.timeStampMs = timeStampMs
	block.minerAddress = *minerAddress

	/* Create a special transaction to reward miner (always as transaction 0) */
	block.Transactions = append(block.Transactions, CreateTransaction(1, 1))
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, blockIdx)
	copy(block.Transactions[0].Inputs[0].PrevtxMap[:], b)

	/* TODO: 100 coins, should be adjusted based on timeStamp */
	block.Transactions[0].Outputs[0].Value = MinerRewardBase
	block.Transactions[0].Outputs[0].Address = *minerAddress

	/* Add real transactions */
	block.AddTransactions(transactions)

	return &block
}

func CreateFirstBlock(timeStampMs uint64, minerAddress *rsa.PublicKey) *Block {
	var prevBlockHash [HashSize]byte /* doesn't matter for the first block*/
	var trans []Transaction
	return createBlock(prevBlockHash, 0, timeStampMs, minerAddress, trans)
}

func CreateNextEmptyBlock(prevBlock *Block, timeStamp uint64, minerAddress *rsa.PublicKey) *Block {
	var trans []Transaction
	return createBlock(prevBlock.hash, prevBlock.blockIdx+1, timeStamp, minerAddress, trans)
}

func CreateNextBlock(prevBlock *Block, timeStamp uint64, minerAddress *rsa.PublicKey, naunce uint64, transactions []Transaction) *Block {
	block := createBlock(prevBlock.hash, prevBlock.blockIdx+1, timeStamp, minerAddress, transactions)

	/* Finalize block */
	block.nuance.data[0] = naunce
	block.hash = sha256.Sum256(block.getRawDataToHash())
	return block
}

func (block *Block) AddTransaction(tran *Transaction) {
	block.Transactions = append(block.Transactions, *tran)
}

func (block *Block) AddTransactions(trans []Transaction) {
	block.Transactions = append(block.Transactions, trans...)
}

func (block *Block) getRawDataToHash() []byte {
	data := block.prevBlockHash[:]
	data = appendUint64(data, block.timeStampMs)
	/*
	 * Don't need to hash blockIdx, blockValue since they
	 * can be derived from prevBlockHash and timeStamp
	 */
	data = appendAddress(data, &block.minerAddress)
	data = appendUint256(data, block.nuance)

	for i := 0; i < len(block.Transactions); i++ {
		data = append(data, block.Transactions[i].GetRawDataToHash()...)
	}
	return data
}

func (block *Block) FinalizeBlockAt(naunce uint64, timeStampMs uint64) {
	block.nuance.data[0] = naunce
	block.timeStampMs = timeStampMs
	block.hash = sha256.Sum256(block.getRawDataToHash())
}

func (block *Block) VerifyBlockHash() bool {
	hash := sha256.Sum256(block.getRawDataToHash())
	return block.hash == hash
}

func (block *Block) Print() string {
	var buffer bytes.Buffer

	for _, tran := range block.Transactions {
		buffer.WriteString(tran.Print())
	}

	return fmt.Sprintf("Block:%p[hash:%x,prevBlockHash:%x,blockIdx:%v,blockValue:%v,timeStampMs:%v,minerAddress:%v,nuance:%v,Transactions:[%s],",
		&block,
		md5.Sum([]byte(string(block.hash[:]))),
		md5.Sum([]byte(string(block.prevBlockHash[:]))),
		block.blockIdx,
		block.blockValue,
		block.timeStampMs,
		util.GetShortIdentity(block.minerAddress),
		block.nuance,
		buffer.String(),
	)
}
