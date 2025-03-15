package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	if os.Getenv("RENDER") == "" { 
		if err := loadEnvFile(); err != nil {
			log.Printf("Warning: %v\n", err)
		}
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing required database connection information in environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	log.Println("Database connection successfully established.")
	return db, nil
}

func loadEnvFile() error {
	rootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	for {
		envPath := fmt.Sprintf("%s/.env", rootDir)
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}
			return nil
		}

		parentDir := filepath.Dir(rootDir)
		if parentDir == rootDir {
			break
		}
		rootDir = parentDir
	}

	return fmt.Errorf(".env file not found")
}
