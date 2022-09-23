// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AdditionalRecipient is an auto generated low-level Go binding around an user-defined struct.
type AdditionalRecipient struct {
	Amount    *big.Int
	Recipient common.Address
}

// BasicOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type BasicOrderParameters struct {
	ConsiderationToken                common.Address
	ConsiderationIdentifier           *big.Int
	ConsiderationAmount               *big.Int
	Offerer                           common.Address
	Zone                              common.Address
	OfferToken                        common.Address
	OfferIdentifier                   *big.Int
	OfferAmount                       *big.Int
	BasicOrderType                    uint8
	StartTime                         *big.Int
	EndTime                           *big.Int
	ZoneHash                          [32]byte
	Salt                              *big.Int
	OffererConduitKey                 [32]byte
	FulfillerConduitKey               [32]byte
	TotalOriginalAdditionalRecipients *big.Int
	AdditionalRecipients              []AdditionalRecipient
	Signature                         []byte
}

// ConsiderationItem is an auto generated low-level Go binding around an user-defined struct.
type ConsiderationItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
	Recipient            common.Address
}

// OfferItem is an auto generated low-level Go binding around an user-defined struct.
type OfferItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Parameters OrderParameters
	Signature  []byte
}

// OrderParameters is an auto generated low-level Go binding around an user-defined struct.
type OrderParameters struct {
	Offerer                         common.Address
	Zone                            common.Address
	Offer                           []OfferItem
	Consideration                   []ConsiderationItem
	OrderType                       uint8
	StartTime                       *big.Int
	EndTime                         *big.Int
	ZoneHash                        [32]byte
	Salt                            *big.Int
	ConduitKey                      [32]byte
	TotalOriginalConsiderationItems *big.Int
}

// XmobExchangeCoreMetaData contains all meta data concerning the XmobExchangeCore contract.
var XmobExchangeCoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"DepositEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MemberJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"WithdrawStake\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"MAGIC_SIGNATURE\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEAPORT_CORE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"buyBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFulFilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"buyOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFulFilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_raisedTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_takeProfitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stopLossPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_raisedAmountDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_mobName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderHashDigest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"joinPay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"memberDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mobItem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mobName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raisedAmountDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raisedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"registerOrderHashDigest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"registerSellOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaport\",\"type\":\"address\"}],\"name\":\"setSeaportAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"takeTransferFee\",\"type\":\"bool\"}],\"name\":\"settlementAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"settlements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopLossPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takeProfitPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"validateSellOrders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// XmobExchangeCoreABI is the input ABI used to generate the binding from.
// Deprecated: Use XmobExchangeCoreMetaData.ABI instead.
var XmobExchangeCoreABI = XmobExchangeCoreMetaData.ABI

// XmobExchangeCore is an auto generated Go binding around an Ethereum contract.
type XmobExchangeCore struct {
	XmobExchangeCoreCaller     // Read-only binding to the contract
	XmobExchangeCoreTransactor // Write-only binding to the contract
	XmobExchangeCoreFilterer   // Log filterer for contract events
}

// XmobExchangeCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type XmobExchangeCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmobExchangeCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XmobExchangeCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmobExchangeCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XmobExchangeCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmobExchangeCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XmobExchangeCoreSession struct {
	Contract     *XmobExchangeCore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XmobExchangeCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XmobExchangeCoreCallerSession struct {
	Contract *XmobExchangeCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// XmobExchangeCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XmobExchangeCoreTransactorSession struct {
	Contract     *XmobExchangeCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// XmobExchangeCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type XmobExchangeCoreRaw struct {
	Contract *XmobExchangeCore // Generic contract binding to access the raw methods on
}

// XmobExchangeCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XmobExchangeCoreCallerRaw struct {
	Contract *XmobExchangeCoreCaller // Generic read-only contract binding to access the raw methods on
}

// XmobExchangeCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XmobExchangeCoreTransactorRaw struct {
	Contract *XmobExchangeCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXmobExchangeCore creates a new instance of XmobExchangeCore, bound to a specific deployed contract.
func NewXmobExchangeCore(address common.Address, backend bind.ContractBackend) (*XmobExchangeCore, error) {
	contract, err := bindXmobExchangeCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCore{XmobExchangeCoreCaller: XmobExchangeCoreCaller{contract: contract}, XmobExchangeCoreTransactor: XmobExchangeCoreTransactor{contract: contract}, XmobExchangeCoreFilterer: XmobExchangeCoreFilterer{contract: contract}}, nil
}

// NewXmobExchangeCoreCaller creates a new read-only instance of XmobExchangeCore, bound to a specific deployed contract.
func NewXmobExchangeCoreCaller(address common.Address, caller bind.ContractCaller) (*XmobExchangeCoreCaller, error) {
	contract, err := bindXmobExchangeCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreCaller{contract: contract}, nil
}

// NewXmobExchangeCoreTransactor creates a new write-only instance of XmobExchangeCore, bound to a specific deployed contract.
func NewXmobExchangeCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*XmobExchangeCoreTransactor, error) {
	contract, err := bindXmobExchangeCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreTransactor{contract: contract}, nil
}

// NewXmobExchangeCoreFilterer creates a new log filterer instance of XmobExchangeCore, bound to a specific deployed contract.
func NewXmobExchangeCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*XmobExchangeCoreFilterer, error) {
	contract, err := bindXmobExchangeCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreFilterer{contract: contract}, nil
}

// bindXmobExchangeCore binds a generic wrapper to an already deployed contract.
func bindXmobExchangeCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XmobExchangeCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XmobExchangeCore *XmobExchangeCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XmobExchangeCore.Contract.XmobExchangeCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XmobExchangeCore *XmobExchangeCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.XmobExchangeCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XmobExchangeCore *XmobExchangeCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.XmobExchangeCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XmobExchangeCore *XmobExchangeCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XmobExchangeCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XmobExchangeCore *XmobExchangeCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XmobExchangeCore *XmobExchangeCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.contract.Transact(opts, method, params...)
}

// MAGICSIGNATURE is a free data retrieval call binding the contract method 0xa2c8d892.
//
// Solidity: function MAGIC_SIGNATURE() view returns(bytes)
func (_XmobExchangeCore *XmobExchangeCoreCaller) MAGICSIGNATURE(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "MAGIC_SIGNATURE")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MAGICSIGNATURE is a free data retrieval call binding the contract method 0xa2c8d892.
//
// Solidity: function MAGIC_SIGNATURE() view returns(bytes)
func (_XmobExchangeCore *XmobExchangeCoreSession) MAGICSIGNATURE() ([]byte, error) {
	return _XmobExchangeCore.Contract.MAGICSIGNATURE(&_XmobExchangeCore.CallOpts)
}

// MAGICSIGNATURE is a free data retrieval call binding the contract method 0xa2c8d892.
//
// Solidity: function MAGIC_SIGNATURE() view returns(bytes)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) MAGICSIGNATURE() ([]byte, error) {
	return _XmobExchangeCore.Contract.MAGICSIGNATURE(&_XmobExchangeCore.CallOpts)
}

// SEAPORTCORE is a free data retrieval call binding the contract method 0x0a8e841c.
//
// Solidity: function SEAPORT_CORE() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCaller) SEAPORTCORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "SEAPORT_CORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEAPORTCORE is a free data retrieval call binding the contract method 0x0a8e841c.
//
// Solidity: function SEAPORT_CORE() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreSession) SEAPORTCORE() (common.Address, error) {
	return _XmobExchangeCore.Contract.SEAPORTCORE(&_XmobExchangeCore.CallOpts)
}

// SEAPORTCORE is a free data retrieval call binding the contract method 0x0a8e841c.
//
// Solidity: function SEAPORT_CORE() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) SEAPORTCORE() (common.Address, error) {
	return _XmobExchangeCore.Contract.SEAPORTCORE(&_XmobExchangeCore.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_XmobExchangeCore *XmobExchangeCoreCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_XmobExchangeCore *XmobExchangeCoreSession) VERSION() (string, error) {
	return _XmobExchangeCore.Contract.VERSION(&_XmobExchangeCore.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) VERSION() (string, error) {
	return _XmobExchangeCore.Contract.VERSION(&_XmobExchangeCore.CallOpts)
}

// AmountTotal is a free data retrieval call binding the contract method 0xa73c7140.
//
// Solidity: function amountTotal() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) AmountTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "amountTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountTotal is a free data retrieval call binding the contract method 0xa73c7140.
//
// Solidity: function amountTotal() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) AmountTotal() (*big.Int, error) {
	return _XmobExchangeCore.Contract.AmountTotal(&_XmobExchangeCore.CallOpts)
}

// AmountTotal is a free data retrieval call binding the contract method 0xa73c7140.
//
// Solidity: function amountTotal() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) AmountTotal() (*big.Int, error) {
	return _XmobExchangeCore.Contract.AmountTotal(&_XmobExchangeCore.CallOpts)
}

// CanClaim is a free data retrieval call binding the contract method 0x6dc7a627.
//
// Solidity: function canClaim() view returns(bool)
func (_XmobExchangeCore *XmobExchangeCoreCaller) CanClaim(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "canClaim")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0x6dc7a627.
//
// Solidity: function canClaim() view returns(bool)
func (_XmobExchangeCore *XmobExchangeCoreSession) CanClaim() (bool, error) {
	return _XmobExchangeCore.Contract.CanClaim(&_XmobExchangeCore.CallOpts)
}

// CanClaim is a free data retrieval call binding the contract method 0x6dc7a627.
//
// Solidity: function canClaim() view returns(bool)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) CanClaim() (bool, error) {
	return _XmobExchangeCore.Contract.CanClaim(&_XmobExchangeCore.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Cost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "cost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) Cost() (*big.Int, error) {
	return _XmobExchangeCore.Contract.Cost(&_XmobExchangeCore.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Cost() (*big.Int, error) {
	return _XmobExchangeCore.Contract.Cost(&_XmobExchangeCore.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreSession) Creator() (common.Address, error) {
	return _XmobExchangeCore.Contract.Creator(&_XmobExchangeCore.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Creator() (common.Address, error) {
	return _XmobExchangeCore.Contract.Creator(&_XmobExchangeCore.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) Deadline() (*big.Int, error) {
	return _XmobExchangeCore.Contract.Deadline(&_XmobExchangeCore.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Deadline() (*big.Int, error) {
	return _XmobExchangeCore.Contract.Deadline(&_XmobExchangeCore.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) Fee() (*big.Int, error) {
	return _XmobExchangeCore.Contract.Fee(&_XmobExchangeCore.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Fee() (*big.Int, error) {
	return _XmobExchangeCore.Contract.Fee(&_XmobExchangeCore.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _orderHashDigest, bytes _signature) view returns(bytes4)
func (_XmobExchangeCore *XmobExchangeCoreCaller) IsValidSignature(opts *bind.CallOpts, _orderHashDigest [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "isValidSignature", _orderHashDigest, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _orderHashDigest, bytes _signature) view returns(bytes4)
func (_XmobExchangeCore *XmobExchangeCoreSession) IsValidSignature(_orderHashDigest [32]byte, _signature []byte) ([4]byte, error) {
	return _XmobExchangeCore.Contract.IsValidSignature(&_XmobExchangeCore.CallOpts, _orderHashDigest, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _orderHashDigest, bytes _signature) view returns(bytes4)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) IsValidSignature(_orderHashDigest [32]byte, _signature []byte) ([4]byte, error) {
	return _XmobExchangeCore.Contract.IsValidSignature(&_XmobExchangeCore.CallOpts, _orderHashDigest, _signature)
}

// MemberDetails is a free data retrieval call binding the contract method 0x9d3b7cc9.
//
// Solidity: function memberDetails(address ) view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) MemberDetails(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "memberDetails", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemberDetails is a free data retrieval call binding the contract method 0x9d3b7cc9.
//
// Solidity: function memberDetails(address ) view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) MemberDetails(arg0 common.Address) (*big.Int, error) {
	return _XmobExchangeCore.Contract.MemberDetails(&_XmobExchangeCore.CallOpts, arg0)
}

// MemberDetails is a free data retrieval call binding the contract method 0x9d3b7cc9.
//
// Solidity: function memberDetails(address ) view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) MemberDetails(arg0 common.Address) (*big.Int, error) {
	return _XmobExchangeCore.Contract.MemberDetails(&_XmobExchangeCore.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreSession) Members(arg0 *big.Int) (common.Address, error) {
	return _XmobExchangeCore.Contract.Members(&_XmobExchangeCore.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Members(arg0 *big.Int) (common.Address, error) {
	return _XmobExchangeCore.Contract.Members(&_XmobExchangeCore.CallOpts, arg0)
}

// MobItem is a free data retrieval call binding the contract method 0xee4fe97a.
//
// Solidity: function mobItem() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) MobItem(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "mobItem")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// MobItem is a free data retrieval call binding the contract method 0xee4fe97a.
//
// Solidity: function mobItem() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) MobItem() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _XmobExchangeCore.Contract.MobItem(&_XmobExchangeCore.CallOpts)
}

// MobItem is a free data retrieval call binding the contract method 0xee4fe97a.
//
// Solidity: function mobItem() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) MobItem() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _XmobExchangeCore.Contract.MobItem(&_XmobExchangeCore.CallOpts)
}

// MobName is a free data retrieval call binding the contract method 0xc335d87a.
//
// Solidity: function mobName() view returns(string)
func (_XmobExchangeCore *XmobExchangeCoreCaller) MobName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "mobName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MobName is a free data retrieval call binding the contract method 0xc335d87a.
//
// Solidity: function mobName() view returns(string)
func (_XmobExchangeCore *XmobExchangeCoreSession) MobName() (string, error) {
	return _XmobExchangeCore.Contract.MobName(&_XmobExchangeCore.CallOpts)
}

// MobName is a free data retrieval call binding the contract method 0xc335d87a.
//
// Solidity: function mobName() view returns(string)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) MobName() (string, error) {
	return _XmobExchangeCore.Contract.MobName(&_XmobExchangeCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreSession) Owner() (common.Address, error) {
	return _XmobExchangeCore.Contract.Owner(&_XmobExchangeCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Owner() (common.Address, error) {
	return _XmobExchangeCore.Contract.Owner(&_XmobExchangeCore.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_XmobExchangeCore *XmobExchangeCoreCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_XmobExchangeCore *XmobExchangeCoreSession) ProxiableUUID() ([32]byte, error) {
	return _XmobExchangeCore.Contract.ProxiableUUID(&_XmobExchangeCore.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) ProxiableUUID() ([32]byte, error) {
	return _XmobExchangeCore.Contract.ProxiableUUID(&_XmobExchangeCore.CallOpts)
}

// RaisedAmountDeadline is a free data retrieval call binding the contract method 0xd205f1de.
//
// Solidity: function raisedAmountDeadline() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) RaisedAmountDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "raisedAmountDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaisedAmountDeadline is a free data retrieval call binding the contract method 0xd205f1de.
//
// Solidity: function raisedAmountDeadline() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) RaisedAmountDeadline() (*big.Int, error) {
	return _XmobExchangeCore.Contract.RaisedAmountDeadline(&_XmobExchangeCore.CallOpts)
}

// RaisedAmountDeadline is a free data retrieval call binding the contract method 0xd205f1de.
//
// Solidity: function raisedAmountDeadline() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) RaisedAmountDeadline() (*big.Int, error) {
	return _XmobExchangeCore.Contract.RaisedAmountDeadline(&_XmobExchangeCore.CallOpts)
}

// RaisedTotal is a free data retrieval call binding the contract method 0x0afb2dfd.
//
// Solidity: function raisedTotal() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) RaisedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "raisedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaisedTotal is a free data retrieval call binding the contract method 0x0afb2dfd.
//
// Solidity: function raisedTotal() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) RaisedTotal() (*big.Int, error) {
	return _XmobExchangeCore.Contract.RaisedTotal(&_XmobExchangeCore.CallOpts)
}

// RaisedTotal is a free data retrieval call binding the contract method 0x0afb2dfd.
//
// Solidity: function raisedTotal() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) RaisedTotal() (*big.Int, error) {
	return _XmobExchangeCore.Contract.RaisedTotal(&_XmobExchangeCore.CallOpts)
}

// RegisterOrderHashDigest is a free data retrieval call binding the contract method 0xfbae1d18.
//
// Solidity: function registerOrderHashDigest(bytes32 ) view returns(bool)
func (_XmobExchangeCore *XmobExchangeCoreCaller) RegisterOrderHashDigest(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "registerOrderHashDigest", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisterOrderHashDigest is a free data retrieval call binding the contract method 0xfbae1d18.
//
// Solidity: function registerOrderHashDigest(bytes32 ) view returns(bool)
func (_XmobExchangeCore *XmobExchangeCoreSession) RegisterOrderHashDigest(arg0 [32]byte) (bool, error) {
	return _XmobExchangeCore.Contract.RegisterOrderHashDigest(&_XmobExchangeCore.CallOpts, arg0)
}

// RegisterOrderHashDigest is a free data retrieval call binding the contract method 0xfbae1d18.
//
// Solidity: function registerOrderHashDigest(bytes32 ) view returns(bool)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) RegisterOrderHashDigest(arg0 [32]byte) (bool, error) {
	return _XmobExchangeCore.Contract.RegisterOrderHashDigest(&_XmobExchangeCore.CallOpts, arg0)
}

// Settlements is a free data retrieval call binding the contract method 0xba25b1b3.
//
// Solidity: function settlements(address ) view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Settlements(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "settlements", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Settlements is a free data retrieval call binding the contract method 0xba25b1b3.
//
// Solidity: function settlements(address ) view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) Settlements(arg0 common.Address) (*big.Int, error) {
	return _XmobExchangeCore.Contract.Settlements(&_XmobExchangeCore.CallOpts, arg0)
}

// Settlements is a free data retrieval call binding the contract method 0xba25b1b3.
//
// Solidity: function settlements(address ) view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Settlements(arg0 common.Address) (*big.Int, error) {
	return _XmobExchangeCore.Contract.Settlements(&_XmobExchangeCore.CallOpts, arg0)
}

// StopLossPrice is a free data retrieval call binding the contract method 0xe0efb1fd.
//
// Solidity: function stopLossPrice() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) StopLossPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "stopLossPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StopLossPrice is a free data retrieval call binding the contract method 0xe0efb1fd.
//
// Solidity: function stopLossPrice() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) StopLossPrice() (*big.Int, error) {
	return _XmobExchangeCore.Contract.StopLossPrice(&_XmobExchangeCore.CallOpts)
}

// StopLossPrice is a free data retrieval call binding the contract method 0xe0efb1fd.
//
// Solidity: function stopLossPrice() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) StopLossPrice() (*big.Int, error) {
	return _XmobExchangeCore.Contract.StopLossPrice(&_XmobExchangeCore.CallOpts)
}

// TakeProfitPrice is a free data retrieval call binding the contract method 0x0cb479d8.
//
// Solidity: function takeProfitPrice() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCaller) TakeProfitPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "takeProfitPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TakeProfitPrice is a free data retrieval call binding the contract method 0x0cb479d8.
//
// Solidity: function takeProfitPrice() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreSession) TakeProfitPrice() (*big.Int, error) {
	return _XmobExchangeCore.Contract.TakeProfitPrice(&_XmobExchangeCore.CallOpts)
}

// TakeProfitPrice is a free data retrieval call binding the contract method 0x0cb479d8.
//
// Solidity: function takeProfitPrice() view returns(uint256)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) TakeProfitPrice() (*big.Int, error) {
	return _XmobExchangeCore.Contract.TakeProfitPrice(&_XmobExchangeCore.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XmobExchangeCore.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreSession) Token() (common.Address, error) {
	return _XmobExchangeCore.Contract.Token(&_XmobExchangeCore.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_XmobExchangeCore *XmobExchangeCoreCallerSession) Token() (common.Address, error) {
	return _XmobExchangeCore.Contract.Token(&_XmobExchangeCore.CallOpts)
}

// BuyBasicOrder is a paid mutator transaction binding the contract method 0xa9254861.
//
// Solidity: function buyBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool isFulFilled)
func (_XmobExchangeCore *XmobExchangeCoreTransactor) BuyBasicOrder(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "buyBasicOrder", parameters)
}

// BuyBasicOrder is a paid mutator transaction binding the contract method 0xa9254861.
//
// Solidity: function buyBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool isFulFilled)
func (_XmobExchangeCore *XmobExchangeCoreSession) BuyBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.BuyBasicOrder(&_XmobExchangeCore.TransactOpts, parameters)
}

// BuyBasicOrder is a paid mutator transaction binding the contract method 0xa9254861.
//
// Solidity: function buyBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool isFulFilled)
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) BuyBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.BuyBasicOrder(&_XmobExchangeCore.TransactOpts, parameters)
}

// BuyOrder is a paid mutator transaction binding the contract method 0xb1d9646b.
//
// Solidity: function buyOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool isFulFilled)
func (_XmobExchangeCore *XmobExchangeCoreTransactor) BuyOrder(opts *bind.TransactOpts, order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "buyOrder", order, fulfillerConduitKey)
}

// BuyOrder is a paid mutator transaction binding the contract method 0xb1d9646b.
//
// Solidity: function buyOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool isFulFilled)
func (_XmobExchangeCore *XmobExchangeCoreSession) BuyOrder(order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.BuyOrder(&_XmobExchangeCore.TransactOpts, order, fulfillerConduitKey)
}

// BuyOrder is a paid mutator transaction binding the contract method 0xb1d9646b.
//
// Solidity: function buyOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) order, bytes32 fulfillerConduitKey) payable returns(bool isFulFilled)
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) BuyOrder(order Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.BuyOrder(&_XmobExchangeCore.TransactOpts, order, fulfillerConduitKey)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) Claim() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Claim(&_XmobExchangeCore.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) Claim() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Claim(&_XmobExchangeCore.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x84b77472.
//
// Solidity: function initialize(address _creator, address _token, uint256 _fee, uint256 _raisedTotal, uint256 _takeProfitPrice, uint256 _stopLossPrice, uint256 _raisedAmountDeadline, uint256 _deadline, string _mobName) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) Initialize(opts *bind.TransactOpts, _creator common.Address, _token common.Address, _fee *big.Int, _raisedTotal *big.Int, _takeProfitPrice *big.Int, _stopLossPrice *big.Int, _raisedAmountDeadline *big.Int, _deadline *big.Int, _mobName string) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "initialize", _creator, _token, _fee, _raisedTotal, _takeProfitPrice, _stopLossPrice, _raisedAmountDeadline, _deadline, _mobName)
}

// Initialize is a paid mutator transaction binding the contract method 0x84b77472.
//
// Solidity: function initialize(address _creator, address _token, uint256 _fee, uint256 _raisedTotal, uint256 _takeProfitPrice, uint256 _stopLossPrice, uint256 _raisedAmountDeadline, uint256 _deadline, string _mobName) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) Initialize(_creator common.Address, _token common.Address, _fee *big.Int, _raisedTotal *big.Int, _takeProfitPrice *big.Int, _stopLossPrice *big.Int, _raisedAmountDeadline *big.Int, _deadline *big.Int, _mobName string) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Initialize(&_XmobExchangeCore.TransactOpts, _creator, _token, _fee, _raisedTotal, _takeProfitPrice, _stopLossPrice, _raisedAmountDeadline, _deadline, _mobName)
}

// Initialize is a paid mutator transaction binding the contract method 0x84b77472.
//
// Solidity: function initialize(address _creator, address _token, uint256 _fee, uint256 _raisedTotal, uint256 _takeProfitPrice, uint256 _stopLossPrice, uint256 _raisedAmountDeadline, uint256 _deadline, string _mobName) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) Initialize(_creator common.Address, _token common.Address, _fee *big.Int, _raisedTotal *big.Int, _takeProfitPrice *big.Int, _stopLossPrice *big.Int, _raisedAmountDeadline *big.Int, _deadline *big.Int, _mobName string) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Initialize(&_XmobExchangeCore.TransactOpts, _creator, _token, _fee, _raisedTotal, _takeProfitPrice, _stopLossPrice, _raisedAmountDeadline, _deadline, _mobName)
}

// JoinPay is a paid mutator transaction binding the contract method 0x69369fa1.
//
// Solidity: function joinPay(address member) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) JoinPay(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "joinPay", member)
}

// JoinPay is a paid mutator transaction binding the contract method 0x69369fa1.
//
// Solidity: function joinPay(address member) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) JoinPay(member common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.JoinPay(&_XmobExchangeCore.TransactOpts, member)
}

// JoinPay is a paid mutator transaction binding the contract method 0x69369fa1.
//
// Solidity: function joinPay(address member) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) JoinPay(member common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.JoinPay(&_XmobExchangeCore.TransactOpts, member)
}

// RegisterSellOrder is a paid mutator transaction binding the contract method 0x8073dfb6.
//
// Solidity: function registerSellOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) RegisterSellOrder(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "registerSellOrder", orders)
}

// RegisterSellOrder is a paid mutator transaction binding the contract method 0x8073dfb6.
//
// Solidity: function registerSellOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) RegisterSellOrder(orders []Order) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.RegisterSellOrder(&_XmobExchangeCore.TransactOpts, orders)
}

// RegisterSellOrder is a paid mutator transaction binding the contract method 0x8073dfb6.
//
// Solidity: function registerSellOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) RegisterSellOrder(orders []Order) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.RegisterSellOrder(&_XmobExchangeCore.TransactOpts, orders)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.RenounceOwnership(&_XmobExchangeCore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.RenounceOwnership(&_XmobExchangeCore.TransactOpts)
}

// SetSeaportAddress is a paid mutator transaction binding the contract method 0xe67ece91.
//
// Solidity: function setSeaportAddress(address seaport) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) SetSeaportAddress(opts *bind.TransactOpts, seaport common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "setSeaportAddress", seaport)
}

// SetSeaportAddress is a paid mutator transaction binding the contract method 0xe67ece91.
//
// Solidity: function setSeaportAddress(address seaport) returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) SetSeaportAddress(seaport common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.SetSeaportAddress(&_XmobExchangeCore.TransactOpts, seaport)
}

// SetSeaportAddress is a paid mutator transaction binding the contract method 0xe67ece91.
//
// Solidity: function setSeaportAddress(address seaport) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) SetSeaportAddress(seaport common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.SetSeaportAddress(&_XmobExchangeCore.TransactOpts, seaport)
}

// SettlementAllocation is a paid mutator transaction binding the contract method 0xea409f88.
//
// Solidity: function settlementAllocation(bool takeTransferFee) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) SettlementAllocation(opts *bind.TransactOpts, takeTransferFee bool) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "settlementAllocation", takeTransferFee)
}

// SettlementAllocation is a paid mutator transaction binding the contract method 0xea409f88.
//
// Solidity: function settlementAllocation(bool takeTransferFee) returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) SettlementAllocation(takeTransferFee bool) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.SettlementAllocation(&_XmobExchangeCore.TransactOpts, takeTransferFee)
}

// SettlementAllocation is a paid mutator transaction binding the contract method 0xea409f88.
//
// Solidity: function settlementAllocation(bool takeTransferFee) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) SettlementAllocation(takeTransferFee bool) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.SettlementAllocation(&_XmobExchangeCore.TransactOpts, takeTransferFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.TransferOwnership(&_XmobExchangeCore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.TransferOwnership(&_XmobExchangeCore.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.UpgradeTo(&_XmobExchangeCore.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.UpgradeTo(&_XmobExchangeCore.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.UpgradeToAndCall(&_XmobExchangeCore.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.UpgradeToAndCall(&_XmobExchangeCore.TransactOpts, newImplementation, data)
}

// ValidateSellOrders is a paid mutator transaction binding the contract method 0xcbc9544c.
//
// Solidity: function validateSellOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool isValidated)
func (_XmobExchangeCore *XmobExchangeCoreTransactor) ValidateSellOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "validateSellOrders", orders)
}

// ValidateSellOrders is a paid mutator transaction binding the contract method 0xcbc9544c.
//
// Solidity: function validateSellOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool isValidated)
func (_XmobExchangeCore *XmobExchangeCoreSession) ValidateSellOrders(orders []Order) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.ValidateSellOrders(&_XmobExchangeCore.TransactOpts, orders)
}

// ValidateSellOrders is a paid mutator transaction binding the contract method 0xcbc9544c.
//
// Solidity: function validateSellOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] orders) returns(bool isValidated)
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) ValidateSellOrders(orders []Order) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.ValidateSellOrders(&_XmobExchangeCore.TransactOpts, orders)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) WithdrawStake() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.WithdrawStake(&_XmobExchangeCore.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.WithdrawStake(&_XmobExchangeCore.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Fallback(&_XmobExchangeCore.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Fallback(&_XmobExchangeCore.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobExchangeCore.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_XmobExchangeCore *XmobExchangeCoreSession) Receive() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Receive(&_XmobExchangeCore.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_XmobExchangeCore *XmobExchangeCoreTransactorSession) Receive() (*types.Transaction, error) {
	return _XmobExchangeCore.Contract.Receive(&_XmobExchangeCore.TransactOpts)
}

// XmobExchangeCoreAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the XmobExchangeCore contract.
type XmobExchangeCoreAdminChangedIterator struct {
	Event *XmobExchangeCoreAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreAdminChanged represents a AdminChanged event raised by the XmobExchangeCore contract.
type XmobExchangeCoreAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*XmobExchangeCoreAdminChangedIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreAdminChangedIterator{contract: _XmobExchangeCore.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreAdminChanged) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreAdminChanged)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseAdminChanged(log types.Log) (*XmobExchangeCoreAdminChanged, error) {
	event := new(XmobExchangeCoreAdminChanged)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the XmobExchangeCore contract.
type XmobExchangeCoreBeaconUpgradedIterator struct {
	Event *XmobExchangeCoreBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreBeaconUpgraded represents a BeaconUpgraded event raised by the XmobExchangeCore contract.
type XmobExchangeCoreBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*XmobExchangeCoreBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreBeaconUpgradedIterator{contract: _XmobExchangeCore.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreBeaconUpgraded)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseBeaconUpgraded(log types.Log) (*XmobExchangeCoreBeaconUpgraded, error) {
	event := new(XmobExchangeCoreBeaconUpgraded)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the XmobExchangeCore contract.
type XmobExchangeCoreBuyIterator struct {
	Event *XmobExchangeCoreBuy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreBuy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreBuy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreBuy represents a Buy event raised by the XmobExchangeCore contract.
type XmobExchangeCoreBuy struct {
	Seller common.Address
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e.
//
// Solidity: event Buy(address indexed seller, uint256 price)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterBuy(opts *bind.FilterOpts, seller []common.Address) (*XmobExchangeCoreBuyIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "Buy", sellerRule)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreBuyIterator{contract: _XmobExchangeCore.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e.
//
// Solidity: event Buy(address indexed seller, uint256 price)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreBuy, seller []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "Buy", sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreBuy)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "Buy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuy is a log parse operation binding the contract event 0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e.
//
// Solidity: event Buy(address indexed seller, uint256 price)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseBuy(log types.Log) (*XmobExchangeCoreBuy, error) {
	event := new(XmobExchangeCoreBuy)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the XmobExchangeCore contract.
type XmobExchangeCoreClaimIterator struct {
	Event *XmobExchangeCoreClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreClaim represents a Claim event raised by the XmobExchangeCore contract.
type XmobExchangeCoreClaim struct {
	Member common.Address
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address member, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterClaim(opts *bind.FilterOpts) (*XmobExchangeCoreClaimIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreClaimIterator{contract: _XmobExchangeCore.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address member, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreClaim) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreClaim)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address member, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseClaim(log types.Log) (*XmobExchangeCoreClaim, error) {
	event := new(XmobExchangeCoreClaim)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreDepositEthIterator is returned from FilterDepositEth and is used to iterate over the raw logs and unpacked data for DepositEth events raised by the XmobExchangeCore contract.
type XmobExchangeCoreDepositEthIterator struct {
	Event *XmobExchangeCoreDepositEth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreDepositEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreDepositEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreDepositEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreDepositEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreDepositEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreDepositEth represents a DepositEth event raised by the XmobExchangeCore contract.
type XmobExchangeCoreDepositEth struct {
	Sender common.Address
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositEth is a free log retrieval operation binding the contract event 0x7034bb05cfe54b0d147fc0574ed166101e7f0313eb404e113974fbe2a998ca83.
//
// Solidity: event DepositEth(address sender, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterDepositEth(opts *bind.FilterOpts) (*XmobExchangeCoreDepositEthIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "DepositEth")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreDepositEthIterator{contract: _XmobExchangeCore.contract, event: "DepositEth", logs: logs, sub: sub}, nil
}

// WatchDepositEth is a free log subscription operation binding the contract event 0x7034bb05cfe54b0d147fc0574ed166101e7f0313eb404e113974fbe2a998ca83.
//
// Solidity: event DepositEth(address sender, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchDepositEth(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreDepositEth) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "DepositEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreDepositEth)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "DepositEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositEth is a log parse operation binding the contract event 0x7034bb05cfe54b0d147fc0574ed166101e7f0313eb404e113974fbe2a998ca83.
//
// Solidity: event DepositEth(address sender, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseDepositEth(log types.Log) (*XmobExchangeCoreDepositEth, error) {
	event := new(XmobExchangeCoreDepositEth)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "DepositEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the XmobExchangeCore contract.
type XmobExchangeCoreInitializedIterator struct {
	Event *XmobExchangeCoreInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreInitialized represents a Initialized event raised by the XmobExchangeCore contract.
type XmobExchangeCoreInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterInitialized(opts *bind.FilterOpts) (*XmobExchangeCoreInitializedIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreInitializedIterator{contract: _XmobExchangeCore.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreInitialized) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreInitialized)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseInitialized(log types.Log) (*XmobExchangeCoreInitialized, error) {
	event := new(XmobExchangeCoreInitialized)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreMemberJoinIterator is returned from FilterMemberJoin and is used to iterate over the raw logs and unpacked data for MemberJoin events raised by the XmobExchangeCore contract.
type XmobExchangeCoreMemberJoinIterator struct {
	Event *XmobExchangeCoreMemberJoin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreMemberJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreMemberJoin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreMemberJoin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreMemberJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreMemberJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreMemberJoin represents a MemberJoin event raised by the XmobExchangeCore contract.
type XmobExchangeCoreMemberJoin struct {
	Member common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberJoin is a free log retrieval operation binding the contract event 0x6caf6186588a289956ca08d23a7cbfe2ca473d5969db8117a4cbbe45a96674d1.
//
// Solidity: event MemberJoin(address member, uint256 value)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterMemberJoin(opts *bind.FilterOpts) (*XmobExchangeCoreMemberJoinIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "MemberJoin")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreMemberJoinIterator{contract: _XmobExchangeCore.contract, event: "MemberJoin", logs: logs, sub: sub}, nil
}

// WatchMemberJoin is a free log subscription operation binding the contract event 0x6caf6186588a289956ca08d23a7cbfe2ca473d5969db8117a4cbbe45a96674d1.
//
// Solidity: event MemberJoin(address member, uint256 value)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchMemberJoin(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreMemberJoin) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "MemberJoin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreMemberJoin)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "MemberJoin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMemberJoin is a log parse operation binding the contract event 0x6caf6186588a289956ca08d23a7cbfe2ca473d5969db8117a4cbbe45a96674d1.
//
// Solidity: event MemberJoin(address member, uint256 value)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseMemberJoin(log types.Log) (*XmobExchangeCoreMemberJoin, error) {
	event := new(XmobExchangeCoreMemberJoin)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "MemberJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the XmobExchangeCore contract.
type XmobExchangeCoreOwnershipTransferredIterator struct {
	Event *XmobExchangeCoreOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreOwnershipTransferred represents a OwnershipTransferred event raised by the XmobExchangeCore contract.
type XmobExchangeCoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*XmobExchangeCoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreOwnershipTransferredIterator{contract: _XmobExchangeCore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreOwnershipTransferred)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseOwnershipTransferred(log types.Log) (*XmobExchangeCoreOwnershipTransferred, error) {
	event := new(XmobExchangeCoreOwnershipTransferred)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the XmobExchangeCore contract.
type XmobExchangeCoreSettlementIterator struct {
	Event *XmobExchangeCoreSettlement // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreSettlement)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreSettlement)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreSettlement represents a Settlement event raised by the XmobExchangeCore contract.
type XmobExchangeCoreSettlement struct {
	Total *big.Int
	Time  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x3b9add234c8c7540648088b823e43c595a23ebe2c2ce32edd5cca48350f6fa97.
//
// Solidity: event Settlement(uint256 total, uint256 time)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterSettlement(opts *bind.FilterOpts) (*XmobExchangeCoreSettlementIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "Settlement")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreSettlementIterator{contract: _XmobExchangeCore.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x3b9add234c8c7540648088b823e43c595a23ebe2c2ce32edd5cca48350f6fa97.
//
// Solidity: event Settlement(uint256 total, uint256 time)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreSettlement) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "Settlement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreSettlement)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "Settlement", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlement is a log parse operation binding the contract event 0x3b9add234c8c7540648088b823e43c595a23ebe2c2ce32edd5cca48350f6fa97.
//
// Solidity: event Settlement(uint256 total, uint256 time)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseSettlement(log types.Log) (*XmobExchangeCoreSettlement, error) {
	event := new(XmobExchangeCoreSettlement)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the XmobExchangeCore contract.
type XmobExchangeCoreUpgradedIterator struct {
	Event *XmobExchangeCoreUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreUpgraded represents a Upgraded event raised by the XmobExchangeCore contract.
type XmobExchangeCoreUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*XmobExchangeCoreUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreUpgradedIterator{contract: _XmobExchangeCore.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreUpgraded)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseUpgraded(log types.Log) (*XmobExchangeCoreUpgraded, error) {
	event := new(XmobExchangeCoreUpgraded)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobExchangeCoreWithdrawStakeIterator is returned from FilterWithdrawStake and is used to iterate over the raw logs and unpacked data for WithdrawStake events raised by the XmobExchangeCore contract.
type XmobExchangeCoreWithdrawStakeIterator struct {
	Event *XmobExchangeCoreWithdrawStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XmobExchangeCoreWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobExchangeCoreWithdrawStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XmobExchangeCoreWithdrawStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XmobExchangeCoreWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobExchangeCoreWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobExchangeCoreWithdrawStake represents a WithdrawStake event raised by the XmobExchangeCore contract.
type XmobExchangeCoreWithdrawStake struct {
	Member common.Address
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawStake is a free log retrieval operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address member, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) FilterWithdrawStake(opts *bind.FilterOpts) (*XmobExchangeCoreWithdrawStakeIterator, error) {

	logs, sub, err := _XmobExchangeCore.contract.FilterLogs(opts, "WithdrawStake")
	if err != nil {
		return nil, err
	}
	return &XmobExchangeCoreWithdrawStakeIterator{contract: _XmobExchangeCore.contract, event: "WithdrawStake", logs: logs, sub: sub}, nil
}

// WatchWithdrawStake is a free log subscription operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address member, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) WatchWithdrawStake(opts *bind.WatchOpts, sink chan<- *XmobExchangeCoreWithdrawStake) (event.Subscription, error) {

	logs, sub, err := _XmobExchangeCore.contract.WatchLogs(opts, "WithdrawStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobExchangeCoreWithdrawStake)
				if err := _XmobExchangeCore.contract.UnpackLog(event, "WithdrawStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawStake is a log parse operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address member, uint256 amt)
func (_XmobExchangeCore *XmobExchangeCoreFilterer) ParseWithdrawStake(log types.Log) (*XmobExchangeCoreWithdrawStake, error) {
	event := new(XmobExchangeCoreWithdrawStake)
	if err := _XmobExchangeCore.contract.UnpackLog(event, "WithdrawStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
