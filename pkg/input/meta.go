package input

import "github.com/Nekroze/tuireframe/pkg/ir"

type MetaInstruction struct {
	X    int     `json:"x"`
	Y    int     `json:"y"`
	Cell ir.Cell `json:"cell"`
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
	screen[m.X][m.Y].Apply(m.Cell)
}
