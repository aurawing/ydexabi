package ydexabi

import (
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateSigForWithdrawal(withdrawal *StructsWithdrawal, sk string) (string, error) {
	nonce := withdrawal.Nonce.Bytes()
	walletAddress := withdrawal.WalletAddress.Bytes()
	var assetInfo []byte
	if withdrawal.WithdrawalType == 0 {
		assetInfo = []byte(withdrawal.AssetSymbol)
	} else if withdrawal.WithdrawalType == 1 {
		assetInfo = withdrawal.AssetAddress.Bytes()
	} else {
		return "", errors.New("wrong withdraw type")
	}
	quantityInPips := []byte(pipToDecimal(withdrawal.QuantityInPips))
	autoDispatchEnabled := []byte{1}
	if !withdrawal.AutoDispatchEnabled {
		//autoDispatchEnabled = []byte{0}
		return "", errors.New("autoDispatchEnabled must be true")
	}
	hash := crypto.Keccak256Hash(nonce, walletAddress, assetInfo, quantityInPips, autoDispatchEnabled)
	prefixedHash := crypto.Keccak256Hash(
		[]byte("\x19Ethereum Signed Message:\n32"),
		hash.Bytes(),
	)
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", err
	}
	sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	sig[64] += 27
	return hex.EncodeToString(sig), nil
}

func GenerateSigForOrder(order *StructsOrder, marketSymbol, sk string) (string, error) {
	signatureHashVersion := []byte{order.SignatureHashVersion}
	nonce := order.Nonce.Bytes()
	walletAddress := order.WalletAddress.Bytes()
	mktSymbol := []byte(marketSymbol)
	orderType := []byte{order.OrderType}
	side := []byte{order.Side}
	quantityInPips := []byte(pipToDecimal(order.QuantityInPips))
	isQuantityInQuote := []byte{0}
	if order.IsQuantityInQuote {
		isQuantityInQuote = []byte{1}
	}
	limitPriceInPips := []byte{}
	if order.LimitPriceInPips > 0 {
		limitPriceInPips = []byte(pipToDecimal(order.LimitPriceInPips))
	}
	stopPriceInPips := []byte{}
	if order.StopPriceInPips > 0 {
		stopPriceInPips = []byte(pipToDecimal(order.StopPriceInPips))
	}
	clientOrderId := []byte(order.ClientOrderId)
	timeInForce := []byte{order.TimeInForce}
	selfTradePrevention := []byte{order.SelfTradePrevention}
	cancelAfter := make([]byte, 8)
	binary.BigEndian.PutUint64(cancelAfter, order.CancelAfter)
	hash := crypto.Keccak256Hash(signatureHashVersion, nonce, walletAddress, mktSymbol, orderType, side, quantityInPips, isQuantityInQuote, limitPriceInPips, stopPriceInPips, clientOrderId, timeInForce, selfTradePrevention, cancelAfter)
	prefixedHash := crypto.Keccak256Hash(
		[]byte("\x19Ethereum Signed Message:\n32"),
		hash.Bytes(),
	)
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", err
	}
	sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	sig[64] += 27
	return hex.EncodeToString(sig), nil
}

func pipToDecimal(pips uint64) string {
	var copy uint64 = pips
	var length uint64 = 0
	for copy != 0 {
		length++
		copy /= 10
	}
	if length < 9 {
		length = 9
	}
	length++

	decimal := make([]uint8, length)
	for i := length; i > 0; i-- {
		if length-i == 8 {
			decimal[i-1] = uint8(46)
		} else {
			decimal[i-1] = uint8(48 + (pips % 10))
			pips /= 10
		}
	}
	return string(decimal)
}
