package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("[%s] \n", hex.EncodeToString(sum[:]))
}
