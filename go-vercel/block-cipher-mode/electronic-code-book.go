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

	// merge blocks into one. create cipher text the size of block length
		// merge blocks into one
	cipherText := utils.MergeBlocksIntoOneString(cipherTextBlocks, len(plainText));
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
	plainText := utils.MergeBlocksIntoOneString(plainTextBlocks, len(cipherText))
	return plainText
}