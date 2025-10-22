package geo_range

type Precision uint

const (
	Country     Precision = 2 // ±630km
	Province    Precision = 3 // ±78km
	City        Precision = 4 // ±20km
	District    Precision = 5 // ±2km
	Street      Precision = 6 // ±500m
	Village     Precision = 7 // less than 100m
	Building    Precision = 8 // ±20m
	HouseNumber Precision = 9 // less than 10m
)

type PrecisionDynamicFunc func(radius float64) (start, max Precision)

// GetPrecisionDynamic
//
// get the start precision and maximum precision,
// this is only a preset configuration, you could replace it with your function
func GetPrecisionDynamic(radius float64) (start, max Precision) {
	switch {
	case radius <= 1000:
		return District, Street
	case radius <= 5000:
		return District, Street
	case radius <= 10000:
		return District, District
	case radius <= 20000:
		return City, District
	case radius <= 50000:
		return City, District
	case radius <= 100000:
		return Province, City
	default:
		return Country, City
	}
}
