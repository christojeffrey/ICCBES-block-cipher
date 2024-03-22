package main

import (
	"encoding/json"
	"net/http"

	"ICCBES/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/:blockCipherMode/:decryptionMode", func(c echo.Context) error {
		blockCipherMode := c.Param("blockCipherMode")
		decryptionMode := c.Param("decryptionMode")

		// validate mode
		validModes := []string{"ecb", "cbc", "cfb", "ofb", "ctr"}
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
		// time the algorithm
		if(blockCipherMode == "cbc" || blockCipherMode == "cfb" || blockCipherMode == "ofb") {
			result = handler.CbcOfbCfbHandler(blockCipherMode, decryptionMode, jsonBody)
		}else if(blockCipherMode == "ecb") {
			result = handler.ECBHandler(decryptionMode, jsonBody)
		}else { // counter
			result = handler.CounterHandler(decryptionMode, jsonBody)
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