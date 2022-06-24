package main

type Elevator struct {
	ID int
	status string
	amountOfFloors int
	currentFloor int
	direction string
	door Door
	overweight bool
	obstruction bool
	completedRequestsList []int
	floorRequestsList []int
}

func NewElevator(ID int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	elevator := &Elevator{ID: ID, status: "idle", amountOfFloors: _amountOfFloors, currentFloor: _currentFloor}

	return elevator
}

// func (e *Elevator) move() {
	
// }