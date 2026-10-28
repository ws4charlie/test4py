package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	impls := map[string]func(int) int{
		"Recursive": Recursive,
		"Memo":      Memo,
		"Iterative": Iterative,
	}
	for name, f := range impls {
		for n, w := range want {
			if got := f(n); got != w {
				t.Errorf("%s(%d) = %d, want %d", name, n, got, w)
			}
		}
	}
	if got := Iterative(90); got != Memo(90) || got != 2880067194370816120 {
		t.Errorf("F(90): Iterative=%d Memo=%d", got, Memo(90))
	}
}
