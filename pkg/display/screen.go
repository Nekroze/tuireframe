package display

import (
	"github.com/Nekroze/tuireframe/pkg/ir"
	"github.com/gdamore/tcell"
)

func drawScreen(faux ir.Screen, screen tcell.Screen) {
	for y, row := range faux {
		for x, cell := range row {
			drawCell(x, y, cell, screen)
		}
	}
}

func drawCell(x, y int, cell *ir.Cell, screen tcell.Screen) {
	st := tcell.StyleDefault
	st = st.Background(tcell.GetColor(cell.Background))
	st = st.Foreground(tcell.GetColor(cell.Foreground))
	screen.SetCell(x, y, st, cell.Character)
}
