package main

import (
	BCA "ICCBES/block-cipher-algorithm"
	cipherMode "ICCBES/block-cipher-mode"
	"ICCBES/lib/constant"
	"ICCBES/lib/utils"
)


func main(){
	key := utils.GenerateRandomByte(constant.KeyByteSize)
	message := utils.GenerateRandomByte(constant.MessageByteSize)
	println("Key: ", string(key))
	println("Message: ", string(message))
	utils.PrintDivider()

	// testing ECB
	println("=== ECB mode ===")

	cipherText := cipherMode.EncryptECB(message, key, BCA.SimpleEncryptionAlgorithm)
	println("CipherText: ", string(cipherText))

	plainText := cipherMode.DecryptECB(cipherText, key, BCA.SimpleDecryptionAlgorithm)
	println("plainText: ", string(plainText))
	utils.PrintDivider()

	// testing CBC
	println("=== CBC mode ===")
	iv := utils.GenerateRandomByte(constant.MessageByteSize)
	cipherText = cipherMode.EncryptCBC(message, key, BCA.SimpleEncryptionAlgorithm, iv)
	println("CipherText: ", string(cipherText))
	plainText = cipherMode.DecryptCBC(cipherText, key, BCA.SimpleDecryptionAlgorithm, iv)
	println("PlainText: ", string(plainText))

}





