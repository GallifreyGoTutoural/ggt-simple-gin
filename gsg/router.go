package gsg

import "net/http"

// HandleFunc defines the request handler used by gsg
type router struct {
	handlers map[string]HandleFunc
}

// newRouter is the constructor of router
func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

// addRoute adds a route and its handler to the router
func (r *router) addRoute(method string, path string, handler HandleFunc) {
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
