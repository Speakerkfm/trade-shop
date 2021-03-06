package flags

type Config struct {
	DatabaseHost     string `long:"database-host" env:"DATABASE_HOST" required:"true"`
	DatabasePort     string `long:"database-port" env:"DATABASE_PORT" default:"3306"`
	DatabaseName     string `long:"database-name" env:"DATABASE_NAME" required:"true"`
	DatabaseUser     string `long:"database-user" env:"DATABASE_USER" required:"true"`
	DatabasePassword string `long:"database-password" env:"DATABASE_PASSWORD" required:"true"`
	RedisHost        string `long:"redis-host" env:"REDIS_HOST" required:"true"`

	RedisStoreSize     string `long:"redis-store-size" env:"REDIS_STORE_SIZE" default:"10"`
	RedisStoreHost     string `long:"redis-store-host" env:"REDIS_STORE_HOST" required:"true"`
	RedisStorePort     string `long:"redis-store-port" env:"REDIS_STORE_PORT" default:"6379"`
	RedisStorePassword string `long:"redis-store-password" env:"REDIS_STORE_PASSWORD" required:"true"`

	AmqpHost     string `long:"amqp-host" env:"AMQP_HOST" required:"true"`
	AmqpPort     int    `long:"amqp-port" env:"AMQP_PORT" required:"true"`
	AmqpUser     string `long:"amqp-user" env:"AMQP_USER" required:"true"`
	AmqpPassword string `long:"amqp-password" env:"AMQP_PASSWORD" required:"true"`
}
