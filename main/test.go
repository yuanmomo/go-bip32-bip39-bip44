package main
//
//import (
//"fmt"
//"github.com/btcsuite/btcutil/hdkeychain"
//"github.com/ethereum/go-ethereum/crypto"
//)
//
//func main() {
//	extendedKey, _ := hdkeychain.NewKeyFromString("xprvA1pqm2imDm4ggksV27rBygcqqgJL4umijYvE2aFteyxhV7AVLTAGV8a9HTz88opMGeRasJQqX75U8tgGuHGZLB3TVuUyV2Qj8dcsnmAPX3j")
//
//	// This gives the path: m/44H
//	acc44H, _ := extendedKey.Child(44)
//
//	// This gives the path: m/44H/60H
//	acc44H60H, _ := acc44H.Child(60)
//
//	// This gives the path: m/44H/60H/0H
//	acc44H60H0H, _ := acc44H60H.Child( 0)
//
//	// This gives the path: m/44H/60H/0H/0
//	acc44H60H0H0, _ := acc44H60H0H.Child(0)
//
//	// This gives the path: m/44H/60H/0H/0/0
//	acc44H60H0H00, _ := acc44H60H0H0.Child(0)
//
//	publicKey, _ := acc44H60H0H00.ECPubKey()
//
//	pub := publicKey.ToECDSA()
//
//	path := "m/44H/60H/0H/0/0"
//
//	fmt.Printf("%s:%s \n", path, crypto.PubkeyToAddress(*pub))
//}
//
