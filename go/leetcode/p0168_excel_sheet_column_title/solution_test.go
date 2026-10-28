package p0168

import "testing"

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "A"},
		{26, "Z"},
		{27, "AA"},
		{28, "AB"},
		{52, "AZ"},
		{701, "ZY"},
		{702, "ZZ"},
		{703, "AAA"},
		{2147483647, "FXSHRXW"},
	}
	for _, tt := range tests {
		if got := ConvertToTitle(tt.n); got != tt.want {
			t.Errorf("ConvertToTitle(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
