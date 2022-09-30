package casting

import (
	"encoding/hex"
	"math/big"
	"strings"
)

func StringToBigInt(num string, base int) *big.Int {
	a := new(big.Int)
	a.SetString(strings.TrimPrefix(num, "0x"), base)
	return a
}

func HexStringToBigInt(hexString string) *big.Int {
	decodedByteArray, err := hex.DecodeString(strings.TrimPrefix(hexString, "0x"))

	if err != nil {
		panic("Unable to convert hex to byte.")
	}
	a := new(big.Int)
	a.SetBytes(decodedByteArray)
	return a
}

func HexStringToByte32(hexString string) [32]byte {
	if hexString == "0x0000000000000000000000000000000000000000000000000000000000000000" {
		return [32]byte{0}
	}

	decodedByteArray, err := hex.DecodeString(strings.TrimPrefix(hexString, "0x"))

	if err != nil {
		panic("Unable to convert hex to byte.")
	}

	var arr [32]byte
	copy(arr[:], decodedByteArray)
	return arr
}

func HexStringToBytes(hexString string) []byte {
	decodedByteArray, err := hex.DecodeString(strings.TrimPrefix(hexString, "0x"))
	if err != nil {
		panic("Unable to convert hex to byte.")
	}
	return decodedByteArray
}

func HexStringToByte65(hexString string) []byte {
	decodedByteArray, err := hex.DecodeString(strings.TrimPrefix(hexString, "0x"))

	if err != nil {
		panic("Unable to convert hex to byte.")
	}

	var arr [65]byte
	copy(arr[:], decodedByteArray)
	return arr[:]
}
