package scalespace

func convolve(signal []uint8, filter []float64) []uint8 {
    var result []float64
    mod := len(filter) % 2

    for i:=0; i<len(signal)-len(filter)+mod; i++ {
        sum := 0.0;
        for j:=0; j<len(filter); j++ {
            sum += float64(signal[i+j]) * filter[j]
        }
        result = append(result, sum)
    }

    return round(result)
}

func round(signal []float64) []uint8 {
    var result []uint8
    for i:=0; i<len(signal); i++ {
        result = append(result, uint8(signal[i]))
    }
    return result;
}
