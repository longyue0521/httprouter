package router

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/longyue0521/httprouter/internel/router/simple"
	"github.com/longyue0521/httprouter/web"
)

type Router struct {
	trees map[string]*simple.Tree
}

func New() *Router {
	return &Router{trees: make(map[string]*simple.Tree)}
}

type Node struct {
	path     string
	children map[string]*Node
	handler  web.HandleFunc
}

func (n *Node) child(path string) *Node {
	if n.children == nil {
		return nil
	}
	node, ok := n.children[path]
	if !ok {
		return nil
	}
	return node
}

type param struct {
	key string
	val string
}

type PathParams []param

func (p *PathParams) Get(key string) string {

	return ""
}

func (p *PathParams) Set(key string, val string) {

}

// Add 注册路由。
// method 是 HTTP 方法
// - 已经注册了的路由，无法被覆盖。例如 /user/home 注册两次，会冲突
// - path 必须以 / 开始并且结尾不能有 /，中间也不允许有连续的 /
// - 不能在同一个位置注册不同的参数路由，例如 /user/:id 和 /user/:name 冲突
// - 不能在同一个位置同时注册通配符路由和参数路由，例如 /user/:id 和 /user/* 冲突
// - 同名路径参数，在路由匹配的时候，值会被覆盖。例如 /user/:id/abc/:id，那么 /user/123/abc/456 最终 id = 456
func (r *Router) Add(method string, path string, handler web.HandleFunc) {

	subPaths := r.parse(path)

	tree, ok := r.trees[method]
	if !ok {
		tree = simple.NewTree("/", nil)
		r.trees[method] = tree
	}

	log.Println("root", tree.Root())

	node := tree.AddNode(subPaths[1:])

	if node.Handler() != nil {
		panic("web: 路由重复")
	}

	node.SetHandler(handler)

	log.Println(subPaths)
}

func (r *Router) Get(method string, path string) (params PathParams, handleFunc web.HandleFunc, err error) {

	defer func() {
		if r := recover(); r != nil {
			params, handleFunc, err = nil, nil, errors.New(fmt.Sprintf("%s", r))
		}
	}()

	subPaths := r.parse(path)

	tree, ok := r.trees[method]
	if !ok {
		panic("web: 非法路由")
	}

	node := tree.GetNode(subPaths[1:])

	if node == nil || node.Handler() == nil {
		panic("web: 非法路由")
	}

	return nil, node.Handler(), nil
}

func (r *Router) parse(path string) []string {

	subPaths := []string{"/"}

	if path == "/" {
		return subPaths
	}

	if path == "" {
		panic("web: 路由不能为空")
	}

	if path[0] != '/' {
		panic("web: 路由必须以'/'开头")
	}

	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 路由不能以'/'结尾")
	}

	for _, subPath := range strings.Split(path[1:], "/") {

		if subPath == "" {
			panic("web: 路由不能有重复'/'")
		}

		if strings.TrimSpace(subPath) != subPath {
			panic("web: 路由不能包含空白字符")
		}

		subPaths = append(subPaths, subPath)
	}

	return subPaths
}
