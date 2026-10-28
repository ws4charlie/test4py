package p0070

import "testing"

func TestClimbStairs(t *testing.T) {
	want := []int{1, 2, 3, 5, 8, 13, 21}
	for i, w := range want {
		if got := ClimbStairs(i + 1); got != w {
			t.Errorf("ClimbStairs(%d) = %d, want %d", i+1, got, w)
		}
	}
	if got := ClimbStairs(45); got != 1836311903 {
		t.Errorf("ClimbStairs(45) = %d, want 1836311903", got)
	}
}
