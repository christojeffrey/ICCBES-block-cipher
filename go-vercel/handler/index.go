package handler

import (
	BCA "ICCBES/block-cipher-algorithm"
	cipherMode "ICCBES/block-cipher-mode"
)

// will handle encryption and decryption for ECB mode mode must be "encrypt" or "decrypt"
func CBCHandler(mode string, body map[string]interface{}) map[string]interface{} {
	// TODO: add limitation such as length and character
	// validate body: text string and iv string, key string
	if !validateBody(body) {
		return map[string]interface{}{"error": "invalid body. required: text, iv, key"}
	}
	// get text and iv from body
	text := []byte(body["text"].(string))
	iv := []byte(body["iv"].(string))
	key := []byte(body["key"].(string))
	// encrypt or decrypt
	var result []byte
	if mode == "encrypt" {
		result = cipherMode.EncryptCBC(text, key, BCA.EncryptionAlgorithm, iv)
	} else {
		result = cipherMode.DecryptCBC(text, key, BCA.DecryptionAlgorithm, iv)
	}
	return map[string]interface{}{"result": string(result)}


	
}
func validateBody(body map[string]interface{}) bool {
	if _, ok := body["text"]; !ok {
		return false
	}
	if _, ok := body["iv"]; !ok {
		return false
	}
	if _, ok := body["key"]; !ok {
		return false
	}

	return true
}