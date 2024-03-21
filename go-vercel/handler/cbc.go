package handler

import (
	BCA "ICCBES/block-cipher-algorithm"
	cipherMode "ICCBES/block-cipher-mode"
	"ICCBES/lib/constant"
	"ICCBES/lib/utils"
	"errors"
	"strconv"
)

// will handle encryption and decryption for ECB mode mode must be "encrypt" or "decrypt"
func CBCHandler(mode string, body map[string]interface{}) map[string]interface{} {
	// TODO: add limitation such as length and character
	// validate body: text string and iv string, key string
	isAutoFill, err := validateCBCBody(body)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	
	var key []byte
	var iv []byte
	// handle autofill
	if(isAutoFill){
		key = utils.GenerateRandomByte(constant.KeyByteSize)
		iv = utils.GenerateRandomByte(constant.MessageBlockByteSize)
	}else{
		// get text and iv from body
		iv = []byte(body["iv"].(string))
		key = []byte(body["key"].(string))
	}
	text := utils.TransmissionDecoding(body["text"].(string))
	
	// encrypt or decrypt
	var result []byte
	if mode == "encrypt" {
		result = cipherMode.EncryptCBC(text, key, BCA.EncryptionAlgorithm, iv)
	} else {
		result = cipherMode.DecryptCBC(text, key, BCA.DecryptionAlgorithm, iv)
	}

	return map[string]interface{}{"result": utils.TransmissionEncoding(result), "key" : string(key), "iv": string(iv)}
}
// if autofil is available, only check text. 
func validateCBCBody(body map[string]interface{})  (bool, error) {
	isAutoFilAvail := false
	_, ok := body["autofill"];

	if ok {
		isAutoFilAvail = bool(body["autofill"].(bool))
	}

	if _, ok := body["text"]; !ok {
		return isAutoFilAvail, errors.New("text is required")
	}
	// if auto fill, skip below

	if(isAutoFilAvail){
		return isAutoFilAvail, nil
	}
	
	if _, ok := body["iv"]; !ok {
		return isAutoFilAvail, errors.New("iv is required")
	}
	
	if _, ok := body["key"]; !ok {
		return isAutoFilAvail, errors.New("key is required")
	}
	// validate their length
	if len(body["key"].(string)) != constant.KeyByteSize {
		return isAutoFilAvail, errors.New("key length must be " + strconv.Itoa(constant.KeyByteSize )+ " bytes (or characters)")
	}
	if len(body["iv"].(string)) != constant.MessageBlockByteSize {
		return isAutoFilAvail, errors.New("iv length must be " + strconv.Itoa(constant.MessageBlockByteSize) + " bytes (or characters)")
	}

	return isAutoFilAvail, nil
}