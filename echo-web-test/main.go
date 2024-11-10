package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dzjyyds666/echo-web-test/config"
	"github.com/dzjyyds666/echo-web-test/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/labstack/gommon/log"
)

func main() {

	//初始化数据库
	config.InitDB()

	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	// })

	// 日志中间件
	e.Use(middleware.Logger())

	// 隐藏横幅
	// e.HideBanner = true

	// 引入路由文件
	router.ApiRouter(e)

	// 自定义服务器
	s := &http.Server{
		Addr: ":9999",
		// 读写超时
		ReadTimeout:  time.Minute * 20,
		WriteTimeout: time.Minute * 20,
	}

	type Users struct {
		Name  string `json:"name" form:"name" query:"name"`
		Email string `json:"email" form:"email" query:"email"`
	}

	type UserDTO struct {
		Name    string
		Email   string
		IsAdmin bool
	}

	e.POST("/user", func(c echo.Context) (err error) {
		u := &Users{}

		// josn数据绑定
		// if err := c.Bind(u); err != nil {
		// 	return c.String(http.StatusBadRequest, "参数错误")
		// }

		// 表单数据绑定
		// err := echo.FormFieldBinder(c).
		// 	String("Name",&u.Name).
		// 	String("Email",&u.Email).
		// 	BindError()

		fmt.Println(*u)

		user := UserDTO{
			Name:    u.Name,
			Email:   u.Email,
			IsAdmin: true,
		}

		return c.JSON(http.StatusOK, user)
	})

	e.GET("/cookietest", func(c echo.Context) error {
		cookie := &http.Cookie{}
		cookie.Name = "test"
		cookie.Value = "hello,world"
		cookie.Expires = time.Now().Add(time.Second * 30)
		c.SetCookie(cookie)
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/getcookie", func(c echo.Context) error {
		cookie, err := c.Cookie("test")
		if err != nil {
			return c.String(http.StatusOK, "cookie不存在")
		}

		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)

		return c.String(http.StatusOK, "ok")
	})

	// 自定义监视器

	e.Logger.Fatal(e.StartServer(s))

}
