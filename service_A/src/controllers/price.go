package controllers

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/labstack/echo/v4"
    "bitcoin_server/models"
)


func GetLastPrice(c echo.Context) (err error) {
    // Read the data from storage
    file, _ := ioutil.ReadFile("bitcoin.json")
 
	var b_data models.BitCoinStorage
 
	_ = json.Unmarshal([]byte(file), &b_data)

    return c.JSON(http.StatusOK, b_data.LastValue)
}
