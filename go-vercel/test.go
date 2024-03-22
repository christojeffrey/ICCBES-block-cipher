package main

import (
	"fmt" // Use fmt for formatted printing

	BCA "ICCBES/block-cipher-algorithm"
	cipherMode "ICCBES/block-cipher-mode"
	"ICCBES/lib/constant"
	"ICCBES/lib/utils"
)

func a() {
	key := utils.GenerateRandomByte(constant.KeyByteSize)
	message := utils.GenerateRandomByte(constant.MessageByteSize)

	fmt.Println("Key:", string(key))
	fmt.Println("Message:", string(message))
	utils.PrintDivider()

	// Testing ECB
	fmt.Println("=== ECB mode ===")
	cipherText := cipherMode.EncryptECB(message, key, BCA.EncryptionAlgorithm)
	fmt.Println("CipherText:", string(cipherText))
	plainText := cipherMode.DecryptECB(cipherText, key, BCA.DecryptionAlgorithm)
	fmt.Println("PlainText:", string(plainText))
	utils.PrintDivider()

	// Testing CBC
	fmt.Println("=== CBC mode ===")
	iv := utils.GenerateRandomByte(constant.MessageBlockByteSize)
	fmt.Println("IV:", string(iv))
	cipherText = cipherMode.EncryptCBC(message, key, BCA.EncryptionAlgorithm, iv)
	fmt.Println("CipherText:", string(cipherText))
	plainText = cipherMode.DecryptCBC(cipherText, key, BCA.DecryptionAlgorithm, iv)
	fmt.Println("PlainText:", string(plainText))
	utils.PrintDivider()

	// Testing CFB
	fmt.Println("=== CFB mode ===")
	iv = utils.GenerateRandomByte(constant.MessageBlockByteSize) // Reuse the generated iv
	cipherText = cipherMode.EncryptCFB(message, key, BCA.EncryptionAlgorithm, iv)
	fmt.Println("CipherText:", string(cipherText))
	plainText = cipherMode.DecryptCFB(cipherText, key, BCA.DecryptionAlgorithm, iv)
	fmt.Println("PlainText:", string(plainText))
	utils.PrintDivider()

	// Testing OFB
	fmt.Println("=== OFB mode ===")
	iv = utils.GenerateRandomByte(constant.MessageBlockByteSize) // Reuse the generated iv (assuming valid for OFB)
	cipherText = cipherMode.EncryptOFB(message, key, BCA.EncryptionAlgorithm, iv)
	fmt.Println("CipherText:", string(cipherText))
	plainText = cipherMode.DecryptOFB(cipherText, key, BCA.DecryptionAlgorithm, iv)
	fmt.Println("PlainText:", string(plainText))
	utils.PrintDivider()

	// Testing CTR
	fmt.Println("=== CTR mode ===")
	iv = utils.GenerateRandomByte(constant.MessageBlockByteSize) // Reuse the generated iv (assuming valid for CTR)
	cipherText = cipherMode.EncryptCounter(message, key, BCA.EncryptionAlgorithm, iv)
	fmt.Println("CipherText:", string(cipherText))
	plainText = cipherMode.DecryptCounter(cipherText, key, BCA.DecryptionAlgorithm, iv)
	fmt.Println("PlainText:", string(plainText))
}
