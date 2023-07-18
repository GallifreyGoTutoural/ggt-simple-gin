package gsg

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParseFullPath(t *testing.T) {
	ok := reflect.DeepEqual(parseFullPath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parseFullPath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parseFullPath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRoute("GET", "/hello/gallifrey")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.fullPath != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if params["name"] != "gallifrey" {
		t.Fatal("name should be equal to 'gallifrey'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.path, params["name"])

}
