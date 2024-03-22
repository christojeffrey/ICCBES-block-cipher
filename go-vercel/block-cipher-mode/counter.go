package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
)

// EncryptCTR encrypts plaintext using CTR mode
func EncryptCounter(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Initialize ciphertext slice
	cipherText := make([]byte, len(plainText))
	// Initialize counter with IV
	counter := make([]byte, len(iv))
	copy(counter, iv)
	// Encrypt plaintext byte by byte using CTR mode
	for i, b := range plainText {
		// Encrypt counter block
		encryptedCounter := encryptionAlgorithm(counter, key)
		// XOR plaintext byte with encrypted counter block
		cipherText[i] = encryptedCounter[0] ^ b
		// Increment counter
		utils.IncrementCounter(counter)
	}
	return cipherText
}

// DecryptCTR decrypts ciphertext using CTR mode
func DecryptCounter(cipherText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Initialize plaintext slice
	plainText := make([]byte, len(cipherText))
	// Initialize counter with IV
	counter := make([]byte, len(iv))
	copy(counter, iv)
	// Decrypt ciphertext byte by byte using CTR mode
	for i, b := range cipherText {
		// Encrypt counter block
		encryptedCounter := encryptionAlgorithm(counter, key)
		// XOR ciphertext byte with encrypted counter block
		plainText[i] = encryptedCounter[0] ^ b
		// Increment counter
		utils.IncrementCounter(counter)
	}
	return plainText
}
