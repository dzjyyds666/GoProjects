package handler

import (
	"net/http"

	"github.com/dzjyyds666/echo-web-test/dao"
	"github.com/labstack/echo"
)

// func GetUser(c echo.Context) error {

// 	name := "杜智军"
// 	age := 18
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"name": name,
// 		"age":  age,
// 	})
// }

func GetUserById(e echo.Context) error {
	id := e.QueryParam("id")
	// if err != nil {
	// 	return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	// }
	user, err := dao.GetUserById(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return e.JSON(http.StatusOK, user)
}
