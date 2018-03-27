package orm

import "time"

type PK struct {
	ID uint `gorm:"primary_key"`
}

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
