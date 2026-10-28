# https://leetcode.com/problems/linked-list-cycle/

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
