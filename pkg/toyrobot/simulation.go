package toyrobot

import (
	"fmt"
	"io"
)

// Simulation ties the whole thing together. In its current form, the
// simulation is little more than a Surface and an uninitialised position
// (representing the robot).
//
// Any non-PLACE commands issued before the first PLACE are discarded.
// Likewise, any unrecognised commands are discarded. It's anyone's guess
// whether this helps or hinders forwards-compatibility -- hopefully the
// spec evolves in a backwards-compatible way. :-)
//
// Reports are written in plaintext with a trailing newline to the reporter.
type Simulation struct {
	robot    position
	surface  Surface
	reporter io.Writer
}

func NewSimulation(surface Surface, reporter io.Writer) *Simulation {
	return &Simulation{
		surface:  surface,
		reporter: reporter,
	}
}

func (s *Simulation) Place(p position) {
	// spec says not to allow robot to be placed in an invalid position
	if s.surface.ValidPosition(p) {
		s.robot = p
	}
}

func (s *Simulation) Move() {
	if !s.placed() {
		return
	}

	proposed := s.robot
	switch s.robot.Orientation {
	case orientationNorth:
		proposed.Y++
	case orientationSouth:
		proposed.Y--
	case orientationEast:
		proposed.X++
	case orientationWest:
		proposed.X--
	}

	if s.surface.ValidPosition(proposed) {
		s.robot = proposed
	}
}

func (s *Simulation) Left() {
	switch s.robot.Orientation {
	case orientationNorth:
		s.robot.Orientation = orientationWest
	case orientationSouth:
		s.robot.Orientation = orientationEast
	case orientationEast:
		s.robot.Orientation = orientationNorth
	case orientationWest:
		s.robot.Orientation = orientationSouth
	}
}

func (s *Simulation) Right() {
	switch s.robot.Orientation {
	case orientationNorth:
		s.robot.Orientation = orientationEast
	case orientationSouth:
		s.robot.Orientation = orientationWest
	case orientationEast:
		s.robot.Orientation = orientationSouth
	case orientationWest:
		s.robot.Orientation = orientationNorth
	}
}

func (s *Simulation) Report() {
	if !s.placed() {
		return
	}
	fmt.Fprintf(s.reporter, "REPORT %d, %d, %s\n", s.robot.X, s.robot.Y, s.robot.Orientation)
}

// Spec doesn't say what to do in case of invalid input, lets just noop.
func (s *Simulation) Invalid(line string) {
}

func (s *Simulation) placed() bool {
	return s.robot.Orientation != orientationUnset
}

// RunStandardSimulation provides a simple implementation conforming
// to the Toy Robot Spec. It reads from an input (typically os.Stdin),
// reports to an output (typically os.Stdout) and processes all input
// commands without interruption. A single robot is placed on a 5x5
// square table with (0,0) being the southwest-most point.
func RunStandardSimulation(input io.Reader, output io.Writer) {
	surface := NewSquareSurface(5)
	simulation := NewSimulation(surface, output)
	parser := NewParser(input, simulation)

	for parser.Next() {
		// just process all commands until EOF
	}
}
