package main

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/go-vp/vp"
	"os"
	"strconv"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	client := binance.NewClient(apiKey, secretKey)
	klines, err := client.NewKlinesService().Symbol("ETHUSDT").
		Interval("15m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	var highs, lows, closes, volumes, takers, makers []float64
	for _, k := range klines {
		highs = append(highs, str2float64(k.High))
		lows = append(lows, str2float64(k.Low))
		closes = append(closes, str2float64(k.Close))
		v := str2float64(k.Volume)
		t := str2float64(k.TakerBuyQuoteAssetVolume)
		volumes = append(volumes, v)
		takers = append(takers, t)
		makers = append(makers, v-t)
		fmt.Printf("%v\n", k)
	}
	p := vp.VP{RowSize: 12, ValueAreaVolume: 0.7}
	profiles, poc, highest, lowest, vah, val, err := p.VolumeProfile(highs, closes, lows, volumes, takers, makers)
	if err != nil {
		panic(err)
	}
	for _, bar := range profiles {
		fmt.Println(bar)
	}
	fmt.Printf("poc = %d\n", poc)
	fmt.Printf("highest = %f, lowest = %f, vah = %f, val = %f\n", highest, lowest, vah, val)
}

func str2float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}
