
# Graph (undirected)
# https://blog.csdn.net/weixin_40314737/article/details/80893507
# https://blog.csdn.net/changyuanchn/article/details/79008760

class Graph:
    def __init__(self, nodes=[], sides=[]):
        self.nodes = {}
        self.visited = set()
        
        for node in nodes:
            n_sides = []
            for side in sides:
                u, v = side
                if node == u:
                    n_sides.append(v)
                elif node == v:
                    n_sides.append(u)
            self.nodes[node] = n_sides

    def add_nodes(self, nodes):
        for node in nodes:
            self.add_node(node)

    def add_node(self, node):
        if not node in self.nodes:
            self.nodes[node] = []

    def add_edge(self, edge):
        u, v = edge
        if(v not in self.nodes[u]) and ( u not in self.nodes[v]):
            self.nodes[u].append(v)

            if u != v:
                self.nodes[v].append(u)

    def nodes(self):
        return self.nodes.keys()

    def _visit(self, node):
        print(node)

    def dfs(self, root=None):
        if not root:
            return

        #visit root
        self._visit(root)
        self.visited.add(root)

        #dfs each neighbors
        for nb in self.nodes[root]:
            if not nb in self.visited:
                self.dfs(nb)

        #对于不连通的结点, 即dfs(root)完仍是没有visit过的单独处理，再做一次dfs
        for nd in self.nodes:
            if not nd in self.visited:
                self.dfs(nd)

    def bfs(self, root=None):
        if not root:
            return

        q = [root]
        while q:#和二叉树的层次遍历一样
            nd = q.pop(0)
            self._visit(nd)
            self.visited.add(nd)

            for nb in self.nodes[nd]:
                if nb not in self.visited and not nb in q:
                    q.append(nb)

        for nd in self.nodes:
            if nd not in self.visited:
                self.bfs(nd)

#  1---2---4---8
#  |   |      /
#  |   |    /
#  3   5  /
#  | \
#  |  \
#  6---7
if __name__ == '__main__':
    g = Graph()
    g.add_nodes([i+1 for i in range(8)])
    g.add_edge((1, 2))
    g.add_edge((1, 3))
    g.add_edge((2, 4))
    g.add_edge((2, 5))
    g.add_edge((4, 8))
    g.add_edge((5, 8))
    g.add_edge((3, 6))
    g.add_edge((3, 7))
    g.add_edge((6, 7))
    print("nodes:", g.nodes)
    
    g.visited = set()
    g.bfs(1)

    g.visited = set()
    g.dfs(1)