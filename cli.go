package ydexabi

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Stub stub of exchange contract
type Stub struct {
	Instance *Ydexabi
}

//NewStub create a new stub of Exchange contract
func NewStub(url, exchangeContractAddr string) (*Stub, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("error when dial network %s: %s\n", url, err.Error())
		return nil, err
	}
	instance, err := NewYdexabi(common.HexToAddress(exchangeContractAddr), client)
	if err != nil {
		log.Printf("error when create instance of Exchange contract: %s\n", err.Error())
		return nil, err
	}
	return &Stub{Instance: instance}, nil
}

//read-only operation

//LoadAssetBySymbol find asset struct by token symbol
func (stub *Stub) LoadAssetBySymbol(ctx context.Context, symbol string, timestampInMs uint64) (*StructsAsset, error) {
	asset, err := stub.Instance.LoadAssetBySymbol(&bind.CallOpts{Context: ctx}, symbol, timestampInMs)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

//LoadBalanceInAssetUnitsByAddress get balance by wallet address and token contract address
func (stub *Stub) LoadBalanceInAssetUnitsByAddress(ctx context.Context, wallet string, assetAddress string) (*big.Int, error) {
	return stub.Instance.LoadBalanceInAssetUnitsByAddress(&bind.CallOpts{Context: ctx}, common.HexToAddress(wallet), common.HexToAddress(assetAddress))
}

//LoadBalanceInAssetUnitsBySymbol get balance by wallet address and token symbol
func (stub *Stub) LoadBalanceInAssetUnitsBySymbol(ctx context.Context, wallet string, assetSymbol string) (*big.Int, error) {
	return stub.Instance.LoadBalanceInAssetUnitsBySymbol(&bind.CallOpts{Context: ctx}, common.HexToAddress(wallet), assetSymbol)
}

//LoadBalanceInPipsByAddress get balance in pips by wallet address and token contract address
func (stub *Stub) LoadBalanceInPipsByAddress(ctx context.Context, wallet string, assetAddress string) (uint64, error) {
	return stub.Instance.LoadBalanceInPipsByAddress(&bind.CallOpts{Context: ctx}, common.HexToAddress(wallet), common.HexToAddress(assetAddress))
}

//LoadBalanceInPipsBySymbol get balance in pips by wallet address and token symbol
func (stub *Stub) LoadBalanceInPipsBySymbol(ctx context.Context, wallet string, assetSymbol string) (uint64, error) {
	return stub.Instance.LoadBalanceInPipsBySymbol(&bind.CallOpts{Context: ctx}, common.HexToAddress(wallet), assetSymbol)
}

//LoadFeeWalet get fee wallet address
func (stub *Stub) LoadFeeWalet(ctx context.Context) (string, error) {
	addr, err := stub.Instance.LoadFeeWallet(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

//LoadPartiallyFilledOrderQuantityInPips load quantity of a partially filled order by order hash
func (stub *Stub) LoadPartiallyFilledOrderQuantityInPips(ctx context.Context, orderHash string) (uint64, error) {
	b, err := hex.DecodeString(orderHash)
	if err != nil {
		return 0, err
	}
	if len(b) != 32 {
		return 0, errors.New("length of order hash must be 32")
	}
	var arr [32]byte
	for index, v := range b {
		arr[index] = v
	}
	return stub.Instance.LoadPartiallyFilledOrderQuantityInPips(&bind.CallOpts{Context: ctx}, arr)
}

//write operation

//SetAdmin set admin account for Exchange contract
func (stub *Stub) SetAdmin(opts *bind.TransactOpts, newAdminAddr string) (*types.Transaction, error) {
	return stub.Instance.SetAdmin(opts, common.HexToAddress(newAdminAddr))
}

//RemoveAdmin remove admin account for Exchange contract
func (stub *Stub) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return stub.Instance.RemoveAdmin(opts)
}

//SetCustodian set address of Custodian contract
func (stub *Stub) SetCustodian(opts *bind.TransactOpts, newCustodianAddr string) (*types.Transaction, error) {
	return stub.Instance.SetCustodian(opts, common.HexToAddress(newCustodianAddr))
}

//SetDispatcher set dispatcher account for Exchange contract
func (stub *Stub) SetDispatcher(opts *bind.TransactOpts, newDispatcherAddr string) (*types.Transaction, error) {
	return stub.Instance.SetDispatcher(opts, common.HexToAddress(newDispatcherAddr))
}

//RemoveDispatcher remove dispatcher account for Exchange contract
func (stub *Stub) RemoveDispatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return stub.Instance.RemoveDispatcher(opts)
}

//SetFeeWallet set fee wallet address for Exchange contract
func (stub *Stub) SetFeeWallet(opts *bind.TransactOpts, newFeeWalletAddr string) (*types.Transaction, error) {
	return stub.Instance.SetFeeWallet(opts, common.HexToAddress(newFeeWalletAddr))
}

//SetChainPropagationPeriod set fee wallet address for Exchange contract
func (stub *Stub) SetChainPropagationPeriod(opts *bind.TransactOpts, newChainPropagationPeriod uint64) (*types.Transaction, error) {
	return stub.Instance.SetChainPropagationPeriod(opts, new(big.Int).SetUint64(newChainPropagationPeriod))
}

//RegisterToken register token information in Exchange contract
func (stub *Stub) RegisterToken(opts *bind.TransactOpts, tokenAddr string, symbol string, decimals uint8) (*types.Transaction, error) {
	return stub.Instance.RegisterToken(opts, common.HexToAddress(tokenAddr), symbol, decimals)
}

//AddTokenSymbol add symbol for one token
func (stub *Stub) AddTokenSymbol(opts *bind.TransactOpts, tokenAddr string, symbol string) (*types.Transaction, error) {
	return stub.Instance.AddTokenSymbol(opts, common.HexToAddress(tokenAddr), symbol)
}

//ConfirmTokenRegistration confirm token information which have registered in Exchange contract
func (stub *Stub) ConfirmTokenRegistration(opts *bind.TransactOpts, tokenAddr string, symbol string, decimals uint8) (*types.Transaction, error) {
	return stub.Instance.ConfirmTokenRegistration(opts, common.HexToAddress(tokenAddr), symbol, decimals)
}

//DepositBNB deposit BNB to Exchange contract
func (stub *Stub) DepositBNB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return stub.Instance.DepositBNB(opts)
}

//DepositTokenByAddress deposit token to Exchange contract by token contract address
func (stub *Stub) DepositTokenByAddress(opts *bind.TransactOpts, tokenAddr string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return stub.Instance.DepositTokenByAddress(opts, common.HexToAddress(tokenAddr), quantityInAssetUnits)
}

//DepositTokenBySymbol deposit token to Exchange contract by token symbol
func (stub *Stub) DepositTokenBySymbol(opts *bind.TransactOpts, assetSymbol string, quantityInAssetUnits *big.Int) (*types.Transaction, error) {
	return stub.Instance.DepositTokenBySymbol(opts, assetSymbol, quantityInAssetUnits)
}

//ExecuteTrade execute trade
func (stub *Stub) ExecuteTrade(opts *bind.TransactOpts, buy *StructsOrder, sell *StructsOrder, trade *StructsTrade) (*types.Transaction, error) {
	return stub.Instance.ExecuteTrade(opts, *buy, *sell, *trade)
}

//InvalidateOrderNonce invalidate order by nonce
func (stub *Stub) InvalidateOrderNonce(opts *bind.TransactOpts, nonce string) (*types.Transaction, error) {
	b, err := hex.DecodeString(nonce)
	if err != nil {
		return nil, err
	}
	return stub.Instance.InvalidateOrderNonce(opts, new(big.Int).SetBytes(b))
}

//Withdraw withdraw token from Exchange contract
func (stub *Stub) Withdraw(opts *bind.TransactOpts, withdrawal *StructsWithdrawal) (*types.Transaction, error) {
	return stub.Instance.Withdraw(opts, *withdrawal)
}

//WithdrawExit withdraw and exit
func (stub *Stub) WithdrawExit(opts *bind.TransactOpts, assetAddr string) (*types.Transaction, error) {
	return stub.Instance.WithdrawExit(opts, common.HexToAddress(assetAddr))
}

//ExitWallet wallet exit
func (stub *Stub) ExitWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return stub.Instance.ExitWallet(opts)
}

//ClearWalletExit clear exit status for current wallet
func (stub *Stub) ClearWalletExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return stub.Instance.ClearWalletExit(opts)
}
