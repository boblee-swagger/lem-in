package src

type Room struct {
	Name  string
	CordX int
	CordY int
	Neighbors []Room
}

func (r *Room) SetNeighbors(ant AntFarm) {
	for _, link := range ant.Links{
		if link.ConnectedRoom[0]== r.Name {
			room := GetRoom(link.ConnectedRoom[1], ant.Rooms)
			r.Neighbors = append(r.Neighbors, room)
		}else if link.ConnectedRoom[1] == r.Name {
			room := GetRoom(link.ConnectedRoom[1], ant.Rooms)
			r.Neighbors = append(r.Neighbors, room)
		}
	}
} 

func (r *Room) GetNeighbors() []Room {
	return r.Neighbors
}

func (r *Room) GetCordX() int {
	return r.CordX
}

func (r *Room) GetCordY() int {
	return r.CordY
}

func (r *Room) SetCordX(x int) {
	r.CordX = x
}

func (r *Room) SetCordY(y int) {
	r.CordY = y
}

func (r *Room) SetName(name string) {
	r.Name = name
}

func (r *Room) GetName() string {
	return r.Name
}
