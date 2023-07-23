package gsg

import (
	"net/http"
	"path"
	"strings"
)

// HandleFunc defines the request handler used by gsg
type HandleFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// RouterGroup is a group of router
type RouterGroup struct {
	prefix      string       // the prefix of the group
	middlewares []HandleFunc // support middleware
	parent      *RouterGroup // support nesting
	engine      *Engine      // all groups share an Engine instance
}

// New is the constructor of gsg.Engine
func New() *Engine {
	engin := &Engine{router: newRouter()}
	engin.RouterGroup = &RouterGroup{engine: engin}
	engin.groups = []*RouterGroup{engin.RouterGroup}
	return engin
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// group add router
func (group *RouterGroup) addRouter(method string, comp string, handle HandleFunc) {
	fullPath := group.prefix + comp
	group.engine.router.addRoute(method, fullPath, handle)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(relativePath string, handle HandleFunc) {
	group.addRouter("GET", relativePath, handle)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(relativePath string, handle HandleFunc) {
	group.addRouter("POST", relativePath, handle)

}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// ServeHTTP implement the interface of http.Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandleFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	engine.router.handle(c)
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	prefixPath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(prefixPath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)

	}
}

// Static serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}
