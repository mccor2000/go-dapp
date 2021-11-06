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

	address := common.HexToAddress("0x2C371eb1C63dD2a3867cD2B596F7D2373A2FDe1A")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
