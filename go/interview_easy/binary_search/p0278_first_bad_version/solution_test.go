package p0278

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct{ n, bad int }{
		{5, 4},
		{1, 1},
		{10, 1},
		{10, 10},
		{2126753390, 1702766719},
	}
	for _, tt := range tests {
		calls := 0
		isBadVersion = func(v int) bool { calls++; return v >= tt.bad }
		if got := FirstBadVersion(tt.n); got != tt.bad {
			t.Errorf("FirstBadVersion(%d) = %d, want %d", tt.n, got, tt.bad)
		}
		if calls > 32 {
			t.Errorf("n=%d: %d calls to isBadVersion, want O(log n)", tt.n, calls)
		}
	}
}
