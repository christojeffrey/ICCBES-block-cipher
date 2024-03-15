package blockCipherAlgorithm

// will XOR bit by bit the plainTextBlock with the key
func SimpleEncryptionAlgorithm(plainTextBlock []byte, key []byte) []byte {
	cipherTextBlock := make([]byte, len(plainTextBlock))
	for i := 0; i < len(plainTextBlock); i++ {
		cipherTextBlock[i] = plainTextBlock[i] ^ key[i]
	}
	return cipherTextBlock
}

// will XOR bit by bit the cipherTextBlock with the key
func SimpleDecryptionAlgorithm(cipherTextBlock []byte, key []byte) []byte {
	plainTextBlock := make([]byte, len(cipherTextBlock))
	for i := 0; i < len(cipherTextBlock); i++ {
		plainTextBlock[i] = cipherTextBlock[i] ^ key[i]
	}
	return plainTextBlock
}
