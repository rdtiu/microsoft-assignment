package main

import (
	"bitcoin_server/controllers"
	"bitcoin_server/models"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var INTERVAL time.Duration = 10 // interval new data should be fetched in seconds
var TARGET_URL = "https://api.coincap.io/v2/assets/bitcoin"
var TEN_MINUTES_DATAPOINTS = 60

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	go loopInterval(INTERVAL)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from BitCoin Server!")
	})

	// Endpoints for liveness and readiness probes
	e.GET("/readyz", controllers.IsReady)
	e.GET("/livez", controllers.IsLive)

	e.GET("/price", controllers.GetLastPrice)
	e.GET("/average", controllers.CalculateAveragePrice)

	e.Logger.Fatal(e.Start(":8080"))
}

func GetBtcPrice() {
	client := http.Client{}
	req, err := http.NewRequest("GET", TARGET_URL, nil)
	if err != nil {
		panic(err.Error())
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data models.BitCoinPriceDetailsData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	// Read the data from storage
	file, _ := ioutil.ReadFile("bitcoin.json")

	var b_data models.BitCoinStorage

	_ = json.Unmarshal([]byte(file), &b_data)

	// Update data and write to storage
	float_last_value, _ := strconv.ParseFloat(data.Data.PriceUsd, 64)

	b_data.LastValue = float_last_value
	if len(b_data.DataPoints) < TEN_MINUTES_DATAPOINTS {
		b_data.DataPoints = append(b_data.DataPoints, float_last_value)
	} else {
		b_data.DataPoints = append(b_data.DataPoints[1:], float_last_value)
	}
	w_file, _ := json.MarshalIndent(b_data, "", " ")

	_ = ioutil.WriteFile("bitcoin.json", w_file, 0644)
}

func loopInterval(interval time.Duration) {
	GetBtcPrice()
	for range time.Tick(time.Second * interval) {
		GetBtcPrice()
	}
}
