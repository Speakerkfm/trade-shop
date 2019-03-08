package store

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Email    string
	Password string
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

func (st *Store) CreateNewUser(email string, password string) (*User, error) {
	id, _ := uuid.NewV4()
	passwordHash := hashAndSalt([]byte(password))
	user := User{ID: id, Email: email, Password: passwordHash}

	err := st.gorm.Create(&user).Error

	return &user, err
}

func (st *Store) IsEmailTaken(email string) bool {
	var user User
	err := st.gorm.First(&user, "email = ?", email).Error

	return found(err)
}

func (st *Store) UserByUserID(userID uuid.UUID) (*User, bool) {
	user := User{
		ID: userID,
	}
	err := st.gorm.First(&user).Error

	return &user, found(err)
}

func (u *User) PasswordValid(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (st *Store) GetUserBill(userID uuid.UUID) float64 {
	user := User{ID: userID}

	st.gorm.First(&user)

	return user.Bill
}

func (st *Store) AddMoneyToUser(db *gorm.DB, userID uuid.UUID, money float64) error {
	user := User{ID: userID}

	if err := db.First(&user).Error; err != nil {
		return nil
	}

	user.Bill += money

	return db.Save(&user).Error
}

func (st *Store) RemoveMoneyFromUser(db *gorm.DB, userID uuid.UUID, money float64) error {
	user := User{ID: userID}

	if err := db.First(&user).Error; err != nil {
		return nil
	}

	if money > user.Bill {
		return fmt.Errorf("not enough money")
	}

	user.Bill -= money

	return db.Save(&user).Error
}
