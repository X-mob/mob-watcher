package utils

import (
	"math/big"
	"testing"
)

func TestCalcOpenSeaFee(t *testing.T) {
	price := big.NewInt(2000000000000000)
	fee := CalcOpenSeaFeeByFixedPrice(price)
	t.Log(fee)
}
