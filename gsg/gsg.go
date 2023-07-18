package gsg

import (
	"fmt"
	"net/http"
)

// HandleFunc defines the request handler used by gsg
type HandleFunc func(w http.ResponseWriter, r *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandleFunc
}

// implement the interface of ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handle, ok := engine.router[key]; ok {
		handle(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// New is the constructor of gsg.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// add router
func (engine *Engine) addRouter(method string, relativePath string, handle HandleFunc) {
	key := method + "-" + relativePath
	engine.router[key] = handle
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
