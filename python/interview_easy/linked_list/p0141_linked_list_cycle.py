"""
141. Linked List Cycle  (Easy)
https://leetcode.com/problems/linked-list-cycle/

Return True if the linked list has a cycle: some node can be reached
again by following next pointers. In the examples, pos is the index the
tail links back to (-1 means no cycle); it is not passed to the function.

Example 1:
    Input:  head = [3,2,0,-4], pos = 1
    Output: True

        3 -> 2 -> 0 -> -4
             ^          |
             |          |
             +----------+

Example 2:
    Input:  head = [1,2], pos = 0
    Output: True

        1 -> 2
        ^    |
        +----+

Example 3:
    Input:  head = [1], pos = -1
    Output: False

Constraints:
    0 <= number of nodes <= 10^4
    -10^5 <= Node.val <= 10^5
    pos is -1 or a valid index in the list.

Follow-up: can you do it with O(1) memory?
"""

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head):
        # Floyd: fast moves 2 steps, slow 1; they meet only inside a cycle
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if slow is fast:
                return True
        return False

def build(vals, pos):
    """Build a list; the tail links back to index pos (-1 = no cycle)."""
    nodes = [ListNode(v) for v in vals]
    for a, b in zip(nodes, nodes[1:]):
        a.next = b
    if nodes and pos >= 0:
        nodes[-1].next = nodes[pos]
    return nodes[0] if nodes else None

if __name__ == "__main__":
    sln = Solution()
    assert sln.hasCycle(build([3, 2, 0, -4], 1))
    assert sln.hasCycle(build([1, 2], 0))
    assert sln.hasCycle(build([1], 0))
    assert not sln.hasCycle(build([1], -1))
    assert not sln.hasCycle(build([], -1))
    print("ok")
