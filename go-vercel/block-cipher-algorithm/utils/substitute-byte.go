package utils

import "ICCBES/block-cipher-algorithm/constant"

func getFirstNibble(b byte) byte {
	return b >> 4
}

func getSecondNibble(b byte) byte {
	return b & 0x0f
}

func SubstituteBytes(word []byte) []byte {
	// Substitute bytes
	for i := 0; i < len(word); i++ {
		word[i] = constant.SBox[getFirstNibble(word[i])][getSecondNibble(word[i])]
	}
	return word
}

func InverseSubstituteBytes(word []byte) []byte {
	// Inverse substitute bytes
	for i := 0; i < len(word); i++ {
		word[i] = constant.InverseSBox[getFirstNibble(word[i])][getSecondNibble(word[i])]
	}
	return word
}