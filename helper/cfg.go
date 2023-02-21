package helper

import (
	"github.com/joho/godotenv"
	"os"
)

func Cfg(key string) string {
	err := godotenv.Load(".env")
	Pie(err)
	return os.Getenv(key)
}
