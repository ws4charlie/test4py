r"""
104. Maximum Depth of Binary Tree  (Easy)
https://leetcode.com/problems/maximum-depth-of-binary-tree/

Return the maximum depth of a binary tree: the number of nodes on the
longest path from the root down to a leaf.

Example 1:
    Input:  root = [3,9,20,null,null,15,7]
    Output: 3

            3
           / \
          9   20
             /  \
            15   7

Example 2:
    Input:  root = [1,null,2]
    Output: 2

Constraints:
    0 <= number of nodes <= 10^4
    -100 <= Node.val <= 100
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
    def maxDepth(self, root):
        if not root:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))

    def maxDepthBFS(self, root):
        # count levels
        depth, level = 0, [root] if root else []
        while level:
            depth += 1
            level = [c for n in level for c in (n.left, n.right) if c]
        return depth

if __name__ == "__main__":
    sln = Solution()
    for f in (sln.maxDepth, sln.maxDepthBFS):
        assert f(build([3, 9, 20, None, None, 15, 7])) == 3
        assert f(build([1, None, 2])) == 2
        assert f(build([])) == 0
    print("ok")
