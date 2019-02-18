package toyrobot

import "strings"

type position struct {
	X           int
	Y           int
	Orientation orientation
}

// Go's support for enums is pretty lacklustre
const (
	orientationUnset = iota
	orientationNorth = iota
	orientationSouth = iota
	orientationEast  = iota
	orientationWest  = iota
)

type orientation int

func (o orientation) String() string {
	switch o {
	case orientationNorth:
		return "NORTH"
	case orientationSouth:
		return "SOUTH"
	case orientationEast:
		return "EAST"
	case orientationWest:
		return "WEST"
	default:
		panic("unknown orientation")
	}
}

func orientationFromString(input string) orientation {
	switch strings.ToUpper(input) {
	case "NORTH":
		return orientationNorth
	case "SOUTH":
		return orientationSouth
	case "EAST":
		return orientationEast
	case "WEST":
		return orientationWest
	default:
		panic("unknown orientation")
	}
}
