package lib

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"runtime"

	"github.com/X-mob/mob-watcher/config"
	"github.com/X-mob/mob-watcher/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/eth-go-bindings/erc721"
)

func GetFeeRate() {
	feeRate, err := XmobManageInstance.FeeRate(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract fee rate:", feeRate)
}

func GetMobById(id *big.Int) *XmobExchangeCore {
	mobAddress, err := XmobManageInstance.MobsById(nil, id)
	if err != nil {
		panic(err)
	}
	return GetMobByAddress(mobAddress.Hex())
}

func GetMobByAddress(address string) *XmobExchangeCore {
	mob, err := NewXmobExchangeCore(common.HexToAddress(address), EthClient)
	if err != nil {
		panic(err)
	}
	return mob
}

func GetMobStatus(address string) uint8 {
	metadata := GetMobMetadata(address)
	return metadata.Status
}

type Metadata struct {
	Name            string
	Creator         common.Address
	Token           common.Address
	TokenId         *big.Int
	RaisedAmount    *big.Int
	RaiseTarget     *big.Int
	TakeProfitPrice *big.Int
	StopLossPrice   *big.Int
	Fee             *big.Int
	Deadline        uint64
	RaiseDeadline   uint64
	TargetMode      uint8
	Status          uint8
}

func GetMobMetadata(address string) Metadata {
	mob := GetMobByAddress(address)
	metadata, err := mob.Metadata(nil)
	if err != nil {
		panic(err)
	}
	return metadata
}

func CreateMob(
	_name string,
	_token string,
	_tokenId int64,
	_raiseTarget string,
	_takeProfitPrice string,
	_stopLossPrice string,
	_raiseDeadline uint64,
	_deadline uint64,
	_targetMode uint8) {

	token := common.HexToAddress(_token)
	tokenId := big.NewInt(_tokenId)
	raiseTarget := utils.StringToBigInt(_raiseTarget)
	takeProfitPrice := big.NewInt(25)
	stopLossPrice := big.NewInt(1)
	raiseDeadline := _raiseDeadline
	deadline := _deadline
	targetMode := _targetMode
	name := _name

	txOpts := NewTxOpts(nil)
	tx, err := XmobManageInstance.CreateMob(txOpts, token, tokenId, raiseTarget, takeProfitPrice, stopLossPrice, raiseDeadline, deadline, targetMode, name)
	CheckAndWaitTx(tx, err)
}

func JoinMob(value string, mobAddress string) {
	var txOpts *bind.TransactOpts = NewTxOpts(utils.StringToBigInt(value))
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.JoinPay(txOpts, txOpts.From)
	CheckAndWaitTx(tx, err)
}

func Claim(mobAddress string) {
	mob := GetMobByAddress(mobAddress)

	// check if claim is allowed
	metadata, err := mob.Metadata(nil)
	if err != nil {
		panic(err)
	}
	canClaim := metadata.Status == 3 // 3 means can claim MobStatus in xmob contract
	if canClaim == false {
		panic("can not claim yet")
	}

	// claim
	txOpts := NewTxOpts(nil)
	tx, err := mob.Claim(txOpts)
	CheckAndWaitTx(tx, err)
}

func Settle(mobAddress string) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.SettlementAllocation(txOpts)
	CheckAndWaitTx(tx, err)
}

func SettleAfterDeadline(mobAddress string) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.SettlementAfterDeadline(txOpts)
	CheckAndWaitTx(tx, err)
}

func SettleAfterBuyFailed(mobAddress string) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.SettlementAfterBuyFailed(txOpts)
	CheckAndWaitTx(tx, err)
}

func BuyBasicOrder(mobAddress string, order BasicOrderParameters) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.BuyBasicOrder(txOpts, order)
	CheckAndWaitTx(tx, err)
}

func BuyOrder(mobAddress string, order Order) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.BuyOrder(txOpts, order, [32]byte{0})
	CheckAndWaitTx(tx, err)
}

func RegisterSellOrder(mobAddress string, orders []Order) {
	isUnowned := IsNFTUnOwned(mobAddress)
	if isUnowned == true {
		fmt.Printf("contract has no nft, skip to sell")
		return
	}

	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.RegisterSellOrder(txOpts, orders)
	CheckAndWaitTx(tx, err)
}

func ValidateSellOrder(mobAddress string, orders []Order) {
	isUnowned := IsNFTUnOwned(mobAddress)
	if isUnowned == true {
		fmt.Printf("contract has no nft, skip to sell")
		return
	}

	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.ValidateSellOrders(txOpts, orders)
	CheckAndWaitTx(tx, err)
}

func IsNFTUnOwned(mobAddress string) bool {
	mob := GetMobByAddress(mobAddress)
	metadata, err := mob.Metadata(nil)
	if err != nil {
	}
	token := metadata.Token.Hex()
	tokenId := metadata.TokenId
	targetMode := metadata.TargetMode
	if targetMode == 0 { // restrict mode
		addr := ERC721OwnerOf(mobAddress, tokenId)
		if addr != mobAddress {
			// check if still holds some other nft id
			bal := ERC721BalanceOf(mobAddress, token)
			if bal.Cmp(big.NewInt(0)) > 0 {
				log.Fatalf("warn: contract holds no nft id, but contains some other nft id, encounter NFT balance attacking?")
			}
		}
		return addr != mobAddress
	}

	if targetMode == 1 { // full open mode
		bal := ERC721BalanceOf(mobAddress, token)
		if bal.Cmp(big.NewInt(0)) > 1 {
			// check if still holds some other nft id
			log.Fatalf("warn: contract holds more than 1 nft id, encounter NFT balance attacking?")
		}
		return bal == big.NewInt(0)
	}

	panic("unknown target mode")
}

func ERC721BalanceOf(owner string, tokenAddress string) *big.Int {
	token, err := erc721.NewErc721(common.HexToAddress(tokenAddress), EthClient)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	bal, err := token.BalanceOf(nil, common.HexToAddress(owner))
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	return bal
}

func ERC721OwnerOf(tokenAddress string, tokenId *big.Int) string {
	token, err := erc721.NewErc721(common.HexToAddress(tokenAddress), EthClient)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	addr, err := token.OwnerOf(nil, tokenId)
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	return addr.Hex()
}

// only for testing, mainnet we don't have to do this
func PatchMobWithSeaport(mobAddress string) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	{
		tx, err := mob.SetSeaportAddress(txOpts, config.GlobalConfig.SeaportAddress)
		CheckAndWaitTx(tx, err)
	}
}

func CheckAndWaitTx(tx *types.Transaction, txErr error) {
	// show caller info
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fmt.Printf("called from %s\n", details.Name())
	}

	// check
	if txErr != nil {
		fmt.Printf("tx failed: %s", txErr.Error())
		return
	}

	fmt.Printf("tx Hash: %s, wait to be mined..", tx.Hash())
	receipt, err := bind.WaitMined(context.Background(), EthClient, tx)
	if err != nil {
		fmt.Printf("get tx receipt failed: %s", err.Error())
		return
	}

	fmt.Printf("tx mined at block %d\n", receipt.BlockNumber.Int64())
}

func NewTxOpts(value *big.Int) *bind.TransactOpts {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)

	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	txOpts.GasPrice = gasPrice

	if value != nil {
		txOpts.Value = value
	}

	return txOpts
}
