package elevator

import (
	"github.com/Sirupsen/logrus"
	"math"
	"sort"
)

const (
	Idle = iota
	Up
	Down
)

type Elevator struct {
	Id        int
	Floor     int
	Direction int
	Goals     requests
}

func NewElevator(id int, floor int) *Elevator {
	return &Elevator{
		Id:        id,
		Floor:     floor,
		Direction: Idle,
		Goals:     requests{},
	}
}

func (e *Elevator) AddGoal(req request) {
	if len(e.Goals) == 0 {
		if req.floor > e.Floor {
			e.Direction = Up
		} else if req.floor < e.Floor {
			e.Direction = Down
		} else {
			e.Direction = req.direction
		}
	}

	e.Goals = append(e.Goals, req)
	sort.Sort(e.Goals)
}

func (e *Elevator) step() {
	// Without goals, there is no change
	if len(e.Goals) == 0 {
		e.Direction = Idle
		return
	}

	// The elevator is on a goal floor
	if indices, ok := sliceContains(e.Goals, e.Floor); ok {
		for _, i := range indices {
			// Visit it, if we are going in the right direction
			if e.Goals[i].direction == e.Direction {
				e.visit(i)
				return
			}
		}
	}

	// Elevator should move towards a goal
	from := e.Floor
	e.move()
	to := e.Floor
	logrus.Infof("Elevator %v moved from %v, to %v", e, from, to)

	lowestGoal := e.Goals[0]
	highestGoal := e.Goals[len(e.Goals)-1]

	// Adjust orientation to that of request
	// at limit of request range
	if e.Floor == highestGoal.floor {
		e.Direction = highestGoal.direction
	} else if e.Floor == lowestGoal.floor {
		e.Direction = lowestGoal.direction
	}
}

func (e *Elevator) visit(goalIndex int) {
	logrus.Infof("Elevator %v, visiting %v", e, e.Goals[goalIndex])

	e.Goals = append(e.Goals[:goalIndex], e.Goals[goalIndex+1:]...)
	if len(e.Goals) == 0 {
		e.Direction = Idle
	}
}

func (e *Elevator) move() {
	lowestFloor := e.Goals[0].floor
	highestFloor := e.Goals[len(e.Goals)-1].floor

	switch e.Direction {
	case Down:
		if lowestFloor < e.Floor {
			e.Floor--
		} else {
			e.Direction = Up
		}
	case Up:
		if highestFloor > e.Floor {
			e.Floor++
		} else {
			e.Direction = Down
		}
	case Idle:
		closestGoal := e.closestGoal()

		if closestGoal.floor < e.Floor {
			e.Direction = Down
		} else {
			e.Direction = Up
		}
	}
}

func (e *Elevator) closestGoal() request {
	distance := math.MaxFloat64
	var closestGoal request

	for _, g := range e.Goals {
		if d := math.Abs(float64(e.Floor - g.floor)); d < distance {
			distance = d
			closestGoal = g
		}
	}

	return closestGoal
}
