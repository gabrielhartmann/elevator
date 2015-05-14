package elevator

type Scheduler interface {
	Schedule(floor int, dir int, els []*Elevator) *Elevator
}

type DistanceScheduler struct {
	FloorCount       int
	GuaranteedWeight float64
	WorstCaseWeight  float64
}

func NewDistanceScheduler(floorCount int, guaranteedWeight float64, worstCaseWeight float64) DistanceScheduler {
	return DistanceScheduler{
		FloorCount:       floorCount,
		GuaranteedWeight: guaranteedWeight,
		WorstCaseWeight:  worstCaseWeight,
	}
}

func (s *DistanceScheduler) Schedule(floor int, dir int, els []*Elevator) *Elevator {
	el, _ := s.scheduleInternal(floor, dir, els)
	return el
}

func (s *DistanceScheduler) scheduleInternal(floor int, dir int, els []*Elevator) (*Elevator, float64) {
	if len(els) == 0 {
		return nil, -1
	}

	elev := els[0]
	dist := s.getDistance(floor, dir, elev)

	for _, el := range els {
		d := s.getDistance(floor, dir, el)

		if d == 0 {
			return el, 0
		}

		if d < dist {
			dist = d
			elev = el
		}
	}

	return elev, dist
}

// A lower distance is better.  It is a weighted combination of the guaranteed
// distance an elevator must travel to service a request, and the worst case
// scenario in which subsequent requests take the elevator on the longest possible
// route to service this request.
func (s *DistanceScheduler) getDistance(floor int, dir int, el *Elevator) float64 {
	gDist := float64(s.getGuaranteedDistance(floor, dir, el))
	wDist := float64(s.getWorstCaseDistance(floor, dir, el))

	return (s.GuaranteedWeight * gDist) + (s.WorstCaseWeight * wDist)
}

// Guaranteed distance is a measure that indicates how far an elevator must travel
// to reach the given request floor if no further changes were made to the elevators
// goals.  There is the additional constraint that an elevator will not change
// direction until all goals in the current direction have been serviced.
// The implication is that in some cases an elevator will pass by a request and return
// to it later, to maintain its direction as much as possible
func (s *DistanceScheduler) getGuaranteedDistance(floor int, dir int, el *Elevator) int {
	if floor == el.Floor && dir == el.Direction {
		return 0
	}

	// Idle elevators can go directly to requests in any direction
	if el.Direction == Idle {
		if el.Floor > floor {
			return el.Floor - floor
		} else {
			return floor - el.Floor
		}
	}

	if dir == Down {
		return s.getGuaranteedDownDistance(floor, el)
	} else {
		return s.getGuaranteedUpDistance(floor, el)
	}

	return 0
}

func (s *DistanceScheduler) getGuaranteedDownDistance(floor int, el *Elevator) int {
	highestFloor := el.Goals[len(el.Goals)-1].floor
	lowestFloor := el.Goals[0].floor
	return getDownDistance(floor, el, lowestFloor, highestFloor)
}

func (s *DistanceScheduler) getGuaranteedUpDistance(floor int, el *Elevator) int {
	highestFloor := el.Goals[len(el.Goals)-1].floor
	lowestFloor := el.Goals[0].floor
	return getUpDistance(floor, el, lowestFloor, highestFloor)
}

func (s *DistanceScheduler) getWorstCaseDistance(floor int, dir int, el *Elevator) int {
	if floor == el.Floor && dir == el.Direction {
		return 0
	}

	// Idle elevators can go directly to requests in any direction
	if el.Direction == Idle {
		if el.Floor > floor {
			return el.Floor - floor
		} else {
			return floor - el.Floor
		}
	}

	if dir == Down {
		return s.getWorstCaseDownDistance(floor, el)
	} else {
		return s.getWorstCaseUpDistance(floor, el)
	}

	return 0

}

func (s *DistanceScheduler) getWorstCaseDownDistance(floor int, el *Elevator) int {
	highestGoal := s.FloorCount - 1
	lowestGoal := 0
	return getDownDistance(floor, el, lowestGoal, highestGoal)
}

func (s *DistanceScheduler) getWorstCaseUpDistance(floor int, el *Elevator) int {
	highestGoal := s.FloorCount - 1
	lowestGoal := 0
	return getUpDistance(floor, el, lowestGoal, highestGoal)
}

func getDownDistance(floor int, el *Elevator, lowestGoal int, highestGoal int) int {
	switch el.Direction {
	case Down:
		if el.Floor > floor {
			return el.Floor - floor
		} else {
			return (floor - lowestGoal) + (el.Floor - lowestGoal)
		}
	case Up:
		if highestGoal < floor {
			return floor - el.Floor
		} else {
			return (highestGoal - floor) + (highestGoal - el.Floor)
		}
	default:
		return 0
	}
}

func getUpDistance(floor int, el *Elevator, lowestGoal int, highestGoal int) int {
	switch el.Direction {
	case Up:
		if el.Floor < floor {
			return floor - el.Floor
		} else {
			return (highestGoal - floor) + (highestGoal - el.Floor)
		}
	case Down:
		if lowestGoal > floor {
			return el.Floor - floor
		} else {
			return (floor - lowestGoal) + (el.Floor - lowestGoal)
		}
	default:
		return 0
	}
}
