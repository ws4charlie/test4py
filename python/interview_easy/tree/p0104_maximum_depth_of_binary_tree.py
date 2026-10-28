# https://leetcode.com/problems/maximum-depth-of-binary-tree/

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
