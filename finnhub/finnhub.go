package finnhub

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "strconv"
    "time"
)

const Url = "https://finnhub.io/api/v1/stock/candle"
const Token = "example"

type Finnhub struct{
    Http HttpClient
}

func init() {
    log.SetFlags(0)
}

func NewFinnhub(httpClient HttpClient) *Finnhub {
    return &Finnhub{
        Http: httpClient,
    }
}

func (f *Finnhub) GetCandles(symbol string) Candles{

    url := getUrl(symbol)
    resp, err := f.Http.Get(url)

    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var candles Candles
    json.Unmarshal([]byte(body), &candles)

    return candles
}

func getUrl(symbol string) string {
    symbol = "?symbol=" + symbol
    resolution := "&resolution=D"
    from := "&from=" + strconv.FormatInt(time.Now().AddDate(-1, 0, 0).Unix(), 10)
    to := "&to=" + strconv.FormatInt(time.Now().Unix(), 10)
    token := "&token=" + Token
    return Url + symbol + resolution + from + to + token
}
