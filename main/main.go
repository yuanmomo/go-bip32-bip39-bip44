package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"go-hd-wallet/bip44"
)

func main()  {
	// Generate a mnemonic for memorization or user-friendly seeds
	//entropy, _ := bip39.NewEntropy(128)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	mnemonic := "pool capital page melody monster divert wealth file coil guide cigar chef"

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Sed ", hex.EncodeToString(seed))
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)


	ak, _ := bip44.NewAccountKeyFromXPubKey(publicKey.String())

	externalAddress, _ := ak.DeriveP2PKAddress(60, 0, bip44.MAINNET)
	internalAddress, _ := ak.DeriveP2PKAddress(60, 0, bip44.MAINNET)

	fmt.Println("External Address: ", externalAddress)
	fmt.Println("Internal Address: ", internalAddress)
}