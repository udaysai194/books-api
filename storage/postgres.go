package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func Configure(envFile string) *Config {
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("error loading in env file")
		log.Fatal(err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return config
}

func Connect(config *Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := pgxpool.Connect(c, dsn)
	if err != nil {
		return db, err
	}
	return db, nil
}
