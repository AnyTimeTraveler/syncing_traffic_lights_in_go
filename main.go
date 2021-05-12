package main

import (
	"fmt"
	"time"
)

type TrafficLight struct {
	controlChannel chan int
	syncChannel    chan int
	cd             CardinalDirection
	c              Colour
}

const CardinalDirections = 4

func main() {
	c := make(chan int)
	syncChannels := [CardinalDirections / 2]chan int{make(chan int), make(chan int)}
	lights := [CardinalDirections]TrafficLight{}
	for i := 0; i < CardinalDirections; i++ {
		lights[i] = TrafficLight{
			cd:             CardinalDirection(i),
			c:              Red,
			controlChannel: c,
			syncChannel:    syncChannels[i%(CardinalDirections/2)],
		}
	}
	for i := 0; i < CardinalDirections; i++ {
		go lights[i].run(NSAxis)
	}

	time.Sleep(1_000_000)
}

func (t *TrafficLight) run(activeAxis Axis) {
	for {
		if t.cd.axis() == activeAxis {
			t.cycle()
			fmt.Printf("\t\t\t\t\t\t\t\t\t[%s] Giving up control!\n", t.cd.toString())
			t.controlChannel <- 0
			t.sync()
		} else {
			fmt.Printf("\t\t\t\t\t\t\t\t\t[%s] Waiting for control...\n", t.cd.toString())
			<-t.controlChannel
			fmt.Printf("\t\t\t\t\t\t\t\t\t[%s] Got control!\n", t.cd.toString())
			t.sync()
		}
		activeAxis = activeAxis.next()
	}
}

func (t *TrafficLight) cycle() {
	for {
		t.show()
		t.sync()
		t.c.next()
		if t.c == Green {
			return
		}
	}
}

func (t *TrafficLight) show() {
	fmt.Printf("[%s] is %s\n", t.cd.toString(), t.c.toString())
}

func (t *TrafficLight) sync() {
	if t.cd/2 == 0 {
		fmt.Printf("\t\t\t\t[%s] Syncing A...\n", t.cd.toString())
		t.syncChannel <- 0
		fmt.Printf("\t\t\t\t[%s] Synced A!\n", t.cd.toString())
	} else {
		fmt.Printf("\t\t\t\t[%s] Syncing B...\n", t.cd.toString())
		<-t.syncChannel
		fmt.Printf("\t\t\t\t[%s] Synced B!\n", t.cd.toString())
	}
}
