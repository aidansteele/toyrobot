package toyrobot

// Surface is the interface that a type must conform to if it is to have
// robots placed upon it. A surface need only concern itself with whether
// a given location is valid for a hypothetical robot.
//
// Non-exhaustive examples of things that a surface needn't worry about:
//
// * If a given position is close enough to a robot's current location
//
// * If a robot is already in a given position
type Surface interface {
	ValidPosition(p position) bool
}

// SquareSurface is the finite square tabletop surface described in the
// Toy Robot Spec. It conforms to the Surface interface.
type SquareSurface struct {
	minX int // inclusive
	minY int // inclusive
	maxX int // exclusive
	maxY int // exclusive
}

func NewSquareSurface(size int) *SquareSurface {
	return &SquareSurface{
		minX: 0,
		minY: 0,
		maxX: size,
		maxY: size,
	}
}

func (s *SquareSurface) ValidPosition(p position) bool {
	return p.X >= s.minX && p.X < s.maxX && p.Y >= s.minY && p.Y < s.maxY
}
