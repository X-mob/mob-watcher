package utils

import "math/big"

func DecimalStringToBigInt(num string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(num, 10)
	if !ok {
		panic("convert failed")
	}
	return n
}
