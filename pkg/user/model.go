package user

import (
	"graphql-assent/pkg/orm"
)

type Model struct {
	orm.PK
	Email       string
	Password    string
	NewPassword string `gorm:"-"`
	orm.Timestamps
}

func (m Model) TableName() string {
	return "users"
}

func (m *Model) BeforeSave() (err error) {
	if m.NewPassword != "" {
		m.Password, err = HashPassword(m.NewPassword)

		if err != nil {
			return err
		}
	}
	return
}
