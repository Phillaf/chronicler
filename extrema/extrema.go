package extrema

import (
    . "github.com/phillaf/chronicler/pyramid"
)

func ScanPyramid(pyramid *Pyramid) {
    for o:=0; o<len(pyramid.Octaves); o++ {
        for dog:=1; dog<len(pyramid.Octaves[o].DiffGauss)-1; dog++ {
            for i:=0; i<len(pyramid.Octaves[o].DiffGauss[dog]); i++ {
                scanMin(pyramid, o, dog, i)
                scanMax(pyramid, o, dog, i)
            }
        }
    }
}

func scanMin(pyramid *Pyramid, o int, dog int, i int) {
    value := pyramid.Octaves[o].DiffGauss[dog][i]

    if i > 0 {
        for j:=dog-1; j<=dog+1; j++ {
            if pyramid.Octaves[o].DiffGauss[j][i-1] <= value {
                return
            }
        }
    }

    if i < len(pyramid.Octaves[o].DiffGauss[dog])-1 {
        for j:=dog-1; j<=dog+1; j++ {
            if pyramid.Octaves[o].DiffGauss[j][i+1] <= value {
                return
            }
        }
    }

    if pyramid.Octaves[o].DiffGauss[dog-1][i] <= value {return}
    if pyramid.Octaves[o].DiffGauss[dog+1][i] <= value {return}

    timestamp := pyramid.Octaves[o].Timestamp[i]
    pyramid.Octaves[o].Extrema = append(pyramid.Octaves[o].Extrema, timestamp)
}

func scanMax(pyramid *Pyramid, o int, dog int, i int) {
    value := pyramid.Octaves[o].DiffGauss[dog][i]

    if i > 0 {
        for j:=dog-1; j<=dog+1; j++ {
            if pyramid.Octaves[o].DiffGauss[j][i-1] >= value {
                return
            }
        }
    }

    if i < len(pyramid.Octaves[o].DiffGauss[dog])-1 {
        for j:=dog-1; j<=dog+1; j++ {
            if pyramid.Octaves[o].DiffGauss[j][i+1] >= value {
                return
            }
        }
    }

    if pyramid.Octaves[o].DiffGauss[dog-1][i] >= value {return}
    if pyramid.Octaves[o].DiffGauss[dog+1][i] >= value {return}

    timestamp := pyramid.Octaves[o].Timestamp[i]
    pyramid.Octaves[o].Extrema = append(pyramid.Octaves[o].Extrema, timestamp)
}
