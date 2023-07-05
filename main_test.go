package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	b := create_board()
	expected := [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
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

func TestPutToken02(t *testing.T) {
    b := create_board()
    b.put(1, 0, "x")
    if b.get(1, 0) != "x" {
        t.Errorf("TestPutToken02 Error")
    }
}

func TestCheckWinner01(t *testing.T) {
    b := create_board()
    b.put(0, 0, "x")
    b.put(0, 1, "x")
    b.put(0, 2, "x")
    if b.check_winner() != "x" {
        t.Errorf("TestCheckWinner01 Error")
    }
}

func TestCheckWinner02(t *testing.T) {
    b := create_board()
    b.put(0, 0, "o")
    b.put(1, 1, "o")
    b.put(2, 2, "o")
    if b.check_winner() != "o" {
        t.Errorf("TestCheckWinner02 Error")
    }
}

func TestCheckWinner03(t *testing.T) {
    b := create_board()
    b.put(0, 0, "o")
    b.put(1, 1, "x")
    b.put(2, 2, "o")
    if b.check_winner() != " " {
        t.Errorf("TestCheckWinner03 Error")
    }
}