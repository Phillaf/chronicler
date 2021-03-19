package finnhub

import (
    "bytes"
    "io/ioutil"
    "log"
    "net/http"
)

type HttpClient interface {
    Get(url string) (*http.Response, error)
}

type RealHttpClient struct {}
type FakeHttpClient struct {}

func (r *RealHttpClient) Get(url string) (*http.Response, error) {
    return http.Get(url)
}

func (r *FakeHttpClient) Get(url string) (*http.Response, error) {
    content, err := ioutil.ReadFile("aapl_fixture")
    if err != nil {
        log.Fatal(err)
    }

    body := ioutil.NopCloser(bytes.NewReader([]byte(content)))
    return &http.Response{
        Body: body,
    }, nil
}
