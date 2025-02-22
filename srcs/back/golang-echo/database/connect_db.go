package database

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST") // Docker Composeのサービス名
	dbPort := "5432"
	dbName := os.Getenv("POSTGRES_DB")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		return nil, fmt.Errorf("there is a problem with the environment variables")
	}

	// PostgreSQL用の接続文字列を構築
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	// データベースに接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	// 接続テスト
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	return db, nil
}
