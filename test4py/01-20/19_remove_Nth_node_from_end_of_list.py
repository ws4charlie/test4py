# https://leetcode.com/problems/remove-nth-node-from-end-of-list/

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        l, r, width = head, head, 1
        while r.next != None:
            if width < n+1:
                r = r.next
                width += 1
            else:
                l = l.next
                r = r.next
        
        if width == n: # remove head, because Nth is the head
            return head.next
        else: # remove Nth node
            l.next = l.next.next
        return head

if __name__ == "__main__":
    head = ListNode(1)
    tail = head
    tail.next = ListNode(2)
    tail = tail.next
    tail.next = ListNode(3)
    tail = tail.next
    tail.next = ListNode(4)
    tail = tail.next
    tail.next = ListNode(5)
    tail = tail.next

    sln = Solution()
    sln.removeNthFromEnd(head, 2)
    while head:
        print(head.val)
        head = head.next