package scalespace

import (
    . "github.com/phillaf/chronicler/pyramid"
    "math"
)

func resampleList (curveList []Curve) []Curve {
    var result []Curve

    for i:=0; i<len(curveList); i++ {
        result = append(result, resampleUint8(curveList[i]))
    }

    return result
}

func resampleUint8 (original []uint8) []uint8 {
    var resampled Curve

    halfSize := int(math.Floor( float64(len(original))/2.0) )
    for j:=0; j<halfSize; j++ {
        resampled = append(resampled, original[j*2])
    }

    return resampled
}

func resampleUint32 (original []uint32) []uint32 {
    var resampled []uint32

    halfSize := int(math.Floor( float64(len(original))/2.0) )
    for j:=0; j<halfSize; j++ {
        resampled = append(resampled, original[j*2])
    }

    return resampled
}
