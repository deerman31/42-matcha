package main

import (
	"database/sql"
	"fmt"
	"golang-echo/middle"
	"golang-echo/validations"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // PostgreSQLドライバー
)

func main() {
	// データベースに接続
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database")

	e := echo.New()

	// ミドルウェアの設定
	middle.SetupMiddleware(e)

	// カスタムバリデータの設定
	e.Validator = validations.NewValidator()

	routing(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}

func connectDB() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST") // Docker Composeのサービス名
	dbPort := "5432"
	dbName := os.Getenv("POSTGRES_DB")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		return nil, fmt.Errorf("There is a problem with the environment variables")
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

//func connectDB()
