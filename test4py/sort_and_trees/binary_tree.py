# https://zh.wikipedia.org/wiki/%E6%A0%91%E7%9A%84%E9%81%8D%E5%8E%86

#---------------------------------------------------------------
# Define plain node data
#---------------------------------------------------------------
class Node:
    def __init__(self, v, lch=None, rch=None):
        self.v = v
        self.lch = lch
        self.rch = rch

#----------------------------------------------------------------
# A flat list of Nodes as input parameter of func 'create_tree'
#----------------------------------------------------------------
tree_nodes = [Node('F', 1, 2),
              Node('B', 3, 4),
              Node('G', None, 5),
              Node('A', None, None),
              Node('D', 6, 7),
              Node('I', 8, None),
              Node('C', None, None),
              Node('E', None, None),
              Node('H', None, None)]

# function to visit node
def vst(node):
    if node:
        print(node.v, end='', flush=True)

#                             A typical Binary Tree
#
#                                       F
#                                   /      \
#                                  B        G
#                               /     \      \
#                              A        D     I
#                                     /   \   /
#                                    C    E   H

#---------------------------------------------------------------
# Binary tree node definition
#---------------------------------------------------------------
class TreeNode:
    def __init__(self, v):
        self.v = v
        self.lch = None
        self.rch = None

def create_tree(nodes: 'list[Node]', root: 'root index'=0) -> 'TreeNode':
    if root == None:
        return None
    
    node = nodes[root]
    head = TreeNode(node.v)
    head.lch = create_tree(nodes, node.lch)
    head.rch = create_tree(nodes, node.rch)
    return head

#---------------------------------------------------------------
# 先序遍历-递归 Pre-order traverse
#---------------------------------------------------------------
def pre_traverse(root):
    if root == None:
        return
    vst(root)
    pre_traverse(root.lch)
    pre_traverse(root.rch)

#---------------------------------------------------------------
# 先序遍历-循环 Pre-order traverse
#---------------------------------------------------------------
def pre_traverse_loop(root):
    if root == None:
        return
    
    s = [root] # use a stack
    while s:
        node = s.pop()
        vst(node)
        if node.rch:
            s.append(node.rch)
        if node.lch:
            s.append(node.lch)

#---------------------------------------------------------------
# 中序遍历-递归 In-order traverse
#---------------------------------------------------------------
def in_traverse(root):
    if root == None:
        return
    in_traverse(root.lch)
    vst(root)
    in_traverse(root.rch)

#---------------------------------------------------------------
# 中序遍历-循环 In-order traverse
#---------------------------------------------------------------
def in_traverse_loop(root):
    if root == None:
        return
    
    s = [root] # use a stack
    while s:
        node = s[-1]
        if node.lch:# append left child all the way reaching left-most leaf
            s.append(node.lch)
        else:# now this node has not left child
            while s:
                node = s.pop()
                vst(node)
                if node.rch:
                    s.append(node.rch)
                    break

#---------------------------------------------------------------
# 后序遍历-递归 Post-order traverse
#---------------------------------------------------------------
def post_traverse(root):
    if root == None:
        return
    post_traverse(root.lch)
    post_traverse(root.rch)
    vst(root)

#---------------------------------------------------------------
# 后序遍历-循环 Post-order traverse
#---------------------------------------------------------------
def post_traverse_loop(root):
    if root == None:
        return
    
    s = [root]     # use a stack
    while s:
        node = s[-1]
        if node.rch:# append its right child, if available
            s.append(node.rch)
        if node.lch:# append its left child, if available
            s.append(node.lch)
        if node.rch == None and node.lch == None:# now this node without child
            node = s.pop()
            vst(node)
            while s:  # pop and visit parent if done with both left and right child
                top = s[-1]
                if top.lch == node or top.rch == node:
                    node = s.pop()
                    vst(node)
                else: # break because top is the node's brother
                    break

#---------------------------------------------------------------
# 后序遍历-循环-简单版(额外空间O(N)存放result) Post-order traverse
#---------------------------------------------------------------
def post_traverse_loop_simple(root):
        if root == None:
            return
        
        s = [root]
        res = []
        while s:
            node = s.pop()
            res.insert(0, node.v)
            if node.lch:
                s.append(node.lch)
            if node.rch:
                s.append(node.rch)
        
        for v in res:
            print(v, end='', flush=True)

#---------------------------------------------------------------
# 广度优先遍历-按层次遍历 Layer traverse
#---------------------------------------------------------------
def layer_traverse(root):
    if root == None:
        return

    s = [root]
    while s:
        node = s.pop(0)
        vst(node)
        if node.lch:
            s.append(node.lch)
        if node.rch:
            s.append(node.rch)

#---------------------------------------------------------------
# Test
#---------------------------------------------------------------
if __name__ == '__main__':
    my_tree = create_tree(tree_nodes, 0)

    pre_traverse(my_tree)
    print('')
    pre_traverse_loop(my_tree)
    print('')

    in_traverse(my_tree)
    print('')
    in_traverse_loop(my_tree)
    print('')

    post_traverse(my_tree)
    print('')
    post_traverse_loop(my_tree)
    print('')
    post_traverse_loop_simple(my_tree)
    print('')

    layer_traverse(my_tree)
    print('')