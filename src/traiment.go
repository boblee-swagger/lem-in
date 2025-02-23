package src

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	CordX int
	CordY int
}

type AntFarmData struct {
	NumberOfAnts  int
	StartingRoom string 
	EndingRoom string
	Rooms     []Room
}


// parses a list of file lines to extract useful data(n of ants...)
func ValidateAntfarmData(fileLines []string) (AntFarmData, error) {
	var data AntFarmData
	number_of_ants , err := strconv.Atoi(fileLines[0])
	if err != nil {
		return AntFarmData{}, errors.New("Error: Invalid data format, invalid number of ants")
	}
	data.NumberOfAnts = number_of_ants
	
	for i:=1; i< len(fileLines); i++ {
		if fileLines[i] == "##start" {
			if i+1 < len(fileLines){
				data.StartingRoom = fileLines[i+1]
				i++
			}
		}else if fileLines[i] == "##end" {
			if i+1 < len(fileLines){
				data.EndingRoom = fileLines[i+1]
				i++
			}
		}else if err := ValidateRoom(fileLines[i]); err != nil {
			return AntFarmData{}, err

		}else if err := ValidateLinks(fileLines[i]); err != nil {
			return AntFarmData{}, err
		}
	}

	if data.StartingRoom == "" || data.EndingRoom == "" {
		return AntFarmData{}, errors.New("Error: invalid data format, starting room or ending room not found")
	}

	return data, nil
}

//Links must have the following syntax : a-b
func ValidateLinks(link string) error {
	rooms := strings.Split(link, "-")
	if len(rooms) != 2 {
		return errors.New("Error: Invalid data format, invalid link format")
	}

	if rooms[0] == rooms[1] {
		return errors.New("Error: Invalid data format, cyclic link")
	}
	return nil
}

//Verify room name's format and its coordonate
func ValidateRoom(line string) error{
	data := strings.Split(line, " ")
	//room format : name cordx cordy
	if len(data) == 3 {
		//room name's never start with # or L
		if strings.HasPrefix(data[0], "#") || strings.HasPrefix(data[0], "L") {
			return errors.New("Error: invalid data format, invalid room name format")
		}

		_, err1 := strconv.Atoi(data[1])
		_, err2 := strconv.Atoi(data[2])
		if err1 != nil || err2 != nil {
			return errors.New("Error: invalid data format, invalid room coordinate")
		}
	}

	return nil
}


func FileContent(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil , errors.New("error on opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string

	for scanner.Scan(){
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines, nil
}