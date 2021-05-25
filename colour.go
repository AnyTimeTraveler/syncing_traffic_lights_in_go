package main

import "fmt"

type Colour int

const (
	Red    Colour = iota
	Green         = 1
	Yellow        = 2
)

func (c *Colour) next() {
	switch *c {
	case Red:
		*c = Green
		return
	case Green:
		*c = Yellow
		return
	case Yellow:
		*c = Red
		return
	}
	fmt.Printf("Unexpected: %s", c)
	panic("UNREACHABLE")
}

func (c Colour) String() string {
	switch c {
	case Red:
		return "   \u001B[1;31mRed\u001B[0m"
	case Green:
		return " \u001B[1;32mGreen\u001B[0m"
	case Yellow:
		return "\u001B[1;33mYellow\u001B[0m"
	}
	fmt.Printf("Unexpected: %d", c)
	panic("UNREACHABLE")
	return "Red"
}
