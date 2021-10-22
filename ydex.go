// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ydexabi

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

// StructsAsset is an auto generated low-level Go binding around an user-defined struct.
type StructsAsset struct {
	Exists                 bool
	AssetAddress           common.Address
	Symbol                 string
	Decimals               uint8
	IsConfirmed            bool
	ConfirmedTimestampInMs uint64
}

// StructsOrder is an auto generated low-level Go binding around an user-defined struct.
type StructsOrder struct {
	SignatureHashVersion uint8
	Nonce                *big.Int
	WalletAddress        common.Address
	OrderType            uint8
	Side                 uint8
	QuantityInPips       uint64
	IsQuantityInQuote    bool
	LimitPriceInPips     uint64
	StopPriceInPips      uint64
	ClientOrderId        string
	TimeInForce          uint8
	SelfTradePrevention  uint8
	CancelAfter          uint64
	WalletSignature      []byte
}

// StructsTrade is an auto generated low-level Go binding around an user-defined struct.
type StructsTrade struct {
	BaseAssetSymbol          string
	QuoteAssetSymbol         string
	BaseAssetAddress         common.Address
	QuoteAssetAddress        common.Address
	GrossBaseQuantityInPips  uint64
	GrossQuoteQuantityInPips uint64
	NetBaseQuantityInPips    uint64
	NetQuoteQuantityInPips   uint64
	MakerFeeAssetAddress     common.Address
	TakerFeeAssetAddress     common.Address
	MakerFeeQuantityInPips   uint64
	TakerFeeQuantityInPips   uint64
	PriceInPips              uint64
	MakerSide                uint8
}

// StructsWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type StructsWithdrawal struct {
	WithdrawalType      uint8
	Nonce               *big.Int
	WalletAddress       common.Address
	AssetSymbol         string
	AssetAddress        common.Address
	QuantityInPips      uint64
	GasFeeInPips        uint64
	AutoDispatchEnabled bool
	WalletSignature     []byte
}

// YdexabiMetaData contains all meta data concerning the Ydexabi contract.
var YdexabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChainPropagationPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"assetSymbolIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"DispatcherChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"FeeWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestampInMs\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlockNumber\",\"type\":\"uint256\"}],\"name\":\"OrderNonceInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"}],\"name\":\"TokenSymbolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"baseAssetSymbolIndex\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"quoteAssetSymbolIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseAssetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"quoteAssetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"baseQuantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quoteQuantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"tradePriceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyOrderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellOrderHash\",\"type\":\"bytes32\"}],\"name\":\"TradeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"WalletExitCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"WalletExitWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlockNumber\",\"type\":\"uint256\"}],\"name\":\"WalletExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addTokenSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearWalletExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"confirmTokenRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantityInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"depositTokenByAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantityInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"depositTokenBySymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"signatureHashVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumEnums.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isQuantityInQuote\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"limitPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stopPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientOrderId\",\"type\":\"string\"},{\"internalType\":\"enumEnums.OrderTimeInForce\",\"name\":\"timeInForce\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSelfTradePrevention\",\"name\":\"selfTradePrevention\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"cancelAfter\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structStructs.Order\",\"name\":\"buy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"signatureHashVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumEnums.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isQuantityInQuote\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"limitPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stopPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientOrderId\",\"type\":\"string\"},{\"internalType\":\"enumEnums.OrderTimeInForce\",\"name\":\"timeInForce\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSelfTradePrevention\",\"name\":\"selfTradePrevention\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"cancelAfter\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structStructs.Order\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"baseAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quoteAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"baseAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"grossBaseQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"grossQuoteQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"netBaseQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"netQuoteQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"makerFeeAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerFeeAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"makerFeeQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"takerFeeQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"priceInPips\",\"type\":\"uint64\"},{\"internalType\":\"enumEnums.OrderSide\",\"name\":\"makerSide\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Trade\",\"name\":\"trade\",\"type\":\"tuple\"}],\"name\":\"executeTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"}],\"name\":\"invalidateOrderNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"isWalletExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"isWalletExitFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestampInMs\",\"type\":\"uint64\"}],\"name\":\"loadAssetBySymbol\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isConfirmed\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"confirmedTimestampInMs\",\"type\":\"uint64\"}],\"internalType\":\"structStructs.Asset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"loadBalanceAssetAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"loadBalanceInAssetUnitsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"}],\"name\":\"loadBalanceInAssetUnitsBySymbol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"loadBalanceInPipsByAddress\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"}],\"name\":\"loadBalanceInPipsBySymbol\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadFeeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"loadPartiallyFilledOrderQuantityInPips\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainPropagationPeriod\",\"type\":\"uint256\"}],\"name\":\"setChainPropagationPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCustodian\",\"type\":\"address\"}],\"name\":\"setCustodian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDispatcherWallet\",\"type\":\"address\"}],\"name\":\"setDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeWallet\",\"type\":\"address\"}],\"name\":\"setFeeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumEnums.WithdrawalType\",\"name\":\"withdrawalType\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasFeeInPips\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"autoDispatchEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structStructs.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"withdrawExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YdexabiABI is the input ABI used to generate the binding from.
// Deprecated: Use YdexabiMetaData.ABI instead.
var YdexabiABI = YdexabiMetaData.ABI

// Ydexabi is an auto generated Go binding around an Ethereum contract.
type Ydexabi struct {
	YdexabiCaller     // Read-only binding to the contract
	YdexabiTransactor // Write-only binding to the contract
	YdexabiFilterer   // Log filterer for contract events
}

// YdexabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type YdexabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YdexabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YdexabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YdexabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YdexabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YdexabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YdexabiSession struct {
	Contract     *Ydexabi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YdexabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YdexabiCallerSession struct {
	Contract *YdexabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// YdexabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YdexabiTransactorSession struct {
	Contract     *YdexabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// YdexabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type YdexabiRaw struct {
	Contract *Ydexabi // Generic contract binding to access the raw methods on
}

// YdexabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YdexabiCallerRaw struct {
	Contract *YdexabiCaller // Generic read-only contract binding to access the raw methods on
}

// YdexabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YdexabiTransactorRaw struct {
	Contract *YdexabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYdexabi creates a new instance of Ydexabi, bound to a specific deployed contract.
func NewYdexabi(address common.Address, backend bind.ContractBackend) (*Ydexabi, error) {
	contract, err := bindYdexabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ydexabi{YdexabiCaller: YdexabiCaller{contract: contract}, YdexabiTransactor: YdexabiTransactor{contract: contract}, YdexabiFilterer: YdexabiFilterer{contract: contract}}, nil
}

// NewYdexabiCaller creates a new read-only instance of Ydexabi, bound to a specific deployed contract.
func NewYdexabiCaller(address common.Address, caller bind.ContractCaller) (*YdexabiCaller, error) {
	contract, err := bindYdexabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YdexabiCaller{contract: contract}, nil
}

// NewYdexabiTransactor creates a new write-only instance of Ydexabi, bound to a specific deployed contract.
func NewYdexabiTransactor(address common.Address, transactor bind.ContractTransactor) (*YdexabiTransactor, error) {
	contract, err := bindYdexabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YdexabiTransactor{contract: contract}, nil
}

// NewYdexabiFilterer creates a new log filterer instance of Ydexabi, bound to a specific deployed contract.
func NewYdexabiFilterer(address common.Address, filterer bind.ContractFilterer) (*YdexabiFilterer, error) {
	contract, err := bindYdexabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YdexabiFilterer{contract: contract}, nil
}

// bindYdexabi binds a generic wrapper to an already deployed contract.
func bindYdexabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YdexabiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ydexabi *YdexabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ydexabi.Contract.YdexabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ydexabi *YdexabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.Contract.YdexabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ydexabi *YdexabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ydexabi.Contract.YdexabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ydexabi *YdexabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ydexabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ydexabi *YdexabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ydexabi *YdexabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ydexabi.Contract.contract.Transact(opts, method, params...)
}

// IsWalletExit is a free data retrieval call binding the contract method 0xdfa51870.
//
// Solidity: function isWalletExit(address wallet) view returns(bool)
func (_Ydexabi *YdexabiCaller) IsWalletExit(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "isWalletExit", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletExit is a free data retrieval call binding the contract method 0xdfa51870.
//
// Solidity: function isWalletExit(address wallet) view returns(bool)
func (_Ydexabi *YdexabiSession) IsWalletExit(wallet common.Address) (bool, error) {
	return _Ydexabi.Contract.IsWalletExit(&_Ydexabi.CallOpts, wallet)
}

// IsWalletExit is a free data retrieval call binding the contract method 0xdfa51870.
//
// Solidity: function isWalletExit(address wallet) view returns(bool)
func (_Ydexabi *YdexabiCallerSession) IsWalletExit(wallet common.Address) (bool, error) {
	return _Ydexabi.Contract.IsWalletExit(&_Ydexabi.CallOpts, wallet)
}

// IsWalletExitFinalized is a free data retrieval call binding the contract method 0x80d0dee7.
//
// Solidity: function isWalletExitFinalized(address wallet) view returns(bool)
func (_Ydexabi *YdexabiCaller) IsWalletExitFinalized(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "isWalletExitFinalized", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletExitFinalized is a free data retrieval call binding the contract method 0x80d0dee7.
//
// Solidity: function isWalletExitFinalized(address wallet) view returns(bool)
func (_Ydexabi *YdexabiSession) IsWalletExitFinalized(wallet common.Address) (bool, error) {
	return _Ydexabi.Contract.IsWalletExitFinalized(&_Ydexabi.CallOpts, wallet)
}

// IsWalletExitFinalized is a free data retrieval call binding the contract method 0x80d0dee7.
//
// Solidity: function isWalletExitFinalized(address wallet) view returns(bool)
func (_Ydexabi *YdexabiCallerSession) IsWalletExitFinalized(wallet common.Address) (bool, error) {
	return _Ydexabi.Contract.IsWalletExitFinalized(&_Ydexabi.CallOpts, wallet)
}

// LoadAssetBySymbol is a free data retrieval call binding the contract method 0x869af212.
//
// Solidity: function loadAssetBySymbol(string assetSymbol, uint64 timestampInMs) view returns((bool,address,string,uint8,bool,uint64))
func (_Ydexabi *YdexabiCaller) LoadAssetBySymbol(opts *bind.CallOpts, assetSymbol string, timestampInMs uint64) (StructsAsset, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadAssetBySymbol", assetSymbol, timestampInMs)

	if err != nil {
		return *new(StructsAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsAsset)).(*StructsAsset)

	return out0, err

}

// LoadAssetBySymbol is a free data retrieval call binding the contract method 0x869af212.
//
// Solidity: function loadAssetBySymbol(string assetSymbol, uint64 timestampInMs) view returns((bool,address,string,uint8,bool,uint64))
func (_Ydexabi *YdexabiSession) LoadAssetBySymbol(assetSymbol string, timestampInMs uint64) (StructsAsset, error) {
	return _Ydexabi.Contract.LoadAssetBySymbol(&_Ydexabi.CallOpts, assetSymbol, timestampInMs)
}

// LoadAssetBySymbol is a free data retrieval call binding the contract method 0x869af212.
//
// Solidity: function loadAssetBySymbol(string assetSymbol, uint64 timestampInMs) view returns((bool,address,string,uint8,bool,uint64))
func (_Ydexabi *YdexabiCallerSession) LoadAssetBySymbol(assetSymbol string, timestampInMs uint64) (StructsAsset, error) {
	return _Ydexabi.Contract.LoadAssetBySymbol(&_Ydexabi.CallOpts, assetSymbol, timestampInMs)
}

// LoadBalanceAssetAddress is a free data retrieval call binding the contract method 0x91ec8c8b.
//
// Solidity: function loadBalanceAssetAddress(address wallet) view returns(address[])
func (_Ydexabi *YdexabiCaller) LoadBalanceAssetAddress(opts *bind.CallOpts, wallet common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadBalanceAssetAddress", wallet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LoadBalanceAssetAddress is a free data retrieval call binding the contract method 0x91ec8c8b.
//
// Solidity: function loadBalanceAssetAddress(address wallet) view returns(address[])
func (_Ydexabi *YdexabiSession) LoadBalanceAssetAddress(wallet common.Address) ([]common.Address, error) {
	return _Ydexabi.Contract.LoadBalanceAssetAddress(&_Ydexabi.CallOpts, wallet)
}

// LoadBalanceAssetAddress is a free data retrieval call binding the contract method 0x91ec8c8b.
//
// Solidity: function loadBalanceAssetAddress(address wallet) view returns(address[])
func (_Ydexabi *YdexabiCallerSession) LoadBalanceAssetAddress(wallet common.Address) ([]common.Address, error) {
	return _Ydexabi.Contract.LoadBalanceAssetAddress(&_Ydexabi.CallOpts, wallet)
}

// LoadBalanceInAssetUnitsByAddress is a free data retrieval call binding the contract method 0x13cfda2c.
//
// Solidity: function loadBalanceInAssetUnitsByAddress(address wallet, address assetAddress) view returns(uint256)
func (_Ydexabi *YdexabiCaller) LoadBalanceInAssetUnitsByAddress(opts *bind.CallOpts, wallet common.Address, assetAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadBalanceInAssetUnitsByAddress", wallet, assetAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoadBalanceInAssetUnitsByAddress is a free data retrieval call binding the contract method 0x13cfda2c.
//
// Solidity: function loadBalanceInAssetUnitsByAddress(address wallet, address assetAddress) view returns(uint256)
func (_Ydexabi *YdexabiSession) LoadBalanceInAssetUnitsByAddress(wallet common.Address, assetAddress common.Address) (*big.Int, error) {
	return _Ydexabi.Contract.LoadBalanceInAssetUnitsByAddress(&_Ydexabi.CallOpts, wallet, assetAddress)
}

// LoadBalanceInAssetUnitsByAddress is a free data retrieval call binding the contract method 0x13cfda2c.
//
// Solidity: function loadBalanceInAssetUnitsByAddress(address wallet, address assetAddress) view returns(uint256)
func (_Ydexabi *YdexabiCallerSession) LoadBalanceInAssetUnitsByAddress(wallet common.Address, assetAddress common.Address) (*big.Int, error) {
	return _Ydexabi.Contract.LoadBalanceInAssetUnitsByAddress(&_Ydexabi.CallOpts, wallet, assetAddress)
}

// LoadBalanceInAssetUnitsBySymbol is a free data retrieval call binding the contract method 0x1abb5832.
//
// Solidity: function loadBalanceInAssetUnitsBySymbol(address wallet, string assetSymbol) view returns(uint256)
func (_Ydexabi *YdexabiCaller) LoadBalanceInAssetUnitsBySymbol(opts *bind.CallOpts, wallet common.Address, assetSymbol string) (*big.Int, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadBalanceInAssetUnitsBySymbol", wallet, assetSymbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoadBalanceInAssetUnitsBySymbol is a free data retrieval call binding the contract method 0x1abb5832.
//
// Solidity: function loadBalanceInAssetUnitsBySymbol(address wallet, string assetSymbol) view returns(uint256)
func (_Ydexabi *YdexabiSession) LoadBalanceInAssetUnitsBySymbol(wallet common.Address, assetSymbol string) (*big.Int, error) {
	return _Ydexabi.Contract.LoadBalanceInAssetUnitsBySymbol(&_Ydexabi.CallOpts, wallet, assetSymbol)
}

// LoadBalanceInAssetUnitsBySymbol is a free data retrieval call binding the contract method 0x1abb5832.
//
// Solidity: function loadBalanceInAssetUnitsBySymbol(address wallet, string assetSymbol) view returns(uint256)
func (_Ydexabi *YdexabiCallerSession) LoadBalanceInAssetUnitsBySymbol(wallet common.Address, assetSymbol string) (*big.Int, error) {
	return _Ydexabi.Contract.LoadBalanceInAssetUnitsBySymbol(&_Ydexabi.CallOpts, wallet, assetSymbol)
}

// LoadBalanceInPipsByAddress is a free data retrieval call binding the contract method 0xdbb36535.
//
// Solidity: function loadBalanceInPipsByAddress(address wallet, address assetAddress) view returns(uint64)
func (_Ydexabi *YdexabiCaller) LoadBalanceInPipsByAddress(opts *bind.CallOpts, wallet common.Address, assetAddress common.Address) (uint64, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadBalanceInPipsByAddress", wallet, assetAddress)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LoadBalanceInPipsByAddress is a free data retrieval call binding the contract method 0xdbb36535.
//
// Solidity: function loadBalanceInPipsByAddress(address wallet, address assetAddress) view returns(uint64)
func (_Ydexabi *YdexabiSession) LoadBalanceInPipsByAddress(wallet common.Address, assetAddress common.Address) (uint64, error) {
	return _Ydexabi.Contract.LoadBalanceInPipsByAddress(&_Ydexabi.CallOpts, wallet, assetAddress)
}

// LoadBalanceInPipsByAddress is a free data retrieval call binding the contract method 0xdbb36535.
//
// Solidity: function loadBalanceInPipsByAddress(address wallet, address assetAddress) view returns(uint64)
func (_Ydexabi *YdexabiCallerSession) LoadBalanceInPipsByAddress(wallet common.Address, assetAddress common.Address) (uint64, error) {
	return _Ydexabi.Contract.LoadBalanceInPipsByAddress(&_Ydexabi.CallOpts, wallet, assetAddress)
}

// LoadBalanceInPipsBySymbol is a free data retrieval call binding the contract method 0xef3b9d4a.
//
// Solidity: function loadBalanceInPipsBySymbol(address wallet, string assetSymbol) view returns(uint64)
func (_Ydexabi *YdexabiCaller) LoadBalanceInPipsBySymbol(opts *bind.CallOpts, wallet common.Address, assetSymbol string) (uint64, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadBalanceInPipsBySymbol", wallet, assetSymbol)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LoadBalanceInPipsBySymbol is a free data retrieval call binding the contract method 0xef3b9d4a.
//
// Solidity: function loadBalanceInPipsBySymbol(address wallet, string assetSymbol) view returns(uint64)
func (_Ydexabi *YdexabiSession) LoadBalanceInPipsBySymbol(wallet common.Address, assetSymbol string) (uint64, error) {
	return _Ydexabi.Contract.LoadBalanceInPipsBySymbol(&_Ydexabi.CallOpts, wallet, assetSymbol)
}

// LoadBalanceInPipsBySymbol is a free data retrieval call binding the contract method 0xef3b9d4a.
//
// Solidity: function loadBalanceInPipsBySymbol(address wallet, string assetSymbol) view returns(uint64)
func (_Ydexabi *YdexabiCallerSession) LoadBalanceInPipsBySymbol(wallet common.Address, assetSymbol string) (uint64, error) {
	return _Ydexabi.Contract.LoadBalanceInPipsBySymbol(&_Ydexabi.CallOpts, wallet, assetSymbol)
}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_Ydexabi *YdexabiCaller) LoadFeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadFeeWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_Ydexabi *YdexabiSession) LoadFeeWallet() (common.Address, error) {
	return _Ydexabi.Contract.LoadFeeWallet(&_Ydexabi.CallOpts)
}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_Ydexabi *YdexabiCallerSession) LoadFeeWallet() (common.Address, error) {
	return _Ydexabi.Contract.LoadFeeWallet(&_Ydexabi.CallOpts)
}

// LoadPartiallyFilledOrderQuantityInPips is a free data retrieval call binding the contract method 0xae0e969e.
//
// Solidity: function loadPartiallyFilledOrderQuantityInPips(bytes32 orderHash) view returns(uint64)
func (_Ydexabi *YdexabiCaller) LoadPartiallyFilledOrderQuantityInPips(opts *bind.CallOpts, orderHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _Ydexabi.contract.Call(opts, &out, "loadPartiallyFilledOrderQuantityInPips", orderHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LoadPartiallyFilledOrderQuantityInPips is a free data retrieval call binding the contract method 0xae0e969e.
//
// Solidity: function loadPartiallyFilledOrderQuantityInPips(bytes32 orderHash) view returns(uint64)
func (_Ydexabi *YdexabiSession) LoadPartiallyFilledOrderQuantityInPips(orderHash [32]byte) (uint64, error) {
	return _Ydexabi.Contract.LoadPartiallyFilledOrderQuantityInPips(&_Ydexabi.CallOpts, orderHash)
}

// LoadPartiallyFilledOrderQuantityInPips is a free data retrieval call binding the contract method 0xae0e969e.
//
// Solidity: function loadPartiallyFilledOrderQuantityInPips(bytes32 orderHash) view returns(uint64)
func (_Ydexabi *YdexabiCallerSession) LoadPartiallyFilledOrderQuantityInPips(orderHash [32]byte) (uint64, error) {
	return _Ydexabi.Contract.LoadPartiallyFilledOrderQuantityInPips(&_Ydexabi.CallOpts, orderHash)
}

// AddTokenSymbol is a paid mutator transaction binding the contract method 0x72e8f08d.
//
// Solidity: function addTokenSymbol(address tokenAddress, string symbol) returns()
func (_Ydexabi *YdexabiTransactor) AddTokenSymbol(opts *bind.TransactOpts, tokenAddress common.Address, symbol string) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "addTokenSymbol", tokenAddress, symbol)
}

// AddTokenSymbol is a paid mutator transaction binding the contract method 0x72e8f08d.
//
// Solidity: function addTokenSymbol(address tokenAddress, string symbol) returns()
func (_Ydexabi *YdexabiSession) AddTokenSymbol(tokenAddress common.Address, symbol string) (*types.Transaction, error) {
	return _Ydexabi.Contract.AddTokenSymbol(&_Ydexabi.TransactOpts, tokenAddress, symbol)
}

// AddTokenSymbol is a paid mutator transaction binding the contract method 0x72e8f08d.
//
// Solidity: function addTokenSymbol(address tokenAddress, string symbol) returns()
func (_Ydexabi *YdexabiTransactorSession) AddTokenSymbol(tokenAddress common.Address, symbol string) (*types.Transaction, error) {
	return _Ydexabi.Contract.AddTokenSymbol(&_Ydexabi.TransactOpts, tokenAddress, symbol)
}

// ClearWalletExit is a paid mutator transaction binding the contract method 0x4a284ef9.
//
// Solidity: function clearWalletExit() returns()
func (_Ydexabi *YdexabiTransactor) ClearWalletExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "clearWalletExit")
}

// ClearWalletExit is a paid mutator transaction binding the contract method 0x4a284ef9.
//
// Solidity: function clearWalletExit() returns()
func (_Ydexabi *YdexabiSession) ClearWalletExit() (*types.Transaction, error) {
	return _Ydexabi.Contract.ClearWalletExit(&_Ydexabi.TransactOpts)
}

// ClearWalletExit is a paid mutator transaction binding the contract method 0x4a284ef9.
//
// Solidity: function clearWalletExit() returns()
func (_Ydexabi *YdexabiTransactorSession) ClearWalletExit() (*types.Transaction, error) {
	return _Ydexabi.Contract.ClearWalletExit(&_Ydexabi.TransactOpts)
}

// ConfirmTokenRegistration is a paid mutator transaction binding the contract method 0x0226b70e.
//
// Solidity: function confirmTokenRegistration(address tokenAddress, string symbol, uint8 decimals) returns()
func (_Ydexabi *YdexabiTransactor) ConfirmTokenRegistration(opts *bind.TransactOpts, tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "confirmTokenRegistration", tokenAddress, symbol, decimals)
}

// ConfirmTokenRegistration is a paid mutator transaction binding the contract method 0x0226b70e.
//
// Solidity: function confirmTokenRegistration(address tokenAddress, string symbol, uint8 decimals) returns()
func (_Ydexabi *YdexabiSession) ConfirmTokenRegistration(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Ydexabi.Contract.ConfirmTokenRegistration(&_Ydexabi.TransactOpts, tokenAddress, symbol, decimals)
}

// ConfirmTokenRegistration is a paid mutator transaction binding the contract method 0x0226b70e.
//
// Solidity: function confirmTokenRegistration(address tokenAddress, string symbol, uint8 decimals) returns()
func (_Ydexabi *YdexabiTransactorSession) ConfirmTokenRegistration(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Ydexabi.Contract.ConfirmTokenRegistration(&_Ydexabi.TransactOpts, tokenAddress, symbol, decimals)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x42220f34.
//
// Solidity: function depositBNB() payable returns()
func (_Ydexabi *YdexabiTransactor) DepositBNB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "depositBNB")
}

// DepositBNB is a paid mutator transaction binding the contract method 0x42220f34.
//
// Solidity: function depositBNB() payable returns()
func (_Ydexabi *YdexabiSession) DepositBNB() (*types.Transaction, error) {
	return _Ydexabi.Contract.DepositBNB(&_Ydexabi.TransactOpts)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x42220f34.
//
// Solidity: function depositBNB() payable returns()
func (_Ydexabi *YdexabiTransactorSession) DepositBNB() (*types.Transaction, error) {
	return _Ydexabi.Contract.DepositBNB(&_Ydexabi.TransactOpts)
}

// DepositTokenByAddress is a paid mutator transaction binding the contract method 0xdcc63490.
//
// Solidity: function depositTokenByAddress(address tokenAddress, uint256 quantityInAssetUnits) returns()
func (_Ydexabi *YdexabiTransactor) DepositTokenByAddress(opts *bind.TransactOpts, tokenAddress common.Address, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "depositTokenByAddress", tokenAddress, quantityInAssetUnits)
}

// DepositTokenByAddress is a paid mutator transaction binding the contract method 0xdcc63490.
//
// Solidity: function depositTokenByAddress(address tokenAddress, uint256 quantityInAssetUnits) returns()
func (_Ydexabi *YdexabiSession) DepositTokenByAddress(tokenAddress common.Address, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.DepositTokenByAddress(&_Ydexabi.TransactOpts, tokenAddress, quantityInAssetUnits)
}

// DepositTokenByAddress is a paid mutator transaction binding the contract method 0xdcc63490.
//
// Solidity: function depositTokenByAddress(address tokenAddress, uint256 quantityInAssetUnits) returns()
func (_Ydexabi *YdexabiTransactorSession) DepositTokenByAddress(tokenAddress common.Address, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.DepositTokenByAddress(&_Ydexabi.TransactOpts, tokenAddress, quantityInAssetUnits)
}

// DepositTokenBySymbol is a paid mutator transaction binding the contract method 0x41968182.
//
// Solidity: function depositTokenBySymbol(string assetSymbol, uint256 quantityInAssetUnits) returns()
func (_Ydexabi *YdexabiTransactor) DepositTokenBySymbol(opts *bind.TransactOpts, assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "depositTokenBySymbol", assetSymbol, quantityInAssetUnits)
}

// DepositTokenBySymbol is a paid mutator transaction binding the contract method 0x41968182.
//
// Solidity: function depositTokenBySymbol(string assetSymbol, uint256 quantityInAssetUnits) returns()
func (_Ydexabi *YdexabiSession) DepositTokenBySymbol(assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.DepositTokenBySymbol(&_Ydexabi.TransactOpts, assetSymbol, quantityInAssetUnits)
}

// DepositTokenBySymbol is a paid mutator transaction binding the contract method 0x41968182.
//
// Solidity: function depositTokenBySymbol(string assetSymbol, uint256 quantityInAssetUnits) returns()
func (_Ydexabi *YdexabiTransactorSession) DepositTokenBySymbol(assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.DepositTokenBySymbol(&_Ydexabi.TransactOpts, assetSymbol, quantityInAssetUnits)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x65fce5b7.
//
// Solidity: function executeTrade((uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) buy, (uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) sell, (string,string,address,address,uint64,uint64,uint64,uint64,address,address,uint64,uint64,uint64,uint8) trade) returns()
func (_Ydexabi *YdexabiTransactor) ExecuteTrade(opts *bind.TransactOpts, buy StructsOrder, sell StructsOrder, trade StructsTrade) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "executeTrade", buy, sell, trade)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x65fce5b7.
//
// Solidity: function executeTrade((uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) buy, (uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) sell, (string,string,address,address,uint64,uint64,uint64,uint64,address,address,uint64,uint64,uint64,uint8) trade) returns()
func (_Ydexabi *YdexabiSession) ExecuteTrade(buy StructsOrder, sell StructsOrder, trade StructsTrade) (*types.Transaction, error) {
	return _Ydexabi.Contract.ExecuteTrade(&_Ydexabi.TransactOpts, buy, sell, trade)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x65fce5b7.
//
// Solidity: function executeTrade((uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) buy, (uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) sell, (string,string,address,address,uint64,uint64,uint64,uint64,address,address,uint64,uint64,uint64,uint8) trade) returns()
func (_Ydexabi *YdexabiTransactorSession) ExecuteTrade(buy StructsOrder, sell StructsOrder, trade StructsTrade) (*types.Transaction, error) {
	return _Ydexabi.Contract.ExecuteTrade(&_Ydexabi.TransactOpts, buy, sell, trade)
}

// ExitWallet is a paid mutator transaction binding the contract method 0xeb5068f2.
//
// Solidity: function exitWallet() returns()
func (_Ydexabi *YdexabiTransactor) ExitWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "exitWallet")
}

// ExitWallet is a paid mutator transaction binding the contract method 0xeb5068f2.
//
// Solidity: function exitWallet() returns()
func (_Ydexabi *YdexabiSession) ExitWallet() (*types.Transaction, error) {
	return _Ydexabi.Contract.ExitWallet(&_Ydexabi.TransactOpts)
}

// ExitWallet is a paid mutator transaction binding the contract method 0xeb5068f2.
//
// Solidity: function exitWallet() returns()
func (_Ydexabi *YdexabiTransactorSession) ExitWallet() (*types.Transaction, error) {
	return _Ydexabi.Contract.ExitWallet(&_Ydexabi.TransactOpts)
}

// InvalidateOrderNonce is a paid mutator transaction binding the contract method 0xd7a6aec7.
//
// Solidity: function invalidateOrderNonce(uint128 nonce) returns()
func (_Ydexabi *YdexabiTransactor) InvalidateOrderNonce(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "invalidateOrderNonce", nonce)
}

// InvalidateOrderNonce is a paid mutator transaction binding the contract method 0xd7a6aec7.
//
// Solidity: function invalidateOrderNonce(uint128 nonce) returns()
func (_Ydexabi *YdexabiSession) InvalidateOrderNonce(nonce *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.InvalidateOrderNonce(&_Ydexabi.TransactOpts, nonce)
}

// InvalidateOrderNonce is a paid mutator transaction binding the contract method 0xd7a6aec7.
//
// Solidity: function invalidateOrderNonce(uint128 nonce) returns()
func (_Ydexabi *YdexabiTransactorSession) InvalidateOrderNonce(nonce *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.InvalidateOrderNonce(&_Ydexabi.TransactOpts, nonce)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x457aa3c6.
//
// Solidity: function registerToken(address tokenAddress, string symbol, uint8 decimals) returns()
func (_Ydexabi *YdexabiTransactor) RegisterToken(opts *bind.TransactOpts, tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "registerToken", tokenAddress, symbol, decimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x457aa3c6.
//
// Solidity: function registerToken(address tokenAddress, string symbol, uint8 decimals) returns()
func (_Ydexabi *YdexabiSession) RegisterToken(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Ydexabi.Contract.RegisterToken(&_Ydexabi.TransactOpts, tokenAddress, symbol, decimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x457aa3c6.
//
// Solidity: function registerToken(address tokenAddress, string symbol, uint8 decimals) returns()
func (_Ydexabi *YdexabiTransactorSession) RegisterToken(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Ydexabi.Contract.RegisterToken(&_Ydexabi.TransactOpts, tokenAddress, symbol, decimals)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Ydexabi *YdexabiTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Ydexabi *YdexabiSession) RemoveAdmin() (*types.Transaction, error) {
	return _Ydexabi.Contract.RemoveAdmin(&_Ydexabi.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Ydexabi *YdexabiTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Ydexabi.Contract.RemoveAdmin(&_Ydexabi.TransactOpts)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x0c187a72.
//
// Solidity: function removeDispatcher() returns()
func (_Ydexabi *YdexabiTransactor) RemoveDispatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "removeDispatcher")
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x0c187a72.
//
// Solidity: function removeDispatcher() returns()
func (_Ydexabi *YdexabiSession) RemoveDispatcher() (*types.Transaction, error) {
	return _Ydexabi.Contract.RemoveDispatcher(&_Ydexabi.TransactOpts)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x0c187a72.
//
// Solidity: function removeDispatcher() returns()
func (_Ydexabi *YdexabiTransactorSession) RemoveDispatcher() (*types.Transaction, error) {
	return _Ydexabi.Contract.RemoveDispatcher(&_Ydexabi.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Ydexabi *YdexabiTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Ydexabi *YdexabiSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetAdmin(&_Ydexabi.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Ydexabi *YdexabiTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetAdmin(&_Ydexabi.TransactOpts, newAdmin)
}

// SetChainPropagationPeriod is a paid mutator transaction binding the contract method 0x20e6a8e3.
//
// Solidity: function setChainPropagationPeriod(uint256 newChainPropagationPeriod) returns()
func (_Ydexabi *YdexabiTransactor) SetChainPropagationPeriod(opts *bind.TransactOpts, newChainPropagationPeriod *big.Int) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "setChainPropagationPeriod", newChainPropagationPeriod)
}

// SetChainPropagationPeriod is a paid mutator transaction binding the contract method 0x20e6a8e3.
//
// Solidity: function setChainPropagationPeriod(uint256 newChainPropagationPeriod) returns()
func (_Ydexabi *YdexabiSession) SetChainPropagationPeriod(newChainPropagationPeriod *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetChainPropagationPeriod(&_Ydexabi.TransactOpts, newChainPropagationPeriod)
}

// SetChainPropagationPeriod is a paid mutator transaction binding the contract method 0x20e6a8e3.
//
// Solidity: function setChainPropagationPeriod(uint256 newChainPropagationPeriod) returns()
func (_Ydexabi *YdexabiTransactorSession) SetChainPropagationPeriod(newChainPropagationPeriod *big.Int) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetChainPropagationPeriod(&_Ydexabi.TransactOpts, newChainPropagationPeriod)
}

// SetCustodian is a paid mutator transaction binding the contract method 0x403f3731.
//
// Solidity: function setCustodian(address newCustodian) returns()
func (_Ydexabi *YdexabiTransactor) SetCustodian(opts *bind.TransactOpts, newCustodian common.Address) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "setCustodian", newCustodian)
}

// SetCustodian is a paid mutator transaction binding the contract method 0x403f3731.
//
// Solidity: function setCustodian(address newCustodian) returns()
func (_Ydexabi *YdexabiSession) SetCustodian(newCustodian common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetCustodian(&_Ydexabi.TransactOpts, newCustodian)
}

// SetCustodian is a paid mutator transaction binding the contract method 0x403f3731.
//
// Solidity: function setCustodian(address newCustodian) returns()
func (_Ydexabi *YdexabiTransactorSession) SetCustodian(newCustodian common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetCustodian(&_Ydexabi.TransactOpts, newCustodian)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address newDispatcherWallet) returns()
func (_Ydexabi *YdexabiTransactor) SetDispatcher(opts *bind.TransactOpts, newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "setDispatcher", newDispatcherWallet)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address newDispatcherWallet) returns()
func (_Ydexabi *YdexabiSession) SetDispatcher(newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetDispatcher(&_Ydexabi.TransactOpts, newDispatcherWallet)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address newDispatcherWallet) returns()
func (_Ydexabi *YdexabiTransactorSession) SetDispatcher(newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetDispatcher(&_Ydexabi.TransactOpts, newDispatcherWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_Ydexabi *YdexabiTransactor) SetFeeWallet(opts *bind.TransactOpts, newFeeWallet common.Address) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "setFeeWallet", newFeeWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_Ydexabi *YdexabiSession) SetFeeWallet(newFeeWallet common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetFeeWallet(&_Ydexabi.TransactOpts, newFeeWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_Ydexabi *YdexabiTransactorSession) SetFeeWallet(newFeeWallet common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.SetFeeWallet(&_Ydexabi.TransactOpts, newFeeWallet)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6bb509f2.
//
// Solidity: function withdraw((uint8,uint128,address,string,address,uint64,uint64,bool,bytes) withdrawal) returns()
func (_Ydexabi *YdexabiTransactor) Withdraw(opts *bind.TransactOpts, withdrawal StructsWithdrawal) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "withdraw", withdrawal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6bb509f2.
//
// Solidity: function withdraw((uint8,uint128,address,string,address,uint64,uint64,bool,bytes) withdrawal) returns()
func (_Ydexabi *YdexabiSession) Withdraw(withdrawal StructsWithdrawal) (*types.Transaction, error) {
	return _Ydexabi.Contract.Withdraw(&_Ydexabi.TransactOpts, withdrawal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6bb509f2.
//
// Solidity: function withdraw((uint8,uint128,address,string,address,uint64,uint64,bool,bytes) withdrawal) returns()
func (_Ydexabi *YdexabiTransactorSession) Withdraw(withdrawal StructsWithdrawal) (*types.Transaction, error) {
	return _Ydexabi.Contract.Withdraw(&_Ydexabi.TransactOpts, withdrawal)
}

// WithdrawAllExit is a paid mutator transaction binding the contract method 0xf34e2ef1.
//
// Solidity: function withdrawAllExit() returns()
func (_Ydexabi *YdexabiTransactor) WithdrawAllExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "withdrawAllExit")
}

// WithdrawAllExit is a paid mutator transaction binding the contract method 0xf34e2ef1.
//
// Solidity: function withdrawAllExit() returns()
func (_Ydexabi *YdexabiSession) WithdrawAllExit() (*types.Transaction, error) {
	return _Ydexabi.Contract.WithdrawAllExit(&_Ydexabi.TransactOpts)
}

// WithdrawAllExit is a paid mutator transaction binding the contract method 0xf34e2ef1.
//
// Solidity: function withdrawAllExit() returns()
func (_Ydexabi *YdexabiTransactorSession) WithdrawAllExit() (*types.Transaction, error) {
	return _Ydexabi.Contract.WithdrawAllExit(&_Ydexabi.TransactOpts)
}

// WithdrawExit is a paid mutator transaction binding the contract method 0x985c4af5.
//
// Solidity: function withdrawExit(address assetAddress) returns()
func (_Ydexabi *YdexabiTransactor) WithdrawExit(opts *bind.TransactOpts, assetAddress common.Address) (*types.Transaction, error) {
	return _Ydexabi.contract.Transact(opts, "withdrawExit", assetAddress)
}

// WithdrawExit is a paid mutator transaction binding the contract method 0x985c4af5.
//
// Solidity: function withdrawExit(address assetAddress) returns()
func (_Ydexabi *YdexabiSession) WithdrawExit(assetAddress common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.WithdrawExit(&_Ydexabi.TransactOpts, assetAddress)
}

// WithdrawExit is a paid mutator transaction binding the contract method 0x985c4af5.
//
// Solidity: function withdrawExit(address assetAddress) returns()
func (_Ydexabi *YdexabiTransactorSession) WithdrawExit(assetAddress common.Address) (*types.Transaction, error) {
	return _Ydexabi.Contract.WithdrawExit(&_Ydexabi.TransactOpts, assetAddress)
}

// YdexabiChainPropagationPeriodChangedIterator is returned from FilterChainPropagationPeriodChanged and is used to iterate over the raw logs and unpacked data for ChainPropagationPeriodChanged events raised by the Ydexabi contract.
type YdexabiChainPropagationPeriodChangedIterator struct {
	Event *YdexabiChainPropagationPeriodChanged // Event containing the contract specifics and raw log

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
func (it *YdexabiChainPropagationPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiChainPropagationPeriodChanged)
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
		it.Event = new(YdexabiChainPropagationPeriodChanged)
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
func (it *YdexabiChainPropagationPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiChainPropagationPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiChainPropagationPeriodChanged represents a ChainPropagationPeriodChanged event raised by the Ydexabi contract.
type YdexabiChainPropagationPeriodChanged struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPropagationPeriodChanged is a free log retrieval operation binding the contract event 0x9a22227d6c0251a79ef8b846202ddcbe9d682ee5482e84abeec6dda096398a6f.
//
// Solidity: event ChainPropagationPeriodChanged(uint256 previousValue, uint256 newValue)
func (_Ydexabi *YdexabiFilterer) FilterChainPropagationPeriodChanged(opts *bind.FilterOpts) (*YdexabiChainPropagationPeriodChangedIterator, error) {

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "ChainPropagationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &YdexabiChainPropagationPeriodChangedIterator{contract: _Ydexabi.contract, event: "ChainPropagationPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchChainPropagationPeriodChanged is a free log subscription operation binding the contract event 0x9a22227d6c0251a79ef8b846202ddcbe9d682ee5482e84abeec6dda096398a6f.
//
// Solidity: event ChainPropagationPeriodChanged(uint256 previousValue, uint256 newValue)
func (_Ydexabi *YdexabiFilterer) WatchChainPropagationPeriodChanged(opts *bind.WatchOpts, sink chan<- *YdexabiChainPropagationPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "ChainPropagationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiChainPropagationPeriodChanged)
				if err := _Ydexabi.contract.UnpackLog(event, "ChainPropagationPeriodChanged", log); err != nil {
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

// ParseChainPropagationPeriodChanged is a log parse operation binding the contract event 0x9a22227d6c0251a79ef8b846202ddcbe9d682ee5482e84abeec6dda096398a6f.
//
// Solidity: event ChainPropagationPeriodChanged(uint256 previousValue, uint256 newValue)
func (_Ydexabi *YdexabiFilterer) ParseChainPropagationPeriodChanged(log types.Log) (*YdexabiChainPropagationPeriodChanged, error) {
	event := new(YdexabiChainPropagationPeriodChanged)
	if err := _Ydexabi.contract.UnpackLog(event, "ChainPropagationPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Ydexabi contract.
type YdexabiDepositedIterator struct {
	Event *YdexabiDeposited // Event containing the contract specifics and raw log

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
func (it *YdexabiDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiDeposited)
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
		it.Event = new(YdexabiDeposited)
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
func (it *YdexabiDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiDeposited represents a Deposited event raised by the Ydexabi contract.
type YdexabiDeposited struct {
	Index                          uint64
	Wallet                         common.Address
	AssetAddress                   common.Address
	AssetSymbolIndex               common.Hash
	AssetSymbol                    string
	QuantityInPips                 uint64
	NewExchangeBalanceInPips       uint64
	NewExchangeBalanceInAssetUnits *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xc2813a84168de59cf710030ba2acb13031b567b8ba76f70e8ab0782aa9469a7a.
//
// Solidity: event Deposited(uint64 index, address indexed wallet, address indexed assetAddress, string indexed assetSymbolIndex, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) FilterDeposited(opts *bind.FilterOpts, wallet []common.Address, assetAddress []common.Address, assetSymbolIndex []string) (*YdexabiDepositedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var assetSymbolIndexRule []interface{}
	for _, assetSymbolIndexItem := range assetSymbolIndex {
		assetSymbolIndexRule = append(assetSymbolIndexRule, assetSymbolIndexItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "Deposited", walletRule, assetAddressRule, assetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiDepositedIterator{contract: _Ydexabi.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc2813a84168de59cf710030ba2acb13031b567b8ba76f70e8ab0782aa9469a7a.
//
// Solidity: event Deposited(uint64 index, address indexed wallet, address indexed assetAddress, string indexed assetSymbolIndex, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *YdexabiDeposited, wallet []common.Address, assetAddress []common.Address, assetSymbolIndex []string) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var assetSymbolIndexRule []interface{}
	for _, assetSymbolIndexItem := range assetSymbolIndex {
		assetSymbolIndexRule = append(assetSymbolIndexRule, assetSymbolIndexItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "Deposited", walletRule, assetAddressRule, assetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiDeposited)
				if err := _Ydexabi.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xc2813a84168de59cf710030ba2acb13031b567b8ba76f70e8ab0782aa9469a7a.
//
// Solidity: event Deposited(uint64 index, address indexed wallet, address indexed assetAddress, string indexed assetSymbolIndex, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) ParseDeposited(log types.Log) (*YdexabiDeposited, error) {
	event := new(YdexabiDeposited)
	if err := _Ydexabi.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiDispatcherChangedIterator is returned from FilterDispatcherChanged and is used to iterate over the raw logs and unpacked data for DispatcherChanged events raised by the Ydexabi contract.
type YdexabiDispatcherChangedIterator struct {
	Event *YdexabiDispatcherChanged // Event containing the contract specifics and raw log

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
func (it *YdexabiDispatcherChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiDispatcherChanged)
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
		it.Event = new(YdexabiDispatcherChanged)
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
func (it *YdexabiDispatcherChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiDispatcherChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiDispatcherChanged represents a DispatcherChanged event raised by the Ydexabi contract.
type YdexabiDispatcherChanged struct {
	PreviousValue common.Address
	NewValue      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDispatcherChanged is a free log retrieval operation binding the contract event 0xddb86df2d5262dd7b44067db5962cc7875e9db409cb21c88adfe3c5760315e39.
//
// Solidity: event DispatcherChanged(address previousValue, address newValue)
func (_Ydexabi *YdexabiFilterer) FilterDispatcherChanged(opts *bind.FilterOpts) (*YdexabiDispatcherChangedIterator, error) {

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "DispatcherChanged")
	if err != nil {
		return nil, err
	}
	return &YdexabiDispatcherChangedIterator{contract: _Ydexabi.contract, event: "DispatcherChanged", logs: logs, sub: sub}, nil
}

// WatchDispatcherChanged is a free log subscription operation binding the contract event 0xddb86df2d5262dd7b44067db5962cc7875e9db409cb21c88adfe3c5760315e39.
//
// Solidity: event DispatcherChanged(address previousValue, address newValue)
func (_Ydexabi *YdexabiFilterer) WatchDispatcherChanged(opts *bind.WatchOpts, sink chan<- *YdexabiDispatcherChanged) (event.Subscription, error) {

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "DispatcherChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiDispatcherChanged)
				if err := _Ydexabi.contract.UnpackLog(event, "DispatcherChanged", log); err != nil {
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

// ParseDispatcherChanged is a log parse operation binding the contract event 0xddb86df2d5262dd7b44067db5962cc7875e9db409cb21c88adfe3c5760315e39.
//
// Solidity: event DispatcherChanged(address previousValue, address newValue)
func (_Ydexabi *YdexabiFilterer) ParseDispatcherChanged(log types.Log) (*YdexabiDispatcherChanged, error) {
	event := new(YdexabiDispatcherChanged)
	if err := _Ydexabi.contract.UnpackLog(event, "DispatcherChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiFeeWalletChangedIterator is returned from FilterFeeWalletChanged and is used to iterate over the raw logs and unpacked data for FeeWalletChanged events raised by the Ydexabi contract.
type YdexabiFeeWalletChangedIterator struct {
	Event *YdexabiFeeWalletChanged // Event containing the contract specifics and raw log

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
func (it *YdexabiFeeWalletChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiFeeWalletChanged)
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
		it.Event = new(YdexabiFeeWalletChanged)
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
func (it *YdexabiFeeWalletChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiFeeWalletChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiFeeWalletChanged represents a FeeWalletChanged event raised by the Ydexabi contract.
type YdexabiFeeWalletChanged struct {
	PreviousValue common.Address
	NewValue      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeWalletChanged is a free log retrieval operation binding the contract event 0x9f4f5dce3c4d197b5d7496cb96e25f0a89809167195964b0daa3ef5fed63c00a.
//
// Solidity: event FeeWalletChanged(address previousValue, address newValue)
func (_Ydexabi *YdexabiFilterer) FilterFeeWalletChanged(opts *bind.FilterOpts) (*YdexabiFeeWalletChangedIterator, error) {

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "FeeWalletChanged")
	if err != nil {
		return nil, err
	}
	return &YdexabiFeeWalletChangedIterator{contract: _Ydexabi.contract, event: "FeeWalletChanged", logs: logs, sub: sub}, nil
}

// WatchFeeWalletChanged is a free log subscription operation binding the contract event 0x9f4f5dce3c4d197b5d7496cb96e25f0a89809167195964b0daa3ef5fed63c00a.
//
// Solidity: event FeeWalletChanged(address previousValue, address newValue)
func (_Ydexabi *YdexabiFilterer) WatchFeeWalletChanged(opts *bind.WatchOpts, sink chan<- *YdexabiFeeWalletChanged) (event.Subscription, error) {

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "FeeWalletChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiFeeWalletChanged)
				if err := _Ydexabi.contract.UnpackLog(event, "FeeWalletChanged", log); err != nil {
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

// ParseFeeWalletChanged is a log parse operation binding the contract event 0x9f4f5dce3c4d197b5d7496cb96e25f0a89809167195964b0daa3ef5fed63c00a.
//
// Solidity: event FeeWalletChanged(address previousValue, address newValue)
func (_Ydexabi *YdexabiFilterer) ParseFeeWalletChanged(log types.Log) (*YdexabiFeeWalletChanged, error) {
	event := new(YdexabiFeeWalletChanged)
	if err := _Ydexabi.contract.UnpackLog(event, "FeeWalletChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiOrderNonceInvalidatedIterator is returned from FilterOrderNonceInvalidated and is used to iterate over the raw logs and unpacked data for OrderNonceInvalidated events raised by the Ydexabi contract.
type YdexabiOrderNonceInvalidatedIterator struct {
	Event *YdexabiOrderNonceInvalidated // Event containing the contract specifics and raw log

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
func (it *YdexabiOrderNonceInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiOrderNonceInvalidated)
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
		it.Event = new(YdexabiOrderNonceInvalidated)
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
func (it *YdexabiOrderNonceInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiOrderNonceInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiOrderNonceInvalidated represents a OrderNonceInvalidated event raised by the Ydexabi contract.
type YdexabiOrderNonceInvalidated struct {
	Wallet               common.Address
	Nonce                *big.Int
	TimestampInMs        *big.Int
	EffectiveBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOrderNonceInvalidated is a free log retrieval operation binding the contract event 0x10cf19671b43c88b1f02d4e94932d7ffaa89c7278bc5b8868fa7b7676210809b.
//
// Solidity: event OrderNonceInvalidated(address indexed wallet, uint128 nonce, uint128 timestampInMs, uint256 effectiveBlockNumber)
func (_Ydexabi *YdexabiFilterer) FilterOrderNonceInvalidated(opts *bind.FilterOpts, wallet []common.Address) (*YdexabiOrderNonceInvalidatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "OrderNonceInvalidated", walletRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiOrderNonceInvalidatedIterator{contract: _Ydexabi.contract, event: "OrderNonceInvalidated", logs: logs, sub: sub}, nil
}

// WatchOrderNonceInvalidated is a free log subscription operation binding the contract event 0x10cf19671b43c88b1f02d4e94932d7ffaa89c7278bc5b8868fa7b7676210809b.
//
// Solidity: event OrderNonceInvalidated(address indexed wallet, uint128 nonce, uint128 timestampInMs, uint256 effectiveBlockNumber)
func (_Ydexabi *YdexabiFilterer) WatchOrderNonceInvalidated(opts *bind.WatchOpts, sink chan<- *YdexabiOrderNonceInvalidated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "OrderNonceInvalidated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiOrderNonceInvalidated)
				if err := _Ydexabi.contract.UnpackLog(event, "OrderNonceInvalidated", log); err != nil {
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

// ParseOrderNonceInvalidated is a log parse operation binding the contract event 0x10cf19671b43c88b1f02d4e94932d7ffaa89c7278bc5b8868fa7b7676210809b.
//
// Solidity: event OrderNonceInvalidated(address indexed wallet, uint128 nonce, uint128 timestampInMs, uint256 effectiveBlockNumber)
func (_Ydexabi *YdexabiFilterer) ParseOrderNonceInvalidated(log types.Log) (*YdexabiOrderNonceInvalidated, error) {
	event := new(YdexabiOrderNonceInvalidated)
	if err := _Ydexabi.contract.UnpackLog(event, "OrderNonceInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the Ydexabi contract.
type YdexabiTokenRegisteredIterator struct {
	Event *YdexabiTokenRegistered // Event containing the contract specifics and raw log

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
func (it *YdexabiTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiTokenRegistered)
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
		it.Event = new(YdexabiTokenRegistered)
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
func (it *YdexabiTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiTokenRegistered represents a TokenRegistered event raised by the Ydexabi contract.
type YdexabiTokenRegistered struct {
	AssetAddress common.Address
	AssetSymbol  string
	Decimals     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xcf66acda7673c71cd9458c59067eb30a473eeaf2fa05d008acf47616cbbbe01c.
//
// Solidity: event TokenRegistered(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_Ydexabi *YdexabiFilterer) FilterTokenRegistered(opts *bind.FilterOpts, assetAddress []common.Address) (*YdexabiTokenRegisteredIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "TokenRegistered", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiTokenRegisteredIterator{contract: _Ydexabi.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xcf66acda7673c71cd9458c59067eb30a473eeaf2fa05d008acf47616cbbbe01c.
//
// Solidity: event TokenRegistered(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_Ydexabi *YdexabiFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *YdexabiTokenRegistered, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "TokenRegistered", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiTokenRegistered)
				if err := _Ydexabi.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0xcf66acda7673c71cd9458c59067eb30a473eeaf2fa05d008acf47616cbbbe01c.
//
// Solidity: event TokenRegistered(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_Ydexabi *YdexabiFilterer) ParseTokenRegistered(log types.Log) (*YdexabiTokenRegistered, error) {
	event := new(YdexabiTokenRegistered)
	if err := _Ydexabi.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiTokenRegistrationConfirmedIterator is returned from FilterTokenRegistrationConfirmed and is used to iterate over the raw logs and unpacked data for TokenRegistrationConfirmed events raised by the Ydexabi contract.
type YdexabiTokenRegistrationConfirmedIterator struct {
	Event *YdexabiTokenRegistrationConfirmed // Event containing the contract specifics and raw log

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
func (it *YdexabiTokenRegistrationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiTokenRegistrationConfirmed)
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
		it.Event = new(YdexabiTokenRegistrationConfirmed)
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
func (it *YdexabiTokenRegistrationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiTokenRegistrationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiTokenRegistrationConfirmed represents a TokenRegistrationConfirmed event raised by the Ydexabi contract.
type YdexabiTokenRegistrationConfirmed struct {
	AssetAddress common.Address
	AssetSymbol  string
	Decimals     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistrationConfirmed is a free log retrieval operation binding the contract event 0xa8968236dc4afb2c189afeb69c91b02445e6bd09f0ffe27256b62838edc99edb.
//
// Solidity: event TokenRegistrationConfirmed(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_Ydexabi *YdexabiFilterer) FilterTokenRegistrationConfirmed(opts *bind.FilterOpts, assetAddress []common.Address) (*YdexabiTokenRegistrationConfirmedIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "TokenRegistrationConfirmed", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiTokenRegistrationConfirmedIterator{contract: _Ydexabi.contract, event: "TokenRegistrationConfirmed", logs: logs, sub: sub}, nil
}

// WatchTokenRegistrationConfirmed is a free log subscription operation binding the contract event 0xa8968236dc4afb2c189afeb69c91b02445e6bd09f0ffe27256b62838edc99edb.
//
// Solidity: event TokenRegistrationConfirmed(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_Ydexabi *YdexabiFilterer) WatchTokenRegistrationConfirmed(opts *bind.WatchOpts, sink chan<- *YdexabiTokenRegistrationConfirmed, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "TokenRegistrationConfirmed", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiTokenRegistrationConfirmed)
				if err := _Ydexabi.contract.UnpackLog(event, "TokenRegistrationConfirmed", log); err != nil {
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

// ParseTokenRegistrationConfirmed is a log parse operation binding the contract event 0xa8968236dc4afb2c189afeb69c91b02445e6bd09f0ffe27256b62838edc99edb.
//
// Solidity: event TokenRegistrationConfirmed(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_Ydexabi *YdexabiFilterer) ParseTokenRegistrationConfirmed(log types.Log) (*YdexabiTokenRegistrationConfirmed, error) {
	event := new(YdexabiTokenRegistrationConfirmed)
	if err := _Ydexabi.contract.UnpackLog(event, "TokenRegistrationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiTokenSymbolAddedIterator is returned from FilterTokenSymbolAdded and is used to iterate over the raw logs and unpacked data for TokenSymbolAdded events raised by the Ydexabi contract.
type YdexabiTokenSymbolAddedIterator struct {
	Event *YdexabiTokenSymbolAdded // Event containing the contract specifics and raw log

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
func (it *YdexabiTokenSymbolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiTokenSymbolAdded)
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
		it.Event = new(YdexabiTokenSymbolAdded)
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
func (it *YdexabiTokenSymbolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiTokenSymbolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiTokenSymbolAdded represents a TokenSymbolAdded event raised by the Ydexabi contract.
type YdexabiTokenSymbolAdded struct {
	AssetAddress common.Address
	AssetSymbol  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenSymbolAdded is a free log retrieval operation binding the contract event 0x7f63a63f4a2ea318da0b031794d0b5a3183a1b2de54e053ceb17cabdba031735.
//
// Solidity: event TokenSymbolAdded(address indexed assetAddress, string assetSymbol)
func (_Ydexabi *YdexabiFilterer) FilterTokenSymbolAdded(opts *bind.FilterOpts, assetAddress []common.Address) (*YdexabiTokenSymbolAddedIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "TokenSymbolAdded", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiTokenSymbolAddedIterator{contract: _Ydexabi.contract, event: "TokenSymbolAdded", logs: logs, sub: sub}, nil
}

// WatchTokenSymbolAdded is a free log subscription operation binding the contract event 0x7f63a63f4a2ea318da0b031794d0b5a3183a1b2de54e053ceb17cabdba031735.
//
// Solidity: event TokenSymbolAdded(address indexed assetAddress, string assetSymbol)
func (_Ydexabi *YdexabiFilterer) WatchTokenSymbolAdded(opts *bind.WatchOpts, sink chan<- *YdexabiTokenSymbolAdded, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "TokenSymbolAdded", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiTokenSymbolAdded)
				if err := _Ydexabi.contract.UnpackLog(event, "TokenSymbolAdded", log); err != nil {
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

// ParseTokenSymbolAdded is a log parse operation binding the contract event 0x7f63a63f4a2ea318da0b031794d0b5a3183a1b2de54e053ceb17cabdba031735.
//
// Solidity: event TokenSymbolAdded(address indexed assetAddress, string assetSymbol)
func (_Ydexabi *YdexabiFilterer) ParseTokenSymbolAdded(log types.Log) (*YdexabiTokenSymbolAdded, error) {
	event := new(YdexabiTokenSymbolAdded)
	if err := _Ydexabi.contract.UnpackLog(event, "TokenSymbolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the Ydexabi contract.
type YdexabiTradeExecutedIterator struct {
	Event *YdexabiTradeExecuted // Event containing the contract specifics and raw log

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
func (it *YdexabiTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiTradeExecuted)
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
		it.Event = new(YdexabiTradeExecuted)
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
func (it *YdexabiTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiTradeExecuted represents a TradeExecuted event raised by the Ydexabi contract.
type YdexabiTradeExecuted struct {
	BuyWallet             common.Address
	SellWallet            common.Address
	BaseAssetSymbolIndex  common.Hash
	QuoteAssetSymbolIndex common.Hash
	BaseAssetSymbol       string
	QuoteAssetSymbol      string
	BaseQuantityInPips    uint64
	QuoteQuantityInPips   uint64
	TradePriceInPips      uint64
	BuyOrderHash          [32]byte
	SellOrderHash         [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTradeExecuted is a free log retrieval operation binding the contract event 0xbd75b2e24c3d43ec9f10f3583265606c0a7d8b67d147f9dc48376b81456823a1.
//
// Solidity: event TradeExecuted(address buyWallet, address sellWallet, string indexed baseAssetSymbolIndex, string indexed quoteAssetSymbolIndex, string baseAssetSymbol, string quoteAssetSymbol, uint64 baseQuantityInPips, uint64 quoteQuantityInPips, uint64 tradePriceInPips, bytes32 buyOrderHash, bytes32 sellOrderHash)
func (_Ydexabi *YdexabiFilterer) FilterTradeExecuted(opts *bind.FilterOpts, baseAssetSymbolIndex []string, quoteAssetSymbolIndex []string) (*YdexabiTradeExecutedIterator, error) {

	var baseAssetSymbolIndexRule []interface{}
	for _, baseAssetSymbolIndexItem := range baseAssetSymbolIndex {
		baseAssetSymbolIndexRule = append(baseAssetSymbolIndexRule, baseAssetSymbolIndexItem)
	}
	var quoteAssetSymbolIndexRule []interface{}
	for _, quoteAssetSymbolIndexItem := range quoteAssetSymbolIndex {
		quoteAssetSymbolIndexRule = append(quoteAssetSymbolIndexRule, quoteAssetSymbolIndexItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "TradeExecuted", baseAssetSymbolIndexRule, quoteAssetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiTradeExecutedIterator{contract: _Ydexabi.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0xbd75b2e24c3d43ec9f10f3583265606c0a7d8b67d147f9dc48376b81456823a1.
//
// Solidity: event TradeExecuted(address buyWallet, address sellWallet, string indexed baseAssetSymbolIndex, string indexed quoteAssetSymbolIndex, string baseAssetSymbol, string quoteAssetSymbol, uint64 baseQuantityInPips, uint64 quoteQuantityInPips, uint64 tradePriceInPips, bytes32 buyOrderHash, bytes32 sellOrderHash)
func (_Ydexabi *YdexabiFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *YdexabiTradeExecuted, baseAssetSymbolIndex []string, quoteAssetSymbolIndex []string) (event.Subscription, error) {

	var baseAssetSymbolIndexRule []interface{}
	for _, baseAssetSymbolIndexItem := range baseAssetSymbolIndex {
		baseAssetSymbolIndexRule = append(baseAssetSymbolIndexRule, baseAssetSymbolIndexItem)
	}
	var quoteAssetSymbolIndexRule []interface{}
	for _, quoteAssetSymbolIndexItem := range quoteAssetSymbolIndex {
		quoteAssetSymbolIndexRule = append(quoteAssetSymbolIndexRule, quoteAssetSymbolIndexItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "TradeExecuted", baseAssetSymbolIndexRule, quoteAssetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiTradeExecuted)
				if err := _Ydexabi.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
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

// ParseTradeExecuted is a log parse operation binding the contract event 0xbd75b2e24c3d43ec9f10f3583265606c0a7d8b67d147f9dc48376b81456823a1.
//
// Solidity: event TradeExecuted(address buyWallet, address sellWallet, string indexed baseAssetSymbolIndex, string indexed quoteAssetSymbolIndex, string baseAssetSymbol, string quoteAssetSymbol, uint64 baseQuantityInPips, uint64 quoteQuantityInPips, uint64 tradePriceInPips, bytes32 buyOrderHash, bytes32 sellOrderHash)
func (_Ydexabi *YdexabiFilterer) ParseTradeExecuted(log types.Log) (*YdexabiTradeExecuted, error) {
	event := new(YdexabiTradeExecuted)
	if err := _Ydexabi.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiWalletExitClearedIterator is returned from FilterWalletExitCleared and is used to iterate over the raw logs and unpacked data for WalletExitCleared events raised by the Ydexabi contract.
type YdexabiWalletExitClearedIterator struct {
	Event *YdexabiWalletExitCleared // Event containing the contract specifics and raw log

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
func (it *YdexabiWalletExitClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiWalletExitCleared)
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
		it.Event = new(YdexabiWalletExitCleared)
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
func (it *YdexabiWalletExitClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiWalletExitClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiWalletExitCleared represents a WalletExitCleared event raised by the Ydexabi contract.
type YdexabiWalletExitCleared struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletExitCleared is a free log retrieval operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address indexed wallet)
func (_Ydexabi *YdexabiFilterer) FilterWalletExitCleared(opts *bind.FilterOpts, wallet []common.Address) (*YdexabiWalletExitClearedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "WalletExitCleared", walletRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiWalletExitClearedIterator{contract: _Ydexabi.contract, event: "WalletExitCleared", logs: logs, sub: sub}, nil
}

// WatchWalletExitCleared is a free log subscription operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address indexed wallet)
func (_Ydexabi *YdexabiFilterer) WatchWalletExitCleared(opts *bind.WatchOpts, sink chan<- *YdexabiWalletExitCleared, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "WalletExitCleared", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiWalletExitCleared)
				if err := _Ydexabi.contract.UnpackLog(event, "WalletExitCleared", log); err != nil {
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

// ParseWalletExitCleared is a log parse operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address indexed wallet)
func (_Ydexabi *YdexabiFilterer) ParseWalletExitCleared(log types.Log) (*YdexabiWalletExitCleared, error) {
	event := new(YdexabiWalletExitCleared)
	if err := _Ydexabi.contract.UnpackLog(event, "WalletExitCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiWalletExitWithdrawnIterator is returned from FilterWalletExitWithdrawn and is used to iterate over the raw logs and unpacked data for WalletExitWithdrawn events raised by the Ydexabi contract.
type YdexabiWalletExitWithdrawnIterator struct {
	Event *YdexabiWalletExitWithdrawn // Event containing the contract specifics and raw log

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
func (it *YdexabiWalletExitWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiWalletExitWithdrawn)
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
		it.Event = new(YdexabiWalletExitWithdrawn)
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
func (it *YdexabiWalletExitWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiWalletExitWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiWalletExitWithdrawn represents a WalletExitWithdrawn event raised by the Ydexabi contract.
type YdexabiWalletExitWithdrawn struct {
	Wallet                         common.Address
	AssetAddress                   common.Address
	AssetSymbol                    string
	QuantityInPips                 uint64
	NewExchangeBalanceInPips       uint64
	NewExchangeBalanceInAssetUnits *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterWalletExitWithdrawn is a free log retrieval operation binding the contract event 0x02e03f232c88c87047c5a9cfbad1213b842503fee6a002ec2eec9f5a64725bf4.
//
// Solidity: event WalletExitWithdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) FilterWalletExitWithdrawn(opts *bind.FilterOpts, wallet []common.Address, assetAddress []common.Address) (*YdexabiWalletExitWithdrawnIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "WalletExitWithdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiWalletExitWithdrawnIterator{contract: _Ydexabi.contract, event: "WalletExitWithdrawn", logs: logs, sub: sub}, nil
}

// WatchWalletExitWithdrawn is a free log subscription operation binding the contract event 0x02e03f232c88c87047c5a9cfbad1213b842503fee6a002ec2eec9f5a64725bf4.
//
// Solidity: event WalletExitWithdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) WatchWalletExitWithdrawn(opts *bind.WatchOpts, sink chan<- *YdexabiWalletExitWithdrawn, wallet []common.Address, assetAddress []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "WalletExitWithdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiWalletExitWithdrawn)
				if err := _Ydexabi.contract.UnpackLog(event, "WalletExitWithdrawn", log); err != nil {
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

// ParseWalletExitWithdrawn is a log parse operation binding the contract event 0x02e03f232c88c87047c5a9cfbad1213b842503fee6a002ec2eec9f5a64725bf4.
//
// Solidity: event WalletExitWithdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) ParseWalletExitWithdrawn(log types.Log) (*YdexabiWalletExitWithdrawn, error) {
	event := new(YdexabiWalletExitWithdrawn)
	if err := _Ydexabi.contract.UnpackLog(event, "WalletExitWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiWalletExitedIterator is returned from FilterWalletExited and is used to iterate over the raw logs and unpacked data for WalletExited events raised by the Ydexabi contract.
type YdexabiWalletExitedIterator struct {
	Event *YdexabiWalletExited // Event containing the contract specifics and raw log

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
func (it *YdexabiWalletExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiWalletExited)
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
		it.Event = new(YdexabiWalletExited)
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
func (it *YdexabiWalletExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiWalletExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiWalletExited represents a WalletExited event raised by the Ydexabi contract.
type YdexabiWalletExited struct {
	Wallet               common.Address
	EffectiveBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterWalletExited is a free log retrieval operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address indexed wallet, uint256 effectiveBlockNumber)
func (_Ydexabi *YdexabiFilterer) FilterWalletExited(opts *bind.FilterOpts, wallet []common.Address) (*YdexabiWalletExitedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "WalletExited", walletRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiWalletExitedIterator{contract: _Ydexabi.contract, event: "WalletExited", logs: logs, sub: sub}, nil
}

// WatchWalletExited is a free log subscription operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address indexed wallet, uint256 effectiveBlockNumber)
func (_Ydexabi *YdexabiFilterer) WatchWalletExited(opts *bind.WatchOpts, sink chan<- *YdexabiWalletExited, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "WalletExited", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiWalletExited)
				if err := _Ydexabi.contract.UnpackLog(event, "WalletExited", log); err != nil {
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

// ParseWalletExited is a log parse operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address indexed wallet, uint256 effectiveBlockNumber)
func (_Ydexabi *YdexabiFilterer) ParseWalletExited(log types.Log) (*YdexabiWalletExited, error) {
	event := new(YdexabiWalletExited)
	if err := _Ydexabi.contract.UnpackLog(event, "WalletExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YdexabiWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Ydexabi contract.
type YdexabiWithdrawnIterator struct {
	Event *YdexabiWithdrawn // Event containing the contract specifics and raw log

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
func (it *YdexabiWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YdexabiWithdrawn)
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
		it.Event = new(YdexabiWithdrawn)
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
func (it *YdexabiWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YdexabiWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YdexabiWithdrawn represents a Withdrawn event raised by the Ydexabi contract.
type YdexabiWithdrawn struct {
	Wallet                         common.Address
	AssetAddress                   common.Address
	AssetSymbol                    string
	QuantityInPips                 uint64
	NewExchangeBalanceInPips       uint64
	NewExchangeBalanceInAssetUnits *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x6960a1f64ecf9da0d1d1bcbfa3dd27f8c1c60de69b13faa28127dafa36c111e4.
//
// Solidity: event Withdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) FilterWithdrawn(opts *bind.FilterOpts, wallet []common.Address, assetAddress []common.Address) (*YdexabiWithdrawnIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.FilterLogs(opts, "Withdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YdexabiWithdrawnIterator{contract: _Ydexabi.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x6960a1f64ecf9da0d1d1bcbfa3dd27f8c1c60de69b13faa28127dafa36c111e4.
//
// Solidity: event Withdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *YdexabiWithdrawn, wallet []common.Address, assetAddress []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Ydexabi.contract.WatchLogs(opts, "Withdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YdexabiWithdrawn)
				if err := _Ydexabi.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x6960a1f64ecf9da0d1d1bcbfa3dd27f8c1c60de69b13faa28127dafa36c111e4.
//
// Solidity: event Withdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Ydexabi *YdexabiFilterer) ParseWithdrawn(log types.Log) (*YdexabiWithdrawn, error) {
	event := new(YdexabiWithdrawn)
	if err := _Ydexabi.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
