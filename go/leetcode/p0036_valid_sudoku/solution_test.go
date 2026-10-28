package p0036

import "testing"

func parse(rows []string) [][]byte {
	board := make([][]byte, len(rows))
	for i, r := range rows {
		board[i] = []byte(r)
	}
	return board
}

var valid = []string{
	"53..7....",
	"6..195...",
	".98....6.",
	"8...6...3",
	"4..8.3..1",
	"7...2...6",
	".6....28.",
	"...419..5",
	"....8..79",
}

func TestIsValidSudoku(t *testing.T) {
	sameBox := append([]string{}, valid...)
	sameBox[0] = "83..7...." // 8 already in the top-left box (and column 0)

	sameRow := append([]string{}, valid...)
	sameRow[0] = "53..7...5"

	tests := []struct {
		name  string
		board []string
		want  bool
	}{
		{"valid", valid, true},
		{"duplicate in box and column", sameBox, false},
		{"duplicate in row", sameRow, false},
	}
	for _, tt := range tests {
		if got := IsValidSudoku(parse(tt.board)); got != tt.want {
			t.Errorf("%s: IsValidSudoku = %v, want %v", tt.name, got, tt.want)
		}
	}
}
