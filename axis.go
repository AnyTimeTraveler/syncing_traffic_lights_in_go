package main

import "fmt"

type Axis int

const (
	NSAxis Axis = iota
	EWAxis      = 1
)

func (a *Axis) next() {
	switch *a {
	case NSAxis:
		*a = EWAxis
		return
	case EWAxis:
		*a = NSAxis
		return
	}
	fmt.Printf("Unexpected: %s", a)
	panic("UNREACHABLE")
}

func (a Axis) String() string {
	switch a {
	case NSAxis:
		return "NSAxis"
	case EWAxis:
		return "EWAxis"
	}
	fmt.Printf("Unexpected: %d", a)
	panic("UNREACHABLE")
}
