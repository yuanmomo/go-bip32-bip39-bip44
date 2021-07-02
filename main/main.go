package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jbenet/go-base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Generate seed by BIP39 Mnemonic and BIP39 Passphrase
	mnemonic := "purse alcohol glimpse oblige birth return urge gift innocent town now patient"
	passphrase := ""
	seed := bip39.NewSeed(mnemonic, passphrase)
	fmt.Printf("mnemonic : %s\n", mnemonic)
	fmt.Printf("Seed : %s\n", hex.EncodeToString(seed))
	fmt.Println()

	// TODO. handle error
	// BIP32 Root Key
	rootPrivateKey, _ := bip32.NewMasterKey(seed)
	rootPubKey := rootPrivateKey.PublicKey()
	fmt.Printf("BIP32 root private key : %s\n", rootPrivateKey.String())
	fmt.Printf("BIP32 root public key : %s\n", rootPubKey.String())
	fmt.Println()

	// Get BIP44 Extended Private Key
	m44H, _ := rootPrivateKey.NewChildKey(bip32.FirstHardenedChild + 44)
	m44H60H, _ := m44H.NewChildKey(bip32.FirstHardenedChild + 60)
	m44H60H0H, _ := m44H60H.NewChildKey(bip32.FirstHardenedChild + 0)
	m44H60H0H0PrivateKey, _ := m44H60H0H.NewChildKey(0)
	// Get BIP44 Extended public Key
	m44H60H0H0PublicKey := m44H60H0H0PrivateKey.PublicKey()
	fmt.Printf("BIP32/BIP44 extended private key : %s\n", m44H60H0H0PrivateKey.String())
	fmt.Printf("BIP32/BIP44 extended public key : %s\n", m44H60H0H0PublicKey.String())
	fmt.Println()

	bip44ExtendedPrivateKey, _ := bip32.B58Deserialize(m44H60H0H0PrivateKey.String())
	bip44ExtendedPublicKey, _ := bip32.B58Deserialize(m44H60H0H0PublicKey.String())

	// Derivation Path m/44'/60'/0'/0:
	for i := 0; i < 5; i++ {
		// Get private key by BIP44 Extended Private Key
		priKeyBIP44, _ := bip44ExtendedPrivateKey.NewChildKey(uint32(i))
		priKeyDecode := base58.Decode(priKeyBIP44.String())
		privateKey := priKeyDecode[46:78]

		// Get public key by BIP44 Extended Private Key
		pubKey, _ := bip44ExtendedPublicKey.NewChildKey(uint32(i))

		// Get address
		publicKeyString := hex.EncodeToString(pubKey.Key);
		publicKeyBytes, _ := hex.DecodeString(publicKeyString)
		ecdsaPub, _ := crypto.DecompressPubkey(publicKeyBytes);
		// To address
		EthAddress := crypto.PubkeyToAddress(*ecdsaPub).String()

		// print
		fmt.Printf("Path : m/44'/60'/0'/0/%d,address : %s, public key: 0x%s, private key : 0x%s\n",
			i, EthAddress, hex.EncodeToString(pubKey.Key), hex.EncodeToString(privateKey))
	}
}
