package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell_Apply(t *testing.T) {
	tests := []struct {
		name string
		in   *Cell
		args Cell
		want *Cell
	}{
		{
			name: "empty to nil is nil",
			in:   nil,
			args: Cell{0, "", ""},
			want: nil,
		},
		{
			name: "char to nil is nil, I wish it wasn't though",
			in:   nil,
			args: Cell{'a', "", ""},
			want: nil,
		},
		{
			name: "empty to empty is empty",
			in:   &Cell{0, "", ""},
			args: Cell{0, "", ""},
			want: &Cell{0, "", ""},
		},
		{
			name: "char to empty has char",
			in:   &Cell{0, "", ""},
			args: Cell{'a', "", ""},
			want: &Cell{'a', "", ""},
		},
		{
			name: "char to full has new char",
			in:   &Cell{0, "black", "white"},
			args: Cell{'b', "", ""},
			want: &Cell{'b', "black", "white"},
		},
		{
			name: "bg to empty has bg",
			in:   &Cell{0, "", ""},
			args: Cell{0, "black", ""},
			want: &Cell{0, "black", ""},
		},
		{
			name: "bg to full has new bg",
			in:   &Cell{'a', "black", "white"},
			args: Cell{0, "darkgreen", ""},
			want: &Cell{'a', "darkgreen", "white"},
		},
		{
			name: "fg to empty has fg",
			in:   &Cell{0, "", ""},
			args: Cell{0, "", "white"},
			want: &Cell{0, "", "white"},
		},
		{
			name: "fg to full has new fg",
			in:   &Cell{'a', "black", "white"},
			args: Cell{0, "", "red"},
			want: &Cell{'a', "black", "red"},
		},
		{
			name: "full to full has new full",
			in:   &Cell{'a', "black", "white"},
			args: Cell{'b', "darkgreen", "red"},
			want: &Cell{'b', "darkgreen", "red"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.Apply(tt.args)
			assert.Equal(t, tt.want, tt.in)
		})
	}
}
