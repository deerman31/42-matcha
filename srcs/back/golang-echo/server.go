package main

import (
	"fmt"
	"golang-echo/app/database"
	"golang-echo/app/middle"
	"golang-echo/app/routes"
	"golang-echo/app/validations"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // PostgreSQLドライバー
)

func main() {
	// データベースに接続
	db, err := database.ConnectDB()
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

	routes.Routing(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
