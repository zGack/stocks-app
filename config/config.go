package config

import (
	"time"

	"github.com/zgack/stocks/internal/env"
)

type Conf struct {
	Server ConfServer
	DB     ConfDB
}

type ConfServer struct {
	Port         int           `env:"SERVER_PORT,default=8080"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,default=30s"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,default=30s"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,default=60s"`
	Debug        bool          `env:"SERVER_DEBUG,default=true"`
	CorsOrigins  []string      `env:"SERVER_CORS_ORIGINS,default=*"`
}

type ConfDB struct {
	DBPath       string        `env:"DB_PATH,default=database.db"`
	MaxOpenConns int           `env:"DB_MAX_OPEN_CONNS,default=10"`
	MaxIdleConns int           `env:"DB_MAX_IDLE_CONNS,default=5"`
	MaxIdleTime  time.Duration `env:"DB_MAX_IDLE_TIME,default=5m"`
	Debug        bool          `env:"SERVER_DEBUG,default=true"`
}

func New() *Conf {
	config := &Conf{
		Server: ConfServer{
			Port:         env.GetInt("SERVER_PORT", 8080),
			TimeoutRead:  time.Second * 30,
			TimeoutWrite: time.Second * 30,
			TimeoutIdle:  time.Minute,
			Debug:        env.GetBool("SERVER_DEBUG", true),
			CorsOrigins:  []string{env.GetString("SERVER_CORS_ORIGINS", "*")},
		},
		DB: ConfDB{
			DBPath: env.GetString("DB_ADDR", "postgresql://root@localhost:26257/defaultdb?sslmode=disable"),
			Debug:  env.GetBool("SERVER_DEBUG", true),
		},
	}

	return config
}
