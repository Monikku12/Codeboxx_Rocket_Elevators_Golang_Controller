package main

type Column struct {
	ID int
	status string
	amountOfFloors int
	servedFloorsList []int
	amountOfElevators int
	isBasement bool
	elevatorsList []Elevator
	callButtonsList []CallButton

}

func NewColumn(_id, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	
}

// //Simulate when a user press a button on a floor to go back to the first floor
// func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	
// }