package main

import (
	"fmt"
)

type Elevator struct {
	ID int
	status string
	amountOfFloors int
	currentFloor int
	direction string
	overweight bool
	obstruction bool
	door Door
	completedRequestsList []int
	floorRequestsList []int
}

func NewElevator(ID int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	
	door := Door{ID: ID, status: "closed"}

	elevator := &Elevator{ID: ID, status: "idle", amountOfFloors: _amountOfFloors, currentFloor: _currentFloor, Door: door}
	
	return elevator
}

func (e *Elevator) move() {
	
	for true{
		if len(e.floorRequestsList) != 0 {
		var destination = e.floorRequestsList[0]
		e.status = "moving"
			if e.currentFloor < destination {
				e.direction = "up"
				e.sortFloorList()
				if e.currentFloor < destination {
					e.currentFloor++
					var screenDisplay = e.currentFloor
					}
			}
		}
	}
}

func (e *Elevator) operateDoors() {

	e.door.status = "opened"
	fmt.Println("Wait 5 seconds")
	if e.overweight == false {
		e.door.status = "closing"
		if e.obstruction == false {
			e.door.status = "closed"
			break
		}
		if e.operateDoors(){
			break
		}
		break
	}
	if e.overweight == true {
		fmt.Println("Activate overweight alarm")
	}
	e.operateDoors()
	break
}











func (e *Elevator) addNewRequest() {

	// if e.floorRequestsList.
}
