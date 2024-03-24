package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
	"time"
)

// EncryptOFB encrypts plaintext using OFB mode. EFB only has encryption, decryption is the same as encryption
func EncryptOFB(plainText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, iv []byte) []byte {
	startTime := time.Now()
	// split plainText into blocks
	plainTextBlocks := utils.TextToBlocks(plainText)
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)
	encryptedIV := iv
	// Encrypt plaintext byte by byte using CFB mode
	for i := 0; i < blockLength; i++ {
		// encrypt iv with key
		encryptedIV := encryptionAlgorithm(encryptedIV, key)
		// update iv
		// XOR with plaintext
		cipherTextBlocks[i] = utils.DoBitXOR(plainTextBlocks[i], encryptedIV)
	}
	// merge blocks into one
	cipherText := utils.MergeBlocksIntoOneString(cipherTextBlocks, len(plainText))
	elapsedTime := time.Since(startTime)
	println("elapsed time EncryptOFB in ns: ", elapsedTime.Nanoseconds())
	return cipherText
}

