package main

import (
	"fmt"
	// "os"
	// "strconv"
)

func main() {
	door := NewDoor(1, "closed")


	fmt.Println(door)

	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	runScenario(scenarioNumber)
	// } else {
	// 	fmt.Println(err)
	// }
}
