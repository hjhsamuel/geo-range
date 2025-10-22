package geo_range

import (
	"errors"
	"github.com/mmcloughlin/geohash"
)

func IsCoordinateValid(lat, lng float64) bool {
	if lat < -90 || lat > 90 {
		return false
	}
	if lng < -180 || lng > 180 {
		return false
	}
	return true
}

func GetHash(lat, lng float64) (string, error) {
	if !IsCoordinateValid(lat, lng) {
		return "", errors.New("invalid coordinate")
	}
	hash := geohash.Encode(lat, lng)
	return hash, nil
}
