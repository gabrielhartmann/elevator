package elevator

var elevatorId = 0

func getNextElevator(f int, d int) *Elevator {
	el := Elevator{
		Id:        elevatorId,
		Floor:     f,
		Direction: d,
		Goals:     requests{},
	}

	elevatorId++
	return &el
}

var distScheduler = &DistanceScheduler{
	FloorCount:       10,
	GuaranteedWeight: 0.70,
	WorstCaseWeight:  0.30,
}
