package main

import (
	"graphql-assent/pkg/conf"

	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"graphql-assent/pkg/user"
)

func main() {
	c := conf.Init("configs", "config.dev")

	db, err := gorm.Open("mysql", c.DB.MarshallString())

	//init1(db)
	if err != nil {
		panic(err)
	}

	var u user.Model

	db.Where("email = ?", "test@example.com").First(&u)

	match, _ := user.CompareHashWithPassword(u.Password, "sd")

	fmt.Println(match)

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
