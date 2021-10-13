package ydexabi

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

var stub *Stub

func TestMain(m *testing.M) {
	stub = NewStub("https://data-seed-prebsc-1-s2.binance.org:8545", "0x5a20630dc55332ba80d947784032c44257c12474")
	os.Exit(m.Run())
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
