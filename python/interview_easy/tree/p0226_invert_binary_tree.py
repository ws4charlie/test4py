r"""
226. Invert Binary Tree  (Easy)
https://leetcode.com/problems/invert-binary-tree/

Mirror a binary tree (swap every node's left and right children) and
return its root.

Example 1:
    Input:  root = [4,2,7,1,3,6,9]
    Output: [4,7,2,9,6,3,1]

              4                     4
             / \                   / \
            /   \                 /   \
           2     7     -->       7     2
          / \   / \             / \   / \
         1   3 6   9           9   6 3   1

Example 2:
    Input:  root = [2,1,3]
    Output: [2,3,1]

          2                2
         / \     -->      / \
        1   3            3   1

Example 3:
    Input:  root = []
    Output: []

Constraints:
    0 <= number of nodes <= 100
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

def level_order(root):
    out, queue = [], [root] if root else []
    while queue:
        node = queue.pop(0)
        out.append(node.val)
        queue += [c for c in (node.left, node.right) if c]
    return out

class Solution:
    def invertTree(self, root):
        if root:
            root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root

if __name__ == "__main__":
    sln = Solution()
    assert level_order(sln.invertTree(build([4, 2, 7, 1, 3, 6, 9]))) == [4, 7, 2, 9, 6, 3, 1]
    assert level_order(sln.invertTree(build([2, 1, 3]))) == [2, 3, 1]
    assert sln.invertTree(build([])) is None
    print("ok")
