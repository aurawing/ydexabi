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

// YDexCliMetaData contains all meta data concerning the YDexCli contract.
var YDexCliMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChainPropagationPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"assetSymbolIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"DispatcherChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"FeeWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestampInMs\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlockNumber\",\"type\":\"uint256\"}],\"name\":\"OrderNonceInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"}],\"name\":\"TokenSymbolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"baseAssetSymbolIndex\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"quoteAssetSymbolIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseAssetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"quoteAssetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"baseQuantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quoteQuantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"tradePriceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyOrderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellOrderHash\",\"type\":\"bytes32\"}],\"name\":\"TradeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"WalletExitCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"WalletExitWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlockNumber\",\"type\":\"uint256\"}],\"name\":\"WalletExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addTokenSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearWalletExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"confirmTokenRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantityInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"depositTokenByAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantityInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"depositTokenBySymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"signatureHashVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumEnums.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isQuantityInQuote\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"limitPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stopPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientOrderId\",\"type\":\"string\"},{\"internalType\":\"enumEnums.OrderTimeInForce\",\"name\":\"timeInForce\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSelfTradePrevention\",\"name\":\"selfTradePrevention\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"cancelAfter\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structStructs.Order\",\"name\":\"buy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"signatureHashVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"enumEnums.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isQuantityInQuote\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"limitPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stopPriceInPips\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"clientOrderId\",\"type\":\"string\"},{\"internalType\":\"enumEnums.OrderTimeInForce\",\"name\":\"timeInForce\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.OrderSelfTradePrevention\",\"name\":\"selfTradePrevention\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"cancelAfter\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structStructs.Order\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"baseAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quoteAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"baseAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"grossBaseQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"grossQuoteQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"netBaseQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"netQuoteQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"makerFeeAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerFeeAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"makerFeeQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"takerFeeQuantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"priceInPips\",\"type\":\"uint64\"},{\"internalType\":\"enumEnums.OrderSide\",\"name\":\"makerSide\",\"type\":\"uint8\"}],\"internalType\":\"structStructs.Trade\",\"name\":\"trade\",\"type\":\"tuple\"}],\"name\":\"executeTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"}],\"name\":\"invalidateOrderNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestampInMs\",\"type\":\"uint64\"}],\"name\":\"loadAssetBySymbol\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isConfirmed\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"confirmedTimestampInMs\",\"type\":\"uint64\"}],\"internalType\":\"structStructs.Asset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"loadBalanceInAssetUnitsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"}],\"name\":\"loadBalanceInAssetUnitsBySymbol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"loadBalanceInPipsByAddress\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"}],\"name\":\"loadBalanceInPipsBySymbol\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadFeeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"loadPartiallyFilledOrderQuantityInPips\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainPropagationPeriod\",\"type\":\"uint256\"}],\"name\":\"setChainPropagationPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCustodian\",\"type\":\"address\"}],\"name\":\"setCustodian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDispatcherWallet\",\"type\":\"address\"}],\"name\":\"setDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeWallet\",\"type\":\"address\"}],\"name\":\"setFeeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumEnums.WithdrawalType\",\"name\":\"withdrawalType\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasFeeInPips\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"autoDispatchEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structStructs.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"withdrawExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YDexCliABI is the input ABI used to generate the binding from.
// Deprecated: Use YDexCliMetaData.ABI instead.
var YDexCliABI = YDexCliMetaData.ABI

// YDexCli is an auto generated Go binding around an Ethereum contract.
type YDexCli struct {
	YDexCliCaller     // Read-only binding to the contract
	YDexCliTransactor // Write-only binding to the contract
	YDexCliFilterer   // Log filterer for contract events
}

// YDexCliCaller is an auto generated read-only Go binding around an Ethereum contract.
type YDexCliCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YDexCliTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YDexCliTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YDexCliFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YDexCliFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YDexCliSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YDexCliSession struct {
	Contract     *YDexCli          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YDexCliCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YDexCliCallerSession struct {
	Contract *YDexCliCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// YDexCliTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YDexCliTransactorSession struct {
	Contract     *YDexCliTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// YDexCliRaw is an auto generated low-level Go binding around an Ethereum contract.
type YDexCliRaw struct {
	Contract *YDexCli // Generic contract binding to access the raw methods on
}

// YDexCliCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YDexCliCallerRaw struct {
	Contract *YDexCliCaller // Generic read-only contract binding to access the raw methods on
}

// YDexCliTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YDexCliTransactorRaw struct {
	Contract *YDexCliTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYDexCli creates a new instance of YDexCli, bound to a specific deployed contract.
func NewYDexCli(address common.Address, backend bind.ContractBackend) (*YDexCli, error) {
	contract, err := bindYDexCli(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YDexCli{YDexCliCaller: YDexCliCaller{contract: contract}, YDexCliTransactor: YDexCliTransactor{contract: contract}, YDexCliFilterer: YDexCliFilterer{contract: contract}}, nil
}

// NewYDexCliCaller creates a new read-only instance of YDexCli, bound to a specific deployed contract.
func NewYDexCliCaller(address common.Address, caller bind.ContractCaller) (*YDexCliCaller, error) {
	contract, err := bindYDexCli(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YDexCliCaller{contract: contract}, nil
}

// NewYDexCliTransactor creates a new write-only instance of YDexCli, bound to a specific deployed contract.
func NewYDexCliTransactor(address common.Address, transactor bind.ContractTransactor) (*YDexCliTransactor, error) {
	contract, err := bindYDexCli(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YDexCliTransactor{contract: contract}, nil
}

// NewYDexCliFilterer creates a new log filterer instance of YDexCli, bound to a specific deployed contract.
func NewYDexCliFilterer(address common.Address, filterer bind.ContractFilterer) (*YDexCliFilterer, error) {
	contract, err := bindYDexCli(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YDexCliFilterer{contract: contract}, nil
}

// bindYDexCli binds a generic wrapper to an already deployed contract.
func bindYDexCli(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YDexCliABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YDexCli *YDexCliRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YDexCli.Contract.YDexCliCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YDexCli *YDexCliRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.Contract.YDexCliTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YDexCli *YDexCliRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YDexCli.Contract.YDexCliTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YDexCli *YDexCliCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YDexCli.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YDexCli *YDexCliTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YDexCli *YDexCliTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YDexCli.Contract.contract.Transact(opts, method, params...)
}

// LoadAssetBySymbol is a free data retrieval call binding the contract method 0x869af212.
//
// Solidity: function loadAssetBySymbol(string assetSymbol, uint64 timestampInMs) view returns((bool,address,string,uint8,bool,uint64))
func (_YDexCli *YDexCliCaller) LoadAssetBySymbol(opts *bind.CallOpts, assetSymbol string, timestampInMs uint64) (StructsAsset, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadAssetBySymbol", assetSymbol, timestampInMs)

	if err != nil {
		return *new(StructsAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsAsset)).(*StructsAsset)

	return out0, err

}

// LoadAssetBySymbol is a free data retrieval call binding the contract method 0x869af212.
//
// Solidity: function loadAssetBySymbol(string assetSymbol, uint64 timestampInMs) view returns((bool,address,string,uint8,bool,uint64))
func (_YDexCli *YDexCliSession) LoadAssetBySymbol(assetSymbol string, timestampInMs uint64) (StructsAsset, error) {
	return _YDexCli.Contract.LoadAssetBySymbol(&_YDexCli.CallOpts, assetSymbol, timestampInMs)
}

// LoadAssetBySymbol is a free data retrieval call binding the contract method 0x869af212.
//
// Solidity: function loadAssetBySymbol(string assetSymbol, uint64 timestampInMs) view returns((bool,address,string,uint8,bool,uint64))
func (_YDexCli *YDexCliCallerSession) LoadAssetBySymbol(assetSymbol string, timestampInMs uint64) (StructsAsset, error) {
	return _YDexCli.Contract.LoadAssetBySymbol(&_YDexCli.CallOpts, assetSymbol, timestampInMs)
}

// LoadBalanceInAssetUnitsByAddress is a free data retrieval call binding the contract method 0x13cfda2c.
//
// Solidity: function loadBalanceInAssetUnitsByAddress(address wallet, address assetAddress) view returns(uint256)
func (_YDexCli *YDexCliCaller) LoadBalanceInAssetUnitsByAddress(opts *bind.CallOpts, wallet common.Address, assetAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadBalanceInAssetUnitsByAddress", wallet, assetAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoadBalanceInAssetUnitsByAddress is a free data retrieval call binding the contract method 0x13cfda2c.
//
// Solidity: function loadBalanceInAssetUnitsByAddress(address wallet, address assetAddress) view returns(uint256)
func (_YDexCli *YDexCliSession) LoadBalanceInAssetUnitsByAddress(wallet common.Address, assetAddress common.Address) (*big.Int, error) {
	return _YDexCli.Contract.LoadBalanceInAssetUnitsByAddress(&_YDexCli.CallOpts, wallet, assetAddress)
}

// LoadBalanceInAssetUnitsByAddress is a free data retrieval call binding the contract method 0x13cfda2c.
//
// Solidity: function loadBalanceInAssetUnitsByAddress(address wallet, address assetAddress) view returns(uint256)
func (_YDexCli *YDexCliCallerSession) LoadBalanceInAssetUnitsByAddress(wallet common.Address, assetAddress common.Address) (*big.Int, error) {
	return _YDexCli.Contract.LoadBalanceInAssetUnitsByAddress(&_YDexCli.CallOpts, wallet, assetAddress)
}

// LoadBalanceInAssetUnitsBySymbol is a free data retrieval call binding the contract method 0x1abb5832.
//
// Solidity: function loadBalanceInAssetUnitsBySymbol(address wallet, string assetSymbol) view returns(uint256)
func (_YDexCli *YDexCliCaller) LoadBalanceInAssetUnitsBySymbol(opts *bind.CallOpts, wallet common.Address, assetSymbol string) (*big.Int, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadBalanceInAssetUnitsBySymbol", wallet, assetSymbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoadBalanceInAssetUnitsBySymbol is a free data retrieval call binding the contract method 0x1abb5832.
//
// Solidity: function loadBalanceInAssetUnitsBySymbol(address wallet, string assetSymbol) view returns(uint256)
func (_YDexCli *YDexCliSession) LoadBalanceInAssetUnitsBySymbol(wallet common.Address, assetSymbol string) (*big.Int, error) {
	return _YDexCli.Contract.LoadBalanceInAssetUnitsBySymbol(&_YDexCli.CallOpts, wallet, assetSymbol)
}

// LoadBalanceInAssetUnitsBySymbol is a free data retrieval call binding the contract method 0x1abb5832.
//
// Solidity: function loadBalanceInAssetUnitsBySymbol(address wallet, string assetSymbol) view returns(uint256)
func (_YDexCli *YDexCliCallerSession) LoadBalanceInAssetUnitsBySymbol(wallet common.Address, assetSymbol string) (*big.Int, error) {
	return _YDexCli.Contract.LoadBalanceInAssetUnitsBySymbol(&_YDexCli.CallOpts, wallet, assetSymbol)
}

// LoadBalanceInPipsByAddress is a free data retrieval call binding the contract method 0xdbb36535.
//
// Solidity: function loadBalanceInPipsByAddress(address wallet, address assetAddress) view returns(uint64)
func (_YDexCli *YDexCliCaller) LoadBalanceInPipsByAddress(opts *bind.CallOpts, wallet common.Address, assetAddress common.Address) (uint64, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadBalanceInPipsByAddress", wallet, assetAddress)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LoadBalanceInPipsByAddress is a free data retrieval call binding the contract method 0xdbb36535.
//
// Solidity: function loadBalanceInPipsByAddress(address wallet, address assetAddress) view returns(uint64)
func (_YDexCli *YDexCliSession) LoadBalanceInPipsByAddress(wallet common.Address, assetAddress common.Address) (uint64, error) {
	return _YDexCli.Contract.LoadBalanceInPipsByAddress(&_YDexCli.CallOpts, wallet, assetAddress)
}

// LoadBalanceInPipsByAddress is a free data retrieval call binding the contract method 0xdbb36535.
//
// Solidity: function loadBalanceInPipsByAddress(address wallet, address assetAddress) view returns(uint64)
func (_YDexCli *YDexCliCallerSession) LoadBalanceInPipsByAddress(wallet common.Address, assetAddress common.Address) (uint64, error) {
	return _YDexCli.Contract.LoadBalanceInPipsByAddress(&_YDexCli.CallOpts, wallet, assetAddress)
}

// LoadBalanceInPipsBySymbol is a free data retrieval call binding the contract method 0xef3b9d4a.
//
// Solidity: function loadBalanceInPipsBySymbol(address wallet, string assetSymbol) view returns(uint64)
func (_YDexCli *YDexCliCaller) LoadBalanceInPipsBySymbol(opts *bind.CallOpts, wallet common.Address, assetSymbol string) (uint64, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadBalanceInPipsBySymbol", wallet, assetSymbol)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LoadBalanceInPipsBySymbol is a free data retrieval call binding the contract method 0xef3b9d4a.
//
// Solidity: function loadBalanceInPipsBySymbol(address wallet, string assetSymbol) view returns(uint64)
func (_YDexCli *YDexCliSession) LoadBalanceInPipsBySymbol(wallet common.Address, assetSymbol string) (uint64, error) {
	return _YDexCli.Contract.LoadBalanceInPipsBySymbol(&_YDexCli.CallOpts, wallet, assetSymbol)
}

// LoadBalanceInPipsBySymbol is a free data retrieval call binding the contract method 0xef3b9d4a.
//
// Solidity: function loadBalanceInPipsBySymbol(address wallet, string assetSymbol) view returns(uint64)
func (_YDexCli *YDexCliCallerSession) LoadBalanceInPipsBySymbol(wallet common.Address, assetSymbol string) (uint64, error) {
	return _YDexCli.Contract.LoadBalanceInPipsBySymbol(&_YDexCli.CallOpts, wallet, assetSymbol)
}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_YDexCli *YDexCliCaller) LoadFeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadFeeWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_YDexCli *YDexCliSession) LoadFeeWallet() (common.Address, error) {
	return _YDexCli.Contract.LoadFeeWallet(&_YDexCli.CallOpts)
}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_YDexCli *YDexCliCallerSession) LoadFeeWallet() (common.Address, error) {
	return _YDexCli.Contract.LoadFeeWallet(&_YDexCli.CallOpts)
}

// LoadPartiallyFilledOrderQuantityInPips is a free data retrieval call binding the contract method 0xae0e969e.
//
// Solidity: function loadPartiallyFilledOrderQuantityInPips(bytes32 orderHash) view returns(uint64)
func (_YDexCli *YDexCliCaller) LoadPartiallyFilledOrderQuantityInPips(opts *bind.CallOpts, orderHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _YDexCli.contract.Call(opts, &out, "loadPartiallyFilledOrderQuantityInPips", orderHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LoadPartiallyFilledOrderQuantityInPips is a free data retrieval call binding the contract method 0xae0e969e.
//
// Solidity: function loadPartiallyFilledOrderQuantityInPips(bytes32 orderHash) view returns(uint64)
func (_YDexCli *YDexCliSession) LoadPartiallyFilledOrderQuantityInPips(orderHash [32]byte) (uint64, error) {
	return _YDexCli.Contract.LoadPartiallyFilledOrderQuantityInPips(&_YDexCli.CallOpts, orderHash)
}

// LoadPartiallyFilledOrderQuantityInPips is a free data retrieval call binding the contract method 0xae0e969e.
//
// Solidity: function loadPartiallyFilledOrderQuantityInPips(bytes32 orderHash) view returns(uint64)
func (_YDexCli *YDexCliCallerSession) LoadPartiallyFilledOrderQuantityInPips(orderHash [32]byte) (uint64, error) {
	return _YDexCli.Contract.LoadPartiallyFilledOrderQuantityInPips(&_YDexCli.CallOpts, orderHash)
}

// AddTokenSymbol is a paid mutator transaction binding the contract method 0x72e8f08d.
//
// Solidity: function addTokenSymbol(address tokenAddress, string symbol) returns()
func (_YDexCli *YDexCliTransactor) AddTokenSymbol(opts *bind.TransactOpts, tokenAddress common.Address, symbol string) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "addTokenSymbol", tokenAddress, symbol)
}

// AddTokenSymbol is a paid mutator transaction binding the contract method 0x72e8f08d.
//
// Solidity: function addTokenSymbol(address tokenAddress, string symbol) returns()
func (_YDexCli *YDexCliSession) AddTokenSymbol(tokenAddress common.Address, symbol string) (*types.Transaction, error) {
	return _YDexCli.Contract.AddTokenSymbol(&_YDexCli.TransactOpts, tokenAddress, symbol)
}

// AddTokenSymbol is a paid mutator transaction binding the contract method 0x72e8f08d.
//
// Solidity: function addTokenSymbol(address tokenAddress, string symbol) returns()
func (_YDexCli *YDexCliTransactorSession) AddTokenSymbol(tokenAddress common.Address, symbol string) (*types.Transaction, error) {
	return _YDexCli.Contract.AddTokenSymbol(&_YDexCli.TransactOpts, tokenAddress, symbol)
}

// ClearWalletExit is a paid mutator transaction binding the contract method 0x4a284ef9.
//
// Solidity: function clearWalletExit() returns()
func (_YDexCli *YDexCliTransactor) ClearWalletExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "clearWalletExit")
}

// ClearWalletExit is a paid mutator transaction binding the contract method 0x4a284ef9.
//
// Solidity: function clearWalletExit() returns()
func (_YDexCli *YDexCliSession) ClearWalletExit() (*types.Transaction, error) {
	return _YDexCli.Contract.ClearWalletExit(&_YDexCli.TransactOpts)
}

// ClearWalletExit is a paid mutator transaction binding the contract method 0x4a284ef9.
//
// Solidity: function clearWalletExit() returns()
func (_YDexCli *YDexCliTransactorSession) ClearWalletExit() (*types.Transaction, error) {
	return _YDexCli.Contract.ClearWalletExit(&_YDexCli.TransactOpts)
}

// ConfirmTokenRegistration is a paid mutator transaction binding the contract method 0x0226b70e.
//
// Solidity: function confirmTokenRegistration(address tokenAddress, string symbol, uint8 decimals) returns()
func (_YDexCli *YDexCliTransactor) ConfirmTokenRegistration(opts *bind.TransactOpts, tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "confirmTokenRegistration", tokenAddress, symbol, decimals)
}

// ConfirmTokenRegistration is a paid mutator transaction binding the contract method 0x0226b70e.
//
// Solidity: function confirmTokenRegistration(address tokenAddress, string symbol, uint8 decimals) returns()
func (_YDexCli *YDexCliSession) ConfirmTokenRegistration(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _YDexCli.Contract.ConfirmTokenRegistration(&_YDexCli.TransactOpts, tokenAddress, symbol, decimals)
}

// ConfirmTokenRegistration is a paid mutator transaction binding the contract method 0x0226b70e.
//
// Solidity: function confirmTokenRegistration(address tokenAddress, string symbol, uint8 decimals) returns()
func (_YDexCli *YDexCliTransactorSession) ConfirmTokenRegistration(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _YDexCli.Contract.ConfirmTokenRegistration(&_YDexCli.TransactOpts, tokenAddress, symbol, decimals)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x42220f34.
//
// Solidity: function depositBNB() payable returns()
func (_YDexCli *YDexCliTransactor) DepositBNB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "depositBNB")
}

// DepositBNB is a paid mutator transaction binding the contract method 0x42220f34.
//
// Solidity: function depositBNB() payable returns()
func (_YDexCli *YDexCliSession) DepositBNB() (*types.Transaction, error) {
	return _YDexCli.Contract.DepositBNB(&_YDexCli.TransactOpts)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x42220f34.
//
// Solidity: function depositBNB() payable returns()
func (_YDexCli *YDexCliTransactorSession) DepositBNB() (*types.Transaction, error) {
	return _YDexCli.Contract.DepositBNB(&_YDexCli.TransactOpts)
}

// DepositTokenByAddress is a paid mutator transaction binding the contract method 0xdcc63490.
//
// Solidity: function depositTokenByAddress(address tokenAddress, uint256 quantityInAssetUnits) returns()
func (_YDexCli *YDexCliTransactor) DepositTokenByAddress(opts *bind.TransactOpts, tokenAddress common.Address, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "depositTokenByAddress", tokenAddress, quantityInAssetUnits)
}

// DepositTokenByAddress is a paid mutator transaction binding the contract method 0xdcc63490.
//
// Solidity: function depositTokenByAddress(address tokenAddress, uint256 quantityInAssetUnits) returns()
func (_YDexCli *YDexCliSession) DepositTokenByAddress(tokenAddress common.Address, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.DepositTokenByAddress(&_YDexCli.TransactOpts, tokenAddress, quantityInAssetUnits)
}

// DepositTokenByAddress is a paid mutator transaction binding the contract method 0xdcc63490.
//
// Solidity: function depositTokenByAddress(address tokenAddress, uint256 quantityInAssetUnits) returns()
func (_YDexCli *YDexCliTransactorSession) DepositTokenByAddress(tokenAddress common.Address, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.DepositTokenByAddress(&_YDexCli.TransactOpts, tokenAddress, quantityInAssetUnits)
}

// DepositTokenBySymbol is a paid mutator transaction binding the contract method 0x41968182.
//
// Solidity: function depositTokenBySymbol(string assetSymbol, uint256 quantityInAssetUnits) returns()
func (_YDexCli *YDexCliTransactor) DepositTokenBySymbol(opts *bind.TransactOpts, assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "depositTokenBySymbol", assetSymbol, quantityInAssetUnits)
}

// DepositTokenBySymbol is a paid mutator transaction binding the contract method 0x41968182.
//
// Solidity: function depositTokenBySymbol(string assetSymbol, uint256 quantityInAssetUnits) returns()
func (_YDexCli *YDexCliSession) DepositTokenBySymbol(assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.DepositTokenBySymbol(&_YDexCli.TransactOpts, assetSymbol, quantityInAssetUnits)
}

// DepositTokenBySymbol is a paid mutator transaction binding the contract method 0x41968182.
//
// Solidity: function depositTokenBySymbol(string assetSymbol, uint256 quantityInAssetUnits) returns()
func (_YDexCli *YDexCliTransactorSession) DepositTokenBySymbol(assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.DepositTokenBySymbol(&_YDexCli.TransactOpts, assetSymbol, quantityInAssetUnits)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x65fce5b7.
//
// Solidity: function executeTrade((uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) buy, (uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) sell, (string,string,address,address,uint64,uint64,uint64,uint64,address,address,uint64,uint64,uint64,uint8) trade) returns()
func (_YDexCli *YDexCliTransactor) ExecuteTrade(opts *bind.TransactOpts, buy StructsOrder, sell StructsOrder, trade StructsTrade) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "executeTrade", buy, sell, trade)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x65fce5b7.
//
// Solidity: function executeTrade((uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) buy, (uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) sell, (string,string,address,address,uint64,uint64,uint64,uint64,address,address,uint64,uint64,uint64,uint8) trade) returns()
func (_YDexCli *YDexCliSession) ExecuteTrade(buy StructsOrder, sell StructsOrder, trade StructsTrade) (*types.Transaction, error) {
	return _YDexCli.Contract.ExecuteTrade(&_YDexCli.TransactOpts, buy, sell, trade)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x65fce5b7.
//
// Solidity: function executeTrade((uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) buy, (uint8,uint128,address,uint8,uint8,uint64,bool,uint64,uint64,string,uint8,uint8,uint64,bytes) sell, (string,string,address,address,uint64,uint64,uint64,uint64,address,address,uint64,uint64,uint64,uint8) trade) returns()
func (_YDexCli *YDexCliTransactorSession) ExecuteTrade(buy StructsOrder, sell StructsOrder, trade StructsTrade) (*types.Transaction, error) {
	return _YDexCli.Contract.ExecuteTrade(&_YDexCli.TransactOpts, buy, sell, trade)
}

// ExitWallet is a paid mutator transaction binding the contract method 0xeb5068f2.
//
// Solidity: function exitWallet() returns()
func (_YDexCli *YDexCliTransactor) ExitWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "exitWallet")
}

// ExitWallet is a paid mutator transaction binding the contract method 0xeb5068f2.
//
// Solidity: function exitWallet() returns()
func (_YDexCli *YDexCliSession) ExitWallet() (*types.Transaction, error) {
	return _YDexCli.Contract.ExitWallet(&_YDexCli.TransactOpts)
}

// ExitWallet is a paid mutator transaction binding the contract method 0xeb5068f2.
//
// Solidity: function exitWallet() returns()
func (_YDexCli *YDexCliTransactorSession) ExitWallet() (*types.Transaction, error) {
	return _YDexCli.Contract.ExitWallet(&_YDexCli.TransactOpts)
}

// InvalidateOrderNonce is a paid mutator transaction binding the contract method 0xd7a6aec7.
//
// Solidity: function invalidateOrderNonce(uint128 nonce) returns()
func (_YDexCli *YDexCliTransactor) InvalidateOrderNonce(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "invalidateOrderNonce", nonce)
}

// InvalidateOrderNonce is a paid mutator transaction binding the contract method 0xd7a6aec7.
//
// Solidity: function invalidateOrderNonce(uint128 nonce) returns()
func (_YDexCli *YDexCliSession) InvalidateOrderNonce(nonce *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.InvalidateOrderNonce(&_YDexCli.TransactOpts, nonce)
}

// InvalidateOrderNonce is a paid mutator transaction binding the contract method 0xd7a6aec7.
//
// Solidity: function invalidateOrderNonce(uint128 nonce) returns()
func (_YDexCli *YDexCliTransactorSession) InvalidateOrderNonce(nonce *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.InvalidateOrderNonce(&_YDexCli.TransactOpts, nonce)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x457aa3c6.
//
// Solidity: function registerToken(address tokenAddress, string symbol, uint8 decimals) returns()
func (_YDexCli *YDexCliTransactor) RegisterToken(opts *bind.TransactOpts, tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "registerToken", tokenAddress, symbol, decimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x457aa3c6.
//
// Solidity: function registerToken(address tokenAddress, string symbol, uint8 decimals) returns()
func (_YDexCli *YDexCliSession) RegisterToken(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _YDexCli.Contract.RegisterToken(&_YDexCli.TransactOpts, tokenAddress, symbol, decimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x457aa3c6.
//
// Solidity: function registerToken(address tokenAddress, string symbol, uint8 decimals) returns()
func (_YDexCli *YDexCliTransactorSession) RegisterToken(tokenAddress common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _YDexCli.Contract.RegisterToken(&_YDexCli.TransactOpts, tokenAddress, symbol, decimals)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_YDexCli *YDexCliTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_YDexCli *YDexCliSession) RemoveAdmin() (*types.Transaction, error) {
	return _YDexCli.Contract.RemoveAdmin(&_YDexCli.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_YDexCli *YDexCliTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _YDexCli.Contract.RemoveAdmin(&_YDexCli.TransactOpts)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x0c187a72.
//
// Solidity: function removeDispatcher() returns()
func (_YDexCli *YDexCliTransactor) RemoveDispatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "removeDispatcher")
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x0c187a72.
//
// Solidity: function removeDispatcher() returns()
func (_YDexCli *YDexCliSession) RemoveDispatcher() (*types.Transaction, error) {
	return _YDexCli.Contract.RemoveDispatcher(&_YDexCli.TransactOpts)
}

// RemoveDispatcher is a paid mutator transaction binding the contract method 0x0c187a72.
//
// Solidity: function removeDispatcher() returns()
func (_YDexCli *YDexCliTransactorSession) RemoveDispatcher() (*types.Transaction, error) {
	return _YDexCli.Contract.RemoveDispatcher(&_YDexCli.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_YDexCli *YDexCliTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_YDexCli *YDexCliSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetAdmin(&_YDexCli.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_YDexCli *YDexCliTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetAdmin(&_YDexCli.TransactOpts, newAdmin)
}

// SetChainPropagationPeriod is a paid mutator transaction binding the contract method 0x20e6a8e3.
//
// Solidity: function setChainPropagationPeriod(uint256 newChainPropagationPeriod) returns()
func (_YDexCli *YDexCliTransactor) SetChainPropagationPeriod(opts *bind.TransactOpts, newChainPropagationPeriod *big.Int) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "setChainPropagationPeriod", newChainPropagationPeriod)
}

// SetChainPropagationPeriod is a paid mutator transaction binding the contract method 0x20e6a8e3.
//
// Solidity: function setChainPropagationPeriod(uint256 newChainPropagationPeriod) returns()
func (_YDexCli *YDexCliSession) SetChainPropagationPeriod(newChainPropagationPeriod *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.SetChainPropagationPeriod(&_YDexCli.TransactOpts, newChainPropagationPeriod)
}

// SetChainPropagationPeriod is a paid mutator transaction binding the contract method 0x20e6a8e3.
//
// Solidity: function setChainPropagationPeriod(uint256 newChainPropagationPeriod) returns()
func (_YDexCli *YDexCliTransactorSession) SetChainPropagationPeriod(newChainPropagationPeriod *big.Int) (*types.Transaction, error) {
	return _YDexCli.Contract.SetChainPropagationPeriod(&_YDexCli.TransactOpts, newChainPropagationPeriod)
}

// SetCustodian is a paid mutator transaction binding the contract method 0x403f3731.
//
// Solidity: function setCustodian(address newCustodian) returns()
func (_YDexCli *YDexCliTransactor) SetCustodian(opts *bind.TransactOpts, newCustodian common.Address) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "setCustodian", newCustodian)
}

// SetCustodian is a paid mutator transaction binding the contract method 0x403f3731.
//
// Solidity: function setCustodian(address newCustodian) returns()
func (_YDexCli *YDexCliSession) SetCustodian(newCustodian common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetCustodian(&_YDexCli.TransactOpts, newCustodian)
}

// SetCustodian is a paid mutator transaction binding the contract method 0x403f3731.
//
// Solidity: function setCustodian(address newCustodian) returns()
func (_YDexCli *YDexCliTransactorSession) SetCustodian(newCustodian common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetCustodian(&_YDexCli.TransactOpts, newCustodian)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address newDispatcherWallet) returns()
func (_YDexCli *YDexCliTransactor) SetDispatcher(opts *bind.TransactOpts, newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "setDispatcher", newDispatcherWallet)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address newDispatcherWallet) returns()
func (_YDexCli *YDexCliSession) SetDispatcher(newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetDispatcher(&_YDexCli.TransactOpts, newDispatcherWallet)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address newDispatcherWallet) returns()
func (_YDexCli *YDexCliTransactorSession) SetDispatcher(newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetDispatcher(&_YDexCli.TransactOpts, newDispatcherWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_YDexCli *YDexCliTransactor) SetFeeWallet(opts *bind.TransactOpts, newFeeWallet common.Address) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "setFeeWallet", newFeeWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_YDexCli *YDexCliSession) SetFeeWallet(newFeeWallet common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetFeeWallet(&_YDexCli.TransactOpts, newFeeWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_YDexCli *YDexCliTransactorSession) SetFeeWallet(newFeeWallet common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.SetFeeWallet(&_YDexCli.TransactOpts, newFeeWallet)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6bb509f2.
//
// Solidity: function withdraw((uint8,uint128,address,string,address,uint64,uint64,bool,bytes) withdrawal) returns()
func (_YDexCli *YDexCliTransactor) Withdraw(opts *bind.TransactOpts, withdrawal StructsWithdrawal) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "withdraw", withdrawal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6bb509f2.
//
// Solidity: function withdraw((uint8,uint128,address,string,address,uint64,uint64,bool,bytes) withdrawal) returns()
func (_YDexCli *YDexCliSession) Withdraw(withdrawal StructsWithdrawal) (*types.Transaction, error) {
	return _YDexCli.Contract.Withdraw(&_YDexCli.TransactOpts, withdrawal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6bb509f2.
//
// Solidity: function withdraw((uint8,uint128,address,string,address,uint64,uint64,bool,bytes) withdrawal) returns()
func (_YDexCli *YDexCliTransactorSession) Withdraw(withdrawal StructsWithdrawal) (*types.Transaction, error) {
	return _YDexCli.Contract.Withdraw(&_YDexCli.TransactOpts, withdrawal)
}

// WithdrawExit is a paid mutator transaction binding the contract method 0x985c4af5.
//
// Solidity: function withdrawExit(address assetAddress) returns()
func (_YDexCli *YDexCliTransactor) WithdrawExit(opts *bind.TransactOpts, assetAddress common.Address) (*types.Transaction, error) {
	return _YDexCli.contract.Transact(opts, "withdrawExit", assetAddress)
}

// WithdrawExit is a paid mutator transaction binding the contract method 0x985c4af5.
//
// Solidity: function withdrawExit(address assetAddress) returns()
func (_YDexCli *YDexCliSession) WithdrawExit(assetAddress common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.WithdrawExit(&_YDexCli.TransactOpts, assetAddress)
}

// WithdrawExit is a paid mutator transaction binding the contract method 0x985c4af5.
//
// Solidity: function withdrawExit(address assetAddress) returns()
func (_YDexCli *YDexCliTransactorSession) WithdrawExit(assetAddress common.Address) (*types.Transaction, error) {
	return _YDexCli.Contract.WithdrawExit(&_YDexCli.TransactOpts, assetAddress)
}

// YDexCliChainPropagationPeriodChangedIterator is returned from FilterChainPropagationPeriodChanged and is used to iterate over the raw logs and unpacked data for ChainPropagationPeriodChanged events raised by the YDexCli contract.
type YDexCliChainPropagationPeriodChangedIterator struct {
	Event *YDexCliChainPropagationPeriodChanged // Event containing the contract specifics and raw log

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
func (it *YDexCliChainPropagationPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliChainPropagationPeriodChanged)
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
		it.Event = new(YDexCliChainPropagationPeriodChanged)
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
func (it *YDexCliChainPropagationPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliChainPropagationPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliChainPropagationPeriodChanged represents a ChainPropagationPeriodChanged event raised by the YDexCli contract.
type YDexCliChainPropagationPeriodChanged struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPropagationPeriodChanged is a free log retrieval operation binding the contract event 0x9a22227d6c0251a79ef8b846202ddcbe9d682ee5482e84abeec6dda096398a6f.
//
// Solidity: event ChainPropagationPeriodChanged(uint256 previousValue, uint256 newValue)
func (_YDexCli *YDexCliFilterer) FilterChainPropagationPeriodChanged(opts *bind.FilterOpts) (*YDexCliChainPropagationPeriodChangedIterator, error) {

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "ChainPropagationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &YDexCliChainPropagationPeriodChangedIterator{contract: _YDexCli.contract, event: "ChainPropagationPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchChainPropagationPeriodChanged is a free log subscription operation binding the contract event 0x9a22227d6c0251a79ef8b846202ddcbe9d682ee5482e84abeec6dda096398a6f.
//
// Solidity: event ChainPropagationPeriodChanged(uint256 previousValue, uint256 newValue)
func (_YDexCli *YDexCliFilterer) WatchChainPropagationPeriodChanged(opts *bind.WatchOpts, sink chan<- *YDexCliChainPropagationPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "ChainPropagationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliChainPropagationPeriodChanged)
				if err := _YDexCli.contract.UnpackLog(event, "ChainPropagationPeriodChanged", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseChainPropagationPeriodChanged(log types.Log) (*YDexCliChainPropagationPeriodChanged, error) {
	event := new(YDexCliChainPropagationPeriodChanged)
	if err := _YDexCli.contract.UnpackLog(event, "ChainPropagationPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the YDexCli contract.
type YDexCliDepositedIterator struct {
	Event *YDexCliDeposited // Event containing the contract specifics and raw log

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
func (it *YDexCliDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliDeposited)
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
		it.Event = new(YDexCliDeposited)
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
func (it *YDexCliDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliDeposited represents a Deposited event raised by the YDexCli contract.
type YDexCliDeposited struct {
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
func (_YDexCli *YDexCliFilterer) FilterDeposited(opts *bind.FilterOpts, wallet []common.Address, assetAddress []common.Address, assetSymbolIndex []string) (*YDexCliDepositedIterator, error) {

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

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "Deposited", walletRule, assetAddressRule, assetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliDepositedIterator{contract: _YDexCli.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc2813a84168de59cf710030ba2acb13031b567b8ba76f70e8ab0782aa9469a7a.
//
// Solidity: event Deposited(uint64 index, address indexed wallet, address indexed assetAddress, string indexed assetSymbolIndex, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_YDexCli *YDexCliFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *YDexCliDeposited, wallet []common.Address, assetAddress []common.Address, assetSymbolIndex []string) (event.Subscription, error) {

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

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "Deposited", walletRule, assetAddressRule, assetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliDeposited)
				if err := _YDexCli.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseDeposited(log types.Log) (*YDexCliDeposited, error) {
	event := new(YDexCliDeposited)
	if err := _YDexCli.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliDispatcherChangedIterator is returned from FilterDispatcherChanged and is used to iterate over the raw logs and unpacked data for DispatcherChanged events raised by the YDexCli contract.
type YDexCliDispatcherChangedIterator struct {
	Event *YDexCliDispatcherChanged // Event containing the contract specifics and raw log

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
func (it *YDexCliDispatcherChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliDispatcherChanged)
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
		it.Event = new(YDexCliDispatcherChanged)
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
func (it *YDexCliDispatcherChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliDispatcherChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliDispatcherChanged represents a DispatcherChanged event raised by the YDexCli contract.
type YDexCliDispatcherChanged struct {
	PreviousValue common.Address
	NewValue      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDispatcherChanged is a free log retrieval operation binding the contract event 0xddb86df2d5262dd7b44067db5962cc7875e9db409cb21c88adfe3c5760315e39.
//
// Solidity: event DispatcherChanged(address previousValue, address newValue)
func (_YDexCli *YDexCliFilterer) FilterDispatcherChanged(opts *bind.FilterOpts) (*YDexCliDispatcherChangedIterator, error) {

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "DispatcherChanged")
	if err != nil {
		return nil, err
	}
	return &YDexCliDispatcherChangedIterator{contract: _YDexCli.contract, event: "DispatcherChanged", logs: logs, sub: sub}, nil
}

// WatchDispatcherChanged is a free log subscription operation binding the contract event 0xddb86df2d5262dd7b44067db5962cc7875e9db409cb21c88adfe3c5760315e39.
//
// Solidity: event DispatcherChanged(address previousValue, address newValue)
func (_YDexCli *YDexCliFilterer) WatchDispatcherChanged(opts *bind.WatchOpts, sink chan<- *YDexCliDispatcherChanged) (event.Subscription, error) {

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "DispatcherChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliDispatcherChanged)
				if err := _YDexCli.contract.UnpackLog(event, "DispatcherChanged", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseDispatcherChanged(log types.Log) (*YDexCliDispatcherChanged, error) {
	event := new(YDexCliDispatcherChanged)
	if err := _YDexCli.contract.UnpackLog(event, "DispatcherChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliFeeWalletChangedIterator is returned from FilterFeeWalletChanged and is used to iterate over the raw logs and unpacked data for FeeWalletChanged events raised by the YDexCli contract.
type YDexCliFeeWalletChangedIterator struct {
	Event *YDexCliFeeWalletChanged // Event containing the contract specifics and raw log

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
func (it *YDexCliFeeWalletChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliFeeWalletChanged)
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
		it.Event = new(YDexCliFeeWalletChanged)
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
func (it *YDexCliFeeWalletChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliFeeWalletChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliFeeWalletChanged represents a FeeWalletChanged event raised by the YDexCli contract.
type YDexCliFeeWalletChanged struct {
	PreviousValue common.Address
	NewValue      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeWalletChanged is a free log retrieval operation binding the contract event 0x9f4f5dce3c4d197b5d7496cb96e25f0a89809167195964b0daa3ef5fed63c00a.
//
// Solidity: event FeeWalletChanged(address previousValue, address newValue)
func (_YDexCli *YDexCliFilterer) FilterFeeWalletChanged(opts *bind.FilterOpts) (*YDexCliFeeWalletChangedIterator, error) {

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "FeeWalletChanged")
	if err != nil {
		return nil, err
	}
	return &YDexCliFeeWalletChangedIterator{contract: _YDexCli.contract, event: "FeeWalletChanged", logs: logs, sub: sub}, nil
}

// WatchFeeWalletChanged is a free log subscription operation binding the contract event 0x9f4f5dce3c4d197b5d7496cb96e25f0a89809167195964b0daa3ef5fed63c00a.
//
// Solidity: event FeeWalletChanged(address previousValue, address newValue)
func (_YDexCli *YDexCliFilterer) WatchFeeWalletChanged(opts *bind.WatchOpts, sink chan<- *YDexCliFeeWalletChanged) (event.Subscription, error) {

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "FeeWalletChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliFeeWalletChanged)
				if err := _YDexCli.contract.UnpackLog(event, "FeeWalletChanged", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseFeeWalletChanged(log types.Log) (*YDexCliFeeWalletChanged, error) {
	event := new(YDexCliFeeWalletChanged)
	if err := _YDexCli.contract.UnpackLog(event, "FeeWalletChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliOrderNonceInvalidatedIterator is returned from FilterOrderNonceInvalidated and is used to iterate over the raw logs and unpacked data for OrderNonceInvalidated events raised by the YDexCli contract.
type YDexCliOrderNonceInvalidatedIterator struct {
	Event *YDexCliOrderNonceInvalidated // Event containing the contract specifics and raw log

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
func (it *YDexCliOrderNonceInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliOrderNonceInvalidated)
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
		it.Event = new(YDexCliOrderNonceInvalidated)
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
func (it *YDexCliOrderNonceInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliOrderNonceInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliOrderNonceInvalidated represents a OrderNonceInvalidated event raised by the YDexCli contract.
type YDexCliOrderNonceInvalidated struct {
	Wallet               common.Address
	Nonce                *big.Int
	TimestampInMs        *big.Int
	EffectiveBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOrderNonceInvalidated is a free log retrieval operation binding the contract event 0x10cf19671b43c88b1f02d4e94932d7ffaa89c7278bc5b8868fa7b7676210809b.
//
// Solidity: event OrderNonceInvalidated(address indexed wallet, uint128 nonce, uint128 timestampInMs, uint256 effectiveBlockNumber)
func (_YDexCli *YDexCliFilterer) FilterOrderNonceInvalidated(opts *bind.FilterOpts, wallet []common.Address) (*YDexCliOrderNonceInvalidatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "OrderNonceInvalidated", walletRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliOrderNonceInvalidatedIterator{contract: _YDexCli.contract, event: "OrderNonceInvalidated", logs: logs, sub: sub}, nil
}

// WatchOrderNonceInvalidated is a free log subscription operation binding the contract event 0x10cf19671b43c88b1f02d4e94932d7ffaa89c7278bc5b8868fa7b7676210809b.
//
// Solidity: event OrderNonceInvalidated(address indexed wallet, uint128 nonce, uint128 timestampInMs, uint256 effectiveBlockNumber)
func (_YDexCli *YDexCliFilterer) WatchOrderNonceInvalidated(opts *bind.WatchOpts, sink chan<- *YDexCliOrderNonceInvalidated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "OrderNonceInvalidated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliOrderNonceInvalidated)
				if err := _YDexCli.contract.UnpackLog(event, "OrderNonceInvalidated", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseOrderNonceInvalidated(log types.Log) (*YDexCliOrderNonceInvalidated, error) {
	event := new(YDexCliOrderNonceInvalidated)
	if err := _YDexCli.contract.UnpackLog(event, "OrderNonceInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the YDexCli contract.
type YDexCliTokenRegisteredIterator struct {
	Event *YDexCliTokenRegistered // Event containing the contract specifics and raw log

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
func (it *YDexCliTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliTokenRegistered)
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
		it.Event = new(YDexCliTokenRegistered)
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
func (it *YDexCliTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliTokenRegistered represents a TokenRegistered event raised by the YDexCli contract.
type YDexCliTokenRegistered struct {
	AssetAddress common.Address
	AssetSymbol  string
	Decimals     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xcf66acda7673c71cd9458c59067eb30a473eeaf2fa05d008acf47616cbbbe01c.
//
// Solidity: event TokenRegistered(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_YDexCli *YDexCliFilterer) FilterTokenRegistered(opts *bind.FilterOpts, assetAddress []common.Address) (*YDexCliTokenRegisteredIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "TokenRegistered", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliTokenRegisteredIterator{contract: _YDexCli.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xcf66acda7673c71cd9458c59067eb30a473eeaf2fa05d008acf47616cbbbe01c.
//
// Solidity: event TokenRegistered(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_YDexCli *YDexCliFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *YDexCliTokenRegistered, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "TokenRegistered", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliTokenRegistered)
				if err := _YDexCli.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseTokenRegistered(log types.Log) (*YDexCliTokenRegistered, error) {
	event := new(YDexCliTokenRegistered)
	if err := _YDexCli.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliTokenRegistrationConfirmedIterator is returned from FilterTokenRegistrationConfirmed and is used to iterate over the raw logs and unpacked data for TokenRegistrationConfirmed events raised by the YDexCli contract.
type YDexCliTokenRegistrationConfirmedIterator struct {
	Event *YDexCliTokenRegistrationConfirmed // Event containing the contract specifics and raw log

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
func (it *YDexCliTokenRegistrationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliTokenRegistrationConfirmed)
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
		it.Event = new(YDexCliTokenRegistrationConfirmed)
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
func (it *YDexCliTokenRegistrationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliTokenRegistrationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliTokenRegistrationConfirmed represents a TokenRegistrationConfirmed event raised by the YDexCli contract.
type YDexCliTokenRegistrationConfirmed struct {
	AssetAddress common.Address
	AssetSymbol  string
	Decimals     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistrationConfirmed is a free log retrieval operation binding the contract event 0xa8968236dc4afb2c189afeb69c91b02445e6bd09f0ffe27256b62838edc99edb.
//
// Solidity: event TokenRegistrationConfirmed(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_YDexCli *YDexCliFilterer) FilterTokenRegistrationConfirmed(opts *bind.FilterOpts, assetAddress []common.Address) (*YDexCliTokenRegistrationConfirmedIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "TokenRegistrationConfirmed", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliTokenRegistrationConfirmedIterator{contract: _YDexCli.contract, event: "TokenRegistrationConfirmed", logs: logs, sub: sub}, nil
}

// WatchTokenRegistrationConfirmed is a free log subscription operation binding the contract event 0xa8968236dc4afb2c189afeb69c91b02445e6bd09f0ffe27256b62838edc99edb.
//
// Solidity: event TokenRegistrationConfirmed(address indexed assetAddress, string assetSymbol, uint8 decimals)
func (_YDexCli *YDexCliFilterer) WatchTokenRegistrationConfirmed(opts *bind.WatchOpts, sink chan<- *YDexCliTokenRegistrationConfirmed, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "TokenRegistrationConfirmed", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliTokenRegistrationConfirmed)
				if err := _YDexCli.contract.UnpackLog(event, "TokenRegistrationConfirmed", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseTokenRegistrationConfirmed(log types.Log) (*YDexCliTokenRegistrationConfirmed, error) {
	event := new(YDexCliTokenRegistrationConfirmed)
	if err := _YDexCli.contract.UnpackLog(event, "TokenRegistrationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliTokenSymbolAddedIterator is returned from FilterTokenSymbolAdded and is used to iterate over the raw logs and unpacked data for TokenSymbolAdded events raised by the YDexCli contract.
type YDexCliTokenSymbolAddedIterator struct {
	Event *YDexCliTokenSymbolAdded // Event containing the contract specifics and raw log

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
func (it *YDexCliTokenSymbolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliTokenSymbolAdded)
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
		it.Event = new(YDexCliTokenSymbolAdded)
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
func (it *YDexCliTokenSymbolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliTokenSymbolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliTokenSymbolAdded represents a TokenSymbolAdded event raised by the YDexCli contract.
type YDexCliTokenSymbolAdded struct {
	AssetAddress common.Address
	AssetSymbol  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenSymbolAdded is a free log retrieval operation binding the contract event 0x7f63a63f4a2ea318da0b031794d0b5a3183a1b2de54e053ceb17cabdba031735.
//
// Solidity: event TokenSymbolAdded(address indexed assetAddress, string assetSymbol)
func (_YDexCli *YDexCliFilterer) FilterTokenSymbolAdded(opts *bind.FilterOpts, assetAddress []common.Address) (*YDexCliTokenSymbolAddedIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "TokenSymbolAdded", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliTokenSymbolAddedIterator{contract: _YDexCli.contract, event: "TokenSymbolAdded", logs: logs, sub: sub}, nil
}

// WatchTokenSymbolAdded is a free log subscription operation binding the contract event 0x7f63a63f4a2ea318da0b031794d0b5a3183a1b2de54e053ceb17cabdba031735.
//
// Solidity: event TokenSymbolAdded(address indexed assetAddress, string assetSymbol)
func (_YDexCli *YDexCliFilterer) WatchTokenSymbolAdded(opts *bind.WatchOpts, sink chan<- *YDexCliTokenSymbolAdded, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "TokenSymbolAdded", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliTokenSymbolAdded)
				if err := _YDexCli.contract.UnpackLog(event, "TokenSymbolAdded", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseTokenSymbolAdded(log types.Log) (*YDexCliTokenSymbolAdded, error) {
	event := new(YDexCliTokenSymbolAdded)
	if err := _YDexCli.contract.UnpackLog(event, "TokenSymbolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the YDexCli contract.
type YDexCliTradeExecutedIterator struct {
	Event *YDexCliTradeExecuted // Event containing the contract specifics and raw log

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
func (it *YDexCliTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliTradeExecuted)
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
		it.Event = new(YDexCliTradeExecuted)
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
func (it *YDexCliTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliTradeExecuted represents a TradeExecuted event raised by the YDexCli contract.
type YDexCliTradeExecuted struct {
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
func (_YDexCli *YDexCliFilterer) FilterTradeExecuted(opts *bind.FilterOpts, baseAssetSymbolIndex []string, quoteAssetSymbolIndex []string) (*YDexCliTradeExecutedIterator, error) {

	var baseAssetSymbolIndexRule []interface{}
	for _, baseAssetSymbolIndexItem := range baseAssetSymbolIndex {
		baseAssetSymbolIndexRule = append(baseAssetSymbolIndexRule, baseAssetSymbolIndexItem)
	}
	var quoteAssetSymbolIndexRule []interface{}
	for _, quoteAssetSymbolIndexItem := range quoteAssetSymbolIndex {
		quoteAssetSymbolIndexRule = append(quoteAssetSymbolIndexRule, quoteAssetSymbolIndexItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "TradeExecuted", baseAssetSymbolIndexRule, quoteAssetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliTradeExecutedIterator{contract: _YDexCli.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0xbd75b2e24c3d43ec9f10f3583265606c0a7d8b67d147f9dc48376b81456823a1.
//
// Solidity: event TradeExecuted(address buyWallet, address sellWallet, string indexed baseAssetSymbolIndex, string indexed quoteAssetSymbolIndex, string baseAssetSymbol, string quoteAssetSymbol, uint64 baseQuantityInPips, uint64 quoteQuantityInPips, uint64 tradePriceInPips, bytes32 buyOrderHash, bytes32 sellOrderHash)
func (_YDexCli *YDexCliFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *YDexCliTradeExecuted, baseAssetSymbolIndex []string, quoteAssetSymbolIndex []string) (event.Subscription, error) {

	var baseAssetSymbolIndexRule []interface{}
	for _, baseAssetSymbolIndexItem := range baseAssetSymbolIndex {
		baseAssetSymbolIndexRule = append(baseAssetSymbolIndexRule, baseAssetSymbolIndexItem)
	}
	var quoteAssetSymbolIndexRule []interface{}
	for _, quoteAssetSymbolIndexItem := range quoteAssetSymbolIndex {
		quoteAssetSymbolIndexRule = append(quoteAssetSymbolIndexRule, quoteAssetSymbolIndexItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "TradeExecuted", baseAssetSymbolIndexRule, quoteAssetSymbolIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliTradeExecuted)
				if err := _YDexCli.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseTradeExecuted(log types.Log) (*YDexCliTradeExecuted, error) {
	event := new(YDexCliTradeExecuted)
	if err := _YDexCli.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliWalletExitClearedIterator is returned from FilterWalletExitCleared and is used to iterate over the raw logs and unpacked data for WalletExitCleared events raised by the YDexCli contract.
type YDexCliWalletExitClearedIterator struct {
	Event *YDexCliWalletExitCleared // Event containing the contract specifics and raw log

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
func (it *YDexCliWalletExitClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliWalletExitCleared)
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
		it.Event = new(YDexCliWalletExitCleared)
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
func (it *YDexCliWalletExitClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliWalletExitClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliWalletExitCleared represents a WalletExitCleared event raised by the YDexCli contract.
type YDexCliWalletExitCleared struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletExitCleared is a free log retrieval operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address indexed wallet)
func (_YDexCli *YDexCliFilterer) FilterWalletExitCleared(opts *bind.FilterOpts, wallet []common.Address) (*YDexCliWalletExitClearedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "WalletExitCleared", walletRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliWalletExitClearedIterator{contract: _YDexCli.contract, event: "WalletExitCleared", logs: logs, sub: sub}, nil
}

// WatchWalletExitCleared is a free log subscription operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address indexed wallet)
func (_YDexCli *YDexCliFilterer) WatchWalletExitCleared(opts *bind.WatchOpts, sink chan<- *YDexCliWalletExitCleared, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "WalletExitCleared", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliWalletExitCleared)
				if err := _YDexCli.contract.UnpackLog(event, "WalletExitCleared", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseWalletExitCleared(log types.Log) (*YDexCliWalletExitCleared, error) {
	event := new(YDexCliWalletExitCleared)
	if err := _YDexCli.contract.UnpackLog(event, "WalletExitCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliWalletExitWithdrawnIterator is returned from FilterWalletExitWithdrawn and is used to iterate over the raw logs and unpacked data for WalletExitWithdrawn events raised by the YDexCli contract.
type YDexCliWalletExitWithdrawnIterator struct {
	Event *YDexCliWalletExitWithdrawn // Event containing the contract specifics and raw log

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
func (it *YDexCliWalletExitWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliWalletExitWithdrawn)
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
		it.Event = new(YDexCliWalletExitWithdrawn)
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
func (it *YDexCliWalletExitWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliWalletExitWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliWalletExitWithdrawn represents a WalletExitWithdrawn event raised by the YDexCli contract.
type YDexCliWalletExitWithdrawn struct {
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
func (_YDexCli *YDexCliFilterer) FilterWalletExitWithdrawn(opts *bind.FilterOpts, wallet []common.Address, assetAddress []common.Address) (*YDexCliWalletExitWithdrawnIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "WalletExitWithdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliWalletExitWithdrawnIterator{contract: _YDexCli.contract, event: "WalletExitWithdrawn", logs: logs, sub: sub}, nil
}

// WatchWalletExitWithdrawn is a free log subscription operation binding the contract event 0x02e03f232c88c87047c5a9cfbad1213b842503fee6a002ec2eec9f5a64725bf4.
//
// Solidity: event WalletExitWithdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_YDexCli *YDexCliFilterer) WatchWalletExitWithdrawn(opts *bind.WatchOpts, sink chan<- *YDexCliWalletExitWithdrawn, wallet []common.Address, assetAddress []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "WalletExitWithdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliWalletExitWithdrawn)
				if err := _YDexCli.contract.UnpackLog(event, "WalletExitWithdrawn", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseWalletExitWithdrawn(log types.Log) (*YDexCliWalletExitWithdrawn, error) {
	event := new(YDexCliWalletExitWithdrawn)
	if err := _YDexCli.contract.UnpackLog(event, "WalletExitWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliWalletExitedIterator is returned from FilterWalletExited and is used to iterate over the raw logs and unpacked data for WalletExited events raised by the YDexCli contract.
type YDexCliWalletExitedIterator struct {
	Event *YDexCliWalletExited // Event containing the contract specifics and raw log

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
func (it *YDexCliWalletExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliWalletExited)
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
		it.Event = new(YDexCliWalletExited)
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
func (it *YDexCliWalletExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliWalletExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliWalletExited represents a WalletExited event raised by the YDexCli contract.
type YDexCliWalletExited struct {
	Wallet               common.Address
	EffectiveBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterWalletExited is a free log retrieval operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address indexed wallet, uint256 effectiveBlockNumber)
func (_YDexCli *YDexCliFilterer) FilterWalletExited(opts *bind.FilterOpts, wallet []common.Address) (*YDexCliWalletExitedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "WalletExited", walletRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliWalletExitedIterator{contract: _YDexCli.contract, event: "WalletExited", logs: logs, sub: sub}, nil
}

// WatchWalletExited is a free log subscription operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address indexed wallet, uint256 effectiveBlockNumber)
func (_YDexCli *YDexCliFilterer) WatchWalletExited(opts *bind.WatchOpts, sink chan<- *YDexCliWalletExited, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "WalletExited", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliWalletExited)
				if err := _YDexCli.contract.UnpackLog(event, "WalletExited", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseWalletExited(log types.Log) (*YDexCliWalletExited, error) {
	event := new(YDexCliWalletExited)
	if err := _YDexCli.contract.UnpackLog(event, "WalletExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YDexCliWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the YDexCli contract.
type YDexCliWithdrawnIterator struct {
	Event *YDexCliWithdrawn // Event containing the contract specifics and raw log

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
func (it *YDexCliWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YDexCliWithdrawn)
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
		it.Event = new(YDexCliWithdrawn)
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
func (it *YDexCliWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YDexCliWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YDexCliWithdrawn represents a Withdrawn event raised by the YDexCli contract.
type YDexCliWithdrawn struct {
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
func (_YDexCli *YDexCliFilterer) FilterWithdrawn(opts *bind.FilterOpts, wallet []common.Address, assetAddress []common.Address) (*YDexCliWithdrawnIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.FilterLogs(opts, "Withdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &YDexCliWithdrawnIterator{contract: _YDexCli.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x6960a1f64ecf9da0d1d1bcbfa3dd27f8c1c60de69b13faa28127dafa36c111e4.
//
// Solidity: event Withdrawn(address indexed wallet, address indexed assetAddress, string assetSymbol, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_YDexCli *YDexCliFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *YDexCliWithdrawn, wallet []common.Address, assetAddress []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _YDexCli.contract.WatchLogs(opts, "Withdrawn", walletRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YDexCliWithdrawn)
				if err := _YDexCli.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_YDexCli *YDexCliFilterer) ParseWithdrawn(log types.Log) (*YDexCliWithdrawn, error) {
	event := new(YDexCliWithdrawn)
	if err := _YDexCli.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
