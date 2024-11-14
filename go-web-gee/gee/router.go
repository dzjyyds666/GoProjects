package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}


//only one * is allowed
func parsePattern(pattern string)[]string{
	vs:=strings.Split(pattern, "/")

	parts:=make([]string,0)

	for _,item :=range vs{
		if item !=""{
			parts=append(parts, item)
			if item[0]=='*'
		}
	}
}


func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handle, ok := r.handlers[key]; ok {
		handle(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
