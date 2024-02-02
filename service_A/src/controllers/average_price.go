package controllers

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/labstack/echo/v4"
	"bitcoin_server/models"
)


func CalculateAveragePrice(c echo.Context) (err error) {
    // Read the data from storage
    file, _ := ioutil.ReadFile("bitcoin.json")
 
	var b_data models.BitCoinStorage
 
	_ = json.Unmarshal([]byte(file), &b_data)
 
	var sum float64 = 0
	var length float64 = 0
	for i := 0; i < len(b_data.DataPoints); i++ {
		sum += b_data.DataPoints[i]
		length += 1
	}

    // Update data and write to storage
    return c.JSON(http.StatusOK, sum / length)
}
