package main

import (
	"fmt"
	// "os"
	// "strconv"
)

func main() {
	// Examples:
	door := NewDoor(1, "closed")
	floorRequestButton := NewFloorRequestButton(1, "OFF", 2, "Down")
	callButton := NewCallButton(1, "OFF", 2, "Up")
	elevator := NewElevator(1, "idle", 60, 6)
	elevator.floorRequestsList = append(elevator.floorRequestsList, 8)
	elevator.move()
	// column.servedFloorsList = append(column.servedFloorsList, 8)

	// column := NewColumn(1, "online", 60, 1, 5, false)


	// Prints
	fmt.Println("door is", door)
	fmt.Println("floorRequestButton", floorRequestButton)
	fmt.Println("callButton", callButton)
	fmt.Println("elevator", elevator)
	fmt.Println("floorRequestsList", elevator.floorRequestsList)
	// fmt.Println("servedFloorsList", column.servedFloorsList)
	// fmt.Println("column", column)
	

	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	runScenario(scenarioNumber)
	// } else {
	// 	fmt.Println(err)
	// }
}
