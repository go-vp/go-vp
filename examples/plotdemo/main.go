package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func openHtml(name string) {
	var err error
	path, _ := os.Getwd()
	var url = filepath.Join(path, name)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var symbol string = "ETHUSDT"
	var timeframe string = "15m"
	var bins int = 13

	data := GetKLineData(symbol, timeframe)
	profiles, poc, _ := ProfileGetter(data, bins)
	kline := KLineGetter(symbol, timeframe, data, poc)
	bar := BarGetter(symbol, timeframe, profiles)

	f, _ := os.Create("results.html")
	kline.Render(f)
	bar.Render(f)

	openHtml(f.Name())
}
