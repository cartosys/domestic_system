package wallet

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

func connect() {
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/API_KEY")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA("PRIVATE_KEY")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

        value := big.NewInt(10000000000000000) // in wei (0.01 eth)
    gasLimit := uint64(21000)               // in units
    tip := big.NewInt(2000000000)           // maxPriorityFeePerGas = 2 Gwei
    feeCap := big.NewInt(20000000000)       // maxFeePerGas = 20 Gwei
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("ADDRESS_TO")
    var data []byte

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    tx := types.NewTx(&types.DynamicFeeTx{
        ChainID:   chainID,
        Nonce:     nonce,
        GasFeeCap: feeCap,
        GasTipCap: tip,
        Gas:       gasLimit,
        To:        &toAddress,
        Value:     value,
        Data:      data,
    })

    signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)

    if err != nil {
        log.Fatal(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Transaction hash: %s", signedTx.Hash().Hex())

}
