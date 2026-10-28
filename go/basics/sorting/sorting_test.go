package sorting

import (
	"math/rand/v2"
	"slices"
	"testing"
)

var sorts = map[string]func([]int){
	"Bubble":         Bubble[int],
	"Insertion":      Insertion[int],
	"Selection":      Selection[int],
	"Merge":          Merge[int],
	"MergeBottomUp":  MergeBottomUp[int],
	"Heap":           Heap[int],
	"Quick":          Quick[int],
	"QuickIterative": QuickIterative[int],
}

func inputs() [][]int {
	in := [][]int{
		nil,
		{1},
		{2, 1},
		{2, 7, 9, 4, 7, 5, 1},
		{5, 5, 5, 5},
	}
	desc := make([]int, 100)
	for i := range desc {
		desc[i] = 100 - i
	}
	in = append(in, desc)
	r := rand.New(rand.NewPCG(1, 2))
	for n := 0; n < 50; n++ {
		a := make([]int, r.IntN(200))
		for i := range a {
			a[i] = r.IntN(50) - 25 // small range forces duplicates
		}
		in = append(in, a)
	}
	return in
}

func TestSorts(t *testing.T) {
	for name, sort := range sorts {
		for _, in := range inputs() {
			want := slices.Clone(in)
			slices.Sort(want)
			got := slices.Clone(in)
			sort(got)
			if !slices.Equal(got, want) {
				t.Fatalf("%s(%v) = %v, want %v", name, in, got, want)
			}
		}
	}
}

func TestShuffle(t *testing.T) {
	a := make([]int, 54) // a deck of cards
	for i := range a {
		a[i] = i
	}
	b := slices.Clone(a)
	Shuffle(b)
	slices.Sort(b)
	if !slices.Equal(a, b) {
		t.Errorf("Shuffle lost or duplicated elements")
	}
}
