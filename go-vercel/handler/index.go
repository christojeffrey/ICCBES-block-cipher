package handler

import (
	BCA "ICCBES/block-cipher-algorithm"
	cipherMode "ICCBES/block-cipher-mode"
	constants "ICCBES/lib/constant"
	"errors"
	"strconv"
)

// will handle encryption and decryption for ECB mode mode must be "encrypt" or "decrypt"
func CBCHandler(mode string, body map[string]interface{}) map[string]interface{} {
	// TODO: add limitation such as length and character
	// validate body: text string and iv string, key string
	err := validateBody(body)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	// get text and iv from body
	text := []byte(body["text"].(string))
	iv := []byte(body["iv"].(string))
	key := []byte(body["key"].(string))
	// encrypt or decrypt
	var result []byte
	println("asdfkadsfljk")
	if mode == "encrypt" {
		result = cipherMode.EncryptCBC(text, key, BCA.EncryptionAlgorithm, iv)
	} else {
		result = cipherMode.DecryptCBC(text, key, BCA.DecryptionAlgorithm, iv)
	}
	println("2asdfkadsfljk")
	return map[string]interface{}{"result": string(result)}


	
}
func validateBody(body map[string]interface{}) error {
	if _, ok := body["text"]; !ok {
		return errors.New("text is required")
	}
	if _, ok := body["iv"]; !ok {
		return errors.New("iv is required")
	}
	if _, ok := body["key"]; !ok {
		return errors.New("key is required")
	}
	// validate their length
	if len(body["key"].(string)) != constants.KeyByteSize {
		return errors.New("key length must be " + strconv.Itoa(constants.KeyByteSize )+ " bytes (or characters)")
	}
	if len(body["iv"].(string)) != constants.MessageBlockByteSize {
		return errors.New("iv length must be " + strconv.Itoa(constants.MessageBlockByteSize) + " bytes (or characters)")
	}

	return nil
}