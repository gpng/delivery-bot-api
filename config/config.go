package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config for app
type Config struct {
	Docs   bool
	Server serverConf
	DB     dbConf
	Bot    botConf
}

type serverConf struct {
	Port string
}

type dbConf struct {
	Name     string
	Password string
	User     string
	Host     string
}

type botConf struct {
	Token string
}

// New app config
func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load godotenv")
	}

	docs := getEnvAsBool("DOCS", false)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")

	botToken := os.Getenv("BOT_TOKEN")

	return &Config{
		Docs: docs,
		Server: serverConf{
			Port: port,
		},
		DB: dbConf{
			Name:     dbName,
			Password: dbPassword,
			User:     dbUser,
			Host:     dbHost,
		},
		Bot: botConf{botToken},
	}
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := os.Getenv(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}
