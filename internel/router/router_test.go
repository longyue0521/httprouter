package router_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/longyue0521/httprouter/internel/router"
	"github.com/longyue0521/httprouter/web"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {

	handler := web.HandleFunc(func(ctx *web.Context) {})
	handler2 := web.HandleFunc(func(ctx *web.Context) {})
	assert.False(t, Equals(handler, handler2))

	t.Run("路径非法", func(t *testing.T) {
		t.Parallel()

		testcases := map[string]struct {
			router   *router.Router
			expected string
			paths    []string
		}{
			"空路由": {
				router:   NewRouter(),
				expected: "web: 路由不能为空",
				paths:    []string{""},
			},
			"不以'/'开头": {
				router:   NewRouter(),
				expected: "web: 路由必须以'/'开头",
				paths:    []string{" ", "a/", "a/b", "a/b/c"},
			},
			"以'/'结尾": {
				router:   NewRouter(),
				expected: "web: 路由不能以'/'结尾",
				paths:    []string{"//", "/a/", "/a/b/", "/a/b/c/"},
			},
			"有'/'重复": {
				router:   NewRouter(),
				expected: "web: 路由不能有重复'/'",
				paths:    []string{"///a//b///c//////d", "//a"},
			},
			"有空白字符": {
				router:   NewRouter(),
				expected: "web: 路由不能包含空白字符",
				paths:    []string{"/\t", "/ ", "/\n", "/\v", "/\f", "/\r", "/\ta /\t\tb\n\f\v"},
			},
		}

		for name, tc := range testcases {
			tc := tc
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				for _, path := range tc.paths {
					assert.PanicsWithValue(t, tc.expected, func() {
						tc.router.Add(http.MethodGet, path, handler)
					})
					_, _, err := tc.router.Get(http.MethodGet, path)
					assert.EqualError(t, err, tc.expected)
				}
			})
		}
	})

	t.Run("静态路由", func(t *testing.T) {
		t.Parallel()

		t.Run("路由重复", func(t *testing.T) {
			t.Parallel()

			paths := []string{"/", "/order", "/order/edit", "/user/add/1", "/user/add/2"}

			for _, path := range paths {

				r := NewRouter()

				r.Add(http.MethodGet, path, handler)

				assert.PanicsWithValue(t, "web: 路由重复", func() {
					r.Add(http.MethodGet, path, handler2)
				})
			}

			for i := len(paths) - 1; i >= 0; i-- {
				r := NewRouter()

				r.Add(http.MethodGet, paths[i], handler)

				assert.PanicsWithValue(t, "web: 路由重复", func() {
					r.Add(http.MethodGet, paths[i], handler2)
				})
			}

		})

		t.Run("正常情况", func(t *testing.T) {
			t.Parallel()

			paths := []string{"/", "/order/create", "/order/edit/1", "/order/edit/2", "/user/add/id/1", "/user/add/id/2"}

			for _, path := range paths {
				r := NewRouter()

				r.Add(http.MethodGet, path, handler)

				_, expected, err := r.Get(http.MethodGet, path)

				assert.NoError(t, err)
				assert.True(t, Equals(expected, handler))
				assert.False(t, Equals(expected, handler2))
			}

			for i := len(paths) - 1; i >= 0; i-- {
				r := NewRouter()

				r.Add(http.MethodGet, paths[i], handler)

				_, expected, err := r.Get(http.MethodGet, paths[i])

				assert.NoError(t, err)
				assert.True(t, Equals(expected, handler))
				assert.False(t, Equals(expected, handler2))
			}

		})

		t.Run("非法路由", func(t *testing.T) {
			t.Parallel()

			r := NewRouter()

			paths := []string{"/", "/order/create", "/order/edit/1", "/order/edit/2", "/user/add/id/1", "/user/add/id/2"}

			for _, path := range paths {
				_, _, err := r.Get(http.MethodGet, path)
				assert.EqualError(t, err, "web: 非法路由")
			}

			for i := len(paths) - 1; i >= 0; i-- {
				_, _, err := r.Get(http.MethodGet, paths[i])
				assert.EqualError(t, err, "web: 非法路由")
			}

		})
	})

	t.Run("通配路由", func(i *testing.T) {

	})

	t.Run("参数路由", func(t *testing.T) {

	})

	t.Run("正则路由", func(t *testing.T) {

	})
}

func NewRouter() *router.Router {
	return router.New()
}

func Equals(handler1 web.HandleFunc, handler2 web.HandleFunc) bool {
	v1 := reflect.ValueOf(handler1)
	v2 := reflect.ValueOf(handler2)
	return v1.UnsafePointer() == v2.UnsafePointer()
}

func TestAdd(t *testing.T) {
	paths := []string{"/"}
	for _, path := range paths[1:] {
		assert.True(t, false, path)
	}
	assert.Equal(t, []string{}, paths[1:])
}
