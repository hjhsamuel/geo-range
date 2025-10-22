package geo_range

import "math"

const EarthRadius = 6371008.7714

func Haversine(lat1, lng1, lat2, lng2 float64) float64 {
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLng := (lng2 - lng1) * math.Pi / 180.0

	lat1Rad := lat1 * math.Pi / 180.0
	lat2Rad := lat2 * math.Pi / 180.0

	out := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	cal := 2 * math.Atan2(math.Sqrt(out), math.Sqrt(1-out))
	return EarthRadius * cal
}
