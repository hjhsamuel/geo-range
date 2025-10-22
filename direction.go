package geo_range

const (
	DirectionBaseAngle  float64 = 45
	DirectionSplitAngle float64 = 22.5
)

type Direction int

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

func GuessDirection(brg float64) Direction {
	idx := int((brg+DirectionSplitAngle)/DirectionBaseAngle) % 8
	return Direction(idx)
}
