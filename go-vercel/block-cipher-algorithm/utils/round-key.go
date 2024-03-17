package utils

import (
	"ICCBES/block-cipher-algorithm/constant"
	"ICCBES/lib/utils"
)

// gFunction will take 4 byte and return 4 byte
func gFunction(word []byte, round int) []byte {
	// Left rotate by 2
	word = utils.LeftRotateNTimes(word, 2)
	// Substitute bytes
	word = SubstituteBytes(word)
	// Xor with round constant
	roundConstantByte := []byte{constant.RoundConstant[round], 0x00, 0x00, 0x00}
	word = utils.DoBitXOR(word, roundConstantByte)
	return word
}

// Takes a 128 bit key and returns 52 words (13 * 4 words, 1 word is 16 byte)
func GenerateRoundKeys(key []byte) [][]byte {
	roundKeys := make([][]byte, 13)
	for i := range roundKeys {
		roundKeys[i] = make([]byte, 16)
	}
	// First 16 bytes is the key, spread it [k1, k4, k8, k12, k2, k5, k9, k13, k3, k6, k10, k14, k4, k7, k11, k15] for easier access later
	for i := 0; i < 4; i++ {
		roundKeys[0][i*4] = key[i]
		roundKeys[0][i*4 + 1] = key[i + 4]
		roundKeys[0][i*4 + 2] = key[i + 8]
		roundKeys[0][i*4 + 3] = key[i + 12]
	}
	// Generate 12 round keys
	for i := 1; i <= 12 ; i++ {
		var tempRoundKey []byte
		// Put last words into g function
		gFuncKey := gFunction(roundKeys[i-1][12:16], i)
		for j := 0; j < 4; j++ {
			if (j == 0) {
				tempRoundKey = append(tempRoundKey, utils.DoBitXOR(gFuncKey, roundKeys[i-1][0:4])...)
			} else {
				tempRoundKey = append(tempRoundKey, utils.DoBitXOR(tempRoundKey[(j-1)*4:(j-1)*4+4], roundKeys[i-1][j*4:j*4+4])...)
			}
		}
		roundKeys[i] = tempRoundKey
	}
 	return roundKeys
}