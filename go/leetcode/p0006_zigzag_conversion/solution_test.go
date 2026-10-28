package p0006

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		want    string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"AB", 1, "AB"},
		{"AB", 3, "AB"},
	}
	for _, tt := range tests {
		if got := Convert(tt.s, tt.numRows); got != tt.want {
			t.Errorf("Convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
		}
	}
}
