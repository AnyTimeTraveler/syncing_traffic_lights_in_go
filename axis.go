package main

import "fmt"

type Axis int

const (
	NSAxis Axis = iota
	EWAxis      = 1
)

func (a Axis) next() Axis {
	switch a {
	case NSAxis:
		return EWAxis
	case EWAxis:
		return NSAxis
	}
	fmt.Printf("Unexpected: %s", a)
	panic("UNREACHABLE")
	return NSAxis
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
	return "NSAxis"
}
