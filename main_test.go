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

func TestGetToken02(t *testing.T) {
    b := create_board()
    b.tokens[1][2] = "x"
    if b.get(1, 2) != "x" {
        t.Errorf("TestGetToken02 Error")
    }
}

func TestPutToken01(t *testing.T) {
    b := create_board()
    b.put(0, 1, "o")
    if b.get(0, 1) != "o" {
        t.Errorf("TestPutToken01 Error")
    }
}
