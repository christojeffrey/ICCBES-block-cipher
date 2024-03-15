package main

import (
	"math/rand"
)

const MessageBitSize = 128
const messageByteSize = MessageBitSize/8
const KeyBitSize = 128
const keyByteSize = KeyBitSize/8
const Rounds = 10

func main(){
	key := generateRandomByte(keyByteSize)
	message := generateRandomByte(messageByteSize)
	println("Key: ", string(key))
	println("Message: ", string(message))
	printDivider()
	println("=== ECB mode ===")

	cipherText := encryptECB(message, key, simpleEncryptionAlgorithm)
	println("CipherText: ", string(cipherText))

	plainText := decryptECB(cipherText, key, simpleDecryptionAlgorithm)
	println("plainText: ", string(plainText))
	printDivider()
	println("=== CBC mode ===")
	iv := generateRandomByte(messageByteSize)
	cipherText = encryptCBC(message, key, simpleEncryptionAlgorithm, iv)
	println("CipherText: ", string(cipherText))
	plainText = decryptCBC(cipherText, key, simpleDecryptionAlgorithm, iv)
	println("PlainText: ", string(plainText))

}
func printDivider() {
	println("--------------------------------------------------")
}
type encryptionAlgorithm func([]byte, []byte) []byte
type decryptionAlgorithm func([]byte, []byte) []byte


// will XOR bit by bit the plainTextBlock with the key
func simpleEncryptionAlgorithm(plainTextBlock []byte, key []byte) []byte {
	cipherTextBlock := make([]byte, len(plainTextBlock))
	for i := 0; i < len(plainTextBlock); i++ {
		cipherTextBlock[i] = plainTextBlock[i] ^ key[i]
	}
	return cipherTextBlock
}

// will XOR bit by bit the cipherTextBlock with the key
func simpleDecryptionAlgorithm(cipherTextBlock []byte, key []byte) []byte {
	plainTextBlock := make([]byte, len(cipherTextBlock))
	for i := 0; i < len(cipherTextBlock); i++ {
		plainTextBlock[i] = cipherTextBlock[i] ^ key[i]
	}
	return plainTextBlock
}
// split either plaintext or ciphertext into blocks. if there are leftover, fill with 0
func textToBlocks(text []byte, blockSize int) [][]byte {
	fillerByte := byte(0)
	// ceiling
	blockLength := (len(text) + blockSize - 1) / blockSize
	blocks := make([][]byte, blockLength)
	for i := 0; i < blockLength; i++ {
		// setup block to be encrypted
		// if block is not full, fill with 0
		block := make([]byte, blockSize)
		for j := 0; j < blockSize; j++ {
			if i*blockSize+j < len(text) {
				block[j] = text[i*blockSize+j]
			} else {
				block[j] = fillerByte
			}
		}
		blocks[i] = block
	}
	return blocks
}

// encrypt plainText with key using CBC mode
func encryptCBC(plainText []byte, key []byte, encryptionAlgorithm encryptionAlgorithm, iv []byte) []byte {

	// split plainText into blocks
	plainTextBlocks := textToBlocks(plainText, len(key))
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)
	previousCipherTextBlock := iv
	for i := 0; i < blockLength; i++ {
		currentBlock := plainTextBlocks[i]
		

		// XOR with previous cipher text
		currentBlock = doBitXOR(currentBlock, previousCipherTextBlock)
		
		// encrypt block
		currentBlock = encryptionAlgorithm(currentBlock, key)

		// save the result
		cipherTextBlocks[i] = currentBlock

		// save the previous cipher text
		previousCipherTextBlock = currentBlock
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
// decrypt cipherText with key using CBC mode
func decryptCBC(cipherText []byte, key []byte, decryptionAlgorithm decryptionAlgorithm, iv []byte) []byte {
	// split cipherText into blocks
	cipherTextBlocks := textToBlocks(cipherText, len(key))
	blockLength := len(cipherTextBlocks)
	plainTextBlocks := make([][]byte, blockLength)
	previousPlainTextBlock := iv
	for i := 0; i < blockLength; i++ {
		currentBlock := cipherTextBlocks[i]
		
		// decyrpt block
		currentBlock = decryptionAlgorithm(currentBlock, key)

		// XOR with previous plain text
		currentBlock = doBitXOR(currentBlock, previousPlainTextBlock)

		// save the result
		plainTextBlocks[i] = currentBlock

		// save the previous cipher text
		previousPlainTextBlock = currentBlock
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
func doBitXOR(a []byte, b []byte) []byte {
	// will XOR bit by bit the a with b
	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}


func encryptECB(message []byte, key []byte, encryptionAlgorithm encryptionAlgorithm) []byte {
	// encrypt message with key using ECB mode
	cipherText := encryptionAlgorithm(message, key)
	return cipherText
}

func decryptECB(cipherText []byte, key []byte, decryptionAlgorithm decryptionAlgorithm) []byte {
	// decrypt message with key using ECB mode
	message := decryptionAlgorithm(cipherText, key)
	return message
}
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomByte(n int) []byte {
	// return random letter as byte with the size of n
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
	return b
}
