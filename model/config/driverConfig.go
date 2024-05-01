package dbconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

var (
	PostgresDriver = getEnv("POSTGRES_DRIVER", "postgres")
	User           = getEnv("POSTGRES_USER", "postgres")
	Host           = getEnv("POSTGRES_HOST", "localhost")
	Port           = getEnv("POSTGRES_PORT", "5432")
	Password       = getEnv("POSTGRES_PASSWORD", "postgres")
	DbName         = getEnv("POSTGRES_DB", "postgres")
	SSLMode        = getEnv("POSTGRES_SSLMODE", "disable")
)

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	Host, Port, User, Password, DbName, SSLMode)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
