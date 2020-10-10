package route

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/fahmisyaifudin/echo-boilerplate/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	// fmt.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	authKey := claims["key"].(string)
	return c.String(http.StatusOK, authKey)
}

func HandleRequest(db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", controllers.ActionLogin(db))

	r := e.Group("/api")

	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.GET("/profile", controllers.ActionProfile(db))

	start := fmt.Sprintf(":%s", os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(start))
}
