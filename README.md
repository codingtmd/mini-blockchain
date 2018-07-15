<img src="https://drive.google.com/uc?export=view&id=1B2ar0Mn8CV6GHCVg76Tw0TZ7a8lGZsPk" width="40%">

## Mini-Blockchain

Mini-Blockchain is a reference design for a blockchain system to demostate a full end2end flow in current blockchain technology. 

There are so many open source projects of blockchain implementation, like ethereum, bitcoin, eos etc.. But even running their code will take u couple of days to setup. 

The goal of the project is to build a "Hello-world" version of blockchain to help people understand how the foundamental technology without wasting time reading in a huge code base, to simulate the different roles, like user, miner, booster, and demostrate how they participate the blockchain network. And it also simplify the deployment, instead of configuring tons of library and debugging a version in a real distributed network, it is single machine version to easy the debug, test, and iterate.

Since it is primarily designed as a learning tool for people to learn the blockchain end to end. It is meant to be as simple to understand as possible while still providing clear structure and workflow. 

## What is this not?

Mini-Blockchain is **not** meant for real usage as it is a simplified version and has many hard coded parameters. It's just meant for learning.

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

**Nonce -** a 32-bit (4-byte) field to random the hash generation. Any change to the block data (such as the nonce) will make the block hash completely different. The resulting hash has to be a value less than the current difficulty and so will have to have a certain number of leading zero bits to be less than that. As this iterative calculation requires time and resources, the presentation of the block with the correct nonce value constitutes proof of work.

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

 - Use msg to communicate infro between miners, users. (Currently just function call)
 - Same-input tramnsaction merging
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

## Reading List
If you are not familiar with blockchain and its technology, below info will help u to ramp up the knowledge.
### Articles: 
+ Vision
	* [https://avc.com/2018/05/is-buying-crypto-assets-investing/](https://avc.com/2018/05/is-buying-crypto-assets-investing/)
	+ [https://news.earn.com/thoughts-on-tokens-436109aabcbe](https://news.earn.com/thoughts-on-tokens-436109aabcbe)
	+ [https://thecontrol.co/cryptoeconomics-101-e5c883e9a8ff](https://thecontrol.co/cryptoeconomics-101-e5c883e9a8ff)
	+ [https://medium.com/@cdixon/why-decentralization-matters-5e3f79f7638e](https://medium.com/@cdixon/why-decentralization-matters-5e3f79f7638e)
	+ [https://continuations.com/post/148098927445/crypto-tokens-and-the-coming-age-of-protocol](https://continuations.com/post/148098927445/crypto-tokens-and-the-coming-age-of-protocol)
	+ [https://medium.com/@cdixon/crypto-tokens-a-breakthrough-in-open-network-design-e600975be2ef](https://medium.com/@cdixon/crypto-tokens-a-breakthrough-in-open-network-design-e600975be2ef)
	+ [https://www.theinformation.com/articles/14-ways-the-cryptocurrency-market-will-change-in-2018](https://www.theinformation.com/articles/14-ways-the-cryptocurrency-market-will-change-in-2018)
	+ [https://medium.com/@FEhrsam/why-decentralized-exchange-protocols-matter-58fb5e08b320](https://medium.com/@FEhrsam/why-decentralized-exchange-protocols-matter-58fb5e08b320)
	+ [https://thecontrol.co/some-blockchain-reading-1d98ec6b2f39](https://thecontrol.co/some-blockchain-reading-1d98ec6b2f39)
+ App Token v.s. Protocol Token
	+ [https://blog.0xproject.com/the-difference-between-app-coins-and-protocol-tokens-7281a428348c](https://blog.0xproject.com/the-difference-between-app-coins-and-protocol-tokens-7281a428348c)
	+ [https://medium.com/blockchannel/protocol-tokens-good-for-greedy-investors-bad-for-business-9002b40cf4cc](https://medium.com/blockchannel/protocol-tokens-good-for-greedy-investors-bad-for-business-9002b40cf4cc)
	+ [https://blog.citowise.com/the-basics-coin-vs-token-what-is-the-difference-5cd270591538](https://blog.citowise.com/the-basics-coin-vs-token-what-is-the-difference-5cd270591538)
+ BaaS/Platform
	+ [https://medium.com/@ACINQ/strike-our-stripe-like-api-for-lightning-is-live-cd1dce76ce2e](https://medium.com/@ACINQ/strike-our-stripe-like-api-for-lightning-is-live-cd1dce76ce2e)
	+ [https://medium.com/@jbackus/blockchain-platform-plays-2827247a9014](https://medium.com/@jbackus/blockchain-platform-plays-2827247a9014)
	+ [https://techcrunch.com/2018/05/22/po-et-launches-lab-for-developers-to-build-apps-on-publishing-blockchain/](https://techcrunch.com/2018/05/22/po-et-launches-lab-for-developers-to-build-apps-on-publishing-blockchain/)
+ Business Model
	+ [https://blog.coinbase.com/app-coins-and-the-dawn-of-the-decentralized-business-model-8b8c951e734](https://blog.coinbase.com/app-coins-and-the-dawn-of-the-decentralized-business-model-8b8c951e734)
+ Consensus mechanism
	+ [https://www.coindesk.com/blockchains-feared-51-attack-now-becoming-regular/](https://www.coindesk.com/blockchains-feared-51-attack-now-becoming-regular/)
+ dApp
	+ [https://medium.com/@FEhrsam/the-dapp-developer-stack-the-blockchain-industry-barometer-8d55ec1c7d4](https://medium.com/@FEhrsam/the-dapp-developer-stack-the-blockchain-industry-barometer-8d55ec1c7d4)
+ Decentralized Exchange
	+ [https://medium.com/@FEhrsam/why-decentralized-exchange-protocols-matter-58fb5e08b320](https://medium.com/@FEhrsam/why-decentralized-exchange-protocols-matter-58fb5e08b320)
	+ [https://www.reuters.com/article/crypto-currencies-coinbase/coinbase-acquires-cryptocurrency-trading-platform-paradex-idUSL2N1SU1KK](https://www.reuters.com/article/crypto-currencies-coinbase/coinbase-acquires-cryptocurrency-trading-platform-paradex-idUSL2N1SU1KK)
+ Digital
	+ [https://medium.com/kinfoundation/kin-blockchain-taking-fate-into-our-own-hands-f5bdfa759502](https://medium.com/kinfoundation/kin-blockchain-taking-fate-into-our-own-hands-f5bdfa759502)
+ ENS
	+ [https://medium.com/the-ethereum-name-service/a-beginners-guide-to-buying-an-ens-domain-3ccac2bdc770](https://medium.com/the-ethereum-name-service/a-beginners-guide-to-buying-an-ens-domain-3ccac2bdc770)
+ ERC-20:
	+ [https://medium.com/@james_3093/ethereum-erc20-tokens-explained-9f7f304055df](https://medium.com/@james_3093/ethereum-erc20-tokens-explained-9f7f304055df)
	+ [https://medium.com/0xcert/fungible-vs-non-fungible-tokens-on-the-blockchain-ab4b12e0181a](https://medium.com/0xcert/fungible-vs-non-fungible-tokens-on-the-blockchain-ab4b12e0181a)
	+ [https://hackernoon.com/an-overview-of-non-fungible-tokens-5f140c32a70a](https://hackernoon.com/an-overview-of-non-fungible-tokens-5f140c32a70a)
+ Fat Protocol
	+ [https://www.usv.com/blog/fat-protocols](https://www.usv.com/blog/fat-protocols)
+ Game
	+ [https://www.usv.com/blog/cryptokitties-1](https://www.usv.com/blog/cryptokitties-1)
	+ [https://techcrunch.com/2018/05/25/gravys-new-mobile-game-show-is-price-is-right-mixed-with-qvc/](https://techcrunch.com/2018/05/25/gravys-new-mobile-game-show-is-price-is-right-mixed-with-qvc/)
	+ [http://www.businessinsider.com/fortnite-esports-prize-pool-money-epic-games-2018-5](http://www.businessinsider.com/fortnite-esports-prize-pool-money-epic-games-2018-5)
+ Login kit
	+ [https://medium.com/cleargraphinc/introducing-cleargraph-4713bc215a77](https://medium.com/cleargraphinc/introducing-cleargraph-4713bc215a77)
	+ [https://tokensale.civic.com/CivicTokenSaleWhitePaper.pdf](https://tokensale.civic.com/CivicTokenSaleWhitePaper.pdf)
+ Loyalty
	+ [https://medium.com/@bitrewards/why-blockchain-is-a-smart-solution-for-loyalty-programs-9443af408f71](https://medium.com/@bitrewards/why-blockchain-is-a-smart-solution-for-loyalty-programs-9443af408f71)
	+ [http://www.oliverwyman.com/our-expertise/insights/2017/mar/Blockchain-Will-Transform-Customer-Loyalty-Programs.html](http://www.oliverwyman.com/our-expertise/insights/2017/mar/Blockchain-Will-Transform-Customer-Loyalty-Programs.html)
	+ [http://www.kaleidoinsights.com/analysis-should-blockchain-power-your-customer-loyalty-program/](http://www.kaleidoinsights.com/analysis-should-blockchain-power-your-customer-loyalty-program/)
+ Proof of work v.s. Proof of stake
	+ [https://medium.com/@robertgreenfieldiv/explaining-proof-of-stake-f1eae6feb26f](https://medium.com/@robertgreenfieldiv/explaining-proof-of-stake-f1eae6feb26f)
+ Recorded video
	+ [https://avc.com/2017/12/video-of-the-week-token-1-0-vs-token-2-0/](https://avc.com/2017/12/video-of-the-week-token-1-0-vs-token-2-0/)
	+ [https://avc.com/2017/12/video-of-the-week-the-token-summit-conversation/](https://avc.com/2017/12/video-of-the-week-the-token-summit-conversation/)
+ Research report:
	+ [https://coincenter.org/report](https://coincenter.org/report)
+ Rewarding system
	+ [https://techcrunch.com/2018/05/22/tango-card-raises-35m-for-its-rewards-as-a-service-gift-card-aggregation-platform/](https://techcrunch.com/2018/05/22/tango-card-raises-35m-for-its-rewards-as-a-service-gift-card-aggregation-platform/)
	+ [https://medium.com/@bitrewards/why-blockchain-is-a-smart-solution-for-loyalty-programs-9443af408f71](https://medium.com/@bitrewards/why-blockchain-is-a-smart-solution-for-loyalty-programs-9443af408f71)
	+ [https://techcrunch.com/2018/05/11/hollywood-producer-plans-to-incentivise-content-viewers-with-tokens/](https://techcrunch.com/2018/05/11/hollywood-producer-plans-to-incentivise-content-viewers-with-tokens/)
+ Security Token
	+ [https://medium.com/@mkogan4/what-the-heck-are-tokenised-securities-7cd1123cbdad](https://medium.com/@mkogan4/what-the-heck-are-tokenised-securities-7cd1123cbdad)
+ Smart Contract
	+ [https://medium.com/@ninosm/squashing-bugs-and-stopping-heists-the-coming-arms-race-in-smart-contract-infrastructure-9666fb830f65](https://medium.com/@ninosm/squashing-bugs-and-stopping-heists-the-coming-arms-race-in-smart-contract-infrastructure-9666fb830f65)
+ Token
	+ [https://thecontrol.co/tokens-tokens-and-more-tokens-d4b177fbb443](https://thecontrol.co/tokens-tokens-and-more-tokens-d4b177fbb443)


### Some white papers
+ [Bitcoin: A Peer-to-Peer Electronic Cash System](https://bitcoin.org/bitcoin.pdf)
+ [Ethereum: A Next Generation Smart Contract and Decentralized Application Platform](https://github.com/ethereum/wiki/wiki/White-Paper)
+ [ETHEREUM: A SECURE DECENTRALISED GENERALISED TRANSACTION LEDGER](https://ethereum.github.io/yellowpaper/paper.pdf)
+ [BeigePaper: A Ether Tech Spec](https://github.com/chronaeon/beigepaper/blob/master/beigepaper.pdf)
+ [Enabling Blockchain Innovations with Pegged Sidechains](https://blockstream.com/sidechains.pdf)
+ [Augur: A Decentralized, Open-source Platform for Prediction Markets](https://bravenewcoin.com/assets/Whitepapers/Augur-A-Decentralized-Open-Source-Platform-for-Prediction-Markets.pdf)
+ [The Dai Stablecoin System](https://github.com/makerdao/docs/blob/master/Dai.md)
+ [Sia: Simple Decentralized Storage](http://www.sia.tech/sia.pdf)
+ [OmniLedger: A Secure, Scale-Out, Decentralized Ledger via Sharding](https://eprint.iacr.org/2017/406.pdf)
+ [Bitcoin UTXO Lifespan Prediction](http://cs229.stanford.edu/proj2015/225_report.pdf)

 ## Authors
Lei Zhang, https://www.linkedin.com/in/codingtmd

## Contact me

Leave me a message if you want to participate for fun.