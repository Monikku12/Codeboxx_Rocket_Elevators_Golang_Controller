package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID int
	status string
	floorRequestButtonID int
	buttonFloor int
	direction string
}

func NewFloorRequestButton(ID int, status string, _buttonFloor int, _direction string) *FloorRequestButton {
	floorRequestButton := &FloorRequestButton{ID: ID, status: status, buttonFloor: _buttonFloor, direction: _direction}

	return floorRequestButton
}
