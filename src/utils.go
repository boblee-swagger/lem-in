package src

import (
	"bufio"
	"errors"
	"os"
	//"fmt"
)


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

func GetRoom(roomName string, rooms []Room) Room {
	for _,room := range rooms{
		if room.Name == roomName {
			return room
		}
	}
	//remind that you can get an empty room
	return Room{}
}

