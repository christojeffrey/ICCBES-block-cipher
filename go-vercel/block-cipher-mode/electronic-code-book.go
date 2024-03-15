package blockCipherMode

import "ICCBES/lib"

func EncryptECB(message []byte, key []byte, encryptionAlgorithm lib.EncryptionAlgorithm) []byte {
	// encrypt message with key using ECB mode
	cipherText := encryptionAlgorithm(message, key)
	return cipherText
}

func DecryptECB(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm) []byte {
	// decrypt message with key using ECB mode
	message := decryptionAlgorithm(cipherText, key)
	return message
}