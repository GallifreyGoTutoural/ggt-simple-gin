package gsg

import (
	"net/http"
)

// HandleFunc defines the request handler used by gsg
type HandleFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// ServeHTTP implement the interface of http.Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

// New is the constructor of gsg.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// add router
func (engine *Engine) addRouter(method string, relativePath string, handle HandleFunc) {
	engine.router.addRoute(method, relativePath, handle)
}

// GET defines the method to add GET request
func (engine *Engine) GET(relativePath string, handle HandleFunc) {
	engine.addRouter("GET", relativePath, handle)
}

// POST defines the method to add POST request
func (engine *Engine) POST(relativePath string, handle HandleFunc) {
	engine.addRouter("POST", relativePath, handle)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
