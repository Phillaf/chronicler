package extrema

import (
    . "github.com/phillaf/chronicler/pyramid"
    "testing"
)

func createPyramid() Pyramid {
    return Pyramid{
        Octaves: []Octave{
            Octave{
                Timestamp: []uint32{1,2,3,4,5,6,7,8,9,10},
                DiffGauss: []Curve{
                    Curve{5,5,5,5,5,5,5,5,5,5},
                    Curve{9,6,6,9,6,7,6,3,6,1},
                    Curve{8,8,8,8,8,8,8,8,8,8},
                },
                Extrema: []uint32{},
            },
            Octave{
            },
        },
    }
}

func TestScanPyramid(t *testing.T) {
    pyramid := createPyramid()
    ScanPyramid(&pyramid)
    expected := []uint32{1,4,8,10}
    for i:=0; i<4; i++ {
        if pyramid.Octaves[0].Extrema[i] != expected[i] {
            t.Fatalf("scan pyramid result unexpected")
        }
    }
}
