package gsg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// Context is the encapsulation of http.ResponseWriter and *http.Request
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandleFunc
	index    int
	// engine pointer
	engine *Engine
}

// newContext is the constructor of Context
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}

// PostForm is a wrapper of r.PostFormValue()
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query is a wrapper of r.URL.Query().Get()
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status sets the status code for response
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets the header for response
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String sets the string for response
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON sets the json for response
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data sets the data for response
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML sets the html for response
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

// Param returns the path parameter by key
func (c *Context) Param(key string) string {

	value, ok := c.Params[key]
	if !ok {
		return ""
	}
	return value
}

// Fail sets the status code and error message for response
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// Next is used to execute the next middleware
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}

}
