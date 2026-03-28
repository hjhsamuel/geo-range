package geo_range

import "math"

type Location struct {
	Lng float64
	Lat float64
}

func round(val float64, length float64) float64 {
	shift := math.Pow(10, length)
	return math.Round(val*shift) / shift
}
