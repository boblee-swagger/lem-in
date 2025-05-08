package main

import (
	"fmt"
	"lem-in/src"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	fileLines, err := src.FileContent(filename)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := src.ValidateAntfarmData(fileLines)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("number of ants: %d\n",data.NumberOfAnts)
	fmt.Printf("starting room: %s\n",data.StartingRoom.Name)
	fmt.Printf("ending room: %s\n",data.EndingRoom.Name)
	fmt.Printf("rooms: %v\n",data.Rooms)
	fmt.Printf("links: %v\n",data.Links)

	

}
