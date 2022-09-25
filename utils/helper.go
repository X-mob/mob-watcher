package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func StringToBigInt(num string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(num, 10)
	if !ok {
		panic("convert failed")
	}
	return n
}

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
	str := "0x42"
	sig := common.Hex2Bytes(str)
	return sig
}

func ZeroAddress() common.Address {
	addrBytes := [20]byte{0}
	addr := common.BytesToAddress(addrBytes[:])
	return addr
}
