package main
//
//import (
//	"fmt"
//	"github.com/btcsuite/btcutil/hdkeychain"
//)
//
//func main() {
//	extKey, _ := hdkeychain.NewKeyFromString("xpub6DubSqbhBUXSF4h3dnE9hgw93qxs7nfrLBc9i2DGwvMKEnpX4Uq9m5btyrTJeMQvS4TeziJBQfiMYoHdAjVGPnEX746rLAXrVWHs5odYNGq")
//
//	path := []uint32{ 44, 60,  0, 0 , 0};
//
//	for _, childNum := range path {
//		extKey, _ = extKey.Derive(childNum)
//	}
//
//	btcecPrivKey, _ := extKey.ECPrivKey()
//	privateKey := btcecPrivKey.ToECDSA()
//
//	//path := "m/44H/60H/0H/0/0"
//
//	fmt.Printf("%s:%s \n", path, privateKey)
//}
//
