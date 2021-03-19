package scalespace

import (
    . "github.com/phillaf/chronicler/pyramid"
)

func BuildPyramid(timestamp []uint32, signal Curve, s int, octaveCount int, sigma float64) Pyramid {
    pyramid := Pyramid {
        Octaves: []Octave{buildOctave(timestamp, signal, s, sigma)},
    }

    for i:=1; i<octaveCount; i++ {
        octave := Octave{
            Timestamp: resampleUint32(pyramid.Octaves[i-1].Timestamp),
            Gauss: resampleList(pyramid.Octaves[i-1].Gauss),
        }
        octave.DiffGauss = buildDiffGauss(octave.Gauss)
        pyramid.Octaves = append(pyramid.Octaves, octave)
    }

    return pyramid
}

func buildOctave(timestamp []uint32, signal Curve, s int, sigma float64) Octave {
    var octave Octave

    octave.Timestamp = timestamp
    octave.Gauss = buildGaussians(signal, s, sigma)
    octave.DiffGauss = buildDiffGauss(octave.Gauss)

    return octave
}
