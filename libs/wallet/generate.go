package wallet

import (
	"fmt"
	"os"
	"github.com/tyler-smith/go-bip39"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

/*func main() {

	fmt.Println(GenerateMnemonic())

	addresses := GenerateAddressRangeFromMnemonic(0, 20)
	fmt.Println(addresses)

}*/

func GenerateMnemonic() string {
	//check if mnemonic exists
	mnemonic, ok := os.LookupEnv("MNEMONICSEEDPHRASE")
	if ok {
		fmt.Println("mnemonic exists continuing...")
	} else {
		fmt.Println("mnemonic is not set generating new one. Write it down or LOSE EVERYTHING")
		// Generate a mnemonic
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ = bip39.NewMnemonic(entropy)

		//set mnemonic env var
		//TODO: this onlyh sets the env var during the run session. It clears when program completes
		os.Setenv("MNEMONICSEEDPHRASE", mnemonic)
	}
	return mnemonic
}

func GenerateAddressRangeFromMnemonic(first int, numberOf int) []string {

	// Create an HD wallet from the mnemonic
	wallet, _ := hdwallet.NewFromMnemonic(os.Getenv("MNEMONICSEEDPHRASE"))

	var addresses []string

	 for i := first; i < first+numberOf; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, _ := wallet.Derive(path, false)
		addresses = append(addresses, account.Address.Hex())
	 }
	 return addresses
}
