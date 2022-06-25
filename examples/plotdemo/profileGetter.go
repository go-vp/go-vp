package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/go-vp/vp"
)

type KLineData struct {
	highs   []float64
	lows    []float64
	closes  []float64
	opens   []float64
	volumes []float64
	takers  []float64
	makers  []float64
	times   []string
}

func str2float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func GetKLineData(symbol string, interval string) (ret KLineData) {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	client := binance.NewClient(apiKey, secretKey)
	klines, err := client.NewKlinesService().Symbol(symbol).
		Interval(interval).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	var highs, lows, closes, opens, volumes, takers, makers []float64
	var times []string
	for _, k := range klines {
		highs = append(highs, str2float64(k.High))
		lows = append(lows, str2float64(k.Low))
		closes = append(closes, str2float64(k.Close))
		opens = append(opens, str2float64(k.Open))
		v := str2float64(k.Volume)
		t := str2float64(k.TakerBuyQuoteAssetVolume)
		volumes = append(volumes, v)
		takers = append(takers, t)
		makers = append(makers, v-t)
		opentime := time.Unix(k.OpenTime/1000.0, 0)
		times = append(times, fmt.Sprintf("%d/%02d/%02d %02d:%02d",
			opentime.Year(),
			opentime.Month(),
			opentime.Day(),
			opentime.Hour(),
			opentime.Minute(),
		))
	}
	var kline KLineData
	kline.highs = highs
	kline.lows = lows
	kline.closes = closes
	kline.opens = opens
	kline.makers = makers
	kline.takers = takers
	kline.times = times
	kline.volumes = volumes
	return kline
}

func ProfileGetter(data KLineData, bins int) (profiles []vp.Profile, poc float64, stddev float64) {

	p := vp.VP{RowSize: bins, ValueAreaVolume: 0.7}
	profiles, pocIndex, _, _, vah, val, err := p.VolumeProfile(data.highs, data.closes, data.lows, data.volumes, data.takers, data.makers)
	if err != nil {
		panic(err)
	}
	var price float64 = (profiles[pocIndex].High + profiles[pocIndex].Low) / 2.0
	var sdv float64 = (vah - val) / 2.0
	return profiles, price, sdv
}
