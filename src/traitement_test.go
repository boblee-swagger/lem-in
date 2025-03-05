package src

import (
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

}

func TestValideLinkFormat(t *testing.T) {
	link :="hello-world";
	
	if _, err := ValideLinkFormat(link); err != nil {
		t.Errorf("ValidLinkFormat(line) = %v %v", link, err)
	}
}

func TestValideRoomFormat(t * testing.T) {
	room := "correct 1 1";
	if _, err := ValideRoomFormat(room); err != nil {
		t.Errorf("ValidRoomFormat(rooms[i]) = %v, %v", room, err)
	}
}

func TestValidLinkedRoom(t * testing.T) {
	rooms := []Room{{"a", 2, 5}, {"hello", 5 , -3}, {"b", 8, 9}}
	correctlink := "a-b";
	if err := ValideLinkedRoom(rooms, correctlink); err != nil {
		t.Errorf("Error: %v", err)
	}
}