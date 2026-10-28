package hanoi

import (
	"slices"
	"testing"
)

func TestSolveSmall(t *testing.T) {
	want := []Move{{"a", "c"}, {"a", "b"}, {"c", "b"}, {"a", "c"}, {"b", "a"}, {"b", "c"}, {"a", "c"}}
	if got := Solve(3, "a", "b", "c"); !slices.Equal(got, want) {
		t.Errorf("Solve(3) = %v, want %v", got, want)
	}
}

// TestSolveValid replays the moves and checks the rules: never put a
// larger disk on a smaller one, and end with every disk on the target peg.
func TestSolveValid(t *testing.T) {
	for n := 0; n <= 10; n++ {
		pegs := map[string][]int{"a": nil, "b": nil, "c": nil}
		for d := n; d >= 1; d-- {
			pegs["a"] = append(pegs["a"], d)
		}
		moves := Solve(n, "a", "b", "c")
		if len(moves) != 1<<n-1 {
			t.Fatalf("n=%d: %d moves, want %d", n, len(moves), 1<<n-1)
		}
		for _, m := range moves {
			src := pegs[m.From]
			if len(src) == 0 {
				t.Fatalf("n=%d: move from empty peg %s", n, m.From)
			}
			disk := src[len(src)-1]
			if dst := pegs[m.To]; len(dst) > 0 && dst[len(dst)-1] < disk {
				t.Fatalf("n=%d: disk %d placed on smaller disk %d", n, disk, dst[len(dst)-1])
			}
			pegs[m.From] = src[:len(src)-1]
			pegs[m.To] = append(pegs[m.To], disk)
		}
		if len(pegs["c"]) != n {
			t.Fatalf("n=%d: %d disks on target peg", n, len(pegs["c"]))
		}
	}
}
