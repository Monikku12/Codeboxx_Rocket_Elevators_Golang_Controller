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

	// Prints
	fmt.Println(door)
	fmt.Println(floorRequestButton)
	fmt.Println(callButton)

	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	runScenario(scenarioNumber)
	// } else {
	// 	fmt.Println(err)
	// }
}
