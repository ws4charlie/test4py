package p0232

import "testing"

func TestMyQueue(t *testing.T) {
	q := Constructor()
	if !q.Empty() {
		t.Fatal("new queue is not empty")
	}
	q.Push(1)
	q.Push(2)
	if v := q.Peek(); v != 1 {
		t.Errorf("Peek() = %d, want 1", v)
	}
	if v := q.Pop(); v != 1 {
		t.Errorf("Pop() = %d, want 1", v)
	}
	q.Push(3) // goes to in while out still holds 2
	for _, want := range []int{2, 3} {
		if v := q.Pop(); v != want {
			t.Errorf("Pop() = %d, want %d", v, want)
		}
	}
	if !q.Empty() {
		t.Error("queue should be empty")
	}
}
