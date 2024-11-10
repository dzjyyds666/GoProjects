package router

import (
	"image-video/handler"

	"github.com/labstack/echo/v4"
)

func ImageRouter(c *echo.Echo) {
	c.POST("/resize", handler.DealWhitImage)
}
