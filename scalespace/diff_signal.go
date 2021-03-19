package scalespace

import (
    . "github.com/phillaf/chronicler/pyramid"
)

func buildDiffGauss (gaussians []Curve) []Curve {
    var diffGauss []Curve

    for i:=1; i<len(gaussians); i++ {
        newDiffGauss := subtract(gaussians[i], gaussians[i-1])
        diffGauss = append(diffGauss, newDiffGauss)
    }

    return diffGauss
}

func subtract(blurry Curve, sharp Curve) Curve {
    var result Curve
    for i:=0; i<len(blurry); i++ {
        result = append(result, 128 + blurry[i] - sharp[i])
    }
    return result
}
