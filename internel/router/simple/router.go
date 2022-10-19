package simple

import (
	"log"
	"strings"

	"github.com/longyue0521/httprouter/web"
)

type Router struct {
}

func New() *Router {
	return &Router{}
}

type Node struct {
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
func (r *Router) Add(method string, path string, handleFunc web.HandleFunc) {

	subpaths := r.checkPath(path)
	log.Println(subpaths)

}

func (r *Router) Get(method string, path string) (params PathParams, handleFunc web.HandleFunc, found bool) {

	subpaths := r.checkPath(path)
	log.Println(subpaths)

	return
}

func (r *Router) checkPath(path string) []string {

	subpaths := []string{"/"}

	if path == "/" {
		return subpaths
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

	for _, subpath := range strings.Split(path[1:], "/") {
		if subpath == "" {
			panic("web: 路由不能有重复'/'")
		}
		log.Println(subpath)
		subpaths = append(subpaths, subpath)
	}
	return subpaths
}
