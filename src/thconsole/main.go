package main

import "github.com/dizzyd/gogotelehash"
import "crypto/rand"
import "crypto/rsa"

func main() {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	s, _ := telehash.NewSwitch("0.0.0.0:0", key)
}
