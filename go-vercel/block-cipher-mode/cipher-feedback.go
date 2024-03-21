package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
)

const (
	CFB_MODE = "CFB"
)

func EncryptCFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	// Split plainText into blocks
	plainTextBlocks := utils.TextToBlocks(plainText, len(key))
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)
	prevCipherBlock := iv
	for i := 0; i < blockLength; i++ {
		currentBlock := plainTextBlocks[i]

		// Generate keystream
		keystream := encryptCFBBlock(prevCipherBlock, encryptionAlgorithm)

		// Encrypt block
		currentBlock = utils.DoBitXOR(currentBlock, keystream)
		prevCipherBlock = currentBlock

		// Save the result
		cipherTextBlocks[i] = currentBlock
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

func DecryptCFB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm, iv []byte) []byte {
	// Split cipherText into blocks
	cipherTextBlocks := utils.TextToBlocks(cipherText, len(key))
	blockLength := len(cipherTextBlocks)
	plainTextBlocks := make([][]byte, blockLength)
	prevCipherBlock := iv
	for i := 0; i < blockLength; i++ {
		currentBlock := cipherTextBlocks[i]

		// Generate keystream
		keystream := encryptCFBBlock(prevCipherBlock, key)

		// Decrypt block
		currentBlock = utils.DoBitXOR(currentBlock, keystream)
		prevCipherBlock = cipherTextBlocks[i]

		// Save the result
		plainTextBlocks[i] = currentBlock
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

func encryptCFBBlock(prevCipherBlock []byte, encryptionAlgorithm lib.EncryptionAlgorithm) []byte {
	// Encrypt previous cipher block
	encryptedBlock := encryptionAlgorithm(prevCipherBlock, []byte{})

	// Return first N bits of encrypted block
	return encryptedBlock[:len(prevCipherBlock)]
}
