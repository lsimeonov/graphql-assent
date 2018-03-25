package user

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"golang.org/x/crypto/argon2"
)

type Model struct {
	gorm.Model
	Email       string
	Password    string
	NewPassword []byte `gorm:"-"`
}

func (m *Model) BeforeSave() (err error) {
	if m.NewPassword != nil {
		p := argon2.Key(m.NewPassword, []byte("test"), 2, 1024, 2, 32)
		fmt.Println(string(p))
		os.Exit(1)
		m.NewPassword = nil
	}
	return
}
