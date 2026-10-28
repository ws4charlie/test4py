// https://leetcode.com/problems/implement-queue-using-stacks/
package p0232

// MyQueue pushes onto in and pops from out, refilling out (which reverses
// the order) only when it is empty. Each element moves at most once, so
// every operation is amortized O(1).
type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	v := q.Peek()
	q.out = q.out[:len(q.out)-1]
	return v
}

func (q *MyQueue) Peek() int {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}
