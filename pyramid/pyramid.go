package pyramid

import (
    "fmt"
    "strings"
)

type Curve []uint8

type Octave struct {
    Timestamp []uint32 `json:"timestamp"`
    Gauss []Curve `json:"gauss"`
    DiffGauss []Curve `json:"diffGauss"`
    Extrema []uint32 `json:"extrema"`
}

type Pyramid struct {
    Octaves []Octave `json:"octaves"`
}

func (c Curve) MarshalJSON() ([]byte, error) {
    var result string
    if c == nil {
        result = "null"
    } else {
        result = strings.Join(strings.Fields(fmt.Sprintf("%d", c)), ",")
    }
    return []byte(result), nil
}
