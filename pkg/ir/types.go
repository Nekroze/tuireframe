package ir

type Cell struct {
	Character  rune   `json:"char"`
	Background string `json:"bg"`
	Foreground string `json:"fg"`
}

type ScreenRow map[int]*Cell

// A screen is a two dimensional map of *Cell's where the
// first index is the y coordinate and the second the x.
type Screen map[int]ScreenRow

var emptyCell = Cell{}

func (c *Cell) Apply(in Cell) {
	if in.Character == emptyCell.Character {
		c.Character = in.Character
	}
}
