package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Tokens struct {
	Token string `json:"token"`
}

type Logins struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "WellCome to me")
}

func login(c echo.Context) error {
	// user := new(Logins)
	// c.Bind(user)

	// if user.Username != "admin" || user.Password != "123" {
	// 	if user.Username != "admin" {
	// 		return c.JSON(http.StatusUnauthorized, "Wrong Username")
	// 	}
	// 	if user.Password != "123" {
	// 		return c.JSON(http.StatusUnauthorized, "Wrong Password")
	// 	}
	// }

	return c.JSON(http.StatusOK, &Tokens{
		Token: "asfj123jyh5u3423hj",
	})

}

func basicAuth(username string, password string, c echo.Context) (bool, error) {
	if username != "admin" || password != "123" {
		return false, nil
	}

	return true, nil
}

func main() {
	e := echo.New()

	group := e.Group("/thienhi")

	group.GET("/hello", hello)

	group.POST("/login", login, middleware.BasicAuth(basicAuth))

	e.Logger.Fatal(e.Start(":3000"))
}
