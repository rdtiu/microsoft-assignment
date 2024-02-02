package models

type BitCoinStorage struct {
    LastValue float64
    DataPoints []float64
}

type BitCoinPriceDetailsData struct {
    Data BitCoinPriceDetails `json:"data"`
    Timestamp int64 `json:"timestamp"`
}

type BitCoinPriceDetails struct {
    Name string `json:"name"`
    Symbol string `json:"symbol"`
    PriceUsd string `json:"priceUsd"`
}
