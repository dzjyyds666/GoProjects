package router

import (
	"github.com/dzjyyds666/echo-web-test/handler"
	"github.com/labstack/echo"
)

func ApiRouter(e *echo.Echo) {
	userRouter(e)
}

func userRouter(e *echo.Echo) {
	e.GET("/getuser", handler.GetUserById)
}
