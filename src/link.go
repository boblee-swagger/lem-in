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

func (e *Link) SetWeight(room1, room2 Room) {
	//like maths we calcualte distance between two points basing on their coordinate
	d := math.Pow(float64(room1.CordX + room2.CordX), 2) + math.Pow(float64(room1.CordY + room2.CordY), 2)
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

