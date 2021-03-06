package store

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/vmihailenco/msgpack"
)

type Store struct {
	codec *cache.Codec
	gorm  *gorm.DB
}

func NewStore(db *gorm.DB, redisClient *redis.Client) *Store {
	codec := &cache.Codec{
		Redis: redisClient,

		Marshal: msgpack.Marshal,

		Unmarshal: msgpack.Unmarshal,
	}
	return &Store{gorm: db, codec: codec}
}

func (s *Store) CreateTransaction() *gorm.DB {
	return s.gorm.Begin()
}

func (s *Store) CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}

func (s *Store) RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
