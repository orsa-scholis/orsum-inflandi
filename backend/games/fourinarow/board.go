package fourinarow

type board struct {
	rows [5][7]int
	game *game
}

func initBoard() *board {
	return &board{
		rows: [5][7]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}
}

func (b *board) rowFull(i int) bool {
	return false
}
