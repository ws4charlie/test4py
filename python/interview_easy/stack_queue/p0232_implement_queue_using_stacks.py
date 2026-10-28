"""
232. Implement Queue using Stacks  (Easy)
https://leetcode.com/problems/implement-queue-using-stacks/

Build a first-in-first-out queue using only two stacks. The only stack
operations allowed are push to top, peek/pop from top, size and is-empty.

    MyQueue()    create an empty queue.
    push(x)      add x to the back of the queue.
    pop()        remove and return the front element.
    peek()       return the front element.
    empty()      return True if the queue is empty.

Example:
    Input:
        ["MyQueue","push","push","peek","pop","empty"]
        [[],[1],[2],[],[],[]]
    Output:
        [null,null,null,1,1,false]

Constraints:
    1 <= x <= 9
    At most 100 calls to push, pop, peek and empty.
    pop and peek are only called on a non-empty queue.

Follow-up: can every operation take amortized O(1) time?
"""

class MyQueue:
    # push onto `inbox`; pop from `outbox`, refilling it (reversed) only
    # when empty. Each element moves at most once: amortized O(1).
    def __init__(self):
        self.inbox = []
        self.outbox = []

    def push(self, x):
        self.inbox.append(x)

    def pop(self):
        self.peek()
        return self.outbox.pop()

    def peek(self):
        if not self.outbox:
            while self.inbox:
                self.outbox.append(self.inbox.pop())
        return self.outbox[-1]

    def empty(self):
        return not self.inbox and not self.outbox

if __name__ == "__main__":
    q = MyQueue()
    q.push(1)
    q.push(2)
    assert q.peek() == 1
    assert q.pop() == 1
    q.push(3)
    assert q.pop() == 2
    assert q.pop() == 3
    assert q.empty()
    print("ok")
