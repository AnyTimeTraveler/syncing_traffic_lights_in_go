package main

import (
	"fmt"
	"time"
)

type TrafficLight struct {
	controlChannel chan bool
	syncChannel    chan bool
	cd             CardinalDirection
	c              Colour
}

func debug(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format, args)
	}
}

const CardinalDirections = 4
const Debug = false

func main() {
	c := make(chan bool)
	syncChannels := [CardinalDirections / 2]chan bool{make(chan bool), make(chan bool)}
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

	time.Sleep(10 * time.Millisecond)
}

func (t *TrafficLight) run(activeAxis Axis) {
	for {
		if t.cd.axis() == activeAxis {
			t.cycle()
			debug("\t\t\t\t\t\t\t\t\t[%s] Waiting to give up control...\n", t.cd)
			t.controlChannel <- true
			debug("\t\t\t\t\t\t\t\t\t[%s] Given up control!\n", t.cd)
			t.sync()
		} else {
			debug("\t\t\t\t\t\t\t\t\t[%s] Waiting for control...\n", t.cd)
			<-t.controlChannel
			debug("\t\t\t\t\t\t\t\t\t[%s] Got control!\n", t.cd)
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
	fmt.Printf("[%s] is %s\n", t.cd, t.c)
}

func (t *TrafficLight) sync() {
	select {
	case msg := <-t.syncChannel:
		if !msg {
			println("This should never happen")
		}
		debug("\t\t\t\t[%s] Synced!\n", t.cd)
	case t.syncChannel <- true:
		debug("\t\t\t\t[%s] Syncing...\n", t.cd)
	}
}
