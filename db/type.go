package db

import (
	"bytes"
	"encoding/gob"
	"math/big"
	"time"

	"github.com/X-mob/mob-watcher/lib"
	"github.com/ethereum/go-ethereum/common"
)

type MobStatus int64

const (
	Raising     MobStatus = iota
	RaiseFailed           // the deadline is reached but raise not enough
	RaiseSuccess
	NftBought
	NftBuyFailed // the deadline is reached but not bought yet
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
	case RaiseSuccess:
		return "RaiseSuccess"
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
	Address         common.Address `json:"address"` // primary key
	Name            string         `json:"name"`
	Creator         common.Address `json:"creator"`
	Token           common.Address `json:"token"`
	TokenId         *big.Int       `json:"token_id"`
	RaiseTarget     *big.Int       `json:"raise_target"`
	RaisedAmount    *big.Int       `json:"raise_amount"`
	TakeProfitPrice *big.Int       `json:"take_profit_price"`
	StopLossPrice   *big.Int       `json:"stop_loss_price"`
	Fee             *big.Int       `json:"fee"`
	Deadline        *big.Int       `json:"deadline"`
	RaiseDeadline   *big.Int       `json:"raise_deadline"`
	Balance         *big.Int       `json:"balance"`
	CreatedTime     int64          `json:"created_time"`
	Status          MobStatus      `json:"status"`
	TargetMode      uint8          `json:"target_mode"`
	Asset           MobAsset       `json:"asset"`
}

type MobAsset struct {
	BackgroundColor      string           `json:"background_color"`
	ImageUrl             string           `json:"image_url"`
	ImagePreviewUrl      string           `json:"image_preview_url"`
	ImageThumbnailUrl    string           `json:"image_thumbnail_url"`
	AnimationUrl         string           `json:"animation_url"`
	AnimationOriginalUrl string           `json:"animation_original_url"`
	AssetContract        MobAssetContract `json:"asset_contract"`
}

type MobAssetContract struct {
	Name              string `json:"name"`
	CreatedDate       string `json:"created_date"`
	Address           string `json:"address"`
	CollectionSlug    string `json:"collection_slug"`
	ImageUrl          string `json:"image_url"`
	SchemaName        string `json:"schema_name"`
	Symbol            string `json:"symbol"`
	TotalSupply       string `json:"total_supply"`
	ExternalLink      string `json:"external_link"`
	Description       string `json:"description"`
	AssetContractType string `json:"asset_contract_type"`
}

type Player struct {
	Address  common.Address // primary key
	JoinMobs []common.Address
}

type JoinDetail struct {
	Player     common.Address `json:"player"` // combined primary key
	Mob        common.Address `json:"mob"`    // combined primary key
	DepositEth *big.Int       `json:"deposit_eth"`
}

type MobMember struct {
	Address common.Address // primary key
	Member  []common.Address
}

func GetMobSaveKey(address string) string {
	return MOB_KEY_PREFIX + address
}

func GetPlayerSaveKey(playerAddress string) string {
	return PLAYER_KEY_PREFIX + playerAddress
}

func GetMobMemberSaveKey(mobAddress string) string {
	return MOB_MEMBER_KEY_PREFIX + mobAddress
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
		Address:         m.Proxy,
		Name:            m.Name,
		Creator:         m.Creator,
		Token:           m.Token,
		TokenId:         m.TokenId,
		RaiseTarget:     m.RaiseTarget,
		RaisedAmount:    big.NewInt(0),
		TakeProfitPrice: m.TakeProfitPrice,
		StopLossPrice:   m.StopLossPrice,
		Fee:             m.Fee,
		Deadline:        m.Deadline,
		RaiseDeadline:   m.RaiseDeadline,
		Balance:         big.NewInt(0),
		CreatedTime:     time.Now().Unix(),
		Status:          Raising,
		TargetMode:      m.TargetMode,
	}
}
