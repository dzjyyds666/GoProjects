package main

import (
	"gee"
	"net/http"
)

func main() {

	engine := gee.New()

	engine.GET("/", indexPage)
	engine.POST("/", POSTtest)

	engine.Run(":9999")
}

func indexPage(c *gee.Context) {
	c.String(http.StatusOK, "Hello Gee\n")
}

func POSTtest(c *gee.Context) {
	c.JSON(http.StatusOK, "POST")
}
