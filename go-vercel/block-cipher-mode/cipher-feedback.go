package blockCipherMode

import (
	"ICCBES/lib"
	"time"
)

// EncryptCFB encrypts plaintext using CFB mode. this CFB will encrypt byte per byte
func EncryptCFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	startTime := time.Now()
	// split plainText into blocks
	cipherText := make([]byte, len(plainText))
	encryptedIV := iv
	for i := 0; i < len(plainText); i++ {
		// encrypt iv with key
		encryptedIV = encryptionAlgorithm(encryptedIV, key)
		// XOR with plaintext
		cipherText[i] = plainText[i] ^ encryptedIV[0]
		// add cipherTextBytes[i] to right most of IV
		encryptedIV = append(encryptedIV[0: len(encryptedIV) - 1], cipherText[i])
	}
	elapsedTime := time.Since(startTime)
	println("elapsed time EncryptCFB in ns: ", elapsedTime.Nanoseconds())
	return cipherText

	
}
// EncryptCFB encrypts plaintext using CFB mode. this CFB will encrypt byte per byte. both uses encryption alogrithm
func DecryptCFB(cipherText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	startTime := time.Now()
	// split plainText into blocks
	plainText := make([]byte, len(cipherText))
	encryptedIV := iv
	for i := 0; i < len(plainText); i++ {
		// encrypt iv with key
		encryptedIV = encryptionAlgorithm(encryptedIV, key)
		// XOR with plaintext
		plainText[i] = cipherText[i] ^ encryptedIV[0]
		// add cipherTextBytes[i] to right most of IV
		encryptedIV = append(encryptedIV[0: len(encryptedIV) - 1], cipherText[i])
	}
	// combine all cipherTextBytes into one
	elapsedTime := time.Since(startTime)
	println("elapsed time DecryptCFB in ns: ", elapsedTime.Nanoseconds())
	return plainText	
}