package toyrobot

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type echoReceiver struct {
	io.Writer
}

func (e *echoReceiver) Place(p position) {
	fmt.Printf("PLACE %d, %d, %s\n", p.X, p.Y, p.Orientation)
}

func (e *echoReceiver) Move() {
	fmt.Println("MOVE")
}

func (e *echoReceiver) Left() {
	fmt.Println("LEFT")
}

func (e *echoReceiver) Right() {
	fmt.Println("RIGHT")
}

func (e *echoReceiver) Report() {
	fmt.Println("REPORT")
}

func (e *echoReceiver) Invalid(line string) {
	fmt.Printf("INVALID: %s\n", line)
}

func ExampleParser_Next() {
	input := `
REPORT
report
repoRt
left  
 right 
	move
place
place 1, 0, east
place 1, 0, noteast
`

	p := NewParser(strings.NewReader(input), &echoReceiver{os.Stdout})

	for p.Next() {
		// here we're just immediately going to the next line
	}

	// Output:
	// REPORT
	// REPORT
	// REPORT
	// LEFT
	// RIGHT
	// MOVE
	// INVALID: place
	// PLACE 1, 0, EAST
	// INVALID: place 1, 0, noteast
}
