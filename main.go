package main

import (
	"fmt"
	"github.com/GallifreyGoTutoural/ggt-simple-gin/gsg"
	"html/template"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gsg.HandleFunc {
	return func(c *gsg.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gsg.New()
	r.Use(gsg.Logger()) // global middleware

	v1 := r.Group("/v1")
	{

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
	v2.Use(onlyForV2()) // v2 group middleware
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

	r.Static("/assets", "./static") // localhost:8080/assets/xxx.png => ./static/xxx.png
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")

	v3 := r.Group("/v3")
	{
		v3.GET("/lovelyPng", func(c *gsg.Context) {
			c.HTML(http.StatusOK, "css.tmpl", nil)
		})
		v3.GET("/students", func(c *gsg.Context) {
			c.HTML(http.StatusOK, "arr.tmpl", gsg.H{
				"title":  "gsg",
				"stuArr": [2]*student{{Name: "gallifrey", Age: 20}, {Name: "Ace", Age: 12}},
			})
		})
		v3.GET("/date", func(c *gsg.Context) {
			c.HTML(http.StatusOK, "func.tmpl", gsg.H{
				"title": "gsg",
				"now":   time.Now(),
			})
		})

	}
	v4 := r.Group("/v4")
	v4.Use(gsg.Recovery(), gsg.Logger())

	{
		v4.GET("/panic", func(c *gsg.Context) {
			names := []string{"gallifrey"}
			c.String(http.StatusOK, names[100])
		})

	}

	r.Run(":8088")

}
