package main

type Column struct {
	ID							int
	status						string
	amountOfFloors				int
	servedFloorsList			[]int
	amountOfElevators			int
	isBasement					bool	
	elevatorsList				[]Elevator
	callButtonsList				[]CallButton
	buttonFloor					int
	bestElevatorInformations	BestElevatorInformations
	elevator					Elevator
}

type BestElevatorInformations struct {
	bestElevator		Elevator
	bestScore			int
	referenceGap		int
}

func NewColumn(ID int, _status string, _amountOfFloors int, _servedFloorsList []int, _amountOfElevators int, _isBasement bool) Column {
	
	column := Column{ID: ID, status: "idle", amountOfFloors: _amountOfFloors, servedFloorsList: _servedFloorsList, amountOfElevators: _amountOfElevators, isBasement: _isBasement}

	column.createElevators(column.amountOfFloors, column.amountOfElevators)
	column.createCallButtons(column.amountOfFloors, column.isBasement)
	
	return column
}

func NewBestElevatorInformations(_bestElevator Elevator, bestScore int, referenceGap int) BestElevatorInformations {

	bestElevatorInformations := BestElevatorInformations{bestElevator: _bestElevator, bestScore: 6, referenceGap: 10000000}
	
	return bestElevatorInformations
}

 // Creates the buttons outside the elevators to call them
func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
            if _isBasement == true {
                c.buttonFloor = -1
                for i := 0; i < _amountOfFloors; i++ {
					callButton := NewCallButton(i + 1, "OFF", c.buttonFloor, "Up")
                    c.callButtonsList = append(c.callButtonsList, callButton)
                    c.buttonFloor--
				}
            }
            if _isBasement == false {
                c.buttonFloor = 1
                for i := 0; i < _amountOfFloors; i++ {
                    callButton := NewCallButton(i + 1, "OFF", c.buttonFloor, "Down")
                    c.callButtonsList = append(c.callButtonsList, callButton)
                    c.buttonFloor++
                }
            }
        }

// Creates the elevators in the columns
func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	for i := 0; i < _amountOfElevators; i++ {
		elevator := NewElevator(i + 1, "idle", _amountOfFloors, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(userPosition int, direction string) Elevator {
	elevator := c.findElevator(userPosition, direction)
	elevator.addNewRequest(userPosition)
	elevator.move()
	elevator.addNewRequest(1) //Always 1 because the user can only go back to the lobby
	elevator.move()
	return elevator
}

//We use a score system depending on the current elevators state. Since the bestScore and the referenceGap are
//higher values than what could be possibly calculated, the first elevator will always become the default bestElevator,
//before being compared with to other elevators. If two elevators get the same score, the nearest one is prioritized. Unlike
//the classic algorithm, the logic isn't exactly the same depending on if the request is done in the lobby or on a floor.
func (c *Column) findElevator(requestedFloor int, requestedDirection string) Elevator {
	bestElevatorInformations := NewBestElevatorInformations(c.elevatorsList[0], 6, 10000000)

	if requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			// The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed.
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is at the lobby and has no requests
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is lower than me and is coming up. It means that I'm requesting an elevator to go to a basement, and the elevator is on it's way to me.
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is above me and is coming down. It means that I'm requesting an elevator to go to a floor, and the elevator is on it's way to me
			} else if 1 < elevator.currentFloor && elevator.direction == "down"	{
				bestElevatorInformations = c.checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is not at the first floor, but doesn't have any request
			} else if elevator.status == "idle"	{
				bestElevatorInformations = c.checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is not available, but still could take the call if nothing better is found
			} else
			{
				bestElevatorInformations = c.checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	
	} else {
		for _, elevator := range c.elevatorsList {
			// The elevator is at the same level as me, and is about to depart to the first floor
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevatorInformations = c.checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is lower than me and is going up. I'm on a basement, and the elevator can pick me up on it's way
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is higher than me and is going down. I'm on a floor, and the elevator can pick me up on it's way
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevatorInformations = c.checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is idle and has no requests
			} else if elevator.status == "idle"	{
				bestElevatorInformations = c.checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			
			// The elevator is not available, but still could take the call if nothing better is found
			} else
			{
				bestElevatorInformations = c.checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	}
	return bestElevatorInformations.bestElevator
}



// Compare the elevators score to choose the best one
func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestElevatorInformations BestElevatorInformations, floor int) BestElevatorInformations {
	if scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = scoreToCheck
		bestElevatorInformations.bestElevator = newElevator
		bestElevatorInformations.referenceGap = Abs(newElevator.currentFloor - floor)
	} else if bestElevatorInformations.bestScore == scoreToCheck {
		gap := Abs(newElevator.currentFloor - floor)
		if bestElevatorInformations.referenceGap > gap {
			bestElevatorInformations.bestElevator = newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return bestElevatorInformations
}