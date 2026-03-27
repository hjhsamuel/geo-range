package geo_range

import (
	"math"

	"github.com/mmcloughlin/geohash"
)

var geohashBase32 = []rune("0123456789bcdefghjkmnpqrstuvwxyz")

func MergeHashes(hashes []string, minMergedCnt int) []string {
	m := make(map[string][]string)
	for _, hash := range hashes {
		lowHash := hash[:len(hash)-1]
		if _, ok := m[lowHash]; !ok {
			m[lowHash] = make([]string, 0)
		}
		m[lowHash] = append(m[lowHash], hash)
	}

	res := make([]string, 0)
	for hash, vs := range m {
		if len(vs) >= minMergedCnt {
			res = append(res, hash)
		} else {
			res = append(res, vs...)
		}
	}
	return res
}

// RadiusSearch Get geohashes by coordinate and radius
//
// Params:
//
//	lat: latitude
//	lng: longitude
//	radius: radius(meter)
//	tolerance: the tolerance to check is rectangle is intersected with circle, (0, 1.0]
//	f: the custom function to get start and maximum precision by radius
func RadiusSearch(lat, lng, radius, tolerance float64, f PrecisionDynamicFunc) []string {
	if tolerance <= 0 || tolerance > 1 {
		return nil
	}

	var (
		startPrecision Precision
		maxPrecision   Precision
	)
	if f == nil {
		startPrecision, maxPrecision = GetPrecisionDynamic(radius)
	} else {
		startPrecision, maxPrecision = f(radius)
	}

	trie := NewGeoTrie()
	root := geohash.EncodeWithPrecision(lat, lng, uint(startPrecision))
	refineAdaptive(root, lat, lng, radius, startPrecision, maxPrecision, trie, tolerance)

	return trie.AllLeaf()
}

func rectIntersectsCircle(minLat, minLng, maxLat, maxLng, lat, lng, radius, tolerance float64) bool {
	closestLat := math.Max(minLat, math.Min(lat, maxLat))
	closestLng := math.Max(minLng, math.Min(lng, maxLng))
	d := Haversine(&Location{Lat: lat, Lng: lng}, &Location{Lat: closestLat, Lng: closestLng})
	return d <= (radius * tolerance)
}

func rectInsideCircle(minLat, minLng, maxLat, maxLng, lat, lng, radius float64) bool {
	corners := [][2]float64{
		{minLat, minLng},
		{minLat, maxLng},
		{maxLat, minLng},
		{maxLat, maxLng},
	}
	for _, c := range corners {
		if Haversine(&Location{Lat: lat, Lng: lng}, &Location{Lat: c[0], Lng: c[1]}) > radius {
			return false
		}
	}
	return true
}

func expandHash(hash string) []string {
	sub := make([]string, len(geohashBase32))
	for i, c := range geohashBase32 {
		sub[i] = hash + string(c)
	}
	return sub
}

func refineAdaptive(hash string, lat, lng, radius float64, startPrecision, maxPrecision Precision, trie *GeoTrieRoot,
	tolerance float64) {

	if trie.Has(hash) {
		return
	}

	box := geohash.BoundingBox(hash)

	if rectInsideCircle(box.MinLat, box.MinLng, box.MaxLat, box.MaxLng, lat, lng, radius) {
		if !trie.Insert(hash) {
			return
		}
		nbs := geohash.Neighbors(hash)
		for _, nb := range nbs {
			refineAdaptive(nb, lat, lng, radius, startPrecision, maxPrecision, trie, tolerance)
		}
		return
	}

	if !rectIntersectsCircle(box.MinLat, box.MinLng, box.MaxLat, box.MaxLng, lat, lng, radius, tolerance) {
		return
	}

	length := uint(len(hash))
	if length < uint(maxPrecision) {
		for _, sub := range expandHash(hash) {
			refineAdaptive(sub, lat, lng, radius, startPrecision, maxPrecision, trie, tolerance)
		}

		nbs := geohash.Neighbors(hash)
		for _, nb := range nbs {
			refineAdaptive(nb, lat, lng, radius, startPrecision, maxPrecision, trie, tolerance)
		}
	} else {
		trie.Insert(hash)
	}
}
