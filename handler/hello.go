package handler

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	Username := claims["username"].(string)
	admin := claims["admin"].(bool)
	password := claims["password"].(string)
	// role := claims["role"].(string)

	mes := fmt.Sprintf("WellCome %s Password %s , %v ", Username, password, admin)

	return c.JSON(http.StatusOK, mes)
}
