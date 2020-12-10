package main

import (
	"thienhi/handler"
	"thienhi/mdw"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	isLogin := middleware.JWT([]byte("secretKey"))
	isAdmin := mdw.AdminMdw

	group := e.Group("/thien", isLogin)

	group.GET("/hello", handler.Hello, isAdmin)

	e.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	e.Logger.Fatal(e.Start(":3000"))
}
