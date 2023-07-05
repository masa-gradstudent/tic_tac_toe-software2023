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
