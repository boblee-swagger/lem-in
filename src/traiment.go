package src

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Room struct {
	Name  string
	CordX int
	CordY int
}

type AntFarm struct {
	NumberOfAnts int
	StartingRoom string
	EndingRoom   string
	Rooms        []Room
	Links        []string
}

// parses a list of file lines to extract useful data(n of ants...)
func ValidateAntfarmData(fileLines []string) (AntFarm, error) {
	var data AntFarm
	var err error
	number_of_ants, err := strconv.Atoi(fileLines[0])
	if err != nil {
		return AntFarm{}, errors.New("error: Invalid data format, invalid number of ants")
	}
	data.NumberOfAnts = number_of_ants
	for i := 1; i < len(fileLines); i++ {

		line := strings.Trim(fileLines[i]," ");
		if line == "##start" {
			if i+1 < len(fileLines) {
				data.StartingRoom = fileLines[i+1]
				i++
			}
		} else if line == "##end" {
			if i+1 < len(fileLines) {
				data.EndingRoom = fileLines[i+1]
				i++
			}
		}else{
			err = ValideRoom(line);
			if err != nil {
				return AntFarm{}, err
			}else{
				room := SetRoom(line)
				data.Rooms = append(data.Rooms, room)
			}
			err = ValideLink(line); 
			if err != nil {
				return AntFarm{}, err
			}else{
				//collect the link
			}
		}
	}
	if data.StartingRoom == "" || data.EndingRoom == "" {
		return AntFarm{}, errors.New("error: invalid data format, starting room or ending room not found")
	}
	return data, nil
}

// Links must have the following syntax : a-b
func ValideLink(link string) error {
	test, err := regexp.MatchString(`^[A-Za-z0-9]+-[A-Za-z0-9]+$`, link)
	if err != nil {
		return errors.New("Error: incorrect pattern for links")
	}

	if test {
		lns := strings.Split(link, "-")
		//the room is link to itself
		if lns[0] == lns[1] {
			return errors.New("Error: invalid link")
		}
	}
	
	return nil
}

// Verify room name's format and its coordonate
func ValideRoom(line string) error {
	//i must typed like this : a b b/c , spaces matter
	test, err := regexp.MatchString(`^*([A-Za-z0-9]+ +[0-9]+ +[0-9]+)$`, line)
	if err != nil {
		return errors.New("Error: Invalid room")
	}
	//room never started with # or L
	
	if test && (strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L")){
		return errors.New("Error: Invalid room")
	}

	return nil
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
