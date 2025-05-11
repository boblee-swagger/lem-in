package main

import (
	"lem-in/src"
	"log"
	"os"
	//"fmt"

)

func main() {
	Start()
}

func Start() {
	filename := os.Args[1]
	fileLines, err := src.FileContent(filename)
	if err != nil {
		log.Fatalln(err)
	}

	antFarm, err := src.ValidateAntfarmData(fileLines)
	if err != nil {
		log.Fatalln(err)
	}

	for i := range antFarm.GetLinks() {
		if i < len(antFarm.Rooms){
			antFarm.GetRooms()[i].SetNeighbors(antFarm.Links, antFarm.Rooms)
		}
		antFarm.GetLinks()[i].SetWeight(antFarm.Rooms)		
	}

}
