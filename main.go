package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	myPrivateKey := [32]byte{191, 145, 93, 226, 52, 69, 63, 94, 153, 130, 232, 193, 74, 144, 81, 137, 132, 134, 145, 160, 130, 210, 154, 23, 221, 139, 188, 40, 59, 38, 146, 143}
	myPublicKey := generatePublicKey(myPrivateKey)
	fmt.Println("My private key: ", myPrivateKey)
	fmt.Println("My public key: ", myPublicKey)

	theirPrivateKey := [32]byte{63, 25, 33, 73, 88, 56, 26, 249, 232, 213, 30, 122, 35, 253, 225, 48, 148, 173, 229, 241, 29, 141, 170, 31, 37, 16, 158, 227, 75, 34, 153, 228}
	theirPublicKey := generatePublicKey(theirPrivateKey)
	fmt.Println("Their private key: ", theirPrivateKey)
	fmt.Println("Their public key: ", theirPublicKey)

	sharedSecret1 := generateSharedSecret(theirPrivateKey, myPublicKey)
	sharedSecret2 := generateSharedSecret(myPrivateKey, theirPublicKey)

	fmt.Println("Shared secret 1: ", sharedSecret1)
	fmt.Println("Shared secret 2: ", sharedSecret2)
}
