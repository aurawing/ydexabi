package ydexabi

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var stub *Stub

func TestMain(m *testing.M) {
	stub, _ = NewStubByURL("https://data-seed-prebsc-1-s2.binance.org:8545", "0x1c74ba5d6a25e26a84f6a3c5279fea3b71bf461e")
	os.Exit(m.Run())
}

func TestRecipent(t *testing.T) {
	ethclient, err := ethclient.Dial("https://data-seed-prebsc-1-s2.binance.org:8545")
	if err != nil {
		panic(err)
	}
	ret, err := ethclient.TransactionReceipt(context.Background(), common.HexToHash("0x3a2205fed612a25a6be94f912b3c2670a92ef5e8c2f1071ebdc70b5b160e04f3"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ret)

}
func TestLoadAssetBySymbol(t *testing.T) {
	asset, err := stub.LoadAssetBySymbol(context.Background(), "BNB", 1634046847000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", asset)
}

func TestLoadBalanceInAssetUnitsByAddress(t *testing.T) {
	balance, err := stub.LoadBalanceInAssetUnitsByAddress(context.Background(), "0xcCB98929A6D118d51224F6451F8CBE599E9343BE", "0x7C28F63e0671D7e7BF1c6A8d81Dd0D7a93004FCe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", balance.String())
}

func TestLoadBalanceInAssetUnitsBySymbol(t *testing.T) {
	balance, err := stub.LoadBalanceInAssetUnitsBySymbol(context.Background(), "0xcCB98929A6D118d51224F6451F8CBE599E9343BE", "TSC2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", balance.String())
}

func TestLoadBalanceInPipsByAddress(t *testing.T) {
	balance, err := stub.LoadBalanceInPipsByAddress(context.Background(), "0x11b5A8efFa4E05EF833C8b987bc0c3336eF81eEf", "0x7C28F63e0671D7e7BF1c6A8d81Dd0D7a93004FCe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", balance)
}

func TestLoadBalanceInPipsBySymbol(t *testing.T) {
	balance, err := stub.LoadBalanceInPipsBySymbol(context.Background(), "0x11b5A8efFa4E05EF833C8b987bc0c3336eF81eEf", "TSC2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", balance)
}

func TestLoadFeeWallet(t *testing.T) {
	addr, err := stub.LoadFeeWalet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", addr)
}

func TestLoadPartiallyFilledOrderQuantityInPips(t *testing.T) {
	quantity, err := stub.LoadPartiallyFilledOrderQuantityInPips(context.Background(), "c0e8d52a58ffeca6f54321bef57dbe2a8c95de4e0a08def5c4dde10ab020d9d1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", quantity)
}

func TestIsWalletExit(t *testing.T) {
	b, err := stub.IsWalletExit(context.Background(), "0xcCB98929A6D118d51224F6451F8CBE599E9343BE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", b)
}

func TestIsWalletAllExit(t *testing.T) {
	b, err := stub.IsWalletExitFinalized(context.Background(), "0xcCB98929A6D118d51224F6451F8CBE599E9343BE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", b)
}

func TestLoadBalanceAssetAddress(t *testing.T) {
	addrs, err := stub.LoadBalanceAssetAddress(context.Background(), "0xcCB98929A6D118d51224F6451F8CBE599E9343BE")
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		fmt.Printf("%s\n", addr)
	}
}

func TestGenerateSigForWithdrawal(t *testing.T) {
	sk := "a793a6351ee63c9cf6278784f7d3939084d81cd2ef7663ffe8c1931ad07a1f7f"
	withdrawal := new(StructsWithdrawal)
	withdrawal.WithdrawalType = 0
	nonce := new(big.Int)
	noncebytes, _ := hex.DecodeString("62e802dd2b3311eca9207085c2b700ea")
	nonce.SetBytes(noncebytes)
	withdrawal.Nonce = nonce
	withdrawal.WalletAddress = common.HexToAddress("0xcCB98929A6D118d51224F6451F8CBE599E9343BE")
	withdrawal.AssetSymbol = "TSC1"
	withdrawal.QuantityInPips = 1000000000000
	withdrawal.AutoDispatchEnabled = true
	fmt.Println(GenerateSigForWithdrawal(withdrawal, sk))
}

func TestGenerateSigForOrder(t *testing.T) {
	// sk := "a793a6351ee63c9cf6278784f7d3939084d81cd2ef7663ffe8c1931ad07a1f7f"
	// order := new(StructsOrder)
	// order.SignatureHashVersion = 1
	// nonce := new(big.Int)
	// noncebytes, _ := hex.DecodeString("b9327c022b3411eca24a7085c2b700ea")
	// nonce.SetBytes(noncebytes)
	// order.Nonce = nonce
	// order.WalletAddress = common.HexToAddress("0xcCB98929A6D118d51224F6451F8CBE599E9343BE")
	// marketSymbol := "TSC1-BNB"
	// order.OrderType = 1
	// order.Side = 0
	// order.QuantityInPips = 10000000000
	// order.IsQuantityInQuote = false
	// order.LimitPriceInPips = 30000
	// order.StopPriceInPips = 0
	// order.ClientOrderId = ""
	// order.TimeInForce = 0
	// order.SelfTradePrevention = 3
	// order.CancelAfter = 0
	// fmt.Println(GenerateSigForOrder(order, marketSymbol, sk))

	// sk := "e701b622b54dee377f974013f3b7e84cfc8cd89f9b62134128ef63eced67545e"
	// order := new(StructsOrder)
	// order.SignatureHashVersion = 1
	// nonce := new(big.Int)
	// noncebytes, _ := hex.DecodeString("1bf801532b3511ec86ef7085c2b700ea")
	// nonce.SetBytes(noncebytes)
	// order.Nonce = nonce
	// order.WalletAddress = common.HexToAddress("0x11b5A8efFa4E05EF833C8b987bc0c3336eF81eEf")
	// marketSymbol := "TSC1-BNB"
	// order.OrderType = 1
	// order.Side = 1
	// order.QuantityInPips = 15000000000
	// order.IsQuantityInQuote = false
	// order.LimitPriceInPips = 30000
	// order.StopPriceInPips = 0
	// order.ClientOrderId = ""
	// order.TimeInForce = 0
	// order.SelfTradePrevention = 3
	// order.CancelAfter = 0
	// fmt.Println(GenerateSigForOrder(order, marketSymbol, sk))

	sk := "a793a6351ee63c9cf6278784f7d3939084d81cd2ef7663ffe8c1931ad07a1f7f"
	order := new(StructsOrder)
	order.SignatureHashVersion = 1
	nonce := new(big.Int)
	noncebytes, _ := hex.DecodeString("c0de054a2b3811ec9c6d7085c2b700ea")
	nonce.SetBytes(noncebytes)
	order.Nonce = nonce
	order.WalletAddress = common.HexToAddress("0xcCB98929A6D118d51224F6451F8CBE599E9343BE")
	marketSymbol := "TSC1-BNB"
	order.OrderType = 0
	order.Side = 0
	order.QuantityInPips = 15000000000
	order.IsQuantityInQuote = true
	order.LimitPriceInPips = 0
	order.StopPriceInPips = 0
	order.ClientOrderId = ""
	order.TimeInForce = 0
	order.SelfTradePrevention = 3
	order.CancelAfter = 0
	fmt.Println(GenerateSigForOrder(order, marketSymbol, sk))
}

func TestFilterDeposited(t *testing.T) {
	//???????????????????????????????????????
	opts := new(bind.FilterOpts)
	//????????????????????????13154871???13243195???End???nil???????????????????????????
	opts.Start = 13154871
	end := uint64(13243195)
	opts.End = &end
	opts.Context = context.Background()
	//FilterDeposited????????????Deposited???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????indexed?????????????????????????????????wallet???assetSymbolIndex???????????????????????????1??????BNB????????????????????????????????????13154871???13243195???
	iter, err := stub.Instance.FilterDeposited(opts, []common.Address{common.HexToAddress("0xcCB98929A6D118d51224F6451F8CBE599E9343BE")}, nil, []string{"BNB"})
	if err != nil {
		panic(err)
	}
	//?????????????????????
	for {
		//Next????????????true????????????????????????????????????????????????Event??????????????????
		if iter.Next() {
			fmt.Printf("Event: %+v\n", iter.Event)
		} else {
			if iter.Error() != nil {
				//Next????????????false???Error()??????????????????nil?????????????????????????????????????????????????????????????????????
				fmt.Printf("error: %s\n", iter.Error())
			} else {
				//Next????????????false???Error()????????????nil????????????????????????????????????????????????
				return
			}
		}
	}
}
