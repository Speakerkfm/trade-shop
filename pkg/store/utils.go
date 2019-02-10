package store

import "github.com/jinzhu/gorm"

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

func check(err error) {
	if err != gorm.ErrRecordNotFound && err != nil {
		panic(err)
	}
}
