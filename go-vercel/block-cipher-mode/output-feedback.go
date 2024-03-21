package blockCipherMode

import "ICCBES/lib"

// EncryptOFB encrypts plaintext using OFB mode
func EncryptOFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Initialize ciphertext slice
	cipherText := make([]byte, len(plainText))
	// Initialize feedback register with IV
	feedback := make([]byte, len(iv))
	copy(feedback, iv)
	// Encrypt plaintext byte by byte using OFB mode
	for i, b := range plainText {
		// Encrypt feedback register
		encryptedFeedback := encryptionAlgorithm(feedback, key)
		// XOR plaintext byte with encrypted feedback
		cipherText[i] = encryptedFeedback[0] ^ b
		// Update feedback register with encrypted feedback
		copy(feedback, encryptedFeedback)
	}
	return cipherText
}

// DecryptOFB decrypts ciphertext using OFB mode
func DecryptOFB(cipherText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Initialize plaintext slice
	plainText := make([]byte, len(cipherText))
	// Initialize feedback register with IV
	feedback := make([]byte, len(iv))
	copy(feedback, iv)
	// Decrypt ciphertext byte by byte using OFB mode
	for i, b := range cipherText {
		// Encrypt feedback register
		encryptedFeedback := encryptionAlgorithm(feedback, key)
		// XOR ciphertext byte with encrypted feedback
		plainText[i] = encryptedFeedback[0] ^ b
		// Update feedback register with encrypted feedback
		copy(feedback, encryptedFeedback)
	}
	return plainText
}
