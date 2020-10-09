package route

import (
	"fmt"
	"os"

	"github.com/fahmisyaifudin/echo-boilerplate/controllers"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func HandleRequest(db *gorm.DB) {
	e := echo.New()

	e.POST("/login", controllers.ActionLogin(db))

	start := fmt.Sprintf(":%s", os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(start))
}
