package elevator

import (
	"testing"
)

func TestStepNaturalDirection(t *testing.T) {
	startFloor := 5
	expectedFloor := startFloor - 1

	el := getNextElevator(startFloor, Down)
	req := request{expectedFloor, Down}
	el.Goals = requests{req}

	ncs := &NearestControlSystem{
		elevators: []*Elevator{el},
	}

	var cs ControlSystem = ncs
	cs.Step()

	actualFloor := cs.Status()[0].Floor

	if actualFloor != expectedFloor {
		t.Errorf("Expected floor: %v, returned %v", expectedFloor, actualFloor)
	}
}

func TestStepOppositeDirection(t *testing.T) {
	startFloor := 5
	expectedFloor := startFloor + 1

	el := getNextElevator(startFloor, Down)
	req := request{expectedFloor, Down}
	el.Goals = requests{req}

	ncs := &NearestControlSystem{
		elevators: []*Elevator{el},
	}

	var cs ControlSystem = ncs
	// Turn around
	cs.Step()
	// Move
	cs.Step()

	actualFloor := cs.Status()[0].Floor

	if actualFloor != expectedFloor {
		t.Errorf("Expected floor: %v, returned %v", expectedFloor, actualFloor)
	}
}

func TestCompleteGoal(t *testing.T) {
	startFloor := 5

	el := getNextElevator(startFloor, Down)
	req := request{startFloor, Down}
	el.Goals = requests{req}

	ncs := &NearestControlSystem{
		elevators: []*Elevator{el},
	}

	var cs ControlSystem = ncs
	cs.Step()

	if len(el.Goals) != 0 {
		t.Errorf("Expected all goal floors to be completed, returned: %v", el.Goals)
	}

	if el.Direction != Idle {
		t.Errorf("Expected the direction of the elevator to be Idle, returned: %v", el.Direction)
	}
}

func TestIdleToMotionTransition(t *testing.T) {
	startFloor := 5
	closestFloor := startFloor - 2
	closeReq := request{closestFloor, Down}
	farthestFloor := startFloor + 3
	farReq := request{farthestFloor, Down}

	el := getNextElevator(startFloor, Idle)
	el.Goals = requests{closeReq, farReq}

	if closeGoal := el.closestGoal(); closeReq != closeGoal {
		t.Errorf("Expected closestFloor: %v, returned %v", closeReq, closeGoal)
	}

	ncs := &NearestControlSystem{
		elevators: []*Elevator{el},
	}

	var cs ControlSystem = ncs
	cs.Step()

	if el.Direction != Down {
		t.Errorf("Expected elevator's direction to be Down, returned: %v", el.Direction)
	}

	cs.Step()
	if el.Floor != startFloor-1 {
		t.Errorf("Expected elevator's floor to be %v, returned: %v", startFloor-1, el.Floor)
	}
}

func TestTurnAround(t *testing.T) {
	el := getNextElevator(5, Up)
	req := request{6, Down}
	el.Goals = requests{req}

	ncs := &NearestControlSystem{
		elevators: []*Elevator{el},
	}

	var cs ControlSystem = ncs
	cs.Step()

	if el.Floor != req.floor {
		t.Errorf("Expected floor %v, returned %v", req.floor, el.Floor)
	}

	if len(el.Goals) != 1 {
		t.Errorf("Expected 1 goal, returned %v", len(el.Goals))
	}

	if el.Direction != Down {
		t.Errorf("Expected elevator direction %v, returned %v", Down, el.Direction)
	}
}
