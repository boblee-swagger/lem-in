package src

import (
	"testing"
)

func TestFileContent(t *testing.T) {
	var tests = []struct {
		name string
		filename string
		wantErr  bool
	}{
		{"file0exists", "../examples/example00", false},
		{"file11notExists", "../examples/example011", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if _, err := FileContent(tt.filename); (err != nil) != tt.wantErr{
				t.Errorf("FileContent(%s) = %s", tt.filename, err)
			}
		})
	}
}

func TestValidateAntfarmData(t *testing.T) {
	var tests = []struct{
		name string
		filename string
		wantErr bool
	}{
		{"normal", "../examples/example01", false},
		{"confuse", "../examples/example03", false},
		{"weird", "../examples/example09", true},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			filelines, _ := FileContent(tt.filename)
			if _, err := ValidateAntfarmData(filelines); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAntfarmData(%s) = %s", tt.filename, err)
			}
		})
	}
}

func TestFormatLink(t *testing.T) {
	var tests = []struct{
		name string
		fromTo []string
		wantErr bool
	}{
		{"correct", []string{"hello", "world"}, false},
		{"cyclic link", []string{"no", "no"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if err := FormatLink(tt.fromTo); (err != nil) != tt.wantErr{
				t.Errorf("ForamtLink(%s) = %s", tt.fromTo, err)
			}
		})
	}
}

func TestValideRoomFormat(t *testing.T) {
	var tests = []struct{
		name string
		room []string
		wantErr bool
	}{
		{"correct", []string{"numORdigit", "2", "4"}, false},
		{"bad coordinate", []string{"mustBeInt", "t", "9"}, true},
		{"not start with L", []string{"Last", "3", "7"}, true},
	}
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			if _, err := FormatRoom(tt.room); (err != nil) != tt.wantErr{
				t.Errorf("FormatRoom(%s) = %s", tt.room, err)
			}	
		})
	}
}
