package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id   int    `json "id"`
	Name string `json "name"`
	Age  int    `json "age"`
}

var users = map[int]*User{}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func createUser(c echo.Context) error {
	user := &User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	users[user.Id] = user

	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	user := new(User)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(user); err != nil {
		return err
	}
	users[id] = user
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.JSON(http.StatusOK, "Delete Users")
}

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	group := e.Group("/k")

	group.GET("/users", getAllUsers)
	group.GET("/users/:id", getUser)
	group.POST("/users", createUser)
	group.PUT("/users/:id", updateUser)
	group.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
