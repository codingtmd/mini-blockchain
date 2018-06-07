package util

import (
	"crypto/md5"
	"encoding/hex"

	"../config"

	"github.com/cnf/structhash"
	"github.com/rs/xid"
)

func Hash(c interface{}) string {
	hash, _ := structhash.Hash(c, 1)
	return hash[len(hash)-5:]
}

func hashBytes(c []byte) string {
	hash := md5.Sum(c)
	md5InString := hex.EncodeToString(hash[:])
	return md5InString[len(md5InString)-5:]
}

func HashBytes(c [config.HashSize]byte) string {
	bytes := []byte(string(c[:]))
	return hashBytes(bytes)
}

func GetShortedUniqueId() string {
	id := xid.New().String()
	return hashBytes([]byte(id))
}
