package core

import (
	"crypto/rsa"
	"encoding/binary"
)

func appendUint32(data []byte, Value uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, Value)
	return append(data, b...)
}

func appendUint64(data []byte, Value uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, Value)
	return append(data, b...)
}

func appendAddress(data []byte, key *rsa.PublicKey) []byte {
	data = appendUint32(data, uint32(key.E))
	keyBytes := key.N.Bytes()
	data = appendUint32(data, uint32(len(keyBytes)))
	return append(data, keyBytes...)
}

func appendUint256(data []byte, Value uint256) []byte {
	for i := 0; i < 4; i++ {
		data = appendUint64(data, Value.data[i])
	}
	return data
}
