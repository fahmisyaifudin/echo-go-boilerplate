package database

import (
	"github.com/fahmisyaifudin/echo-boilerplate/function"
	"github.com/thanhpk/randstr"
)

func RunSeeder() {
	db := Connect()
	users := []User{
		{Email: "fahmisyaifudin@gmail.com", Password: function.HashPassword("secret"), FirstName: "Fahmi", LastName: "Syaifudin", AuthKey: randstr.String(32)},
	}
	db.Create(&users)
}
