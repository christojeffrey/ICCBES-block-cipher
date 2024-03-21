package blockCipherAlgorithm

import (
	BCAUtils "ICCBES/block-cipher-algorithm/utils"
	"ICCBES/lib/utils"
)

// TODO:
// 1. Make F function more complicated

func divideLeftRightBlock(block []byte) ([]byte, []byte) {
	return block[:len(block)/2], block[len(block)/2:]
}

func swapLeftRightBlock(leftBlock []byte, rightBlock []byte) ([]byte, []byte) {
	return rightBlock, leftBlock
}

// This function receive 128 bit key and 64 bit block, returns a 64 bit block
func F(key []byte, block []byte) []byte {
	// Divide key into two parts
	leftKey, rightKey := divideLeftRightBlock(key)
	// Xor the block with left and right part of key
	block = utils.DoBitXOR(block, leftKey)
	block = utils.DoBitXOR(block, rightKey)

	// Rotate the block 3 times to the left (permutation)
	block = utils.LeftRotateNTimes(block, 3)

	// Substitute the block using S-box
	block = BCAUtils.SubstituteBytes(block)

	return block
}
// Encrypt the plainTextBlock using the key by feistel network
func EncryptionAlgorithm(plainTextBlock []byte, key []byte) []byte {
	// Divide the block
	leftBlock, rightBlock := divideLeftRightBlock(plainTextBlock)
	// Generate round keys
	roundKeys := BCAUtils.GenerateRoundKeys(key)

	// Feistel network
	for i := 0; i < 12; i++ {
		// F function
		fResult := F(roundKeys[i + 1], rightBlock) // We don't need the first round key
		// XOR with left block
		leftBlock = utils.DoBitXOR(leftBlock, fResult)
		// Swap left and right block
		leftBlock, rightBlock = swapLeftRightBlock(leftBlock, rightBlock)
	}
	
	// Combine the block
	return append(leftBlock, rightBlock...)
}

// Decrypt the cipherTextBlock using the key by feistel network
func DecryptionAlgorithm(cipherTextBlock []byte, key []byte) []byte {
	// Divide the block
	leftBlock, rightBlock := divideLeftRightBlock(cipherTextBlock)
	// Generate round keys
	roundKeys := BCAUtils.GenerateRoundKeys(key)

	// Feistel network
	for i := 0; i < 12; i++ {
		// F function
		fResult := F(roundKeys[12 - i], leftBlock)
		// XOR with right block
		rightBlock = utils.DoBitXOR(rightBlock, fResult)
		// Swap left and right block
		leftBlock, rightBlock = swapLeftRightBlock(leftBlock, rightBlock)
	}
	
	// Combine the block
	return append(leftBlock, rightBlock...)
}
