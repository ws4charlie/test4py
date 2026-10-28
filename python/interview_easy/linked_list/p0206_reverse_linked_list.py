# https://leetcode.com/problems/reverse-linked-list/

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseList(self, head):
        # iterative: point each node back at the previous one
        prev = None
        while head:
            head.next, prev, head = prev, head, head.next
        return prev

    def reverseListRecursive(self, head):
        if not head or not head.next:
            return head
        new_head = self.reverseListRecursive(head.next)
        head.next.next = head
        head.next = None
        return new_head

def build(vals):
    head = None
    for v in reversed(vals):
        node = ListNode(v)
        node.next = head
        head = node
    return head

def to_list(head):
    out = []
    while head:
        out.append(head.val)
        head = head.next
    return out

if __name__ == "__main__":
    sln = Solution()
    for f in (sln.reverseList, sln.reverseListRecursive):
        assert to_list(f(build([1, 2, 3, 4, 5]))) == [5, 4, 3, 2, 1]
        assert to_list(f(build([1, 2]))) == [2, 1]
        assert to_list(f(build([]))) == []
    print("ok")
