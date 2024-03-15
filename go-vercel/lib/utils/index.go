package utils

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomByte(n int) []byte {
	// return random letter as byte with the size of n
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
	return b
}

// split either plaintext or ciphertext into blocks. if there are leftover, fill with 0
func TextToBlocks(text []byte, blockSize int) [][]byte {
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

func DoBitXOR(a []byte, b []byte) []byte {
	// will XOR bit by bit the a with b
	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func PrintDivider() {
	println("--------------------------------------------------")
}