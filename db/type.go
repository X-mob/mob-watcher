package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MobStatus int64

const (
	Raising     MobStatus = iota
	RaiseFailed           // the deadline is reached but raise not enough
	RaiseFinished
	NftBought
	NftSold
	CanClaim
	Claimed
)

type Mob struct {
	Name                 string
	Creator              common.Address
	Token                common.Address
	TokenId              *big.Int
	Address              common.Address
	RaisedTotal          *big.Int
	TakeProfitPrice      *big.Int
	StopLossPrice        *big.Int
	Fee                  *big.Int
	Deadline             *big.Int
	RaisedAmountDeadline *big.Int
	Balance              *big.Int
	CreatedTime          int64
	Status               MobStatus
}

func SerializeMob(mob Mob) string {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(mob)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Bytes())
	return data.String()
}

func DeSerializeMob(data []byte) Mob {
	dataBuf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(dataBuf)
	var mob Mob
	err := dec.Decode(&mob)
	if err != nil {
		panic(err)
	}
	return mob
}

type Player struct {
	// todo
}
