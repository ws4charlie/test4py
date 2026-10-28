// Package p0588 solves LeetCode 588. Design In-Memory File System (Hard, LeetCode Premium).
// https://leetcode.com/problems/design-in-memory-file-system/
//
// Design an in-memory file system with this API:
//
//	FileSystem()                    start with an empty file system.
//	ls(path)                        if path is a file, return [its name];
//	                                if it's a directory, return the names
//	                                inside it, sorted lexicographically.
//	mkdir(path)                     create the directory path, plus any
//	                                missing parent directories.
//	addContentToFile(path, content) create the file if it doesn't exist,
//	                                otherwise append content to it.
//	readContentFromFile(path)       return the file's content.
//
// Example:
//
//	Input:
//	    ["FileSystem","ls","mkdir","addContentToFile","ls","readContentFromFile"]
//	    [[],["/"],["/a/b/c"],["/a/b/c/d","hello"],["/"],["/a/b/c/d"]]
//	Output:
//	    [null,[],null,null,["a"],"hello"]
//
//	After the calls the tree looks like this:
//
//	    /
//	    `-- a/
//	        `-- b/
//	            `-- c/
//	                `-- d        file, content "hello"
//
// Constraints:
//
//	1 <= path.length, filePath.length <= 100
//	Paths are absolute: they start with "/" and don't end with "/",
//	except the root "/" itself.
//	Directory and file names are lowercase letters; names in the same
//	directory are unique.
//	readContentFromFile and ls are only called on paths that exist.
//	1 <= content.length <= 50
//	At most 300 calls in total.
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
