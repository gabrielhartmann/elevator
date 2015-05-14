package main

import (
	"fmt"
	. "github.com/gabrielhartmann/elevator/elevator"
)

const (
	MaxElevatorCount = 16
	MaxFloorCount    = 5
)

func main() {
	// Create an elevator on the top and bottom floors
	el1 := NewElevator(1, 0)
	el2 := NewElevator(2, 4)
	els := []*Elevator{el1, el2}

	// Create a scheduler
	sched := NewDistanceScheduler(MaxFloorCount, 0.7, 0.3)

	// Create a control system
	ncs := NewNearestControlSystem(els, &sched)
	var cs ControlSystem = &ncs

	// Print the status of the system
	fmt.Printf("Created an elevator on the top and bottom floors\n")
	printStatus(cs)

	// Press all the up buttons
	// Note: no UP button on the top floor
	for i := 0; i < MaxFloorCount-1; i++ {
		cs.Pickup(i, Up)
	}

	// Press all the down buttons
	// Note: no DOWN button on the bottom floor
	for i := 1; i < MaxFloorCount; i++ {
		cs.Pickup(i, Down)
	}

	// Print the status of the system
	fmt.Printf("Pressed all the buttons\n")
	printStatus(cs)

	// Step the simulation until all requests have been serviced
	fmt.Printf("Running the simulation until all requests are fulfilled\n")
	for requestsOutstanding(cs) {
		cs.Step()
	}
}

func printStatus(cs ControlSystem) {
	for _, el := range cs.Status() {
		fmt.Printf("Elevator Id: %v, Floor: %v, Direction %v, Requests: %v\n", el.Id, el.Floor, translateDirection(el.Direction), el.Goals)
	}

	fmt.Printf("\n")
}

func translateDirection(dir int) string {
	switch dir {
	case Idle:
		return "idle"
	case Up:
		return "up"
	case Down:
		return "down"
	default:
		return "invalid"
	}
}

func requestsOutstanding(cs ControlSystem) bool {
	for _, el := range cs.Status() {
		if len(el.Goals) > 0 {
			return true
		}
	}

	return false
}
