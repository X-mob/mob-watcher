package task

import (
	"bytes"
	"encoding/gob"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	GethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TxDataHash = string // TxData's hash: tx.to + tx.data + tx.value

type TxData struct {
	To    common.Address
	Data  []byte
	Value *big.Int
}

func (txData TxData) Serialize() []byte {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(txData)
	if err != nil {
		panic(err)
	}
	return data.Bytes()
}

func (txData TxData) Hash() common.Hash {
	hash := crypto.Keccak256Hash(txData.Serialize())
	return hash
}

type TxStatus int

const (
	NotFound TxStatus = iota
	NotSend
	Sent
	OnChain
	Failed // fail tx
)

func (s TxStatus) String() string {
	switch s {
	case NotFound:
		return "NotFound"
	case NotSend:
		return "NotSend"
	case Sent:
		return "Sent"
	case OnChain:
		return "OnChain"
	case Failed:
		return "Failed"
	}
	return "unknown"
}

type Tx struct {
	DataHash TxDataHash // primary key, id
	Hash     string     // same with ethereum txHash
	Status   TxStatus
	Receipt  *GethTypes.Receipt
}

type Pool = map[TxDataHash]*Tx

type TxManagerInterface interface {
	GetPool() Pool
	SavePoolSnapshot()
	Add(tx TxData)
	Check(id TxDataHash) TxStatus
	Update(id TxDataHash, newTxStatus TxStatus)
}

type TxManager struct {
	pool Pool
}

func (txManager TxManager) SavePoolSnapshot() {

}

func (txManager TxManager) Add(txData TxData) {
	tx := Tx{
		DataHash: txData.Hash().Hex(),
		Status:   NotSend,
	}
	txManager.pool[txData.Hash().Hex()] = &tx
}

func (txManager TxManager) Check(id TxDataHash) TxStatus {
	tx := txManager.pool[id]
	if tx == nil {
		return NotFound
	}

	return tx.Status
}

func (txManager TxManager) Update(id TxDataHash, newTxStatus TxStatus) {
	tx := txManager.pool[id]
	if tx == nil {
		panic("tx not found")
	}

	tx.Status = newTxStatus
	txManager.pool[id] = tx
}

func (txManager TxManager) GetTx(id TxDataHash) *Tx {
	tx := txManager.pool[id]
	return tx
}

func (txManager TxManager) UpdateTx(tx *Tx) {
	id := tx.DataHash
	txManager.pool[id] = tx
}

func NewTxManager() *TxManager {
	txManager := new(TxManager)
	txManager.pool = Pool{}
	return txManager
}
