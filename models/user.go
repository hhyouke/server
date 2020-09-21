package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// User model
type User struct {
	ID         int64
	Nickname   string
	RealName   string
	Mail       string
	Phone      string
	Gender     bool
	Dob        *time.Time
	Avatar     string
	CreateUser int64
	UpdateUser int64
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
}

// TableName db table name
func (u *User) TableName() string {
	return "t_user"
}

// Create create new user
func (u *User) Create(tx *sqlx.Tx) error {
	iq := "INSERT into t_user(id, nickname, create_user, update_user, create_time, update_time) VALUES (:id, :nickname, :create_user, :update_user, :create_time, :update_time)"
	if _, err := tx.NamedExec(iq, u); err != nil {
		return err
	}
	return nil
}
