package main

import (
	"fmt"
	"github.com/GallifreyGoTutoural/ggt-simple-gin/gsg"
	"net/http"
)

func main() {
	route := gsg.New()
	route.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	route.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	route.Run(":8088")

}
