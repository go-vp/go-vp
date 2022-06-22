package vp

import "math"

func Max(high []float64) float64 {
	var h float64
	for _, hh := range high {
		if hh > h {
			h = hh
		}
	}
	return h
}

func Min(low []float64) float64 {
	if len(low) == 0 {
		return 0
	}
	l := low[0]
	for _, ll := range low {
		if ll < l {
			l = ll
		}
	}
	return l
}

type VP struct {
	RowSize         int
	ValueAreaVolume float64
}

type Profile struct {
	Index      int
	High       float64
	Low        float64
	BuyVolume  float64
	SellVolume float64
	Volume     float64
	ValueArea  bool
}

func (vp VP) VolumeProfile(highs, closes, lows, volumes, buys, sells []float64) (profiles []Profile, poc int, highest, lowest, vah, val float64, err error) {
	highest = Max(highs)
	lowest = Min(lows)
	var step float64
	profiles, step = makeProfile(highest, lowest, vp.RowSize)
	for i := 0; i < len(closes); i++ {
		index := int(math.Floor((closes[i] - lowest) / step))
		profiles[index].BuyVolume += buys[i]
		profiles[index].SellVolume += sells[i]
		profiles[index].Volume += volumes[i]
	}

	var hv, total float64
	for i, p := range profiles {
		total += p.Volume
		if hv < p.Volume {
			hv = p.Volume
			poc = i
		}
	}
	profiles[poc].ValueArea = true
	vaVolumeCap := total * vp.ValueAreaVolume

	upper := poc + 1
	lower := poc - 1
	va := hv
	for va < vaVolumeCap {
		uv := math.MaxFloat64
		if upper < len(profiles) {
			uv = profiles[upper].Volume
		}
		lv := math.MaxFloat64
		if lower >= 0 {
			lv = profiles[lower].Volume
		}
		if uv > lv {
			if va+lv > vaVolumeCap {
				break
			} else {
				va += lv
				profiles[lower].ValueArea = true
				lower--
			}
		} else {
			if va+uv > vaVolumeCap {
				break
			} else {
				va += uv
				profiles[upper].ValueArea = true
				upper++
			}
		}
	}
	vah = profiles[upper-1].High
	val = profiles[lower+1].Low
	return
}

func makeProfile(highest, lowest float64, size int) ([]Profile, float64) {
	step := (highest - lowest) / float64(size)
	ps := make([]Profile, size)
	for i := 0; i < len(ps); i++ {
		ps[i].Index = i
		ps[i].Low = lowest + step*float64(i)
		ps[i].High = ps[i].Low + step
	}
	return ps, step
}
