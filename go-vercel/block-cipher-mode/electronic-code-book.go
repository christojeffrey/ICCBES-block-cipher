package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
	"time"
)

func EncryptECB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm) []byte {
	startTime := time.Now()
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
	elapsedTime := time.Since(startTime)
	println("elapsed time EncryptECB in ns: ", elapsedTime.Nanoseconds())
	return cipherText
}

func DecryptECB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm) []byte {
	startTime := time.Now()
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
	elapsedTime := time.Since(startTime)
	println("elapsed time DecryptECB in ns: ", elapsedTime.Nanoseconds())
	return plainText
}