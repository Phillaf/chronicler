package main

import (
    "github.com/phillaf/chronicler/finnhub"
    "github.com/phillaf/chronicler/normalizer"
    "github.com/phillaf/chronicler/scalespace"
    "github.com/phillaf/chronicler/extrema"

    "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func home(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")

    if (r.Method == "OPTIONS") {
        return
    }

    api := finnhub.NewFinnhub(&finnhub.RealHttpClient{})
    candles := api.GetCandles("AAPL")
    normalized := normalizer.Normalize(candles.Open)

    pyramid := scalespace.BuildPyramid(candles.Timestamp, normalized, 2, 4, 1.6);
    extrema.ScanPyramid(&pyramid)

    //mycurve := Curve{
    //    Dates: candles.Timestamp,
    //    Raw: candles.Open,
    //    Normalized: normalized,
    //}

    json.NewEncoder(w).Encode(pyramid)
}
