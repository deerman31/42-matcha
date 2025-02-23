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

	_ "golang-echo/docs" // swag initで生成されるdocsパッケージ

	echoSwagger "github.com/swaggo/echo-swagger" // Swaggerパッケージの追加
)

// @title あなたのAPIタイトル
// @version 1.0
// @description APIの説明をここに書く
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// データベースに接続
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database")

	e := echo.New()

	// Swaggerのルート設定
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// ミドルウェアの設定
	middle.SetupMiddleware(e)

	// カスタムバリデータの設定
	e.Validator = validations.NewValidator()

	routes.Routing(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
