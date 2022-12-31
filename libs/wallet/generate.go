/*package main

import (
  "fmt"
  "github.com/tyler-smith/go-bip39"
  //"github.com/tyler-smith/go-bip32"
)

func main(){
  // Generate a mnemonic for memorization or user-friendly seeds
  entropy, _ := bip39.NewEntropy(256)
  mnemonic, _ := bip39.NewMnemonic(entropy)

  // Generate a Bip32 HD wallet for the mnemonic and a user supplied password
  //seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

  //masterKey, _ := bip32.NewMasterKey(seed)
  //publicKey := masterKey.PublicKey()

  // Display mnemonic and keys
  fmt.Println("", mnemonic)
  //fmt.Println("Master private key: ", masterKey)
  //fmt.Println("Master public key: ", publicKey)
}*/

//from chat.openai.com  bot
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha3"
	"fmt"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Generate a new seed phrase with 256 bits of entropy
	mnemonic, err := bip39.NewMnemonic(256)
	if err != nil {
		// handle error
	}

	// Derive a seed from the mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Generate 10 deterministic keys and use them to generate
	// public/private key pairs
	for i := 0; i < 10; i++ {
		key := sha3.GenerateDeterministicKey(seed, 32)
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), bytes.NewReader(key))
		if err != nil {
			// handle error
		}
		publicKey := privateKey.Public()
		fmt.Printf("Private key: %x\n", privateKey)
		fmt.Printf("Public key: %x\n", publicKey)
	}
}
