package store

import uuid "github.com/satori/go.uuid"

type Item struct {
	ID   uuid.UUID `gorm:"primary_key"`
	Name string
}
