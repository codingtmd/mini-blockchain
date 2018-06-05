<img src="https://drive.google.com/uc?export=view&id=1B2ar0Mn8CV6GHCVg76Tw0TZ7a8lGZsPk" width="40%">

## Mini-Blockchain

Mini-Blockchain is a reference design for a blockchain system to demostate a full end2end flow in current blockchain technology. 

There are so many open source projects of blockchain implementation, like ethereum, bitcoin, eos etc.. But even running their code will take u couple of days to setup. 

The goal of the project is to build a "Hello-world" version of blockchain to help people understand how the foundamental technology without wasting time reading in a huge code base, to simulate the different roles, like user, miner, booster, and demostrate how they participate the blockchain network. And it also simplify the deployment, instead of configuring tons of library and debugging a version in a real distributed network, it is single machine version to easy the debug, test, and iterate.

Since it is primarily designed as a learning tool for people to learn the blockchain end to end. It is meant to be as simple to understand as possible while still providing clear structure and workflow. 

## What is this not?

Mini-Blockchain is **not** meant for real usage as it is a simplified veraion and has many hard coded parameters. It's meant for 

It also may just simply stop working at any time since there are some depandences.

## What is included?

This reference system is written in Go.(Don't ask me why use Go. I just fucking lost my mind.)
 - a main.go to boost the blockchian and simnulate the workflow
 - ./core to implement the blockchain
 - ./role to implement different actors in the blockchain ecosystem. 
 - ./test to implement some unit tests to keep the code in some quality(even very little)


## Parts of the system

This is how bitcoin mining works. I follow the same flow except checking the 2016th block.

<img src="https://drive.google.com/uc?export=view&id=1gzntnj7ZSDGAZCuxTarCNFHbKEI2TgOs" width="60%">
SOME DEFINITIONS:
**User -** anyone who as an account(hash address) is a user. User holds the coins in their account and is able to send/receive coins to/from other users.

**Transaction -** any coin transfer between two users is a transaction. The activities of users will generate transaction, which is the source of transaction pool.

**Miner -** a special role, who doesn't generate transaction, but collect/validate transaction from the pool.

**Block -** a collection of validated transactions. Every miner can propose a block, but only the one acknowledged by most of the miners will be the official block in the chian.

**Difficulty -**  a measure of how difficult it is to find a hash below a given target. Valid blocks must have a hash below this target. Mining pools also have a pool-specific share difficulty setting a lower limit for shares. In bitcoin, the network difficulty changes every 2016 blocks. For my implementation, the difficulty changes every block.

**Nonce -** a 32-bit (4-byte) field to ramdom the hash generation. Any change to the block data (such as the nonce) will make the block hash completely different. The resulting hash has to be a value less than the current difficulty and so will have to have a certain number of leading zero bits to be less than that. As this iterative calculation requires time and resources, the presentation of the block with the correct nonce value constitutes proof of work.

**Reward -** when a block is discovered, the miner may award themselves a certain number of bitcoins, which is agreed-upon by everyone in the network. Normally the rewarding transaction is the first transaction in a block proposed by the miner.

**Fee -** The miner is also awarded the fees paid by users sending transactions. The fee is an incentive for the miner to include the transaction in their block. In the future, as the number of new bitcoins miners are allowed to create in each block dwindles, the fees will make up a much more important percentage of mining income. Ethereum is a good example of fee usage.

## How does the simulated workflow work?

First, it will start to initialize an empty blockchain. Note, for every blockchain project, boosting it from beginning is the tricky part. In this implementation, I created a empty chian and then add a empty block as the head.

Second, it creates 1 miner to represent Consortium Blockchain(if you don't know what is consortium blockchain, read [this](https://www.blockchaindailynews.com/The-difference-between-a-Private-Public-Consortium-Blockchain_a24681.html "this")). The miner will submit several empty blocks and earn rewards for each block.

Third, it creates 10 users and vest some coins for users to trade with each other. The vesting is implemented by transferring the coins from miner to each user.

Forth, each user will start to randomly trade with each other and submit his transaction to the trasnactuon pool for mining.

Fifty, miner keeps mining, validate transaction, confirm block.

 Cool, we got that? Great. Let's actually build this thing now. 


## Building the code

First, you need to install Golang
if Mac

	brew install go 

if Ubuntu

	sudo apt-get update && sudo apt-get -y upgrade && sudo apt-get install -y golang-go

Once you have installed and rebooted, log in, then open up the program “terminal.” Now run the command…

	sudo apt-get update && sudo apt-get -y upgrade && sudo apt-get -y install git

Now you can clone the repository to get the full code.

	cd ~/ && git clone https://github.com/codingtmd/mini-blockchain.git 

Do not forget to download the missing library:

	go get -d -v .


Go to the directory and build the code 

	go build

Then run it with fun

	go run main.go

And you will see the console output as below
![workflow](https://drive.google.com/uc?export=view&id=1SDnBbREANWRk2DnqipcTDwInm-EeI1Vk)

If you need an IDE, normally I use [microsoft visual studio code](https://code.visualstudio.com/download "microsoft visual studio code") and [Go Plugin](https://code.visualstudio.com/docs/languages/go "Go Plugin")


## Things to know

The logging uses loggo. Plz check the configuration and usage here: https://github.com/juju/loggo

## Cool future work / Areas you can contribute / TODOs

 - Use msg to communicate infro between miner, user. (Currently just function call)
 - Add multi-miner support
 - A simple script to initialize a blockchain.
 - Create a config for all hard code parameters.
 - Better logging.
 - Make a web UI to look into operation details, like etherscan.io
 - Create PoS support
 - Create a wallet implementation and web UI
 - Add some animations that make it easier to understand.
 - Build a dApp
 - Add ICO simulation like how to vest coins to user
