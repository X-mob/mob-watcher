package db

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/X-mob/mob-watcher/config"
	badger "github.com/dgraph-io/badger/v3"
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

func GetAllMobAddress() []string {
	keys := GetKeyWithPrefix(MOB_KEY_PREFIX)
	var addrs []string
	for _, key := range keys {
		addr := strings.Join(strings.Split(key, MOB_KEY_PREFIX), "")
		addrs = append(addrs, addr)
	}
	return addrs
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

func SetMob(mob Mob) {
	value := SerializeMob(mob)
	key := GetMobSaveKey(mob.Address.Hex())
	SetKV(key, value)
}

func GetMob(address string) Mob {
	key := GetMobSaveKey(address)
	data := GetKVWithBytes(key)
	return DeSerializeMob(data)
}

func UpdateMobStatus(address string, newStatus MobStatus) {
	mob := GetMob(address)
	mob.Status = newStatus
	SetMob(mob)
}

func UpdateMobBalance(address string, newBalance big.Int) {
	mob := GetMob(address)
	mob.Balance = &newBalance
	SetMob(mob)
}

func UpdateMobAmountTotal(address string, amountTotal big.Int) {
	mob := GetMob(address)
	mob.AmountTotal = &amountTotal
	SetMob(mob)
}

func SetPlayer(playerAddress string, mobAddress string) {

}

func StringToBigInt(num string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(num, 10)
	if !ok {
		panic("convert failed")
	}
	return n
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
