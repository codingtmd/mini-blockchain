package role

import (
	"crypto/rsa"
	"time"

	"github.com/juju/loggo"

	"../core"
	"../util"
)

type Miner struct {
	chain core.Blockchain
	key   *rsa.PrivateKey

	Address         rsa.PublicKey
	TransactionPool []*core.Transaction
}

func CreateMiner(chain core.Blockchain) *Miner {
	user := CreateUser(chain)

	var miner Miner
	miner.Address = user.Address
	miner.key = user.key
	miner.chain = chain

	return &miner
}

/*
 * StartMining starts to propose and confirm block in the chain
 * Here assume we only have one miner, otherwise this function need to handle multi-threading
 */
func (miner *Miner) StartMining() {
	miner.getLogger().Infof("Miner %v starts mining\n", miner.GetShortIdentity())
	for i := 0; true; i++ {
		block := core.CreateNextEmptyBlock(miner.chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000), &miner.Address)
		for _, tran := range miner.chain.TransactionPool {
			block.AddTransaction(tran)
			miner.getLogger().Debugf("Added transaction %s\n", tran.Print())
		}

		var nuance uint64
		startTime := time.Now()

		for true {
			//miner.getLogger().Infof("current transaction pool %v\n", miner.chain.PrintTransactionPool())

			block.FinalizeBlockAt(nuance, uint64(time.Now().UnixNano()/1000000))
			if miner.chain.ReachDifficulty(block) {
				miner.getLogger().Debugf("Current chain:%s\n", miner.chain.Print())
				miner.getLogger().Debugf("Start to confirm block: %s\n", block.Print())
				err := miner.chain.AddBlock(block)
				if err != nil {
					miner.getLogger().Errorf("Failed to add a valid block: %s\n", err)
					continue
				}
				miner.getLogger().Infof("Confimed Block %s\n", util.Hash(block))

				break
			} else {
				//miner.getLogger().Debugf("Not meet difficulty and sleep 1s\n")
				time.Sleep(1000 * time.Millisecond)
			}

			nuance++
		}

		miner.getLogger().Infof("Mined %d th block at %s (used time (ms) %d, nuance %d)\n",
			i+1, time.Now(), (time.Now().UnixNano()-startTime.UnixNano())/1000000, nuance)
		miner.getLogger().Infof("New difficulty: %s \n", miner.chain.GetDifficulty().Print())
	}
}

func (miner *Miner) GetShortIdentity() string {
	return util.GetShortIdentity(miner.Address)
}

func (miner *Miner) GetPrivateKey() *rsa.PrivateKey {
	return miner.key
}

func (miner *Miner) GetBlockChain() *core.Blockchain {
	return &miner.chain
}

func (miner *Miner) getLogger() loggo.Logger {
	return util.GetMinerLogger(miner.GetShortIdentity())
}

func (miner *Miner) SendTo(receipt *User, amount uint64, fee uint64) {
	tran, err := miner.chain.TransferCoin(&miner.Address, &receipt.Address, amount, fee)
	if err != nil {
		miner.getLogger().Errorf("Failed to create transaction: %v\n", err)
		return
	}

	tran.SignTransaction([]*rsa.PrivateKey{miner.GetPrivateKey()})

	miner.getLogger().Debugf("%s\n", tran.Print())
	miner.chain.AcceptBroadcastedTransaction(tran)
	miner.getLogger().Infof("User %v sends %d coins to user %v\n", miner.GetShortIdentity(), amount, receipt.GetShortIdentity())
}
