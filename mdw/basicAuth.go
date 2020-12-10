package mdw

import "github.com/labstack/echo/v4"

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "123" {
		c.Set("username", username)
		c.Set("admin", true)
		c.Set("password", password)
		return true, nil
	}

	if username == "hi" && password == "123" {
		c.Set("username", username)
		c.Set("admin", false)
		c.Set("password", password)
		return true, nil
	}

	return false, nil
}
