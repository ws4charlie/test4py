package p0038

import "testing"

func TestCountAndSay(t *testing.T) {
	want := []string{"1", "11", "21", "1211", "111221", "312211", "13112221"}
	for i, w := range want {
		if got := CountAndSay(i + 1); got != w {
			t.Errorf("CountAndSay(%d) = %q, want %q", i+1, got, w)
		}
	}
}
