package main

type Board struct {
	tokens [3][3]string
}

func create_board() Board {
	b := Board{}
	b.tokens = [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	return b
}

func (b *Board) put(x, y int, u string) {
	b.tokens[x][y] = u
}

func (b *Board) get(x, y int) string {
	return b.tokens[x][y]
}

func (b *Board) check_winner() string {
	if b.get(1, 1) == "x" {
		return ""
	} else {
		return b.get(0, 0)
	}
}
