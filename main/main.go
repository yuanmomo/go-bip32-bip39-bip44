package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tyler-smith/go-bip32"
)

func main()  {
	extendedMasterPublic, _ := bip32.B58Deserialize("xpub6E3rJViLG4vuxAFtecHm7iTkbUvgudwH8CwYZkAp3z5GJxheVM1giUuSvvzmGA4Tzs9HbEd2skmqbwPtLNuNY6sKdYZvUCr8VdNduPsmYAT")

	pubKey, _ := extendedMasterPublic.NewChildKey(0)
	fmt.Printf("%s \n",hex.EncodeToString(pubKey.Key))


	pubKey3, _ := extendedMasterPublic.NewChildKey(3)
	fmt.Printf("%s \n",hex.EncodeToString(pubKey3.Key))
}