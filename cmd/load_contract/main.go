package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mccor2000/go-dapp/pkg/store"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/c3a654f66f8c4f04827f14d20e5dadaf")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x9Fe1439d69a59A2589648C12c0dF2E089E8Ea9F7")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
