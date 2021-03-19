package finnhub

import (
    "regexp"
    "testing"
    //"fmt"
    //"io/ioutil"
)

func TestGetUrl(t *testing.T) {
    url := getUrl("aapl")
    want := regexp.MustCompile(`.*finnhub.*candle\?symbol.*resolution.*from.*to.*token`);
    if !want.MatchString(url) {
        t.Fatalf(`getUrl(aapl) = %q, want match for %#q`, url, want)
    }
}

func TestGetCandles (t *testing.T) {
    finnhub := NewFinnhub(&FakeHttpClient{})
    candles := finnhub.GetCandles("aapl")

    expectedCount := 253
    candlesCount := len(candles.Open)

    if len(candles.Open) != expectedCount {
        t.Fatalf(`Expecting %d candles, got %d instead`, expectedCount, candlesCount)
    }
}
