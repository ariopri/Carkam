package config

import (
	"github.com/joho/godotenv"
	"os"
)

import (
	"fmt"
	"strconv"
	"time"
)

type Config struct {
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	ServerPort     string
	TokenSecret    string
	TokenExpiresIn time.Duration
	TokenMaxAge    int
}

func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return 0
	}
	return duration
}

func parseInt(intStr string) int {
	intValue, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Println("Error parsing integer:", err)
		return 0
	}
	return intValue
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	config = Config{
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBName:         os.Getenv("DB_NAME"),
		ServerPort:     os.Getenv("PORT"),
		TokenSecret:    os.Getenv("TOKEN_SECRET"),
		TokenExpiresIn: parseDuration(os.Getenv("TOKEN_EXPIRED_IN")),
		TokenMaxAge:    parseInt(os.Getenv("TOKEN_MAXAGE")),
	}
	return config, nil
}
