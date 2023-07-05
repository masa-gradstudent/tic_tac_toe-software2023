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
	return ""
}
