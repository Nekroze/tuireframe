package input

import (
	"bufio"
	"io"

	"github.com/Nekroze/tuireframe/pkg/ir"
)

func ReadScreen(screen io.Reader) ir.Screen {
	out := ir.Screen{}
	scanner := bufio.NewScanner(screen)
	row := 0
	for scanner.Scan() {
		out[row] = stringToScreenRow(scanner.Text())
		row++
	}
	return out
}

func stringToScreenRow(in string) ir.ScreenRow {
	out := ir.ScreenRow{}
	for col, char := range in {
		out[col] = &ir.Cell{
			Character: char,
		}
	}
	return out
}
