package store

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Email    *string
	Password *string
	Bill     float64
}

func (User) TableName() string {
	return "users"
}

func (st *Store) UserByEmail(email string) (*User, bool) {
	var user User
	err := st.gorm.First(&user, "email = ?", email).Error

	return &user, found(err)
}

func (u *User) PasswordValid(password string) bool {
	//err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password))
	//return err == nil
	return *u.Password == password
}

func (st *Store) AddMoneyToUser(db *gorm.DB, userID uuid.UUID, money float64) error {
	user := User{ID: userID}

	if err := db.First(&user).Error; err != nil {
		return nil
	}

	user.Bill = user.Bill + money

	return db.Save(&user).Error
}

func (st *Store) RemoveMoneyFromUser(db *gorm.DB, userID uuid.UUID, money float64) error {
	user := User{ID: userID}

	if err := db.First(&user).Error; err != nil {
		return nil
	}

	if money > user.Bill {
		return errors.New("NotEnoughMoney")
	}

	user.Bill = user.Bill - money

	return db.Save(&user).Error
}
