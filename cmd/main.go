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

	fmt.Println(data.EndingRoom)

}
