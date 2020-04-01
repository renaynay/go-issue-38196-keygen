package main

import (
	"fmt"

	crypto "github.com/karalabe/go-issue-38196-crypto"
)

func main() {
	fmt.Println("Generating crypto key:", crypto.GenerateKey())
}
