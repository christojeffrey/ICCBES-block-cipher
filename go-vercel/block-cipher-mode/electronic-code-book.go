package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
)

func EncryptECB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm) []byte {
	// split plainText into blocks
	plainTextBlocks := utils.TextToBlocks(plainText)
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)
	for i := 0; i < blockLength; i++ {
		currentBlock := plainTextBlocks[i]
	
		// encrypt block
		currentBlock = encryptionAlgorithm(currentBlock, key)

		// save the result
		cipherTextBlocks[i] = currentBlock
	}

	// merge blocks into one
	cipherText := make([]byte, len(plainText))
	for i := 0; i < blockLength; i++ {
		for j := 0; j < len(key); j++ {
			cipherText[i*len(key)+j] = cipherTextBlocks[i][j]
		}
	}
	return cipherText
}

func DecryptECB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm) []byte {
	// split cipherText into blocks
	cipherTextBlocks := utils.TextToBlocks(cipherText)
	blockLength := len(cipherTextBlocks)
	plainTextBlocks := make([][]byte, blockLength)
	for i := 0; i < blockLength; i++ {
		currentBlock := cipherTextBlocks[i]
		
		// decyrpt block
		currentBlock = decryptionAlgorithm(currentBlock, key)

		// save the result
		plainTextBlocks[i] = currentBlock
	}

	// merge blocks into one
	plainText := make([]byte, len(cipherText))
	for i := 0; i < blockLength; i++ {
		for j := 0; j < len(key); j++ {
			plainText[i*len(key)+j] = plainTextBlocks[i][j]
		}
	}
	return plainText
}