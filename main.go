package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"./config"
	"./core"
	"./role"
	"./util"
)

var chain core.Blockchain
var users []*role.User
var miner *role.Miner

const user_count = 4

func boostNetwork() {
	// 1. create the initial user of blockchain
	firstUser := role.CreateBoostUser()

	// 2. boost the blockchain with initial user
	diff := core.CreateMADifficulty(10000, 0.2, 16)
	chain = core.InitializeBlockchainWithDiff(&firstUser.Address, diff)
}

func boostUsers() {
	// create 10 users
	for i := 0; i < user_count; i++ {
		user := role.CreateUser(chain)
		users = append(users, user)
		util.GetMainLogger().Infof("User %v created\n", user.GetShortIdentity())
	}
}

func startTrading() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	block := miner.GetBlockChain().GetLatestBlock()
	var usage [user_count + 1]int // miner is the last one
	for i := 0; i <= user_count; i++ {
		usage[user_count] = 0
	}

	var from, to int
	for i := 0; true; i++ {
		from = r1.Intn(user_count)
		if from < user_count/2 {
			to = user_count/2 + r1.Intn(user_count/2)
		} else {
			to = r1.Intn(user_count / 2)
		}

		//if block.GetBlockHash() != chain.GetLatestBlock().GetBlockHash() {
		if util.Hash(*block) != util.Hash(miner.GetBlockChain().GetLatestBlock()) {
			util.GetMainLogger().Infof("Chain confirmed a new block. Clean the usage\n")
			for i := 0; i <= user_count; i++ {
				usage[user_count] = 0
			}
			block = miner.GetBlockChain().GetLatestBlock()
		}

		if block != nil {
			//util.GetMainLogger().Debugf("Verify %s, %s\n", util.Hash(*block), util.Hash(miner.GetBlockChain().GetLatestBlock()))
			//util.GetMainLogger().Debugf("outdated blockchain %s,\n", new_chain.Print())
			time.Sleep(500 * time.Millisecond)
		}

		amount := r1.Intn(config.MinerRewardBase / 1000)
		fee := r1.Intn(10)
		if usage[user_count] == 0 && int(miner.GetBlockChain().BalanceOf(&miner.Address)) > amount {
			//util.GetMainLogger().Debugf("Verify %s, %s\n", util.Hash(block), util.Hash(miner.GetBlockChain().GetLatestBlock()))
			miner.SendTo(users[to], uint64(amount), uint64(fee))
			time.Sleep(1 * time.Second)
			usage[user_count] = 1
		}

		amount = r1.Intn(config.MinerRewardBase / 1000)
		fee = r1.Intn(user_count)
		if usage[from] == 0 && int(miner.GetBlockChain().BalanceOf(&users[from].Address)) > amount {
			//util.GetMainLogger().Debugf("Verify %s, %s\n", util.Hash(block), util.Hash(miner.GetBlockChain().GetLatestBlock()))
			users[from].SendTo(users[to], uint64(amount), uint64(fee))
			time.Sleep(1 * time.Second)
			usage[from] = 1
		}
	}
}

func initializeOneMinerAndStartMining() {
	miner = role.CreateMiner(chain)
	miner.StartMining()
}

func printStatus() {
	for {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("Miner[%s:%d]] ", miner.GetShortIdentity(), miner.GetBlockChain().BalanceOf(&miner.Address)))

		for i := 0; i < user_count; i++ {
			buffer.WriteString(fmt.Sprintf("User[%s:%d]] ", users[i].GetShortIdentity(), miner.GetBlockChain().BalanceOf(&users[i].Address)))
		}

		util.GetMainLogger().Debugf("Account Status: %s\n", buffer.String())
		//util.GetMainLogger().Debugf("Chain Status: %s\n", miner.GetBlockChain().Print())

		time.Sleep(1 * time.Second)
	}
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

	//time.Sleep(10 * time.Second)

	// 2. boost users
	util.GetMainLogger().Infof("Start to boost users \n")
	boostUsers()
	util.GetMainLogger().Infof("Finished boosting users \n")

	// 3. print status
	go printStatus()

	// 4. use miner to vest coins to user. Like user buy coins from exchange

	//time.Sleep(10 * time.Second)
	// 5. initialize a few users for generating random transactions
	util.GetMainLogger().Infof("Start to boost trading \n")
	go startTrading()

	for {
		time.Sleep(10 * time.Second)
	}
}

func main() {
	runSimulator()
}
