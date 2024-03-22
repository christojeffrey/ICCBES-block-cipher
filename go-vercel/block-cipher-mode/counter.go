package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/utils"
	"time"
)

// EncryptCTR encrypts plaintext using CTR mode. decrypt and encrypt is the same in counter mode
func EncryptCounter(inputText []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm, counter []byte) []byte {	
	startTime := time.Now()
	// split inputText into blocks
	inputTextBlocks := utils.TextToBlocks(inputText)
	blockLength := len(inputTextBlocks)
	outputTextBlocks := make([][]byte, blockLength)
	encryptedCounter := counter
	for i := 0; i < blockLength; i++ {
		// encrypt counter with key
		encryptedCounter = encryptionAlgorithm(encryptedCounter, key)
		// XOR with inputText
		outputTextBlocks[i] = utils.DoBitXOR(inputTextBlocks[i], encryptedCounter)
		// increment counter
		encryptedCounter = incrementCounter(encryptedCounter)
	}
	// merge blocks into one
	outputText := utils.MergeBlocksIntoOneString(outputTextBlocks, len(inputText))
	println("outputText: ", string(outputText))

	elapsedTime := time.Since(startTime)
	println("elapsed time EncryptCounter in ns: ", elapsedTime.Nanoseconds())
	return outputText
}


func incrementCounter(counter []byte)[]byte{
	// add every byte in counter by 1
	for i := len(counter) - 1; i >= 0; i-- {
		counter[i]++
		if counter[i] != 0 {
			break
		}
	}
	return counter
}
