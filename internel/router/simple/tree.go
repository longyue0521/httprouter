package simple

import (
	"github.com/longyue0521/httprouter/web"
)

type Node struct {
	path     string
	handler  web.HandleFunc
	children map[string]*Node
}

func (n *Node) Handler() web.HandleFunc {
	return n.handler
}

func (n *Node) SetHandler(handler web.HandleFunc) {
	n.handler = handler
}

type Tree struct {
	root *Node
}

func NewTree(path string, handler web.HandleFunc) *Tree {
	return &Tree{root: &Node{path: path, handler: handler, children: make(map[string]*Node)}}
}

func (t *Tree) Root() *Node {
	return t.root
}

// AddNode 通过路径切片添加结点
func (t *Tree) AddNode(paths []string) *Node {
	createFunc := func(path string) *Node {
		// todo: 需要传入path变量，留着它是因为没有测试用例覆盖，节点上的路径标记到底有啥用
		return &Node{path: "", children: make(map[string]*Node)}
	}
	return t.walk(paths, createFunc)
}

// GetNode 通过路径切片获取结点
// paths 中不包含任何"/"
func (t *Tree) GetNode(paths []string) *Node {
	findFunc := func(path string) *Node { return nil }
	return t.walk(paths, findFunc)
}

func (t *Tree) walk(paths []string, fn func(path string) *Node) *Node {
	root := t.root
	for _, path := range paths {
		child, ok := root.children[path]
		if !ok {
			node := fn(path)
			if node == nil {
				return nil
			}
			root.children[path] = node
			root = node
			continue
		}
		root = child
	}
	return root
}
