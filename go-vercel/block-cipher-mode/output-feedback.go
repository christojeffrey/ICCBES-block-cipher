package blockCipherMode

import (
	"ICCBES/lib" // Likely incorrect, replace with the actual package name
	"ICCBES/lib/utils"
)

func EncryptOFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Split plainText into blocks (same size as key)
	plainTextBlocks := utils.TextToBlocks(plainText, len(key))
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)
	prevCipherBlock := iv

	for i := 0; i < blockLength; i++ {
		currentBlock := plainTextBlocks[i]

		// Generate keystream (entire encrypted previous block)
		keystream := encryptOFBBlock(prevCipherBlock, encryptionAlgorithm)

		// Encrypt block using XOR with keystream
		currentBlock = utils.DoBitXOR(currentBlock, keystream)

		// Save the result
		cipherTextBlocks[i] = currentBlock
		prevCipherBlock = keystream // Update for next iteration (OFB characteristic)
	}

	// Merge blocks into one
	cipherText := make([]byte, len(plainText))
	for i := 0; i < blockLength; i++ {
		for j := 0; j < len(key); j++ {
			cipherText[i*len(key)+j] = cipherTextBlocks[i][j]
		}
	}
	return cipherText
}

func DecryptOFB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm, iv []byte) []byte {
	// Split cipherText into blocks (same size as key)
	cipherTextBlocks := utils.TextToBlocks(cipherText, len(key))
	blockLength := len(cipherTextBlocks)
	plainTextBlocks := make([][]byte, blockLength)
	prevCipherBlock := iv

	for i := 0; i < blockLength; i++ {
		currentBlock := cipherTextBlocks[i]

		// Generate keystream (entire encrypted previous block)
		keystream := decryptOFBBlock(prevCipherBlock, decryptionAlgorithm)

		// Decrypt block using XOR with keystream
		currentBlock = utils.DoBitXOR(currentBlock, keystream)

		// Save the result
		plainTextBlocks[i] = currentBlock
		prevCipherBlock = keystream // Update for next iteration (OFB characteristic)
	}

	// Merge blocks into one
	plainText := make([]byte, len(cipherText))
	for i := 0; i < blockLength; i++ {
		for j := 0; j < len(key); j++ {
			plainText[i*len(key)+j] = plainTextBlocks[i][j]
		}
	}
	return plainText
}

func encryptOFBBlock(prevCipherBlock []byte, encryptionAlgorithm lib.EncryptionAlgorithm) []byte {
	// Encrypt previous cipher block
	encryptedBlock := encryptionAlgorithm(prevCipherBlock, []byte{})

	// Return entire encrypted block as keystream
	return encryptedBlock
}
func decryptOFBBlock(prevCipherBlock []byte, decryptionAlgorithm lib.DecryptionAlgorithm) []byte {
	// Encrypt previous cipher block
	encryptedBlock := decryptionAlgorithm(prevCipherBlock, []byte{})

	// Return entire encrypted block as keystream
	return encryptedBlock
}
