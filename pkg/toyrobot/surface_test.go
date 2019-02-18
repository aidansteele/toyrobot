package toyrobot

import "fmt"

func ExampleSquareSurface_ValidPosition() {
	table := NewSquareSurface(5)

	fmt.Printf("%+v\n", table.ValidPosition(position{0, 0, 0}))
	fmt.Printf("%+v\n", table.ValidPosition(position{4, 4, 0}))
	fmt.Printf("%+v\n", table.ValidPosition(position{-1, 4, 0}))
	fmt.Printf("%+v\n", table.ValidPosition(position{6, 4, 0}))

	// Output:
	// true
	// true
	// false
	// false
}
