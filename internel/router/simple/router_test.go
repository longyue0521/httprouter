package simple_test

import (
	"net/http"
	"testing"

	"github.com/longyue0521/httprouter/internel/router/simple"
	"github.com/longyue0521/httprouter/web"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {

	t.Run("静态路由", func(t *testing.T) {

		handler := web.HandleFunc(func(ctx *web.Context) {

		})
		expectedHandlerPtr := &handler

		assert.Equal(t, &handler, expectedHandlerPtr)

		// case /
		t.Run("路径", func(t *testing.T) {

			t.Run("必须以'/'开头", func(t *testing.T) {
				r := NewRouter()

				assert.PanicsWithValue(t, "web: 路由必须以'/'开头", func() {
					r.Add(http.MethodGet, "a", handler)
				})

				assert.PanicsWithValue(t, "web: 路由必须以'/'开头", func() {
					r.Get(http.MethodGet, "/")
				})

			})

			t.Run("不能以'/'结尾", func(t *testing.T) {

			})

			t.Run("中间无重复'/'", func(t *testing.T) {

			})
		})

		// /user/add/1
		// /user/fnd/1
	})

	t.Run("通配路由", func(i *testing.T) {

	})

	t.Run("参数路由", func(t *testing.T) {

	})

	t.Run("正则路由", func(t *testing.T) {

	})
}

func NewRouter() *simple.Router {
	return simple.New()
}