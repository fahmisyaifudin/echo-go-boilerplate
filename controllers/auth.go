package controllers

import (
	"net/http"

	"github.com/fahmisyaifudin/echo-boilerplate/database"
	"github.com/fahmisyaifudin/echo-boilerplate/function"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type response struct {
	Message string        `json:"message"`
	Data    database.User `json:"data"`
}

func ActionLogin(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) (err error) {
		input := new(database.User)
		if err = c.Bind(input); err != nil {
			return
		}

		var users database.User

		db.Where("email = ? ", input.Email).First(&users)

		if function.CheckPasswordHash(input.Password, users.Password) {
			return c.JSON(http.StatusOK, &response{Message: "Success", Data: users})
		} else {
			return c.JSON(http.StatusUnauthorized, &response{Message: "Password and email didnt match"})
		}
	}
}
