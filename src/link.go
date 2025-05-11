package src

import (
	"math"
)

type Link struct {
	ConnectedRoom []string
	Weight int
}

func (e *Link) New (room1_name, room2_name string, weight int) Link{
	return Link{
		ConnectedRoom: []string{room1_name, room2_name},
		Weight: weight,
	}
}

func (e *Link) SetWeight(rooms []Room) {
	source := GetRoom(e.ConnectedRoom[0], rooms)
	dest := GetRoom(e.ConnectedRoom[1], rooms)

	d := math.Pow(float64(source.GetCordX() + dest.GetCordX()), 2) + math.Pow(float64(source.GetCordY() + dest.GetCordY()), 2)
	e.Weight = int(math.Sqrt(d))	
}

func (e *Link) GetWeight() int {
	return e.Weight
}

func (e *Link) GetConnectedRooms() []string {
	return e.ConnectedRoom
}

func (e * Link) SetConnectedRooms(rooms1_name, room2_name string) {
	e.ConnectedRoom = []string{rooms1_name, room2_name}
}

