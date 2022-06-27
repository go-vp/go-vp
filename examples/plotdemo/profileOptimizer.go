package main

import (
	"math"

	"github.com/chobie/go-gaussian"
	"github.com/go-vp/vp"
	"gonum.org/v1/gonum/stat"
)

func FitNormal(mu float64, sigma float64, profiles []vp.Profile) (normal []float64) {
	yh := []float64{}
	y := []float64{}
	dist := gaussian.NewGaussian(mu, math.Pow(sigma, 2))
	for _, p := range profiles {
		yh = append(yh, dist.Pdf((p.High+p.Low)/2.0))
		y = append(y, p.BuyVolume-p.SellVolume)
	}
	_, beta := stat.LinearRegression(yh, y, nil, true)
	for i, _ := range yh {
		yh[i] = beta * yh[i]
	}

	return yh
}
