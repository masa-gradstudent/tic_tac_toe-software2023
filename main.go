package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	tokens [3][3]string
}

func create_board() Board {
	b := Board{}
	b.tokens = [3][3]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}
	return b
}

func (b *Board) put(x, y int, u string) bool {
	if b.tokens[x][y] != "." {
		return false
	}
	b.tokens[x][y] = u
	return true
}

func (b *Board) get(x, y int) string {
	return b.tokens[x][y]
}

func (b *Board) check_winner() string {
	for i := 0; i < 3; i++ {
		if b.get(i, 0) == b.get(i, 1) && b.get(i, 1) == b.get(i, 2) {
			return b.get(i, 0)
		}
		if b.get(0, i) == b.get(1, i) && b.get(1, i) == b.get(2, i) {
			return b.get(0, i)
		}
	}
	if b.get(0, 0) == b.get(1, 1) && b.get(1, 1) == b.get(2, 2) {
		return b.get(0, 0)
	}
	if b.get(0, 2) == b.get(1, 1) && b.get(1, 1) == b.get(2, 0) {
		return b.get(0, 2)
	}
	return "."
}

func (b *Board) print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", b.get(i, j))
		}
		fmt.Printf("\n")
	}
}

func convert_coordinate(input string) (int, int, bool) {
	inputs := strings.Split(input, ",")
	if len(inputs) != 2 {
		return 0, 0, false
	}
	x, _ := strconv.Atoi(inputs[0])
	y, _ := strconv.Atoi(inputs[1])
	return x, y, true
}

func main() {
	b := create_board()
	winner := ""
	current := "o"
	next := "x"
	i := 0
	for i < 9 {
		var input string
		fmt.Printf("Player %s: Input (x,y) ", current)
		fmt.Scan(&input)
		x, y, is_convert := convert_coordinate(input)
		if !is_convert {
			fmt.Printf("Invalid string(Ex. 1,2)\n")
			continue
		}
		is_put := b.put(x, y, current)
		if !is_put {
			fmt.Printf("Already put on board[%d][%d]\n", x, y)
			continue
		}
		b.print()
		winner = b.check_winner()
		if winner != "." {
			break
		}
		current, next = next, current
		i++
	}
	if winner == "." {
		fmt.Printf("Even\n")
	} else {
		fmt.Printf("Player %s won\n", winner)
	}
}
