package store

import (
	"github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Email    *string
	Password *string
	Bill     float64
}

func (st *Store) UserByEmail(email string) (*User, bool) {
	var user User
	err := st.gorm.First(&user, "email = ?", email).Error

	return &user, found(err)
}
