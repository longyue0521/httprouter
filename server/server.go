package server

import (
	"net/http"

	"github.com/longyue0521/httprouter/router"
	"github.com/longyue0521/httprouter/web"
)

type Server interface {
	http.Handler
	// Start 启动服务器
	// addr 是监听地址。如果只指定端口，可以使用 ":8081"
	// 或者 "localhost:8082"
	Start(addr string) error

	// AddRoute 注册一个路由
	// method 是 HTTP 方法
	AddRoute(method string, path string, handler web.HandleFunc)
	// 我们并不采取这种设计方案
	// addRoute(method string, path string, handlers... HandleFunc)
}

// 确保 HTTPServer 肯定实现了 Server 接口
var _ Server = &HTTPServer{}

type HTTPServer struct {
	router.Router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Router: router.New(),
	}
}

// ServeHTTP HTTPServer 处理请求的入口
func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &web.Context{
		Req:  request,
		Resp: writer,
	}
	s.serve(ctx)
}

// Start 启动服务器
func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *HTTPServer) Post(path string, handler web.HandleFunc) {
	s.AddRoute(http.MethodPost, path, handler)
}

func (s *HTTPServer) Get(path string, handler web.HandleFunc) {
	s.AddRoute(http.MethodGet, path, handler)
}

func (s *HTTPServer) serve(ctx *web.Context) {
	mi, ok := s.FindRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || mi.Node == nil || mi.Handler() == nil {
		ctx.Resp.WriteHeader(404)
		ctx.Resp.Write([]byte("Not Found"))
		return
	}
	ctx.PathParams = mi.PathParams
	mi.Handler()(ctx)
}
