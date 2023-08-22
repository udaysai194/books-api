package storage

import (
	"books-api/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func ConfigPostgres(envFile string) (*models.Config, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("error loading in env file")
		log.Fatal(err)
	}

	config := &models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return config, err
}

func InitPostgres(config *models.Config) (Postgres, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return db, err
	}

	return db, nil
}
