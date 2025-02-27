package src

import (
	"strconv"
	"strings"
)

func SetRoom(str string) Room {
	room := Room{}

	data := strings.Split(str, " ")
	room.Name = data[0]
	room.CordX, _ = strconv.Atoi(data[1])
	room.CordY, _ = strconv.Atoi(data[2])

	return room
}
