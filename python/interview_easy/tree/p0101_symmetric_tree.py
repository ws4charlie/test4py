r"""
101. Symmetric Tree  (Easy)
https://leetcode.com/problems/symmetric-tree/

Return True if a binary tree is a mirror image of itself around its
center.

Example 1:
    Input:  root = [1,2,2,3,4,4,3]
    Output: True

              1
             / \
            /   \
           2     2
          / \   / \
         3   4 4   3

Example 2:
    Input:  root = [1,2,2,null,3,null,3]
    Output: False

            1
           / \
          2   2
           \   \
            3   3

Constraints:
    1 <= number of nodes <= 1000
    -100 <= Node.val <= 100

Follow-up: can you solve it both recursively and iteratively?
"""

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def build(vals):
    """Build a tree from LeetCode level order, None for a missing node."""
    if not vals or vals[0] is None:
        return None
    root = TreeNode(vals[0])
    queue, i = [root], 1
    while queue and i < len(vals):
        node = queue.pop(0)
        for side in ("left", "right"):
            if i < len(vals) and vals[i] is not None:
                child = TreeNode(vals[i])
                setattr(node, side, child)
                queue.append(child)
            i += 1
    return root

class Solution:
    def isSymmetric(self, root):
        # two subtrees mirror each other if their roots match and
        # outer pair and inner pair of children mirror each other
        def mirror(a, b):
            if not a or not b:
                return a is b
            return a.val == b.val and mirror(a.left, b.right) and mirror(a.right, b.left)
        return mirror(root.left, root.right) if root else True

if __name__ == "__main__":
    sln = Solution()
    assert sln.isSymmetric(build([1, 2, 2, 3, 4, 4, 3]))
    assert not sln.isSymmetric(build([1, 2, 2, None, 3, None, 3]))
    assert sln.isSymmetric(build([1]))
    assert not sln.isSymmetric(build([1, 2, 3]))
    print("ok")
