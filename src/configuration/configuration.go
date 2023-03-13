package configuration

import (
	"fmt"
)

var DBConfig *DatabaseConfig

type DatabaseConfig struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
}

func (db *DatabaseConfig) DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		db.PostgresUser,
		db.PostgresPassword,
		db.PostgresHost,
		db.PostgresPort,
		db.PostgresDB,
	)
}

func initDBConfig() {
	prefix := "postgres_"
	DBConfig = &DatabaseConfig{
		PostgresUser:     getEnvOrDefault(prefix, "user", "postgres"),
		PostgresPassword: getEnvOrDefault(prefix, "password", "postgres"),
		PostgresDB:       getEnvOrDefault(prefix, "database", "projects"),
		PostgresHost:     getEnvOrDefault(prefix, "host", "localhost"),
		PostgresPort:     getEnvOrDefault(prefix, "port", "5432"),
	}
}

func InitConfigurations() {
	initDBConfig()

}
