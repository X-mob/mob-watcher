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

// we use fixed salt for selling order from mob contract
// since the signature is not related to salt value,
// using fixed value is helpful to avoid re-submit the same listing
const FixedSalt32BytesHexNumber = "0x0000000000000000000000000000000000000000000000000000000000000001"

// minus the 2.5% fee directly from the selling price
// 	fee = price * 0.025
// 	earning price = price * 0.975
func CalcOpenSeaFeeByBasePrice(price *big.Int) string {
	priceFloat := new(big.Float).SetInt(price)
	fee := new(big.Float).Mul(priceFloat, big.NewFloat(0.025))

	// up to more than 1, otherwise opensea will complain
	fee = fee.Add(fee, big.NewFloat(0.999999))

	feeRound := new(big.Int)
	feeRound, _ = fee.Int(feeRound)
	return feeRound.String()
}

// fixed desire selling price, add more to enable 2.5% fee taken by OpenSea
// 	earning = price
// 	fee = price / 0.975 * 0.025
func CalcOpenSeaFeeByFixedPrice(price *big.Int) string {
	priceFloat := new(big.Float).SetInt(price)

	fee := new(big.Float).Quo(priceFloat, big.NewFloat(0.975))
	fee = fee.Mul(fee, big.NewFloat(0.025))

	// up to more than 1, otherwise opensea will complain
	fee = fee.Add(fee, big.NewFloat(0.999999))

	feeRound := new(big.Int)
	feeRound, _ = fee.Int(feeRound)
	return feeRound.String()
}
