package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	ID int
	status string
	floor int
	direction string
}

func NewCallButton(ID int, _status string, _floor int, _direction string) CallButton {
	callButton := CallButton{ID: ID, status: _status, floor: _floor, direction: _direction}

	return callButton
}