// Package p0232 solves LeetCode 232. Implement Queue using Stacks (Easy).
// https://leetcode.com/problems/implement-queue-using-stacks/
//
// Build a first-in-first-out queue using only two stacks. The only stack
// operations allowed are push to top, peek/pop from top, size and is-empty.
//
//	MyQueue()    create an empty queue.
//	push(x)      add x to the back of the queue.
//	pop()        remove and return the front element.
//	peek()       return the front element.
//	empty()      return true if the queue is empty.
//
// Example:
//
//	Input:
//	    ["MyQueue","push","push","peek","pop","empty"]
//	    [[],[1],[2],[],[],[]]
//	Output:
//	    [null,null,null,1,1,false]
//
// Constraints:
//
//	1 <= x <= 9
//	At most 100 calls to push, pop, peek and empty.
//	pop and peek are only called on a non-empty queue.
//
// Follow-up: can every operation take amortized O(1) time?
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
