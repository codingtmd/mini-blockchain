package role

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/juju/loggo"

	"../core"
	"../util"
)

type User struct {
	chain core.Blockchain

	key     *rsa.PrivateKey
	Address rsa.PublicKey
}

/*
 * CreateBoostUser to create the first user before boosting the chain
 */
func CreateBoostUser() *User {
	return createUser()
}

/*
 * RegisterBoostUser used to register the boost user after initialize the blockchain
 */
func (user *User) RegisterBoostUser(chain core.Blockchain) {
	user.chain = chain

	utxoMap := make(map[core.UTXO]bool)
	chain.RegisterUser(user.Address, utxoMap)
	user.getLogger().Debugf("Register boost user %v\n", user.GetShortIdentity())

}

func CreateUser(chain core.Blockchain) *User {
	user := createUser()
	user.chain = chain

	utxoMap := make(map[core.UTXO]bool)
	chain.RegisterUser(user.Address, utxoMap)

	return user
}

func createUser() *User {
	var user User

	account, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		user.getLogger().Errorf("Cannot create account: %s\n", err)
		return nil
	}

	user.Address = account.PublicKey
	user.key = account

	user.getLogger().Debugf("Created a user at %v\n", user.GetShortIdentity())
	return &user
}

func (user *User) Balance() uint64 {
	return user.chain.BalanceOf(&user.Address)
}

func (user *User) SendTo(receipt *User, amount uint64, fee uint64) {
	tran, err := user.chain.TransferCoin(&user.Address, &receipt.Address, amount, fee)
	if err != nil {
		user.getLogger().Errorf("Failed to create transaction: %v\n", err)
		return
	}

	tran.SignTransaction([]*rsa.PrivateKey{user.GetPrivateKey()})

	user.getLogger().Debugf("%s\n", tran.Print())
	user.chain.AcceptBroadcastedTransaction(tran)
	user.getLogger().Infof("User %v sends %d coins to user %v\n", user.GetShortIdentity(), amount, receipt.GetShortIdentity())
}

/*
 *  BroadcastTransaction broadcasts the transaction to all miners
 *  TODO: The broadcast should be based on msg in real world
 */
func (user *User) BroadcastTransaction(tran *core.Transaction) {
	user.chain.AcceptBroadcastedTransaction(tran)
}

func (user *User) GetShortIdentity() string {
	return util.GetShortIdentity(user.Address)
}

func (user *User) GetPrivateKey() *rsa.PrivateKey {
	return user.key
}

func (user *User) getLogger() loggo.Logger {
	return util.GetUserLogger(user.GetShortIdentity())
}
