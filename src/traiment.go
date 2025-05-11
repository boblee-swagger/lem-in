package src

import (
	"errors"
	"strconv"
	"strings"
)

// parses a list of file lines to extract useful data(n of ants...)
func ValidateAntfarmData(fileLines []string) (AntFarm, error) {
	var data AntFarm
	var links []Link
	var rooms []Room
	var StartingRoom, EndingRoom Room
	var numAnts int
	var nextRoomType string

	for i := range fileLines {
		line := strings.TrimSpace(fileLines[i])
		if line == "" {
			continue
		}

		if numAnts == 0 {
			var err error
			numAnts, err = strconv.Atoi(line)
			if err == nil {
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
				StartingRoom = room
				nextRoomType = "" // Reset the flag
			} else if nextRoomType == "end" {
				EndingRoom = room
				nextRoomType = "" // Reset the flag
			}
			rooms = append(rooms, room)
			continue
		}

		linkParts := strings.Split(line, "-")
		if len(linkParts) == 2 {
			err := FormatLink(linkParts)
			if err != nil {
				return AntFarm{}, err
			}

			var link = Link{
				ConnectedRoom: linkParts,
			}
			links = append(links, link)
		}
	}

	// Validate the required fields
	if StartingRoom.Name == "" || EndingRoom.Name == "" || numAnts <= 0 {
		return AntFarm{}, errors.New("error: invalid data format")
	}

	return data.New(numAnts, StartingRoom, EndingRoom, rooms, links), nil
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

func FormatLink(link []string) (error) {
	//cyclic link
	if link[0] == link[1] {
		return errors.New("error: invalid link format")
	}
	return nil
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
	room.SetCordX(X)
	room.SetCordY(Y)
	return room, nil
}

