package db

import (
	"fmt"
	"math/big"

	"github.com/X-mob/mob-watcher/config"
	badger "github.com/dgraph-io/badger/v3"
)

var DbClient *badger.DB
var DefaultStorePath = "/tmp/badger"

const MOB_KEY_PREFIX = "mob:"

func init() {
	path := config.GlobalConfig.StorePath
	if path == "" {
		path = DefaultStorePath
	}
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}
	DbClient = db
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
		handle(err)
		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		handle(err)
		fmt.Printf("The answer is: %s\n", valCopy)
		return nil
	})
	handle(err)
	return valCopy
}

func SetMob(mob Mob) {
	value := SerializeMob(mob)
	key := getMobSaveKey(mob.Address.Hex())
	SetKV(key, value)
}

func GetMob(address string) Mob {
	key := getMobSaveKey(address)
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

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func getMobSaveKey(address string) string {
	return MOB_KEY_PREFIX + address
}
