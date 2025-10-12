package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//infura
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b73024675d6e4633a6452256aae40925")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("区块高度：", block.Number())
	fmt.Println("区块哈希：", block.Hash())
	fmt.Println("区块时间：", block.Time())
	fmt.Println("区块交易数：", len(block.Transactions()))
	fmt.Println("区块矿工：", block.Coinbase())
	fmt.Println("区块难度：", block.Difficulty())
	fmt.Println("区块gasLimit：", block.GasLimit())
	fmt.Println("区块gasUsed：", block.GasUsed())
	fmt.Println("区块nonce：", block.Nonce())

	privateKey, err := crypto.HexToECDSA("74928517cafb1f4b596acb36e52c79c2eab33e1f6acddeab4911038c13c09777")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	pubAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), pubAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gaslimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{})
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(1_000_000_000_000_000)
	var data []byte
	toAddress := common.HexToAddress("0x86ea349D89C26E5BD569Eb47f073C0a9eaBD8955")
	tx := types.NewTransaction(nonce, toAddress, amount, gaslimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println("交易详情：", signedTx)
	fmt.Println("交易哈希：", signedTx.Hash().Hex())
	fmt.Println("交易发送方：", pubAddress.Hex())
	fmt.Println("交易接收方：", signedTx.To().Hex())

}
