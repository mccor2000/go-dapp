package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/mccor2000/go-dapp/pkg/store"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/c3a654f66f8c4f04827f14d20e5dadaf")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f4f71829ee79b042e3da5c7e481402b22029170d447816378358d0a589d3603b")
	if err != nil {
		log.Fatal(err)
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1).Mul(gasPrice, big.NewInt(2))

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Println("deploying...")
		log.Fatal(err)
	}

	fmt.Println("contract address:", address.Hex())
	fmt.Println("transaction hash:", tx.Hash().Hex())

	_ = instance
}
