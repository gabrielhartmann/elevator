package elevator

import (
	"testing"
)

// Null test.  Make sure stepping the system with
// no requests results in no movement.
func TestSingleElevatorNull(t *testing.T) {
	startFloor := 5
	el := getNextElevator(startFloor, Idle)
	els := []*Elevator{el}
	ncs := &NearestControlSystem{
		elevators: els,
		scheduler: distScheduler,
	}

	var cs ControlSystem = ncs
	cs.Step()

	if len(el.Goals) != 0 {
		t.Errorf("Expected 0 goals, returned %d", len(el.Goals))
	}

	if el.Floor != startFloor {
		t.Errorf("Expected floor %d, returned %d", startFloor, el.Floor)
	}

	// Receive single request
	cs.Pickup(2, Up)
	cs.Step()

	if el.Direction != Down {
		t.Errorf("Expected direction %d, returned %d", Down, el.Direction)
	}

	if el.Floor != startFloor-1 {
		t.Errorf("Expected floor %d, returned %d", startFloor-1, el.Floor)
	}
}

func TestSingleElevatorSingleRequest(t *testing.T) {
	startFloor := 5
	el := getNextElevator(startFloor, Idle)
	els := []*Elevator{el}
	ncs := &NearestControlSystem{
		elevators: els,
		scheduler: distScheduler,
	}

	var cs ControlSystem = ncs

	cs.Pickup(2, Up)
	cs.Step()

	if el.Direction != Down {
		t.Errorf("Expected direction %d, returned %d", Down, el.Direction)
	}

	if el.Floor != startFloor-1 {
		t.Errorf("Expected floor %d, returned %d", startFloor-1, el.Floor)
	}
}

func TestTwoElevatorTwoOppositeRequests(t *testing.T) {
	startFloor := 5
	el1 := getNextElevator(startFloor, Idle)
	el2 := getNextElevator(startFloor, Idle)

	els := []*Elevator{el1, el2}
	ncs := &NearestControlSystem{
		elevators: els,
		scheduler: distScheduler,
	}

	var cs ControlSystem = ncs

	cs.Pickup(3, Down)
	cs.Pickup(7, Up)
	cs.Step()

	if el1.Direction == Idle || el2.Direction == Idle {
		t.Errorf("Both elevators should be busy, returned: el1: %v, el2: %v", el1, el2)
	}

	if el1.Direction == el2.Direction {
		t.Errorf("Expected elevators to head in opposite directions, returned: el1: %v, el2: %v", el1, el2)
	}
}

func TestEndToEnd(t *testing.T) {
	el := getNextElevator(5, Idle)
	els := []*Elevator{el}

	ncs := &NearestControlSystem{
		elevators: els,
		scheduler: distScheduler,
	}

	var cs ControlSystem = ncs
	cs.Pickup(4, Down)
	cs.Pickup(3, Down)
	cs.Pickup(2, Down)
	cs.Pickup(1, Up)
	cs.Pickup(9, Down)
	cs.Pickup(8, Down)
	cs.Pickup(7, Down)

	if len(el.Goals) != 7 {
		t.Errorf("Expected 7 goals, returned %v", len(el.Goals))
	}

	for i := 0; i < 21; i++ {
		cs.Step()
	}

	if len(el.Goals) != 0 {
		t.Errorf("Expected goals to be completed, returned %v", el)
	}

	if el.Direction != Idle {
		t.Errorf("Expected direction to be Idle, returned %v", el.Direction)
	}
}
