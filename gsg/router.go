package gsg

import (
	"strings"
)

// HandleFunc defines the request handler used by gsg
type router struct {
	root     map[string]*node
	handlers map[string]HandleFunc
}

/*
roots key eg, roots['GET'] roots['POST']
handlers key eg, handlers['GET-/p/:lang/doc'] handlers['POST-/p/book']
newRouter is the constructor of router
*/
func newRouter() *router {
	return &router{
		root:     map[string]*node{},
		handlers: make(map[string]HandleFunc),
	}
}

// getRoute returns the matching node and its parameters
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.fullPath
		// execute the handler
		r.handlers[key](c)
	} else {
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}

// parseFullPath parses a full path to a string slice -- paths
func parseFullPath(fullPath string) []string {
	vs := strings.Split(fullPath, "/")
	paths := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			paths = append(paths, item)
			// Only one wildcard(*) is allowed in a path
			if item[0] == '*' {
				break
			}
		}
	}
	return paths
}

// addRoute adds a route and its handler to the router
func (r *router) addRoute(method string, fullPath string, handler HandleFunc) {
	paths := parseFullPath(fullPath)
	key := method + "-" + fullPath
	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}
	r.root[method].insert(fullPath, paths, 0)
	r.handlers[key] = handler
}

/*
getRoute eg:
1 normal: fullPath=[/p/go/doc] math [/p/:lang/doc]  => {"lang": "go"}
2 [/]: fullPath=[/static/css/gallifrey.css] math [/static/*filepath]  => {"filepath": "css/gallifrey.css"}
3 [*]: fullPath=[/p/go] math [/p/:lang] => {"lang": "go"}
*/
// getRoute finds the matching route according to the method and path and returns a map of parameters
func (r *router) getRoute(method string, fullPath string) (*node, map[string]string) {
	searchPaths := parseFullPath(fullPath)
	params := make(map[string]string)
	root, ok := r.root[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchPaths, 0)
	if n != nil {
		paths := parseFullPath(n.fullPath)
		for index, path := range paths {
			if path[0] == ':' {
				params[path[1:]] = searchPaths[index]
			}
			if path[0] == '*' && len(path) > 1 {
				params[path[1:]] = strings.Join(searchPaths[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
