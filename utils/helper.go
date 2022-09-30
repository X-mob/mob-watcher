package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GenRandomSalt(length uint8) *big.Int {
	salt := new(big.Int)
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("gen randomBytes %d failed", length)
		fmt.Println(err)
		panic("gen randomBytes failed")
	}
	salt.SetBytes(randomBytes[:])
	return salt
}

func MagicSignature() []byte {
	sig := common.Hex2Bytes(MAGIC_SIG_STR)
	return sig
}

func ZeroAddress() common.Address {
	addrBytes := [20]byte{0}
	addr := common.BytesToAddress(addrBytes[:])
	return addr
}

func Zero32BytesHexString() string {
	return "0x0000000000000000000000000000000000000000000000000000000000000000"
}
