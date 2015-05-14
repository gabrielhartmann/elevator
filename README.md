# Elevator Control System

![alt tag](https://raw.github.com/gabrielhartmann/elevator/master/elevator_demo.png)

This software models an elevator control system as a discrete time simulation.  At each time step requests for service can be made, elevators may move, and elevators may visit floors.  The systems also works to adhere to the standard expected constraints of an elevator system.

1. Once an elevator is serviciving requests in a particular direction (e.g. going down), it continues in this direction until all requests in this direction are honored.  Note that the source of the request for service, internal to the elevator or a button press from a floor, is irrelevant.

2. It does not take passengers in the wrong direction.  That is, for example, if the elevator is heading up, it will not open its doors to service a request to go down even if the elevator is on the floor of the request.

3. All requests for service are honored.  This may sound trivial, but the system must guarantee that no customers are ever starved of service due to higher priority requests.  To avoid starvation in this implementation all requests are immediately assigned to elevators upon arrival.  All elevators are guaranteed to service all requests, so no requests are starved


The three major components are the control system, the scheduler, and the elevators themselves.

1. The control system serves mostly as an interface and point of common contact for the other two components.  On pickup requests for example it consults the scheduler and uses this information to inform an elevator of a new request.  It also allows requests to query the status of all the elevators, and provides a means for causing the simulation to move forward in time.

2. The scheduler, when given a request from a particular floor in a given direction (e.g. Floor 2, Up) generates a score for each elevator.  The elevator with the lowest score is assigned the request and will service it as soon as possible, subject to the constraints mentioned above.  The score is derived from two components: guaranteed, and worst-case distance to the request.  The guaranteed distance is how far an elevator must travel if nothing changes before it can service the request.  The worst-case distance is how far an elevator could travel if customer requests conspire to delay the elevator's progress towards the request in question.  These are weighted in the example code in a 70%, 30% split to the guaranteed and worst-case distances respectively.

3. The elevators respond to step calls from the control system, and provide an interface for adding goals.  The move in one direction as long as there are requests to service in that direction, and they only visit floors when they are going in the requested direction. 

The combination of the scheduler and elevator behaviors attempts to minimize the distance traveled by elevators and thus the time customers wait for requests to be serviced, and the time they are in transit.
