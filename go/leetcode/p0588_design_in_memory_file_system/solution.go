// https://leetcode.com/problems/design-in-memory-file-system/
package p0588

import (
	"slices"
	"strings"
)

// node is a trie node: a directory has children, a file has content.
type node struct {
	children map[string]*node
	isFile   bool
	content  strings.Builder
}

func newNode() *node { return &node{children: make(map[string]*node)} }

type FileSystem struct {
	root *node
}

func Constructor() FileSystem {
	return FileSystem{root: newNode()}
}

// walk follows path from the root, creating missing nodes along the way.
func (fs *FileSystem) walk(path string) *node {
	cur := fs.root
	for _, name := range strings.Split(path, "/") {
		if name == "" {
			continue
		}
		next, ok := cur.children[name]
		if !ok {
			next = newNode()
			cur.children[name] = next
		}
		cur = next
	}
	return cur
}

// Ls returns [name] for a file, or the sorted entries of a directory.
func (fs *FileSystem) Ls(path string) []string {
	n := fs.walk(path)
	if n.isFile {
		return []string{path[strings.LastIndex(path, "/")+1:]}
	}
	names := make([]string, 0, len(n.children))
	for name := range n.children {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

func (fs *FileSystem) Mkdir(path string) {
	fs.walk(path)
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	n := fs.walk(filePath)
	n.isFile = true
	n.content.WriteString(content)
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	return fs.walk(filePath).content.String()
}
