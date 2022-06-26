package main

type Battery struct {
	ID							int
	amountOfElevators			int
	currentFloor				int
	amountOfColumns				int
	amountOfFloors				int
	amountOfBasements			int
	amountOfElevatorPerColumn	int
	status						string
	elevator					[]Elevator
	isBasement					bool
	column						[]Column
	columnsList					[]Column
	servedFloorsList			[]Column
	floorRequestsButtonsList	[]FloorRequestButton
}

func NewBattery(ID int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int, _isBasement bool) *Battery {
	battery := &Battery{ID: ID, status: "online", amountOfColumns: _amountOfColumns, amountOfFloors: _amountOfFloors, amountOfBasements: _amountOfBasements, amountOfElevatorPerColumn: _amountOfElevatorPerColumn, servedFloorsList: _servedFloorsList, amountOfElevators: _amountOfElevators, isBasement: _isBasement}

	if _amountOfBasements > 0 {
		// battery.createBasementFloorRequestButtons(_amountOfBasements)
		// battery.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	// battery.createFloorRequestButtons(_amountOfFloors)
	// battery.createColumns(_amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn)

	return battery
}

// Creates the column in charge of the basement floors.
func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	List<int> servedFloorsList = new List<int>()
	int floor = -1
	for (int i = 0; i < _amountOfBasements; i++)
	{
		servedFloorsList.Add(floor)
		floor--
	}
	Column column = new Column(ID, "online", _amountOfBasements, servedFloorsList, _amountOfElevatorPerColumn, true)
	this.columnsList.Add(column)
}

// // Creates the column in charge of the upper floors.
// public void createColumns(int _amountOfColumns, int _amountOfFloors, int _amountOfBasements, int _amountOfElevatorPerColumn)
// {
// 	double d = _amountOfFloors / _amountOfColumns;
// 	int amountOfFloorsPerColumn = (int)Math.Ceiling(d);
// 	int floor = 1;
// 	for (int i = 0; i < _amountOfColumns; i++)
// 	{

// 		List<int> servedFloors = new List<int>();
// 		for (int a = 0; a < amountOfFloorsPerColumn; a++)
// 		{
// 			if (floor <= _amountOfFloors)
// 			{
// 				servedFloors.Add(floor);
// 				floor++;
// 			}
// 		}

// 		Column column = new Column(i + 1, "online", _amountOfFloors, servedFloors, _amountOfElevatorPerColumn, false);
// 		this.columnsList.Add(column);
// 	}
// }

// // Creates the central panel in the lobby to call an elevator to go to a specific upper floor.
// public void createFloorRequestButtons(int _amountOfFloors)
// {
// 	int buttonFloor = 1;
// 	for (int i = 0; i < _amountOfFloors; i++)
// 	{
// 		FloorRequestButton floorRequestButton = new FloorRequestButton(i, "OFF", buttonFloor, "Up");
// 		this.floorRequestsButtonsList.Add(floorRequestButton);
// 		buttonFloor++;
// 	}
// }

// // Creates the central panel in the lobby to call an elevator to go to a specific basement floor.
// public void createBasementFloorRequestButtons(int _amountOfBasements)
// {
// 	int buttonFloor = -1;
// 	for (int i = 0; i < _amountOfBasements; i++)
// 	{
// 		FloorRequestButton floorRequestButton = new FloorRequestButton(i, "OFF", buttonFloor, "Down");
// 		this.floorRequestsButtonsList.Add(floorRequestButton);
// 		buttonFloor--;
// 	}
// }







// func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	
// }


// //Simulate when a user press a button at the lobby
// func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

// }