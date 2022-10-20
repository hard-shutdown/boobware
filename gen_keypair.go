package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	priv, _, _, _ := elliptic.GenerateKey(elliptic.P521(), rand.Reader)
	fmt.Println(hex.EncodeToString(priv))
}
