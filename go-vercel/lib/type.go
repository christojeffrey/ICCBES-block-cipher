package lib

type EncryptionAlgorithm func(plainText []byte, key []byte) []byte
type DecryptionAlgorithm func(cipherText []byte, key []byte) []byte