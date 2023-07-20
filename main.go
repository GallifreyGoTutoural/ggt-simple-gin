package main

import (
	"github.com/GallifreyGoTutoural/ggt-simple-gin/gsg"
	"net/http"
)

func main() {
	r := gsg.New()

	r.GET("/index", func(c *gsg.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gsg.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		hello := v1.Group("/hello")
		{
			hello.GET("/:name", func(c *gsg.Context) {
				// expect /hello/gallifrey
				c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
			})
			hello.GET("/", func(c *gsg.Context) {
				// expect /hello?name=gallifrey
				c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
			})
		}
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gsg.Context) {
			// expect /hello/gallifrey
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gsg.Context) {
			c.JSON(http.StatusOK, gsg.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":8088")

}
