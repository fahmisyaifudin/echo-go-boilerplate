package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fahmisyaifudin/echo-boilerplate/database"
	"github.com/fahmisyaifudin/echo-boilerplate/function"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ActionLogin(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		input := new(database.User)
		if err = c.Bind(input); err != nil {
			return
		}

		var users database.User

		db.Where("email = ? ", input.Email).First(&users)

		if function.CheckPasswordHash(input.Password, users.Password) {
			token, _ := JwtEncoded(users)
			return c.JSON(http.StatusOK, map[string]string{
				"message": "success",
				"token":   token,
			})
		} else {
			return echo.ErrUnauthorized
		}
	}
}

func ActionProfile(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		var users database.User
		db.Where("auth_key = ? ", JwtDecoded(c)).First(&users)

		var response struct {
			Message string        `json:"message"`
			Data    database.User `json:"data"`
		}
		response.Message = "success"
		response.Data = users

		return c.JSON(http.StatusOK, response)
	}
}

func JwtDecoded(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authKey := claims["key"].(string)
	return authKey
}

func JwtEncoded(users database.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["key"] = users.AuthKey
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, err
}
