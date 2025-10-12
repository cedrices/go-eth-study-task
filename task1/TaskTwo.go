package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/cedric/go-eth-study-task/counter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//部署智能合约
	// client, err := ethclient.Dial("https://sepolia.infura.io/v3/b73024675d6e4633a6452256aae40925")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Close()
	// privateKey, err := crypto.HexToECDSA("74928517cafb1f4b596acb36e52c79c2eab33e1f6acddeab4911038c13c09777")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// gaslimit := uint64(3000000)
	// chainID, err := client.NetworkID(context.Background())
	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)
	// auth.GasLimit = gaslimit
	// auth.GasPrice = gasPrice
	// contractAddress, tx, counterInstance, err := counter.DeployCounter(auth, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Contract deployed at address:", contractAddress.Hex())
	// log.Println("tx:", tx.Hash().Hex())

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b73024675d6e4633a6452256aae40925")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	contractAddress := common.HexToAddress("0xaFb6b3FdF5A52CD90C6dfeb7926D5048241Eb5c5")
	counterInstance, err := counter.NewCounter(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	// auth := getAuth(client)
	// counterInstance.Add(auth, big.NewInt(20))
	// time.Sleep(5 * time.Second)
	counts, err := counterInstance.Get(&bind.CallOpts{Context: context.Background()})

	log.Println("counts:", counts)
}

func getAuth(client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("74928517cafb1f4b596acb36e52c79c2eab33e1f6acddeab4911038c13c09777")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
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
	gasPrice = gasPrice.Add(gasPrice, big.NewInt(10000000))
	gaslimit := uint64(3000000) * 3
	chainID, err := client.NetworkID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gaslimit
	auth.GasPrice = gasPrice
	return auth
}
