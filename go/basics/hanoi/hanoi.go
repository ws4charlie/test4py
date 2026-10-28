// Package hanoi solves the Tower of Hanoi, ported from
// python/basics/hanoi.py.
package hanoi

// Move is one disk moving from peg From to peg To.
type Move struct{ From, To string }

// Solve moves n disks from `from` to `to` using `via`: move n-1 disks out
// of the way to via, move the largest to `to`, then move the n-1 back on
// top of it. Always 2^n - 1 moves.
func Solve(n int, from, via, to string) []Move {
	var moves []Move
	var move func(n int, from, via, to string)
	move = func(n int, from, via, to string) {
		if n == 0 {
			return
		}
		move(n-1, from, to, via)
		moves = append(moves, Move{from, to})
		move(n-1, via, from, to)
	}
	move(n, from, via, to)
	return moves
}
