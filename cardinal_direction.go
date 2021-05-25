package main

import "fmt"

type CardinalDirection int

const (
	North CardinalDirection = iota
	East                    = 1
	South                   = 2
	West                    = 3
)

func (cd CardinalDirection) next() CardinalDirection {
	switch cd {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	panic("UNREACHABLE")
	return North
}

func (cd CardinalDirection) axis() Axis {
	switch cd {
	case North:
		return NSAxis
	case South:
		return NSAxis
	case East:
		return EWAxis
	case West:
		return EWAxis
	}
	fmt.Printf("Unexpected: %s", cd)
	panic("UNREACHABLE")
	return NSAxis
}

func (cd CardinalDirection) String() string {
	switch cd {
	case North:
		return "\u001B[1;34mNorth\u001B[0m"
	case East:
		return "\u001B[1;35mEast \u001B[0m"
	case South:
		return "\u001B[1;36mSouth\u001B[0m"
	case West:
		return "\u001B[1;37mWest \u001B[0m"
	}
	fmt.Printf("Unexpected: %d", cd)
	panic("UNREACHABLE")
	return "North"
}
