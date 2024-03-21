package main

import (
	"encoding/json"
	"net/http"

	"ICCBES/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func temp() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/:mode/encrypt", func(c echo.Context) error {
		mode := c.Param("mode")

		// validate mode
		validModes := []string{"ecb", "cbc", "cfb", "ofb", "ctr"}
		isValidMode := false
		for _, validMode := range validModes {
			if mode == validMode {
				isValidMode = true
				break
			}
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
		if(mode == "cbc") {
			result = handler.CBCHandler("encrypt", jsonBody)
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