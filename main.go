package main

import (
	"github.com/GallifreyGoTutoural/ggt-simple-gin/gsg"
	"net/http"
)

func main() {
	r := gsg.New()

	r.GET("/", func(c *gsg.Context) {
		c.HTML(http.StatusOK, "<h1>Hello GSG</h1>")
	})

	r.GET("/hello", func(c *gsg.Context) {
		// expect /hello?name=gallifrey
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gsg.Context) {
		c.JSON(http.StatusOK, gsg.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8088")

}
