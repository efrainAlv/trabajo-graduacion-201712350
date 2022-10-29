package analyzer

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "errors"
    // "log"
    ts "kafka/types"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func OpenFile(price ts.HistoryPrice) (string, error){

    fmt.Println(price.TimePeriodStart+","+price.TimePeriodEnd+","+price.TimeOpen+","+price.TimeClose+","+price.RateOpen+","+price.RateHigh+","+price.RateLow+","+price.RateClose)

    resp, err := http.Get("https://storage.googleapis.com/crypto-forecast-models/arima/btcusd-history-price_daily.csv")
    jsonBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return "", errors.New("Failed to read file")
    }

    return string(jsonBody), nil
}