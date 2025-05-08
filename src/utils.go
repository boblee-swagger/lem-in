package src

import (
	"bufio"
	"errors"
	"os"
	"strings"
	//"fmt"
)

// make connections between different connected rooms
func DefineTheGraph(rooms []Room, links []string) [][]string {
	connections := make([][]string, len(rooms))
	for i := range rooms {
		//The first element on the connection is the room, and there are the connected rooms.
		connections[i] = append(connections[i], rooms[i].Name)
		for j := range links {
			room := rooms[i]
			link := strings.Split(links[j], "-")

			if room.Name == link[0] {
				connections[i] = append(connections[i], link[1])
			} else if room.Name == link[1] {
				connections[i] = append(connections[i], link[0])
			}
		}
	}
	return connections
}

func FileContent(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error on opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines, nil
}
