package src

import (
	"errors"
	"strconv"
	"strings"
)

// parses a list of file lines to extract useful data(n of ants...)
func ValidateAntfarmData(fileLines []string) (AntFarm, error) {
	var data = AntFarm{}
	var nextRoomType string

	for i := 0; i < len(fileLines); i++ {
		line := strings.TrimSpace(fileLines[i])
		if line == "" {
			continue
		}

		if data.NumberOfAnts == 0 {
			numAnts, err := strconv.Atoi(line)
			if err == nil {
				data.NumberOfAnts = numAnts
				continue
			}
	}

		// Check for special commands
		if line == "##start" {
			nextRoomType = "start"
			continue
		} else if line == "##end" {
			nextRoomType = "end"
			continue
		}

		//check if line is a comment
		parts := strings.Fields(line)
		if len(parts) == 3 {
			room, err := FormatRoom(parts)
			if err != nil {
				return AntFarm{}, err
			}
			
			// Assign room based on previous command
			if nextRoomType == "start" {
				data.StartingRoom = room
				nextRoomType = "" // Reset the flag
			} else if nextRoomType == "end" {
				data.EndingRoom = room
				nextRoomType = "" // Reset the flag
			} 
			data.Rooms = append(data.Rooms, room)
			continue
		}

		linkParts := strings.Split(line, "-")
		if len(linkParts) == 2 {
			link, err := FormatLink(linkParts)
			if err != nil {
				return AntFarm{}, err
			}
			data.Links = append(data.Links, link)
			continue
		}
	}
	
	// Validate the required fields
	if data.StartingRoom.Name == "" || data.EndingRoom.Name == "" || data.NumberOfAnts <= 0 {
		return AntFarm{}, errors.New("error: invalid data format")
	}
	return data, nil
}

// Verify room name's format and its coordonate
func FormatRoom(room []string) (Room, error) {
	//room never started with # or L
	if strings.HasPrefix(room[0], "L") || strings.HasPrefix(room[0], "#") {
		return Room{}, errors.New("error: invalid room format")
	}
	r, err := ParseRoom(room)
	if err != nil {
		return Room{}, err
	}
	return r, nil
}

func FormatLink(link []string) (Link, error) {
	//cyclic link
	if link[0] == link[1] {
		return Link{}, errors.New("error: invalid link format")
	}
	result, err := ParseLink(link)
	if err != nil {
		return Link{}, err
	}

	return result, nil
}

func ParseRoom(data []string) (Room, error) {
	room := Room{}
	room.Name = data[0]
	X, errX := strconv.Atoi(data[1])
	Y, errY := strconv.Atoi(data[2])
	//room coordinate must be integer
	if errX != nil || errY != nil {
		return Room{}, errors.New("error: invalid room coordinates")
	}
	room.CordX = X
	room.CordY = Y
	return room, nil
}

func ParseLink(data []string) (Link, error) {
	link := Link{}
	link.ConnectedRooms = append(link.ConnectedRooms, data[0])
	link.ConnectedRooms = append(link.ConnectedRooms, data[1])
	return link, nil
}
