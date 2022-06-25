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
	elevator.move()


	// Prints
	fmt.Println("door is", door)
	fmt.Println("floorRequestButton", floorRequestButton)
	fmt.Println("callButton", callButton)
	fmt.Println("elevator", elevator)

	

	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	runScenario(scenarioNumber)
	// } else {
	// 	fmt.Println(err)
	// }
}
