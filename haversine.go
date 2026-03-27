package geo_range

import "math"

const EarthRadius = 6371008.7714

// Haversine
//
// # Calculate the distance between two points
func Haversine(a, b *Location) float64 {
	dLat := (b.Lat - a.Lat) * math.Pi / 180.0
	dLng := (b.Lng - a.Lng) * math.Pi / 180.0

	lat1Rad := a.Lat * math.Pi / 180.0
	lat2Rad := b.Lat * math.Pi / 180.0

	out := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	cal := 2 * math.Atan2(math.Sqrt(out), math.Sqrt(1-out))
	return EarthRadius * cal
}

// PointToSegmentDistance
//
// # The distance from point to line
//
// Param:
//
//	point: the location of point
//	start: the start location of line
//	end: the end location of line
func PointToSegmentDistance(point, start, end *Location) float64 {
	ax, ay := start.Lng, start.Lat
	bx, by := end.Lng, end.Lat
	px, py := point.Lng, point.Lat

	dx := bx - ax
	dy := by - ay

	if dx == 0 && dy == 0 {
		return Haversine(point, start)
	}

	t := ((px-ax)*dx + (py-ay)*dy) / (dx*dx + dy*dy)

	if t < 0 {
		return Haversine(point, start)
	} else if t > 1 {
		return Haversine(point, end)
	}

	proj := &Location{
		Lat: ay + t*dy,
		Lng: ax + t*dx,
	}

	return Haversine(point, proj)
}

// IsNearPolyline
//
// # Is point near the line
//
// Param:
//
//	point: the location of point
//	polyline: the list of point of line
//	threshold: distance(meter)
func IsNearPolyline(point *Location, polyline []*Location, threshold float64) bool {
	for i := 0; i < len(polyline)-1; i += 1 {
		d := PointToSegmentDistance(point, polyline[i], polyline[i+1])
		if d <= threshold {
			return true
		}
	}
	return false
}

// SplitLine
//
// # Split a long line to multiple short lines
//
// Param:
//
//	start: the start point location of line
//	end : the end point location of line
//	distance: the max distance of each short line(meter)
func SplitLine(start, end *Location, distance float64) []*Location {
	wholeDistance := Haversine(start, end)
	n := int(math.Ceil(wholeDistance / distance)) // 段数

	points := make([]*Location, 0)
	for i := 0; i < n; i++ {
		t := float64(i) / float64(n)

		lng := start.Lng + (end.Lng-start.Lng)*t
		lat := start.Lat + (end.Lat-start.Lat)*t

		points = append(points, &Location{Lng: lng, Lat: lat})
	}
	points = append(points, end)

	return points
}

func GetPointAtDistance(point *Location, radius float64, angle float64) *Location {
	// 转换为弧度
	lat1 := point.Lat * math.Pi / 180
	lng1 := point.Lng * math.Pi / 180
	brng := angle * math.Pi / 180
	dist := radius / EarthRadius

	lat2 := math.Asin(math.Sin(lat1)*math.Cos(dist) +
		math.Cos(lat1)*math.Sin(dist)*math.Cos(brng))

	lng2 := lng1 + math.Atan2(math.Sin(brng)*math.Sin(dist)*math.Cos(lat1),
		math.Cos(dist)-math.Sin(lat1)*math.Sin(lat2))

	return &Location{
		Lat: lat2 * 180 / math.Pi,
		Lng: lng2 * 180 / math.Pi,
	}
}
