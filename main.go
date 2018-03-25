package main

import (
	"graphql-assent/pkg/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"graphql-assent/pkg/user"
)

func main() {
	c := conf.Init("configs", "config.dev")

	db, err := gorm.Open("mysql", c.DB.MarshallString())

	db.AutoMigrate(&user.Model{})

	u := user.Model{
		Email:       "test@example.com",
		NewPassword: []byte("test"),
	}

	db.NewRecord(u)
	db.Create(&u)

	if (err != nil) {
		panic(err)
	}

	defer db.Close()
}
