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
	err := st.gorm.Table("users").First(&user, "email = ?", email).Error

	return &user, found(err)
}

func (u *User) PasswordValid(password string) bool {
	//err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password))
	//return err == nil
	return *u.Password == password
}

func (st *Store) UpdateUserBill(userID uuid.UUID, newBill float64) bool {
	err := st.gorm.
		Table("users").
		Where("id = ?", userID.String()).
		Updates(map[string]interface{}{"bill": newBill}).Error

	return found(err)
}
