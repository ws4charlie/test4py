package coffeemachine

import (
	"math/rand/v2"
	"testing"
)

func grid(rows ...string) [][]bool {
	g := make([][]bool, len(rows))
	for i, r := range rows {
		for _, ch := range r {
			g[i] = append(g[i], ch == 'D')
		}
	}
	return g
}

func TestLocate(t *testing.T) {
	office := grid(
		"D...D",
		".....",
		"..D..",
	)
	r, c, ok := Locate(office)
	if !ok || r != 0 || c != 2 {
		t.Errorf("Locate = (%d, %d, %v), want (0, 2, true)", r, c, ok)
	}
	if got := Cost(office, r, c); got != 6 {
		t.Errorf("Cost = %d, want 6", got)
	}

	if _, _, ok := Locate(grid("...", "...")); ok {
		t.Errorf("Locate on an empty office: ok = true, want false")
	}
}

// TestLocateIsOptimal compares against trying every cell.
func TestLocateIsOptimal(t *testing.T) {
	rnd := rand.New(rand.NewPCG(7, 8))
	for n := 0; n < 200; n++ {
		R, C := 1+rnd.IntN(8), 1+rnd.IntN(8)
		office := make([][]bool, R)
		hasDev := false
		for r := range office {
			office[r] = make([]bool, C)
			for c := range office[r] {
				office[r][c] = rnd.IntN(4) == 0
				hasDev = hasDev || office[r][c]
			}
		}
		r, c, ok := Locate(office)
		if ok != hasDev {
			t.Fatalf("ok = %v, want %v", ok, hasDev)
		}
		if !ok {
			continue
		}
		best := Cost(office, 0, 0)
		for rr := 0; rr < R; rr++ {
			for cc := 0; cc < C; cc++ {
				best = min(best, Cost(office, rr, cc))
			}
		}
		if got := Cost(office, r, c); got != best {
			t.Fatalf("Locate cost %d, brute force best %d, office %v", got, best, office)
		}
	}
}
