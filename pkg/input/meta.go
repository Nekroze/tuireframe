package input

import (
	"github.com/Nekroze/tuireframe/pkg/ir"
)

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type MetaInstruction struct {
	Position Coordinates `json:"position"`
	Size     Dimensions  `json:"size"`
	Cell     ir.Cell     `json:"cell"`
}

type MetaFile struct {
	Description  string            `json:"description"`
	Instructions []MetaInstruction `json:"instructions"`
}

func (mf MetaFile) ApplyToScreen(screen ir.Screen) {
	for _, instruction := range mf.Instructions {
		instruction.applyToScreen(screen)
	}
}

func (m MetaInstruction) applyToScreen(screen ir.Screen) {
	for x := 0; x < m.Size.Width; x++ {
		for y := 0; y < m.Size.Height; y++ {
			screen[m.Position.Y+y][m.Position.X+x].Apply(m.Cell)
		}
	}
}
