package utils

import (
	"ICCBES/lib/constant"
	"encoding/base64"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomByte(n int) []byte {
	// return random letter as byte with the size of n
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
	return b
}

func TransmissionEncoding(text []byte) string{
	// return string(text)
	return base64.StdEncoding.EncodeToString(text)
}
func TransmissionDecoding(text string) []byte{
	// return []byte(text)
	decoded, _ := base64.StdEncoding.DecodeString(text)
	return decoded
}

// split either plaintext or ciphertext into blocks. if there are leftover, fill with 0
func TextToBlocks(text []byte) [][]byte {
	blockSize := constant.MessageBlockByteSize
	fillerByte := byte(0)
	// ceiling
	blockLength := (len(text) + blockSize - 1) / blockSize
	blocks := make([][]byte, blockLength)
	totalSize := constant.MessageBlockByteSize * blockLength

	counter := 0
	for i := 0; i < blockLength; i++ {
		// setup block to be encrypted
		// if block is not full, fill with 0
		block := make([]byte, blockSize)
		for j := 0; j < blockSize; j++ {
			if totalSize - (i*blockSize+j) <= len(text) {
				block[j] = text[counter]
				counter++
			} else {
				block[j] = fillerByte
			}
		}
		blocks[i] = block
	}
	return blocks
}

func DoBitXOR(a []byte, b []byte) []byte {
	// will XOR bit by bit the a with b
	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func LeftRotateNTimes(word []byte, n int) []byte {
	// Left rotate byte order n times
	for i := 0; i < n; i++ {
		temp := word[0]
		for j := 0; j < len(word)-1; j++ {
			word[j] = word[j+1]
		}
		word[len(word)-1] = temp
	}
	return word
}

func PrintDivider() {
	println("--------------------------------------------------")
}

func MergeBlocksIntoOneString(blocks [][]byte, outputTextLength int) []byte{
	// merge blocks into one
	outputText := make([]byte, constant.MessageBlockByteSize * len(blocks))

	for p := 0; p < constant.MessageBlockByteSize * len(blocks); p++ {
		i := p / constant.MessageBlockByteSize;
		j := p % constant.MessageBlockByteSize;
		outputText[p] = blocks[i][j];			
	}

	// remove byte(0) in front of the outputText
	// find start from
	startFrom := 0
	for i := 0; i < len(outputText); i++ {
		if outputText[i] != byte(0) {
			startFrom = i
			break
		}
	}
	newOutputText := outputText[startFrom:]
	// return the result
	return newOutputText

}