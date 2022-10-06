package db

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/X-mob/mob-watcher/config"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
)

var DbClient *badger.DB
var DefaultStorePath = "/tmp/badger"

func init() {
	path := config.GlobalConfig.StorePath
	if path == "" {
		path = DefaultStorePath
	}
	options := badger.DefaultOptions(path)
	options.Logger = nil
	db, err := badger.Open(options)
	if err != nil {
		panic(err)
	}
	DbClient = db
}

func GetKeyWithPrefix(prefix string) []string {
	var keys []string
	err := DbClient.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(prefix)

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				return nil
			})
			if err != nil {
				return err
			}
			keys = append(keys, string(k))
		}
		return nil
	})
	handle(err)
	return keys
}

func SetKV(key string, value string) {
	err := DbClient.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	if err != nil {
		panic(err)
	}
}

func GetKV(key string) string {
	data := GetKVWithBytes(key)
	return string(data[:])
}

func GetKVWithBytes(key string) []byte {
	var valCopy []byte
	err := DbClient.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			fmt.Println("get key err", err)
			return nil
		}

		valCopy, err = item.ValueCopy(nil)
		handle(err)
		return nil
	})
	handle(err)
	return valCopy
}

func IsKeyExits(key string) bool {
	err := DbClient.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if err != nil {
			fmt.Println("get key err", err)
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}

func IsMobExits(address string) bool {
	key := GetMobSaveKey(address)
	return IsKeyExits(key)
}

func GetAllMobAddress() []string {
	keys := GetKeyWithPrefix(MOB_KEY_PREFIX)
	var addrs []string
	for _, key := range keys {
		addr := strings.Join(strings.Split(key, MOB_KEY_PREFIX), "")
		addrs = append(addrs, addr)
	}
	return addrs
}

func GetMobAddressByStatus(status MobStatus) []string {
	addrs := GetAllMobAddress()
	var target []string
	for _, addr := range addrs {
		mob := GetMob(addr)
		if mob.Status == status {
			target = append(target, mob.Address.Hex())
		}
	}
	return target
}

func GetMobsWithStatus(status MobStatus) []Mob {
	addrs := GetAllMobAddress()
	var mobs []Mob
	for _, addr := range addrs {
		mob := GetMob(addr)
		if mob.Status == status {
			mobs = append(mobs, mob)
		}
	}
	return mobs
}

func AddMob(mob Mob) {
	value := SerializeMob(mob)
	key := GetMobSaveKey(mob.Address.Hex())
	SetKV(key, value)
}

func GetMob(address string) Mob {
	key := GetMobSaveKey(address)
	data := GetKVWithBytes(key)
	return DeSerializeMob(data)
}

// todo: use generic to refactor the following method
func UpdateMobStatus(address string, newStatus MobStatus) {
	mob := GetMob(address)
	mob.Status = newStatus
	AddMob(mob)
}

func UpdateMobTokenId(address string, newTokenId *big.Int) {
	mob := GetMob(address)
	mob.TokenId = newTokenId
	AddMob(mob)
}

func UpdateMobBalance(address string, newBalance big.Int) {
	mob := GetMob(address)
	mob.Balance = &newBalance
	AddMob(mob)
}

func UpdateMobRaisedAmount(address string, raisedAmount big.Int) {
	mob := GetMob(address)
	mob.RaisedAmount = &raisedAmount
	AddMob(mob)
}

func GetMobMember(mobAddress string) MobMember {
	key := GetMobMemberSaveKey(mobAddress)
	data := GetKVWithBytes(key)
	return DeSerializeMobMember(data)
}

func GetJoinDetail(playerAddress string, mobAddress string) JoinDetail {
	key := GetJoinDetailSaveKey(playerAddress, mobAddress)
	data := GetKVWithBytes(key)
	return DeSerializeJoinDetail(data)
}

func GetJoinDetailsByMobAddress(mobAddress string) []JoinDetail {
	mobMembers := GetMobMember(mobAddress)
	var details = []JoinDetail{}
	for _, m := range mobMembers.Member {
		detail := GetJoinDetail(m.Hex(), mobAddress)
		details = append(details, detail)
	}
	return details
}

// include repeat item settlement
func AddNewMember(playerAddress string, mobAddress string, value *big.Int) {
	AddPlayer(playerAddress, mobAddress)
	AddJoinDetail(playerAddress, mobAddress, value)
	AddMobMember(playerAddress, mobAddress)
}

func AddPlayer(playerAddress string, mobAddress string) {
	key := GetPlayerSaveKey(playerAddress)
	isExist := IsKeyExits(key)
	if isExist {
		player := DeSerializePlayer(GetKVWithBytes(key))
		mobAddr := common.HexToAddress(mobAddress)
		if ContainsAddress(player.JoinMobs, mobAddr) {
			// already contains, do nothing..
			return
		}

		player.JoinMobs = append(player.JoinMobs, mobAddr)
		SetKV(key, SerializePlayer(player))
	} else {
		player := Player{
			Address: common.HexToAddress(playerAddress),
			JoinMobs: []common.Address{
				common.HexToAddress(mobAddress),
			},
		}
		SetKV(key, SerializePlayer(player))
	}
}

func AddJoinDetail(playerAddress string, mobAddress string, value *big.Int) {
	key := GetJoinDetailSaveKey(playerAddress, mobAddress)
	isExist := IsKeyExits(key)
	if isExist {
		detail := DeSerializeJoinDetail(GetKVWithBytes(key))

		// add deposit value
		detail.DepositEth = detail.DepositEth.Add(detail.DepositEth, value)
	} else {
		detail := JoinDetail{
			Player:     common.HexToAddress(playerAddress),
			Mob:        common.HexToAddress(mobAddress),
			DepositEth: value,
		}
		SetKV(key, SerializeJoinDetail(detail))
	}
}

func AddMobMember(playerAddress string, mobAddress string) {
	key := GetMobMemberSaveKey(mobAddress)
	isExist := IsKeyExits(key)
	playerAddr := common.HexToAddress(playerAddress)
	if isExist {
		mobMember := DeSerializeMobMember(GetKVWithBytes(key))
		if ContainsAddress(mobMember.Member, playerAddr) {
			// already contains, do nothing..
			return
		}

		mobMember.Member = append(mobMember.Member, playerAddr)
		SetKV(key, SerializeMobMember(mobMember))
	} else {
		mobMember := MobMember{
			Address: common.HexToAddress(mobAddress),
			Member: []common.Address{
				playerAddr,
			},
		}
		SetKV(key, SerializeMobMember(mobMember))
	}
}

/**
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
*/

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// todo change to generic after go 1.8
func ContainsAddress(s []common.Address, e common.Address) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
