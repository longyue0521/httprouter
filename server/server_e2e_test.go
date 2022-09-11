package server

import (
	"testing"

	"github.com/longyue0521/httprouter/web"
)

// 这里放着端到端测试的代码

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	s.Get("/", func(ctx *web.Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *web.Context) {
		ctx.Resp.Write([]byte("hello, user"))
	})

	s.Start(":8081")
}
