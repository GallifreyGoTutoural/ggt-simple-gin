package gsg

import "strings"

/*
	example trie tree：
	[/] has three node children: [/:lang],[/about],[/p]
	[/:lang] has two node children:  [/doc],[/tutorial],[/intro]
	[/about] has no node children
	[/p] has one node children: [/blog],[/related]
	lean more about this router trie tree: https://geektutu.com/post/gee-day3.html
*/

// router trie tree‘s node
type node struct {
	fullPath string  // waiting to match route, such as [/p/:lang]
	path     string  // part of fullPath, such as [:lang]
	children []*node // child nodes, such as [doc], [tutorial], [intro]
	isWild   bool    // whether fuzzy match or not, path contains [:] or [*] is true
}

// matchChild finds the first child node that matches the path, used in insert
func (n *node) matchChild(path string) *node {
	for _, child := range n.children {
		if child.path == path || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren finds all child nodes that match the path, used in search
func (n *node) matchChildren(path string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.path == path || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert inserts paths nodes into the trie tree
func (n *node) insert(fullPath string, paths []string, height int) {
	// recursion exit
	if len(paths) == height {
		n.fullPath = fullPath
		return
	}

	path := paths[height]
	child := n.matchChild(path)
	// if no child node matches, then create a new node
	if child == nil {
		child = &node{path: path, isWild: path[0] == ':' || path[0] == '*'}
		n.children = append(n.children, child)
	}
	// recursion
	child.insert(fullPath, paths, height+1)
}

// search searches the node that matches the paths
func (n *node) search(paths []string, height int) *node {
	// recursion exit
	if len(paths) == height || strings.HasPrefix(n.path, "*") {
		if n.fullPath == "" {
			return nil
		}
		return n
	}
	// recursion
	path := paths[height]
	children := n.matchChildren(path)
	for _, child := range children {
		result := child.search(paths, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
