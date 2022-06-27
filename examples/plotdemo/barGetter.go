package main

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-vp/vp"
)

func VPBarGetter(symbol string, timeframe string, profiles []vp.Profile, normal []float64) (chart *charts.Bar) {

	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Volume Profile [ mu=PoC, sigma=(VAH-VAL)/2 ]",
		Subtitle: symbol + ", " + timeframe,
	}))

	buyVolume := make([]opts.BarData, 0)
	totalVolume := make([]opts.BarData, 0)
	normalFit := make([]opts.LineData, 0)
	var xaxis []string
	for i, p := range profiles {
		totalVolume = append(totalVolume, opts.BarData{Value: p.BuyVolume - p.SellVolume})
		buyVolume = append(buyVolume, opts.BarData{Value: p.BuyVolume})
		xaxis = append(xaxis, fmt.Sprintf("%.2f", (p.High+p.Low)/2.0))
		normalFit = append(normalFit, opts.LineData{Value: normal[i]})
	}

	line := charts.NewLine()
	line.SetXAxis(xaxis).AddSeries("Normal Fit", normalFit).
		SetSeriesOptions(charts.WithLineStyleOpts(opts.LineStyle{Color: "lightblue"}))

	// Put data into instance
	bar.SetXAxis(xaxis).
		AddSeries("Total Volume", totalVolume).
		XYReversal().
		Overlap(line)

	return bar
}

func LiquidityBarGetter(symbol string, timeframe string, profiles []vp.Profile, normal []float64) (chart *charts.Bar) {

	liqBar := charts.NewBar()

	liqBar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Liquidity",
		Subtitle: symbol + ", " + timeframe,
	}))

	highLiq := make([]opts.BarData, 0)
	lowLiq := make([]opts.BarData, 0)
	var xaxis []string
	for i, p := range profiles {
		net := normal[i] - (p.BuyVolume - p.SellVolume)
		if net >= 0 {
			highLiq = append(highLiq, opts.BarData{Value: net})
			lowLiq = append(lowLiq, opts.BarData{Value: 0})
		} else {
			highLiq = append(highLiq, opts.BarData{Value: 0})
			lowLiq = append(lowLiq, opts.BarData{Value: net})
		}
		xaxis = append(xaxis, fmt.Sprintf("%.2f", (p.High+p.Low)/2.0))
	}

	// Put data into instance
	liqBar.SetXAxis(xaxis).
		AddSeries("Higher Liquidity", highLiq).
		AddSeries("Lower Liquidity", lowLiq).
		XYReversal()

	return liqBar
}
