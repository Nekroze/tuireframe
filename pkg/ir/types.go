package ir

type Cell struct {
	Character  rune   `json:"char"`
	Background string `json:"bg"`
	Foreground string `json:"fg"`
}

type ScreenRow map[int]*Cell
type Screen map[int]ScreenRow

var emptyCell = Cell{}

func (c *Cell) Apply(in Cell) {
	if in.Character == emptyCell.Character {
		c.Character = in.Character
	}
}
