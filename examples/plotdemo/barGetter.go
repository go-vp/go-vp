package main

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-vp/vp"
)

func BarGetter(symbol string, timeframe string, profiles []vp.Profile) (chart *charts.Bar) {

	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Volume Profile",
		Subtitle: symbol + ", " + timeframe,
	}))

	buyVolume := make([]opts.BarData, 0)
	totalVolume := make([]opts.BarData, 0)
	var xaxis []string
	for _, p := range profiles {
		totalVolume = append(totalVolume, opts.BarData{Value: p.BuyVolume - p.SellVolume})
		buyVolume = append(buyVolume, opts.BarData{Value: p.BuyVolume})
		xaxis = append(xaxis, fmt.Sprintf("%.2f", (p.High+p.Low)/2.0))
	}

	// Put data into instance
	bar.SetXAxis(xaxis).
		AddSeries("Buy Volume", buyVolume).
		AddSeries("Total Volume", totalVolume).
		XYReversal()

	return bar
}
