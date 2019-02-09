package flags

type Config struct {
	DatabaseHost     string `long:"database-host" env:"DATABASE_HOST" required:"true"`
	DatabasePort     string `long:"database-port" env:"DATABASE_PORT" default:"3306"`
	DatabaseName     string `long:"database-name" env:"DATABASE_NAME" required:"true"`
	DatabaseUser     string `long:"database-user" env:"DATABASE_USER" required:"true"`
	DatabasePassword string `long:"database-password" env:"DATABASE_PASSWORD" required:"true"`
}
