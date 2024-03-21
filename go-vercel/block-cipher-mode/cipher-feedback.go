package blockCipherMode

import (
	"ICCBES/lib"
)

// EncryptCFB encrypts plaintext using CFB mode
func EncryptCFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Initialize ciphertext slice
	cipherText := make([]byte, len(plainText))
	// Initialize feedback block with IV
	feedbackBlock := iv
	// Encrypt plaintext byte by byte using CFB mode
	for i, b := range plainText {
		// Encrypt feedback block
		encryptedFeedbackBlock := encryptionAlgorithm(feedbackBlock, key)
		// XOR plaintext byte with encrypted feedback block
		cipherText[i] = encryptedFeedbackBlock[0] ^ b
		// Update feedback block by removing the first byte and appending the ciphertext byte
		feedbackBlock = append(feedbackBlock[1:], cipherText[i])
	}
	return cipherText
}

// DecryptCFB decrypts ciphertext using CFB mode
func DecryptCFB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm, iv []byte) []byte {
	// Initialize plaintext slice
	plainText := make([]byte, len(cipherText))
	// Initialize feedback block with IV
	feedbackBlock := iv
	// Decrypt ciphertext byte by byte using CFB mode
	for i, b := range cipherText {
		// Encrypt feedback block
		encryptedFeedbackBlock := decryptionAlgorithm(feedbackBlock, key)
		// XOR ciphertext byte with encrypted feedback block
		plainText[i] = encryptedFeedbackBlock[0] ^ b
		// Update feedback block by removing the first byte and appending the decrypted ciphertext byte
		feedbackBlock = append(feedbackBlock[1:], b)
	}
	return plainText
}
