package main

import (
	"math/rand"

	"golang.org/x/crypto/curve25519"
)

func generatePublicKey(privateKey [32]byte) [32]byte {
	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	return publicKey
}

func generateRandomBuffer(n int) []byte {
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, byte(rand.Intn(0xff-0+1)))
	}
	return result
}
func generateSharedSecret(myPrivateKey [32]byte, theirPublicKey [32]byte) []byte {
	var sharedSecret [32]byte
	curve25519.ScalarMult(&sharedSecret, &myPrivateKey, &theirPublicKey)
	return sharedSecret[:]
}
