package p0012

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"},
	}
	for _, tt := range tests {
		if got := IntToRoman(tt.num); got != tt.want {
			t.Errorf("IntToRoman(%d) = %q, want %q", tt.num, got, tt.want)
		}
	}
}
