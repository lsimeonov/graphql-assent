package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"graphql-assent/pkg/conf"
	"graphql-assent/pkg/user"
)

func main() {
	c := conf.Init("configs", "config.dev")

	pa, err := user.NewPwdAlgo(c.General.PasswordAlgo)
	if err != nil {
		panic(err)
	}
	user.Services.Pwd = pa

	db, err := gorm.Open("mysql", c.DB.MarshallString())

	if err != nil {
		panic(err)
	}
	init1(db)
	//
	//var u user.Model
	//
	//db.Where("email = ?", "test@example.com").First(&u)
	//
	//match, _ := user.Compare(u.Password, "sd")
	//
	//fmt.Println(match)

	defer db.Close()
}

func init1(db *gorm.DB) {
	db.AutoMigrate(&user.Model{})
	u := user.Model{
		Email:       "test@example.com",
		NewPassword: "test",
	}
	db.NewRecord(u)
	db.Create(&u)
}
