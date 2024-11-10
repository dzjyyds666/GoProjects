package main

import (
	"image-video/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	router.ImageRouter(e)

	//日志中间件
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":9999"))
}
