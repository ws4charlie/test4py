r"""
543. Diameter of Binary Tree  (Easy)
https://leetcode.com/problems/diameter-of-binary-tree/

Return the diameter of a binary tree: the number of edges on the longest
path between any two nodes. The path may or may not pass through the
root.

Example 1:
    Input:  root = [1,2,3,4,5]
    Output: 3
    Explanation: 4 -> 2 -> 1 -> 3 (or 5 -> 2 -> 1 -> 3) has 3 edges.

              1
             / \
            /   \
           2     3
          / \
         4   5

Example 2:
    Input:  root = [1,2]
    Output: 1

Constraints:
    1 <= number of nodes <= 10^4
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
    def diameterOfBinaryTree(self, root):
        # the longest path through a node is depth(left) + depth(right);
        # compute depths bottom-up and keep the best sum seen
        best = 0
        def depth(node):
            nonlocal best
            if not node:
                return 0
            l, r = depth(node.left), depth(node.right)
            best = max(best, l + r)
            return 1 + max(l, r)
        depth(root)
        return best

if __name__ == "__main__":
    sln = Solution()
    assert sln.diameterOfBinaryTree(build([1, 2, 3, 4, 5])) == 3
    assert sln.diameterOfBinaryTree(build([1, 2])) == 1
    assert sln.diameterOfBinaryTree(build([1])) == 0
    # the longest path doesn't go through the root
    assert sln.diameterOfBinaryTree(build([1, 2, None, 3, 4, 5, None, None, 6])) == 4
    print("ok")
