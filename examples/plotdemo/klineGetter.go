package main

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-vp/vp"
)

func KLineGetter(symbol string, timeframe string, klines KLineData, poc float64, profiles []vp.Profile, normal []float64) (chart *charts.Kline) {

	kline := charts.NewKLine()

	var klineData [][4]float64
	var klineTime []string
	var beginTime string = klines.times[0]
	var endTime string = klines.times[len(klines.times)-1]
	for i, _ := range klines.volumes {
		klineData = append(klineData, [4]float64{klines.opens[i], klines.closes[i], klines.lows[i], klines.highs[i]})
		klineTime = append(klineTime, klines.times[i])
	}

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	nx := len(klineTime)
	for i := 0; i < nx; i++ {
		x = append(x, klineTime[i])
		y = append(y, opts.KlineData{Value: klineData[i]})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    symbol + " Price & PoC",
			Subtitle: beginTime + "~" + endTime + ", " + symbol + ", " + timeframe,
		}),
		charts.WithSingleAxisOpts(opts.SingleAxis{
			Min: profiles[0].Low,
			Max: profiles[len(profiles)-1].High,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      0,
			End:        float32(nx),
			XAxisIndex: []int{0},
		}),
	)
	kline.SetXAxis(x).AddSeries("kline", y).
		SetSeriesOptions(
			charts.WithMarkLineNameYAxisItemOpts(opts.MarkLineNameYAxisItem{
				Name:  "POC",
				YAxis: fmt.Sprintf("%f", poc),
			}),
		)

	liq := make([]float64, len(profiles))
	for i, p := range profiles {
		liq[i] = normal[i] - (p.BuyVolume - p.SellVolume)
	}

	for i, p := range profiles {
		if i > 0 {
			if liq[i]*liq[i-1] < 0 {
				kline.SetSeriesOptions(
					charts.WithMarkLineNameYAxisItemOpts(opts.MarkLineNameYAxisItem{
						Name:  "Liq",
						YAxis: fmt.Sprintf("%.2f", p.Low),
					}),
				)
			}
		}
	}
	return kline
}
