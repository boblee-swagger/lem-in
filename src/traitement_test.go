package src

import (
	"fmt"
	"testing"
)

func TestFileContent(t *testing.T){
	filename := "../examples/example00"
	content, err := FileContent(filename);
	if err != nil {
		t.Errorf(`FileContent(filename) = %s`, err)
	}
	fmt.Println(content)
}


func TestValidateAntfarmData(t * testing.T){

}