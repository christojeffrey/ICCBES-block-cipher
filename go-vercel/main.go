package main

import "math/rand"

const MessageBitSize = 128
const KeyBitSize = 128
const Rounds = 10

func main(){
	key := generateRandomByte(KeyBitSize/8)
	message := generateRandomByte(MessageBitSize/8)
	println("Key: ", string(key))
	println("Message: ", string(message))

	cipherText := encryptECB(message, key, simpleEncryptionAlgorithm)
	println("CipherText: ", string(cipherText))

	decryptedMessage := decryptECB(cipherText, key, simpleDecryptionAlgorithm)
	println("DecryptedMessage: ", string(decryptedMessage))

}

type encryptionAlgorithm func([]byte, []byte) []byte
// type decryptionAlgorithm func([]byte, []byte) []byte
func simpleEncryptionAlgorithm(message []byte, key []byte) []byte {
	// will XOR bit by bit the message with the key
	cipherText := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		cipherText[i] = message[i] ^ key[i]
	}
	return cipherText
}

func simpleDecryptionAlgorithm(cipherText []byte, key []byte) []byte {
	// will XOR bit by bit the cipherText with the key
	message := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i++ {
		message[i] = cipherText[i] ^ key[i]
	}
	return message
}



func encryptECB(message []byte, key []byte, encryptionAlgorithm encryptionAlgorithm) []byte {
	// encrypt message with key using ECB mode
	cipherText := encryptionAlgorithm(message, key)
	return cipherText
}

func decryptECB(cipherText []byte, key []byte, decryptionAlgorithm encryptionAlgorithm) []byte {
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
