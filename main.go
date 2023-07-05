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
}

func (b *Board) get(x, y int) string {
	return ""
}
