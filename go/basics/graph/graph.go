// Package graph is an undirected graph stored as adjacency lists, with
// DFS, BFS and connected components; ported from python/basics/graph.py.
// Nodes and neighbors keep insertion order, so traversals are
// deterministic.
package graph

type Graph struct {
	adj   map[int][]int
	order []int // nodes in insertion order
}

func New() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddNode(v int) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = nil
		g.order = append(g.order, v)
	}
}

// AddEdge adds an undirected edge u-v, adding missing nodes. Duplicate
// edges are ignored.
func (g *Graph) AddEdge(u, v int) {
	g.AddNode(u)
	g.AddNode(v)
	for _, w := range g.adj[u] {
		if w == v {
			return
		}
	}
	g.adj[u] = append(g.adj[u], v)
	if u != v {
		g.adj[v] = append(g.adj[v], u)
	}
}

// Nodes returns all nodes in insertion order.
func (g *Graph) Nodes() []int { return append([]int(nil), g.order...) }

// DFS returns the nodes reachable from start, depth first.
func (g *Graph) DFS(start int) []int {
	return g.dfs(start, make(map[int]bool))
}

func (g *Graph) dfs(start int, seen map[int]bool) []int {
	var res []int
	var visit func(v int)
	visit = func(v int) {
		seen[v] = true
		res = append(res, v)
		for _, w := range g.adj[v] {
			if !seen[w] {
				visit(w)
			}
		}
	}
	visit(start)
	return res
}

// BFS returns the nodes reachable from start, nearest first. A node is
// marked seen when it is queued, so it is never queued twice.
func (g *Graph) BFS(start int) []int {
	seen := map[int]bool{start: true}
	queue := []int{start}
	var res []int
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		res = append(res, v)
		for _, w := range g.adj[v] {
			if !seen[w] {
				seen[w] = true
				queue = append(queue, w)
			}
		}
	}
	return res
}

// Components runs DFS from every node not seen yet; each run is one
// connected component.
func (g *Graph) Components() [][]int {
	seen := make(map[int]bool)
	var res [][]int
	for _, v := range g.order {
		if !seen[v] {
			res = append(res, g.dfs(v, seen))
		}
	}
	return res
}
