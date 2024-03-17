package blockCipherAlgorithm
/*
Requirements
1. Menerapkan prinsip diffusion and confusion
2. Mendefinisikan fungsi putaran f yang berisi jaringan substitusi dan permutasi
3. Operasi substitusi dan transposisi (Pake tabel S-box)
4. Menerapkan cipher berulang 10 - 16 kali. Setiap putaran kuncinya beda. Kunci putaran dibangkitkan dari kunci eksternal
5. Panjang kunci 128 bit
*/

// TODO:
// 1. Create a substitution box [x] and permutation box
// 2. Create a function to generate a key for each round [x]
// 3. Create a round function (this is f)
// 4. Create a feistel function that contains the substitution and permutation network (implementation of number 3)
// 5. Create a function to encrypt and decrypt the message


// will XOR bit by bit the plainTextBlock with the key
func SimpleEncryptionAlgorithm(plainTextBlock []byte, key []byte) []byte {
	cipherTextBlock := make([]byte, len(plainTextBlock))
	for i := 0; i < len(plainTextBlock); i++ {
		cipherTextBlock[i] = plainTextBlock[i] ^ key[i]
	}
	return cipherTextBlock
}

// will XOR bit by bit the cipherTextBlock with the key
func SimpleDecryptionAlgorithm(cipherTextBlock []byte, key []byte) []byte {
	plainTextBlock := make([]byte, len(cipherTextBlock))
	for i := 0; i < len(cipherTextBlock); i++ {
		plainTextBlock[i] = cipherTextBlock[i] ^ key[i]
	}
	return plainTextBlock
}
