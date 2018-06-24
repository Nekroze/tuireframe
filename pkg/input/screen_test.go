package input

import (
	"strings"
	"testing"

	"github.com/Nekroze/tuireframe/pkg/ir"
	"github.com/stretchr/testify/assert"
)

func Test_stringToScreenRow(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want ir.ScreenRow
	}{
		{
			name: "empty gets emtpy",
			arg:  "",
			want: map[int]*ir.Cell{},
		},
		{
			name: "single gets single",
			arg:  "a",
			want: map[int]*ir.Cell{
				0: {Character: 'a'},
			},
		},
		{
			name: "multiple gets multiple",
			arg:  "abcdef",
			want: map[int]*ir.Cell{
				0: {Character: 'a'},
				1: {Character: 'b'},
				2: {Character: 'c'},
				3: {Character: 'd'},
				4: {Character: 'e'},
				5: {Character: 'f'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, stringToScreenRow(tt.arg))
		})
	}
}

func TestReadScreen(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want ir.Screen
	}{
		{
			name: "empty gets empty",
			arg:  "",
			want: ir.Screen{},
		},
		{
			name: "single gets single row and column",
			arg:  "a",
			want: ir.Screen{
				0: ir.ScreenRow{
					0: {Character: 'a'},
				},
			},
		},
		{
			name: "multiple gets single row and multiple columns",
			arg:  "abc",
			want: ir.Screen{
				0: ir.ScreenRow{
					0: {Character: 'a'},
					1: {Character: 'b'},
					2: {Character: 'c'},
				},
			},
		},
		{
			name: "multiple lines get multiple rows and single column",
			arg:  "a\nb",
			want: ir.Screen{
				0: ir.ScreenRow{
					0: {Character: 'a'},
				},
				1: ir.ScreenRow{
					0: {Character: 'b'},
				},
			},
		},
		{
			name: "square gets multiple rows and multiple columns",
			arg:  "ab\ncd",
			want: ir.Screen{
				0: ir.ScreenRow{
					0: {Character: 'a'},
					1: {Character: 'b'},
				},
				1: ir.ScreenRow{
					0: {Character: 'c'},
					1: {Character: 'd'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.arg)
			assert.Equal(t, tt.want, ReadScreen(reader))
		})
	}
}
