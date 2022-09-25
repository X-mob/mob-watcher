package casting

import "math/big"

func StringToBigInt(num string, base int) *big.Int {
	a := new(big.Int)
	a.SetString(num, base)
	return a
}

func StringToByte32(str string) [32]byte {
	var arr [32]byte
	copy(arr[:], str)
	return arr
}
