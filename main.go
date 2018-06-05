package main

import (
	"math/rand"
	"time"

	"./core"
	"./role"
	"./util"
)

var chain core.Blockchain
var users []*role.User
var miner *role.Miner

func boostNetwork() {
	// 1. create the initial user of blockchain
	firstUser := role.CreateBoostUser()

	// 2. boost the blockchain with initial user
	diff := core.CreateMADifficulty(10000, 0.05, 16)
	chain = core.InitializeBlockchainWithDiff(&firstUser.Address, diff)

	// 3. create an empty block as the first block
	/* miner_temp := role.CreateMiner(chain)
	block := core.CreateNextEmptyBlock(chain.GetLatestBlock(), uint64(time.Now().UnixNano()/1000000), &miner_temp.Address)

	var nuance uint64
	block.FinalizeBlockAt(nuance, uint64(time.Now().UnixNano()/1000000))
	chain.AddBlock(block) */
}

func boostUsers() {
	// create 10 users
	for i := 0; i < 10; i++ {
		user := role.CreateUser(chain)
		chain.PopulateICOTransaction(miner.GetPrivateKey(), user.Address, core.MinerRewardBase/100)
		users = append(users, user)
		//util.GetMainLogger().Infof("User %v created", user.GetShortIdentity())
	}
}

func startTrading() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var from, to int
	for i := 0; true; i++ {
		from = r1.Intn(10)
		if from < 5 {
			to = 5 + r1.Intn(5)
		} else {
			to = r1.Intn(5)
		}

		amount := r1.Intn(100)
		fee := r1.Intn(10)
		users[from].SendTo(users[to], uint64(amount), uint64(fee))

		time.Sleep(5 * time.Second)
	}
}

func initializeOneMinerAndStartMining() {
	miner = role.CreateMiner(chain)

	miner.StartMining()
}

/*
 * This function is to simulate the blochchain workflow
 */
func runSimulator() {
	// 1. boost the blochchain
	util.GetMainLogger().Infof("Start to boost blockchain \n")
	boostNetwork()
	util.GetMainLogger().Infof("Finished boosting blockchain \n")

	// 3. initialize a miner to mine the trasaction and generate block
	util.GetMainLogger().Infof("Start to boost miner \n")
	go initializeOneMinerAndStartMining()

	time.Sleep(10 * time.Second)

	// 2. boost users
	util.GetMainLogger().Infof("Start to boost users \n")
	boostUsers()
	util.GetMainLogger().Infof("Finished boosting users \n")

	// 4. use miner to vest coins to user. Like user buy coins from exchange

	time.Sleep(50 * time.Second)
	// 4. initialize a few users for generating random transactions
	util.GetMainLogger().Infof("Start to boost users \n")
	startTrading()

}

func main() {
	runSimulator()
}
