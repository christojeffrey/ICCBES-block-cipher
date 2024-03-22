package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
	"time"
)

// encrypt plainText with key using CBC mode
func EncryptCBC(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	startTime := time.Now()
	// split plainText into blocks
	plainTextBlocks := utils.TextToBlocks(plainText)
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)
	previousCipherTextBlock := iv
	for i := 0; i < blockLength; i++ {
		currentBlock := plainTextBlocks[i]


		// XOR with previous cipher text
		currentBlock = utils.DoBitXOR(currentBlock, previousCipherTextBlock)
		
		// encrypt block
		currentBlock = encryptionAlgorithm(currentBlock, key)

		// save the result
		cipherTextBlocks[i] = currentBlock

		// save the previous cipher text
		previousCipherTextBlock = currentBlock
	}

	// merge blocks into one
	cipherText := utils.MergeBlocksIntoOneString(cipherTextBlocks, len(plainText));

	elapsedTime := time.Since(startTime)
	println("elapsed time EncryptCBC in ns: ", elapsedTime.Nanoseconds())
	return cipherText
}
// decrypt cipherText with key using CBC mode
func DecryptCBC(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm, iv []byte) []byte {
	startTime := time.Now()
	// split cipherText into blocks
	cipherTextBlocks := utils.TextToBlocks(cipherText)
	blockLength := len(cipherTextBlocks)
	plainTextBlocks := make([][]byte, blockLength)
	previousPlainTextBlock := iv
	for i := 0; i < blockLength; i++ {
		currentBlock := cipherTextBlocks[i]
		
		// decyrpt block
		currentBlock = decryptionAlgorithm(currentBlock, key)

		// XOR with previous plain text
		currentBlock = utils.DoBitXOR(currentBlock, previousPlainTextBlock)

		// save the result
		plainTextBlocks[i] = currentBlock

		// save the previous cipher text
		previousPlainTextBlock = cipherTextBlocks[i]
	}

	// merge blocks into one
	plainText := utils.MergeBlocksIntoOneString(plainTextBlocks, len(cipherText))

	elapsedTime := time.Since(startTime)
	println("elapsed time DecryptCBC in ns: ", elapsedTime.Nanoseconds())
	return plainText
}