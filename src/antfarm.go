package src

type AntFarm struct {
	NumberOfAnts int
	StartingRoom Room
	EndingRoom   Room
	Rooms        []Room
	Links        []Link
}

func (a *AntFarm) New (numAnts int, startingRoom, endingRoom Room, rooms []Room, links []Link) AntFarm{
	return AntFarm{
		NumberOfAnts: numAnts,
		StartingRoom: startingRoom,
		EndingRoom: endingRoom,
		Rooms: rooms,
		Links: links,
	}
}

func (a *AntFarm) SetNumberOfAnts(ants int) {
	a.NumberOfAnts = ants
}

func (a *AntFarm) GetNumberOfAnts() int {
	return a.NumberOfAnts
}

func (a *AntFarm) SetStartingRoom(room Room) {
	a.StartingRoom = room
}

func (a *AntFarm) GetStartingRoom() Room {
	return a.StartingRoom
}

func (a *AntFarm) SetEndingRoom(room Room) {
	a.EndingRoom = room
}

func (a *AntFarm) GetEndingRoom() Room {
	return a.EndingRoom
}

func (a *AntFarm) GetRooms() []Room {
	return a.Rooms
}


func (a *AntFarm) GetLinks() []Link {
	return a.Links
}

