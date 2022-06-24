package main

type Door struct {
	ID int
	status string
}

func NewDoor(ID int, status string) *Door {
	door := &Door{ID: ID, status: "closed"}

	return door
}