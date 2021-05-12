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
	fmt.Printf("Unexpected: %s", a.toString())
	panic("UNREACHABLE")
	return NSAxis
}

func (a Axis) toString() string {
	switch a {
	case NSAxis:
		return "NSAxis"
	case EWAxis:
		return "EWAxis"
	}
	fmt.Printf("Unexpected: %v", a)
	panic("UNREACHABLE")
	return "NSAxis"
}
