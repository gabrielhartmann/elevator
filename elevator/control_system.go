package elevator

type ControlSystem interface {
	Status() []*Elevator
	Pickup(floor int, dir int)
	Step()
}

type NearestControlSystem struct {
	elevators []*Elevator
	scheduler Scheduler
}

func NewNearestControlSystem(elevators []*Elevator, scheduler Scheduler) NearestControlSystem {
	return NearestControlSystem{
		elevators: elevators,
		scheduler: scheduler,
	}
}

func (cs *NearestControlSystem) Status() []*Elevator {
	return cs.elevators
}

func (cs *NearestControlSystem) Pickup(floor int, dir int) {
	el := cs.scheduler.Schedule(floor, dir, cs.elevators)
	el.AddGoal(request{floor, dir})
}

func (cs *NearestControlSystem) Step() {
	for _, elev := range cs.elevators {
		elev.step()
	}
}
