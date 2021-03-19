package scalespace

import (
    "encoding/json"
    //"fmt"
    "io/ioutil"
    "testing"
    //"github.com/phillaf/chronicler/finnhub"
    //"github.com/phillaf/chronicler/normalizer"
)

var signal Curve

func init() {
    data, _ := ioutil.ReadFile("curve_fixture.json")
    json.Unmarshal([]byte(data), &signal)
}

func TestConvolveKeepsSize(t *testing.T) {
    filter := gaussFilter(1)
    padded := pad(signal, 1)
    result := convolve(padded, filter)
    if len(result) != len(signal) {
        t.Fatalf("convolution should preserve size. expecting %d, got %d", len(signal), len(result))
    }
}

func TestGaussianAreSameSize(t *testing.T) {
    gaussians := buildGaussians(signal, 2, 1)
    size := len(gaussians[0])

    for i:=0; i<len(gaussians); i++ {
        if size != len(gaussians[i]) {
            t.Fatal("gaussians of an octave should be the same size")
        }
    }
}

func TestBuildPyramid(t *testing.T) {
    result := buildPyramid(signal, 2, 4, 1.6)
    result = result
}

//func TestGenerateFixture(t *testing.T) {
//
//    content, err := ioutil.ReadFile("../finnhub/aapl_fixture")
//    if err != nil {
//        t.Fatal(err)
//    }
//
//    var candles finnhub.Candles
//    json.Unmarshal([]byte(content), &candles)
//
//    normalized := normalizer.Normalize(candles.Open)
//
//    curve, err := json.Marshal(Curve(normalized))
//    if err != nil {
//        t.Fatal(err)
//    }
//
//    fmt.Println(string(curve))
//
//    ioutil.WriteFile("curve_fixture.json", curve, 0644)
//
//}
