# https://leetcode.com/problems/add-two-numbers/

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head, tail = None, None
        
        carry = 0
        while l1 or l2 or carry > 0:
            sum = carry
            if l1:
                sum += l1.val
                l1 = l1.next
            if l2:
                sum += l2.val
                l2 = l2.next
            carry = sum // 10
            sum = sum % 10
            
            node = ListNode(sum)
            if not head:
                head, tail = node, node
            else:
                tail.next = node
                tail = node
                
        return head

l1 = ListNode(2)
l1.next = ListNode(4)
l1.next.next = ListNode(3)
l2 = ListNode(5)
l2.next = ListNode(6)
l2.next.next = ListNode(4)
                
if __name__ == '__main__':
    sln = Solution()
    res = sln.addTwoNumbers(l1, l2)
    while (res):
        print(res.val, end=' ')
        res = res.next