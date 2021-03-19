package normalizer

import ()

func Normalize(input []float32) []uint8 {

    var output []uint8
    min, max := minMax(input)
    max = max - min

    for _, value := range input {
        output = append(output, uint8((value-min)*255/max))
    }

    return output
}

func minMax(input []float32) (float32, float32) {
    min := input[0]
    max := input[0]

    for _, value := range input {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }

    return min, max
}
