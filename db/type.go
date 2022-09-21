package db

import (
	"bytes"
	"encoding/gob"
	"math/big"
	"time"

	"github.com/X-mob/mob-watcher/lib"
	"github.com/ethereum/go-ethereum/common"
)

const MOB_KEY_PREFIX = "mob:"
const PLAYER_KEY_PREFIX = "player:"
const JOIN_DETAIL_PREFIX = "join:detail:"
const SCAN_MOB_CURSOR_KEY = "cursor:mob:blockNum"
const SCAN_JOIN_CURSOR_KEY = "cursor:join:blockNum"

type MobStatus int64

const (
	Raising     MobStatus = iota
	RaiseFailed           // the deadline is reached but raise not enough
	RaiseFinished
	NftBought
	NftSold
	CanClaim
	AllClaimed
)

func (s MobStatus) String() string {
	switch s {
	case Raising:
		return "Raising"
	case RaiseFailed:
		return "RaiseFailed"
	case RaiseFinished:
		return "RaiseFinished"
	case NftBought:
		return "NftBought"
	case NftSold:
		return "NftSold"
	case CanClaim:
		return "CanClaim"
	case AllClaimed:
		return "AllClaimed"
	}
	return "unknown"
}

type Mob struct {
	Address              common.Address // primary key
	Name                 string
	Creator              common.Address
	Token                common.Address
	TokenId              *big.Int
	RaisedTotal          *big.Int
	AmountTotal          *big.Int
	TakeProfitPrice      *big.Int
	StopLossPrice        *big.Int
	Fee                  *big.Int
	Deadline             *big.Int
	RaisedAmountDeadline *big.Int
	Balance              *big.Int
	CreatedTime          int64
	Status               MobStatus
}

type Player struct {
	Address  common.Address // primary key
	JoinMobs []common.Address
}

type JoinDetail struct {
	Player     common.Address // combined primary key
	Mob        common.Address // combined primary key
	DepositEth big.Int
}

type MobMember struct {
	address common.Address // primary key
	member  []common.Address
}

func GetMobSaveKey(address string) string {
	return MOB_KEY_PREFIX + address
}

func GetPlayerSaveKey(address string) string {
	return PLAYER_KEY_PREFIX + address
}

func GetJoinDetailSaveKey(playerAddress string, mobAddress string) string {
	return JOIN_DETAIL_PREFIX + playerAddress + ":" + mobAddress
}

// todo: use generic after update go 1.8
func SerializeMob(mob Mob) string {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(mob)
	if err != nil {
		panic(err)
	}
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

func SerializePlayer(player Player) string {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(player)
	if err != nil {
		panic(err)
	}
	return data.String()
}

func DeSerializePlayer(data []byte) Player {
	dataBuf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(dataBuf)
	var player Player
	err := dec.Decode(&player)
	if err != nil {
		panic(err)
	}
	return player
}

func SerializeJoinDetail(joinDetail JoinDetail) string {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(joinDetail)
	if err != nil {
		panic(err)
	}
	return data.String()
}

func DeSerializeJoinDetail(data []byte) JoinDetail {
	dataBuf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(dataBuf)
	var joinDetail JoinDetail
	err := dec.Decode(&joinDetail)
	if err != nil {
		panic(err)
	}
	return joinDetail
}

func SerializeMobMember(mobMember MobMember) string {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(mobMember)
	if err != nil {
		panic(err)
	}
	return data.String()
}

func DeSerializeMobMember(data []byte) MobMember {
	dataBuf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(dataBuf)
	var mobMember MobMember
	err := dec.Decode(&mobMember)
	if err != nil {
		panic(err)
	}
	return mobMember
}

// convertor
func MobCreateToMob(m lib.XmobManageMobCreate) Mob {
	return Mob{
		Address:              m.Proxy,
		Name:                 m.Name,
		Creator:              m.Creator,
		Token:                m.Token,
		TokenId:              m.TokenId,
		RaisedTotal:          m.RaisedTotal,
		AmountTotal:          big.NewInt(0),
		TakeProfitPrice:      m.TakeProfitPrice,
		StopLossPrice:        m.StopLossPrice,
		Fee:                  m.Fee,
		Deadline:             m.Deadline,
		RaisedAmountDeadline: m.RasiedAmountDeadline,
		Balance:              big.NewInt(0),
		CreatedTime:          time.Now().Unix(),
		Status:               Raising,
	}
}
