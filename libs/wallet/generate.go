package wallet

import (
	"fmt"
	"log"
	"os"

	//for creating & writing to ssh key file
	"os/user"
	"bufio"

	"github.com/tyler-smith/go-bip39"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	qrcode "github.com/skip2/go-qrcode"
	//"github.com/charmbracelet/melt"
)

/*func main() {

	fmt.Println(GenerateMnemonic())

	addresses := GenerateAddressRangeFromMnemonic(0, 20)
	fmt.Println(addresses)

}*/
func Txt2QR(text string) string {
	// Generate the QR code with the given text
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		panic(err)
	}

	// Get the ASCII representation of the QR code
	ascii := qr.ToSmallString(false)

	// Print the ASCII QR code to the console
	return ascii
}
func GenerateMnemonic() string {
	//check if mnemonic exists
	mnemonic, ok := os.LookupEnv("MNEMONICSEEDPHRASE")
	if ok && len(mnemonic) > 0 {
		fmt.Println("mnemonic exists continuing...")
	} else {
		fmt.Println("mnemonic is not set generating new one. Write it down or LOSE EVERYTHING")
		// Generate a mnemonic
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ = bip39.NewMnemonic(entropy)

		//set mnemonic env var
		//TODO: this only sets the env var during the run session. It clears when program completes
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

		fmt.Printf("Account %d address: %s\n", i, account.Address.Hex())
		privateKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			log.Fatal(err)
		}

		// Print the private key
		fmt.Printf("Private key: %s\n", privateKey)
		//fmt.Printf("Private key: %s\n", common.ToHex(crypto.FromECDSA(privateKey)))

		addresses = append(addresses, account.Address.Hex())

	}
	return addresses
}

func KeyFileCheck() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	fileName := "domestic_system.txt"
	filePath := dir + "/" + fileName

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Create file if it does not exist
			file, createErr := os.Create(filePath)
			if createErr != nil {
				 fmt.Println("Error creating ssh key file:", createErr)
				 return
			}
			datawriter := bufio.NewWriter(file)
 			datawriter.WriteString(GenerateMnemonic() + "\n")
			datawriter.Flush()

			file.Close()

			fmt.Printf("key file %s has been created", filePath)
	} else {
		 	fmt.Println("%s already exists in the home directory.", filePath)

			file, err := os.Open(filePath)
			if err != nil {
			  fmt.Println(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			//for scanner.Scan() {
        //fmt.Println(scanner.Text())
    	//}
			scanner.Scan()
			os.Setenv("MNEMONICSEEDPHRASE", scanner.Text())
			file.Close()
	}
	mnemonic := os.Getenv("MNEMONICSEEDPHRASE")
	fmt.Println(Txt2QR(mnemonic))
}
