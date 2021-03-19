package finnhub

import ()

type Candles struct {
    Open []float32 `json:"o"`
    High []float32 `json:"h"`
    Low []float32 `json:"l"`
    Closer []float32 `json:"c"`
    Volume []uint32 `json:"v"`
    Timestamp []uint32 `json:"t"`
    Status string `json:"s"`
}
