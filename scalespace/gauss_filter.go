package scalespace

import (
    . "github.com/phillaf/chronicler/pyramid"
    "math"
)

func buildGaussians(signal Curve, s int, sigma float64) []Curve {
    gaussians := []Curve{signal}
    k := math.Exp2(1/float64(s))
    filter := gaussFilter(k*sigma)

    for i:=0; i<s+2; i++ {
        padded := pad(gaussians[len(gaussians)-1], k*sigma)
        blurred := convolve(padded, filter)
        gaussians = append(gaussians, blurred)
    }

    return gaussians
}

func gaussFilter(sigma float64) []float64{
    var filter []float64
    filterSize := 2*int(math.Ceil(3*sigma))+1 // 3 std devs on each side + one for the peak

    filterSum := 0.0
    for i:=0; i<filterSize; i++ {
        value := gauss(sigma, i-int(filterSize/2))
        filter = append(filter, value)
        filterSum += value
    }

    for i:=0; i<filterSize; i++ {
        filter[i] = filter[i]/filterSum // filter normalization
    }

    return filter
}

func gauss(sigma float64, x int) float64 {
    sigSq := math.Pow(sigma, 2)
    left := 1.0 / math.Sqrt(2 * math.Pi * sigSq)
    exp := - (math.Pow(float64(x), 2) / (2 * sigSq))
    right := math.Pow(math.E, exp)
    return left * right
}

func pad(signal []uint8, sigma float64) []uint8 {
    var padded []uint8
    borderSize := int(math.Ceil(3*sigma))

    for i:=0; i<borderSize; i++ {
        padded = append(padded, signal[0])
    }

    for i:=0; i<len(signal); i++ {
        padded = append(padded, signal[i])
    }

    for i:=len(signal); i<len(signal)+borderSize; i++ {
        padded = append(padded, signal[len(signal)-1])
    }

    return padded
}
