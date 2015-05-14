package elevator

import (
	"math"
	"testing"
)

var scheduler Scheduler = distScheduler

func TestGetGuaranteedDistance(t *testing.T) {
	el := getNextElevator(5, Idle)
	guaranteedDistanceHelper(t, 3, Down, el, 2)
	guaranteedDistanceHelper(t, 8, Down, el, 3)
	guaranteedDistanceHelper(t, 3, Up, el, 2)
	guaranteedDistanceHelper(t, 8, Up, el, 3)

	// Tests below are in groups of three with requests:
	// 1. On route to a goal
	// 2. Passed the farthest goal in the current direction
	// 3. Behind the elevator

	req := request{3, Down}
	el.Goals = requests{req}
	el.Direction = Down
	guaranteedDistanceHelper(t, 4, Down, el, 1)
	guaranteedDistanceHelper(t, 2, Down, el, 3)
	guaranteedDistanceHelper(t, 7, Down, el, 6)

	guaranteedDistanceHelper(t, 4, Up, el, 3)
	guaranteedDistanceHelper(t, 2, Up, el, 3)
	guaranteedDistanceHelper(t, 7, Up, el, 6)

	req = request{7, Down}
	el.Goals = requests{req}
	el.Direction = Up
	guaranteedDistanceHelper(t, 6, Up, el, 1)
	guaranteedDistanceHelper(t, 8, Up, el, 3)
	guaranteedDistanceHelper(t, 3, Up, el, 6)

	guaranteedDistanceHelper(t, 6, Down, el, 3)
	guaranteedDistanceHelper(t, 8, Down, el, 3)
	guaranteedDistanceHelper(t, 3, Down, el, 6)
}

func guaranteedDistanceHelper(t *testing.T, floor int, dir int, el *Elevator, expectedDistance int) {
	gDist := distScheduler.getGuaranteedDistance(floor, dir, el)

	if expectedDistance != gDist {
		t.Errorf("Expected guaranteed distance: %v, returned: %v", expectedDistance, gDist)
	}
}

func TestGetWorstCaseDistance(t *testing.T) {
	el := getNextElevator(5, Idle)
	worstCaseDistanceHelper(t, 3, Down, el, 2)
	worstCaseDistanceHelper(t, 8, Down, el, 3)
	worstCaseDistanceHelper(t, 3, Up, el, 2)
	worstCaseDistanceHelper(t, 8, Up, el, 3)

	// Tests below are in groups of three with requests:
	// 1. On route to a goal
	// 2. Passed the farthest goal in the current direction
	// 3. Behind the elevator

	req := request{3, Down}
	el.Goals = requests{req}
	el.Direction = Down
	worstCaseDistanceHelper(t, 4, Down, el, 1)
	worstCaseDistanceHelper(t, 2, Down, el, 3)
	worstCaseDistanceHelper(t, 7, Down, el, 12)

	worstCaseDistanceHelper(t, 4, Up, el, 9)
	worstCaseDistanceHelper(t, 2, Up, el, 7)
	worstCaseDistanceHelper(t, 7, Up, el, 12)

	req = request{7, Down}
	el.Goals = requests{req}
	el.Direction = Up
	worstCaseDistanceHelper(t, 6, Up, el, 1)
	worstCaseDistanceHelper(t, 8, Up, el, 3)
	worstCaseDistanceHelper(t, 3, Up, el, 10)

	worstCaseDistanceHelper(t, 6, Down, el, 7)
	worstCaseDistanceHelper(t, 8, Down, el, 5)
	worstCaseDistanceHelper(t, 3, Down, el, 10)
}

func worstCaseDistanceHelper(t *testing.T, floor int, dir int, el *Elevator, expectedDistance int) {
	worstDist := distScheduler.getWorstCaseDistance(floor, dir, el)

	if expectedDistance != worstDist {
		t.Errorf("Expected worst case distance: %v, returned: %v", expectedDistance, worstDist)
	}
}

func TestGetDistance(t *testing.T) {
	el := getNextElevator(5, Idle)
	distanceHelper(t, 3, Down, el, 2)
	distanceHelper(t, 8, Up, el, 3)

	req := request{3, Down}
	el.Goals = requests{req}
	el.Direction = Down
	distanceHelper(t, 4, Down, el, 1)
	distanceHelper(t, 7, Down, el, 7.8)
	distanceHelper(t, 4, Up, el, 4.8)
	distanceHelper(t, 7, Up, el, 7.8)

	req = request{7, Down}
	el.Goals = requests{req}
	el.Direction = Up
	distanceHelper(t, 6, Down, el, 4.2)
	distanceHelper(t, 3, Down, el, 7.2)
	distanceHelper(t, 6, Up, el, 1)
	distanceHelper(t, 3, Up, el, 7.2)
}

func distanceHelper(t *testing.T, floor int, dir int, el *Elevator, expectedDistance float64) {
	dist := distScheduler.getDistance(floor, dir, el)

	var epsilon float64 = 0.0001

	if math.Abs(expectedDistance-dist) > epsilon {
		t.Errorf("Expected distance: %v, returned: %v", expectedDistance, dist)
	}
}

// All elevators and requests going Down
func TestScheduleDownDownDown(t *testing.T) {
	el1 := getNextElevator(5, Down)
	req := request{2, Down}
	el1.Goals = requests{req}

	el2 := getNextElevator(8, Down)
	req = request{2, Down}
	el2.Goals = requests{req}

	elSched := scheduler.Schedule(6, Down, []*Elevator{el1, el2})

	if elSched != el2 {
		t.Errorf("Expected elevator %v, returned %v", el2, elSched)
	}
}

// Elevators are going down, but the request is up
func TestScheduleDownUpDown(t *testing.T) {
	el1 := getNextElevator(5, Down)
	req := request{2, Down}
	el1.Goals = requests{req}

	el2 := getNextElevator(8, Down)
	req = request{2, Down}
	el2.Goals = requests{req}

	elSched := scheduler.Schedule(6, Up, []*Elevator{el1, el2})

	if elSched != el1 {
		t.Errorf("Expected elevator %v, returned %v", el2, elSched)
	}
}
