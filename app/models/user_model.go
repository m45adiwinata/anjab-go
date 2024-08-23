package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Email     string     `json:"email" db:"email"`
	Password  *string    `json:"password,omitempty" db:"password"`
	Fullname  string     `json:"fullname" db:"fullname"`
	Nickname  string     `json:"nickname" db:"nickname"`
	DOB       string     `json:"dob,omitempty" db:"dob"`
	SkpdId    int        `json:"skpd_id" db:"skpd_id"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (r *User) TableName() string {
	return "users"
}

func (r *User) Connection() string {
	return "mysql"
}
