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
	if c == nil {
		c = &Cell{}
	}
	if in.Character != emptyCell.Character {
		c.Character = in.Character
	}
	if in.Background != emptyCell.Background {
		c.Background = in.Background
	}
	if in.Foreground != emptyCell.Foreground {
		c.Foreground = in.Foreground
	}
}
