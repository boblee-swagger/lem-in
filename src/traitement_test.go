package src

import (
	//"fmt"
	"fmt"
	"testing"
)

func TestFileContent(t *testing.T) {
	filename := "../examples/example00"
	_, err := FileContent(filename)
	if err != nil {
		t.Errorf(`FileContent(filename) = %s`, err)
	}
	
}

func TestValidateAntfarmData(t *testing.T) {
	filename := "../examples/example01"
	fileLines, _ := FileContent(filename)
	_, err := ValidateAntfarmData(fileLines)
	if err != nil {
		t.Errorf(`ValidateAnteFarmData(fileLines) = %v %v`, fileLines, err)
	}
	fmt.Println("test passed")
}

func TestValideLinkFormat(t *testing.T) {
	link := []string{"hello", "world"}

	if _, err := FormatLink(link); err != nil {
		t.Errorf("FormatLink(line) = %v %v", link, err)
	}
	fmt.Println("test passed")
}

func TestValideRoomFormat(t *testing.T) {
	room := []string{"c", "1", "2"}
	if _, err := FormatRoom(room); err != nil {
		t.Errorf("FormatRoom(rooms[i]) = %v, %v", room, err)
	}
	fmt.Println("test passed")
}
