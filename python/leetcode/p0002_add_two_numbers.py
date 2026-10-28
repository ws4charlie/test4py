"""
2. Add Two Numbers  (Medium)
https://leetcode.com/problems/add-two-numbers/

Two non-empty linked lists each hold a non-negative integer, one digit per
node, with the digits in reverse order (ones digit first). Return the sum
of the two numbers as a linked list in the same reverse order. Neither
number has leading zeros, except the number 0 itself.

Example 1:

      2 -> 4 -> 3          342
    + 5 -> 6 -> 4        + 465
    -------------        -----
      7 -> 0 -> 8          807

    Input:  l1 = [2,4,3], l2 = [5,6,4]
    Output: [7,0,8]
    Explanation: 342 + 465 = 807

Example 2:
    Input:  l1 = [0], l2 = [0]
    Output: [0]

Example 3:
    Input:  l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
    Output: [8,9,9,9,0,0,0,1]

Constraints:
    Each list has 1 to 100 nodes.
    0 <= Node.val <= 9
    The numbers have no leading zeros, except the number 0 itself.
"""

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