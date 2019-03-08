package store

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/vmihailenco/msgpack"
	"os"
	"time"
)

var (
	Gorm        *gorm.DB
	Codec       *cache.Codec
	RedisClient *redis.Client
)

func init() {
	mysqlConf := mysql.NewConfig()
	mysqlConf.Net = "tcp"
	mysqlConf.Addr = os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT")
	mysqlConf.User = os.Getenv("DATABASE_USER")
	mysqlConf.Passwd = os.Getenv("DATABASE_PASSWORD")
	mysqlConf.DBName = os.Getenv("DATABASE_NAME")
	mysqlConf.MultiStatements = true
	mysqlConf.ParseTime = true
	mysqlConf.Loc = time.Local

	db, err := gorm.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	Gorm = db

	redisOpt := redis.Options{Addr: os.Getenv("REDIS_HOST")}
	redisClient := redis.NewClient(&redisOpt)
	if _, err := redisClient.Ping().Result(); err != nil {
		panic(err)
	}

	RedisClient = redisClient

	codec := &cache.Codec{
		Redis: redisClient,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	Codec = codec
}
