package main

import (
	"fmt"
	"os"
	"github.com/tyler-smith/go-bip39"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {

	//check if mnemonic exists
	mnemonic, ok := os.LookupEnv("MNEMONICSEEDPHRASE")
	if ok {
		fmt.Println("mnemonic exists continuing...")
		mnemonic = os.Getenv("MNEMONICSEEDPHRASE")
	} else {
		fmt.Println("mnemonic is not set generating new one")
		// Generate a mnemonic
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ = bip39.NewMnemonic(entropy)

		//set mnemonic env var
		//os.Setenv("MNEMONICSEEDPHRASE", mnemonic)
	}

	fmt.Println(mnemonic)

	// Create an HD wallet from the mnemonic
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	for i := 0; i < 10; i++ {
		// Generate an Ethereum account
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, _ := wallet.Derive(path, false)
		fmt.Printf("Account %d address: %s\n", i, account.Address.Hex())
	}
}
