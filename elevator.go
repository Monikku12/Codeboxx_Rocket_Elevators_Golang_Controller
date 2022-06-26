package main

import (
	"fmt"
	"sort"
)

// Elevators in a column
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
	screenDisplay int
}

func NewElevator(ID int, _status string, _amountOfFloors int, _currentFloor int) Elevator {
	
	door := Door{ID: ID, status: "closed"}

	elevator := Elevator{ID: ID, status: "idle", amountOfFloors: _amountOfFloors, currentFloor: _currentFloor, door: door}
	
	return elevator
}

 // Function in charge of moving the elevator in the columns
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
					e.screenDisplay = e.currentFloor
					}
					break
			}
			if e.currentFloor > destination {
				e.direction = "down"
				e.sortFloorList()
				if e.currentFloor > destination {
					e.currentFloor--
					e.screenDisplay = e.currentFloor
					}
					break
			}
			e.status = "stopped"
			e.operateDoors()
			e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
			e.floorRequestsList = append([]int{1}, e.floorRequestsList...)
		}
		e.status = "idle"
	}
}

// Function in charge of sorting the floors in the right order depending if the elevator is going up or down
func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)
	}
	if e.direction != "up" {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}

// Function in charge of opening and closing the elevator doors
func (e *Elevator) operateDoors() {

	e.door.status = "opened"
	fmt.Println("Wait 5 seconds")
	if e.overweight == false {
		e.door.status = "closing"
		if e.obstruction == false {
			e.door.status = "closed"
		}
		if e.obstruction == true {
			e.operateDoors() 
		}
	}
	if e.overweight == true {
		fmt.Println("Activate overweight alarm")
	}
	e.operateDoors()	
}

func (e *Elevator) addNewRequest(requestedFloor int) {

	if len(e.floorRequestsList) == 0 {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor{
		e.direction = "down"
	}
}