package store

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func found(err error) bool {
	switch err {
	case nil:
		return true
	case gorm.ErrRecordNotFound:
		return false
	}

	panic(err)
}

func notFound(err error) bool {
	switch err {
	case nil:
		return false
	case gorm.ErrRecordNotFound:
		return true
	}

	panic(err)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}
