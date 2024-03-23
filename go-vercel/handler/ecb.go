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
func ECBHandler(mode string, body map[string]interface{}) map[string]interface{} {
	// TODO: add limitation such as length and character
	// validate body: text string and iv string, key string
	isAutoFill, err := validateECBBody(body)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	
	var key []byte
	var iv []byte
	// handle autofill
	if(isAutoFill){
		key = utils.GenerateRandomByte(constant.KeyByteSize)
	}else{
		// get text and iv from body
	key = []byte(body["key"].(string))
	}
	text := utils.TransmissionDecoding(body["text"].(string))

	// encrypt or decrypt
	var result []byte
	if mode == "encrypt" {
		result = cipherMode.EncryptECB(text, key, BCA.EncryptionAlgorithm)
	} else {
		result = cipherMode.DecryptECB(text, key, BCA.DecryptionAlgorithm)
	}
	println("RESULTNYA", result)
	return map[string]interface{}{"result": utils.TransmissionEncoding(result), "key" : string(key), "iv": string(iv)}
}
// if autofil is available, only check text

func validateECBBody(body map[string]interface{})  (bool, error) {
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
	
	if _, ok := body["key"]; !ok {
		return isAutoFilAvail, errors.New("key is required")
	}
	// validate their length
	if len(body["key"].(string)) != constant.KeyByteSize {
		return isAutoFilAvail, errors.New("key length must be " + strconv.Itoa(constant.KeyByteSize )+ " bytes (or characters)")
	}

	return isAutoFilAvail, nil
}