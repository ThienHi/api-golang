package handler

import (
	"log"
	"net/http"
	"thienhi/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)
	admin := c.Get("admin").(bool)
	password := c.Get("password").(string)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	t, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		log.Printf("Error parsing token %v", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.Tokens{
		Token: t,
	})
}
