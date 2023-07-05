package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	b := create_board()
	expected := [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	if b.tokens != expected {
		t.Errorf("Create Error")
	}
}

func TestGetToken01(t *testing.T) {
    b := create_board()
    b.tokens[0][1] = "o"
    if b.get(0, 1) != "o" {
        t.Errorf("TestGetToken01 Error")
    }
}
