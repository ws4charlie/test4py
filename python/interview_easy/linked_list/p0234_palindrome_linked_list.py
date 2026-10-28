# https://leetcode.com/problems/palindrome-linked-list/

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def isPalindrome(self, head):
        # O(1) space: find the middle, reverse the second half, compare
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        prev = None
        while slow:
            slow.next, prev, slow = prev, slow, slow.next
        while prev:
            if prev.val != head.val:
                return False
            prev, head = prev.next, head.next
        return True

def build(vals):
    head = None
    for v in reversed(vals):
        node = ListNode(v)
        node.next = head
        head = node
    return head

if __name__ == "__main__":
    sln = Solution()
    assert sln.isPalindrome(build([1, 2, 2, 1]))
    assert sln.isPalindrome(build([1, 2, 3, 2, 1]))
    assert not sln.isPalindrome(build([1, 2]))
    assert sln.isPalindrome(build([1]))
    assert sln.isPalindrome(build([]))
    print("ok")
