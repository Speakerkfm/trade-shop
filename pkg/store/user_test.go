package store

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestStore_UserByEmail(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()
	user := User{ID: uID, Email: "asdf@mail.com"}

	Gorm.Table("users").Create(&user)

	res, found := s.UserByEmail("asdf@mail.com")
	assert.True(t, found)
	assert.True(t, res.ID == uID)

	Gorm.Table("users").Delete(&user)
}

func TestStore_UserByUserID(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()
	user := User{ID: uID, Email: "asdf@mail.com"}

	Gorm.Table("users").Create(&user)

	res, found := s.UserByUserID(uID)
	assert.True(t, found)
	assert.True(t, res.Email == "asdf@mail.com")

	Gorm.Table("users").Delete(&user)
}

func TestStore_GetUserBill(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()
	user := User{ID: uID, Email: "asdf@mail.com", Bill: 1234.5}

	Gorm.Table("users").Create(&user)

	bill := s.GetUserBill(uID)
	assert.True(t, bill == 1234.5)

	Gorm.Table("users").Delete(&user)
}

func TestStore_AddMoneyToUser(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()
	user := User{ID: uID, Email: "asdf@mail.com", Bill: 1234.5}

	Gorm.Table("users").Create(&user)

	err := s.AddMoneyToUser(Gorm, uID, 10.2)
	assert.Nil(t, err)

	Gorm.Table("users").First(&user)
	assert.True(t, user.Bill == 1244.7)

	Gorm.Table("users").Delete(&user)
}

func TestStore_RemoveMoneyFromUser(t *testing.T) {
	s := NewStore(Gorm, RedisClient)

	uID, _ := uuid.NewV4()
	user := User{ID: uID, Email: "asdf@mail.com", Bill: 1234.5}

	Gorm.Table("users").Create(&user)

	err := s.RemoveMoneyFromUser(Gorm, uID, 10.2)
	assert.Nil(t, err)

	Gorm.Table("users").First(&user)
	assert.True(t, user.Bill == 1224.3)

	err = s.RemoveMoneyFromUser(Gorm, uID, 5000.1)
	assert.NotNil(t, err)

	fmt.Println(user)
	Gorm.Table("users").First(&user)
	assert.True(t, user.Bill == 1224.3)

	Gorm.Table("users").Delete(&user)
}
