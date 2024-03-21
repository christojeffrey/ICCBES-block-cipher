package blockCipherMode

import (
	"ICCBES/lib" // Likely incorrect, replace with the actual package name
	"ICCBES/lib/utils"
)

func EncryptOFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Split plainText into blocks
	plainTextBlocks := utils.TextToBlocks(plainText)
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
	cipherText := utils.MergeBlocksIntoOneString(cipherTextBlocks, len(plainText));

	return cipherText
}

func DecryptOFB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm, iv []byte) []byte {
	// Split cipherText into blocks
	cipherTextBlocks := utils.TextToBlocks(cipherText)
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
	plainText := utils.MergeBlocksIntoOneString(plainTextBlocks, len(cipherText));
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
