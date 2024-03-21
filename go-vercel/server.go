package main

import (
	"encoding/json"
	"net/http"

	"ICCBES/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/:blockCipherMode/:decryptionMode", func(c echo.Context) error {
		blockCipherMode := c.Param("blockCipherMode")
		decryptionMode := c.Param("decryptionMode")

		// validate mode
		validModes := []string{"ecb", "cbc"}
		isValidMode := false
		for _, validMode := range validModes {
			if blockCipherMode == validMode {
				isValidMode = true
				break
			}
		}
		if(decryptionMode != "encrypt" && decryptionMode != "decrypt"){
			return c.JSON(http.StatusBadRequest, "invalid mode. must be encrypt or decrypt")
		}

		if !isValidMode {
			return c.JSON(http.StatusBadRequest, "invalid mode. valid modes: ecb, cbc, cfb, ofb, ctr")
		}
		// get json body
		jsonBody := GetJSONRawBody(c)
		if jsonBody == nil {
			return c.JSON(http.StatusBadRequest, "empty json body")
		}
		var result map[string]interface{}
		// give to handler
		if(blockCipherMode == "cbc") {
			result = handler.CBCHandler(decryptionMode, jsonBody)
		}
		if(blockCipherMode == "ecb") {
			result = handler.ECBHandler(decryptionMode, jsonBody)
		}
		

		return c.JSON(http.StatusOK, result)

	})

	e.Logger.Fatal(e.Start(":1323"))
}


func GetJSONRawBody(c echo.Context) map[string]interface{}  {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

   return jsonBody
}