# https://leetcode.com/problems/merge-two-sorted-lists/

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        head, tail = None, None
        
        while l1 or l2:
            min = None;
            if l1 is None:
                min = l2
                l2 = l2.next
            elif l2 is None:
                min = l1
                l1 = l1.next
            else:
                if l1.val < l2.val:
                    min = l1
                    l1 = l1.next
                else:
                    min = l2
                    l2 = l2.next
            if not head:
                head, tail = min, min
            else:
                tail.next = min
                tail = min
                
        return head

if __name__ == "__main__":
    sln = Solution()
    l1 = ListNode(1)
    l1.next = ListNode(2)
    l1.next.next = ListNode(4)
    l2 = ListNode(1)
    l2.next = ListNode(3)
    l2.next.next = ListNode(4)
    l = sln.mergeTwoLists(l1, l2)
    while l:
        print(l.val)
        l = l.next
