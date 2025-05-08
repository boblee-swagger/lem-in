package src

type Link struct {
	ConnectedRooms []string
	Weight int
}

func (e *Link) SetLink(room1, room2 string, weight int) {
	e.ConnectedRooms = append(e.ConnectedRooms, room1, room2)
	e.Weight = weight
}

func (e *Link) SetWeight(weight int) {
	e.Weight = weight
}

func (e *Link) GetWeight() int {
	return e.Weight
}

func (e *Link) GetConnectedRooms() []string {
	return e.ConnectedRooms
}

