package util

import (
	"crypto/md5"
	"encoding/hex"

	"../config"

	"github.com/cnf/structhash"
)

func Hash(c interface{}) string {
	hash, _ := structhash.Hash(c, 1)
	return hash[len(hash)-5:]
}

func HashBytes(c [config.HashSize]byte) string {
	hash := md5.Sum([]byte(string(c[:])))
	md5InString := hex.EncodeToString(hash[:])
	return md5InString[len(md5InString)-5:]
}
