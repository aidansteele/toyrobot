package toyrobot

import (
	"os"
	"strings"
)

func Example() {
	input := strings.NewReader(`
LEFT # all this is ignored as it's before the first valid PLACE
MOVE
MOVE
MOVE
PLACE 6, 1, EAST
# ^ that was also ignored
PLACE 0, 1, EAST
MOVE
report
# ^ case insensitive
MOVE
REPORT
MOVE
REPORT
LEFT
movE
# ^ typo insensitive
REPORT
MOVE
REPORT
MOVE
REPORT
MOVE
REPORT
`)

	RunStandardSimulation(input, os.Stdout)

	// Output:
	// REPORT 1, 1, EAST
	// REPORT 2, 1, EAST
	// REPORT 3, 1, EAST
	// REPORT 3, 2, NORTH
	// REPORT 3, 3, NORTH
	// REPORT 3, 4, NORTH
	// REPORT 3, 4, NORTH
}
