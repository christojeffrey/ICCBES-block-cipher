package lib

type EncryptionAlgorithm func([]byte, []byte) []byte
type DecryptionAlgorithm func([]byte, []byte) []byte